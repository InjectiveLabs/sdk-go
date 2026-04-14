package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/InjectiveLabs/sdk-go/client"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
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
		os.Getenv("INJECTIVE_PRIVATE_KEY"), // keyring will be used if pk not provided
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

	chainClient, err := chainclient.NewChainClientV2(
		clientCtx,
		network,
		common.OptionGasPrices(client.DefaultGasPriceWithDenom),
	)

	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	marketId := "0x0611780ba69656949525013d947713300f56c37b6175e02f26bffa495c3208fe"
	subaccountId := chainClient.Subaccount(senderAddress, 0)
	orderHashes := []string{"0x57a01cd26f1e2080860af3264e865d7c9c034a701e30946d01c1dc7a303cf2c1"}

	res, err := chainClient.FetchChainSpotOrdersByHashes(ctx, marketId, subaccountId.Hex(), orderHashes)
	if err != nil {
		fmt.Println(err)
	}

	str, _ := json.MarshalIndent(res, "", "\t")
	fmt.Print(string(str))

}
