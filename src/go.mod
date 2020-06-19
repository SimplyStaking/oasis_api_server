module github.com/SimplyVC/oasis_api_server/src

go 1.13

replace (
	github.com/tendermint/tendermint => github.com/oasisprotocol/tendermint v0.33.4-oasis2
	golang.org/x/crypto/curve25519 => github.com/oasisprotocol/ed25519/extra/x25519 v0.0.0-20200528083105-55566edd6df0
	golang.org/x/crypto/ed25519 => github.com/oasisprotocol/ed25519 v0.0.0-20200528083105-55566edd6df0
)

require (
	github.com/blevesearch/bleve v1.0.9
	github.com/cenkalti/backoff/v4 v4.0.0
	github.com/claudetech/ini v0.0.0-20140910072410-73e6100d9d51
	github.com/dgraph-io/badger/v2 v2.0.3
	github.com/eapache/channels v1.1.0
	github.com/fxamacker/cbor/v2 v2.2.0
	github.com/go-kit/kit v0.10.0
	github.com/golang/protobuf v1.4.0
	github.com/golang/snappy v0.0.1
	github.com/gorilla/mux v1.7.4
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.1-0.20190118093823-f849b5445de4
	github.com/hpcloud/tail v1.0.0
	github.com/kainlite/grpc-ping v0.0.0-20190623201649-d8f897f70424
	github.com/libp2p/go-libp2p v0.9.1
	github.com/libp2p/go-libp2p-core v0.5.6
	github.com/mackerelio/go-osstat v0.1.0
	github.com/multiformats/go-multiaddr v0.2.2
	github.com/multiformats/go-multiaddr-net v0.1.5
	github.com/oasisprotocol/deoxysii v0.0.0-20200527154044-851aec403956
	github.com/oasisprotocol/ed25519 v0.0.0-20200528083105-55566edd6df0
	github.com/oasisprotocol/oasis-core v0.0.0-20200616160851-aa117564d654 // indirect
	github.com/oasisprotocol/oasis-core/go v0.0.0-20200618144736-02a945839e9b
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.5.1
	github.com/prometheus/common v0.9.1
	github.com/seccomp/libseccomp-golang v0.9.1
	github.com/sparrc/go-ping v0.0.0-20190613174326-4e5b6552494c
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.0
	github.com/stretchr/testify v1.5.1
	github.com/tendermint/iavl v0.12.2
	github.com/tendermint/tendermint v0.33.4
	github.com/tendermint/tm-db v0.5.1
	github.com/thepudds/fzgo v0.2.2
	github.com/uber/jaeger-client-go v2.16.0+incompatible
	github.com/whyrusleeping/go-logging v0.0.1
	github.com/whyrusleeping/go-notifier v0.0.0-20170827234753-097c5d47330f // indirect
	github.com/zenazn/goji v0.9.0
	github.com/zondax/ledger-oasis-go v0.3.0
	gitlab.com/yawning/dynlib.git v0.0.0-20200603163025-35fe007b0761
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37
	golang.org/x/net v0.0.0-20200520004742-59133d7f0dd7
	google.golang.org/genproto v0.0.0-20200313141609-30c55424f95d
	google.golang.org/grpc v1.29.1
)
