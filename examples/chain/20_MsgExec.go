package main

import (
	"fmt"
	"os"
	"github.com/InjectiveLabs/sdk-go/client/common"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	authz "github.com/cosmos/cosmos-sdk/codec/types"
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
		os.Getenv("HOME") + "/.injectived",
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

	marketId := "0x0511ddc4e6586f3bfe1acb2dd905f8b8a82c97e1edaef654b12ca7e6031ca0fa"
	amount := decimal.NewFromFloat(2)
	price := cosmtypes.MustNewDecFromStr("22")
	orderSize := chainClient.GetSpotQuantity(amount, cosmtypes.MustNewDecFromStr("10000"), 6)

	order := chainClient.SpotOrder(defaultSubaccountID, &chainclient.SpotOrderData{
		OrderType:    1,
		Quantity:     orderSize,
		Price:        price,
		FeeRecipient: senderAddress.String(),
		MarketId: marketId,
	})

	msg0 := new(exchangetypes.MsgCreateSpotLimitOrder)
	msg0.Sender = senderAddress.String()
	msg0.Order = exchangetypes.SpotOrder(*order)

	grantee := "inj1hkhdaj2a2clmq5jq6mspsggqs32vynpk228q3r"

	msg := &authztypes.MsgExec{
		Grantee: grantee,
		Msgs: []*authz.Any{msg0},
	}

	for i:=0; i<1; i++ {
		err := chainClient.QueueBroadcastMsg(msg)
		if err != nil {
			fmt.Println(err)
		}
	}
	time.Sleep(time.Second * 5)

}
