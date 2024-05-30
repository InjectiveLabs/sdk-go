package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authzcdc "github.com/cosmos/cosmos-sdk/x/authz/codec"
	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

const ModuleName = "permissions"

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgUpdateParams{}, "permissions/MsgUpdateParams", nil)
	cdc.RegisterConcrete(&MsgCreateNamespace{}, "permissions/MsgCreateNamespace", nil)
	cdc.RegisterConcrete(&MsgDeleteNamespace{}, "permissions/MsgDeleteNamespace", nil)
	cdc.RegisterConcrete(&MsgUpdateNamespace{}, "permissions/MsgUpdateNamespace", nil)
	cdc.RegisterConcrete(&MsgUpdateNamespaceRoles{}, "permissions/MsgUpdateNamespaceRoles", nil)
	cdc.RegisterConcrete(&MsgRevokeNamespaceRoles{}, "permissions/MsgRevokeNamespaceRoles", nil)
	cdc.RegisterConcrete(&MsgClaimVoucher{}, "permissions/MsgClaimVoucher", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
		&MsgCreateNamespace{},
		&MsgDeleteNamespace{},
		&MsgUpdateNamespace{},
		&MsgUpdateNamespaceRoles{},
		&MsgRevokeNamespaceRoles{},
		&MsgClaimVoucher{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	ModuleCdc = codec.NewLegacyAmino()
)

func init() {
	RegisterCodec(ModuleCdc)
	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	sdk.RegisterLegacyAminoCodec(ModuleCdc)
	RegisterCodec(authzcdc.Amino)
	ModuleCdc.Seal()
}
