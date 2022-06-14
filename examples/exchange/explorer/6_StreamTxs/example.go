package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/InjectiveLabs/sdk-go/client/common"
	exchangeclient "github.com/InjectiveLabs/sdk-go/client/exchange"
)

func main() {
	//network := common.LoadNetwork("mainnet", "k8s")
	network := common.LoadNetwork("testnet", "k8s")
	exchangeClient, err := exchangeclient.NewExchangeClient(network.ExchangeGrpcEndpoint, common.OptionTLSCert(network.ExchangeTlsCert))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	stream, err := exchangeClient.StreamTxs(ctx)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case <-ctx.Done():
			return
		default:
			res, err := stream.Recv()
			if err != nil {
				panic(err)
				return
			}
			str, _ := json.MarshalIndent(res, "", " ")
			fmt.Print(string(str))
		}
	}
}
