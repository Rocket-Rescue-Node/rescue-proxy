package router

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Rocket-Rescue-Node/credentials"
	"github.com/Rocket-Rescue-Node/credentials/pb"
	gbp "github.com/Rocket-Rescue-Node/guarded-beacon-proxy"
	"github.com/Rocket-Rescue-Node/rescue-proxy/config"
	"github.com/Rocket-Rescue-Node/rescue-proxy/consensuslayer"
	"github.com/Rocket-Rescue-Node/rescue-proxy/executionlayer"
	"github.com/Rocket-Rescue-Node/rescue-proxy/metrics"
	"github.com/ethereum/go-ethereum/common"
	rptypes "github.com/rocket-pool/smartnode/bindings/types"
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

type authInfo struct {
	nodeAddr     common.Address
	operatorType credentials.OperatorType
}

// Used to avoid collisions in context.WithValue()
// see: https://pkg.go.dev/context#WithValue
type prContextKey string

const prContextOperatorTypeKey = prContextKey("operator_type")
const prContextNodeAddrKey = prContextKey("node")

func (pr *ProxyRouter) logCredentialSharing(authInfo authInfo,
	rpInfo *executionlayer.RPInfo,
	validatorInfo *consensuslayer.ValidatorInfo,
) {

	var chainNodeAddress common.Address
	operatorType := authInfo.operatorType
	credNodeAddr := authInfo.nodeAddr

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

func (pr *ProxyRouter) readContext(ctx context.Context) (authInfo, error) {
	// Grab the authorized node address, only used for metrics/logging
	authedNode, ok := ctx.Value(prContextNodeAddrKey).([]byte)
	if !ok {
		return authInfo{}, fmt.Errorf("unable to retrieve node address")
	}

	// Grab the credential type, only used for metrics/logging
	operatorType, ok := ctx.Value(prContextOperatorTypeKey).(credentials.OperatorType)
	if !ok {
		return authInfo{}, fmt.Errorf("unable to retrieve operator_type")
	}

	return authInfo{
		nodeAddr:     common.BytesToAddress(authedNode),
		operatorType: operatorType,
	}, nil
}

func (pr *ProxyRouter) rocketPoolPBPGuard(authInfo authInfo, validatorInfo *consensuslayer.ValidatorInfo, feeRecipient string) (gbp.AuthenticationStatus, error) {
	pubkey := validatorInfo.Pubkey

	rpInfo, err := pr.EL.GetRPInfo(pubkey)
	if err != nil {
		pr.Logger.Panic("error querying cache", zap.Error(err))
		return gbp.InternalError, fmt.Errorf("error with cache, please report it to Rescue Node maintainers")
	}

	pr.logCredentialSharing(authInfo, rpInfo, validatorInfo)

	if rpInfo == nil {
		// This is not a RP validator, so we can return with no error
		// But don't claim it's authorized, because it may not be- simply, it isn't authorized
		// by this function
		return gbp.Unauthorized, nil
	}

	// If the fee recipient matches expectations, good, move on to the next one
	if strings.EqualFold(rpInfo.ExpectedFeeRecipient.String(), feeRecipient) {
		pr.m.Counter("prepare_beacon_correct_fee_recipient").Inc()
		metrics.ObserveValidator(rpInfo.NodeAddress, pubkey)
		return gbp.Allowed, nil
	}

	// rETH address is a 'safe' default fee recipient, and should be allowed.
	// However, it does indicate a misconfigured node, so log it.
	if strings.EqualFold(pr.EL.REthAddress().String(), feeRecipient) {
		pr.m.Counter("prepare_beacon_reth_fee_recipient").Inc()
		pr.Logger.Warn("prepare_beacon_proposer called with rETH fee recipient",
			zap.String("expected", rpInfo.ExpectedFeeRecipient.String()),
			zap.String("node", rpInfo.NodeAddress.String()))
		metrics.ObserveValidator(rpInfo.NodeAddress, pubkey)
		return gbp.Allowed, nil
	}

	pr.m.Counter("prepare_beacon_incorrect_fee_recipient").Inc()
	pr.Logger.Warn("prepare_beacon_proposer called with unexpected fee recipient",
		zap.String("expected", rpInfo.ExpectedFeeRecipient.String()), zap.String("got", feeRecipient))
	return gbp.Conflict, fmt.Errorf("actual fee recipient %s didn't match expected fee recipient %s",
		feeRecipient,
		rpInfo.ExpectedFeeRecipient.String(),
	)
}

func (pr *ProxyRouter) stakewisePBPGuard(_ authInfo, validatorInfo *consensuslayer.ValidatorInfo, feeRecipient string) (gbp.AuthenticationStatus, error) {
	// Create a context with a 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check the vault registry for the validator's 0x01/2 address
	stakewiseFeeRecipient, err := pr.EL.StakewiseFeeRecipient(ctx, validatorInfo.WithdrawalAddress)
	if err != nil {
		pr.Logger.Error("Error while querying stakewise for fee recipient", zap.Error(err))
		return gbp.InternalError, fmt.Errorf("error querying stakewise for fee recipient: %w", err)
	}

	if stakewiseFeeRecipient == nil {
		// This is not a stakewise validator, so we can return with no error
		// But don't claim it's authorized, because it may not be- simply, it isn't authorized
		// by this function
		return gbp.Unauthorized, nil
	}

	if strings.EqualFold(stakewiseFeeRecipient.String(), feeRecipient) {
		pr.m.Counter("prepare_beacon_correct_fee_recipient").Inc()
		// Increment the solo metrics- we're not tracking stakewise use, we just don't want to see
		// legitimate users with stakewise validators shoot themselves in the foot by using the rescue node.
		metrics.ObserveSoloValidator(*stakewiseFeeRecipient, validatorInfo.Pubkey)
		return gbp.Allowed, nil
	}

	return gbp.Conflict, fmt.Errorf("validator is stakewise, but fee recipient doesn't match the mevEscrow contract address at %s", stakewiseFeeRecipient.String())
}

func (pr *ProxyRouter) soloPBPGuard(_ authInfo, validatorInfo *consensuslayer.ValidatorInfo, feeRecipient string) (gbp.AuthenticationStatus, error) {
	if !validatorInfo.IsELWithdrawal {
		pr.m.Counter("prepare_beacon_0x00_solo").Inc()
		return gbp.Forbidden,
			fmt.Errorf("attempting to set fee recipient for pubkey %s, but validator has no EL withdrawal address",
				validatorInfo.Pubkey,
			)
	}
	if !strings.EqualFold(validatorInfo.WithdrawalAddress.String(), feeRecipient) {
		pr.m.Counter("prepare_beacon_0x00_solo").Inc()
		return gbp.Forbidden,
			fmt.Errorf("attempting to set fee recipient to %s differs from 0x01 credential withdrawal address %x",
				feeRecipient,
				validatorInfo.WithdrawalAddress,
			)
	}

	// If the fee recipient is equal to the withdrawal address, allow it
	pr.m.Counter("prepare_beacon_correct_fee_recipient_solo").Inc()
	metrics.ObserveSoloValidator(validatorInfo.WithdrawalAddress, validatorInfo.Pubkey)
	return gbp.Allowed, nil
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

	authInfo, err := pr.readContext(ctx)
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

		authStatus, err := pr.rocketPoolPBPGuard(authInfo, validatorInfo, proposer.FeeRecipient)
		if err != nil {
			return authStatus, err
		}

		if authStatus == gbp.Allowed {
			continue
		}

		// Rocket pool check didn't forbid this fee recipient, but also didn't authorize it.
		// Try further checks.

		authStatus, err = pr.stakewisePBPGuard(authInfo, validatorInfo, proposer.FeeRecipient)
		if err != nil {
			return authStatus, err
		}

		if authStatus == gbp.Allowed {
			continue
		}

		// Finally, check solo validators fee recipients match their withdrawal address
		authStatus, err = pr.soloPBPGuard(authInfo, validatorInfo, proposer.FeeRecipient)
		if err != nil {
			return authStatus, err
		}

		if authStatus == gbp.Allowed {
			continue
		} else {
			pr.Logger.Panic("soloPBPGuard should always return allowed or an error")
		}
	}

	// At this point all the fee recipients match our expectations. Proxy the request
	return gbp.Allowed, nil
}

func (pr *ProxyRouter) registerValidatorGuard(validators gbp.RegisterValidatorRequest, ctx context.Context) (gbp.AuthenticationStatus, error) {
	pr.m.Counter("register_validator").Inc()

	authInfo, err := pr.readContext(ctx)
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

		pr.logCredentialSharing(authInfo, rpInfo, nil)
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
