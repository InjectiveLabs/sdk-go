package chain

import (
	"context"
	"errors"
	"time"

	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"

	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"

	tokenfactorytypes "github.com/InjectiveLabs/sdk-go/chain/tokenfactory/types"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"

	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	chainstreamtypes "github.com/InjectiveLabs/sdk-go/chain/stream/types"
	"github.com/InjectiveLabs/sdk-go/client/common"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	eth "github.com/ethereum/go-ethereum/common"
	"google.golang.org/grpc"
)

type MockChainClient struct {
	DenomsMetadataResponses []*banktypes.QueryDenomsMetadataResponse
}

func (c *MockChainClient) CanSignTransactions() bool {
	return true
}

func (c *MockChainClient) FromAddress() sdk.AccAddress {
	return sdk.AccAddress{}
}

func (c *MockChainClient) QueryClient() *grpc.ClientConn {
	return &grpc.ClientConn{}
}

func (c *MockChainClient) ClientContext() client.Context {
	return client.Context{}
}

func (c *MockChainClient) GetAccNonce() (accNum uint64, accSeq uint64) {
	return 1, 2
}

func (c *MockChainClient) SimulateMsg(clientCtx client.Context, msgs ...sdk.Msg) (*txtypes.SimulateResponse, error) {
	return &txtypes.SimulateResponse{}, nil
}

func (c *MockChainClient) AsyncBroadcastMsg(msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error) {
	return &txtypes.BroadcastTxResponse{}, nil
}

func (c *MockChainClient) SyncBroadcastMsg(msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error) {
	return &txtypes.BroadcastTxResponse{}, nil
}

func (c *MockChainClient) BuildSignedTx(clientCtx client.Context, accNum, accSeq, initialGas uint64, msg ...sdk.Msg) ([]byte, error) {
	return *new([]byte), nil
}

func (c *MockChainClient) SyncBroadcastSignedTx(tyBytes []byte) (*txtypes.BroadcastTxResponse, error) {
	return &txtypes.BroadcastTxResponse{}, nil
}

func (c *MockChainClient) AsyncBroadcastSignedTx(txBytes []byte) (*txtypes.BroadcastTxResponse, error) {
	return &txtypes.BroadcastTxResponse{}, nil
}

func (c *MockChainClient) QueueBroadcastMsg(msgs ...sdk.Msg) error {
	return nil
}

func (c *MockChainClient) GetBankBalances(ctx context.Context, address string) (*banktypes.QueryAllBalancesResponse, error) {
	return &banktypes.QueryAllBalancesResponse{}, nil
}

func (c *MockChainClient) GetBankBalance(ctx context.Context, address string, denom string) (*banktypes.QueryBalanceResponse, error) {
	return &banktypes.QueryBalanceResponse{}, nil
}

func (c *MockChainClient) GetBankSpendableBalances(ctx context.Context, address string, pagination *query.PageRequest) (*banktypes.QuerySpendableBalancesResponse, error) {
	return &banktypes.QuerySpendableBalancesResponse{}, nil
}

func (c *MockChainClient) GetBankSpendableBalancesByDenom(ctx context.Context, address string, denom string) (*banktypes.QuerySpendableBalanceByDenomResponse, error) {
	return &banktypes.QuerySpendableBalanceByDenomResponse{}, nil
}

func (c *MockChainClient) GetBankTotalSupply(ctx context.Context, pagination *query.PageRequest) (*banktypes.QueryTotalSupplyResponse, error) {
	return &banktypes.QueryTotalSupplyResponse{}, nil
}

func (c *MockChainClient) GetBankSupplyOf(ctx context.Context, denom string) (*banktypes.QuerySupplyOfResponse, error) {
	return &banktypes.QuerySupplyOfResponse{}, nil
}

func (c *MockChainClient) GetDenomMetadata(ctx context.Context, denom string) (*banktypes.QueryDenomMetadataResponse, error) {
	return &banktypes.QueryDenomMetadataResponse{}, nil
}

