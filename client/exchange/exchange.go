package exchange

import (
	"context"
	"fmt"
	"github.com/InjectiveLabs/sdk-go/client/common"
	accountPB "github.com/InjectiveLabs/sdk-go/exchange/accounts_rpc/pb"
	auctionPB "github.com/InjectiveLabs/sdk-go/exchange/auction_rpc/pb"
	derivativeExchangePB "github.com/InjectiveLabs/sdk-go/exchange/derivative_exchange_rpc/pb"
	insurancePB "github.com/InjectiveLabs/sdk-go/exchange/insurance_rpc/pb"
	metaPB "github.com/InjectiveLabs/sdk-go/exchange/meta_rpc/pb"
	oraclePB "github.com/InjectiveLabs/sdk-go/exchange/oracle_rpc/pb"
	spotExchangePB "github.com/InjectiveLabs/sdk-go/exchange/spot_exchange_rpc/pb"
	"google.golang.org/grpc/metadata"

	"github.com/pkg/errors"
	log "github.com/xlab/suplog"
	"google.golang.org/grpc"
)

type ExchangeClient interface {
	QueryClient() *grpc.ClientConn
	GetMarket(ctx context.Context, marketId string) (derivativeExchangePB.MarketResponse, error)
	GetOrderbook(ctx context.Context, marketId string) (derivativeExchangePB.OrderbookResponse, error)
	GetOrderbooks(ctx context.Context, marketIds []string) (derivativeExchangePB.OrderbooksResponse, error)
	StreamOrderbook(ctx context.Context, marketIds []string) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamOrderbookClient, error)
	StreamMarket(ctx context.Context, marketIds []string) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamMarketClient, error)
	GetOrders(ctx context.Context, req derivativeExchangePB.OrdersRequest) (derivativeExchangePB.OrdersResponse, error)
	GetMarkets(ctx context.Context, req derivativeExchangePB.MarketsRequest) (derivativeExchangePB.MarketsResponse, error)
	GetPositions(ctx context.Context, req derivativeExchangePB.PositionsRequest) (derivativeExchangePB.PositionsResponse, error)
	StreamPositions(ctx context.Context, req derivativeExchangePB.StreamPositionsRequest) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamPositionsClient, error)
	StreamOrders(ctx context.Context, req derivativeExchangePB.StreamOrdersRequest) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamOrdersClient, error)
	GetTrades(ctx context.Context, req derivativeExchangePB.TradesRequest) (derivativeExchangePB.TradesResponse, error)
	StreamTrades(ctx context.Context, req derivativeExchangePB.StreamTradesRequest) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamTradesClient, error)
	GetSubaccountOrdersList(ctx context.Context, req derivativeExchangePB.SubaccountOrdersListRequest) (derivativeExchangePB.SubaccountOrdersListResponse, error)
	GetSubaccountTradesList(ctx context.Context, req derivativeExchangePB.SubaccountTradesListRequest) (derivativeExchangePB.SubaccountTradesListResponse, error)
	GetFundingPayments(ctx context.Context, req derivativeExchangePB.FundingPaymentsRequest) (derivativeExchangePB.FundingPaymentsResponse, error)
	GetFundingRates(ctx context.Context, req derivativeExchangePB.FundingRatesRequest) (derivativeExchangePB.FundingRatesResponse, error)
	GetPrice(ctx context.Context, baseSymbol string, quoteSymbol string, oracleType string, oracleScaleFactor uint32) (oraclePB.PriceResponse, error)
	GetOracleList(ctx context.Context) (oraclePB.OracleListResponse, error)
	StreamPrices(ctx context.Context, baseSymbol string, quoteSymbol string, oracleType string) (oraclePB.InjectiveOracleRPC_StreamPricesClient, error)
	GetAuction(ctx context.Context, round int64) (auctionPB.AuctionResponse, error)
	GetAuctions(ctx context.Context) (auctionPB.AuctionsResponse, error)
	StreamBids(ctx context.Context) (auctionPB.InjectiveAuctionRPC_StreamBidsClient, error)
	Close()
}

