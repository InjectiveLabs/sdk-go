package chain

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	sdkmath "cosmossdk.io/math"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	log "github.com/InjectiveLabs/suplog"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	"github.com/cosmos/cosmos-sdk/client/tx"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/gogoproto/proto"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	ibcconnectiontypes "github.com/cosmos/ibc-go/v8/modules/core/03-connection/types"
	ibcchanneltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	exchangev2types "github.com/InjectiveLabs/sdk-go/chain/exchange/types/v2"
	permissionstypes "github.com/InjectiveLabs/sdk-go/chain/permissions/types"
	chainstreamtypes "github.com/InjectiveLabs/sdk-go/chain/stream/types"
	chainstreamv2types "github.com/InjectiveLabs/sdk-go/chain/stream/types/v2"
	tokenfactorytypes "github.com/InjectiveLabs/sdk-go/chain/tokenfactory/types"
	txfeestypes "github.com/InjectiveLabs/sdk-go/chain/txfees/types"
	"github.com/InjectiveLabs/sdk-go/client"
	"github.com/InjectiveLabs/sdk-go/client/common"
)

type OrderbookType string

const (
	msgCommitBatchSizeLimit          = 1024
	msgCommitBatchTimeLimit          = 500 * time.Millisecond
	defaultBroadcastStatusPoll       = 100 * time.Millisecond
	defaultBroadcastTimeout          = 40 * time.Second
	defaultTimeoutHeight             = 20
	defaultTimeoutHeightSyncInterval = 10 * time.Second
	SpotOrderbook                    = "injective.exchange.v1beta1.EventOrderbookUpdate.spot_orderbooks"
	DerivativeOrderbook              = "injective.exchange.v1beta1.EventOrderbookUpdate.derivative_orderbooks"
)

var (
	ErrTimedOut       = errors.New("tx timed out")
	ErrQueueClosed    = errors.New("queue is closed")
	ErrEnqueueTimeout = errors.New("enqueue timeout")
	ErrReadOnly       = errors.New("client is in read-only mode")
)

