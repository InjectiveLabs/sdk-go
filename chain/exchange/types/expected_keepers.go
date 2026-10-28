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
	GetReferencePrice(ctx sdk.Context, oracletype oracletypes.OracleType, base, quote string) *sdkmath.LegacyDec
	GetPricePairState(
		ctx sdk.Context,
		oracletype oracletypes.OracleType,
		base, quote string, scalingOptions *oracletypes.ScalingOptions,
	) *oracletypes.PricePairState
	GetCumulativePrice(
		ctx sdk.Context,
		oracleType oracletypes.OracleType,
		base, quote string,
	) (baseCumulative, quoteCumulative *sdkmath.LegacyDec)
	GetProviderInfo(ctx sdk.Context, provider string) *oracletypes.ProviderInfo
	GetProviderPrice(ctx sdk.Context, provider, symbol string) *sdkmath.LegacyDec
	// GetPythPriceState exposes the raw per-priceID Pyth state (spot price,
	// EMA price, confidence, publish time) for consumers that need more than
	// the spot pairing — the reference-price gate's PYTH_EMA_LONG source reads
	// EmaPrice and PublishTime.
	GetPythPriceState(ctx sdk.Context, priceID common.Hash) *oracletypes.PythPriceState
	GetPythProPriceState(ctx sdk.Context, feedID uint32) *oracletypes.PythProPriceState
	GetChainlinkDataStreamsPriceState(ctx sdk.Context, feedID string) *oracletypes.ChainlinkDataStreamsPriceState
	GetSedaFastPriceState(ctx sdk.Context, feedID string) *oracletypes.SedaFastPriceState
}

// InsuranceKeeper defines the expected insurance keeper methods.
type InsuranceKeeper interface {
	// HasInsuranceFund returns true if InsuranceFund for the given marketID exists.
	HasInsuranceFund(ctx sdk.Context, marketID common.Hash) bool
	// GetInsuranceFund returns the insurance fund corresponding to the given marketID.
	GetInsuranceFund(ctx sdk.Context, marketID common.Hash) *insurancetypes.InsuranceFund
	GetActiveInsuranceFund(ctx sdk.Context, marketID common.Hash) *insurancetypes.InsuranceFund
	HasActiveInsuranceFund(ctx sdk.Context, marketID common.Hash) bool
	// DepositIntoInsuranceFund increments the insurance fund balance by amount.
	DepositIntoInsuranceFund(ctx sdk.Context, marketID common.Hash, amount sdkmath.Int) error
	// WithdrawFromInsuranceFund decrements the insurance fund balance by amount and sends
	WithdrawFromInsuranceFund(ctx sdk.Context, marketID common.Hash, amount sdkmath.Int) error
	// UpdateInsuranceFundOracleParams updates the insurance fund's oracle parameters
	UpdateInsuranceFundOracleParams(ctx sdk.Context, marketID common.Hash, oracleParams *OracleParams) error
	HasConsolidatedInsurancePool(ctx sdk.Context, denom string) bool
	HasConsolidatedInsurancePolicy(ctx sdk.Context, marketID common.Hash) bool
	GetConsolidatedInsurancePool(ctx sdk.Context, denom string) *insurancetypes.ConsolidatedInsurancePool
	GetConsolidatedInsurancePolicy(ctx sdk.Context, marketID common.Hash) *insurancetypes.ConsolidatedInsurancePolicy
	EnrollConsolidatedInsuranceMarket(ctx sdk.Context, marketID common.Hash, denom string, quoteDecimals uint32) error
	MigrateLegacyInsuranceMarketToConsolidated(ctx sdk.Context, marketID common.Hash, denom string, quoteDecimals uint32) error
	GetConsolidatedInsuranceAvailable(ctx sdk.Context, marketID common.Hash) sdkmath.Int
	GetConsolidatedInsuranceCapacity(
		ctx sdk.Context,
		marketID common.Hash,
	) *insurancetypes.ConsolidatedInsuranceCapacity
	ValidateConsolidatedInsuranceClaims(ctx sdk.Context, marketIDs []common.Hash, requested []sdkmath.Int) error
	RecordConsolidatedMarketReservation(ctx sdk.Context, marketID common.Hash, balance sdkmath.LegacyDec)
	ProcessConsolidatedReservations(ctx sdk.Context)
	SetConsolidatedCollateralCursor(ctx sdk.Context, marketID common.Hash, cursor []byte)
	ReleaseConsolidatedSettlementReservation(ctx sdk.Context, marketID common.Hash, remainingDeficit sdkmath.Int) error
	ValidateConsolidatedMarketReservations(ctx sdk.Context, marketIDs []common.Hash, balances []sdkmath.LegacyDec) error
	ValidateConsolidatedReservationState(ctx sdk.Context) error
	WithdrawFromConsolidatedInsurancePool(ctx sdk.Context, marketID common.Hash, requested sdkmath.Int) (sdkmath.Int, error)
	WithdrawFromConsolidatedInsurancePoolBatch(
		ctx sdk.Context,
		marketIDs []common.Hash,
		requested []sdkmath.Int,
	) ([]sdkmath.Int, error)
	WithdrawConsolidatedInsuranceForSettlement(
		ctx sdk.Context, marketIDs []common.Hash, requested []sdkmath.Int,
	) ([]sdkmath.Int, error)
	CreditConsolidatedInsurancePoolForMarket(ctx sdk.Context, marketID common.Hash, amount sdkmath.Int) error
	// GetConsolidatedNetPrefunding and AdjustConsolidatedNetPrefunding track pool
	// cash that a live market ledger paid ahead of losing positions' debts.
	GetConsolidatedNetPrefunding(ctx sdk.Context, marketID common.Hash) sdkmath.Int
	AdjustConsolidatedNetPrefunding(ctx sdk.Context, marketID common.Hash, delta sdkmath.Int) error
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
	SetDelegationTransferReceiver(ctx context.Context, receiver sdk.AccAddress)
}

type WasmViewKeeper interface {
	wasmtypes.ViewKeeper
}

type WasmContractOpsKeeper interface {
	wasmtypes.ContractOpsKeeper
}

type DowntimeKeeper interface {
	GetLastDowntimeOfLength(ctx sdk.Context, dur downtimetypes.Downtime) (time.Time, error)
	GetLastBlockTime(ctx sdk.Context) (time.Time, error)
}

type PermissionsKeeper interface {
	IsEnforcedRestrictionsDenom(ctx sdk.Context, denom string) bool
	SendRestrictionFn(ctx context.Context, fromAddr, toAddr sdk.AccAddress, amount sdk.Coin) (newToAddr sdk.AccAddress, err error)
}
