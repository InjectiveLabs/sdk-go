package main

import (
	"fmt"
	"github.com/InjectiveLabs/sdk-go/client/common"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	"github.com/shopspring/decimal"
	"time"
)

func main() {
	network := common.LoadNetwork("testnet", "k8s")
	tmRPC, err := rpchttp.New(network.TmEndpoint, "/websocket")
	if err != nil {
		fmt.Println(err)
	}

	senderAddress, cosmosKeyring, err := chainclient.InitCosmosKeyring(
		"/Users/akalantzis/.injectived",
		"injectived",
		"file",
		"inj-user",
		"12345678",
		"5d386fbdbf11f1141010f81a46b40f94887367562bd33b452bbaa6ce1cd1381e", // keyring will be used if pk not provided
		false,
	)

	clientCtx, err := chainclient.NewClientContext(
		network.ChainId,
		senderAddress.String(),
		cosmosKeyring,
	)
	clientCtx.WithNodeURI(network.TmEndpoint)
	clientCtx = clientCtx.WithClient(tmRPC)

	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network.ChainGrpcEndpoint,
		common.OptionTLSCert(network.ChainTlsCert),
		common.OptionGasPrices("500000000inj"),
	)

	defaultSubaccountID := chainClient.DefaultSubaccount(senderAddress)

	marketId := "0x0511ddc4e6586f3bfe1acb2dd905f8b8a82c97e1edaef654b12ca7e6031ca0fa"
	amount := decimal.NewFromFloat(2)
	price := cosmtypes.MustNewDecFromStr("30")
	orderSize := chainClient.GetSpotQuantity(amount, cosmtypes.MustNewDecFromStr("10000"), 6)


	order := chainClient.SpotOrder(defaultSubaccountID, &chainclient.SpotOrderData{
		OrderType:    exchangetypes.OrderType_BUY,
		Quantity:     orderSize,
		Price:        price,
		FeeRecipient: senderAddress.String(),
		MarketId: marketId,
	})

	msg := new(exchangetypes.MsgCreateSpotMarketOrder)
	msg.Sender = senderAddress.String()
	msg.Order = exchangetypes.SpotOrder(*order)
	CosMsgs := []cosmtypes.Msg{msg}
	for i:=0; i<1; i++ {
		err := chainClient.QueueBroadcastMsg(CosMsgs...)
		if err != nil {
			fmt.Println(err)
		}
	}
	time.Sleep(time.Second * 5)

}
