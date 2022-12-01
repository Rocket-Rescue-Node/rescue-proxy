package router

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/execution_layer"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type proxyRouter struct {
	proxy  *httputil.ReverseProxy
	logger *zap.Logger
	el     *execution_layer.ExecutionLayer
}

type handler func(http.ResponseWriter, *http.Request)

func (pr *proxyRouter) prepareBeaconProposer() handler {
	return func(w http.ResponseWriter, r *http.Request) {
		pr.logger.Debug("Proxying Guarded URL", zap.String("url", r.URL.String()))
		// For now, simply proxy the request
		// TODO: inspect request `r` and enforce the dang rules
		pr.proxy.ServeHTTP(w, r)
	}
}

func (pr *proxyRouter) registerValidator() handler {
	return func(w http.ResponseWriter, r *http.Request) {
		pr.logger.Debug("Proxying Guarded URL", zap.String("url", r.URL.String()))
		// For now, simply proxy the request
		// TODO: inspect request `r` and enforce the dang rules
		pr.proxy.ServeHTTP(w, r)
	}
}

// Adds authentication to any handler.
// TODO: Implement, lol
func (pr *proxyRouter) authenticated(h handler) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		// Authenticate the request here, return 403 and exit early as needed.

		// If auth succeeds:
		pr.logger.Debug("Authentication succeeded")
		h(w, r)
	}
}

func NewProxyRouter(beaconNode *url.URL, el *execution_layer.ExecutionLayer, logger *zap.Logger) *mux.Router {
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
	}

	// Path to check the status of the rescue node. Simply 200 OK.
	out.Path("/status").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("Received healthcheck, replying 200 OK")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK\n"))
		return
	})

	// Some paths are "guarded"- we need custom logic to enforce the rules
	out.Path("/eth/v1/validator/prepare_beacon_proposer").
		HandlerFunc(pr.authenticated(pr.prepareBeaconProposer()))

	out.Path("/eth/v1/validator/register_validator").
		HandlerFunc(pr.authenticated(pr.registerValidator()))

	// By default, simply reverse-proxy every request
	out.PathPrefix("/").HandlerFunc(pr.authenticated(func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("Proxying unprotected URL", zap.String("url", r.URL.String()))
		proxy.ServeHTTP(w, r)
	}))

	return out
}
