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
	"time"

	sdkmath "cosmossdk.io/math"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	log "github.com/InjectiveLabs/suplog"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
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
	gethsigner "github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	erc20types "github.com/InjectiveLabs/sdk-go/chain/erc20/types"
	evmtypes "github.com/InjectiveLabs/sdk-go/chain/evm/types"
	exchangev2types "github.com/InjectiveLabs/sdk-go/chain/exchange/types/v2"
	permissionstypes "github.com/InjectiveLabs/sdk-go/chain/permissions/types"
	chainstreamv2types "github.com/InjectiveLabs/sdk-go/chain/stream/types/v2"
	tokenfactorytypes "github.com/InjectiveLabs/sdk-go/chain/tokenfactory/types"
	txfeestypes "github.com/InjectiveLabs/sdk-go/chain/txfees/types"
	"github.com/InjectiveLabs/sdk-go/client"
	"github.com/InjectiveLabs/sdk-go/client/common"
)

type ChainClientV2 interface {
	CanSignTransactions() bool
	FromAddress() sdk.AccAddress
	QueryClient() *grpc.ClientConn
	ClientContext() sdkclient.Context
	// return account number and sequence without increasing sequence
	GetAccNonce() (accNum uint64, accSeq uint64)

	SimulateMsg(ctx context.Context, msgs ...sdk.Msg) (*txtypes.SimulateResponse, error)
	AsyncBroadcastMsg(ctx context.Context, msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error)
	SyncBroadcastMsg(ctx context.Context, pollInterval *time.Duration, maxRetries int, msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error)
	BroadcastMsg(ctx context.Context, broadcastMode txtypes.BroadcastMode, msgs ...sdk.Msg) (*txtypes.BroadcastTxRequest, *txtypes.BroadcastTxResponse, error)

	// Build signed tx with given accNum and accSeq, useful for offline siging
	// If simulate is set to false, initialGas will be used
	BuildSignedTx(ctx context.Context, accNum, accSeq, initialGas uint64, gasPrice uint64, msg ...sdk.Msg) ([]byte, error)
	SyncBroadcastSignedTx(ctx context.Context, txBytes []byte, pollInterval *time.Duration, maxRetries int) (*txtypes.BroadcastTxResponse, error)
	AsyncBroadcastSignedTx(ctx context.Context, txBytes []byte) (*txtypes.BroadcastTxResponse, error)
	BroadcastSignedTx(ctx context.Context, txBytes []byte, broadcastMode txtypes.BroadcastMode) (*txtypes.BroadcastTxResponse, error)

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

	UpdateSubaccountNonceFromChain() error
	SynchronizeSubaccountNonce(subaccountId ethcommon.Hash) error
	ComputeOrderHashes(spotOrders []exchangev2types.SpotOrder, derivativeOrders []exchangev2types.DerivativeOrder, subaccountId ethcommon.Hash) (OrderHashes, error)

	CreateSpotOrderV2(defaultSubaccountID ethcommon.Hash, d *SpotOrderData) *exchangev2types.SpotOrder
	CreateDerivativeOrderV2(defaultSubaccountID ethcommon.Hash, d *DerivativeOrderData) *exchangev2types.DerivativeOrder
	OrderCancelV2(defaultSubaccountID ethcommon.Hash, d *OrderCancelData) *exchangev2types.OrderData

	GetGasFee() (string, error)

	StreamEventOrderFail(sender string, failEventCh chan map[string]uint)
	StreamEventOrderFailWithWebsocket(sender string, websocket *rpchttp.HTTP, failEventCh chan map[string]uint)
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
	FetchSubaccountDeposits(ctx context.Context, subaccountID string) (*exchangev2types.QuerySubaccountDepositsResponse, error)
	FetchSubaccountDeposit(ctx context.Context, subaccountId string, denom string) (*exchangev2types.QuerySubaccountDepositResponse, error)
	FetchExchangeBalances(ctx context.Context) (*exchangev2types.QueryExchangeBalancesResponse, error)
	FetchAggregateVolume(ctx context.Context, account string) (*exchangev2types.QueryAggregateVolumeResponse, error)
	FetchAggregateVolumes(ctx context.Context, accounts []string, marketIDs []string) (*exchangev2types.QueryAggregateVolumesResponse, error)
	FetchAggregateMarketVolume(ctx context.Context, marketId string) (*exchangev2types.QueryAggregateMarketVolumeResponse, error)
	FetchAggregateMarketVolumes(ctx context.Context, marketIDs []string) (*exchangev2types.QueryAggregateMarketVolumesResponse, error)
	FetchDenomDecimal(ctx context.Context, denom string) (*exchangev2types.QueryDenomDecimalResponse, error)
	FetchDenomDecimals(ctx context.Context, denoms []string) (*exchangev2types.QueryDenomDecimalsResponse, error)
	FetchChainSpotMarkets(ctx context.Context, status string, marketIDs []string) (*exchangev2types.QuerySpotMarketsResponse, error)
	FetchChainSpotMarket(ctx context.Context, marketId string) (*exchangev2types.QuerySpotMarketResponse, error)
	FetchChainFullSpotMarkets(ctx context.Context, status string, marketIDs []string, withMidPriceAndTob bool) (*exchangev2types.QueryFullSpotMarketsResponse, error)
	FetchChainFullSpotMarket(ctx context.Context, marketId string, withMidPriceAndTob bool) (*exchangev2types.QueryFullSpotMarketResponse, error)
	FetchChainSpotOrderbook(ctx context.Context, marketId string, limit uint64, orderSide exchangev2types.OrderSide, limitCumulativeNotional sdkmath.LegacyDec, limitCumulativeQuantity sdkmath.LegacyDec) (*exchangev2types.QuerySpotOrderbookResponse, error)
	FetchChainTraderSpotOrders(ctx context.Context, marketId string, subaccountId string) (*exchangev2types.QueryTraderSpotOrdersResponse, error)
	FetchChainAccountAddressSpotOrders(ctx context.Context, marketId string, address string) (*exchangev2types.QueryAccountAddressSpotOrdersResponse, error)
	FetchChainSpotOrdersByHashes(ctx context.Context, marketId string, subaccountId string, orderHashes []string) (*exchangev2types.QuerySpotOrdersByHashesResponse, error)
	FetchChainSubaccountOrders(ctx context.Context, subaccountId string, marketId string) (*exchangev2types.QuerySubaccountOrdersResponse, error)
	FetchChainTraderSpotTransientOrders(ctx context.Context, marketId string, subaccountId string) (*exchangev2types.QueryTraderSpotOrdersResponse, error)
	FetchSpotMidPriceAndTOB(ctx context.Context, marketId string) (*exchangev2types.QuerySpotMidPriceAndTOBResponse, error)
	FetchDerivativeMidPriceAndTOB(ctx context.Context, marketId string) (*exchangev2types.QueryDerivativeMidPriceAndTOBResponse, error)
	FetchChainDerivativeOrderbook(ctx context.Context, marketId string, limit uint64, limitCumulativeNotional sdkmath.LegacyDec) (*exchangev2types.QueryDerivativeOrderbookResponse, error)
	FetchChainTraderDerivativeOrders(ctx context.Context, marketId string, subaccountId string) (*exchangev2types.QueryTraderDerivativeOrdersResponse, error)
	FetchChainAccountAddressDerivativeOrders(ctx context.Context, marketId string, address string) (*exchangev2types.QueryAccountAddressDerivativeOrdersResponse, error)
	FetchChainDerivativeOrdersByHashes(ctx context.Context, marketId string, subaccountId string, orderHashes []string) (*exchangev2types.QueryDerivativeOrdersByHashesResponse, error)
	FetchChainTraderDerivativeTransientOrders(ctx context.Context, marketId string, subaccountId string) (*exchangev2types.QueryTraderDerivativeOrdersResponse, error)
	FetchChainDerivativeMarkets(ctx context.Context, status string, marketIDs []string, withMidPriceAndTob bool) (*exchangev2types.QueryDerivativeMarketsResponse, error)
	FetchChainDerivativeMarket(ctx context.Context, marketId string) (*exchangev2types.QueryDerivativeMarketResponse, error)
	FetchDerivativeMarketAddress(ctx context.Context, marketId string) (*exchangev2types.QueryDerivativeMarketAddressResponse, error)
	FetchSubaccountTradeNonce(ctx context.Context, subaccountId string) (*exchangev2types.QuerySubaccountTradeNonceResponse, error)
	FetchChainPositions(ctx context.Context) (*exchangev2types.QueryPositionsResponse, error)
	FetchChainPositionsInMarket(ctx context.Context, marketId string) (*exchangev2types.QueryPositionsInMarketResponse, error)
	FetchChainSubaccountPositions(ctx context.Context, subaccountId string) (*exchangev2types.QuerySubaccountPositionsResponse, error)
	FetchChainSubaccountPositionInMarket(ctx context.Context, subaccountId string, marketId string) (*exchangev2types.QuerySubaccountPositionInMarketResponse, error)
	FetchChainSubaccountEffectivePositionInMarket(ctx context.Context, subaccountId string, marketId string) (*exchangev2types.QuerySubaccountEffectivePositionInMarketResponse, error)
	FetchChainPerpetualMarketInfo(ctx context.Context, marketId string) (*exchangev2types.QueryPerpetualMarketInfoResponse, error)
	FetchChainExpiryFuturesMarketInfo(ctx context.Context, marketId string) (*exchangev2types.QueryExpiryFuturesMarketInfoResponse, error)
	FetchChainPerpetualMarketFunding(ctx context.Context, marketId string) (*exchangev2types.QueryPerpetualMarketFundingResponse, error)
	FetchSubaccountOrderMetadata(ctx context.Context, subaccountId string) (*exchangev2types.QuerySubaccountOrderMetadataResponse, error)
	FetchTradeRewardPoints(ctx context.Context, accounts []string) (*exchangev2types.QueryTradeRewardPointsResponse, error)
	FetchPendingTradeRewardPoints(ctx context.Context, accounts []string) (*exchangev2types.QueryTradeRewardPointsResponse, error)
	FetchFeeDiscountAccountInfo(ctx context.Context, account string) (*exchangev2types.QueryFeeDiscountAccountInfoResponse, error)
	FetchTradeRewardCampaign(ctx context.Context) (*exchangev2types.QueryTradeRewardCampaignResponse, error)
	FetchFeeDiscountSchedule(ctx context.Context) (*exchangev2types.QueryFeeDiscountScheduleResponse, error)
	FetchBalanceMismatches(ctx context.Context, dustFactor int64) (*exchangev2types.QueryBalanceMismatchesResponse, error)
	FetchBalanceWithBalanceHolds(ctx context.Context) (*exchangev2types.QueryBalanceWithBalanceHoldsResponse, error)
	FetchFeeDiscountTierStatistics(ctx context.Context) (*exchangev2types.QueryFeeDiscountTierStatisticsResponse, error)
	FetchMitoVaultInfos(ctx context.Context) (*exchangev2types.MitoVaultInfosResponse, error)
	FetchMarketIDFromVault(ctx context.Context, vaultAddress string) (*exchangev2types.QueryMarketIDFromVaultResponse, error)
	FetchHistoricalTradeRecords(ctx context.Context, marketId string) (*exchangev2types.QueryHistoricalTradeRecordsResponse, error)
	FetchIsOptedOutOfRewards(ctx context.Context, account string) (*exchangev2types.QueryIsOptedOutOfRewardsResponse, error)
	FetchOptedOutOfRewardsAccounts(ctx context.Context) (*exchangev2types.QueryOptedOutOfRewardsAccountsResponse, error)
	FetchMarketVolatility(ctx context.Context, marketId string, tradeHistoryOptions *exchangev2types.TradeHistoryOptions) (*exchangev2types.QueryMarketVolatilityResponse, error)
	FetchChainBinaryOptionsMarkets(ctx context.Context, status string) (*exchangev2types.QueryBinaryMarketsResponse, error)
	FetchTraderDerivativeConditionalOrders(ctx context.Context, subaccountId string, marketId string) (*exchangev2types.QueryTraderDerivativeConditionalOrdersResponse, error)
	FetchMarketAtomicExecutionFeeMultiplier(ctx context.Context, marketId string) (*exchangev2types.QueryMarketAtomicExecutionFeeMultiplierResponse, error)
	FetchActiveStakeGrant(ctx context.Context, grantee string) (*exchangev2types.QueryActiveStakeGrantResponse, error)
	FetchGrantAuthorization(ctx context.Context, granter, grantee string) (*exchangev2types.QueryGrantAuthorizationResponse, error)
	FetchGrantAuthorizations(ctx context.Context, granter string) (*exchangev2types.QueryGrantAuthorizationsResponse, error)
	FetchL3DerivativeOrderbook(ctx context.Context, marketId string) (*exchangev2types.QueryFullDerivativeOrderbookResponse, error)
	FetchL3SpotOrderbook(ctx context.Context, marketId string) (*exchangev2types.QueryFullSpotOrderbookResponse, error)
	FetchMarketBalance(ctx context.Context, marketId string) (*exchangev2types.QueryMarketBalanceResponse, error)
	FetchMarketBalances(ctx context.Context) (*exchangev2types.QueryMarketBalancesResponse, error)
	FetchDenomMinNotional(ctx context.Context, denom string) (*exchangev2types.QueryDenomMinNotionalResponse, error)
	FetchDenomMinNotionals(ctx context.Context) (*exchangev2types.QueryDenomMinNotionalsResponse, error)

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

	// ERC20 module
	FetchAllTokenPairs(ctx context.Context, pagination *query.PageRequest) (*erc20types.QueryAllTokenPairsResponse, error)
	FetchTokenPairByDenom(ctx context.Context, bankDenom string) (*erc20types.QueryTokenPairByDenomResponse, error)
	FetchTokenPairByERC20Address(ctx context.Context, erc20Address string) (*erc20types.QueryTokenPairByERC20AddressResponse, error)

	// EVM module
	FetchEVMAccount(ctx context.Context, evmAddress string) (*evmtypes.QueryAccountResponse, error)
	FetchEVMCosmosAccount(ctx context.Context, address string) (*evmtypes.QueryCosmosAccountResponse, error)
	FetchEVMValidatorAccount(ctx context.Context, consAddress string) (*evmtypes.QueryValidatorAccountResponse, error)
	FetchEVMBalance(ctx context.Context, address string) (*evmtypes.QueryBalanceResponse, error)
	FetchEVMStorage(ctx context.Context, address string, key *string) (*evmtypes.QueryStorageResponse, error)
	FetchEVMCode(ctx context.Context, address string) (*evmtypes.QueryCodeResponse, error)
	FetchEVMBaseFee(ctx context.Context) (*evmtypes.QueryBaseFeeResponse, error)

	CurrentChainGasPrice(ctx context.Context) int64
	SetGasPrice(gasPrice int64)

	GetNetwork() common.Network
	Close()
}

