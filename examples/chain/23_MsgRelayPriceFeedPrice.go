package main

import (
	"fmt"
	"os"
	"github.com/InjectiveLabs/sdk-go/client/common"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	oracletypes "github.com/InjectiveLabs/sdk-go/chain/oracle/types"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
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

	clientCtx, err := chainclient.NewClientContext(
		network.ChainId,
		senderAddress.String(),
		cosmosKeyring,
	)
	clientCtx.WithNodeURI(network.TmEndpoint)
	clientCtx = clientCtx.WithClient(tmRPC)

	price := []cosmtypes.Dec{cosmtypes.MustNewDecFromStr("100")}
	base := []string{"BAYC"}
	quote := []string{"WETH"}

	msg := &oracletypes.MsgRelayPriceFeedPrice{
		Sender: senderAddress.String(),
		Price: price,
		Base: base,
		Quote: quote,
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
