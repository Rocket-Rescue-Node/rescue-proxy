package router

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/Rocket-Pool-Rescue-Node/credentials"
	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/consensuslayer"
	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/executionlayer"
	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/metrics"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
	rptypes "github.com/rocket-pool/rocketpool-go/types"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type ProxyRouter struct {
	proxy              *httputil.ReverseProxy
	Logger             *zap.Logger
	EL                 *executionlayer.ExecutionLayer
	CL                 *consensuslayer.ConsensusLayer
	CM                 *credentials.CredentialManager
	AuthValidityWindow time.Duration
	m                  *metrics.MetricsRegistry
}

// Used to avoid collisions in context.WithValue()
// see: https://pkg.go.dev/context#WithValue
type prContextKey string

func cloneRequestBody(r *http.Request) (io.ReadCloser, error) {
	// Read the body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	original := ioutil.NopCloser(bytes.NewBuffer(buf))
	clone := ioutil.NopCloser(bytes.NewBuffer(buf))
	r.Body = original
	return clone, nil
}

func (pr *ProxyRouter) prepareBeaconProposer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pr.m.Counter("prepare_beacon_proposer").Inc()
		// Clone the request body so it can still be proxied
		buf, err := cloneRequestBody(r)
		if err != nil {
			pr.Logger.Warn("Error cloning prepare_beacon_proposers request body", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Parse JSON body of request
		var proposers consensuslayer.PrepareBeaconProposerRequest
		if err := json.NewDecoder(buf).Decode(&proposers); err != nil {
			pr.Logger.Warn("Malformed prepare_beacon_proposers request", zap.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

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
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Grab the authorized node address
		authedNode, ok := r.Context().Value(prContextKey("node")).([]byte)
		if !ok {
			pr.Logger.Warn("Unable to retrieve node address cached on request context")
			w.WriteHeader(http.StatusInternalServerError)
			return
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
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			// Next we need to get the expected fee recipient for the pubkey
			expectedFeeRecipient, unowned := pr.EL.ValidatorFeeRecipient(pubkey, &authedNodeAddr)
			if expectedFeeRecipient == nil {
				pr.m.Counter("prepare_beacon_proposer_unowned").Inc()
				pr.Logger.Warn("Pubkey not found in EL cache, or wasn't owned by the user",
					zap.String("key", pubkey.String()),
					zap.Bool("someone else's validator", unowned))
				w.WriteHeader(http.StatusForbidden)
				return
			}
			if !strings.EqualFold(expectedFeeRecipient.String(), proposer.FeeRecipient) {
				// Looks like a cheater- fee recipient doesn't match expectations
				pr.m.Counter("prepare_beacon_incorrect_fee_recipient").Inc()
				pr.Logger.Warn("prepare_beacon_proposer called with unexpected fee recipient",
					zap.String("expected", expectedFeeRecipient.String()), zap.String("got", proposer.FeeRecipient))
				w.WriteHeader(http.StatusConflict)
				return
			}

			pr.m.Counter("prepare_beacon_correct_fee_recipient").Inc()
			metrics.ObserveValidator(authedNodeAddr, pubkey)
		}

		// At this point all the fee recipients match our expectations. Proxy the request
		pr.proxy.ServeHTTP(w, r)
	}
}

func (pr *ProxyRouter) registerValidator() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pr.m.Counter("register_validator").Inc()
		// Clone the request body so it can still be proxied
		buf, err := cloneRequestBody(r)
		if err != nil {
			pr.Logger.Warn("Error cloning register_validator request body", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Parse JSON body of request
		var validators consensuslayer.RegisterValidatorRequest
		if err := json.NewDecoder(buf).Decode(&validators); err != nil {
			pr.Logger.Warn("Malformed register_validator request", zap.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Grab the authorized node address
		authedNode, ok := r.Context().Value(prContextKey("node")).([]byte)
		if !ok {
			pr.Logger.Warn("Unable to retrieve node address cached on request context")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		authedNodeAddr := common.BytesToAddress(authedNode)

		for _, validator := range validators {
			pubkeyStr := strings.TrimPrefix(validator.Message.Pubkey, "0x")

			pubkey, err := rptypes.HexToValidatorPubkey(pubkeyStr)
			if err != nil {
				pr.Logger.Warn("Malformed pubkey in register_validator_request", zap.Error(err), zap.String("pubkey", pubkeyStr))
				w.WriteHeader(http.StatusInternalServerError)
				return
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
				w.WriteHeader(http.StatusForbidden)
				return
			}

			if !strings.EqualFold(expectedFeeRecipient.String(), validator.Message.FeeRecipient) {
				pr.m.Counter("register_validator_incorrect_fee_recipient").Inc()
				pr.Logger.Warn("register_validator called with unexpected fee recipient",
					zap.String("expected", expectedFeeRecipient.String()), zap.String("got", validator.Message.FeeRecipient))
				w.WriteHeader(http.StatusConflict)
				return
			}

			// This fee recipient matches expectations, carry on to the next validator
			pr.m.Counter("register_validator_correct_fee_recipient").Inc()
			metrics.ObserveValidator(authedNodeAddr, pubkey)
		}

		// At this point all the fee recipients match our expectations. Proxy the request
		pr.proxy.ServeHTTP(w, r)
	}
}

// Adds authentication to any handler.
func (pr *ProxyRouter) authenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If this is an "internal" request, do not bother with auth
		if strings.HasPrefix(r.RequestURI, "/_/") {
			pr.Logger.Debug("Request on unauthenticated endpoint", zap.String("uri", r.RequestURI))
			next.ServeHTTP(w, r)
			return
		}

		// Authenticate the request here, return 403 and exit early as needed.
		// Start by grabbing basicauth
		username, password, ok := r.BasicAuth()
		if !ok {
			pr.m.Counter("missing_credentials").Inc()
			pr.Logger.Debug("Received request with no credentials on guarded endpoint")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// The username is just the base64 encoding of the node id. We'll want to check it against the credential, so decode it
		decoder := base64.NewDecoder(base64.URLEncoding, bytes.NewReader([]byte(username)))
		nodeID, err := io.ReadAll(decoder)
		if err != nil {
			pr.m.Counter("malformed_username").Inc()
			pr.Logger.Debug("Received request with malformed username on guarded endpoint", zap.Error(err))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// The password is our hmac credential
		decoder = base64.NewDecoder(base64.URLEncoding, bytes.NewReader([]byte(password)))
		decoded, err := io.ReadAll(decoder)
		if err != nil {
			pr.m.Counter("malformed_password").Inc()
			pr.Logger.Debug("Received request with malformed password on guarded endpoint", zap.Error(err))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Unmarshal the protobuf
		unmarshaled := &credentials.AuthenticatedCredential{}
		err = proto.Unmarshal(decoded, unmarshaled.Pb())
		if err != nil {
			pr.m.Counter("malformed_protobuf").Inc()
			pr.Logger.Debug("Received request with malformed password protobuf on guarded endpoint", zap.Error(err))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// We don't require the nodeID is present in the password, since it is passed as the user,
		// however, it was used to generate the hmac, so repopulate it
		// This also ensures that the password was generated for the requesting node
		unmarshaled.Credential.NodeId = nodeID
		err = pr.CM.Verify(unmarshaled)
		if err != nil {
			pr.m.Counter("invalid_hmac").Inc()
			pr.Logger.Debug("Unable to verify hmac on guarded endpoint", zap.Error(err))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Grab the timestamp and make sure the credential is recent enough
		ts := time.Unix(unmarshaled.Credential.Timestamp, 0)
		now := time.Now()

		if ts.Before(now) && now.Sub(ts) > pr.AuthValidityWindow {
			pr.m.Counter("expired_credentials").Inc()
			pr.Logger.Debug("Stale credential seen on guarded endpoint")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// If auth succeeds:
		pr.m.Counter("auth_ok").Inc()
		pr.Logger.Debug("Proxying Guarded URI", zap.String("uri", r.RequestURI))
		// Add the node address to the request context
		ctx := context.WithValue(r.Context(), prContextKey("node"), nodeID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (pr *ProxyRouter) Init(beaconNode *url.URL) {

	// Create the reverse proxy.
	pr.proxy = httputil.NewSingleHostReverseProxy(beaconNode)

	pr.m = metrics.NewMetricsRegistry("http_proxy")

	router := mux.NewRouter()

	// Path to check the status of the rescue node. Simply 200 OK.
	router.Path("/_/status").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pr.Logger.Debug("Received healthcheck, replying 200 OK")
		_, err := w.Write([]byte("OK\n"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		pr.m.Counter("status").Inc()
		w.WriteHeader(http.StatusOK)
	})

	router.Path("/eth/v1/validator/prepare_beacon_proposer").
		HandlerFunc(pr.prepareBeaconProposer())

	router.Path("/eth/v1/validator/register_validator").
		HandlerFunc(pr.registerValidator())

	// By default, simply reverse-proxy every request
	router.PathPrefix("/").Handler(pr.proxy)

	// Install the authentication middleware
	router.Use(pr.authenticationMiddleware)
	http.Handle("/", router)
}
