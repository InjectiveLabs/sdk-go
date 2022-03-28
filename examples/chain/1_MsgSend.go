package main

import (
	"fmt"
	"os"
	"github.com/InjectiveLabs/sdk-go/client/common"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	"time"
)

func main() {
	network := common.LoadNetwork("testnet", "k8s")
	tmRPC, err := rpchttp.New(network.TmEndpoint, "/websocket")
	if err != nil {
		fmt.Println(err)
	}

	senderAddress, cosmosKeyring, err := chainclient.InitCosmosKeyring(
		os.Getenv("HOME") + "/.injectived",
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
		fmt.Println(err)
	}

	clientCtx.WithNodeURI(network.TmEndpoint)
	clientCtx = clientCtx.WithClient(tmRPC)

	msg := &banktypes.MsgSend{
		FromAddress: senderAddress.String(),
		ToAddress:   "inj1hkhdaj2a2clmq5jq6mspsggqs32vynpk228q3r",
		Amount: []sdktypes.Coin{{
			Denom: "inj", Amount: sdktypes.NewInt(1000000000000000000)}, // 1 INJ
		},
	}

	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network.ChainGrpcEndpoint,
		common.OptionTLSCert(network.ChainTlsCert),
		common.OptionGasPrices("500000000inj"),
	)

	if err != nil {
		fmt.Println(err)
	}

	for i:=0; i<1; i++ {
		err := chainClient.QueueBroadcastMsg(msg)
		if err != nil {
			fmt.Println(err)
		}
	}
	time.Sleep(time.Second * 5)
}
