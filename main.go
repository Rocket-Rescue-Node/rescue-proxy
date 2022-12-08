package main

import (
	"context"
	"crypto/sha256"
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/Rocket-Pool-Rescue-Node/credentials"
	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/api"
	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/consensuslayer"
	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/executionlayer"
	"github.com/Rocket-Pool-Rescue-Node/rescue-proxy/router"
	"go.uber.org/zap"
)

var logger *zap.Logger

type config struct {
	BeaconURL          *url.URL
	ExecutionURL       *url.URL
	ListenAddr         string
	APIListenAddr      string
	RocketStorageAddr  string
	CredentialSecret   string
	AuthValidityWindow time.Duration
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
	apiAddrURLFlag := flag.String("api-addr", "0.0.0.0:8080", "Address on which to reply to gRPC API requests")
	rocketStorageAddrFlag := flag.String("rocketstorage-addr", "0x1d8f8f00cfa6758d7bE78336684788Fb0ee0Fa46", "Address of the Rocket Storage contract. Defaults to mainnet")
	debug := flag.Bool("debug", false, "Whether to enable verbose logging")
	credentialSecretFlag := flag.String("hmac-secret", "test-secret", "The secret to use for HMAC")
	authValidityWindowFlag := flag.String("auth-valid-for", "360h", "The duration after which a credential should be considered invalid, eg, 360h for 15 days")

	flag.Parse()

	if err := initLogger(*debug); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing logger: %v\n", err)
		os.Exit(1)
		return
	}
	defer logger.Sync()

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

	// TODO: Support grpc:// as a protocol type
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

	config.ListenAddr = *addrURLFlag
	config.APIListenAddr = *apiAddrURLFlag
	config.RocketStorageAddr = *rocketStorageAddrFlag
	config.CredentialSecret = *credentialSecretFlag
	return
}

func blockUntilSIGINT() {

	// Trap SIGINT
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until SIGINT is received
	<-c

	// Allow subsequent SIGINT to quickly shut down by removing the trap
	signal.Reset()
	close(c)
}

func main() {

	// Initialize config
	config := initFlags()
	logger.Info("Starting up the rescue node proxy...")

	// Listen on the provided address
	listener, err := net.Listen("tcp", config.ListenAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to listen on provided address %s\n%v\n", config.ListenAddr, err)
		os.Exit(1)
		return
	}

	// Connect to and initialize the execution layer
	el := executionlayer.NewExecutionLayer(config.ExecutionURL, config.RocketStorageAddr, logger)

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

	// Spin up the server on a different goroutine, since it blocks.
	var serverWaitGroup sync.WaitGroup
	serverWaitGroup.Add(1)
	server := http.Server{}
	go func() {
		router := &router.ProxyRouter{
			EL:                 el,
			CL:                 cl,
			CM:                 cm,
			Logger:             logger,
			AuthValidityWindow: config.AuthValidityWindow,
		}
		router.Init(config.BeaconURL)
		logger.Info("Starting http server", zap.String("url", config.ListenAddr))
		if err := server.Serve(listener); err != nil {
			logger.Info("Server stopped", zap.Error(err))
		}
		serverWaitGroup.Done()
	}()

	api := api.NewAPI(config.APIListenAddr, el, logger)
	if err := api.Init(); err != nil {
		logger.Error("Unable to start grpc server", zap.Error(err))
		os.Exit(1)
		return
	}

	blockUntilSIGINT()

	// Shut down gracefully
	logger.Debug("Received SIGINT, shutting down")
	server.Shutdown(context.Background())
	listener.Close()

	api.Deinit()

	// Wait for the listener/server to exit
	serverWaitGroup.Wait()

	// Disconnect from the execution client
	el.Deinit()
	cl.Deinit()
}
