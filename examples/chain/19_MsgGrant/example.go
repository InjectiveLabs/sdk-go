package main

import (
	"fmt"
	"os"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"

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

	// build generic authz msg
	grantee := "inj1hkhdaj2a2clmq5jq6mspsggqs32vynpk228q3r"
	auth := authztypes.NewGenericAuthorization("/injective.exchange.v1beta1.MsgCreateSpotLimitOrder")
	authAny := codectypes.UnsafePackAny(auth)
	expireIn := time.Now().AddDate(1, 0, 0)

	msg := &authztypes.MsgGrant{
		Granter: senderAddress.String(),
		Grantee: grantee,
		Grant: authztypes.Grant{
			Authorization: authAny,
			Expiration:    expireIn,
		},
	}

	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network.ChainGrpcEndpoint,
		common.OptionTLSCert(network.ChainTlsCert),
		common.OptionGasPrices("500000000inj"),
	)

	if err != nil {
		fmt.Println(err)
	}

	err = chainClient.QueueBroadcastMsg(msg)

	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(time.Second * 5)
}
