package main

import (
	"fmt"
	"os"
	"time"

	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/shopspring/decimal"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"

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

	clientCtx.WithNodeURI(network.TmEndpoint)
	clientCtx = clientCtx.WithClient(tmRPC)

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

	smarketId := "0x0511ddc4e6586f3bfe1acb2dd905f8b8a82c97e1edaef654b12ca7e6031ca0fa"
	samount := decimal.NewFromFloat(2)
	sprice := decimal.NewFromFloat(22.5)
	smarketIds := []string{"0xa508cb32923323679f29a032c70342c147c17d0145625922b0ef22e955c844c0"}

	spot_order := chainClient.SpotOrder(defaultSubaccountID, network, &chainclient.SpotOrderData{
		OrderType:    exchangetypes.OrderType_BUY,
		Quantity:     samount,
		Price:        sprice,
		FeeRecipient: senderAddress.String(),
		MarketId:     smarketId,
	})

	dmarketId := "0x4ca0f92fc28be0c9761326016b5a1a2177dd6375558365116b5bdda9abc229ce"
	damount := decimal.NewFromFloat(2)
	dprice := cosmtypes.MustNewDecFromStr("31000000000") //31,000
	dleverage := cosmtypes.MustNewDecFromStr("2.5")
	dmarketIds := []string{"0x4ca0f92fc28be0c9761326016b5a1a2177dd6375558365116b5bdda9abc229ce"}

	derivative_order := chainClient.DerivativeOrder(defaultSubaccountID, network, &chainclient.DerivativeOrderData{
		OrderType:    exchangetypes.OrderType_BUY,
		Quantity:     damount,
		Price:        dprice,
		Leverage:     dleverage,
		FeeRecipient: senderAddress.String(),
		MarketId:     dmarketId,
	})

	msg := new(exchangetypes.MsgBatchUpdateOrders)
	msg.Sender = senderAddress.String()
	msg.SubaccountId = defaultSubaccountID.Hex()
	msg.SpotOrdersToCreate = []*exchangetypes.SpotOrder{spot_order}
	msg.DerivativeOrdersToCreate = []*exchangetypes.DerivativeOrder{derivative_order}
	msg.SpotMarketIdsToCancelAll = smarketIds
	msg.DerivativeMarketIdsToCancelAll = dmarketIds

	CosMsgs := []cosmtypes.Msg{msg}
	for i := 0; i < 1; i++ {
		err := chainClient.QueueBroadcastMsg(CosMsgs...)
		if err != nil {
			fmt.Println(err)
		}
	}
	time.Sleep(time.Second * 5)

}
