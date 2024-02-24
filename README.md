[![Rescue Node Proxy](https://github.com/Rocket-Rescue-Node/rescue-proxy/actions/workflows/tests.yml/badge.svg)](https://github.com/Rocket-Rescue-Node/rescue-proxy/actions/workflows/tests.yml) [![golangci-lint](https://github.com/Rocket-Rescue-Node/rescue-proxy/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/Rocket-Rescue-Node/rescue-proxy/actions/workflows/golangci-lint.yml) [![GoReportCard](https://goreportcard.com/badge/github.com/Rocket-Rescue-Node/rescue-proxy)](https://goreportcard.com/report/github.com/Rocket-Rescue-Node/rescue-proxy)

# Rescue-Proxy

[Rocket Rescue Node](https://rescuenode.com)'s Rescue-Proxy is a custom reverse proxy meant to sit between a shared beacon node and its downstream users. It behaves like a normal reverse proxy with the following added features and protections:

1. HMAC authentication via HTTP Basic Auth / GRPC headers
1. Fee Recipient validation for Rocket Pool validator clients
1. Credential expiration
1. Robust caching for frequently accessed immutable and mutable chain data
1. GRPC support for Prysm; HTTP support for Nimbus, Lighthouse, Lodestar, and Teku
1. Fee Recipient validation for Solo staker validator clients
1. Detailed logging and [metrics](https://status.rescuenode.com).
1. A gRPC API allowing [rescue-api](https://github.com/Rocket-Rescue-Node/rescue-api) to see views of currently active solo and Rocket Pool node operators

## Usage

```
Usage of ./rescue-proxy:
  -addr string
        Address on which to reply to HTTP requests (default "0.0.0.0:80")
  -admin-addr string
        Address on which to reply to admin/metrics requests (default "0.0.0.0:8000")
  -api-addr string
        Address on which to reply to gRPC API requests (default "0.0.0.0:8080")
  -bn-url string
        URL to the beacon node to proxy, eg, http://localhost:5052
  -cache-path string
        A path to cache EL data in. Leave blank to disable caching.
  -debug
        Whether to enable verbose logging
  -ec-url string
        URL to the execution client to use, eg, http://localhost:8545
  -enable-solo-validators
        Whether or not to allow solo validators access. (default true)
  -grpc-addr string
        Address on which to reply to gRPC requests
  -grpc-beacon-addr string
        Address to the beacon node to proxy for gRPC, eg, localhost:4000
  -grpc-tls-cert-file string
        Optional TLS Certificate for the gRPC host
  -grpc-tls-key-file string
        Optional TLS Key for the gRPC host
  -hmac-secret string
        The secret to use for HMAC (default "test-secret")
        Can be passed multiple times. Credentials are considered valid if they were generated with any supplied secret.
  -rocketstorage-addr string
        Address of the Rocket Storage contract. Defaults to mainnet (default "0x1d8f8f00cfa6758d7bE78336684788Fb0ee0Fa46")
```

  * The `-grpc` flags should only be used with a Prysm beacon node.
    * The user must pass `--grpc-headers=rprnauth=USERNAME:PASSWORD` in this case.
  * The user should use Basic Auth for access (e.g. Beacon Node URL `http://USERNAME:PASSWORD@0.0.0.0:80`)
  * `-hmac-secret` must match the one used with the [Credentials](https://github.com/Rocket-Rescue-Node/credentials) library that generated the username, password

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[AGPL](https://www.gnu.org/licenses/agpl-3.0.en.html)  
Copyright (C) 2022 Jacob Shufro and João Poupino

