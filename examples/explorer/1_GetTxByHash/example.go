package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/InjectiveLabs/sdk-go/client/common"
	explorerclient "github.com/InjectiveLabs/sdk-go/client/explorer"
)

func main() {
	network := common.LoadNetwork("mainnet", "sentry")
	explorerClient, err := explorerclient.NewExplorerClient(network)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	hash := "E5DCF04CC670A0567F58683409F7DAFC49754278DAAD507FE6EB40DFBFD71830"
	res, err := explorerClient.GetTxByTxHash(ctx, hash)
	if err != nil {
		fmt.Println(err)
	}

	str, _ := json.MarshalIndent(res, "", "\t")
	fmt.Print(string(str))
}
