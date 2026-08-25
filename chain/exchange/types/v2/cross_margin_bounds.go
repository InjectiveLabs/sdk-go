package v2

import (
	"fmt"
	"math/big"

	sdkmath "cosmossdk.io/math"

	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
)

const (
	// MaxCrossMarginDerivativeCarriers is the hard per-subaccount cardinality
	// bound over positions, persistent derivative orders, and all current-block
	// transient derivative indicators.
	MaxCrossMarginDerivativeCarriers uint32 = 64
	// MaxCrossMarginDerivativeCarriersPerQuote is the hard per-quote bound for
	// non-binary derivative carriers.
	MaxCrossMarginDerivativeCarriersPerQuote uint32 = 8
	// MaxCrossMarginCancellableOrdersPerQuote is the hard cleanup-work bound over
	// derivative orders plus spot orders locking the same quote denom.
	MaxCrossMarginCancellableOrdersPerQuote uint32 = 32
	// MaxCrossMarginCanonicalChunksPerQuote is the hard RFQ canonical chunk bound.
	MaxCrossMarginCanonicalChunksPerQuote uint32 = 32
)

var (
	maxOrderQuantityRaw        = new(big.Int).Set(exchangetypes.MaxOrderQuantity.BigInt())
	legacyPrecisionRaw         = new(big.Int).Exp(big.NewInt(10), big.NewInt(sdkmath.LegacyPrecision), nil)
	legacyHalfPrecisionRaw     = new(big.Int).Quo(new(big.Int).Set(legacyPrecisionRaw), big.NewInt(2))
	legacyDecRangeExclusiveRaw = new(big.Int).Mul(
		new(big.Int).Lsh(big.NewInt(1), 256),
		new(big.Int).Set(legacyPrecisionRaw),
	)
)

// CrossMarginCanonicalChunkCountRaw computes
// ceil((abs(position)+queuedVanilla)/MaxOrderQuantity) using raw integers only.
// It never constructs an intermediate LegacyDec, so an adversarial aggregate
// cannot overflow before the hard bound is checked.
func CrossMarginCanonicalChunkCountRaw(positionRaw, queuedVanillaRaw *big.Int) (*big.Int, error) {
	if positionRaw == nil || queuedVanillaRaw == nil {
		return nil, fmt.Errorf("cross-margin chunk input is nil")
	}
	if queuedVanillaRaw.Sign() < 0 {
		return nil, fmt.Errorf("cross-margin queued vanilla quantity is negative")
	}

	numerator := new(big.Int).Abs(new(big.Int).Set(positionRaw))
	numerator.Add(numerator, queuedVanillaRaw)
	if numerator.Sign() == 0 {
		return new(big.Int), nil
	}

	quotient, remainder := new(big.Int), new(big.Int)
	quotient.QuoRem(numerator, maxOrderQuantityRaw, remainder)
	if remainder.Sign() != 0 {
		quotient.Add(quotient, big.NewInt(1))
	}
	return quotient, nil
}

// CrossMarginAtomicPositionEnvelopeRaw returns
// max(abs(position), abs(position+signedRequestedQuantity)) in raw units.
func CrossMarginAtomicPositionEnvelopeRaw(positionRaw, signedRequestedQuantityRaw *big.Int) (*big.Int, error) {
	if positionRaw == nil || signedRequestedQuantityRaw == nil {
		return nil, fmt.Errorf("cross-margin atomic envelope input is nil")
	}
	before := new(big.Int).Abs(new(big.Int).Set(positionRaw))
	after := new(big.Int).Add(new(big.Int).Set(positionRaw), signedRequestedQuantityRaw)
	after.Abs(after)
	if after.Cmp(before) > 0 {
		return after, nil
	}
	return before, nil
}

