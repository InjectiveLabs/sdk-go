package exchange

import (
	"context"
	"fmt"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/InjectiveLabs/sdk-go/client/common"
	accountPB "github.com/InjectiveLabs/sdk-go/exchange/accounts_rpc/pb"
	auctionPB "github.com/InjectiveLabs/sdk-go/exchange/auction_rpc/pb"
	derivativeExchangePB "github.com/InjectiveLabs/sdk-go/exchange/derivative_exchange_rpc/pb"
	insurancePB "github.com/InjectiveLabs/sdk-go/exchange/insurance_rpc/pb"
	metaPB "github.com/InjectiveLabs/sdk-go/exchange/meta_rpc/pb"
	oraclePB "github.com/InjectiveLabs/sdk-go/exchange/oracle_rpc/pb"
	portfolioExchangePB "github.com/InjectiveLabs/sdk-go/exchange/portfolio_rpc/pb"
	spotExchangePB "github.com/InjectiveLabs/sdk-go/exchange/spot_exchange_rpc/pb"
	log "github.com/InjectiveLabs/suplog"
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
	Close()
}

func NewExchangeClient(network common.Network, options ...common.ClientOption) (ExchangeClient, error) {
	// process options
	opts := common.DefaultClientOptions()
	if network.ChainTlsCert != nil {
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
		accountClient:            accountPB.NewInjectiveAccountsRPCClient(conn),
		auctionClient:            auctionPB.NewInjectiveAuctionRPCClient(conn),
		oracleClient:             oraclePB.NewInjectiveOracleRPCClient(conn),
		insuranceClient:          insurancePB.NewInjectiveInsuranceRPCClient(conn),
		spotExchangeClient:       spotExchangePB.NewInjectiveSpotExchangeRPCClient(conn),
		derivativeExchangeClient: derivativeExchangePB.NewInjectiveDerivativeExchangeRPCClient(conn),
		portfolioExchangeClient:  portfolioExchangePB.NewInjectivePortfolioRPCClient(conn),

		logger: log.WithFields(log.Fields{
			"module": "sdk-go",
			"svc":    "exchangeClient",
		}),
	}

	return cc, nil
}

type exchangeClient struct {
	opts    *common.ClientOptions
	network common.Network
	conn    *grpc.ClientConn
	logger  log.Logger

	metaClient               metaPB.InjectiveMetaRPCClient
	accountClient            accountPB.InjectiveAccountsRPCClient
	auctionClient            auctionPB.InjectiveAuctionRPCClient
	oracleClient             oraclePB.InjectiveOracleRPCClient
	insuranceClient          insurancePB.InjectiveInsuranceRPCClient
	spotExchangeClient       spotExchangePB.InjectiveSpotExchangeRPCClient
	derivativeExchangeClient derivativeExchangePB.InjectiveDerivativeExchangeRPCClient
	portfolioExchangeClient  portfolioExchangePB.InjectivePortfolioRPCClient
}

func (c *exchangeClient) QueryClient() *grpc.ClientConn {
	return c.conn
}

// Derivatives RPC

func (c *exchangeClient) GetDerivativeOrders(ctx context.Context, req *derivativeExchangePB.OrdersRequest) (*derivativeExchangePB.OrdersResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.derivativeExchangeClient.Orders, req)
	if err != nil {
		fmt.Println(err)
		return &derivativeExchangePB.OrdersResponse{}, err
	}

	return res, nil
}

