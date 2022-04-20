package main

import (
	"fmt"
	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/shopspring/decimal"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	"os"
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

	// build orders
	defaultSubaccountID := chainClient.DefaultSubaccount(senderAddress)
	spotOrder := chainClient.SpotOrder(defaultSubaccountID, network, &chainclient.SpotOrderData{
		OrderType:    exchangetypes.OrderType_BUY,
		Quantity:     decimal.NewFromFloat(2),
		Price:        decimal.NewFromFloat(22.55),
		FeeRecipient: senderAddress.String(),
		MarketId:     "0x0511ddc4e6586f3bfe1acb2dd905f8b8a82c97e1edaef654b12ca7e6031ca0fa",
	})
	spotOrders := []exchangetypes.SpotOrder{*spotOrder}

	derivativeOrder := chainClient.DerivativeOrder(defaultSubaccountID, network, &chainclient.DerivativeOrderData{
		OrderType:    exchangetypes.OrderType_BUY,
		Quantity:     decimal.NewFromFloat(2),
		Price:        cosmtypes.MustNewDecFromStr("31000000000"),
		Leverage:     cosmtypes.MustNewDecFromStr("2.5"),
		FeeRecipient: senderAddress.String(),
		MarketId:     "0x4ca0f92fc28be0c9761326016b5a1a2177dd6375558365116b5bdda9abc229ce",
	})
	derivativeOrders := []exchangetypes.DerivativeOrder{*derivativeOrder, *derivativeOrder}

	orderHashes, err := chainClient.ComputeOrderHashes(spotOrders, derivativeOrders)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("computed spot order hashes", orderHashes.Spot)
	fmt.Println("computed derivative order hashes", orderHashes.Derivative)

	// simulate to compare with local order hashes
	msg1 := exchangetypes.MsgBatchCreateSpotLimitOrders{
		Sender: senderAddress.String(),
		Orders: spotOrders,
	}
	msg2 := exchangetypes.MsgBatchCreateDerivativeLimitOrders{
		Sender: senderAddress.String(),
		Orders: derivativeOrders,
	}
	simRes, err := chainClient.SimulateMsg(clientCtx, &msg1, &msg2)
	if err != nil {
		fmt.Println(err)
	}
	simResMsgs := common.MsgResponse(simRes.Result.Data)
	spotResponse := exchangetypes.MsgBatchCreateSpotLimitOrdersResponse{}
	derivativeResponse := exchangetypes.MsgBatchCreateDerivativeLimitOrdersResponse{}
	spotResponse.Unmarshal(simResMsgs[0].Data)
	derivativeResponse.Unmarshal(simResMsgs[1].Data)
	fmt.Println("simulated spot order hashes", spotResponse.OrderHashes)
	fmt.Println("simulated derivative order hashes", derivativeResponse.OrderHashes)
}
