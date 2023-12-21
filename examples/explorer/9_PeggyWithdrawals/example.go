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
	network := common.LoadNetwork("testnet", "lb")
	explorerClient, err := explorerclient.NewExplorerClient(network)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	sender := "inj14au322k9munkmx5wrchz9q30juf5wjgz2cfqku"

	req := explorerPB.GetPeggyWithdrawalTxsRequest{
		Sender: sender,
	}

	res, err := explorerClient.GetPeggyWithdrawals(ctx, &req)
	if err != nil {
		fmt.Println(err)
	}

	str, _ := json.MarshalIndent(res, "", " ")
	fmt.Print(string(str))
}