type ChainClient interface {
	CanSignTransactions() bool
	FromAddress() sdk.AccAddress
	QueryClient() *grpc.ClientConn
	ClientContext() sdkclient.Context
	// return account number and sequence without increasing sequence
	GetAccNonce() (accNum uint64, accSeq uint64)

	SimulateMsg(clientCtx sdkclient.Context, msgs ...sdk.Msg) (*txtypes.SimulateResponse, error)
	AsyncBroadcastMsg(msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error)
	SyncBroadcastMsg(msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error)
	BroadcastMsg(broadcastMode txtypes.BroadcastMode, msgs ...sdk.Msg) (*txtypes.BroadcastTxRequest, *txtypes.BroadcastTxResponse, error)

	// Build signed tx with given accNum and accSeq, useful for offline siging
	// If simulate is set to false, initialGas will be used
	BuildSignedTx(clientCtx sdkclient.Context, accNum, accSeq, initialGas uint64, gasPrice uint64, msg ...sdk.Msg) ([]byte, error)
	SyncBroadcastSignedTx(tyBytes []byte) (*txtypes.BroadcastTxResponse, error)
	AsyncBroadcastSignedTx(txBytes []byte) (*txtypes.BroadcastTxResponse, error)
	BroadcastSignedTx(txBytes []byte, broadcastMode txtypes.BroadcastMode) (*txtypes.BroadcastTxResponse, error)
	QueueBroadcastMsg(msgs ...sdk.Msg) error

	// Bank Module
	GetBankBalances(ctx context.Context, address string) (*banktypes.QueryAllBalancesResponse, error)
	GetBankBalance(ctx context.Context, address string, denom string) (*banktypes.QueryBalanceResponse, error)
	GetBankSpendableBalances(ctx context.Context, address string, pagination *query.PageRequest) (*banktypes.QuerySpendableBalancesResponse, error)
	GetBankSpendableBalancesByDenom(ctx context.Context, address string, denom string) (*banktypes.QuerySpendableBalanceByDenomResponse, error)
	GetBankTotalSupply(ctx context.Context, pagination *query.PageRequest) (*banktypes.QueryTotalSupplyResponse, error)
	GetBankSupplyOf(ctx context.Context, denom string) (*banktypes.QuerySupplyOfResponse, error)
	GetDenomMetadata(ctx context.Context, denom string) (*banktypes.QueryDenomMetadataResponse, error)
	GetDenomsMetadata(ctx context.Context, pagination *query.PageRequest) (*banktypes.QueryDenomsMetadataResponse, error)
	GetDenomOwners(ctx context.Context, denom string, pagination *query.PageRequest) (*banktypes.QueryDenomOwnersResponse, error)
	GetBankSendEnabled(ctx context.Context, denoms []string, pagination *query.PageRequest) (*banktypes.QuerySendEnabledResponse, error)

	GetAuthzGrants(ctx context.Context, req authztypes.QueryGrantsRequest) (*authztypes.QueryGrantsResponse, error)
	GetAccount(ctx context.Context, address string) (*authtypes.QueryAccountResponse, error)

	BuildGenericAuthz(granter string, grantee string, msgtype string, expireIn time.Time) *authztypes.MsgGrant
	BuildExchangeAuthz(granter string, grantee string, authzType ExchangeAuthz, subaccountId string, markets []string, expireIn time.Time) *authztypes.MsgGrant
	BuildExchangeBatchUpdateOrdersAuthz(
		granter string,
		grantee string,
		subaccountId string,
		spotMarkets []string,
		derivativeMarkets []string,
		expireIn time.Time,
	) *authztypes.MsgGrant

	DefaultSubaccount(acc sdk.AccAddress) ethcommon.Hash
	Subaccount(account sdk.AccAddress, index int) ethcommon.Hash

	// Deprecated: use GetSubAccountNonceV2 instead
	GetSubAccountNonce(ctx context.Context, subaccountId ethcommon.Hash) (*exchangetypes.QuerySubaccountTradeNonceResponse, error)
	// Deprecated: use GetFeeDiscountInfoV2 instead
	GetFeeDiscountInfo(ctx context.Context, account string) (*exchangetypes.QueryFeeDiscountAccountInfoResponse, error)

	UpdateSubaccountNonceFromChain() error
	SynchronizeSubaccountNonce(subaccountId ethcommon.Hash) error
	ComputeOrderHashes(spotOrders []exchangev2types.SpotOrder, derivativeOrders []exchangev2types.DerivativeOrder, subaccountId ethcommon.Hash) (OrderHashes, error)

	// Deprecated: use CreateSpotOrderV2 instead
	CreateSpotOrder(defaultSubaccountID ethcommon.Hash, d *SpotOrderData, marketsAssistant MarketsAssistant) *exchangetypes.SpotOrder
	// Deprecated: use CreateDerivativeOrderV2 instead
	CreateDerivativeOrder(defaultSubaccountID ethcommon.Hash, d *DerivativeOrderData, marketAssistant MarketsAssistant) *exchangetypes.DerivativeOrder
	// Deprecated: use OrderCancelV2 instead
	OrderCancel(defaultSubaccountID ethcommon.Hash, d *OrderCancelData) *exchangetypes.OrderData

	CreateSpotOrderV2(defaultSubaccountID ethcommon.Hash, d *SpotOrderData) *exchangev2types.SpotOrder
	CreateDerivativeOrderV2(defaultSubaccountID ethcommon.Hash, d *DerivativeOrderData) *exchangev2types.DerivativeOrder
	OrderCancelV2(defaultSubaccountID ethcommon.Hash, d *OrderCancelData) *exchangev2types.OrderData

	GetGasFee() (string, error)

	StreamEventOrderFail(sender string, failEventCh chan map[string]uint)
	StreamEventOrderFailWithWebsocket(sender string, websocket *rpchttp.HTTP, failEventCh chan map[string]uint)
	// Deprecated: use the chain stream instead
	StreamOrderbookUpdateEvents(orderbookType OrderbookType, marketIDs []string, orderbookCh chan exchangetypes.Orderbook)
	// Deprecated: use the chain stream instead
	StreamOrderbookUpdateEventsWithWebsocket(orderbookType OrderbookType, marketIDs []string, websocket *rpchttp.HTTP, orderbookCh chan exchangetypes.Orderbook)

	// Deprecated: use ChainStreamV2 instead
	ChainStream(ctx context.Context, req chainstreamtypes.StreamRequest) (chainstreamtypes.Stream_StreamClient, error)
	ChainStreamV2(ctx context.Context, req chainstreamv2types.StreamRequest) (chainstreamv2types.Stream_StreamV2Client, error)

	// get tx from chain node
	GetTx(ctx context.Context, txHash string) (*txtypes.GetTxResponse, error)

	// wasm module
	FetchContractInfo(ctx context.Context, address string) (*wasmtypes.QueryContractInfoResponse, error)
	FetchContractHistory(ctx context.Context, address string, pagination *query.PageRequest) (*wasmtypes.QueryContractHistoryResponse, error)
	FetchContractsByCode(ctx context.Context, codeId uint64, pagination *query.PageRequest) (*wasmtypes.QueryContractsByCodeResponse, error)
	FetchAllContractsState(ctx context.Context, address string, pagination *query.PageRequest) (*wasmtypes.QueryAllContractStateResponse, error)
	RawContractState(
		ctx context.Context,
		contractAddress string,
		queryData []byte,
	) (*wasmtypes.QueryRawContractStateResponse, error)
	SmartContractState(
		ctx context.Context,
		contractAddress string,
		queryData []byte,
	) (*wasmtypes.QuerySmartContractStateResponse, error)
	FetchCode(ctx context.Context, codeId uint64) (*wasmtypes.QueryCodeResponse, error)
	FetchCodes(ctx context.Context, pagination *query.PageRequest) (*wasmtypes.QueryCodesResponse, error)
	FetchPinnedCodes(ctx context.Context, pagination *query.PageRequest) (*wasmtypes.QueryPinnedCodesResponse, error)
	FetchContractsByCreator(ctx context.Context, creator string, pagination *query.PageRequest) (*wasmtypes.QueryContractsByCreatorResponse, error)

	// tokenfactory module
	FetchDenomAuthorityMetadata(ctx context.Context, creator string, subDenom string) (*tokenfactorytypes.QueryDenomAuthorityMetadataResponse, error)
	FetchDenomsFromCreator(ctx context.Context, creator string) (*tokenfactorytypes.QueryDenomsFromCreatorResponse, error)
	FetchTokenfactoryModuleState(ctx context.Context) (*tokenfactorytypes.QueryModuleStateResponse, error)

	// distribution module
	FetchValidatorDistributionInfo(ctx context.Context, validatorAddress string) (*distributiontypes.QueryValidatorDistributionInfoResponse, error)
	FetchValidatorOutstandingRewards(ctx context.Context, validatorAddress string) (*distributiontypes.QueryValidatorOutstandingRewardsResponse, error)
	FetchValidatorCommission(ctx context.Context, validatorAddress string) (*distributiontypes.QueryValidatorCommissionResponse, error)
	FetchValidatorSlashes(ctx context.Context, validatorAddress string, startingHeight uint64, endingHeight uint64, pagination *query.PageRequest) (*distributiontypes.QueryValidatorSlashesResponse, error)
	FetchDelegationRewards(ctx context.Context, delegatorAddress string, validatorAddress string) (*distributiontypes.QueryDelegationRewardsResponse, error)
	FetchDelegationTotalRewards(ctx context.Context, delegatorAddress string) (*distributiontypes.QueryDelegationTotalRewardsResponse, error)
	FetchDelegatorValidators(ctx context.Context, delegatorAddress string) (*distributiontypes.QueryDelegatorValidatorsResponse, error)
	FetchDelegatorWithdrawAddress(ctx context.Context, delegatorAddress string) (*distributiontypes.QueryDelegatorWithdrawAddressResponse, error)
	FetchCommunityPool(ctx context.Context) (*distributiontypes.QueryCommunityPoolResponse, error)

	// chain exchange module

	// Deprecated: use FetchSubaccountDepositsV2 instead
	FetchSubaccountDeposits(ctx context.Context, subaccountID string) (*exchangetypes.QuerySubaccountDepositsResponse, error)
	// Deprecated: use FetchSubaccountDepositV2 instead
	FetchSubaccountDeposit(ctx context.Context, subaccountId string, denom string) (*exchangetypes.QuerySubaccountDepositResponse, error)
	// Deprecated: use FetchExchangeBalancesV2 instead
	FetchExchangeBalances(ctx context.Context) (*exchangetypes.QueryExchangeBalancesResponse, error)
	// Deprecated: use FetchAggregateVolumeV2 instead
	FetchAggregateVolume(ctx context.Context, account string) (*exchangetypes.QueryAggregateVolumeResponse, error)
	// Deprecated: use FetchAggregateVolumesV2 instead
	FetchAggregateVolumes(ctx context.Context, accounts []string, marketIDs []string) (*exchangetypes.QueryAggregateVolumesResponse, error)
	// Deprecated: use FetchAggregateMarketVolumeV2 instead
	FetchAggregateMarketVolume(ctx context.Context, marketId string) (*exchangetypes.QueryAggregateMarketVolumeResponse, error)
	// Deprecated: use FetchAggregateMarketVolumesV2 instead
	FetchAggregateMarketVolumes(ctx context.Context, marketIDs []string) (*exchangetypes.QueryAggregateMarketVolumesResponse, error)
	// Deprecated: use FetchDenomDecimalV2 instead
	FetchDenomDecimal(ctx context.Context, denom string) (*exchangetypes.QueryDenomDecimalResponse, error)
	// Deprecated: use FetchDenomDecimalsV2 instead
	FetchDenomDecimals(ctx context.Context, denoms []string) (*exchangetypes.QueryDenomDecimalsResponse, error)
	// Deprecated: use FetchChainSpotMarketsV2 instead
	FetchChainSpotMarkets(ctx context.Context, status string, marketIDs []string) (*exchangetypes.QuerySpotMarketsResponse, error)
	// Deprecated: use FetchChainSpotMarketV2 instead
	FetchChainSpotMarket(ctx context.Context, marketId string) (*exchangetypes.QuerySpotMarketResponse, error)
	// Deprecated: use FetchChainFullSpotMarketsV2 instead
	FetchChainFullSpotMarkets(ctx context.Context, status string, marketIDs []string, withMidPriceAndTob bool) (*exchangetypes.QueryFullSpotMarketsResponse, error)
	// Deprecated: use FetchChainFullSpotMarketV2 instead
	FetchChainFullSpotMarket(ctx context.Context, marketId string, withMidPriceAndTob bool) (*exchangetypes.QueryFullSpotMarketResponse, error)
	// Deprecated: use FetchChainSpotOrderbookV2 instead
	FetchChainSpotOrderbook(ctx context.Context, marketId string, limit uint64, orderSide exchangetypes.OrderSide, limitCumulativeNotional sdkmath.LegacyDec, limitCumulativeQuantity sdkmath.LegacyDec) (*exchangetypes.QuerySpotOrderbookResponse, error)
	// Deprecated: use FetchChainTraderSpotOrdersV2 instead
	FetchChainTraderSpotOrders(ctx context.Context, marketId string, subaccountId string) (*exchangetypes.QueryTraderSpotOrdersResponse, error)
	// Deprecated: use FetchChainAccountAddressSpotOrdersV2 instead
	FetchChainAccountAddressSpotOrders(ctx context.Context, marketId string, address string) (*exchangetypes.QueryAccountAddressSpotOrdersResponse, error)
	// Deprecated: use FetchChainSpotOrdersByHashesV2 instead
	FetchChainSpotOrdersByHashes(ctx context.Context, marketId string, subaccountId string, orderHashes []string) (*exchangetypes.QuerySpotOrdersByHashesResponse, error)
	// Deprecated: use FetchChainSubaccountOrdersV2 instead
	FetchChainSubaccountOrders(ctx context.Context, subaccountId string, marketId string) (*exchangetypes.QuerySubaccountOrdersResponse, error)
	// Deprecated: use FetchChainTraderSpotTransientOrdersV2 instead
	FetchChainTraderSpotTransientOrders(ctx context.Context, marketId string, subaccountId string) (*exchangetypes.QueryTraderSpotOrdersResponse, error)
	// Deprecated: use FetchSpotMidPriceAndTOBV2 instead
	FetchSpotMidPriceAndTOB(ctx context.Context, marketId string) (*exchangetypes.QuerySpotMidPriceAndTOBResponse, error)
	// Deprecated: use FetchDerivativeMidPriceAndTOBV2 instead
	FetchDerivativeMidPriceAndTOB(ctx context.Context, marketId string) (*exchangetypes.QueryDerivativeMidPriceAndTOBResponse, error)
	// Deprecated: use FetchChainDerivativeOrderbookV2 instead
	FetchChainDerivativeOrderbook(ctx context.Context, marketId string, limit uint64, limitCumulativeNotional sdkmath.LegacyDec) (*exchangetypes.QueryDerivativeOrderbookResponse, error)
	// Deprecated: use FetchChainTraderDerivativeOrdersV2 instead
	FetchChainTraderDerivativeOrders(ctx context.Context, marketId string, subaccountId string) (*exchangetypes.QueryTraderDerivativeOrdersResponse, error)
	// Deprecated: use FetchChainAccountAddressDerivativeOrdersV2 instead
	FetchChainAccountAddressDerivativeOrders(ctx context.Context, marketId string, address string) (*exchangetypes.QueryAccountAddressDerivativeOrdersResponse, error)
	// Deprecated: use FetchChainDerivativeOrdersByHashesV2 instead
	FetchChainDerivativeOrdersByHashes(ctx context.Context, marketId string, subaccountId string, orderHashes []string) (*exchangetypes.QueryDerivativeOrdersByHashesResponse, error)
	// Deprecated: use FetchChainTraderDerivativeTransientOrdersV2 instead
	FetchChainTraderDerivativeTransientOrders(ctx context.Context, marketId string, subaccountId string) (*exchangetypes.QueryTraderDerivativeOrdersResponse, error)
	// Deprecated: use FetchChainDerivativeMarketV2 instead
	FetchChainDerivativeMarkets(ctx context.Context, status string, marketIDs []string, withMidPriceAndTob bool) (*exchangetypes.QueryDerivativeMarketsResponse, error)
	// Deprecated: use FetchChainDerivativeMarketV2 instead
	FetchChainDerivativeMarket(ctx context.Context, marketId string) (*exchangetypes.QueryDerivativeMarketResponse, error)
	// Deprecated: use FetchDerivativeMarketAddressV2 instead
	FetchDerivativeMarketAddress(ctx context.Context, marketId string) (*exchangetypes.QueryDerivativeMarketAddressResponse, error)
	// Deprecated: use FetchSubaccountTradeNonceV2 instead
	FetchSubaccountTradeNonce(ctx context.Context, subaccountId string) (*exchangetypes.QuerySubaccountTradeNonceResponse, error)
	// Deprecated: use FetchChainPositionsV2 instead
	FetchChainPositions(ctx context.Context) (*exchangetypes.QueryPositionsResponse, error)
	// Deprecated: use FetchChainSubaccountPositionsV2 instead
	FetchChainSubaccountPositions(ctx context.Context, subaccountId string) (*exchangetypes.QuerySubaccountPositionsResponse, error)
	// Deprecated: use FetchChainSubaccountPositionInMarketV2 instead
	FetchChainSubaccountPositionInMarket(ctx context.Context, subaccountId string, marketId string) (*exchangetypes.QuerySubaccountPositionInMarketResponse, error)
	// Deprecated: use FetchChainSubaccountEffectivePositionInMarketV2 instead
	FetchChainSubaccountEffectivePositionInMarket(ctx context.Context, subaccountId string, marketId string) (*exchangetypes.QuerySubaccountEffectivePositionInMarketResponse, error)
	// Deprecated: use FetchChainPerpetualMarketInfoV2 instead
	FetchChainPerpetualMarketInfo(ctx context.Context, marketId string) (*exchangetypes.QueryPerpetualMarketInfoResponse, error)
	// Deprecated: use FetchChainExpiryFuturesMarketInfoV2 instead
	FetchChainExpiryFuturesMarketInfo(ctx context.Context, marketId string) (*exchangetypes.QueryExpiryFuturesMarketInfoResponse, error)
	// Deprecated: use FetchChainPerpetualMarketFundingV2 instead
	FetchChainPerpetualMarketFunding(ctx context.Context, marketId string) (*exchangetypes.QueryPerpetualMarketFundingResponse, error)
	// Deprecated: use FetchSubaccountOrderMetadataV2 instead
	FetchSubaccountOrderMetadata(ctx context.Context, subaccountId string) (*exchangetypes.QuerySubaccountOrderMetadataResponse, error)
	// Deprecated: use FetchTradeRewardPointsV2 instead
	FetchTradeRewardPoints(ctx context.Context, accounts []string) (*exchangetypes.QueryTradeRewardPointsResponse, error)
	// Deprecated: use FetchPendingTradeRewardPointsV2 instead
	FetchPendingTradeRewardPoints(ctx context.Context, accounts []string) (*exchangetypes.QueryTradeRewardPointsResponse, error)
	// Deprecated: use FetchFeeDiscountAccountInfoV2 instead
	FetchFeeDiscountAccountInfo(ctx context.Context, account string) (*exchangetypes.QueryFeeDiscountAccountInfoResponse, error)
	// Deprecated: use FetchTradeRewardCampaignV2 instead
	FetchTradeRewardCampaign(ctx context.Context) (*exchangetypes.QueryTradeRewardCampaignResponse, error)
	// Deprecated: use FetchFeeDiscountScheduleV2 instead
	FetchFeeDiscountSchedule(ctx context.Context) (*exchangetypes.QueryFeeDiscountScheduleResponse, error)
	// Deprecated: use FetchBalanceMismatchesV2 instead
	FetchBalanceMismatches(ctx context.Context, dustFactor int64) (*exchangetypes.QueryBalanceMismatchesResponse, error)
	// Deprecated: use FetchBalanceWithBalanceHoldsV2 instead
	FetchBalanceWithBalanceHolds(ctx context.Context) (*exchangetypes.QueryBalanceWithBalanceHoldsResponse, error)
	// Deprecated: use FetchFeeDiscountTierStatisticsV2 instead
	FetchFeeDiscountTierStatistics(ctx context.Context) (*exchangetypes.QueryFeeDiscountTierStatisticsResponse, error)
	// Deprecated: use FetchMitoVaultInfosV2 instead
	FetchMitoVaultInfos(ctx context.Context) (*exchangetypes.MitoVaultInfosResponse, error)
	// Deprecated: use FetchMarketIDFromVaultV2 instead
	FetchMarketIDFromVault(ctx context.Context, vaultAddress string) (*exchangetypes.QueryMarketIDFromVaultResponse, error)
	// Deprecated: use FetchHistoricalTradeRecordsV2 instead
	FetchHistoricalTradeRecords(ctx context.Context, marketId string) (*exchangetypes.QueryHistoricalTradeRecordsResponse, error)
	// Deprecated: use FetchIsOptedOutOfRewardsV2 instead
	FetchIsOptedOutOfRewards(ctx context.Context, account string) (*exchangetypes.QueryIsOptedOutOfRewardsResponse, error)
	// Deprecated: use FetchOptedOutOfRewardsAccountsV2 instead
	FetchOptedOutOfRewardsAccounts(ctx context.Context) (*exchangetypes.QueryOptedOutOfRewardsAccountsResponse, error)
	// Deprecated: use FetchMarketVolatilityV2 instead
	FetchMarketVolatility(ctx context.Context, marketId string, tradeHistoryOptions *exchangetypes.TradeHistoryOptions) (*exchangetypes.QueryMarketVolatilityResponse, error)
	// Deprecated: use FetchChainBinaryOptionsMarketsV2 instead
	FetchChainBinaryOptionsMarkets(ctx context.Context, status string) (*exchangetypes.QueryBinaryMarketsResponse, error)
	// Deprecated: use FetchTraderDerivativeConditionalOrdersV2 instead
	FetchTraderDerivativeConditionalOrders(ctx context.Context, subaccountId string, marketId string) (*exchangetypes.QueryTraderDerivativeConditionalOrdersResponse, error)
	// Deprecated: use FetchMarketAtomicExecutionFeeMultiplierV2 instead
	FetchMarketAtomicExecutionFeeMultiplier(ctx context.Context, marketId string) (*exchangetypes.QueryMarketAtomicExecutionFeeMultiplierResponse, error)
	// Deprecated: use FetchL3DerivativeOrderbookV2 instead
	FetchL3DerivativeOrderBook(ctx context.Context, marketId string) (*exchangetypes.QueryFullDerivativeOrderbookResponse, error)
	// Deprecated: use FetchL3SpotOrderbookV2 instead
	FetchL3SpotOrderBook(ctx context.Context, marketId string) (*exchangetypes.QueryFullSpotOrderbookResponse, error)
	// Deprecated: use FetchMarketBalanceV2 instead
	FetchMarketBalance(ctx context.Context, marketId string) (*exchangetypes.QueryMarketBalanceResponse, error)
	// Deprecated: use FetchMarketBalancesV2 instead
	FetchMarketBalances(ctx context.Context) (*exchangetypes.QueryMarketBalancesResponse, error)
	// Deprecated: use FetchDenomMinNotionalV2 instead
	FetchDenomMinNotional(ctx context.Context, denom string) (*exchangetypes.QueryDenomMinNotionalResponse, error)
	// Deprecated: use FetchDenomMinNotionalsV2 instead
	FetchDenomMinNotionals(ctx context.Context) (*exchangetypes.QueryDenomMinNotionalsResponse, error)

	// chain exchange v2 module
	FetchSubaccountDepositsV2(ctx context.Context, subaccountID string) (*exchangev2types.QuerySubaccountDepositsResponse, error)
	FetchSubaccountDepositV2(ctx context.Context, subaccountId string, denom string) (*exchangev2types.QuerySubaccountDepositResponse, error)
	FetchExchangeBalancesV2(ctx context.Context) (*exchangev2types.QueryExchangeBalancesResponse, error)
	FetchAggregateVolumeV2(ctx context.Context, account string) (*exchangev2types.QueryAggregateVolumeResponse, error)
	FetchAggregateVolumesV2(ctx context.Context, accounts []string, marketIDs []string) (*exchangev2types.QueryAggregateVolumesResponse, error)
	FetchAggregateMarketVolumeV2(ctx context.Context, marketId string) (*exchangev2types.QueryAggregateMarketVolumeResponse, error)
	FetchAggregateMarketVolumesV2(ctx context.Context, marketIDs []string) (*exchangev2types.QueryAggregateMarketVolumesResponse, error)
	FetchDenomDecimalV2(ctx context.Context, denom string) (*exchangev2types.QueryDenomDecimalResponse, error)
	FetchDenomDecimalsV2(ctx context.Context, denoms []string) (*exchangev2types.QueryDenomDecimalsResponse, error)
	FetchChainSpotMarketsV2(ctx context.Context, status string, marketIDs []string) (*exchangev2types.QuerySpotMarketsResponse, error)
	FetchChainSpotMarketV2(ctx context.Context, marketId string) (*exchangev2types.QuerySpotMarketResponse, error)
	FetchChainFullSpotMarketsV2(ctx context.Context, status string, marketIDs []string, withMidPriceAndTob bool) (*exchangev2types.QueryFullSpotMarketsResponse, error)
	FetchChainFullSpotMarketV2(ctx context.Context, marketId string, withMidPriceAndTob bool) (*exchangev2types.QueryFullSpotMarketResponse, error)
	FetchChainSpotOrderbookV2(ctx context.Context, marketId string, limit uint64, orderSide exchangev2types.OrderSide, limitCumulativeNotional sdkmath.LegacyDec, limitCumulativeQuantity sdkmath.LegacyDec) (*exchangev2types.QuerySpotOrderbookResponse, error)
	FetchChainTraderSpotOrdersV2(ctx context.Context, marketId string, subaccountId string) (*exchangev2types.QueryTraderSpotOrdersResponse, error)
	FetchChainAccountAddressSpotOrdersV2(ctx context.Context, marketId string, address string) (*exchangev2types.QueryAccountAddressSpotOrdersResponse, error)
	FetchChainSpotOrdersByHashesV2(ctx context.Context, marketId string, subaccountId string, orderHashes []string) (*exchangev2types.QuerySpotOrdersByHashesResponse, error)
	FetchChainSubaccountOrdersV2(ctx context.Context, subaccountId string, marketId string) (*exchangev2types.QuerySubaccountOrdersResponse, error)
	FetchChainTraderSpotTransientOrdersV2(ctx context.Context, marketId string, subaccountId string) (*exchangev2types.QueryTraderSpotOrdersResponse, error)
	FetchSpotMidPriceAndTOBV2(ctx context.Context, marketId string) (*exchangev2types.QuerySpotMidPriceAndTOBResponse, error)
	FetchDerivativeMidPriceAndTOBV2(ctx context.Context, marketId string) (*exchangev2types.QueryDerivativeMidPriceAndTOBResponse, error)
	FetchChainDerivativeOrderbookV2(ctx context.Context, marketId string, limit uint64, limitCumulativeNotional sdkmath.LegacyDec) (*exchangev2types.QueryDerivativeOrderbookResponse, error)
	FetchChainTraderDerivativeOrdersV2(ctx context.Context, marketId string, subaccountId string) (*exchangev2types.QueryTraderDerivativeOrdersResponse, error)
	FetchChainAccountAddressDerivativeOrdersV2(ctx context.Context, marketId string, address string) (*exchangev2types.QueryAccountAddressDerivativeOrdersResponse, error)
	FetchChainDerivativeOrdersByHashesV2(ctx context.Context, marketId string, subaccountId string, orderHashes []string) (*exchangev2types.QueryDerivativeOrdersByHashesResponse, error)
	FetchChainTraderDerivativeTransientOrdersV2(ctx context.Context, marketId string, subaccountId string) (*exchangev2types.QueryTraderDerivativeOrdersResponse, error)
	FetchChainDerivativeMarketsV2(ctx context.Context, status string, marketIDs []string, withMidPriceAndTob bool) (*exchangev2types.QueryDerivativeMarketsResponse, error)
	FetchChainDerivativeMarketV2(ctx context.Context, marketId string) (*exchangev2types.QueryDerivativeMarketResponse, error)
	FetchDerivativeMarketAddressV2(ctx context.Context, marketId string) (*exchangev2types.QueryDerivativeMarketAddressResponse, error)
	FetchSubaccountTradeNonceV2(ctx context.Context, subaccountId string) (*exchangev2types.QuerySubaccountTradeNonceResponse, error)
	FetchChainPositionsV2(ctx context.Context) (*exchangev2types.QueryPositionsResponse, error)
	FetchChainSubaccountPositionsV2(ctx context.Context, subaccountId string) (*exchangev2types.QuerySubaccountPositionsResponse, error)
	FetchChainSubaccountPositionInMarketV2(ctx context.Context, subaccountId string, marketId string) (*exchangev2types.QuerySubaccountPositionInMarketResponse, error)
	FetchChainSubaccountEffectivePositionInMarketV2(ctx context.Context, subaccountId string, marketId string) (*exchangev2types.QuerySubaccountEffectivePositionInMarketResponse, error)
	FetchChainPerpetualMarketInfoV2(ctx context.Context, marketId string) (*exchangev2types.QueryPerpetualMarketInfoResponse, error)
	FetchChainExpiryFuturesMarketInfoV2(ctx context.Context, marketId string) (*exchangev2types.QueryExpiryFuturesMarketInfoResponse, error)
	FetchChainPerpetualMarketFundingV2(ctx context.Context, marketId string) (*exchangev2types.QueryPerpetualMarketFundingResponse, error)
	FetchSubaccountOrderMetadataV2(ctx context.Context, subaccountId string) (*exchangev2types.QuerySubaccountOrderMetadataResponse, error)
	FetchTradeRewardPointsV2(ctx context.Context, accounts []string) (*exchangev2types.QueryTradeRewardPointsResponse, error)
	FetchPendingTradeRewardPointsV2(ctx context.Context, accounts []string) (*exchangev2types.QueryTradeRewardPointsResponse, error)
	FetchFeeDiscountAccountInfoV2(ctx context.Context, account string) (*exchangev2types.QueryFeeDiscountAccountInfoResponse, error)
	FetchTradeRewardCampaignV2(ctx context.Context) (*exchangev2types.QueryTradeRewardCampaignResponse, error)
	FetchFeeDiscountScheduleV2(ctx context.Context) (*exchangev2types.QueryFeeDiscountScheduleResponse, error)
	FetchBalanceMismatchesV2(ctx context.Context, dustFactor int64) (*exchangev2types.QueryBalanceMismatchesResponse, error)
	FetchBalanceWithBalanceHoldsV2(ctx context.Context) (*exchangev2types.QueryBalanceWithBalanceHoldsResponse, error)
	FetchFeeDiscountTierStatisticsV2(ctx context.Context) (*exchangev2types.QueryFeeDiscountTierStatisticsResponse, error)
	FetchMitoVaultInfosV2(ctx context.Context) (*exchangev2types.MitoVaultInfosResponse, error)
	FetchMarketIDFromVaultV2(ctx context.Context, vaultAddress string) (*exchangev2types.QueryMarketIDFromVaultResponse, error)
	FetchHistoricalTradeRecordsV2(ctx context.Context, marketId string) (*exchangev2types.QueryHistoricalTradeRecordsResponse, error)
	FetchIsOptedOutOfRewardsV2(ctx context.Context, account string) (*exchangev2types.QueryIsOptedOutOfRewardsResponse, error)
	FetchOptedOutOfRewardsAccountsV2(ctx context.Context) (*exchangev2types.QueryOptedOutOfRewardsAccountsResponse, error)
	FetchMarketVolatilityV2(ctx context.Context, marketId string, tradeHistoryOptions *exchangev2types.TradeHistoryOptions) (*exchangev2types.QueryMarketVolatilityResponse, error)
	FetchChainBinaryOptionsMarketsV2(ctx context.Context, status string) (*exchangev2types.QueryBinaryMarketsResponse, error)
	FetchTraderDerivativeConditionalOrdersV2(ctx context.Context, subaccountId string, marketId string) (*exchangev2types.QueryTraderDerivativeConditionalOrdersResponse, error)
	FetchMarketAtomicExecutionFeeMultiplierV2(ctx context.Context, marketId string) (*exchangev2types.QueryMarketAtomicExecutionFeeMultiplierResponse, error)
	FetchActiveStakeGrantV2(ctx context.Context, grantee string) (*exchangev2types.QueryActiveStakeGrantResponse, error)
	FetchGrantAuthorizationV2(ctx context.Context, granter, grantee string) (*exchangev2types.QueryGrantAuthorizationResponse, error)
	FetchGrantAuthorizationsV2(ctx context.Context, granter string) (*exchangev2types.QueryGrantAuthorizationsResponse, error)
	FetchL3DerivativeOrderbookV2(ctx context.Context, marketId string) (*exchangev2types.QueryFullDerivativeOrderbookResponse, error)
	FetchL3SpotOrderbookV2(ctx context.Context, marketId string) (*exchangev2types.QueryFullSpotOrderbookResponse, error)
	FetchMarketBalanceV2(ctx context.Context, marketId string) (*exchangev2types.QueryMarketBalanceResponse, error)
	FetchMarketBalancesV2(ctx context.Context) (*exchangev2types.QueryMarketBalancesResponse, error)
	FetchDenomMinNotionalV2(ctx context.Context, denom string) (*exchangev2types.QueryDenomMinNotionalResponse, error)
	FetchDenomMinNotionalsV2(ctx context.Context) (*exchangev2types.QueryDenomMinNotionalsResponse, error)

	// Tendermint module
	FetchNodeInfo(ctx context.Context) (*cmtservice.GetNodeInfoResponse, error)
	FetchSyncing(ctx context.Context) (*cmtservice.GetSyncingResponse, error)
	FetchLatestBlock(ctx context.Context) (*cmtservice.GetLatestBlockResponse, error)
	FetchBlockByHeight(ctx context.Context, height int64) (*cmtservice.GetBlockByHeightResponse, error)
	FetchLatestValidatorSet(ctx context.Context) (*cmtservice.GetLatestValidatorSetResponse, error)
	FetchValidatorSetByHeight(ctx context.Context, height int64, pagination *query.PageRequest) (*cmtservice.GetValidatorSetByHeightResponse, error)
	ABCIQuery(ctx context.Context, path string, data []byte, height int64, prove bool) (*cmtservice.ABCIQueryResponse, error)

	// IBC Transfer module
	FetchDenomTrace(ctx context.Context, hash string) (*ibctransfertypes.QueryDenomTraceResponse, error)
	FetchDenomTraces(ctx context.Context, pagination *query.PageRequest) (*ibctransfertypes.QueryDenomTracesResponse, error)
	FetchDenomHash(ctx context.Context, trace string) (*ibctransfertypes.QueryDenomHashResponse, error)
	FetchEscrowAddress(ctx context.Context, portId string, channelId string) (*ibctransfertypes.QueryEscrowAddressResponse, error)
	FetchTotalEscrowForDenom(ctx context.Context, denom string) (*ibctransfertypes.QueryTotalEscrowForDenomResponse, error)

	// IBC Core Channel module
	FetchIBCChannel(ctx context.Context, portId string, channelId string) (*ibcchanneltypes.QueryChannelResponse, error)
	FetchIBCChannels(ctx context.Context, pagination *query.PageRequest) (*ibcchanneltypes.QueryChannelsResponse, error)
	FetchIBCConnectionChannels(ctx context.Context, connection string, pagination *query.PageRequest) (*ibcchanneltypes.QueryConnectionChannelsResponse, error)
	FetchIBCChannelClientState(ctx context.Context, portId string, channelId string) (*ibcchanneltypes.QueryChannelClientStateResponse, error)
	FetchIBCChannelConsensusState(ctx context.Context, portId string, channelId string, revisionNumber uint64, revisionHeight uint64) (*ibcchanneltypes.QueryChannelConsensusStateResponse, error)
	FetchIBCPacketCommitment(ctx context.Context, portId string, channelId string, sequence uint64) (*ibcchanneltypes.QueryPacketCommitmentResponse, error)
	FetchIBCPacketCommitments(ctx context.Context, portId string, channelId string, pagination *query.PageRequest) (*ibcchanneltypes.QueryPacketCommitmentsResponse, error)
	FetchIBCPacketReceipt(ctx context.Context, portId string, channelId string, sequence uint64) (*ibcchanneltypes.QueryPacketReceiptResponse, error)
	FetchIBCPacketAcknowledgement(ctx context.Context, portId string, channelId string, sequence uint64) (*ibcchanneltypes.QueryPacketAcknowledgementResponse, error)
	FetchIBCPacketAcknowledgements(ctx context.Context, portId string, channelId string, packetCommitmentSequences []uint64, pagination *query.PageRequest) (*ibcchanneltypes.QueryPacketAcknowledgementsResponse, error)
	FetchIBCUnreceivedPackets(ctx context.Context, portId string, channelId string, packetCommitmentSequences []uint64) (*ibcchanneltypes.QueryUnreceivedPacketsResponse, error)
	FetchIBCUnreceivedAcks(ctx context.Context, portId string, channelId string, packetAckSequences []uint64) (*ibcchanneltypes.QueryUnreceivedAcksResponse, error)
	FetchIBCNextSequenceReceive(ctx context.Context, portId string, channelId string) (*ibcchanneltypes.QueryNextSequenceReceiveResponse, error)

	// IBC Core Chain module
	FetchIBCClientState(ctx context.Context, clientId string) (*ibcclienttypes.QueryClientStateResponse, error)
	FetchIBCClientStates(ctx context.Context, pagination *query.PageRequest) (*ibcclienttypes.QueryClientStatesResponse, error)
	FetchIBCConsensusState(ctx context.Context, clientId string, revisionNumber uint64, revisionHeight uint64, latestHeight bool) (*ibcclienttypes.QueryConsensusStateResponse, error)
	FetchIBCConsensusStates(ctx context.Context, clientId string, pagination *query.PageRequest) (*ibcclienttypes.QueryConsensusStatesResponse, error)
	FetchIBCConsensusStateHeights(ctx context.Context, clientId string, pagination *query.PageRequest) (*ibcclienttypes.QueryConsensusStateHeightsResponse, error)
	FetchIBCClientStatus(ctx context.Context, clientId string) (*ibcclienttypes.QueryClientStatusResponse, error)
	FetchIBCClientParams(ctx context.Context) (*ibcclienttypes.QueryClientParamsResponse, error)
	FetchIBCUpgradedClientState(ctx context.Context) (*ibcclienttypes.QueryUpgradedClientStateResponse, error)
	FetchIBCUpgradedConsensusState(ctx context.Context) (*ibcclienttypes.QueryUpgradedConsensusStateResponse, error)

	// IBC Core Connection module
	FetchIBCConnection(ctx context.Context, connectionId string) (*ibcconnectiontypes.QueryConnectionResponse, error)
	FetchIBCConnections(ctx context.Context, pagination *query.PageRequest) (*ibcconnectiontypes.QueryConnectionsResponse, error)
	FetchIBCClientConnections(ctx context.Context, clientId string) (*ibcconnectiontypes.QueryClientConnectionsResponse, error)
	FetchIBCConnectionClientState(ctx context.Context, connectionId string) (*ibcconnectiontypes.QueryConnectionClientStateResponse, error)
	FetchIBCConnectionConsensusState(ctx context.Context, connectionId string, revisionNumber uint64, revisionHeight uint64) (*ibcconnectiontypes.QueryConnectionConsensusStateResponse, error)
	FetchIBCConnectionParams(ctx context.Context) (*ibcconnectiontypes.QueryConnectionParamsResponse, error)

	// Permissions module
	FetchPermissionsNamespaceDenoms(ctx context.Context) (*permissionstypes.QueryNamespaceDenomsResponse, error)
	FetchPermissionsNamespaces(ctx context.Context) (*permissionstypes.QueryNamespacesResponse, error)
	FetchPermissionsNamespace(ctx context.Context, denom string) (*permissionstypes.QueryNamespaceResponse, error)
	FetchPermissionsRolesByActor(ctx context.Context, denom, actor string) (*permissionstypes.QueryRolesByActorResponse, error)
	FetchPermissionsActorsByRole(ctx context.Context, denom, role string) (*permissionstypes.QueryActorsByRoleResponse, error)
	FetchPermissionsRoleManagers(ctx context.Context, denom string) (*permissionstypes.QueryRoleManagersResponse, error)
	FetchPermissionsRoleManager(ctx context.Context, denom, manager string) (*permissionstypes.QueryRoleManagerResponse, error)
	FetchPermissionsPolicyStatuses(ctx context.Context, denom string) (*permissionstypes.QueryPolicyStatusesResponse, error)
	FetchPermissionsPolicyManagerCapabilities(ctx context.Context, denom string) (*permissionstypes.QueryPolicyManagerCapabilitiesResponse, error)
	FetchPermissionsVouchers(ctx context.Context, denom string) (*permissionstypes.QueryVouchersResponse, error)
	FetchPermissionsVoucher(ctx context.Context, denom, address string) (*permissionstypes.QueryVoucherResponse, error)
	FetchPermissionsModuleState(ctx context.Context) (*permissionstypes.QueryModuleStateResponse, error)

	// TxFees module
	FetchTxFeesParams(ctx context.Context) (*txfeestypes.QueryParamsResponse, error)
	FetchEipBaseFee(ctx context.Context) (*txfeestypes.QueryEipBaseFeeResponse, error)

	CurrentChainGasPrice() int64
	SetGasPrice(gasPrice int64)

	GetNetwork() common.Network
	Close()
}

