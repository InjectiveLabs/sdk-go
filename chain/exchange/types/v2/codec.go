package v2

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzcdc "github.com/cosmos/cosmos-sdk/x/authz/codec"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/cosmos/gogoproto/proto"
)

// todo: do we need this module cdc?

var (
	// ModuleCdc references the global x/exchange module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/exchange and
	// defined at the application level.
	ModuleCdc = codec.NewLegacyAmino()
)

func init() {
	RegisterLegacyAminoCodec(ModuleCdc)
	cryptocodec.RegisterCrypto(ModuleCdc)
	RegisterLegacyAminoCodec(authzcdc.Amino)

	ModuleCdc.Seal()
}

// RegisterLegacyAminoCodec registers the necessary x/exchange interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgDeposit{}, "exchange/v2/MsgDeposit", nil)
	cdc.RegisterConcrete(&MsgWithdraw{}, "exchange/v2/MsgWithdraw", nil)
	cdc.RegisterConcrete(&MsgInstantSpotMarketLaunch{}, "exchange/v2/MsgInstantSpotMarketLaunch", nil)
	cdc.RegisterConcrete(&MsgInstantPerpetualMarketLaunch{}, "exchange/v2/MsgInstantPerpetualMarketLaunch", nil)
	cdc.RegisterConcrete(&MsgInstantExpiryFuturesMarketLaunch{}, "exchange/v2/MsgInstantExpiryFuturesMarketLaunch", nil)
	cdc.RegisterConcrete(&MsgCreateSpotLimitOrder{}, "exchange/v2/MsgCreateSpotLimitOrder", nil)
	cdc.RegisterConcrete(&MsgBatchCreateSpotLimitOrders{}, "exchange/v2/MsgBatchCreateSpotLimitOrders", nil)
	cdc.RegisterConcrete(&MsgCreateSpotMarketOrder{}, "exchange/v2/MsgCreateSpotMarketOrder", nil)
	cdc.RegisterConcrete(&MsgCancelSpotOrder{}, "exchange/v2/MsgCancelSpotOrder", nil)
	cdc.RegisterConcrete(&MsgBatchCancelSpotOrders{}, "exchange/v2/MsgBatchCancelSpotOrders", nil)
	cdc.RegisterConcrete(&MsgCreateDerivativeLimitOrder{}, "exchange/v2/MsgCreateDerivativeLimitOrder", nil)
	cdc.RegisterConcrete(&MsgBatchCreateDerivativeLimitOrders{}, "exchange/v2/MsgBatchCreateDerivativeLimitOrders", nil)
	cdc.RegisterConcrete(&MsgCreateDerivativeMarketOrder{}, "exchange/v2/MsgCreateDerivativeMarketOrder", nil)
	cdc.RegisterConcrete(&MsgCancelDerivativeOrder{}, "exchange/v2/MsgCancelDerivativeOrder", nil)
	cdc.RegisterConcrete(&MsgBatchCancelDerivativeOrders{}, "exchange/v2/MsgBatchCancelDerivativeOrders", nil)
	cdc.RegisterConcrete(&MsgBatchCancelBinaryOptionsOrders{}, "exchange/v2/MsgBatchCancelBinaryOptionsOrders", nil)
	cdc.RegisterConcrete(&MsgSubaccountTransfer{}, "exchange/v2/MsgSubaccountTransfer", nil)
	cdc.RegisterConcrete(&MsgExternalTransfer{}, "exchange/v2/MsgExternalTransfer", nil)
	cdc.RegisterConcrete(&MsgIncreasePositionMargin{}, "exchange/v2/MsgIncreasePositionMargin", nil)
	cdc.RegisterConcrete(&MsgLiquidatePosition{}, "exchange/v2/MsgLiquidatePosition", nil)
	cdc.RegisterConcrete(&MsgBatchUpdateOrders{}, "exchange/v2/MsgBatchUpdateOrders", nil)
	cdc.RegisterConcrete(&MsgPrivilegedExecuteContract{}, "exchange/v2/MsgPrivilegedExecuteContract", nil)
	cdc.RegisterConcrete(&MsgRewardsOptOut{}, "exchange/v2/MsgRewardsOptOut", nil)
	cdc.RegisterConcrete(&MsgInstantBinaryOptionsMarketLaunch{}, "exchange/v2/MsgInstantBinaryOptionsMarketLaunch", nil)
	cdc.RegisterConcrete(&MsgCreateBinaryOptionsLimitOrder{}, "exchange/v2/MsgCreateBinaryOptionsLimitOrder", nil)
	cdc.RegisterConcrete(&MsgCreateBinaryOptionsMarketOrder{}, "exchange/v2/MsgCreateBinaryOptionsMarketOrder", nil)
	cdc.RegisterConcrete(&MsgCancelBinaryOptionsOrder{}, "exchange/v2/MsgCancelBinaryOptionsOrder", nil)
	cdc.RegisterConcrete(&MsgAdminUpdateBinaryOptionsMarket{}, "exchange/v2/MsgAdminUpdateBinaryOptionsMarket", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "exchange/v2/MsgUpdateParams", nil)
	cdc.RegisterConcrete(&MsgUpdateSpotMarket{}, "exchange/v2/MsgUpdateSpotMarket", nil)
	cdc.RegisterConcrete(&MsgUpdateDerivativeMarket{}, "exchange/v2/MsgUpdateDerivativeMarket", nil)

	cdc.RegisterConcrete(&ExchangeEnableProposal{}, "exchange/v2/ExchangeEnableProposal", nil)
	cdc.RegisterConcrete(&BatchExchangeModificationProposal{}, "exchange/v2/BatchExchangeModificationProposal", nil)
	cdc.RegisterConcrete(&SpotMarketParamUpdateProposal{}, "exchange/v2/SpotMarketParamUpdateProposal", nil)
	cdc.RegisterConcrete(&SpotMarketLaunchProposal{}, "exchange/v2/SpotMarketLaunchProposal", nil)
	cdc.RegisterConcrete(&PerpetualMarketLaunchProposal{}, "exchange/v2/PerpetualMarketLaunchProposal", nil)
	cdc.RegisterConcrete(&ExpiryFuturesMarketLaunchProposal{}, "exchange/v2/ExpiryFuturesMarketLaunchProposal", nil)
	cdc.RegisterConcrete(&DerivativeMarketParamUpdateProposal{}, "exchange/v2/DerivativeMarketParamUpdateProposal", nil)
	cdc.RegisterConcrete(&MarketForcedSettlementProposal{}, "exchange/v2/MarketForcedSettlementProposal", nil)
	cdc.RegisterConcrete(&UpdateDenomDecimalsProposal{}, "exchange/v2/UpdateDenomDecimalsProposal", nil)
	cdc.RegisterConcrete(&TradingRewardCampaignLaunchProposal{}, "exchange/v2/TradingRewardCampaignLaunchProposal", nil)
	cdc.RegisterConcrete(&TradingRewardCampaignUpdateProposal{}, "exchange/v2/TradingRewardCampaignUpdateProposal", nil)
	cdc.RegisterConcrete(&TradingRewardPendingPointsUpdateProposal{}, "exchange/v2/TradingRewardPendingPointsUpdateProposal", nil)
	cdc.RegisterConcrete(&FeeDiscountProposal{}, "exchange/v2/FeeDiscountProposal", nil)
	cdc.RegisterConcrete(&BatchCommunityPoolSpendProposal{}, "exchange/v2/BatchCommunityPoolSpendProposal", nil)
	cdc.RegisterConcrete(&BinaryOptionsMarketParamUpdateProposal{}, "exchange/v2/BinaryOptionsMarketParamUpdateProposal", nil)
	cdc.RegisterConcrete(&BinaryOptionsMarketLaunchProposal{}, "exchange/v2/BinaryOptionsMarketLaunchProposal", nil)
	cdc.RegisterConcrete(&AtomicMarketOrderFeeMultiplierScheduleProposal{}, "exchange/v2/AtomicMarketOrderFeeMultiplierScheduleProposal", nil)

	cdc.RegisterConcrete(&CreateSpotLimitOrderAuthz{}, "exchange/v2/CreateSpotLimitOrderAuthz", nil)
	cdc.RegisterConcrete(&CreateSpotMarketOrderAuthz{}, "exchange/v2/CreateSpotMarketOrderAuthz", nil)
	cdc.RegisterConcrete(&BatchCreateSpotLimitOrdersAuthz{}, "exchange/v2/BatchCreateSpotLimitOrdersAuthz", nil)
	cdc.RegisterConcrete(&CancelSpotOrderAuthz{}, "exchange/v2/CancelSpotOrderAuthz", nil)
	cdc.RegisterConcrete(&BatchCancelSpotOrdersAuthz{}, "exchange/v2/BatchCancelSpotOrdersAuthz", nil)
	cdc.RegisterConcrete(&CreateDerivativeLimitOrderAuthz{}, "exchange/v2/CreateDerivativeLimitOrderAuthz", nil)
	cdc.RegisterConcrete(&CreateDerivativeMarketOrderAuthz{}, "exchange/v2/CreateDerivativeMarketOrderAuthz", nil)
	cdc.RegisterConcrete(&BatchCreateDerivativeLimitOrdersAuthz{}, "exchange/v2/BatchCreateDerivativeLimitOrdersAuthz", nil)
	cdc.RegisterConcrete(&CancelDerivativeOrderAuthz{}, "exchange/v2/CancelDerivativeOrderAuthz", nil)
	cdc.RegisterConcrete(&BatchCancelDerivativeOrdersAuthz{}, "exchange/v2/BatchCancelDerivativeOrdersAuthz", nil)
	cdc.RegisterConcrete(&BatchUpdateOrdersAuthz{}, "exchange/v2/BatchUpdateOrdersAuthz", nil)

	cdc.RegisterConcrete(&Params{}, "exchange/v2/Params", nil)
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
		&MsgPrivilegedExecuteContract{},
		&MsgRewardsOptOut{},
		&MsgInstantBinaryOptionsMarketLaunch{},
		&MsgCreateBinaryOptionsLimitOrder{},
		&MsgCreateBinaryOptionsMarketOrder{},
		&MsgCancelBinaryOptionsOrder{},
		&MsgAdminUpdateBinaryOptionsMarket{},
		&MsgUpdateParams{},
		&MsgUpdateSpotMarket{},
		&MsgUpdateDerivativeMarket{},
	)

	registry.RegisterImplementations(
		(*proto.Message)(nil),
		&MsgCreateSpotLimitOrderResponse{},
		&MsgCreateSpotMarketOrderResponse{},
		&MsgBatchCreateSpotLimitOrdersResponse{},
		&MsgCreateDerivativeLimitOrderResponse{},
		&MsgCreateDerivativeMarketOrderResponse{},
		&MsgBatchCreateDerivativeLimitOrdersResponse{},
		&MsgCreateBinaryOptionsLimitOrderResponse{},
		&MsgCreateBinaryOptionsMarketOrderResponse{},
		&MsgBatchUpdateOrdersResponse{},
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