var _ ChainClientV2 = &chainClientV2{}

type chainClientV2 struct {
	ctx             sdkclient.Context
	network         common.Network
	opts            *common.ClientOptions
	logger          log.Logger
	conn            *grpc.ClientConn
	chainStreamConn *grpc.ClientConn
	txFactory       tx.Factory

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
	chainStreamV2Client      chainstreamv2types.StreamClient
	distributionQueryClient  distributiontypes.QueryClient
	erc20QueryClient         erc20types.QueryClient
	evmQueryClient           evmtypes.QueryClient
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

func NewChainClientV2(
	ctx sdkclient.Context,
	network common.Network,
	options ...common.ClientOption,
) (ChainClientV2, error) {

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
	protoCodec, ok := ctx.Codec.(*codec.ProtoCodec)
	if !ok {
		return nil, errors.New("codec is not a proto codec")
	}

	var conn *grpc.ClientConn
	var err error
	stickySessionEnabled := true
	if opts.TLSCert != nil {
		conn, err = grpc.NewClient(
			network.ChainGrpcEndpoint,
			grpc.WithTransportCredentials(opts.TLSCert),
			grpc.WithContextDialer(common.DialerFunc),
			grpc.WithDefaultCallOptions(grpc.ForceCodec(protoCodec.GRPCCodec())),
		)
	} else {
		conn, err = grpc.NewClient(
			network.ChainGrpcEndpoint,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithContextDialer(common.DialerFunc),
			grpc.WithDefaultCallOptions(grpc.ForceCodec(protoCodec.GRPCCodec())),
		)
		stickySessionEnabled = false
	}
	if err != nil {
		err = errors.Wrapf(err, "failed to connect to the gRPC: %s", network.ChainGrpcEndpoint)
		return nil, err
	}

	var chainStreamConn *grpc.ClientConn
	if opts.TLSCert != nil {
		chainStreamConn, err = grpc.NewClient(
			network.ChainStreamGrpcEndpoint,
			grpc.WithTransportCredentials(opts.TLSCert),
			grpc.WithContextDialer(common.DialerFunc),
			grpc.WithDefaultCallOptions(grpc.ForceCodec(protoCodec.GRPCCodec())),
		)

	} else {
		chainStreamConn, err = grpc.NewClient(
			network.ChainStreamGrpcEndpoint,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithContextDialer(common.DialerFunc),
			grpc.WithDefaultCallOptions(grpc.ForceCodec(protoCodec.GRPCCodec())),
		)
	}
	if err != nil {
		err = errors.Wrapf(err, "failed to connect to the chain stream gRPC: %s", network.ChainStreamGrpcEndpoint)
		return nil, err
	}

	cancelCtx, cancelFn := context.WithCancel(context.Background())
	// build client
	cc := &chainClientV2{
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
		cancelCtx:       cancelCtx,
		cancelFn:        cancelFn,

		sessionEnabled: stickySessionEnabled,

		authQueryClient:          authtypes.NewQueryClient(conn),
		authzQueryClient:         authztypes.NewQueryClient(conn),
		bankQueryClient:          banktypes.NewQueryClient(conn),
		chainStreamV2Client:      chainstreamv2types.NewStreamClient(chainStreamConn),
		distributionQueryClient:  distributiontypes.NewQueryClient(conn),
		erc20QueryClient:         erc20types.NewQueryClient(conn),
		evmQueryClient:           evmtypes.NewQueryClient(conn),
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
		go cc.syncTimeoutHeight()
	}

	return cc, nil
}

func (c *chainClientV2) syncNonce() {
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

func (c *chainClientV2) syncTimeoutHeight() {
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

func (c *chainClientV2) getAccSeq() uint64 {
	defer func() {
		c.accSeq += 1
	}()
	return c.accSeq
}

func (c *chainClientV2) GetAccNonce() (accNum, accSeq uint64) {
	return c.accNum, c.accSeq
}

func (c *chainClientV2) QueryClient() *grpc.ClientConn {
	return c.conn
}

func (c *chainClientV2) ClientContext() sdkclient.Context {
	return c.ctx
}

func (c *chainClientV2) CanSignTransactions() bool {
	return c.canSign
}

func (c *chainClientV2) FromAddress() sdk.AccAddress {
	if !c.canSign {
		return sdk.AccAddress{}
	}

	return c.ctx.FromAddress
}

func (c *chainClientV2) Close() {
	if !c.canSign {
		return
	}

	if c.cancelFn != nil {
		c.cancelFn()
	}

	if c.conn != nil {
		c.conn.Close()
	}
	if c.chainStreamConn != nil {
		c.chainStreamConn.Close()
	}
}

// Bank Module

func (c *chainClientV2) GetBankBalances(ctx context.Context, address string) (*banktypes.QueryAllBalancesResponse, error) {
	req := &banktypes.QueryAllBalancesRequest{
		Address: address,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.AllBalances, req)

	return res, err
}

func (c *chainClientV2) GetBankBalance(ctx context.Context, address, denom string) (*banktypes.QueryBalanceResponse, error) {
	req := &banktypes.QueryBalanceRequest{
		Address: address,
		Denom:   denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.Balance, req)

	return res, err
}

func (c *chainClientV2) GetBankSpendableBalances(ctx context.Context, address string, pagination *query.PageRequest) (*banktypes.QuerySpendableBalancesResponse, error) {
	req := &banktypes.QuerySpendableBalancesRequest{
		Address:    address,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.SpendableBalances, req)

	return res, err
}

func (c *chainClientV2) GetBankSpendableBalancesByDenom(ctx context.Context, address, denom string) (*banktypes.QuerySpendableBalanceByDenomResponse, error) {
	req := &banktypes.QuerySpendableBalanceByDenomRequest{
		Address: address,
		Denom:   denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.SpendableBalanceByDenom, req)

	return res, err
}

func (c *chainClientV2) GetBankTotalSupply(ctx context.Context, pagination *query.PageRequest) (*banktypes.QueryTotalSupplyResponse, error) {
	req := &banktypes.QueryTotalSupplyRequest{Pagination: pagination}
	resp, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.TotalSupply, req)

	return resp, err
}

func (c *chainClientV2) GetBankSupplyOf(ctx context.Context, denom string) (*banktypes.QuerySupplyOfResponse, error) {
	req := &banktypes.QuerySupplyOfRequest{Denom: denom}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.SupplyOf, req)

	return res, err
}

func (c *chainClientV2) GetDenomMetadata(ctx context.Context, denom string) (*banktypes.QueryDenomMetadataResponse, error) {
	req := &banktypes.QueryDenomMetadataRequest{Denom: denom}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.DenomMetadata, req)

	return res, err
}

func (c *chainClientV2) GetDenomsMetadata(ctx context.Context, pagination *query.PageRequest) (*banktypes.QueryDenomsMetadataResponse, error) {
	req := &banktypes.QueryDenomsMetadataRequest{Pagination: pagination}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.DenomsMetadata, req)

	return res, err
}

func (c *chainClientV2) GetDenomOwners(ctx context.Context, denom string, pagination *query.PageRequest) (*banktypes.QueryDenomOwnersResponse, error) {
	req := &banktypes.QueryDenomOwnersRequest{
		Denom:      denom,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.DenomOwners, req)

	return res, err
}

func (c *chainClientV2) GetBankSendEnabled(ctx context.Context, denoms []string, pagination *query.PageRequest) (*banktypes.QuerySendEnabledResponse, error) {
	req := &banktypes.QuerySendEnabledRequest{
		Denoms:     denoms,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.SendEnabled, req)

	return res, err
}

// Auth Module

func (c *chainClientV2) GetAccount(ctx context.Context, address string) (*authtypes.QueryAccountResponse, error) {
	req := &authtypes.QueryAccountRequest{
		Address: address,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.authQueryClient.Account, req)

	return res, err
}

func (c *chainClientV2) SimulateMsg(ctx context.Context, msgs ...sdk.Msg) (*txtypes.SimulateResponse, error) {
	c.txFactory = c.txFactory.WithSequence(c.accSeq)
	c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
	txf, err := PrepareFactory(c.ctx, c.txFactory)
	if err != nil {
		err = errors.Wrap(err, "failed to prepareFactory")
		return nil, err
	}

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

	return simRes, nil
}

func (c *chainClientV2) BuildSignedTx(ctx context.Context, accNum, accSeq, initialGas uint64, gasPrice uint64, msgs ...sdk.Msg) ([]byte, error) {
	txf := NewTxFactory(c.ctx).WithSequence(accSeq).WithAccountNumber(accNum)
	txf = txf.WithGas(initialGas)

	gasPriceWithDenom := fmt.Sprintf("%d%s", gasPrice, client.InjDenom)
	txf = txf.WithGasPrices(gasPriceWithDenom)

	return c.buildSignedTx(ctx, txf, msgs...)
}

