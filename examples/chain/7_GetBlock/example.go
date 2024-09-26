package main

import (
	"context"
	"fmt"

	"github.com/InjectiveLabs/sdk-go/client/common"
	tmclient "github.com/InjectiveLabs/sdk-go/client/tm"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")
	tmClient := tmclient.NewRPCClient(network.TmEndpoint)
	clientCtx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	latestBlockHeight, err := tmClient.GetLatestBlockHeight(clientCtx)
	if err != nil {
		fmt.Println(err)
	}
	res, err := tmClient.GetBlock(clientCtx, int64(float64(latestBlockHeight)*0.99))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.Block.Txs)
}
