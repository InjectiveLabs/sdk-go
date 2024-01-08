package core

import "github.com/shopspring/decimal"

type Token struct {
	Name     string
	Symbol   string
	Denom    string
	Address  string
	Decimals int32
	Logo     string
	Updated  int64
}

func (t Token) ChainFormattedValue(humanReadableValue decimal.Decimal) decimal.Decimal {
	multiplier := decimal.New(1, t.Decimals)
	return humanReadableValue.Mul(multiplier)
}
