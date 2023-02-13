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
	network := common.LoadNetwork("testnet", "k8s")
	explorerClient, err := explorerclient.NewExplorerClient(network.ExplorerGrpcEndpoint, common.OptionTLSCert(network.ExplorerTlsCert))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	before := uint64(7158400)

	req := explorerPB.GetTxsRequest{
		Before: before,
	}

	res, err := explorerClient.GetTxs(ctx, req)
	if err != nil {
		fmt.Println(err)
	}

	str, _ := json.MarshalIndent(res, "", " ")
	fmt.Print(string(str))
}
