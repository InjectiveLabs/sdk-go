package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"cosmossdk.io/math"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")
	tmClient, err := rpchttp.New(network.TmEndpoint)
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

	// initialize grpc client

	clientCtx, err := chainclient.NewClientContext(
		network.ChainId,
		senderAddress.String(),
		cosmosKeyring,
	)

	if err != nil {
		panic(err)
	}

	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint).WithClient(tmClient)

	chainClient, err := chainclient.NewChainClientV2(
		clientCtx,
		network,
	)

	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	gasPrice := chainClient.CurrentChainGasPrice(ctx)
	// adjust gas price to make it valid even if it changes between the time it is requested and the TX is broadcasted
	gasPrice = int64(float64(gasPrice) * 1.1)
	chainClient.SetGasPrice(gasPrice)

	// prepare tx msg

	msg := banktypes.MsgMultiSend{
		Inputs: []banktypes.Input{
			{
				Address: senderAddress.String(),
				Coins: []sdktypes.Coin{{
					Denom: "inj", Amount: math.NewInt(1000000000000000000)}, // 1 INJ
				},
			},
			{
				Address: senderAddress.String(),
				Coins: []sdktypes.Coin{{
					Denom: "peggy0x87aB3B4C8661e07D6372361211B96ed4Dc36B1B5", Amount: math.NewInt(1000000)}, // 1 USDT
				},
			},
		},
		Outputs: []banktypes.Output{
			{
				Address: "inj1hkhdaj2a2clmq5jq6mspsggqs32vynpk228q3r",
				Coins: []sdktypes.Coin{{
					Denom: "inj", Amount: math.NewInt(1000000000000000000)}, // 1 INJ
				},
			},
			{
				Address: "inj1hkhdaj2a2clmq5jq6mspsggqs32vynpk228q3r",
				Coins: []sdktypes.Coin{{
					Denom: "peggy0x87aB3B4C8661e07D6372361211B96ed4Dc36B1B5", Amount: math.NewInt(1000000)}, // 1 USDT
				},
			},
		},
	}

	// AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	_, response, err := chainClient.BroadcastMsg(ctx, txtypes.BroadcastMode_BROADCAST_MODE_SYNC, &msg)

	if err != nil {
		panic(err)
	}

	str, _ := json.MarshalIndent(response, "", "\t")
	fmt.Print(string(str))

	gasPrice = chainClient.CurrentChainGasPrice(ctx)
	// adjust gas price to make it valid even if it changes between the time it is requested and the TX is broadcasted
	gasPrice = int64(float64(gasPrice) * 1.1)
	chainClient.SetGasPrice(gasPrice)
}
