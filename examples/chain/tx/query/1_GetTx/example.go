package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/InjectiveLabs/sdk-go/client"
	"github.com/InjectiveLabs/sdk-go/client/common"
	"github.com/joho/godotenv"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
)

func main() {
	_ = godotenv.Load()
	// network := common.LoadNetwork("mainnet", "k8s")
	network := common.LoadNetwork("mainnet", "lb")
	tmRPC, err := rpchttp.New(network.TmEndpoint)

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

	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint).WithClient(tmRPC)

	chainClient, err := chainclient.NewChainClientV2(
		clientCtx,
		network,
		common.OptionGasPrices(client.DefaultGasPriceWithDenom),
	)

	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	timeOutCtx, cancelFn := context.WithTimeout(ctx, 30*time.Second)
	defer cancelFn()

	resp, err := chainClient.GetTx(timeOutCtx, "A2B2B971C690AE7977451D24D6F450AECE6BCCB271E91E32C2563342DDA5254B")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.TxResponse)
}
