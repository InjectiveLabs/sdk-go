package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcdc "github.com/cosmos/cosmos-sdk/x/authz/codec"
)

// RegisterLegacyAminoCodec registers the necessary x/insurance interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateInsuranceFund{}, "insurance/MsgCreateInsuranceFund", nil)
	cdc.RegisterConcrete(&MsgUnderwrite{}, "insurance/MsgUnderwrite", nil)
	cdc.RegisterConcrete(&MsgRequestRedemption{}, "insurance/MsgRequestRedemption", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "insurance/MsgUpdateParams", nil)
	cdc.RegisterConcrete(&Params{}, "insurance/Params", nil)

}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateInsuranceFund{},
		&MsgUnderwrite{},
		&MsgRequestRedemption{},
		&MsgUpdateParams{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	// ModuleCdc references the global x/insurance module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/insurance and
	// defined at the application level.
	ModuleCdc = codec.NewLegacyAmino()
)

func init() {
	RegisterLegacyAminoCodec(ModuleCdc)
	cryptocodec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()

	RegisterLegacyAminoCodec(authzcdc.Amino)
}
