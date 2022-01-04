module github.com/SimplyVC/oasis_api_server/src

go 1.15

replace (
	github.com/tendermint/tendermint => github.com/oasisprotocol/tendermint v0.34.9-oasis2
	golang.org/x/crypto/curve25519 => github.com/oasisprotocol/ed25519/extra/x25519 v0.0.0-20210127160119-f7017427c1ea
	golang.org/x/crypto/ed25519 => github.com/oasisprotocol/ed25519 v0.0.0-20210127160119-f7017427c1ea
)

require (
	github.com/claudetech/ini v0.0.0-20140910072410-73e6100d9d51
	github.com/gorilla/mux v1.7.4
	github.com/mackerelio/go-osstat v0.1.0
	github.com/oasisprotocol/oasis-core/go v0.2103.7
	github.com/prometheus/common v0.31.0
	github.com/tendermint/tendermint v0.34.9
	github.com/zenazn/goji v0.9.0
	google.golang.org/grpc v1.41.0
)
