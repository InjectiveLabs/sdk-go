package main

import (
	"encoding/json"
	"fmt"
	"os"

	"cosmossdk.io/math"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
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

	minPriceTickSize := math.LegacyMustNewDecFromStr("0.01")
	minQuantityTickSize := math.LegacyMustNewDecFromStr("0.01")
	minNotional := math.LegacyMustNewDecFromStr("2")

	chainMinPriceTickSize := minPriceTickSize.Mul(math.LegacyNewDecFromIntWithPrec(math.NewInt(1), int64(6)))
	chainMinPriceTickSize = chainMinPriceTickSize.Quo(math.LegacyNewDecFromIntWithPrec(math.NewInt(1), int64(18)))

	chainMinQuantityTickSize := minQuantityTickSize.Mul(math.LegacyNewDecFromIntWithPrec(math.NewInt(1), int64(18)))
	chainMinNotional := minNotional.Mul(math.LegacyNewDecFromIntWithPrec(math.NewInt(1), int64(6)))

	msg := &exchangetypes.MsgUpdateSpotMarket{
		Admin:                  senderAddress.String(),
		MarketId:               "0x215970bfdea5c94d8e964a759d3ce6eae1d113900129cc8428267db5ccdb3d1a",
		NewTicker:              "INJ/USDC 2",
		NewMinPriceTickSize:    chainMinPriceTickSize,
		NewMinQuantityTickSize: chainMinQuantityTickSize,
		NewMinNotional:         chainMinNotional,
	}

	// AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	response, err := chainClient.AsyncBroadcastMsg(msg)

	if err != nil {
		panic(err)
	}

	str, _ := json.MarshalIndent(response, "", " ")
	fmt.Print(string(str))

	gasPrice = chainClient.CurrentChainGasPrice()
	// adjust gas price to make it valid even if it changes between the time it is requested and the TX is broadcasted
	gasPrice = int64(float64(gasPrice) * 1.1)
	chainClient.SetGasPrice(gasPrice)
}
