package types

import (
	"strings"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
)

const QuoteUSD = "USD"
const TwapWindow = int64(5 * 60)              // 5 minute TWAP window
const BandPriceMultiplier uint64 = 1000000000 // 1e9

// MaxHistoricalPriceRecordAge is the maximum age of oracle price records to track.
const MaxHistoricalPriceRecordAge = 60 * 5

func GetOracleType(oracleTypeStr string) (OracleType, error) {
	oracleTypeStr = strings.ToLower(oracleTypeStr)
	var oracleType OracleType

	switch oracleTypeStr {
	case "band":
		oracleType = OracleType_Band
	case "bandibc":
		oracleType = OracleType_BandIBC
	case "pricefeed":
		oracleType = OracleType_PriceFeed
	case "coinbase":
		oracleType = OracleType_Coinbase
	case "provider":
		oracleType = OracleType_Provider
	case "pyth":
		oracleType = OracleType_Pyth
	default:
		return OracleType_Band, errors.Wrapf(ErrUnsupportedOracleType, "%s", oracleTypeStr)
	}
	return oracleType, nil
}

func (o *OracleType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(OracleType_value, data, "OracleType")
	if err != nil {
		return err
	}
	*o = OracleType(value)
	return nil
}

func (c *CoinbasePriceState) GetDecPrice() sdk.Dec {
	// nolint:all
	// price = price/10^6
	return sdk.NewDec(int64(c.Value)).QuoTruncate(sdk.NewDec(10).Power(6))
}

func NewPriceState(price sdk.Dec, timestamp int64) *PriceState {
	return &PriceState{
		Price:           price,
		CumulativePrice: sdk.ZeroDec(),
		Timestamp:       timestamp,
	}
}

func NewProviderPriceState(symbol string, price sdk.Dec, timestamp int64) *ProviderPriceState {
	return &ProviderPriceState{
		Symbol: symbol,
		State:  NewPriceState(price, timestamp),
	}
}

func (p *PriceState) UpdatePrice(price sdk.Dec, timestamp int64) {
	cumulativePriceDelta := sdk.NewDec(timestamp - p.Timestamp).Mul(p.Price)
	p.CumulativePrice = p.CumulativePrice.Add(cumulativePriceDelta)
	p.Timestamp = timestamp
	p.Price = price
}

type SymbolPriceTimestamps []*SymbolPriceTimestamp

func (s SymbolPriceTimestamps) SetTimestamp(oracleType OracleType, symbol string, ts int64) SymbolPriceTimestamps {
	for _, entry := range s {
		if entry.SymbolId == symbol {
			entry.Timestamp = ts
			return s
		}
	}

	s = append(s, &SymbolPriceTimestamp{
		Oracle:    oracleType,
		SymbolId:  symbol,
		Timestamp: ts,
	})

	return s
}

func (s SymbolPriceTimestamps) GetTimestamp(oracleType OracleType, symbol string) (ts int64, ok bool) {
	for i := range s {
		if s[i].Oracle == oracleType && s[i].SymbolId == symbol {
			return s[i].Timestamp, true
		}
	}

	return -1, false
}

// CheckPriceFeedThreshold returns true if the newPrice has changed beyond 100x or less than 1% of the last price
func CheckPriceFeedThreshold(lastPrice, newPrice sdk.Dec) bool {
	return newPrice.GT(lastPrice.Mul(sdk.NewDec(100))) || newPrice.LT(lastPrice.Quo(sdk.NewDec(100)))
}

func IsLegacySchemeOracleScript(scriptID int64, params BandIBCParams) bool {
	for _, id := range params.LegacyOracleIds {
		if id == scriptID {
			return true
		}
	}

	return false
}
