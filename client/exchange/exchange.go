package exchange

import (
	"context"
	"fmt"

	"github.com/InjectiveLabs/sdk-go/client/common"
	accountPB "github.com/InjectiveLabs/sdk-go/exchange/accounts_rpc/pb"
	auctionPB "github.com/InjectiveLabs/sdk-go/exchange/auction_rpc/pb"
	derivativeExchangePB "github.com/InjectiveLabs/sdk-go/exchange/derivative_exchange_rpc/pb"
	explorerPB "github.com/InjectiveLabs/sdk-go/exchange/explorer_rpc/pb"
	insurancePB "github.com/InjectiveLabs/sdk-go/exchange/insurance_rpc/pb"
	metaPB "github.com/InjectiveLabs/sdk-go/exchange/meta_rpc/pb"
	oraclePB "github.com/InjectiveLabs/sdk-go/exchange/oracle_rpc/pb"
	spotExchangePB "github.com/InjectiveLabs/sdk-go/exchange/spot_exchange_rpc/pb"
	"google.golang.org/grpc/metadata"

	log "github.com/InjectiveLabs/suplog"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type ExchangeClient interface {
	QueryClient() *grpc.ClientConn
	GetDerivativeMarket(ctx context.Context, marketId string) (derivativeExchangePB.MarketResponse, error)
	GetDerivativeOrderbook(ctx context.Context, marketId string) (derivativeExchangePB.OrderbookResponse, error)
	GetDerivativeOrderbooks(ctx context.Context, marketIds []string) (derivativeExchangePB.OrderbooksResponse, error)
	GetDerivativeOrderbookV2(ctx context.Context, marketId string) (derivativeExchangePB.OrderbookV2Response, error)
	GetDerivativeOrderbooksV2(ctx context.Context, marketIds []string) (derivativeExchangePB.OrderbooksV2Response, error)
	// StreamDerivativeOrderbook deprecated API
	StreamDerivativeOrderbook(ctx context.Context, marketIds []string) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamOrderbookClient, error)
	StreamDerivativeOrderbookV2(ctx context.Context, marketIds []string) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamOrderbookV2Client, error)
	StreamDerivativeOrderbookUpdate(ctx context.Context, marketIds []string) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamOrderbookUpdateClient, error)
	StreamDerivativeMarket(ctx context.Context, marketIds []string) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamMarketClient, error)
	GetDerivativeOrders(ctx context.Context, req derivativeExchangePB.OrdersRequest) (derivativeExchangePB.OrdersResponse, error)
	GetDerivativeMarkets(ctx context.Context, req derivativeExchangePB.MarketsRequest) (derivativeExchangePB.MarketsResponse, error)
	GetDerivativePositions(ctx context.Context, req derivativeExchangePB.PositionsRequest) (derivativeExchangePB.PositionsResponse, error)
	StreamDerivativePositions(ctx context.Context, req derivativeExchangePB.StreamPositionsRequest) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamPositionsClient, error)
	StreamDerivativeOrders(ctx context.Context, req derivativeExchangePB.StreamOrdersRequest) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamOrdersClient, error)
	GetDerivativeTrades(ctx context.Context, req derivativeExchangePB.TradesRequest) (derivativeExchangePB.TradesResponse, error)
	StreamDerivativeTrades(ctx context.Context, req derivativeExchangePB.StreamTradesRequest) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamTradesClient, error)
	GetSubaccountDerivativeOrdersList(ctx context.Context, req derivativeExchangePB.SubaccountOrdersListRequest) (derivativeExchangePB.SubaccountOrdersListResponse, error)
	GetSubaccountDerivativeTradesList(ctx context.Context, req derivativeExchangePB.SubaccountTradesListRequest) (derivativeExchangePB.SubaccountTradesListResponse, error)
	GetDerivativeFundingPayments(ctx context.Context, req derivativeExchangePB.FundingPaymentsRequest) (derivativeExchangePB.FundingPaymentsResponse, error)
	GetDerivativeFundingRates(ctx context.Context, req derivativeExchangePB.FundingRatesRequest) (derivativeExchangePB.FundingRatesResponse, error)
	GetPrice(ctx context.Context, baseSymbol string, quoteSymbol string, oracleType string, oracleScaleFactor uint32) (oraclePB.PriceResponse, error)
	GetOracleList(ctx context.Context) (oraclePB.OracleListResponse, error)
	StreamPrices(ctx context.Context, baseSymbol string, quoteSymbol string, oracleType string) (oraclePB.InjectiveOracleRPC_StreamPricesClient, error)
	GetAuction(ctx context.Context, round int64) (auctionPB.AuctionEndpointResponse, error)
	GetAuctions(ctx context.Context) (auctionPB.AuctionsResponse, error)
	StreamBids(ctx context.Context) (auctionPB.InjectiveAuctionRPC_StreamBidsClient, error)
	GetSubaccountsList(ctx context.Context, accountAddress string) (accountPB.SubaccountsListResponse, error)
	GetSubaccountBalance(ctx context.Context, subaccountId string, denom string) (accountPB.SubaccountBalanceEndpointResponse, error)
	StreamSubaccountBalance(ctx context.Context, subaccountId string) (accountPB.InjectiveAccountsRPC_StreamSubaccountBalanceClient, error)
	GetSubaccountBalancesList(ctx context.Context, subaccountId string) (accountPB.SubaccountBalancesListResponse, error)
	GetSubaccountHistory(ctx context.Context, req accountPB.SubaccountHistoryRequest) (accountPB.SubaccountHistoryResponse, error)
	GetSubaccountOrderSummary(ctx context.Context, req accountPB.SubaccountOrderSummaryRequest) (accountPB.SubaccountOrderSummaryResponse, error)
	GetOrderStates(ctx context.Context, req accountPB.OrderStatesRequest) (accountPB.OrderStatesResponse, error)
	GetPortfolio(ctx context.Context, accountAddress string) (accountPB.PortfolioResponse, error)
	GetRewards(ctx context.Context, req accountPB.RewardsRequest) (accountPB.RewardsResponse, error)
	GetSpotOrders(ctx context.Context, req spotExchangePB.OrdersRequest) (spotExchangePB.OrdersResponse, error)
	GetSpotOrderbook(ctx context.Context, marketId string) (spotExchangePB.OrderbookResponse, error)
	GetSpotOrderbooks(ctx context.Context, marketIds []string) (spotExchangePB.OrderbooksResponse, error)
	GetSpotOrderbookV2(ctx context.Context, marketId string) (spotExchangePB.OrderbookV2Response, error)
	GetSpotOrderbooksV2(ctx context.Context, marketIds []string) (spotExchangePB.OrderbooksV2Response, error)
	// StreamSpotOrderbook deprecated API
	StreamSpotOrderbook(ctx context.Context, marketIds []string) (spotExchangePB.InjectiveSpotExchangeRPC_StreamOrderbookClient, error)
	StreamSpotOrderbookV2(ctx context.Context, marketIds []string) (spotExchangePB.InjectiveSpotExchangeRPC_StreamOrderbookV2Client, error)
	StreamSpotOrderbookUpdate(ctx context.Context, marketIds []string) (spotExchangePB.InjectiveSpotExchangeRPC_StreamOrderbookUpdateClient, error)
	GetSpotMarkets(ctx context.Context, req spotExchangePB.MarketsRequest) (spotExchangePB.MarketsResponse, error)
	GetSpotMarket(ctx context.Context, marketId string) (spotExchangePB.MarketResponse, error)
	StreamSpotMarket(ctx context.Context, marketIds []string) (spotExchangePB.InjectiveSpotExchangeRPC_StreamMarketsClient, error)
	StreamSpotOrders(ctx context.Context, req spotExchangePB.StreamOrdersRequest) (spotExchangePB.InjectiveSpotExchangeRPC_StreamOrdersClient, error)
	GetSpotTrades(ctx context.Context, req spotExchangePB.TradesRequest) (spotExchangePB.TradesResponse, error)
	StreamSpotTrades(ctx context.Context, req spotExchangePB.StreamTradesRequest) (spotExchangePB.InjectiveSpotExchangeRPC_StreamTradesClient, error)
	GetSubaccountSpotOrdersList(ctx context.Context, req spotExchangePB.SubaccountOrdersListRequest) (spotExchangePB.SubaccountOrdersListResponse, error)
	GetSubaccountSpotTradesList(ctx context.Context, req spotExchangePB.SubaccountTradesListRequest) (spotExchangePB.SubaccountTradesListResponse, error)
	GetInsuranceFunds(ctx context.Context, req insurancePB.FundsRequest) (insurancePB.FundsResponse, error)
	GetRedemptions(ctx context.Context, req insurancePB.RedemptionsRequest) (insurancePB.RedemptionsResponse, error)
	StreamKeepalive(ctx context.Context) (metaPB.InjectiveMetaRPC_StreamKeepaliveClient, error)
	GetInfo(ctx context.Context, req metaPB.InfoRequest) (metaPB.InfoResponse, error)
	GetVersion(ctx context.Context, req metaPB.VersionRequest) (metaPB.VersionResponse, error)
	Ping(ctx context.Context, req metaPB.PingRequest) (metaPB.PingResponse, error)
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
		explorerClient:           explorerPB.NewInjectiveExplorerRPCClient(conn),
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
	explorerClient           explorerPB.InjectiveExplorerRPCClient
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

