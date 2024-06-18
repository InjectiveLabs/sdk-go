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
		Timestamp: timestamp,
		Symbol: symbol,
		Value: price,
		PriceState:  *NewPriceState(price, blockTime),
	}
}
