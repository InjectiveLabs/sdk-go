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
	denom := "inj"
	subaccountId := "0xaf79152ac5df276d9a8e1e2e22822f9713474902000000000000000000000000"
	transferTypes := []string{"deposit"}

	req := accountPB.SubaccountHistoryRequest{
		Denom:         denom,
		SubaccountId:  subaccountId,
		TransferTypes: transferTypes,
	}

	res, err := exchangeClient.GetSubaccountHistory(ctx, req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
