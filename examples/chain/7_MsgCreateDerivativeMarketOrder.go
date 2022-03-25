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

	marketId := "0x4ca0f92fc28be0c9761326016b5a1a2177dd6375558365116b5bdda9abc229ce"
	amount := decimal.NewFromFloat(2)
	price := cosmtypes.MustNewDecFromStr("31000000000") //31,000
	leverage := cosmtypes.MustNewDecFromStr("2.5")
	margin := cosmtypes.MustNewDecFromStr(fmt.Sprint(amount)).Mul(price).Quo(leverage)

	orderSize := chainClient.GetDerivativeQuantity(amount, cosmtypes.MustNewDecFromStr("0.0001"))
	orderPrice := chainClient.GetDerivativePrice(price, cosmtypes.MustNewDecFromStr("1000"))
	orderMargin := chainClient.GetDerivativePrice(margin, cosmtypes.MustNewDecFromStr("1000"))

	order := chainClient.DerivativeOrder(defaultSubaccountID, &chainclient.DerivativeOrderData{
		OrderType:    exchangetypes.OrderType_BUY,
		Quantity:     orderSize,
		Price:        orderPrice,
		Margin: orderMargin,
		FeeRecipient: senderAddress.String(),
		MarketId: marketId,
	})

	msg := new(exchangetypes.MsgCreateDerivativeMarketOrder)
	msg.Sender = senderAddress.String()
	msg.Order = exchangetypes.DerivativeOrder(*order)
	CosMsgs := []cosmtypes.Msg{msg}
	for i:=0; i<1; i++ {
		err := chainClient.QueueBroadcastMsg(CosMsgs...)
		if err != nil {
			fmt.Println(err)
		}
	}
	time.Sleep(time.Second * 5)

}
