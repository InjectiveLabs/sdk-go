package main

import (
	"fmt"
	"os"
	"time"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	exchangev2types "github.com/InjectiveLabs/sdk-go/chain/exchange/types/v2"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")
	remoteAddress := fmt.Sprintf("%s/websocket", network.TmEndpoint)
	tmClient, err := rpchttp.New(remoteAddress)
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

	// prepare tx msg
	defaultSubaccountID := chainClient.Subaccount(senderAddress, 1)

	spotOrder := chainClient.CreateSpotOrderV2(
		defaultSubaccountID,
		&chainclient.SpotOrderData{
			OrderType:    exchangev2types.OrderType_BUY,
			Quantity:     decimal.NewFromFloat(2),
			Price:        decimal.NewFromFloat(22.55),
			FeeRecipient: senderAddress.String(),
			MarketId:     "0x0611780ba69656949525013d947713300f56c37b6175e02f26bffa495c3208fe",
			Cid:          uuid.NewString(),
		},
	)

	derivativeOrder := chainClient.CreateDerivativeOrderV2(
		defaultSubaccountID,
		&chainclient.DerivativeOrderData{
			OrderType:    exchangev2types.OrderType_BUY,
			Quantity:     decimal.NewFromFloat(2),
			Price:        decimal.RequireFromString("31"),
			Leverage:     decimal.RequireFromString("2.5"),
			FeeRecipient: senderAddress.String(),
			MarketId:     "0x17ef48032cb24375ba7c2e39f384e56433bcab20cbee9a7357e4cba2eb00abe6",
			Cid:          uuid.NewString(),
		},
	)

	msg := new(exchangev2types.MsgBatchCreateSpotLimitOrders)
	msg.Sender = senderAddress.String()
	msg.Orders = []exchangev2types.SpotOrder{*spotOrder}

	msg1 := new(exchangev2types.MsgBatchCreateDerivativeLimitOrders)
	msg1.Sender = senderAddress.String()
	msg1.Orders = []exchangev2types.DerivativeOrder{*derivativeOrder}

	// compute local order hashes
	orderHashes, err := chainClient.ComputeOrderHashes(msg.Orders, msg1.Orders, defaultSubaccountID)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("computed spot order hashes: ", orderHashes.Spot)
	fmt.Println("computed derivative order hashes: ", orderHashes.Derivative)

	// AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	err = chainClient.QueueBroadcastMsg(msg, msg1)

	if err != nil {
		fmt.Println(err)
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
