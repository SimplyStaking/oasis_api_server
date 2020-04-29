module github.com/SimplyVC/oasis_api_server/src

go 1.13

replace (
	github.com/tendermint/iavl => github.com/oasislabs/iavl v0.12.0-ekiden3
	github.com/tendermint/tendermint => github.com/oasislabs/tendermint v0.32.10-oasis1
	golang.org/x/crypto/curve25519 => github.com/oasislabs/ed25519/extra/x25519 v0.0.0-20191022155220-a426dcc8ad5f
	golang.org/x/crypto/ed25519 => github.com/oasislabs/ed25519 v0.0.0-20191109133925-b197a691e30d
)

require (
	github.com/blevesearch/bleve v0.8.0
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
	github.com/libp2p/go-libp2p v0.1.1
	github.com/libp2p/go-libp2p-core v0.0.3
	github.com/mackerelio/go-osstat v0.1.0
	github.com/multiformats/go-multiaddr v0.0.4
	github.com/multiformats/go-multiaddr-net v0.0.1
	github.com/oasislabs/deoxysii v0.0.0-20190807103041-6159f99c2236
	github.com/oasislabs/ed25519 v0.0.0-20191122104632-9d9ffc15f526
	github.com/oasislabs/oasis-core/go v0.0.0-20200416125006-75b2bfa01f8d
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.5.1
	github.com/prometheus/common v0.9.1
	github.com/seccomp/libseccomp-golang v0.9.1
	github.com/sparrc/go-ping v0.0.0-20190613174326-4e5b6552494c
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.6.3
	github.com/stretchr/testify v1.5.1
	github.com/tendermint/go-amino v0.15.0 // indirect
	github.com/tendermint/iavl v0.12.2
	github.com/tendermint/tendermint v0.32.8
	github.com/tendermint/tm-db v0.5.1
	github.com/thepudds/fzgo v0.2.2
	github.com/uber/jaeger-client-go v2.16.0+incompatible
	github.com/whyrusleeping/go-logging v0.0.0-20170515211332-0457bb6b88fc
	github.com/zenazn/goji v0.9.0
	github.com/zondax/ledger-oasis-go v0.3.0
	gitlab.com/yawning/dynlib.git v0.0.0-20190911075527-1e6ab3739fd8
	golang.org/x/crypto v0.0.0-20200406173513-056763e48d71
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e
	google.golang.org/genproto v0.0.0-20200313141609-30c55424f95d
	google.golang.org/grpc v1.28.1
)
