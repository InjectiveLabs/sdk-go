package main

import (
	"fmt"
	"os"

	"github.com/InjectiveLabs/sdk-go/client/common"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
)

func mustNewIntFromString(s string) cosmtypes.Int {
	newInt, ok := cosmtypes.NewIntFromString(s)
	if !ok {
		panic(fmt.Errorf("cannot new int from %s", s))
	}

	return newInt
}

func main() {
	// network := common.LoadNetwork("mainnet", "k8s")
	network := common.LoadNetwork("testnet", "k8s")
	tmRPC, err := rpchttp.New(network.TmEndpoint, "/websocket")
	if err != nil {
		fmt.Println(err)
		return
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
	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint).WithClient(tmRPC)
	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network.ChainGrpcEndpoint,
		common.OptionTLSCert(network.ChainTlsCert),
		common.OptionGasPrices("500000000inj"),
	)
	if err != nil {
		panic(err)
	}

	defaultSubaccountID := chainClient.DefaultSubaccount(senderAddress)
	offchainContract := "inj1v4s3v3nmsx6szyxqgcnqdfseucu4ccm65sh0xr"
	// redeem 0.001 offchain vault LP token
	msgSubscribe := &exchangetypes.MsgPrivilegedExecuteContract{
		Sender:          senderAddress.String(),
		ContractAddress: offchainContract,
		Data: fmt.Sprintf(`
			{
				"args":{
					"Redeem":{
						"args":{
							"redeemer_subaccount_id":"%s",
							"redemption_type":{
								"DerivativeRedemptionType":"PositionAndQuote"
							},
							"slippage":{"max_penalty":"0.1"}
						}
					}
				},
				"name":"VaultRedeem",
				"origin":"%s"
			}`, defaultSubaccountID.String(), senderAddress.String()),
		Funds: cosmtypes.NewCoins(
			cosmtypes.NewCoin(fmt.Sprintf("factory/%s/lp", offchainContract), mustNewIntFromString("1000000000000000")),
		).String(),
	}

	broadcastResult, err := chainClient.SyncBroadcastMsg(msgSubscribe)
	if err != nil {
		panic(err)
	}

	fmt.Println("offchain redeem tx:", broadcastResult.TxResponse.TxHash)
}
