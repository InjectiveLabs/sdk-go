package v2

import (
	"errors"

	"cosmossdk.io/math"

	"github.com/InjectiveLabs/sdk-go/chain/exchange/types"
)

// IsolatedOrderMarginForFill allocates margin to a fill from the decrease in
// its remaining cancellation claim. Ordinary pure-isolated journal execution
// and its closing-prefix probes share this allocation, including rounding at
// tiny quantities. Legacy cross/hybrid/liquidation allocation and the position
// close/open split within one fill retain their existing arithmetic.
func IsolatedOrderMarginForFill(margin, orderQuantity, remaining, fillQuantity math.LegacyDec) (math.LegacyDec, error) {
	for _, value := range []math.LegacyDec{margin, orderQuantity, remaining, fillQuantity} {
		if value.IsNil() || !value.IsInValidRange() {
			return math.LegacyDec{}, errors.New("invalid isolated order margin allocation input")
		}
	}
	if margin.IsNegative() || !orderQuantity.IsPositive() || remaining.IsNegative() || fillQuantity.IsNegative() ||
		remaining.GT(orderQuantity) || fillQuantity.GT(remaining) {
		return math.LegacyDec{}, errors.New("invalid isolated order margin allocation quantity or margin")
	}
	before := proportionalOrderMarginRaw(margin, remaining, orderQuantity)
	after := proportionalOrderMarginRaw(margin, remaining.Sub(fillQuantity), orderQuantity)
	// The validated claims decrease from at most margin to zero. Their
	// difference therefore fits LegacyDec even when the raw products do not.
	return math.LegacyNewDecFromBigIntWithPrec(before.Sub(before, after), math.LegacyPrecision), nil
}

// CanonicalIsolatedLimitHold uses the cancellation formula, checking raw results
// before constructing LegacyDec values.
func CanonicalIsolatedLimitHold(order *DerivativeLimitOrder, market DerivativeMarketI, feeRate math.LegacyDec) (math.LegacyDec, error) {
	return canonicalIsolatedLimitHold(order, market, feeRate, false)
}

// LegacyIsolatedLimitHoldForMigration reconstructs the pre-upgrade cancellation
// claim. Only upgrade reconciliation may use it: surviving partial orders require
// an available-balance adjustment of oldHold-newHold before switching formulas.
func LegacyIsolatedLimitHoldForMigration(order *DerivativeLimitOrder, market DerivativeMarketI, feeRate math.LegacyDec) (math.LegacyDec, error) {
	return canonicalIsolatedLimitHold(order, market, feeRate, true)
}

//nolint:revive // Share validation, fee arithmetic and conversion with the upgrade's legacy margin reader.
func canonicalIsolatedLimitHold(order *DerivativeLimitOrder, market DerivativeMarketI, feeRate math.LegacyDec, legacyMargin bool) (math.LegacyDec, error) {
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
		marginRaw = proportionalOrderMarginRaw(order.Margin, order.Fillable, order.OrderInfo.Quantity)
		if legacyMargin {
			marginRaw = proportionalPositionMarginRaw(order.Margin, order.Fillable, order.OrderInfo.Quantity)
		}
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