var _ ChainClient = &chainClient{}

type chainClient struct {
	ctx             sdkclient.Context
	network         common.Network
	opts            *common.ClientOptions
	logger          log.Logger
	conn            *grpc.ClientConn
	chainStreamConn *grpc.ClientConn
	txFactory       tx.Factory

	doneC   chan bool
	msgC    chan sdk.Msg
	syncMux *sync.Mutex

	cancelCtx context.Context
	cancelFn  func()

	accNum    uint64
	accSeq    uint64
	gasWanted uint64
	gasFee    string

	sessionEnabled bool

	ofacChecker *OfacChecker

	authQueryClient          authtypes.QueryClient
	authzQueryClient         authztypes.QueryClient
	bankQueryClient          banktypes.QueryClient
	chainStreamClient        chainstreamtypes.StreamClient
	chainStreamV2Client      chainstreamv2types.StreamClient
	distributionQueryClient  distributiontypes.QueryClient
	exchangeQueryClient      exchangetypes.QueryClient
	exchangeV2QueryClient    exchangev2types.QueryClient
	ibcChannelQueryClient    ibcchanneltypes.QueryClient
	ibcClientQueryClient     ibcclienttypes.QueryClient
	ibcConnectionQueryClient ibcconnectiontypes.QueryClient
	ibcTransferQueryClient   ibctransfertypes.QueryClient
	permissionsQueryClient   permissionstypes.QueryClient
	tendermintQueryClient    cmtservice.ServiceClient
	tokenfactoryQueryClient  tokenfactorytypes.QueryClient
	txfeesQueryClient        txfeestypes.QueryClient
	txClient                 txtypes.ServiceClient
	wasmQueryClient          wasmtypes.QueryClient
	subaccountToNonce        map[ethcommon.Hash]uint32

	closed  int64
	canSign bool
}

func NewChainClient(
	ctx sdkclient.Context,
	network common.Network,
	options ...common.ClientOption,
) (ChainClient, error) {

	// process options
	opts := common.DefaultClientOptions()

	if network.ChainTLSCert != nil {
		options = append(options, common.OptionTLSCert(network.ChainTLSCert))
	}
	for _, opt := range options {
		if err := opt(opts); err != nil {
			err = errors.Wrap(err, "error in client option")
			return nil, err
		}
	}

	// init tx factory
	var txFactory tx.Factory
	if opts.TxFactory == nil {
		txFactory = NewTxFactory(ctx)
		if opts.GasPrices != "" {
			txFactory = txFactory.WithGasPrices(opts.GasPrices)
		}
	} else {
		txFactory = *opts.TxFactory
	}

	// init grpc connection
	var conn *grpc.ClientConn
	var err error
	stickySessionEnabled := true
	if opts.TLSCert != nil {
		conn, err = grpc.NewClient(network.ChainGrpcEndpoint, grpc.WithTransportCredentials(opts.TLSCert), grpc.WithContextDialer(common.DialerFunc))
	} else {
		conn, err = grpc.NewClient(network.ChainGrpcEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(common.DialerFunc))
		stickySessionEnabled = false
	}
	if err != nil {
		err = errors.Wrapf(err, "failed to connect to the gRPC: %s", network.ChainGrpcEndpoint)
		return nil, err
	}

	var chainStreamConn *grpc.ClientConn
	if opts.TLSCert != nil {
		chainStreamConn, err = grpc.NewClient(network.ChainStreamGrpcEndpoint, grpc.WithTransportCredentials(opts.TLSCert), grpc.WithContextDialer(common.DialerFunc))
	} else {
		chainStreamConn, err = grpc.NewClient(network.ChainStreamGrpcEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(common.DialerFunc))
	}
	if err != nil {
		err = errors.Wrapf(err, "failed to connect to the chain stream gRPC: %s", network.ChainStreamGrpcEndpoint)
		return nil, err
	}

	cancelCtx, cancelFn := context.WithCancel(context.Background())
	// build client
	cc := &chainClient{
		ctx:     ctx,
		network: network,
		opts:    opts,

		logger: log.WithFields(log.Fields{
			"module": "sdk-go",
			"svc":    "chainClient",
		}),

		conn:            conn,
		chainStreamConn: chainStreamConn,
		txFactory:       txFactory,
		canSign:         ctx.Keyring != nil,
		syncMux:         new(sync.Mutex),
		msgC:            make(chan sdk.Msg, msgCommitBatchSizeLimit),
		doneC:           make(chan bool, 1),
		cancelCtx:       cancelCtx,
		cancelFn:        cancelFn,

		sessionEnabled: stickySessionEnabled,

		authQueryClient:          authtypes.NewQueryClient(conn),
		authzQueryClient:         authztypes.NewQueryClient(conn),
		bankQueryClient:          banktypes.NewQueryClient(conn),
		chainStreamClient:        chainstreamtypes.NewStreamClient(chainStreamConn),
		chainStreamV2Client:      chainstreamv2types.NewStreamClient(chainStreamConn),
		distributionQueryClient:  distributiontypes.NewQueryClient(conn),
		exchangeQueryClient:      exchangetypes.NewQueryClient(conn),
		exchangeV2QueryClient:    exchangev2types.NewQueryClient(conn),
		ibcChannelQueryClient:    ibcchanneltypes.NewQueryClient(conn),
		ibcClientQueryClient:     ibcclienttypes.NewQueryClient(conn),
		ibcConnectionQueryClient: ibcconnectiontypes.NewQueryClient(conn),
		ibcTransferQueryClient:   ibctransfertypes.NewQueryClient(conn),
		permissionsQueryClient:   permissionstypes.NewQueryClient(conn),
		tendermintQueryClient:    cmtservice.NewServiceClient(conn),
		tokenfactoryQueryClient:  tokenfactorytypes.NewQueryClient(conn),
		txfeesQueryClient:        txfeestypes.NewQueryClient(conn),
		txClient:                 txtypes.NewServiceClient(conn),
		wasmQueryClient:          wasmtypes.NewQueryClient(conn),
		subaccountToNonce:        make(map[ethcommon.Hash]uint32),
	}

	cc.ofacChecker, err = NewOfacChecker()
	if err != nil {
		return nil, errors.Wrap(err, "Error creating OFAC checker")
	}
	if cc.canSign {
		var err error
		account, err := cc.txFactory.AccountRetriever().GetAccount(ctx, ctx.GetFromAddress())
		if err != nil {
			err = errors.Wrapf(err, "failed to get account")
			return nil, err
		}
		if cc.ofacChecker.IsBlacklisted(account.GetAddress().String()) {
			return nil, errors.Errorf("Address %s is in the OFAC list", account.GetAddress())
		}
		cc.accNum, cc.accSeq = account.GetAccountNumber(), account.GetSequence()
		go cc.runBatchBroadcast()
		go cc.syncTimeoutHeight()
	}

	return cc, nil
}

func (c *chainClient) syncNonce() {
	num, seq, err := c.txFactory.AccountRetriever().GetAccountNumberSequence(c.ctx, c.ctx.GetFromAddress())
	if err != nil {
		c.logger.WithError(err).Errorln("failed to get account seq")
		return
	} else if num != c.accNum {
		c.logger.WithFields(log.Fields{
			"expected": c.accNum,
			"actual":   num,
		}).Panic("account number changed during nonce sync")
	}

	c.accSeq = seq
}

func (c *chainClient) syncTimeoutHeight() {
	t := time.NewTicker(defaultTimeoutHeightSyncInterval)
	defer t.Stop()

	for {
		block, err := c.ctx.Client.Block(c.cancelCtx, nil)
		if err != nil {
			c.logger.WithError(err).Errorln("failed to get current block")
			return
		}
		c.txFactory.WithTimeoutHeight(uint64(block.Block.Height) + defaultTimeoutHeight)

		select {
		case <-c.cancelCtx.Done():
			return
		case <-t.C:
			continue
		}
	}
}

// PrepareFactory ensures the account defined by ctx.GetFromAddress() exists and
// if the account number and/or the account sequence number are zero (not set),
// they will be queried for and set on the provided Factory. A new Factory with
// the updated fields will be returned.
func PrepareFactory(clientCtx sdkclient.Context, txf tx.Factory) (tx.Factory, error) {
	from := clientCtx.GetFromAddress()

	if err := txf.AccountRetriever().EnsureExists(clientCtx, from); err != nil {
		return txf, err
	}

	initNum, initSeq := txf.AccountNumber(), txf.Sequence()
	if initNum == 0 || initSeq == 0 {
		num, seq, err := txf.AccountRetriever().GetAccountNumberSequence(clientCtx, from)
		if err != nil {
			return txf, err
		}

		if initNum == 0 {
			txf = txf.WithAccountNumber(num)
		}

		if initSeq == 0 {
			txf = txf.WithSequence(seq)
		}
	}

	return txf, nil
}

func (c *chainClient) getAccSeq() uint64 {
	defer func() {
		c.accSeq += 1
	}()
	return c.accSeq
}

func (c *chainClient) GetAccNonce() (accNum, accSeq uint64) {
	return c.accNum, c.accSeq
}

func (c *chainClient) QueryClient() *grpc.ClientConn {
	return c.conn
}

func (c *chainClient) ClientContext() sdkclient.Context {
	return c.ctx
}

func (c *chainClient) CanSignTransactions() bool {
	return c.canSign
}

func (c *chainClient) FromAddress() sdk.AccAddress {
	if !c.canSign {
		return sdk.AccAddress{}
	}

	return c.ctx.FromAddress
}

func (c *chainClient) Close() {
	if !c.canSign {
		return
	}
	if atomic.CompareAndSwapInt64(&c.closed, 0, 1) {
		close(c.msgC)
	}

	if c.cancelFn != nil {
		c.cancelFn()
	}
	<-c.doneC
	if c.conn != nil {
		c.conn.Close()
	}
	if c.chainStreamConn != nil {
		c.chainStreamConn.Close()
	}
}

// Bank Module

func (c *chainClient) GetBankBalances(ctx context.Context, address string) (*banktypes.QueryAllBalancesResponse, error) {
	req := &banktypes.QueryAllBalancesRequest{
		Address: address,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.AllBalances, req)

	return res, err
}

func (c *chainClient) GetBankBalance(ctx context.Context, address, denom string) (*banktypes.QueryBalanceResponse, error) {
	req := &banktypes.QueryBalanceRequest{
		Address: address,
		Denom:   denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.Balance, req)

	return res, err
}

func (c *chainClient) GetBankSpendableBalances(ctx context.Context, address string, pagination *query.PageRequest) (*banktypes.QuerySpendableBalancesResponse, error) {
	req := &banktypes.QuerySpendableBalancesRequest{
		Address:    address,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.SpendableBalances, req)

	return res, err
}

func (c *chainClient) GetBankSpendableBalancesByDenom(ctx context.Context, address, denom string) (*banktypes.QuerySpendableBalanceByDenomResponse, error) {
	req := &banktypes.QuerySpendableBalanceByDenomRequest{
		Address: address,
		Denom:   denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.SpendableBalanceByDenom, req)

	return res, err
}

func (c *chainClient) GetBankTotalSupply(ctx context.Context, pagination *query.PageRequest) (*banktypes.QueryTotalSupplyResponse, error) {
	req := &banktypes.QueryTotalSupplyRequest{Pagination: pagination}
	resp, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.TotalSupply, req)

	return resp, err
}

func (c *chainClient) GetBankSupplyOf(ctx context.Context, denom string) (*banktypes.QuerySupplyOfResponse, error) {
	req := &banktypes.QuerySupplyOfRequest{Denom: denom}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.SupplyOf, req)

	return res, err
}

func (c *chainClient) GetDenomMetadata(ctx context.Context, denom string) (*banktypes.QueryDenomMetadataResponse, error) {
	req := &banktypes.QueryDenomMetadataRequest{Denom: denom}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.DenomMetadata, req)

	return res, err
}

func (c *chainClient) GetDenomsMetadata(ctx context.Context, pagination *query.PageRequest) (*banktypes.QueryDenomsMetadataResponse, error) {
	req := &banktypes.QueryDenomsMetadataRequest{Pagination: pagination}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.DenomsMetadata, req)

	return res, err
}

func (c *chainClient) GetDenomOwners(ctx context.Context, denom string, pagination *query.PageRequest) (*banktypes.QueryDenomOwnersResponse, error) {
	req := &banktypes.QueryDenomOwnersRequest{
		Denom:      denom,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.DenomOwners, req)

	return res, err
}

func (c *chainClient) GetBankSendEnabled(ctx context.Context, denoms []string, pagination *query.PageRequest) (*banktypes.QuerySendEnabledResponse, error) {
	req := &banktypes.QuerySendEnabledRequest{
		Denoms:     denoms,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.SendEnabled, req)

	return res, err
}

// Auth Module

func (c *chainClient) GetAccount(ctx context.Context, address string) (*authtypes.QueryAccountResponse, error) {
	req := &authtypes.QueryAccountRequest{
		Address: address,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.authQueryClient.Account, req)

	return res, err
}

func (c *chainClient) SimulateMsg(clientCtx sdkclient.Context, msgs ...sdk.Msg) (*txtypes.SimulateResponse, error) {
	c.txFactory = c.txFactory.WithSequence(c.accSeq)
	c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
	txf, err := PrepareFactory(clientCtx, c.txFactory)
	if err != nil {
		err = errors.Wrap(err, "failed to prepareFactory")
		return nil, err
	}

	simTxBytes, err := txf.BuildSimTx(msgs...)
	if err != nil {
		err = errors.Wrap(err, "failed to build sim tx bytes")
		return nil, err
	}

	ctx := context.Background()
	req := &txtypes.SimulateRequest{TxBytes: simTxBytes}
	simRes, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.txClient.Simulate, req)

	if err != nil {
		err = errors.Wrap(err, "failed to CalculateGas")
		return nil, err
	}

	return simRes, nil
}

func (c *chainClient) BuildSignedTx(clientCtx sdkclient.Context, accNum, accSeq, initialGas uint64, gasPrice uint64, msgs ...sdk.Msg) ([]byte, error) {
	txf := NewTxFactory(clientCtx).WithSequence(accSeq).WithAccountNumber(accNum)
	txf = txf.WithGas(initialGas)

	gasPriceWithDenom := fmt.Sprintf("%d%s", gasPrice, client.InjDenom)
	txf = txf.WithGasPrices(gasPriceWithDenom)

	return c.buildSignedTx(clientCtx, txf, msgs...)
}

func (c *chainClient) buildSignedTx(clientCtx sdkclient.Context, txf tx.Factory, msgs ...sdk.Msg) ([]byte, error) {
	ctx := context.Background()
	if clientCtx.Simulate {
		simTxBytes, err := txf.BuildSimTx(msgs...)
		if err != nil {
			err = errors.Wrap(err, "failed to build sim tx bytes")
			return nil, err
		}

		req := &txtypes.SimulateRequest{TxBytes: simTxBytes}
		simRes, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.txClient.Simulate, req)

		if err != nil {
			err = errors.Wrap(err, "failed to CalculateGas")
			return nil, err
		}

		adjustedGas := uint64(txf.GasAdjustment() * float64(simRes.GasInfo.GasUsed))
		txf = txf.WithGas(adjustedGas)

		c.gasWanted = adjustedGas
	}

	txf, err := PrepareFactory(clientCtx, txf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepareFactory")
	}

	txn, err := txf.BuildUnsignedTx(msgs...)
	if err != nil {
		err = errors.Wrap(err, "failed to BuildUnsignedTx")
		return nil, err
	}

	txn.SetFeeGranter(clientCtx.GetFeeGranterAddress())
	err = tx.Sign(ctx, txf, clientCtx.GetFromName(), txn, true)
	if err != nil {
		err = errors.Wrap(err, "failed to Sign Tx")
		return nil, err
	}

	return clientCtx.TxConfig.TxEncoder()(txn.GetTx())
}

func (c *chainClient) SyncBroadcastSignedTx(txBytes []byte) (*txtypes.BroadcastTxResponse, error) {
	res, err := c.BroadcastSignedTx(txBytes, txtypes.BroadcastMode_BROADCAST_MODE_SYNC)
	if err != nil || res.TxResponse.Code != 0 {
		return res, err
	}

	awaitCtx, cancelFn := context.WithTimeout(context.Background(), defaultBroadcastTimeout)
	defer cancelFn()

	txHash, _ := hex.DecodeString(res.TxResponse.TxHash)
	t := time.NewTimer(defaultBroadcastStatusPoll)

	for {
		select {
		case <-awaitCtx.Done():
			err := errors.Wrapf(ErrTimedOut, "%s", res.TxResponse.TxHash)
			t.Stop()
			return nil, err
		case <-t.C:
			resultTx, err := c.ctx.Client.Tx(awaitCtx, txHash, false)
			if err != nil {
				if errRes := sdkclient.CheckCometError(err, txBytes); errRes != nil {
					return &txtypes.BroadcastTxResponse{TxResponse: errRes}, err
				}

				t.Reset(defaultBroadcastStatusPoll)
				continue

			} else if resultTx.Height > 0 {
				resResultTx := sdk.NewResponseResultTx(resultTx, res.TxResponse.Tx, res.TxResponse.Timestamp)
				res = &txtypes.BroadcastTxResponse{TxResponse: resResultTx}
				t.Stop()
				return res, err
			}

			t.Reset(defaultBroadcastStatusPoll)
		}
	}
}