// CrossMarginRawDec copies the raw fixed-point integer without performing
// arithmetic in LegacyDec.
func CrossMarginRawDec(value sdkmath.LegacyDec) (*big.Int, error) {
	if value.IsNil() {
		return nil, fmt.Errorf("cross-margin decimal is nil")
	}
	return new(big.Int).Set(value.BigInt()), nil
}

func checkedLegacyMulRaw(a, b sdkmath.LegacyDec) (*big.Int, error) {
	if a.IsNil() || b.IsNil() {
		return nil, fmt.Errorf("nil decimal multiplication input")
	}
	product := new(big.Int).Mul(a.BigInt(), b.BigInt())
	sign := product.Sign()
	product.Abs(product)
	quotient, remainder := new(big.Int), new(big.Int)
	quotient.QuoRem(product, legacyPrecisionRaw, remainder)
	if cmp := remainder.Cmp(legacyHalfPrecisionRaw); cmp > 0 || cmp == 0 && quotient.Bit(0) == 1 {
		quotient.Add(quotient, big.NewInt(1))
	}
	if sign < 0 {
		quotient.Neg(quotient)
	}
	if new(big.Int).Abs(new(big.Int).Set(quotient)).Cmp(legacyDecRangeExclusiveRaw) >= 0 {
		return nil, fmt.Errorf("decimal multiplication is out of range")
	}
	return quotient, nil
}

func checkedLegacyAddRaw(a, b *big.Int) (*big.Int, error) {
	if a == nil || b == nil {
		return nil, fmt.Errorf("nil decimal addition input")
	}
	sum := new(big.Int).Add(new(big.Int).Set(a), b)
	if new(big.Int).Abs(new(big.Int).Set(sum)).Cmp(legacyDecRangeExclusiveRaw) >= 0 {
		return nil, fmt.Errorf("decimal addition is out of range")
	}
	return sum, nil
}

// CanonicalCrossMarginSpotLimitHold returns the chain-format physical hold for
// a live spot limit order. GetUnfilledMarginHoldAndMarginDenom deliberately
// performs notional and fee operations in the runtime order; callers must not
// replace this with an algebraically collapsed product.
func CanonicalCrossMarginSpotLimitHold(
	order *SpotLimitOrder,
	market *SpotMarket,
	isTransient bool,
) (sdkmath.LegacyDec, string, error) {
	if order == nil || market == nil {
		return sdkmath.LegacyDec{}, "", fmt.Errorf("missing canonical spot limit hold input")
	}
	if order.Fillable.IsNil() || order.Fillable.IsNegative() || !order.Fillable.IsInValidRange() {
		return sdkmath.LegacyDec{}, "", fmt.Errorf("invalid canonical spot limit fillable")
	}
	if order.IsBuy() {
		feeRate := market.MakerFeeRate
		if isTransient {
			feeRate = market.TakerFeeRate
		}
		if feeRate.IsNil() || !feeRate.IsInValidRange() {
			return sdkmath.LegacyDec{}, "", fmt.Errorf("invalid canonical spot buy hold input")
		}
		if !isTransient {
			feeRate = sdkmath.LegacyMaxDec(sdkmath.LegacyZeroDec(), feeRate)
		}
		if order.OrderInfo.Price.IsNil() || order.OrderInfo.Price.IsNegative() ||
			!order.OrderInfo.Price.IsInValidRange() ||
			market.QuoteDenom == "" || market.QuoteDecimals > exchangetypes.MaxDecimals {
			return sdkmath.LegacyDec{}, "", fmt.Errorf("invalid canonical spot buy hold input")
		}
		notionalRaw, err := checkedLegacyMulRaw(order.Fillable, order.OrderInfo.Price)
		if err != nil {
			return sdkmath.LegacyDec{}, "", fmt.Errorf("canonical spot buy notional: %w", err)
		}
		notional := sdkmath.LegacyNewDecFromBigIntWithPrec(notionalRaw, sdkmath.LegacyPrecision)
		feeRaw, err := checkedLegacyMulRaw(notional, feeRate)
		if err != nil {
			return sdkmath.LegacyDec{}, "", fmt.Errorf("canonical spot buy fee: %w", err)
		}
		if _, err := checkedLegacyAddRaw(notionalRaw, feeRaw); err != nil {
			return sdkmath.LegacyDec{}, "", fmt.Errorf("canonical spot buy hold: %w", err)
		}
	} else if market.BaseDenom == "" || market.BaseDecimals > exchangetypes.MaxDecimals {
		return sdkmath.LegacyDec{}, "", fmt.Errorf("invalid canonical spot sell hold input")
	}
	hold, denom := order.GetUnfilledMarginHoldAndMarginDenom(market, isTransient)
	if hold.IsNil() || hold.IsNegative() {
		return sdkmath.LegacyDec{}, "", fmt.Errorf("invalid canonical spot limit hold")
	}
	var chainHold sdkmath.LegacyDec
	if order.IsBuy() {
		if !exchangetypes.CanRepresentNotionalInChainFormat(hold, market.QuoteDecimals) {
			return sdkmath.LegacyDec{}, "", fmt.Errorf("canonical spot limit hold is out of chain-format range")
		}
		chainHold = market.NotionalToChainFormat(hold)
	} else {
		if !exchangetypes.CanRepresentNotionalInChainFormat(hold, market.BaseDecimals) {
			return sdkmath.LegacyDec{}, "", fmt.Errorf("canonical spot limit hold is out of chain-format range")
		}
		chainHold = market.QuantityToChainFormat(hold)
	}
	if chainHold.IsNil() || chainHold.IsNegative() || !chainHold.IsInValidRange() {
		return sdkmath.LegacyDec{}, "", fmt.Errorf("canonical spot limit hold is out of range")
	}
	return chainHold, denom, nil
}

