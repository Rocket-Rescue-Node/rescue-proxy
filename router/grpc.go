package router

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/Rocket-Pool-Rescue-Node/credentials"
	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/consensuslayer"
	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/executionlayer"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mwitkow/grpc-proxy/proxy"
	prysmpb "github.com/prysmaticlabs/prysm/v3/proto/prysm/v1alpha1"
	rptypes "github.com/rocket-pool/rocketpool-go/types"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func getMessages() []string {
	return []string{
		"GetAttesterDuties",
		"GetProposerDuties",
		"GetSyncCommitteeDuties",
		"GetSyncStatus",
		"ProduceBlockV2",
		"ProduceBlockV2SSZ",
		"ProduceBlindedBlock",
		"ProduceBlindedBlockSSZ",
		"PrepareBeaconProposer",
		"SubmitValidatorRegistration",
		"ProduceAttestationData",
		"GetAggregateAttestation",
		"SubmitAggregateAndProofs",
		"SubmitBeaconCommitteeSubscription",
		"SubmitSyncCommitteeSubscription",
		"ProduceSyncCommitteeContribution",
		"SubmitContributionAndProofs",
		"GetLiveness",
		"WaitForActivation",
	}
}

type GRPCRouter struct {
	proxy              *grpc.Server
	upstream           *grpc.ClientConn
	listener           net.Listener
	Logger             *zap.Logger
	EL                 *executionlayer.ExecutionLayer
	CL                 *consensuslayer.ConsensusLayer
	CM                 *credentials.CredentialManager
	AuthValidityWindow time.Duration
}

type validationCb func(proto.Message, common.Address) error

type guardedServerStream struct {
	grpc.ServerStream
	router     *GRPCRouter
	svcName    string
	credential credentials.AuthenticatedCredential
	cb         validationCb
	nodeAddr   common.Address
}

func (g *GRPCRouter) validatePrepareBeaconProposer(m proto.Message, nodeAddr common.Address) error {

	pbp := &prysmpb.PrepareBeaconProposerRequest{}

	unknown := []byte(m.ProtoReflect().GetUnknown())
	proto.Unmarshal(unknown, pbp)

	// Create a slice of the indices
	indices := make([]string, 0, len(pbp.Recipients))

	for _, proposer := range pbp.Recipients {
		indices = append(indices, strconv.FormatUint(uint64(proposer.ValidatorIndex), 10))
	}

	// Get the index->pubkey map
	pubkeyMap, err := g.CL.GetValidatorPubkey(indices)
	if err != nil {
		g.Logger.Error("Error while querying CL for validator pubkeys", zap.Error(err))
		return status.Errorf(codes.Internal, "internal error")
	}

	// Iterate the results and check the fee recipients against our expected values
	// Note: we iterate the map from the HTTP request to ensure every key is present in the
	// response from the consensuslayer abstraction
	for _, proposer := range pbp.Recipients {
		index := strconv.FormatUint(uint64(proposer.ValidatorIndex), 10)
		pubkey, found := pubkeyMap[index]
		if !found {
			g.Logger.Warn("Pubkey for index not found in response from cl.",
				zap.String("requested index", index))
			return status.Error(codes.PermissionDenied, "pubkey isn't owned by node")
		}

		// Next we need to get the expected fee recipient for the pubkey
		expectedFeeRecipient, unowned := g.EL.ValidatorFeeRecipient(pubkey, &nodeAddr)
		if expectedFeeRecipient == nil {
			g.Logger.Warn("Pubkey not found in EL cache, or wasn't owned by the user",
				zap.String("key", pubkey.String()),
				zap.Bool("someone else's validator", unowned))
			return status.Error(codes.PermissionDenied, "pubkey belongs to someone else or isn't owned by a rp node")
		}

		if !bytes.Equal(expectedFeeRecipient.Bytes(), proposer.FeeRecipient) {
			// Looks like a cheater- fee recipient doesn't match expectations
			g.Logger.Warn("prepare_beacon_proposer called with unexpected fee recipient",
				zap.String("expected", expectedFeeRecipient.String()), zap.String("got", hex.EncodeToString(proposer.FeeRecipient)))
			return status.Error(codes.PermissionDenied, "incorrect fee recipient")
		}
	}

	return nil
}

func (g *GRPCRouter) validateRegisterValidators(m proto.Message, nodeAddr common.Address) error {

	rv := &prysmpb.SignedValidatorRegistrationsV1{}

	unknown := []byte(m.ProtoReflect().GetUnknown())
	proto.Unmarshal(unknown, rv)

	for _, registration := range rv.Messages {
		pubkey := (*rptypes.ValidatorPubkey)(registration.Message.Pubkey)

		// Grab the expected fee recipient for the pubkey
		expectedFeeRecipient, unowned := g.EL.ValidatorFeeRecipient(*pubkey, &nodeAddr)
		if expectedFeeRecipient == nil {
			// When unowned is true for register_validators, it means the pubkey was someone else's minipool
			// we still want that to get rejected... however, if unowned is false and expectedFeeRecipient is nil,
			// it means we're seeing a solo validator using mev-boost. Since register_validator requires a signature,
			// we can allow this fee recipient.
			if !unowned {
				// Move on to the next pubkey
				continue
			}
			g.Logger.Warn("Pubkey not found in EL cache. Not an RP validator?", zap.String("key", pubkey.String()))
			return status.Error(codes.PermissionDenied, "pubkey belongs to someone else")
		}

		if !bytes.Equal(expectedFeeRecipient.Bytes(), registration.Message.FeeRecipient) {
			g.Logger.Warn("register_validator called with unexpected fee recipient",
				zap.String("expected", expectedFeeRecipient.String()),
				zap.String("got", hex.EncodeToString(registration.Message.FeeRecipient)))
			return status.Error(codes.PermissionDenied, "incorrect fee recipient")
		}

	}

	return nil
}

