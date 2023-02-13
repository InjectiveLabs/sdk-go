package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/InjectiveLabs/sdk-go/client/common"
	exchangeclient "github.com/InjectiveLabs/sdk-go/client/exchange"
	spotExchangePB "github.com/InjectiveLabs/sdk-go/exchange/spot_exchange_rpc/pb"
	"time"
)

func main() {
	//network := common.LoadNetwork("mainnet", "k8s")
	network := common.LoadNetwork("mainnet", "k8s")
	exchangeClient, err := exchangeclient.NewExchangeClient(network.ExchangeGrpcEndpoint, common.OptionTLSCert(network.ExchangeTlsCert))
	if err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()
	marketId := "0xda0bb7a7d8361d17a9d2327ed161748f33ecbf02738b45a7dd1d812735d1531c"
	subaccountId := "0x9d9db6d545b8d6d231dfe36423086ed745c462a8000000000000000000000000"
	orderSide := "sell"

	req := spotExchangePB.StreamOrdersRequest{
		MarketId:     marketId,
		SubaccountId: subaccountId,
		OrderSide:    orderSide,
	}

	startTime := time.Now()
	go func() {
		for {
			time.Sleep(10 * time.Second)
			fmt.Println(time.Now().Sub(startTime).String())
		}
	}()

	stream, err := exchangeClient.StreamSpotOrders(ctx, req)
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
