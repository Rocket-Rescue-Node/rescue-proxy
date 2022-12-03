package router

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/consensuslayer"
	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/executionlayer"
	"github.com/gorilla/mux"
	rptypes "github.com/rocket-pool/rocketpool-go/types"
	"go.uber.org/zap"
)

type proxyRouter struct {
	proxy  *httputil.ReverseProxy
	logger *zap.Logger
	el     *executionlayer.ExecutionLayer
	cl     *consensuslayer.ConsensusLayer
}

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

func (pr *proxyRouter) prepareBeaconProposer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Clone the request body so it can still be proxied
		buf, err := cloneRequestBody(r)
		if err != nil {
			pr.logger.Warn("Error cloning prepare_beacon_proposers request body", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Parse JSON body of request
		var proposers consensuslayer.PrepareBeaconProposerRequest
		if err := json.NewDecoder(buf).Decode(&proposers); err != nil {
			pr.logger.Warn("Malformed prepare_beacon_proposers request", zap.Error(err))
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
		pubkeyMap, err := pr.cl.GetValidatorPubkey(indices)
		if err != nil {
			pr.logger.Error("Error while querying CL for validator pubkeys", zap.Error(err))
		}

		// Iterate the results and check the fee recipients against our expected values
		// Note: we iterate the map from the HTTP request to ensure every key is present in the
		// response from the consensuslayer abstraction
		for _, proposer := range proposers {
			pubkey, found := pubkeyMap[proposer.ValidatorIndex]
			if !found {
				pr.logger.Warn("Pubkey for index not found in response from cl.",
					zap.String("requested index", proposer.ValidatorIndex))
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			// Next we need to get the expected fee recipient for the pubkey
			expectedFeeRecipient := pr.el.ValidatorFeeRecipient(pubkey)
			if expectedFeeRecipient == nil {
				pr.logger.Warn("Pubkey not found in EL cache. Not an RP validator?", zap.String("key", pubkey.String()))
				w.WriteHeader(http.StatusForbidden)
				return
			}
			if !strings.EqualFold(expectedFeeRecipient.String(), proposer.FeeRecipient) {
				// Looks like a cheater- fee recipient doesn't match expectations
				pr.logger.Warn("prepare_beacon_proposer called with unexpected fee recipient",
					zap.String("expected", expectedFeeRecipient.String()), zap.String("got", proposer.FeeRecipient))
				w.WriteHeader(http.StatusConflict)
				return
			}
		}

		// At this point all the fee recipients match our expectations. Proxy the request
		pr.proxy.ServeHTTP(w, r)
	}
}

func (pr *proxyRouter) registerValidator() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Clone the request body so it can still be proxied
		buf, err := cloneRequestBody(r)
		if err != nil {
			pr.logger.Warn("Error cloning register_validator request body", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Parse JSON body of request
		var validators consensuslayer.RegisterValidatorRequest
		if err := json.NewDecoder(buf).Decode(&validators); err != nil {
			pr.logger.Warn("Malformed register_validator request", zap.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		for _, validator := range validators {
			pubkeyStr := strings.TrimPrefix(validator.Message.Pubkey, "0x")

			pubkey, err := rptypes.HexToValidatorPubkey(pubkeyStr)
			if err != nil {
				pr.logger.Warn("Malformed pubkey in register_validator_request", zap.Error(err), zap.String("pubkey", pubkeyStr))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Grab the expected fee recipient for the pubkey
			expectedFeeRecipient := pr.el.ValidatorFeeRecipient(pubkey)
			if expectedFeeRecipient == nil {
				pr.logger.Warn("Pubkey not found in EL cache. Not an RP validator?", zap.String("key", pubkey.String()))
				w.WriteHeader(http.StatusForbidden)
				return
			}

			if !strings.EqualFold(expectedFeeRecipient.String(), validator.Message.FeeRecipient) {
				pr.logger.Warn("register_validator called with unexpected fee recipient",
					zap.String("expected", expectedFeeRecipient.String()), zap.String("got", validator.Message.FeeRecipient))
				w.WriteHeader(http.StatusConflict)
				return
			}

			// This fee recipient matches expectations, carry on to the next validator
		}

		// At this point all the fee recipients match our expectations. Proxy the request
		pr.proxy.ServeHTTP(w, r)
	}
}

// Adds authentication to any handler.
// TODO: Implement, lol
func (pr *proxyRouter) authenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If this is an "internal" request, do not bother with auth
		if strings.HasPrefix(r.RequestURI, "/_/") {
			pr.logger.Debug("Request on unauthenticated endpoint", zap.String("uri", r.RequestURI))
			next.ServeHTTP(w, r)
			return
		}

		// Authenticate the request here, return 403 and exit early as needed.

		// If auth succeeds:
		pr.logger.Debug("Proxying Guarded URI", zap.String("uri", r.RequestURI))
		next.ServeHTTP(w, r)
	})
}

// NewProxyRouter creates a new mux.router with the provided beacon node URL, execution layer, and logger
func NewProxyRouter(beaconNode *url.URL, el *executionlayer.ExecutionLayer, cl *consensuslayer.ConsensusLayer, logger *zap.Logger) *mux.Router {
	out := mux.NewRouter()

	// Create the reverse proxy.
	proxy := httputil.NewSingleHostReverseProxy(beaconNode)

	// Create the go 'receiver' for convenience
	// Enables the prepareBeaconProposer closures et al to access
	// 'proxy' and 'logger' without explicitly passing them.
	pr := &proxyRouter{
		proxy,
		logger,
		el,
		cl,
	}

	// Path to check the status of the rescue node. Simply 200 OK.
	out.Path("/_/status").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("Received healthcheck, replying 200 OK")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK\n"))
	})

	out.Path("/eth/v1/validator/prepare_beacon_proposer").
		HandlerFunc(pr.prepareBeaconProposer())

	out.Path("/eth/v1/validator/register_validator").
		HandlerFunc(pr.registerValidator())

	// By default, simply reverse-proxy every request
	out.PathPrefix("/").Handler(proxy)

	// Install the authentication middleware
	out.Use(pr.authenticationMiddleware)

	return out
}
