package main

import (
	"context"
	"fmt"
	"github.com/InjectiveLabs/sdk-go/client/common"
	exchangeclient "github.com/InjectiveLabs/sdk-go/client/exchange"
)

func main() {
	network := common.LoadNetwork("mainnet", "lb")
	exchangeClient, err := exchangeclient.NewExchangeClient(network.ExchangeGrpcEndpoint, common.OptionTLSCert(network.ExchangeTlsCert))
	if err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()
	res, err := exchangeClient.GetOracleList(ctx)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