func (c *exchangeClient) GetDerivativeOrders(ctx context.Context, req derivativeExchangePB.OrdersRequest) (derivativeExchangePB.OrdersResponse, error) {
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

func (c *exchangeClient) GetDerivativePositions(ctx context.Context, req derivativeExchangePB.PositionsRequest) (derivativeExchangePB.PositionsResponse, error) {
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

func (c *exchangeClient) GetDerivativeOrderbook(ctx context.Context, marketId string) (derivativeExchangePB.OrderbookResponse, error) {
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

func (c *exchangeClient) GetDerivativeOrderbooks(ctx context.Context, marketIds []string) (derivativeExchangePB.OrderbooksResponse, error) {
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

func (c *exchangeClient) GetDerivativeOrderbookV2(ctx context.Context, marketId string) (derivativeExchangePB.OrderbookV2Response, error) {
	req := derivativeExchangePB.OrderbookV2Request{
		MarketId: marketId,
	}

	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.OrderbookV2(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return derivativeExchangePB.OrderbookV2Response{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) GetDerivativeOrderbooksV2(ctx context.Context, marketIds []string) (derivativeExchangePB.OrderbooksV2Response, error) {
	req := derivativeExchangePB.OrderbooksV2Request{
		MarketIds: marketIds,
	}

	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.OrderbooksV2(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return derivativeExchangePB.OrderbooksV2Response{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) StreamDerivativeOrderbook(ctx context.Context, marketIds []string) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamOrderbookClient, error) {
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

func (c *exchangeClient) StreamDerivativeOrderbookV2(ctx context.Context, marketIds []string) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamOrderbookV2Client, error) {
	req := derivativeExchangePB.StreamOrderbookV2Request{
		MarketIds: marketIds,
	}

	ctx = c.getCookie(ctx)
	stream, err := c.derivativeExchangeClient.StreamOrderbookV2(ctx, &req)
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

func (c *exchangeClient) StreamDerivativeOrderbookUpdate(ctx context.Context, marketIds []string) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamOrderbookUpdateClient, error) {
	req := derivativeExchangePB.StreamOrderbookUpdateRequest{
		MarketIds: marketIds,
	}

	ctx = c.getCookie(ctx)
	stream, err := c.derivativeExchangeClient.StreamOrderbookUpdate(ctx, &req)
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

func (c *exchangeClient) GetDerivativeMarkets(ctx context.Context, req derivativeExchangePB.MarketsRequest) (derivativeExchangePB.MarketsResponse, error) {
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

func (c *exchangeClient) GetDerivativeMarket(ctx context.Context, marketId string) (derivativeExchangePB.MarketResponse, error) {
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

func (c *exchangeClient) StreamDerivativeMarket(ctx context.Context, marketIds []string) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamMarketClient, error) {
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

func (c *exchangeClient) StreamDerivativePositions(ctx context.Context, req derivativeExchangePB.StreamPositionsRequest) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamPositionsClient, error) {
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

func (c *exchangeClient) StreamDerivativeOrders(ctx context.Context, req derivativeExchangePB.StreamOrdersRequest) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamOrdersClient, error) {
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

func (c *exchangeClient) GetDerivativeTrades(ctx context.Context, req derivativeExchangePB.TradesRequest) (derivativeExchangePB.TradesResponse, error) {
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

func (c *exchangeClient) StreamDerivativeTrades(ctx context.Context, req derivativeExchangePB.StreamTradesRequest) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamTradesClient, error) {
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

func (c *exchangeClient) GetSubaccountDerivativeOrdersList(ctx context.Context, req derivativeExchangePB.SubaccountOrdersListRequest) (derivativeExchangePB.SubaccountOrdersListResponse, error) {
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

func (c *exchangeClient) GetSubaccountDerivativeTradesList(ctx context.Context, req derivativeExchangePB.SubaccountTradesListRequest) (derivativeExchangePB.SubaccountTradesListResponse, error) {
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

func (c *exchangeClient) GetDerivativeFundingPayments(ctx context.Context, req derivativeExchangePB.FundingPaymentsRequest) (derivativeExchangePB.FundingPaymentsResponse, error) {
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

func (c *exchangeClient) GetDerivativeFundingRates(ctx context.Context, req derivativeExchangePB.FundingRatesRequest) (derivativeExchangePB.FundingRatesResponse, error) {
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

func (c *exchangeClient) GetAuction(ctx context.Context, round int64) (auctionPB.AuctionEndpointResponse, error) {
	req := auctionPB.AuctionEndpointRequest{
		Round: round,
	}

	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.auctionClient.AuctionEndpoint(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return auctionPB.AuctionEndpointResponse{}, err
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

// Accounts RPC

func (c *exchangeClient) GetSubaccountsList(ctx context.Context, accountAddress string) (accountPB.SubaccountsListResponse, error) {
	req := accountPB.SubaccountsListRequest{
		AccountAddress: accountAddress,
	}

	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.accountClient.SubaccountsList(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return accountPB.SubaccountsListResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) GetSubaccountBalance(ctx context.Context, subaccountId string, denom string) (accountPB.SubaccountBalanceEndpointResponse, error) {
	req := accountPB.SubaccountBalanceEndpointRequest{
		SubaccountId: subaccountId,
		Denom:        denom,
	}

	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.accountClient.SubaccountBalanceEndpoint(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return accountPB.SubaccountBalanceEndpointResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) StreamSubaccountBalance(ctx context.Context, subaccountId string) (accountPB.InjectiveAccountsRPC_StreamSubaccountBalanceClient, error) {
	req := accountPB.StreamSubaccountBalanceRequest{
		SubaccountId: subaccountId,
	}

	ctx = c.getCookie(ctx)
	stream, err := c.accountClient.StreamSubaccountBalance(ctx, &req)
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

func (c *exchangeClient) GetSubaccountBalancesList(ctx context.Context, subaccountId string) (accountPB.SubaccountBalancesListResponse, error) {
	req := accountPB.SubaccountBalancesListRequest{
		SubaccountId: subaccountId,
	}

	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.accountClient.SubaccountBalancesList(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return accountPB.SubaccountBalancesListResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) GetSubaccountHistory(ctx context.Context, req accountPB.SubaccountHistoryRequest) (accountPB.SubaccountHistoryResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.accountClient.SubaccountHistory(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return accountPB.SubaccountHistoryResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) GetSubaccountOrderSummary(ctx context.Context, req accountPB.SubaccountOrderSummaryRequest) (accountPB.SubaccountOrderSummaryResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.accountClient.SubaccountOrderSummary(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return accountPB.SubaccountOrderSummaryResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) GetOrderStates(ctx context.Context, req accountPB.OrderStatesRequest) (accountPB.OrderStatesResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.accountClient.OrderStates(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return accountPB.OrderStatesResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) GetPortfolio(ctx context.Context, accountAddress string) (accountPB.PortfolioResponse, error) {
	req := accountPB.PortfolioRequest{
		AccountAddress: accountAddress,
	}

	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.accountClient.Portfolio(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return accountPB.PortfolioResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) GetRewards(ctx context.Context, req accountPB.RewardsRequest) (accountPB.RewardsResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.accountClient.Rewards(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return accountPB.RewardsResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

// Spot RPC

func (c *exchangeClient) GetSpotOrders(ctx context.Context, req spotExchangePB.OrdersRequest) (spotExchangePB.OrdersResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.spotExchangeClient.Orders(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return spotExchangePB.OrdersResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) GetSpotOrderbook(ctx context.Context, marketId string) (spotExchangePB.OrderbookResponse, error) {
	req := spotExchangePB.OrderbookRequest{
		MarketId: marketId,
	}

	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.spotExchangeClient.Orderbook(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return spotExchangePB.OrderbookResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) GetSpotOrderbookV2(ctx context.Context, marketId string) (spotExchangePB.OrderbookV2Response, error) {
	req := spotExchangePB.OrderbookV2Request{
		MarketId: marketId,
	}

	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.spotExchangeClient.OrderbookV2(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return spotExchangePB.OrderbookV2Response{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) GetSpotOrderbooks(ctx context.Context, marketIds []string) (spotExchangePB.OrderbooksResponse, error) {
	req := spotExchangePB.OrderbooksRequest{
		MarketIds: marketIds,
	}

	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.spotExchangeClient.Orderbooks(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return spotExchangePB.OrderbooksResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) GetSpotOrderbooksV2(ctx context.Context, marketIds []string) (spotExchangePB.OrderbooksV2Response, error) {
	req := spotExchangePB.OrderbooksV2Request{
		MarketIds: marketIds,
	}

	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.spotExchangeClient.OrderbooksV2(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return spotExchangePB.OrderbooksV2Response{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) StreamSpotOrderbookUpdate(ctx context.Context, marketIds []string) (spotExchangePB.InjectiveSpotExchangeRPC_StreamOrderbookUpdateClient, error) {
	req := spotExchangePB.StreamOrderbookUpdateRequest{
		MarketIds: marketIds,
	}

	ctx = c.getCookie(ctx)
	stream, err := c.spotExchangeClient.StreamOrderbookUpdate(ctx, &req)
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

func (c *exchangeClient) StreamSpotOrderbook(ctx context.Context, marketIds []string) (spotExchangePB.InjectiveSpotExchangeRPC_StreamOrderbookClient, error) {
	req := spotExchangePB.StreamOrderbookRequest{
		MarketIds: marketIds,
	}

	ctx = c.getCookie(ctx)
	stream, err := c.spotExchangeClient.StreamOrderbook(ctx, &req)
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

func (c *exchangeClient) StreamSpotOrderbookV2(ctx context.Context, marketIds []string) (spotExchangePB.InjectiveSpotExchangeRPC_StreamOrderbookV2Client, error) {
	req := spotExchangePB.StreamOrderbookV2Request{
		MarketIds: marketIds,
	}

	ctx = c.getCookie(ctx)
	stream, err := c.spotExchangeClient.StreamOrderbookV2(ctx, &req)
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

func (c *exchangeClient) GetSpotMarkets(ctx context.Context, req spotExchangePB.MarketsRequest) (spotExchangePB.MarketsResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.spotExchangeClient.Markets(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return spotExchangePB.MarketsResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) GetSpotMarket(ctx context.Context, marketId string) (spotExchangePB.MarketResponse, error) {
	req := spotExchangePB.MarketRequest{
		MarketId: marketId,
	}

	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.spotExchangeClient.Market(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return spotExchangePB.MarketResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) StreamSpotMarket(ctx context.Context, marketIds []string) (spotExchangePB.InjectiveSpotExchangeRPC_StreamMarketsClient, error) {
	req := spotExchangePB.StreamMarketsRequest{
		MarketIds: marketIds,
	}

	ctx = c.getCookie(ctx)
	stream, err := c.spotExchangeClient.StreamMarkets(ctx, &req)
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

func (c *exchangeClient) StreamSpotOrders(ctx context.Context, req spotExchangePB.StreamOrdersRequest) (spotExchangePB.InjectiveSpotExchangeRPC_StreamOrdersClient, error) {
	ctx = c.getCookie(ctx)
	stream, err := c.spotExchangeClient.StreamOrders(ctx, &req)
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

func (c *exchangeClient) GetSpotTrades(ctx context.Context, req spotExchangePB.TradesRequest) (spotExchangePB.TradesResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.spotExchangeClient.Trades(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return spotExchangePB.TradesResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) StreamSpotTrades(ctx context.Context, req spotExchangePB.StreamTradesRequest) (spotExchangePB.InjectiveSpotExchangeRPC_StreamTradesClient, error) {
	ctx = c.getCookie(ctx)
	stream, err := c.spotExchangeClient.StreamTrades(ctx, &req)
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

func (c *exchangeClient) GetSubaccountSpotOrdersList(ctx context.Context, req spotExchangePB.SubaccountOrdersListRequest) (spotExchangePB.SubaccountOrdersListResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.spotExchangeClient.SubaccountOrdersList(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return spotExchangePB.SubaccountOrdersListResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) GetSubaccountSpotTradesList(ctx context.Context, req spotExchangePB.SubaccountTradesListRequest) (spotExchangePB.SubaccountTradesListResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.spotExchangeClient.SubaccountTradesList(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return spotExchangePB.SubaccountTradesListResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) GetInsuranceFunds(ctx context.Context, req insurancePB.FundsRequest) (insurancePB.FundsResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.insuranceClient.Funds(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return insurancePB.FundsResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) GetRedemptions(ctx context.Context, req insurancePB.RedemptionsRequest) (insurancePB.RedemptionsResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.insuranceClient.Redemptions(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return insurancePB.RedemptionsResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) Ping(ctx context.Context, req metaPB.PingRequest) (metaPB.PingResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.metaClient.Ping(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return metaPB.PingResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) GetVersion(ctx context.Context, req metaPB.VersionRequest) (metaPB.VersionResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.metaClient.Version(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return metaPB.VersionResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) GetInfo(ctx context.Context, req metaPB.InfoRequest) (metaPB.InfoResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.metaClient.Info(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return metaPB.InfoResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) StreamKeepalive(ctx context.Context) (metaPB.InjectiveMetaRPC_StreamKeepaliveClient, error) {
	req := metaPB.StreamKeepaliveRequest{}

	ctx = c.getCookie(ctx)
	stream, err := c.metaClient.StreamKeepalive(ctx, &req)
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
	c.conn.Close()
}