// Deprecated: Use GetDerivativePositionsV2 instead.
func (c *exchangeClient) GetDerivativePositions(ctx context.Context, req *derivativeExchangePB.PositionsRequest) (*derivativeExchangePB.PositionsResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.derivativeExchangeClient.Positions, req)
	if err != nil {
		fmt.Println(err)
		return &derivativeExchangePB.PositionsResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetDerivativePositionsV2(ctx context.Context, req *derivativeExchangePB.PositionsV2Request) (*derivativeExchangePB.PositionsV2Response, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.derivativeExchangeClient.PositionsV2, req)
	if err != nil {
		fmt.Println(err)
		return &derivativeExchangePB.PositionsV2Response{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetDerivativeLiquidablePositions(ctx context.Context, req *derivativeExchangePB.LiquidablePositionsRequest) (*derivativeExchangePB.LiquidablePositionsResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.derivativeExchangeClient.LiquidablePositions, req)
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

	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.derivativeExchangeClient.OrderbookV2, &req)
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

	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.derivativeExchangeClient.OrderbooksV2, &req)
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

	stream, err := common.ExecuteStreamCall(ctx, c.network.ExchangeCookieAssistant, c.derivativeExchangeClient.StreamOrderbookV2, &req)

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

	stream, err := common.ExecuteStreamCall(ctx, c.network.ExchangeCookieAssistant, c.derivativeExchangeClient.StreamOrderbookUpdate, &req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

func (c *exchangeClient) GetDerivativeMarkets(ctx context.Context, req *derivativeExchangePB.MarketsRequest) (*derivativeExchangePB.MarketsResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.derivativeExchangeClient.Markets, req)
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

	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.derivativeExchangeClient.Market, &req)
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

	stream, err := common.ExecuteStreamCall(ctx, c.network.ExchangeCookieAssistant, c.derivativeExchangeClient.StreamMarket, &req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

func (c *exchangeClient) StreamDerivativePositions(ctx context.Context, req *derivativeExchangePB.StreamPositionsRequest) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamPositionsClient, error) {
	stream, err := common.ExecuteStreamCall(ctx, c.network.ExchangeCookieAssistant, c.derivativeExchangeClient.StreamPositions, req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

func (c *exchangeClient) StreamDerivativeOrders(ctx context.Context, req *derivativeExchangePB.StreamOrdersRequest) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamOrdersClient, error) {
	stream, err := common.ExecuteStreamCall(ctx, c.network.ExchangeCookieAssistant, c.derivativeExchangeClient.StreamOrders, req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

func (c *exchangeClient) GetDerivativeTrades(ctx context.Context, req *derivativeExchangePB.TradesRequest) (*derivativeExchangePB.TradesResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.derivativeExchangeClient.Trades, req)
	if err != nil {
		fmt.Println(err)
		return &derivativeExchangePB.TradesResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetDerivativeTradesV2(ctx context.Context, req *derivativeExchangePB.TradesV2Request) (*derivativeExchangePB.TradesV2Response, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.derivativeExchangeClient.TradesV2, req)
	if err != nil {
		fmt.Println(err)
		return &derivativeExchangePB.TradesV2Response{}, err
	}

	return res, nil
}

func (c *exchangeClient) StreamDerivativeTrades(ctx context.Context, req *derivativeExchangePB.StreamTradesRequest) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamTradesClient, error) {
	stream, err := common.ExecuteStreamCall(ctx, c.network.ExchangeCookieAssistant, c.derivativeExchangeClient.StreamTrades, req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

func (c *exchangeClient) StreamDerivativeV2Trades(ctx context.Context, req *derivativeExchangePB.StreamTradesV2Request) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamTradesV2Client, error) {
	stream, err := common.ExecuteStreamCall(ctx, c.network.ExchangeCookieAssistant, c.derivativeExchangeClient.StreamTradesV2, req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

func (c *exchangeClient) GetSubaccountDerivativeOrdersList(ctx context.Context, req *derivativeExchangePB.SubaccountOrdersListRequest) (*derivativeExchangePB.SubaccountOrdersListResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.derivativeExchangeClient.SubaccountOrdersList, req)

	if err != nil {
		fmt.Println(err)
		return &derivativeExchangePB.SubaccountOrdersListResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetSubaccountDerivativeTradesList(ctx context.Context, req *derivativeExchangePB.SubaccountTradesListRequest) (*derivativeExchangePB.SubaccountTradesListResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.derivativeExchangeClient.SubaccountTradesList, req)

	if err != nil {
		fmt.Println(err)
		return &derivativeExchangePB.SubaccountTradesListResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetHistoricalDerivativeOrders(ctx context.Context, req *derivativeExchangePB.OrdersHistoryRequest) (*derivativeExchangePB.OrdersHistoryResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.derivativeExchangeClient.OrdersHistory, req)
	if err != nil {
		return &derivativeExchangePB.OrdersHistoryResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) StreamHistoricalDerivativeOrders(ctx context.Context, req *derivativeExchangePB.StreamOrdersHistoryRequest) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamOrdersHistoryClient, error) {
	stream, err := common.ExecuteStreamCall(ctx, c.network.ExchangeCookieAssistant, c.derivativeExchangeClient.StreamOrdersHistory, req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

func (c *exchangeClient) GetDerivativeFundingPayments(ctx context.Context, req *derivativeExchangePB.FundingPaymentsRequest) (*derivativeExchangePB.FundingPaymentsResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.derivativeExchangeClient.FundingPayments, req)
	if err != nil {
		fmt.Println(err)
		return &derivativeExchangePB.FundingPaymentsResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetDerivativeFundingRates(ctx context.Context, req *derivativeExchangePB.FundingRatesRequest) (*derivativeExchangePB.FundingRatesResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.derivativeExchangeClient.FundingRates, req)
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

	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.oracleClient.Price, &req)

	if err != nil {
		fmt.Println(err)
		return &oraclePB.PriceResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetOracleList(ctx context.Context) (*oraclePB.OracleListResponse, error) {
	req := oraclePB.OracleListRequest{}

	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.oracleClient.OracleList, &req)
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

	stream, err := common.ExecuteStreamCall(ctx, c.network.ExchangeCookieAssistant, c.oracleClient.StreamPrices, &req)

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

	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.auctionClient.AuctionEndpoint, &req)
	if err != nil {
		fmt.Println(err)
		return &auctionPB.AuctionEndpointResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetAuctions(ctx context.Context) (*auctionPB.AuctionsResponse, error) {
	req := auctionPB.AuctionsRequest{}

	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.auctionClient.Auctions, &req)
	if err != nil {
		fmt.Println(err)
		return &auctionPB.AuctionsResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) StreamBids(ctx context.Context) (auctionPB.InjectiveAuctionRPC_StreamBidsClient, error) {
	req := auctionPB.StreamBidsRequest{}

	stream, err := common.ExecuteStreamCall(ctx, c.network.ExchangeCookieAssistant, c.auctionClient.StreamBids, &req)

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

	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.accountClient.SubaccountsList, &req)

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

	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.accountClient.SubaccountBalanceEndpoint, &req)

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

	stream, err := common.ExecuteStreamCall(ctx, c.network.ExchangeCookieAssistant, c.accountClient.StreamSubaccountBalance, &req)

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

	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.accountClient.SubaccountBalancesList, &req)

	if err != nil {
		fmt.Println(err)
		return &accountPB.SubaccountBalancesListResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetSubaccountHistory(ctx context.Context, req *accountPB.SubaccountHistoryRequest) (*accountPB.SubaccountHistoryResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.accountClient.SubaccountHistory, req)

	if err != nil {
		fmt.Println(err)
		return &accountPB.SubaccountHistoryResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetSubaccountOrderSummary(ctx context.Context, req *accountPB.SubaccountOrderSummaryRequest) (*accountPB.SubaccountOrderSummaryResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.accountClient.SubaccountOrderSummary, req)

	if err != nil {
		fmt.Println(err)
		return &accountPB.SubaccountOrderSummaryResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetOrderStates(ctx context.Context, req *accountPB.OrderStatesRequest) (*accountPB.OrderStatesResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.accountClient.OrderStates, req)
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

	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.accountClient.Portfolio, &req)

	if err != nil {
		fmt.Println(err)
		return &accountPB.PortfolioResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetRewards(ctx context.Context, req *accountPB.RewardsRequest) (*accountPB.RewardsResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.accountClient.Rewards, req)

	if err != nil {
		fmt.Println(err)
		return &accountPB.RewardsResponse{}, err
	}

	return res, nil
}

// Spot RPC

func (c *exchangeClient) GetSpotOrders(ctx context.Context, req *spotExchangePB.OrdersRequest) (*spotExchangePB.OrdersResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.spotExchangeClient.Orders, req)

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

	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.spotExchangeClient.OrderbookV2, &req)

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

	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.spotExchangeClient.OrderbooksV2, &req)

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

	stream, err := common.ExecuteStreamCall(ctx, c.network.ExchangeCookieAssistant, c.spotExchangeClient.StreamOrderbookUpdate, &req)

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

	stream, err := common.ExecuteStreamCall(ctx, c.network.ExchangeCookieAssistant, c.spotExchangeClient.StreamOrderbookV2, &req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

func (c *exchangeClient) GetSpotMarkets(ctx context.Context, req *spotExchangePB.MarketsRequest) (*spotExchangePB.MarketsResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.spotExchangeClient.Markets, req)

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

	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.spotExchangeClient.Market, &req)

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

	stream, err := common.ExecuteStreamCall(ctx, c.network.ExchangeCookieAssistant, c.spotExchangeClient.StreamMarkets, &req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

func (c *exchangeClient) StreamSpotOrders(ctx context.Context, req *spotExchangePB.StreamOrdersRequest) (spotExchangePB.InjectiveSpotExchangeRPC_StreamOrdersClient, error) {
	stream, err := common.ExecuteStreamCall(ctx, c.network.ExchangeCookieAssistant, c.spotExchangeClient.StreamOrders, req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

func (c *exchangeClient) GetSpotTrades(ctx context.Context, req *spotExchangePB.TradesRequest) (*spotExchangePB.TradesResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.spotExchangeClient.Trades, req)

	if err != nil {
		fmt.Println(err)
		return &spotExchangePB.TradesResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetSpotTradesV2(ctx context.Context, req *spotExchangePB.TradesV2Request) (*spotExchangePB.TradesV2Response, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.spotExchangeClient.TradesV2, req)

	if err != nil {
		fmt.Println(err)
		return &spotExchangePB.TradesV2Response{}, err
	}

	return res, nil
}

func (c *exchangeClient) StreamSpotTrades(ctx context.Context, req *spotExchangePB.StreamTradesRequest) (spotExchangePB.InjectiveSpotExchangeRPC_StreamTradesClient, error) {
	stream, err := common.ExecuteStreamCall(ctx, c.network.ExchangeCookieAssistant, c.spotExchangeClient.StreamTrades, req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

func (c *exchangeClient) StreamSpotTradesV2(ctx context.Context, req *spotExchangePB.StreamTradesV2Request) (spotExchangePB.InjectiveSpotExchangeRPC_StreamTradesV2Client, error) {
	stream, err := common.ExecuteStreamCall(ctx, c.network.ExchangeCookieAssistant, c.spotExchangeClient.StreamTradesV2, req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

func (c *exchangeClient) GetSubaccountSpotOrdersList(ctx context.Context, req *spotExchangePB.SubaccountOrdersListRequest) (*spotExchangePB.SubaccountOrdersListResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.spotExchangeClient.SubaccountOrdersList, req)

	if err != nil {
		fmt.Println(err)
		return &spotExchangePB.SubaccountOrdersListResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetSubaccountSpotTradesList(ctx context.Context, req *spotExchangePB.SubaccountTradesListRequest) (*spotExchangePB.SubaccountTradesListResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.spotExchangeClient.SubaccountTradesList, req)

	if err != nil {
		fmt.Println(err)
		return &spotExchangePB.SubaccountTradesListResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetHistoricalSpotOrders(ctx context.Context, req *spotExchangePB.OrdersHistoryRequest) (*spotExchangePB.OrdersHistoryResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.spotExchangeClient.OrdersHistory, req)
	if err != nil {
		return &spotExchangePB.OrdersHistoryResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) StreamHistoricalSpotOrders(ctx context.Context, req *spotExchangePB.StreamOrdersHistoryRequest) (spotExchangePB.InjectiveSpotExchangeRPC_StreamOrdersHistoryClient, error) {
	stream, err := common.ExecuteStreamCall(ctx, c.network.ExchangeCookieAssistant, c.spotExchangeClient.StreamOrdersHistory, req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

func (c *exchangeClient) GetInsuranceFunds(ctx context.Context, req *insurancePB.FundsRequest) (*insurancePB.FundsResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.insuranceClient.Funds, req)
	if err != nil {
		fmt.Println(err)
		return &insurancePB.FundsResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetRedemptions(ctx context.Context, req *insurancePB.RedemptionsRequest) (*insurancePB.RedemptionsResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.insuranceClient.Redemptions, req)

	if err != nil {
		fmt.Println(err)
		return &insurancePB.RedemptionsResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) Ping(ctx context.Context, req *metaPB.PingRequest) (*metaPB.PingResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.metaClient.Ping, req)

	if err != nil {
		fmt.Println(err)
		return &metaPB.PingResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetVersion(ctx context.Context, req *metaPB.VersionRequest) (*metaPB.VersionResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.metaClient.Version, req)

	if err != nil {
		fmt.Println(err)
		return &metaPB.VersionResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetInfo(ctx context.Context, req *metaPB.InfoRequest) (*metaPB.InfoResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.metaClient.Info, req)
	if err != nil {
		fmt.Println(err)
		return &metaPB.InfoResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) StreamKeepalive(ctx context.Context) (metaPB.InjectiveMetaRPC_StreamKeepaliveClient, error) {
	req := metaPB.StreamKeepaliveRequest{}

	stream, err := common.ExecuteStreamCall(ctx, c.network.ExchangeCookieAssistant, c.metaClient.StreamKeepalive, &req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

// Deprecated: Use GetAccountPortfolioBalances instead.
func (c *exchangeClient) GetAccountPortfolio(ctx context.Context, accountAddress string) (*portfolioExchangePB.AccountPortfolioResponse, error) {
	req := &portfolioExchangePB.AccountPortfolioRequest{
		AccountAddress: accountAddress,
	}
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.portfolioExchangeClient.AccountPortfolio, req)
	if err != nil {
		fmt.Println(err)
		return &portfolioExchangePB.AccountPortfolioResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) GetAccountPortfolioBalances(ctx context.Context, accountAddress string) (*portfolioExchangePB.AccountPortfolioBalancesResponse, error) {
	req := &portfolioExchangePB.AccountPortfolioBalancesRequest{
		AccountAddress: accountAddress,
	}
	res, err := common.ExecuteCall(ctx, c.network.ExchangeCookieAssistant, c.portfolioExchangeClient.AccountPortfolioBalances, req)
	if err != nil {
		fmt.Println(err)
		return &portfolioExchangePB.AccountPortfolioBalancesResponse{}, err
	}

	return res, nil
}

func (c *exchangeClient) StreamAccountPortfolio(ctx context.Context, accountAddress string, subaccountId, balanceType string) (portfolioExchangePB.InjectivePortfolioRPC_StreamAccountPortfolioClient, error) {
	req := &portfolioExchangePB.StreamAccountPortfolioRequest{
		AccountAddress: accountAddress,
		SubaccountId:   subaccountId,
		Type:           balanceType,
	}
	stream, err := common.ExecuteStreamCall(ctx, c.network.ExchangeCookieAssistant, c.portfolioExchangeClient.StreamAccountPortfolio, req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

func (c *exchangeClient) Close() {
	c.conn.Close()
}
