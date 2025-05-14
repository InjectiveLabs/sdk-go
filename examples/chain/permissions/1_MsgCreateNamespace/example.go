package main

import (
	"encoding/json"
	"fmt"
	"os"

	permissionstypes "github.com/InjectiveLabs/sdk-go/chain/permissions/types"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
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

	denom := "factory/inj1hkhdaj2a2clmq5jq6mspsggqs32vynpk228q3r/inj_test"
	role1 := permissionstypes.Role{
		Name:        "EVERYONE",
		RoleId:      0,
		Permissions: uint32(permissionstypes.Action_RECEIVE | permissionstypes.Action_SEND),
	}
	role2 := permissionstypes.Role{
		Name:        "admin",
		RoleId:      1,
		Permissions: uint32(permissionstypes.Action_MODIFY_ROLE_PERMISSIONS),
	}
	role3 := permissionstypes.Role{
		Name:   "user",
		RoleId: 2,
		Permissions: uint32(
			permissionstypes.Action_MINT |
				permissionstypes.Action_RECEIVE |
				permissionstypes.Action_BURN |
				permissionstypes.Action_SEND),
	}

	actor_role1 := permissionstypes.ActorRoles{
		Actor: "inj1specificactoraddress",
		Roles: []string{"admin"},
	}
	actor_role2 := permissionstypes.ActorRoles{
		Actor: "inj1anotheractoraddress",
		Roles: []string{"user"},
	}

	role_manager := permissionstypes.RoleManager{
		Manager: "inj1manageraddress",
		Roles:   []string{"admin"},
	}

	policy_status1 := permissionstypes.PolicyStatus{
		Action:     permissionstypes.Action_MINT,
		IsDisabled: false,
		IsSealed:   false,
	}
	policy_status2 := permissionstypes.PolicyStatus{
		Action:     permissionstypes.Action_BURN,
		IsDisabled: false,
		IsSealed:   false,
	}

	policy_manager_capability := permissionstypes.PolicyManagerCapability{
		Manager:    "inj1policymanageraddress",
		Action:     permissionstypes.Action_MODIFY_CONTRACT_HOOK,
		CanDisable: true,
		CanSeal:    false,
	}

	namespace := permissionstypes.Namespace{
		Denom:                     denom,
		ContractHook:              "",
		RolePermissions:           []*permissionstypes.Role{&role1, &role2, &role3},
		ActorRoles:                []*permissionstypes.ActorRoles{&actor_role1, &actor_role2},
		RoleManagers:              []*permissionstypes.RoleManager{&role_manager},
		PolicyStatuses:            []*permissionstypes.PolicyStatus{&policy_status1, &policy_status2},
		PolicyManagerCapabilities: []*permissionstypes.PolicyManagerCapability{&policy_manager_capability},
	}

	msg := &permissionstypes.MsgCreateNamespace{
		Sender:    senderAddress.String(),
		Namespace: namespace,
	}

	//AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	response, err := chainClient.SyncBroadcastMsg(msg)

	if err != nil {
		panic(err)
	}

	str, _ := json.MarshalIndent(response, "", " ")
	fmt.Print(string(str))

	gasPrice = chainClient.CurrentChainGasPrice()
	// adjust gas price to make it valid even if it changes between the time it is requested and the TX is broadcasted
	gasPrice = int64(float64(gasPrice) * 1.1)
	chainClient.SetGasPrice(gasPrice)
}
