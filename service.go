package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/admin"
	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/api"
	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/consensuslayer"
	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/executionlayer"
	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/router"
	"go.uber.org/zap"
)

// Service is a rescue-proxy service. It runs several goroutines that implement the features of
// rescue-proxy.
type Service struct {
	ctx    context.Context
	cancel context.CancelFunc

	// A [zap.Logger] to use for logging. If not provided, one will be initialized from the [Config]
	Logger *zap.Logger
	// A [Config] to use for initialization.
	Config *Config

	// Sub-services
	admin *admin.AdminApi
	el    *executionlayer.ExecutionLayer
	cl    *consensuslayer.ConsensusLayer
	r     *router.ProxyRouter
	a     *api.API

	// error reporting channel
	errs chan error
}

// NewService creates a [Service] from a given [Config].func NewService(config *Config) *Service {
func NewService(config *Config) *Service {
	return &Service{
		Config: config,
	}
}

// Run initializes the [Service] without blocking.
// Callers should read from the returned channel to detect errors.
// The provided [context.Context] cam becanceled to initiate a graceful shutdown,
// and the returned channel will be closed.
func (s *Service) Run(ctx context.Context) chan error {
	out := make(chan error, 32)
	go s.run(ctx, out)
	return out
}

func (s *Service) run(ctx context.Context, errs chan error) {
	s.errs = errs

	s.ctx, s.cancel = context.WithCancel(ctx)
	defer s.cancel()

	if s.Logger == nil {
		s.Logger = initLogger(s.Config)
	}
	s.Logger.Info("Starting up the rescue node proxy...")

	// Create the admin-only http server
	s.admin = new(admin.AdminApi)
	go func() {
		s.Logger.Info("Starting admin API")
		if err := s.admin.Start(s.Config.AdminListenAddr); err != nil {
			s.errs <- err
		}
	}()

	// Connect to and initialize the execution layer
	s.el = &executionlayer.ExecutionLayer{
		ECURL:             s.Config.ExecutionURL,
		RocketStorageAddr: s.Config.RocketStorageAddr,
		Logger:            s.Logger,
		CachePath:         s.Config.CachePath,
	}
	// Init() blocks until the cache is warmed up. This is good, we don't want to
	// start accepting http requests on the proxy until we're ready to handle them.
	if err := s.el.Init(); err != nil {
		s.errs <- fmt.Errorf("unable to init Execution Layer client: %v", err)
		return
	}
	// After Init() we still have to call Start() to subscribe to new blocks
	go func() {
		s.Logger.Info("Starting EL monitor")
		if err := s.el.Start(); err != nil {
			s.errs <- fmt.Errorf("EL error: %v", err)
		}
	}()

	// Connect to and initialize the consensus layer
	s.cl = consensuslayer.NewConsensusLayer(s.Config.BeaconURL, s.Logger)
	s.Logger.Info("Starting CL monitor")
	// Consensus Layer is non-blocking/synchronous only.
	// On Init() it will create the client and warm the validator key cache, which
	// is needed to serve responses to rescue-api
	if err := s.cl.Init(s.ctx); err != nil {
		// Optimization: serialize the EL cache by stopping it now so we can recover
		// faster.
		s.el.Stop()
		// Only write the error to the channel after so we don't panic while writing
		// the cache to disk
		s.errs <- fmt.Errorf("unable to init Consensus Layer client: %v", err)
		return
	}

	s.r = &router.ProxyRouter{
		Addr:                 s.Config.ListenAddr,
		BeaconURL:            s.Config.BeaconURL,
		GRPCAddr:             s.Config.GRPCListenAddr,
		GRPCBeaconURL:        s.Config.GRPCBeaconAddr,
		TLSCertFile:          s.Config.GRPCTLSCertFile,
		TLSKeyFile:           s.Config.GRPCTLSKeyFile,
		Logger:               s.Logger,
		EL:                   s.el,
		CL:                   s.cl,
		AuthValidityWindow:   s.Config.AuthValidityWindow,
		EnableSoloValidators: s.Config.EnableSoloValidators,
		CredentialSecret:     s.Config.CredentialSecret,
	}
	// Spin up the rest of the servers on different goroutines, since they block.
	go func() {
		s.Logger.Info("Starting http server", zap.String("url", s.Config.ListenAddr))
		if err := s.r.Start(); err != nil {
			s.errs <- err
		}
	}()

	s.a = &api.API{
		EL:         s.el,
		CL:         s.cl,
		Logger:     s.Logger,
		ListenAddr: s.Config.APIListenAddr,
	}
	go func() {
		s.Logger.Info("Starting rescue-api endpoint")
		if err := s.a.Init(); err != nil {
			s.errs <- err
		}
	}()

	<-s.ctx.Done()

	// Create a context for things that require one for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Stop the api used by rescue-api
	s.a.Deinit()
	s.Logger.Info("Stopped API")

	// Stop the proxy
	s.r.Stop(ctx)
	s.Logger.Info("Stopped router")

	// Shut down metrics server
	s.admin.Close()
	s.Logger.Info("Stopped internal API")

	// Disconnect from the execution client as soon as feasible after shutting down http
	// handlers so that we can serialize the cache
	s.el.Stop()
	s.Logger.Info("Stopped executionlayer")

	// Disconnect from the consensus client
	s.cl.Deinit()
	s.Logger.Info("Stopped consensuslayer")

	close(s.errs)
}

func (s *Service) Stop() error {
	var out error

	s.cancel()

	for err := range s.errs {
		if err == http.ErrServerClosed {
			// This error is expected
			continue
		}
		out = errors.Join(out, err)
	}

	return out
}
