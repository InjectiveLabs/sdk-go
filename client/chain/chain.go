package chain

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/cosmos/cosmos-sdk/types/query"

	"google.golang.org/grpc/credentials/insecure"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/gogoproto/proto"
	eth "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	chainstreamtypes "github.com/InjectiveLabs/sdk-go/chain/stream/types"
	tokenfactorytypes "github.com/InjectiveLabs/sdk-go/chain/tokenfactory/types"
	"github.com/InjectiveLabs/sdk-go/client/common"

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
	defaultTimeoutHeightSyncInterval = 30 * time.Second
	defaultSessionRenewalOffset      = 120
	defaultBlockTime                 = 3 * time.Second
	defaultChainCookieName           = ".chain_cookie"

	MaxGasFee = "1000000000inj"
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

	GetBlockHeight() (int64, error)

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

	DefaultSubaccount(acc cosmtypes.AccAddress) eth.Hash
	Subaccount(account cosmtypes.AccAddress, index int) eth.Hash

	GetSubAccountNonce(ctx context.Context, subaccountId eth.Hash) (*exchangetypes.QuerySubaccountTradeNonceResponse, error)
	GetFeeDiscountInfo(ctx context.Context, account string) (*exchangetypes.QueryFeeDiscountAccountInfoResponse, error)

	UpdateSubaccountNonceFromChain() error
	SynchronizeSubaccountNonce(subaccountId eth.Hash) error
	ComputeOrderHashes(spotOrders []exchangetypes.SpotOrder, derivativeOrders []exchangetypes.DerivativeOrder, subaccountId eth.Hash) (OrderHashes, error)

	SpotOrder(defaultSubaccountID eth.Hash, network common.Network, d *SpotOrderData) *exchangetypes.SpotOrder
	CreateSpotOrder(defaultSubaccountID eth.Hash, network common.Network, d *SpotOrderData, marketsAssistant MarketsAssistant) *exchangetypes.SpotOrder
	DerivativeOrder(defaultSubaccountID eth.Hash, network common.Network, d *DerivativeOrderData) *exchangetypes.DerivativeOrder
	CreateDerivativeOrder(defaultSubaccountID eth.Hash, network common.Network, d *DerivativeOrderData, marketAssistant MarketsAssistant) *exchangetypes.DerivativeOrder
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

	AdjustGasPricesToMax()
	AdjustedGasPricesToOrigin()

	Close()
}

type chainClient struct {
	ctx             client.Context
	network         common.Network
	opts            *common.ClientOptions
	logger          *logrus.Logger
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

	nonce                   uint32
	closeRoutineUpdateNonce context.CancelFunc

	sessionCookie  string
	sessionEnabled bool

	txClient                txtypes.ServiceClient
	authQueryClient         authtypes.QueryClient
	exchangeQueryClient     exchangetypes.QueryClient
	bankQueryClient         banktypes.QueryClient
	authzQueryClient        authztypes.QueryClient
	wasmQueryClient         wasmtypes.QueryClient
	chainStreamClient       chainstreamtypes.StreamClient
	tokenfactoryQueryClient tokenfactorytypes.QueryClient
	subaccountToNonce       map[ethcommon.Hash]uint32

	closed  int64
	canSign bool

	isDynamicGasPrices bool
}