func (c *MockChainClient) GetDenomsMetadata(ctx context.Context, pagination *query.PageRequest) (*banktypes.QueryDenomsMetadataResponse, error) {
	var response *banktypes.QueryDenomsMetadataResponse
	var localError error
	if len(c.DenomsMetadataResponses) > 0 {
		response = c.DenomsMetadataResponses[0]
		c.DenomsMetadataResponses = c.DenomsMetadataResponses[1:]
		localError = nil
	} else {
		response = &banktypes.QueryDenomsMetadataResponse{}
		localError = errors.New("there are no responses configured")
	}

	return response, localError
}

func (c *MockChainClient) GetDenomOwners(ctx context.Context, denom string, pagination *query.PageRequest) (*banktypes.QueryDenomOwnersResponse, error) {
	return &banktypes.QueryDenomOwnersResponse{}, nil
}

func (c *MockChainClient) GetBankSendEnabled(ctx context.Context, denoms []string, pagination *query.PageRequest) (*banktypes.QuerySendEnabledResponse, error) {
	return &banktypes.QuerySendEnabledResponse{}, nil
}

func (c *MockChainClient) GetAuthzGrants(ctx context.Context, req authztypes.QueryGrantsRequest) (*authztypes.QueryGrantsResponse, error) {
	return &authztypes.QueryGrantsResponse{}, nil
}

func (c *MockChainClient) GetAccount(ctx context.Context, address string) (*authtypes.QueryAccountResponse, error) {
	return &authtypes.QueryAccountResponse{}, nil
}

func (c *MockChainClient) BuildGenericAuthz(granter string, grantee string, msgtype string, expireIn time.Time) *authztypes.MsgGrant {
	return &authztypes.MsgGrant{}
}

func (c *MockChainClient) BuildExchangeAuthz(granter string, grantee string, authzType ExchangeAuthz, subaccountId string, markets []string, expireIn time.Time) *authztypes.MsgGrant {
	return &authztypes.MsgGrant{}
}

func (c *MockChainClient) BuildExchangeBatchUpdateOrdersAuthz(
	granter string,
	grantee string,
	subaccountId string,
	spotMarkets []string,
	derivativeMarkets []string,
	expireIn time.Time,
) *authztypes.MsgGrant {
	return &authztypes.MsgGrant{}
}

func (c *MockChainClient) DefaultSubaccount(acc sdk.AccAddress) eth.Hash {
	return eth.HexToHash("")
}

func (c *MockChainClient) Subaccount(account sdk.AccAddress, index int) eth.Hash {
	return eth.HexToHash("")
}

func (c *MockChainClient) GetSubAccountNonce(ctx context.Context, subaccountId eth.Hash) (*exchangetypes.QuerySubaccountTradeNonceResponse, error) {
	return &exchangetypes.QuerySubaccountTradeNonceResponse{}, nil
}

func (c *MockChainClient) GetFeeDiscountInfo(ctx context.Context, account string) (*exchangetypes.QueryFeeDiscountAccountInfoResponse, error) {
	return &exchangetypes.QueryFeeDiscountAccountInfoResponse{}, nil
}

func (c *MockChainClient) UpdateSubaccountNonceFromChain() error {
	return nil
}

func (c *MockChainClient) SynchronizeSubaccountNonce(subaccountId eth.Hash) error {
	return nil
}

func (c *MockChainClient) ComputeOrderHashes(spotOrders []exchangetypes.SpotOrder, derivativeOrders []exchangetypes.DerivativeOrder, subaccountId eth.Hash) (OrderHashes, error) {
	return OrderHashes{}, nil
}

func (c *MockChainClient) SpotOrder(defaultSubaccountID eth.Hash, network common.Network, d *SpotOrderData) *exchangetypes.SpotOrder {
	return c.CreateSpotOrder(defaultSubaccountID, d, MarketsAssistant{})
}

func (c *MockChainClient) CreateSpotOrder(defaultSubaccountID eth.Hash, d *SpotOrderData, marketsAssistant MarketsAssistant) *exchangetypes.SpotOrder {
	return &exchangetypes.SpotOrder{}
}

