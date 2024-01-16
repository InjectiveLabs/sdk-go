package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/InjectiveLabs/sdk-go/client/common"
	exchangeclient "github.com/InjectiveLabs/sdk-go/client/exchange"
	accountPB "github.com/InjectiveLabs/sdk-go/exchange/accounts_rpc/pb"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")
	exchangeClient, err := exchangeclient.NewExchangeClient(network)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	spotOrderHashes := []string{"0x0b156df549747187210ca5381f0291f179d76d613d0bae1a3c4fd2e3c0504b7c"}
	derivativeOrderHashes := []string{"0x82113f3998999bdc3892feaab2c4e53ba06c5fe887a2d5f9763397240f24da50"}

	req := accountPB.OrderStatesRequest{
		SpotOrderHashes:       spotOrderHashes,
		DerivativeOrderHashes: derivativeOrderHashes,
	}

	res, err := exchangeClient.GetOrderStates(ctx, &req)
	if err != nil {
		fmt.Println(err)
	}

	str, _ := json.MarshalIndent(res, "", " ")
	fmt.Print(string(str))
}