func NewChainClient(
	ctx client.Context,
	network common.Network,
	options ...common.ClientOption,
) (ChainClient, error) {
	var err error
	// process options
	opts := common.DefaultClientOptions()

	if network.ChainTlsCert != nil {
		options = append(options, common.OptionTLSCert(network.ChainTlsCert))
	}
	for _, opt := range options {
		if err = opt(opts); err != nil {
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
	defer func() {
		if err != nil {
			conn.Close()
		}
	}()

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
	defer func() {
		if err != nil {
			chainStreamConn.Close()
		}
	}()

	cancelCtx, cancelFn := context.WithCancel(context.Background())
	// build client
	cc := &chainClient{
		ctx:     ctx,
		network: network,
		opts:    opts,

		logger: opts.Logger,

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

		txClient:                txtypes.NewServiceClient(conn),
		authQueryClient:         authtypes.NewQueryClient(conn),
		exchangeQueryClient:     exchangetypes.NewQueryClient(conn),
		bankQueryClient:         banktypes.NewQueryClient(conn),
		authzQueryClient:        authztypes.NewQueryClient(conn),
		wasmQueryClient:         wasmtypes.NewQueryClient(conn),
		chainStreamClient:       chainstreamtypes.NewStreamClient(chainStreamConn),
		tokenfactoryQueryClient: tokenfactorytypes.NewQueryClient(conn),
		subaccountToNonce:       make(map[ethcommon.Hash]uint32),

		isDynamicGasPrices: opts.IsDynamicGasPrices,
	}
	defer func() {
		if err != nil {
			cc.Close()
		}
	}()

	// routine upate nonce
	if cc.canSign {
		closeRoutineUpdateNonce := cc.RoutineUpdateNounce()
		cc.closeRoutineUpdateNonce = closeRoutineUpdateNonce
	}

	if cc.canSign {

		cc.accNum, cc.accSeq, err = cc.txFactory.AccountRetriever().GetAccountNumberSequence(ctx, ctx.GetFromAddress())
		if err != nil {
			err = errors.Wrap(err, "failed to get initial account num and seq")
			return nil, err
		}

		go cc.runBatchBroadcast()
		go cc.syncTimeoutHeight()
	}

	// create file if not exist
	var cookieFile *os.File
	cookieFile, err = os.OpenFile(defaultChainCookieName, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		cc.logger.Errorln(err)
	} else {
		defer cookieFile.Close()
	}

	// attempt to load from disk
	var data []byte
	data, err = os.ReadFile(defaultChainCookieName)
	if err != nil {
		cc.logger.Errorf("[INJ-GO-SDK] Failed to read default chain cookie %q: %v", defaultChainCookieName, err)
	} else {
		cc.sessionCookie = string(data)
		cc.logger.Infoln("[INJ-GO-SDK] Chain session cookie loaded from disk")
	}

	return cc, nil
}

func (c *chainClient) syncNonce() {
	num, seq, err := c.txFactory.AccountRetriever().GetAccountNumberSequence(c.ctx, c.ctx.GetFromAddress())
	if err != nil {
		c.logger.Errorln("[INJ-GO-SDK] Failed to get account seq: ", err)
		return
	} else if num != c.accNum {
		c.logger.Errorf("[INJ-GO-SDK] Account number changed during nonce sync: expected: %s,actual: %s", c.accNum, num)
	}

	c.accSeq = seq
}

func (c *chainClient) syncTimeoutHeight() {
	t := time.NewTicker(defaultTimeoutHeightSyncInterval)
	defer t.Stop()

	for {
		block, err := c.ctx.Client.Block(c.cancelCtx, nil)
		if err != nil {
			c.logger.Errorln("[INJ-GO-SDK] Failed to get current block: ", err)
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

// test
func (c *chainClient) GetBlockHeight() (int64, error) {
	ctx := context.Background()
	if c.ctx.Client == nil {
		return 0, errors.New("client is nil")
	}

	block, err := c.ctx.Client.Block(ctx, nil)
	if err != nil {
		return 0, err
	}
	return block.Block.Height, nil
}

// prepareFactory ensures the account defined by ctx.GetFromAddress() exists and
// if the account number and/or the account sequence number are zero (not set),
// they will be queried for and set on the provided Factory. A new Factory with
// the updated fields will be returned.
func (c *chainClient) prepareFactory(clientCtx client.Context, txf tx.Factory) (tx.Factory, error) {
	from := clientCtx.GetFromAddress()

	// if err := txf.AccountRetriever().EnsureExists(clientCtx, from); err != nil {
	// 	return txf, err
	// }

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

func (c *chainClient) requestCookie() metadata.MD {
	var header metadata.MD
	_, err := c.bankQueryClient.Params(context.Background(), &banktypes.QueryParamsRequest{}, grpc.Header(&header))
	if err != nil {
		c.logger.Errorln("[INJ-GO-SDK] Failed to get chain cookie: ", err)
	}
	return header
}

func (c *chainClient) getCookie(ctx context.Context) context.Context {
	provider := common.NewMetadataProvider(c.requestCookie)
	cookie, _ := c.network.ChainMetadata(provider)
	md := metadata.Pairs("cookie", cookie)
	return metadata.NewOutgoingContext(ctx, md)
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
	// if !c.canSign {
	// 	return
	// }
	if c.msgC != nil {
		if atomic.CompareAndSwapInt64(&c.closed, 0, 1) {
			close(c.msgC)
		}
	}

	if c.cancelFn != nil {
		c.cancelFn()
	}
	// <-c.doneC // not used
	if c.conn != nil {
		c.conn.Close()
	}

	if c.closeRoutineUpdateNonce != nil {
		c.closeRoutineUpdateNonce()
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
	return c.bankQueryClient.AllBalances(ctx, req)
}

func (c *chainClient) GetBankBalance(ctx context.Context, address string, denom string) (*banktypes.QueryBalanceResponse, error) {
	req := &banktypes.QueryBalanceRequest{
		Address: address,
		Denom:   denom,
	}
	return c.bankQueryClient.Balance(ctx, req)
}

func (c *chainClient) GetBankSpendableBalances(ctx context.Context, address string, pagination *query.PageRequest) (*banktypes.QuerySpendableBalancesResponse, error) {
	req := &banktypes.QuerySpendableBalancesRequest{
		Address:    address,
		Pagination: pagination,
	}
	return c.bankQueryClient.SpendableBalances(ctx, req)
}

func (c *chainClient) GetBankSpendableBalancesByDenom(ctx context.Context, address string, denom string) (*banktypes.QuerySpendableBalanceByDenomResponse, error) {
	req := &banktypes.QuerySpendableBalanceByDenomRequest{
		Address: address,
		Denom:   denom,
	}
	return c.bankQueryClient.SpendableBalanceByDenom(ctx, req)
}

func (c *chainClient) GetBankTotalSupply(ctx context.Context, pagination *query.PageRequest) (*banktypes.QueryTotalSupplyResponse, error) {
	req := &banktypes.QueryTotalSupplyRequest{Pagination: pagination}
	return c.bankQueryClient.TotalSupply(ctx, req)
}

func (c *chainClient) GetBankSupplyOf(ctx context.Context, denom string) (*banktypes.QuerySupplyOfResponse, error) {
	req := &banktypes.QuerySupplyOfRequest{Denom: denom}
	return c.bankQueryClient.SupplyOf(ctx, req)
}

func (c *chainClient) GetDenomMetadata(ctx context.Context, denom string) (*banktypes.QueryDenomMetadataResponse, error) {
	req := &banktypes.QueryDenomMetadataRequest{Denom: denom}
	return c.bankQueryClient.DenomMetadata(ctx, req)
}

func (c *chainClient) GetDenomsMetadata(ctx context.Context, pagination *query.PageRequest) (*banktypes.QueryDenomsMetadataResponse, error) {
	req := &banktypes.QueryDenomsMetadataRequest{Pagination: pagination}
	return c.bankQueryClient.DenomsMetadata(ctx, req)
}

func (c *chainClient) GetDenomOwners(ctx context.Context, denom string, pagination *query.PageRequest) (*banktypes.QueryDenomOwnersResponse, error) {
	req := &banktypes.QueryDenomOwnersRequest{
		Denom:      denom,
		Pagination: pagination,
	}
	return c.bankQueryClient.DenomOwners(ctx, req)
}

func (c *chainClient) GetBankSendEnabled(ctx context.Context, denoms []string, pagination *query.PageRequest) (*banktypes.QuerySendEnabledResponse, error) {
	req := &banktypes.QuerySendEnabledRequest{
		Denoms:     denoms,
		Pagination: pagination,
	}
	return c.bankQueryClient.SendEnabled(ctx, req)
}

// Auth Module

func (c *chainClient) GetAccount(ctx context.Context, address string) (*authtypes.QueryAccountResponse, error) {
	req := &authtypes.QueryAccountRequest{
		Address: address,
	}
	return c.authQueryClient.Account(ctx, req)
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
			c.logger.Debugln("[INJ-GO-SDK] Retrying broadcastTx with nonce: ", c.accSeq)
			res, err = c.broadcastTx(c.ctx, c.txFactory, true, msgs...)
		}
		if err != nil {
			resJSON, _ := json.MarshalIndent(res, "", "\t")
			c.logger.Errorf("[INJ-GO-SDK] Failed to commit msg batch: %s, size: %d: %v", string(resJSON), len(msgs), err)
			return nil, err
		}
	}

	return res, nil
}

func (c *chainClient) GetFeeDiscountInfo(ctx context.Context, account string) (*exchangetypes.QueryFeeDiscountAccountInfoResponse, error) {
	req := &exchangetypes.QueryFeeDiscountAccountInfoRequest{
		Account: account,
	}
	return c.exchangeQueryClient.FeeDiscountAccountInfo(ctx, req)
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
	ctx = c.getCookie(ctx)
	simRes, err := c.txClient.Simulate(ctx, &txtypes.SimulateRequest{TxBytes: simTxBytes})
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
	broadcastFunc := func() (*txtypes.BroadcastTxResponse, []byte, error) {
		c.syncMux.Lock()
		defer c.syncMux.Unlock()

		c.txFactory = c.txFactory.WithSequence(c.accSeq)
		c.txFactory = c.txFactory.WithAccountNumber(c.accNum)

		c.logger.Infoln("[INJ-GO-SDK] Sending chain msg: ", time.Now().Format("2006-01-02 15:04:05"))

		res, txBytes, err := c.broadcastTxAsync(c.ctx, c.txFactory, false, msgs...)
		if err != nil {
			if strings.Contains(err.Error(), "account sequence mismatch") {
				c.syncNonce()

				c.txFactory = c.txFactory.WithSequence(c.accSeq)
				c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
				c.logger.Infof("[INJ-GO-SDK] Retrying broadcastTx with nonce: %d ,err: %v", c.accSeq, err)
				res, txBytes, err = c.broadcastTxAsync(c.ctx, c.txFactory, false, msgs...)
			}
			if err != nil {
				resJSON, _ := json.MarshalIndent(res, "", "\t")
				c.logger.Errorf("[INJ-GO-SDK] Failed to commit msg batch: %s, size: %d: %v", string(resJSON), len(msgs), err)
				return nil, nil, err
			}
		}

		c.accSeq++

		c.logger.Debugln("[INJ-GO-SDK] Nonce incremented to ", c.accSeq)
		c.logger.Debugln("[INJ-GO-SDK] Gas wanted: ", c.gasWanted)

		return res, txBytes, nil
	}

	res, txBytes, err := broadcastFunc()
	if err != nil {
		return nil, err
	}

	res, err = c.PollTxResults(res, c.ctx, txBytes)
	if err != nil {
		c.accSeq--
		if strings.Contains(err.Error(), ErrTimedOut.Error()) && c.isDynamicGasPrices {
			c.logger.Debugln("[INJ-GO-SDK] Update gas to max gas price")
			c.txFactory = c.txFactory.WithGasPrices(MaxGasFee)
		}
		return res, err
	} else if res != nil && res.TxResponse != nil {
		c.txFactory = c.txFactory.WithGasPrices(c.opts.GasPrices)
		if res.TxResponse.Code != 0 {
			err = errors.Errorf("error %d (%s): %s", res.TxResponse.Code, res.TxResponse.Codespace, res.TxResponse.RawLog)
			c.logger.Errorf("[INJ-GO-SDK] Failed to commit msg batch, txHash: %s with err: %v", res.TxResponse.TxHash, err)
			return res, err
		} else {
			c.logger.Debugln("[INJ-GO-SDK] Msg batch committed successfully at height: ", res.TxResponse.Height)
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
		ctx := c.getCookie(context.Background())
		var header metadata.MD
		simRes, err := c.txClient.Simulate(ctx, &txtypes.SimulateRequest{TxBytes: simTxBytes}, grpc.Header(&header))
		if err != nil {
			err = errors.Wrap(err, "failed to CalculateGas")
			return nil, err
		}

		adjustedGas := uint64(txf.GasAdjustment()*float64(simRes.GasInfo.GasUsed)) * 12 / 10
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
	ctx = c.getCookie(ctx)
	res, err := c.txClient.BroadcastTx(ctx, &req)
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
	// use our own client to broadcast tx
	ctx = c.getCookie(ctx)
	res, err := c.txClient.BroadcastTx(ctx, &req)
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
		ctx := c.getCookie(ctx)
		simRes, err := c.txClient.Simulate(ctx, &txtypes.SimulateRequest{TxBytes: simTxBytes})
		if err != nil {
			err = errors.Wrap(err, "failed to CalculateGas")
			return nil, err
		}

		adjustedGas := uint64(txf.GasAdjustment()*float64(simRes.GasInfo.GasUsed)) * 12 / 10
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
	// use our own client to broadcast tx
	ctx = c.getCookie(ctx)
	res, err := c.txClient.BroadcastTx(ctx, &req)
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

				c.logger.Debugf("Tx Error for Hash: %s: %v", res.TxResponse.TxHash, err)

				// log.WithError(err).Warningln("Tx Error for Hash:", res.TxHash)

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

func (c *chainClient) broadcastTxAsync(
	clientCtx client.Context,
	txf tx.Factory,
	await bool,
	msgs ...sdk.Msg,
) (*txtypes.BroadcastTxResponse, []byte, error) {
	txf, err := c.prepareFactory(clientCtx, txf)
	if err != nil {
		err = errors.Wrap(err, "failed to prepareFactory")
		return nil, nil, err
	}
	ctx := context.Background()
	if clientCtx.Simulate {
		simTxBytes, err := txf.BuildSimTx(msgs...)
		if err != nil {
			err = errors.Wrap(err, "failed to build sim tx bytes")
			return nil, nil, err
		}
		ctx := c.getCookie(ctx)
		var header metadata.MD
		simRes, err := c.txClient.Simulate(ctx, &txtypes.SimulateRequest{TxBytes: simTxBytes}, grpc.Header(&header))
		if err != nil {
			err = errors.Wrap(err, "failed to CalculateGas")
			return nil, nil, err
		}

		adjustedGas := uint64(txf.GasAdjustment()*float64(simRes.GasInfo.GasUsed)) * 12 / 10
		txf = txf.WithGas(adjustedGas)

		c.gasWanted = adjustedGas
	}

	txn, err := txf.BuildUnsignedTx(msgs...)

	if err != nil {
		err = errors.Wrap(err, "failed to BuildUnsignedTx")
		return nil, nil, err
	}

	txn.SetFeeGranter(clientCtx.GetFeeGranterAddress())
	err = tx.Sign(txf, clientCtx.GetFromName(), txn, true)
	if err != nil {
		err = errors.Wrap(err, "failed to Sign Tx")
		return nil, nil, err
	}

	txBytes, err := clientCtx.TxConfig.TxEncoder()(txn.GetTx())
	if err != nil {
		err = errors.Wrap(err, "failed TxEncoder to encode Tx")
		return nil, nil, err
	}

	req := txtypes.BroadcastTxRequest{
		txBytes,
		txtypes.BroadcastMode_BROADCAST_MODE_SYNC,
	}
	// use our own client to broadcast tx
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.txClient.BroadcastTx(ctx, &req, grpc.Header(&header))
	return res, txBytes, err

}

func (c *chainClient) PollTxResults(res *txtypes.BroadcastTxResponse, clientCtx client.Context, txBytes []byte) (*txtypes.BroadcastTxResponse, error) {
	awaitCtx, cancelFn := context.WithTimeout(context.Background(), defaultBroadcastTimeout)
	defer cancelFn()

	txHash, _ := hex.DecodeString(res.TxResponse.TxHash)
	t := time.NewTimer(defaultBroadcastStatusPoll)

	for {
		// do action immediately before first ticker is triggered
		resultTx, err := clientCtx.Client.Tx(awaitCtx, txHash, false)
		if err != nil {
			if errRes := client.CheckTendermintError(err, txBytes); errRes != nil {
				return &txtypes.BroadcastTxResponse{TxResponse: errRes}, err
			} // else continue trying to get result.
		} else if resultTx.Height > 0 {
			resResultTx := sdk.NewResponseResultTx(resultTx, res.TxResponse.Tx, res.TxResponse.Timestamp)
			res = &txtypes.BroadcastTxResponse{TxResponse: resResultTx}
			t.Stop()
			// Tx failed, resync nonce.
			if res.TxResponse.Data == "" {
				c.updateNounce()
			}

			return res, err
		}

		t.Reset(defaultBroadcastStatusPoll)

		select {
		case <-awaitCtx.Done():
			err := errors.Wrapf(ErrTimedOut, "%s", res.TxResponse.TxHash)
			t.Stop()
			return nil, err
		case <-t.C:
			continue
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
		c.logger.Debugln("[INJ-GO-SDK] BroadcastTx with nonce", c.accSeq)
		res, err := c.broadcastTx(c.ctx, c.txFactory, true, toSubmit...)
		if err != nil {
			if strings.Contains(err.Error(), "account sequence mismatch") {
				c.syncNonce()
				sequence := c.getAccSeq()
				c.txFactory = c.txFactory.WithSequence(sequence)
				c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
				c.logger.Debugln("[INJ-GO-SDK] Retrying broadcastTx with nonce: ", c.accSeq)
				res, err = c.broadcastTx(c.ctx, c.txFactory, true, toSubmit...)
			}
			if err != nil {
				resJSON, _ := json.MarshalIndent(res, "", "\t")
				c.logger.Errorf("[INJ-GO-SDK] Failed to commit msg batch: %s, size: %d: %v", string(resJSON), len(toSubmit), err)
				return
			}
		}

		if res.TxResponse.Code != 0 {
			err = errors.Errorf("error %d (%s): %s", res.TxResponse.Code, res.TxResponse.Codespace, res.TxResponse.RawLog)
			c.logger.Errorf("[INJ-GO-SDK] Failed to commit msg batch, txHash: %s with err: %v", res.TxResponse.TxHash, err)
		} else {
			c.logger.Debugln("[INJ-GO-SDK] Msg batch committed successfully at height: ", res.TxResponse.Height)
		}

		c.accSeq++
		c.logger.Debugln("[INJ-GO-SDK] Nonce incremented to ", c.accSeq)
		c.logger.Debugln("[INJ-GO-SDK] Gas wanted: ", c.gasWanted)
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

func (c *chainClient) DefaultSubaccount(acc cosmtypes.AccAddress) eth.Hash {
	return c.Subaccount(acc, 0)
}

func (c *chainClient) Subaccount(account cosmtypes.AccAddress, index int) eth.Hash {
	ethAddress := eth.BytesToAddress(account.Bytes())
	ethLowerAddress := strings.ToLower(ethAddress.String())

	subaccountId := fmt.Sprintf("%s%024x", ethLowerAddress, index)
	return eth.HexToHash(subaccountId)
}

func (c *chainClient) GetSubAccountNonce(ctx context.Context, subaccountId eth.Hash) (*exchangetypes.QuerySubaccountTradeNonceResponse, error) {
	req := &exchangetypes.QuerySubaccountTradeNonceRequest{SubaccountId: subaccountId.String()}
	return c.exchangeQueryClient.SubaccountTradeNonce(ctx, req)
}

// Deprecated: Use CreateSpotOrder instead
func (c *chainClient) SpotOrder(defaultSubaccountID eth.Hash, network common.Network, d *SpotOrderData) *exchangetypes.SpotOrder {
	assistant, err := NewMarketsAssistant(network.Name)
	if err != nil {
		c.logger.Errorln("[INJ-GO-SDK] ", err)
	}

	return c.CreateSpotOrder(defaultSubaccountID, network, d, assistant)
}

func (c *chainClient) CreateSpotOrder(defaultSubaccountID eth.Hash, network common.Network, d *SpotOrderData, marketsAssistant MarketsAssistant) *exchangetypes.SpotOrder {

	market, isPresent := marketsAssistant.AllSpotMarkets()[d.MarketId]
	if !isPresent {
		c.logger.Errorln("[INJ-GO-SDK] ", errors.Errorf("Invalid spot market id for %s network (%s)", c.network.Name, d.MarketId))
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
		c.logger.Errorln("[INJ-GO-SDK] ", err)
	}

	return c.CreateDerivativeOrder(defaultSubaccountID, network, d, assistant)
}

func (c *chainClient) CreateDerivativeOrder(defaultSubaccountID eth.Hash, network common.Network, d *DerivativeOrderData, marketAssistant MarketsAssistant) *exchangetypes.DerivativeOrder {
	market, isPresent := marketAssistant.AllDerivativeMarkets()[d.MarketId]
	if !isPresent {
		c.logger.Errorln("[INJ-GO-SDK] ", errors.Errorf("Invalid derivative market id for %s network (%s)", c.network.Name, d.MarketId))
	}

	orderSize := market.QuantityToChainFormat(d.Quantity)
	orderPrice := market.PriceToChainFormat(d.Price)
	orderMargin := cosmtypes.MustNewDecFromStr("0")

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
	return c.authzQueryClient.Grants(ctx, &req)
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
		c.logger.Errorln("[INJ-GO-SDK] ", "please use BuildExchangeBatchUpdateOrdersAuthz for BatchUpdateOrdersAuthz")
	default:
		c.logger.Errorln("[INJ-GO-SDK] ", "unsupported exchange authz type")
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
		c.logger.Errorln("[INJ-GO-SDK] ", err)
	}

	if !cometbftClient.IsRunning() {
		err = cometbftClient.Start()
		if err != nil {
			c.logger.Errorln("[INJ-GO-SDK] ", err)
		}
	}
	defer func() {
		err := cometbftClient.Stop()
		if err != nil {
			c.logger.Errorln("[INJ-GO-SDK] ", err)
		}
	}()

	c.StreamEventOrderFailWithWebsocket(sender, cometbftClient, failEventCh)
}

func (c *chainClient) StreamEventOrderFailWithWebsocket(sender string, websocket *rpchttp.HTTP, failEventCh chan map[string]uint) {
	filter := fmt.Sprintf("tm.event='Tx' AND message.sender='%s' AND message.action='/injective.exchange.v1beta1.MsgBatchUpdateOrders' AND injective.exchange.v1beta1.EventOrderFail.flags EXISTS", sender)
	eventCh, err := websocket.Subscribe(context.Background(), "OrderFail", filter, 10000)
	if err != nil {
		c.logger.Errorln("[INJ-GO-SDK] ", err)
	}

	// stream and extract fail events
	for {
		e := <-eventCh

		var failedOrderHashes []string
		err = json.Unmarshal([]byte(e.Events["injective.exchange.v1beta1.EventOrderFail.hashes"][0]), &failedOrderHashes)
		if err != nil {
			c.logger.Errorln("[INJ-GO-SDK] ", err)
		}

		var failedOrderCodes []uint
		err = json.Unmarshal([]byte(e.Events["injective.exchange.v1beta1.EventOrderFail.flags"][0]), &failedOrderCodes)
		if err != nil {
			c.logger.Errorln("[INJ-GO-SDK] ", err)
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
		c.logger.Errorln("[INJ-GO-SDK] ", err)
	}

	if !cometbftClient.IsRunning() {
		err = cometbftClient.Start()
		if err != nil {
			c.logger.Errorln("[INJ-GO-SDK] ", err)
		}
	}
	defer func() {
		err := cometbftClient.Stop()
		if err != nil {
			c.logger.Errorln("[INJ-GO-SDK] ", err)
		}
	}()

	c.StreamOrderbookUpdateEventsWithWebsocket(orderbookType, marketIds, cometbftClient, orderbookCh)

}

func (c *chainClient) StreamOrderbookUpdateEventsWithWebsocket(orderbookType OrderbookType, marketIds []string, websocket *rpchttp.HTTP, orderbookCh chan exchangetypes.Orderbook) {
	filter := fmt.Sprintf("tm.event='NewBlock' AND %s EXISTS", orderbookType)
	eventCh, err := websocket.Subscribe(context.Background(), "OrderbookUpdate", filter, 10000)
	if err != nil {
		c.logger.Errorln("[INJ-GO-SDK] ", err)
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
				c.logger.Errorln("[INJ-GO-SDK] ", err)
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
	return c.txClient.GetTx(ctx, &txtypes.GetTxRequest{
		Hash: txHash,
	})
}

func (c *chainClient) ChainStream(ctx context.Context, req chainstreamtypes.StreamRequest) (chainstreamtypes.Stream_StreamClient, error) {
	ctx = c.getCookie(ctx)
	stream, err := c.chainStreamClient.Stream(ctx, &req)
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
	return c.wasmQueryClient.ContractInfo(ctx, req)
}

func (c *chainClient) FetchContractHistory(ctx context.Context, address string, pagination *query.PageRequest) (*wasmtypes.QueryContractHistoryResponse, error) {
	req := &wasmtypes.QueryContractHistoryRequest{
		Address:    address,
		Pagination: pagination,
	}
	return c.wasmQueryClient.ContractHistory(ctx, req)
}

func (c *chainClient) FetchContractsByCode(ctx context.Context, codeId uint64, pagination *query.PageRequest) (*wasmtypes.QueryContractsByCodeResponse, error) {
	req := &wasmtypes.QueryContractsByCodeRequest{
		CodeId:     codeId,
		Pagination: pagination,
	}
	return c.wasmQueryClient.ContractsByCode(ctx, req)
}

func (c *chainClient) FetchAllContractsState(ctx context.Context, address string, pagination *query.PageRequest) (*wasmtypes.QueryAllContractStateResponse, error) {
	req := &wasmtypes.QueryAllContractStateRequest{
		Address:    address,
		Pagination: pagination,
	}
	return c.wasmQueryClient.AllContractState(ctx, req)
}

func (c *chainClient) RawContractState(
	ctx context.Context,
	contractAddress string,
	queryData []byte,
) (*wasmtypes.QueryRawContractStateResponse, error) {
	return c.wasmQueryClient.RawContractState(
		ctx,
		&wasmtypes.QueryRawContractStateRequest{
			Address:   contractAddress,
			QueryData: queryData,
		},
	)
}

func (c *chainClient) SmartContractState(
	ctx context.Context,
	contractAddress string,
	queryData []byte,
) (*wasmtypes.QuerySmartContractStateResponse, error) {
	return c.wasmQueryClient.SmartContractState(
		ctx,
		&wasmtypes.QuerySmartContractStateRequest{
			Address:   contractAddress,
			QueryData: queryData,
		},
	)
}

func (c *chainClient) FetchCode(ctx context.Context, codeId uint64) (*wasmtypes.QueryCodeResponse, error) {
	req := &wasmtypes.QueryCodeRequest{
		CodeId: codeId,
	}
	return c.wasmQueryClient.Code(ctx, req)
}

func (c *chainClient) FetchCodes(ctx context.Context, pagination *query.PageRequest) (*wasmtypes.QueryCodesResponse, error) {
	req := &wasmtypes.QueryCodesRequest{
		Pagination: pagination,
	}
	return c.wasmQueryClient.Codes(ctx, req)
}

func (c *chainClient) FetchPinnedCodes(ctx context.Context, pagination *query.PageRequest) (*wasmtypes.QueryPinnedCodesResponse, error) {
	req := &wasmtypes.QueryPinnedCodesRequest{
		Pagination: pagination,
	}
	return c.wasmQueryClient.PinnedCodes(ctx, req)
}

func (c *chainClient) FetchContractsByCreator(ctx context.Context, creator string, pagination *query.PageRequest) (*wasmtypes.QueryContractsByCreatorResponse, error) {
	req := &wasmtypes.QueryContractsByCreatorRequest{
		CreatorAddress: creator,
		Pagination:     pagination,
	}
	return c.wasmQueryClient.ContractsByCreator(ctx, req)
}

// Tokenfactory module

func (c *chainClient) FetchDenomAuthorityMetadata(ctx context.Context, creator string, subDenom string) (*tokenfactorytypes.QueryDenomAuthorityMetadataResponse, error) {
	req := &tokenfactorytypes.QueryDenomAuthorityMetadataRequest{
		Creator: creator,
	}

	if subDenom != "" {
		req.SubDenom = subDenom
	}

	return c.tokenfactoryQueryClient.DenomAuthorityMetadata(ctx, req)
}

func (c *chainClient) FetchDenomsFromCreator(ctx context.Context, creator string) (*tokenfactorytypes.QueryDenomsFromCreatorResponse, error) {
	req := &tokenfactorytypes.QueryDenomsFromCreatorRequest{
		Creator: creator,
	}

	return c.tokenfactoryQueryClient.DenomsFromCreator(ctx, req)
}

func (c *chainClient) FetchTokenfactoryModuleState(ctx context.Context) (*tokenfactorytypes.QueryModuleStateResponse, error) {
	req := &tokenfactorytypes.QueryModuleStateRequest{}

	return c.tokenfactoryQueryClient.TokenfactoryModuleState(ctx, req)
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

func (c *chainClient) AdjustGasPricesToMax() {
	c.txFactory = c.txFactory.WithGasPrices(MaxGasFee)
}

func (c *chainClient) AdjustedGasPricesToOrigin() {
	c.txFactory = c.txFactory.WithGasPrices(c.opts.GasPrices)
}
