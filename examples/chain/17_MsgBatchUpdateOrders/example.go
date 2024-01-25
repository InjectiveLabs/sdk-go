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
		fmt.Println(err)
		return
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
		fmt.Println(err)
		return
	}

	defaultSubaccountID := chainClient.DefaultSubaccount(senderAddress)

	smarketId := "0x0511ddc4e6586f3bfe1acb2dd905f8b8a82c97e1edaef654b12ca7e6031ca0fa"
	samount := decimal.NewFromFloat(2)
	sprice := decimal.NewFromFloat(22.5)
	smarketIds := []string{"0xa508cb32923323679f29a032c70342c147c17d0145625922b0ef22e955c844c0"}

	spot_order := chainClient.CreateSpotOrder(
		defaultSubaccountID,
		&chainclient.SpotOrderData{
			OrderType:    exchangetypes.OrderType_BUY, //BUY SELL BUY_PO SELL_PO
			Quantity:     samount,
			Price:        sprice,
			FeeRecipient: senderAddress.String(),
			MarketId:     smarketId,
			Cid:          uuid.NewString(),
		},
		marketsAssistant,
	)

	dmarketId := "0x4ca0f92fc28be0c9761326016b5a1a2177dd6375558365116b5bdda9abc229ce"
	damount := decimal.NewFromFloat(0.01)
	dprice := decimal.RequireFromString("31000") //31,000
	dleverage := decimal.RequireFromString("2")
	dmarketIds := []string{"0x4ca0f92fc28be0c9761326016b5a1a2177dd6375558365116b5bdda9abc229ce"}

	derivative_order := chainClient.CreateDerivativeOrder(
		defaultSubaccountID,
		&chainclient.DerivativeOrderData{
			OrderType:    exchangetypes.OrderType_BUY, //BUY SELL BUY_PO SELL_PO
			Quantity:     damount,
			Price:        dprice,
			Leverage:     dleverage,
			FeeRecipient: senderAddress.String(),
			MarketId:     dmarketId,
			IsReduceOnly: false,
			Cid:          uuid.NewString(),
		},
		marketsAssistant,
	)

	msg := new(exchangetypes.MsgBatchUpdateOrders)
	msg.Sender = senderAddress.String()
	msg.SubaccountId = defaultSubaccountID.Hex()
	msg.SpotOrdersToCreate = []*exchangetypes.SpotOrder{spot_order}
	msg.DerivativeOrdersToCreate = []*exchangetypes.DerivativeOrder{derivative_order}
	msg.SpotMarketIdsToCancelAll = smarketIds
	msg.DerivativeMarketIdsToCancelAll = dmarketIds

	simRes, err := chainClient.SimulateMsg(clientCtx, msg)

	if err != nil {
		fmt.Println(err)
		return
	}

	MsgBatchUpdateOrdersResponse := exchangetypes.MsgBatchUpdateOrdersResponse{}
	MsgBatchUpdateOrdersResponse.Unmarshal(simRes.Result.MsgResponses[0].Value)

	fmt.Println("simulated spot order hashes", MsgBatchUpdateOrdersResponse.SpotOrderHashes)

	fmt.Println("simulated derivative order hashes", MsgBatchUpdateOrdersResponse.DerivativeOrderHashes)

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
