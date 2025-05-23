package router

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"

	"github.com/Rocket-Rescue-Node/credentials"
	"github.com/Rocket-Rescue-Node/credentials/pb"
	gbp "github.com/Rocket-Rescue-Node/guarded-beacon-proxy"
	"github.com/Rocket-Rescue-Node/rescue-proxy/config"
	"github.com/Rocket-Rescue-Node/rescue-proxy/consensuslayer"
	"github.com/Rocket-Rescue-Node/rescue-proxy/executionlayer"
	"github.com/Rocket-Rescue-Node/rescue-proxy/metrics"
	"github.com/ethereum/go-ethereum/common"
	rptypes "github.com/rocket-pool/rocketpool-go/types"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type ProxyRouter struct {
	Addr                 string
	BeaconURL            *url.URL
	GRPCAddr             string
	GRPCBeaconURL        string
	TLSCertFile          string
	TLSKeyFile           string
	Logger               *zap.Logger
	EL                   executionlayer.ExecutionLayer
	CL                   consensuslayer.ConsensusLayer
	CredentialSecrets    config.CredentialSecrets
	EnableSoloValidators bool

	gbp  *gbp.GuardedBeaconProxy
	m    *metrics.MetricsRegistry
	gm   *metrics.MetricsRegistry
	auth *auth
}

// Used to avoid collisions in context.WithValue()
// see: https://pkg.go.dev/context#WithValue
type prContextKey string

const prContextOperatorTypeKey = prContextKey("operator_type")
const prContextNodeAddrKey = prContextKey("node")

func (pr *ProxyRouter) logCredentialSharing(operatorType credentials.OperatorType,
	rpInfo *executionlayer.RPInfo,
	validatorInfo *consensuslayer.ValidatorInfo,
	credNodeAddr common.Address) {

	var chainNodeAddress common.Address

	if rpInfo == nil {
		if operatorType == pb.OperatorType_OT_ROCKETPOOL {
			// The credential was issued to a rp node, but the validator appears to be solo.
			pr.m.Counter("credential_sharing_rp_to_solo").Inc()
			return
		}
		// Solo validators on rv will not be looked up
		if validatorInfo == nil {
			return
		}
		// On pbp we will, so we can check
		chainNodeAddress = validatorInfo.WithdrawalAddress
	} else {
		if operatorType == pb.OperatorType_OT_SOLO {
			// The credential was issued to a solo validator, but appears to be used on a rp validator.
			pr.m.Counter("credential_sharing_solo_to_rp").Inc()
			return
		}

		chainNodeAddress = rpInfo.NodeAddress
	}

	if bytes.Equal(credNodeAddr[:], chainNodeAddress[:]) {
		return
	}

	// Someone got a credential with one node and used it on a different node.
	switch operatorType {
	case pb.OperatorType_OT_SOLO:
		pr.m.Counter("credential_sharing_solo_to_solo").Inc()
	case pb.OperatorType_OT_ROCKETPOOL:
		pr.m.Counter("credential_sharing_rp_to_rp").Inc()
	}
}

func (pr *ProxyRouter) readContext(ctx context.Context) ([]byte, credentials.OperatorType, error) {
	// Grab the authorized node address, only used for metrics/logging
	authedNode, ok := ctx.Value(prContextNodeAddrKey).([]byte)
	if !ok {
		return nil, 0, fmt.Errorf("unable to retrieve node address")
	}

	// Grab the credential type, only used for metrics/logging
	operatorType, ok := ctx.Value(prContextOperatorTypeKey).(credentials.OperatorType)
	if !ok {
		return nil, 0, fmt.Errorf("unable to retrieve operator_type")
	}

	return authedNode, operatorType, nil
}