func (c *chainClientV2) buildSignedTx(ctx context.Context, txf tx.Factory, msgs ...sdk.Msg) ([]byte, error) {
	if c.ctx.Simulate {
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

	txf, err := PrepareFactory(c.ctx, txf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepareFactory")
	}

	txn, err := txf.BuildUnsignedTx(msgs...)
	if err != nil {
		err = errors.Wrap(err, "failed to BuildUnsignedTx")
		return nil, err
	}

	txn.SetFeeGranter(c.ctx.GetFeeGranterAddress())
	err = tx.Sign(ctx, txf, c.ctx.GetFromName(), txn, true)
	if err != nil {
		err = errors.Wrap(err, "failed to Sign Tx")
		return nil, err
	}

	return c.ctx.TxConfig.TxEncoder()(txn.GetTx())
}

func (c *chainClientV2) SyncBroadcastSignedTx(ctx context.Context, txBytes []byte, pollInterval *time.Duration, maxRetries int) (*txtypes.BroadcastTxResponse, error) {
	res, err := c.BroadcastSignedTx(ctx, txBytes, txtypes.BroadcastMode_BROADCAST_MODE_SYNC)
	if err != nil || res.TxResponse.Code != 0 {
		return res, err
	}

	txHash, _ := hex.DecodeString(res.TxResponse.TxHash)

	statusPollInterval := defaultBroadcastStatusPoll
	if pollInterval != nil {
		statusPollInterval = *pollInterval
	}

	totalAttempts := 0
	t := time.NewTimer(statusPollInterval)

	for {
		select {
		case <-ctx.Done():
			err := errors.Wrapf(ErrTimedOut, "%s", res.TxResponse.TxHash)
			t.Stop()
			return nil, err
		case <-t.C:
			totalAttempts++
			resultTx, txErr := c.ctx.Client.Tx(ctx, txHash, false)

			if txErr != nil {
				// Check if this is a fatal error that shouldn't be retried
				if errRes := sdkclient.CheckCometError(txErr, txBytes); errRes != nil {
					return &txtypes.BroadcastTxResponse{TxResponse: errRes}, txErr
				}

				// If we've reached max retries, return error
				if totalAttempts >= maxRetries {
					t.Stop()
					return nil, errors.Wrapf(txErr, "failed to get transaction after %d retries: %s", maxRetries, res.TxResponse.TxHash)
				}

				// Continue retrying with same interval
				t.Reset(statusPollInterval)
				continue
			} else if resultTx.Height > 0 {
				resResultTx := sdk.NewResponseResultTx(resultTx, res.TxResponse.Tx, res.TxResponse.Timestamp)
				res = &txtypes.BroadcastTxResponse{TxResponse: resResultTx}
				t.Stop()
				return res, err
			}

			// Transaction not yet in block, continue polling
			t.Reset(statusPollInterval)
		}
	}
}

func (c *chainClientV2) AsyncBroadcastSignedTx(ctx context.Context, txBytes []byte) (*txtypes.BroadcastTxResponse, error) {
	return c.BroadcastSignedTx(ctx, txBytes, txtypes.BroadcastMode_BROADCAST_MODE_ASYNC)
}

func (c *chainClientV2) BroadcastSignedTx(ctx context.Context, txBytes []byte, broadcastMode txtypes.BroadcastMode) (*txtypes.BroadcastTxResponse, error) {
	req := txtypes.BroadcastTxRequest{
		TxBytes: txBytes,
		Mode:    broadcastMode,
	}

	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.txClient.BroadcastTx, &req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *chainClientV2) broadcastTx(
	ctx context.Context,
	txf tx.Factory,
	broadcastMode txtypes.BroadcastMode,
	msgs ...sdk.Msg,
) (*txtypes.BroadcastTxRequest, *txtypes.BroadcastTxResponse, error) {
	txBytes, err := c.buildSignedTx(ctx, txf, msgs...)
	if err != nil {
		err = errors.Wrap(err, "failed to build signed Tx")
		return nil, nil, err
	}

	req := txtypes.BroadcastTxRequest{
		TxBytes: txBytes,
		Mode:    broadcastMode,
	}

	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.txClient.BroadcastTx, &req)
	return &req, res, err

}

func (c *chainClientV2) GetGasFee() (string, error) {
	gasPrices := strings.TrimSuffix(c.txFactory.GasPrices().String(), client.InjDenom)

	gas, err := strconv.ParseFloat(gasPrices, 64)

	if err != nil {
		return "", err
	}

	gasFeeAdjusted := gas * float64(c.gasWanted) / math.Pow(10, 18)
	gasFeeFormatted := strconv.FormatFloat(gasFeeAdjusted, 'f', -1, 64)
	c.gasFee = gasFeeFormatted

	return c.gasFee, err
}

func (c *chainClientV2) DefaultSubaccount(acc sdk.AccAddress) ethcommon.Hash {
	return c.Subaccount(acc, 0)
}

func (c *chainClientV2) Subaccount(account sdk.AccAddress, index int) ethcommon.Hash {
	ethAddress := ethcommon.BytesToAddress(account.Bytes())
	ethLowerAddress := strings.ToLower(ethAddress.String())

	subaccountId := fmt.Sprintf("%s%024x", ethLowerAddress, index)
	return ethcommon.HexToHash(subaccountId)
}

func (c *chainClientV2) CreateSpotOrderV2(defaultSubaccountID ethcommon.Hash, d *SpotOrderData) *exchangev2types.SpotOrder {
	return &exchangev2types.SpotOrder{
		MarketId:  d.MarketId,
		OrderType: exchangev2types.OrderType(d.OrderType),
		OrderInfo: exchangev2types.OrderInfo{
			SubaccountId: defaultSubaccountID.Hex(),
			FeeRecipient: d.FeeRecipient,
			Price:        sdkmath.LegacyMustNewDecFromStr(d.Price.String()),
			Quantity:     sdkmath.LegacyMustNewDecFromStr(d.Quantity.String()),
			Cid:          d.Cid,
		},
		ExpirationBlock: d.ExpirationBlock,
	}
}

func (c *chainClientV2) CreateDerivativeOrderV2(defaultSubaccountID ethcommon.Hash, d *DerivativeOrderData) *exchangev2types.DerivativeOrder {
	orderMargin := sdkmath.LegacyMustNewDecFromStr("0")

	if !d.IsReduceOnly {
		orderMargin = sdkmath.LegacyMustNewDecFromStr(d.Quantity.Mul(d.Price).Div(d.Leverage).String())
	}

	return &exchangev2types.DerivativeOrder{
		MarketId:  d.MarketId,
		OrderType: exchangev2types.OrderType(d.OrderType),
		Margin:    orderMargin,
		OrderInfo: exchangev2types.OrderInfo{
			SubaccountId: defaultSubaccountID.Hex(),
			FeeRecipient: d.FeeRecipient,
			Price:        sdkmath.LegacyMustNewDecFromStr(d.Price.String()),
			Quantity:     sdkmath.LegacyMustNewDecFromStr(d.Quantity.String()),
			Cid:          d.Cid,
		},
		ExpirationBlock: d.ExpirationBlock,
	}
}

func (c *chainClientV2) OrderCancelV2(defaultSubaccountID ethcommon.Hash, d *OrderCancelData) *exchangev2types.OrderData {
	return &exchangev2types.OrderData{
		MarketId:     d.MarketId,
		OrderHash:    d.OrderHash,
		SubaccountId: defaultSubaccountID.Hex(),
		Cid:          d.Cid,
	}
}

func (c *chainClientV2) GetAuthzGrants(ctx context.Context, req authztypes.QueryGrantsRequest) (*authztypes.QueryGrantsResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.authzQueryClient.Grants, &req)

	return res, err
}