func (c *chainClient) AsyncBroadcastSignedTx(txBytes []byte) (*txtypes.BroadcastTxResponse, error) {
	return c.BroadcastSignedTx(txBytes, txtypes.BroadcastMode_BROADCAST_MODE_ASYNC)
}

func (c *chainClient) BroadcastSignedTx(txBytes []byte, broadcastMode txtypes.BroadcastMode) (*txtypes.BroadcastTxResponse, error) {
	req := txtypes.BroadcastTxRequest{
		TxBytes: txBytes,
		Mode:    broadcastMode,
	}

	ctx := context.Background()
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.txClient.BroadcastTx, &req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *chainClient) broadcastTx(
	clientCtx sdkclient.Context,
	txf tx.Factory,
	broadcastMode txtypes.BroadcastMode,
	msgs ...sdk.Msg,
) (*txtypes.BroadcastTxRequest, *txtypes.BroadcastTxResponse, error) {
	txBytes, err := c.buildSignedTx(clientCtx, txf, msgs...)
	if err != nil {
		err = errors.Wrap(err, "failed to build signed Tx")
		return nil, nil, err
	}

	req := txtypes.BroadcastTxRequest{
		TxBytes: txBytes,
		Mode:    broadcastMode,
	}

	res, err := common.ExecuteCall(context.Background(), c.network.ChainCookieAssistant, c.txClient.BroadcastTx, &req)
	return &req, res, err

}

// QueueBroadcastMsg enqueues a list of messages. Messages will added to the queue
// and grouped into Txns in chunks. Use this method to mass broadcast Txns with efficiency.
func (c *chainClient) QueueBroadcastMsg(msgs ...sdk.Msg) error {
	if !c.canSign {
		return ErrReadOnly
	} else if atomic.LoadInt64(&c.closed) == 1 {
		return ErrQueueClosed
	}

	t := time.NewTimer(10 * time.Second)
	for _, msg := range msgs {
		select {
		case <-t.C:
			return ErrEnqueueTimeout
		case c.msgC <- msg:
		}
	}
	t.Stop()

	return nil
}

func (c *chainClient) runBatchBroadcast() {
	expirationTimer := time.NewTimer(msgCommitBatchTimeLimit)
	msgBatch := make([]sdk.Msg, 0, msgCommitBatchSizeLimit)

	submitBatch := func(toSubmit []sdk.Msg) {
		res, err := c.SyncBroadcastMsg(toSubmit...)

		if err != nil {
			c.logger.WithError(err)
		} else {
			if res.TxResponse.Code != 0 {
				err = errors.Errorf("error %d (%s): %s", res.TxResponse.Code, res.TxResponse.Codespace, res.TxResponse.RawLog)
				c.logger.WithField("txHash", res.TxResponse.TxHash).WithError(err).Errorln("failed to broadcast messages batch")
			} else {
				c.logger.WithField("txHash", res.TxResponse.TxHash).Debugln("msg batch broadcasted successfully at height", res.TxResponse.Height)
			}
		}

		c.logger.Debugln("gas wanted: ", c.gasWanted)
	}

	for {
		select {
		case msg, ok := <-c.msgC:
			if !ok {
				// exit required
				if len(msgBatch) > 0 {
					submitBatch(msgBatch)
				}

				close(c.doneC)
				return
			}

			msgBatch = append(msgBatch, msg)

			if len(msgBatch) >= msgCommitBatchSizeLimit {
				toSubmit := msgBatch
				msgBatch = msgBatch[:0]
				expirationTimer.Reset(msgCommitBatchTimeLimit)

				submitBatch(toSubmit)
			}
		case <-expirationTimer.C:
			if len(msgBatch) > 0 {
				toSubmit := msgBatch
				msgBatch = msgBatch[:0]
				expirationTimer.Reset(msgCommitBatchTimeLimit)
				submitBatch(toSubmit)
			} else {
				expirationTimer.Reset(msgCommitBatchTimeLimit)
			}
		}
	}
}

func (c *chainClient) GetGasFee() (string, error) {
	gasPrices := strings.Trim(c.opts.GasPrices, "inj")

	gas, err := strconv.ParseFloat(gasPrices, 64)

	if err != nil {
		return "", err
	}

	gasFeeAdjusted := gas * float64(c.gasWanted) / math.Pow(10, 18)
	gasFeeFormatted := strconv.FormatFloat(gasFeeAdjusted, 'f', -1, 64)
	c.gasFee = gasFeeFormatted

	return c.gasFee, err
}

func (c *chainClient) DefaultSubaccount(acc sdk.AccAddress) ethcommon.Hash {
	return c.Subaccount(acc, 0)
}

func (c *chainClient) Subaccount(account sdk.AccAddress, index int) ethcommon.Hash {
	ethAddress := ethcommon.BytesToAddress(account.Bytes())
	ethLowerAddress := strings.ToLower(ethAddress.String())

	subaccountId := fmt.Sprintf("%s%024x", ethLowerAddress, index)
	return ethcommon.HexToHash(subaccountId)
}

// Deprecated: use GetSubAccountNonceV2 instead
func (c *chainClient) GetSubAccountNonce(ctx context.Context, subaccountId ethcommon.Hash) (*exchangetypes.QuerySubaccountTradeNonceResponse, error) {
	req := &exchangetypes.QuerySubaccountTradeNonceRequest{SubaccountId: subaccountId.String()}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.SubaccountTradeNonce, req)

	return res, err
}

// Deprecated: use GetFeeDiscountInfoV2 instead
func (c *chainClient) GetFeeDiscountInfo(ctx context.Context, account string) (*exchangetypes.QueryFeeDiscountAccountInfoResponse, error) {
	req := &exchangetypes.QueryFeeDiscountAccountInfoRequest{
		Account: account,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.FeeDiscountAccountInfo, req)

	return res, err
}

func (c *chainClient) GetSubAccountNonceV2(ctx context.Context, subaccountId ethcommon.Hash) (*exchangev2types.QuerySubaccountTradeNonceResponse, error) {
	req := &exchangev2types.QuerySubaccountTradeNonceRequest{SubaccountId: subaccountId.String()}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.SubaccountTradeNonce, req)

	return res, err
}

func (c *chainClient) GetFeeDiscountInfoV2(ctx context.Context, account string) (*exchangev2types.QueryFeeDiscountAccountInfoResponse, error) {
	req := &exchangev2types.QueryFeeDiscountAccountInfoRequest{
		Account: account,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.FeeDiscountAccountInfo, req)

	return res, err
}

// Deprecated: use CreateSpotOrderV2 instead
func (c *chainClient) CreateSpotOrder(defaultSubaccountID ethcommon.Hash, d *SpotOrderData, marketsAssistant MarketsAssistant) *exchangetypes.SpotOrder {

	market, isPresent := marketsAssistant.AllSpotMarkets()[d.MarketId]
	if !isPresent {
		panic(errors.Errorf("Invalid spot market id for %s network (%s)", c.network.Name, d.MarketId))
	}

	orderSize := market.QuantityToChainFormat(d.Quantity)
	orderPrice := market.PriceToChainFormat(d.Price)

	return &exchangetypes.SpotOrder{
		MarketId:  d.MarketId,
		OrderType: exchangetypes.OrderType(d.OrderType),
		OrderInfo: exchangetypes.OrderInfo{
			SubaccountId: defaultSubaccountID.Hex(),
			FeeRecipient: d.FeeRecipient,
			Price:        orderPrice,
			Quantity:     orderSize,
			Cid:          d.Cid,
		},
	}
}

// Deprecated: use CreateDerivativeOrderV2 instead
func (c *chainClient) CreateDerivativeOrder(defaultSubaccountID ethcommon.Hash, d *DerivativeOrderData, marketAssistant MarketsAssistant) *exchangetypes.DerivativeOrder {
	market, isPresent := marketAssistant.AllDerivativeMarkets()[d.MarketId]
	if !isPresent {
		panic(errors.Errorf("Invalid derivative market id for %s network (%s)", c.network.Name, d.MarketId))
	}

	orderSize := market.QuantityToChainFormat(d.Quantity)
	orderPrice := market.PriceToChainFormat(d.Price)
	orderMargin := sdkmath.LegacyMustNewDecFromStr("0")

	if !d.IsReduceOnly {
		orderMargin = market.CalculateMarginInChainFormat(d.Quantity, d.Price, d.Leverage)
	}

	return &exchangetypes.DerivativeOrder{
		MarketId:  d.MarketId,
		OrderType: exchangetypes.OrderType(d.OrderType),
		Margin:    orderMargin,
		OrderInfo: exchangetypes.OrderInfo{
			SubaccountId: defaultSubaccountID.Hex(),
			FeeRecipient: d.FeeRecipient,
			Price:        orderPrice,
			Quantity:     orderSize,
			Cid:          d.Cid,
		},
	}
}

// Deprecated: use OrderCancelV2 instead
func (c *chainClient) OrderCancel(defaultSubaccountID ethcommon.Hash, d *OrderCancelData) *exchangetypes.OrderData {
	return &exchangetypes.OrderData{
		MarketId:     d.MarketId,
		OrderHash:    d.OrderHash,
		SubaccountId: defaultSubaccountID.Hex(),
		Cid:          d.Cid,
	}
}

func (c *chainClient) CreateSpotOrderV2(defaultSubaccountID ethcommon.Hash, d *SpotOrderData) *exchangev2types.SpotOrder {
	return &exchangev2types.SpotOrder{
		MarketId:  d.MarketId,
		OrderType: d.OrderType,
		OrderInfo: exchangev2types.OrderInfo{
			SubaccountId: defaultSubaccountID.Hex(),
			FeeRecipient: d.FeeRecipient,
			Price:        sdkmath.LegacyMustNewDecFromStr(d.Price.String()),
			Quantity:     sdkmath.LegacyMustNewDecFromStr(d.Quantity.String()),
			Cid:          d.Cid,
		},
	}
}

func (c *chainClient) CreateDerivativeOrderV2(defaultSubaccountID ethcommon.Hash, d *DerivativeOrderData) *exchangev2types.DerivativeOrder {
	orderMargin := sdkmath.LegacyMustNewDecFromStr("0")

	if !d.IsReduceOnly {
		orderMargin = sdkmath.LegacyMustNewDecFromStr(d.Quantity.Mul(d.Price).Div(d.Leverage).String())
	}

	return &exchangev2types.DerivativeOrder{
		MarketId:  d.MarketId,
		OrderType: d.OrderType,
		Margin:    orderMargin,
		OrderInfo: exchangev2types.OrderInfo{
			SubaccountId: defaultSubaccountID.Hex(),
			FeeRecipient: d.FeeRecipient,
			Price:        sdkmath.LegacyMustNewDecFromStr(d.Price.String()),
			Quantity:     sdkmath.LegacyMustNewDecFromStr(d.Quantity.String()),
			Cid:          d.Cid,
		},
	}
}

func (c *chainClient) OrderCancelV2(defaultSubaccountID ethcommon.Hash, d *OrderCancelData) *exchangev2types.OrderData {
	return &exchangev2types.OrderData{
		MarketId:     d.MarketId,
		OrderHash:    d.OrderHash,
		SubaccountId: defaultSubaccountID.Hex(),
		Cid:          d.Cid,
	}
}

func (c *chainClient) GetAuthzGrants(ctx context.Context, req authztypes.QueryGrantsRequest) (*authztypes.QueryGrantsResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.authzQueryClient.Grants, &req)

	return res, err
}

func (c *chainClient) BuildGenericAuthz(granter, grantee, msgtype string, expireIn time.Time) *authztypes.MsgGrant {
	if c.ofacChecker.IsBlacklisted(granter) {
		panic("Address is in the OFAC list") // panics should generally be avoided, but otherwise function signature should be changed
	}
	authz := authztypes.NewGenericAuthorization(msgtype)
	authzAny := codectypes.UnsafePackAny(authz)
	return &authztypes.MsgGrant{
		Granter: granter,
		Grantee: grantee,
		Grant: authztypes.Grant{
			Authorization: authzAny,
			Expiration:    &expireIn,
		},
	}
}

type ExchangeAuthz string

var (
	CreateSpotLimitOrderAuthz       = ExchangeAuthz("/" + proto.MessageName(&exchangev2types.CreateSpotLimitOrderAuthz{}))
	CreateSpotMarketOrderAuthz      = ExchangeAuthz("/" + proto.MessageName(&exchangev2types.CreateSpotMarketOrderAuthz{}))
	BatchCreateSpotLimitOrdersAuthz = ExchangeAuthz("/" + proto.MessageName(&exchangev2types.BatchCreateSpotLimitOrdersAuthz{}))
	CancelSpotOrderAuthz            = ExchangeAuthz("/" + proto.MessageName(&exchangev2types.CancelSpotOrderAuthz{}))
	BatchCancelSpotOrdersAuthz      = ExchangeAuthz("/" + proto.MessageName(&exchangev2types.BatchCancelSpotOrdersAuthz{}))

	CreateDerivativeLimitOrderAuthz       = ExchangeAuthz("/" + proto.MessageName(&exchangev2types.CreateDerivativeLimitOrderAuthz{}))
	CreateDerivativeMarketOrderAuthz      = ExchangeAuthz("/" + proto.MessageName(&exchangev2types.CreateDerivativeMarketOrderAuthz{}))
	BatchCreateDerivativeLimitOrdersAuthz = ExchangeAuthz("/" + proto.MessageName(&exchangev2types.BatchCreateDerivativeLimitOrdersAuthz{}))
	CancelDerivativeOrderAuthz            = ExchangeAuthz("/" + proto.MessageName(&exchangev2types.CancelDerivativeOrderAuthz{}))
	BatchCancelDerivativeOrdersAuthz      = ExchangeAuthz("/" + proto.MessageName(&exchangev2types.BatchCancelDerivativeOrdersAuthz{}))

	BatchUpdateOrdersAuthz = ExchangeAuthz("/" + proto.MessageName(&exchangev2types.BatchUpdateOrdersAuthz{}))
)

func (c *chainClient) BuildExchangeAuthz(granter, grantee string, authzType ExchangeAuthz, subaccountId string, markets []string, expireIn time.Time) *authztypes.MsgGrant {
	if c.ofacChecker.IsBlacklisted(granter) {
		panic("Address is in the OFAC list") // panics should generally be avoided, but otherwise function signature should be changed
	}
	var typedAuthzAny codectypes.Any
	var typedAuthzBytes []byte
	switch authzType {
	// spot msgs
	case CreateSpotLimitOrderAuthz:
		typedAuthz := &exchangev2types.CreateSpotLimitOrderAuthz{
			SubaccountId: subaccountId,
			MarketIds:    markets,
		}
		typedAuthzBytes, _ = typedAuthz.Marshal()
	case CreateSpotMarketOrderAuthz:
		typedAuthz := &exchangev2types.CreateSpotMarketOrderAuthz{
			SubaccountId: subaccountId,
			MarketIds:    markets,
		}
		typedAuthzBytes, _ = typedAuthz.Marshal()
	case BatchCreateSpotLimitOrdersAuthz:
		typedAuthz := &exchangev2types.BatchCreateSpotLimitOrdersAuthz{
			SubaccountId: subaccountId,
			MarketIds:    markets,
		}
		typedAuthzBytes, _ = typedAuthz.Marshal()
	case CancelSpotOrderAuthz:
		typedAuthz := &exchangev2types.CancelSpotOrderAuthz{
			SubaccountId: subaccountId,
			MarketIds:    markets,
		}
		typedAuthzBytes, _ = typedAuthz.Marshal()
	case BatchCancelSpotOrdersAuthz:
		typedAuthz := &exchangev2types.BatchCancelSpotOrdersAuthz{
			SubaccountId: subaccountId,
			MarketIds:    markets,
		}
		typedAuthzBytes, _ = typedAuthz.Marshal()
	// derivative msgs
	case CreateDerivativeLimitOrderAuthz:
		typedAuthz := &exchangev2types.CreateDerivativeLimitOrderAuthz{
			SubaccountId: subaccountId,
			MarketIds:    markets,
		}
		typedAuthzBytes, _ = typedAuthz.Marshal()
	case CreateDerivativeMarketOrderAuthz:
		typedAuthz := &exchangev2types.CreateDerivativeMarketOrderAuthz{
			SubaccountId: subaccountId,
			MarketIds:    markets,
		}
		typedAuthzBytes, _ = typedAuthz.Marshal()
	case BatchCreateDerivativeLimitOrdersAuthz:
		typedAuthz := &exchangev2types.BatchCreateDerivativeLimitOrdersAuthz{
			SubaccountId: subaccountId,
			MarketIds:    markets,
		}
		typedAuthzBytes, _ = typedAuthz.Marshal()
	case CancelDerivativeOrderAuthz:
		typedAuthz := &exchangev2types.CancelDerivativeOrderAuthz{
			SubaccountId: subaccountId,
			MarketIds:    markets,
		}
		typedAuthzBytes, _ = typedAuthz.Marshal()
	case BatchCancelDerivativeOrdersAuthz:
		typedAuthz := &exchangev2types.BatchCancelDerivativeOrdersAuthz{
			SubaccountId: subaccountId,
			MarketIds:    markets,
		}
		typedAuthzBytes, _ = typedAuthz.Marshal()
	// common msgs
	case BatchUpdateOrdersAuthz:
		panic("please use BuildExchangeBatchUpdateOrdersAuthz for BatchUpdateOrdersAuthz")
	default:
		panic("unsupported exchange authz type")
	}

	typedAuthzAny = codectypes.Any{
		TypeUrl: string(authzType),
		Value:   typedAuthzBytes,
	}

	return &authztypes.MsgGrant{
		Granter: granter,
		Grantee: grantee,
		Grant: authztypes.Grant{
			Authorization: &typedAuthzAny,
			Expiration:    &expireIn,
		},
	}
}

func (c *chainClient) BuildExchangeBatchUpdateOrdersAuthz(
	granter string,
	grantee string,
	subaccountId string,
	spotMarkets []string,
	derivativeMarkets []string,
	expireIn time.Time,
) *authztypes.MsgGrant {
	if c.ofacChecker.IsBlacklisted(granter) {
		panic("Address is in the OFAC list") // panics should generally be avoided, but otherwise function signature should be changed
	}
	typedAuthz := &exchangev2types.BatchUpdateOrdersAuthz{
		SubaccountId:      subaccountId,
		SpotMarkets:       spotMarkets,
		DerivativeMarkets: derivativeMarkets,
	}
	typedAuthzBytes, _ := typedAuthz.Marshal()
	typedAuthzAny := codectypes.Any{
		TypeUrl: string(BatchUpdateOrdersAuthz),
		Value:   typedAuthzBytes,
	}
	return &authztypes.MsgGrant{
		Granter: granter,
		Grantee: grantee,
		Grant: authztypes.Grant{
			Authorization: &typedAuthzAny,
			Expiration:    &expireIn,
		},
	}
}

func (c *chainClient) StreamEventOrderFail(sender string, failEventCh chan map[string]uint) {
	var cometbftClient *rpchttp.HTTP
	var err error

	cometbftClient, err = rpchttp.New(c.network.TmEndpoint)
	if err != nil {
		panic(err)
	}

	if !cometbftClient.IsRunning() {
		err = cometbftClient.Start()
		if err != nil {
			panic(err)
		}
	}
	defer func() {
		err := cometbftClient.Stop()
		if err != nil {
			panic(err)
		}
	}()

	c.StreamEventOrderFailWithWebsocket(sender, cometbftClient, failEventCh)
}

