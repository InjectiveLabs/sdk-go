package main

import (
	"encoding/json"
	"fmt"
	"os"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	exchangev2types "github.com/InjectiveLabs/sdk-go/chain/exchange/types/v2"
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
	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint).WithClient(tmClient).WithSimulation(false)

	txFactory := chainclient.NewTxFactory(clientCtx)
	txFactory = txFactory.WithGas(uint64(txFactory.GasAdjustment() * 140000))

	clientInstance, err := chainclient.NewChainClientV2(
		clientCtx,
		network,
		common.OptionTxFactory(&txFactory),
	)

	if err != nil {
		panic(err)
	}

	gasPrice := clientInstance.CurrentChainGasPrice()
	// adjust gas price to make it valid even if it changes between the time it is requested and the TX is broadcasted
	gasPrice = int64(float64(gasPrice) * 1.1)
	clientInstance.SetGasPrice(gasPrice)

	defaultSubaccountID := clientInstance.DefaultSubaccount(senderAddress)

	marketId := "0x0611780ba69656949525013d947713300f56c37b6175e02f26bffa495c3208fe"

	amount := decimal.NewFromFloat(1)
	price := decimal.NewFromFloat(4.55)

	order := clientInstance.CreateSpotOrderV2(
		defaultSubaccountID,
		&chainclient.SpotOrderData{
			OrderType:    int32(exchangev2types.OrderType_BUY), //BUY SELL BUY_PO SELL_PO
			Quantity:     amount,
			Price:        price,
			FeeRecipient: senderAddress.String(),
			MarketId:     marketId,
			Cid:          uuid.NewString(),
		},
	)

	msg := new(exchangev2types.MsgCreateSpotLimitOrder)
	msg.Sender = senderAddress.String()
	msg.Order = *order

	_, response, err := clientInstance.BroadcastMsg(txtypes.BroadcastMode_BROADCAST_MODE_SYNC, msg)

	if err != nil {
		panic(err)
	}

	str, _ := json.MarshalIndent(response, "", "\t")
	fmt.Print(string(str))

	gasPrice = clientInstance.CurrentChainGasPrice()
	// adjust gas price to make it valid even if it changes between the time it is requested and the TX is broadcasted
	gasPrice = int64(float64(gasPrice) * 1.1)
	clientInstance.SetGasPrice(gasPrice)

}
