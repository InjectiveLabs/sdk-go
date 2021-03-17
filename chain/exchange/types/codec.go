package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

var (
	// ModuleCdc references the global exchange module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding.
	//
	// The actual codec used for serialization should be provided to modules/exchange and
	// defined at the application level.
	ModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())
)

// RegisterInterfaces registers concrete types on the Amino codec
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgDeposit{},
		&MsgWithdraw{},
		&MsgInstantSpotMarketLaunch{},
		&MsgCreateSpotLimitOrder{},
		&MsgCreateSpotMarketOrder{},
		&MsgCancelSpotOrder{},
		&MsgCreateDerivativeLimitOrder{},
		&MsgCreateDerivativeMarketOrder{},
		&MsgCancelDerivativeOrder{},
		&MsgSubaccountTransfer{},
		&MsgExternalTransfer{},
	)

	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&SpotMarketParamUpdateProposal{},
		&SpotMarketLaunchProposal{},
		&SpotMarketStatusSetProposal{},
		&PerpetualMarketLaunchProposal{},
		&ExpiryFuturesMarketLaunchProposal{},
		&DerivativeMarketParamUpdateProposal{},
		&DerivativeMarketStatusSetProposal{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
