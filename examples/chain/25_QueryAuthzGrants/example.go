package main

import (
	"context"
	"fmt"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
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

	clientCtx.WithNodeURI(network.TmEndpoint)
	clientCtx = clientCtx.WithClient(tmRPC)

	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network.ChainGrpcEndpoint,
		common.OptionTLSCert(network.ChainTlsCert),
		common.OptionGasPrices("500000000inj"),
	)

	ctx := context.Background()

	granter := "inj14au322k9munkmx5wrchz9q30juf5wjgz2cfqku"
	grantee := "inj1hkhdaj2a2clmq5jq6mspsggqs32vynpk228q3r"
	msg_type_url := "/injective.exchange.v1beta1.MsgCreateSpotLimitOrder"

	req := authztypes.QueryGrantsRequest{
		Granter:    granter,
		Grantee:    grantee,
		MsgTypeUrl: msg_type_url,
	}

	res, err := chainClient.GetAuthzGrants(ctx, req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

}
