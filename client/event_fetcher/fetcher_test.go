package event_fetcher

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"testing"

	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/cosmos/gogoproto/proto"
	"github.com/ethereum/go-ethereum/common"
)

var Unmarshaler = &jsonpb.Unmarshaler{
	AllowUnknownFields: true,
}

type SpotTrade struct {
	exchangetypes.EventBatchSpotExecution
	EventIdx int         `json:"eventIdx,omitempty" bson:"eventIdx"`
	TxHash   common.Hash `json:"txHash,omitempty" bson:"txHash,omitempty"`
}

func (e *SpotTrade) New() interface{} {
	return &SpotTrade{}
}

func (e *SpotTrade) Type() string {
	return proto.MessageName(&e.EventBatchSpotExecution)
}

func (e *SpotTrade) Unmarshal(version string, data []byte) error {
	// proto msg
	if strings.Contains(e.Type(), ".") {
		switch version {
		case "1":
			return Unmarshaler.Unmarshal(bytes.NewReader(data), &e.EventBatchSpotExecution)
		case "2":
			return proto.Unmarshal(data, &e.EventBatchSpotExecution)
		default:
			return errors.New("invalid tm response version")
		}
	}
	// json msg
	return json.Unmarshal(data, e)
}

func TestFetcherComet(t *testing.T) {
	tradeEvent := &SpotTrade{}
	client := NewCometFetcher("https://testnet.tm.injective.network:443", FilterMap{
		tradeEvent.Type(): tradeEvent,
	})

	events, err := client.Fetch(context.Background(), 18732454)
	if err != nil {
		panic(err)
	}

	for _, e := range events {
		fmt.Println("event:", e.Type())
	}
}

func TestFetcherProvider(t *testing.T) {
	tradeEvent := &SpotTrade{}
	client, err := NewEventProviderFetcher("tcp://localhost:9912", FilterMap{
		tradeEvent.Type(): tradeEvent,
	})
	if err != nil {
		panic(err)
	}

	events, err := client.Fetch(context.Background(), 135136)
	if err != nil {
		panic(err)
	}

	for _, e := range events {
		fmt.Println("event:", e.Type())
	}
}
