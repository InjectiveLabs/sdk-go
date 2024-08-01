package types

import (
	"cosmossdk.io/x/tx/signing"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/cosmos/gogoproto/proto"

	injcodec "github.com/InjectiveLabs/sdk-go/chain/codec"
)

type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Codec             codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

// MakeEncodingConfig creates an EncodingConfig for testing
func MakeEncodingConfig() EncodingConfig {
	cdc := codec.NewLegacyAmino()
	interfaceRegistry := NewInterfaceRegistry()
	appCodec := codec.NewProtoCodec(interfaceRegistry)
	encodingConfig := EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Codec:             appCodec,
		TxConfig:          tx.NewTxConfig(appCodec, tx.DefaultSignModes),
		Amino:             cdc,
	}

	injcodec.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	injcodec.RegisterLegacyAminoCodec(encodingConfig.Amino)
	return encodingConfig
}

// NewInterfaceRegistry returns a new InterfaceRegistry
func NewInterfaceRegistry() types.InterfaceRegistry {
	registry, err := types.NewInterfaceRegistryWithOptions(types.InterfaceRegistryOptions{
		ProtoFiles: proto.HybridResolver,
		SigningOptions: signing.Options{
			AddressCodec: address.Bech32Codec{
				Bech32Prefix: sdktypes.GetConfig().GetBech32AccountAddrPrefix(),
			},
			ValidatorAddressCodec: address.Bech32Codec{
				Bech32Prefix: sdktypes.GetConfig().GetBech32ValidatorAddrPrefix(),
			},
		},
	})
	if err != nil {
		panic(err)
	}
	return registry
}
