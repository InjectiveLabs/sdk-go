package main

import (
	"encoding/json"
	"fmt"
	"os"

	wasmxtypes "github.com/InjectiveLabs/sdk-go/chain/wasmx/types"
	"github.com/InjectiveLabs/sdk-go/client"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
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
		"f9db9bf330e23cb7839039e944adef6e9df447b90b503d5b4464c90bea9022f3", // keyring will be used if pk not provided
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
		fmt.Println(err)
		return
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

	firstAmount := 69
	firstToken := "factory/inj1hdvy6tl89llqy3ze8lv6mz5qh66sx9enn0jxg6/inj12ngevx045zpvacus9s6anr258gkwpmthnz80e9"
	secondAmount := 420
	secondToken := "peggy0x44C21afAaF20c270EBbF5914Cfc3b5022173FEB7"
	thirdAmount := 1
	thirdToken := "peggy0x87aB3B4C8661e07D6372361211B96ed4Dc36B1B5"
	funds := fmt.Sprintf(
		"%v%s,%v%s,%v%s",
		firstAmount,
		firstToken,
		secondAmount,
		secondToken,
		thirdAmount,
		thirdToken,
	)

	message := wasmxtypes.MsgExecuteContractCompat{
		Sender:   senderAddress.String(),
		Contract: "inj1ady3s7whq30l4fx8sj3x6muv5mx4dfdlcpv8n7",
		Msg:      "{\"increment\": {}}",
		Funds:    funds,
	}

	//AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	response, err := chainClient.AsyncBroadcastMsg(&message)

	if err != nil {
		panic(err)
	}

	str, _ := json.MarshalIndent(response, "", " ")
	fmt.Println(string(str))
}
