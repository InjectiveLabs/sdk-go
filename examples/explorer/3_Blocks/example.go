package main

import (
	"context"
	"encoding/json"
	"fmt"

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
	res, err := explorerClient.GetBlocks(ctx)
	if err != nil {
		fmt.Println(err)
	}

	str, _ := json.MarshalIndent(res, "", " ")
	fmt.Print(string(str))
}
