package codec

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"

	cryptocodec "github.com/InjectiveLabs/sdk-go/chain/crypto/codec"
	wasmxtypes "github.com/InjectiveLabs/sdk-go/chain/wasmx/types"
	injective "github.com/InjectiveLabs/sdk-go/chain/types"
)

// RegisterLegacyAminoCodec registers Interfaces from types, crypto, and SDK std.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptocodec.RegisterCrypto(cdc)
	codec.RegisterEvidences(cdc)
	// Historical decoding only: wasmx has no module or service registration.
	wasmxtypes.RegisterLegacyAminoCodec(cdc)
}

// RegisterInterfaces registers Interfaces from types, crypto, and SDK std.
func RegisterInterfaces(interfaceRegistry codectypes.InterfaceRegistry) {
	std.RegisterInterfaces(interfaceRegistry)
	cryptocodec.RegisterInterfaces(interfaceRegistry)
	injective.RegisterInterfaces(interfaceRegistry)
	wasmxtypes.RegisterInterfaces(interfaceRegistry)
}
