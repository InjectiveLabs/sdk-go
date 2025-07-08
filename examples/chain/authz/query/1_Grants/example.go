package main

import (
	"context"
	"fmt"

	"github.com/InjectiveLabs/sdk-go/client"

	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"

	"os"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")
	tmClient, err := rpchttp.New(network.TmEndpoint)
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

	chainClient, err := chainclient.NewChainClientV2(
		clientCtx,
		network,
		common.OptionGasPrices(client.DefaultGasPriceWithDenom),
	)

	if err != nil {
		panic(err)
	}

	granter := "inj14au322k9munkmx5wrchz9q30juf5wjgz2cfqku"
	grantee := "inj1hkhdaj2a2clmq5jq6mspsggqs32vynpk228q3r"
	msg_type_url := "/injective.exchange.v1beta1.MsgCreateSpotLimitOrder"

	req := authztypes.QueryGrantsRequest{
		Granter:    granter,
		Grantee:    grantee,
		MsgTypeUrl: msg_type_url,
	}

	ctx := context.Background()

	res, err := chainClient.GetAuthzGrants(ctx, req)
	if err != nil {
		panic(err)
	}

	jsonResponse, err := clientCtx.Codec.MarshalJSON(res)
	if err != nil {
		panic(err)
	}

	fmt.Print(string(jsonResponse))
}