func (c *chainClientV2) BuildGenericAuthz(granter, grantee, msgtype string, expireIn time.Time) *authztypes.MsgGrant {
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

var (
	CreateSpotLimitOrderAuthzV2       = ExchangeAuthz("/" + proto.MessageName(&exchangev2types.CreateSpotLimitOrderAuthz{}))
	CreateSpotMarketOrderAuthzV2      = ExchangeAuthz("/" + proto.MessageName(&exchangev2types.CreateSpotMarketOrderAuthz{}))
	BatchCreateSpotLimitOrdersAuthzV2 = ExchangeAuthz("/" + proto.MessageName(&exchangev2types.BatchCreateSpotLimitOrdersAuthz{}))
	CancelSpotOrderAuthzV2            = ExchangeAuthz("/" + proto.MessageName(&exchangev2types.CancelSpotOrderAuthz{}))
	BatchCancelSpotOrdersAuthzV2      = ExchangeAuthz("/" + proto.MessageName(&exchangev2types.BatchCancelSpotOrdersAuthz{}))

	CreateDerivativeLimitOrderAuthzV2       = ExchangeAuthz("/" + proto.MessageName(&exchangev2types.CreateDerivativeLimitOrderAuthz{}))
	CreateDerivativeMarketOrderAuthzV2      = ExchangeAuthz("/" + proto.MessageName(&exchangev2types.CreateDerivativeMarketOrderAuthz{}))
	BatchCreateDerivativeLimitOrdersAuthzV2 = ExchangeAuthz("/" + proto.MessageName(&exchangev2types.BatchCreateDerivativeLimitOrdersAuthz{}))
	CancelDerivativeOrderAuthzV2            = ExchangeAuthz("/" + proto.MessageName(&exchangev2types.CancelDerivativeOrderAuthz{}))
	BatchCancelDerivativeOrdersAuthzV2      = ExchangeAuthz("/" + proto.MessageName(&exchangev2types.BatchCancelDerivativeOrdersAuthz{}))

	BatchUpdateOrdersAuthzV2 = ExchangeAuthz("/" + proto.MessageName(&exchangev2types.BatchUpdateOrdersAuthz{}))
)

func (c *chainClientV2) BuildExchangeAuthz(granter, grantee string, authzType ExchangeAuthz, subaccountId string, markets []string, expireIn time.Time) *authztypes.MsgGrant {
	if c.ofacChecker.IsBlacklisted(granter) {
		panic("Address is in the OFAC list") // panics should generally be avoided, but otherwise function signature should be changed
	}
	var typedAuthzAny codectypes.Any
	var typedAuthzBytes []byte
	switch authzType {
	// spot msgs
	case CreateSpotLimitOrderAuthzV2:
		typedAuthz := &exchangev2types.CreateSpotLimitOrderAuthz{
			SubaccountId: subaccountId,
			MarketIds:    markets,
		}
		typedAuthzBytes, _ = typedAuthz.Marshal()
	case CreateSpotMarketOrderAuthzV2:
		typedAuthz := &exchangev2types.CreateSpotMarketOrderAuthz{
			SubaccountId: subaccountId,
			MarketIds:    markets,
		}
		typedAuthzBytes, _ = typedAuthz.Marshal()
	case BatchCreateSpotLimitOrdersAuthzV2:
		typedAuthz := &exchangev2types.BatchCreateSpotLimitOrdersAuthz{
			SubaccountId: subaccountId,
			MarketIds:    markets,
		}
		typedAuthzBytes, _ = typedAuthz.Marshal()
	case CancelSpotOrderAuthzV2:
		typedAuthz := &exchangev2types.CancelSpotOrderAuthz{
			SubaccountId: subaccountId,
			MarketIds:    markets,
		}
		typedAuthzBytes, _ = typedAuthz.Marshal()
	case BatchCancelSpotOrdersAuthzV2:
		typedAuthz := &exchangev2types.BatchCancelSpotOrdersAuthz{
			SubaccountId: subaccountId,
			MarketIds:    markets,
		}
		typedAuthzBytes, _ = typedAuthz.Marshal()
	// derivative msgs
	case CreateDerivativeLimitOrderAuthzV2:
		typedAuthz := &exchangev2types.CreateDerivativeLimitOrderAuthz{
			SubaccountId: subaccountId,
			MarketIds:    markets,
		}
		typedAuthzBytes, _ = typedAuthz.Marshal()
	case CreateDerivativeMarketOrderAuthzV2:
		typedAuthz := &exchangev2types.CreateDerivativeMarketOrderAuthz{
			SubaccountId: subaccountId,
			MarketIds:    markets,
		}
		typedAuthzBytes, _ = typedAuthz.Marshal()
	case BatchCreateDerivativeLimitOrdersAuthzV2:
		typedAuthz := &exchangev2types.BatchCreateDerivativeLimitOrdersAuthz{
			SubaccountId: subaccountId,
			MarketIds:    markets,
		}
		typedAuthzBytes, _ = typedAuthz.Marshal()
	case CancelDerivativeOrderAuthzV2:
		typedAuthz := &exchangev2types.CancelDerivativeOrderAuthz{
			SubaccountId: subaccountId,
			MarketIds:    markets,
		}
		typedAuthzBytes, _ = typedAuthz.Marshal()
	case BatchCancelDerivativeOrdersAuthzV2:
		typedAuthz := &exchangev2types.BatchCancelDerivativeOrdersAuthz{
			SubaccountId: subaccountId,
			MarketIds:    markets,
		}
		typedAuthzBytes, _ = typedAuthz.Marshal()
	// common msgs
	case BatchUpdateOrdersAuthzV2:
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

func (c *chainClientV2) BuildExchangeBatchUpdateOrdersAuthz(
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

func (c *chainClientV2) StreamEventOrderFail(sender string, failEventCh chan map[string]uint) {
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

func (c *chainClientV2) StreamEventOrderFailWithWebsocket(sender string, websocket *rpchttp.HTTP, failEventCh chan map[string]uint) {
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

func (c *chainClientV2) GetTx(ctx context.Context, txHash string) (*txtypes.GetTxResponse, error) {
	req := &txtypes.GetTxRequest{
		Hash: txHash,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.txClient.GetTx, req)

	return res, err
}

func (c *chainClientV2) ChainStreamV2(ctx context.Context, req chainstreamv2types.StreamRequest) (chainstreamv2types.Stream_StreamV2Client, error) {
	stream, err := common.ExecuteStreamCall(ctx, c.network.ChainCookieAssistant, c.chainStreamV2Client.StreamV2, &req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

// wasm module

func (c *chainClientV2) FetchContractInfo(ctx context.Context, address string) (*wasmtypes.QueryContractInfoResponse, error) {
	req := &wasmtypes.QueryContractInfoRequest{
		Address: address,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.wasmQueryClient.ContractInfo, req)

	return res, err
}

func (c *chainClientV2) FetchContractHistory(ctx context.Context, address string, pagination *query.PageRequest) (*wasmtypes.QueryContractHistoryResponse, error) {
	req := &wasmtypes.QueryContractHistoryRequest{
		Address:    address,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.wasmQueryClient.ContractHistory, req)

	return res, err
}

func (c *chainClientV2) FetchContractsByCode(ctx context.Context, codeId uint64, pagination *query.PageRequest) (*wasmtypes.QueryContractsByCodeResponse, error) {
	req := &wasmtypes.QueryContractsByCodeRequest{
		CodeId:     codeId,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.wasmQueryClient.ContractsByCode, req)

	return res, err
}

func (c *chainClientV2) FetchAllContractsState(ctx context.Context, address string, pagination *query.PageRequest) (*wasmtypes.QueryAllContractStateResponse, error) {
	req := &wasmtypes.QueryAllContractStateRequest{
		Address:    address,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.wasmQueryClient.AllContractState, req)

	return res, err
}

func (c *chainClientV2) RawContractState(
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

func (c *chainClientV2) SmartContractState(
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

func (c *chainClientV2) FetchCode(ctx context.Context, codeId uint64) (*wasmtypes.QueryCodeResponse, error) {
	req := &wasmtypes.QueryCodeRequest{
		CodeId: codeId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.wasmQueryClient.Code, req)

	return res, err
}

func (c *chainClientV2) FetchCodes(ctx context.Context, pagination *query.PageRequest) (*wasmtypes.QueryCodesResponse, error) {
	req := &wasmtypes.QueryCodesRequest{
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.wasmQueryClient.Codes, req)

	return res, err
}

func (c *chainClientV2) FetchPinnedCodes(ctx context.Context, pagination *query.PageRequest) (*wasmtypes.QueryPinnedCodesResponse, error) {
	req := &wasmtypes.QueryPinnedCodesRequest{
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.wasmQueryClient.PinnedCodes, req)

	return res, err
}

func (c *chainClientV2) FetchContractsByCreator(ctx context.Context, creator string, pagination *query.PageRequest) (*wasmtypes.QueryContractsByCreatorResponse, error) {
	req := &wasmtypes.QueryContractsByCreatorRequest{
		CreatorAddress: creator,
		Pagination:     pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.wasmQueryClient.ContractsByCreator, req)

	return res, err
}

// Tokenfactory module

func (c *chainClientV2) FetchDenomAuthorityMetadata(ctx context.Context, creator, subDenom string) (*tokenfactorytypes.QueryDenomAuthorityMetadataResponse, error) {
	req := &tokenfactorytypes.QueryDenomAuthorityMetadataRequest{
		Creator: creator,
	}

	if subDenom != "" {
		req.SubDenom = subDenom
	}

	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tokenfactoryQueryClient.DenomAuthorityMetadata, req)

	return res, err
}

func (c *chainClientV2) FetchDenomsFromCreator(ctx context.Context, creator string) (*tokenfactorytypes.QueryDenomsFromCreatorResponse, error) {
	req := &tokenfactorytypes.QueryDenomsFromCreatorRequest{
		Creator: creator,
	}

	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tokenfactoryQueryClient.DenomsFromCreator, req)

	return res, err
}

func (c *chainClientV2) FetchTokenfactoryModuleState(ctx context.Context) (*tokenfactorytypes.QueryModuleStateResponse, error) {
	req := &tokenfactorytypes.QueryModuleStateRequest{}

	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tokenfactoryQueryClient.TokenfactoryModuleState, req)

	return res, err
}

// Distribution module
func (c *chainClientV2) FetchValidatorDistributionInfo(ctx context.Context, validatorAddress string) (*distributiontypes.QueryValidatorDistributionInfoResponse, error) {
	req := &distributiontypes.QueryValidatorDistributionInfoRequest{
		ValidatorAddress: validatorAddress,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.distributionQueryClient.ValidatorDistributionInfo, req)

	return res, err
}

func (c *chainClientV2) FetchValidatorOutstandingRewards(ctx context.Context, validatorAddress string) (*distributiontypes.QueryValidatorOutstandingRewardsResponse, error) {
	req := &distributiontypes.QueryValidatorOutstandingRewardsRequest{
		ValidatorAddress: validatorAddress,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.distributionQueryClient.ValidatorOutstandingRewards, req)

	return res, err
}

func (c *chainClientV2) FetchValidatorCommission(ctx context.Context, validatorAddress string) (*distributiontypes.QueryValidatorCommissionResponse, error) {
	req := &distributiontypes.QueryValidatorCommissionRequest{
		ValidatorAddress: validatorAddress,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.distributionQueryClient.ValidatorCommission, req)

	return res, err
}

func (c *chainClientV2) FetchValidatorSlashes(ctx context.Context, validatorAddress string, startingHeight, endingHeight uint64, pagination *query.PageRequest) (*distributiontypes.QueryValidatorSlashesResponse, error) {
	req := &distributiontypes.QueryValidatorSlashesRequest{
		ValidatorAddress: validatorAddress,
		StartingHeight:   startingHeight,
		EndingHeight:     endingHeight,
		Pagination:       pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.distributionQueryClient.ValidatorSlashes, req)

	return res, err
}

func (c *chainClientV2) FetchDelegationRewards(ctx context.Context, delegatorAddress, validatorAddress string) (*distributiontypes.QueryDelegationRewardsResponse, error) {
	req := &distributiontypes.QueryDelegationRewardsRequest{
		DelegatorAddress: delegatorAddress,
		ValidatorAddress: validatorAddress,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.distributionQueryClient.DelegationRewards, req)

	return res, err
}

func (c *chainClientV2) FetchDelegationTotalRewards(ctx context.Context, delegatorAddress string) (*distributiontypes.QueryDelegationTotalRewardsResponse, error) {
	req := &distributiontypes.QueryDelegationTotalRewardsRequest{
		DelegatorAddress: delegatorAddress,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.distributionQueryClient.DelegationTotalRewards, req)

	return res, err
}

func (c *chainClientV2) FetchDelegatorValidators(ctx context.Context, delegatorAddress string) (*distributiontypes.QueryDelegatorValidatorsResponse, error) {
	req := &distributiontypes.QueryDelegatorValidatorsRequest{
		DelegatorAddress: delegatorAddress,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.distributionQueryClient.DelegatorValidators, req)

	return res, err
}

func (c *chainClientV2) FetchDelegatorWithdrawAddress(ctx context.Context, delegatorAddress string) (*distributiontypes.QueryDelegatorWithdrawAddressResponse, error) {
	req := &distributiontypes.QueryDelegatorWithdrawAddressRequest{
		DelegatorAddress: delegatorAddress,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.distributionQueryClient.DelegatorWithdrawAddress, req)

	return res, err
}

func (c *chainClientV2) FetchCommunityPool(ctx context.Context) (*distributiontypes.QueryCommunityPoolResponse, error) {
	req := &distributiontypes.QueryCommunityPoolRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.distributionQueryClient.CommunityPool, req)

	return res, err
}

// Chain exchange module

func (c *chainClientV2) FetchSubaccountDeposits(ctx context.Context, subaccountId string) (*exchangev2types.QuerySubaccountDepositsResponse, error) {
	req := &exchangev2types.QuerySubaccountDepositsRequest{
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.SubaccountDeposits, req)

	return res, err
}

func (c *chainClientV2) FetchSubaccountDeposit(ctx context.Context, subaccountId, denom string) (*exchangev2types.QuerySubaccountDepositResponse, error) {
	req := &exchangev2types.QuerySubaccountDepositRequest{
		SubaccountId: subaccountId,
		Denom:        denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.SubaccountDeposit, req)

	return res, err
}

func (c *chainClientV2) FetchExchangeBalances(ctx context.Context) (*exchangev2types.QueryExchangeBalancesResponse, error) {
	req := &exchangev2types.QueryExchangeBalancesRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.ExchangeBalances, req)

	return res, err
}

func (c *chainClientV2) FetchAggregateVolume(ctx context.Context, account string) (*exchangev2types.QueryAggregateVolumeResponse, error) {
	req := &exchangev2types.QueryAggregateVolumeRequest{
		Account: account,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.AggregateVolume, req)

	return res, err
}

func (c *chainClientV2) FetchAggregateVolumes(ctx context.Context, accounts, marketIDs []string) (*exchangev2types.QueryAggregateVolumesResponse, error) {
	req := &exchangev2types.QueryAggregateVolumesRequest{
		Accounts:  accounts,
		MarketIds: marketIDs,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.AggregateVolumes, req)

	return res, err
}

func (c *chainClientV2) FetchAggregateMarketVolume(ctx context.Context, marketId string) (*exchangev2types.QueryAggregateMarketVolumeResponse, error) {
	req := &exchangev2types.QueryAggregateMarketVolumeRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.AggregateMarketVolume, req)

	return res, err
}

func (c *chainClientV2) FetchAggregateMarketVolumes(ctx context.Context, marketIDs []string) (*exchangev2types.QueryAggregateMarketVolumesResponse, error) {
	req := &exchangev2types.QueryAggregateMarketVolumesRequest{
		MarketIds: marketIDs,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.AggregateMarketVolumes, req)

	return res, err
}

func (c *chainClientV2) FetchDenomDecimal(ctx context.Context, denom string) (*exchangev2types.QueryDenomDecimalResponse, error) {
	req := &exchangev2types.QueryDenomDecimalRequest{
		Denom: denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.DenomDecimal, req)

	return res, err
}

func (c *chainClientV2) FetchDenomDecimals(ctx context.Context, denoms []string) (*exchangev2types.QueryDenomDecimalsResponse, error) {
	req := &exchangev2types.QueryDenomDecimalsRequest{
		Denoms: denoms,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.DenomDecimals, req)

	return res, err
}

func (c *chainClientV2) FetchChainSpotMarkets(ctx context.Context, status string, marketIDs []string) (*exchangev2types.QuerySpotMarketsResponse, error) {
	req := &exchangev2types.QuerySpotMarketsRequest{
		MarketIds: marketIDs,
	}
	if status != "" {
		req.Status = status
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.SpotMarkets, req)

	return res, err
}

func (c *chainClientV2) FetchChainSpotMarket(ctx context.Context, marketId string) (*exchangev2types.QuerySpotMarketResponse, error) {
	req := &exchangev2types.QuerySpotMarketRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.SpotMarket, req)

	return res, err
}

func (c *chainClientV2) FetchChainFullSpotMarkets(ctx context.Context, status string, marketIDs []string, withMidPriceAndTob bool) (*exchangev2types.QueryFullSpotMarketsResponse, error) {
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

func (c *chainClientV2) FetchChainFullSpotMarket(ctx context.Context, marketId string, withMidPriceAndTob bool) (*exchangev2types.QueryFullSpotMarketResponse, error) {
	req := &exchangev2types.QueryFullSpotMarketRequest{
		MarketId:           marketId,
		WithMidPriceAndTob: withMidPriceAndTob,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.FullSpotMarket, req)

	return res, err
}

func (c *chainClientV2) FetchChainSpotOrderbook(ctx context.Context, marketId string, limit uint64, orderSide exchangev2types.OrderSide, limitCumulativeNotional, limitCumulativeQuantity sdkmath.LegacyDec) (*exchangev2types.QuerySpotOrderbookResponse, error) {
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

func (c *chainClientV2) FetchChainTraderSpotOrders(ctx context.Context, marketId, subaccountId string) (*exchangev2types.QueryTraderSpotOrdersResponse, error) {
	req := &exchangev2types.QueryTraderSpotOrdersRequest{
		MarketId:     marketId,
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.TraderSpotOrders, req)

	return res, err
}

func (c *chainClientV2) FetchChainAccountAddressSpotOrders(ctx context.Context, marketId, address string) (*exchangev2types.QueryAccountAddressSpotOrdersResponse, error) {
	req := &exchangev2types.QueryAccountAddressSpotOrdersRequest{
		MarketId:       marketId,
		AccountAddress: address,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.AccountAddressSpotOrders, req)

	return res, err
}

func (c *chainClientV2) FetchChainSpotOrdersByHashes(ctx context.Context, marketId, subaccountId string, orderHashes []string) (*exchangev2types.QuerySpotOrdersByHashesResponse, error) {
	req := &exchangev2types.QuerySpotOrdersByHashesRequest{
		MarketId:     marketId,
		SubaccountId: subaccountId,
		OrderHashes:  orderHashes,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.SpotOrdersByHashes, req)

	return res, err
}

func (c *chainClientV2) FetchChainSubaccountOrders(ctx context.Context, subaccountId, marketId string) (*exchangev2types.QuerySubaccountOrdersResponse, error) {
	req := &exchangev2types.QuerySubaccountOrdersRequest{
		SubaccountId: subaccountId,
		MarketId:     marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.SubaccountOrders, req)

	return res, err
}

func (c *chainClientV2) FetchChainTraderSpotTransientOrders(ctx context.Context, marketId, subaccountId string) (*exchangev2types.QueryTraderSpotOrdersResponse, error) {
	req := &exchangev2types.QueryTraderSpotOrdersRequest{
		MarketId:     marketId,
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.TraderSpotTransientOrders, req)

	return res, err
}

func (c *chainClientV2) FetchSpotMidPriceAndTOB(ctx context.Context, marketId string) (*exchangev2types.QuerySpotMidPriceAndTOBResponse, error) {
	req := &exchangev2types.QuerySpotMidPriceAndTOBRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.SpotMidPriceAndTOB, req)

	return res, err
}

func (c *chainClientV2) FetchDerivativeMidPriceAndTOB(ctx context.Context, marketId string) (*exchangev2types.QueryDerivativeMidPriceAndTOBResponse, error) {
	req := &exchangev2types.QueryDerivativeMidPriceAndTOBRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.DerivativeMidPriceAndTOB, req)

	return res, err
}

func (c *chainClientV2) FetchChainDerivativeOrderbook(ctx context.Context, marketId string, limit uint64, limitCumulativeNotional sdkmath.LegacyDec) (*exchangev2types.QueryDerivativeOrderbookResponse, error) {
	req := &exchangev2types.QueryDerivativeOrderbookRequest{
		MarketId:                marketId,
		Limit:                   limit,
		LimitCumulativeNotional: &limitCumulativeNotional,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.DerivativeOrderbook, req)

	return res, err
}

func (c *chainClientV2) FetchChainTraderDerivativeOrders(ctx context.Context, marketId, subaccountId string) (*exchangev2types.QueryTraderDerivativeOrdersResponse, error) {
	req := &exchangev2types.QueryTraderDerivativeOrdersRequest{
		MarketId:     marketId,
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.TraderDerivativeOrders, req)

	return res, err
}

func (c *chainClientV2) FetchChainAccountAddressDerivativeOrders(ctx context.Context, marketId, address string) (*exchangev2types.QueryAccountAddressDerivativeOrdersResponse, error) {
	req := &exchangev2types.QueryAccountAddressDerivativeOrdersRequest{
		MarketId:       marketId,
		AccountAddress: address,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.AccountAddressDerivativeOrders, req)

	return res, err
}

func (c *chainClientV2) FetchChainDerivativeOrdersByHashes(ctx context.Context, marketId, subaccountId string, orderHashes []string) (*exchangev2types.QueryDerivativeOrdersByHashesResponse, error) {
	req := &exchangev2types.QueryDerivativeOrdersByHashesRequest{
		MarketId:     marketId,
		SubaccountId: subaccountId,
		OrderHashes:  orderHashes,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.DerivativeOrdersByHashes, req)

	return res, err
}

func (c *chainClientV2) FetchChainTraderDerivativeTransientOrders(ctx context.Context, marketId, subaccountId string) (*exchangev2types.QueryTraderDerivativeOrdersResponse, error) {
	req := &exchangev2types.QueryTraderDerivativeOrdersRequest{
		MarketId:     marketId,
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.TraderDerivativeTransientOrders, req)

	return res, err
}

func (c *chainClientV2) FetchChainDerivativeMarkets(ctx context.Context, status string, marketIDs []string, withMidPriceAndTob bool) (*exchangev2types.QueryDerivativeMarketsResponse, error) {
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

func (c *chainClientV2) FetchChainDerivativeMarket(ctx context.Context, marketId string) (*exchangev2types.QueryDerivativeMarketResponse, error) {
	req := &exchangev2types.QueryDerivativeMarketRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.DerivativeMarket, req)

	return res, err
}

func (c *chainClientV2) FetchDerivativeMarketAddress(ctx context.Context, marketId string) (*exchangev2types.QueryDerivativeMarketAddressResponse, error) {
	req := &exchangev2types.QueryDerivativeMarketAddressRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.DerivativeMarketAddress, req)

	return res, err
}

func (c *chainClientV2) FetchSubaccountTradeNonce(ctx context.Context, subaccountId string) (*exchangev2types.QuerySubaccountTradeNonceResponse, error) {
	req := &exchangev2types.QuerySubaccountTradeNonceRequest{
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.SubaccountTradeNonce, req)

	return res, err
}

func (c *chainClientV2) FetchChainPositions(ctx context.Context) (*exchangev2types.QueryPositionsResponse, error) {
	req := &exchangev2types.QueryPositionsRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.Positions, req)

	return res, err
}

func (c *chainClientV2) FetchChainPositionsInMarket(ctx context.Context, marketId string) (*exchangev2types.QueryPositionsInMarketResponse, error) {
	req := &exchangev2types.QueryPositionsInMarketRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.PositionsInMarket, req)

	return res, err
}

