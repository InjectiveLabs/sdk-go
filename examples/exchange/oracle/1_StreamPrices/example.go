package main

import (
	"context"
	"encoding/json"
	"fmt"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"

	"github.com/InjectiveLabs/sdk-go/client/common"
	exchangeclient "github.com/InjectiveLabs/sdk-go/client/exchange"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")
	exchangeClient, err := exchangeclient.NewExchangeClient(network)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	marketsAssistant, err := chainclient.NewMarketsAssistantInitializedFromChain(ctx, exchangeClient)
	market := marketsAssistant.AllDerivativeMarkets()["0x17ef48032cb24375ba7c2e39f384e56433bcab20cbee9a7357e4cba2eb00abe6"]

	baseSymbol := market.OracleBase
	quoteSymbol := market.OracleQuote
	oracleType := market.OracleType
	stream, err := exchangeClient.StreamPrices(ctx, baseSymbol, quoteSymbol, oracleType)
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
				fmt.Println(err)
				return
			}
			str, _ := json.MarshalIndent(res, "", " ")
			fmt.Print(string(str))
		}
	}
}
