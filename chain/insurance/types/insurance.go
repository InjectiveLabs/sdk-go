package types

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"

	oracletypes "github.com/InjectiveLabs/sdk-go/chain/oracle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var InsuranceFundInitialSupply = sdk.NewIntWithDecimal(1, 18)

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
		Balance:                        sdk.ZeroInt(),
		TotalShare:                     sdk.ZeroInt(),
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

func (fund *InsuranceFund) AddTotalShare(shares sdk.Int) {
	fund.TotalShare = fund.TotalShare.Add(shares)
}

func (fund *InsuranceFund) SubTotalShare(shares sdk.Int) {
	fund.TotalShare = fund.TotalShare.Sub(shares)
}

func ShareDenomFromId(id uint64) string {
	return fmt.Sprintf("share%d", id)
}
