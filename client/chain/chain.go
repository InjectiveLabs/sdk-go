package chain

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	"github.com/InjectiveLabs/sdk-go/client/common"
	log "github.com/InjectiveLabs/suplog"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	eth "github.com/ethereum/go-ethereum/common"
	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

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
	defaultSessionRenewalOffset      = 120
	defaultBlockTime                 = 3 * time.Second
	defaultChainCookieName           = ".chain_cookie"
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

	GetBankBalances(ctx context.Context, address string) (*banktypes.QueryAllBalancesResponse, error)
	GetBankBalance(ctx context.Context, address string, denom string) (*banktypes.QueryBalanceResponse, error)
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

	GetSubAccountNonce(ctx context.Context, subaccountId eth.Hash) (*exchangetypes.QuerySubaccountTradeNonceResponse, error)
	GetFeeDiscountInfo(ctx context.Context, account string) (*exchangetypes.QueryFeeDiscountAccountInfoResponse, error)
	ComputeOrderHashes(spotOrders []exchangetypes.SpotOrder, derivativeOrders []exchangetypes.DerivativeOrder) (OrderHashes, error)

	SpotOrder(defaultSubaccountID eth.Hash, network common.Network, d *SpotOrderData) *exchangetypes.SpotOrder
	DerivativeOrder(defaultSubaccountID eth.Hash, network common.Network, d *DerivativeOrderData) *exchangetypes.DerivativeOrder
	OrderCancel(defaultSubaccountID eth.Hash, d *OrderCancelData) *exchangetypes.OrderData

	SmartContractState(
		ctx context.Context,
		contractAddress string,
		queryData []byte,
	) (*wasmtypes.QuerySmartContractStateResponse, error)
	RawContractState(
		ctx context.Context,
		contractAddress string,
		queryData []byte,
	) (*wasmtypes.QueryRawContractStateResponse, error)

	GetGasFee() (string, error)

	StreamEventOrderFail(sender string, failEventCh chan map[string]uint)
	StreamOrderbookUpdateEvents(orderbookType OrderbookType, marketIds []string, orderbookCh chan exchangetypes.Orderbook)

	Close()
}

type chainClient struct {
	ctx       client.Context
	opts      *common.ClientOptions
	logger    log.Logger
	conn      *grpc.ClientConn
	txFactory tx.Factory

	fromAddress sdk.AccAddress
	doneC       chan bool
	msgC        chan sdk.Msg
	syncMux     *sync.Mutex

	accNum    uint64
	accSeq    uint64
	gasWanted uint64
	gasFee    string

	sessionCookie  string
	sessionEnabled bool

	txClient            txtypes.ServiceClient
	authQueryClient     authtypes.QueryClient
	exchangeQueryClient exchangetypes.QueryClient
	bankQueryClient     banktypes.QueryClient
	authzQueryClient    authztypes.QueryClient
	wasmQueryClient     wasmtypes.QueryClient

	closed  int64
	canSign bool
}

