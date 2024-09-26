package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/InjectiveLabs/sdk-go/client/common"
	exchangeclient "github.com/InjectiveLabs/sdk-go/client/exchange"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")
	exchangeClient, err := exchangeclient.NewExchangeClient(network)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	res, err := exchangeClient.GetDerivativeMarket(ctx, "0x95698a9d8ba11660f44d7001d8c6fb191552ece5d9141a05c5d9128711cdc2e0")
	if err != nil {
		fmt.Println(err)
	}
	str, _ := json.MarshalIndent(res, "", " ")
	fmt.Print(string(str))
}
