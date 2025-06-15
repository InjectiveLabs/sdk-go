package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"google.golang.org/grpc/metadata"

	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

	"github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	"github.com/InjectiveLabs/sdk-go/client"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
	tmclient "github.com/InjectiveLabs/sdk-go/client/tm"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")
	tmClient, err := rpchttp.New(network.TmEndpoint, "/websocket")
	if err != nil {
		panic(err)
	}

	senderAddress, cosmosKeyring, err := chainclient.InitCosmosKeyring(
		os.Getenv("HOME")+"/.injectived",
		"injectived",
		"file",
		"inj-user",
		"12345678",
		"5d386fbdbf11f1141010f81a46b40f94887367562bd33b452bbaa6ce1cd1381e", // keyring will be used if pk not provided
		false,
	)

	if err != nil {
		panic(err)
	}

	clientCtx, err := chainclient.NewClientContext(
		network.ChainId,
		senderAddress.String(),
		cosmosKeyring,
	)

	if err != nil {
		panic(err)
	}

	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint).WithClient(tmClient)

	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network,
		common.OptionGasPrices(client.DefaultGasPriceWithDenom),
	)

	if err != nil {
		panic(err)
	}

	status := "Active"
	marketIds := []string{"0x0611780ba69656949525013d947713300f56c37b6175e02f26bffa495c3208fe"}

	rpcClient := tmclient.NewRPCClient(network.TmEndpoint)
	latestBlockHeight, err := rpcClient.GetLatestBlockHeight(context.Background())
	if err != nil {
		panic(err)
	}

	{
		res, err := fetchSpotMarketsAtHeight(
			chainClient,
			latestBlockHeight,
			status,
			marketIds,
		)
		if err != nil {
			panic(err)
		}
		str, _ := json.MarshalIndent(res, "", " ")
		fmt.Println(string(str))
	}

	{
		res, err := fetchSpotMarketsAtHeight(
			chainClient,
			latestBlockHeight-1,
			status,
			marketIds,
		)
		if err != nil {
			panic(err)
		}
		str, _ := json.MarshalIndent(res, "", " ")
		fmt.Println(string(str))
	}

	{
		_, err := fetchSpotMarketsAtHeight(
			chainClient,
			10,
			status,
			marketIds,
		)
		if err == nil {
			panic("Expected error for old block height")
		}
		fmt.Println("Expected error for old block height:", err)
	}
}

func fetchSpotMarketsAtHeight(
	chainClient chainclient.ChainClient,
	height int64,
	status string,
	marketIds []string,
) (*types.QuerySpotMarketsResponse, error) {
	ctx := context.Background()

	ctxWithHeight := metadata.AppendToOutgoingContext(
		ctx,
		grpctypes.GRPCBlockHeightHeader, strconv.FormatInt(height, 10),
	)

	return chainClient.FetchChainSpotMarkets(ctxWithHeight, status, marketIds)
}
