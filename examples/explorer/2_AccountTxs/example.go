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

	address := "inj1akxycslq8cjt0uffw4rjmfm3echchptu52a2dq"
	after := uint64(14112176)

	req := explorerPB.GetAccountTxsRequest{
		After:   after,
		Address: address,
	}

	ctx := context.Background()
	res, err := explorerClient.GetAccountTxs(ctx, req)
	if err != nil {
		fmt.Println(err)
	}

	str, _ := json.MarshalIndent(res, "", " ")
	fmt.Print(string(str))
}
