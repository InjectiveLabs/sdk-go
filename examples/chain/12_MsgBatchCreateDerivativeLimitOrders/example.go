package main

import (
	"fmt"
	"github.com/InjectiveLabs/sdk-go/client"
	"os"
	"time"

	"github.com/InjectiveLabs/sdk-go/client/common"
	"github.com/shopspring/decimal"

	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
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
		common.OptionGasPrices(client.DefaultGasPriceWithDenom),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	defaultSubaccountID := chainClient.DefaultSubaccount(senderAddress)

	marketId := "0x4ca0f92fc28be0c9761326016b5a1a2177dd6375558365116b5bdda9abc229ce"
	amount := decimal.NewFromFloat(2)
	price := cosmtypes.MustNewDecFromStr("31000000000") //31,000
	leverage := cosmtypes.MustNewDecFromStr("2.5")

	order := chainClient.DerivativeOrder(defaultSubaccountID, network, &chainclient.DerivativeOrderData{
		OrderType:    exchangetypes.OrderType_BUY, //BUY SELL BUY_PO SELL_PO
		Quantity:     amount,
		Price:        price,
		Leverage:     leverage,
		FeeRecipient: senderAddress.String(),
		MarketId:     marketId,
		IsReduceOnly: true,
	})

	msg := new(exchangetypes.MsgBatchCreateDerivativeLimitOrders)
	msg.Sender = senderAddress.String()
	msg.Orders = []exchangetypes.DerivativeOrder{*order}

	simRes, err := chainClient.SimulateMsg(clientCtx, msg)

	if err != nil {
		fmt.Println(err)
		return
	}

	msgBatchCreateDerivativeLimitOrdersResponse := exchangetypes.MsgBatchCreateDerivativeLimitOrdersResponse{}
	err = msgBatchCreateDerivativeLimitOrdersResponse.Unmarshal(simRes.Result.MsgResponses[0].Value)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("simulated order hashes", msgBatchCreateDerivativeLimitOrdersResponse.OrderHashes)

	//AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
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
}
