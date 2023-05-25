package main

import (
	"fmt"

	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
)

func main() {
	network := common.LoadNetwork("devnet", "")
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

	//0x74b17b0d6855feba39f1f7ab1e8bad0363bd510ee1dcc74e40c2adfe1502f781
	//0x74ee114ad750f8429a97e07b5e73e145724e9b21670a7666625ddacc03d6758d
	//0x26413a70c9b78a495023e5ab8003c9cf963ef963f6755f8b57255feb5744bf31
	marketIds := []string{
		"0xa508cb32923323679f29a032c70342c147c17d0145625922b0ef22e955c844c0",
	}

	orderbookCh := make(chan exchangetypes.Orderbook, 10000)
	go chainClient.StreamOrderbookUpdateEvents(chainclient.SpotOrderbook, marketIds, orderbookCh)
	for {
		ob := <-orderbookCh
		fmt.Println(ob)
	}
}
