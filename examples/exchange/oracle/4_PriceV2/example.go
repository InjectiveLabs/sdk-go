package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

	"github.com/InjectiveLabs/sdk-go/client"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
	"github.com/InjectiveLabs/sdk-go/client/core"
	exchangeclient "github.com/InjectiveLabs/sdk-go/client/exchange"
	oraclepb "github.com/InjectiveLabs/sdk-go/exchange/oracle_rpc/pb"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")
	tmClient, err := rpchttp.New(network.TmEndpoint)
	if err != nil {
		panic(err)
	}

	senderAddress, cosmosKeyring, err := chainclient.InitCosmosKeyring(
		os.Getenv("HOME")+"/.injectived",
		"injectived",
		"file",
		"inj-user",
		"12345678",
		"5d386fbdbf11f1141010f81a46b40f94887367562bd33b452bbaa6ce1cd1381e", // keyring will be used if pk not provided
		false,
	)

	if err != nil {
		panic(err)
	}

	clientCtx, err := chainclient.NewClientContext(
		network.ChainId,
		senderAddress.String(),
		cosmosKeyring,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint).WithClient(tmClient)

	chainClient, err := chainclient.NewChainClientV2(
		clientCtx,
		network,
		common.OptionGasPrices(client.DefaultGasPriceWithDenom),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	ctx := context.Background()
	marketsAssistant, err := chainclient.NewHumanReadableMarketsAssistant(ctx, chainClient)
	if err != nil {
		panic(err)
	}

	exchangeClient, err := exchangeclient.NewExchangeClient(network)
	if err != nil {
		panic(err)
	}

	market := marketsAssistant.AllDerivativeMarkets()["0x17ef48032cb24375ba7c2e39f384e56433bcab20cbee9a7357e4cba2eb00abe6"]
	derivativeMarket, ok := market.(core.DerivativeMarketV2)
	if !ok {
		panic("market is not a derivative market")
	}

	baseSymbol := derivativeMarket.OracleBase
	quoteSymbol := derivativeMarket.OracleQuote
	oracleType := derivativeMarket.OracleType
	oracleScaleFactor := uint32(0)
	filters := []*oraclepb.PricePayloadV2{
		{
			BaseSymbol:        baseSymbol,
			QuoteSymbol:       quoteSymbol,
			OracleType:        oracleType,
			OracleScaleFactor: oracleScaleFactor,
		},
	}

	res, err := exchangeClient.FetchPriceV2(ctx, filters)
	if err != nil {
		fmt.Println(err)
	}

	str, _ := json.MarshalIndent(res, "", "\t")
	fmt.Print(string(str))
}
