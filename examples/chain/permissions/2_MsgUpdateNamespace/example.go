package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

	permissionstypes "github.com/InjectiveLabs/sdk-go/chain/permissions/types"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
)

func main() {
	network := common.LoadNetwork("devnet", "lb")
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

	chainClient, err := chainclient.NewChainClientV2(
		clientCtx,
		network,
	)

	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	gasPrice := chainClient.CurrentChainGasPrice(ctx)
	// adjust gas price to make it valid even if it changes between the time it is requested and the TX is broadcasted
	gasPrice = int64(float64(gasPrice) * 1.1)
	chainClient.SetGasPrice(gasPrice)

	denom := "factory/inj1hkhdaj2a2clmq5jq6mspsggqs32vynpk228q3r/inj_test"

	role1 := permissionstypes.Role{
		Name:        "EVERYONE",
		RoleId:      0,
		Permissions: uint32(permissionstypes.Action_UNSPECIFIED),
	}
	role2 := permissionstypes.Role{
		Name:        "user",
		RoleId:      2,
		Permissions: uint32(permissionstypes.Action_RECEIVE | permissionstypes.Action_SEND),
	}

	role_manager := permissionstypes.RoleManager{
		Manager: "inj1manageraddress",
		Roles:   []string{"admin", "user"},
	}

	policy_status1 := permissionstypes.PolicyStatus{
		Action:     permissionstypes.Action_MINT,
		IsDisabled: true,
		IsSealed:   false,
	}
	policy_status2 := permissionstypes.PolicyStatus{
		Action:     permissionstypes.Action_BURN,
		IsDisabled: false,
		IsSealed:   true,
	}

	policy_manager_capability := permissionstypes.PolicyManagerCapability{
		Manager:    "inj1policymanageraddress",
		Action:     permissionstypes.Action_MODIFY_ROLE_PERMISSIONS,
		CanDisable: true,
		CanSeal:    false,
	}

	msg := &permissionstypes.MsgUpdateNamespace{
		Sender:                    senderAddress.String(),
		Denom:                     denom,
		ContractHook:              &permissionstypes.MsgUpdateNamespace_SetContractHook{NewValue: "inj19ld6swyldyujcn72j7ugnu9twafhs9wxlyye5m"},
		RolePermissions:           []*permissionstypes.Role{&role1, &role2},
		RoleManagers:              []*permissionstypes.RoleManager{&role_manager},
		PolicyStatuses:            []*permissionstypes.PolicyStatus{&policy_status1, &policy_status2},
		PolicyManagerCapabilities: []*permissionstypes.PolicyManagerCapability{&policy_manager_capability},
	}

	//AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	pollInterval := 100 * time.Millisecond
	response, err := chainClient.SyncBroadcastMsg(ctx, &pollInterval, 5, msg)

	if err != nil {
		panic(err)
	}

	str, _ := json.MarshalIndent(response, "", "\t")
	fmt.Print(string(str))

	gasPrice = chainClient.CurrentChainGasPrice(ctx)
	// adjust gas price to make it valid even if it changes between the time it is requested and the TX is broadcasted
	gasPrice = int64(float64(gasPrice) * 1.1)
	chainClient.SetGasPrice(gasPrice)
}
