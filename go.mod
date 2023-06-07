module github.com/Rocket-Pool-Rescue-Node/rescue-proxy

go 1.19

require (
	github.com/Rocket-Pool-Rescue-Node/credentials v0.0.0-20221210220221-3e4b9363005f
	github.com/Rocket-Rescue-Node/guarded-beacon-proxy v0.0.0-20230607151709-4da8427484e7
	github.com/allegro/bigcache/v3 v3.1.0
	github.com/attestantio/go-eth2-client v0.14.5
	github.com/ethereum/go-ethereum v1.12.0
	github.com/gorilla/mux v1.8.0
	github.com/mattn/go-sqlite3 v1.11.0
	github.com/mwitkow/grpc-proxy v0.0.0-20220126150247-db34e7bfee32
	github.com/prometheus/client_golang v1.14.0
	github.com/prysmaticlabs/prysm/v3 v3.2.0
	github.com/rocket-pool/rocketpool-go v1.6.0
	github.com/rocket-pool/smartnode v1.9.5
	github.com/rs/zerolog v1.26.1
	go.uber.org/zap v1.24.0
	google.golang.org/grpc v1.52.3
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/alessio/shellescape v1.4.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.3.2 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/deckarep/golang-set/v2 v2.3.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.2.0 // indirect
	github.com/fatih/color v1.14.1 // indirect
	github.com/ferranbt/fastssz v0.1.3 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/goccy/go-yaml v1.9.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.0.1 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/holiman/uint256 v1.2.2 // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/minio/sha256-simd v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pbnjay/memory v0.0.0-20210728143218-7b4eea64cf58 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.39.0 // indirect
	github.com/prometheus/procfs v0.9.0 // indirect
	github.com/prysmaticlabs/fastssz v0.0.0-20221107182844-78142813af44 // indirect
	github.com/prysmaticlabs/go-bitfield v0.0.0-20210809151128-385d8c5e3fb7 // indirect
	github.com/prysmaticlabs/gohashtree v0.0.2-alpha // indirect
	github.com/r3labs/sse/v2 v2.7.4 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/thomaso-mirodin/intmath v0.0.0-20160323211736-5dc6d854e46e // indirect
	github.com/tklauser/go-sysconf v0.3.11 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/urfave/cli v1.22.12 // indirect
	github.com/yusufpapurcu/wmi v1.2.3 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	golang.org/x/crypto v0.9.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sync v0.2.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	gonum.org/v1/gonum v0.13.0 // indirect
	google.golang.org/genproto v0.0.0-20221118155620-16455021b5e6 // indirect
	gopkg.in/cenkalti/backoff.v1 v1.1.0 // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/mwitkow/grpc-proxy => github.com/jshufro/grpc-proxy v0.0.2

replace github.com/Rocket-Rescue-Node/guarded-beacon-proxy => ../guarded-beacon-proxy
