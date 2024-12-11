package types

import (
	"cosmossdk.io/math"
)

type FeeDiscountRates struct {
	MakerDiscountRate math.LegacyDec
	TakerDiscountRate math.LegacyDec
}

type FeeDiscountRatesMap map[uint64]*FeeDiscountRates
