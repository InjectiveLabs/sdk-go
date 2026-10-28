package v2

import (
	"errors"

	"cosmossdk.io/math"

	"github.com/InjectiveLabs/sdk-go/chain/exchange/types"
)

// CanonicalIsolatedLimitHold uses the cancellation formula and its rounding
// order, checking raw results before constructing LegacyDec values.
//
//nolint:revive // Keep input validation and the exact cancellation arithmetic together.
func CanonicalIsolatedLimitHold(order *DerivativeLimitOrder, market DerivativeMarketI, feeRate math.LegacyDec) (math.LegacyDec, error) {
	if order == nil || market == nil {
		return math.LegacyDec{}, errors.New("missing isolated limit hold input")
	}
	for _, value := range []math.LegacyDec{order.Fillable, order.OrderInfo.Quantity, order.OrderInfo.Price, order.Margin, feeRate} {
		if value.IsNil() || !value.IsInValidRange() {
			return math.LegacyDec{}, errors.New("invalid isolated limit hold input")
		}
	}
	if order.Fillable.IsNegative() || !order.OrderInfo.Quantity.IsPositive() ||
		order.Fillable.GT(order.OrderInfo.Quantity) || order.OrderInfo.Price.IsNegative() || order.Margin.IsNegative() {
		return math.LegacyDec{}, errors.New("invalid isolated limit hold quantity, price or margin")
	}
	if !order.IsVanilla() {
		return math.LegacyZeroDec(), nil
	}
	marginRaw := order.Margin.BigInt()
	if !order.Fillable.Equal(order.OrderInfo.Quantity) {
		marginRaw = proportionalPositionMarginRaw(order.Margin, order.Fillable, order.OrderInfo.Quantity)
	}
	if marginRaw.Cmp(legacyDecRangeExclusiveRaw) >= 0 {
		return math.LegacyDec{}, errors.New("isolated margin refund is out of range")
	}
	notionalRaw, err := checkedLegacyMulRaw(order.Fillable, order.OrderInfo.Price)
	if err != nil {
		return math.LegacyDec{}, err
	}
	feeRaw, err := checkedLegacyMulRaw(math.LegacyNewDecFromBigIntWithPrec(notionalRaw, math.LegacyPrecision), math.LegacyMaxDec(feeRate, math.LegacyZeroDec()))
	if err != nil {
		return math.LegacyDec{}, err
	}
	holdRaw, err := checkedLegacyAddRaw(marginRaw, feeRaw)
	if err != nil {
		return math.LegacyDec{}, err
	}
	hold, err := CanonicalIsolatedMarketHold(math.LegacyNewDecFromBigIntWithPrec(holdRaw, math.LegacyPrecision), market)
	if err != nil {
		return math.LegacyDec{}, err
	}

	return hold, nil
}

// HasExactFullMarginAllocation reports whether matching an entire order preserves
// its margin under the existing proportional allocation arithmetic.
func (m *DerivativeLimitOrder) HasExactFullMarginAllocation() bool {
	if m == nil || m.Margin.IsNil() || m.OrderInfo.Quantity.IsNil() || !m.OrderInfo.Quantity.IsPositive() {
		return false
	}
	return proportionalPositionMarginRaw(m.Margin, m.OrderInfo.Quantity, m.OrderInfo.Quantity).Cmp(m.Margin.BigInt()) == 0
}

// CanonicalIsolatedMarketHold converts a stored human-format MarginHold once.
func CanonicalIsolatedMarketHold(hold math.LegacyDec, market DerivativeMarketI) (math.LegacyDec, error) {
	if market == nil || hold.IsNil() || hold.IsNegative() || !hold.IsInValidRange() ||
		market.GetQuoteDecimals() > types.MaxDecimals || !types.CanRepresentNotionalInChainFormat(hold, market.GetQuoteDecimals()) {
		return math.LegacyDec{}, errors.New("invalid or unrepresentable isolated derivative hold")
	}
	return market.NotionalToChainFormat(hold), nil
}
