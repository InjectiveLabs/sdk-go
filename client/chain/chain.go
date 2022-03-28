package chain

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"math/big"
	common "github.com/InjectiveLabs/sdk-go/client/common"
	chaintypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	eth "github.com/ethereum/go-ethereum/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"github.com/shopspring/decimal"
	"strings"
	"sync"
	"fmt"
	"sync/atomic"
	"time"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	log "github.com/xlab/suplog"
)

const (
	msgCommitBatchSizeLimit = 1024
	msgCommitBatchTimeLimit = 500 * time.Millisecond
	defaultBroadcastStatusPoll = 100 * time.Millisecond
	defaultBroadcastTimeout    = 40 * time.Second
)

var (
	ErrTimedOut = errors.New("tx timed out")
	ErrQueueClosed    = errors.New("queue is closed")
	ErrEnqueueTimeout = errors.New("enqueue timeout")
	ErrReadOnly       = errors.New("client is in read-only mode")
)

type ChainClient interface {
	CanSignTransactions() bool
	FromAddress() sdk.AccAddress
	QueryClient() *grpc.ClientConn
	ClientContext() client.Context

	SyncBroadcastMsg(msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error)
	AsyncBroadcastMsg(msgs ...sdk.Msg) (*txtypes.BroadcastTxResponse, error)
	QueueBroadcastMsg(msgs ...sdk.Msg) error

	GetBankBalances(ctx context.Context, address string) (*banktypes.QueryAllBalancesResponse, error)

	DefaultSubaccount(acc cosmtypes.AccAddress) eth.Hash
	GetSpotQuantity(value decimal.Decimal, minTickSize cosmtypes.Dec, baseDecimals int) (qty cosmtypes.Dec)
	GetSpotPrice(price decimal.Decimal, baseDecimals int, quoteDecimals int, minPriceTickSize cosmtypes.Dec) cosmtypes.Dec
	GetDerivativeQuantity(value decimal.Decimal, minTickSize cosmtypes.Dec) (qty cosmtypes.Dec)
	GetDerivativePrice(value, tickSize cosmtypes.Dec) cosmtypes.Dec

	SpotOrder(defaultSubaccountID eth.Hash, d *SpotOrderData) *exchangetypes.SpotOrder
	DerivativeOrder(defaultSubaccountID eth.Hash, d *DerivativeOrderData) *exchangetypes.DerivativeOrder
	OrderCancel(defaultSubaccountID eth.Hash, d *OrderCancelData) *exchangetypes.OrderData

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

	accNum uint64
	accSeq uint64

	sessionCookie string

	txClient txtypes.ServiceClient
	authQueryClient authtypes.QueryClient
	chainQueryClient chaintypes.QueryClient
	bankQueryClient  banktypes.QueryClient

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

	txFactory := NewTxFactory(ctx)
	if len(opts.GasPrices) > 0 {
		txFactory = txFactory.WithGasPrices(opts.GasPrices)
	}

	var conn *grpc.ClientConn
	var err error
	if opts.TLSCert != nil {
		conn, err = grpc.Dial(protoAddr, grpc.WithTransportCredentials(opts.TLSCert), grpc.WithContextDialer(common.DialerFunc))
	} else {
		conn, err = grpc.Dial(protoAddr, grpc.WithInsecure(), grpc.WithContextDialer(common.DialerFunc))
	}
	if err != nil {
		err := errors.Wrapf(err, "failed to connect to the gRPC: %s", protoAddr)
		return nil, err
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

		txClient: txtypes.NewServiceClient(conn),
		authQueryClient: authtypes.NewQueryClient(conn),
		chainQueryClient: chaintypes.NewQueryClient(conn),
		bankQueryClient:  banktypes.NewQueryClient(conn),
	}

	if cc.canSign {
		var err error

		cc.accNum, cc.accSeq, err = cc.txFactory.AccountRetriever().GetAccountNumberSequence(ctx, ctx.GetFromAddress())
		if err != nil {
			err = errors.Wrap(err, "failed to get initial account num and seq")
			return nil, err
		}

		go cc.runBatchBroadcast()
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
	md := metadata.Get("set-cookie")
	if len(md) > 0 {
		c.sessionCookie = md[0]
	}
}

func (c *chainClient) getCookie(ctx context.Context) context.Context {
	md := metadata.Pairs("cookie",c.sessionCookie)
	return metadata.NewOutgoingContext(ctx, md)
	//return metadata.AppendToOutgoingContext(ctx, "cookie", c.sessionCookie)
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

//AsyncBroadcastMsg sends Tx to chain and doesn't wait until Tx is included in block. This method
//cannot be used for rapid Tx sending, it is expected that you wait for transaction status with
//external tools. If you want sdk to wait for it, use SyncBroadcastMsg.
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

	// attempt to load cookie
	ctx := context.Background()

	if txf.SimulateAndExecute() || clientCtx.Simulate {
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
		c.setCookie(header)

		adjustedGas := uint64(txf.GasAdjustment() * float64(simRes.GasInfo.GasUsed))
		txf = txf.WithGas(adjustedGas)
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
	c.setCookie(header)

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

//QueueBroadcastMsg enqueues a list of messages. Messages will added to the queue
//and grouped into Txns in chunks. Use this method to mass broadcast Txns with efficiency.
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
			log.WithField("txHash", res.TxResponse.TxHash).Debugln("msg batch committed successfully")
		}

		c.accSeq++
		log.Debugln("nonce incremented to", c.accSeq)
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

func (c *chainClient) DefaultSubaccount(acc cosmtypes.AccAddress) eth.Hash {
	return eth.BytesToHash(eth.RightPadBytes(acc.Bytes(), 32))
}


func formatPriceToTickSize(value, tickSize cosmtypes.Dec) cosmtypes.Dec {
    residue := new(big.Int).Mod(value.BigInt(), tickSize.BigInt())
    formattedValue := new(big.Int).Sub(value.BigInt(), residue)
    p := decimal.NewFromBigInt(formattedValue, -18).StringFixed(18)
    realValue, _ := cosmtypes.NewDecFromStr(p)
    return realValue
}


func (c *chainClient) GetSpotQuantity(value decimal.Decimal, minTickSize cosmtypes.Dec, baseDecimals int) (qty cosmtypes.Dec) {
	mid, _ := cosmtypes.NewDecFromStr(value.String())
	bStr := decimal.New(1, int32(baseDecimals)).String()
	baseDec, _ := cosmtypes.NewDecFromStr(bStr)
	scale := baseDec.Quo(minTickSize)
	midScaledInt := mid.Mul(scale).TruncateDec()
	qty = minTickSize.Mul(midScaledInt)
	return qty
}

func (c *chainClient) GetSpotPrice(price decimal.Decimal, baseDecimals int, quoteDecimals int, minPriceTickSize cosmtypes.Dec) cosmtypes.Dec {
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

func (c *chainClient) GetDerivativeQuantity(value decimal.Decimal, minTickSize cosmtypes.Dec) (qty cosmtypes.Dec) {
	mid := cosmtypes.MustNewDecFromStr(value.StringFixed(18))
	baseDec := cosmtypes.OneDec()
	scale := baseDec.Quo(minTickSize)
	midScaledInt := mid.Mul(scale).TruncateDec()
	qty = minTickSize.Mul(midScaledInt)
	return qty
}

func (c *chainClient) GetDerivativePrice(value, tickSize cosmtypes.Dec) cosmtypes.Dec {
	residue := new(big.Int).Mod(value.BigInt(), tickSize.BigInt())
	formattedValue := new(big.Int).Sub(value.BigInt(), residue)
	p := decimal.NewFromBigInt(formattedValue, -18).String()
	realValue, _ := cosmtypes.NewDecFromStr(p)
	return realValue
}


func (c *chainClient) SpotOrder(defaultSubaccountID eth.Hash, d *SpotOrderData) *exchangetypes.SpotOrder {
	return &exchangetypes.SpotOrder{
		MarketId:  d.MarketId,
		OrderType: d.OrderType,
		OrderInfo: exchangetypes.OrderInfo{
			SubaccountId: defaultSubaccountID.Hex(),
			FeeRecipient: d.FeeRecipient,
			Price:        d.Price,
			Quantity:     d.Quantity,
		},
	}
}

func (c *chainClient) DerivativeOrder(defaultSubaccountID eth.Hash, d *DerivativeOrderData) *exchangetypes.DerivativeOrder {
	return &exchangetypes.DerivativeOrder{
		MarketId:  d.MarketId,
		OrderType: d.OrderType,
		Margin: d.Margin,
		OrderInfo: exchangetypes.OrderInfo{
			SubaccountId: defaultSubaccountID.Hex(),
			FeeRecipient: d.FeeRecipient,
			Price:        d.Price,
			Quantity:     d.Quantity,
		},
	}
}


func (c *chainClient) OrderCancel(defaultSubaccountID eth.Hash, d *OrderCancelData) *exchangetypes.OrderData {
	return &exchangetypes.OrderData{
		MarketId:  d.MarketId,
		OrderHash: d.OrderHash,
		SubaccountId: defaultSubaccountID.Hex(),
	}
}


type DerivativeOrderData struct {
	OrderType    exchangetypes.OrderType
	Price        cosmtypes.Dec
	Margin     	 cosmtypes.Dec
	Quantity     cosmtypes.Dec
	Leverage	 cosmtypes.Dec
	FeeRecipient string
	MarketId string
}


type SpotOrderData struct {
	OrderType    exchangetypes.OrderType
	Price        cosmtypes.Dec
	Quantity     cosmtypes.Dec
	FeeRecipient string
	MarketId string
}


type OrderCancelData struct {
	MarketId string
	OrderHash string
}
