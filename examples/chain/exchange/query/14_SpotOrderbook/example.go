package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"cosmossdk.io/math"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

	v2 "github.com/InjectiveLabs/sdk-go/chain/exchange/types/v2"
	"github.com/InjectiveLabs/sdk-go/client"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
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

	ctx := context.Background()

	marketId := "0x0611780ba69656949525013d947713300f56c37b6175e02f26bffa495c3208fe"
	limit := uint64(2)
	orderSide := v2.OrderSide_Buy
	limitCumulativeNotional := math.LegacyDec{}
	limitCumulativeQuantity := math.LegacyDec{}

	res, err := chainClient.FetchChainSpotOrderbookV2(ctx, marketId, limit, orderSide, limitCumulativeNotional, limitCumulativeQuantity)
	if err != nil {
		fmt.Println(err)
	}

	str, _ := json.MarshalIndent(res, "", " ")
	fmt.Print(string(str))

}