func NewExchangeClient(protoAddr string, options ...common.ClientOption) (ExchangeClient, error) {
	// process options
	opts := common.DefaultClientOptions()
	for _, opt := range options {
		if err := opt(opts); err != nil {
			err = errors.Wrap(err, "error in client option")
			return nil, err
		}
	}

	// create grpc client
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
	cc := &exchangeClient{
		opts: opts,
		conn: conn,

		metaClient:               metaPB.NewInjectiveMetaRPCClient(conn),
		accountClient:            accountPB.NewInjectiveAccountsRPCClient(conn),
		auctionClient:            auctionPB.NewInjectiveAuctionRPCClient(conn),
		oracleClient:             oraclePB.NewInjectiveOracleRPCClient(conn),
		insuranceClient:          insurancePB.NewInjectiveInsuranceRPCClient(conn),
		spotExchangeClient:       spotExchangePB.NewInjectiveSpotExchangeRPCClient(conn),
		derivativeExchangeClient: derivativeExchangePB.NewInjectiveDerivativeExchangeRPCClient(conn),

		logger: log.WithFields(log.Fields{
			"module": "sdk-go",
			"svc":    "exchangeClient",
		}),
	}

	return cc, nil
}

type exchangeClient struct {
	opts   *common.ClientOptions
	conn   *grpc.ClientConn
	logger log.Logger
	client *grpc.ClientConn

	sessionCookie string

	metaClient               metaPB.InjectiveMetaRPCClient
	accountClient            accountPB.InjectiveAccountsRPCClient
	auctionClient            auctionPB.InjectiveAuctionRPCClient
	oracleClient             oraclePB.InjectiveOracleRPCClient
	insuranceClient          insurancePB.InjectiveInsuranceRPCClient
	spotExchangeClient       spotExchangePB.InjectiveSpotExchangeRPCClient
	derivativeExchangeClient derivativeExchangePB.InjectiveDerivativeExchangeRPCClient

	closed int64
}

func (c *exchangeClient) setCookie(metadata metadata.MD) {
	md := metadata.Get("set-cookie")
	if len(md) > 0 {
		c.sessionCookie = md[0]
	}
}

func (c *exchangeClient) getCookie(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "cookie", c.sessionCookie)
}

func (c *exchangeClient) QueryClient() *grpc.ClientConn {
	return c.conn
}

// Derivatives RPC

