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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Example contract address (replace with an actual contract address)
	contractAddress := "inj1ady3s7whq30l4fx8sj3x6muv5mx4dfdlcpv8n7"

	req := &explorerPB.GetContractTxsRequest{
		Address: contractAddress,
		Limit:   10, // Fetch 10 transactions
	}

	response, err := explorerClient.FetchContractTxs(ctx, req)
	if err != nil {
		log.Fatalf("Failed to fetch contract transactions: %v", err)
	}

	fmt.Println("Total Contract Transactions:", len(response.Data))
	for _, tx := range response.Data {
		fmt.Printf("Tx Hash: %s, Block: %d\n", tx.Hash, tx.BlockNumber)
	}

	fmt.Printf("\n\n")
	fmt.Println("Full response:")
	str, _ := json.MarshalIndent(response, "", " ")
	fmt.Print(string(str))
}