func (pr *ProxyRouter) prepareBeaconProposerGuard(proposers gbp.PrepareBeaconProposerRequest, ctx context.Context) (gbp.AuthenticationStatus, error) {
	pr.m.Counter("prepare_beacon_proposer").Inc()

	// Create a slice of the indices
	indices := make([]string, 0, len(proposers))

	// Validate each proposer's fee recipient
	for _, proposer := range proposers {
		indices = append(indices, proposer.ValidatorIndex)
	}

	// Get the index->info map
	validatorMap, err := pr.CL.GetValidatorInfo(indices)
	if err != nil {
		pr.Logger.Error("Error while querying CL for validator info", zap.Error(err))
		return gbp.InternalError, nil
	}

	authedNode, operatorType, err := pr.readContext(ctx)
	if err != nil {
		pr.Logger.Warn("Error reading cached data from request context", zap.Error(err))
		return gbp.InternalError, nil
	}

	// Iterate the results and check the fee recipients against our expected values
	// Note: we iterate the map from the HTTP request to ensure every key is present in the
	// response from the consensuslayer abstraction
	for _, proposer := range proposers {
		validatorInfo, found := validatorMap[proposer.ValidatorIndex]
		if !found {
			pr.Logger.Warn("Pubkey for index not found in response from cl.",
				zap.String("requested index", proposer.ValidatorIndex))
			return gbp.BadRequest, fmt.Errorf("unknown validator index %s", proposer.ValidatorIndex)
		}

		pubkey := validatorInfo.Pubkey

		// Next we need to get the expected fee recipient for the pubkey
		rpInfo, err := pr.EL.GetRPInfo(pubkey)
		if err != nil {
			pr.Logger.Panic("error querying cache", zap.Error(err))
			return gbp.InternalError, fmt.Errorf("error with cache, please report it to Rescue Node maintainers")
		}

		pr.logCredentialSharing(operatorType, rpInfo, validatorInfo, common.BytesToAddress(authedNode))

		if rpInfo == nil {
			// Solo validators may only use their withdrawal credential in prepare_beacon_proposer
			if !validatorInfo.IsELWithdrawal ||
				!strings.EqualFold(validatorInfo.WithdrawalAddress.String(), proposer.FeeRecipient) {

				pr.m.Counter("prepare_beacon_incorrect_fee_recipient_solo").Inc()
				return gbp.Forbidden,
					fmt.Errorf("attempting to set fee recipient to %s differs from 0x01 credential withdrawal address %x",
						proposer.FeeRecipient,
						validatorInfo.WithdrawalAddress,
					)
			}

			pr.m.Counter("prepare_beacon_correct_fee_recipient_solo").Inc()
			metrics.ObserveSoloValidator(validatorInfo.WithdrawalAddress, validatorInfo.Pubkey)
			continue
		}

		// If the fee recipient matches expectations, good, move on to the next one
		if strings.EqualFold(rpInfo.ExpectedFeeRecipient.String(), proposer.FeeRecipient) {
			pr.m.Counter("prepare_beacon_correct_fee_recipient").Inc()
			metrics.ObserveValidator(rpInfo.NodeAddress, pubkey)
			continue
		}

		// rETH address is a 'safe' default fee recipient, and should be allowed.
		// However, it does indicate a misconfigured node, so log it.
		if strings.EqualFold(pr.EL.REthAddress().String(), proposer.FeeRecipient) {
			pr.m.Counter("prepare_beacon_reth_fee_recipient").Inc()
			pr.Logger.Warn("prepare_beacon_proposer called with rETH fee recipient",
				zap.String("expected", rpInfo.ExpectedFeeRecipient.String()),
				zap.String("node", rpInfo.NodeAddress.String()))
			continue
		}

		// Looks like a cheater- fee recipient doesn't match expectations
		pr.m.Counter("prepare_beacon_incorrect_fee_recipient").Inc()
		pr.Logger.Warn("prepare_beacon_proposer called with unexpected fee recipient",
			zap.String("expected", rpInfo.ExpectedFeeRecipient.String()), zap.String("got", proposer.FeeRecipient))
		return gbp.Conflict, fmt.Errorf("actual fee recipient %s didn't match expected fee recipient %s",
			proposer.FeeRecipient,
			rpInfo.ExpectedFeeRecipient.String(),
		)
	}

	// At this point all the fee recipients match our expectations. Proxy the request
	return gbp.Allowed, nil
}

