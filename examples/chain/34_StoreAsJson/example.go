package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/shopspring/decimal"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

type Msg struct {
	Type       string          `json:"@type"`
	RawMessage json.RawMessage `json:"rawMessage"`
	AccSeq     uint64          `json:"accSeq"`
	AccNum     uint64          `json:"accNum"`
	Msg        sdk.Msg         `json:"-"`
}

func GetMessageFromJson(msg []byte) (*Msg, error) {
	var message Msg
	err := json.Unmarshal(msg, &message)
	if err != nil {
		return nil, err
	}

	switch message.Type {
	// TODO: Support more types here
	case "MsgCreateSpotLimitOrder":
		message.Msg = &exchangetypes.MsgCreateSpotLimitOrder{}
	}

	if message.Msg == nil {
		return nil, errors.New("unsupported type")
	}

	err = json.Unmarshal(message.RawMessage, message.Msg)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

func BuildMsg(msgType string, msg sdk.Msg, accSeq, accNum uint64) (*Msg, error) {
	sdkMsg, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	message := &Msg{
		Type:       msgType,
		RawMessage: sdkMsg,
		AccSeq:     accSeq,
		AccNum:     accNum,
	}

	return message, nil
}

func LoadFromFile(path string) (*Msg, error) {
	r, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("read file err: %w", err)
	}

	return GetMessageFromJson(b)
}

func SaveToFile(path string, message *Msg) error {
	b, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, b, 0755)
}

func main() {
	network := common.LoadNetwork("testnet", "k8s")
	tmRPC, err := rpchttp.New(network.TmEndpoint, "/websocket")
	if err != nil {
		panic(err)
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

	defaultSubaccountID := chainClient.DefaultSubaccount(senderAddress)
	marketId := "0x0611780ba69656949525013d947713300f56c37b6175e02f26bffa495c3208fe"

	amount := decimal.NewFromFloat(0.001)
	price := decimal.NewFromFloat(500000000.55)

	order := chainClient.SpotOrder(defaultSubaccountID, network, &chainclient.SpotOrderData{
		OrderType:    exchangetypes.OrderType_BUY, //BUY SELL BUY_PO SELL_PO
		Quantity:     amount,
		Price:        price,
		FeeRecipient: senderAddress.String(),
		MarketId:     marketId,
	})

	msg := new(exchangetypes.MsgCreateSpotLimitOrder)
	msg.Sender = senderAddress.String()
	msg.Order = exchangetypes.SpotOrder(*order)

	offlineMsg, err := BuildMsg("MsgCreateSpotLimitOrder", msg, 104, 24)
	if err != nil {
		panic(err)
	}

	// save to file
	if err := SaveToFile("msg.json", offlineMsg); err != nil {
		panic(err)
	}

	// load from the same file
	theMessage, err := LoadFromFile("msg.json")
	if err != nil {
		panic(err)
	}

	fmt.Println("price:", theMessage.Msg.(*exchangetypes.MsgCreateSpotLimitOrder).Order.GetPrice())
	fmt.Println("sender addr", senderAddress.String())
	chainClient.SetAccSeq(theMessage.AccSeq)
	chainClient.SetAccNum(theMessage.AccNum)
	txResp, err := chainClient.SyncBroadcastMsg(theMessage.Msg)
	if err != nil {
		panic(err)
	}
	fmt.Println("result:", txResp)
}
