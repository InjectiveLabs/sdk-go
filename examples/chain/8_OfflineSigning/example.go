// example for how to sign tx offline, store to file and load + broadcast later
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/shopspring/decimal"

	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	"github.com/InjectiveLabs/sdk-go/client"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
)

func StoreTxToFile(fileName string, txBytes []byte) error {
	return os.WriteFile(fileName, txBytes, 0755)
}

func LoadTxFromFile(fileName string) ([]byte, error) {
	return os.ReadFile(fileName)
}

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

	ctx := context.Background()
	marketsAssistant, err := chainclient.NewMarketsAssistant(ctx, chainClient)
	if err != nil {
		panic(err)
	}

	defaultSubaccountID := chainClient.DefaultSubaccount(senderAddress)
	marketId := "0x0611780ba69656949525013d947713300f56c37b6175e02f26bffa495c3208fe"
	amount := decimal.NewFromFloat(2)
	price := decimal.NewFromFloat(1.02)

	order := chainClient.CreateSpotOrder(
		defaultSubaccountID,
		&chainclient.SpotOrderData{
			OrderType:    exchangetypes.OrderType_BUY, //BUY SELL BUY_PO SELL_PO
			Quantity:     amount,
			Price:        price,
			FeeRecipient: senderAddress.String(),
			MarketId:     marketId,
		},
		marketsAssistant,
	)

	msg := new(exchangetypes.MsgCreateSpotLimitOrder)
	msg.Sender = senderAddress.String()
	msg.Order = exchangetypes.SpotOrder(*order)

	accNum, accSeq := chainClient.GetAccNonce()
	signedTx, err := chainClient.BuildSignedTx(clientCtx, accNum, accSeq, 20000, client.DefaultGasPrice, msg)
	if err != nil {
		panic(err)
	}

	// store signed tx into file
	err = StoreTxToFile("msg.dat", signedTx)
	if err != nil {
		panic(err)
	}

	// Broadcast the signed tx using BroadcastSignedTx, AsyncBroadcastSignedTx, or SyncBroadcastSignedTx
	response, err := chainClient.BroadcastSignedTx(signedTx, txtypes.BroadcastMode_BROADCAST_MODE_SYNC)
	if err != nil {
		panic(err)
	}

	fmt.Printf("tx hash: %s\n", response.TxResponse.TxHash)
	fmt.Printf("tx code: %v\n\n", response.TxResponse.Code)

	str, _ := json.MarshalIndent(response, "", " ")
	fmt.Print(string(str))
}
