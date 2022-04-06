package main

func main() {
	// network := common.LoadNetwork("mainnet", "k8s")
	//network := common.LoadNetwork("testnet", "k8s")
	//tmRPC, err := rpchttp.New(network.TmEndpoint, "/websocket")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//senderAddress, cosmosKeyring, err := chainclient.InitCosmosKeyring(
	//	os.Getenv("HOME")+"/.injectived",
	//	"injectived",
	//	"file",
	//	"inj-user",
	//	"12345678",
	//	"f9db9bf330e23cb7839039e944adef6e9df447b90b503d5b4464c90bea9022f3", // keyring will be used if pk not provided
	//	false,
	//)
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//clientCtx, err := chainclient.NewClientContext(
	//	network.ChainId,
	//	senderAddress.String(),
	//	cosmosKeyring,
	//)
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//clientCtx.WithNodeURI(network.TmEndpoint)
	//clientCtx = clientCtx.WithClient(tmRPC)
	//
	//chainClient, err := chainclient.NewChainClient(
	//	clientCtx,
	//	network.ChainGrpcEndpoint,
	//	common.OptionTLSCert(network.ChainTlsCert),
	//	common.OptionGasPrices("500000000inj"),
	//)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//// note that we use grantee keyring to send the msg on behalf of granter here
	//// sender, subaccount are from granter
	//granter := "inj14au322k9munkmx5wrchz9q30juf5wjgz2cfqku"
	//grantee := "inj1hkhdaj2a2clmq5jq6mspsggqs32vynpk228q3r"
	//granterAcc, _ := sdk.AccAddressFromBech32(granter)
	//defaultSubaccountID := chainClient.DefaultSubaccount(granterAcc)
	//
	//marketId := "0x0511ddc4e6586f3bfe1acb2dd905f8b8a82c97e1edaef654b12ca7e6031ca0fa"
	//amount := decimal.NewFromFloat(2)
	//price := cosmtypes.MustNewDecFromStr("22")
	//quantity := cosmtypes.MustNewDecFromStr("10000")
	//order := chainClient.SpotOrder(defaultSubaccountID, &chainclient.SpotOrderData{
	//	OrderType:    exchangetypes.OrderType_BUY,
	//	Quantity:     quantity,
	//	Price:        price,
	//	FeeRecipient: senderAddress.String(),
	//	MarketId:     marketId,
	//})
	//
	//// manually pack msg into Any type
	//msg0 := exchangetypes.MsgCreateSpotLimitOrder{
	//	Sender: granter,
	//	Order:  *order,
	//}
	//msg0Bytes, _ := msg0.Marshal()
	//msg0Any := &codectypes.Any{}
	//msg0Any.TypeUrl = sdk.MsgTypeURL(&msg0)
	//msg0Any.Value = msg0Bytes
	//
	//msg := &authztypes.MsgExec{
	//	Grantee: grantee,
	//	Msgs:    []*codectypes.Any{msg0Any},
	//}
	//
	//err = chainClient.QueueBroadcastMsg(msg)
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//time.Sleep(time.Second * 5)
}
