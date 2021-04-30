module github.com/SimplyVC/oasis_api_server/src

go 1.15

replace (
	github.com/tendermint/tendermint => github.com/oasisprotocol/tendermint v0.34.9-oasis2
	golang.org/x/crypto/curve25519 => github.com/oasisprotocol/ed25519/extra/x25519 v0.0.0-20210127160119-f7017427c1ea
	golang.org/x/crypto/ed25519 => github.com/oasisprotocol/ed25519 v0.0.0-20210127160119-f7017427c1ea
)

require (
	github.com/Kubuxu/go-os-helper v0.0.1 // indirect
	github.com/blevesearch/bleve v1.0.14
	github.com/cenkalti/backoff/v4 v4.1.0
	github.com/claudetech/ini v0.0.0-20140910072410-73e6100d9d51
	github.com/davidlazar/go-crypto v0.0.0-20190912175916-7055855a373f // indirect
	github.com/dgraph-io/badger/v2 v2.2007.2
	github.com/eapache/channels v1.1.0
	github.com/fxamacker/cbor/v2 v2.2.1-0.20200820021930-bafca87fa6db
	github.com/go-check/check v0.0.0-20180628173108-788fd7840127 // indirect
	github.com/go-kit/kit v0.10.0
	github.com/golang/snappy v0.0.3
	github.com/gorilla/mux v1.7.4
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/hpcloud/tail v1.0.0
	github.com/jackpal/gateway v1.0.5 // indirect
	github.com/kainlite/grpc-ping v0.0.0-20190623201649-d8f897f70424
	github.com/libp2p/go-libp2p v0.13.0
	github.com/libp2p/go-libp2p-core v0.8.5
	github.com/libp2p/go-libp2p-peer v0.2.0 // indirect
	github.com/libp2p/go-sockaddr v0.1.0 // indirect
	github.com/libp2p/go-stream-muxer v0.0.1 // indirect
	github.com/mackerelio/go-osstat v0.1.0
	github.com/multiformats/go-multiaddr v0.3.1
	github.com/multiformats/go-multiaddr-net v0.2.0
	github.com/oasisprotocol/deoxysii v0.0.0-20200527154044-851aec403956
	github.com/oasisprotocol/ed25519 v0.0.0-20210127160119-f7017427c1ea
	github.com/oasisprotocol/oasis-core/go v0.2101.1
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.10.0
	github.com/prometheus/common v0.20.0
	github.com/seccomp/libseccomp-golang v0.9.1
	github.com/sparrc/go-ping v0.0.0-20190613174326-4e5b6552494c
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/iavl v0.12.2
	github.com/tendermint/tendermint v0.34.9
	github.com/tendermint/tm-db v0.6.4
	github.com/thepudds/fzgo v0.2.2
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/whyrusleeping/go-logging v0.0.1
	github.com/whyrusleeping/go-notifier v0.0.0-20170827234753-097c5d47330f // indirect
	github.com/whyrusleeping/mafmt v1.2.8 // indirect
	github.com/zenazn/goji v0.9.0
	github.com/zondax/ledger-oasis-go v0.3.0
	gitlab.com/yawning/dynlib.git v0.0.0-20200603163025-35fe007b0761
	golang.org/x/crypto v0.0.0-20201117144127-c1f2f97bffc9
	golang.org/x/net v0.0.0-20201021035429-f5854403a974
	google.golang.org/genproto v0.0.0-20201119123407-9b1e624d6bc4
	google.golang.org/grpc v1.36.1
	google.golang.org/protobuf v1.26.0
)
