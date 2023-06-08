package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Rocket-Pool-Rescue-Node/credentials"
	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/admin"
	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/api"
	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/consensuslayer"
	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/executionlayer"
	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/metrics"
	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/router"
	"go.uber.org/zap"
)

var logger *zap.Logger

type config struct {
	BeaconURL          *url.URL
	ExecutionURL       *url.URL
	ListenAddr         string
	APIListenAddr      string
	AdminListenAddr    string
	GRPCListenAddr     string
	GRPCBeaconAddr     string
	GRPCTLSCertFile    string
	GRPCTLSKeyFile     string
	RocketStorageAddr  string
	CredentialSecret   string
	AuthValidityWindow time.Duration
	CachePath          string
}

func initLogger(debug bool) error {
	var cfg zap.Config
	var err error

	if debug {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}

	logger, err = cfg.Build()
	return err
}

func initFlags() (config config) {
	bnURLFlag := flag.String("bn-url", "", "URL to the beacon node to proxy, eg, http://localhost:5052")
	ecURLFlag := flag.String("ec-url", "", "URL to the execution client to use, eg, http://localhost:8545")
	addrURLFlag := flag.String("addr", "0.0.0.0:80", "Address on which to reply to HTTP requests")
	adminAddrURLFlag := flag.String("admin-addr", "0.0.0.0:8000", "Address on which to reply to admin/metrics requests")
	apiAddrURLFlag := flag.String("api-addr", "0.0.0.0:8080", "Address on which to reply to gRPC API requests")
	grpcAddrFlag := flag.String("grpc-addr", "", "Address on which to reply to gRPC requests")
	grpcBeaconAddrFlag := flag.String("grpc-beacon-addr", "", "Address to the beacon node to proxy for gRPC, eg, localhost:4000")
	grpcTLSCertFileFlag := flag.String("grpc-tls-cert-file", "", "Optional TLS Certificate for the gRPC host")
	grpcTLSKeyFileFlag := flag.String("grpc-tls-key-file", "", "Optional TLS Key for the gRPC host")
	rocketStorageAddrFlag := flag.String("rocketstorage-addr", "0x1d8f8f00cfa6758d7bE78336684788Fb0ee0Fa46", "Address of the Rocket Storage contract. Defaults to mainnet")
	debug := flag.Bool("debug", false, "Whether to enable verbose logging")
	credentialSecretFlag := flag.String("hmac-secret", "test-secret", "The secret to use for HMAC")
	authValidityWindowFlag := flag.String("auth-valid-for", "360h", "The duration after which a credential should be considered invalid, eg, 360h for 15 days")
	cachePathFlag := flag.String("cache-path", "", "A path to cache EL data in. Leave blank to disable caching.")

	flag.Parse()

	if err := initLogger(*debug); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing logger: %v\n", err)
		os.Exit(1)
		return
	}

	if *bnURLFlag == "" {
		fmt.Fprintf(os.Stderr, "Invalid -bn-url:\n")
		flag.PrintDefaults()
		os.Exit(1)
		return
	}

	if *ecURLFlag == "" {
		fmt.Fprintf(os.Stderr, "Invalid -ec-url:\n")
		flag.PrintDefaults()
		os.Exit(1)
		return
	}

	base, err := url.Parse(*bnURLFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid -bn-url: %s\n %v\n", *bnURLFlag, err)
		os.Exit(1)
		return
	}
	config.BeaconURL = base

	base, err = url.Parse(*ecURLFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid -ec-url: %s\n %v\n", *ecURLFlag, err)
		os.Exit(1)
		return
	}
	config.ExecutionURL = base

	if config.BeaconURL.Scheme != "http" && config.BeaconURL.Scheme != "https" {
		fmt.Fprintf(os.Stderr, "Invalid -bn-url: %s\nOnly http and https Beacon Nodes are supported right now.\n", *bnURLFlag)
		os.Exit(1)
		return
	}

	// We must use websockets to subscribe to events
	if config.ExecutionURL.Scheme != "ws" {
		fmt.Fprintf(os.Stderr, "Invalid -ec-url: %s\nOnly ws Execution Clients are supported right now.\n", *ecURLFlag)
		os.Exit(1)
		return
	}

	if *addrURLFlag == "" {
		fmt.Fprintf(os.Stderr, "Invalid -addr:\n")
		os.Exit(1)
		return
	}

	if *adminAddrURLFlag == "" {
		fmt.Fprintf(os.Stderr, "Invalid -admin-addr:\n")
		os.Exit(1)
		return
	}

	if *apiAddrURLFlag == "" {
		fmt.Fprintf(os.Stderr, "Invalid -api-addr:\n")
		os.Exit(1)
		return
	}

	if *credentialSecretFlag == "" {
		fmt.Fprintf(os.Stderr, "Invalid -hmac-secret:\n")
		os.Exit(1)
		return
	}

	if *authValidityWindowFlag == "" {
		fmt.Fprintf(os.Stderr, "Invalid -auth-valid-for:\n")
		os.Exit(1)
		return
	}

	config.AuthValidityWindow, err = time.ParseDuration(*authValidityWindowFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid -auth-valid-for:\n%v\n", err)
		os.Exit(1)
		return
	}

	config.GRPCTLSCertFile = *grpcTLSCertFileFlag
	config.GRPCTLSKeyFile = *grpcTLSKeyFileFlag
	if (config.GRPCTLSCertFile == "" && config.GRPCTLSKeyFile != "") ||
		(config.GRPCTLSCertFile != "" && config.GRPCTLSKeyFile == "") {

		fmt.Fprintf(os.Stderr, "If either --grpc-tls-key-file or --grpc-tls-cert-file is set, both must be set")
		os.Exit(1)
		return
	}

	config.AdminListenAddr = *adminAddrURLFlag
	config.APIListenAddr = *apiAddrURLFlag
	config.CredentialSecret = *credentialSecretFlag
	config.CachePath = *cachePathFlag
	config.GRPCListenAddr = *grpcAddrFlag
	config.GRPCBeaconAddr = *grpcBeaconAddrFlag
	config.ListenAddr = *addrURLFlag
	config.RocketStorageAddr = *rocketStorageAddrFlag
	return
}

