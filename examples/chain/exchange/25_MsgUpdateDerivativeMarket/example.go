package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"cosmossdk.io/math"

	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	exchangeclient "github.com/InjectiveLabs/sdk-go/client/exchange"

	"github.com/InjectiveLabs/sdk-go/client"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")
	tmClient, err := rpchttp.New(network.TmEndpoint, "/websocket")
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
		panic(err)
	}

	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint).WithClient(tmClient)

	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network,
		common.OptionGasPrices(client.DefaultGasPriceWithDenom),
	)
	if err != nil {
		panic(err)
	}

	exchangeClient, err := exchangeclient.NewExchangeClient(network)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	marketsAssistant, err := chainclient.NewMarketsAssistantInitializedFromChain(ctx, exchangeClient)
	if err != nil {
		panic(err)
	}

	quoteToken := marketsAssistant.AllTokens()["USDT"]
	minPriceTickSize := math.LegacyMustNewDecFromStr("0.1")
	minQuantityTickSize := math.LegacyMustNewDecFromStr("0.1")
	minNotional := math.LegacyMustNewDecFromStr("2")

	chainMinPriceTickSize := minPriceTickSize.Mul(math.LegacyNewDecFromIntWithPrec(math.NewInt(1), int64(quoteToken.Decimals)))
	chainMinQuantityTickSize := minQuantityTickSize
	chainMinNotional := minNotional.Mul(math.LegacyNewDecFromIntWithPrec(math.NewInt(1), int64(quoteToken.Decimals)))

	msg := &exchangetypes.MsgUpdateDerivativeMarket{
		Admin:                     senderAddress.String(),
		MarketId:                  "0x17ef48032cb24375ba7c2e39f384e56433bcab20cbee9a7357e4cba2eb00abe6",
		NewTicker:                 "INJ/USDT PERP 2",
		NewMinPriceTickSize:       chainMinPriceTickSize,
		NewMinQuantityTickSize:    chainMinQuantityTickSize,
		NewMinNotional:            chainMinNotional,
		NewInitialMarginRatio:     math.LegacyMustNewDecFromStr("0.4"),
		NewMaintenanceMarginRatio: math.LegacyMustNewDecFromStr("0.085"),
	}

	// AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	response, err := chainClient.AsyncBroadcastMsg(msg)

	if err != nil {
		panic(err)
	}

	str, _ := json.MarshalIndent(response, "", " ")
	fmt.Print(string(str))
}