func (c *MockChainClient) DerivativeOrder(defaultSubaccountID eth.Hash, network common.Network, d *DerivativeOrderData) *exchangetypes.DerivativeOrder {
	return c.CreateDerivativeOrder(defaultSubaccountID, d, MarketsAssistant{})
}

func (c *MockChainClient) CreateDerivativeOrder(defaultSubaccountID eth.Hash, d *DerivativeOrderData, marketAssistant MarketsAssistant) *exchangetypes.DerivativeOrder {
	return &exchangetypes.DerivativeOrder{}
}

func (c *MockChainClient) OrderCancel(defaultSubaccountID eth.Hash, d *OrderCancelData) *exchangetypes.OrderData {
	return &exchangetypes.OrderData{}
}

func (c *MockChainClient) StreamEventOrderFail(sender string, failEventCh chan map[string]uint) {}

func (c *MockChainClient) StreamEventOrderFailWithWebsocket(sender string, websocket *rpchttp.HTTP, failEventCh chan map[string]uint) {
}

func (c *MockChainClient) StreamOrderbookUpdateEvents(orderbookType OrderbookType, marketIds []string, orderbookCh chan exchangetypes.Orderbook) {
}

func (c *MockChainClient) StreamOrderbookUpdateEventsWithWebsocket(orderbookType OrderbookType, marketIds []string, websocket *rpchttp.HTTP, orderbookCh chan exchangetypes.Orderbook) {
}

func (c *MockChainClient) ChainStream(ctx context.Context, req chainstreamtypes.StreamRequest) (chainstreamtypes.Stream_StreamClient, error) {
	return nil, nil
}

func (c *MockChainClient) GetTx(ctx context.Context, txHash string) (*txtypes.GetTxResponse, error) {
	return &txtypes.GetTxResponse{}, nil
}

func (c *MockChainClient) Close() {}

func (c *MockChainClient) GetGasFee() (string, error) {
	return "", nil
}

func (c *MockChainClient) FetchContractInfo(ctx context.Context, address string) (*wasmtypes.QueryContractInfoResponse, error) {
	return &wasmtypes.QueryContractInfoResponse{}, nil
}

func (c *MockChainClient) FetchContractHistory(ctx context.Context, address string, pagination *query.PageRequest) (*wasmtypes.QueryContractHistoryResponse, error) {
	return &wasmtypes.QueryContractHistoryResponse{}, nil
}

func (c *MockChainClient) FetchContractsByCode(ctx context.Context, codeId uint64, pagination *query.PageRequest) (*wasmtypes.QueryContractsByCodeResponse, error) {
	return &wasmtypes.QueryContractsByCodeResponse{}, nil
}

func (c *MockChainClient) FetchAllContractsState(ctx context.Context, address string, pagination *query.PageRequest) (*wasmtypes.QueryAllContractStateResponse, error) {
	return &wasmtypes.QueryAllContractStateResponse{}, nil
}

func (c *MockChainClient) SmartContractState(
	ctx context.Context,
	contractAddress string,
	queryData []byte,
) (*wasmtypes.QuerySmartContractStateResponse, error) {
	return &wasmtypes.QuerySmartContractStateResponse{}, nil
}

func (c *MockChainClient) RawContractState(
	ctx context.Context,
	contractAddress string,
	queryData []byte,
) (*wasmtypes.QueryRawContractStateResponse, error) {
	return &wasmtypes.QueryRawContractStateResponse{}, nil
}

func (c *MockChainClient) FetchCode(ctx context.Context, codeId uint64) (*wasmtypes.QueryCodeResponse, error) {
	return &wasmtypes.QueryCodeResponse{}, nil
}

func (c *MockChainClient) FetchCodes(ctx context.Context, pagination *query.PageRequest) (*wasmtypes.QueryCodesResponse, error) {
	return &wasmtypes.QueryCodesResponse{}, nil
}

func (c *MockChainClient) FetchPinnedCodes(ctx context.Context, pagination *query.PageRequest) (*wasmtypes.QueryPinnedCodesResponse, error) {
	return &wasmtypes.QueryPinnedCodesResponse{}, nil
}

