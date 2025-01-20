package exchange

import (
	"context"
	"errors"

	"github.com/InjectiveLabs/sdk-go/client/common"

	accountPB "github.com/InjectiveLabs/sdk-go/exchange/accounts_rpc/pb"
	auctionPB "github.com/InjectiveLabs/sdk-go/exchange/auction_rpc/pb"
	derivativeExchangePB "github.com/InjectiveLabs/sdk-go/exchange/derivative_exchange_rpc/pb"
	insurancePB "github.com/InjectiveLabs/sdk-go/exchange/insurance_rpc/pb"
	metaPB "github.com/InjectiveLabs/sdk-go/exchange/meta_rpc/pb"
	oraclePB "github.com/InjectiveLabs/sdk-go/exchange/oracle_rpc/pb"
	portfolioExchangePB "github.com/InjectiveLabs/sdk-go/exchange/portfolio_rpc/pb"
	spotExchangePB "github.com/InjectiveLabs/sdk-go/exchange/spot_exchange_rpc/pb"
	"google.golang.org/grpc"
)

type MockExchangeClient struct {
	Network                    common.Network
	SpotMarketsResponses       []*spotExchangePB.MarketsResponse
	DerivativeMarketsResponses []*derivativeExchangePB.MarketsResponse
}

var _ ExchangeClient = &MockExchangeClient{}

func (e *MockExchangeClient) QueryClient() *grpc.ClientConn {
	dummyConnection := grpc.ClientConn{}
	return &dummyConnection
}

func (e *MockExchangeClient) GetDerivativeMarket(ctx context.Context, marketId string) (*derivativeExchangePB.MarketResponse, error) {
	return &derivativeExchangePB.MarketResponse{}, nil
}

func (e *MockExchangeClient) GetDerivativeOrderbookV2(ctx context.Context, marketId string) (*derivativeExchangePB.OrderbookV2Response, error) {
	return &derivativeExchangePB.OrderbookV2Response{}, nil
}

func (e *MockExchangeClient) GetDerivativeOrderbooksV2(ctx context.Context, marketIDs []string) (*derivativeExchangePB.OrderbooksV2Response, error) {
	return &derivativeExchangePB.OrderbooksV2Response{}, nil
}

func (e *MockExchangeClient) StreamDerivativeOrderbookV2(ctx context.Context, marketIDs []string) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamOrderbookV2Client, error) {
	return nil, nil
}

func (e *MockExchangeClient) StreamDerivativeOrderbookUpdate(ctx context.Context, marketIDs []string) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamOrderbookUpdateClient, error) {
	return nil, nil
}

func (e *MockExchangeClient) StreamDerivativeMarket(ctx context.Context, marketIDs []string) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamMarketClient, error) {
	return nil, nil
}

func (e *MockExchangeClient) GetDerivativeOrders(ctx context.Context, req *derivativeExchangePB.OrdersRequest) (*derivativeExchangePB.OrdersResponse, error) {
	return &derivativeExchangePB.OrdersResponse{}, nil
}

func (e *MockExchangeClient) GetDerivativeMarkets(ctx context.Context, req *derivativeExchangePB.MarketsRequest) (*derivativeExchangePB.MarketsResponse, error) {
	var response *derivativeExchangePB.MarketsResponse
	var localError error
	if len(e.DerivativeMarketsResponses) > 0 {
		response = e.DerivativeMarketsResponses[0]
		e.DerivativeMarketsResponses = e.DerivativeMarketsResponses[1:]
		localError = nil
	} else {
		response = &derivativeExchangePB.MarketsResponse{}
		localError = errors.New("There are no responses configured")
	}

	return response, localError
}

func (e *MockExchangeClient) GetDerivativePositions(ctx context.Context, req *derivativeExchangePB.PositionsRequest) (*derivativeExchangePB.PositionsResponse, error) {
	return &derivativeExchangePB.PositionsResponse{}, nil
}

func (e *MockExchangeClient) GetDerivativePositionsV2(ctx context.Context, req *derivativeExchangePB.PositionsV2Request) (*derivativeExchangePB.PositionsV2Response, error) {
	return &derivativeExchangePB.PositionsV2Response{}, nil
}

func (e *MockExchangeClient) GetDerivativeLiquidablePositions(ctx context.Context, req *derivativeExchangePB.LiquidablePositionsRequest) (*derivativeExchangePB.LiquidablePositionsResponse, error) {
	return &derivativeExchangePB.LiquidablePositionsResponse{}, nil
}

