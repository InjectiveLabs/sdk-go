package main

import (
	"context"
	"fmt"
	"os"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	"github.com/InjectiveLabs/sdk-go/client"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
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
	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint).WithClient(tmClient).WithSimulation(false)

	txFactory := chainclient.NewTxFactory(clientCtx)
	txFactory = txFactory.WithGasPrices(client.DefaultGasPriceWithDenom)
	txFactory = txFactory.WithGas(uint64(txFactory.GasAdjustment() * 140000))

	clientInstance, err := chainclient.NewChainClient(
		clientCtx,
		network,
		common.OptionTxFactory(&txFactory),
	)

	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	marketsAssistant, err := chainclient.NewMarketsAssistant(ctx, clientInstance)
	if err != nil {
		panic(err)
	}

	defaultSubaccountID := clientInstance.DefaultSubaccount(senderAddress)

	marketId := "0x0611780ba69656949525013d947713300f56c37b6175e02f26bffa495c3208fe"

	amount := decimal.NewFromFloat(1)
	price := decimal.NewFromFloat(4.55)

	order := clientInstance.CreateSpotOrder(
		defaultSubaccountID,
		&chainclient.SpotOrderData{
			OrderType:    exchangetypes.OrderType_BUY, //BUY SELL BUY_PO SELL_PO
			Quantity:     amount,
			Price:        price,
			FeeRecipient: senderAddress.String(),
			MarketId:     marketId,
			Cid:          uuid.NewString(),
		},
		marketsAssistant,
	)

	msg := new(exchangetypes.MsgCreateSpotLimitOrder)
	msg.Sender = senderAddress.String()
	msg.Order = exchangetypes.SpotOrder(*order)

	result, err := clientInstance.SyncBroadcastMsg(msg)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Broadcast result: %s\n", result)

}
