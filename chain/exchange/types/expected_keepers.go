package types

import (
	"context"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"

	downtimetypes "github.com/InjectiveLabs/sdk-go/chain/downtime-detector/types"
	insurancetypes "github.com/InjectiveLabs/sdk-go/chain/insurance/types"
	oracletypes "github.com/InjectiveLabs/sdk-go/chain/oracle/types"
	wasmxtypes "github.com/InjectiveLabs/sdk-go/chain/wasmx/types"
)

// BankKeeper defines the expected bank keeper methods.
type BankKeeper interface {
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	GetSupply(ctx sdk.Context) sdk.Coin
}

// OracleKeeper defines the expected oracle keeper methods.
type OracleKeeper interface {
	GetPrice(ctx sdk.Context, oracletype oracletypes.OracleType, base string, quote string) *sdkmath.LegacyDec
	GetPricePairState(ctx sdk.Context, oracletype oracletypes.OracleType, base, quote string, scalingOptions *oracletypes.ScalingOptions) *oracletypes.PricePairState
	GetCumulativePrice(ctx sdk.Context, oracleType oracletypes.OracleType, base string, quote string) *sdkmath.LegacyDec
	GetHistoricalPriceRecords(ctx sdk.Context, oracleType oracletypes.OracleType, symbol string, from int64) (entry *oracletypes.PriceRecords, omitted bool)
	GetMixedHistoricalPriceRecords(ctx sdk.Context, baseOracleType, quoteOracleType oracletypes.OracleType, baseSymbol, quoteSymbol string, from int64) (mixed *oracletypes.PriceRecords, ok bool)
	GetStandardDeviationForPriceRecords(priceRecords []*oracletypes.PriceRecord) *sdkmath.LegacyDec
	GetProviderInfo(ctx sdk.Context, provider string) *oracletypes.ProviderInfo
	GetProviderPrice(ctx sdk.Context, provider, symbol string) *sdkmath.LegacyDec
	GetProviderPriceState(ctx sdk.Context, provider, symbol string) *oracletypes.ProviderPriceState
}

// InsuranceKeeper defines the expected insurance keeper methods.
type InsuranceKeeper interface {
	// HasInsuranceFund returns true if InsuranceFund for the given marketID exists.
	HasInsuranceFund(ctx sdk.Context, marketID common.Hash) bool
	// GetInsuranceFund returns the insurance fund corresponding to the given marketID.
	GetInsuranceFund(ctx sdk.Context, marketID common.Hash) *insurancetypes.InsuranceFund
	// DepositIntoInsuranceFund increments the insurance fund balance by amount.
	DepositIntoInsuranceFund(ctx sdk.Context, marketID common.Hash, amount sdkmath.Int) error
	// WithdrawFromInsuranceFund decrements the insurance fund balance by amount and sends
	WithdrawFromInsuranceFund(ctx sdk.Context, marketID common.Hash, amount sdkmath.Int) error
	// UpdateInsuranceFundOracleParams updates the insurance fund's oracle parameters
	UpdateInsuranceFundOracleParams(ctx sdk.Context, marketID common.Hash, oracleParams *OracleParams) error
}

type GovKeeper interface {
	IterateActiveProposalsQueue(ctx context.Context, endTime time.Time, cb func(proposal v1.Proposal) (stop bool))
	GetParams(ctx context.Context) v1.Params
}

type DistributionKeeper interface {
	GetFeePool(ctx sdk.Context) (feePool types.FeePool)
	DistributeFromFeePool(ctx sdk.Context, amount sdk.Coins, receiveAddr sdk.AccAddress) error
	FundCommunityPool(ctx sdk.Context, amount sdk.Coins, sender sdk.AccAddress) error
}

type StakingKeeper interface {
	GetDelegatorDelegations(ctx context.Context, delegator sdk.AccAddress, maxRetrieve uint16) (delegations []stakingtypes.Delegation, err error)
	Validator(context.Context, sdk.ValAddress) (stakingtypes.ValidatorI, error) // get a particular validator by operator address
}

type WasmViewKeeper interface {
	wasmtypes.ViewKeeper
}

type WasmContractOpsKeeper interface {
	wasmtypes.ContractOpsKeeper
}

type WasmxExecutionKeeper interface {
	InjectiveExec(ctx sdk.Context, contractAddress sdk.AccAddress, funds sdk.Coins, msg *wasmxtypes.InjectiveExecMsg) ([]byte, error)
}

type DowntimeKeeper interface {
	GetLastDowntimeOfLength(ctx sdk.Context, dur downtimetypes.Downtime) (time.Time, error)
	GetLastBlockTime(ctx sdk.Context) (time.Time, error)
}