func (pr *ProxyRouter) registerValidatorGuard(validators gbp.RegisterValidatorRequest, ctx context.Context) (gbp.AuthenticationStatus, error) {
	pr.m.Counter("register_validator").Inc()

	authedNode, operatorType, err := pr.readContext(ctx)
	if err != nil {
		pr.Logger.Warn("Error reading cached data from request context", zap.Error(err))
		return gbp.InternalError, nil
	}

	for _, validator := range validators {
		pubkeyStr := strings.TrimPrefix(validator.Message.Pubkey, "0x")

		pubkey, err := rptypes.HexToValidatorPubkey(pubkeyStr)
		if err != nil {
			pr.Logger.Warn("Malformed pubkey in register_validator_request", zap.Error(err), zap.String("pubkey", pubkeyStr))
			return gbp.BadRequest, fmt.Errorf("error parsing pubkey from request body: %v", err)
		}

		// Grab the expected fee recipient for the pubkey
		rpInfo, err := pr.EL.GetRPInfo(pubkey)
		if err != nil {
			pr.Logger.Panic("error querying cache", zap.Error(err))
			return gbp.InternalError, fmt.Errorf("error with cache, please report it to Rescue Node maintainers")
		}

		pr.logCredentialSharing(operatorType, rpInfo, nil, common.BytesToAddress(authedNode))
		if rpInfo == nil {
			// Solo validators can do whatever they want in register_validator.
			// The endpoint requires a signature, which the BN will validate, so
			// we know for sure that the downstream user has custody of the BLS keys.

			// The only thing to do is record some metrics

			feeRecipient := common.HexToAddress(validator.Message.FeeRecipient)
			metrics.ObserveSoloValidator(feeRecipient, pubkey)
			continue
		}

		if strings.EqualFold(rpInfo.ExpectedFeeRecipient.String(), validator.Message.FeeRecipient) {
			// This fee recipient matches expectations, carry on to the next validator
			pr.m.Counter("register_validator_correct_fee_recipient").Inc()
			metrics.ObserveValidator(rpInfo.NodeAddress, pubkey)
			continue
		}

		if strings.EqualFold(pr.EL.REthAddress().String(), validator.Message.FeeRecipient) {
			// rETH address is a 'safe' default fee recipient, and should be allowed.
			// However, it does indicate a misconfigured node, so log it.
			pr.m.Counter("register_validator_reth_fee_recipient").Inc()
			pr.Logger.Warn("register_validator called with rETH fee recipient",
				zap.String("expected", rpInfo.ExpectedFeeRecipient.String()),
				zap.String("node", rpInfo.NodeAddress.String()))
			continue
		}

		pr.m.Counter("register_validator_incorrect_fee_recipient").Inc()
		pr.Logger.Warn("register_validator called with unexpected fee recipient",
			zap.String("expected", rpInfo.ExpectedFeeRecipient.String()),
			zap.String("got", validator.Message.FeeRecipient),
		)
		return gbp.Conflict, fmt.Errorf("actual fee recipient %s didn't match expected fee recipient %s",
			validator.Message.FeeRecipient,
			rpInfo.ExpectedFeeRecipient.String(),
		)

	}

	// At this point all the fee recipients match our expectations. Proxy the request
	return gbp.Allowed, nil
}

