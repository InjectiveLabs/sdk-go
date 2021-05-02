package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
)

const Int64Max int64 = 9223372036854775807

func NewMarketOrderForLiquidation(position *Position, positionSubaccountID common.Hash, liquidator sdk.Address) *DerivativeMarketOrder {

	// if long position, market sell order at price 0
	// if short position, market buy order at price infinity
	var worstPrice sdk.Dec
	var orderType OrderType
	if position.IsLong {
		worstPrice = sdk.ZeroDec()
		orderType = OrderType_SELL
	} else {
		// TODO: should use biggest sdk.Dec instead, which can be larger than Int64Max
		worstPrice = sdk.NewDec(Int64Max)
		orderType = OrderType_BUY
	}

	order := DerivativeMarketOrder{
		OrderInfo: OrderInfo{
			SubaccountId: positionSubaccountID.Hex(),
			FeeRecipient: SdkAddressToEthAddress(liquidator).Hex(),
			Price:        worstPrice,
			Quantity:     position.Quantity,
		},
		OrderType:    orderType,
		Margin:       sdk.ZeroDec(),
		MarginHold:   sdk.ZeroDec(),
		TriggerPrice: nil,
	}

	return &order
}

func (o *DerivativeMarketOrderCancel) GetCancelDepositDelta() *DepositDelta {
	order := o.MarketOrder
	// no market order quantity was executed, so refund the entire margin hold
	if order.IsVanilla() && o.CancelQuantity.Equal(order.OrderInfo.Quantity) {
		return &DepositDelta{
			AvailableBalanceDelta: order.MarginHold,
			TotalBalanceDelta:     sdk.ZeroDec(),
		}
	}
	// TODO: double check that partial market order executions are covered earlier upstream
	return nil
}

func (o *DerivativeMarketOrderCancel) ApplyDerivativeMarketCancellation(
	depositDeltas DepositDeltas,
	positionStates PositionStates,
) {
	order := o.MarketOrder
	subaccountID := order.SubaccountID()
	// For vanilla orders that were not executed at all, increment the available balance
	// For reduce-only orders, free the position hold quantity
	if order.IsVanilla() {
		depositDelta := o.GetCancelDepositDelta()
		if depositDelta != nil {
			depositDeltas.ApplyDepositDelta(subaccountID, depositDelta)
		}
	} else if order.IsReduceOnly() {
		position := positionStates[subaccountID].Position
		position.HoldQuantity = position.HoldQuantity.Sub(o.CancelQuantity)
	}
}

func (o *DerivativeOrder) GetNewDerivativeLimitOrder(orderHash common.Hash) *DerivativeLimitOrder {
	return &DerivativeLimitOrder{
		OrderInfo:    o.OrderInfo,
		OrderType:    o.OrderType,
		Margin:       o.Margin,
		Fillable:     o.OrderInfo.Quantity,
		TriggerPrice: o.TriggerPrice,
		OrderHash:    orderHash.Bytes(),
	}
}

func (o *DerivativeLimitOrder) GetCancelDepositDelta(makerFeeRate sdk.Dec) *DepositDelta {
	depositDelta := NewDepositDelta()
	if o.IsVanilla() {
		// Refund = (Fillable / Quantity) * (Margin + Price * Quantity * MakerFeeRate)
		fillableFraction := o.Fillable.Quo(o.OrderInfo.Quantity)
		notional := o.OrderInfo.Price.Mul(o.OrderInfo.Quantity)
		marginHoldRefund := fillableFraction.Mul(o.Margin.Add(notional.Mul(makerFeeRate)))
		depositDelta.AvailableBalanceDelta = marginHoldRefund
	}
	return depositDelta
}

func (o *DerivativeOrder) CheckTickSize(minPriceTickSize, minQuantityTickSize sdk.Dec) error {
	// reject order if non-zero price decimals or trigger price decimals exceeds market.MaxPriceScaleDecimals or quantity decimals exceeds market.MaxQuantityScaleDecimals
	if breachesMinimumTickSize(o.OrderInfo.Price, minPriceTickSize) {
		return sdkerrors.Wrapf(ErrInvalidPrice, "price exceeds market maximum price decimals %s", minPriceTickSize.String())
	}
	if breachesMinimumTickSize(o.OrderInfo.Quantity, minQuantityTickSize) {
		return sdkerrors.Wrapf(ErrInvalidQuantity, "quantity exceeds market maximum quantity decimals %s", minQuantityTickSize.String())
	}
	return nil
}