func (g *guardedServerStream) SendMsg(m interface{}) error {
	return g.ServerStream.SendMsg(m)
}

func (g *guardedServerStream) RecvMsg(m interface{}) error {
	pbMsg, ok := m.(proto.Message)
	if !ok {
		g.router.Logger.Warn("Unable to capture proto message from gRPC request")
		return status.Errorf(codes.Internal, "invalid request")
	}

	g.router.Logger.Debug("intercepted proto request", zap.String("svc", g.svcName))
	if err := g.cb(pbMsg, g.nodeAddr); err != nil {
		return err
	}
	return g.ServerStream.RecvMsg(m)
}

func (g *GRPCRouter) payloadInterceptor() grpc.StreamServerInterceptor {
	services := map[string]any{
		"ethereum.eth.v1alpha1.BeaconChain":         struct{}{},
		"ethereum.eth.v1alpha1.BeaconNodeValidator": struct{}{},
		"ethereum.eth.v1alpha1.Node":                struct{}{},
	}

	msgCbs := map[string]validationCb{
		"PrepareBeaconProposer":        g.validatePrepareBeaconProposer,
		"SubmitValidatorRegistrations": g.validateRegisterValidators,
	}

	return func(srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {

		method := strings.Split(info.FullMethod, "/")
		_, matched := services[method[1]]
		if !matched {
			g.Logger.Warn("unknown service", zap.String("service", method[1]))
			return status.Errorf(codes.Unimplemented, "unknown service %s", method[1])
		}

		// Get the authentication header
		ctx := stream.Context()
		md, exists := metadata.FromIncomingContext(ctx)
		if !exists {
			g.Logger.Warn("no metadata on inbound request", zap.String("service", method[1]))
			return status.Errorf(codes.Unauthenticated, "no metadata on inbound request")
		}

		// See https://github.com/prysmaticlabs/prysm/issues/11765
		var nodeAddr common.Address
		if method[2] != "DomainData" && method[2] != "SubscribeCommitteeSubnets" {
			val, exists := md["rprnauth"]
			if !exists || len(val) < 1 {
				g.Logger.Debug("grpc access without auth header", zap.String("service", method[1]), zap.String("method", method[2]), zap.Bool("exists", exists))
				return status.Errorf(codes.Unauthenticated, "headers missing")
			}

			auth := strings.Split(val[0], ":")
			if len(auth) != 2 {
				g.Logger.Debug("grpc access with invalid auth header")
				return status.Errorf(codes.Unauthenticated, "headers invalid")
			}

			ac := credentials.AuthenticatedCredential{}
			ac.Base64URLDecode(auth[0], auth[1])

			err := g.CM.Verify(&ac)
			if err != nil {
				g.Logger.Debug("Unable to verify hmac on guarded endpoint", zap.Error(err))
				return status.Errorf(codes.Unauthenticated, "auth header could not be verified")
			}

			// Grab the timestamp and make sure the credential is recent enough
			ts := time.Unix(ac.Credential.Timestamp, 0)
			now := time.Now()

			if ts.Before(now) && now.Sub(ts) > g.AuthValidityWindow {
				g.Logger.Debug("Stale credential seen on guarded endpoint")
				return status.Errorf(codes.PermissionDenied, "credential expired")
			}

			nodeAddr = common.BytesToAddress(ac.Credential.NodeId)
		}

		if cb, matched := msgCbs[method[2]]; matched {
			wrapper := &guardedServerStream{
				ServerStream: stream,
				router:       g,
				svcName:      method[2],
				cb:           cb,
				nodeAddr:     nodeAddr}

			return handler(srv, wrapper)
		}
		return handler(srv, stream)
	}
}

func (g *GRPCRouter) director() proxy.StreamDirector {
	return func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {

		g.Logger.Warn("director invoked", zap.String("service", fullMethodName))
		return ctx, nil, fmt.Errorf("director should not be invoked")
	}
}

func (g *GRPCRouter) Init(listenAddr string, beaconAddr string) error {
	var err error
	g.listener, err = net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	g.Logger.Info("Starting grpc server", zap.String("url", listenAddr), zap.String("upstream", beaconAddr))

	g.upstream, err = grpc.Dial(beaconAddr,
		grpc.WithInsecure())
	if err != nil {
		return err
	}

	director := g.director()

	g.proxy = proxy.NewProxy(g.upstream,
		grpc.UnknownServiceHandler(proxy.TransparentHandler(director)),
		grpc.StreamInterceptor(g.payloadInterceptor()))

	go func() {
		server := g.proxy
		if err := server.Serve(g.listener); err != nil {
			g.Logger.Panic("gRPC proxy server stopped", zap.Error(err))
		}
	}()

	return nil
}

func (g *GRPCRouter) Deinit() {
	g.Logger.Debug("Stopping grpc proxy")
	// GracefulStop doesn't close streams opened by the upstream, so call Stop instead
	g.proxy.Stop()
	g.listener.Close()
	g.upstream.Close()
}