// https://github.com/ChainSafe/lodestar/issues/6154
func (pr *ProxyRouter) urlDecode(username, password string) (string, string) {

	u, err := url.QueryUnescape(username)
	if err != nil {
		u = username
	}

	p, err := url.QueryUnescape(password)
	if err != nil {
		p = password
	}

	if username != u || password != p {
		pr.m.Counter("url_decoded").Inc()
	}

	return u, p
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

	// https://github.com/ChainSafe/lodestar/issues/6154
	username, password = pr.urlDecode(username, password)

	ac, err := pr.auth.authenticate(username, password)
	if err != nil {
		pr.m.Counter("unauthed").Inc()
		pr.Logger.Debug("Unable to authenticate credentials", zap.Error(err))
		return err.gbpStatus, nil, err
	}

	if ac.id.Equals(pr.auth.credentialManager.ID()) {
		pr.m.Counter("own_hmac").Inc()
	} else {
		pr.Logger.Debug(
			"authenticated request from partner cluster",
			zap.Binary("node_id", ac.Credential.NodeId),
			zap.String("secret", ac.id.String()),
		)
		pr.m.Counter("partner_hmac").Inc()
	}

	// If auth succeeds:
	if ac.Credential.OperatorType == pb.OperatorType_OT_ROCKETPOOL {
		pr.m.Counter("auth_ok").Inc()
	} else {
		// If we're dropping solo traffic, 429 it here
		if !pr.EnableSoloValidators {
			return gbp.TooManyRequests, nil, fmt.Errorf("solo validator support was manually disabled, but may be restored later")
		}
		pr.m.Counter("auth_ok_solo").Inc()
	}
	pr.Logger.Debug("Proxying Guarded URI", zap.String("uri", r.RequestURI))
	// Add the node address to the request context
	ctx := context.WithValue(r.Context(), prContextNodeAddrKey, ac.Credential.NodeId)
	// Add the operator type to the request context
	ctx = context.WithValue(ctx, prContextOperatorTypeKey, ac.Credential.OperatorType)
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

	ac, err := pr.auth.authenticate(auth[0], auth[1])
	if err != nil {
		pr.gm.Counter("unauthed").Inc()
		pr.Logger.Debug("Unable to authenticate credentials", zap.Error(err))
		return err.gbpStatus, nil, err
	}

	if ac.Credential.OperatorType == pb.OperatorType_OT_ROCKETPOOL {
		pr.gm.Counter("auth_ok").Inc()
	} else {
		// If we're dropping solo traffic, 429 it here
		if !pr.EnableSoloValidators {
			return gbp.TooManyRequests, nil, fmt.Errorf("solo validator support was manually disabled, but may be restored later")
		}
		pr.gm.Counter("auth_ok_solo").Inc()
	}

	ctx := context.WithValue(context.Background(), prContextNodeAddrKey, ac.Credential.NodeId)
	ctx = context.WithValue(ctx, prContextOperatorTypeKey, ac.Credential.OperatorType)
	return gbp.Allowed, ctx, nil
}

func (pr *ProxyRouter) Init() {
	// Initialize the auth handler
	pr.auth = initAuth(pr.CredentialSecrets)
	for _, id := range pr.auth.credentialManager.PartnerIDs() {
		pr.Logger.Info(
			"Loaded partner secret",
			zap.String("id", id.String()),
		)
	}
	pr.Logger.Info(
		"Initialized HMAC credentials",
		zap.Int("num", len(pr.CredentialSecrets)),
		zap.String("primary id", pr.auth.credentialManager.ID().String()),
	)

	// Create the reverse proxy.
	pr.gbp = &gbp.GuardedBeaconProxy{
		BeaconURL:                  pr.BeaconURL,
		GRPCBeaconURL:              pr.GRPCBeaconURL,
		HTTPAuthenticator:          pr.authenticate,
		GRPCAuthenticator:          pr.grpcAuthenticate,
		PrepareBeaconProposerGuard: pr.prepareBeaconProposerGuard,
		RegisterValidatorGuard:     pr.registerValidatorGuard,
	}
	pr.gbp.TLS.CertFile = pr.TLSCertFile
	pr.gbp.TLS.KeyFile = pr.TLSKeyFile

	pr.m = metrics.NewMetricsRegistry("http_proxy")
	pr.gm = metrics.NewMetricsRegistry("grpc_proxy")
}

func (pr *ProxyRouter) Start() error {
	pr.gbp.Addr = pr.Addr
	pr.gbp.GRPCAddr = pr.GRPCAddr
	return pr.gbp.ListenAndServe()
}

func (pr *ProxyRouter) Serve(httpListener net.Listener, grpcListener net.Listener) error {

	pr.gbp.Addr = httpListener.Addr().String()
	if grpcListener != nil {
		pr.gbp.GRPCAddr = grpcListener.Addr().String()
		return pr.gbp.Serve(httpListener, &grpcListener)
	}
	return pr.gbp.Serve(httpListener, nil)
}

func (pr *ProxyRouter) Stop(ctx context.Context) {
	pr.gbp.Stop(ctx)
}
