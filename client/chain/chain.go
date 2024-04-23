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

	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"

	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"

	"github.com/cosmos/cosmos-sdk/types/query"

	"google.golang.org/grpc/credentials/insecure"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	chainstreamtypes "github.com/InjectiveLabs/sdk-go/chain/stream/types"
	tokenfactorytypes "github.com/InjectiveLabs/sdk-go/chain/tokenfactory/types"
	"github.com/InjectiveLabs/sdk-go/client/common"
	log "github.com/InjectiveLabs/suplog"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/gogoproto/proto"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	ibcchanneltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	eth "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc"

	ethcommon "github.com/ethereum/go-ethereum/common"
)

type OrderbookType string

const (
	SpotOrderbook       = "injective.exchange.v1beta1.EventOrderbookUpdate.spot_orderbooks"
	DerivativeOrderbook = "injective.exchange.v1beta1.EventOrderbookUpdate.derivative_orderbooks"
)

const (
	msgCommitBatchSizeLimit          = 1024
	msgCommitBatchTimeLimit          = 500 * time.Millisecond
	defaultBroadcastStatusPoll       = 100 * time.Millisecond
	defaultBroadcastTimeout          = 40 * time.Second
	defaultTimeoutHeight             = 20
	defaultTimeoutHeightSyncInterval = 10 * time.Second
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
	ClientContext() client.Context
	// return account number and sequence without increasing sequence
	GetAccNonce() (accNum uint64, accSeq uint64)

	SimulateMsg(clientCtx client.Context, msgs ...sdk.Msg) (*txtypes.SimulateResponse, error)
	AsyncBroadcastMsg(msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error)
	SyncBroadcastMsg(msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error)

	// Build signed tx with given accNum and accSeq, useful for offline siging
	// If simulate is set to false, initialGas will be used
	BuildSignedTx(clientCtx client.Context, accNum, accSeq, initialGas uint64, msg ...sdk.Msg) ([]byte, error)
	SyncBroadcastSignedTx(tyBytes []byte) (*txtypes.BroadcastTxResponse, error)
	AsyncBroadcastSignedTx(txBytes []byte) (*txtypes.BroadcastTxResponse, error)
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

	DefaultSubaccount(acc sdk.AccAddress) eth.Hash
	Subaccount(account sdk.AccAddress, index int) eth.Hash

	GetSubAccountNonce(ctx context.Context, subaccountId eth.Hash) (*exchangetypes.QuerySubaccountTradeNonceResponse, error)
	GetFeeDiscountInfo(ctx context.Context, account string) (*exchangetypes.QueryFeeDiscountAccountInfoResponse, error)

	UpdateSubaccountNonceFromChain() error
	SynchronizeSubaccountNonce(subaccountId eth.Hash) error
	ComputeOrderHashes(spotOrders []exchangetypes.SpotOrder, derivativeOrders []exchangetypes.DerivativeOrder, subaccountId eth.Hash) (OrderHashes, error)

	SpotOrder(defaultSubaccountID eth.Hash, network common.Network, d *SpotOrderData) *exchangetypes.SpotOrder
	CreateSpotOrder(defaultSubaccountID eth.Hash, d *SpotOrderData, marketsAssistant MarketsAssistant) *exchangetypes.SpotOrder
	DerivativeOrder(defaultSubaccountID eth.Hash, network common.Network, d *DerivativeOrderData) *exchangetypes.DerivativeOrder
	CreateDerivativeOrder(defaultSubaccountID eth.Hash, d *DerivativeOrderData, marketAssistant MarketsAssistant) *exchangetypes.DerivativeOrder
	OrderCancel(defaultSubaccountID eth.Hash, d *OrderCancelData) *exchangetypes.OrderData

	GetGasFee() (string, error)

	StreamEventOrderFail(sender string, failEventCh chan map[string]uint)
	StreamEventOrderFailWithWebsocket(sender string, websocket *rpchttp.HTTP, failEventCh chan map[string]uint)
	StreamOrderbookUpdateEvents(orderbookType OrderbookType, marketIds []string, orderbookCh chan exchangetypes.Orderbook)
	StreamOrderbookUpdateEventsWithWebsocket(orderbookType OrderbookType, marketIds []string, websocket *rpchttp.HTTP, orderbookCh chan exchangetypes.Orderbook)

	ChainStream(ctx context.Context, req chainstreamtypes.StreamRequest) (chainstreamtypes.Stream_StreamClient, error)

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
	FetchSubaccountDeposits(ctx context.Context, subaccountID string) (*exchangetypes.QuerySubaccountDepositsResponse, error)
	FetchSubaccountDeposit(ctx context.Context, subaccountId string, denom string) (*exchangetypes.QuerySubaccountDepositResponse, error)
	FetchExchangeBalances(ctx context.Context) (*exchangetypes.QueryExchangeBalancesResponse, error)
	FetchAggregateVolume(ctx context.Context, account string) (*exchangetypes.QueryAggregateVolumeResponse, error)
	FetchAggregateVolumes(ctx context.Context, accounts []string, marketIds []string) (*exchangetypes.QueryAggregateVolumesResponse, error)
	FetchAggregateMarketVolume(ctx context.Context, marketId string) (*exchangetypes.QueryAggregateMarketVolumeResponse, error)
	FetchAggregateMarketVolumes(ctx context.Context, marketIds []string) (*exchangetypes.QueryAggregateMarketVolumesResponse, error)
	FetchDenomDecimal(ctx context.Context, denom string) (*exchangetypes.QueryDenomDecimalResponse, error)
	FetchDenomDecimals(ctx context.Context, denoms []string) (*exchangetypes.QueryDenomDecimalsResponse, error)
	FetchChainSpotMarkets(ctx context.Context, status string, marketIds []string) (*exchangetypes.QuerySpotMarketsResponse, error)
	FetchChainSpotMarket(ctx context.Context, marketId string) (*exchangetypes.QuerySpotMarketResponse, error)
	FetchChainFullSpotMarkets(ctx context.Context, status string, marketIds []string, withMidPriceAndTob bool) (*exchangetypes.QueryFullSpotMarketsResponse, error)
	FetchChainFullSpotMarket(ctx context.Context, marketId string, withMidPriceAndTob bool) (*exchangetypes.QueryFullSpotMarketResponse, error)
	FetchChainSpotOrderbook(ctx context.Context, marketId string, limit uint64, orderSide exchangetypes.OrderSide, limitCumulativeNotional sdk.Dec, limitCumulativeQuantity sdk.Dec) (*exchangetypes.QuerySpotOrderbookResponse, error)
	FetchChainTraderSpotOrders(ctx context.Context, marketId string, subaccountId string) (*exchangetypes.QueryTraderSpotOrdersResponse, error)
	FetchChainAccountAddressSpotOrders(ctx context.Context, marketId string, address string) (*exchangetypes.QueryAccountAddressSpotOrdersResponse, error)
	FetchChainSpotOrdersByHashes(ctx context.Context, marketId string, subaccountId string, orderHashes []string) (*exchangetypes.QuerySpotOrdersByHashesResponse, error)
	FetchChainSubaccountOrders(ctx context.Context, subaccountId string, marketId string) (*exchangetypes.QuerySubaccountOrdersResponse, error)
	FetchChainTraderSpotTransientOrders(ctx context.Context, marketId string, subaccountId string) (*exchangetypes.QueryTraderSpotOrdersResponse, error)
	FetchSpotMidPriceAndTOB(ctx context.Context, marketId string) (*exchangetypes.QuerySpotMidPriceAndTOBResponse, error)
	FetchDerivativeMidPriceAndTOB(ctx context.Context, marketId string) (*exchangetypes.QueryDerivativeMidPriceAndTOBResponse, error)
	FetchChainDerivativeOrderbook(ctx context.Context, marketId string, limit uint64, limitCumulativeNotional sdk.Dec) (*exchangetypes.QueryDerivativeOrderbookResponse, error)
	FetchChainTraderDerivativeOrders(ctx context.Context, marketId string, subaccountId string) (*exchangetypes.QueryTraderDerivativeOrdersResponse, error)
	FetchChainAccountAddressDerivativeOrders(ctx context.Context, marketId string, address string) (*exchangetypes.QueryAccountAddressDerivativeOrdersResponse, error)
	FetchChainDerivativeOrdersByHashes(ctx context.Context, marketId string, subaccountId string, orderHashes []string) (*exchangetypes.QueryDerivativeOrdersByHashesResponse, error)
	FetchChainTraderDerivativeTransientOrders(ctx context.Context, marketId string, subaccountId string) (*exchangetypes.QueryTraderDerivativeOrdersResponse, error)
	FetchChainDerivativeMarkets(ctx context.Context, status string, marketIds []string, withMidPriceAndTob bool) (*exchangetypes.QueryDerivativeMarketsResponse, error)
	FetchChainDerivativeMarket(ctx context.Context, marketId string) (*exchangetypes.QueryDerivativeMarketResponse, error)
	FetchDerivativeMarketAddress(ctx context.Context, marketId string) (*exchangetypes.QueryDerivativeMarketAddressResponse, error)
	FetchSubaccountTradeNonce(ctx context.Context, subaccountId string) (*exchangetypes.QuerySubaccountTradeNonceResponse, error)
	FetchChainPositions(ctx context.Context) (*exchangetypes.QueryPositionsResponse, error)
	FetchChainSubaccountPositions(ctx context.Context, subaccountId string) (*exchangetypes.QuerySubaccountPositionsResponse, error)
	FetchChainSubaccountPositionInMarket(ctx context.Context, subaccountId string, marketId string) (*exchangetypes.QuerySubaccountPositionInMarketResponse, error)
	FetchChainSubaccountEffectivePositionInMarket(ctx context.Context, subaccountId string, marketId string) (*exchangetypes.QuerySubaccountEffectivePositionInMarketResponse, error)
	FetchChainPerpetualMarketInfo(ctx context.Context, marketId string) (*exchangetypes.QueryPerpetualMarketInfoResponse, error)
	FetchChainExpiryFuturesMarketInfo(ctx context.Context, marketId string) (*exchangetypes.QueryExpiryFuturesMarketInfoResponse, error)
	FetchChainPerpetualMarketFunding(ctx context.Context, marketId string) (*exchangetypes.QueryPerpetualMarketFundingResponse, error)
	FetchSubaccountOrderMetadata(ctx context.Context, subaccountId string) (*exchangetypes.QuerySubaccountOrderMetadataResponse, error)
	FetchTradeRewardPoints(ctx context.Context, accounts []string) (*exchangetypes.QueryTradeRewardPointsResponse, error)
	FetchPendingTradeRewardPoints(ctx context.Context, accounts []string) (*exchangetypes.QueryTradeRewardPointsResponse, error)
	FetchFeeDiscountAccountInfo(ctx context.Context, account string) (*exchangetypes.QueryFeeDiscountAccountInfoResponse, error)
	FetchTradeRewardCampaign(ctx context.Context) (*exchangetypes.QueryTradeRewardCampaignResponse, error)
	FetchFeeDiscountSchedule(ctx context.Context) (*exchangetypes.QueryFeeDiscountScheduleResponse, error)
	FetchBalanceMismatches(ctx context.Context, dustFactor int64) (*exchangetypes.QueryBalanceMismatchesResponse, error)
	FetchBalanceWithBalanceHolds(ctx context.Context) (*exchangetypes.QueryBalanceWithBalanceHoldsResponse, error)
	FetchFeeDiscountTierStatistics(ctx context.Context) (*exchangetypes.QueryFeeDiscountTierStatisticsResponse, error)
	FetchMitoVaultInfos(ctx context.Context) (*exchangetypes.MitoVaultInfosResponse, error)
	FetchMarketIDFromVault(ctx context.Context, vaultAddress string) (*exchangetypes.QueryMarketIDFromVaultResponse, error)
	FetchHistoricalTradeRecords(ctx context.Context, marketId string) (*exchangetypes.QueryHistoricalTradeRecordsResponse, error)
	FetchIsOptedOutOfRewards(ctx context.Context, account string) (*exchangetypes.QueryIsOptedOutOfRewardsResponse, error)
	FetchOptedOutOfRewardsAccounts(ctx context.Context) (*exchangetypes.QueryOptedOutOfRewardsAccountsResponse, error)
	FetchMarketVolatility(ctx context.Context, marketId string, tradeHistoryOptions *exchangetypes.TradeHistoryOptions) (*exchangetypes.QueryMarketVolatilityResponse, error)
	FetchChainBinaryOptionsMarkets(ctx context.Context, status string) (*exchangetypes.QueryBinaryMarketsResponse, error)
	FetchTraderDerivativeConditionalOrders(ctx context.Context, subaccountId string, marketId string) (*exchangetypes.QueryTraderDerivativeConditionalOrdersResponse, error)
	FetchMarketAtomicExecutionFeeMultiplier(ctx context.Context, marketId string) (*exchangetypes.QueryMarketAtomicExecutionFeeMultiplierResponse, error)

	// Tendermint module
	FetchNodeInfo(ctx context.Context) (*tmservice.GetNodeInfoResponse, error)
	FetchSyncing(ctx context.Context) (*tmservice.GetSyncingResponse, error)
	FetchLatestBlock(ctx context.Context) (*tmservice.GetLatestBlockResponse, error)
	FetchBlockByHeight(ctx context.Context, height int64) (*tmservice.GetBlockByHeightResponse, error)
	FetchLatestValidatorSet(ctx context.Context) (*tmservice.GetLatestValidatorSetResponse, error)
	FetchValidatorSetByHeight(ctx context.Context, height int64, pagination *query.PageRequest) (*tmservice.GetValidatorSetByHeightResponse, error)
	ABCIQuery(ctx context.Context, path string, data []byte, height int64, prove bool) (*tmservice.ABCIQueryResponse, error)

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

	Close()
}

