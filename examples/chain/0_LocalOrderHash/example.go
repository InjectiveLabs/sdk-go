package main

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"

	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	"github.com/InjectiveLabs/sdk-go/client"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/shopspring/decimal"
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

	// initialize grpc client
	clientCtx, err := chainclient.NewClientContext(
		network.ChainId,
		senderAddress.String(),
		cosmosKeyring,
	)

	if err != nil {
		fmt.Println(err)
	}

	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint).WithClient(tmClient)

	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network,
		common.OptionGasPrices(client.DefaultGasPriceWithDenom),
	)

	if err != nil {
		fmt.Println(err)
	}

	// prepare tx msg
	defaultSubaccountID := chainClient.Subaccount(senderAddress, 1)

	spotOrder := chainClient.SpotOrder(defaultSubaccountID, network, &chainclient.SpotOrderData{
		OrderType:    exchangetypes.OrderType_BUY,
		Quantity:     decimal.NewFromFloat(2),
		Price:        decimal.NewFromFloat(22.55),
		FeeRecipient: senderAddress.String(),
		MarketId:     "0x0511ddc4e6586f3bfe1acb2dd905f8b8a82c97e1edaef654b12ca7e6031ca0fa",
		Cid:          uuid.NewString(),
	})

	derivativeOrder := chainClient.DerivativeOrder(defaultSubaccountID, network, &chainclient.DerivativeOrderData{
		OrderType:    exchangetypes.OrderType_BUY,
		Quantity:     decimal.NewFromFloat(2),
		Price:        cosmtypes.MustNewDecFromStr("31000000000"),
		Leverage:     cosmtypes.MustNewDecFromStr("2.5"),
		FeeRecipient: senderAddress.String(),
		MarketId:     "0x4ca0f92fc28be0c9761326016b5a1a2177dd6375558365116b5bdda9abc229ce",
		Cid:          uuid.NewString(),
	})

	msg := new(exchangetypes.MsgBatchCreateSpotLimitOrders)
	msg.Sender = senderAddress.String()
	msg.Orders = []exchangetypes.SpotOrder{*spotOrder}

	msg1 := new(exchangetypes.MsgBatchCreateDerivativeLimitOrders)
	msg1.Sender = senderAddress.String()
	msg1.Orders = []exchangetypes.DerivativeOrder{*derivativeOrder, *derivativeOrder}

	// compute local order hashes
	orderHashes, err := chainClient.ComputeOrderHashes(msg.Orders, msg1.Orders, defaultSubaccountID)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("computed spot order hashes: ", orderHashes.Spot)
	fmt.Println("computed derivative order hashes: ", orderHashes.Derivative)

	//AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
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
}
