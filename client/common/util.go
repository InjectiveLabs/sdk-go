package common

import (
	"encoding/hex"
	"strings"

	"github.com/shopspring/decimal"
)

func HexToBytes(str string) ([]byte, error) {
	str = strings.TrimPrefix(str, "0x")

	data, err := hex.DecodeString(str)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func RemoveExtraDecimals(value decimal.Decimal, decimalsToRemove int32) decimal.Decimal {
	return value.Div(decimal.New(1, decimalsToRemove))
}