type chainClient struct {
	ctx             client.Context
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

	authQueryClient         authtypes.QueryClient
	authzQueryClient        authztypes.QueryClient
	bankQueryClient         banktypes.QueryClient
	chainStreamClient       chainstreamtypes.StreamClient
	distributionQueryClient distributiontypes.QueryClient
	exchangeQueryClient     exchangetypes.QueryClient
	ibcChannelQueryClient   ibcchanneltypes.QueryClient
	ibcClientQueryClient    ibcclienttypes.QueryClient
	ibcTransferQueryClient  ibctransfertypes.QueryClient
	tendermintQueryClient   tmservice.ServiceClient
	tokenfactoryQueryClient tokenfactorytypes.QueryClient
	txClient                txtypes.ServiceClient
	wasmQueryClient         wasmtypes.QueryClient
	subaccountToNonce       map[ethcommon.Hash]uint32

	closed  int64
	canSign bool
}

func NewChainClient(
	ctx client.Context,
	network common.Network,
	options ...common.ClientOption,
) (ChainClient, error) {

	// process options
	opts := common.DefaultClientOptions()

	if network.ChainTlsCert != nil {
		options = append(options, common.OptionTLSCert(network.ChainTlsCert))
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
		if len(opts.GasPrices) > 0 {
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
		conn, err = grpc.Dial(network.ChainGrpcEndpoint, grpc.WithTransportCredentials(opts.TLSCert), grpc.WithContextDialer(common.DialerFunc))
	} else {
		conn, err = grpc.Dial(network.ChainGrpcEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(common.DialerFunc))
		stickySessionEnabled = false
	}
	if err != nil {
		err = errors.Wrapf(err, "failed to connect to the gRPC: %s", network.ChainGrpcEndpoint)
		return nil, err
	}

	var chainStreamConn *grpc.ClientConn
	if opts.TLSCert != nil {
		chainStreamConn, err = grpc.Dial(network.ChainStreamGrpcEndpoint, grpc.WithTransportCredentials(opts.TLSCert), grpc.WithContextDialer(common.DialerFunc))
	} else {
		chainStreamConn, err = grpc.Dial(network.ChainStreamGrpcEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(common.DialerFunc))
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

		authQueryClient:         authtypes.NewQueryClient(conn),
		authzQueryClient:        authztypes.NewQueryClient(conn),
		bankQueryClient:         banktypes.NewQueryClient(conn),
		chainStreamClient:       chainstreamtypes.NewStreamClient(chainStreamConn),
		distributionQueryClient: distributiontypes.NewQueryClient(conn),
		exchangeQueryClient:     exchangetypes.NewQueryClient(conn),
		ibcChannelQueryClient:   ibcchanneltypes.NewQueryClient(conn),
		ibcClientQueryClient:    ibcclienttypes.NewQueryClient(conn),
		ibcTransferQueryClient:  ibctransfertypes.NewQueryClient(conn),
		tendermintQueryClient:   tmservice.NewServiceClient(conn),
		tokenfactoryQueryClient: tokenfactorytypes.NewQueryClient(conn),
		txClient:                txtypes.NewServiceClient(conn),
		wasmQueryClient:         wasmtypes.NewQueryClient(conn),
		subaccountToNonce:       make(map[ethcommon.Hash]uint32),
	}

	if cc.canSign {
		var err error

		cc.accNum, cc.accSeq, err = cc.txFactory.AccountRetriever().GetAccountNumberSequence(ctx, ctx.GetFromAddress())
		if err != nil {
			err = errors.Wrap(err, "failed to get initial account num and seq")
			return nil, err
		}

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

// prepareFactory ensures the account defined by ctx.GetFromAddress() exists and
// if the account number and/or the account sequence number are zero (not set),
// they will be queried for and set on the provided Factory. A new Factory with
// the updated fields will be returned.
func (c *chainClient) prepareFactory(clientCtx client.Context, txf tx.Factory) (tx.Factory, error) {
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

func (c *chainClient) GetAccNonce() (accNum uint64, accSeq uint64) {
	return c.accNum, c.accSeq
}

func (c *chainClient) QueryClient() *grpc.ClientConn {
	return c.conn
}

func (c *chainClient) ClientContext() client.Context {
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

//Bank Module

func (c *chainClient) GetBankBalances(ctx context.Context, address string) (*banktypes.QueryAllBalancesResponse, error) {
	req := &banktypes.QueryAllBalancesRequest{
		Address: address,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.bankQueryClient.AllBalances, req)

	return res, err
}

func (c *chainClient) GetBankBalance(ctx context.Context, address string, denom string) (*banktypes.QueryBalanceResponse, error) {
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

func (c *chainClient) GetBankSpendableBalancesByDenom(ctx context.Context, address string, denom string) (*banktypes.QuerySpendableBalanceByDenomResponse, error) {
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

// SyncBroadcastMsg sends Tx to chain and waits until Tx is included in block.
func (c *chainClient) SyncBroadcastMsg(msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error) {
	c.syncMux.Lock()
	defer c.syncMux.Unlock()

	sequence := c.getAccSeq()
	c.txFactory = c.txFactory.WithSequence(sequence)
	c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
	res, err := c.broadcastTx(c.ctx, c.txFactory, true, msgs...)

	if err != nil {
		if strings.Contains(err.Error(), "account sequence mismatch") {
			c.syncNonce()
			sequence := c.getAccSeq()
			c.txFactory = c.txFactory.WithSequence(sequence)
			c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
			log.Debugln("retrying broadcastTx with nonce", sequence)
			res, err = c.broadcastTx(c.ctx, c.txFactory, true, msgs...)
		}
		if err != nil {
			resJSON, _ := json.MarshalIndent(res, "", "\t")
			c.logger.WithField("size", len(msgs)).WithError(err).Errorln("failed synchronously broadcast messages:", string(resJSON))
			return nil, err
		}
	}

	return res, nil
}

func (c *chainClient) GetFeeDiscountInfo(ctx context.Context, account string) (*exchangetypes.QueryFeeDiscountAccountInfoResponse, error) {
	req := &exchangetypes.QueryFeeDiscountAccountInfoRequest{
		Account: account,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.FeeDiscountAccountInfo, req)

	return res, err
}

func (c *chainClient) SimulateMsg(clientCtx client.Context, msgs ...sdk.Msg) (*txtypes.SimulateResponse, error) {
	c.txFactory = c.txFactory.WithSequence(c.accSeq)
	c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
	txf, err := c.prepareFactory(clientCtx, c.txFactory)
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

// AsyncBroadcastMsg sends Tx to chain and doesn't wait until Tx is included in block. This method
// cannot be used for rapid Tx sending, it is expected that you wait for transaction status with
// external tools. If you want sdk to wait for it, use SyncBroadcastMsg.
func (c *chainClient) AsyncBroadcastMsg(msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error) {
	c.syncMux.Lock()
	defer c.syncMux.Unlock()

	sequence := c.getAccSeq()
	c.txFactory = c.txFactory.WithSequence(sequence)
	c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
	res, err := c.broadcastTx(c.ctx, c.txFactory, false, msgs...)
	if err != nil {
		if strings.Contains(err.Error(), "account sequence mismatch") {
			c.syncNonce()
			sequence := c.getAccSeq()
			c.txFactory = c.txFactory.WithSequence(sequence)
			c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
			log.Debugln("retrying broadcastTx with nonce", sequence)
			res, err = c.broadcastTx(c.ctx, c.txFactory, false, msgs...)
		}
		if err != nil {
			resJSON, _ := json.MarshalIndent(res, "", "\t")
			c.logger.WithField("size", len(msgs)).WithError(err).Errorln("failed to asynchronously broadcast messagess:", string(resJSON))
			return nil, err
		}
	}

	return res, nil
}

func (c *chainClient) BuildSignedTx(clientCtx client.Context, accNum, accSeq, initialGas uint64, msgs ...sdk.Msg) ([]byte, error) {
	txf := NewTxFactory(clientCtx).WithSequence(accSeq).WithAccountNumber(accNum).WithGas(initialGas)

	if clientCtx.Simulate {
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

		adjustedGas := uint64(txf.GasAdjustment() * float64(simRes.GasInfo.GasUsed))
		txf = txf.WithGas(adjustedGas)

		c.gasWanted = adjustedGas
	}

	txf, err := c.prepareFactory(clientCtx, txf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepareFactory")
	}

	txn, err := txf.BuildUnsignedTx(msgs...)
	if err != nil {
		err = errors.Wrap(err, "failed to BuildUnsignedTx")
		return nil, err
	}

	txn.SetFeeGranter(clientCtx.GetFeeGranterAddress())
	err = tx.Sign(txf, clientCtx.GetFromName(), txn, true)
	if err != nil {
		err = errors.Wrap(err, "failed to Sign Tx")
		return nil, err
	}

	return clientCtx.TxConfig.TxEncoder()(txn.GetTx())
}

func (c *chainClient) SyncBroadcastSignedTx(txBytes []byte) (*txtypes.BroadcastTxResponse, error) {
	req := txtypes.BroadcastTxRequest{
		TxBytes: txBytes,
		Mode:    txtypes.BroadcastMode_BROADCAST_MODE_SYNC,
	}

	ctx := context.Background()

	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.txClient.BroadcastTx, &req)
	if err != nil {
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
				if errRes := client.CheckTendermintError(err, txBytes); errRes != nil {
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
	req := txtypes.BroadcastTxRequest{
		TxBytes: txBytes,
		Mode:    txtypes.BroadcastMode_BROADCAST_MODE_SYNC,
	}

	ctx := context.Background()
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.txClient.BroadcastTx, &req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *chainClient) broadcastTx(
	clientCtx client.Context,
	txf tx.Factory,
	await bool,
	msgs ...sdk.Msg,
) (*txtypes.BroadcastTxResponse, error) {
	txf, err := c.prepareFactory(clientCtx, txf)
	if err != nil {
		err = errors.Wrap(err, "failed to prepareFactory")
		return nil, err
	}
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

	txn, err := txf.BuildUnsignedTx(msgs...)

	if err != nil {
		err = errors.Wrap(err, "failed to BuildUnsignedTx")
		return nil, err
	}

	txn.SetFeeGranter(clientCtx.GetFeeGranterAddress())
	err = tx.Sign(txf, clientCtx.GetFromName(), txn, true)
	if err != nil {
		err = errors.Wrap(err, "failed to Sign Tx")
		return nil, err
	}

	txBytes, err := clientCtx.TxConfig.TxEncoder()(txn.GetTx())
	if err != nil {
		err = errors.Wrap(err, "failed TxEncoder to encode Tx")
		return nil, err
	}

	req := txtypes.BroadcastTxRequest{
		TxBytes: txBytes,
		Mode:    txtypes.BroadcastMode_BROADCAST_MODE_SYNC,
	}

	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.txClient.BroadcastTx, &req)
	if !await || err != nil {
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
			resultTx, err := clientCtx.Client.Tx(awaitCtx, txHash, false)
			if err != nil {
				if errRes := client.CheckTendermintError(err, txBytes); errRes != nil {
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
		c.syncMux.Lock()
		defer c.syncMux.Unlock()
		sequence := c.getAccSeq()
		c.txFactory = c.txFactory.WithSequence(sequence)
		c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
		log.Debugln("broadcastTx with nonce", sequence)
		res, err := c.broadcastTx(c.ctx, c.txFactory, true, toSubmit...)
		if err != nil {
			if strings.Contains(err.Error(), "account sequence mismatch") {
				c.syncNonce()
				sequence := c.getAccSeq()
				c.txFactory = c.txFactory.WithSequence(sequence)
				c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
				log.Debugln("retrying broadcastTx with nonce", sequence)
				res, err = c.broadcastTx(c.ctx, c.txFactory, true, toSubmit...)
			}
			if err != nil {
				resJSON, _ := json.MarshalIndent(res, "", "\t")
				c.logger.WithField("size", len(toSubmit)).WithError(err).Errorln("failed to broadcast messages batch:", string(resJSON))
				return
			}
		}

		if res.TxResponse.Code != 0 {
			err = errors.Errorf("error %d (%s): %s", res.TxResponse.Code, res.TxResponse.Codespace, res.TxResponse.RawLog)
			log.WithField("txHash", res.TxResponse.TxHash).WithError(err).Errorln("failed to broadcast messages batch")
		} else {
			log.WithField("txHash", res.TxResponse.TxHash).Debugln("msg batch broadcasted successfully at height", res.TxResponse.Height)
		}

		log.Debugln("gas wanted: ", c.gasWanted)
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

func (c *chainClient) DefaultSubaccount(acc sdk.AccAddress) eth.Hash {
	return c.Subaccount(acc, 0)
}

func (c *chainClient) Subaccount(account sdk.AccAddress, index int) eth.Hash {
	ethAddress := eth.BytesToAddress(account.Bytes())
	ethLowerAddress := strings.ToLower(ethAddress.String())

	subaccountId := fmt.Sprintf("%s%024x", ethLowerAddress, index)
	return eth.HexToHash(subaccountId)
}

func (c *chainClient) GetSubAccountNonce(ctx context.Context, subaccountId eth.Hash) (*exchangetypes.QuerySubaccountTradeNonceResponse, error) {
	req := &exchangetypes.QuerySubaccountTradeNonceRequest{SubaccountId: subaccountId.String()}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.SubaccountTradeNonce, req)

	return res, err
}

// Deprecated: Use CreateSpotOrder instead
func (c *chainClient) SpotOrder(defaultSubaccountID eth.Hash, network common.Network, d *SpotOrderData) *exchangetypes.SpotOrder {
	assistant, err := NewMarketsAssistant(network.Name)
	if err != nil {
		panic(err)
	}

	return c.CreateSpotOrder(defaultSubaccountID, d, assistant)
}

func (c *chainClient) CreateSpotOrder(defaultSubaccountID eth.Hash, d *SpotOrderData, marketsAssistant MarketsAssistant) *exchangetypes.SpotOrder {

	market, isPresent := marketsAssistant.AllSpotMarkets()[d.MarketId]
	if !isPresent {
		panic(errors.Errorf("Invalid spot market id for %s network (%s)", c.network.Name, d.MarketId))
	}

	orderSize := market.QuantityToChainFormat(d.Quantity)
	orderPrice := market.PriceToChainFormat(d.Price)

	return &exchangetypes.SpotOrder{
		MarketId:  d.MarketId,
		OrderType: d.OrderType,
		OrderInfo: exchangetypes.OrderInfo{
			SubaccountId: defaultSubaccountID.Hex(),
			FeeRecipient: d.FeeRecipient,
			Price:        orderPrice,
			Quantity:     orderSize,
			Cid:          d.Cid,
		},
	}
}

// Deprecated: Use CreateDerivativeOrder instead
func (c *chainClient) DerivativeOrder(defaultSubaccountID eth.Hash, network common.Network, d *DerivativeOrderData) *exchangetypes.DerivativeOrder {

	assistant, err := NewMarketsAssistant(network.Name)
	if err != nil {
		panic(err)
	}

	return c.CreateDerivativeOrder(defaultSubaccountID, d, assistant)
}

func (c *chainClient) CreateDerivativeOrder(defaultSubaccountID eth.Hash, d *DerivativeOrderData, marketAssistant MarketsAssistant) *exchangetypes.DerivativeOrder {
	market, isPresent := marketAssistant.AllDerivativeMarkets()[d.MarketId]
	if !isPresent {
		panic(errors.Errorf("Invalid derivative market id for %s network (%s)", c.network.Name, d.MarketId))
	}

	orderSize := market.QuantityToChainFormat(d.Quantity)
	orderPrice := market.PriceToChainFormat(d.Price)
	orderMargin := sdk.MustNewDecFromStr("0")

	if !d.IsReduceOnly {
		orderMargin = market.CalculateMarginInChainFormat(d.Quantity, d.Price, d.Leverage)
	}

	return &exchangetypes.DerivativeOrder{
		MarketId:  d.MarketId,
		OrderType: d.OrderType,
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

func (c *chainClient) OrderCancel(defaultSubaccountID eth.Hash, d *OrderCancelData) *exchangetypes.OrderData {
	return &exchangetypes.OrderData{
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

func (c *chainClient) BuildGenericAuthz(granter string, grantee string, msgtype string, expireIn time.Time) *authztypes.MsgGrant {
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
	CreateSpotLimitOrderAuthz       = ExchangeAuthz("/" + proto.MessageName(&exchangetypes.CreateSpotLimitOrderAuthz{}))
	CreateSpotMarketOrderAuthz      = ExchangeAuthz("/" + proto.MessageName(&exchangetypes.CreateSpotMarketOrderAuthz{}))
	BatchCreateSpotLimitOrdersAuthz = ExchangeAuthz("/" + proto.MessageName(&exchangetypes.BatchCreateSpotLimitOrdersAuthz{}))
	CancelSpotOrderAuthz            = ExchangeAuthz("/" + proto.MessageName(&exchangetypes.CancelSpotOrderAuthz{}))
	BatchCancelSpotOrdersAuthz      = ExchangeAuthz("/" + proto.MessageName(&exchangetypes.BatchCancelSpotOrdersAuthz{}))

	CreateDerivativeLimitOrderAuthz       = ExchangeAuthz("/" + proto.MessageName(&exchangetypes.CreateDerivativeLimitOrderAuthz{}))
	CreateDerivativeMarketOrderAuthz      = ExchangeAuthz("/" + proto.MessageName(&exchangetypes.CreateDerivativeMarketOrderAuthz{}))
	BatchCreateDerivativeLimitOrdersAuthz = ExchangeAuthz("/" + proto.MessageName(&exchangetypes.BatchCreateDerivativeLimitOrdersAuthz{}))
	CancelDerivativeOrderAuthz            = ExchangeAuthz("/" + proto.MessageName(&exchangetypes.CancelDerivativeOrderAuthz{}))
	BatchCancelDerivativeOrdersAuthz      = ExchangeAuthz("/" + proto.MessageName(&exchangetypes.BatchCancelDerivativeOrdersAuthz{}))

	BatchUpdateOrdersAuthz = ExchangeAuthz("/" + proto.MessageName(&exchangetypes.BatchUpdateOrdersAuthz{}))
)

func (c *chainClient) BuildExchangeAuthz(granter string, grantee string, authzType ExchangeAuthz, subaccountId string, markets []string, expireIn time.Time) *authztypes.MsgGrant {
	var typedAuthzAny codectypes.Any
	var typedAuthzBytes []byte
	switch authzType {
	// spot msgs
	case CreateSpotLimitOrderAuthz:
		typedAuthz := &exchangetypes.CreateSpotLimitOrderAuthz{
			SubaccountId: subaccountId,
			MarketIds:    markets,
		}
		typedAuthzBytes, _ = typedAuthz.Marshal()
	case CreateSpotMarketOrderAuthz:
		typedAuthz := &exchangetypes.CreateSpotMarketOrderAuthz{
			SubaccountId: subaccountId,
			MarketIds:    markets,
		}
		typedAuthzBytes, _ = typedAuthz.Marshal()
	case BatchCreateSpotLimitOrdersAuthz:
		typedAuthz := &exchangetypes.BatchCreateSpotLimitOrdersAuthz{
			SubaccountId: subaccountId,
			MarketIds:    markets,
		}
		typedAuthzBytes, _ = typedAuthz.Marshal()
	case CancelSpotOrderAuthz:
		typedAuthz := &exchangetypes.CancelSpotOrderAuthz{
			SubaccountId: subaccountId,
			MarketIds:    markets,
		}
		typedAuthzBytes, _ = typedAuthz.Marshal()
	case BatchCancelSpotOrdersAuthz:
		typedAuthz := &exchangetypes.BatchCancelSpotOrdersAuthz{
			SubaccountId: subaccountId,
			MarketIds:    markets,
		}
		typedAuthzBytes, _ = typedAuthz.Marshal()
	// derivative msgs
	case CreateDerivativeLimitOrderAuthz:
		typedAuthz := &exchangetypes.CreateDerivativeLimitOrderAuthz{
			SubaccountId: subaccountId,
			MarketIds:    markets,
		}
		typedAuthzBytes, _ = typedAuthz.Marshal()
	case CreateDerivativeMarketOrderAuthz:
		typedAuthz := &exchangetypes.CreateDerivativeMarketOrderAuthz{
			SubaccountId: subaccountId,
			MarketIds:    markets,
		}
		typedAuthzBytes, _ = typedAuthz.Marshal()
	case BatchCreateDerivativeLimitOrdersAuthz:
		typedAuthz := &exchangetypes.BatchCreateDerivativeLimitOrdersAuthz{
			SubaccountId: subaccountId,
			MarketIds:    markets,
		}
		typedAuthzBytes, _ = typedAuthz.Marshal()
	case CancelDerivativeOrderAuthz:
		typedAuthz := &exchangetypes.CancelDerivativeOrderAuthz{
			SubaccountId: subaccountId,
			MarketIds:    markets,
		}
		typedAuthzBytes, _ = typedAuthz.Marshal()
	case BatchCancelDerivativeOrdersAuthz:
		typedAuthz := &exchangetypes.BatchCancelDerivativeOrdersAuthz{
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
	typedAuthz := &exchangetypes.BatchUpdateOrdersAuthz{
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

	cometbftClient, err = rpchttp.New(c.network.TmEndpoint, "/websocket")
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

func (c *chainClient) StreamOrderbookUpdateEvents(orderbookType OrderbookType, marketIds []string, orderbookCh chan exchangetypes.Orderbook) {
	var cometbftClient *rpchttp.HTTP
	var err error

	cometbftClient, err = rpchttp.New(c.network.TmEndpoint, "/websocket")
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

	c.StreamOrderbookUpdateEventsWithWebsocket(orderbookType, marketIds, cometbftClient, orderbookCh)

}

func (c *chainClient) StreamOrderbookUpdateEventsWithWebsocket(orderbookType OrderbookType, marketIds []string, websocket *rpchttp.HTTP, orderbookCh chan exchangetypes.Orderbook) {
	filter := fmt.Sprintf("tm.event='NewBlock' AND %s EXISTS", orderbookType)
	eventCh, err := websocket.Subscribe(context.Background(), "OrderbookUpdate", filter, 10000)
	if err != nil {
		panic(err)
	}

	// turn array into map for convenient lookup
	marketIdsMap := map[string]bool{}
	for _, id := range marketIds {
		marketIdsMap[id] = true
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
				if marketIdsMap[id] {
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

func (c *chainClient) ChainStream(ctx context.Context, req chainstreamtypes.StreamRequest) (chainstreamtypes.Stream_StreamClient, error) {
	stream, err := common.ExecuteStreamCall(ctx, c.network.ChainCookieAssistant, c.chainStreamClient.Stream, &req)

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

func (c *chainClient) FetchDenomAuthorityMetadata(ctx context.Context, creator string, subDenom string) (*tokenfactorytypes.QueryDenomAuthorityMetadataResponse, error) {
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
	OrderType    exchangetypes.OrderType
	Price        decimal.Decimal
	Quantity     decimal.Decimal
	Leverage     decimal.Decimal
	FeeRecipient string
	MarketId     string
	IsReduceOnly bool
	Cid          string
}

type SpotOrderData struct {
	OrderType    exchangetypes.OrderType
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

func (c *chainClient) FetchValidatorSlashes(ctx context.Context, validatorAddress string, startingHeight uint64, endingHeight uint64, pagination *query.PageRequest) (*distributiontypes.QueryValidatorSlashesResponse, error) {
	req := &distributiontypes.QueryValidatorSlashesRequest{
		ValidatorAddress: validatorAddress,
		StartingHeight:   startingHeight,
		EndingHeight:     endingHeight,
		Pagination:       pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.distributionQueryClient.ValidatorSlashes, req)

	return res, err
}

func (c *chainClient) FetchDelegationRewards(ctx context.Context, delegatorAddress string, validatorAddress string) (*distributiontypes.QueryDelegationRewardsResponse, error) {
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
func (c *chainClient) FetchSubaccountDeposits(ctx context.Context, subaccountId string) (*exchangetypes.QuerySubaccountDepositsResponse, error) {
	req := &exchangetypes.QuerySubaccountDepositsRequest{
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.SubaccountDeposits, req)

	return res, err
}

func (c *chainClient) FetchSubaccountDeposit(ctx context.Context, subaccountId string, denom string) (*exchangetypes.QuerySubaccountDepositResponse, error) {
	req := &exchangetypes.QuerySubaccountDepositRequest{
		SubaccountId: subaccountId,
		Denom:        denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.SubaccountDeposit, req)

	return res, err
}

func (c *chainClient) FetchExchangeBalances(ctx context.Context) (*exchangetypes.QueryExchangeBalancesResponse, error) {
	req := &exchangetypes.QueryExchangeBalancesRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.ExchangeBalances, req)

	return res, err
}

func (c *chainClient) FetchAggregateVolume(ctx context.Context, account string) (*exchangetypes.QueryAggregateVolumeResponse, error) {
	req := &exchangetypes.QueryAggregateVolumeRequest{
		Account: account,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.AggregateVolume, req)

	return res, err
}

func (c *chainClient) FetchAggregateVolumes(ctx context.Context, accounts []string, marketIds []string) (*exchangetypes.QueryAggregateVolumesResponse, error) {
	req := &exchangetypes.QueryAggregateVolumesRequest{
		Accounts:  accounts,
		MarketIds: marketIds,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.AggregateVolumes, req)

	return res, err
}

func (c *chainClient) FetchAggregateMarketVolume(ctx context.Context, marketId string) (*exchangetypes.QueryAggregateMarketVolumeResponse, error) {
	req := &exchangetypes.QueryAggregateMarketVolumeRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.AggregateMarketVolume, req)

	return res, err
}

func (c *chainClient) FetchAggregateMarketVolumes(ctx context.Context, marketIds []string) (*exchangetypes.QueryAggregateMarketVolumesResponse, error) {
	req := &exchangetypes.QueryAggregateMarketVolumesRequest{
		MarketIds: marketIds,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.AggregateMarketVolumes, req)

	return res, err
}

func (c *chainClient) FetchDenomDecimal(ctx context.Context, denom string) (*exchangetypes.QueryDenomDecimalResponse, error) {
	req := &exchangetypes.QueryDenomDecimalRequest{
		Denom: denom,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.DenomDecimal, req)

	return res, err
}

func (c *chainClient) FetchDenomDecimals(ctx context.Context, denoms []string) (*exchangetypes.QueryDenomDecimalsResponse, error) {
	req := &exchangetypes.QueryDenomDecimalsRequest{
		Denoms: denoms,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.DenomDecimals, req)

	return res, err
}

func (c *chainClient) FetchChainSpotMarkets(ctx context.Context, status string, marketIds []string) (*exchangetypes.QuerySpotMarketsResponse, error) {
	req := &exchangetypes.QuerySpotMarketsRequest{
		MarketIds: marketIds,
	}
	if status != "" {
		req.Status = status
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.SpotMarkets, req)

	return res, err
}

func (c *chainClient) FetchChainSpotMarket(ctx context.Context, marketId string) (*exchangetypes.QuerySpotMarketResponse, error) {
	req := &exchangetypes.QuerySpotMarketRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.SpotMarket, req)

	return res, err
}

func (c *chainClient) FetchChainFullSpotMarkets(ctx context.Context, status string, marketIds []string, withMidPriceAndTob bool) (*exchangetypes.QueryFullSpotMarketsResponse, error) {
	req := &exchangetypes.QueryFullSpotMarketsRequest{
		MarketIds:          marketIds,
		WithMidPriceAndTob: withMidPriceAndTob,
	}
	if status != "" {
		req.Status = status
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.FullSpotMarkets, req)

	return res, err
}

func (c *chainClient) FetchChainFullSpotMarket(ctx context.Context, marketId string, withMidPriceAndTob bool) (*exchangetypes.QueryFullSpotMarketResponse, error) {
	req := &exchangetypes.QueryFullSpotMarketRequest{
		MarketId:           marketId,
		WithMidPriceAndTob: withMidPriceAndTob,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.FullSpotMarket, req)

	return res, err
}

func (c *chainClient) FetchChainSpotOrderbook(ctx context.Context, marketId string, limit uint64, orderSide exchangetypes.OrderSide, limitCumulativeNotional sdk.Dec, limitCumulativeQuantity sdk.Dec) (*exchangetypes.QuerySpotOrderbookResponse, error) {
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

func (c *chainClient) FetchChainTraderSpotOrders(ctx context.Context, marketId string, subaccountId string) (*exchangetypes.QueryTraderSpotOrdersResponse, error) {
	req := &exchangetypes.QueryTraderSpotOrdersRequest{
		MarketId:     marketId,
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.TraderSpotOrders, req)

	return res, err
}

func (c *chainClient) FetchChainAccountAddressSpotOrders(ctx context.Context, marketId string, address string) (*exchangetypes.QueryAccountAddressSpotOrdersResponse, error) {
	req := &exchangetypes.QueryAccountAddressSpotOrdersRequest{
		MarketId:       marketId,
		AccountAddress: address,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.AccountAddressSpotOrders, req)

	return res, err
}

func (c *chainClient) FetchChainSpotOrdersByHashes(ctx context.Context, marketId string, subaccountId string, orderHashes []string) (*exchangetypes.QuerySpotOrdersByHashesResponse, error) {
	req := &exchangetypes.QuerySpotOrdersByHashesRequest{
		MarketId:     marketId,
		SubaccountId: subaccountId,
		OrderHashes:  orderHashes,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.SpotOrdersByHashes, req)

	return res, err
}

func (c *chainClient) FetchChainSubaccountOrders(ctx context.Context, subaccountId string, marketId string) (*exchangetypes.QuerySubaccountOrdersResponse, error) {
	req := &exchangetypes.QuerySubaccountOrdersRequest{
		SubaccountId: subaccountId,
		MarketId:     marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.SubaccountOrders, req)

	return res, err
}

func (c *chainClient) FetchChainTraderSpotTransientOrders(ctx context.Context, marketId string, subaccountId string) (*exchangetypes.QueryTraderSpotOrdersResponse, error) {
	req := &exchangetypes.QueryTraderSpotOrdersRequest{
		MarketId:     marketId,
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.TraderSpotTransientOrders, req)

	return res, err
}

func (c *chainClient) FetchSpotMidPriceAndTOB(ctx context.Context, marketId string) (*exchangetypes.QuerySpotMidPriceAndTOBResponse, error) {
	req := &exchangetypes.QuerySpotMidPriceAndTOBRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.SpotMidPriceAndTOB, req)

	return res, err
}

func (c *chainClient) FetchDerivativeMidPriceAndTOB(ctx context.Context, marketId string) (*exchangetypes.QueryDerivativeMidPriceAndTOBResponse, error) {
	req := &exchangetypes.QueryDerivativeMidPriceAndTOBRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.DerivativeMidPriceAndTOB, req)

	return res, err
}

func (c *chainClient) FetchChainDerivativeOrderbook(ctx context.Context, marketId string, limit uint64, limitCumulativeNotional sdk.Dec) (*exchangetypes.QueryDerivativeOrderbookResponse, error) {
	req := &exchangetypes.QueryDerivativeOrderbookRequest{
		MarketId:                marketId,
		Limit:                   limit,
		LimitCumulativeNotional: &limitCumulativeNotional,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.DerivativeOrderbook, req)

	return res, err
}

func (c *chainClient) FetchChainTraderDerivativeOrders(ctx context.Context, marketId string, subaccountId string) (*exchangetypes.QueryTraderDerivativeOrdersResponse, error) {
	req := &exchangetypes.QueryTraderDerivativeOrdersRequest{
		MarketId:     marketId,
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.TraderDerivativeOrders, req)

	return res, err
}

func (c *chainClient) FetchChainAccountAddressDerivativeOrders(ctx context.Context, marketId string, address string) (*exchangetypes.QueryAccountAddressDerivativeOrdersResponse, error) {
	req := &exchangetypes.QueryAccountAddressDerivativeOrdersRequest{
		MarketId:       marketId,
		AccountAddress: address,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.AccountAddressDerivativeOrders, req)

	return res, err
}

func (c *chainClient) FetchChainDerivativeOrdersByHashes(ctx context.Context, marketId string, subaccountId string, orderHashes []string) (*exchangetypes.QueryDerivativeOrdersByHashesResponse, error) {
	req := &exchangetypes.QueryDerivativeOrdersByHashesRequest{
		MarketId:     marketId,
		SubaccountId: subaccountId,
		OrderHashes:  orderHashes,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.DerivativeOrdersByHashes, req)

	return res, err
}

func (c *chainClient) FetchChainTraderDerivativeTransientOrders(ctx context.Context, marketId string, subaccountId string) (*exchangetypes.QueryTraderDerivativeOrdersResponse, error) {
	req := &exchangetypes.QueryTraderDerivativeOrdersRequest{
		MarketId:     marketId,
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.TraderDerivativeTransientOrders, req)

	return res, err
}

func (c *chainClient) FetchChainDerivativeMarkets(ctx context.Context, status string, marketIds []string, withMidPriceAndTob bool) (*exchangetypes.QueryDerivativeMarketsResponse, error) {
	req := &exchangetypes.QueryDerivativeMarketsRequest{
		MarketIds:          marketIds,
		WithMidPriceAndTob: withMidPriceAndTob,
	}
	if status != "" {
		req.Status = status
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.DerivativeMarkets, req)

	return res, err
}

func (c *chainClient) FetchChainDerivativeMarket(ctx context.Context, marketId string) (*exchangetypes.QueryDerivativeMarketResponse, error) {
	req := &exchangetypes.QueryDerivativeMarketRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.DerivativeMarket, req)

	return res, err
}

func (c *chainClient) FetchDerivativeMarketAddress(ctx context.Context, marketId string) (*exchangetypes.QueryDerivativeMarketAddressResponse, error) {
	req := &exchangetypes.QueryDerivativeMarketAddressRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.DerivativeMarketAddress, req)

	return res, err
}

func (c *chainClient) FetchSubaccountTradeNonce(ctx context.Context, subaccountId string) (*exchangetypes.QuerySubaccountTradeNonceResponse, error) {
	req := &exchangetypes.QuerySubaccountTradeNonceRequest{
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.SubaccountTradeNonce, req)

	return res, err
}

func (c *chainClient) FetchChainPositions(ctx context.Context) (*exchangetypes.QueryPositionsResponse, error) {
	req := &exchangetypes.QueryPositionsRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.Positions, req)

	return res, err
}

func (c *chainClient) FetchChainSubaccountPositions(ctx context.Context, subaccountId string) (*exchangetypes.QuerySubaccountPositionsResponse, error) {
	req := &exchangetypes.QuerySubaccountPositionsRequest{
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.SubaccountPositions, req)

	return res, err
}

func (c *chainClient) FetchChainSubaccountPositionInMarket(ctx context.Context, subaccountId string, marketId string) (*exchangetypes.QuerySubaccountPositionInMarketResponse, error) {
	req := &exchangetypes.QuerySubaccountPositionInMarketRequest{
		SubaccountId: subaccountId,
		MarketId:     marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.SubaccountPositionInMarket, req)

	return res, err
}

func (c *chainClient) FetchChainSubaccountEffectivePositionInMarket(ctx context.Context, subaccountId string, marketId string) (*exchangetypes.QuerySubaccountEffectivePositionInMarketResponse, error) {
	req := &exchangetypes.QuerySubaccountEffectivePositionInMarketRequest{
		SubaccountId: subaccountId,
		MarketId:     marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.SubaccountEffectivePositionInMarket, req)

	return res, err
}

func (c *chainClient) FetchChainPerpetualMarketInfo(ctx context.Context, marketId string) (*exchangetypes.QueryPerpetualMarketInfoResponse, error) {
	req := &exchangetypes.QueryPerpetualMarketInfoRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.PerpetualMarketInfo, req)

	return res, err
}

func (c *chainClient) FetchChainExpiryFuturesMarketInfo(ctx context.Context, marketId string) (*exchangetypes.QueryExpiryFuturesMarketInfoResponse, error) {
	req := &exchangetypes.QueryExpiryFuturesMarketInfoRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.ExpiryFuturesMarketInfo, req)

	return res, err
}

func (c *chainClient) FetchChainPerpetualMarketFunding(ctx context.Context, marketId string) (*exchangetypes.QueryPerpetualMarketFundingResponse, error) {
	req := &exchangetypes.QueryPerpetualMarketFundingRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.PerpetualMarketFunding, req)

	return res, err
}

func (c *chainClient) FetchSubaccountOrderMetadata(ctx context.Context, subaccountId string) (*exchangetypes.QuerySubaccountOrderMetadataResponse, error) {
	req := &exchangetypes.QuerySubaccountOrderMetadataRequest{
		SubaccountId: subaccountId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.SubaccountOrderMetadata, req)

	return res, err
}

func (c *chainClient) FetchTradeRewardPoints(ctx context.Context, accounts []string) (*exchangetypes.QueryTradeRewardPointsResponse, error) {
	req := &exchangetypes.QueryTradeRewardPointsRequest{
		Accounts: accounts,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.TradeRewardPoints, req)

	return res, err
}

func (c *chainClient) FetchPendingTradeRewardPoints(ctx context.Context, accounts []string) (*exchangetypes.QueryTradeRewardPointsResponse, error) {
	req := &exchangetypes.QueryTradeRewardPointsRequest{
		Accounts: accounts,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.PendingTradeRewardPoints, req)

	return res, err
}

func (c *chainClient) FetchTradeRewardCampaign(ctx context.Context) (*exchangetypes.QueryTradeRewardCampaignResponse, error) {
	req := &exchangetypes.QueryTradeRewardCampaignRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.TradeRewardCampaign, req)

	return res, err
}

func (c *chainClient) FetchFeeDiscountAccountInfo(ctx context.Context, account string) (*exchangetypes.QueryFeeDiscountAccountInfoResponse, error) {
	req := &exchangetypes.QueryFeeDiscountAccountInfoRequest{
		Account: account,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.FeeDiscountAccountInfo, req)

	return res, err
}

func (c *chainClient) FetchFeeDiscountSchedule(ctx context.Context) (*exchangetypes.QueryFeeDiscountScheduleResponse, error) {
	req := &exchangetypes.QueryFeeDiscountScheduleRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.FeeDiscountSchedule, req)

	return res, err
}

func (c *chainClient) FetchBalanceMismatches(ctx context.Context, dustFactor int64) (*exchangetypes.QueryBalanceMismatchesResponse, error) {
	req := &exchangetypes.QueryBalanceMismatchesRequest{
		DustFactor: dustFactor,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.BalanceMismatches, req)

	return res, err
}

func (c *chainClient) FetchBalanceWithBalanceHolds(ctx context.Context) (*exchangetypes.QueryBalanceWithBalanceHoldsResponse, error) {
	req := &exchangetypes.QueryBalanceWithBalanceHoldsRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.BalanceWithBalanceHolds, req)

	return res, err
}

func (c *chainClient) FetchFeeDiscountTierStatistics(ctx context.Context) (*exchangetypes.QueryFeeDiscountTierStatisticsResponse, error) {
	req := &exchangetypes.QueryFeeDiscountTierStatisticsRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.FeeDiscountTierStatistics, req)

	return res, err
}

func (c *chainClient) FetchMitoVaultInfos(ctx context.Context) (*exchangetypes.MitoVaultInfosResponse, error) {
	req := &exchangetypes.MitoVaultInfosRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.MitoVaultInfos, req)

	return res, err
}

func (c *chainClient) FetchMarketIDFromVault(ctx context.Context, vaultAddress string) (*exchangetypes.QueryMarketIDFromVaultResponse, error) {
	req := &exchangetypes.QueryMarketIDFromVaultRequest{
		VaultAddress: vaultAddress,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.QueryMarketIDFromVault, req)

	return res, err
}

func (c *chainClient) FetchHistoricalTradeRecords(ctx context.Context, marketId string) (*exchangetypes.QueryHistoricalTradeRecordsResponse, error) {
	req := &exchangetypes.QueryHistoricalTradeRecordsRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.HistoricalTradeRecords, req)

	return res, err
}

func (c *chainClient) FetchIsOptedOutOfRewards(ctx context.Context, account string) (*exchangetypes.QueryIsOptedOutOfRewardsResponse, error) {
	req := &exchangetypes.QueryIsOptedOutOfRewardsRequest{
		Account: account,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.IsOptedOutOfRewards, req)

	return res, err
}

func (c *chainClient) FetchOptedOutOfRewardsAccounts(ctx context.Context) (*exchangetypes.QueryOptedOutOfRewardsAccountsResponse, error) {
	req := &exchangetypes.QueryOptedOutOfRewardsAccountsRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.OptedOutOfRewardsAccounts, req)

	return res, err
}

func (c *chainClient) FetchMarketVolatility(ctx context.Context, marketId string, tradeHistoryOptions *exchangetypes.TradeHistoryOptions) (*exchangetypes.QueryMarketVolatilityResponse, error) {
	req := &exchangetypes.QueryMarketVolatilityRequest{
		MarketId:            marketId,
		TradeHistoryOptions: tradeHistoryOptions,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.MarketVolatility, req)

	return res, err
}

func (c *chainClient) FetchChainBinaryOptionsMarkets(ctx context.Context, status string) (*exchangetypes.QueryBinaryMarketsResponse, error) {
	req := &exchangetypes.QueryBinaryMarketsRequest{}
	if status != "" {
		req.Status = status
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.BinaryOptionsMarkets, req)

	return res, err
}

func (c *chainClient) FetchTraderDerivativeConditionalOrders(ctx context.Context, subaccountId string, marketId string) (*exchangetypes.QueryTraderDerivativeConditionalOrdersResponse, error) {
	req := &exchangetypes.QueryTraderDerivativeConditionalOrdersRequest{
		SubaccountId: subaccountId,
		MarketId:     marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.TraderDerivativeConditionalOrders, req)

	return res, err
}

func (c *chainClient) FetchMarketAtomicExecutionFeeMultiplier(ctx context.Context, marketId string) (*exchangetypes.QueryMarketAtomicExecutionFeeMultiplierResponse, error) {
	req := &exchangetypes.QueryMarketAtomicExecutionFeeMultiplierRequest{
		MarketId: marketId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.exchangeQueryClient.MarketAtomicExecutionFeeMultiplier, req)

	return res, err
}

// Tendermint module

func (c *chainClient) FetchNodeInfo(ctx context.Context) (*tmservice.GetNodeInfoResponse, error) {
	req := &tmservice.GetNodeInfoRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tendermintQueryClient.GetNodeInfo, req)

	return res, err
}

func (c *chainClient) FetchSyncing(ctx context.Context) (*tmservice.GetSyncingResponse, error) {
	req := &tmservice.GetSyncingRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tendermintQueryClient.GetSyncing, req)

	return res, err
}

func (c *chainClient) FetchLatestBlock(ctx context.Context) (*tmservice.GetLatestBlockResponse, error) {
	req := &tmservice.GetLatestBlockRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tendermintQueryClient.GetLatestBlock, req)

	return res, err
}

func (c *chainClient) FetchBlockByHeight(ctx context.Context, height int64) (*tmservice.GetBlockByHeightResponse, error) {
	req := &tmservice.GetBlockByHeightRequest{
		Height: height,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tendermintQueryClient.GetBlockByHeight, req)

	return res, err
}

func (c *chainClient) FetchLatestValidatorSet(ctx context.Context) (*tmservice.GetLatestValidatorSetResponse, error) {
	req := &tmservice.GetLatestValidatorSetRequest{}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tendermintQueryClient.GetLatestValidatorSet, req)

	return res, err
}

func (c *chainClient) FetchValidatorSetByHeight(ctx context.Context, height int64, pagination *query.PageRequest) (*tmservice.GetValidatorSetByHeightResponse, error) {
	req := &tmservice.GetValidatorSetByHeightRequest{
		Height:     height,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.tendermintQueryClient.GetValidatorSetByHeight, req)

	return res, err
}

func (c *chainClient) ABCIQuery(ctx context.Context, path string, data []byte, height int64, prove bool) (*tmservice.ABCIQueryResponse, error) {
	req := &tmservice.ABCIQueryRequest{
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

func (c *chainClient) FetchEscrowAddress(ctx context.Context, portId string, channelId string) (*ibctransfertypes.QueryEscrowAddressResponse, error) {
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
func (c *chainClient) FetchIBCChannel(ctx context.Context, portId string, channelId string) (*ibcchanneltypes.QueryChannelResponse, error) {
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

func (c *chainClient) FetchIBCChannelClientState(ctx context.Context, portId string, channelId string) (*ibcchanneltypes.QueryChannelClientStateResponse, error) {
	req := &ibcchanneltypes.QueryChannelClientStateRequest{
		PortId:    portId,
		ChannelId: channelId,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.ChannelClientState, req)

	return res, err
}

func (c *chainClient) FetchIBCChannelConsensusState(ctx context.Context, portId string, channelId string, revisionNumber uint64, revisionHeight uint64) (*ibcchanneltypes.QueryChannelConsensusStateResponse, error) {
	req := &ibcchanneltypes.QueryChannelConsensusStateRequest{
		PortId:         portId,
		ChannelId:      channelId,
		RevisionNumber: revisionNumber,
		RevisionHeight: revisionHeight,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.ChannelConsensusState, req)

	return res, err
}

func (c *chainClient) FetchIBCPacketCommitment(ctx context.Context, portId string, channelId string, sequence uint64) (*ibcchanneltypes.QueryPacketCommitmentResponse, error) {
	req := &ibcchanneltypes.QueryPacketCommitmentRequest{
		PortId:    portId,
		ChannelId: channelId,
		Sequence:  sequence,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.PacketCommitment, req)

	return res, err
}

func (c *chainClient) FetchIBCPacketCommitments(ctx context.Context, portId string, channelId string, pagination *query.PageRequest) (*ibcchanneltypes.QueryPacketCommitmentsResponse, error) {
	req := &ibcchanneltypes.QueryPacketCommitmentsRequest{
		PortId:     portId,
		ChannelId:  channelId,
		Pagination: pagination,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.PacketCommitments, req)

	return res, err
}

func (c *chainClient) FetchIBCPacketReceipt(ctx context.Context, portId string, channelId string, sequence uint64) (*ibcchanneltypes.QueryPacketReceiptResponse, error) {
	req := &ibcchanneltypes.QueryPacketReceiptRequest{
		PortId:    portId,
		ChannelId: channelId,
		Sequence:  sequence,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.PacketReceipt, req)

	return res, err
}

func (c *chainClient) FetchIBCPacketAcknowledgement(ctx context.Context, portId string, channelId string, sequence uint64) (*ibcchanneltypes.QueryPacketAcknowledgementResponse, error) {
	req := &ibcchanneltypes.QueryPacketAcknowledgementRequest{
		PortId:    portId,
		ChannelId: channelId,
		Sequence:  sequence,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.PacketAcknowledgement, req)

	return res, err
}

func (c *chainClient) FetchIBCPacketAcknowledgements(ctx context.Context, portId string, channelId string, packetCommitmentSequences []uint64, pagination *query.PageRequest) (*ibcchanneltypes.QueryPacketAcknowledgementsResponse, error) {
	req := &ibcchanneltypes.QueryPacketAcknowledgementsRequest{
		PortId:                    portId,
		ChannelId:                 channelId,
		Pagination:                pagination,
		PacketCommitmentSequences: packetCommitmentSequences,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.PacketAcknowledgements, req)

	return res, err
}

func (c *chainClient) FetchIBCUnreceivedPackets(ctx context.Context, portId string, channelId string, packetCommitmentSequences []uint64) (*ibcchanneltypes.QueryUnreceivedPacketsResponse, error) {
	req := &ibcchanneltypes.QueryUnreceivedPacketsRequest{
		PortId:                    portId,
		ChannelId:                 channelId,
		PacketCommitmentSequences: packetCommitmentSequences,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.UnreceivedPackets, req)

	return res, err
}

func (c *chainClient) FetchIBCUnreceivedAcks(ctx context.Context, portId string, channelId string, packetAckSequences []uint64) (*ibcchanneltypes.QueryUnreceivedAcksResponse, error) {
	req := &ibcchanneltypes.QueryUnreceivedAcksRequest{
		PortId:             portId,
		ChannelId:          channelId,
		PacketAckSequences: packetAckSequences,
	}
	res, err := common.ExecuteCall(ctx, c.network.ChainCookieAssistant, c.ibcChannelQueryClient.UnreceivedAcks, req)

	return res, err
}

func (c *chainClient) FetchIBCNextSequenceReceive(ctx context.Context, portId string, channelId string) (*ibcchanneltypes.QueryNextSequenceReceiveResponse, error) {
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

func (c *chainClient) FetchIBCConsensusState(ctx context.Context, clientId string, revisionNumber uint64, revisionHeight uint64, latestHeight bool) (*ibcclienttypes.QueryConsensusStateResponse, error) {
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
