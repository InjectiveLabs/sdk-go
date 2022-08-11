package main

import (
	"context"
	"fmt"
	"os"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"

	"github.com/InjectiveLabs/sdk-go/chain/types"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
)

func main() {
	// network := common.LoadNetwork("mainnet", "k8s")
	network := common.LoadNetwork("mainnet", "sentry0")
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

	queryClient := authtypes.NewQueryClient(clientCtx)

	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	address := "inj1akxycslq8cjt0uffw4rjmfm3echchptu52a2dq"

	res, err := queryClient.Account(ctx, &authtypes.QueryAccountRequest{
		Address: address,
	})
	if err != nil {
		fmt.Println(err)
	}

	var account types.EthAccount

	clientCtx.Codec.MustUnmarshal(res.GetAccount().GetValue(), &account)

	fmt.Println(account.String())

}
