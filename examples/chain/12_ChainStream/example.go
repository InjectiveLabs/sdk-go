package main

import (
	"context"
	"encoding/json"
	"fmt"

	chainstreamv2 "github.com/InjectiveLabs/sdk-go/chain/stream/types/v2"
	"github.com/InjectiveLabs/sdk-go/client"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
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

	chainClient, err := chainclient.NewChainClientV2(
		clientCtx,
		network,
		common.OptionGasPrices(client.DefaultGasPriceWithDenom),
	)

	if err != nil {
		panic(err)
	}

	subaccountId := "0xbdaedec95d563fb05240d6e01821008454c24c36000000000000000000000000"

	injUsdtMarket := "0x0611780ba69656949525013d947713300f56c37b6175e02f26bffa495c3208fe"
	injUsdtPerpMarket := "0x17ef48032cb24375ba7c2e39f384e56433bcab20cbee9a7357e4cba2eb00abe6"

	req := chainstreamv2.StreamRequest{
		BankBalancesFilter: &chainstreamv2.BankBalancesFilter{
			Accounts: []string{"*"},
		},
		SpotOrdersFilter: &chainstreamv2.OrdersFilter{
			MarketIds:     []string{injUsdtMarket},
			SubaccountIds: []string{subaccountId},
		},
		DerivativeOrdersFilter: &chainstreamv2.OrdersFilter{
			MarketIds:     []string{injUsdtPerpMarket},
			SubaccountIds: []string{subaccountId},
		},
		SpotTradesFilter: &chainstreamv2.TradesFilter{
			MarketIds:     []string{injUsdtMarket},
			SubaccountIds: []string{"*"},
		},
		SubaccountDepositsFilter: &chainstreamv2.SubaccountDepositsFilter{
			SubaccountIds: []string{subaccountId},
		},
		DerivativeOrderbooksFilter: &chainstreamv2.OrderbookFilter{
			MarketIds: []string{injUsdtPerpMarket},
		},
		SpotOrderbooksFilter: &chainstreamv2.OrderbookFilter{
			MarketIds: []string{injUsdtMarket},
		},
		PositionsFilter: &chainstreamv2.PositionsFilter{
			SubaccountIds: []string{subaccountId},
			MarketIds:     []string{injUsdtPerpMarket},
		},
		DerivativeTradesFilter: &chainstreamv2.TradesFilter{
			SubaccountIds: []string{"*"},
			MarketIds:     []string{injUsdtPerpMarket},
		},
		OraclePriceFilter: &chainstreamv2.OraclePriceFilter{
			Symbol: []string{"INJ", "USDT"},
		},
	}

	ctx := context.Background()

	stream, err := chainClient.ChainStreamV2(ctx, req)
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
			}
			str, _ := json.MarshalIndent(res, "", "\t")
			fmt.Print(string(str))
		}
	}
}
