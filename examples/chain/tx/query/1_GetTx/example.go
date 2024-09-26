package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/InjectiveLabs/sdk-go/client"
	"github.com/InjectiveLabs/sdk-go/client/common"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")
	tmRPC, err := rpchttp.New(network.TmEndpoint, "/websocket")

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

	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint).WithClient(tmRPC)

	chainClient, err := chainclient.NewChainClient(
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

	resp, err := chainClient.GetTx(timeOutCtx, "4446C993253F02F1ECF563477C48611859A36615C05262C2C74826E795BBF586")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.TxResponse)
}
