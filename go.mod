module github.com/InjectiveLabs/sdk-go

go 1.16

require (
	github.com/InjectiveLabs/suplog v1.3.3
	github.com/allegro/bigcache v1.2.1 // indirect
	github.com/bandprotocol/bandchain-packet v0.0.2
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/btcsuite/btcd v0.22.0-beta
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce
	github.com/bugsnag/panicwrap v1.2.0 // indirect
	github.com/cespare/cp v1.1.1 // indirect
	github.com/cosmos/cosmos-sdk v0.45.2
	github.com/cosmos/ibc-go/v2 v2.0.2
	github.com/ethereum/go-ethereum v1.10.17
	github.com/gballet/go-libpcsclite v0.0.0-20191108122812-4678299bea08 // indirect
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.3
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/tsdb v0.10.0 // indirect
	github.com/regen-network/cosmos-proto v0.3.1
	github.com/rjeczalik/notify v0.9.2 // indirect
	github.com/shopspring/decimal v1.2.0
	github.com/spf13/cobra v1.3.0
	github.com/status-im/keycard-go v0.0.0-20200402102358-957c09536969 // indirect
	github.com/tendermint/tendermint v0.34.16
	github.com/tyler-smith/go-bip39 v1.0.2
	golang.org/x/crypto v0.0.0-20211215165025-cf75a172585e
	google.golang.org/genproto v0.0.0-20220118154757-00ab72f36ad5
	google.golang.org/grpc v1.44.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/ini.v1 v1.66.4
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/btcsuite/btcutil => github.com/btcsuite/btcutil v1.0.2

replace github.com/cosmos/cosmos-sdk => github.com/InjectiveLabs/cosmos-sdk v0.45.2-inj
