package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

var (
	// ModuleCdc references the global erc20bridge module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding.
	//
	// The actual codec used for serialization should be provided to modules/erc20bridge and
	// defined at the application level.
	ModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())
)

// RegisterLegacyAminoCodec registers concrete types on the LegacyAmino codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(TokenMapping{}, "cosmos-sdk/TokenMapping", nil)
	cdc.RegisterConcrete(&RegisterTokenMappingProposal{}, "cosmos-sdk/RegisterTokenMappingProposal", nil)
}

// RegisterInterfaces register implementations
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgERC20BridgeMint{},
		&MsgInitHub{},
	)
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&RegisterTokenMappingProposal{},
	)
}
