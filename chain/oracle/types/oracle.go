package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
	"strings"
)

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
