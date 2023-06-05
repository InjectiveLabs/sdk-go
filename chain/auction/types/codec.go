package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcdc "github.com/cosmos/cosmos-sdk/x/authz/codec"
)

// RegisterLegacyAminoCodec registers the necessary x/auction interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgBid{}, "auction/MsgBid", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "auction/MsgUpdateParams", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBid{},
		&MsgUpdateParams{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/auction module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/auction and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)

	RegisterLegacyAminoCodec(authzcdc.Amino)
	amino.Seal()
}
