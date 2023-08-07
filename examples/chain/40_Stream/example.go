package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/InjectiveLabs/injective-core/injective-chain/stream/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	tlsConfig := tls.Config{InsecureSkipVerify: true}
	creds := credentials.NewTLS(&tlsConfig)
	cc, err := grpc.Dial("staging.stream.injective.network:443", grpc.WithTransportCredentials(creds))
	defer cc.Close()
	if err != nil {
		fmt.Println(err)
	}

	client := types.NewStreamClient(cc)

	ctx := context.Background()
	stream, err := client.Stream(ctx, &types.StreamRequest{
		BankBalancesFilter: &types.BankBalancesFilter{
			Accounts: []string{"*"},
		},
		SpotOrdersFilter: &types.OrdersFilter{
			MarketIds:     []string{"*"},
			SubaccountIds: []string{"*"},
		},
		DerivativeOrdersFilter: &types.OrdersFilter{
			MarketIds:     []string{"*"},
			SubaccountIds: []string{"*"},
		},
		SpotTradesFilter: &types.TradesFilter{
			MarketIds:     []string{"*"},
			SubaccountIds: []string{"*"},
		},
		SubaccountDepositsFilter: &types.SubaccountDepositsFilter{
			SubaccountIds: []string{"*"},
		},
		DerivativeOrderbooksFilter: &types.OrderbookFilter{
			MarketIds: []string{"*"},
		},
		SpotOrderbooksFilter: &types.OrderbookFilter{
			MarketIds: []string{"*"},
		},
		PositionsFilter: &types.PositionsFilter{
			SubaccountIds: []string{"*"},
			MarketIds:     []string{"*"},
		},
		DerivativeTradesFilter: &types.TradesFilter{
			SubaccountIds: []string{"*"},
			MarketIds:     []string{"*"},
		},
		OraclePriceFilter: &types.OraclePriceFilter{
			Symbol: []string{"*"},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			fmt.Println(err)
		}
		bz, _ := json.Marshal(res)
		fmt.Println(string(bz))
	}
}
