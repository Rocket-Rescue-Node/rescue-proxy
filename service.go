package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/Rocket-Rescue-Node/rescue-proxy/admin"
	"github.com/Rocket-Rescue-Node/rescue-proxy/api"
	"github.com/Rocket-Rescue-Node/rescue-proxy/config"
	"github.com/Rocket-Rescue-Node/rescue-proxy/consensuslayer"
	"github.com/Rocket-Rescue-Node/rescue-proxy/executionlayer"
	"github.com/Rocket-Rescue-Node/rescue-proxy/router"
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
	Config *config.Config

	// Sub-services
	admin *admin.AdminApi
	el    executionlayer.ExecutionLayer
	cl    consensuslayer.ConsensusLayer
	r     *router.ProxyRouter
	a     *api.API

	// error reporting channel
	errs chan error
}

// NewService creates a [Service] from a given [Config].func NewService(config *Config) *Service {
func NewService(config *config.Config) *Service {
	return &Service{
		Config: config,
	}
}

// Run initializes the [Service] without blocking.
// Callers should read from the returned channel to detect errors.
// The provided [context.Context] can be canceled to initiate a graceful shutdown,
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
	// This initializes metrics, so do it first.
	s.admin = new(admin.AdminApi)
	if err := s.admin.Init("rescue_proxy"); err != nil {
		s.errs <- fmt.Errorf("unable to init admin api (metrics): %w", err)
		return
	}
	if listener, err := net.Listen("tcp", s.Config.AdminListenAddr); err != nil {
		s.errs <- fmt.Errorf("unable to init admin api (metrics): %w", err)
		return
	} else {
		go func() {
			s.Logger.Info("Starting admin API", zap.String("addr", s.Config.AdminListenAddr))
			if err := s.admin.Serve(listener); err != nil && err != http.ErrServerClosed {
				s.errs <- fmt.Errorf("unable to init admin api (metrics): %w", err)
			}
		}()
	}

	// Connect to and initialize the execution layer
	el := &executionlayer.CachingExecutionLayer{
		ECURL:                s.Config.ExecutionURL,
		RocketStorageAddr:    s.Config.RocketStorageAddr,
		Logger:               s.Logger,
		SWVaultsRegistryAddr: s.Config.SWVaultsRegistryAddr,
	}
	s.el = el
	// Init() blocks until the cache is warmed up. This is good, we don't want to
	// start accepting http requests on the proxy until we're ready to handle them.
	if err := el.Init(); err != nil {
		s.errs <- fmt.Errorf("unable to init Execution Layer client: %w", err)
		return
	}
	// After Init() we still have to call Start() to subscribe to new blocks
	go func() {
		s.Logger.Info("Starting EL monitor")
		if err := el.Start(); err != nil {
			s.errs <- fmt.Errorf("EL error: %w", err)
		}
	}()

	// Connect to and initialize the consensus layer
	cl := consensuslayer.NewCachingConsensusLayer(s.Config.BeaconURL, s.Logger, s.Config.ForceBNJSON)
	s.cl = cl
	s.Logger.Info("Starting CL monitor")
	// Consensus Layer is non-blocking/synchronous only.
	// On Init() it will create the client and warm the validator key cache, which
	// is needed to serve responses to rescue-api
	if err := cl.Init(s.ctx); err != nil {
		// Optimization: serialize the EL cache by stopping it now so we can recover
		// faster.
		el.Stop()
		// Only write the error to the channel after so we don't panic while writing
		// the cache to disk
		s.errs <- fmt.Errorf("unable to init Consensus Layer client: %w", err)
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
		EnableSoloValidators: s.Config.EnableSoloValidators,
		CredentialSecrets:    s.Config.CredentialSecrets,
	}
	s.r.Init()
	// Spin up the rest of the servers on different goroutines, since they block.
	go func() {
		s.Logger.Info("Starting http server", zap.String("url", s.Config.ListenAddr))
		if err := s.r.Start(); err != nil {
			s.errs <- fmt.Errorf("unable to start http server: %w", err)
		}
	}()

	s.a = &api.API{
		EL:     s.el,
		CL:     s.cl,
		Logger: s.Logger,
	}
	go func() {
		s.Logger.Info("Starting rescue-api endpoint")

		listener, err := net.Listen("tcp", s.Config.APIListenAddr)
		if err != nil {
			s.errs <- fmt.Errorf("unable to listen on api endpoint: %w", err)
			return
		}

		if err := s.a.Init(listener); err != nil {
			s.errs <- fmt.Errorf("unable to init api: %w", err)
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
	if err := s.admin.Shutdown(ctx); err != nil {
		s.Logger.Info("Error stopping internal API", zap.Error(err))
	}
	s.Logger.Info("Stopped internal API")

	// Disconnect from the execution client as soon as feasible after shutting down http
	// handlers so that we can serialize the cache
	el.Stop()
	s.Logger.Info("Stopped executionlayer")

	// Disconnect from the consensus client
	cl.Deinit()
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
