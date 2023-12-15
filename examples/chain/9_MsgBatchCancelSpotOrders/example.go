package main

import (
	"fmt"
	"github.com/InjectiveLabs/sdk-go/client"
	"os"
	"time"

	"github.com/InjectiveLabs/sdk-go/client/common"

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

	defaultSubaccountID := chainClient.DefaultSubaccount(senderAddress)

	marketId := "0xa508cb32923323679f29a032c70342c147c17d0145625922b0ef22e955c844c0"
	orderHash := "0x17196096ffc32ad088ef959ad95b4cc247a87c7c9d45a2500b81ab8f5a71da5a"
	cid := "666C5452-B563-445F-977B-539D79C482EE"

	orderWithHash := chainClient.OrderCancel(defaultSubaccountID, &chainclient.OrderCancelData{
		MarketId:  marketId,
		OrderHash: orderHash,
	})

	orderWithCid := chainClient.OrderCancel(defaultSubaccountID, &chainclient.OrderCancelData{
		MarketId: marketId,
		Cid:      cid,
	})

	msg := new(exchangetypes.MsgBatchCancelSpotOrders)
	msg.Sender = senderAddress.String()
	msg.Data = []exchangetypes.OrderData{*orderWithHash, *orderWithCid}
	CosMsgs := []cosmtypes.Msg{msg}

	// AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	err = chainClient.QueueBroadcastMsg(CosMsgs...)

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
