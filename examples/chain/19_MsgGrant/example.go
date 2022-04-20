package main

import (
	"fmt"
	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	"github.com/gogo/protobuf/proto"
	"os"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"

	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
)

func buildGenericAuthz(granter string, grantee string) *authztypes.MsgGrant {
	authz := authztypes.NewGenericAuthorization("/injective.exchange.v1beta1.MsgCreateSpotLimitOrder")
	authzAny := codectypes.UnsafePackAny(authz)
	expireIn := time.Now().AddDate(1, 0, 0)
	return &authztypes.MsgGrant{
		Granter: granter,
		Grantee: grantee,
		Grant: authztypes.Grant{
			Authorization: authzAny,
			Expiration:    expireIn,
		},
	}
}

func buildTypedAuthz(granter string, grantee string, subaccountId string, markets []string) *authztypes.MsgGrant {
	typedAuthz := exchangetypes.CreateSpotLimitOrderAuthz{
		SubaccountId: subaccountId,
		MarketIds: markets,
	}
	typedAuthzBytes, _ := typedAuthz.Marshal()
	typedAuthzAny := &codectypes.Any{
		TypeUrl: "/" + proto.MessageName(&typedAuthz),
		Value: typedAuthzBytes,
	}

	expireIn := time.Now().AddDate(1, 0, 0)
	return &authztypes.MsgGrant{
		Granter: granter,
		Grantee: grantee,
		Grant: authztypes.Grant{
			Authorization: typedAuthzAny,
			Expiration:    expireIn,
		},
	}
}

func main() {
	// network := common.LoadNetwork("mainnet", "k8s")
	network := common.LoadNetwork("testnet", "k8s")
	tmRPC, err := rpchttp.New(network.TmEndpoint, "/websocket")
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
	}
	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint).WithClient(tmRPC)

	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network.ChainGrpcEndpoint,
		common.OptionTLSCert(network.ChainTlsCert),
		common.OptionGasPrices("500000000inj"),
	)
	if err != nil {
		fmt.Println(err)
	}

	granter := senderAddress.String()
	grantee := "inj1hkhdaj2a2clmq5jq6mspsggqs32vynpk228q3r"

	//msg := buildGenericAuthz(granter, grantee)
	msg := buildTypedAuthz(
		granter,
		grantee,
		chainClient.DefaultSubaccount(senderAddress).String(),
		[]string{"0xe0dc13205fb8b23111d8555a6402681965223135d368eeeb964681f9ff12eb2a"},
	)

	err = chainClient.QueueBroadcastMsg(msg)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second * 5)
}
