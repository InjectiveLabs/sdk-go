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

	erc20types "github.com/InjectiveLabs/sdk-go/chain/erc20/types"
	evmtypes "github.com/InjectiveLabs/sdk-go/chain/evm/types"
	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	exchangev2types "github.com/InjectiveLabs/sdk-go/chain/exchange/types/v2"
	permissionstypes "github.com/InjectiveLabs/sdk-go/chain/permissions/types"
	chainstreamv2types "github.com/InjectiveLabs/sdk-go/chain/stream/types/v2"
	tokenfactorytypes "github.com/InjectiveLabs/sdk-go/chain/tokenfactory/types"
	txfeestypes "github.com/InjectiveLabs/sdk-go/chain/txfees/types"
	injectiveclient "github.com/InjectiveLabs/sdk-go/client"
	"github.com/InjectiveLabs/sdk-go/client/common"
)

var _ ChainClientV2 = &MockChainClientV2{}

type MockChainClientV2 struct {
	Network                           common.Network
	DenomsMetadataResponses           []*banktypes.QueryDenomsMetadataResponse
	QuerySpotMarketsV2Responses       []*exchangev2types.QuerySpotMarketsResponse
	QueryDerivativeMarketsV2Responses []*exchangev2types.QueryDerivativeMarketsResponse
	QueryBinaryMarketsV2Responses     []*exchangev2types.QueryBinaryMarketsResponse
}

func (c *MockChainClientV2) CanSignTransactions() bool {
	return true
}

func (c *MockChainClientV2) FromAddress() sdk.AccAddress {
	return sdk.AccAddress{}
}

func (c *MockChainClientV2) QueryClient() *grpc.ClientConn {
	return &grpc.ClientConn{}
}

func (c *MockChainClientV2) ClientContext() client.Context {
	return client.Context{}
}

func (c *MockChainClientV2) GetAccNonce() (accNum, accSeq uint64) {
	return 1, 2
}

func (c *MockChainClientV2) SimulateMsg(clientCtx client.Context, msgs ...sdk.Msg) (*txtypes.SimulateResponse, error) {
	return &txtypes.SimulateResponse{}, nil
}

func (c *MockChainClientV2) AsyncBroadcastMsg(msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error) {
	return &txtypes.BroadcastTxResponse{}, nil
}

func (c *MockChainClientV2) SyncBroadcastMsg(msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error) {
	return &txtypes.BroadcastTxResponse{}, nil
}

func (c *MockChainClientV2) BroadcastMsg(broadcastMode txtypes.BroadcastMode, msgs ...sdk.Msg) (*txtypes.BroadcastTxRequest, *txtypes.BroadcastTxResponse, error) {
	return &txtypes.BroadcastTxRequest{}, &txtypes.BroadcastTxResponse{}, nil
}

func (c *MockChainClientV2) BuildSignedTx(clientCtx client.Context, accNum, accSeq, initialGas uint64, gasPrice uint64, msg ...sdk.Msg) ([]byte, error) {
	return []byte(nil), nil
}

func (c *MockChainClientV2) SyncBroadcastSignedTx(tyBytes []byte) (*txtypes.BroadcastTxResponse, error) {
	return &txtypes.BroadcastTxResponse{}, nil
}

func (c *MockChainClientV2) AsyncBroadcastSignedTx(txBytes []byte) (*txtypes.BroadcastTxResponse, error) {
	return &txtypes.BroadcastTxResponse{}, nil
}

func (c *MockChainClientV2) BroadcastSignedTx(txBytes []byte, broadcastMode txtypes.BroadcastMode) (*txtypes.BroadcastTxResponse, error) {
	return &txtypes.BroadcastTxResponse{}, nil
}

func (c *MockChainClientV2) QueueBroadcastMsg(msgs ...sdk.Msg) error {
	return nil
}

func (c *MockChainClientV2) GetBankBalances(ctx context.Context, address string) (*banktypes.QueryAllBalancesResponse, error) {
	return &banktypes.QueryAllBalancesResponse{}, nil
}

func (c *MockChainClientV2) GetBankBalance(ctx context.Context, address, denom string) (*banktypes.QueryBalanceResponse, error) {
	return &banktypes.QueryBalanceResponse{}, nil
}

func (c *MockChainClientV2) GetBankSpendableBalances(ctx context.Context, address string, pagination *query.PageRequest) (*banktypes.QuerySpendableBalancesResponse, error) {
	return &banktypes.QuerySpendableBalancesResponse{}, nil
}

func (c *MockChainClientV2) GetBankSpendableBalancesByDenom(ctx context.Context, address, denom string) (*banktypes.QuerySpendableBalanceByDenomResponse, error) {
	return &banktypes.QuerySpendableBalanceByDenomResponse{}, nil
}

func (c *MockChainClientV2) GetBankTotalSupply(ctx context.Context, pagination *query.PageRequest) (*banktypes.QueryTotalSupplyResponse, error) {
	return &banktypes.QueryTotalSupplyResponse{}, nil
}

func (c *MockChainClientV2) GetBankSupplyOf(ctx context.Context, denom string) (*banktypes.QuerySupplyOfResponse, error) {
	return &banktypes.QuerySupplyOfResponse{}, nil
}

func (c *MockChainClientV2) GetDenomMetadata(ctx context.Context, denom string) (*banktypes.QueryDenomMetadataResponse, error) {
	return &banktypes.QueryDenomMetadataResponse{}, nil
}