// NewCosmosClient creates a new gRPC client that communicates with gRPC server at protoAddr.
// protoAddr must be in form "tcp://127.0.0.1:8080" or "unix:///tmp/test.sock", protocol is required.
func NewChainClient(
	ctx client.Context,
	protoAddr string,
	options ...common.ClientOption,
) (ChainClient, error) {
	// process options
	opts := common.DefaultClientOptions()
	for _, opt := range options {
		if err := opt(opts); err != nil {
			err = errors.Wrap(err, "error in client option")
			return nil, err
		}
	}

	// init tx factory
	txFactory := NewTxFactory(ctx)
	if len(opts.GasPrices) > 0 {
		txFactory = txFactory.WithGasPrices(opts.GasPrices)
	}

	// init grpc connection
	var conn *grpc.ClientConn
	var err error
	stickySessionEnabled := true
	if opts.TLSCert != nil {
		conn, err = grpc.Dial(protoAddr, grpc.WithTransportCredentials(opts.TLSCert), grpc.WithContextDialer(common.DialerFunc))
	} else {
		conn, err = grpc.Dial(protoAddr, grpc.WithInsecure(), grpc.WithContextDialer(common.DialerFunc))
		stickySessionEnabled = false
	}
	if err != nil {
		err = errors.Wrapf(err, "failed to connect to the gRPC: %s", protoAddr)
		return nil, err
	}

	// init tm websocket
	if ctx.Client != nil {
		err = ctx.Client.Start()
		if err != nil {
			panic(err)
		}
	}

	// build client
	cc := &chainClient{
		ctx:  ctx,
		opts: opts,

		logger: log.WithFields(log.Fields{
			"module": "sdk-go",
			"svc":    "chainClient",
		}),

		conn:      conn,
		txFactory: txFactory,
		canSign:   ctx.Keyring != nil,
		syncMux:   new(sync.Mutex),
		msgC:      make(chan sdk.Msg, msgCommitBatchSizeLimit),
		doneC:     make(chan bool, 1),

		sessionEnabled: stickySessionEnabled,

		txClient:            txtypes.NewServiceClient(conn),
		authQueryClient:     authtypes.NewQueryClient(conn),
		exchangeQueryClient: exchangetypes.NewQueryClient(conn),
		bankQueryClient:     banktypes.NewQueryClient(conn),
		authzQueryClient:    authztypes.NewQueryClient(conn),
		wasmQueryClient:     wasmtypes.NewQueryClient(conn),
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

	// create file if not exist
	os.OpenFile(defaultChainCookieName, os.O_RDONLY|os.O_CREATE, 0666)

	// attempt to load from disk
	data, err := os.ReadFile(defaultChainCookieName)
	if err != nil {
		cc.logger.Errorln(err)
	} else {
		cc.sessionCookie = string(data)
		cc.logger.Infoln("chain session cookie loaded from disk")
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
	for {
		ctx := context.Background()
		block, err := c.ctx.Client.Block(ctx, nil)
		if err != nil {
			c.logger.WithError(err).Errorln("failed to get current block")
			return
		}
		c.txFactory.WithTimeoutHeight(uint64(block.Block.Height) + defaultTimeoutHeight)
		time.Sleep(defaultTimeoutHeightSyncInterval)
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

func (c *chainClient) setCookie(metadata metadata.MD) {
	if !c.sessionEnabled {
		return
	}
	md := metadata.Get("set-cookie")
	if len(md) > 0 {
		// write to client instance
		c.sessionCookie = md[0]
		// write to disk
		err := os.WriteFile(defaultChainCookieName, []byte(md[0]), 0644)
		if err != nil {
			c.logger.Errorln(err)
			return
		}
		c.logger.Infoln("chain session cookie saved to disk")
	}
}

func (c *chainClient) fetchCookie(ctx context.Context) context.Context {
	var header metadata.MD
	c.txClient.GetTx(context.Background(), &txtypes.GetTxRequest{}, grpc.Header(&header))
	c.setCookie(header)
	time.Sleep(defaultBlockTime)
	return metadata.NewOutgoingContext(ctx, metadata.Pairs("cookie", c.sessionCookie))
}

func cookieByName(cookies []*http.Cookie, key string) *http.Cookie {
	for _, c := range cookies {
		if c.Name == key {
			return c
		}
	}
	return nil
}

func (c *chainClient) getCookieExpirationTime(cookies []*http.Cookie) (time.Time, error) {
	var expiresAt string
	if cookieByName(cookies, "GCLB") != nil {
		// parse global load balance cookie timestamp
		cookie := cookieByName(cookies, "expires")
		expiresAt = strings.Replace(cookie.Value, "-", " ", -1)
	} else {
		cookie := cookieByName(cookies, "Expires")
		expiresAt = strings.Replace(cookie.Value, "-", " ", -1)
		yyyy := fmt.Sprintf("20%s", expiresAt[12:14])
		expiresAt = expiresAt[:12] + yyyy + expiresAt[14:]
	}

	return time.Parse(time.RFC1123, expiresAt)
}

func (c *chainClient) getCookie(ctx context.Context) context.Context {
	md := metadata.Pairs("cookie", c.sessionCookie)
	if !c.sessionEnabled {
		return metadata.NewOutgoingContext(ctx, md)
	}

	// borrow http request to parse cookie
	header := http.Header{}
	header.Add("Cookie", c.sessionCookie)
	request := http.Request{Header: header}
	cookies := request.Cookies()

	if len(cookies) > 0 {
		// parse expire field into unix timestamp
		expiresTimestamp, err := c.getCookieExpirationTime(cookies)
		if err != nil {
			panic(err)
		}

		// renew session if timestamp diff < offset
		timestampDiff := expiresTimestamp.Unix() - time.Now().Unix()
		if timestampDiff < defaultSessionRenewalOffset {
			return c.fetchCookie(ctx)
		}
	} else {
		return c.fetchCookie(ctx)
	}

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
	if !c.canSign {
		return
	}
	if atomic.CompareAndSwapInt64(&c.closed, 0, 1) {
		close(c.msgC)
	}
	<-c.doneC
	if c.conn != nil {
		c.conn.Close()
	}
}

func (c *chainClient) GetBankBalances(ctx context.Context, address string) (*banktypes.QueryAllBalancesResponse, error) {
	req := &banktypes.QueryAllBalancesRequest{
		Address: address,
	}
	return c.bankQueryClient.AllBalances(ctx, req)
}

func (c *chainClient) GetAccount(ctx context.Context, address string) (*authtypes.QueryAccountResponse, error) {
	req := &authtypes.QueryAccountRequest{
		Address: address,
	}
	return c.authQueryClient.Account(ctx, req)
}

func (c *chainClient) GetBankBalance(ctx context.Context, address string, denom string) (*banktypes.QueryBalanceResponse, error) {
	req := &banktypes.QueryBalanceRequest{
		Address: address,
		Denom:   denom,
	}
	return c.bankQueryClient.Balance(ctx, req)
}

// SyncBroadcastMsg sends Tx to chain and waits until Tx is included in block.
func (c *chainClient) SyncBroadcastMsg(msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error) {
	c.syncMux.Lock()
	defer c.syncMux.Unlock()

	c.txFactory = c.txFactory.WithSequence(c.accSeq)
	c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
	res, err := c.broadcastTx(c.ctx, c.txFactory, true, msgs...)

	if err != nil {
		if strings.Contains(err.Error(), "account sequence mismatch") {
			c.syncNonce()
			c.txFactory = c.txFactory.WithSequence(c.accSeq)
			c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
			log.Debugln("retrying broadcastTx with nonce", c.accSeq)
			res, err = c.broadcastTx(c.ctx, c.txFactory, true, msgs...)
		}
		if err != nil {
			resJSON, _ := json.MarshalIndent(res, "", "\t")
			c.logger.WithField("size", len(msgs)).WithError(err).Errorln("failed to commit msg batch:", string(resJSON))
			return nil, err
		}
	}

	c.accSeq++

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

	simTxBytes, err := tx.BuildSimTx(txf, msgs...)
	if err != nil {
		err = errors.Wrap(err, "failed to build sim tx bytes")
		return nil, err
	}

	ctx := context.Background()
	ctx = c.getCookie(ctx)
	var header metadata.MD
	simRes, err := c.txClient.Simulate(ctx, &txtypes.SimulateRequest{TxBytes: simTxBytes}, grpc.Header(&header))
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

	c.txFactory = c.txFactory.WithSequence(c.accSeq)
	c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
	res, err := c.broadcastTx(c.ctx, c.txFactory, false, msgs...)
	if err != nil {
		if strings.Contains(err.Error(), "account sequence mismatch") {
			c.syncNonce()
			c.txFactory = c.txFactory.WithSequence(c.accSeq)
			c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
			log.Debugln("retrying broadcastTx with nonce", c.accSeq)
			res, err = c.broadcastTx(c.ctx, c.txFactory, false, msgs...)
		}
		if err != nil {
			resJSON, _ := json.MarshalIndent(res, "", "\t")
			c.logger.WithField("size", len(msgs)).WithError(err).Errorln("failed to commit msg batch:", string(resJSON))
			return nil, err
		}
	}

	c.accSeq++

	return res, nil
}

func (c *chainClient) BuildSignedTx(clientCtx client.Context, accNum, accSeq, initialGas uint64, msgs ...sdk.Msg) ([]byte, error) {
	txf := NewTxFactory(clientCtx).WithSequence(accSeq).WithAccountNumber(accNum).WithGas(initialGas)

	if clientCtx.Simulate {
		simTxBytes, err := tx.BuildSimTx(txf, msgs...)
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

		adjustedGas := uint64(txf.GasAdjustment() * float64(simRes.GasInfo.GasUsed))
		txf = txf.WithGas(adjustedGas)

		c.gasWanted = adjustedGas
	}

	txf, err := c.prepareFactory(clientCtx, txf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepareFactory")
	}

	txn, err := tx.BuildUnsignedTx(txf, msgs...)
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
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.txClient.BroadcastTx(ctx, &req, grpc.Header(&header))
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

func (c *chainClient) AsyncBroadcastSignedTx(txBytes []byte) (*txtypes.BroadcastTxResponse, error) {
	req := txtypes.BroadcastTxRequest{
		TxBytes: txBytes,
		Mode:    txtypes.BroadcastMode_BROADCAST_MODE_SYNC,
	}

	ctx := context.Background()
	// use our own client to broadcast tx
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.txClient.BroadcastTx(ctx, &req, grpc.Header(&header))
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
		simTxBytes, err := tx.BuildSimTx(txf, msgs...)
		if err != nil {
			err = errors.Wrap(err, "failed to build sim tx bytes")
			return nil, err
		}
		ctx := c.getCookie(ctx)
		var header metadata.MD
		simRes, err := c.txClient.Simulate(ctx, &txtypes.SimulateRequest{TxBytes: simTxBytes}, grpc.Header(&header))
		if err != nil {
			err = errors.Wrap(err, "failed to CalculateGas")
			return nil, err
		}

		adjustedGas := uint64(txf.GasAdjustment() * float64(simRes.GasInfo.GasUsed))
		txf = txf.WithGas(adjustedGas)

		c.gasWanted = adjustedGas
	}

	txn, err := tx.BuildUnsignedTx(txf, msgs...)

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
		txBytes,
		txtypes.BroadcastMode_BROADCAST_MODE_SYNC,
	}
	// use our own client to broadcast tx
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.txClient.BroadcastTx(ctx, &req, grpc.Header(&header))
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
		c.txFactory = c.txFactory.WithSequence(c.accSeq)
		c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
		log.Debugln("broadcastTx with nonce", c.accSeq)
		res, err := c.broadcastTx(c.ctx, c.txFactory, true, toSubmit...)
		if err != nil {
			if strings.Contains(err.Error(), "account sequence mismatch") {
				c.syncNonce()
				c.txFactory = c.txFactory.WithSequence(c.accSeq)
				c.txFactory = c.txFactory.WithAccountNumber(c.accNum)
				log.Debugln("retrying broadcastTx with nonce", c.accSeq)
				res, err = c.broadcastTx(c.ctx, c.txFactory, true, toSubmit...)
			}
			if err != nil {
				resJSON, _ := json.MarshalIndent(res, "", "\t")
				c.logger.WithField("size", len(toSubmit)).WithError(err).Errorln("failed to commit msg batch:", string(resJSON))
				return
			}
		}

		if res.TxResponse.Code != 0 {
			err = errors.Errorf("error %d (%s): %s", res.TxResponse.Code, res.TxResponse.Codespace, res.TxResponse.RawLog)
			log.WithField("txHash", res.TxResponse.TxHash).WithError(err).Errorln("failed to commit msg batch")
		} else {
			log.WithField("txHash", res.TxResponse.TxHash).Debugln("msg batch committed successfully at height", res.TxResponse.Height)
		}

		c.accSeq++
		log.Debugln("nonce incremented to", c.accSeq)
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

func (c *chainClient) DefaultSubaccount(acc cosmtypes.AccAddress) eth.Hash {
	return eth.BytesToHash(eth.RightPadBytes(acc.Bytes(), 32))
}

func (c *chainClient) GetSubAccountNonce(ctx context.Context, subaccountId eth.Hash) (*exchangetypes.QuerySubaccountTradeNonceResponse, error) {
	req := &exchangetypes.QuerySubaccountTradeNonceRequest{SubaccountId: subaccountId.String()}
	return c.exchangeQueryClient.SubaccountTradeNonce(ctx, req)
}

func formatPriceToTickSize(value, tickSize cosmtypes.Dec) cosmtypes.Dec {
	residue := new(big.Int).Mod(value.BigInt(), tickSize.BigInt())
	formattedValue := new(big.Int).Sub(value.BigInt(), residue)
	p := decimal.NewFromBigInt(formattedValue, -18).StringFixed(18)
	realValue, _ := cosmtypes.NewDecFromStr(p)
	return realValue
}

func GetSpotQuantity(value decimal.Decimal, minTickSize cosmtypes.Dec, baseDecimals int) (qty cosmtypes.Dec) {
	mid, _ := cosmtypes.NewDecFromStr(value.String())
	bStr := decimal.New(1, int32(baseDecimals)).String()
	baseDec, _ := cosmtypes.NewDecFromStr(bStr)
	scale := baseDec.Quo(minTickSize)
	midScaledInt := mid.Mul(scale).TruncateDec()
	qty = minTickSize.Mul(midScaledInt)
	return qty
}

func GetSpotPrice(price decimal.Decimal, baseDecimals int, quoteDecimals int, minPriceTickSize cosmtypes.Dec) cosmtypes.Dec {
	scale := decimal.New(1, int32(quoteDecimals-baseDecimals))
	priceStr := scale.Mul(price).StringFixed(18)
	decPrice, err := cosmtypes.NewDecFromStr(priceStr)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(priceStr, scale.String(), price.String())
		fmt.Println(decPrice.String())
	}
	realPrice := formatPriceToTickSize(decPrice, minPriceTickSize)
	return realPrice
}

func GetDerivativeQuantity(value decimal.Decimal, minTickSize cosmtypes.Dec) (qty cosmtypes.Dec) {
	mid := cosmtypes.MustNewDecFromStr(value.StringFixed(18))
	baseDec := cosmtypes.OneDec()
	scale := baseDec.Quo(minTickSize)
	midScaledInt := mid.Mul(scale).TruncateDec()
	qty = minTickSize.Mul(midScaledInt)
	return qty
}

func GetDerivativePrice(value, tickSize cosmtypes.Dec) cosmtypes.Dec {
	residue := new(big.Int).Mod(value.BigInt(), tickSize.BigInt())
	formattedValue := new(big.Int).Sub(value.BigInt(), residue)
	p := decimal.NewFromBigInt(formattedValue, -18).String()
	realValue, _ := cosmtypes.NewDecFromStr(p)
	return realValue
}

func (c *chainClient) SpotOrder(defaultSubaccountID eth.Hash, network common.Network, d *SpotOrderData) *exchangetypes.SpotOrder {

	baseDecimals := common.LoadMetadata(network, d.MarketId).Base
	quoteDecimals := common.LoadMetadata(network, d.MarketId).Quote
	minPriceTickSize := common.LoadMetadata(network, d.MarketId).MinPriceTickSize
	minQuantityTickSize := common.LoadMetadata(network, d.MarketId).MinQuantityTickSize

	orderSize := GetSpotQuantity(d.Quantity, cosmtypes.MustNewDecFromStr(strconv.FormatFloat(minQuantityTickSize, 'f', -1, 64)), baseDecimals)
	orderPrice := GetSpotPrice(d.Price, baseDecimals, quoteDecimals, cosmtypes.MustNewDecFromStr(strconv.FormatFloat(minPriceTickSize, 'f', -1, 64)))

	return &exchangetypes.SpotOrder{
		MarketId:  d.MarketId,
		OrderType: d.OrderType,
		OrderInfo: exchangetypes.OrderInfo{
			SubaccountId: defaultSubaccountID.Hex(),
			FeeRecipient: d.FeeRecipient,
			Price:        orderPrice,
			Quantity:     orderSize,
		},
	}
}

func (c *chainClient) DerivativeOrder(defaultSubaccountID eth.Hash, network common.Network, d *DerivativeOrderData) *exchangetypes.DerivativeOrder {

	margin := cosmtypes.MustNewDecFromStr(fmt.Sprint(d.Quantity)).Mul(d.Price).Quo(d.Leverage)

	if d.IsReduceOnly == true {
		margin = cosmtypes.MustNewDecFromStr("0")
	}

	minPriceTickSize := common.LoadMetadata(network, d.MarketId).MinPriceTickSize
	minQuantityTickSize := common.LoadMetadata(network, d.MarketId).MinQuantityTickSize

	orderSize := GetDerivativeQuantity(d.Quantity, cosmtypes.MustNewDecFromStr(strconv.FormatFloat(minQuantityTickSize, 'f', -1, 64)))
	orderPrice := GetDerivativePrice(d.Price, cosmtypes.MustNewDecFromStr(strconv.FormatFloat(minPriceTickSize, 'f', -1, 64)))
	orderMargin := GetDerivativePrice(margin, cosmtypes.MustNewDecFromStr(strconv.FormatFloat(minPriceTickSize, 'f', -1, 64)))

	return &exchangetypes.DerivativeOrder{
		MarketId:  d.MarketId,
		OrderType: d.OrderType,
		Margin:    orderMargin,
		OrderInfo: exchangetypes.OrderInfo{
			SubaccountId: defaultSubaccountID.Hex(),
			FeeRecipient: d.FeeRecipient,
			Price:        orderPrice,
			Quantity:     orderSize,
		},
	}
}

func (c *chainClient) OrderCancel(defaultSubaccountID eth.Hash, d *OrderCancelData) *exchangetypes.OrderData {
	return &exchangetypes.OrderData{
		MarketId:     d.MarketId,
		OrderHash:    d.OrderHash,
		SubaccountId: defaultSubaccountID.Hex(),
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
			Expiration:    expireIn,
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
			Expiration:    expireIn,
		},
	}
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
			Expiration:    expireIn,
		},
	}
}

func (c *chainClient) StreamEventOrderFail(sender string, failEventCh chan map[string]uint) {
	filter := fmt.Sprintf("tm.event='Tx' AND message.sender='%s' AND message.action='/injective.exchange.v1beta1.MsgBatchUpdateOrders' AND injective.exchange.v1beta1.EventOrderFail.flags EXISTS", sender)
	eventCh, err := c.ctx.Client.Subscribe(context.Background(), "OrderFail", filter, 10000)
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
	filter := fmt.Sprintf("tm.event='NewBlock' AND %s EXISTS", orderbookType)
	eventCh, err := c.ctx.Client.Subscribe(context.Background(), "OrderbookUpdate", filter, 10000)
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

type DerivativeOrderData struct {
	OrderType    exchangetypes.OrderType
	Price        cosmtypes.Dec
	Quantity     decimal.Decimal
	Leverage     cosmtypes.Dec
	FeeRecipient string
	MarketId     string
	IsReduceOnly bool
}

type SpotOrderData struct {
	OrderType    exchangetypes.OrderType
	Price        decimal.Decimal
	Quantity     decimal.Decimal
	FeeRecipient string
	MarketId     string
}

type OrderCancelData struct {
	MarketId  string
	OrderHash string
}
