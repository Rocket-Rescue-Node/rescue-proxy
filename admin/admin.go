package admin

import (
	"context"
	"net"
	"net/http"

	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/metrics"
	"github.com/gorilla/mux"
)

type AdminApi struct {
	http.Server

	metricsHandler http.Handler
}

func (a *AdminApi) Init(name string) error {
	var err error

	// Initialize metrics globals
	a.metricsHandler, err = metrics.Init(name)

	return err
}

func (a *AdminApi) Start(listenAddr string) error {

	router := mux.NewRouter()

	a.Handler = router

	// Add admin handlers to the admin only http server and start it
	router.Path("/metrics").Handler(a.metricsHandler)
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	a.Addr = listener.Addr().String()
	return a.Serve(listener)
}

func (a *AdminApi) Stop(ctx context.Context) error {
	// Attempt a graceful stop
	err := a.Shutdown(ctx)
	if err == nil {
		return nil
	}

	// Shutdown immediately if the context deadline is exceeded.
	return a.Close()
}