func (c *MockChainClientV2) GetDenomsMetadata(ctx context.Context, pagination *query.PageRequest) (*banktypes.QueryDenomsMetadataResponse, error) {
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

func (c *MockChainClientV2) GetDenomOwners(ctx context.Context, denom string, pagination *query.PageRequest) (*banktypes.QueryDenomOwnersResponse, error) {
	return &banktypes.QueryDenomOwnersResponse{}, nil
}

func (c *MockChainClientV2) GetBankSendEnabled(ctx context.Context, denoms []string, pagination *query.PageRequest) (*banktypes.QuerySendEnabledResponse, error) {
	return &banktypes.QuerySendEnabledResponse{}, nil
}

func (c *MockChainClientV2) GetAuthzGrants(ctx context.Context, req authztypes.QueryGrantsRequest) (*authztypes.QueryGrantsResponse, error) {
	return &authztypes.QueryGrantsResponse{}, nil
}

func (c *MockChainClientV2) GetAccount(ctx context.Context, address string) (*authtypes.QueryAccountResponse, error) {
	return &authtypes.QueryAccountResponse{}, nil
}

func (c *MockChainClientV2) BuildGenericAuthz(granter, grantee, msgtype string, expireIn time.Time) *authztypes.MsgGrant {
	return &authztypes.MsgGrant{}
}

func (c *MockChainClientV2) BuildExchangeAuthz(granter, grantee string, authzType ExchangeAuthz, subaccountId string, markets []string, expireIn time.Time) *authztypes.MsgGrant {
	return &authztypes.MsgGrant{}
}

func (c *MockChainClientV2) BuildExchangeBatchUpdateOrdersAuthz(
	granter string,
	grantee string,
	subaccountId string,
	spotMarkets []string,
	derivativeMarkets []string,
	expireIn time.Time,
) *authztypes.MsgGrant {
	return &authztypes.MsgGrant{}
}

func (c *MockChainClientV2) DefaultSubaccount(acc sdk.AccAddress) ethcommon.Hash {
	return ethcommon.HexToHash("")
}

func (c *MockChainClientV2) Subaccount(account sdk.AccAddress, index int) ethcommon.Hash {
	return ethcommon.HexToHash("")
}

func (c *MockChainClientV2) GetSubAccountNonce(ctx context.Context, subaccountId ethcommon.Hash) (*exchangev2types.QuerySubaccountTradeNonceResponse, error) {
	return &exchangev2types.QuerySubaccountTradeNonceResponse{}, nil
}

func (c *MockChainClientV2) GetFeeDiscountInfo(ctx context.Context, account string) (*exchangev2types.QueryFeeDiscountAccountInfoResponse, error) {
	return &exchangev2types.QueryFeeDiscountAccountInfoResponse{}, nil
}

func (c *MockChainClientV2) UpdateSubaccountNonceFromChain() error {
	return nil
}

func (c *MockChainClientV2) SynchronizeSubaccountNonce(subaccountId ethcommon.Hash) error {
	return nil
}

func (c *MockChainClientV2) ComputeOrderHashes(spotOrders []exchangev2types.SpotOrder, derivativeOrders []exchangev2types.DerivativeOrder, subaccountId ethcommon.Hash) (OrderHashes, error) {
	return OrderHashes{}, nil
}

func (c *MockChainClientV2) CreateSpotOrder(defaultSubaccountID ethcommon.Hash, d *SpotOrderData, marketAssistant MarketsAssistant) *exchangetypes.SpotOrder {
	return &exchangetypes.SpotOrder{}
}

func (c *MockChainClientV2) CreateDerivativeOrder(defaultSubaccountID ethcommon.Hash, d *DerivativeOrderData, marketAssistant MarketsAssistant) *exchangetypes.DerivativeOrder {
	return &exchangetypes.DerivativeOrder{}
}

func (c *MockChainClientV2) OrderCancel(defaultSubaccountID ethcommon.Hash, d *OrderCancelData) *exchangetypes.OrderData {
	return &exchangetypes.OrderData{}
}

func (c *MockChainClientV2) CreateSpotOrderV2(defaultSubaccountID ethcommon.Hash, d *SpotOrderData) *exchangev2types.SpotOrder {
	return &exchangev2types.SpotOrder{}
}

func (c *MockChainClientV2) CreateDerivativeOrderV2(defaultSubaccountID ethcommon.Hash, d *DerivativeOrderData) *exchangev2types.DerivativeOrder {
	return &exchangev2types.DerivativeOrder{}
}

func (c *MockChainClientV2) OrderCancelV2(defaultSubaccountID ethcommon.Hash, d *OrderCancelData) *exchangev2types.OrderData {
	return &exchangev2types.OrderData{}
}

func (c *MockChainClientV2) StreamEventOrderFail(sender string, failEventCh chan map[string]uint) {}

func (c *MockChainClientV2) StreamEventOrderFailWithWebsocket(sender string, websocket *rpchttp.HTTP, failEventCh chan map[string]uint) {
}

func (c *MockChainClientV2) StreamOrderbookUpdateEvents(orderbookType OrderbookType, marketIDs []string, orderbookCh chan exchangetypes.Orderbook) {
}

func (c *MockChainClientV2) StreamOrderbookUpdateEventsWithWebsocket(orderbookType OrderbookType, marketIDs []string, websocket *rpchttp.HTTP, orderbookCh chan exchangetypes.Orderbook) {
}

func (c *MockChainClientV2) ChainStreamV2(ctx context.Context, req chainstreamv2types.StreamRequest) (chainstreamv2types.Stream_StreamV2Client, error) {
	return nil, nil
}

func (c *MockChainClientV2) GetTx(ctx context.Context, txHash string) (*txtypes.GetTxResponse, error) {
	return &txtypes.GetTxResponse{}, nil
}

func (c *MockChainClientV2) Close() {}

func (c *MockChainClientV2) GetGasFee() (string, error) {
	return "", nil
}

func (c *MockChainClientV2) FetchContractInfo(ctx context.Context, address string) (*wasmtypes.QueryContractInfoResponse, error) {
	return &wasmtypes.QueryContractInfoResponse{}, nil
}

func (c *MockChainClientV2) FetchContractHistory(ctx context.Context, address string, pagination *query.PageRequest) (*wasmtypes.QueryContractHistoryResponse, error) {
	return &wasmtypes.QueryContractHistoryResponse{}, nil
}

func (c *MockChainClientV2) FetchContractsByCode(ctx context.Context, codeId uint64, pagination *query.PageRequest) (*wasmtypes.QueryContractsByCodeResponse, error) {
	return &wasmtypes.QueryContractsByCodeResponse{}, nil
}

func (c *MockChainClientV2) FetchAllContractsState(ctx context.Context, address string, pagination *query.PageRequest) (*wasmtypes.QueryAllContractStateResponse, error) {
	return &wasmtypes.QueryAllContractStateResponse{}, nil
}

func (c *MockChainClientV2) SmartContractState(
	ctx context.Context,
	contractAddress string,
	queryData []byte,
) (*wasmtypes.QuerySmartContractStateResponse, error) {
	return &wasmtypes.QuerySmartContractStateResponse{}, nil
}

func (c *MockChainClientV2) RawContractState(
	ctx context.Context,
	contractAddress string,
	queryData []byte,
) (*wasmtypes.QueryRawContractStateResponse, error) {
	return &wasmtypes.QueryRawContractStateResponse{}, nil
}

func (c *MockChainClientV2) FetchCode(ctx context.Context, codeId uint64) (*wasmtypes.QueryCodeResponse, error) {
	return &wasmtypes.QueryCodeResponse{}, nil
}

func (c *MockChainClientV2) FetchCodes(ctx context.Context, pagination *query.PageRequest) (*wasmtypes.QueryCodesResponse, error) {
	return &wasmtypes.QueryCodesResponse{}, nil
}

func (c *MockChainClientV2) FetchPinnedCodes(ctx context.Context, pagination *query.PageRequest) (*wasmtypes.QueryPinnedCodesResponse, error) {
	return &wasmtypes.QueryPinnedCodesResponse{}, nil
}

func (c *MockChainClientV2) FetchContractsByCreator(ctx context.Context, creator string, pagination *query.PageRequest) (*wasmtypes.QueryContractsByCreatorResponse, error) {
	return &wasmtypes.QueryContractsByCreatorResponse{}, nil
}

func (c *MockChainClientV2) FetchDenomAuthorityMetadata(ctx context.Context, creator, subDenom string) (*tokenfactorytypes.QueryDenomAuthorityMetadataResponse, error) {
	return &tokenfactorytypes.QueryDenomAuthorityMetadataResponse{}, nil
}

func (c *MockChainClientV2) FetchDenomsFromCreator(ctx context.Context, creator string) (*tokenfactorytypes.QueryDenomsFromCreatorResponse, error) {
	return &tokenfactorytypes.QueryDenomsFromCreatorResponse{}, nil
}

func (c *MockChainClientV2) FetchTokenfactoryModuleState(ctx context.Context) (*tokenfactorytypes.QueryModuleStateResponse, error) {
	return &tokenfactorytypes.QueryModuleStateResponse{}, nil
}

// Distribution module
func (c *MockChainClientV2) FetchValidatorDistributionInfo(ctx context.Context, validatorAddress string) (*distributiontypes.QueryValidatorDistributionInfoResponse, error) {
	return &distributiontypes.QueryValidatorDistributionInfoResponse{}, nil
}

func (c *MockChainClientV2) FetchValidatorOutstandingRewards(ctx context.Context, validatorAddress string) (*distributiontypes.QueryValidatorOutstandingRewardsResponse, error) {
	return &distributiontypes.QueryValidatorOutstandingRewardsResponse{}, nil
}

func (c *MockChainClientV2) FetchValidatorCommission(ctx context.Context, validatorAddress string) (*distributiontypes.QueryValidatorCommissionResponse, error) {
	return &distributiontypes.QueryValidatorCommissionResponse{}, nil
}

func (c *MockChainClientV2) FetchValidatorSlashes(ctx context.Context, validatorAddress string, startingHeight, endingHeight uint64, pagination *query.PageRequest) (*distributiontypes.QueryValidatorSlashesResponse, error) {
	return &distributiontypes.QueryValidatorSlashesResponse{}, nil
}

func (c *MockChainClientV2) FetchDelegationRewards(ctx context.Context, delegatorAddress, validatorAddress string) (*distributiontypes.QueryDelegationRewardsResponse, error) {
	return &distributiontypes.QueryDelegationRewardsResponse{}, nil
}

func (c *MockChainClientV2) FetchDelegationTotalRewards(ctx context.Context, delegatorAddress string) (*distributiontypes.QueryDelegationTotalRewardsResponse, error) {
	return &distributiontypes.QueryDelegationTotalRewardsResponse{}, nil
}

func (c *MockChainClientV2) FetchDelegatorValidators(ctx context.Context, delegatorAddress string) (*distributiontypes.QueryDelegatorValidatorsResponse, error) {
	return &distributiontypes.QueryDelegatorValidatorsResponse{}, nil
}

func (c *MockChainClientV2) FetchDelegatorWithdrawAddress(ctx context.Context, delegatorAddress string) (*distributiontypes.QueryDelegatorWithdrawAddressResponse, error) {
	return &distributiontypes.QueryDelegatorWithdrawAddressResponse{}, nil
}

func (c *MockChainClientV2) FetchCommunityPool(ctx context.Context) (*distributiontypes.QueryCommunityPoolResponse, error) {
	return &distributiontypes.QueryCommunityPoolResponse{}, nil
}

// Chain exchange V2 module
func (c *MockChainClientV2) FetchSubaccountDeposits(ctx context.Context, subaccountId string) (*exchangev2types.QuerySubaccountDepositsResponse, error) {
	return &exchangev2types.QuerySubaccountDepositsResponse{}, nil
}

func (c *MockChainClientV2) FetchSubaccountDeposit(ctx context.Context, subaccountId, denom string) (*exchangev2types.QuerySubaccountDepositResponse, error) {
	return &exchangev2types.QuerySubaccountDepositResponse{}, nil
}

func (c *MockChainClientV2) FetchExchangeBalances(_ context.Context) (*exchangev2types.QueryExchangeBalancesResponse, error) {
	return &exchangev2types.QueryExchangeBalancesResponse{}, nil
}

func (c *MockChainClientV2) FetchAggregateVolume(ctx context.Context, account string) (*exchangev2types.QueryAggregateVolumeResponse, error) {
	return &exchangev2types.QueryAggregateVolumeResponse{}, nil
}

func (c *MockChainClientV2) FetchAggregateVolumes(ctx context.Context, accounts, marketIDs []string) (*exchangev2types.QueryAggregateVolumesResponse, error) {
	return &exchangev2types.QueryAggregateVolumesResponse{}, nil
}

func (c *MockChainClientV2) FetchAggregateMarketVolume(ctx context.Context, marketId string) (*exchangev2types.QueryAggregateMarketVolumeResponse, error) {
	return &exchangev2types.QueryAggregateMarketVolumeResponse{}, nil
}

func (c *MockChainClientV2) FetchAggregateMarketVolumes(ctx context.Context, marketIDs []string) (*exchangev2types.QueryAggregateMarketVolumesResponse, error) {
	return &exchangev2types.QueryAggregateMarketVolumesResponse{}, nil
}

func (c *MockChainClientV2) FetchDenomDecimal(ctx context.Context, denom string) (*exchangev2types.QueryDenomDecimalResponse, error) {
	return &exchangev2types.QueryDenomDecimalResponse{}, nil
}

func (c *MockChainClientV2) FetchDenomDecimals(ctx context.Context, denoms []string) (*exchangev2types.QueryDenomDecimalsResponse, error) {
	return &exchangev2types.QueryDenomDecimalsResponse{}, nil
}

func (c *MockChainClientV2) FetchChainSpotMarkets(ctx context.Context, status string, marketIDs []string) (*exchangev2types.QuerySpotMarketsResponse, error) {
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

func (c *MockChainClientV2) FetchChainSpotMarket(ctx context.Context, marketId string) (*exchangev2types.QuerySpotMarketResponse, error) {
	return &exchangev2types.QuerySpotMarketResponse{}, nil
}

func (c *MockChainClientV2) FetchChainFullSpotMarkets(ctx context.Context, status string, marketIDs []string, withMidPriceAndTob bool) (*exchangev2types.QueryFullSpotMarketsResponse, error) {
	return &exchangev2types.QueryFullSpotMarketsResponse{}, nil
}

func (c *MockChainClientV2) FetchChainFullSpotMarket(ctx context.Context, marketId string, withMidPriceAndTob bool) (*exchangev2types.QueryFullSpotMarketResponse, error) {
	return &exchangev2types.QueryFullSpotMarketResponse{}, nil
}

func (c *MockChainClientV2) FetchChainSpotOrderbook(ctx context.Context, marketId string, limit uint64, orderSide exchangev2types.OrderSide, limitCumulativeNotional, limitCumulativeQuantity sdkmath.LegacyDec) (*exchangev2types.QuerySpotOrderbookResponse, error) {
	return &exchangev2types.QuerySpotOrderbookResponse{}, nil
}

func (c *MockChainClientV2) FetchChainTraderSpotOrders(ctx context.Context, marketId, subaccountId string) (*exchangev2types.QueryTraderSpotOrdersResponse, error) {
	return &exchangev2types.QueryTraderSpotOrdersResponse{}, nil
}

func (c *MockChainClientV2) FetchChainAccountAddressSpotOrders(ctx context.Context, marketId, address string) (*exchangev2types.QueryAccountAddressSpotOrdersResponse, error) {
	return &exchangev2types.QueryAccountAddressSpotOrdersResponse{}, nil
}

func (c *MockChainClientV2) FetchChainSpotOrdersByHashes(ctx context.Context, marketId, subaccountId string, orderHashes []string) (*exchangev2types.QuerySpotOrdersByHashesResponse, error) {
	return &exchangev2types.QuerySpotOrdersByHashesResponse{}, nil
}

func (c *MockChainClientV2) FetchChainSubaccountOrders(ctx context.Context, subaccountId, marketId string) (*exchangev2types.QuerySubaccountOrdersResponse, error) {
	return &exchangev2types.QuerySubaccountOrdersResponse{}, nil
}

func (c *MockChainClientV2) FetchChainTraderSpotTransientOrders(ctx context.Context, marketId, subaccountId string) (*exchangev2types.QueryTraderSpotOrdersResponse, error) {
	return &exchangev2types.QueryTraderSpotOrdersResponse{}, nil
}

func (c *MockChainClientV2) FetchSpotMidPriceAndTOB(ctx context.Context, marketId string) (*exchangev2types.QuerySpotMidPriceAndTOBResponse, error) {
	return &exchangev2types.QuerySpotMidPriceAndTOBResponse{}, nil
}

func (c *MockChainClientV2) FetchDerivativeMidPriceAndTOB(ctx context.Context, marketId string) (*exchangev2types.QueryDerivativeMidPriceAndTOBResponse, error) {
	return &exchangev2types.QueryDerivativeMidPriceAndTOBResponse{}, nil
}

func (c *MockChainClientV2) FetchChainDerivativeOrderbook(ctx context.Context, marketId string, limit uint64, limitCumulativeNotional sdkmath.LegacyDec) (*exchangev2types.QueryDerivativeOrderbookResponse, error) {
	return &exchangev2types.QueryDerivativeOrderbookResponse{}, nil
}

func (c *MockChainClientV2) FetchChainTraderDerivativeOrders(ctx context.Context, marketId, subaccountId string) (*exchangev2types.QueryTraderDerivativeOrdersResponse, error) {
	return &exchangev2types.QueryTraderDerivativeOrdersResponse{}, nil
}

func (c *MockChainClientV2) FetchChainAccountAddressDerivativeOrders(ctx context.Context, marketId, address string) (*exchangev2types.QueryAccountAddressDerivativeOrdersResponse, error) {
	return &exchangev2types.QueryAccountAddressDerivativeOrdersResponse{}, nil
}

func (c *MockChainClientV2) FetchChainDerivativeOrdersByHashes(ctx context.Context, marketId, subaccountId string, orderHashes []string) (*exchangev2types.QueryDerivativeOrdersByHashesResponse, error) {
	return &exchangev2types.QueryDerivativeOrdersByHashesResponse{}, nil
}

func (c *MockChainClientV2) FetchChainTraderDerivativeTransientOrders(ctx context.Context, marketId, subaccountId string) (*exchangev2types.QueryTraderDerivativeOrdersResponse, error) {
	return &exchangev2types.QueryTraderDerivativeOrdersResponse{}, nil
}

func (c *MockChainClientV2) FetchChainDerivativeMarkets(ctx context.Context, status string, marketIDs []string, withMidPriceAndTob bool) (*exchangev2types.QueryDerivativeMarketsResponse, error) {
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

func (c *MockChainClientV2) FetchChainDerivativeMarket(ctx context.Context, marketId string) (*exchangev2types.QueryDerivativeMarketResponse, error) {
	return &exchangev2types.QueryDerivativeMarketResponse{}, nil
}

func (c *MockChainClientV2) FetchDerivativeMarketAddress(ctx context.Context, marketId string) (*exchangev2types.QueryDerivativeMarketAddressResponse, error) {
	return &exchangev2types.QueryDerivativeMarketAddressResponse{}, nil
}

func (c *MockChainClientV2) FetchSubaccountTradeNonce(ctx context.Context, subaccountId string) (*exchangev2types.QuerySubaccountTradeNonceResponse, error) {
	return &exchangev2types.QuerySubaccountTradeNonceResponse{}, nil
}

func (c *MockChainClientV2) FetchChainPositions(ctx context.Context) (*exchangev2types.QueryPositionsResponse, error) {
	return &exchangev2types.QueryPositionsResponse{}, nil
}

func (c *MockChainClientV2) FetchChainPositionsInMarket(ctx context.Context, marketId string) (*exchangev2types.QueryPositionsInMarketResponse, error) {
	return &exchangev2types.QueryPositionsInMarketResponse{}, nil
}

func (c *MockChainClientV2) FetchChainSubaccountPositions(ctx context.Context, subaccountId string) (*exchangev2types.QuerySubaccountPositionsResponse, error) {
	return &exchangev2types.QuerySubaccountPositionsResponse{}, nil
}

func (c *MockChainClientV2) FetchChainSubaccountPositionInMarket(ctx context.Context, subaccountId, marketId string) (*exchangev2types.QuerySubaccountPositionInMarketResponse, error) {
	return &exchangev2types.QuerySubaccountPositionInMarketResponse{}, nil
}

func (c *MockChainClientV2) FetchChainSubaccountEffectivePositionInMarket(ctx context.Context, subaccountId, marketId string) (*exchangev2types.QuerySubaccountEffectivePositionInMarketResponse, error) {
	return &exchangev2types.QuerySubaccountEffectivePositionInMarketResponse{}, nil
}

func (c *MockChainClientV2) FetchChainPerpetualMarketInfo(ctx context.Context, marketId string) (*exchangev2types.QueryPerpetualMarketInfoResponse, error) {
	return &exchangev2types.QueryPerpetualMarketInfoResponse{}, nil
}

func (c *MockChainClientV2) FetchChainExpiryFuturesMarketInfo(ctx context.Context, marketId string) (*exchangev2types.QueryExpiryFuturesMarketInfoResponse, error) {
	return &exchangev2types.QueryExpiryFuturesMarketInfoResponse{}, nil
}

func (c *MockChainClientV2) FetchChainPerpetualMarketFunding(ctx context.Context, marketId string) (*exchangev2types.QueryPerpetualMarketFundingResponse, error) {
	return &exchangev2types.QueryPerpetualMarketFundingResponse{}, nil
}

func (c *MockChainClientV2) FetchSubaccountOrderMetadata(ctx context.Context, subaccountId string) (*exchangev2types.QuerySubaccountOrderMetadataResponse, error) {
	return &exchangev2types.QuerySubaccountOrderMetadataResponse{}, nil
}

func (c *MockChainClientV2) FetchTradeRewardPoints(ctx context.Context, accounts []string) (*exchangev2types.QueryTradeRewardPointsResponse, error) {
	return &exchangev2types.QueryTradeRewardPointsResponse{}, nil
}

func (c *MockChainClientV2) FetchPendingTradeRewardPoints(ctx context.Context, accounts []string) (*exchangev2types.QueryTradeRewardPointsResponse, error) {
	return &exchangev2types.QueryTradeRewardPointsResponse{}, nil
}

func (c *MockChainClientV2) FetchTradeRewardCampaign(ctx context.Context) (*exchangev2types.QueryTradeRewardCampaignResponse, error) {
	return &exchangev2types.QueryTradeRewardCampaignResponse{}, nil
}

func (c *MockChainClientV2) FetchFeeDiscountAccountInfo(ctx context.Context, account string) (*exchangev2types.QueryFeeDiscountAccountInfoResponse, error) {
	return &exchangev2types.QueryFeeDiscountAccountInfoResponse{}, nil
}

func (c *MockChainClientV2) FetchFeeDiscountSchedule(ctx context.Context) (*exchangev2types.QueryFeeDiscountScheduleResponse, error) {
	return &exchangev2types.QueryFeeDiscountScheduleResponse{}, nil
}

func (c *MockChainClientV2) FetchBalanceMismatches(ctx context.Context, dustFactor int64) (*exchangev2types.QueryBalanceMismatchesResponse, error) {
	return &exchangev2types.QueryBalanceMismatchesResponse{}, nil
}

func (c *MockChainClientV2) FetchBalanceWithBalanceHolds(ctx context.Context) (*exchangev2types.QueryBalanceWithBalanceHoldsResponse, error) {
	return &exchangev2types.QueryBalanceWithBalanceHoldsResponse{}, nil
}

func (c *MockChainClientV2) FetchFeeDiscountTierStatistics(ctx context.Context) (*exchangev2types.QueryFeeDiscountTierStatisticsResponse, error) {
	return &exchangev2types.QueryFeeDiscountTierStatisticsResponse{}, nil
}

func (c *MockChainClientV2) FetchMitoVaultInfos(ctx context.Context) (*exchangev2types.MitoVaultInfosResponse, error) {
	return &exchangev2types.MitoVaultInfosResponse{}, nil
}

func (c *MockChainClientV2) FetchMarketIDFromVault(ctx context.Context, vaultAddress string) (*exchangev2types.QueryMarketIDFromVaultResponse, error) {
	return &exchangev2types.QueryMarketIDFromVaultResponse{}, nil
}

func (c *MockChainClientV2) FetchHistoricalTradeRecords(ctx context.Context, marketId string) (*exchangev2types.QueryHistoricalTradeRecordsResponse, error) {
	return &exchangev2types.QueryHistoricalTradeRecordsResponse{}, nil
}

func (c *MockChainClientV2) FetchIsOptedOutOfRewards(ctx context.Context, account string) (*exchangev2types.QueryIsOptedOutOfRewardsResponse, error) {
	return &exchangev2types.QueryIsOptedOutOfRewardsResponse{}, nil
}

func (c *MockChainClientV2) FetchOptedOutOfRewardsAccounts(ctx context.Context) (*exchangev2types.QueryOptedOutOfRewardsAccountsResponse, error) {
	return &exchangev2types.QueryOptedOutOfRewardsAccountsResponse{}, nil
}

func (c *MockChainClientV2) FetchMarketVolatility(ctx context.Context, marketId string, tradeHistoryOptions *exchangev2types.TradeHistoryOptions) (*exchangev2types.QueryMarketVolatilityResponse, error) {
	return &exchangev2types.QueryMarketVolatilityResponse{}, nil
}

func (c *MockChainClientV2) FetchChainBinaryOptionsMarkets(ctx context.Context, status string) (*exchangev2types.QueryBinaryMarketsResponse, error) {
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

func (c *MockChainClientV2) FetchTraderDerivativeConditionalOrders(ctx context.Context, subaccountId, marketId string) (*exchangev2types.QueryTraderDerivativeConditionalOrdersResponse, error) {
	return &exchangev2types.QueryTraderDerivativeConditionalOrdersResponse{}, nil
}

func (c *MockChainClientV2) FetchMarketAtomicExecutionFeeMultiplier(ctx context.Context, marketId string) (*exchangev2types.QueryMarketAtomicExecutionFeeMultiplierResponse, error) {
	return &exchangev2types.QueryMarketAtomicExecutionFeeMultiplierResponse{}, nil
}

func (c *MockChainClientV2) FetchActiveStakeGrant(ctx context.Context, grantee string) (*exchangev2types.QueryActiveStakeGrantResponse, error) {
	return &exchangev2types.QueryActiveStakeGrantResponse{}, nil
}

func (c *MockChainClientV2) FetchGrantAuthorization(ctx context.Context, granter, grantee string) (*exchangev2types.QueryGrantAuthorizationResponse, error) {
	return &exchangev2types.QueryGrantAuthorizationResponse{}, nil
}

func (c *MockChainClientV2) FetchGrantAuthorizations(ctx context.Context, granter string) (*exchangev2types.QueryGrantAuthorizationsResponse, error) {
	return &exchangev2types.QueryGrantAuthorizationsResponse{}, nil
}

func (c *MockChainClientV2) FetchL3DerivativeOrderbook(ctx context.Context, marketId string) (*exchangev2types.QueryFullDerivativeOrderbookResponse, error) {
	return &exchangev2types.QueryFullDerivativeOrderbookResponse{}, nil
}

func (c *MockChainClientV2) FetchL3SpotOrderbook(ctx context.Context, marketId string) (*exchangev2types.QueryFullSpotOrderbookResponse, error) {
	return &exchangev2types.QueryFullSpotOrderbookResponse{}, nil
}

func (c *MockChainClientV2) FetchMarketBalance(ctx context.Context, marketId string) (*exchangev2types.QueryMarketBalanceResponse, error) {
	return &exchangev2types.QueryMarketBalanceResponse{}, nil
}

func (c *MockChainClientV2) FetchMarketBalances(ctx context.Context) (*exchangev2types.QueryMarketBalancesResponse, error) {
	return &exchangev2types.QueryMarketBalancesResponse{}, nil
}

func (c *MockChainClientV2) FetchDenomMinNotional(ctx context.Context, denom string) (*exchangev2types.QueryDenomMinNotionalResponse, error) {
	return &exchangev2types.QueryDenomMinNotionalResponse{}, nil
}

func (c *MockChainClientV2) FetchDenomMinNotionals(ctx context.Context) (*exchangev2types.QueryDenomMinNotionalsResponse, error) {
	return &exchangev2types.QueryDenomMinNotionalsResponse{}, nil
}

// Tendermint module

func (c *MockChainClientV2) FetchNodeInfo(ctx context.Context) (*cmtservice.GetNodeInfoResponse, error) {
	return &cmtservice.GetNodeInfoResponse{}, nil
}

func (c *MockChainClientV2) FetchSyncing(ctx context.Context) (*cmtservice.GetSyncingResponse, error) {
	return &cmtservice.GetSyncingResponse{}, nil
}

func (c *MockChainClientV2) FetchLatestBlock(ctx context.Context) (*cmtservice.GetLatestBlockResponse, error) {
	return &cmtservice.GetLatestBlockResponse{}, nil
}

func (c *MockChainClientV2) FetchBlockByHeight(ctx context.Context, height int64) (*cmtservice.GetBlockByHeightResponse, error) {
	return &cmtservice.GetBlockByHeightResponse{}, nil
}

func (c *MockChainClientV2) FetchLatestValidatorSet(ctx context.Context) (*cmtservice.GetLatestValidatorSetResponse, error) {
	return &cmtservice.GetLatestValidatorSetResponse{}, nil
}

func (c *MockChainClientV2) FetchValidatorSetByHeight(ctx context.Context, height int64, pagination *query.PageRequest) (*cmtservice.GetValidatorSetByHeightResponse, error) {
	return &cmtservice.GetValidatorSetByHeightResponse{}, nil
}

func (c *MockChainClientV2) ABCIQuery(ctx context.Context, path string, data []byte, height int64, prove bool) (*cmtservice.ABCIQueryResponse, error) {
	return &cmtservice.ABCIQueryResponse{}, nil
}

// IBC Transfer module
func (c *MockChainClientV2) FetchDenomTrace(ctx context.Context, hash string) (*ibctransfertypes.QueryDenomTraceResponse, error) {
	return &ibctransfertypes.QueryDenomTraceResponse{}, nil
}

func (c *MockChainClientV2) FetchDenomTraces(ctx context.Context, pagination *query.PageRequest) (*ibctransfertypes.QueryDenomTracesResponse, error) {
	return &ibctransfertypes.QueryDenomTracesResponse{}, nil
}

func (c *MockChainClientV2) FetchDenomHash(ctx context.Context, trace string) (*ibctransfertypes.QueryDenomHashResponse, error) {
	return &ibctransfertypes.QueryDenomHashResponse{}, nil
}

func (c *MockChainClientV2) FetchEscrowAddress(ctx context.Context, portId, channelId string) (*ibctransfertypes.QueryEscrowAddressResponse, error) {
	return &ibctransfertypes.QueryEscrowAddressResponse{}, nil
}

func (c *MockChainClientV2) FetchTotalEscrowForDenom(ctx context.Context, denom string) (*ibctransfertypes.QueryTotalEscrowForDenomResponse, error) {
	return &ibctransfertypes.QueryTotalEscrowForDenomResponse{}, nil
}

// IBC Core Channel module
func (c *MockChainClientV2) FetchIBCChannel(ctx context.Context, portId, channelId string) (*ibcchanneltypes.QueryChannelResponse, error) {
	return &ibcchanneltypes.QueryChannelResponse{}, nil
}

func (c *MockChainClientV2) FetchIBCChannels(ctx context.Context, pagination *query.PageRequest) (*ibcchanneltypes.QueryChannelsResponse, error) {
	return &ibcchanneltypes.QueryChannelsResponse{}, nil
}

func (c *MockChainClientV2) FetchIBCConnectionChannels(ctx context.Context, connection string, pagination *query.PageRequest) (*ibcchanneltypes.QueryConnectionChannelsResponse, error) {
	return &ibcchanneltypes.QueryConnectionChannelsResponse{}, nil
}

func (c *MockChainClientV2) FetchIBCChannelClientState(ctx context.Context, portId, channelId string) (*ibcchanneltypes.QueryChannelClientStateResponse, error) {
	return &ibcchanneltypes.QueryChannelClientStateResponse{}, nil
}

func (c *MockChainClientV2) FetchIBCChannelConsensusState(ctx context.Context, portId, channelId string, revisionNumber, revisionHeight uint64) (*ibcchanneltypes.QueryChannelConsensusStateResponse, error) {
	return &ibcchanneltypes.QueryChannelConsensusStateResponse{}, nil
}

func (c *MockChainClientV2) FetchIBCPacketCommitment(ctx context.Context, portId, channelId string, sequence uint64) (*ibcchanneltypes.QueryPacketCommitmentResponse, error) {
	return &ibcchanneltypes.QueryPacketCommitmentResponse{}, nil
}

func (c *MockChainClientV2) FetchIBCPacketCommitments(ctx context.Context, portId, channelId string, pagination *query.PageRequest) (*ibcchanneltypes.QueryPacketCommitmentsResponse, error) {
	return &ibcchanneltypes.QueryPacketCommitmentsResponse{}, nil
}

func (c *MockChainClientV2) FetchIBCPacketReceipt(ctx context.Context, portId, channelId string, sequence uint64) (*ibcchanneltypes.QueryPacketReceiptResponse, error) {
	return &ibcchanneltypes.QueryPacketReceiptResponse{}, nil
}

func (c *MockChainClientV2) FetchIBCPacketAcknowledgement(ctx context.Context, portId, channelId string, sequence uint64) (*ibcchanneltypes.QueryPacketAcknowledgementResponse, error) {
	return &ibcchanneltypes.QueryPacketAcknowledgementResponse{}, nil
}

func (c *MockChainClientV2) FetchIBCPacketAcknowledgements(ctx context.Context, portId, channelId string, packetCommitmentSequences []uint64, pagination *query.PageRequest) (*ibcchanneltypes.QueryPacketAcknowledgementsResponse, error) {
	return &ibcchanneltypes.QueryPacketAcknowledgementsResponse{}, nil
}

func (c *MockChainClientV2) FetchIBCUnreceivedPackets(ctx context.Context, portId, channelId string, packetCommitmentSequences []uint64) (*ibcchanneltypes.QueryUnreceivedPacketsResponse, error) {
	return &ibcchanneltypes.QueryUnreceivedPacketsResponse{}, nil
}

func (c *MockChainClientV2) FetchIBCUnreceivedAcks(ctx context.Context, portId, channelId string, packetAckSequences []uint64) (*ibcchanneltypes.QueryUnreceivedAcksResponse, error) {
	return &ibcchanneltypes.QueryUnreceivedAcksResponse{}, nil
}

func (c *MockChainClientV2) FetchIBCNextSequenceReceive(ctx context.Context, portId, channelId string) (*ibcchanneltypes.QueryNextSequenceReceiveResponse, error) {
	return &ibcchanneltypes.QueryNextSequenceReceiveResponse{}, nil
}

// IBC Core Chain module
func (c *MockChainClientV2) FetchIBCClientState(ctx context.Context, clientId string) (*ibcclienttypes.QueryClientStateResponse, error) {
	return &ibcclienttypes.QueryClientStateResponse{}, nil
}

func (c *MockChainClientV2) FetchIBCClientStates(ctx context.Context, pagination *query.PageRequest) (*ibcclienttypes.QueryClientStatesResponse, error) {
	return &ibcclienttypes.QueryClientStatesResponse{}, nil
}

func (c *MockChainClientV2) FetchIBCConsensusState(ctx context.Context, clientId string, revisionNumber, revisionHeight uint64, latestHeight bool) (*ibcclienttypes.QueryConsensusStateResponse, error) {
	return &ibcclienttypes.QueryConsensusStateResponse{}, nil
}

func (c *MockChainClientV2) FetchIBCConsensusStates(ctx context.Context, clientId string, pagination *query.PageRequest) (*ibcclienttypes.QueryConsensusStatesResponse, error) {
	return &ibcclienttypes.QueryConsensusStatesResponse{}, nil
}

func (c *MockChainClientV2) FetchIBCConsensusStateHeights(ctx context.Context, clientId string, pagination *query.PageRequest) (*ibcclienttypes.QueryConsensusStateHeightsResponse, error) {
	return &ibcclienttypes.QueryConsensusStateHeightsResponse{}, nil
}

func (c *MockChainClientV2) FetchIBCClientStatus(ctx context.Context, clientId string) (*ibcclienttypes.QueryClientStatusResponse, error) {
	return &ibcclienttypes.QueryClientStatusResponse{}, nil
}

func (c *MockChainClientV2) FetchIBCClientParams(ctx context.Context) (*ibcclienttypes.QueryClientParamsResponse, error) {
	return &ibcclienttypes.QueryClientParamsResponse{}, nil
}

func (c *MockChainClientV2) FetchIBCUpgradedClientState(ctx context.Context) (*ibcclienttypes.QueryUpgradedClientStateResponse, error) {
	return &ibcclienttypes.QueryUpgradedClientStateResponse{}, nil
}

func (c *MockChainClientV2) FetchIBCUpgradedConsensusState(ctx context.Context) (*ibcclienttypes.QueryUpgradedConsensusStateResponse, error) {
	return &ibcclienttypes.QueryUpgradedConsensusStateResponse{}, nil
}

// IBC Core Connection module
func (c *MockChainClientV2) FetchIBCConnection(ctx context.Context, connectionId string) (*ibcconnectiontypes.QueryConnectionResponse, error) {
	return &ibcconnectiontypes.QueryConnectionResponse{}, nil
}

func (c *MockChainClientV2) FetchIBCConnections(ctx context.Context, pagination *query.PageRequest) (*ibcconnectiontypes.QueryConnectionsResponse, error) {
	return &ibcconnectiontypes.QueryConnectionsResponse{}, nil
}

func (c *MockChainClientV2) FetchIBCClientConnections(ctx context.Context, clientId string) (*ibcconnectiontypes.QueryClientConnectionsResponse, error) {
	return &ibcconnectiontypes.QueryClientConnectionsResponse{}, nil
}

func (c *MockChainClientV2) FetchIBCConnectionClientState(ctx context.Context, connectionId string) (*ibcconnectiontypes.QueryConnectionClientStateResponse, error) {
	return &ibcconnectiontypes.QueryConnectionClientStateResponse{}, nil
}

func (c *MockChainClientV2) FetchIBCConnectionConsensusState(ctx context.Context, connectionId string, revisionNumber, revisionHeight uint64) (*ibcconnectiontypes.QueryConnectionConsensusStateResponse, error) {
	return &ibcconnectiontypes.QueryConnectionConsensusStateResponse{}, nil
}

func (c *MockChainClientV2) FetchIBCConnectionParams(ctx context.Context) (*ibcconnectiontypes.QueryConnectionParamsResponse, error) {
	return &ibcconnectiontypes.QueryConnectionParamsResponse{}, nil
}

// Permissions module

func (c *MockChainClientV2) FetchPermissionsNamespaceDenoms(ctx context.Context) (*permissionstypes.QueryNamespaceDenomsResponse, error) {
	return &permissionstypes.QueryNamespaceDenomsResponse{}, nil
}

func (c *MockChainClientV2) FetchPermissionsNamespaces(ctx context.Context) (*permissionstypes.QueryNamespacesResponse, error) {
	return &permissionstypes.QueryNamespacesResponse{}, nil
}

func (c *MockChainClientV2) FetchPermissionsNamespace(ctx context.Context, denom string) (*permissionstypes.QueryNamespaceResponse, error) {
	return &permissionstypes.QueryNamespaceResponse{}, nil
}

func (c *MockChainClientV2) FetchPermissionsRolesByActor(ctx context.Context, denom, actor string) (*permissionstypes.QueryRolesByActorResponse, error) {
	return &permissionstypes.QueryRolesByActorResponse{}, nil
}

func (c *MockChainClientV2) FetchPermissionsActorsByRole(ctx context.Context, denom, role string) (*permissionstypes.QueryActorsByRoleResponse, error) {
	return &permissionstypes.QueryActorsByRoleResponse{}, nil
}

func (c *MockChainClientV2) FetchPermissionsRoleManagers(ctx context.Context, denom string) (*permissionstypes.QueryRoleManagersResponse, error) {
	return &permissionstypes.QueryRoleManagersResponse{}, nil
}

func (c *MockChainClientV2) FetchPermissionsRoleManager(ctx context.Context, denom, manager string) (*permissionstypes.QueryRoleManagerResponse, error) {
	return &permissionstypes.QueryRoleManagerResponse{}, nil
}

func (c *MockChainClientV2) FetchPermissionsPolicyStatuses(ctx context.Context, denom string) (*permissionstypes.QueryPolicyStatusesResponse, error) {
	return &permissionstypes.QueryPolicyStatusesResponse{}, nil
}

func (c *MockChainClientV2) FetchPermissionsPolicyManagerCapabilities(ctx context.Context, denom string) (*permissionstypes.QueryPolicyManagerCapabilitiesResponse, error) {
	return &permissionstypes.QueryPolicyManagerCapabilitiesResponse{}, nil
}

func (c *MockChainClientV2) FetchPermissionsVouchers(ctx context.Context, denom string) (*permissionstypes.QueryVouchersResponse, error) {
	return &permissionstypes.QueryVouchersResponse{}, nil
}

func (c *MockChainClientV2) FetchPermissionsVoucher(ctx context.Context, denom, address string) (*permissionstypes.QueryVoucherResponse, error) {
	return &permissionstypes.QueryVoucherResponse{}, nil
}

func (c *MockChainClientV2) FetchPermissionsModuleState(ctx context.Context) (*permissionstypes.QueryModuleStateResponse, error) {
	return &permissionstypes.QueryModuleStateResponse{}, nil
}

// TxFees module
func (c *MockChainClientV2) FetchTxFeesParams(ctx context.Context) (*txfeestypes.QueryParamsResponse, error) {
	return &txfeestypes.QueryParamsResponse{}, nil
}

func (c *MockChainClientV2) FetchEipBaseFee(ctx context.Context) (*txfeestypes.QueryEipBaseFeeResponse, error) {
	return &txfeestypes.QueryEipBaseFeeResponse{}, nil
}

// ERC20 module
func (c *MockChainClientV2) FetchAllTokenPairs(ctx context.Context, pagination *query.PageRequest) (*erc20types.QueryAllTokenPairsResponse, error) {
	return &erc20types.QueryAllTokenPairsResponse{}, nil
}

func (c *MockChainClientV2) FetchTokenPairByDenom(ctx context.Context, bankDenom string) (*erc20types.QueryTokenPairByDenomResponse, error) {
	return &erc20types.QueryTokenPairByDenomResponse{}, nil
}

func (c *MockChainClientV2) FetchTokenPairByERC20Address(ctx context.Context, erc20Address string) (*erc20types.QueryTokenPairByERC20AddressResponse, error) {
	return &erc20types.QueryTokenPairByERC20AddressResponse{}, nil
}

// EVM module
func (c *MockChainClientV2) FetchEVMAccount(ctx context.Context, address string) (*evmtypes.QueryAccountResponse, error) {
	return &evmtypes.QueryAccountResponse{}, nil
}

func (c *MockChainClientV2) FetchEVMCosmosAccount(ctx context.Context, address string) (*evmtypes.QueryCosmosAccountResponse, error) {
	return &evmtypes.QueryCosmosAccountResponse{}, nil
}

func (c *MockChainClientV2) FetchEVMValidatorAccount(ctx context.Context, consAddress string) (*evmtypes.QueryValidatorAccountResponse, error) {
	return &evmtypes.QueryValidatorAccountResponse{}, nil
}

func (c *MockChainClientV2) FetchEVMBalance(ctx context.Context, address string) (*evmtypes.QueryBalanceResponse, error) {
	return &evmtypes.QueryBalanceResponse{}, nil
}

func (c *MockChainClientV2) FetchEVMStorage(ctx context.Context, address string, key *string) (*evmtypes.QueryStorageResponse, error) {
	return &evmtypes.QueryStorageResponse{}, nil
}

func (c *MockChainClientV2) FetchEVMCode(ctx context.Context, address string) (*evmtypes.QueryCodeResponse, error) {
	return &evmtypes.QueryCodeResponse{}, nil
}

func (c *MockChainClientV2) FetchEVMBaseFee(ctx context.Context) (*evmtypes.QueryBaseFeeResponse, error) {
	return &evmtypes.QueryBaseFeeResponse{}, nil
}

func (c *MockChainClientV2) CurrentChainGasPrice() int64 {
	return int64(injectiveclient.DefaultGasPrice)
}

func (c *MockChainClientV2) SetGasPrice(gasPrice int64) {
	// do nothing
}

func (c *MockChainClientV2) GetNetwork() common.Network {
	return c.Network
}