func (c *MockChainClient) FetchContractsByCreator(ctx context.Context, creator string, pagination *query.PageRequest) (*wasmtypes.QueryContractsByCreatorResponse, error) {
	return &wasmtypes.QueryContractsByCreatorResponse{}, nil
}

func (c *MockChainClient) FetchDenomAuthorityMetadata(ctx context.Context, creator string, subDenom string) (*tokenfactorytypes.QueryDenomAuthorityMetadataResponse, error) {
	return &tokenfactorytypes.QueryDenomAuthorityMetadataResponse{}, nil
}

func (c *MockChainClient) FetchDenomsFromCreator(ctx context.Context, creator string) (*tokenfactorytypes.QueryDenomsFromCreatorResponse, error) {
	return &tokenfactorytypes.QueryDenomsFromCreatorResponse{}, nil
}

func (c *MockChainClient) FetchTokenfactoryModuleState(ctx context.Context) (*tokenfactorytypes.QueryModuleStateResponse, error) {
	return &tokenfactorytypes.QueryModuleStateResponse{}, nil
}

// Distribution module
func (c *MockChainClient) FetchValidatorDistributionInfo(ctx context.Context, validatorAddress string) (*distributiontypes.QueryValidatorDistributionInfoResponse, error) {
	return &distributiontypes.QueryValidatorDistributionInfoResponse{}, nil
}

func (c *MockChainClient) FetchValidatorOutstandingRewards(ctx context.Context, validatorAddress string) (*distributiontypes.QueryValidatorOutstandingRewardsResponse, error) {
	return &distributiontypes.QueryValidatorOutstandingRewardsResponse{}, nil
}

func (c *MockChainClient) FetchValidatorCommission(ctx context.Context, validatorAddress string) (*distributiontypes.QueryValidatorCommissionResponse, error) {
	return &distributiontypes.QueryValidatorCommissionResponse{}, nil
}

func (c *MockChainClient) FetchValidatorSlashes(ctx context.Context, validatorAddress string, startingHeight uint64, endingHeight uint64, pagination *query.PageRequest) (*distributiontypes.QueryValidatorSlashesResponse, error) {
	return &distributiontypes.QueryValidatorSlashesResponse{}, nil
}

func (c *MockChainClient) FetchDelegationRewards(ctx context.Context, delegatorAddress string, validatorAddress string) (*distributiontypes.QueryDelegationRewardsResponse, error) {
	return &distributiontypes.QueryDelegationRewardsResponse{}, nil
}

func (c *MockChainClient) FetchDelegationTotalRewards(ctx context.Context, delegatorAddress string) (*distributiontypes.QueryDelegationTotalRewardsResponse, error) {
	return &distributiontypes.QueryDelegationTotalRewardsResponse{}, nil
}

func (c *MockChainClient) FetchDelegatorValidators(ctx context.Context, delegatorAddress string) (*distributiontypes.QueryDelegatorValidatorsResponse, error) {
	return &distributiontypes.QueryDelegatorValidatorsResponse{}, nil
}

func (c *MockChainClient) FetchDelegatorWithdrawAddress(ctx context.Context, delegatorAddress string) (*distributiontypes.QueryDelegatorWithdrawAddressResponse, error) {
	return &distributiontypes.QueryDelegatorWithdrawAddressResponse{}, nil
}

func (c *MockChainClient) FetchCommunityPool(ctx context.Context) (*distributiontypes.QueryCommunityPoolResponse, error) {
	return &distributiontypes.QueryCommunityPoolResponse{}, nil
}

// Chain exchange module
func (c *MockChainClient) FetchSubaccountDeposits(ctx context.Context, subaccountId string) (*exchangetypes.QuerySubaccountDepositsResponse, error) {
	return &exchangetypes.QuerySubaccountDepositsResponse{}, nil
}

func (c *MockChainClient) FetchSubaccountDeposit(ctx context.Context, subaccountId string, denom string) (*exchangetypes.QuerySubaccountDepositResponse, error) {
	return &exchangetypes.QuerySubaccountDepositResponse{}, nil
}