func (c *chainClientV2) FetchChainSubaccountPositions(ctx context.Context, subaccountId string) (*exchangev2types.QuerySubaccountPositionsResponse, error) {
	req := &exchangev2types.QuerySubaccountPositionsRequest{
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.SubaccountPositions, req)

	return res, err
}

func (c *chainClientV2) FetchChainSubaccountPositionInMarket(ctx context.Context, subaccountId, marketId string) (*exchangev2types.QuerySubaccountPositionInMarketResponse, error) {
	req := &exchangev2types.QuerySubaccountPositionInMarketRequest{
		SubaccountId: subaccountId,
		MarketId:     marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.SubaccountPositionInMarket, req)

	return res, err
}

func (c *chainClientV2) FetchChainSubaccountEffectivePositionInMarket(ctx context.Context, subaccountId, marketId string) (*exchangev2types.QuerySubaccountEffectivePositionInMarketResponse, error) {
	req := &exchangev2types.QuerySubaccountEffectivePositionInMarketRequest{
		SubaccountId: subaccountId,
		MarketId:     marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.SubaccountEffectivePositionInMarket, req)

	return res, err
}

func (c *chainClientV2) FetchChainPerpetualMarketInfo(ctx context.Context, marketId string) (*exchangev2types.QueryPerpetualMarketInfoResponse, error) {
	req := &exchangev2types.QueryPerpetualMarketInfoRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.PerpetualMarketInfo, req)

	return res, err
}

func (c *chainClientV2) FetchChainExpiryFuturesMarketInfo(ctx context.Context, marketId string) (*exchangev2types.QueryExpiryFuturesMarketInfoResponse, error) {
	req := &exchangev2types.QueryExpiryFuturesMarketInfoRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.ExpiryFuturesMarketInfo, req)

	return res, err
}

func (c *chainClientV2) FetchChainPerpetualMarketFunding(ctx context.Context, marketId string) (*exchangev2types.QueryPerpetualMarketFundingResponse, error) {
	req := &exchangev2types.QueryPerpetualMarketFundingRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.PerpetualMarketFunding, req)

	return res, err
}

func (c *chainClientV2) FetchSubaccountOrderMetadata(ctx context.Context, subaccountId string) (*exchangev2types.QuerySubaccountOrderMetadataResponse, error) {
	req := &exchangev2types.QuerySubaccountOrderMetadataRequest{
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.SubaccountOrderMetadata, req)

	return res, err
}

func (c *chainClientV2) FetchTradeRewardPoints(ctx context.Context, accounts []string) (*exchangev2types.QueryTradeRewardPointsResponse, error) {
	req := &exchangev2types.QueryTradeRewardPointsRequest{
		Accounts: accounts,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.TradeRewardPoints, req)

	return res, err
}

func (c *chainClientV2) FetchPendingTradeRewardPoints(ctx context.Context, accounts []string) (*exchangev2types.QueryTradeRewardPointsResponse, error) {
	req := &exchangev2types.QueryTradeRewardPointsRequest{
		Accounts: accounts,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.PendingTradeRewardPoints, req)

	return res, err
}

func (c *chainClientV2) FetchTradeRewardCampaign(ctx context.Context) (*exchangev2types.QueryTradeRewardCampaignResponse, error) {
	req := &exchangev2types.QueryTradeRewardCampaignRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.TradeRewardCampaign, req)

	return res, err
}

func (c *chainClientV2) FetchFeeDiscountAccountInfo(ctx context.Context, account string) (*exchangev2types.QueryFeeDiscountAccountInfoResponse, error) {
	req := &exchangev2types.QueryFeeDiscountAccountInfoRequest{
		Account: account,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.FeeDiscountAccountInfo, req)

	return res, err
}

func (c *chainClientV2) FetchFeeDiscountSchedule(ctx context.Context) (*exchangev2types.QueryFeeDiscountScheduleResponse, error) {
	req := &exchangev2types.QueryFeeDiscountScheduleRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.FeeDiscountSchedule, req)

	return res, err
}

func (c *chainClientV2) FetchBalanceMismatches(ctx context.Context, dustFactor int64) (*exchangev2types.QueryBalanceMismatchesResponse, error) {
	req := &exchangev2types.QueryBalanceMismatchesRequest{
		DustFactor: dustFactor,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.BalanceMismatches, req)

	return res, err
}

