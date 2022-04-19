package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	"os"
	"time"

	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
)

func main() {
	// network := common.LoadNetwork("mainnet", "k8s")
	network := common.LoadNetwork("testnet", "k8s")
	tmRPC, err := rpchttp.New(network.TmEndpoint, "/websocket")
	if err != nil {
		fmt.Println(err)
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

	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint).WithClient(tmRPC)

	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network.ChainGrpcEndpoint,
		common.OptionTLSCert(network.ChainTlsCert),
		common.OptionGasPrices("500000000inj"),
	)

	if err != nil {
		fmt.Println(err)
	}

	defaultSubaccountID := chainClient.DefaultSubaccount(senderAddress)

	marketId := "0x0511ddc4e6586f3bfe1acb2dd905f8b8a82c97e1edaef654b12ca7e6031ca0fa"

	amount := decimal.NewFromFloat(2)
	price := decimal.NewFromFloat(22.53)

	order := chainClient.SpotOrder(defaultSubaccountID, network, &chainclient.SpotOrderData{
		OrderType:    exchangetypes.OrderType_BUY,
		Quantity:     amount,
		Price:        price,
		FeeRecipient: senderAddress.String(),
		MarketId:     marketId,
	})

	msg := new(exchangetypes.MsgCreateSpotLimitOrder)
	msg.Sender = senderAddress.String()
	msg.Order = exchangetypes.SpotOrder(*order)

	simRes, err := chainClient.SimulateMsg(clientCtx, msg)
	if err != nil {
		fmt.Println(err)
	}
	simResMsgs := common.MsgResponse(simRes.Result.Data)
	msgRes := exchangetypes.MsgCreateSpotLimitOrderResponse{}
	msgRes.Unmarshal(simResMsgs[0].Data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("simulated order hash", msgRes.OrderHash)

	localOrderHashes, err := chainClient.ComputeSpotOrderHash([]exchangetypes.SpotOrder{msg.Order})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("computed order hash", localOrderHashes)

	err = chainClient.QueueBroadcastMsg(msg)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second * 5)
}
