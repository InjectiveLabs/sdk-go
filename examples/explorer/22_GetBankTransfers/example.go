package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/InjectiveLabs/sdk-go/client/common"
	"github.com/InjectiveLabs/sdk-go/client/explorer"
	explorerPB "github.com/InjectiveLabs/sdk-go/exchange/explorer_rpc/pb"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")

	explorerClient, err := explorer.NewExplorerClient(network)
	if err != nil {
		log.Fatalf("Failed to create explorer client: %v", err)
	}
	defer explorerClient.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	req := &explorerPB.GetBankTransfersRequest{
		Senders: []string{"inj17xpfvakm2amg962yls6f84z3kell8c5l6s5ye9"},
		Limit:   5, // Limit number of transfers
	}

	response, err := explorerClient.FetchBankTransfers(ctx, req)
	if err != nil {
		log.Panicf("Failed to fetch bank transfers: %v", err)
	}

	fmt.Println("Bank transfers:")
	str, _ := json.MarshalIndent(response, "", " ")
	fmt.Print(string(str))
}
