package main

import (
	"context"
	"encoding/json"
	"fmt"

	chainStreamModule "github.com/InjectiveLabs/sdk-go/chain/stream/types"
	"github.com/InjectiveLabs/sdk-go/client"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
	"github.com/sirupsen/logrus"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")

	clientCtx, err := chainclient.NewClientContext(
		network.ChainId,
		"",
		nil,
	)
	if err != nil {
		panic(err)
	}
	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint)

	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network,
		common.OptionGasPrices(client.DefaultGasPriceWithDenom),
		common.OptionLogger(logrus.New()),
	)

	if err != nil {
		panic(err)
	}

	subaccountId := "0xbdaedec95d563fb05240d6e01821008454c24c36000000000000000000000000"

	injUsdtMarket := "0x0611780ba69656949525013d947713300f56c37b6175e02f26bffa495c3208fe"
	injUsdtPerpMarket := "0x17ef48032cb24375ba7c2e39f384e56433bcab20cbee9a7357e4cba2eb00abe6"

	req := chainStreamModule.StreamRequest{
		BankBalancesFilter: &chainStreamModule.BankBalancesFilter{
			Accounts: []string{"*"},
		},
		SpotOrdersFilter: &chainStreamModule.OrdersFilter{
			MarketIds:     []string{injUsdtMarket},
			SubaccountIds: []string{subaccountId},
		},
		DerivativeOrdersFilter: &chainStreamModule.OrdersFilter{
			MarketIds:     []string{injUsdtPerpMarket},
			SubaccountIds: []string{subaccountId},
		},
		SpotTradesFilter: &chainStreamModule.TradesFilter{
			MarketIds:     []string{injUsdtMarket},
			SubaccountIds: []string{"*"},
		},
		SubaccountDepositsFilter: &chainStreamModule.SubaccountDepositsFilter{
			SubaccountIds: []string{subaccountId},
		},
		DerivativeOrderbooksFilter: &chainStreamModule.OrderbookFilter{
			MarketIds: []string{injUsdtPerpMarket},
		},
		SpotOrderbooksFilter: &chainStreamModule.OrderbookFilter{
			MarketIds: []string{injUsdtMarket},
		},
		PositionsFilter: &chainStreamModule.PositionsFilter{
			SubaccountIds: []string{subaccountId},
			MarketIds:     []string{injUsdtPerpMarket},
		},
		DerivativeTradesFilter: &chainStreamModule.TradesFilter{
			SubaccountIds: []string{"*"},
			MarketIds:     []string{injUsdtPerpMarket},
		},
		OraclePriceFilter: &chainStreamModule.OraclePriceFilter{
			Symbol: []string{"INJ", "USDT"},
		},
	}

	ctx := context.Background()

	stream, err := chainClient.ChainStream(ctx, req)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case <-ctx.Done():
			return
		default:
			res, err := stream.Recv()
			if err != nil {
				panic(err)
				return
			}
			str, _ := json.MarshalIndent(res, "", " ")
			fmt.Print(string(str))
		}
	}
}
