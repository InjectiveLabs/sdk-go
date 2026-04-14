package main

import (
	"context"
	"fmt"
	"os"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/joho/godotenv"

	"github.com/InjectiveLabs/sdk-go/chain/types"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
)

func main() {
	_ = godotenv.Load()
	// network := common.LoadNetwork("mainnet", "k8s")
	network := common.LoadNetwork("mainnet", "sentry0")
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
		os.Getenv("INJECTIVE_PRIVATE_KEY"), // keyring will be used if pk not provided
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
