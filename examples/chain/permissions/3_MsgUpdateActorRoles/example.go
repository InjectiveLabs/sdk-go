package main

import (
	"encoding/json"
	"fmt"
	"os"

	permissionstypes "github.com/InjectiveLabs/sdk-go/chain/permissions/types"
	"github.com/InjectiveLabs/sdk-go/client"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
)

func main() {
	network := common.LoadNetwork("devnet", "lb")
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
		"f9db9bf330e23cb7839039e944adef6e9df447b90b503d5b4464c90bea9022f3", // keyring will be used if pk not provided
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

	denom := "factory/inj1hkhdaj2a2clmq5jq6mspsggqs32vynpk228q3r/inj_test"
	roleActors1 := permissionstypes.RoleActors{
		Role:   "admin",
		Actors: []string{"inj1actoraddress1", "inj1actoraddress2"},
	}
	roleActors2 := permissionstypes.RoleActors{
		Role:   "user",
		Actors: []string{"inj1actoraddress3"},
	}
	roleActors3 := permissionstypes.RoleActors{
		Role:   "user",
		Actors: []string{"inj1actoraddress4"},
	}
	roleActors4 := permissionstypes.RoleActors{
		Role:   "admin",
		Actors: []string{"inj1actoraddress5"},
	}

	msg := &permissionstypes.MsgUpdateActorRoles{
		Sender:             senderAddress.String(),
		Denom:              denom,
		RoleActorsToAdd:    []*permissionstypes.RoleActors{&roleActors1, &roleActors2},
		RoleActorsToRevoke: []*permissionstypes.RoleActors{&roleActors3, &roleActors4},
	}

	//AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	response, err := chainClient.SyncBroadcastMsg(msg)

	if err != nil {
		panic(err)
	}

	str, _ := json.MarshalIndent(response, "", " ")
	fmt.Print(string(str))
}
