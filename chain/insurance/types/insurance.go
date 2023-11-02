package types

import (
	"fmt"
	"time"

	"cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common"

	oracletypes "github.com/InjectiveLabs/sdk-go/chain/oracle/types"
)

var InsuranceFundInitialSupply = math.NewIntWithDecimal(1, 18)

func NewInsuranceFund(
	marketID common.Hash,
	depositDenom, poolTokenDenom string,
	redemptionNoticePeriodDuration time.Duration,
	ticker, oracleBase, oracleQuote string, oracleType oracletypes.OracleType, expiry int64,
) *InsuranceFund {

	return &InsuranceFund{
		DepositDenom:                   depositDenom,
		InsurancePoolTokenDenom:        poolTokenDenom,
		RedemptionNoticePeriodDuration: redemptionNoticePeriodDuration,
		Balance:                        math.ZeroInt(),
		TotalShare:                     math.ZeroInt(),
		MarketId:                       marketID.Hex(),
		MarketTicker:                   ticker,
		OracleBase:                     oracleBase,
		OracleQuote:                    oracleQuote,
		OracleType:                     oracleType,
		Expiry:                         expiry,
	}

}

func (fund InsuranceFund) ShareDenom() string {
	return fund.InsurancePoolTokenDenom
}

func (fund *InsuranceFund) AddTotalShare(shares math.Int) {
	fund.TotalShare = fund.TotalShare.Add(shares)
}

func (fund *InsuranceFund) SubTotalShare(shares math.Int) {
	fund.TotalShare = fund.TotalShare.Sub(shares)
}

func ShareDenomFromId(id uint64) string {
	return fmt.Sprintf("share%d", id)
}