func (c *chainClient) StreamEventOrderFailWithWebsocket(sender string, websocket *rpchttp.HTTP, failEventCh chan map[string]uint) {
	filter := fmt.Sprintf("tm.event='Tx' AND message.sender='%s' AND message.action='/injective.exchange.v1beta1.MsgBatchUpdateOrders' AND injective.exchange.v1beta1.EventOrderFail.flags EXISTS", sender)
	eventCh, err := websocket.Subscribe(context.Background(), "OrderFail", filter, 10000)
	if err != nil {
		panic(err)
	}

	// stream and extract fail events
	for {
		e := <-eventCh

		var failedOrderHashes []string
		err = json.Unmarshal([]byte(e.Events["injective.exchange.v1beta1.EventOrderFail.hashes"][0]), &failedOrderHashes)
		if err != nil {
			panic(err)
		}

		var failedOrderCodes []uint
		err = json.Unmarshal([]byte(e.Events["injective.exchange.v1beta1.EventOrderFail.flags"][0]), &failedOrderCodes)
		if err != nil {
			panic(err)
		}

		results := map[string]uint{}
		for i, hash := range failedOrderHashes {
			orderHashBytes, _ := base64.StdEncoding.DecodeString(hash)
			orderHash := ethcommon.BytesToHash(orderHashBytes).String()
			results[orderHash] = failedOrderCodes[i]
		}

		failEventCh <- results
	}
}

// Deprecated: use the chain stream instead
func (c *chainClient) StreamOrderbookUpdateEvents(orderbookType OrderbookType, marketIDs []string, orderbookCh chan exchangetypes.Orderbook) {
	var cometbftClient *rpchttp.HTTP
	var err error

	cometbftClient, err = rpchttp.New(c.network.TmEndpoint)
	if err != nil {
		panic(err)
	}

	if !cometbftClient.IsRunning() {
		err = cometbftClient.Start()
		if err != nil {
			panic(err)
		}
	}
	defer func() {
		err := cometbftClient.Stop()
		if err != nil {
			panic(err)
		}
	}()

	c.StreamOrderbookUpdateEventsWithWebsocket(orderbookType, marketIDs, cometbftClient, orderbookCh)

}

// Deprecated: use the chain stream instead
func (c *chainClient) StreamOrderbookUpdateEventsWithWebsocket(orderbookType OrderbookType, marketIDs []string, websocket *rpchttp.HTTP, orderbookCh chan exchangetypes.Orderbook) {
	filter := fmt.Sprintf("tm.event='NewBlock' AND %s EXISTS", orderbookType)
	eventCh, err := websocket.Subscribe(context.Background(), "OrderbookUpdate", filter, 10000)
	if err != nil {
		panic(err)
	}

	// turn array into map for convenient lookup
	marketIDsMap := map[string]bool{}
	for _, id := range marketIDs {
		marketIDsMap[id] = true
	}

	filteredOrderbookUpdateCh := make(chan exchangetypes.Orderbook, 10000)

	// stream and filter orderbooks
	go func() {
		for {
			e := <-eventCh

			var allOrderbookUpdates []exchangetypes.Orderbook
			err = json.Unmarshal([]byte(e.Events[string(orderbookType)][0]), &allOrderbookUpdates)
			if err != nil {
				panic(err)
			}

			for _, ob := range allOrderbookUpdates {
				id := ethcommon.BytesToHash(ob.MarketId).String()
				if marketIDsMap[id] {
					filteredOrderbookUpdateCh <- ob
				}
			}
		}
	}()

	// fetch the orderbooks

	// consume from filtered orderbooks channel
	for {
		ob := <-filteredOrderbookUpdateCh

		// skip update id until it's good to consume

		// construct up-to-date orderbook

		// send results to channel
		orderbookCh <- ob
	}
}

func (c *chainClient) GetTx(ctx context.Context, txHash string) (*txtypes.GetTxResponse, error) {
	req := &txtypes.GetTxRequest{
		Hash: txHash,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.txClient.GetTx, req)

	return res, err
}

// Deprecated: use ChainStreamV2 instead
func (c *chainClient) ChainStream(ctx context.Context, req chainstreamtypes.StreamRequest) (chainstreamtypes.Stream_StreamClient, error) {
	stream, err := common.ExecuteStreamCall(ctx, c.network.ChainCookieAssistant, c.chainStreamClient.Stream, &req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

func (c *chainClient) ChainStreamV2(ctx context.Context, req chainstreamv2types.StreamRequest) (chainstreamv2types.Stream_StreamV2Client, error) {
	stream, err := common.ExecuteStreamCall(ctx, c.network.ChainCookieAssistant, c.chainStreamV2Client.StreamV2, &req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

// wasm module

func (c *chainClient) FetchContractInfo(ctx context.Context, address string) (*wasmtypes.QueryContractInfoResponse, error) {
	req := &wasmtypes.QueryContractInfoRequest{
		Address: address,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.wasmQueryClient.ContractInfo, req)

	return res, err
}

func (c *chainClient) FetchContractHistory(ctx context.Context, address string, pagination *query.PageRequest) (*wasmtypes.QueryContractHistoryResponse, error) {
	req := &wasmtypes.QueryContractHistoryRequest{
		Address:    address,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.wasmQueryClient.ContractHistory, req)

	return res, err
}

func (c *chainClient) FetchContractsByCode(ctx context.Context, codeId uint64, pagination *query.PageRequest) (*wasmtypes.QueryContractsByCodeResponse, error) {
	req := &wasmtypes.QueryContractsByCodeRequest{
		CodeId:     codeId,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.wasmQueryClient.ContractsByCode, req)

	return res, err
}

func (c *chainClient) FetchAllContractsState(ctx context.Context, address string, pagination *query.PageRequest) (*wasmtypes.QueryAllContractStateResponse, error) {
	req := &wasmtypes.QueryAllContractStateRequest{
		Address:    address,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.wasmQueryClient.AllContractState, req)

	return res, err
}

func (c *chainClient) RawContractState(
	ctx context.Context,
	contractAddress string,
	queryData []byte,
) (*wasmtypes.QueryRawContractStateResponse, error) {
	req := &wasmtypes.QueryRawContractStateRequest{
		Address:   contractAddress,
		QueryData: queryData,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.wasmQueryClient.RawContractState, req)

	return res, err
}

func (c *chainClient) SmartContractState(
	ctx context.Context,
	contractAddress string,
	queryData []byte,
) (*wasmtypes.QuerySmartContractStateResponse, error) {
	req := &wasmtypes.QuerySmartContractStateRequest{
		Address:   contractAddress,
		QueryData: queryData,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.wasmQueryClient.SmartContractState, req)

	return res, err
}

func (c *chainClient) FetchCode(ctx context.Context, codeId uint64) (*wasmtypes.QueryCodeResponse, error) {
	req := &wasmtypes.QueryCodeRequest{
		CodeId: codeId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.wasmQueryClient.Code, req)

	return res, err
}

func (c *chainClient) FetchCodes(ctx context.Context, pagination *query.PageRequest) (*wasmtypes.QueryCodesResponse, error) {
	req := &wasmtypes.QueryCodesRequest{
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.wasmQueryClient.Codes, req)

	return res, err
}

func (c *chainClient) FetchPinnedCodes(ctx context.Context, pagination *query.PageRequest) (*wasmtypes.QueryPinnedCodesResponse, error) {
	req := &wasmtypes.QueryPinnedCodesRequest{
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.wasmQueryClient.PinnedCodes, req)

	return res, err
}

func (c *chainClient) FetchContractsByCreator(ctx context.Context, creator string, pagination *query.PageRequest) (*wasmtypes.QueryContractsByCreatorResponse, error) {
	req := &wasmtypes.QueryContractsByCreatorRequest{
		CreatorAddress: creator,
		Pagination:     pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.wasmQueryClient.ContractsByCreator, req)

	return res, err
}

// Tokenfactory module

func (c *chainClient) FetchDenomAuthorityMetadata(ctx context.Context, creator, subDenom string) (*tokenfactorytypes.QueryDenomAuthorityMetadataResponse, error) {
	req := &tokenfactorytypes.QueryDenomAuthorityMetadataRequest{
		Creator: creator,
	}

	if subDenom != "" {
		req.SubDenom = subDenom
	}

	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tokenfactoryQueryClient.DenomAuthorityMetadata, req)

	return res, err
}

func (c *chainClient) FetchDenomsFromCreator(ctx context.Context, creator string) (*tokenfactorytypes.QueryDenomsFromCreatorResponse, error) {
	req := &tokenfactorytypes.QueryDenomsFromCreatorRequest{
		Creator: creator,
	}

	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tokenfactoryQueryClient.DenomsFromCreator, req)

	return res, err
}

func (c *chainClient) FetchTokenfactoryModuleState(ctx context.Context) (*tokenfactorytypes.QueryModuleStateResponse, error) {
	req := &tokenfactorytypes.QueryModuleStateRequest{}

	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tokenfactoryQueryClient.TokenfactoryModuleState, req)

	return res, err
}

type DerivativeOrderData struct {
	OrderType    exchangev2types.OrderType
	Price        decimal.Decimal
	Quantity     decimal.Decimal
	Leverage     decimal.Decimal
	FeeRecipient string
	MarketId     string
	IsReduceOnly bool
	Cid          string
}

type SpotOrderData struct {
	OrderType    exchangev2types.OrderType
	Price        decimal.Decimal
	Quantity     decimal.Decimal
	FeeRecipient string
	MarketId     string
	Cid          string
}

type OrderCancelData struct {
	MarketId  string
	OrderHash string
	Cid       string
}

// Distribution module
func (c *chainClient) FetchValidatorDistributionInfo(ctx context.Context, validatorAddress string) (*distributiontypes.QueryValidatorDistributionInfoResponse, error) {
	req := &distributiontypes.QueryValidatorDistributionInfoRequest{
		ValidatorAddress: validatorAddress,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.distributionQueryClient.ValidatorDistributionInfo, req)

	return res, err
}

func (c *chainClient) FetchValidatorOutstandingRewards(ctx context.Context, validatorAddress string) (*distributiontypes.QueryValidatorOutstandingRewardsResponse, error) {
	req := &distributiontypes.QueryValidatorOutstandingRewardsRequest{
		ValidatorAddress: validatorAddress,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.distributionQueryClient.ValidatorOutstandingRewards, req)

	return res, err
}

func (c *chainClient) FetchValidatorCommission(ctx context.Context, validatorAddress string) (*distributiontypes.QueryValidatorCommissionResponse, error) {
	req := &distributiontypes.QueryValidatorCommissionRequest{
		ValidatorAddress: validatorAddress,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.distributionQueryClient.ValidatorCommission, req)

	return res, err
}

func (c *chainClient) FetchValidatorSlashes(ctx context.Context, validatorAddress string, startingHeight, endingHeight uint64, pagination *query.PageRequest) (*distributiontypes.QueryValidatorSlashesResponse, error) {
	req := &distributiontypes.QueryValidatorSlashesRequest{
		ValidatorAddress: validatorAddress,
		StartingHeight:   startingHeight,
		EndingHeight:     endingHeight,
		Pagination:       pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.distributionQueryClient.ValidatorSlashes, req)

	return res, err
}

func (c *chainClient) FetchDelegationRewards(ctx context.Context, delegatorAddress, validatorAddress string) (*distributiontypes.QueryDelegationRewardsResponse, error) {
	req := &distributiontypes.QueryDelegationRewardsRequest{
		DelegatorAddress: delegatorAddress,
		ValidatorAddress: validatorAddress,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.distributionQueryClient.DelegationRewards, req)

	return res, err
}

func (c *chainClient) FetchDelegationTotalRewards(ctx context.Context, delegatorAddress string) (*distributiontypes.QueryDelegationTotalRewardsResponse, error) {
	req := &distributiontypes.QueryDelegationTotalRewardsRequest{
		DelegatorAddress: delegatorAddress,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.distributionQueryClient.DelegationTotalRewards, req)

	return res, err
}

func (c *chainClient) FetchDelegatorValidators(ctx context.Context, delegatorAddress string) (*distributiontypes.QueryDelegatorValidatorsResponse, error) {
	req := &distributiontypes.QueryDelegatorValidatorsRequest{
		DelegatorAddress: delegatorAddress,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.distributionQueryClient.DelegatorValidators, req)

	return res, err
}

func (c *chainClient) FetchDelegatorWithdrawAddress(ctx context.Context, delegatorAddress string) (*distributiontypes.QueryDelegatorWithdrawAddressResponse, error) {
	req := &distributiontypes.QueryDelegatorWithdrawAddressRequest{
		DelegatorAddress: delegatorAddress,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.distributionQueryClient.DelegatorWithdrawAddress, req)

	return res, err
}

func (c *chainClient) FetchCommunityPool(ctx context.Context) (*distributiontypes.QueryCommunityPoolResponse, error) {
	req := &distributiontypes.QueryCommunityPoolRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.distributionQueryClient.CommunityPool, req)

	return res, err
}

// Chain exchange module

// Deprecated: use FetchSubaccountDepositsV2 instead
func (c *chainClient) FetchSubaccountDeposits(ctx context.Context, subaccountId string) (*exchangetypes.QuerySubaccountDepositsResponse, error) {
	req := &exchangetypes.QuerySubaccountDepositsRequest{
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.SubaccountDeposits, req)

	return res, err
}

// Deprecated: use FetchSubaccountDepositV2 instead
func (c *chainClient) FetchSubaccountDeposit(ctx context.Context, subaccountId, denom string) (*exchangetypes.QuerySubaccountDepositResponse, error) {
	req := &exchangetypes.QuerySubaccountDepositRequest{
		SubaccountId: subaccountId,
		Denom:        denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.SubaccountDeposit, req)

	return res, err
}

// Deprecated: use FetchExchangeBalancesV2 instead
func (c *chainClient) FetchExchangeBalances(ctx context.Context) (*exchangetypes.QueryExchangeBalancesResponse, error) {
	req := &exchangetypes.QueryExchangeBalancesRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.ExchangeBalances, req)

	return res, err
}

// Deprecated: use FetchAggregateVolumeV2 instead
func (c *chainClient) FetchAggregateVolume(ctx context.Context, account string) (*exchangetypes.QueryAggregateVolumeResponse, error) {
	req := &exchangetypes.QueryAggregateVolumeRequest{
		Account: account,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.AggregateVolume, req)

	return res, err
}

// Deprecated: use FetchAggregateVolumesV2 instead
func (c *chainClient) FetchAggregateVolumes(ctx context.Context, accounts, marketIDs []string) (*exchangetypes.QueryAggregateVolumesResponse, error) {
	req := &exchangetypes.QueryAggregateVolumesRequest{
		Accounts:  accounts,
		MarketIds: marketIDs,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.AggregateVolumes, req)

	return res, err
}

// Deprecated: use FetchAggregateMarketVolumeV2 instead
func (c *chainClient) FetchAggregateMarketVolume(ctx context.Context, marketId string) (*exchangetypes.QueryAggregateMarketVolumeResponse, error) {
	req := &exchangetypes.QueryAggregateMarketVolumeRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.AggregateMarketVolume, req)

	return res, err
}

// Deprecated: use FetchAggregateMarketVolumesV2 instead
func (c *chainClient) FetchAggregateMarketVolumes(ctx context.Context, marketIDs []string) (*exchangetypes.QueryAggregateMarketVolumesResponse, error) {
	req := &exchangetypes.QueryAggregateMarketVolumesRequest{
		MarketIds: marketIDs,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.AggregateMarketVolumes, req)

	return res, err
}

// Deprecated: use FetchDenomDecimalV2 instead
func (c *chainClient) FetchDenomDecimal(ctx context.Context, denom string) (*exchangetypes.QueryDenomDecimalResponse, error) {
	req := &exchangetypes.QueryDenomDecimalRequest{
		Denom: denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.DenomDecimal, req)

	return res, err
}

// Deprecated: use FetchDenomDecimalsV2 instead
func (c *chainClient) FetchDenomDecimals(ctx context.Context, denoms []string) (*exchangetypes.QueryDenomDecimalsResponse, error) {
	req := &exchangetypes.QueryDenomDecimalsRequest{
		Denoms: denoms,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.DenomDecimals, req)

	return res, err
}

// Deprecated: use FetchChainSpotMarketsV2 instead
func (c *chainClient) FetchChainSpotMarkets(ctx context.Context, status string, marketIDs []string) (*exchangetypes.QuerySpotMarketsResponse, error) {
	req := &exchangetypes.QuerySpotMarketsRequest{
		MarketIds: marketIDs,
	}
	if status != "" {
		req.Status = status
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.SpotMarkets, req)

	return res, err
}

// Deprecated: use FetchChainSpotMarketV2 instead
func (c *chainClient) FetchChainSpotMarket(ctx context.Context, marketId string) (*exchangetypes.QuerySpotMarketResponse, error) {
	req := &exchangetypes.QuerySpotMarketRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.SpotMarket, req)

	return res, err
}

// Deprecated: use FetchChainFullSpotMarketsV2 instead
func (c *chainClient) FetchChainFullSpotMarkets(ctx context.Context, status string, marketIDs []string, withMidPriceAndTob bool) (*exchangetypes.QueryFullSpotMarketsResponse, error) {
	req := &exchangetypes.QueryFullSpotMarketsRequest{
		MarketIds:          marketIDs,
		WithMidPriceAndTob: withMidPriceAndTob,
	}
	if status != "" {
		req.Status = status
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.FullSpotMarkets, req)

	return res, err
}

// Deprecated: use FetchChainFullSpotMarketV2 instead
func (c *chainClient) FetchChainFullSpotMarket(ctx context.Context, marketId string, withMidPriceAndTob bool) (*exchangetypes.QueryFullSpotMarketResponse, error) {
	req := &exchangetypes.QueryFullSpotMarketRequest{
		MarketId:           marketId,
		WithMidPriceAndTob: withMidPriceAndTob,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.FullSpotMarket, req)

	return res, err
}

// Deprecated: use FetchChainSpotOrderbookV2 instead
func (c *chainClient) FetchChainSpotOrderbook(ctx context.Context, marketId string, limit uint64, orderSide exchangetypes.OrderSide, limitCumulativeNotional, limitCumulativeQuantity sdkmath.LegacyDec) (*exchangetypes.QuerySpotOrderbookResponse, error) {
	req := &exchangetypes.QuerySpotOrderbookRequest{
		MarketId:                marketId,
		Limit:                   limit,
		OrderSide:               orderSide,
		LimitCumulativeNotional: &limitCumulativeNotional,
		LimitCumulativeQuantity: &limitCumulativeQuantity,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.SpotOrderbook, req)

	return res, err
}

// Deprecated: use FetchChainTraderSpotOrdersV2 instead
func (c *chainClient) FetchChainTraderSpotOrders(ctx context.Context, marketId, subaccountId string) (*exchangetypes.QueryTraderSpotOrdersResponse, error) {
	req := &exchangetypes.QueryTraderSpotOrdersRequest{
		MarketId:     marketId,
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.TraderSpotOrders, req)

	return res, err
}

// Deprecated: use FetchChainAccountAddressSpotOrdersV2 instead
func (c *chainClient) FetchChainAccountAddressSpotOrders(ctx context.Context, marketId, address string) (*exchangetypes.QueryAccountAddressSpotOrdersResponse, error) {
	req := &exchangetypes.QueryAccountAddressSpotOrdersRequest{
		MarketId:       marketId,
		AccountAddress: address,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.AccountAddressSpotOrders, req)

	return res, err
}

// Deprecated: use FetchChainSpotOrdersByHashesV2 instead
func (c *chainClient) FetchChainSpotOrdersByHashes(ctx context.Context, marketId, subaccountId string, orderHashes []string) (*exchangetypes.QuerySpotOrdersByHashesResponse, error) {
	req := &exchangetypes.QuerySpotOrdersByHashesRequest{
		MarketId:     marketId,
		SubaccountId: subaccountId,
		OrderHashes:  orderHashes,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.SpotOrdersByHashes, req)

	return res, err
}

// Deprecated: use FetchChainSubaccountOrdersV2 instead
func (c *chainClient) FetchChainSubaccountOrders(ctx context.Context, subaccountId, marketId string) (*exchangetypes.QuerySubaccountOrdersResponse, error) {
	req := &exchangetypes.QuerySubaccountOrdersRequest{
		SubaccountId: subaccountId,
		MarketId:     marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.SubaccountOrders, req)

	return res, err
}

// Deprecated: use FetchChainTraderSpotTransientOrdersV2 instead
func (c *chainClient) FetchChainTraderSpotTransientOrders(ctx context.Context, marketId, subaccountId string) (*exchangetypes.QueryTraderSpotOrdersResponse, error) {
	req := &exchangetypes.QueryTraderSpotOrdersRequest{
		MarketId:     marketId,
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.TraderSpotTransientOrders, req)

	return res, err
}

