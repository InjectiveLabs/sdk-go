package exchange

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/InjectiveLabs/sdk-go/client/common"
	accountPB "github.com/InjectiveLabs/sdk-go/exchange/accounts_rpc/pb"
	auctionPB "github.com/InjectiveLabs/sdk-go/exchange/auction_rpc/pb"
	derivativeExchangePB "github.com/InjectiveLabs/sdk-go/exchange/derivative_exchange_rpc/pb"
	explorerPB "github.com/InjectiveLabs/sdk-go/exchange/explorer_rpc/pb"
	insurancePB "github.com/InjectiveLabs/sdk-go/exchange/insurance_rpc/pb"
	metaPB "github.com/InjectiveLabs/sdk-go/exchange/meta_rpc/pb"
	oraclePB "github.com/InjectiveLabs/sdk-go/exchange/oracle_rpc/pb"
	portfolioExchangePB "github.com/InjectiveLabs/sdk-go/exchange/portfolio_rpc/pb"
	spotExchangePB "github.com/InjectiveLabs/sdk-go/exchange/spot_exchange_rpc/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type ExchangeClient interface {
	QueryClient() *grpc.ClientConn
	GetDerivativeMarket(ctx context.Context, marketId string) (*derivativeExchangePB.MarketResponse, error)
	GetDerivativeOrderbookV2(ctx context.Context, marketId string) (*derivativeExchangePB.OrderbookV2Response, error)
	GetDerivativeOrderbooksV2(ctx context.Context, marketIds []string) (*derivativeExchangePB.OrderbooksV2Response, error)
	// StreamDerivativeOrderbook deprecated API
	StreamDerivativeOrderbookV2(ctx context.Context, marketIds []string) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamOrderbookV2Client, error)
	StreamDerivativeOrderbookUpdate(ctx context.Context, marketIds []string) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamOrderbookUpdateClient, error)
	StreamDerivativeMarket(ctx context.Context, marketIds []string) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamMarketClient, error)
	GetDerivativeOrders(ctx context.Context, req *derivativeExchangePB.OrdersRequest) (*derivativeExchangePB.OrdersResponse, error)
	GetDerivativeMarkets(ctx context.Context, req *derivativeExchangePB.MarketsRequest) (*derivativeExchangePB.MarketsResponse, error)
	GetDerivativePositions(ctx context.Context, req *derivativeExchangePB.PositionsRequest) (*derivativeExchangePB.PositionsResponse, error)
	GetDerivativePositionsV2(ctx context.Context, req *derivativeExchangePB.PositionsV2Request) (*derivativeExchangePB.PositionsV2Response, error)
	GetDerivativeLiquidablePositions(ctx context.Context, req *derivativeExchangePB.LiquidablePositionsRequest) (*derivativeExchangePB.LiquidablePositionsResponse, error)
	StreamDerivativePositions(ctx context.Context, req *derivativeExchangePB.StreamPositionsRequest) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamPositionsClient, error)
	StreamDerivativeOrders(ctx context.Context, req *derivativeExchangePB.StreamOrdersRequest) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamOrdersClient, error)
	GetDerivativeTrades(ctx context.Context, req *derivativeExchangePB.TradesRequest) (*derivativeExchangePB.TradesResponse, error)
	GetDerivativeTradesV2(ctx context.Context, req *derivativeExchangePB.TradesV2Request) (*derivativeExchangePB.TradesV2Response, error)
	StreamDerivativeTrades(ctx context.Context, req *derivativeExchangePB.StreamTradesRequest) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamTradesClient, error)
	StreamDerivativeV2Trades(ctx context.Context, req *derivativeExchangePB.StreamTradesV2Request) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamTradesV2Client, error)
	GetSubaccountDerivativeOrdersList(ctx context.Context, req *derivativeExchangePB.SubaccountOrdersListRequest) (*derivativeExchangePB.SubaccountOrdersListResponse, error)
	GetSubaccountDerivativeTradesList(ctx context.Context, req *derivativeExchangePB.SubaccountTradesListRequest) (*derivativeExchangePB.SubaccountTradesListResponse, error)
	GetHistoricalDerivativeOrders(ctx context.Context, req *derivativeExchangePB.OrdersHistoryRequest) (*derivativeExchangePB.OrdersHistoryResponse, error)
	StreamHistoricalDerivativeOrders(ctx context.Context, req *derivativeExchangePB.StreamOrdersHistoryRequest) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamOrdersHistoryClient, error)
	GetDerivativeFundingPayments(ctx context.Context, req *derivativeExchangePB.FundingPaymentsRequest) (*derivativeExchangePB.FundingPaymentsResponse, error)
	GetDerivativeFundingRates(ctx context.Context, req *derivativeExchangePB.FundingRatesRequest) (*derivativeExchangePB.FundingRatesResponse, error)
	GetPrice(ctx context.Context, baseSymbol string, quoteSymbol string, oracleType string, oracleScaleFactor uint32) (*oraclePB.PriceResponse, error)
	GetOracleList(ctx context.Context) (*oraclePB.OracleListResponse, error)
	StreamPrices(ctx context.Context, baseSymbol string, quoteSymbol string, oracleType string) (oraclePB.InjectiveOracleRPC_StreamPricesClient, error)
	GetAuction(ctx context.Context, round int64) (*auctionPB.AuctionEndpointResponse, error)
	GetAuctions(ctx context.Context) (*auctionPB.AuctionsResponse, error)
	StreamBids(ctx context.Context) (auctionPB.InjectiveAuctionRPC_StreamBidsClient, error)
	GetSubaccountsList(ctx context.Context, accountAddress string) (*accountPB.SubaccountsListResponse, error)
	GetSubaccountBalance(ctx context.Context, subaccountId string, denom string) (*accountPB.SubaccountBalanceEndpointResponse, error)
	StreamSubaccountBalance(ctx context.Context, subaccountId string) (accountPB.InjectiveAccountsRPC_StreamSubaccountBalanceClient, error)
	GetSubaccountBalancesList(ctx context.Context, subaccountId string) (*accountPB.SubaccountBalancesListResponse, error)
	GetSubaccountHistory(ctx context.Context, req *accountPB.SubaccountHistoryRequest) (*accountPB.SubaccountHistoryResponse, error)
	GetSubaccountOrderSummary(ctx context.Context, req *accountPB.SubaccountOrderSummaryRequest) (*accountPB.SubaccountOrderSummaryResponse, error)
	GetOrderStates(ctx context.Context, req *accountPB.OrderStatesRequest) (*accountPB.OrderStatesResponse, error)
	GetPortfolio(ctx context.Context, accountAddress string) (*accountPB.PortfolioResponse, error)
	GetRewards(ctx context.Context, req *accountPB.RewardsRequest) (*accountPB.RewardsResponse, error)
	GetSpotOrders(ctx context.Context, req *spotExchangePB.OrdersRequest) (*spotExchangePB.OrdersResponse, error)
	GetSpotOrderbookV2(ctx context.Context, marketId string) (*spotExchangePB.OrderbookV2Response, error)
	GetSpotOrderbooksV2(ctx context.Context, marketIds []string) (*spotExchangePB.OrderbooksV2Response, error)
	// StreamSpotOrderbook deprecated API
	StreamSpotOrderbookV2(ctx context.Context, marketIds []string) (spotExchangePB.InjectiveSpotExchangeRPC_StreamOrderbookV2Client, error)
	StreamSpotOrderbookUpdate(ctx context.Context, marketIds []string) (spotExchangePB.InjectiveSpotExchangeRPC_StreamOrderbookUpdateClient, error)
	GetSpotMarkets(ctx context.Context, req *spotExchangePB.MarketsRequest) (*spotExchangePB.MarketsResponse, error)
	GetSpotMarket(ctx context.Context, marketId string) (*spotExchangePB.MarketResponse, error)
	StreamSpotMarket(ctx context.Context, marketIds []string) (spotExchangePB.InjectiveSpotExchangeRPC_StreamMarketsClient, error)
	StreamSpotOrders(ctx context.Context, req *spotExchangePB.StreamOrdersRequest) (spotExchangePB.InjectiveSpotExchangeRPC_StreamOrdersClient, error)
	GetSpotTrades(ctx context.Context, req *spotExchangePB.TradesRequest) (*spotExchangePB.TradesResponse, error)
	GetSpotTradesV2(ctx context.Context, req *spotExchangePB.TradesV2Request) (*spotExchangePB.TradesV2Response, error)
	StreamSpotTrades(ctx context.Context, req *spotExchangePB.StreamTradesRequest) (spotExchangePB.InjectiveSpotExchangeRPC_StreamTradesClient, error)
	StreamSpotTradesV2(ctx context.Context, req *spotExchangePB.StreamTradesV2Request) (spotExchangePB.InjectiveSpotExchangeRPC_StreamTradesV2Client, error)
	GetSubaccountSpotOrdersList(ctx context.Context, req *spotExchangePB.SubaccountOrdersListRequest) (*spotExchangePB.SubaccountOrdersListResponse, error)
	GetSubaccountSpotTradesList(ctx context.Context, req *spotExchangePB.SubaccountTradesListRequest) (*spotExchangePB.SubaccountTradesListResponse, error)
	GetHistoricalSpotOrders(ctx context.Context, req *spotExchangePB.OrdersHistoryRequest) (*spotExchangePB.OrdersHistoryResponse, error)
	StreamHistoricalSpotOrders(ctx context.Context, req *spotExchangePB.StreamOrdersHistoryRequest) (spotExchangePB.InjectiveSpotExchangeRPC_StreamOrdersHistoryClient, error)
	GetInsuranceFunds(ctx context.Context, req *insurancePB.FundsRequest) (*insurancePB.FundsResponse, error)
	GetRedemptions(ctx context.Context, req *insurancePB.RedemptionsRequest) (*insurancePB.RedemptionsResponse, error)

	GetAccountPortfolio(ctx context.Context, accountAddress string) (*portfolioExchangePB.AccountPortfolioResponse, error)
	GetAccountPortfolioBalances(ctx context.Context, accountAddress string) (*portfolioExchangePB.AccountPortfolioBalancesResponse, error)
	StreamAccountPortfolio(ctx context.Context, accountAddress string, subaccountId, balanceType string) (portfolioExchangePB.InjectivePortfolioRPC_StreamAccountPortfolioClient, error)

	StreamKeepalive(ctx context.Context) (metaPB.InjectiveMetaRPC_StreamKeepaliveClient, error)
	GetInfo(ctx context.Context, req *metaPB.InfoRequest) (*metaPB.InfoResponse, error)
	GetVersion(ctx context.Context, req *metaPB.VersionRequest) (*metaPB.VersionResponse, error)
	Ping(ctx context.Context, req *metaPB.PingRequest) (*metaPB.PingResponse, error)
	OrdersHistory(ctx context.Context, in *derivativeExchangePB.OrdersHistoryRequest) (*derivativeExchangePB.OrdersHistoryResponse, error)
	Close()
}