func (c *MockChainClient) FetchExchangeBalances(ctx context.Context) (*exchangetypes.QueryExchangeBalancesResponse, error) {
	return &exchangetypes.QueryExchangeBalancesResponse{}, nil
}

func (c *MockChainClient) FetchAggregateVolume(ctx context.Context, account string) (*exchangetypes.QueryAggregateVolumeResponse, error) {
	return &exchangetypes.QueryAggregateVolumeResponse{}, nil
}

func (c *MockChainClient) FetchAggregateVolumes(ctx context.Context, accounts []string, marketIds []string) (*exchangetypes.QueryAggregateVolumesResponse, error) {
	return &exchangetypes.QueryAggregateVolumesResponse{}, nil
}

func (c *MockChainClient) FetchAggregateMarketVolume(ctx context.Context, marketId string) (*exchangetypes.QueryAggregateMarketVolumeResponse, error) {
	return &exchangetypes.QueryAggregateMarketVolumeResponse{}, nil
}

func (c *MockChainClient) FetchAggregateMarketVolumes(ctx context.Context, marketIds []string) (*exchangetypes.QueryAggregateMarketVolumesResponse, error) {
	return &exchangetypes.QueryAggregateMarketVolumesResponse{}, nil
}

func (c *MockChainClient) FetchDenomDecimal(ctx context.Context, denom string) (*exchangetypes.QueryDenomDecimalResponse, error) {
	return &exchangetypes.QueryDenomDecimalResponse{}, nil
}

func (c *MockChainClient) FetchDenomDecimals(ctx context.Context, denoms []string) (*exchangetypes.QueryDenomDecimalsResponse, error) {
	return &exchangetypes.QueryDenomDecimalsResponse{}, nil
}

func (c *MockChainClient) FetchChainSpotMarkets(ctx context.Context, status string, marketIds []string) (*exchangetypes.QuerySpotMarketsResponse, error) {
	return &exchangetypes.QuerySpotMarketsResponse{}, nil
}

func (c *MockChainClient) FetchChainSpotMarket(ctx context.Context, marketId string) (*exchangetypes.QuerySpotMarketResponse, error) {
	return &exchangetypes.QuerySpotMarketResponse{}, nil
}

func (c *MockChainClient) FetchChainFullSpotMarkets(ctx context.Context, status string, marketIds []string, withMidPriceAndTob bool) (*exchangetypes.QueryFullSpotMarketsResponse, error) {
	return &exchangetypes.QueryFullSpotMarketsResponse{}, nil
}

func (c *MockChainClient) FetchChainFullSpotMarket(ctx context.Context, marketId string, withMidPriceAndTob bool) (*exchangetypes.QueryFullSpotMarketResponse, error) {
	return &exchangetypes.QueryFullSpotMarketResponse{}, nil
}

func (c *MockChainClient) FetchChainSpotOrderbook(ctx context.Context, marketId string, limit uint64, orderSide exchangetypes.OrderSide, limitCumulativeNotional sdk.Dec, limitCumulativeQuantity sdk.Dec) (*exchangetypes.QuerySpotOrderbookResponse, error) {
	return &exchangetypes.QuerySpotOrderbookResponse{}, nil
}

func (c *MockChainClient) FetchChainTraderSpotOrders(ctx context.Context, marketId string, subaccountId string) (*exchangetypes.QueryTraderSpotOrdersResponse, error) {
	return &exchangetypes.QueryTraderSpotOrdersResponse{}, nil
}

func (c *MockChainClient) FetchChainAccountAddressSpotOrders(ctx context.Context, marketId string, address string) (*exchangetypes.QueryAccountAddressSpotOrdersResponse, error) {
	return &exchangetypes.QueryAccountAddressSpotOrdersResponse{}, nil
}

func (c *MockChainClient) FetchChainSpotOrdersByHashes(ctx context.Context, marketId string, subaccountId string, orderHashes []string) (*exchangetypes.QuerySpotOrdersByHashesResponse, error) {
	return &exchangetypes.QuerySpotOrdersByHashesResponse{}, nil
}