// Deprecated: use FetchSpotMidPriceAndTOBV2 instead
func (c *chainClient) FetchSpotMidPriceAndTOB(ctx context.Context, marketId string) (*exchangetypes.QuerySpotMidPriceAndTOBResponse, error) {
	req := &exchangetypes.QuerySpotMidPriceAndTOBRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.SpotMidPriceAndTOB, req)

	return res, err
}

// Deprecated: use FetchDerivativeMidPriceAndTOBV2 instead
func (c *chainClient) FetchDerivativeMidPriceAndTOB(ctx context.Context, marketId string) (*exchangetypes.QueryDerivativeMidPriceAndTOBResponse, error) {
	req := &exchangetypes.QueryDerivativeMidPriceAndTOBRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.DerivativeMidPriceAndTOB, req)

	return res, err
}

// Deprecated: use FetchChainDerivativeOrderbookV2 instead
func (c *chainClient) FetchChainDerivativeOrderbook(ctx context.Context, marketId string, limit uint64, limitCumulativeNotional sdkmath.LegacyDec) (*exchangetypes.QueryDerivativeOrderbookResponse, error) {
	req := &exchangetypes.QueryDerivativeOrderbookRequest{
		MarketId:                marketId,
		Limit:                   limit,
		LimitCumulativeNotional: &limitCumulativeNotional,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.DerivativeOrderbook, req)

	return res, err
}

// Deprecated: use FetchChainTraderDerivativeOrdersV2 instead
func (c *chainClient) FetchChainTraderDerivativeOrders(ctx context.Context, marketId, subaccountId string) (*exchangetypes.QueryTraderDerivativeOrdersResponse, error) {
	req := &exchangetypes.QueryTraderDerivativeOrdersRequest{
		MarketId:     marketId,
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.TraderDerivativeOrders, req)

	return res, err
}

// Deprecated: use FetchChainAccountAddressDerivativeOrdersV2 instead
func (c *chainClient) FetchChainAccountAddressDerivativeOrders(ctx context.Context, marketId, address string) (*exchangetypes.QueryAccountAddressDerivativeOrdersResponse, error) {
	req := &exchangetypes.QueryAccountAddressDerivativeOrdersRequest{
		MarketId:       marketId,
		AccountAddress: address,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.AccountAddressDerivativeOrders, req)

	return res, err
}

// Deprecated: use FetchChainDerivativeOrdersByHashesV2 instead
func (c *chainClient) FetchChainDerivativeOrdersByHashes(ctx context.Context, marketId, subaccountId string, orderHashes []string) (*exchangetypes.QueryDerivativeOrdersByHashesResponse, error) {
	req := &exchangetypes.QueryDerivativeOrdersByHashesRequest{
		MarketId:     marketId,
		SubaccountId: subaccountId,
		OrderHashes:  orderHashes,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.DerivativeOrdersByHashes, req)

	return res, err
}

// Deprecated: use FetchChainTraderDerivativeTransientOrdersV2 instead
func (c *chainClient) FetchChainTraderDerivativeTransientOrders(ctx context.Context, marketId, subaccountId string) (*exchangetypes.QueryTraderDerivativeOrdersResponse, error) {
	req := &exchangetypes.QueryTraderDerivativeOrdersRequest{
		MarketId:     marketId,
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.TraderDerivativeTransientOrders, req)

	return res, err
}

// Deprecated: use FetchChainDerivativeMarketsV2 instead
func (c *chainClient) FetchChainDerivativeMarkets(ctx context.Context, status string, marketIDs []string, withMidPriceAndTob bool) (*exchangetypes.QueryDerivativeMarketsResponse, error) {
	req := &exchangetypes.QueryDerivativeMarketsRequest{
		MarketIds:          marketIDs,
		WithMidPriceAndTob: withMidPriceAndTob,
	}
	if status != "" {
		req.Status = status
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.DerivativeMarkets, req)

	return res, err
}

// Deprecated: use FetchChainDerivativeMarketV2 instead
func (c *chainClient) FetchChainDerivativeMarket(ctx context.Context, marketId string) (*exchangetypes.QueryDerivativeMarketResponse, error) {
	req := &exchangetypes.QueryDerivativeMarketRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.DerivativeMarket, req)

	return res, err
}

// Deprecated: use FetchDerivativeMarketAddressV2 instead
func (c *chainClient) FetchDerivativeMarketAddress(ctx context.Context, marketId string) (*exchangetypes.QueryDerivativeMarketAddressResponse, error) {
	req := &exchangetypes.QueryDerivativeMarketAddressRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.DerivativeMarketAddress, req)

	return res, err
}

// Deprecated: use FetchSubaccountTradeNonceV2 instead
func (c *chainClient) FetchSubaccountTradeNonce(ctx context.Context, subaccountId string) (*exchangetypes.QuerySubaccountTradeNonceResponse, error) {
	req := &exchangetypes.QuerySubaccountTradeNonceRequest{
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.SubaccountTradeNonce, req)

	return res, err
}

// Deprecated: use FetchChainPositionsV2 instead
func (c *chainClient) FetchChainPositions(ctx context.Context) (*exchangetypes.QueryPositionsResponse, error) {
	req := &exchangetypes.QueryPositionsRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.Positions, req)

	return res, err
}

// Deprecated: use FetchChainSubaccountPositionsV2 instead
func (c *chainClient) FetchChainSubaccountPositions(ctx context.Context, subaccountId string) (*exchangetypes.QuerySubaccountPositionsResponse, error) {
	req := &exchangetypes.QuerySubaccountPositionsRequest{
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.SubaccountPositions, req)

	return res, err
}

// Deprecated: use FetchChainSubaccountPositionInMarketV2 instead
func (c *chainClient) FetchChainSubaccountPositionInMarket(ctx context.Context, subaccountId, marketId string) (*exchangetypes.QuerySubaccountPositionInMarketResponse, error) {
	req := &exchangetypes.QuerySubaccountPositionInMarketRequest{
		SubaccountId: subaccountId,
		MarketId:     marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.SubaccountPositionInMarket, req)

	return res, err
}

// Deprecated: use FetchChainSubaccountEffectivePositionInMarketV2 instead
func (c *chainClient) FetchChainSubaccountEffectivePositionInMarket(ctx context.Context, subaccountId, marketId string) (*exchangetypes.QuerySubaccountEffectivePositionInMarketResponse, error) {
	req := &exchangetypes.QuerySubaccountEffectivePositionInMarketRequest{
		SubaccountId: subaccountId,
		MarketId:     marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.SubaccountEffectivePositionInMarket, req)

	return res, err
}

// Deprecated: use FetchChainPerpetualMarketInfoV2 instead
func (c *chainClient) FetchChainPerpetualMarketInfo(ctx context.Context, marketId string) (*exchangetypes.QueryPerpetualMarketInfoResponse, error) {
	req := &exchangetypes.QueryPerpetualMarketInfoRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.PerpetualMarketInfo, req)

	return res, err
}

// Deprecated: use FetchChainExpiryFuturesMarketInfoV2 instead
func (c *chainClient) FetchChainExpiryFuturesMarketInfo(ctx context.Context, marketId string) (*exchangetypes.QueryExpiryFuturesMarketInfoResponse, error) {
	req := &exchangetypes.QueryExpiryFuturesMarketInfoRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.ExpiryFuturesMarketInfo, req)

	return res, err
}

// Deprecated: use FetchChainPerpetualMarketFundingV2 instead
func (c *chainClient) FetchChainPerpetualMarketFunding(ctx context.Context, marketId string) (*exchangetypes.QueryPerpetualMarketFundingResponse, error) {
	req := &exchangetypes.QueryPerpetualMarketFundingRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.PerpetualMarketFunding, req)

	return res, err
}

// Deprecated: use FetchSubaccountOrderMetadataV2 instead
func (c *chainClient) FetchSubaccountOrderMetadata(ctx context.Context, subaccountId string) (*exchangetypes.QuerySubaccountOrderMetadataResponse, error) {
	req := &exchangetypes.QuerySubaccountOrderMetadataRequest{
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.SubaccountOrderMetadata, req)

	return res, err
}

// Deprecated: use FetchTradeRewardPointsV2 instead
func (c *chainClient) FetchTradeRewardPoints(ctx context.Context, accounts []string) (*exchangetypes.QueryTradeRewardPointsResponse, error) {
	req := &exchangetypes.QueryTradeRewardPointsRequest{
		Accounts: accounts,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.TradeRewardPoints, req)

	return res, err
}

// Deprecated: use FetchPendingTradeRewardPointsV2 instead
func (c *chainClient) FetchPendingTradeRewardPoints(ctx context.Context, accounts []string) (*exchangetypes.QueryTradeRewardPointsResponse, error) {
	req := &exchangetypes.QueryTradeRewardPointsRequest{
		Accounts: accounts,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.PendingTradeRewardPoints, req)

	return res, err
}

// Deprecated: use FetchTradeRewardCampaignV2 instead
func (c *chainClient) FetchTradeRewardCampaign(ctx context.Context) (*exchangetypes.QueryTradeRewardCampaignResponse, error) {
	req := &exchangetypes.QueryTradeRewardCampaignRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.TradeRewardCampaign, req)

	return res, err
}

// Deprecated: use FetchFeeDiscountAccountInfoV2 instead
func (c *chainClient) FetchFeeDiscountAccountInfo(ctx context.Context, account string) (*exchangetypes.QueryFeeDiscountAccountInfoResponse, error) {
	req := &exchangetypes.QueryFeeDiscountAccountInfoRequest{
		Account: account,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.FeeDiscountAccountInfo, req)

	return res, err
}

// Deprecated: use FetchFeeDiscountScheduleV2 instead
func (c *chainClient) FetchFeeDiscountSchedule(ctx context.Context) (*exchangetypes.QueryFeeDiscountScheduleResponse, error) {
	req := &exchangetypes.QueryFeeDiscountScheduleRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.FeeDiscountSchedule, req)

	return res, err
}

// Deprecated: use FetchBalanceMismatchesV2 instead
func (c *chainClient) FetchBalanceMismatches(ctx context.Context, dustFactor int64) (*exchangetypes.QueryBalanceMismatchesResponse, error) {
	req := &exchangetypes.QueryBalanceMismatchesRequest{
		DustFactor: dustFactor,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.BalanceMismatches, req)

	return res, err
}

// Deprecated: use FetchBalanceWithBalanceHoldsV2 instead
func (c *chainClient) FetchBalanceWithBalanceHolds(ctx context.Context) (*exchangetypes.QueryBalanceWithBalanceHoldsResponse, error) {
	req := &exchangetypes.QueryBalanceWithBalanceHoldsRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.BalanceWithBalanceHolds, req)

	return res, err
}

// Deprecated: use FetchFeeDiscountTierStatisticsV2 instead
func (c *chainClient) FetchFeeDiscountTierStatistics(ctx context.Context) (*exchangetypes.QueryFeeDiscountTierStatisticsResponse, error) {
	req := &exchangetypes.QueryFeeDiscountTierStatisticsRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.FeeDiscountTierStatistics, req)

	return res, err
}

// Deprecated: use FetchMitoVaultInfosV2 instead
func (c *chainClient) FetchMitoVaultInfos(ctx context.Context) (*exchangetypes.MitoVaultInfosResponse, error) {
	req := &exchangetypes.MitoVaultInfosRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.MitoVaultInfos, req)

	return res, err
}

// Deprecated: use FetchMarketIDFromVaultV2 instead
func (c *chainClient) FetchMarketIDFromVault(ctx context.Context, vaultAddress string) (*exchangetypes.QueryMarketIDFromVaultResponse, error) {
	req := &exchangetypes.QueryMarketIDFromVaultRequest{
		VaultAddress: vaultAddress,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.QueryMarketIDFromVault, req)

	return res, err
}

// Deprecated: use FetchHistoricalTradeRecordsV2 instead
func (c *chainClient) FetchHistoricalTradeRecords(ctx context.Context, marketId string) (*exchangetypes.QueryHistoricalTradeRecordsResponse, error) {
	req := &exchangetypes.QueryHistoricalTradeRecordsRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.HistoricalTradeRecords, req)

	return res, err
}

// Deprecated: use FetchIsOptedOutOfRewardsV2 instead
func (c *chainClient) FetchIsOptedOutOfRewards(ctx context.Context, account string) (*exchangetypes.QueryIsOptedOutOfRewardsResponse, error) {
	req := &exchangetypes.QueryIsOptedOutOfRewardsRequest{
		Account: account,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.IsOptedOutOfRewards, req)

	return res, err
}

// Deprecated: use FetchOptedOutOfRewardsAccountsV2 instead
func (c *chainClient) FetchOptedOutOfRewardsAccounts(ctx context.Context) (*exchangetypes.QueryOptedOutOfRewardsAccountsResponse, error) {
	req := &exchangetypes.QueryOptedOutOfRewardsAccountsRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.OptedOutOfRewardsAccounts, req)

	return res, err
}

// Deprecated: use FetchMarketVolatilityV2 instead
func (c *chainClient) FetchMarketVolatility(ctx context.Context, marketId string, tradeHistoryOptions *exchangetypes.TradeHistoryOptions) (*exchangetypes.QueryMarketVolatilityResponse, error) {
	req := &exchangetypes.QueryMarketVolatilityRequest{
		MarketId:            marketId,
		TradeHistoryOptions: tradeHistoryOptions,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.MarketVolatility, req)

	return res, err
}

// Deprecated: use FetchChainBinaryOptionsMarketsV2 instead
func (c *chainClient) FetchChainBinaryOptionsMarkets(ctx context.Context, status string) (*exchangetypes.QueryBinaryMarketsResponse, error) {
	req := &exchangetypes.QueryBinaryMarketsRequest{}
	if status != "" {
		req.Status = status
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.BinaryOptionsMarkets, req)

	return res, err
}

// Deprecated: use FetchTraderDerivativeConditionalOrdersV2 instead
func (c *chainClient) FetchTraderDerivativeConditionalOrders(ctx context.Context, subaccountId, marketId string) (*exchangetypes.QueryTraderDerivativeConditionalOrdersResponse, error) {
	req := &exchangetypes.QueryTraderDerivativeConditionalOrdersRequest{
		SubaccountId: subaccountId,
		MarketId:     marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.TraderDerivativeConditionalOrders, req)

	return res, err
}

// Deprecated: use FetchMarketAtomicExecutionFeeMultiplierV2 instead
func (c *chainClient) FetchMarketAtomicExecutionFeeMultiplier(ctx context.Context, marketId string) (*exchangetypes.QueryMarketAtomicExecutionFeeMultiplierResponse, error) {
	req := &exchangetypes.QueryMarketAtomicExecutionFeeMultiplierRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.MarketAtomicExecutionFeeMultiplier, req)

	return res, err
}

// Deprecated: use FetchL3DerivativeOrderbookV2 instead
func (c *chainClient) FetchL3DerivativeOrderBook(ctx context.Context, marketId string) (*exchangetypes.QueryFullDerivativeOrderbookResponse, error) {
	req := &exchangetypes.QueryFullDerivativeOrderbookRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.L3DerivativeOrderBook, req)

	return res, err
}

// Deprecated: use FetchL3SpotOrderbookV2 instead
func (c *chainClient) FetchL3SpotOrderBook(ctx context.Context, marketId string) (*exchangetypes.QueryFullSpotOrderbookResponse, error) {
	req := &exchangetypes.QueryFullSpotOrderbookRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.L3SpotOrderBook, req)

	return res, err
}

// Deprecated: use FetchMarketBalanceV2 instead
func (c *chainClient) FetchMarketBalance(ctx context.Context, marketId string) (*exchangetypes.QueryMarketBalanceResponse, error) {
	req := &exchangetypes.QueryMarketBalanceRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.MarketBalance, req)

	return res, err
}

// Deprecated: use FetchMarketBalancesV2 instead
func (c *chainClient) FetchMarketBalances(ctx context.Context) (*exchangetypes.QueryMarketBalancesResponse, error) {
	req := &exchangetypes.QueryMarketBalancesRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.MarketBalances, req)

	return res, err
}

// Deprecated: use FetchDenomMinNotionalV2 instead
func (c *chainClient) FetchDenomMinNotional(ctx context.Context, denom string) (*exchangetypes.QueryDenomMinNotionalResponse, error) {
	req := &exchangetypes.QueryDenomMinNotionalRequest{
		Denom: denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.DenomMinNotional, req)

	return res, err
}

// Deprecated: use FetchDenomMinNotionalsV2 instead
func (c *chainClient) FetchDenomMinNotionals(ctx context.Context) (*exchangetypes.QueryDenomMinNotionalsResponse, error) {
	req := &exchangetypes.QueryDenomMinNotionalsRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.DenomMinNotionals, req)

	return res, err
}

// Chain exchange V2 module
func (c *chainClient) FetchSubaccountDepositsV2(ctx context.Context, subaccountId string) (*exchangev2types.QuerySubaccountDepositsResponse, error) {
	req := &exchangev2types.QuerySubaccountDepositsRequest{
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.SubaccountDeposits, req)

	return res, err
}

func (c *chainClient) FetchSubaccountDepositV2(ctx context.Context, subaccountId, denom string) (*exchangev2types.QuerySubaccountDepositResponse, error) {
	req := &exchangev2types.QuerySubaccountDepositRequest{
		SubaccountId: subaccountId,
		Denom:        denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.SubaccountDeposit, req)

	return res, err
}

func (c *chainClient) FetchExchangeBalancesV2(ctx context.Context) (*exchangev2types.QueryExchangeBalancesResponse, error) {
	req := &exchangev2types.QueryExchangeBalancesRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.ExchangeBalances, req)

	return res, err
}

func (c *chainClient) FetchAggregateVolumeV2(ctx context.Context, account string) (*exchangev2types.QueryAggregateVolumeResponse, error) {
	req := &exchangev2types.QueryAggregateVolumeRequest{
		Account: account,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.AggregateVolume, req)

	return res, err
}

func (c *chainClient) FetchAggregateVolumesV2(ctx context.Context, accounts, marketIDs []string) (*exchangev2types.QueryAggregateVolumesResponse, error) {
	req := &exchangev2types.QueryAggregateVolumesRequest{
		Accounts:  accounts,
		MarketIds: marketIDs,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.AggregateVolumes, req)

	return res, err
}

func (c *chainClient) FetchAggregateMarketVolumeV2(ctx context.Context, marketId string) (*exchangev2types.QueryAggregateMarketVolumeResponse, error) {
	req := &exchangev2types.QueryAggregateMarketVolumeRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.AggregateMarketVolume, req)

	return res, err
}

func (c *chainClient) FetchAggregateMarketVolumesV2(ctx context.Context, marketIDs []string) (*exchangev2types.QueryAggregateMarketVolumesResponse, error) {
	req := &exchangev2types.QueryAggregateMarketVolumesRequest{
		MarketIds: marketIDs,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.AggregateMarketVolumes, req)

	return res, err
}

func (c *chainClient) FetchDenomDecimalV2(ctx context.Context, denom string) (*exchangev2types.QueryDenomDecimalResponse, error) {
	req := &exchangev2types.QueryDenomDecimalRequest{
		Denom: denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.DenomDecimal, req)

	return res, err
}

func (c *chainClient) FetchDenomDecimalsV2(ctx context.Context, denoms []string) (*exchangev2types.QueryDenomDecimalsResponse, error) {
	req := &exchangev2types.QueryDenomDecimalsRequest{
		Denoms: denoms,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.DenomDecimals, req)

	return res, err
}

func (c *chainClient) FetchChainSpotMarketsV2(ctx context.Context, status string, marketIDs []string) (*exchangev2types.QuerySpotMarketsResponse, error) {
	req := &exchangev2types.QuerySpotMarketsRequest{
		MarketIds: marketIDs,
	}
	if status != "" {
		req.Status = status
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.SpotMarkets, req)

	return res, err
}

func (c *chainClient) FetchChainSpotMarketV2(ctx context.Context, marketId string) (*exchangev2types.QuerySpotMarketResponse, error) {
	req := &exchangev2types.QuerySpotMarketRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.SpotMarket, req)

	return res, err
}

