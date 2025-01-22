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

	// Example validator address (replace with an actual validator address)
	validatorAddress := "injvaloper1kk523rsm9pey740cx4plalp40009ncs0wrchfe"

	response, err := explorerClient.FetchValidator(ctx, validatorAddress)
	if err != nil {
		log.Fatalf("Failed to fetch validator: %v", err)
	}

	fmt.Println("Validator:")
	str, _ := json.MarshalIndent(response, "", " ")
	fmt.Print(string(str))
}
