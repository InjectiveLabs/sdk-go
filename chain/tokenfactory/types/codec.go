package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authzcdc "github.com/cosmos/cosmos-sdk/x/authz/codec"
	govcdc "github.com/cosmos/cosmos-sdk/x/gov/codec"
	groupcdc "github.com/cosmos/cosmos-sdk/x/group/codec"

	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateDenom{}, "injective/tokenfactory/create-denom", nil)
	cdc.RegisterConcrete(&MsgMint{}, "injective/tokenfactory/mint", nil)
	cdc.RegisterConcrete(&MsgBurn{}, "injective/tokenfactory/burn", nil)
	// nolint:all
	// cdc.RegisterConcrete(&MsgForceTransfer{}, "injective/tokenfactory/force-transfer", nil)
	cdc.RegisterConcrete(&MsgChangeAdmin{}, "injective/tokenfactory/change-admin", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "injective/tokenfactory/update-params", nil)
	cdc.RegisterConcrete(&MsgSetDenomMetadata{}, "injective/tokenfactory/set-denom-metadata", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDenom{},
		&MsgMint{},
		&MsgBurn{},
		// &MsgForceTransfer{},
		&MsgChangeAdmin{},
		&MsgUpdateParams{},
		&MsgSetDenomMetadata{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterCodec(amino)
	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	sdk.RegisterLegacyAminoCodec(amino)
	RegisterCodec(govcdc.Amino)
	RegisterCodec(authzcdc.Amino)
	RegisterCodec(groupcdc.Amino)

	amino.Seal()
}
