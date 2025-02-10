package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/InjectiveLabs/sdk-go/client/common"
	"github.com/InjectiveLabs/sdk-go/client/exchange"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")
	exchangeClient, err := exchange.NewExchangeClient(network)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Fetch INJ burnt details
	injBurntResponse, err := exchangeClient.FetchInjBurnt(ctx)
	if err != nil {
		fmt.Printf("Failed to fetch INJ burnt details: %v\n", err)
		return
	}

	// Print JSON representation of the response
	jsonResponse, err := json.MarshalIndent(injBurntResponse, "", "  ")
	if err != nil {
		fmt.Printf("Failed to marshal response to JSON: %v\n", err)
		return
	}

	fmt.Println("INJ Burnt Details:")
	fmt.Println(string(jsonResponse))
}
