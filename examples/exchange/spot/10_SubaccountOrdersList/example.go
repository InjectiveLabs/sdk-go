package main

import (
	"context"
	"fmt"

	"github.com/InjectiveLabs/sdk-go/client/common"
	exchangeclient "github.com/InjectiveLabs/sdk-go/client/exchange"
	spotExchangePB "github.com/InjectiveLabs/sdk-go/exchange/spot_exchange_rpc/pb"
)

func main() {
	//network := common.LoadNetwork("mainnet", "k8s")
	network := common.LoadNetwork("testnet", "k8s")
	exchangeClient, err := exchangeclient.NewExchangeClient(network.ExchangeGrpcEndpoint, common.OptionTLSCert(network.ExchangeTlsCert))
	if err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()
	marketId := "0xa508cb32923323679f29a032c70342c147c17d0145625922b0ef22e955c844c0"
	subaccountId := "0x0b46e339708ea4d87bd2fcc61614e109ac374bbe000000000000000000000000"
	skip := uint64(0)
	limit := int32(10)

	req := spotExchangePB.SubaccountOrdersListRequest{
		MarketId:     marketId,
		SubaccountId: subaccountId,
		Skip:         skip,
		Limit:        limit,
	}

	res, err := exchangeClient.GetSubaccountSpotOrdersList(ctx, req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
