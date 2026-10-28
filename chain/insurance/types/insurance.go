package types

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common"

	oracletypes "github.com/InjectiveLabs/sdk-go/chain/oracle/types"
)

var (
	InsuranceFundInitialSupply                = math.NewIntWithDecimal(1, 18)
	InsuranceFundProtocolOwnedLiquiditySupply = math.NewIntWithDecimal(1, 16)
	InsuranceFundCreatorSupply                = InsuranceFundInitialSupply.Sub(InsuranceFundProtocolOwnedLiquiditySupply)
)

// ConsolidatedInsuranceCapacity separates a market's policy headroom from
// the shared pool cash that may back several markets.
type ConsolidatedInsuranceCapacity struct {
	PoolDenom      string
	RemainingLimit math.Int
	PoolBalance    math.Int
	// Reserved is this market's commitment; PoolBalance excludes all commitments.
	Reserved math.Int
}

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

// ParseShareDenomID recognizes exactly the share denomination namespace owned
// by the insurance module. Leading-zero aliases are deliberately excluded.
func ParseShareDenomID(denom string) (uint64, bool) {
	if !strings.HasPrefix(denom, "share") || len(denom) == len("share") {
		return 0, false
	}
	id, err := strconv.ParseUint(denom[len("share"):], 10, 64)
	if err != nil || ShareDenomFromId(id) != denom {
		return 0, false
	}
	return id, true
}
