package main

import (
	"fmt"
	"github.com/InjectiveLabs/sdk-go/client"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

func main() {
	network := client.LoadNetwork("testnet", "k8s")
	tmRPC, err := rpchttp.New(network.TmEndpoint, "/websocket")
	if err != nil {
		fmt.Println(err)
	}

	senderAddress, cosmosKeyring, err := client.InitCosmosKeyring(
		"/Users/nam/.injectived",
		"injectived",
		"file",
		"inj-user",
		"12345678",
		"5d386fbdbf11f1141010f81a46b40f94887367562bd33b452bbaa6ce1cd1381e", // keyring will be used if pk not provided
		false,
	)

	clientCtx, err := client.NewClientContext(
		network.ChainId,
		senderAddress.String(),
		cosmosKeyring,
	)
	clientCtx.WithNodeURI(network.TmEndpoint)
	clientCtx = clientCtx.WithClient(tmRPC)

	msg := &banktypes.MsgSend{
		FromAddress: senderAddress.String(),
		ToAddress:   "inj1hkhdaj2a2clmq5jq6mspsggqs32vynpk228q3r",
		Amount: []sdktypes.Coin{
			{Denom: "inj", Amount: sdktypes.NewInt(1)},
		},
	}

	chainClient, err := client.NewChainClient(
		clientCtx,
		network.ChainGrpcEndpoint,
		client.OptionTLSCert(network.ChainTlsCert),
		client.OptionGasPrices("500000000inj"),
	)

	res, err := chainClient.SyncBroadcastTx(msg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

}
