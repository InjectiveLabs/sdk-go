package main

import (
	"context"
	"fmt"
	"github.com/InjectiveLabs/sdk-go/client"
)

func main() {
	network := client.LoadNetwork("mainnet", "lb")
	exchangeClient, err := client.NewExchangeClient(network.ExchangeGrpcEndpoint, client.OptionExchangeTLSCert(network.ExchangeTlsCert))
	if err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()
	marketIds := []string{"0x4ca0f92fc28be0c9761326016b5a1a2177dd6375558365116b5bdda9abc229ce"}
	stream, err := exchangeClient.StreamOrderbook(ctx, marketIds)
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
			fmt.Println(res)
		}
	}
}