func (c *MockChainClient) FetchChainSubaccountOrders(ctx context.Context, subaccountId string, marketId string) (*exchangetypes.QuerySubaccountOrdersResponse, error) {
	return &exchangetypes.QuerySubaccountOrdersResponse{}, nil
}

func (c *MockChainClient) FetchChainTraderSpotTransientOrders(ctx context.Context, marketId string, subaccountId string) (*exchangetypes.QueryTraderSpotOrdersResponse, error) {
	return &exchangetypes.QueryTraderSpotOrdersResponse{}, nil
}

func (c *MockChainClient) FetchSpotMidPriceAndTOB(ctx context.Context, marketId string) (*exchangetypes.QuerySpotMidPriceAndTOBResponse, error) {
	return &exchangetypes.QuerySpotMidPriceAndTOBResponse{}, nil
}

func (c *MockChainClient) FetchDerivativeMidPriceAndTOB(ctx context.Context, marketId string) (*exchangetypes.QueryDerivativeMidPriceAndTOBResponse, error) {
	return &exchangetypes.QueryDerivativeMidPriceAndTOBResponse{}, nil
}

func (c *MockChainClient) FetchChainDerivativeOrderbook(ctx context.Context, marketId string, limit uint64, limitCumulativeNotional sdk.Dec) (*exchangetypes.QueryDerivativeOrderbookResponse, error) {
	return &exchangetypes.QueryDerivativeOrderbookResponse{}, nil
}

func (c *MockChainClient) FetchChainTraderDerivativeOrders(ctx context.Context, marketId string, subaccountId string) (*exchangetypes.QueryTraderDerivativeOrdersResponse, error) {
	return &exchangetypes.QueryTraderDerivativeOrdersResponse{}, nil
}

func (c *MockChainClient) FetchChainAccountAddressDerivativeOrders(ctx context.Context, marketId string, address string) (*exchangetypes.QueryAccountAddressDerivativeOrdersResponse, error) {
	return &exchangetypes.QueryAccountAddressDerivativeOrdersResponse{}, nil
}

func (c *MockChainClient) FetchChainDerivativeOrdersByHashes(ctx context.Context, marketId string, subaccountId string, orderHashes []string) (*exchangetypes.QueryDerivativeOrdersByHashesResponse, error) {
	return &exchangetypes.QueryDerivativeOrdersByHashesResponse{}, nil
}

func (c *MockChainClient) FetchChainTraderDerivativeTransientOrders(ctx context.Context, marketId string, subaccountId string) (*exchangetypes.QueryTraderDerivativeOrdersResponse, error) {
	return &exchangetypes.QueryTraderDerivativeOrdersResponse{}, nil
}

func (c *MockChainClient) FetchChainDerivativeMarkets(ctx context.Context, status string, marketIds []string, withMidPriceAndTob bool) (*exchangetypes.QueryDerivativeMarketsResponse, error) {
	return &exchangetypes.QueryDerivativeMarketsResponse{}, nil
}

func (c *MockChainClient) FetchChainDerivativeMarket(ctx context.Context, marketId string) (*exchangetypes.QueryDerivativeMarketResponse, error) {
	return &exchangetypes.QueryDerivativeMarketResponse{}, nil
}

func (c *MockChainClient) FetchDerivativeMarketAddress(ctx context.Context, marketId string) (*exchangetypes.QueryDerivativeMarketAddressResponse, error) {
	return &exchangetypes.QueryDerivativeMarketAddressResponse{}, nil
}

func (c *MockChainClient) FetchSubaccountTradeNonce(ctx context.Context, subaccountId string) (*exchangetypes.QuerySubaccountTradeNonceResponse, error) {
	return &exchangetypes.QuerySubaccountTradeNonceResponse{}, nil
}

func (c *MockChainClient) FetchChainPositions(ctx context.Context) (*exchangetypes.QueryPositionsResponse, error) {
	return &exchangetypes.QueryPositionsResponse{}, nil
}

