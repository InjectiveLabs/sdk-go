package main

import (
	"context"
	"fmt"
	"os"
	"time"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
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
		return
	}

	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint).WithClient(tmClient)

	chainClient, err := chainclient.NewChainClient(
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

	ctx := context.Background()
	marketsAssistant, err := chainclient.NewMarketsAssistant(ctx, chainClient)
	if err != nil {
		panic(err)
	}

	defaultSubaccountID := chainClient.DefaultSubaccount(senderAddress)

	marketId := "0x4ca0f92fc28be0c9761326016b5a1a2177dd6375558365116b5bdda9abc229ce"
	amount := decimal.NewFromFloat(0.01)
	price := decimal.RequireFromString("33000") //33,000
	leverage := decimal.RequireFromString("2.5")

	order := chainClient.CreateDerivativeOrder(
		defaultSubaccountID,
		&chainclient.DerivativeOrderData{
			OrderType:    exchangetypes.OrderType_SELL, //BUY SELL
			Quantity:     amount,
			Price:        price,
			Leverage:     leverage,
			FeeRecipient: senderAddress.String(),
			MarketId:     marketId,
			IsReduceOnly: true,
			Cid:          uuid.NewString(),
		},
		marketsAssistant,
	)

	msg := new(exchangetypes.MsgCreateDerivativeMarketOrder)
	msg.Sender = senderAddress.String()
	msg.Order = exchangetypes.DerivativeOrder(*order)

	simRes, err := chainClient.SimulateMsg(clientCtx, msg)

	if err != nil {
		panic(err)
	}

	msgCreateDerivativeMarketOrderResponse := exchangetypes.MsgCreateDerivativeMarketOrderResponse{}
	err = msgCreateDerivativeMarketOrderResponse.Unmarshal(simRes.Result.MsgResponses[0].Value)

	if err != nil {
		panic(err)
	}

	fmt.Println("simulated order hash", msgCreateDerivativeMarketOrderResponse.OrderHash)

	// AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	err = chainClient.QueueBroadcastMsg(msg)

	if err != nil {
		fmt.Println(err)
		return
	}

	time.Sleep(time.Second * 5)

	gasFee, err := chainClient.GetGasFee()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("gas fee:", gasFee, "INJ")

	gasPrice = chainClient.CurrentChainGasPrice()
	// adjust gas price to make it valid even if it changes between the time it is requested and the TX is broadcasted
	gasPrice = int64(float64(gasPrice) * 1.1)
	chainClient.SetGasPrice(gasPrice)
}
