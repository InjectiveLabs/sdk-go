package types

import (
	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common"
)

func GetScaledPrice(price math.LegacyDec, scaleFactor uint32) math.LegacyDec {
	return price.Mul(math.LegacyNewDec(10).Power(uint64(scaleFactor)))
}

func (o *DerivativeOrder) MarketID() common.Hash {
	return common.HexToHash(o.MarketId)
}

func (o *DerivativeLimitOrder) Cid() string {
	return o.OrderInfo.GetCid()
}

func (o *DerivativeMarketOrder) Cid() string {
	return o.OrderInfo.GetCid()
}

func (o *DerivativeOrder) Cid() string {
	return o.OrderInfo.GetCid()
}

func (m *DerivativeMarketOrder) Price() math.LegacyDec {
	return m.OrderInfo.Price
}

func (m *DerivativeLimitOrder) Price() math.LegacyDec {
	return m.OrderInfo.Price
}

func (m *DerivativeOrder) Price() math.LegacyDec {
	return m.OrderInfo.Price
}

func (o *DerivativeOrder) SubaccountID() common.Hash {
	return o.OrderInfo.SubaccountID()
}

func (o *DerivativeMarketOrder) SubaccountID() common.Hash {
	return o.OrderInfo.SubaccountID()
}

func (m *DerivativeLimitOrder) SubaccountID() common.Hash {
	return m.OrderInfo.SubaccountID()
}

func (m *OrderInfo) SubaccountID() common.Hash {
	return common.HexToHash(m.SubaccountId)
}

// Test Code Only (for v1 tests)
func (o *DerivativeMarketOrder) ComputeOrderHash(nonce uint32, marketId string) (common.Hash, error) {
	triggerPrice := ""
	if o.TriggerPrice != nil {
		triggerPrice = o.TriggerPrice.String()
	}

	return ComputeOrderHash(marketId, o.OrderInfo.SubaccountId, o.OrderInfo.FeeRecipient, o.OrderInfo.Price.String(), o.OrderInfo.Quantity.String(), o.Margin.String(), triggerPrice, string(o.OrderType), nonce)
}

// Test Code Only (for v1 tests)
func (o *DerivativeOrder) ComputeOrderHash(nonce, scaleFactor uint32) (common.Hash, error) {
	triggerPrice := ""
	if o.TriggerPrice != nil {
		triggerPrice = o.TriggerPrice.String()
	}

	o.OrderInfo.Price = PriceFromChainFormat(o.OrderInfo.Price, 0, scaleFactor)
	o.Margin = NotionalFromChainFormat(o.Margin, scaleFactor)

	return ComputeOrderHash(o.MarketId, o.OrderInfo.SubaccountId, o.OrderInfo.FeeRecipient, o.OrderInfo.Price.String(), o.OrderInfo.Quantity.String(), o.Margin.String(), triggerPrice, string(o.OrderType), nonce)
}

// Test Code Only (for v1 tests)
func (o *DerivativeLimitOrder) GetCancelDepositDelta(feeRate math.LegacyDec) *DepositDelta {
	return &DepositDelta{
		AvailableBalanceDelta: o.GetCancelRefundAmount(feeRate),
		TotalBalanceDelta:     math.LegacyZeroDec(),
	}
}

// Test Code Only (for v1 tests)
func (o *DerivativeLimitOrder) GetCancelRefundAmount(feeRate math.LegacyDec) math.LegacyDec {
	marginHoldRefund := math.LegacyZeroDec()
	if !o.Margin.IsZero() {
		positiveFeePart := math.LegacyMaxDec(math.LegacyZeroDec(), feeRate)
		notional := o.OrderInfo.Price.Mul(o.OrderInfo.Quantity)
		marginHoldRefund = o.Fillable.Mul(o.Margin.Add(notional.Mul(positiveFeePart))).Quo(o.OrderInfo.Quantity)
	}
	return marginHoldRefund
}

// Test Code Only (for v1 tests)
func (o *DerivativeOrder) CheckInitialMarginRequirementMarkPriceThreshold(initialMarginRatio, markPrice math.LegacyDec) (err error) {
	// For Buys: MarkPrice ≥ (Margin - Price * Quantity) / ((InitialMarginRatio - 1) * Quantity)
	// For Sells: MarkPrice ≤ (Margin + Price * Quantity) / ((1 + InitialMarginRatio) * Quantity)
	markPriceThreshold := o.ComputeInitialMarginRequirementMarkPriceThreshold(initialMarginRatio)
	return CheckInitialMarginMarkPriceRequirement(o.OrderType.IsBuy(), markPriceThreshold, markPrice)
}

// Test Code Only (for v1 tests)
func CheckInitialMarginMarkPriceRequirement(isBuyOrLong bool, markPriceThreshold, markPrice math.LegacyDec) error {
	if isBuyOrLong && markPrice.LT(markPriceThreshold) {
		return errors.Wrapf(ErrInsufficientMargin, "Buy MarkPriceThreshold Check: mark/trigger price %s must be GTE %s", markPrice.String(), markPriceThreshold.String())
	} else if !isBuyOrLong && markPrice.GT(markPriceThreshold) {
		return errors.Wrapf(ErrInsufficientMargin, "Sell MarkPriceThreshold Check: mark/trigger price %s must be LTE %s", markPrice.String(), markPriceThreshold.String())
	}
	return nil
}

// Test Code Only (for v1 tests)
func (o *DerivativeOrder) ComputeInitialMarginRequirementMarkPriceThreshold(initialMarginRatio math.LegacyDec) math.LegacyDec {
	return ComputeMarkPriceThreshold(o.OrderType.IsBuy(), o.Price(), o.OrderInfo.Quantity, o.Margin, initialMarginRatio)
}

// Test Code Only (for v1 tests)
func ComputeMarkPriceThreshold(isBuyOrLong bool, price, quantity, margin, initialMarginRatio math.LegacyDec) math.LegacyDec {
	notional := price.Mul(quantity)
	var numerator, denominator math.LegacyDec
	if isBuyOrLong {
		numerator = margin.Sub(notional)
		denominator = initialMarginRatio.Sub(math.LegacyOneDec()).Mul(quantity)
	} else {
		numerator = margin.Add(notional)
		denominator = initialMarginRatio.Add(math.LegacyOneDec()).Mul(quantity)
	}
	return numerator.Quo(denominator)
}

// Test Code Only (for v1 tests)
func ResizeReduceOnlyOrder(o IMutableDerivativeOrder, newQuantity math.LegacyDec) error {
	if o.IsVanilla() {
		return ErrOrderInvalid.Wrap("ResizeReduceOnlyOrder should only be used for reduce only orders!")
	}

	quantityDecrement := o.GetQuantity().Sub(newQuantity)
	if !quantityDecrement.IsPositive() {
		return nil
	}

	o.SetQuantity(newQuantity)
	return nil
}
