package main

import (
	"fmt"
	"os"
	"time"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")
	remoteAddress := fmt.Sprintf("%s/websocket", network.TmEndpoint)
	tmClient, err := rpchttp.New(remoteAddress)
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

	granter := senderAddress.String()
	grantee := "inj1hkhdaj2a2clmq5jq6mspsggqs32vynpk228q3r"
	expireIn := time.Now().AddDate(1, 0, 0) // years months days

	// GENERIC AUTHZ
	// msgtype := "/injective.exchange.v1beta1.MsgCreateSpotLimitOrder"
	// msg := chainClient.BuildGenericAuthz(granter, grantee, msgtype, expireIn)

	// TYPED AUTHZ
	msg := chainClient.BuildExchangeAuthz(
		granter,
		grantee,
		chainclient.CreateSpotLimitOrderAuthz,
		chainClient.DefaultSubaccount(senderAddress).String(),
		[]string{"0xe0dc13205fb8b23111d8555a6402681965223135d368eeeb964681f9ff12eb2a"},
		expireIn,
	)

	// AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	err = chainClient.QueueBroadcastMsg(msg)

	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(time.Second * 5)

	gasFee, err := chainClient.GetGasFee()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("gas fee:", gasFee, "INJ")

	gasPrice = chainClient.CurrentChainGasPrice()
	// adjust gas price to make it valid even if it changes between the time it is requested and the TX is broadcasted
	gasPrice = int64(float64(gasPrice) * 1.1)
	chainClient.SetGasPrice(gasPrice)
}
