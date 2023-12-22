package core

import (
	"github.com/huandu/go-assert"
	"github.com/shopspring/decimal"
	"testing"
)

func createINJToken() Token {
	token := Token{
		Name:     "Injective Protocol",
		Symbol:   "INJ",
		Denom:    "inj",
		Address:  "0xe28b3B32B6c345A34Ff64674606124Dd5Aceca30",
		Decimals: 18,
		Logo:     "https://static.alchemyapi.io/images/assets/7226.png",
		Updated:  1681739137644,
	}

	return token
}

func createUSDTToken() Token {
	token := Token{
		Name:     "USDT",
		Symbol:   "USDT",
		Denom:    "peggy0x87aB3B4C8661e07D6372361211B96ed4Dc36B1B5",
		Address:  "0x0000000000000000000000000000000000000000",
		Decimals: 6,
		Logo:     "https://static.alchemyapi.io/images/assets/825.png",
		Updated:  1681739137645,
	}

	return token
}

func TestChainFormattedValue(t *testing.T) {
	value := decimal.RequireFromString("1.3456")
	token := createINJToken()

	chainFormattedValue := token.ChainFormattedValue(value)
	multiplier := decimal.New(1, int32(token.Decimals))
	expectedValue := value.Mul(multiplier)

	assert.Equal(t, chainFormattedValue, expectedValue)
}
