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
	}
	simResMsgs := common.MsgResponse(simRes.Result.Data)
	msgBatchCreateDerivativeLimitOrdersResponse := exchangetypes.MsgBatchCreateDerivativeLimitOrdersResponse{}
	msgBatchCreateDerivativeLimitOrdersResponse.Unmarshal(simResMsgs[0].Data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("simulated order hashes", msgBatchCreateDerivativeLimitOrdersResponse.OrderHashes)

	err = chainClient.QueueBroadcastMsg(msg)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second * 5)
}