func (c *chainClientV2) FetchBalanceWithBalanceHolds(ctx context.Context) (*exchangev2types.QueryBalanceWithBalanceHoldsResponse, error) {
	req := &exchangev2types.QueryBalanceWithBalanceHoldsRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.BalanceWithBalanceHolds, req)

	return res, err
}

func (c *chainClientV2) FetchFeeDiscountTierStatistics(ctx context.Context) (*exchangev2types.QueryFeeDiscountTierStatisticsResponse, error) {
	req := &exchangev2types.QueryFeeDiscountTierStatisticsRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.FeeDiscountTierStatistics, req)

	return res, err
}

func (c *chainClientV2) FetchMitoVaultInfos(ctx context.Context) (*exchangev2types.MitoVaultInfosResponse, error) {
	req := &exchangev2types.MitoVaultInfosRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.MitoVaultInfos, req)

	return res, err
}

func (c *chainClientV2) FetchMarketIDFromVault(ctx context.Context, vaultAddress string) (*exchangev2types.QueryMarketIDFromVaultResponse, error) {
	req := &exchangev2types.QueryMarketIDFromVaultRequest{
		VaultAddress: vaultAddress,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.QueryMarketIDFromVault, req)

	return res, err
}

func (c *chainClientV2) FetchHistoricalTradeRecords(ctx context.Context, marketId string) (*exchangev2types.QueryHistoricalTradeRecordsResponse, error) {
	req := &exchangev2types.QueryHistoricalTradeRecordsRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.HistoricalTradeRecords, req)

	return res, err
}

func (c *chainClientV2) FetchIsOptedOutOfRewards(ctx context.Context, account string) (*exchangev2types.QueryIsOptedOutOfRewardsResponse, error) {
	req := &exchangev2types.QueryIsOptedOutOfRewardsRequest{
		Account: account,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.IsOptedOutOfRewards, req)

	return res, err
}

func (c *chainClientV2) FetchOptedOutOfRewardsAccounts(ctx context.Context) (*exchangev2types.QueryOptedOutOfRewardsAccountsResponse, error) {
	req := &exchangev2types.QueryOptedOutOfRewardsAccountsRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.OptedOutOfRewardsAccounts, req)

	return res, err
}

func (c *chainClientV2) FetchMarketVolatility(ctx context.Context, marketId string, tradeHistoryOptions *exchangev2types.TradeHistoryOptions) (*exchangev2types.QueryMarketVolatilityResponse, error) {
	req := &exchangev2types.QueryMarketVolatilityRequest{
		MarketId:            marketId,
		TradeHistoryOptions: tradeHistoryOptions,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.MarketVolatility, req)

	return res, err
}

func (c *chainClientV2) FetchChainBinaryOptionsMarkets(ctx context.Context, status string) (*exchangev2types.QueryBinaryMarketsResponse, error) {
	req := &exchangev2types.QueryBinaryMarketsRequest{}
	if status != "" {
		req.Status = status
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.BinaryOptionsMarkets, req)

	return res, err
}

func (c *chainClientV2) FetchTraderDerivativeConditionalOrders(ctx context.Context, subaccountId, marketId string) (*exchangev2types.QueryTraderDerivativeConditionalOrdersResponse, error) {
	req := &exchangev2types.QueryTraderDerivativeConditionalOrdersRequest{
		SubaccountId: subaccountId,
		MarketId:     marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.TraderDerivativeConditionalOrders, req)

	return res, err
}

func (c *chainClientV2) FetchMarketAtomicExecutionFeeMultiplier(ctx context.Context, marketId string) (*exchangev2types.QueryMarketAtomicExecutionFeeMultiplierResponse, error) {
	req := &exchangev2types.QueryMarketAtomicExecutionFeeMultiplierRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.MarketAtomicExecutionFeeMultiplier, req)

	return res, err
}

func (c *chainClientV2) FetchActiveStakeGrant(ctx context.Context, grantee string) (*exchangev2types.QueryActiveStakeGrantResponse, error) {
	req := &exchangev2types.QueryActiveStakeGrantRequest{
		Grantee: grantee,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.ActiveStakeGrant, req)

	return res, err
}

func (c *chainClientV2) FetchGrantAuthorization(ctx context.Context, granter, grantee string) (*exchangev2types.QueryGrantAuthorizationResponse, error) {
	req := &exchangev2types.QueryGrantAuthorizationRequest{
		Granter: granter,
		Grantee: grantee,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.GrantAuthorization, req)

	return res, err
}

func (c *chainClientV2) FetchGrantAuthorizations(ctx context.Context, granter string) (*exchangev2types.QueryGrantAuthorizationsResponse, error) {
	req := &exchangev2types.QueryGrantAuthorizationsRequest{
		Granter: granter,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.GrantAuthorizations, req)

	return res, err
}

func (c *chainClientV2) FetchL3DerivativeOrderbook(ctx context.Context, marketId string) (*exchangev2types.QueryFullDerivativeOrderbookResponse, error) {
	req := &exchangev2types.QueryFullDerivativeOrderbookRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.L3DerivativeOrderBook, req)

	return res, err
}

func (c *chainClientV2) FetchL3SpotOrderbook(ctx context.Context, marketId string) (*exchangev2types.QueryFullSpotOrderbookResponse, error) {
	req := &exchangev2types.QueryFullSpotOrderbookRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.L3SpotOrderBook, req)

	return res, err
}

func (c *chainClientV2) FetchMarketBalance(ctx context.Context, marketId string) (*exchangev2types.QueryMarketBalanceResponse, error) {
	req := &exchangev2types.QueryMarketBalanceRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.MarketBalance, req)

	return res, err
}

func (c *chainClientV2) FetchMarketBalances(ctx context.Context) (*exchangev2types.QueryMarketBalancesResponse, error) {
	req := &exchangev2types.QueryMarketBalancesRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.MarketBalances, req)

	return res, err
}

func (c *chainClientV2) FetchDenomMinNotional(ctx context.Context, denom string) (*exchangev2types.QueryDenomMinNotionalResponse, error) {
	req := &exchangev2types.QueryDenomMinNotionalRequest{
		Denom: denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.DenomMinNotional, req)

	return res, err
}

func (c *chainClientV2) FetchDenomMinNotionals(ctx context.Context) (*exchangev2types.QueryDenomMinNotionalsResponse, error) {
	req := &exchangev2types.QueryDenomMinNotionalsRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeV2QueryClient.DenomMinNotionals, req)

	return res, err
}

// Tendermint module

func (c *chainClientV2) FetchNodeInfo(ctx context.Context) (*cmtservice.GetNodeInfoResponse, error) {
	req := &cmtservice.GetNodeInfoRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tendermintQueryClient.GetNodeInfo, req)

	return res, err
}

func (c *chainClientV2) FetchSyncing(ctx context.Context) (*cmtservice.GetSyncingResponse, error) {
	req := &cmtservice.GetSyncingRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tendermintQueryClient.GetSyncing, req)

	return res, err
}

func (c *chainClientV2) FetchLatestBlock(ctx context.Context) (*cmtservice.GetLatestBlockResponse, error) {
	req := &cmtservice.GetLatestBlockRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tendermintQueryClient.GetLatestBlock, req)

	return res, err
}

func (c *chainClientV2) FetchBlockByHeight(ctx context.Context, height int64) (*cmtservice.GetBlockByHeightResponse, error) {
	req := &cmtservice.GetBlockByHeightRequest{
		Height: height,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tendermintQueryClient.GetBlockByHeight, req)

	return res, err
}

func (c *chainClientV2) FetchLatestValidatorSet(ctx context.Context) (*cmtservice.GetLatestValidatorSetResponse, error) {
	req := &cmtservice.GetLatestValidatorSetRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tendermintQueryClient.GetLatestValidatorSet, req)

	return res, err
}

func (c *chainClientV2) FetchValidatorSetByHeight(ctx context.Context, height int64, pagination *query.PageRequest) (*cmtservice.GetValidatorSetByHeightResponse, error) {
	req := &cmtservice.GetValidatorSetByHeightRequest{
		Height:     height,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tendermintQueryClient.GetValidatorSetByHeight, req)

	return res, err
}

