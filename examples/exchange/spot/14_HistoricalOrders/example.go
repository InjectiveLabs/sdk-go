package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/InjectiveLabs/sdk-go/client/common"
	exchangeclient "github.com/InjectiveLabs/sdk-go/client/exchange"
	spotExchangePB "github.com/InjectiveLabs/sdk-go/exchange/spot_exchange_rpc/pb"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")
	exchangeClient, err := exchangeclient.NewExchangeClient(network)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	marketId := "0x0611780ba69656949525013d947713300f56c37b6175e02f26bffa495c3208fe"
	subaccountId := "0xaf79152ac5df276d9a8e1e2e22822f9713474902000000000000000000000000"
	skip := uint64(0)
	limit := int32(10)
	orderTypes := []string{"buy_po"}

	req := spotExchangePB.OrdersHistoryRequest{
		SubaccountId: subaccountId,
		MarketId:     marketId,
		Skip:         skip,
		Limit:        limit,
		OrderTypes:   orderTypes,
	}

	res, err := exchangeClient.GetHistoricalSpotOrders(ctx, req)
	if err != nil {
		panic(err)
	}

	str, _ := json.MarshalIndent(res, "", " ")
	fmt.Print(string(str))
}
