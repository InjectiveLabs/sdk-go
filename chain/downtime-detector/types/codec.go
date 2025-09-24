package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authzcdc "github.com/cosmos/cosmos-sdk/x/authz/codec"
)

func RegisterLegacyAminoCodec(*codec.LegacyAmino) {}

func RegisterInterfaces(types.InterfaceRegistry) {}

var (
	ModuleCdc = codec.NewLegacyAmino()
)

func init() {
	RegisterLegacyAminoCodec(ModuleCdc)
	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	sdk.RegisterLegacyAminoCodec(ModuleCdc)
	RegisterLegacyAminoCodec(authzcdc.Amino)

	ModuleCdc.Seal()
}
