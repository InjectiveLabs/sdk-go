package main

import (
	"fmt"
	"os"
	"time"

	exchangev2types "github.com/InjectiveLabs/sdk-go/chain/exchange/types/v2"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
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

	defaultSubaccountID := chainClient.DefaultSubaccount(senderAddress)

	marketId := "0xa508cb32923323679f29a032c70342c147c17d0145625922b0ef22e955c844c0"
	orderHash := "0x17196096ffc32ad088ef959ad95b4cc247a87c7c9d45a2500b81ab8f5a71da5a"
	cid := "666C5452-B563-445F-977B-539D79C482EE"

	orderWithHash := chainClient.OrderCancelV2(defaultSubaccountID, &chainclient.OrderCancelData{
		MarketId:  marketId,
		OrderHash: orderHash,
	})

	orderWithCid := chainClient.OrderCancelV2(defaultSubaccountID, &chainclient.OrderCancelData{
		MarketId: marketId,
		Cid:      cid,
	})

	msg := new(exchangev2types.MsgBatchCancelSpotOrders)
	msg.Sender = senderAddress.String()
	msg.Data = []exchangev2types.OrderData{*orderWithHash, *orderWithCid}
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

	gasPrice = chainClient.CurrentChainGasPrice()
	// adjust gas price to make it valid even if it changes between the time it is requested and the TX is broadcasted
	gasPrice = int64(float64(gasPrice) * 1.1)
	chainClient.SetGasPrice(gasPrice)
}
