package main

import (
	"context"
	"fmt"
	"os"
	"time"

	exchangeclient "github.com/InjectiveLabs/sdk-go/client/exchange"
	"github.com/google/uuid"

	"github.com/InjectiveLabs/sdk-go/client"
	"github.com/InjectiveLabs/sdk-go/client/common"
	"github.com/shopspring/decimal"

	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
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

	marketId := "0x0511ddc4e6586f3bfe1acb2dd905f8b8a82c97e1edaef654b12ca7e6031ca0fa"
	amount := decimal.NewFromFloat(0.1)
	price := decimal.NewFromFloat(22)

	order := chainClient.CreateSpotOrder(
		defaultSubaccountID,
		network,
		&chainclient.SpotOrderData{
			OrderType:    exchangetypes.OrderType_SELL, //BUY SELL
			Quantity:     amount,
			Price:        price,
			FeeRecipient: senderAddress.String(),
			MarketId:     marketId,
			Cid:          uuid.NewString(),
		},
		marketsAssistant,
	)

	msg := new(exchangetypes.MsgCreateSpotMarketOrder)
	msg.Sender = senderAddress.String()
	msg.Order = exchangetypes.SpotOrder(*order)

	simRes, err := chainClient.SimulateMsg(clientCtx, msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	msgCreateSpotMarketOrderResponse := exchangetypes.MsgCreateSpotMarketOrderResponse{}
	err = msgCreateSpotMarketOrderResponse.Unmarshal(simRes.Result.MsgResponses[0].Value)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("simulated order hash", msgCreateSpotMarketOrderResponse.OrderHash)

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
