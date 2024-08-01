// example for how to sign tx offline, store to file and load + broadcast later
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	exchangeclient "github.com/InjectiveLabs/sdk-go/client/exchange"

	"github.com/InjectiveLabs/sdk-go/client"
	"github.com/InjectiveLabs/sdk-go/client/common"
	"github.com/shopspring/decimal"

	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
)

func StoreTxToFile(fileName string, txBytes []byte) error {
	return ioutil.WriteFile(fileName, txBytes, 0755)
}

func LoadTxFromFile(fileName string) ([]byte, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}

func main() {
	network := common.LoadNetwork("devnet", "")
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

	exchangeClient, err := exchangeclient.NewExchangeClient(network)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	marketsAssistant, err := chainclient.NewMarketsAssistantInitializedFromChain(ctx, exchangeClient)
	if err != nil {
		panic(err)
	}

	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network,
		common.OptionGasPrices(client.DefaultGasPriceWithDenom),
	)

	if err != nil {
		panic(err)
	}

	defaultSubaccountID := chainClient.DefaultSubaccount(senderAddress)
	marketId := "0xa508cb32923323679f29a032c70342c147c17d0145625922b0ef22e955c844c0"
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
	signedTx, err := chainClient.BuildSignedTx(clientCtx, accNum, accSeq, 20000, msg)
	if err != nil {
		panic(err)
	}

	// store signed tx into file
	err = StoreTxToFile("msg.dat", signedTx)
	if err != nil {
		panic(err)
	}

	// load from file and broadcast the signed tx
	signedTxBytesFromFile, err := LoadTxFromFile("msg.dat")
	if err != nil {
		panic(err)
	}

	txResult, err := chainClient.SyncBroadcastSignedTx(signedTxBytesFromFile)
	if err != nil {
		panic(err)
	}

	fmt.Println("txResult:", txResult)
}
