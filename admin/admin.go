package admin

import (
	"net"
	"net/http"

	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/metrics"
	"github.com/gorilla/mux"
)

type AdminApi struct {
	http.Server
}

func (a *AdminApi) Init(name string) error {
	var err error

	// Initialize metrics globals
	metricsHandler, err := metrics.Init(name)

	router := mux.NewRouter()

	a.Handler = router

	// Add admin handlers to the admin only http server and start it
	router.Path("/metrics").Handler(metricsHandler)

	return err
}

func (a *AdminApi) Serve(l net.Listener) error {
	a.Addr = l.Addr().String()
	return a.Server.Serve(l)
}
