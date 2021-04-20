package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
	default:
		return OracleType_Band, sdkerrors.Wrapf(ErrUnsupportedOracleType, "%s", oracleTypeStr)
	}
	return oracleType, nil
}