func (e *MockExchangeClient) StreamDerivativePositions(ctx context.Context, req *derivativeExchangePB.StreamPositionsRequest) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamPositionsClient, error) {
	return nil, nil
}

func (e *MockExchangeClient) StreamDerivativePositionsV2(ctx context.Context, req *derivativeExchangePB.StreamPositionsV2Request) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamPositionsV2Client, error) {
	return nil, nil
}

func (e *MockExchangeClient) StreamDerivativeOrders(ctx context.Context, req *derivativeExchangePB.StreamOrdersRequest) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamOrdersClient, error) {
	return nil, nil
}

func (e *MockExchangeClient) GetDerivativeTrades(ctx context.Context, req *derivativeExchangePB.TradesRequest) (*derivativeExchangePB.TradesResponse, error) {
	return &derivativeExchangePB.TradesResponse{}, nil
}

func (e *MockExchangeClient) GetDerivativeTradesV2(ctx context.Context, req *derivativeExchangePB.TradesV2Request) (*derivativeExchangePB.TradesV2Response, error) {
	return &derivativeExchangePB.TradesV2Response{}, nil
}

func (e *MockExchangeClient) StreamDerivativeTrades(ctx context.Context, req *derivativeExchangePB.StreamTradesRequest) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamTradesClient, error) {
	return nil, nil
}

func (e *MockExchangeClient) StreamDerivativeV2Trades(ctx context.Context, req *derivativeExchangePB.StreamTradesV2Request) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamTradesV2Client, error) {
	return nil, nil
}

func (e *MockExchangeClient) GetSubaccountDerivativeOrdersList(ctx context.Context, req *derivativeExchangePB.SubaccountOrdersListRequest) (*derivativeExchangePB.SubaccountOrdersListResponse, error) {
	return &derivativeExchangePB.SubaccountOrdersListResponse{}, nil
}

func (e *MockExchangeClient) GetSubaccountDerivativeTradesList(ctx context.Context, req *derivativeExchangePB.SubaccountTradesListRequest) (*derivativeExchangePB.SubaccountTradesListResponse, error) {
	return &derivativeExchangePB.SubaccountTradesListResponse{}, nil
}

func (e *MockExchangeClient) GetHistoricalDerivativeOrders(ctx context.Context, req *derivativeExchangePB.OrdersHistoryRequest) (*derivativeExchangePB.OrdersHistoryResponse, error) {
	return &derivativeExchangePB.OrdersHistoryResponse{}, nil
}

func (e *MockExchangeClient) StreamHistoricalDerivativeOrders(ctx context.Context, req *derivativeExchangePB.StreamOrdersHistoryRequest) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamOrdersHistoryClient, error) {
	return nil, nil
}

func (e *MockExchangeClient) GetDerivativeFundingPayments(ctx context.Context, req *derivativeExchangePB.FundingPaymentsRequest) (*derivativeExchangePB.FundingPaymentsResponse, error) {
	return &derivativeExchangePB.FundingPaymentsResponse{}, nil
}

func (e *MockExchangeClient) GetDerivativeFundingRates(ctx context.Context, req *derivativeExchangePB.FundingRatesRequest) (*derivativeExchangePB.FundingRatesResponse, error) {
	return &derivativeExchangePB.FundingRatesResponse{}, nil
}

func (e *MockExchangeClient) GetPrice(ctx context.Context, baseSymbol, quoteSymbol, oracleType string, oracleScaleFactor uint32) (*oraclePB.PriceResponse, error) {
	return &oraclePB.PriceResponse{}, nil
}

func (e *MockExchangeClient) GetOracleList(ctx context.Context) (*oraclePB.OracleListResponse, error) {
	return &oraclePB.OracleListResponse{}, nil
}

func (e *MockExchangeClient) StreamPrices(ctx context.Context, baseSymbol, quoteSymbol, oracleType string) (oraclePB.InjectiveOracleRPC_StreamPricesClient, error) {
	return nil, nil
}

func (e *MockExchangeClient) GetAuction(ctx context.Context, round int64) (*auctionPB.AuctionEndpointResponse, error) {
	return &auctionPB.AuctionEndpointResponse{}, nil
}

func (e *MockExchangeClient) GetAuctions(ctx context.Context) (*auctionPB.AuctionsResponse, error) {
	return &auctionPB.AuctionsResponse{}, nil
}

func (e *MockExchangeClient) StreamBids(ctx context.Context) (auctionPB.InjectiveAuctionRPC_StreamBidsClient, error) {
	return nil, nil
}

