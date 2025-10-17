package main

import (
	"context"
	"encoding/json"
	"fmt"

	explorerPB "github.com/InjectiveLabs/sdk-go/exchange/explorer_rpc/pb"

	"github.com/InjectiveLabs/sdk-go/client/common"
	explorerclient "github.com/InjectiveLabs/sdk-go/client/explorer"
)

func main() {
	network := common.LoadNetwork("mainnet", "lb")
	explorerClient, err := explorerclient.NewExplorerClient(network)
	if err != nil {
		panic(err)
	}

	address := "inj1ghlynf7z25zql6kpu958wqlvmlwrhpp0a4cu9p"
	after := uint64(137677300)

	req := explorerPB.GetAccountTxsRequest{
		After:   after,
		Address: address,
	}

	ctx := context.Background()
	res, err := explorerClient.GetAccountTxs(ctx, &req)
	if err != nil {
		fmt.Println(err)
	}

	str, _ := json.MarshalIndent(res, "", "\t")
	fmt.Print(string(str))
}