func (c *chainClient) FetchChainFullSpotMarketsV2(ctx context.Context, status string, marketIDs []string, withMidPriceAndTob bool) (*exchangev2types.QueryFullSpotMarketsResponse, error) {
	req := &exchangev2types.QueryFullSpotMarketsRequest{
		MarketIds:          marketIDs,
		WithMidPriceAndTob: withMidPriceAndTob,
	}
	if status != "" {
		req.Status = status
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.FullSpotMarkets, req)

	return res, err
}

func (c *chainClient) FetchChainFullSpotMarketV2(ctx context.Context, marketId string, withMidPriceAndTob bool) (*exchangev2types.QueryFullSpotMarketResponse, error) {
	req := &exchangev2types.QueryFullSpotMarketRequest{
		MarketId:           marketId,
		WithMidPriceAndTob: withMidPriceAndTob,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.FullSpotMarket, req)

	return res, err
}

func (c *chainClient) FetchChainSpotOrderbookV2(ctx context.Context, marketId string, limit uint64, orderSide exchangev2types.OrderSide, limitCumulativeNotional, limitCumulativeQuantity sdkmath.LegacyDec) (*exchangev2types.QuerySpotOrderbookResponse, error) {
	req := &exchangev2types.QuerySpotOrderbookRequest{
		MarketId:                marketId,
		Limit:                   limit,
		OrderSide:               orderSide,
		LimitCumulativeNotional: &limitCumulativeNotional,
		LimitCumulativeQuantity: &limitCumulativeQuantity,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.SpotOrderbook, req)

	return res, err
}

func (c *chainClient) FetchChainTraderSpotOrdersV2(ctx context.Context, marketId, subaccountId string) (*exchangev2types.QueryTraderSpotOrdersResponse, error) {
	req := &exchangev2types.QueryTraderSpotOrdersRequest{
		MarketId:     marketId,
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.TraderSpotOrders, req)

	return res, err
}

func (c *chainClient) FetchChainAccountAddressSpotOrdersV2(ctx context.Context, marketId, address string) (*exchangev2types.QueryAccountAddressSpotOrdersResponse, error) {
	req := &exchangev2types.QueryAccountAddressSpotOrdersRequest{
		MarketId:       marketId,
		AccountAddress: address,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.AccountAddressSpotOrders, req)

	return res, err
}

func (c *chainClient) FetchChainSpotOrdersByHashesV2(ctx context.Context, marketId, subaccountId string, orderHashes []string) (*exchangev2types.QuerySpotOrdersByHashesResponse, error) {
	req := &exchangev2types.QuerySpotOrdersByHashesRequest{
		MarketId:     marketId,
		SubaccountId: subaccountId,
		OrderHashes:  orderHashes,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.SpotOrdersByHashes, req)

	return res, err
}

func (c *chainClient) FetchChainSubaccountOrdersV2(ctx context.Context, subaccountId, marketId string) (*exchangev2types.QuerySubaccountOrdersResponse, error) {
	req := &exchangev2types.QuerySubaccountOrdersRequest{
		SubaccountId: subaccountId,
		MarketId:     marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.SubaccountOrders, req)

	return res, err
}

func (c *chainClient) FetchChainTraderSpotTransientOrdersV2(ctx context.Context, marketId, subaccountId string) (*exchangev2types.QueryTraderSpotOrdersResponse, error) {
	req := &exchangev2types.QueryTraderSpotOrdersRequest{
		MarketId:     marketId,
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.TraderSpotTransientOrders, req)

	return res, err
}

func (c *chainClient) FetchSpotMidPriceAndTOBV2(ctx context.Context, marketId string) (*exchangev2types.QuerySpotMidPriceAndTOBResponse, error) {
	req := &exchangev2types.QuerySpotMidPriceAndTOBRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.SpotMidPriceAndTOB, req)

	return res, err
}

func (c *chainClient) FetchDerivativeMidPriceAndTOBV2(ctx context.Context, marketId string) (*exchangev2types.QueryDerivativeMidPriceAndTOBResponse, error) {
	req := &exchangev2types.QueryDerivativeMidPriceAndTOBRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.DerivativeMidPriceAndTOB, req)

	return res, err
}

func (c *chainClient) FetchChainDerivativeOrderbookV2(ctx context.Context, marketId string, limit uint64, limitCumulativeNotional sdkmath.LegacyDec) (*exchangev2types.QueryDerivativeOrderbookResponse, error) {
	req := &exchangev2types.QueryDerivativeOrderbookRequest{
		MarketId:                marketId,
		Limit:                   limit,
		LimitCumulativeNotional: &limitCumulativeNotional,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.DerivativeOrderbook, req)

	return res, err
}

func (c *chainClient) FetchChainTraderDerivativeOrdersV2(ctx context.Context, marketId, subaccountId string) (*exchangev2types.QueryTraderDerivativeOrdersResponse, error) {
	req := &exchangev2types.QueryTraderDerivativeOrdersRequest{
		MarketId:     marketId,
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.TraderDerivativeOrders, req)

	return res, err
}

func (c *chainClient) FetchChainAccountAddressDerivativeOrdersV2(ctx context.Context, marketId, address string) (*exchangev2types.QueryAccountAddressDerivativeOrdersResponse, error) {
	req := &exchangev2types.QueryAccountAddressDerivativeOrdersRequest{
		MarketId:       marketId,
		AccountAddress: address,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.AccountAddressDerivativeOrders, req)

	return res, err
}

func (c *chainClient) FetchChainDerivativeOrdersByHashesV2(ctx context.Context, marketId, subaccountId string, orderHashes []string) (*exchangev2types.QueryDerivativeOrdersByHashesResponse, error) {
	req := &exchangev2types.QueryDerivativeOrdersByHashesRequest{
		MarketId:     marketId,
		SubaccountId: subaccountId,
		OrderHashes:  orderHashes,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.DerivativeOrdersByHashes, req)

	return res, err
}

func (c *chainClient) FetchChainTraderDerivativeTransientOrdersV2(ctx context.Context, marketId, subaccountId string) (*exchangev2types.QueryTraderDerivativeOrdersResponse, error) {
	req := &exchangev2types.QueryTraderDerivativeOrdersRequest{
		MarketId:     marketId,
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.TraderDerivativeTransientOrders, req)

	return res, err
}

func (c *chainClient) FetchChainDerivativeMarketsV2(ctx context.Context, status string, marketIDs []string, withMidPriceAndTob bool) (*exchangev2types.QueryDerivativeMarketsResponse, error) {
	req := &exchangev2types.QueryDerivativeMarketsRequest{
		MarketIds:          marketIDs,
		WithMidPriceAndTob: withMidPriceAndTob,
	}
	if status != "" {
		req.Status = status
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.DerivativeMarkets, req)

	return res, err
}

func (c *chainClient) FetchChainDerivativeMarketV2(ctx context.Context, marketId string) (*exchangev2types.QueryDerivativeMarketResponse, error) {
	req := &exchangev2types.QueryDerivativeMarketRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.DerivativeMarket, req)

	return res, err
}

func (c *chainClient) FetchDerivativeMarketAddressV2(ctx context.Context, marketId string) (*exchangev2types.QueryDerivativeMarketAddressResponse, error) {
	req := &exchangev2types.QueryDerivativeMarketAddressRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.DerivativeMarketAddress, req)

	return res, err
}

func (c *chainClient) FetchSubaccountTradeNonceV2(ctx context.Context, subaccountId string) (*exchangev2types.QuerySubaccountTradeNonceResponse, error) {
	req := &exchangev2types.QuerySubaccountTradeNonceRequest{
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.SubaccountTradeNonce, req)

	return res, err
}

func (c *chainClient) FetchChainPositionsV2(ctx context.Context) (*exchangev2types.QueryPositionsResponse, error) {
	req := &exchangev2types.QueryPositionsRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.Positions, req)

	return res, err
}

func (c *chainClient) FetchChainSubaccountPositionsV2(ctx context.Context, subaccountId string) (*exchangev2types.QuerySubaccountPositionsResponse, error) {
	req := &exchangev2types.QuerySubaccountPositionsRequest{
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.SubaccountPositions, req)

	return res, err
}

func (c *chainClient) FetchChainSubaccountPositionInMarketV2(ctx context.Context, subaccountId, marketId string) (*exchangev2types.QuerySubaccountPositionInMarketResponse, error) {
	req := &exchangev2types.QuerySubaccountPositionInMarketRequest{
		SubaccountId: subaccountId,
		MarketId:     marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.SubaccountPositionInMarket, req)

	return res, err
}

func (c *chainClient) FetchChainSubaccountEffectivePositionInMarketV2(ctx context.Context, subaccountId, marketId string) (*exchangev2types.QuerySubaccountEffectivePositionInMarketResponse, error) {
	req := &exchangev2types.QuerySubaccountEffectivePositionInMarketRequest{
		SubaccountId: subaccountId,
		MarketId:     marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.SubaccountEffectivePositionInMarket, req)

	return res, err
}

func (c *chainClient) FetchChainPerpetualMarketInfoV2(ctx context.Context, marketId string) (*exchangev2types.QueryPerpetualMarketInfoResponse, error) {
	req := &exchangev2types.QueryPerpetualMarketInfoRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.PerpetualMarketInfo, req)

	return res, err
}

func (c *chainClient) FetchChainExpiryFuturesMarketInfoV2(ctx context.Context, marketId string) (*exchangev2types.QueryExpiryFuturesMarketInfoResponse, error) {
	req := &exchangev2types.QueryExpiryFuturesMarketInfoRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.ExpiryFuturesMarketInfo, req)

	return res, err
}

func (c *chainClient) FetchChainPerpetualMarketFundingV2(ctx context.Context, marketId string) (*exchangev2types.QueryPerpetualMarketFundingResponse, error) {
	req := &exchangev2types.QueryPerpetualMarketFundingRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.PerpetualMarketFunding, req)

	return res, err
}

func (c *chainClient) FetchSubaccountOrderMetadataV2(ctx context.Context, subaccountId string) (*exchangev2types.QuerySubaccountOrderMetadataResponse, error) {
	req := &exchangev2types.QuerySubaccountOrderMetadataRequest{
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.SubaccountOrderMetadata, req)

	return res, err
}

func (c *chainClient) FetchTradeRewardPointsV2(ctx context.Context, accounts []string) (*exchangev2types.QueryTradeRewardPointsResponse, error) {
	req := &exchangev2types.QueryTradeRewardPointsRequest{
		Accounts: accounts,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.TradeRewardPoints, req)

	return res, err
}

func (c *chainClient) FetchPendingTradeRewardPointsV2(ctx context.Context, accounts []string) (*exchangev2types.QueryTradeRewardPointsResponse, error) {
	req := &exchangev2types.QueryTradeRewardPointsRequest{
		Accounts: accounts,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.PendingTradeRewardPoints, req)

	return res, err
}

func (c *chainClient) FetchTradeRewardCampaignV2(ctx context.Context) (*exchangev2types.QueryTradeRewardCampaignResponse, error) {
	req := &exchangev2types.QueryTradeRewardCampaignRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.TradeRewardCampaign, req)

	return res, err
}

func (c *chainClient) FetchFeeDiscountAccountInfoV2(ctx context.Context, account string) (*exchangev2types.QueryFeeDiscountAccountInfoResponse, error) {
	req := &exchangev2types.QueryFeeDiscountAccountInfoRequest{
		Account: account,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.FeeDiscountAccountInfo, req)

	return res, err
}

func (c *chainClient) FetchFeeDiscountScheduleV2(ctx context.Context) (*exchangev2types.QueryFeeDiscountScheduleResponse, error) {
	req := &exchangev2types.QueryFeeDiscountScheduleRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.FeeDiscountSchedule, req)

	return res, err
}

func (c *chainClient) FetchBalanceMismatchesV2(ctx context.Context, dustFactor int64) (*exchangev2types.QueryBalanceMismatchesResponse, error) {
	req := &exchangev2types.QueryBalanceMismatchesRequest{
		DustFactor: dustFactor,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.BalanceMismatches, req)

	return res, err
}

func (c *chainClient) FetchBalanceWithBalanceHoldsV2(ctx context.Context) (*exchangev2types.QueryBalanceWithBalanceHoldsResponse, error) {
	req := &exchangev2types.QueryBalanceWithBalanceHoldsRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.BalanceWithBalanceHolds, req)

	return res, err
}

func (c *chainClient) FetchFeeDiscountTierStatisticsV2(ctx context.Context) (*exchangev2types.QueryFeeDiscountTierStatisticsResponse, error) {
	req := &exchangev2types.QueryFeeDiscountTierStatisticsRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.FeeDiscountTierStatistics, req)

	return res, err
}

func (c *chainClient) FetchMitoVaultInfosV2(ctx context.Context) (*exchangev2types.MitoVaultInfosResponse, error) {
	req := &exchangev2types.MitoVaultInfosRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.MitoVaultInfos, req)

	return res, err
}

func (c *chainClient) FetchMarketIDFromVaultV2(ctx context.Context, vaultAddress string) (*exchangev2types.QueryMarketIDFromVaultResponse, error) {
	req := &exchangev2types.QueryMarketIDFromVaultRequest{
		VaultAddress: vaultAddress,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.QueryMarketIDFromVault, req)

	return res, err
}

func (c *chainClient) FetchHistoricalTradeRecordsV2(ctx context.Context, marketId string) (*exchangev2types.QueryHistoricalTradeRecordsResponse, error) {
	req := &exchangev2types.QueryHistoricalTradeRecordsRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.HistoricalTradeRecords, req)

	return res, err
}

func (c *chainClient) FetchIsOptedOutOfRewardsV2(ctx context.Context, account string) (*exchangev2types.QueryIsOptedOutOfRewardsResponse, error) {
	req := &exchangev2types.QueryIsOptedOutOfRewardsRequest{
		Account: account,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.IsOptedOutOfRewards, req)

	return res, err
}

func (c *chainClient) FetchOptedOutOfRewardsAccountsV2(ctx context.Context) (*exchangev2types.QueryOptedOutOfRewardsAccountsResponse, error) {
	req := &exchangev2types.QueryOptedOutOfRewardsAccountsRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.OptedOutOfRewardsAccounts, req)

	return res, err
}

func (c *chainClient) FetchMarketVolatilityV2(ctx context.Context, marketId string, tradeHistoryOptions *exchangev2types.TradeHistoryOptions) (*exchangev2types.QueryMarketVolatilityResponse, error) {
	req := &exchangev2types.QueryMarketVolatilityRequest{
		MarketId:            marketId,
		TradeHistoryOptions: tradeHistoryOptions,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.MarketVolatility, req)

	return res, err
}

func (c *chainClient) FetchChainBinaryOptionsMarketsV2(ctx context.Context, status string) (*exchangev2types.QueryBinaryMarketsResponse, error) {
	req := &exchangev2types.QueryBinaryMarketsRequest{}
	if status != "" {
		req.Status = status
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.BinaryOptionsMarkets, req)

	return res, err
}

func (c *chainClient) FetchTraderDerivativeConditionalOrdersV2(ctx context.Context, subaccountId, marketId string) (*exchangev2types.QueryTraderDerivativeConditionalOrdersResponse, error) {
	req := &exchangev2types.QueryTraderDerivativeConditionalOrdersRequest{
		SubaccountId: subaccountId,
		MarketId:     marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.TraderDerivativeConditionalOrders, req)

	return res, err
}

func (c *chainClient) FetchMarketAtomicExecutionFeeMultiplierV2(ctx context.Context, marketId string) (*exchangev2types.QueryMarketAtomicExecutionFeeMultiplierResponse, error) {
	req := &exchangev2types.QueryMarketAtomicExecutionFeeMultiplierRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.MarketAtomicExecutionFeeMultiplier, req)

	return res, err
}

func (c *chainClient) FetchActiveStakeGrantV2(ctx context.Context, grantee string) (*exchangev2types.QueryActiveStakeGrantResponse, error) {
	req := &exchangev2types.QueryActiveStakeGrantRequest{
		Grantee: grantee,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.ActiveStakeGrant, req)

	return res, err
}

func (c *chainClient) FetchGrantAuthorizationV2(ctx context.Context, granter, grantee string) (*exchangev2types.QueryGrantAuthorizationResponse, error) {
	req := &exchangev2types.QueryGrantAuthorizationRequest{
		Granter: granter,
		Grantee: grantee,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.GrantAuthorization, req)

	return res, err
}

func (c *chainClient) FetchGrantAuthorizationsV2(ctx context.Context, granter string) (*exchangev2types.QueryGrantAuthorizationsResponse, error) {
	req := &exchangev2types.QueryGrantAuthorizationsRequest{
		Granter: granter,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.GrantAuthorizations, req)

	return res, err
}

func (c *chainClient) FetchL3DerivativeOrderbookV2(ctx context.Context, marketId string) (*exchangev2types.QueryFullDerivativeOrderbookResponse, error) {
	req := &exchangev2types.QueryFullDerivativeOrderbookRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.L3DerivativeOrderBook, req)

	return res, err
}

func (c *chainClient) FetchL3SpotOrderbookV2(ctx context.Context, marketId string) (*exchangev2types.QueryFullSpotOrderbookResponse, error) {
	req := &exchangev2types.QueryFullSpotOrderbookRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.L3SpotOrderBook, req)

	return res, err
}

func (c *chainClient) FetchMarketBalanceV2(ctx context.Context, marketId string) (*exchangev2types.QueryMarketBalanceResponse, error) {
	req := &exchangev2types.QueryMarketBalanceRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.MarketBalance, req)

	return res, err
}

func (c *chainClient) FetchMarketBalancesV2(ctx context.Context) (*exchangev2types.QueryMarketBalancesResponse, error) {
	req := &exchangev2types.QueryMarketBalancesRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.MarketBalances, req)

	return res, err
}

func (c *chainClient) FetchDenomMinNotionalV2(ctx context.Context, denom string) (*exchangev2types.QueryDenomMinNotionalResponse, error) {
	req := &exchangev2types.QueryDenomMinNotionalRequest{
		Denom: denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.DenomMinNotional, req)

	return res, err
}

func (c *chainClient) FetchDenomMinNotionalsV2(ctx context.Context) (*exchangev2types.QueryDenomMinNotionalsResponse, error) {
	req := &exchangev2types.QueryDenomMinNotionalsRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.DenomMinNotionals, req)

	return res, err
}

// Tendermint module

func (c *chainClient) FetchNodeInfo(ctx context.Context) (*cmtservice.GetNodeInfoResponse, error) {
	req := &cmtservice.GetNodeInfoRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tendermintQueryClient.GetNodeInfo, req)

	return res, err
}

func (c *chainClient) FetchSyncing(ctx context.Context) (*cmtservice.GetSyncingResponse, error) {
	req := &cmtservice.GetSyncingRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tendermintQueryClient.GetSyncing, req)

	return res, err
}

func (c *chainClient) FetchLatestBlock(ctx context.Context) (*cmtservice.GetLatestBlockResponse, error) {
	req := &cmtservice.GetLatestBlockRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tendermintQueryClient.GetLatestBlock, req)

	return res, err
}

func (c *chainClient) FetchBlockByHeight(ctx context.Context, height int64) (*cmtservice.GetBlockByHeightResponse, error) {
	req := &cmtservice.GetBlockByHeightRequest{
		Height: height,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tendermintQueryClient.GetBlockByHeight, req)

	return res, err
}

func (c *chainClient) FetchLatestValidatorSet(ctx context.Context) (*cmtservice.GetLatestValidatorSetResponse, error) {
	req := &cmtservice.GetLatestValidatorSetRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tendermintQueryClient.GetLatestValidatorSet, req)

	return res, err
}

func (c *chainClient) FetchValidatorSetByHeight(ctx context.Context, height int64, pagination *query.PageRequest) (*cmtservice.GetValidatorSetByHeightResponse, error) {
	req := &cmtservice.GetValidatorSetByHeightRequest{
		Height:     height,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tendermintQueryClient.GetValidatorSetByHeight, req)

	return res, err
}

func (c *chainClient) ABCIQuery(ctx context.Context, path string, data []byte, height int64, prove bool) (*cmtservice.ABCIQueryResponse, error) {
	req := &cmtservice.ABCIQueryRequest{
		Path:   path,
		Data:   data,
		Height: height,
		Prove:  prove,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tendermintQueryClient.ABCIQuery, req)

	return res, err
}