func (c *chainClientV2) ABCIQuery(ctx context.Context, path string, data []byte, height int64, prove bool) (*cmtservice.ABCIQueryResponse, error) {
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
func (c *chainClientV2) FetchDenomTrace(ctx context.Context, hash string) (*ibctransfertypes.QueryDenomTraceResponse, error) {
	req := &ibctransfertypes.QueryDenomTraceRequest{
		Hash: hash,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcTransferQueryClient.DenomTrace, req)

	return res, err
}

func (c *chainClientV2) FetchDenomTraces(ctx context.Context, pagination *query.PageRequest) (*ibctransfertypes.QueryDenomTracesResponse, error) {
	req := &ibctransfertypes.QueryDenomTracesRequest{
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcTransferQueryClient.DenomTraces, req)

	return res, err
}

func (c *chainClientV2) FetchDenomHash(ctx context.Context, trace string) (*ibctransfertypes.QueryDenomHashResponse, error) {
	req := &ibctransfertypes.QueryDenomHashRequest{
		Trace: trace,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcTransferQueryClient.DenomHash, req)

	return res, err
}

func (c *chainClientV2) FetchEscrowAddress(ctx context.Context, portId, channelId string) (*ibctransfertypes.QueryEscrowAddressResponse, error) {
	req := &ibctransfertypes.QueryEscrowAddressRequest{
		PortId:    portId,
		ChannelId: channelId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcTransferQueryClient.EscrowAddress, req)

	return res, err
}

func (c *chainClientV2) FetchTotalEscrowForDenom(ctx context.Context, denom string) (*ibctransfertypes.QueryTotalEscrowForDenomResponse, error) {
	req := &ibctransfertypes.QueryTotalEscrowForDenomRequest{
		Denom: denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcTransferQueryClient.TotalEscrowForDenom, req)

	return res, err
}

// IBC Core Channel module
func (c *chainClientV2) FetchIBCChannel(ctx context.Context, portId, channelId string) (*ibcchanneltypes.QueryChannelResponse, error) {
	req := &ibcchanneltypes.QueryChannelRequest{
		PortId:    portId,
		ChannelId: channelId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.Channel, req)

	return res, err
}

func (c *chainClientV2) FetchIBCChannels(ctx context.Context, pagination *query.PageRequest) (*ibcchanneltypes.QueryChannelsResponse, error) {
	req := &ibcchanneltypes.QueryChannelsRequest{
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.Channels, req)

	return res, err
}

func (c *chainClientV2) FetchIBCConnectionChannels(ctx context.Context, connection string, pagination *query.PageRequest) (*ibcchanneltypes.QueryConnectionChannelsResponse, error) {
	req := &ibcchanneltypes.QueryConnectionChannelsRequest{
		Connection: connection,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.ConnectionChannels, req)

	return res, err
}

func (c *chainClientV2) FetchIBCChannelClientState(ctx context.Context, portId, channelId string) (*ibcchanneltypes.QueryChannelClientStateResponse, error) {
	req := &ibcchanneltypes.QueryChannelClientStateRequest{
		PortId:    portId,
		ChannelId: channelId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.ChannelClientState, req)

	return res, err
}

func (c *chainClientV2) FetchIBCChannelConsensusState(ctx context.Context, portId, channelId string, revisionNumber, revisionHeight uint64) (*ibcchanneltypes.QueryChannelConsensusStateResponse, error) {
	req := &ibcchanneltypes.QueryChannelConsensusStateRequest{
		PortId:         portId,
		ChannelId:      channelId,
		RevisionNumber: revisionNumber,
		RevisionHeight: revisionHeight,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.ChannelConsensusState, req)

	return res, err
}

func (c *chainClientV2) FetchIBCPacketCommitment(ctx context.Context, portId, channelId string, sequence uint64) (*ibcchanneltypes.QueryPacketCommitmentResponse, error) {
	req := &ibcchanneltypes.QueryPacketCommitmentRequest{
		PortId:    portId,
		ChannelId: channelId,
		Sequence:  sequence,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.PacketCommitment, req)

	return res, err
}

func (c *chainClientV2) FetchIBCPacketCommitments(ctx context.Context, portId, channelId string, pagination *query.PageRequest) (*ibcchanneltypes.QueryPacketCommitmentsResponse, error) {
	req := &ibcchanneltypes.QueryPacketCommitmentsRequest{
		PortId:     portId,
		ChannelId:  channelId,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.PacketCommitments, req)

	return res, err
}

func (c *chainClientV2) FetchIBCPacketReceipt(ctx context.Context, portId, channelId string, sequence uint64) (*ibcchanneltypes.QueryPacketReceiptResponse, error) {
	req := &ibcchanneltypes.QueryPacketReceiptRequest{
		PortId:    portId,
		ChannelId: channelId,
		Sequence:  sequence,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.PacketReceipt, req)

	return res, err
}

func (c *chainClientV2) FetchIBCPacketAcknowledgement(ctx context.Context, portId, channelId string, sequence uint64) (*ibcchanneltypes.QueryPacketAcknowledgementResponse, error) {
	req := &ibcchanneltypes.QueryPacketAcknowledgementRequest{
		PortId:    portId,
		ChannelId: channelId,
		Sequence:  sequence,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.PacketAcknowledgement, req)

	return res, err
}

func (c *chainClientV2) FetchIBCPacketAcknowledgements(ctx context.Context, portId, channelId string, packetCommitmentSequences []uint64, pagination *query.PageRequest) (*ibcchanneltypes.QueryPacketAcknowledgementsResponse, error) {
	req := &ibcchanneltypes.QueryPacketAcknowledgementsRequest{
		PortId:                    portId,
		ChannelId:                 channelId,
		Pagination:                pagination,
		PacketCommitmentSequences: packetCommitmentSequences,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.PacketAcknowledgements, req)

	return res, err
}

func (c *chainClientV2) FetchIBCUnreceivedPackets(ctx context.Context, portId, channelId string, packetCommitmentSequences []uint64) (*ibcchanneltypes.QueryUnreceivedPacketsResponse, error) {
	req := &ibcchanneltypes.QueryUnreceivedPacketsRequest{
		PortId:                    portId,
		ChannelId:                 channelId,
		PacketCommitmentSequences: packetCommitmentSequences,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.UnreceivedPackets, req)

	return res, err
}

func (c *chainClientV2) FetchIBCUnreceivedAcks(ctx context.Context, portId, channelId string, packetAckSequences []uint64) (*ibcchanneltypes.QueryUnreceivedAcksResponse, error) {
	req := &ibcchanneltypes.QueryUnreceivedAcksRequest{
		PortId:             portId,
		ChannelId:          channelId,
		PacketAckSequences: packetAckSequences,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.UnreceivedAcks, req)

	return res, err
}

func (c *chainClientV2) FetchIBCNextSequenceReceive(ctx context.Context, portId, channelId string) (*ibcchanneltypes.QueryNextSequenceReceiveResponse, error) {
	req := &ibcchanneltypes.QueryNextSequenceReceiveRequest{
		PortId:    portId,
		ChannelId: channelId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.NextSequenceReceive, req)

	return res, err
}

// IBC Core Chain module
func (c *chainClientV2) FetchIBCClientState(ctx context.Context, clientId string) (*ibcclienttypes.QueryClientStateResponse, error) {
	req := &ibcclienttypes.QueryClientStateRequest{
		ClientId: clientId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcClientQueryClient.ClientState, req)

	return res, err
}

func (c *chainClientV2) FetchIBCClientStates(ctx context.Context, pagination *query.PageRequest) (*ibcclienttypes.QueryClientStatesResponse, error) {
	req := &ibcclienttypes.QueryClientStatesRequest{
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcClientQueryClient.ClientStates, req)

	return res, err
}

func (c *chainClientV2) FetchIBCConsensusState(ctx context.Context, clientId string, revisionNumber, revisionHeight uint64, latestHeight bool) (*ibcclienttypes.QueryConsensusStateResponse, error) {
	req := &ibcclienttypes.QueryConsensusStateRequest{
		ClientId:       clientId,
		RevisionNumber: revisionNumber,
		RevisionHeight: revisionHeight,
		LatestHeight:   latestHeight,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcClientQueryClient.ConsensusState, req)

	return res, err
}

func (c *chainClientV2) FetchIBCConsensusStates(ctx context.Context, clientId string, pagination *query.PageRequest) (*ibcclienttypes.QueryConsensusStatesResponse, error) {
	req := &ibcclienttypes.QueryConsensusStatesRequest{
		ClientId:   clientId,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcClientQueryClient.ConsensusStates, req)

	return res, err
}

func (c *chainClientV2) FetchIBCConsensusStateHeights(ctx context.Context, clientId string, pagination *query.PageRequest) (*ibcclienttypes.QueryConsensusStateHeightsResponse, error) {
	req := &ibcclienttypes.QueryConsensusStateHeightsRequest{
		ClientId:   clientId,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcClientQueryClient.ConsensusStateHeights, req)

	return res, err
}

func (c *chainClientV2) FetchIBCClientStatus(ctx context.Context, clientId string) (*ibcclienttypes.QueryClientStatusResponse, error) {
	req := &ibcclienttypes.QueryClientStatusRequest{
		ClientId: clientId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcClientQueryClient.ClientStatus, req)

	return res, err
}

func (c *chainClientV2) FetchIBCClientParams(ctx context.Context) (*ibcclienttypes.QueryClientParamsResponse, error) {
	req := &ibcclienttypes.QueryClientParamsRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcClientQueryClient.ClientParams, req)

	return res, err
}

func (c *chainClientV2) FetchIBCUpgradedClientState(ctx context.Context) (*ibcclienttypes.QueryUpgradedClientStateResponse, error) {
	req := &ibcclienttypes.QueryUpgradedClientStateRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcClientQueryClient.UpgradedClientState, req)

	return res, err
}

func (c *chainClientV2) FetchIBCUpgradedConsensusState(ctx context.Context) (*ibcclienttypes.QueryUpgradedConsensusStateResponse, error) {
	req := &ibcclienttypes.QueryUpgradedConsensusStateRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcClientQueryClient.UpgradedConsensusState, req)

	return res, err
}

// IBC Core Connection module
func (c *chainClientV2) FetchIBCConnection(ctx context.Context, connectionId string) (*ibcconnectiontypes.QueryConnectionResponse, error) {
	req := &ibcconnectiontypes.QueryConnectionRequest{
		ConnectionId: connectionId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcConnectionQueryClient.Connection, req)

	return res, err
}

func (c *chainClientV2) FetchIBCConnections(ctx context.Context, pagination *query.PageRequest) (*ibcconnectiontypes.QueryConnectionsResponse, error) {
	req := &ibcconnectiontypes.QueryConnectionsRequest{
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcConnectionQueryClient.Connections, req)

	return res, err
}

func (c *chainClientV2) FetchIBCClientConnections(ctx context.Context, clientId string) (*ibcconnectiontypes.QueryClientConnectionsResponse, error) {
	req := &ibcconnectiontypes.QueryClientConnectionsRequest{
		ClientId: clientId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcConnectionQueryClient.ClientConnections, req)

	return res, err
}

func (c *chainClientV2) FetchIBCConnectionClientState(ctx context.Context, connectionId string) (*ibcconnectiontypes.QueryConnectionClientStateResponse, error) {
	req := &ibcconnectiontypes.QueryConnectionClientStateRequest{
		ConnectionId: connectionId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcConnectionQueryClient.ConnectionClientState, req)

	return res, err
}

func (c *chainClientV2) FetchIBCConnectionConsensusState(ctx context.Context, connectionId string, revisionNumber, revisionHeight uint64) (*ibcconnectiontypes.QueryConnectionConsensusStateResponse, error) {
	req := &ibcconnectiontypes.QueryConnectionConsensusStateRequest{
		ConnectionId:   connectionId,
		RevisionNumber: revisionNumber,
		RevisionHeight: revisionHeight,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcConnectionQueryClient.ConnectionConsensusState, req)

	return res, err
}

func (c *chainClientV2) FetchIBCConnectionParams(ctx context.Context) (*ibcconnectiontypes.QueryConnectionParamsResponse, error) {
	req := &ibcconnectiontypes.QueryConnectionParamsRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcConnectionQueryClient.ConnectionParams, req)

	return res, err
}

// Permissions module

func (c *chainClientV2) FetchPermissionsNamespaceDenoms(ctx context.Context) (*permissionstypes.QueryNamespaceDenomsResponse, error) {
	req := &permissionstypes.QueryNamespaceDenomsRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.permissionsQueryClient.NamespaceDenoms, req)

	return res, err
}

func (c *chainClientV2) FetchPermissionsNamespaces(ctx context.Context) (*permissionstypes.QueryNamespacesResponse, error) {
	req := &permissionstypes.QueryNamespacesRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.permissionsQueryClient.Namespaces, req)

	return res, err
}

func (c *chainClientV2) FetchPermissionsNamespace(ctx context.Context, denom string) (*permissionstypes.QueryNamespaceResponse, error) {
	req := &permissionstypes.QueryNamespaceRequest{
		Denom: denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.permissionsQueryClient.Namespace, req)

	return res, err
}

func (c *chainClientV2) FetchPermissionsRolesByActor(ctx context.Context, denom, actor string) (*permissionstypes.QueryRolesByActorResponse, error) {
	req := &permissionstypes.QueryRolesByActorRequest{
		Denom: denom,
		Actor: actor,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.permissionsQueryClient.RolesByActor, req)

	return res, err
}

func (c *chainClientV2) FetchPermissionsActorsByRole(ctx context.Context, denom, role string) (*permissionstypes.QueryActorsByRoleResponse, error) {
	req := &permissionstypes.QueryActorsByRoleRequest{
		Denom: denom,
		Role:  role,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.permissionsQueryClient.ActorsByRole, req)

	return res, err
}

func (c *chainClientV2) FetchPermissionsRoleManagers(ctx context.Context, denom string) (*permissionstypes.QueryRoleManagersResponse, error) {
	req := &permissionstypes.QueryRoleManagersRequest{
		Denom: denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.permissionsQueryClient.RoleManagers, req)

	return res, err
}

func (c *chainClientV2) FetchPermissionsRoleManager(ctx context.Context, denom, manager string) (*permissionstypes.QueryRoleManagerResponse, error) {
	req := &permissionstypes.QueryRoleManagerRequest{
		Denom:   denom,
		Manager: manager,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.permissionsQueryClient.RoleManager, req)

	return res, err
}

func (c *chainClientV2) FetchPermissionsPolicyStatuses(ctx context.Context, denom string) (*permissionstypes.QueryPolicyStatusesResponse, error) {
	req := &permissionstypes.QueryPolicyStatusesRequest{
		Denom: denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.permissionsQueryClient.PolicyStatuses, req)

	return res, err
}

func (c *chainClientV2) FetchPermissionsPolicyManagerCapabilities(ctx context.Context, denom string) (*permissionstypes.QueryPolicyManagerCapabilitiesResponse, error) {
	req := &permissionstypes.QueryPolicyManagerCapabilitiesRequest{
		Denom: denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.permissionsQueryClient.PolicyManagerCapabilities, req)

	return res, err
}

func (c *chainClientV2) FetchPermissionsVouchers(ctx context.Context, denom string) (*permissionstypes.QueryVouchersResponse, error) {
	req := &permissionstypes.QueryVouchersRequest{
		Denom: denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.permissionsQueryClient.Vouchers, req)

	return res, err
}

func (c *chainClientV2) FetchPermissionsVoucher(ctx context.Context, denom, address string) (*permissionstypes.QueryVoucherResponse, error) {
	req := &permissionstypes.QueryVoucherRequest{
		Denom:   denom,
		Address: address,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.permissionsQueryClient.Voucher, req)

	return res, err
}

func (c *chainClientV2) FetchPermissionsModuleState(ctx context.Context) (*permissionstypes.QueryModuleStateResponse, error) {
	req := &permissionstypes.QueryModuleStateRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.permissionsQueryClient.PermissionsModuleState, req)

	return res, err
}

