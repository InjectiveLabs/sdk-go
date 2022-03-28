package main

import (
	"context"
	"fmt"
	"github.com/InjectiveLabs/sdk-go/client/common"
	exchangeclient "github.com/InjectiveLabs/sdk-go/client/exchange"
)

func main() {
	network := common.LoadNetwork("testnet", "k8s")
	exchangeClient, err := exchangeclient.NewExchangeClient(network.ExchangeGrpcEndpoint, common.OptionTLSCert(network.ExchangeTlsCert))
	if err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()
	accountAddress := "inj14au322k9munkmx5wrchz9q30juf5wjgz2cfqku"
	res, err := exchangeClient.GetSubaccountsList(ctx, accountAddress)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