func (c *exchangeClient) GetOrders(ctx context.Context, req derivativeExchangePB.OrdersRequest) (derivativeExchangePB.OrdersResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.Orders(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return derivativeExchangePB.OrdersResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) GetPositions(ctx context.Context, req derivativeExchangePB.PositionsRequest) (derivativeExchangePB.PositionsResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.Positions(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return derivativeExchangePB.PositionsResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) GetOrderbook(ctx context.Context, marketId string) (derivativeExchangePB.OrderbookResponse, error) {
	req := derivativeExchangePB.OrderbookRequest{
		MarketId: marketId,
	}

	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.Orderbook(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return derivativeExchangePB.OrderbookResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) GetOrderbooks(ctx context.Context, marketIds []string) (derivativeExchangePB.OrderbooksResponse, error) {
	req := derivativeExchangePB.OrderbooksRequest{
		MarketIds: marketIds,
	}

	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.Orderbooks(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return derivativeExchangePB.OrderbooksResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) StreamOrderbook(ctx context.Context, marketIds []string) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamOrderbookClient, error) {
	req := derivativeExchangePB.StreamOrderbookRequest{
		MarketIds: marketIds,
	}

	ctx = c.getCookie(ctx)
	stream, err := c.derivativeExchangeClient.StreamOrderbook(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	header, err := stream.Header()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	c.setCookie(header)

	return stream, nil
}

func (c *exchangeClient) GetMarkets(ctx context.Context, req derivativeExchangePB.MarketsRequest) (derivativeExchangePB.MarketsResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.Markets(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return derivativeExchangePB.MarketsResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) GetMarket(ctx context.Context, marketId string) (derivativeExchangePB.MarketResponse, error) {
	req := derivativeExchangePB.MarketRequest{
		MarketId: marketId,
	}

	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.Market(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return derivativeExchangePB.MarketResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) StreamMarket(ctx context.Context, marketIds []string) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamMarketClient, error) {
	req := derivativeExchangePB.StreamMarketRequest{
		MarketIds: marketIds,
	}

	ctx = c.getCookie(ctx)
	stream, err := c.derivativeExchangeClient.StreamMarket(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	header, err := stream.Header()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	c.setCookie(header)

	return stream, nil
}

func (c *exchangeClient) StreamPositions(ctx context.Context, req derivativeExchangePB.StreamPositionsRequest) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamPositionsClient, error) {
	ctx = c.getCookie(ctx)
	stream, err := c.derivativeExchangeClient.StreamPositions(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	header, err := stream.Header()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	c.setCookie(header)

	return stream, nil
}

func (c *exchangeClient) StreamOrders(ctx context.Context, req derivativeExchangePB.StreamOrdersRequest) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamOrdersClient, error) {
	ctx = c.getCookie(ctx)
	stream, err := c.derivativeExchangeClient.StreamOrders(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	header, err := stream.Header()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	c.setCookie(header)

	return stream, nil
}

func (c *exchangeClient) GetTrades(ctx context.Context, req derivativeExchangePB.TradesRequest) (derivativeExchangePB.TradesResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.Trades(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return derivativeExchangePB.TradesResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) StreamTrades(ctx context.Context, req derivativeExchangePB.StreamTradesRequest) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamTradesClient, error) {
	ctx = c.getCookie(ctx)
	stream, err := c.derivativeExchangeClient.StreamTrades(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	header, err := stream.Header()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	c.setCookie(header)

	return stream, nil
}

func (c *exchangeClient) GetSubaccountOrdersList(ctx context.Context, req derivativeExchangePB.SubaccountOrdersListRequest) (derivativeExchangePB.SubaccountOrdersListResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.SubaccountOrdersList(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return derivativeExchangePB.SubaccountOrdersListResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) GetSubaccountTradesList(ctx context.Context, req derivativeExchangePB.SubaccountTradesListRequest) (derivativeExchangePB.SubaccountTradesListResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.SubaccountTradesList(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return derivativeExchangePB.SubaccountTradesListResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) GetFundingPayments(ctx context.Context, req derivativeExchangePB.FundingPaymentsRequest) (derivativeExchangePB.FundingPaymentsResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.FundingPayments(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return derivativeExchangePB.FundingPaymentsResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) GetFundingRates(ctx context.Context, req derivativeExchangePB.FundingRatesRequest) (derivativeExchangePB.FundingRatesResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.FundingRates(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return derivativeExchangePB.FundingRatesResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

// Oracle RPC

func (c *exchangeClient) GetPrice(ctx context.Context, baseSymbol string, quoteSymbol string, oracleType string, oracleScaleFactor uint32) (oraclePB.PriceResponse, error) {
	req := oraclePB.PriceRequest{
		BaseSymbol:        baseSymbol,
		QuoteSymbol:       quoteSymbol,
		OracleType:        oracleType,
		OracleScaleFactor: oracleScaleFactor,
	}

	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.oracleClient.Price(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return oraclePB.PriceResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) GetOracleList(ctx context.Context) (oraclePB.OracleListResponse, error) {
	req := oraclePB.OracleListRequest{}

	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.oracleClient.OracleList(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return oraclePB.OracleListResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) StreamPrices(ctx context.Context, baseSymbol string, quoteSymbol string, oracleType string) (oraclePB.InjectiveOracleRPC_StreamPricesClient, error) {
	req := oraclePB.StreamPricesRequest{
		BaseSymbol:  baseSymbol,
		QuoteSymbol: quoteSymbol,
		OracleType:  oracleType,
	}

	ctx = c.getCookie(ctx)
	stream, err := c.oracleClient.StreamPrices(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	header, err := stream.Header()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	c.setCookie(header)

	return stream, nil
}

// Auction RPC

func (c *exchangeClient) GetAuction(ctx context.Context, round int64) (auctionPB.AuctionResponse, error) {
	req := auctionPB.AuctionRequest{
		Round: round,
	}

	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.auctionClient.AuctionEndpoint(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return auctionPB.AuctionResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) GetAuctions(ctx context.Context) (auctionPB.AuctionsResponse, error) {
	req := auctionPB.AuctionsRequest{}

	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.auctionClient.Auctions(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return auctionPB.AuctionsResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) StreamBids(ctx context.Context) (auctionPB.InjectiveAuctionRPC_StreamBidsClient, error) {
	req := auctionPB.StreamBidsRequest{}

	ctx = c.getCookie(ctx)
	stream, err := c.auctionClient.StreamBids(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	header, err := stream.Header()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	c.setCookie(header)

	return stream, nil
}

func (c *exchangeClient) Close() {
	c.Close()
}
