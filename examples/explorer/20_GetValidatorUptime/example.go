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
	validatorAddress := "injvaloper1..."

	response, err := explorerClient.FetchValidatorUptime(ctx, validatorAddress)
	if err != nil {
		log.Fatalf("Failed to fetch validator uptime: %v", err)
	}

	fmt.Println("Validator uptime:")
	str, _ := json.MarshalIndent(response, "", " ")
	fmt.Print(string(str))
}
