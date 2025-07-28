package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/InjectiveLabs/sdk-go/client/common"
	"github.com/InjectiveLabs/sdk-go/client/explorer"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")

	explorerClient, err := explorer.NewExplorerClient(network)
	if err != nil {
		log.Fatalf("Failed to create explorer client: %v", err)
	}
	defer explorerClient.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := explorerClient.FetchValidators(ctx)
	if err != nil {
		log.Panicf("Failed to fetch validators: %v", err)
	}

	fmt.Println("Full response:")
	str, _ := json.MarshalIndent(response, "", "\t")
	fmt.Print(string(str))
}
