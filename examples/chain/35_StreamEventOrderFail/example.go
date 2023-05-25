package main

import (
	"fmt"

	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
)

func main() {
	network := common.LoadNetwork("mainnet", "sentry0")
	tmClient, err := rpchttp.New(network.TmEndpoint, "/websocket")
	if err != nil {
		panic(err)
	}

	clientCtx, err := chainclient.NewClientContext(
		network.ChainId,
		"",
		nil,
	)
	if err != nil {
		panic(err)
	}
	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint).WithClient(tmClient)

	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network.ChainGrpcEndpoint,
		common.OptionTLSCert(network.ChainTlsCert),
		common.OptionGasPrices("500000000inj"),
	)
	if err != nil {
		panic(err)
	}

	failEventCh := make(chan map[string]uint, 10000)
	go chainClient.StreamEventOrderFail("inj1rwv4zn3jptsqs7l8lpa3uvzhs57y8duemete9e", failEventCh)
	for {
		e := <-failEventCh
		fmt.Println(e)
	}
}
