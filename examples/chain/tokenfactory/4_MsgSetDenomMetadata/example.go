package main

import (
	"encoding/json"
	"fmt"
	"os"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	tokenfactorytypes "github.com/InjectiveLabs/sdk-go/chain/tokenfactory/types"
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

	chainClient, err := chainclient.NewChainClientV2(
		clientCtx,
		network,
	)

	if err != nil {
		panic(err)
	}

	gasPrice := chainClient.CurrentChainGasPrice()
	// adjust gas price to make it valid even if it changes between the time it is requested and the TX is broadcasted
	gasPrice = int64(float64(gasPrice) * 1.1)
	chainClient.SetGasPrice(gasPrice)

	denom := "factory/inj1hkhdaj2a2clmq5jq6mspsggqs32vynpk228q3r/inj_test"
	subdenom := "inj_test"
	tokenDecimals := uint32(6)

	microDenomUnit := banktypes.DenomUnit{
		Denom:    denom,
		Exponent: 0,
		Aliases:  []string{fmt.Sprintf("micro%s", subdenom)},
	}
	denomUnit := banktypes.DenomUnit{
		Denom:    subdenom,
		Exponent: tokenDecimals,
		Aliases:  []string{subdenom},
	}

	metadata := banktypes.Metadata{
		Description: "Injective Test Token",
		DenomUnits:  []*banktypes.DenomUnit{&microDenomUnit, &denomUnit},
		Base:        denom,
		Display:     subdenom,
		Name:        "Injective Test",
		Symbol:      "INJTEST",
		URI:         "http://injective-test.com/icon.jpg",
		URIHash:     "",
		Decimals:    tokenDecimals,
	}

	message := new(tokenfactorytypes.MsgSetDenomMetadata)
	message.Sender = senderAddress.String()
	message.Metadata = metadata

	// AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	response, err := chainClient.AsyncBroadcastMsg(message)

	if err != nil {
		panic(err)
	}

	str, _ := json.MarshalIndent(response, "", "\t")
	fmt.Print(string(str))

	gasPrice = chainClient.CurrentChainGasPrice()
	// adjust gas price to make it valid even if it changes between the time it is requested and the TX is broadcasted
	gasPrice = int64(float64(gasPrice) * 1.1)
	chainClient.SetGasPrice(gasPrice)
}