func (c *MockChainClient) FetchChainSubaccountPositions(ctx context.Context, subaccountId string) (*exchangetypes.QuerySubaccountPositionsResponse, error) {
	return &exchangetypes.QuerySubaccountPositionsResponse{}, nil
}

func (c *MockChainClient) FetchChainSubaccountPositionInMarket(ctx context.Context, subaccountId string, marketId string) (*exchangetypes.QuerySubaccountPositionInMarketResponse, error) {
	return &exchangetypes.QuerySubaccountPositionInMarketResponse{}, nil
}

func (c *MockChainClient) FetchChainSubaccountEffectivePositionInMarket(ctx context.Context, subaccountId string, marketId string) (*exchangetypes.QuerySubaccountEffectivePositionInMarketResponse, error) {
	return &exchangetypes.QuerySubaccountEffectivePositionInMarketResponse{}, nil
}

func (c *MockChainClient) FetchChainPerpetualMarketInfo(ctx context.Context, marketId string) (*exchangetypes.QueryPerpetualMarketInfoResponse, error) {
	return &exchangetypes.QueryPerpetualMarketInfoResponse{}, nil
}

func (c *MockChainClient) FetchChainExpiryFuturesMarketInfo(ctx context.Context, marketId string) (*exchangetypes.QueryExpiryFuturesMarketInfoResponse, error) {
	return &exchangetypes.QueryExpiryFuturesMarketInfoResponse{}, nil
}

func (c *MockChainClient) FetchChainPerpetualMarketFunding(ctx context.Context, marketId string) (*exchangetypes.QueryPerpetualMarketFundingResponse, error) {
	return &exchangetypes.QueryPerpetualMarketFundingResponse{}, nil
}

func (c *MockChainClient) FetchSubaccountOrderMetadata(ctx context.Context, subaccountId string) (*exchangetypes.QuerySubaccountOrderMetadataResponse, error) {
	return &exchangetypes.QuerySubaccountOrderMetadataResponse{}, nil
}

func (c *MockChainClient) FetchTradeRewardPoints(ctx context.Context, accounts []string) (*exchangetypes.QueryTradeRewardPointsResponse, error) {
	return &exchangetypes.QueryTradeRewardPointsResponse{}, nil
}

func (c *MockChainClient) FetchPendingTradeRewardPoints(ctx context.Context, accounts []string) (*exchangetypes.QueryTradeRewardPointsResponse, error) {
	return &exchangetypes.QueryTradeRewardPointsResponse{}, nil
}

func (c *MockChainClient) FetchTradeRewardCampaign(ctx context.Context) (*exchangetypes.QueryTradeRewardCampaignResponse, error) {
	return &exchangetypes.QueryTradeRewardCampaignResponse{}, nil
}

func (c *MockChainClient) FetchFeeDiscountAccountInfo(ctx context.Context, account string) (*exchangetypes.QueryFeeDiscountAccountInfoResponse, error) {
	return &exchangetypes.QueryFeeDiscountAccountInfoResponse{}, nil
}

func (c *MockChainClient) FetchFeeDiscountSchedule(ctx context.Context) (*exchangetypes.QueryFeeDiscountScheduleResponse, error) {
	return &exchangetypes.QueryFeeDiscountScheduleResponse{}, nil
}

func (c *MockChainClient) FetchBalanceMismatches(ctx context.Context, dustFactor int64) (*exchangetypes.QueryBalanceMismatchesResponse, error) {
	return &exchangetypes.QueryBalanceMismatchesResponse{}, nil
}

func (c *MockChainClient) FetchBalanceWithBalanceHolds(ctx context.Context) (*exchangetypes.QueryBalanceWithBalanceHoldsResponse, error) {
	return &exchangetypes.QueryBalanceWithBalanceHoldsResponse{}, nil
}

func (c *MockChainClient) FetchFeeDiscountTierStatistics(ctx context.Context) (*exchangetypes.QueryFeeDiscountTierStatisticsResponse, error) {
	return &exchangetypes.QueryFeeDiscountTierStatisticsResponse{}, nil
}

