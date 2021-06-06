package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
)

const QuoteUSD = "USD"
const TwapWindow = int64(5 * 60) // 5 minute TWAP window

func GetOracleType(oracleTypeStr string) (OracleType, error) {
	oracleTypeStr = strings.ToLower(oracleTypeStr)
	var oracleType OracleType

	switch oracleTypeStr {
	case "band":
		oracleType = OracleType_Band
	case "pricefeed":
		oracleType = OracleType_PriceFeed
	case "coinbase":
		oracleType = OracleType_Coinbase
	default:
		return OracleType_Band, sdkerrors.Wrapf(ErrUnsupportedOracleType, "%s", oracleTypeStr)
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

func (p *PriceState) UpdatePrice(price sdk.Dec, timestamp int64) {
	cumulativePriceDelta := sdk.NewDec(timestamp - p.Timestamp).Mul(p.Price)
	p.CumulativePrice = p.CumulativePrice.Add(cumulativePriceDelta)
	p.Timestamp = timestamp
	p.Price = price
}