func NewExchangeClient(network common.Network, options ...common.ClientOption) (ExchangeClient, error) {
	// process options
	opts := common.DefaultClientOptions()
	if network.ExchangeTlsCert != nil {
		options = append(options, common.OptionTLSCert(network.ExchangeTlsCert))
	}
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
		conn, err = grpc.Dial(network.ExchangeGrpcEndpoint, grpc.WithTransportCredentials(opts.TLSCert), grpc.WithContextDialer(common.DialerFunc))
	} else {

		conn, err = grpc.Dial(network.ExchangeGrpcEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(common.DialerFunc))
	}
	if err != nil {
		err := errors.Wrapf(err, "failed to connect to the gRPC: %s", network.ExchangeGrpcEndpoint)
		return nil, err
	}

	// build client
	cc := &exchangeClient{
		opts:    opts,
		network: network,
		conn:    conn,

		metaClient:               metaPB.NewInjectiveMetaRPCClient(conn),
		explorerClient:           explorerPB.NewInjectiveExplorerRPCClient(conn),
		accountClient:            accountPB.NewInjectiveAccountsRPCClient(conn),
		auctionClient:            auctionPB.NewInjectiveAuctionRPCClient(conn),
		oracleClient:             oraclePB.NewInjectiveOracleRPCClient(conn),
		insuranceClient:          insurancePB.NewInjectiveInsuranceRPCClient(conn),
		spotExchangeClient:       spotExchangePB.NewInjectiveSpotExchangeRPCClient(conn),
		derivativeExchangeClient: derivativeExchangePB.NewInjectiveDerivativeExchangeRPCClient(conn),
		portfolioExchangeClient:  portfolioExchangePB.NewInjectivePortfolioRPCClient(conn),

		logger: opts.Logger,
	}

	return cc, nil
}

