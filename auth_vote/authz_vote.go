package main

import (
	"encoding/json"
	"fmt"
	"os"

	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

//revive:disable:function-length // this is an example script
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
		"gov_account",
		"",
		"", // keyring will be used if pk not provided
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

	txFactory := chainclient.NewTxFactory(clientCtx)
	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network,
		common.OptionTxFactory(&txFactory),
	)

	if err != nil {
		panic(err)
	}

	gasPrice := chainClient.CurrentChainGasPrice()
	// adjust gas price to make it valid even if it changes between the time it is requested and the TX is broadcasted
	gasPrice = int64(float64(gasPrice) * 1.1)
	chainClient.SetGasPrice(gasPrice)

	// note that we use grantee keyring to send the msg on behalf of granter here
	// sender, subaccount are from granter
	validators := []string{"inj156t3yxd4udv0h9gwagfcmwnmm3quy0npqc7pks", "inj16nd8yqxe9p6ggnrz58qr7dxn5y2834yendward"}
	grantee := senderAddress.String()
	proposalId := uint64(375)
	var msgs = make([]sdk.Msg, 0)

	for _, validator := range validators {
		msgVote := v1beta1.MsgVote{
			ProposalId: proposalId,
			Voter:      validator,
			Option:     v1beta1.OptionYes,
		}

		msg0Bytes, _ := msgVote.Marshal()
		msg0Any := &codectypes.Any{}
		msg0Any.TypeUrl = sdk.MsgTypeURL(&msgVote)
		msg0Any.Value = msg0Bytes

		msg := &authztypes.MsgExec{
			Grantee: grantee,
			Msgs:    []*codectypes.Any{msg0Any},
		}

		sdkMsg := sdk.Msg(msg)
		msgs = append(msgs, sdkMsg)
	}

	// AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	response, err := chainClient.AsyncBroadcastMsg(msgs...)

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
