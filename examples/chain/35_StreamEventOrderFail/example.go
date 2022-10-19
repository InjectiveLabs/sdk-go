package main

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/InjectiveLabs/sdk-go/client/common"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

func main() {
	network := common.LoadNetwork("mainnet", "sentry0")
	tmRPC, err := rpchttp.New(network.TmEndpoint, "/websocket")
	if err != nil {
		panic(err)
	}

	defer tmRPC.WSEvents.Stop()
	tmRPC.WSEvents.Start()

	eventFilter := "tm.event='Tx' AND message.sender='inj1rwv4zn3jptsqs7l8lpa3uvzhs57y8duemete9e' AND message.action='/injective.exchange.v1beta1.MsgBatchUpdateOrders' AND injective.exchange.v1beta1.EventOrderFail.flags EXISTS"
	eventCh, err := tmRPC.WSEvents.Subscribe(context.Background(), "OrderFail", eventFilter, 100)
	if err != nil {
		panic(err)
	}

	for {
		e := <-eventCh

		var failedOrderHashes []string
		err = json.Unmarshal([]byte(e.Events["injective.exchange.v1beta1.EventOrderFail.hashes"][0]), &failedOrderHashes)
		if err != nil {
			panic(err)
		}

		var failedOrderCodes []uint
		err = json.Unmarshal([]byte(e.Events["injective.exchange.v1beta1.EventOrderFail.flags"][0]), &failedOrderCodes)
		if err != nil {
			panic(err)
		}

		results := map[string]uint{}
		for i, hash := range failedOrderHashes {
			orderHashBytes, _ := base64.StdEncoding.DecodeString(hash)
			orderHash := "0x" + hex.EncodeToString(orderHashBytes)
			results[orderHash] = failedOrderCodes[i]
		}

		fmt.Println(results)
	}

}