func (c *MockChainClient) FetchMitoVaultInfos(ctx context.Context) (*exchangetypes.MitoVaultInfosResponse, error) {
	return &exchangetypes.MitoVaultInfosResponse{}, nil
}

func (c *MockChainClient) FetchMarketIDFromVault(ctx context.Context, vaultAddress string) (*exchangetypes.QueryMarketIDFromVaultResponse, error) {
	return &exchangetypes.QueryMarketIDFromVaultResponse{}, nil
}

func (c *MockChainClient) FetchHistoricalTradeRecords(ctx context.Context, marketId string) (*exchangetypes.QueryHistoricalTradeRecordsResponse, error) {
	return &exchangetypes.QueryHistoricalTradeRecordsResponse{}, nil
}

func (c *MockChainClient) FetchIsOptedOutOfRewards(ctx context.Context, account string) (*exchangetypes.QueryIsOptedOutOfRewardsResponse, error) {
	return &exchangetypes.QueryIsOptedOutOfRewardsResponse{}, nil
}

func (c *MockChainClient) FetchOptedOutOfRewardsAccounts(ctx context.Context) (*exchangetypes.QueryOptedOutOfRewardsAccountsResponse, error) {
	return &exchangetypes.QueryOptedOutOfRewardsAccountsResponse{}, nil
}

func (c *MockChainClient) FetchMarketVolatility(ctx context.Context, marketId string, tradeHistoryOptions *exchangetypes.TradeHistoryOptions) (*exchangetypes.QueryMarketVolatilityResponse, error) {
	return &exchangetypes.QueryMarketVolatilityResponse{}, nil
}

func (c *MockChainClient) FetchChainBinaryOptionsMarkets(ctx context.Context, status string) (*exchangetypes.QueryBinaryMarketsResponse, error) {
	return &exchangetypes.QueryBinaryMarketsResponse{}, nil
}

func (c *MockChainClient) FetchTraderDerivativeConditionalOrders(ctx context.Context, subaccountId string, marketId string) (*exchangetypes.QueryTraderDerivativeConditionalOrdersResponse, error) {
	return &exchangetypes.QueryTraderDerivativeConditionalOrdersResponse{}, nil
}

func (c *MockChainClient) FetchMarketAtomicExecutionFeeMultiplier(ctx context.Context, marketId string) (*exchangetypes.QueryMarketAtomicExecutionFeeMultiplierResponse, error) {
	return &exchangetypes.QueryMarketAtomicExecutionFeeMultiplierResponse{}, nil
}

// Tendermint module

func (c *MockChainClient) FetchNodeInfo(ctx context.Context) (*tmservice.GetNodeInfoResponse, error) {
	return &tmservice.GetNodeInfoResponse{}, nil
}

func (c *MockChainClient) FetchSyncing(ctx context.Context) (*tmservice.GetSyncingResponse, error) {
	return &tmservice.GetSyncingResponse{}, nil
}

func (c *MockChainClient) FetchLatestBlock(ctx context.Context) (*tmservice.GetLatestBlockResponse, error) {
	return &tmservice.GetLatestBlockResponse{}, nil
}

func (c *MockChainClient) FetchBlockByHeight(ctx context.Context, height int64) (*tmservice.GetBlockByHeightResponse, error) {
	return &tmservice.GetBlockByHeightResponse{}, nil
}

func (c *MockChainClient) FetchLatestValidatorSet(ctx context.Context) (*tmservice.GetLatestValidatorSetResponse, error) {
	return &tmservice.GetLatestValidatorSetResponse{}, nil
}

func (c *MockChainClient) FetchValidatorSetByHeight(ctx context.Context, height int64, pagination *query.PageRequest) (*tmservice.GetValidatorSetByHeightResponse, error) {
	return &tmservice.GetValidatorSetByHeightResponse{}, nil
}

func (c *MockChainClient) ABCIQuery(ctx context.Context, path string, data []byte, height int64, prove bool) (*tmservice.ABCIQueryResponse, error) {
	return &tmservice.ABCIQueryResponse{}, nil
}
