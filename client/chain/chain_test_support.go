package chain

import (
	"context"
	"errors"
	"time"

	sdkmath "cosmossdk.io/math"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	ibcconnectiontypes "github.com/cosmos/ibc-go/v8/modules/core/03-connection/types"
	ibcchanneltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"google.golang.org/grpc"

	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	exchangev2types "github.com/InjectiveLabs/sdk-go/chain/exchange/types/v2"
	permissionstypes "github.com/InjectiveLabs/sdk-go/chain/permissions/types"
	chainstreamtypes "github.com/InjectiveLabs/sdk-go/chain/stream/types"
	chainstreamv2types "github.com/InjectiveLabs/sdk-go/chain/stream/types/v2"
	tokenfactorytypes "github.com/InjectiveLabs/sdk-go/chain/tokenfactory/types"
	"github.com/InjectiveLabs/sdk-go/client/common"
)

var _ ChainClient = &MockChainClient{}

type MockChainClient struct {
	Network                           common.Network
	DenomsMetadataResponses           []*banktypes.QueryDenomsMetadataResponse
	QuerySpotMarketsResponses         []*exchangetypes.QuerySpotMarketsResponse
	QueryDerivativeMarketsResponses   []*exchangetypes.QueryDerivativeMarketsResponse
	QueryBinaryMarketsResponses       []*exchangetypes.QueryBinaryMarketsResponse
	QuerySpotMarketsV2Responses       []*exchangev2types.QuerySpotMarketsResponse
	QueryDerivativeMarketsV2Responses []*exchangev2types.QueryDerivativeMarketsResponse
	QueryBinaryMarketsV2Responses     []*exchangev2types.QueryBinaryMarketsResponse
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

func (c *MockChainClient) GetAccNonce() (accNum, accSeq uint64) {
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
	return []byte(nil), nil
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

func (c *MockChainClient) GetBankBalance(ctx context.Context, address, denom string) (*banktypes.QueryBalanceResponse, error) {
	return &banktypes.QueryBalanceResponse{}, nil
}

func (c *MockChainClient) GetBankSpendableBalances(ctx context.Context, address string, pagination *query.PageRequest) (*banktypes.QuerySpendableBalancesResponse, error) {
	return &banktypes.QuerySpendableBalancesResponse{}, nil
}

func (c *MockChainClient) GetBankSpendableBalancesByDenom(ctx context.Context, address, denom string) (*banktypes.QuerySpendableBalanceByDenomResponse, error) {
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

func (c *MockChainClient) BuildGenericAuthz(granter, grantee, msgtype string, expireIn time.Time) *authztypes.MsgGrant {
	return &authztypes.MsgGrant{}
}

func (c *MockChainClient) BuildExchangeAuthz(granter, grantee string, authzType ExchangeAuthz, subaccountId string, markets []string, expireIn time.Time) *authztypes.MsgGrant {
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

func (c *MockChainClient) DefaultSubaccount(acc sdk.AccAddress) ethcommon.Hash {
	return ethcommon.HexToHash("")
}

func (c *MockChainClient) Subaccount(account sdk.AccAddress, index int) ethcommon.Hash {
	return ethcommon.HexToHash("")
}

func (c *MockChainClient) GetSubAccountNonce(ctx context.Context, subaccountId ethcommon.Hash) (*exchangetypes.QuerySubaccountTradeNonceResponse, error) {
	return &exchangetypes.QuerySubaccountTradeNonceResponse{}, nil
}

func (c *MockChainClient) GetFeeDiscountInfo(ctx context.Context, account string) (*exchangetypes.QueryFeeDiscountAccountInfoResponse, error) {
	return &exchangetypes.QueryFeeDiscountAccountInfoResponse{}, nil
}

func (c *MockChainClient) GetSubAccountNonceV2(ctx context.Context, subaccountId ethcommon.Hash) (*exchangev2types.QuerySubaccountTradeNonceResponse, error) {
	return &exchangev2types.QuerySubaccountTradeNonceResponse{}, nil
}

func (c *MockChainClient) GetFeeDiscountInfoV2(ctx context.Context, account string) (*exchangev2types.QueryFeeDiscountAccountInfoResponse, error) {
	return &exchangev2types.QueryFeeDiscountAccountInfoResponse{}, nil
}

func (c *MockChainClient) UpdateSubaccountNonceFromChain() error {
	return nil
}

func (c *MockChainClient) SynchronizeSubaccountNonce(subaccountId ethcommon.Hash) error {
	return nil
}

func (c *MockChainClient) ComputeOrderHashes(spotOrders []exchangev2types.SpotOrder, derivativeOrders []exchangev2types.DerivativeOrder, subaccountId ethcommon.Hash) (OrderHashes, error) {
	return OrderHashes{}, nil
}

func (c *MockChainClient) CreateSpotOrder(defaultSubaccountID ethcommon.Hash, d *SpotOrderData, marketAssistant MarketsAssistant) *exchangetypes.SpotOrder {
	return &exchangetypes.SpotOrder{}
}

func (c *MockChainClient) CreateDerivativeOrder(defaultSubaccountID ethcommon.Hash, d *DerivativeOrderData, marketAssistant MarketsAssistant) *exchangetypes.DerivativeOrder {
	return &exchangetypes.DerivativeOrder{}
}

func (c *MockChainClient) OrderCancel(defaultSubaccountID ethcommon.Hash, d *OrderCancelData) *exchangetypes.OrderData {
	return &exchangetypes.OrderData{}
}

func (c *MockChainClient) CreateSpotOrderV2(defaultSubaccountID ethcommon.Hash, d *SpotOrderData) *exchangev2types.SpotOrder {
	return &exchangev2types.SpotOrder{}
}

func (c *MockChainClient) CreateDerivativeOrderV2(defaultSubaccountID ethcommon.Hash, d *DerivativeOrderData) *exchangev2types.DerivativeOrder {
	return &exchangev2types.DerivativeOrder{}
}

func (c *MockChainClient) OrderCancelV2(defaultSubaccountID ethcommon.Hash, d *OrderCancelData) *exchangev2types.OrderData {
	return &exchangev2types.OrderData{}
}

func (c *MockChainClient) StreamEventOrderFail(sender string, failEventCh chan map[string]uint) {}

func (c *MockChainClient) StreamEventOrderFailWithWebsocket(sender string, websocket *rpchttp.HTTP, failEventCh chan map[string]uint) {
}

func (c *MockChainClient) StreamOrderbookUpdateEvents(orderbookType OrderbookType, marketIDs []string, orderbookCh chan exchangetypes.Orderbook) {
}

func (c *MockChainClient) StreamOrderbookUpdateEventsWithWebsocket(orderbookType OrderbookType, marketIDs []string, websocket *rpchttp.HTTP, orderbookCh chan exchangetypes.Orderbook) {
}

func (c *MockChainClient) ChainStream(ctx context.Context, req chainstreamtypes.StreamRequest) (chainstreamtypes.Stream_StreamClient, error) {
	return nil, nil
}

func (c *MockChainClient) ChainStreamV2(ctx context.Context, req chainstreamv2types.StreamRequest) (chainstreamv2types.Stream_StreamV2Client, error) {
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

func (c *MockChainClient) FetchDenomAuthorityMetadata(ctx context.Context, creator, subDenom string) (*tokenfactorytypes.QueryDenomAuthorityMetadataResponse, error) {
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

func (c *MockChainClient) FetchValidatorSlashes(ctx context.Context, validatorAddress string, startingHeight, endingHeight uint64, pagination *query.PageRequest) (*distributiontypes.QueryValidatorSlashesResponse, error) {
	return &distributiontypes.QueryValidatorSlashesResponse{}, nil
}

func (c *MockChainClient) FetchDelegationRewards(ctx context.Context, delegatorAddress, validatorAddress string) (*distributiontypes.QueryDelegationRewardsResponse, error) {
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

func (c *MockChainClient) FetchSubaccountDeposit(ctx context.Context, subaccountId, denom string) (*exchangetypes.QuerySubaccountDepositResponse, error) {
	return &exchangetypes.QuerySubaccountDepositResponse{}, nil
}

func (c *MockChainClient) FetchExchangeBalances(ctx context.Context) (*exchangetypes.QueryExchangeBalancesResponse, error) {
	return &exchangetypes.QueryExchangeBalancesResponse{}, nil
}

func (c *MockChainClient) FetchAggregateVolume(ctx context.Context, account string) (*exchangetypes.QueryAggregateVolumeResponse, error) {
	return &exchangetypes.QueryAggregateVolumeResponse{}, nil
}

func (c *MockChainClient) FetchAggregateVolumes(ctx context.Context, accounts, marketIDs []string) (*exchangetypes.QueryAggregateVolumesResponse, error) {
	return &exchangetypes.QueryAggregateVolumesResponse{}, nil
}

func (c *MockChainClient) FetchAggregateMarketVolume(ctx context.Context, marketId string) (*exchangetypes.QueryAggregateMarketVolumeResponse, error) {
	return &exchangetypes.QueryAggregateMarketVolumeResponse{}, nil
}

func (c *MockChainClient) FetchAggregateMarketVolumes(ctx context.Context, marketIDs []string) (*exchangetypes.QueryAggregateMarketVolumesResponse, error) {
	return &exchangetypes.QueryAggregateMarketVolumesResponse{}, nil
}

func (c *MockChainClient) FetchDenomDecimal(ctx context.Context, denom string) (*exchangetypes.QueryDenomDecimalResponse, error) {
	return &exchangetypes.QueryDenomDecimalResponse{}, nil
}

func (c *MockChainClient) FetchDenomDecimals(ctx context.Context, denoms []string) (*exchangetypes.QueryDenomDecimalsResponse, error) {
	return &exchangetypes.QueryDenomDecimalsResponse{}, nil
}

func (c *MockChainClient) FetchChainSpotMarkets(ctx context.Context, status string, marketIDs []string) (*exchangetypes.QuerySpotMarketsResponse, error) {
	var response *exchangetypes.QuerySpotMarketsResponse
	var localError error
	if len(c.QuerySpotMarketsResponses) > 0 {
		response = c.QuerySpotMarketsResponses[0]
		c.QuerySpotMarketsResponses = c.QuerySpotMarketsResponses[1:]
		localError = nil
	} else {
		response = &exchangetypes.QuerySpotMarketsResponse{}
		localError = errors.New("there are no responses configured")
	}

	return response, localError
}

func (c *MockChainClient) FetchChainSpotMarket(ctx context.Context, marketId string) (*exchangetypes.QuerySpotMarketResponse, error) {
	return &exchangetypes.QuerySpotMarketResponse{}, nil
}

func (c *MockChainClient) FetchChainFullSpotMarkets(ctx context.Context, status string, marketIDs []string, withMidPriceAndTob bool) (*exchangetypes.QueryFullSpotMarketsResponse, error) {
	return &exchangetypes.QueryFullSpotMarketsResponse{}, nil
}

func (c *MockChainClient) FetchChainFullSpotMarket(ctx context.Context, marketId string, withMidPriceAndTob bool) (*exchangetypes.QueryFullSpotMarketResponse, error) {
	return &exchangetypes.QueryFullSpotMarketResponse{}, nil
}

func (c *MockChainClient) FetchChainSpotOrderbook(ctx context.Context, marketId string, limit uint64, orderSide exchangetypes.OrderSide, limitCumulativeNotional, limitCumulativeQuantity sdkmath.LegacyDec) (*exchangetypes.QuerySpotOrderbookResponse, error) {
	return &exchangetypes.QuerySpotOrderbookResponse{}, nil
}

func (c *MockChainClient) FetchChainTraderSpotOrders(ctx context.Context, marketId, subaccountId string) (*exchangetypes.QueryTraderSpotOrdersResponse, error) {
	return &exchangetypes.QueryTraderSpotOrdersResponse{}, nil
}

func (c *MockChainClient) FetchChainAccountAddressSpotOrders(ctx context.Context, marketId, address string) (*exchangetypes.QueryAccountAddressSpotOrdersResponse, error) {
	return &exchangetypes.QueryAccountAddressSpotOrdersResponse{}, nil
}

func (c *MockChainClient) FetchChainSpotOrdersByHashes(ctx context.Context, marketId, subaccountId string, orderHashes []string) (*exchangetypes.QuerySpotOrdersByHashesResponse, error) {
	return &exchangetypes.QuerySpotOrdersByHashesResponse{}, nil
}

func (c *MockChainClient) FetchChainSubaccountOrders(ctx context.Context, subaccountId, marketId string) (*exchangetypes.QuerySubaccountOrdersResponse, error) {
	return &exchangetypes.QuerySubaccountOrdersResponse{}, nil
}

func (c *MockChainClient) FetchChainTraderSpotTransientOrders(ctx context.Context, marketId, subaccountId string) (*exchangetypes.QueryTraderSpotOrdersResponse, error) {
	return &exchangetypes.QueryTraderSpotOrdersResponse{}, nil
}

func (c *MockChainClient) FetchSpotMidPriceAndTOB(ctx context.Context, marketId string) (*exchangetypes.QuerySpotMidPriceAndTOBResponse, error) {
	return &exchangetypes.QuerySpotMidPriceAndTOBResponse{}, nil
}

func (c *MockChainClient) FetchDerivativeMidPriceAndTOB(ctx context.Context, marketId string) (*exchangetypes.QueryDerivativeMidPriceAndTOBResponse, error) {
	return &exchangetypes.QueryDerivativeMidPriceAndTOBResponse{}, nil
}

func (c *MockChainClient) FetchChainDerivativeOrderbook(ctx context.Context, marketId string, limit uint64, limitCumulativeNotional sdkmath.LegacyDec) (*exchangetypes.QueryDerivativeOrderbookResponse, error) {
	return &exchangetypes.QueryDerivativeOrderbookResponse{}, nil
}

func (c *MockChainClient) FetchChainTraderDerivativeOrders(ctx context.Context, marketId, subaccountId string) (*exchangetypes.QueryTraderDerivativeOrdersResponse, error) {
	return &exchangetypes.QueryTraderDerivativeOrdersResponse{}, nil
}

func (c *MockChainClient) FetchChainAccountAddressDerivativeOrders(ctx context.Context, marketId, address string) (*exchangetypes.QueryAccountAddressDerivativeOrdersResponse, error) {
	return &exchangetypes.QueryAccountAddressDerivativeOrdersResponse{}, nil
}

func (c *MockChainClient) FetchChainDerivativeOrdersByHashes(ctx context.Context, marketId, subaccountId string, orderHashes []string) (*exchangetypes.QueryDerivativeOrdersByHashesResponse, error) {
	return &exchangetypes.QueryDerivativeOrdersByHashesResponse{}, nil
}

func (c *MockChainClient) FetchChainTraderDerivativeTransientOrders(ctx context.Context, marketId, subaccountId string) (*exchangetypes.QueryTraderDerivativeOrdersResponse, error) {
	return &exchangetypes.QueryTraderDerivativeOrdersResponse{}, nil
}

func (c *MockChainClient) FetchChainDerivativeMarkets(ctx context.Context, status string, marketIDs []string, withMidPriceAndTob bool) (*exchangetypes.QueryDerivativeMarketsResponse, error) {
	var response *exchangetypes.QueryDerivativeMarketsResponse
	var localError error
	if len(c.QueryDerivativeMarketsResponses) > 0 {
		response = c.QueryDerivativeMarketsResponses[0]
		c.QueryDerivativeMarketsResponses = c.QueryDerivativeMarketsResponses[1:]
		localError = nil
	} else {
		response = &exchangetypes.QueryDerivativeMarketsResponse{}
		localError = errors.New("there are no responses configured")
	}

	return response, localError
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

func (c *MockChainClient) FetchChainSubaccountPositionInMarket(ctx context.Context, subaccountId, marketId string) (*exchangetypes.QuerySubaccountPositionInMarketResponse, error) {
	return &exchangetypes.QuerySubaccountPositionInMarketResponse{}, nil
}

func (c *MockChainClient) FetchChainSubaccountEffectivePositionInMarket(ctx context.Context, subaccountId, marketId string) (*exchangetypes.QuerySubaccountEffectivePositionInMarketResponse, error) {
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
	var response *exchangetypes.QueryBinaryMarketsResponse
	var localError error
	if len(c.QueryBinaryMarketsResponses) > 0 {
		response = c.QueryBinaryMarketsResponses[0]
		c.QueryBinaryMarketsResponses = c.QueryBinaryMarketsResponses[1:]
		localError = nil
	} else {
		response = &exchangetypes.QueryBinaryMarketsResponse{}
		localError = errors.New("there are no responses configured")
	}

	return response, localError
}

func (c *MockChainClient) FetchTraderDerivativeConditionalOrders(ctx context.Context, subaccountId, marketId string) (*exchangetypes.QueryTraderDerivativeConditionalOrdersResponse, error) {
	return &exchangetypes.QueryTraderDerivativeConditionalOrdersResponse{}, nil
}

func (c *MockChainClient) FetchMarketAtomicExecutionFeeMultiplier(ctx context.Context, marketId string) (*exchangetypes.QueryMarketAtomicExecutionFeeMultiplierResponse, error) {
	return &exchangetypes.QueryMarketAtomicExecutionFeeMultiplierResponse{}, nil
}

// Chain exchange V2 module
func (c *MockChainClient) FetchSubaccountDepositsV2(ctx context.Context, subaccountId string) (*exchangev2types.QuerySubaccountDepositsResponse, error) {
	return &exchangev2types.QuerySubaccountDepositsResponse{}, nil
}

func (c *MockChainClient) FetchSubaccountDepositV2(ctx context.Context, subaccountId, denom string) (*exchangev2types.QuerySubaccountDepositResponse, error) {
	return &exchangev2types.QuerySubaccountDepositResponse{}, nil
}

func (c *MockChainClient) FetchExchangeBalancesV2(_ context.Context) (*exchangev2types.QueryExchangeBalancesResponse, error) {
	return &exchangev2types.QueryExchangeBalancesResponse{}, nil
}

func (c *MockChainClient) FetchAggregateVolumeV2(ctx context.Context, account string) (*exchangev2types.QueryAggregateVolumeResponse, error) {
	return &exchangev2types.QueryAggregateVolumeResponse{}, nil
}

func (c *MockChainClient) FetchAggregateVolumesV2(ctx context.Context, accounts, marketIDs []string) (*exchangev2types.QueryAggregateVolumesResponse, error) {
	return &exchangev2types.QueryAggregateVolumesResponse{}, nil
}

func (c *MockChainClient) FetchAggregateMarketVolumeV2(ctx context.Context, marketId string) (*exchangev2types.QueryAggregateMarketVolumeResponse, error) {
	return &exchangev2types.QueryAggregateMarketVolumeResponse{}, nil
}

func (c *MockChainClient) FetchAggregateMarketVolumesV2(ctx context.Context, marketIDs []string) (*exchangev2types.QueryAggregateMarketVolumesResponse, error) {
	return &exchangev2types.QueryAggregateMarketVolumesResponse{}, nil
}

func (c *MockChainClient) FetchDenomDecimalV2(ctx context.Context, denom string) (*exchangev2types.QueryDenomDecimalResponse, error) {
	return &exchangev2types.QueryDenomDecimalResponse{}, nil
}

func (c *MockChainClient) FetchDenomDecimalsV2(ctx context.Context, denoms []string) (*exchangev2types.QueryDenomDecimalsResponse, error) {
	return &exchangev2types.QueryDenomDecimalsResponse{}, nil
}

func (c *MockChainClient) FetchChainSpotMarketsV2(ctx context.Context, status string, marketIDs []string) (*exchangev2types.QuerySpotMarketsResponse, error) {
	var response *exchangev2types.QuerySpotMarketsResponse
	var localError error
	if len(c.QuerySpotMarketsV2Responses) > 0 {
		response = c.QuerySpotMarketsV2Responses[0]
		c.QuerySpotMarketsV2Responses = c.QuerySpotMarketsV2Responses[1:]
		localError = nil
	} else {
		response = &exchangev2types.QuerySpotMarketsResponse{}
		localError = errors.New("there are no responses configured")
	}

	return response, localError
}

func (c *MockChainClient) FetchChainSpotMarketV2(ctx context.Context, marketId string) (*exchangev2types.QuerySpotMarketResponse, error) {
	return &exchangev2types.QuerySpotMarketResponse{}, nil
}

func (c *MockChainClient) FetchChainFullSpotMarketsV2(ctx context.Context, status string, marketIDs []string, withMidPriceAndTob bool) (*exchangev2types.QueryFullSpotMarketsResponse, error) {
	return &exchangev2types.QueryFullSpotMarketsResponse{}, nil
}

func (c *MockChainClient) FetchChainFullSpotMarketV2(ctx context.Context, marketId string, withMidPriceAndTob bool) (*exchangev2types.QueryFullSpotMarketResponse, error) {
	return &exchangev2types.QueryFullSpotMarketResponse{}, nil
}

func (c *MockChainClient) FetchChainSpotOrderbookV2(ctx context.Context, marketId string, limit uint64, orderSide exchangev2types.OrderSide, limitCumulativeNotional, limitCumulativeQuantity sdkmath.LegacyDec) (*exchangev2types.QuerySpotOrderbookResponse, error) {
	return &exchangev2types.QuerySpotOrderbookResponse{}, nil
}

func (c *MockChainClient) FetchChainTraderSpotOrdersV2(ctx context.Context, marketId, subaccountId string) (*exchangev2types.QueryTraderSpotOrdersResponse, error) {
	return &exchangev2types.QueryTraderSpotOrdersResponse{}, nil
}

func (c *MockChainClient) FetchChainAccountAddressSpotOrdersV2(ctx context.Context, marketId, address string) (*exchangev2types.QueryAccountAddressSpotOrdersResponse, error) {
	return &exchangev2types.QueryAccountAddressSpotOrdersResponse{}, nil
}

func (c *MockChainClient) FetchChainSpotOrdersByHashesV2(ctx context.Context, marketId, subaccountId string, orderHashes []string) (*exchangev2types.QuerySpotOrdersByHashesResponse, error) {
	return &exchangev2types.QuerySpotOrdersByHashesResponse{}, nil
}

func (c *MockChainClient) FetchChainSubaccountOrdersV2(ctx context.Context, subaccountId, marketId string) (*exchangev2types.QuerySubaccountOrdersResponse, error) {
	return &exchangev2types.QuerySubaccountOrdersResponse{}, nil
}

func (c *MockChainClient) FetchChainTraderSpotTransientOrdersV2(ctx context.Context, marketId, subaccountId string) (*exchangev2types.QueryTraderSpotOrdersResponse, error) {
	return &exchangev2types.QueryTraderSpotOrdersResponse{}, nil
}

func (c *MockChainClient) FetchSpotMidPriceAndTOBV2(ctx context.Context, marketId string) (*exchangev2types.QuerySpotMidPriceAndTOBResponse, error) {
	return &exchangev2types.QuerySpotMidPriceAndTOBResponse{}, nil
}

func (c *MockChainClient) FetchDerivativeMidPriceAndTOBV2(ctx context.Context, marketId string) (*exchangev2types.QueryDerivativeMidPriceAndTOBResponse, error) {
	return &exchangev2types.QueryDerivativeMidPriceAndTOBResponse{}, nil
}

func (c *MockChainClient) FetchChainDerivativeOrderbookV2(ctx context.Context, marketId string, limit uint64, limitCumulativeNotional sdkmath.LegacyDec) (*exchangev2types.QueryDerivativeOrderbookResponse, error) {
	return &exchangev2types.QueryDerivativeOrderbookResponse{}, nil
}

func (c *MockChainClient) FetchChainTraderDerivativeOrdersV2(ctx context.Context, marketId, subaccountId string) (*exchangev2types.QueryTraderDerivativeOrdersResponse, error) {
	return &exchangev2types.QueryTraderDerivativeOrdersResponse{}, nil
}

func (c *MockChainClient) FetchChainAccountAddressDerivativeOrdersV2(ctx context.Context, marketId, address string) (*exchangev2types.QueryAccountAddressDerivativeOrdersResponse, error) {
	return &exchangev2types.QueryAccountAddressDerivativeOrdersResponse{}, nil
}

func (c *MockChainClient) FetchChainDerivativeOrdersByHashesV2(ctx context.Context, marketId, subaccountId string, orderHashes []string) (*exchangev2types.QueryDerivativeOrdersByHashesResponse, error) {
	return &exchangev2types.QueryDerivativeOrdersByHashesResponse{}, nil
}

func (c *MockChainClient) FetchChainTraderDerivativeTransientOrdersV2(ctx context.Context, marketId, subaccountId string) (*exchangev2types.QueryTraderDerivativeOrdersResponse, error) {
	return &exchangev2types.QueryTraderDerivativeOrdersResponse{}, nil
}

func (c *MockChainClient) FetchChainDerivativeMarketsV2(ctx context.Context, status string, marketIDs []string, withMidPriceAndTob bool) (*exchangev2types.QueryDerivativeMarketsResponse, error) {
	var response *exchangev2types.QueryDerivativeMarketsResponse
	var localError error
	if len(c.QueryDerivativeMarketsV2Responses) > 0 {
		response = c.QueryDerivativeMarketsV2Responses[0]
		c.QueryDerivativeMarketsV2Responses = c.QueryDerivativeMarketsV2Responses[1:]
		localError = nil
	} else {
		response = &exchangev2types.QueryDerivativeMarketsResponse{}
		localError = errors.New("there are no responses configured")
	}

	return response, localError
}

func (c *MockChainClient) FetchChainDerivativeMarketV2(ctx context.Context, marketId string) (*exchangev2types.QueryDerivativeMarketResponse, error) {
	return &exchangev2types.QueryDerivativeMarketResponse{}, nil
}

func (c *MockChainClient) FetchDerivativeMarketAddressV2(ctx context.Context, marketId string) (*exchangev2types.QueryDerivativeMarketAddressResponse, error) {
	return &exchangev2types.QueryDerivativeMarketAddressResponse{}, nil
}

func (c *MockChainClient) FetchSubaccountTradeNonceV2(ctx context.Context, subaccountId string) (*exchangev2types.QuerySubaccountTradeNonceResponse, error) {
	return &exchangev2types.QuerySubaccountTradeNonceResponse{}, nil
}

func (c *MockChainClient) FetchChainPositionsV2(ctx context.Context) (*exchangev2types.QueryPositionsResponse, error) {
	return &exchangev2types.QueryPositionsResponse{}, nil
}

func (c *MockChainClient) FetchChainSubaccountPositionsV2(ctx context.Context, subaccountId string) (*exchangev2types.QuerySubaccountPositionsResponse, error) {
	return &exchangev2types.QuerySubaccountPositionsResponse{}, nil
}

func (c *MockChainClient) FetchChainSubaccountPositionInMarketV2(ctx context.Context, subaccountId, marketId string) (*exchangev2types.QuerySubaccountPositionInMarketResponse, error) {
	return &exchangev2types.QuerySubaccountPositionInMarketResponse{}, nil
}

func (c *MockChainClient) FetchChainSubaccountEffectivePositionInMarketV2(ctx context.Context, subaccountId, marketId string) (*exchangev2types.QuerySubaccountEffectivePositionInMarketResponse, error) {
	return &exchangev2types.QuerySubaccountEffectivePositionInMarketResponse{}, nil
}

func (c *MockChainClient) FetchChainPerpetualMarketInfoV2(ctx context.Context, marketId string) (*exchangev2types.QueryPerpetualMarketInfoResponse, error) {
	return &exchangev2types.QueryPerpetualMarketInfoResponse{}, nil
}

func (c *MockChainClient) FetchChainExpiryFuturesMarketInfoV2(ctx context.Context, marketId string) (*exchangev2types.QueryExpiryFuturesMarketInfoResponse, error) {
	return &exchangev2types.QueryExpiryFuturesMarketInfoResponse{}, nil
}

func (c *MockChainClient) FetchChainPerpetualMarketFundingV2(ctx context.Context, marketId string) (*exchangev2types.QueryPerpetualMarketFundingResponse, error) {
	return &exchangev2types.QueryPerpetualMarketFundingResponse{}, nil
}

func (c *MockChainClient) FetchSubaccountOrderMetadataV2(ctx context.Context, subaccountId string) (*exchangev2types.QuerySubaccountOrderMetadataResponse, error) {
	return &exchangev2types.QuerySubaccountOrderMetadataResponse{}, nil
}

func (c *MockChainClient) FetchTradeRewardPointsV2(ctx context.Context, accounts []string) (*exchangev2types.QueryTradeRewardPointsResponse, error) {
	return &exchangev2types.QueryTradeRewardPointsResponse{}, nil
}

func (c *MockChainClient) FetchPendingTradeRewardPointsV2(ctx context.Context, accounts []string) (*exchangev2types.QueryTradeRewardPointsResponse, error) {
	return &exchangev2types.QueryTradeRewardPointsResponse{}, nil
}

func (c *MockChainClient) FetchTradeRewardCampaignV2(ctx context.Context) (*exchangev2types.QueryTradeRewardCampaignResponse, error) {
	return &exchangev2types.QueryTradeRewardCampaignResponse{}, nil
}

func (c *MockChainClient) FetchFeeDiscountAccountInfoV2(ctx context.Context, account string) (*exchangev2types.QueryFeeDiscountAccountInfoResponse, error) {
	return &exchangev2types.QueryFeeDiscountAccountInfoResponse{}, nil
}

func (c *MockChainClient) FetchFeeDiscountScheduleV2(ctx context.Context) (*exchangev2types.QueryFeeDiscountScheduleResponse, error) {
	return &exchangev2types.QueryFeeDiscountScheduleResponse{}, nil
}

func (c *MockChainClient) FetchBalanceMismatchesV2(ctx context.Context, dustFactor int64) (*exchangev2types.QueryBalanceMismatchesResponse, error) {
	return &exchangev2types.QueryBalanceMismatchesResponse{}, nil
}

func (c *MockChainClient) FetchBalanceWithBalanceHoldsV2(ctx context.Context) (*exchangev2types.QueryBalanceWithBalanceHoldsResponse, error) {
	return &exchangev2types.QueryBalanceWithBalanceHoldsResponse{}, nil
}

func (c *MockChainClient) FetchFeeDiscountTierStatisticsV2(ctx context.Context) (*exchangev2types.QueryFeeDiscountTierStatisticsResponse, error) {
	return &exchangev2types.QueryFeeDiscountTierStatisticsResponse{}, nil
}

func (c *MockChainClient) FetchMitoVaultInfosV2(ctx context.Context) (*exchangev2types.MitoVaultInfosResponse, error) {
	return &exchangev2types.MitoVaultInfosResponse{}, nil
}

func (c *MockChainClient) FetchMarketIDFromVaultV2(ctx context.Context, vaultAddress string) (*exchangev2types.QueryMarketIDFromVaultResponse, error) {
	return &exchangev2types.QueryMarketIDFromVaultResponse{}, nil
}

func (c *MockChainClient) FetchHistoricalTradeRecordsV2(ctx context.Context, marketId string) (*exchangev2types.QueryHistoricalTradeRecordsResponse, error) {
	return &exchangev2types.QueryHistoricalTradeRecordsResponse{}, nil
}

func (c *MockChainClient) FetchIsOptedOutOfRewardsV2(ctx context.Context, account string) (*exchangev2types.QueryIsOptedOutOfRewardsResponse, error) {
	return &exchangev2types.QueryIsOptedOutOfRewardsResponse{}, nil
}

func (c *MockChainClient) FetchOptedOutOfRewardsAccountsV2(ctx context.Context) (*exchangev2types.QueryOptedOutOfRewardsAccountsResponse, error) {
	return &exchangev2types.QueryOptedOutOfRewardsAccountsResponse{}, nil
}

func (c *MockChainClient) FetchMarketVolatilityV2(ctx context.Context, marketId string, tradeHistoryOptions *exchangev2types.TradeHistoryOptions) (*exchangev2types.QueryMarketVolatilityResponse, error) {
	return &exchangev2types.QueryMarketVolatilityResponse{}, nil
}

func (c *MockChainClient) FetchChainBinaryOptionsMarketsV2(ctx context.Context, status string) (*exchangev2types.QueryBinaryMarketsResponse, error) {
	var response *exchangev2types.QueryBinaryMarketsResponse
	var localError error
	if len(c.QueryBinaryMarketsV2Responses) > 0 {
		response = c.QueryBinaryMarketsV2Responses[0]
		c.QueryBinaryMarketsV2Responses = c.QueryBinaryMarketsV2Responses[1:]
		localError = nil
	} else {
		response = &exchangev2types.QueryBinaryMarketsResponse{}
		localError = errors.New("there are no responses configured")
	}

	return response, localError
}

func (c *MockChainClient) FetchTraderDerivativeConditionalOrdersV2(ctx context.Context, subaccountId, marketId string) (*exchangev2types.QueryTraderDerivativeConditionalOrdersResponse, error) {
	return &exchangev2types.QueryTraderDerivativeConditionalOrdersResponse{}, nil
}

func (c *MockChainClient) FetchMarketAtomicExecutionFeeMultiplierV2(ctx context.Context, marketId string) (*exchangev2types.QueryMarketAtomicExecutionFeeMultiplierResponse, error) {
	return &exchangev2types.QueryMarketAtomicExecutionFeeMultiplierResponse{}, nil
}

func (c *MockChainClient) FetchActiveStakeGrantV2(ctx context.Context, grantee string) (*exchangev2types.QueryActiveStakeGrantResponse, error) {
	return &exchangev2types.QueryActiveStakeGrantResponse{}, nil
}

func (c *MockChainClient) FetchGrantAuthorizationV2(ctx context.Context, granter, grantee string) (*exchangev2types.QueryGrantAuthorizationResponse, error) {
	return &exchangev2types.QueryGrantAuthorizationResponse{}, nil
}

func (c *MockChainClient) FetchGrantAuthorizationsV2(ctx context.Context, granter string) (*exchangev2types.QueryGrantAuthorizationsResponse, error) {
	return &exchangev2types.QueryGrantAuthorizationsResponse{}, nil
}

func (c *MockChainClient) FetchL3DerivativeOrderbookV2(ctx context.Context, marketId string) (*exchangev2types.QueryFullDerivativeOrderbookResponse, error) {
	return &exchangev2types.QueryFullDerivativeOrderbookResponse{}, nil
}

func (c *MockChainClient) FetchL3SpotOrderbookV2(ctx context.Context, marketId string) (*exchangev2types.QueryFullSpotOrderbookResponse, error) {
	return &exchangev2types.QueryFullSpotOrderbookResponse{}, nil
}

// Tendermint module

func (c *MockChainClient) FetchNodeInfo(ctx context.Context) (*cmtservice.GetNodeInfoResponse, error) {
	return &cmtservice.GetNodeInfoResponse{}, nil
}

func (c *MockChainClient) FetchSyncing(ctx context.Context) (*cmtservice.GetSyncingResponse, error) {
	return &cmtservice.GetSyncingResponse{}, nil
}

func (c *MockChainClient) FetchLatestBlock(ctx context.Context) (*cmtservice.GetLatestBlockResponse, error) {
	return &cmtservice.GetLatestBlockResponse{}, nil
}

func (c *MockChainClient) FetchBlockByHeight(ctx context.Context, height int64) (*cmtservice.GetBlockByHeightResponse, error) {
	return &cmtservice.GetBlockByHeightResponse{}, nil
}

func (c *MockChainClient) FetchLatestValidatorSet(ctx context.Context) (*cmtservice.GetLatestValidatorSetResponse, error) {
	return &cmtservice.GetLatestValidatorSetResponse{}, nil
}

func (c *MockChainClient) FetchValidatorSetByHeight(ctx context.Context, height int64, pagination *query.PageRequest) (*cmtservice.GetValidatorSetByHeightResponse, error) {
	return &cmtservice.GetValidatorSetByHeightResponse{}, nil
}

func (c *MockChainClient) ABCIQuery(ctx context.Context, path string, data []byte, height int64, prove bool) (*cmtservice.ABCIQueryResponse, error) {
	return &cmtservice.ABCIQueryResponse{}, nil
}

// IBC Transfer module
func (c *MockChainClient) FetchDenomTrace(ctx context.Context, hash string) (*ibctransfertypes.QueryDenomTraceResponse, error) {
	return &ibctransfertypes.QueryDenomTraceResponse{}, nil
}

func (c *MockChainClient) FetchDenomTraces(ctx context.Context, pagination *query.PageRequest) (*ibctransfertypes.QueryDenomTracesResponse, error) {
	return &ibctransfertypes.QueryDenomTracesResponse{}, nil
}

func (c *MockChainClient) FetchDenomHash(ctx context.Context, trace string) (*ibctransfertypes.QueryDenomHashResponse, error) {
	return &ibctransfertypes.QueryDenomHashResponse{}, nil
}

func (c *MockChainClient) FetchEscrowAddress(ctx context.Context, portId, channelId string) (*ibctransfertypes.QueryEscrowAddressResponse, error) {
	return &ibctransfertypes.QueryEscrowAddressResponse{}, nil
}

func (c *MockChainClient) FetchTotalEscrowForDenom(ctx context.Context, denom string) (*ibctransfertypes.QueryTotalEscrowForDenomResponse, error) {
	return &ibctransfertypes.QueryTotalEscrowForDenomResponse{}, nil
}

// IBC Core Channel module
func (c *MockChainClient) FetchIBCChannel(ctx context.Context, portId, channelId string) (*ibcchanneltypes.QueryChannelResponse, error) {
	return &ibcchanneltypes.QueryChannelResponse{}, nil
}

func (c *MockChainClient) FetchIBCChannels(ctx context.Context, pagination *query.PageRequest) (*ibcchanneltypes.QueryChannelsResponse, error) {
	return &ibcchanneltypes.QueryChannelsResponse{}, nil
}

func (c *MockChainClient) FetchIBCConnectionChannels(ctx context.Context, connection string, pagination *query.PageRequest) (*ibcchanneltypes.QueryConnectionChannelsResponse, error) {
	return &ibcchanneltypes.QueryConnectionChannelsResponse{}, nil
}

func (c *MockChainClient) FetchIBCChannelClientState(ctx context.Context, portId, channelId string) (*ibcchanneltypes.QueryChannelClientStateResponse, error) {
	return &ibcchanneltypes.QueryChannelClientStateResponse{}, nil
}

func (c *MockChainClient) FetchIBCChannelConsensusState(ctx context.Context, portId, channelId string, revisionNumber, revisionHeight uint64) (*ibcchanneltypes.QueryChannelConsensusStateResponse, error) {
	return &ibcchanneltypes.QueryChannelConsensusStateResponse{}, nil
}

func (c *MockChainClient) FetchIBCPacketCommitment(ctx context.Context, portId, channelId string, sequence uint64) (*ibcchanneltypes.QueryPacketCommitmentResponse, error) {
	return &ibcchanneltypes.QueryPacketCommitmentResponse{}, nil
}

func (c *MockChainClient) FetchIBCPacketCommitments(ctx context.Context, portId, channelId string, pagination *query.PageRequest) (*ibcchanneltypes.QueryPacketCommitmentsResponse, error) {
	return &ibcchanneltypes.QueryPacketCommitmentsResponse{}, nil
}

func (c *MockChainClient) FetchIBCPacketReceipt(ctx context.Context, portId, channelId string, sequence uint64) (*ibcchanneltypes.QueryPacketReceiptResponse, error) {
	return &ibcchanneltypes.QueryPacketReceiptResponse{}, nil
}

func (c *MockChainClient) FetchIBCPacketAcknowledgement(ctx context.Context, portId, channelId string, sequence uint64) (*ibcchanneltypes.QueryPacketAcknowledgementResponse, error) {
	return &ibcchanneltypes.QueryPacketAcknowledgementResponse{}, nil
}

func (c *MockChainClient) FetchIBCPacketAcknowledgements(ctx context.Context, portId, channelId string, packetCommitmentSequences []uint64, pagination *query.PageRequest) (*ibcchanneltypes.QueryPacketAcknowledgementsResponse, error) {
	return &ibcchanneltypes.QueryPacketAcknowledgementsResponse{}, nil
}

func (c *MockChainClient) FetchIBCUnreceivedPackets(ctx context.Context, portId, channelId string, packetCommitmentSequences []uint64) (*ibcchanneltypes.QueryUnreceivedPacketsResponse, error) {
	return &ibcchanneltypes.QueryUnreceivedPacketsResponse{}, nil
}

func (c *MockChainClient) FetchIBCUnreceivedAcks(ctx context.Context, portId, channelId string, packetAckSequences []uint64) (*ibcchanneltypes.QueryUnreceivedAcksResponse, error) {
	return &ibcchanneltypes.QueryUnreceivedAcksResponse{}, nil
}

func (c *MockChainClient) FetchIBCNextSequenceReceive(ctx context.Context, portId, channelId string) (*ibcchanneltypes.QueryNextSequenceReceiveResponse, error) {
	return &ibcchanneltypes.QueryNextSequenceReceiveResponse{}, nil
}

// IBC Core Chain module
func (c *MockChainClient) FetchIBCClientState(ctx context.Context, clientId string) (*ibcclienttypes.QueryClientStateResponse, error) {
	return &ibcclienttypes.QueryClientStateResponse{}, nil
}

func (c *MockChainClient) FetchIBCClientStates(ctx context.Context, pagination *query.PageRequest) (*ibcclienttypes.QueryClientStatesResponse, error) {
	return &ibcclienttypes.QueryClientStatesResponse{}, nil
}

func (c *MockChainClient) FetchIBCConsensusState(ctx context.Context, clientId string, revisionNumber, revisionHeight uint64, latestHeight bool) (*ibcclienttypes.QueryConsensusStateResponse, error) {
	return &ibcclienttypes.QueryConsensusStateResponse{}, nil
}

func (c *MockChainClient) FetchIBCConsensusStates(ctx context.Context, clientId string, pagination *query.PageRequest) (*ibcclienttypes.QueryConsensusStatesResponse, error) {
	return &ibcclienttypes.QueryConsensusStatesResponse{}, nil
}

func (c *MockChainClient) FetchIBCConsensusStateHeights(ctx context.Context, clientId string, pagination *query.PageRequest) (*ibcclienttypes.QueryConsensusStateHeightsResponse, error) {
	return &ibcclienttypes.QueryConsensusStateHeightsResponse{}, nil
}

func (c *MockChainClient) FetchIBCClientStatus(ctx context.Context, clientId string) (*ibcclienttypes.QueryClientStatusResponse, error) {
	return &ibcclienttypes.QueryClientStatusResponse{}, nil
}

func (c *MockChainClient) FetchIBCClientParams(ctx context.Context) (*ibcclienttypes.QueryClientParamsResponse, error) {
	return &ibcclienttypes.QueryClientParamsResponse{}, nil
}

func (c *MockChainClient) FetchIBCUpgradedClientState(ctx context.Context) (*ibcclienttypes.QueryUpgradedClientStateResponse, error) {
	return &ibcclienttypes.QueryUpgradedClientStateResponse{}, nil
}

func (c *MockChainClient) FetchIBCUpgradedConsensusState(ctx context.Context) (*ibcclienttypes.QueryUpgradedConsensusStateResponse, error) {
	return &ibcclienttypes.QueryUpgradedConsensusStateResponse{}, nil
}

// IBC Core Connection module
func (c *MockChainClient) FetchIBCConnection(ctx context.Context, connectionId string) (*ibcconnectiontypes.QueryConnectionResponse, error) {
	return &ibcconnectiontypes.QueryConnectionResponse{}, nil
}

func (c *MockChainClient) FetchIBCConnections(ctx context.Context, pagination *query.PageRequest) (*ibcconnectiontypes.QueryConnectionsResponse, error) {
	return &ibcconnectiontypes.QueryConnectionsResponse{}, nil
}

func (c *MockChainClient) FetchIBCClientConnections(ctx context.Context, clientId string) (*ibcconnectiontypes.QueryClientConnectionsResponse, error) {
	return &ibcconnectiontypes.QueryClientConnectionsResponse{}, nil
}

func (c *MockChainClient) FetchIBCConnectionClientState(ctx context.Context, connectionId string) (*ibcconnectiontypes.QueryConnectionClientStateResponse, error) {
	return &ibcconnectiontypes.QueryConnectionClientStateResponse{}, nil
}

func (c *MockChainClient) FetchIBCConnectionConsensusState(ctx context.Context, connectionId string, revisionNumber, revisionHeight uint64) (*ibcconnectiontypes.QueryConnectionConsensusStateResponse, error) {
	return &ibcconnectiontypes.QueryConnectionConsensusStateResponse{}, nil
}

func (c *MockChainClient) FetchIBCConnectionParams(ctx context.Context) (*ibcconnectiontypes.QueryConnectionParamsResponse, error) {
	return &ibcconnectiontypes.QueryConnectionParamsResponse{}, nil
}

// Permissions module

func (c *MockChainClient) FetchAllNamespaces(ctx context.Context) (*permissionstypes.QueryAllNamespacesResponse, error) {
	return &permissionstypes.QueryAllNamespacesResponse{}, nil
}

func (c *MockChainClient) FetchNamespaceByDenom(ctx context.Context, denom string, includeRoles bool) (*permissionstypes.QueryNamespaceByDenomResponse, error) {
	return &permissionstypes.QueryNamespaceByDenomResponse{}, nil
}

func (c *MockChainClient) FetchAddressRoles(ctx context.Context, denom, address string) (*permissionstypes.QueryAddressRolesResponse, error) {
	return &permissionstypes.QueryAddressRolesResponse{}, nil
}

func (c *MockChainClient) FetchAddressesByRole(ctx context.Context, denom, role string) (*permissionstypes.QueryAddressesByRoleResponse, error) {
	return &permissionstypes.QueryAddressesByRoleResponse{}, nil
}

func (c *MockChainClient) FetchVouchersForAddress(ctx context.Context, address string) (*permissionstypes.QueryVouchersForAddressResponse, error) {
	return &permissionstypes.QueryVouchersForAddressResponse{}, nil
}

func (c *MockChainClient) GetNetwork() common.Network {
	return c.Network
}
