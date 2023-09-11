package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/InjectiveLabs/sdk-go/client/common"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	cometbfttypes "github.com/cometbft/cometbft/abci/types"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
)

func mustNewIntFromString(s string) cosmtypes.Int {
	newInt, ok := cosmtypes.NewIntFromString(s)
	if !ok {
		panic(fmt.Errorf("cannot new int from %s", s))
	}

	return newInt
}

func mintAmountFromResponse(resp *tx.GetTxResponse) cosmtypes.Int {
	var lpBalanceChangeEvent cometbfttypes.Event
	for _, ev := range resp.TxResponse.Events {
		if ev.Type == "wasm-lp_balance_changed" {
			lpBalanceChangeEvent = ev
			break
		}
	}

	for _, attribute := range lpBalanceChangeEvent.Attributes {
		if attribute.Key == "mint_amount" {
			return mustNewIntFromString(attribute.Value)
		}
	}
	return cosmtypes.ZeroInt()
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
	offchainContract := "inj1ckmdhdz7r8glfurckgtg0rt7x9uvner4ygqhlv"

	// deposit 10 INJ and 80 USDT (testnet)
	msgSubscribe := &exchangetypes.MsgPrivilegedExecuteContract{
		Sender:          senderAddress.String(),
		ContractAddress: offchainContract,
		Data: fmt.Sprintf(`
			{
				"origin":"%s",
				"name":"Subscribe",
				"args":{
					"Subscribe":{
						"args":{
							"subscriber_subaccount_id":"%s"
						}
					}
				}
			}`, senderAddress.String(), defaultSubaccountID.String()),
		Funds: cosmtypes.NewCoins(
			cosmtypes.NewCoin("inj", mustNewIntFromString("10000000000000000000")),
			cosmtypes.NewCoin("peggy0x87aB3B4C8661e07D6372361211B96ed4Dc36B1B5", mustNewIntFromString("80000000")),
		).String(),
	}

	broadcastResult, err := chainClient.SyncBroadcastMsg(msgSubscribe)
	if err != nil {
		panic(err)
	}

	fmt.Println("offchain subscribe tx:", broadcastResult.TxResponse.TxHash)

	// tx result -> minted amount
	retryTime := 5
	var (
		lastErr    error
		txResponse *tx.GetTxResponse
	)
	for i := 0; i < retryTime; i++ {
		txResponse, lastErr = chainClient.GetTx(context.Background(), broadcastResult.TxResponse.TxHash)
		if lastErr == nil {
			break
		}
		time.Sleep(200 * time.Millisecond)
	}
	if lastErr != nil {
		panic(err)
	}

	fmt.Println("minted amount:", mintAmountFromResponse(txResponse).String()+fmt.Sprintf("lp/%s/lp", offchainContract))
}
