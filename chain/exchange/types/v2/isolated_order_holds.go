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
	marginRaw := proportionalPositionMarginRaw(order.Margin, order.Fillable, order.OrderInfo.Quantity)
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
	if market.GetMarketType().IsBinaryOptions() {
		required, err := CanonicalBinaryOptionsLimitExecutionHold(order, market, feeRate)
		if err != nil {
			return math.LegacyDec{}, err
		}
		if hold.LT(required) {
			return math.LegacyDec{}, errors.New("binary options remaining hold does not cover execution collateral")
		}
	}
	return hold, nil
}

// CanonicalBinaryOptionsLimitExecutionHold returns the checked chain-format
// collateral and positive fees required to fill the remainder at its limit
// price. It does not assume proportional order margin covers that remainder.
//
//nolint:revive // Keep checked BO collateral arithmetic and input guards together.
func CanonicalBinaryOptionsLimitExecutionHold(order *DerivativeLimitOrder, market DerivativeMarketI, feeRate math.LegacyDec) (math.LegacyDec, error) {
	if order == nil || market == nil || !market.GetMarketType().IsBinaryOptions() || market.GetOracleScaleFactor() > types.MaxOracleScaleFactor {
		return math.LegacyDec{}, errors.New("invalid binary options execution hold input")
	}
	for _, value := range []math.LegacyDec{order.Fillable, order.OrderInfo.Quantity, order.OrderInfo.Price, order.Margin, feeRate} {
		if value.IsNil() || !value.IsInValidRange() {
			return math.LegacyDec{}, errors.New("invalid binary options execution hold input")
		}
	}
	maxPrice := types.GetScaledPrice(math.LegacyOneDec(), market.GetOracleScaleFactor())
	if order.Fillable.IsNegative() || !order.OrderInfo.Quantity.IsPositive() || order.Fillable.GT(order.OrderInfo.Quantity) ||
		order.OrderInfo.Price.IsNegative() || order.OrderInfo.Price.GT(maxPrice) || order.Margin.IsNegative() {
		return math.LegacyDec{}, errors.New("invalid binary options execution hold quantity, price or margin")
	}
	if !order.IsVanilla() {
		return math.LegacyZeroDec(), nil
	}
	marginPrice := order.OrderInfo.Price
	if !order.IsBuy() {
		marginPrice = maxPrice.Sub(marginPrice)
	}
	marginRaw, err := checkedLegacyMulRaw(order.Fillable, marginPrice)
	if err != nil {
		return math.LegacyDec{}, err
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
	return CanonicalIsolatedMarketHold(math.LegacyNewDecFromBigIntWithPrec(holdRaw, math.LegacyPrecision), market)
}

// CanonicalIsolatedMarketHold converts a stored human-format MarginHold once.
func CanonicalIsolatedMarketHold(hold math.LegacyDec, market DerivativeMarketI) (math.LegacyDec, error) {
	if market == nil || hold.IsNil() || hold.IsNegative() || !hold.IsInValidRange() ||
		market.GetQuoteDecimals() > types.MaxDecimals || !types.CanRepresentNotionalInChainFormat(hold, market.GetQuoteDecimals()) {
		return math.LegacyDec{}, errors.New("invalid or unrepresentable isolated derivative hold")
	}
	return market.NotionalToChainFormat(hold), nil
}
