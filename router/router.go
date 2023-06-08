package router

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/consensuslayer"
	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/executionlayer"
	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/metrics"
	gbp "github.com/Rocket-Rescue-Node/guarded-beacon-proxy"
	"github.com/ethereum/go-ethereum/common"
	rptypes "github.com/rocket-pool/rocketpool-go/types"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type ProxyRouter struct {
	Addr               string
	GRPCAddr           string
	GRPCBeaconURL      *url.URL
	TLSCertFile        string
	TLSKeyFile         string
	gbp                *gbp.GuardedBeaconProxy
	Logger             *zap.Logger
	EL                 *executionlayer.ExecutionLayer
	CL                 *consensuslayer.ConsensusLayer
	AuthValidityWindow time.Duration
	m                  *metrics.MetricsRegistry
	gm                 *metrics.MetricsRegistry
}

// Used to avoid collisions in context.WithValue()
// see: https://pkg.go.dev/context#WithValue
type prContextKey string

func (pr *ProxyRouter) prepareBeaconProposerGuard(proposers gbp.PrepareBeaconProposerRequest, ctx context.Context) (gbp.AuthenticationStatus, error) {
	pr.m.Counter("prepare_beacon_proposer").Inc()

	// Create a slice of the indices
	indices := make([]string, 0, len(proposers))

	// Validate each proposer's fee recipient
	for _, proposer := range proposers {
		indices = append(indices, proposer.ValidatorIndex)
	}

	// Get the index->pubkey map
	pubkeyMap, err := pr.CL.GetValidatorPubkey(indices)
	if err != nil {
		pr.Logger.Error("Error while querying CL for validator pubkeys", zap.Error(err))
		return gbp.InternalError, nil
	}

	// Grab the authorized node address
	authedNode, ok := ctx.Value(prContextKey("node")).([]byte)
	if !ok {
		pr.Logger.Warn("Unable to retrieve node address cached on request context")
		return gbp.InternalError, nil
	}
	authedNodeAddr := common.BytesToAddress(authedNode)

	// Iterate the results and check the fee recipients against our expected values
	// Note: we iterate the map from the HTTP request to ensure every key is present in the
	// response from the consensuslayer abstraction
	for _, proposer := range proposers {
		pubkey, found := pubkeyMap[proposer.ValidatorIndex]
		if !found {
			pr.Logger.Warn("Pubkey for index not found in response from cl.",
				zap.String("requested index", proposer.ValidatorIndex))
			return gbp.BadRequest, fmt.Errorf("unknown validator index %s", proposer.ValidatorIndex)
		}

		// Next we need to get the expected fee recipient for the pubkey
		expectedFeeRecipient, unowned := pr.EL.ValidatorFeeRecipient(pubkey, &authedNodeAddr)
		if expectedFeeRecipient == nil {
			pr.m.Counter("prepare_beacon_proposer_unowned").Inc()
			pr.Logger.Warn("Pubkey not found in EL cache, or wasn't owned by the user",
				zap.String("key", pubkey.String()),
				zap.Bool("someone else's validator", unowned))
			return gbp.Forbidden, fmt.Errorf("attempting to set fee recipient for unowned minipool")
		}

		// If the fee recipient matches expectations, good, move on to the next one
		if strings.EqualFold(expectedFeeRecipient.String(), proposer.FeeRecipient) {
			pr.m.Counter("prepare_beacon_correct_fee_recipient").Inc()
			metrics.ObserveValidator(authedNodeAddr, pubkey)
			continue
		}

		// rETH address is a 'safe' default fee recipient, and should be allowed.
		// However, it does indicate a misconfigured node, so log it.
		if strings.EqualFold(pr.EL.REthAddress().String(), proposer.FeeRecipient) {
			pr.m.Counter("prepare_beacon_reth_fee_recipient").Inc()
			pr.Logger.Warn("prepare_beacon_proposer called with rETH fee recipient",
				zap.String("expected", expectedFeeRecipient.String()),
				zap.String("node", authedNodeAddr.String()))
			continue
		}

		// Looks like a cheater- fee recipient doesn't match expectations
		pr.m.Counter("prepare_beacon_incorrect_fee_recipient").Inc()
		pr.Logger.Warn("prepare_beacon_proposer called with unexpected fee recipient",
			zap.String("expected", expectedFeeRecipient.String()), zap.String("got", proposer.FeeRecipient))
		return gbp.Conflict, fmt.Errorf("actual fee recipient %s didn't match expected fee recipient %s", proposer.FeeRecipient, expectedFeeRecipient.String())
	}

	// At this point all the fee recipients match our expectations. Proxy the request
	return gbp.Allowed, nil
}

