package types

import (
	fmt "fmt"
	"time"

	oracletypes "github.com/InjectiveLabs/sdk-go/chain/oracle/types"
	"github.com/ethereum/go-ethereum/common"

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

func (fund *InsuranceFund) AddTotalShare(amt sdk.Int) {
	fund.TotalShare = fund.TotalShare.Add(amt)
}

func (fund *InsuranceFund) SubTotalShare(amt sdk.Int) {
	fund.TotalShare = fund.TotalShare.Sub(amt)
}

func ShareDenomFromId(id uint64) string {
	return fmt.Sprintf("share%d", id)
}
