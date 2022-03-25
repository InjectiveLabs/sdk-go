package main

import (
	"fmt"
	"github.com/InjectiveLabs/sdk-go/client/common"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
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

	msg := &exchangetypes.MsgIncreasePositionMargin{
		Sender: senderAddress.String(),
		MarketId: "0x4ca0f92fc28be0c9761326016b5a1a2177dd6375558365116b5bdda9abc229ce",
		SourceSubaccountId: "0xaf79152ac5df276d9a8e1e2e22822f9713474902000000000000000000000000",
		DestinationSubaccountId: "0xaf79152ac5df276d9a8e1e2e22822f9713474902000000000000000000000000",
		Amount: cosmtypes.MustNewDecFromStr("100000000"), //100 USDT
	}

	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network.ChainGrpcEndpoint,
		common.OptionTLSCert(network.ChainTlsCert),
		common.OptionGasPrices("500000000inj"),
	)

	for i:=0; i<1; i++ {
		err := chainClient.QueueBroadcastMsg(msg)
		if err != nil {
			fmt.Println(err)
		}
	}
	time.Sleep(time.Second * 5)
}