func (e *MockExchangeClient) FetchInjBurnt(ctx context.Context) (*auctionPB.InjBurntEndpointResponse, error) {
	return &auctionPB.InjBurntEndpointResponse{}, nil
}

func (e *MockExchangeClient) GetSubaccountsList(ctx context.Context, accountAddress string) (*accountPB.SubaccountsListResponse, error) {
	return &accountPB.SubaccountsListResponse{}, nil
}

func (e *MockExchangeClient) GetSubaccountBalance(ctx context.Context, subaccountId, denom string) (*accountPB.SubaccountBalanceEndpointResponse, error) {
	return &accountPB.SubaccountBalanceEndpointResponse{}, nil
}

func (e *MockExchangeClient) StreamSubaccountBalance(ctx context.Context, subaccountId string) (accountPB.InjectiveAccountsRPC_StreamSubaccountBalanceClient, error) {
	return nil, nil
}

func (e *MockExchangeClient) GetSubaccountBalancesList(ctx context.Context, subaccountId string) (*accountPB.SubaccountBalancesListResponse, error) {
	return &accountPB.SubaccountBalancesListResponse{}, nil
}

func (e *MockExchangeClient) GetSubaccountHistory(ctx context.Context, req *accountPB.SubaccountHistoryRequest) (*accountPB.SubaccountHistoryResponse, error) {
	return &accountPB.SubaccountHistoryResponse{}, nil
}

func (e *MockExchangeClient) GetSubaccountOrderSummary(ctx context.Context, req *accountPB.SubaccountOrderSummaryRequest) (*accountPB.SubaccountOrderSummaryResponse, error) {
	return &accountPB.SubaccountOrderSummaryResponse{}, nil
}

func (e *MockExchangeClient) GetOrderStates(ctx context.Context, req *accountPB.OrderStatesRequest) (*accountPB.OrderStatesResponse, error) {
	return &accountPB.OrderStatesResponse{}, nil
}

func (e *MockExchangeClient) GetPortfolio(ctx context.Context, accountAddress string) (*accountPB.PortfolioResponse, error) {
	return &accountPB.PortfolioResponse{}, nil
}

func (e *MockExchangeClient) GetRewards(ctx context.Context, req *accountPB.RewardsRequest) (*accountPB.RewardsResponse, error) {
	return &accountPB.RewardsResponse{}, nil
}

func (e *MockExchangeClient) GetSpotOrders(ctx context.Context, req *spotExchangePB.OrdersRequest) (*spotExchangePB.OrdersResponse, error) {
	return &spotExchangePB.OrdersResponse{}, nil
}

func (e *MockExchangeClient) GetSpotOrderbookV2(ctx context.Context, marketId string) (*spotExchangePB.OrderbookV2Response, error) {
	return &spotExchangePB.OrderbookV2Response{}, nil
}

func (e *MockExchangeClient) GetSpotOrderbooksV2(ctx context.Context, marketIDs []string) (*spotExchangePB.OrderbooksV2Response, error) {
	return &spotExchangePB.OrderbooksV2Response{}, nil
}

func (e *MockExchangeClient) StreamSpotOrderbookV2(ctx context.Context, marketIDs []string) (spotExchangePB.InjectiveSpotExchangeRPC_StreamOrderbookV2Client, error) {
	return nil, nil
}

func (e *MockExchangeClient) StreamSpotOrderbookUpdate(ctx context.Context, marketIDs []string) (spotExchangePB.InjectiveSpotExchangeRPC_StreamOrderbookUpdateClient, error) {
	return nil, nil
}

func (e *MockExchangeClient) GetSpotMarkets(ctx context.Context, req *spotExchangePB.MarketsRequest) (*spotExchangePB.MarketsResponse, error) {
	var response *spotExchangePB.MarketsResponse
	var localError error
	if len(e.SpotMarketsResponses) > 0 {
		response = e.SpotMarketsResponses[0]
		e.SpotMarketsResponses = e.SpotMarketsResponses[1:]
		localError = nil
	} else {
		response = &spotExchangePB.MarketsResponse{}
		localError = errors.New("There are no responses configured")
	}

	return response, localError
}

func (e *MockExchangeClient) GetSpotMarket(ctx context.Context, marketId string) (*spotExchangePB.MarketResponse, error) {
	return &spotExchangePB.MarketResponse{}, nil
}

func (e *MockExchangeClient) StreamSpotMarket(ctx context.Context, marketIDs []string) (spotExchangePB.InjectiveSpotExchangeRPC_StreamMarketsClient, error) {
	return nil, nil
}

