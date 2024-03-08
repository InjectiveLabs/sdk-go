package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	oracletypes "github.com/InjectiveLabs/sdk-go/chain/oracle/types"
	exchangeclient "github.com/InjectiveLabs/sdk-go/client/exchange"

	"github.com/cosmos/cosmos-sdk/types"

	"github.com/InjectiveLabs/sdk-go/client"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
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
		common.OptionGasPrices(client.DefaultGasPriceWithDenom),
	)
	if err != nil {
		panic(err)
	}

	exchangeClient, err := exchangeclient.NewExchangeClient(network)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	marketsAssistant, err := chainclient.NewMarketsAssistantInitializedFromChain(ctx, exchangeClient)
	if err != nil {
		panic(err)
	}

	quoteToken := marketsAssistant.AllTokens()["USDC"]
	minPriceTickSize := types.MustNewDecFromStr("0.01")
	minQuantityTickSize := types.MustNewDecFromStr("0.001")

	chainMinPriceTickSize := minPriceTickSize.Mul(types.NewDecFromIntWithPrec(types.NewInt(1), int64(quoteToken.Decimals)))
	chainMinQuantityTickSize := minQuantityTickSize

	msg := &exchangetypes.MsgInstantPerpetualMarketLaunch{
		Sender:                 senderAddress.String(),
		Ticker:                 "INJ/USDC PERP",
		QuoteDenom:             quoteToken.Denom,
		OracleBase:             "INJ",
		OracleQuote:            "USDC",
		OracleScaleFactor:      6,
		OracleType:             oracletypes.OracleType_Band,
		MakerFeeRate:           types.MustNewDecFromStr("-0.0001"),
		TakerFeeRate:           types.MustNewDecFromStr("0.001"),
		InitialMarginRatio:     types.MustNewDecFromStr("0.33"),
		MaintenanceMarginRatio: types.MustNewDecFromStr("0.095"),
		MinPriceTickSize:       chainMinPriceTickSize,
		MinQuantityTickSize:    chainMinQuantityTickSize,
	}

	//AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	response, err := chainClient.AsyncBroadcastMsg(msg)

	if err != nil {
		panic(err)
	}

	str, _ := json.MarshalIndent(response, "", " ")
	fmt.Print(string(str))
}