func (o *DerivativeOrder) CheckMarginAndGetMarginHold(market *DerivativeMarket, markPrice sdk.Dec) (marginHold sdk.Dec, err error) {
	notional := o.OrderInfo.Price.Mul(o.OrderInfo.Quantity)
	feeAmount := notional.Mul(market.TakerFeeRate)

	// Margin ≥ InitialMarginRatio * Price * Quantity
	if o.Margin.LT(market.InitialMarginRatio.Mul(notional)) {
		return sdk.Dec{}, sdkerrors.Wrapf(ErrInsufficientOrderMargin, "InitialMarginRatio Check: need at least %s but got %s", market.InitialMarginRatio.Mul(notional).String(), o.Margin.String())
	}

	markPriceThreshold := o.ComputeInitialMarginRequirementMarkPriceThreshold(market.InitialMarginRatio)
	// For Buys: MarkPrice ≥ (Margin - Price * Quantity) / ((InitialMarginRatio - 1) * Quantity)
	// For Sells: MarkPrice ≤ (Margin + Price * Quantity) / ((1+ InitialMarginRatio) * Quantity)
	if o.OrderType.IsBuy() && markPrice.LT(markPriceThreshold) {
		return sdk.Dec{}, sdkerrors.Wrapf(ErrInsufficientOrderMargin, "Buy MarkPriceThreshold Check: mark price %s must be GTE %s", markPrice.String(), markPriceThreshold.String())
	} else if !o.OrderType.IsBuy() && markPrice.GT(markPriceThreshold) {

		return sdk.Dec{}, sdkerrors.Wrapf(ErrInsufficientOrderMargin, "Sell MarkPriceThreshold Check: mark price %s must be LTE %s", markPrice.String(), markPriceThreshold.String())
	}

	return o.Margin.Add(feeAmount), nil
}

func (o *DerivativeOrder) ComputeInitialMarginRequirementMarkPriceThreshold(initialMarginRatio sdk.Dec) sdk.Dec {
	notional := o.OrderInfo.Price.Mul(o.OrderInfo.Quantity)
	var numerator, denominator sdk.Dec
	if o.OrderType.IsBuy() {
		numerator = o.Margin.Sub(notional)
		denominator = initialMarginRatio.Sub(sdk.OneDec()).Mul(o.OrderInfo.Quantity)
	} else {
		numerator = o.Margin.Add(notional)
		denominator = initialMarginRatio.Add(sdk.OneDec()).Mul(o.OrderInfo.Quantity)
	}
	return numerator.Quo(denominator)
}

func (o *DerivativeMarketOrder) ComputeOrderHash(nonce uint32, marketId string) (common.Hash, error) {
	triggerPrice := ""
	if o.TriggerPrice != nil {
		triggerPrice = o.TriggerPrice.String()
	}

	return computeOrderHash(marketId, o.OrderInfo.SubaccountId, o.OrderInfo.FeeRecipient, o.OrderInfo.Price.String(), o.OrderInfo.Quantity.String(), o.Margin.String(), triggerPrice, string(o.OrderType), nonce)
}

// ComputeOrderHash computes the order hash for given derivative limit order
func (o *DerivativeOrder) ComputeOrderHash(nonce uint32) (common.Hash, error) {
	triggerPrice := ""
	if o.TriggerPrice != nil {
		triggerPrice = o.TriggerPrice.String()
	}
	return computeOrderHash(o.MarketId, o.OrderInfo.SubaccountId, o.OrderInfo.FeeRecipient, o.OrderInfo.Price.String(), o.OrderInfo.Quantity.String(), o.Margin.String(), triggerPrice, string(o.OrderType), nonce)
}

func (o *DerivativeOrder) IsReduceOnly() bool {
	return o.Margin.IsZero()
}

func (o *DerivativeMarketOrder) IsReduceOnly() bool {
	return o.Margin.IsZero()
}

func (o *DerivativeLimitOrder) IsReduceOnly() bool {
	return o.Margin.IsZero()
}

func (o *DerivativeOrder) IsVanilla() bool {
	return !o.IsReduceOnly()
}

func (o *DerivativeMarketOrder) IsVanilla() bool {
	return !o.IsReduceOnly()
}

func (o *DerivativeLimitOrder) IsVanilla() bool {
	return !o.IsReduceOnly()
}

func (m *DerivativeMarketOrder) IsBuy() bool {
	return m.OrderType.IsBuy()
}

func (m *DerivativeLimitOrder) IsBuy() bool {
	return m.OrderType.IsBuy()
}

func (o *DerivativeOrder) SubaccountID() common.Hash {
	return o.OrderInfo.SubaccountID()
}

func (o *DerivativeMarketOrder) SubaccountID() common.Hash {
	return o.OrderInfo.SubaccountID()
}

func (o *DerivativeLimitOrder) SubaccountID() common.Hash {
	return o.OrderInfo.SubaccountID()
}

func (o *OrderInfo) SubaccountID() common.Hash {
	return common.HexToHash(o.SubaccountId)
}
