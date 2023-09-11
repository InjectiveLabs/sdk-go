package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"cosmossdk.io/math"
	"github.com/InjectiveLabs/sdk-go/client/common"
	"github.com/cosmos/cosmos-sdk/types/tx"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

	"github.com/CosmWasm/wasmd/x/wasm/types"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
)

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

	msgInstantiateContract := &types.MsgInstantiateContract{
		Sender: senderAddress.String(),
		Admin:  senderAddress.String(),
		CodeID: 2711,
		Label:  "spot-offchain-vault example",
		Msg: types.RawContractMessage(
			fmt.Sprintf(`{"admin":"%s","vault_type":{"Spot":{"oracle_type":9,"base_oracle_symbol":"0x2d9315a88f3019f8efa88dfe9c0f0843712da0bac814461e27733f6b83eb51b3","quote_oracle_symbol":"0x1fc18861232290221461220bd4e2acd1dcdfbc89c84092c93c18bdc7756c1588","base_decimals":18,"quote_decimals":6}},"market_id":"0x0611780ba69656949525013d947713300f56c37b6175e02f26bffa495c3208fe","oracle_stale_time":3600,"notional_value_cap":"5000000000000"}`, senderAddress.String()),
		),
		Funds: cosmtypes.NewCoins(cosmtypes.NewCoin("inj", math.NewIntFromBigInt(cosmtypes.MustNewDecFromStr("10").BigInt()))),
	}

	//AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	broadcastResult, err := chainClient.SyncBroadcastMsg(msgInstantiateContract)
	if err != nil {
		panic(err)
	}
	fmt.Println("instantiate tx hash:", broadcastResult.TxResponse.TxHash)

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

	dataHex, err := hex.DecodeString(txResponse.TxResponse.Data)
	if err != nil {
		panic(err)
	}

	var txResult cosmtypes.TxMsgData
	err = clientCtx.Codec.Unmarshal(dataHex, &txResult)
	if err != nil {
		panic(err)
	}

	var instantiateResponse types.MsgInstantiateContractResponse
	err = clientCtx.Codec.Unmarshal(txResult.MsgResponses[0].Value, &instantiateResponse)
	if err != nil {
		panic(err)
	}

	fmt.Println("offchain contract address:", instantiateResponse.Address)
	fmt.Println("gas used:", txResponse.TxResponse.GasUsed, ", fee:", cosmtypes.NewDecCoinFromCoin(txResponse.Tx.AuthInfo.Fee.Amount[0]).Amount.Quo(cosmtypes.NewDec(10).Power(18)).String()+"INJ")
}