type exchangeClient struct {
	opts    *common.ClientOptions
	network common.Network
	conn    *grpc.ClientConn
	logger  *logrus.Logger

	sessionCookie string

	metaClient               metaPB.InjectiveMetaRPCClient
	explorerClient           explorerPB.InjectiveExplorerRPCClient
	accountClient            accountPB.InjectiveAccountsRPCClient
	auctionClient            auctionPB.InjectiveAuctionRPCClient
	oracleClient             oraclePB.InjectiveOracleRPCClient
	insuranceClient          insurancePB.InjectiveInsuranceRPCClient
	spotExchangeClient       spotExchangePB.InjectiveSpotExchangeRPCClient
	derivativeExchangeClient derivativeExchangePB.InjectiveDerivativeExchangeRPCClient
	portfolioExchangeClient  portfolioExchangePB.InjectivePortfolioRPCClient
}

// test
func (c *exchangeClient) OrdersHistory(ctx context.Context, in *derivativeExchangePB.OrdersHistoryRequest) (*derivativeExchangePB.OrdersHistoryResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.OrdersHistory(ctx, in, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	c.setCookie(header)

	return res, nil
}

func (c *exchangeClient) setCookie(metadata metadata.MD) {
	md := metadata.Get("set-cookie")
	if len(md) > 0 {
		c.sessionCookie = md[0]
	}
}

func (c *exchangeClient) requestCookie() metadata.MD {
	var header metadata.MD
	req := metaPB.InfoRequest{Timestamp: time.Now().UnixMilli()}
	_, err := c.metaClient.Info(context.Background(), &req, grpc.Header(&header))
	if err != nil {
		c.logger.Errorln("[INJ-GO-SDK] Failed to get cookie from exchange: ", err)
	}
	return header
}

func (c *exchangeClient) getCookie(ctx context.Context) context.Context {
	provider := common.NewMetadataProvider(c.requestCookie)
	cookie, _ := c.network.ExchangeMetadata(provider)
	return metadata.AppendToOutgoingContext(ctx, "cookie", cookie)
}

func (c *exchangeClient) QueryClient() *grpc.ClientConn {
	return c.conn
}

// Derivatives RPC

func (c *exchangeClient) GetDerivativeOrders(ctx context.Context, req *derivativeExchangePB.OrdersRequest) (*derivativeExchangePB.OrdersResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.Orders(ctx, req)
	if err != nil {
		fmt.Println(err)
		return &derivativeExchangePB.OrdersResponse{}, err
	}

	return res, nil
}

// Deprecated: Use GetDerivativePositionsV2 instead.
func (c *exchangeClient) GetDerivativePositions(ctx context.Context, req *derivativeExchangePB.PositionsRequest) (*derivativeExchangePB.PositionsResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.Positions(ctx, req)
	if err != nil {
		fmt.Println(err)
		return &derivativeExchangePB.PositionsResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetDerivativePositionsV2(ctx context.Context, req *derivativeExchangePB.PositionsV2Request) (*derivativeExchangePB.PositionsV2Response, error) {
	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.PositionsV2(ctx, req)
	if err != nil {
		fmt.Println(err)
		return &derivativeExchangePB.PositionsV2Response{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetDerivativeLiquidablePositions(ctx context.Context, req *derivativeExchangePB.LiquidablePositionsRequest) (*derivativeExchangePB.LiquidablePositionsResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.LiquidablePositions(ctx, req)
	if err != nil {
		fmt.Println(err)
		return &derivativeExchangePB.LiquidablePositionsResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetDerivativeOrderbookV2(ctx context.Context, marketId string) (*derivativeExchangePB.OrderbookV2Response, error) {
	req := derivativeExchangePB.OrderbookV2Request{
		MarketId: marketId,
	}

	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.OrderbookV2(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return &derivativeExchangePB.OrderbookV2Response{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetDerivativeOrderbooksV2(ctx context.Context, marketIds []string) (*derivativeExchangePB.OrderbooksV2Response, error) {
	req := derivativeExchangePB.OrderbooksV2Request{
		MarketIds: marketIds,
	}

	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.OrderbooksV2(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return &derivativeExchangePB.OrderbooksV2Response{}, err
	}

	return res, nil
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

	return stream, nil
}

func (c *exchangeClient) GetDerivativeMarkets(ctx context.Context, req *derivativeExchangePB.MarketsRequest) (*derivativeExchangePB.MarketsResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.Markets(ctx, req)
	if err != nil {
		fmt.Println(err)
		return &derivativeExchangePB.MarketsResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetDerivativeMarket(ctx context.Context, marketId string) (*derivativeExchangePB.MarketResponse, error) {
	req := derivativeExchangePB.MarketRequest{
		MarketId: marketId,
	}

	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.Market(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return &derivativeExchangePB.MarketResponse{}, err
	}

	return res, nil
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

	return stream, nil
}

func (c *exchangeClient) StreamDerivativePositions(ctx context.Context, req *derivativeExchangePB.StreamPositionsRequest) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamPositionsClient, error) {
	ctx = c.getCookie(ctx)
	stream, err := c.derivativeExchangeClient.StreamPositions(ctx, req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

func (c *exchangeClient) StreamDerivativeOrders(ctx context.Context, req *derivativeExchangePB.StreamOrdersRequest) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamOrdersClient, error) {
	ctx = c.getCookie(ctx)
	stream, err := c.derivativeExchangeClient.StreamOrders(ctx, req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

func (c *exchangeClient) GetDerivativeTrades(ctx context.Context, req *derivativeExchangePB.TradesRequest) (*derivativeExchangePB.TradesResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.Trades(ctx, req)
	if err != nil {
		fmt.Println(err)
		return &derivativeExchangePB.TradesResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetDerivativeTradesV2(ctx context.Context, req *derivativeExchangePB.TradesV2Request) (*derivativeExchangePB.TradesV2Response, error) {
	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.TradesV2(ctx, req)
	if err != nil {
		fmt.Println(err)
		return &derivativeExchangePB.TradesV2Response{}, err
	}

	return res, nil
}

func (c *exchangeClient) StreamDerivativeTrades(ctx context.Context, req *derivativeExchangePB.StreamTradesRequest) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamTradesClient, error) {
	ctx = c.getCookie(ctx)
	stream, err := c.derivativeExchangeClient.StreamTrades(ctx, req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

func (c *exchangeClient) StreamDerivativeV2Trades(ctx context.Context, req *derivativeExchangePB.StreamTradesV2Request) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamTradesV2Client, error) {
	ctx = c.getCookie(ctx)
	stream, err := c.derivativeExchangeClient.StreamTradesV2(ctx, req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

func (c *exchangeClient) GetSubaccountDerivativeOrdersList(ctx context.Context, req *derivativeExchangePB.SubaccountOrdersListRequest) (*derivativeExchangePB.SubaccountOrdersListResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.SubaccountOrdersList(ctx, req)
	if err != nil {
		fmt.Println(err)
		return &derivativeExchangePB.SubaccountOrdersListResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetSubaccountDerivativeTradesList(ctx context.Context, req *derivativeExchangePB.SubaccountTradesListRequest) (*derivativeExchangePB.SubaccountTradesListResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.SubaccountTradesList(ctx, req)
	if err != nil {
		fmt.Println(err)
		return &derivativeExchangePB.SubaccountTradesListResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetHistoricalDerivativeOrders(ctx context.Context, req *derivativeExchangePB.OrdersHistoryRequest) (*derivativeExchangePB.OrdersHistoryResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.OrdersHistory(ctx, req)
	if err != nil {
		return &derivativeExchangePB.OrdersHistoryResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) StreamHistoricalDerivativeOrders(ctx context.Context, req *derivativeExchangePB.StreamOrdersHistoryRequest) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamOrdersHistoryClient, error) {
	ctx = c.getCookie(ctx)
	stream, err := c.derivativeExchangeClient.StreamOrdersHistory(ctx, req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

func (c *exchangeClient) GetDerivativeFundingPayments(ctx context.Context, req *derivativeExchangePB.FundingPaymentsRequest) (*derivativeExchangePB.FundingPaymentsResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.FundingPayments(ctx, req)
	if err != nil {
		fmt.Println(err)
		return &derivativeExchangePB.FundingPaymentsResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetDerivativeFundingRates(ctx context.Context, req *derivativeExchangePB.FundingRatesRequest) (*derivativeExchangePB.FundingRatesResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.FundingRates(ctx, req)
	if err != nil {
		fmt.Println(err)
		return &derivativeExchangePB.FundingRatesResponse{}, err
	}

	return res, nil
}

// Oracle RPC

func (c *exchangeClient) GetPrice(ctx context.Context, baseSymbol string, quoteSymbol string, oracleType string, oracleScaleFactor uint32) (*oraclePB.PriceResponse, error) {
	req := oraclePB.PriceRequest{
		BaseSymbol:        baseSymbol,
		QuoteSymbol:       quoteSymbol,
		OracleType:        oracleType,
		OracleScaleFactor: oracleScaleFactor,
	}

	ctx = c.getCookie(ctx)
	res, err := c.oracleClient.Price(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return &oraclePB.PriceResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetOracleList(ctx context.Context) (*oraclePB.OracleListResponse, error) {
	req := oraclePB.OracleListRequest{}

	ctx = c.getCookie(ctx)
	res, err := c.oracleClient.OracleList(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return &oraclePB.OracleListResponse{}, err
	}

	return res, nil
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

	return stream, nil
}

// Auction RPC

func (c *exchangeClient) GetAuction(ctx context.Context, round int64) (*auctionPB.AuctionEndpointResponse, error) {
	req := auctionPB.AuctionEndpointRequest{
		Round: round,
	}

	ctx = c.getCookie(ctx)
	res, err := c.auctionClient.AuctionEndpoint(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return &auctionPB.AuctionEndpointResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetAuctions(ctx context.Context) (*auctionPB.AuctionsResponse, error) {
	req := auctionPB.AuctionsRequest{}

	ctx = c.getCookie(ctx)
	res, err := c.auctionClient.Auctions(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return &auctionPB.AuctionsResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) StreamBids(ctx context.Context) (auctionPB.InjectiveAuctionRPC_StreamBidsClient, error) {
	req := auctionPB.StreamBidsRequest{}

	ctx = c.getCookie(ctx)
	stream, err := c.auctionClient.StreamBids(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

// Accounts RPC

func (c *exchangeClient) GetSubaccountsList(ctx context.Context, accountAddress string) (*accountPB.SubaccountsListResponse, error) {
	req := accountPB.SubaccountsListRequest{
		AccountAddress: accountAddress,
	}

	ctx = c.getCookie(ctx)
	res, err := c.accountClient.SubaccountsList(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return &accountPB.SubaccountsListResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetSubaccountBalance(ctx context.Context, subaccountId string, denom string) (*accountPB.SubaccountBalanceEndpointResponse, error) {
	req := accountPB.SubaccountBalanceEndpointRequest{
		SubaccountId: subaccountId,
		Denom:        denom,
	}

	ctx = c.getCookie(ctx)
	res, err := c.accountClient.SubaccountBalanceEndpoint(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return &accountPB.SubaccountBalanceEndpointResponse{}, err
	}

	return res, nil
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

	return stream, nil
}

func (c *exchangeClient) GetSubaccountBalancesList(ctx context.Context, subaccountId string) (*accountPB.SubaccountBalancesListResponse, error) {
	req := accountPB.SubaccountBalancesListRequest{
		SubaccountId: subaccountId,
	}

	ctx = c.getCookie(ctx)
	res, err := c.accountClient.SubaccountBalancesList(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return &accountPB.SubaccountBalancesListResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetSubaccountHistory(ctx context.Context, req *accountPB.SubaccountHistoryRequest) (*accountPB.SubaccountHistoryResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.accountClient.SubaccountHistory(ctx, req)
	if err != nil {
		fmt.Println(err)
		return &accountPB.SubaccountHistoryResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetSubaccountOrderSummary(ctx context.Context, req *accountPB.SubaccountOrderSummaryRequest) (*accountPB.SubaccountOrderSummaryResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.accountClient.SubaccountOrderSummary(ctx, req)
	if err != nil {
		fmt.Println(err)
		return &accountPB.SubaccountOrderSummaryResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetOrderStates(ctx context.Context, req *accountPB.OrderStatesRequest) (*accountPB.OrderStatesResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.accountClient.OrderStates(ctx, req)
	if err != nil {
		fmt.Println(err)
		return &accountPB.OrderStatesResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetPortfolio(ctx context.Context, accountAddress string) (*accountPB.PortfolioResponse, error) {
	req := accountPB.PortfolioRequest{
		AccountAddress: accountAddress,
	}

	ctx = c.getCookie(ctx)
	res, err := c.accountClient.Portfolio(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return &accountPB.PortfolioResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetRewards(ctx context.Context, req *accountPB.RewardsRequest) (*accountPB.RewardsResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.accountClient.Rewards(ctx, req)
	if err != nil {
		fmt.Println(err)
		return &accountPB.RewardsResponse{}, err
	}

	return res, nil
}

// Spot RPC

func (c *exchangeClient) GetSpotOrders(ctx context.Context, req *spotExchangePB.OrdersRequest) (*spotExchangePB.OrdersResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.spotExchangeClient.Orders(ctx, req)
	if err != nil {
		fmt.Println(err)
		return &spotExchangePB.OrdersResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetSpotOrderbookV2(ctx context.Context, marketId string) (*spotExchangePB.OrderbookV2Response, error) {
	req := spotExchangePB.OrderbookV2Request{
		MarketId: marketId,
	}

	ctx = c.getCookie(ctx)
	res, err := c.spotExchangeClient.OrderbookV2(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return &spotExchangePB.OrderbookV2Response{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetSpotOrderbooksV2(ctx context.Context, marketIds []string) (*spotExchangePB.OrderbooksV2Response, error) {
	req := spotExchangePB.OrderbooksV2Request{
		MarketIds: marketIds,
	}

	ctx = c.getCookie(ctx)
	res, err := c.spotExchangeClient.OrderbooksV2(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return &spotExchangePB.OrderbooksV2Response{}, err
	}

	return res, nil
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

	return stream, nil
}

func (c *exchangeClient) GetSpotMarkets(ctx context.Context, req *spotExchangePB.MarketsRequest) (*spotExchangePB.MarketsResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.spotExchangeClient.Markets(ctx, req)
	if err != nil {
		fmt.Println(err)
		return &spotExchangePB.MarketsResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetSpotMarket(ctx context.Context, marketId string) (*spotExchangePB.MarketResponse, error) {
	req := spotExchangePB.MarketRequest{
		MarketId: marketId,
	}

	ctx = c.getCookie(ctx)
	res, err := c.spotExchangeClient.Market(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return &spotExchangePB.MarketResponse{}, err
	}

	return res, nil
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

	return stream, nil
}

func (c *exchangeClient) StreamSpotOrders(ctx context.Context, req *spotExchangePB.StreamOrdersRequest) (spotExchangePB.InjectiveSpotExchangeRPC_StreamOrdersClient, error) {
	ctx = c.getCookie(ctx)
	stream, err := c.spotExchangeClient.StreamOrders(ctx, req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

func (c *exchangeClient) GetSpotTrades(ctx context.Context, req *spotExchangePB.TradesRequest) (*spotExchangePB.TradesResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.spotExchangeClient.Trades(ctx, req)
	if err != nil {
		fmt.Println(err)
		return &spotExchangePB.TradesResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetSpotTradesV2(ctx context.Context, req *spotExchangePB.TradesV2Request) (*spotExchangePB.TradesV2Response, error) {
	ctx = c.getCookie(ctx)
	res, err := c.spotExchangeClient.TradesV2(ctx, req)
	if err != nil {
		fmt.Println(err)
		return &spotExchangePB.TradesV2Response{}, err
	}

	return res, nil
}

func (c *exchangeClient) StreamSpotTrades(ctx context.Context, req *spotExchangePB.StreamTradesRequest) (spotExchangePB.InjectiveSpotExchangeRPC_StreamTradesClient, error) {
	ctx = c.getCookie(ctx)
	stream, err := c.spotExchangeClient.StreamTrades(ctx, req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

func (c *exchangeClient) StreamSpotTradesV2(ctx context.Context, req *spotExchangePB.StreamTradesV2Request) (spotExchangePB.InjectiveSpotExchangeRPC_StreamTradesV2Client, error) {
	ctx = c.getCookie(ctx)
	stream, err := c.spotExchangeClient.StreamTradesV2(ctx, req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

func (c *exchangeClient) GetSubaccountSpotOrdersList(ctx context.Context, req *spotExchangePB.SubaccountOrdersListRequest) (*spotExchangePB.SubaccountOrdersListResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.spotExchangeClient.SubaccountOrdersList(ctx, req)
	if err != nil {
		fmt.Println(err)
		return &spotExchangePB.SubaccountOrdersListResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetSubaccountSpotTradesList(ctx context.Context, req *spotExchangePB.SubaccountTradesListRequest) (*spotExchangePB.SubaccountTradesListResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.spotExchangeClient.SubaccountTradesList(ctx, req)
	if err != nil {
		fmt.Println(err)
		return &spotExchangePB.SubaccountTradesListResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetHistoricalSpotOrders(ctx context.Context, req *spotExchangePB.OrdersHistoryRequest) (*spotExchangePB.OrdersHistoryResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.spotExchangeClient.OrdersHistory(ctx, req)
	if err != nil {
		return &spotExchangePB.OrdersHistoryResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) StreamHistoricalSpotOrders(ctx context.Context, req *spotExchangePB.StreamOrdersHistoryRequest) (spotExchangePB.InjectiveSpotExchangeRPC_StreamOrdersHistoryClient, error) {
	ctx = c.getCookie(ctx)
	stream, err := c.spotExchangeClient.StreamOrdersHistory(ctx, req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

func (c *exchangeClient) GetInsuranceFunds(ctx context.Context, req *insurancePB.FundsRequest) (*insurancePB.FundsResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.insuranceClient.Funds(ctx, req)
	if err != nil {
		fmt.Println(err)
		return &insurancePB.FundsResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetRedemptions(ctx context.Context, req *insurancePB.RedemptionsRequest) (*insurancePB.RedemptionsResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.insuranceClient.Redemptions(ctx, req)
	if err != nil {
		fmt.Println(err)
		return &insurancePB.RedemptionsResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) Ping(ctx context.Context, req *metaPB.PingRequest) (*metaPB.PingResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.metaClient.Ping(ctx, req)
	if err != nil {
		fmt.Println(err)
		return &metaPB.PingResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetVersion(ctx context.Context, req *metaPB.VersionRequest) (*metaPB.VersionResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.metaClient.Version(ctx, req)
	if err != nil {
		fmt.Println(err)
		return &metaPB.VersionResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetInfo(ctx context.Context, req *metaPB.InfoRequest) (*metaPB.InfoResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.metaClient.Info(ctx, req)
	if err != nil {
		fmt.Println(err)
		return &metaPB.InfoResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) StreamKeepalive(ctx context.Context) (metaPB.InjectiveMetaRPC_StreamKeepaliveClient, error) {
	req := metaPB.StreamKeepaliveRequest{}

	ctx = c.getCookie(ctx)
	stream, err := c.metaClient.StreamKeepalive(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

// Deprecated: Use GetAccountPortfolioBalances instead.
func (c *exchangeClient) GetAccountPortfolio(ctx context.Context, accountAddress string) (*portfolioExchangePB.AccountPortfolioResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.portfolioExchangeClient.AccountPortfolio(ctx, &portfolioExchangePB.AccountPortfolioRequest{
		AccountAddress: accountAddress,
	})
	if err != nil {
		fmt.Println(err)
		return &portfolioExchangePB.AccountPortfolioResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetAccountPortfolioBalances(ctx context.Context, accountAddress string) (*portfolioExchangePB.AccountPortfolioBalancesResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.portfolioExchangeClient.AccountPortfolioBalances(ctx, &portfolioExchangePB.AccountPortfolioBalancesRequest{
		AccountAddress: accountAddress,
	})
	if err != nil {
		fmt.Println(err)
		return &portfolioExchangePB.AccountPortfolioBalancesResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) StreamAccountPortfolio(ctx context.Context, accountAddress string, subaccountId, balanceType string) (portfolioExchangePB.InjectivePortfolioRPC_StreamAccountPortfolioClient, error) {
	ctx = c.getCookie(ctx)
	stream, err := c.portfolioExchangeClient.StreamAccountPortfolio(ctx, &portfolioExchangePB.StreamAccountPortfolioRequest{
		AccountAddress: accountAddress,
		SubaccountId:   subaccountId,
		Type:           balanceType,
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

func (c *exchangeClient) Close() {
	c.conn.Close()
}
