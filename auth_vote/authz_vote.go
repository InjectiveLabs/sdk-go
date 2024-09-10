package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/InjectiveLabs/sdk-go/client"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")
	tmClient, err := rpchttp.New(network.TmEndpoint, "/websocket")
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
	txFactory = txFactory.WithGasPrices(client.DefaultGasPriceWithDenom)
	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network,
		common.OptionTxFactory(&txFactory),
	)

	if err != nil {
		panic(err)
	}

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
}
