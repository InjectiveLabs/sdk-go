package main

import (
	"context"
	"fmt"
	"github.com/InjectiveLabs/sdk-go/client"
)

func main() {
	network := client.LoadNetwork("mainnet", "lb")
	exchangeClient, err := client.NewExchangeClient(network.ExchangeGrpcEndpoint, client.OptionTLSCert(network.ExchangeTlsCert))
	if err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()
	marketId := "0x4ca0f92fc28be0c9761326016b5a1a2177dd6375558365116b5bdda9abc229ce"
	res, err := exchangeClient.GetOrderbook(ctx, marketId)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
