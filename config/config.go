package config

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

type CredentialSecrets [][]byte

func (c *CredentialSecrets) String() string {
	if c == nil {
		return ""
	}

	out := make([]string, 0, len(*c))
	for _, bytes := range *c {
		out = append(out, string(bytes))
	}

	return strings.Join(out, ",")
}

func (c *CredentialSecrets) Set(arg string) error {
	*c = append(*c, []byte(arg))
	return nil
}

type Config struct {
	BeaconURL            *url.URL
	ExecutionURL         *url.URL
	ListenAddr           string
	APIListenAddr        string
	AdminListenAddr      string
	GRPCListenAddr       string
	GRPCBeaconAddr       string
	GRPCTLSCertFile      string
	GRPCTLSKeyFile       string
	RocketStorageAddr    string
	CredentialSecrets    CredentialSecrets
	CachePath            string
	EnableSoloValidators bool
	Debug                bool
	ForceBNJSON          bool
}

func InitFlags() *Config {
	config := new(Config)

	credentialSecrets := make(CredentialSecrets, 0)
	flag.Var(&credentialSecrets, "hmac-secret", "The secret to use for HMAC. Defaults to 'test-secret'")

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
	cachePathFlag := flag.String("cache-path", "", "A path to cache EL data in. Leave blank to disable caching.")
	enableSoloValidatorsFlag := flag.Bool("enable-solo-validators", true, "Whether or not to allow solo validators access.")
	forceBNJSONFlag := flag.Bool("force-bn-json", false, "Disables SSZ in the BN.")

	flag.Parse()

	if len(credentialSecrets) == 0 {
		credentialSecrets = CredentialSecrets{[]byte("test-secret")}
	}

	if *bnURLFlag == "" {
		fmt.Fprintf(os.Stderr, "Invalid -bn-url:\n")
		flag.PrintDefaults()
		os.Exit(1)
		return nil
	}

	if *ecURLFlag == "" {
		fmt.Fprintf(os.Stderr, "Invalid -ec-url:\n")
		flag.PrintDefaults()
		os.Exit(1)
		return nil
	}

	base, err := url.Parse(*bnURLFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid -bn-url: %s\n %v\n", *bnURLFlag, err)
		os.Exit(1)
		return nil
	}
	config.BeaconURL = base

	base, err = url.Parse(*ecURLFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid -ec-url: %s\n %v\n", *ecURLFlag, err)
		os.Exit(1)
		return nil
	}
	config.ExecutionURL = base

	if config.BeaconURL.Scheme != "http" && config.BeaconURL.Scheme != "https" {
		fmt.Fprintf(os.Stderr, "Invalid -bn-url: %s\nOnly http and https Beacon Nodes are supported right now.\n", *bnURLFlag)
		os.Exit(1)
		return nil
	}

	// We must use websockets to subscribe to events
	if config.ExecutionURL.Scheme != "ws" {
		fmt.Fprintf(os.Stderr, "Invalid -ec-url: %s\nOnly ws Execution Clients are supported right now.\n", *ecURLFlag)
		os.Exit(1)
		return nil
	}

	if *addrURLFlag == "" {
		fmt.Fprintf(os.Stderr, "Invalid -addr:\n")
		os.Exit(1)
		return nil
	}

	if *adminAddrURLFlag == "" {
		fmt.Fprintf(os.Stderr, "Invalid -admin-addr:\n")
		os.Exit(1)
		return nil
	}

	if *apiAddrURLFlag == "" {
		fmt.Fprintf(os.Stderr, "Invalid -api-addr:\n")
		os.Exit(1)
		return nil
	}

	config.GRPCTLSCertFile = *grpcTLSCertFileFlag
	config.GRPCTLSKeyFile = *grpcTLSKeyFileFlag
	if (config.GRPCTLSCertFile == "" && config.GRPCTLSKeyFile != "") ||
		(config.GRPCTLSCertFile != "" && config.GRPCTLSKeyFile == "") {

		fmt.Fprintf(os.Stderr, "If either --grpc-tls-key-file or --grpc-tls-cert-file is set, both must be set")
		os.Exit(1)
		return nil
	}

	config.AdminListenAddr = *adminAddrURLFlag
	config.APIListenAddr = *apiAddrURLFlag
	config.CredentialSecrets = credentialSecrets
	config.CachePath = *cachePathFlag
	config.GRPCListenAddr = *grpcAddrFlag
	config.GRPCBeaconAddr = *grpcBeaconAddrFlag
	config.ListenAddr = *addrURLFlag
	config.RocketStorageAddr = *rocketStorageAddrFlag
	config.EnableSoloValidators = *enableSoloValidatorsFlag
	config.Debug = *debug
	config.ForceBNJSON = *forceBNJSONFlag
	return config
}