// TxFees module
func (c *chainClientV2) FetchTxFeesParams(ctx context.Context) (*txfeestypes.QueryParamsResponse, error) {
	req := &txfeestypes.QueryParamsRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.txfeesQueryClient.Params, req)

	return res, err
}

func (c *chainClientV2) FetchEipBaseFee(ctx context.Context) (*txfeestypes.QueryEipBaseFeeResponse, error) {
	req := &txfeestypes.QueryEipBaseFeeRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.txfeesQueryClient.GetEipBaseFee, req)

	return res, err
}

func (c *chainClientV2) GetNetwork() common.Network {
	return c.network
}

// ERC20 module

func (c *chainClientV2) FetchAllTokenPairs(ctx context.Context, pagination *query.PageRequest) (*erc20types.QueryAllTokenPairsResponse, error) {
	req := &erc20types.QueryAllTokenPairsRequest{
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.erc20QueryClient.AllTokenPairs, req)

	return res, err
}

func (c *chainClientV2) FetchTokenPairByDenom(ctx context.Context, bankDenom string) (*erc20types.QueryTokenPairByDenomResponse, error) {
	req := &erc20types.QueryTokenPairByDenomRequest{
		BankDenom: bankDenom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.erc20QueryClient.TokenPairByDenom, req)

	return res, err
}

func (c *chainClientV2) FetchTokenPairByERC20Address(ctx context.Context, erc20Address string) (*erc20types.QueryTokenPairByERC20AddressResponse, error) {
	req := &erc20types.QueryTokenPairByERC20AddressRequest{
		Erc20Address: erc20Address,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.erc20QueryClient.TokenPairByERC20Address, req)

	return res, err
}

// EVM module

func (c *chainClientV2) FetchEVMAccount(ctx context.Context, address string) (*evmtypes.QueryAccountResponse, error) {
	req := &evmtypes.QueryAccountRequest{
		Address: address,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.evmQueryClient.Account, req)
	return res, err
}

func (c *chainClientV2) FetchEVMCosmosAccount(ctx context.Context, address string) (*evmtypes.QueryCosmosAccountResponse, error) {
	req := &evmtypes.QueryCosmosAccountRequest{
		Address: address,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.evmQueryClient.CosmosAccount, req)
	return res, err
}

func (c *chainClientV2) FetchEVMValidatorAccount(ctx context.Context, consAddress string) (*evmtypes.QueryValidatorAccountResponse, error) {
	req := &evmtypes.QueryValidatorAccountRequest{
		ConsAddress: consAddress,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.evmQueryClient.ValidatorAccount, req)
	return res, err
}

func (c *chainClientV2) FetchEVMBalance(ctx context.Context, address string) (*evmtypes.QueryBalanceResponse, error) {
	req := &evmtypes.QueryBalanceRequest{
		Address: address,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.evmQueryClient.Balance, req)
	return res, err
}

func (c *chainClientV2) FetchEVMStorage(ctx context.Context, address string, key *string) (*evmtypes.QueryStorageResponse, error) {
	req := &evmtypes.QueryStorageRequest{
		Address: address,
	}

	if key != nil {
		req.Key = *key
	}

	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.evmQueryClient.Storage, req)
	return res, err
}

func (c *chainClientV2) FetchEVMCode(ctx context.Context, address string) (*evmtypes.QueryCodeResponse, error) {
	req := &evmtypes.QueryCodeRequest{
		Address: address,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.evmQueryClient.Code, req)
	return res, err
}

func (c *chainClientV2) FetchEVMBaseFee(ctx context.Context) (*evmtypes.QueryBaseFeeResponse, error) {
	req := &evmtypes.QueryBaseFeeRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.evmQueryClient.BaseFee, req)
	return res, err
}

// SyncBroadcastMsg sends Tx to chain and waits until Tx is included in block.
func (c *chainClientV2) SyncBroadcastMsg(ctx context.Context, pollInterval *time.Duration, maxRetries int, msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error) {
	req, res, err := c.BroadcastMsg(ctx, txtypes.BroadcastMode_BROADCAST_MODE_SYNC, msgs...)

	if err != nil || res.TxResponse.Code != 0 {
		return res, err
	}

	txHash, _ := hex.DecodeString(res.TxResponse.TxHash)

	statusPollInterval := defaultBroadcastStatusPoll
	if pollInterval != nil {
		statusPollInterval = *pollInterval
	}

	totalAttempts := 0
	t := time.NewTimer(statusPollInterval)

	for {
		select {
		case <-ctx.Done():
			err := errors.Wrapf(ErrTimedOut, "%s", res.TxResponse.TxHash)
			t.Stop()
			return nil, err
		case <-t.C:
			totalAttempts++
			resultTx, txErr := c.ctx.Client.Tx(ctx, txHash, false)

			if txErr != nil {
				// Check if this is a fatal error that shouldn't be retried
				if errRes := sdkclient.CheckCometError(txErr, req.TxBytes); errRes != nil {
					return &txtypes.BroadcastTxResponse{TxResponse: errRes}, txErr
				}

				// If we've reached max retries, return error
				if totalAttempts >= maxRetries {
					t.Stop()
					return nil, errors.Wrapf(txErr, "failed to get transaction after %d retries: %s", maxRetries, res.TxResponse.TxHash)
				}

				// Continue retrying with same interval
				t.Reset(statusPollInterval)
				continue
			} else if resultTx.Height > 0 {
				resResultTx := sdk.NewResponseResultTx(resultTx, res.TxResponse.Tx, res.TxResponse.Timestamp)
				res = &txtypes.BroadcastTxResponse{TxResponse: resResultTx}
				t.Stop()
				return res, err
			}

			// Transaction not yet in block, continue polling
			t.Reset(statusPollInterval)
		}
	}
}

// AsyncBroadcastMsg sends Tx to chain and doesn't wait until Tx is included in block. This method
// cannot be used for rapid Tx sending, it is expected that you wait for transaction status with
// external tools. If you want sdk to wait for it, use SyncBroadcastMsg.
func (c *chainClientV2) AsyncBroadcastMsg(ctx context.Context, msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error) {
	_, res, err := c.BroadcastMsg(ctx, txtypes.BroadcastMode_BROADCAST_MODE_ASYNC, msgs...)
	return res, err
}

// BroadcastMsg submits a group of messages in one transaction to the chain
// The function uses the broadcast mode specified with the broadcastMode parameter
func (c *chainClientV2) BroadcastMsg(ctx context.Context, broadcastMode txtypes.BroadcastMode, msgs ...sdk.Msg) (*txtypes.BroadcastTxRequest, *txtypes.BroadcastTxResponse, error) {
	c.syncMux.Lock()
	defer c.syncMux.Unlock()

	sequence := c.getAccSeq()
	c.txFactory = c.txFactory.WithSequence(sequence)
	c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
	req, res, err := c.broadcastTx(ctx, c.txFactory, broadcastMode, msgs...)
	if err != nil {
		if c.opts.ShouldFixSequenceMismatch && strings.Contains(err.Error(), "account sequence mismatch") {
			c.syncNonce()
			sequence := c.getAccSeq()
			c.txFactory = c.txFactory.WithSequence(sequence)
			c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
			c.logger.Debugln("retrying broadcastTx with nonce", sequence)
			req, res, err = c.broadcastTx(ctx, c.txFactory, broadcastMode, msgs...)
		}
		if err != nil {
			resJSON, _ := json.MarshalIndent(res, "", "\t")
			c.logger.WithField("size", len(msgs)).WithError(err).Errorln("failed to asynchronously broadcast messagess:", string(resJSON))
			return nil, nil, err
		}
	}

	return req, res, nil
}

func (c *chainClientV2) UpdateSubaccountNonceFromChain() error {
	for subaccountId := range c.subaccountToNonce {
		err := c.SynchronizeSubaccountNonce(subaccountId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *chainClientV2) SynchronizeSubaccountNonce(subaccountId ethcommon.Hash) error {
	res, err := c.FetchSubaccountTradeNonce(context.Background(), subaccountId.Hex())
	if err != nil {
		return err
	}
	c.subaccountToNonce[subaccountId] = res.Nonce
	return nil
}

func (c *chainClientV2) ComputeOrderHashes(spotOrders []exchangev2types.SpotOrder, derivativeOrders []exchangev2types.DerivativeOrder, subaccountId ethcommon.Hash) (OrderHashes, error) {
	if len(spotOrders)+len(derivativeOrders) == 0 {
		return OrderHashes{}, nil
	}

	orderHashes := OrderHashes{}
	// get nonce
	if _, exist := c.subaccountToNonce[subaccountId]; !exist {
		if err := c.SynchronizeSubaccountNonce(subaccountId); err != nil {
			return OrderHashes{}, err
		}
	}

	nonce := c.subaccountToNonce[subaccountId]
	for _, o := range spotOrders {
		nonce += 1
		triggerPrice := ""
		if o.TriggerPrice != nil {
			triggerPrice = o.TriggerPrice.String()
		}
		message := map[string]interface{}{
			"MarketId": o.MarketId,
			"OrderInfo": map[string]interface{}{
				"SubaccountId": o.OrderInfo.SubaccountId,
				"FeeRecipient": o.OrderInfo.FeeRecipient,
				"Price":        o.OrderInfo.Price.String(),
				"Quantity":     o.OrderInfo.Quantity.String(),
			},
			"Salt":         strconv.Itoa(int(nonce)),
			"OrderType":    string(o.OrderType),
			"TriggerPrice": triggerPrice,
		}
		typedData := gethsigner.TypedData{
			Types:       eip712OrderTypes,
			PrimaryType: "SpotOrder",
			Domain:      domain,
			Message:     message,
		}
		domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
		if err != nil {
			return OrderHashes{}, err
		}
		typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
		if err != nil {
			return OrderHashes{}, err
		}

		w := sha3.NewLegacyKeccak256()
		w.Write([]byte("\x19\x01"))
		w.Write([]byte(domainSeparator))
		w.Write([]byte(typedDataHash))

		hash := ethcommon.BytesToHash(w.Sum(nil))
		orderHashes.Spot = append(orderHashes.Spot, hash)
	}

	for _, o := range derivativeOrders {
		nonce += 1
		triggerPrice := ""
		if o.TriggerPrice != nil {
			triggerPrice = o.TriggerPrice.String()
		}
		message := map[string]interface{}{
			"MarketId": o.MarketId,
			"OrderInfo": map[string]interface{}{
				"SubaccountId": o.OrderInfo.SubaccountId,
				"FeeRecipient": o.OrderInfo.FeeRecipient,
				"Price":        o.OrderInfo.Price.String(),
				"Quantity":     o.OrderInfo.Quantity.String(),
			},
			"Margin":       o.Margin.String(),
			"OrderType":    string(o.OrderType),
			"TriggerPrice": triggerPrice,
			"Salt":         strconv.Itoa(int(nonce)),
		}
		typedData := gethsigner.TypedData{
			Types:       eip712OrderTypes,
			PrimaryType: "DerivativeOrder",
			Domain:      domain,
			Message:     message,
		}
		domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
		if err != nil {
			return OrderHashes{}, err
		}
		typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
		if err != nil {
			return OrderHashes{}, err
		}

		w := sha3.NewLegacyKeccak256()
		w.Write([]byte("\x19\x01"))
		w.Write([]byte(domainSeparator))
		w.Write([]byte(typedDataHash))

		hash := ethcommon.BytesToHash(w.Sum(nil))
		orderHashes.Derivative = append(orderHashes.Derivative, hash)
	}

	c.subaccountToNonce[subaccountId] = nonce

	return orderHashes, nil
}

func (c *chainClientV2) CurrentChainGasPrice(ctx context.Context) int64 {
	gasPrice := int64(client.DefaultGasPrice)
	eipBaseFee, err := c.FetchEipBaseFee(ctx)

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

func (c *chainClientV2) SetGasPrice(gasPrice int64) {
	gasPrices := fmt.Sprintf("%v%s", gasPrice, client.InjDenom)

	c.txFactory = c.txFactory.WithGasPrices(gasPrices)
}
