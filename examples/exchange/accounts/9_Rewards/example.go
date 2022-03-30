package main

import (
	"context"
	"fmt"

	"github.com/InjectiveLabs/sdk-go/client/common"
	exchangeclient "github.com/InjectiveLabs/sdk-go/client/exchange"
	accountPB "github.com/InjectiveLabs/sdk-go/exchange/accounts_rpc/pb"
)

func main() {
	network := common.LoadNetwork("testnet", "k8s")
	exchangeClient, err := exchangeclient.NewExchangeClient(network.ExchangeGrpcEndpoint, common.OptionTLSCert(network.ExchangeTlsCert))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	accountAddress := "inj1rwv4zn3jptsqs7l8lpa3uvzhs57y8duemete9e"
	epoch := int64(1)

	req := accountPB.RewardsRequest{
		Epoch:          epoch,
		AccountAddress: accountAddress,
	}

	res, err := exchangeClient.GetRewards(ctx, req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