func waitForSignals(signals ...os.Signal) {

	c := make(chan os.Signal, 1)

	// Always wait for SIGTERM at a minimum
	signal.Notify(c, syscall.SIGTERM)

	if len(signals) != 0 {
		for _, s := range signals {
			if s == syscall.SIGTERM {
				continue
			}
			signal.Notify(c, s)
		}
	}

	// Block until signal is received
	<-c

	// Allow subsequent signals to quickly shut down by removing the trap
	signal.Reset()
	close(c)
}

func main() {

	// Initialize config
	config := initFlags()
	logger.Info("Starting up the rescue node proxy...")

	// Initialize metrics globals
	metricsHTTPHandler, err := metrics.Init("rescue_proxy")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to initialize admin api\n%v\n", err)
		os.Exit(1)
		return
	}

	// Initialize collection of node and validator count metrics
	metrics.InitEpochMetrics()

	// Create the admin-only http server
	adminServer := admin.AdminApi{}
	adminServer.Init(config.AdminListenAddr)

	// Add admin handlers to the admin only http server and start it
	adminServer.Handle("/metrics", metricsHTTPHandler)
	err = adminServer.Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to start admin api\n%v\n", err)
		os.Exit(1)
		return
	}

	// Pick a cache
	var cache executionlayer.Cache
	if config.CachePath == "" {
		cache = &executionlayer.MapsCache{}
	} else {
		cache = &executionlayer.SqliteCache{
			Path: config.CachePath,
		}
	}

	// Connect to and initialize the execution layer
	el := executionlayer.NewExecutionLayer(config.ExecutionURL, config.RocketStorageAddr, cache, logger)

	err = el.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to init Execution Layer client. \n%v\n", err)
		os.Exit(1)
		return
	}

	// Connect to and initialize the consensus layer
	cl := consensuslayer.NewConsensusLayer(config.BeaconURL, logger)

	err = cl.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to init Consensus Layer client. \n%v\n", err)
		os.Exit(1)
		return
	}

	// Create a credential manager
	cm := credentials.NewCredentialManager(sha256.New, []byte(config.CredentialSecret))

	// Initialize the authentication library
	router.InitAuth(cm, config.AuthValidityWindow)

	// Spin up the server on a different goroutine, since it blocks.
	var r *router.ProxyRouter
	go func() {
		r = &router.ProxyRouter{
			EL:                 el,
			CL:                 cl,
			Logger:             logger,
			AuthValidityWindow: config.AuthValidityWindow,
			Addr:               config.ListenAddr,
			GRPCAddr:           config.GRPCListenAddr,
			TLSCertFile:        config.GRPCTLSCertFile,
			TLSKeyFile:         config.GRPCTLSKeyFile,
		}

		if config.GRPCBeaconAddr != "" {
			u, err := url.Parse(config.GRPCBeaconAddr)
			if err != nil {
				logger.Error("Unable to parse grpc beacon address", zap.String("url", config.GRPCBeaconAddr))
				os.Exit(1)
				return
			}

			r.GRPCBeaconURL = u
		}

		logger.Info("Starting http server", zap.String("url", config.ListenAddr))
		err := r.Init(config.BeaconURL)
		if err != nil {
			logger.Error("Unable to start grpc server", zap.Error(err))
			os.Exit(1)
			return
		}
	}()

	api := api.NewAPI(config.APIListenAddr, el, logger)
	if err := api.Init(); err != nil {
		logger.Error("Unable to start grpc server", zap.Error(err))
		os.Exit(1)
		return
	}

	logger.Debug("Trapping SIGTERM and SIGINT")
	waitForSignals(os.Interrupt)

	// Shut down gracefully
	logger.Info("Received signal, shutting down")
	// If gracefully shutting down the http server takes too long,
	// forge ahead without finishing.
	r.Deinit(3 * time.Second)

	api.Deinit()
	logger.Debug("Stopped API")

	// Disconnect from the execution client
	el.Deinit()
	logger.Debug("Stopped executionlayer")
	cl.Deinit()
	logger.Debug("Stopped consensuslayer")

	// Shut down admin server
	adminServer.Close()
	logger.Debug("Stopped internal API")

	logger.Debug("Flushing logs")
	_ = logger.Sync()
}
