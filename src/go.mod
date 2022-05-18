module github.com/SimplyVC/oasis_api_server/src

go 1.17

replace (
	github.com/coreos/etcd => github.com/coreos/etcd v3.3.25+incompatible
	github.com/google/flatbuffers => github.com/google/flatbuffers v1.12.1
	github.com/tendermint/tendermint => github.com/oasisprotocol/tendermint v0.34.15-oasis1
	github.com/whyrusleeping/timecache => github.com/oasisprotocol/timecache v0.0.0-20220102191729-558b1c931038
	golang.org/x/crypto/curve25519 => github.com/oasisprotocol/curve25519-voi/primitives/x25519 v0.0.0-20210505121811-294cf0fbfb43
	golang.org/x/crypto/ed25519 => github.com/oasisprotocol/curve25519-voi/primitives/ed25519 v0.0.0-20210505121811-294cf0fbfb43
)

require (
	github.com/claudetech/ini v0.0.0-20140910072410-73e6100d9d51
	github.com/gorilla/mux v1.7.4
	github.com/mackerelio/go-osstat v0.1.0
<<<<<<< HEAD
	github.com/oasisprotocol/oasis-core/go v0.2201.3
	github.com/prometheus/common v0.32.1
	github.com/tendermint/tendermint v0.34.15
	github.com/zenazn/goji v0.9.0
	google.golang.org/grpc v1.45.0
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/btcsuite/btcd v0.22.0-beta // indirect
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce // indirect
	github.com/cenkalti/backoff/v4 v4.1.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/eapache/channels v1.1.0 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/fxamacker/cbor/v2 v2.4.0 // indirect
	github.com/go-kit/kit v0.12.0 // indirect
	github.com/go-kit/log v0.2.0 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gtank/merlin v0.1.1 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/ipfs/go-log/v2 v2.5.0 // indirect
	github.com/libp2p/go-buffer-pool v0.0.2 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/mimoo/StrobeGo v0.0.0-20181016162300-f8f6d4d2b643 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/oasisprotocol/curve25519-voi v0.0.0-20211219162838-e9a669f65da9 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/petermattis/goid v0.0.0-20180202154549-b0b1615b78e5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.12.1 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/sasha-s/go-deadlock v0.2.1-0.20190427202633-1595213edefa // indirect
	github.com/spf13/afero v1.6.0 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.10.1 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.19.1 // indirect
	golang.org/x/crypto v0.0.0-20210915214749-c084706c2272 // indirect
	golang.org/x/net v0.0.0-20211005001312-d4b1ae081e3b // indirect
	golang.org/x/sys v0.0.0-20220227234510-4e6760a101f9 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa // indirect
	google.golang.org/grpc/security/advancedtls v0.0.0-20200902210233-8630cac324bf // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/ini.v1 v1.66.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
=======
	github.com/oasisprotocol/oasis-core/go v0.2103.7
	github.com/prometheus/common v0.31.0
	github.com/tendermint/tendermint v0.34.9
	github.com/zenazn/goji v0.9.0
	google.golang.org/grpc v1.41.0
>>>>>>> ad2c5a67840ea7a8f0b5c321264df108c698ab91
)
