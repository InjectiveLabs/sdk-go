package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcdc "github.com/cosmos/cosmos-sdk/x/authz/codec"
)

var (
	AminoCodec = codec.NewLegacyAmino()
)

func init() {
	RegisterLegacyAminoCodec(AminoCodec)
	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	sdk.RegisterLegacyAminoCodec(AminoCodec)
	RegisterLegacyAminoCodec(authzcdc.Amino)

	AminoCodec.Seal()
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateRateLimit{}, "ibc-rate-limits/MsgCreateRateLimit", nil)
	cdc.RegisterConcrete(&MsgUpdateRateLimit{}, "ibc-rate-limits/MsgUpdateRateLimit", nil)
	cdc.RegisterConcrete(&MsgRemoveRateLimit{}, "ibc-rate-limits/MsgRemoveRateLimit", nil)
}

func RegisterInterfaces(reg codectypes.InterfaceRegistry) {
	reg.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateRateLimit{},
		&MsgUpdateRateLimit{},
		&MsgRemoveRateLimit{},
	)

	msgservice.RegisterMsgServiceDesc(reg, &_Msg_serviceDesc)
}
