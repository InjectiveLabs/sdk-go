package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
)

func NewMarketOrderForLiquidation(position *Position, positionSubaccountID common.Hash, liquidator sdk.AccAddress) *DerivativeMarketOrder {

	// if long position, market sell order at price 0
	// if short position, market buy order at price infinity
	var worstPrice sdk.Dec
	var orderType OrderType
	if position.IsLong {
		worstPrice = sdk.ZeroDec()
		orderType = OrderType_SELL
	} else {
		// cap the worst price to sell a position at 200x the entry price
		worstPrice = position.EntryPrice.MulInt64(200)
		orderType = OrderType_BUY
	}

	order := DerivativeMarketOrder{
		OrderInfo: OrderInfo{
			SubaccountId: positionSubaccountID.Hex(),
			FeeRecipient: liquidator.String(),
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

func (m *DerivativeLimitOrder) ToTrimmed() *TrimmedDerivativeLimitOrder {
	return &TrimmedDerivativeLimitOrder{
		Price:     m.OrderInfo.Price,
		Quantity:  m.OrderInfo.Quantity,
		Margin:    m.Margin,
		Fillable:  m.Fillable,
		IsBuy:     m.IsBuy(),
		OrderHash: common.Bytes2Hex(m.OrderHash),
	}
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
) {
	order := o.MarketOrder
	subaccountID := order.SubaccountID()
	// For vanilla orders that were not executed at all, increment the available balance
	if order.IsVanilla() {
		depositDelta := o.GetCancelDepositDelta()
		if depositDelta != nil {
			depositDeltas.ApplyDepositDelta(subaccountID, depositDelta)
		}
	}
}

func NewDerivativeMarketOrder(o *DerivativeOrder, orderHash common.Hash) *DerivativeMarketOrder {
	return &DerivativeMarketOrder{
		OrderInfo:    o.OrderInfo,
		OrderType:    o.OrderType,
		Margin:       o.Margin,
		MarginHold:   sdk.ZeroDec(),
		TriggerPrice: o.TriggerPrice,
		OrderHash:    orderHash.Bytes(),
	}
}
func NewDerivativeLimitOrder(o *DerivativeOrder, orderHash common.Hash) *DerivativeLimitOrder {
	return &DerivativeLimitOrder{
		OrderInfo:    o.OrderInfo,
		OrderType:    o.OrderType,
		Margin:       o.Margin,
		Fillable:     o.OrderInfo.Quantity,
		TriggerPrice: o.TriggerPrice,
		OrderHash:    orderHash.Bytes(),
	}
}

func (o *DerivativeLimitOrder) ToDerivativeOrder(marketID string) *DerivativeOrder {
	return &DerivativeOrder{
		MarketId:     marketID,
		OrderInfo:    o.OrderInfo,
		OrderType:    o.OrderType,
		Margin:       o.Margin,
		TriggerPrice: o.TriggerPrice,
	}
}
func (o *DerivativeMarketOrder) ToDerivativeOrder(marketID string) *DerivativeOrder {
	return &DerivativeOrder{
		MarketId:     marketID,
		OrderInfo:    o.OrderInfo,
		OrderType:    o.OrderType,
		Margin:       o.Margin,
		TriggerPrice: o.TriggerPrice,
	}
}

func (o *DerivativeLimitOrder) GetCancelDepositDelta(feeRate sdk.Dec) *DepositDelta {
	depositDelta := NewDepositDelta()
	if o.IsVanilla() {
		// Refund = (FillableQuantity / Quantity) * (Margin + Price * Quantity * feeRate)
		notional := o.OrderInfo.Price.Mul(o.OrderInfo.Quantity)
		marginHoldRefund := o.Fillable.Mul(o.Margin.Add(notional.Mul(feeRate))).Quo(o.OrderInfo.Quantity)
		depositDelta.AvailableBalanceDelta = marginHoldRefund
	}
	return depositDelta
}

func (o *DerivativeOrder) CheckTickSize(minPriceTickSize, minQuantityTickSize sdk.Dec) error {
	if BreachesMinimumTickSize(o.OrderInfo.Price, minPriceTickSize) {
		return sdkerrors.Wrapf(ErrInvalidPrice, "price %s must be a multiple of the minimum price tick size %s", o.OrderInfo.Price.String(), minPriceTickSize.String())
	}
	if BreachesMinimumTickSize(o.OrderInfo.Quantity, minQuantityTickSize) {
		return sdkerrors.Wrapf(ErrInvalidQuantity, "quantity %s must be a multiple of the minimum quantity tick size %s", o.OrderInfo.Quantity.String(), minQuantityTickSize.String())
	}
	if !o.Margin.IsZero() {
		if BreachesMinimumTickSize(o.Margin, minQuantityTickSize) {
			return sdkerrors.Wrapf(ErrInvalidMargin, "margin %s must be a multiple of the minimum price tick size %s", o.Margin.String(), minPriceTickSize.String())
		}
	}
	return nil
}

func (o *DerivativeOrder) CheckMarginAndGetMarginHold(market *DerivativeMarket, markPrice, feeRate sdk.Dec) (marginHold sdk.Dec, err error) {
	notional := o.OrderInfo.Price.Mul(o.OrderInfo.Quantity)
	feeAmount := notional.Mul(feeRate)

	// Margin ≥ InitialMarginRatio * Price * Quantity
	if o.Margin.LT(market.InitialMarginRatio.Mul(notional)) {
		return sdk.Dec{}, sdkerrors.Wrapf(ErrInsufficientOrderMargin, "InitialMarginRatio Check: need at least %s but got %s", market.InitialMarginRatio.Mul(notional).String(), o.Margin.String())
	}

	if err = o.CheckInitialMarginRequirementMarkPriceThreshold(market.InitialMarginRatio, markPrice); err != nil {
		return sdk.Dec{}, err
	}

	return o.Margin.Add(feeAmount), nil
}

func (o *DerivativeOrder) CheckInitialMarginRequirementMarkPriceThreshold(initialMarginRatio, markPrice sdk.Dec) (err error) {
	markPriceThreshold := o.ComputeInitialMarginRequirementMarkPriceThreshold(initialMarginRatio)
	// For Buys: MarkPrice ≥ (Margin - Price * Quantity) / ((InitialMarginRatio - 1) * Quantity)
	// For Sells: MarkPrice ≤ (Margin + Price * Quantity) / ((1+ InitialMarginRatio) * Quantity)
	if o.OrderType.IsBuy() && markPrice.LT(markPriceThreshold) {
		return sdkerrors.Wrapf(ErrInsufficientOrderMargin, "Buy MarkPriceThreshold Check: mark price %s must be GTE %s", markPrice.String(), markPriceThreshold.String())
	} else if !o.OrderType.IsBuy() && markPrice.GT(markPriceThreshold) {
		return sdkerrors.Wrapf(ErrInsufficientOrderMargin, "Sell MarkPriceThreshold Check: mark price %s must be LTE %s", markPrice.String(), markPriceThreshold.String())
	}

	return nil
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

func (o *DerivativeLimitOrder) CheckInitialMarginRequirementMarkPriceThreshold(initialMarginRatio, markPrice sdk.Dec) (err error) {
	return o.ToDerivativeOrder("").CheckInitialMarginRequirementMarkPriceThreshold(initialMarginRatio, markPrice)
}

func (o *DerivativeMarketOrder) CheckInitialMarginRequirementMarkPriceThreshold(initialMarginRatio, markPrice sdk.Dec) (err error) {
	return o.ToDerivativeOrder("").CheckInitialMarginRequirementMarkPriceThreshold(initialMarginRatio, markPrice)
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

func (o *DerivativeLimitOrder) Hash() common.Hash {
	return common.BytesToHash(o.OrderHash)
}

func (o *DerivativeMarketOrder) Hash() common.Hash {
	return common.BytesToHash(o.OrderHash)
}

func (o *DerivativeLimitOrder) FeeRecipient() common.Address {
	return o.OrderInfo.FeeRecipientAddress()
}

func (o *DerivativeMarketOrder) FeeRecipient() common.Address {
	return o.OrderInfo.FeeRecipientAddress()
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

func (m *DerivativeOrder) IsBuy() bool {
	return m.OrderType.IsBuy()
}

func (m *DerivativeMarketOrder) Quantity() sdk.Dec {
	return m.OrderInfo.Quantity
}

func (m *DerivativeMarketOrder) FillableQuantity() sdk.Dec {
	return m.OrderInfo.Quantity
}

func (m *DerivativeMarketOrder) Price() sdk.Dec {
	return m.OrderInfo.Price
}

func (m *DerivativeLimitOrder) Price() sdk.Dec {
	return m.OrderInfo.Price
}

func (m *DerivativeOrder) Price() sdk.Dec {
	return m.OrderInfo.Price
}

func (o *DerivativeOrder) SubaccountID() common.Hash {
	return o.OrderInfo.SubaccountID()
}

func (o *DerivativeOrder) MarketID() common.Hash {
	return common.HexToHash(o.MarketId)
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

func (o *OrderInfo) FeeRecipientAddress() common.Address {
	address, _ := sdk.AccAddressFromBech32(o.FeeRecipient)
	return common.BytesToAddress(address.Bytes())
}

func (o *DerivativeLimitOrder) SdkAccAddress() sdk.AccAddress {
	return sdk.AccAddress(o.SubaccountID().Bytes()[:common.AddressLength])
}

func (o *DerivativeMarketOrder) SdkAccAddress() sdk.AccAddress {
	return sdk.AccAddress(o.SubaccountID().Bytes()[:common.AddressLength])
}
