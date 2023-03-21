package main

import (
	"context"
	"encoding/json"
	"fmt"

	exchangeclient "github.com/InjectiveLabs/sdk-go/client/exchange"
)

func main() {
	// select network: local, testnet, mainnet
	//network := common.LoadNetwork("testnet", "k8s")
	exchangeClient, err := exchangeclient.NewExchangeClient("localhost:9907")
	if err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()

	stream, err := exchangeClient.StreamAccountPortfolio(ctx, "inj1rgxjfea3y2e7n0frz5syly8n5zulagy3fc56jy", "", "")
	if err != nil {
		fmt.Println(err)
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