// IBC Transfer module
func (c *chainClient) FetchDenomTrace(ctx context.Context, hash string) (*ibctransfertypes.QueryDenomTraceResponse, error) {
	req := &ibctransfertypes.QueryDenomTraceRequest{
		Hash: hash,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcTransferQueryClient.DenomTrace, req)

	return res, err
}

func (c *chainClient) FetchDenomTraces(ctx context.Context, pagination *query.PageRequest) (*ibctransfertypes.QueryDenomTracesResponse, error) {
	req := &ibctransfertypes.QueryDenomTracesRequest{
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcTransferQueryClient.DenomTraces, req)

	return res, err
}

func (c *chainClient) FetchDenomHash(ctx context.Context, trace string) (*ibctransfertypes.QueryDenomHashResponse, error) {
	req := &ibctransfertypes.QueryDenomHashRequest{
		Trace: trace,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcTransferQueryClient.DenomHash, req)

	return res, err
}

func (c *chainClient) FetchEscrowAddress(ctx context.Context, portId, channelId string) (*ibctransfertypes.QueryEscrowAddressResponse, error) {
	req := &ibctransfertypes.QueryEscrowAddressRequest{
		PortId:    portId,
		ChannelId: channelId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcTransferQueryClient.EscrowAddress, req)

	return res, err
}

func (c *chainClient) FetchTotalEscrowForDenom(ctx context.Context, denom string) (*ibctransfertypes.QueryTotalEscrowForDenomResponse, error) {
	req := &ibctransfertypes.QueryTotalEscrowForDenomRequest{
		Denom: denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcTransferQueryClient.TotalEscrowForDenom, req)

	return res, err
}

// IBC Core Channel module
func (c *chainClient) FetchIBCChannel(ctx context.Context, portId, channelId string) (*ibcchanneltypes.QueryChannelResponse, error) {
	req := &ibcchanneltypes.QueryChannelRequest{
		PortId:    portId,
		ChannelId: channelId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.Channel, req)

	return res, err
}

func (c *chainClient) FetchIBCChannels(ctx context.Context, pagination *query.PageRequest) (*ibcchanneltypes.QueryChannelsResponse, error) {
	req := &ibcchanneltypes.QueryChannelsRequest{
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.Channels, req)

	return res, err
}

func (c *chainClient) FetchIBCConnectionChannels(ctx context.Context, connection string, pagination *query.PageRequest) (*ibcchanneltypes.QueryConnectionChannelsResponse, error) {
	req := &ibcchanneltypes.QueryConnectionChannelsRequest{
		Connection: connection,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.ConnectionChannels, req)

	return res, err
}

func (c *chainClient) FetchIBCChannelClientState(ctx context.Context, portId, channelId string) (*ibcchanneltypes.QueryChannelClientStateResponse, error) {
	req := &ibcchanneltypes.QueryChannelClientStateRequest{
		PortId:    portId,
		ChannelId: channelId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.ChannelClientState, req)

	return res, err
}

func (c *chainClient) FetchIBCChannelConsensusState(ctx context.Context, portId, channelId string, revisionNumber, revisionHeight uint64) (*ibcchanneltypes.QueryChannelConsensusStateResponse, error) {
	req := &ibcchanneltypes.QueryChannelConsensusStateRequest{
		PortId:         portId,
		ChannelId:      channelId,
		RevisionNumber: revisionNumber,
		RevisionHeight: revisionHeight,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.ChannelConsensusState, req)

	return res, err
}

func (c *chainClient) FetchIBCPacketCommitment(ctx context.Context, portId, channelId string, sequence uint64) (*ibcchanneltypes.QueryPacketCommitmentResponse, error) {
	req := &ibcchanneltypes.QueryPacketCommitmentRequest{
		PortId:    portId,
		ChannelId: channelId,
		Sequence:  sequence,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.PacketCommitment, req)

	return res, err
}

func (c *chainClient) FetchIBCPacketCommitments(ctx context.Context, portId, channelId string, pagination *query.PageRequest) (*ibcchanneltypes.QueryPacketCommitmentsResponse, error) {
	req := &ibcchanneltypes.QueryPacketCommitmentsRequest{
		PortId:     portId,
		ChannelId:  channelId,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.PacketCommitments, req)

	return res, err
}

func (c *chainClient) FetchIBCPacketReceipt(ctx context.Context, portId, channelId string, sequence uint64) (*ibcchanneltypes.QueryPacketReceiptResponse, error) {
	req := &ibcchanneltypes.QueryPacketReceiptRequest{
		PortId:    portId,
		ChannelId: channelId,
		Sequence:  sequence,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.PacketReceipt, req)

	return res, err
}

func (c *chainClient) FetchIBCPacketAcknowledgement(ctx context.Context, portId, channelId string, sequence uint64) (*ibcchanneltypes.QueryPacketAcknowledgementResponse, error) {
	req := &ibcchanneltypes.QueryPacketAcknowledgementRequest{
		PortId:    portId,
		ChannelId: channelId,
		Sequence:  sequence,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.PacketAcknowledgement, req)

	return res, err
}

func (c *chainClient) FetchIBCPacketAcknowledgements(ctx context.Context, portId, channelId string, packetCommitmentSequences []uint64, pagination *query.PageRequest) (*ibcchanneltypes.QueryPacketAcknowledgementsResponse, error) {
	req := &ibcchanneltypes.QueryPacketAcknowledgementsRequest{
		PortId:                    portId,
		ChannelId:                 channelId,
		Pagination:                pagination,
		PacketCommitmentSequences: packetCommitmentSequences,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.PacketAcknowledgements, req)

	return res, err
}

func (c *chainClient) FetchIBCUnreceivedPackets(ctx context.Context, portId, channelId string, packetCommitmentSequences []uint64) (*ibcchanneltypes.QueryUnreceivedPacketsResponse, error) {
	req := &ibcchanneltypes.QueryUnreceivedPacketsRequest{
		PortId:                    portId,
		ChannelId:                 channelId,
		PacketCommitmentSequences: packetCommitmentSequences,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.UnreceivedPackets, req)

	return res, err
}

func (c *chainClient) FetchIBCUnreceivedAcks(ctx context.Context, portId, channelId string, packetAckSequences []uint64) (*ibcchanneltypes.QueryUnreceivedAcksResponse, error) {
	req := &ibcchanneltypes.QueryUnreceivedAcksRequest{
		PortId:             portId,
		ChannelId:          channelId,
		PacketAckSequences: packetAckSequences,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.UnreceivedAcks, req)

	return res, err
}

func (c *chainClient) FetchIBCNextSequenceReceive(ctx context.Context, portId, channelId string) (*ibcchanneltypes.QueryNextSequenceReceiveResponse, error) {
	req := &ibcchanneltypes.QueryNextSequenceReceiveRequest{
		PortId:    portId,
		ChannelId: channelId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.NextSequenceReceive, req)

	return res, err
}

// IBC Core Chain module
func (c *chainClient) FetchIBCClientState(ctx context.Context, clientId string) (*ibcclienttypes.QueryClientStateResponse, error) {
	req := &ibcclienttypes.QueryClientStateRequest{
		ClientId: clientId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcClientQueryClient.ClientState, req)

	return res, err
}

func (c *chainClient) FetchIBCClientStates(ctx context.Context, pagination *query.PageRequest) (*ibcclienttypes.QueryClientStatesResponse, error) {
	req := &ibcclienttypes.QueryClientStatesRequest{
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcClientQueryClient.ClientStates, req)

	return res, err
}

func (c *chainClient) FetchIBCConsensusState(ctx context.Context, clientId string, revisionNumber, revisionHeight uint64, latestHeight bool) (*ibcclienttypes.QueryConsensusStateResponse, error) {
	req := &ibcclienttypes.QueryConsensusStateRequest{
		ClientId:       clientId,
		RevisionNumber: revisionNumber,
		RevisionHeight: revisionHeight,
		LatestHeight:   latestHeight,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcClientQueryClient.ConsensusState, req)

	return res, err
}

func (c *chainClient) FetchIBCConsensusStates(ctx context.Context, clientId string, pagination *query.PageRequest) (*ibcclienttypes.QueryConsensusStatesResponse, error) {
	req := &ibcclienttypes.QueryConsensusStatesRequest{
		ClientId:   clientId,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcClientQueryClient.ConsensusStates, req)

	return res, err
}

func (c *chainClient) FetchIBCConsensusStateHeights(ctx context.Context, clientId string, pagination *query.PageRequest) (*ibcclienttypes.QueryConsensusStateHeightsResponse, error) {
	req := &ibcclienttypes.QueryConsensusStateHeightsRequest{
		ClientId:   clientId,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcClientQueryClient.ConsensusStateHeights, req)

	return res, err
}

func (c *chainClient) FetchIBCClientStatus(ctx context.Context, clientId string) (*ibcclienttypes.QueryClientStatusResponse, error) {
	req := &ibcclienttypes.QueryClientStatusRequest{
		ClientId: clientId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcClientQueryClient.ClientStatus, req)

	return res, err
}

func (c *chainClient) FetchIBCClientParams(ctx context.Context) (*ibcclienttypes.QueryClientParamsResponse, error) {
	req := &ibcclienttypes.QueryClientParamsRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcClientQueryClient.ClientParams, req)

	return res, err
}

func (c *chainClient) FetchIBCUpgradedClientState(ctx context.Context) (*ibcclienttypes.QueryUpgradedClientStateResponse, error) {
	req := &ibcclienttypes.QueryUpgradedClientStateRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcClientQueryClient.UpgradedClientState, req)

	return res, err
}

func (c *chainClient) FetchIBCUpgradedConsensusState(ctx context.Context) (*ibcclienttypes.QueryUpgradedConsensusStateResponse, error) {
	req := &ibcclienttypes.QueryUpgradedConsensusStateRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcClientQueryClient.UpgradedConsensusState, req)

	return res, err
}

// IBC Core Connection module
func (c *chainClient) FetchIBCConnection(ctx context.Context, connectionId string) (*ibcconnectiontypes.QueryConnectionResponse, error) {
	req := &ibcconnectiontypes.QueryConnectionRequest{
		ConnectionId: connectionId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcConnectionQueryClient.Connection, req)

	return res, err
}

func (c *chainClient) FetchIBCConnections(ctx context.Context, pagination *query.PageRequest) (*ibcconnectiontypes.QueryConnectionsResponse, error) {
	req := &ibcconnectiontypes.QueryConnectionsRequest{
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcConnectionQueryClient.Connections, req)

	return res, err
}

func (c *chainClient) FetchIBCClientConnections(ctx context.Context, clientId string) (*ibcconnectiontypes.QueryClientConnectionsResponse, error) {
	req := &ibcconnectiontypes.QueryClientConnectionsRequest{
		ClientId: clientId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcConnectionQueryClient.ClientConnections, req)

	return res, err
}

func (c *chainClient) FetchIBCConnectionClientState(ctx context.Context, connectionId string) (*ibcconnectiontypes.QueryConnectionClientStateResponse, error) {
	req := &ibcconnectiontypes.QueryConnectionClientStateRequest{
		ConnectionId: connectionId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcConnectionQueryClient.ConnectionClientState, req)

	return res, err
}

func (c *chainClient) FetchIBCConnectionConsensusState(ctx context.Context, connectionId string, revisionNumber, revisionHeight uint64) (*ibcconnectiontypes.QueryConnectionConsensusStateResponse, error) {
	req := &ibcconnectiontypes.QueryConnectionConsensusStateRequest{
		ConnectionId:   connectionId,
		RevisionNumber: revisionNumber,
		RevisionHeight: revisionHeight,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcConnectionQueryClient.ConnectionConsensusState, req)

	return res, err
}

func (c *chainClient) FetchIBCConnectionParams(ctx context.Context) (*ibcconnectiontypes.QueryConnectionParamsResponse, error) {
	req := &ibcconnectiontypes.QueryConnectionParamsRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcConnectionQueryClient.ConnectionParams, req)

	return res, err
}

// Permissions module

func (c *chainClient) FetchPermissionsNamespaceDenoms(ctx context.Context) (*permissionstypes.QueryNamespaceDenomsResponse, error) {
	req := &permissionstypes.QueryNamespaceDenomsRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.permissionsQueryClient.NamespaceDenoms, req)

	return res, err
}

func (c *chainClient) FetchPermissionsNamespaces(ctx context.Context) (*permissionstypes.QueryNamespacesResponse, error) {
	req := &permissionstypes.QueryNamespacesRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.permissionsQueryClient.Namespaces, req)

	return res, err
}

func (c *chainClient) FetchPermissionsNamespace(ctx context.Context, denom string) (*permissionstypes.QueryNamespaceResponse, error) {
	req := &permissionstypes.QueryNamespaceRequest{
		Denom: denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.permissionsQueryClient.Namespace, req)

	return res, err
}

func (c *chainClient) FetchPermissionsRolesByActor(ctx context.Context, denom, actor string) (*permissionstypes.QueryRolesByActorResponse, error) {
	req := &permissionstypes.QueryRolesByActorRequest{
		Denom: denom,
		Actor: actor,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.permissionsQueryClient.RolesByActor, req)

	return res, err
}

func (c *chainClient) FetchPermissionsActorsByRole(ctx context.Context, denom, role string) (*permissionstypes.QueryActorsByRoleResponse, error) {
	req := &permissionstypes.QueryActorsByRoleRequest{
		Denom: denom,
		Role:  role,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.permissionsQueryClient.ActorsByRole, req)

	return res, err
}

func (c *chainClient) FetchPermissionsRoleManagers(ctx context.Context, denom string) (*permissionstypes.QueryRoleManagersResponse, error) {
	req := &permissionstypes.QueryRoleManagersRequest{
		Denom: denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.permissionsQueryClient.RoleManagers, req)

	return res, err
}

func (c *chainClient) FetchPermissionsRoleManager(ctx context.Context, denom, manager string) (*permissionstypes.QueryRoleManagerResponse, error) {
	req := &permissionstypes.QueryRoleManagerRequest{
		Denom:   denom,
		Manager: manager,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.permissionsQueryClient.RoleManager, req)

	return res, err
}

func (c *chainClient) FetchPermissionsPolicyStatuses(ctx context.Context, denom string) (*permissionstypes.QueryPolicyStatusesResponse, error) {
	req := &permissionstypes.QueryPolicyStatusesRequest{
		Denom: denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.permissionsQueryClient.PolicyStatuses, req)

	return res, err
}

func (c *chainClient) FetchPermissionsPolicyManagerCapabilities(ctx context.Context, denom string) (*permissionstypes.QueryPolicyManagerCapabilitiesResponse, error) {
	req := &permissionstypes.QueryPolicyManagerCapabilitiesRequest{
		Denom: denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.permissionsQueryClient.PolicyManagerCapabilities, req)

	return res, err
}

func (c *chainClient) FetchPermissionsVouchers(ctx context.Context, denom string) (*permissionstypes.QueryVouchersResponse, error) {
	req := &permissionstypes.QueryVouchersRequest{
		Denom: denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.permissionsQueryClient.Vouchers, req)

	return res, err
}

func (c *chainClient) FetchPermissionsVoucher(ctx context.Context, denom, address string) (*permissionstypes.QueryVoucherResponse, error) {
	req := &permissionstypes.QueryVoucherRequest{
		Denom:   denom,
		Address: address,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.permissionsQueryClient.Voucher, req)

	return res, err
}

func (c *chainClient) FetchPermissionsModuleState(ctx context.Context) (*permissionstypes.QueryModuleStateResponse, error) {
	req := &permissionstypes.QueryModuleStateRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.permissionsQueryClient.PermissionsModuleState, req)

	return res, err
}

// TxFees module
func (c *chainClient) FetchTxFeesParams(ctx context.Context) (*txfeestypes.QueryParamsResponse, error) {
	req := &txfeestypes.QueryParamsRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.txfeesQueryClient.Params, req)

	return res, err
}

func (c *chainClient) FetchEipBaseFee(ctx context.Context) (*txfeestypes.QueryEipBaseFeeResponse, error) {
	req := &txfeestypes.QueryEipBaseFeeRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.txfeesQueryClient.GetEipBaseFee, req)

	return res, err
}

func (c *chainClient) GetNetwork() common.Network {
	return c.network
}

// SyncBroadcastMsg sends Tx to chain and waits until Tx is included in block.
func (c *chainClient) SyncBroadcastMsg(msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error) {
	req, res, err := c.BroadcastMsg(txtypes.BroadcastMode_BROADCAST_MODE_SYNC, msgs...)

	if err != nil || res.TxResponse.Code != 0 {
		return res, err
	}

	awaitCtx, cancelFn := context.WithTimeout(context.Background(), defaultBroadcastTimeout)
	defer cancelFn()

	txHash, _ := hex.DecodeString(res.TxResponse.TxHash)
	t := time.NewTimer(defaultBroadcastStatusPoll)

	for {
		select {
		case <-awaitCtx.Done():
			err := errors.Wrapf(ErrTimedOut, "%s", res.TxResponse.TxHash)
			t.Stop()
			return nil, err
		case <-t.C:
			resultTx, err := c.ctx.Client.Tx(awaitCtx, txHash, false)
			if err != nil {
				if errRes := sdkclient.CheckCometError(err, req.TxBytes); errRes != nil {
					return &txtypes.BroadcastTxResponse{TxResponse: errRes}, err
				}

				t.Reset(defaultBroadcastStatusPoll)
				continue

			} else if resultTx.Height > 0 {
				resResultTx := sdk.NewResponseResultTx(resultTx, res.TxResponse.Tx, res.TxResponse.Timestamp)
				res = &txtypes.BroadcastTxResponse{TxResponse: resResultTx}
				t.Stop()
				return res, err
			}

			t.Reset(defaultBroadcastStatusPoll)
		}
	}
}

// AsyncBroadcastMsg sends Tx to chain and doesn't wait until Tx is included in block. This method
// cannot be used for rapid Tx sending, it is expected that you wait for transaction status with
// external tools. If you want sdk to wait for it, use SyncBroadcastMsg.
func (c *chainClient) AsyncBroadcastMsg(msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error) {
	_, res, err := c.BroadcastMsg(txtypes.BroadcastMode_BROADCAST_MODE_ASYNC, msgs...)
	return res, err
}

// BroadcastMsg submits a group of messages in one transaction to the chain
// The function uses the broadcast mode specified with the broadcastMode parameter
func (c *chainClient) BroadcastMsg(broadcastMode txtypes.BroadcastMode, msgs ...sdk.Msg) (*txtypes.BroadcastTxRequest, *txtypes.BroadcastTxResponse, error) {
	c.syncMux.Lock()
	defer c.syncMux.Unlock()

	sequence := c.getAccSeq()
	c.txFactory = c.txFactory.WithSequence(sequence)
	c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
	req, res, err := c.broadcastTx(c.ctx, c.txFactory, broadcastMode, msgs...)
	if err != nil {
		if c.opts.ShouldFixSequenceMismatch && strings.Contains(err.Error(), "account sequence mismatch") {
			c.syncNonce()
			sequence := c.getAccSeq()
			c.txFactory = c.txFactory.WithSequence(sequence)
			c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
			c.logger.Debugln("retrying broadcastTx with nonce", sequence)
			req, res, err = c.broadcastTx(c.ctx, c.txFactory, broadcastMode, msgs...)
		}
		if err != nil {
			resJSON, _ := json.MarshalIndent(res, "", "\t")
			c.logger.WithField("size", len(msgs)).WithError(err).Errorln("failed to asynchronously broadcast messagess:", string(resJSON))
			return nil, nil, err
		}
	}

	return req, res, nil
}

func (c *chainClient) CurrentChainGasPrice() int64 {
	gasPrice := int64(client.DefaultGasPrice)
	eipBaseFee, err := c.FetchEipBaseFee(context.Background())

	if err != nil {
		c.logger.Error("an error occurred when querying the gas price from the chain, using the default gas price")
		c.logger.Debugf("error querying the gas price from chain %s", err)
	} else {
		if !eipBaseFee.BaseFee.BaseFee.IsNil() {
			gasPrice = eipBaseFee.BaseFee.BaseFee.TruncateInt64()
		}
	}

	return gasPrice
}

func (c *chainClient) SetGasPrice(gasPrice int64) {
	gasPrices := fmt.Sprintf("%v%s", gasPrice, client.InjDenom)

	c.txFactory = c.txFactory.WithGasPrices(gasPrices)
}
