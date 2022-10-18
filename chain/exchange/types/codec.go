package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	"github.com/cosmos/cosmos-sdk/x/authz"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// RegisterLegacyAminoCodec registers the necessary x/exchange interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgDeposit{}, "exchange/MsgDeposit", nil)
	cdc.RegisterConcrete(&MsgWithdraw{}, "exchange/MsgWithdraw", nil)
	cdc.RegisterConcrete(&MsgInstantSpotMarketLaunch{}, "exchange/MsgInstantSpotMarketLaunch", nil)
	cdc.RegisterConcrete(&MsgInstantPerpetualMarketLaunch{}, "exchange/MsgInstantPerpetualMarketLaunch", nil)
	cdc.RegisterConcrete(&MsgInstantExpiryFuturesMarketLaunch{}, "exchange/MsgInstantExpiryFuturesMarketLaunch", nil)
	cdc.RegisterConcrete(&MsgCreateSpotLimitOrder{}, "exchange/MsgCreateSpotLimitOrder", nil)
	cdc.RegisterConcrete(&MsgBatchCreateSpotLimitOrders{}, "exchange/MsgBatchCreateSpotLimitOrders", nil)
	cdc.RegisterConcrete(&MsgCreateSpotMarketOrder{}, "exchange/MsgCreateSpotMarketOrder", nil)
	cdc.RegisterConcrete(&MsgCancelSpotOrder{}, "exchange/MsgCancelSpotOrder", nil)
	cdc.RegisterConcrete(&MsgBatchCancelSpotOrders{}, "exchange/MsgBatchCancelSpotOrders", nil)
	cdc.RegisterConcrete(&MsgCreateDerivativeLimitOrder{}, "exchange/MsgCreateDerivativeLimitOrder", nil)
	cdc.RegisterConcrete(&MsgBatchCreateDerivativeLimitOrders{}, "exchange/MsgBatchCreateDerivativeLimitOrders", nil)
	cdc.RegisterConcrete(&MsgCreateDerivativeMarketOrder{}, "exchange/MsgCreateDerivativeMarketOrder", nil)
	cdc.RegisterConcrete(&MsgCancelDerivativeOrder{}, "exchange/MsgCancelDerivativeOrder", nil)
	cdc.RegisterConcrete(&MsgBatchCancelDerivativeOrders{}, "exchange/MsgBatchCancelDerivativeOrders", nil)
	cdc.RegisterConcrete(&MsgBatchCancelBinaryOptionsOrders{}, "exchange/MsgBatchCancelBinaryOptionsOrders", nil)
	cdc.RegisterConcrete(&MsgSubaccountTransfer{}, "exchange/MsgSubaccountTransfer", nil)
	cdc.RegisterConcrete(&MsgExternalTransfer{}, "exchange/MsgExternalTransfer", nil)
	cdc.RegisterConcrete(&MsgIncreasePositionMargin{}, "exchange/MsgIncreasePositionMargin", nil)
	cdc.RegisterConcrete(&MsgLiquidatePosition{}, "exchange/MsgLiquidatePosition", nil)
	cdc.RegisterConcrete(&MsgBatchUpdateOrders{}, "exchange/MsgBatchUpdateOrders", nil)
	cdc.RegisterConcrete(&MsgExec{}, "exchange/MsgExec", nil)
	cdc.RegisterConcrete(&MsgRewardsOptOut{}, "exchange/MsgRewardsOptOut", nil)
	cdc.RegisterConcrete(&MsgInstantBinaryOptionsMarketLaunch{}, "exchange/MsgInstantBinaryOptionsMarketLaunch", nil)
	cdc.RegisterConcrete(&MsgCreateBinaryOptionsLimitOrder{}, "exchange/MsgCreateBinaryOptionsLimitOrder", nil)
	cdc.RegisterConcrete(&MsgCreateBinaryOptionsMarketOrder{}, "exchange/MsgCreateBinaryOptionsMarketOrder", nil)
	cdc.RegisterConcrete(&MsgCancelBinaryOptionsOrder{}, "exchange/MsgCancelBinaryOptionsOrder", nil)
	cdc.RegisterConcrete(&MsgAdminUpdateBinaryOptionsMarket{}, "exchange/MsgAdminUpdateBinaryOptionsMarket", nil)

	cdc.RegisterConcrete(&ExchangeEnableProposal{}, "exchange/ExchangeEnableProposal", nil)
	cdc.RegisterConcrete(&BatchExchangeModificationProposal{}, "exchange/BatchExchangeModificationProposal", nil)
	cdc.RegisterConcrete(&SpotMarketParamUpdateProposal{}, "exchange/SpotMarketParamUpdateProposal", nil)
	cdc.RegisterConcrete(&SpotMarketLaunchProposal{}, "exchange/SpotMarketLaunchProposal", nil)
	cdc.RegisterConcrete(&PerpetualMarketLaunchProposal{}, "exchange/PerpetualMarketLaunchProposal", nil)
	cdc.RegisterConcrete(&ExpiryFuturesMarketLaunchProposal{}, "exchange/ExpiryFuturesMarketLaunchProposal", nil)
	cdc.RegisterConcrete(&DerivativeMarketParamUpdateProposal{}, "exchange/DerivativeMarketParamUpdateProposal", nil)
	cdc.RegisterConcrete(&MarketForcedSettlementProposal{}, "exchange/MarketForcedSettlementProposal", nil)
	cdc.RegisterConcrete(&UpdateDenomDecimalsProposal{}, "exchange/UpdateDenomDecimalsProposal", nil)
	cdc.RegisterConcrete(&TradingRewardCampaignLaunchProposal{}, "exchange/TradingRewardCampaignLaunchProposal", nil)
	cdc.RegisterConcrete(&TradingRewardCampaignUpdateProposal{}, "exchange/TradingRewardCampaignUpdateProposal", nil)
	cdc.RegisterConcrete(&TradingRewardPendingPointsUpdateProposal{}, "exchange/TradingRewardPendingPointsUpdateProposal", nil)
	cdc.RegisterConcrete(&FeeDiscountProposal{}, "exchange/FeeDiscountProposal", nil)
	cdc.RegisterConcrete(&BatchCommunityPoolSpendProposal{}, "exchange/BatchCommunityPoolSpendProposal", nil)
	cdc.RegisterConcrete(&BinaryOptionsMarketParamUpdateProposal{}, "exchange/BinaryOptionsMarketParamUpdateProposal", nil)
	cdc.RegisterConcrete(&BinaryOptionsMarketLaunchProposal{}, "exchange/BinaryOptionsMarketLaunchProposal", nil)
	cdc.RegisterConcrete(&AtomicMarketOrderFeeMultiplierScheduleProposal{}, "exchange/AtomicMarketOrderFeeMultiplierScheduleProposal", nil)

	cdc.RegisterConcrete(&CreateSpotLimitOrderAuthz{}, "exchange/CreateSpotLimitOrderAuthz", nil)
	cdc.RegisterConcrete(&CreateSpotMarketOrderAuthz{}, "exchange/CreateSpotMarketOrderAuthz", nil)
	cdc.RegisterConcrete(&BatchCreateSpotLimitOrdersAuthz{}, "exchange/BatchCreateSpotLimitOrdersAuthz", nil)
	cdc.RegisterConcrete(&CancelSpotOrderAuthz{}, "exchange/CancelSpotOrderAuthz", nil)
	cdc.RegisterConcrete(&BatchCancelSpotOrdersAuthz{}, "exchange/BatchCancelSpotOrdersAuthz", nil)
	cdc.RegisterConcrete(&CreateDerivativeLimitOrderAuthz{}, "exchange/CreateDerivativeLimitOrderAuthz", nil)
	cdc.RegisterConcrete(&CreateDerivativeMarketOrderAuthz{}, "exchange/CreateDerivativeMarketOrderAuthz", nil)
	cdc.RegisterConcrete(&BatchCreateDerivativeLimitOrdersAuthz{}, "exchange/BatchCreateDerivativeLimitOrdersAuthz", nil)
	cdc.RegisterConcrete(&CancelDerivativeOrderAuthz{}, "exchange/CancelDerivativeOrderAuthz", nil)
	cdc.RegisterConcrete(&BatchCancelDerivativeOrdersAuthz{}, "exchange/BatchCancelDerivativeOrdersAuthz", nil)
	cdc.RegisterConcrete(&BatchUpdateOrdersAuthz{}, "exchange/BatchUpdateOrdersAuthz", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeposit{},
		&MsgWithdraw{},
		&MsgInstantSpotMarketLaunch{},
		&MsgInstantPerpetualMarketLaunch{},
		&MsgInstantExpiryFuturesMarketLaunch{},
		&MsgCreateSpotLimitOrder{},
		&MsgBatchCreateSpotLimitOrders{},
		&MsgCreateSpotMarketOrder{},
		&MsgCancelSpotOrder{},
		&MsgBatchCancelSpotOrders{},
		&MsgCreateDerivativeLimitOrder{},
		&MsgBatchCreateDerivativeLimitOrders{},
		&MsgCreateDerivativeMarketOrder{},
		&MsgCancelDerivativeOrder{},
		&MsgBatchCancelDerivativeOrders{},
		&MsgBatchCancelBinaryOptionsOrders{},
		&MsgSubaccountTransfer{},
		&MsgExternalTransfer{},
		&MsgIncreasePositionMargin{},
		&MsgLiquidatePosition{},
		&MsgBatchUpdateOrders{},
		&MsgExec{},
		&MsgRewardsOptOut{},
		&MsgInstantBinaryOptionsMarketLaunch{},
		&MsgCreateBinaryOptionsLimitOrder{},
		&MsgCreateBinaryOptionsMarketOrder{},
		&MsgCancelBinaryOptionsOrder{},
		&MsgAdminUpdateBinaryOptionsMarket{},
	)

	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&ExchangeEnableProposal{},
		&BatchExchangeModificationProposal{},
		&SpotMarketParamUpdateProposal{},
		&SpotMarketLaunchProposal{},
		&PerpetualMarketLaunchProposal{},
		&ExpiryFuturesMarketLaunchProposal{},
		&DerivativeMarketParamUpdateProposal{},
		&MarketForcedSettlementProposal{},
		&UpdateDenomDecimalsProposal{},
		&TradingRewardCampaignLaunchProposal{},
		&TradingRewardCampaignUpdateProposal{},
		&TradingRewardPendingPointsUpdateProposal{},
		&FeeDiscountProposal{},
		&BatchCommunityPoolSpendProposal{},
		&BinaryOptionsMarketParamUpdateProposal{},
		&BinaryOptionsMarketLaunchProposal{},
		&AtomicMarketOrderFeeMultiplierScheduleProposal{},
	)

	registry.RegisterImplementations(
		(*authz.Authorization)(nil),
		// spot authz
		&CreateSpotLimitOrderAuthz{},
		&CreateSpotMarketOrderAuthz{},
		&BatchCreateSpotLimitOrdersAuthz{},
		&CancelSpotOrderAuthz{},
		&BatchCancelSpotOrdersAuthz{},
		// derivative authz
		&CreateDerivativeLimitOrderAuthz{},
		&CreateDerivativeMarketOrderAuthz{},
		&BatchCreateDerivativeLimitOrdersAuthz{},
		&CancelDerivativeOrderAuthz{},
		&BatchCancelDerivativeOrdersAuthz{},
		// common spot, derivative authz
		&BatchUpdateOrdersAuthz{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/exchange module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/exchange and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