// CanonicalCrossMarginSpotMarketHold returns the chain-format physical hold
// persisted in a queued transient spot market order.
func CanonicalCrossMarginSpotMarketHold(
	order *SpotMarketOrder,
	market *SpotMarket,
) (sdkmath.LegacyDec, string, error) {
	if order == nil || market == nil || order.BalanceHold.IsNil() || order.BalanceHold.IsNegative() ||
		!order.BalanceHold.IsInValidRange() {
		return sdkmath.LegacyDec{}, "", fmt.Errorf("invalid canonical spot market hold input")
	}
	denom := market.BaseDenom
	var chainHold sdkmath.LegacyDec
	if order.IsBuy() {
		if market.QuoteDenom == "" || market.QuoteDecimals > exchangetypes.MaxDecimals {
			return sdkmath.LegacyDec{}, "", fmt.Errorf("invalid canonical spot market buy input")
		}
		denom = market.QuoteDenom
		if !exchangetypes.CanRepresentNotionalInChainFormat(order.BalanceHold, market.QuoteDecimals) {
			return sdkmath.LegacyDec{}, "", fmt.Errorf("canonical spot market hold is out of chain-format range")
		}
		chainHold = market.NotionalToChainFormat(order.BalanceHold)
	} else {
		if market.BaseDenom == "" || market.BaseDecimals > exchangetypes.MaxDecimals {
			return sdkmath.LegacyDec{}, "", fmt.Errorf("invalid canonical spot market sell input")
		}
		if !exchangetypes.CanRepresentNotionalInChainFormat(order.BalanceHold, market.BaseDecimals) {
			return sdkmath.LegacyDec{}, "", fmt.Errorf("canonical spot market hold is out of chain-format range")
		}
		chainHold = market.QuantityToChainFormat(order.BalanceHold)
	}
	if chainHold.IsNil() || chainHold.IsNegative() || !chainHold.IsInValidRange() {
		return sdkmath.LegacyDec{}, "", fmt.Errorf("canonical spot market hold is out of range")
	}
	return chainHold, denom, nil
}
