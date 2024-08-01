package types

import (
	"cosmossdk.io/math"
)

func NewStorkPriceState(
	price math.LegacyDec,
	timestamp uint64,
	symbol string,
	blockTime int64,
) *StorkPriceState {
	return &StorkPriceState{
		Timestamp:  timestamp,
		Symbol:     symbol,
		Value:      price,
		PriceState: *NewPriceState(price, blockTime),
	}
}

func (s *StorkPriceState) Update(price math.LegacyDec, timestamp uint64, blockTime int64) {
	s.Value = price
	s.Timestamp = timestamp
	s.PriceState.UpdatePrice(price, blockTime)
}

func ScaleStorkPrice(price math.LegacyDec) math.LegacyDec {
	return price.QuoTruncate(EighteenDecimals)
}