func (e *MockExchangeClient) StreamSpotOrders(ctx context.Context, req *spotExchangePB.StreamOrdersRequest) (spotExchangePB.InjectiveSpotExchangeRPC_StreamOrdersClient, error) {
	return nil, nil
}

func (e *MockExchangeClient) GetSpotTrades(ctx context.Context, req *spotExchangePB.TradesRequest) (*spotExchangePB.TradesResponse, error) {
	return &spotExchangePB.TradesResponse{}, nil
}

func (e *MockExchangeClient) GetSpotTradesV2(ctx context.Context, req *spotExchangePB.TradesV2Request) (*spotExchangePB.TradesV2Response, error) {
	return &spotExchangePB.TradesV2Response{}, nil
}

func (e *MockExchangeClient) StreamSpotTrades(ctx context.Context, req *spotExchangePB.StreamTradesRequest) (spotExchangePB.InjectiveSpotExchangeRPC_StreamTradesClient, error) {
	return nil, nil
}

func (e *MockExchangeClient) StreamSpotTradesV2(ctx context.Context, req *spotExchangePB.StreamTradesV2Request) (spotExchangePB.InjectiveSpotExchangeRPC_StreamTradesV2Client, error) {
	return nil, nil
}

func (e *MockExchangeClient) GetSubaccountSpotOrdersList(ctx context.Context, req *spotExchangePB.SubaccountOrdersListRequest) (*spotExchangePB.SubaccountOrdersListResponse, error) {
	return &spotExchangePB.SubaccountOrdersListResponse{}, nil
}

func (e *MockExchangeClient) GetSubaccountSpotTradesList(ctx context.Context, req *spotExchangePB.SubaccountTradesListRequest) (*spotExchangePB.SubaccountTradesListResponse, error) {
	return &spotExchangePB.SubaccountTradesListResponse{}, nil
}

func (e *MockExchangeClient) GetHistoricalSpotOrders(ctx context.Context, req *spotExchangePB.OrdersHistoryRequest) (*spotExchangePB.OrdersHistoryResponse, error) {
	return &spotExchangePB.OrdersHistoryResponse{}, nil
}

func (e *MockExchangeClient) StreamHistoricalSpotOrders(ctx context.Context, req *spotExchangePB.StreamOrdersHistoryRequest) (spotExchangePB.InjectiveSpotExchangeRPC_StreamOrdersHistoryClient, error) {
	return nil, nil
}

func (e *MockExchangeClient) GetInsuranceFunds(ctx context.Context, req *insurancePB.FundsRequest) (*insurancePB.FundsResponse, error) {
	return &insurancePB.FundsResponse{}, nil
}

func (e *MockExchangeClient) GetRedemptions(ctx context.Context, req *insurancePB.RedemptionsRequest) (*insurancePB.RedemptionsResponse, error) {
	return &insurancePB.RedemptionsResponse{}, nil
}

func (e *MockExchangeClient) GetAccountPortfolio(ctx context.Context, accountAddress string) (*portfolioExchangePB.AccountPortfolioResponse, error) {
	return &portfolioExchangePB.AccountPortfolioResponse{}, nil
}

func (e *MockExchangeClient) GetAccountPortfolioBalances(ctx context.Context, accountAddress string) (*portfolioExchangePB.AccountPortfolioBalancesResponse, error) {
	return &portfolioExchangePB.AccountPortfolioBalancesResponse{}, nil
}

func (e *MockExchangeClient) StreamAccountPortfolio(ctx context.Context, accountAddress, subaccountId, balanceType string) (portfolioExchangePB.InjectivePortfolioRPC_StreamAccountPortfolioClient, error) {
	return nil, nil
}

func (e *MockExchangeClient) StreamKeepalive(ctx context.Context) (metaPB.InjectiveMetaRPC_StreamKeepaliveClient, error) {
	return nil, nil
}

func (e *MockExchangeClient) GetInfo(ctx context.Context, req *metaPB.InfoRequest) (*metaPB.InfoResponse, error) {
	return &metaPB.InfoResponse{}, nil
}

func (e *MockExchangeClient) GetVersion(ctx context.Context, req *metaPB.VersionRequest) (*metaPB.VersionResponse, error) {
	return &metaPB.VersionResponse{}, nil
}

func (e *MockExchangeClient) Ping(ctx context.Context, req *metaPB.PingRequest) (*metaPB.PingResponse, error) {
	return &metaPB.PingResponse{}, nil
}

func (e *MockExchangeClient) GetNetwork() common.Network {
	return e.Network
}

func (e *MockExchangeClient) Close() {

}