func (pr *ProxyRouter) registerValidatorGuard(validators gbp.RegisterValidatorRequest, ctx context.Context) (gbp.AuthenticationStatus, error) {
	pr.m.Counter("register_validator").Inc()

	// Grab the authorized node address
	authedNode, ok := ctx.Value(prContextKey("node")).([]byte)
	if !ok {
		pr.Logger.Warn("Unable to retrieve node address cached on request context")
		return gbp.InternalError, nil
	}
	authedNodeAddr := common.BytesToAddress(authedNode)

	for _, validator := range validators {
		pubkeyStr := strings.TrimPrefix(validator.Message.Pubkey, "0x")

		pubkey, err := rptypes.HexToValidatorPubkey(pubkeyStr)
		if err != nil {
			pr.Logger.Warn("Malformed pubkey in register_validator_request", zap.Error(err), zap.String("pubkey", pubkeyStr))
			return gbp.InternalError, nil
		}

		// Grab the expected fee recipient for the pubkey
		expectedFeeRecipient, unowned := pr.EL.ValidatorFeeRecipient(pubkey, &authedNodeAddr)
		if expectedFeeRecipient == nil {
			// When unowned is true for register_validators, it means the pubkey was someone else's minipool
			// we still want that to get rejected... however, if unowned is false and expectedFeeRecipient is nil,
			// it means we're seeing a solo validator using mev-boost. Since register_validator requires a signature,
			// we can allow this fee recipient.
			if !unowned {
				pr.m.Counter("register_validator_not_minipool").Inc()
				metrics.ObserveValidator(authedNodeAddr, pubkey)
				// Move on to the next pubkey
				continue
			}
			pr.Logger.Warn("Pubkey not found in EL cache. Not an RP validator?", zap.String("key", pubkey.String()))
			return gbp.Forbidden, fmt.Errorf("unknown validator %s", pubkey.String())
		}

		if strings.EqualFold(expectedFeeRecipient.String(), validator.Message.FeeRecipient) {
			// This fee recipient matches expectations, carry on to the next validator
			pr.m.Counter("register_validator_correct_fee_recipient").Inc()
			metrics.ObserveValidator(authedNodeAddr, pubkey)
			continue
		}

		if strings.EqualFold(pr.EL.REthAddress().String(), validator.Message.FeeRecipient) {
			// rETH address is a 'safe' default fee recipient, and should be allowed.
			// However, it does indicate a misconfigured node, so log it.
			pr.m.Counter("register_validator_reth_fee_recipient").Inc()
			pr.Logger.Warn("register_validator called with rETH fee recipient",
				zap.String("expected", expectedFeeRecipient.String()),
				zap.String("node", authedNodeAddr.String()))
			continue
		}

		pr.m.Counter("register_validator_incorrect_fee_recipient").Inc()
		pr.Logger.Warn("register_validator called with unexpected fee recipient",
			zap.String("expected", expectedFeeRecipient.String()), zap.String("got", validator.Message.FeeRecipient))
		return gbp.Conflict, fmt.Errorf("actual fee recipient %s didn't match expected fee recipient %s", validator.Message.FeeRecipient, expectedFeeRecipient.String())

	}

	// At this point all the fee recipients match our expectations. Proxy the request
	return gbp.Allowed, nil
}

// Adds authentication to any handler.
func (pr *ProxyRouter) authenticate(r *http.Request) (gbp.AuthenticationStatus, context.Context, error) {

	// Authenticate the request here, return 403 and exit early as needed.
	// Start by grabbing basicauth
	username, password, ok := r.BasicAuth()
	if !ok {
		pr.m.Counter("missing_credentials").Inc()
		pr.Logger.Debug("Received request with no credentials on guarded endpoint")
		return gbp.Unauthorized, nil, fmt.Errorf("missing credentials")
	}

	ac, err := authenticate(username, password)
	if err != nil {
		pr.m.Counter("unauthed").Inc()
		pr.Logger.Debug("Unable to authenticate credentials", zap.Error(err))
		return err.gbpStatus, nil, err
	}

	// If auth succeeds:
	pr.m.Counter("auth_ok").Inc()
	pr.Logger.Debug("Proxying Guarded URI", zap.String("uri", r.RequestURI))
	// Add the node address to the request context
	ctx := context.WithValue(r.Context(), prContextKey("node"), ac.Credential.NodeId)
	return gbp.Allowed, ctx, nil
}

func (pr *ProxyRouter) grpcAuthenticate(md metadata.MD) (gbp.AuthenticationStatus, context.Context, error) {
	val, exists := md["rprnauth"]
	if !exists || len(val) < 1 {
		pr.gm.Counter("auth_header_missing").Inc()
		pr.Logger.Debug("grpc access without auth header", zap.Bool("exists", exists))
		return gbp.Unauthorized, nil, fmt.Errorf("headers missing")
	}

	auth := strings.Split(val[0], ":")
	if len(auth) != 2 {
		pr.gm.Counter("auth_header_malformed").Inc()
		pr.Logger.Debug("grpc access with invalid auth header")
		return gbp.Unauthorized, nil, fmt.Errorf("headers invalid")
	}

	ac, err := authenticate(auth[0], auth[1])
	if err != nil {
		pr.gm.Counter("unauthed").Inc()
		pr.Logger.Debug("Unable to authenticate credentials", zap.Error(err))
		return err.gbpStatus, nil, err
	}

	pr.gm.Counter("auth_ok").Inc()

	ctx := context.WithValue(context.Background(), prContextKey("node"), ac.Credential.NodeId)
	return gbp.Allowed, ctx, nil
}

func (pr *ProxyRouter) Init(beaconNode *url.URL) error {

	// Create the reverse proxy.
	pr.gbp = &gbp.GuardedBeaconProxy{
		BeaconURL:                  beaconNode,
		Addr:                       pr.Addr,
		HTTPAuthenticator:          pr.authenticate,
		GRPCAuthenticator:          pr.grpcAuthenticate,
		PrepareBeaconProposerGuard: pr.prepareBeaconProposerGuard,
		RegisterValidatorGuard:     pr.registerValidatorGuard,
	}
	pr.gbp.GRPCAddr = pr.GRPCAddr
	pr.gbp.GRPCBeaconURL = pr.GRPCBeaconURL
	pr.gbp.TLS.CertFile = pr.TLSCertFile
	pr.gbp.TLS.KeyFile = pr.TLSKeyFile

	pr.m = metrics.NewMetricsRegistry("http_proxy")
	pr.gm = metrics.NewMetricsRegistry("grpc_proxy")
	return pr.gbp.ListenAndServe()
}

func (pr *ProxyRouter) Deinit(grace time.Duration) {
	pr.gbp.Stop(grace)
}
