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
	response, err := chainClient.SyncBroadcastMsg(msg)

	if err != nil {
		panic(err)
	}

	str, _ := json.MarshalIndent(response, "", " ")
	fmt.Print(string(str))
}
