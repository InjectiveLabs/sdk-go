package main

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
	"unicode/utf8"

	abcitypes "github.com/cometbft/cometbft/api/cometbft/abci/v1"
	sdkcodec "github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	mqtypes "github.com/InjectiveLabs/sdk-go/chain/mq/types"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
)

type blockEventsFile struct {
	CapturedAt            string               `json:"captured_at"`
	BlockHeight           int64                `json:"block_height"`
	AppHash               string               `json:"app_hash"`
	LastAppHash           string               `json:"last_app_hash"`
	BlockEvents           eventSetEventsFile   `json:"block_events"`
	TxEvents              []eventSetEventsFile `json:"tx_events"`
	AdditionalDataEntries []additionalDataFile `json:"additional_data_entries,omitempty"`
}

type eventSetEventsFile struct {
	TrueOrders    []string           `json:"true_orders"`
	ABCIEvents    []abcitypes.Event  `json:"abci_events"`
	PublishEvents []publishEventFile `json:"publish_events"`
}

type publishEventFile struct {
	TypeURL string          `json:"type_url"`
	Value   json.RawMessage `json:"value"`
}

type encodedBytesFile struct {
	Base64 string `json:"base64"`
	UTF8   string `json:"utf8,omitempty"`
}

type additionalDataFile struct {
	Type   string `json:"type"`
	Base64 string `json:"base64"`
	UTF8   string `json:"utf8,omitempty"`
}

type publishEventDecoder struct {
	cdc *sdkcodec.ProtoCodec
}

func newPublishEventDecoder() (*publishEventDecoder, error) {
	encodingConfig := chainclient.NewEncodingConfig()
	cdc, ok := encodingConfig.Marshaler.(*sdkcodec.ProtoCodec)
	if !ok {
		return nil, errors.New("encoding config marshaler is not a proto codec")
	}

	return &publishEventDecoder{cdc: cdc}, nil
}

func shortHash(b []byte) string {
	if len(b) == 0 {
		return "-"
	}
	if len(b) > 8 {
		b = b[:8]
	}
	return hex.EncodeToString(b)
}

func encodeBytes(b []byte) encodedBytesFile {
	encoded := encodedBytesFile{Base64: base64.StdEncoding.EncodeToString(b)}
	if utf8.Valid(b) {
		encoded.UTF8 = string(b)
	}
	return encoded
}

func transformEventSet(events mqtypes.EventSet, decoder *publishEventDecoder) eventSetEventsFile {
	trueOrders := make([]string, 0, len(events.TrueOrders))
	for _, order := range events.TrueOrders {
		trueOrders = append(trueOrders, order.String())
	}

	publishEvents := make([]publishEventFile, 0, len(events.PublishedEvents))
	for idx, event := range events.PublishedEvents {
		publishEvent, err := decoder.decodePublishEvent(event)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error decoding publish event %d: %v\n", idx, err)
			continue
		}
		publishEvents = append(publishEvents, publishEvent)
	}

	return eventSetEventsFile{
		TrueOrders:    trueOrders,
		ABCIEvents:    events.AbciEvents,
		PublishEvents: publishEvents,
	}
}

func transformAdditionalData(entries []*mqtypes.AdditionalDataEntry) []additionalDataFile {
	result := make([]additionalDataFile, 0, len(entries))
	for _, entry := range entries {
		data := encodeBytes(entry.Data)
		result = append(result, additionalDataFile{
			Type:   entry.Type.String(),
			Base64: data.Base64,
			UTF8:   data.UTF8,
		})
	}

	return result
}

func writeEventsFile(eventsDir string, res *mqtypes.EventStreamResponse, decoder *publishEventDecoder) error {
	if eventsDir == "" {
		return nil
	}

	txEvents := make([]eventSetEventsFile, 0, len(res.TxEvents))
	for _, events := range res.TxEvents {
		txEventSet := transformEventSet(events, decoder)
		txEvents = append(txEvents, txEventSet)
	}

	blockEvents := transformEventSet(res.BlockEvents, decoder)

	out := blockEventsFile{
		CapturedAt:            time.Now().Format(time.RFC3339),
		BlockHeight:           res.BlockHeight,
		AppHash:               hex.EncodeToString(res.AppHash),
		LastAppHash:           hex.EncodeToString(res.LastAppHash),
		BlockEvents:           blockEvents,
		TxEvents:              txEvents,
		AdditionalDataEntries: transformAdditionalData(res.AdditionalDataEntries),
	}

	if err := os.MkdirAll(eventsDir, 0o750); err != nil {
		return fmt.Errorf("create events dir: %w", err)
	}

	bz, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal events file: %w", err)
	}

	filename := filepath.Join(eventsDir, fmt.Sprintf("block-%012d.json", res.BlockHeight))
	if err := os.WriteFile(filename, append(bz, '\n'), 0o600); err != nil {
		return fmt.Errorf("write events file: %w", err)
	}

	return nil
}

func printMinimal(res *mqtypes.EventStreamResponse) {
	_, _ = fmt.Printf(
		"%s height=%d tx_events=%d abci_events=%d publish_events=%d app_hash=%s last_app_hash=%s\n",
		time.Now().Format(time.RFC3339),
		res.BlockHeight,
		len(res.TxEvents),
		len(res.BlockEvents.AbciEvents),
		len(res.BlockEvents.PublishedEvents),
		shortHash(res.AppHash),
		shortHash(res.LastAppHash),
	)
}

func printVerbose(res *mqtypes.EventStreamResponse, decoder *publishEventDecoder) {
	_, _ = fmt.Println("---------------------")
	_, _ = fmt.Println("height : ", res.BlockHeight)
	_, _ = fmt.Println("app hash : ", hex.EncodeToString(res.AppHash))
	_, _ = fmt.Println("last app hash : ", hex.EncodeToString(res.LastAppHash))

	_, _ = fmt.Println("block publish events : ")
	for _, e := range res.BlockEvents.PublishedEvents {
		printPublishEvent(e, decoder)
	}

	_, _ = fmt.Println("block events true order : ", res.BlockEvents.TrueOrders)

	for i, tx := range res.TxEvents {
		_, _ = fmt.Println("publish events for tx index : ", i)
		for _, e := range tx.PublishedEvents {
			printPublishEvent(e, decoder)
		}
		_, _ = fmt.Println("tx events true order : ", tx.TrueOrders)
	}
}

func printPublishEvent(event []byte, decoder *publishEventDecoder) {
	publishEvent, err := decoder.decodePublishEvent(event)
	if err != nil {
		_, _ = fmt.Printf("error decoding publish event: %v\n", err)
		return
	}

	bz, err := json.MarshalIndent(publishEvent, "", "  ")
	if err != nil {
		_, _ = fmt.Printf("error formatting publish event: %v\n", err)
		return
	}
	_, _ = fmt.Println(string(bz))
}

func (d *publishEventDecoder) decodePublishEvent(event []byte) (publishEventFile, error) {
	if d == nil || d.cdc == nil {
		return publishEventFile{}, errors.New("publish event decoder is not initialized")
	}

	var anyEvent codectypes.Any
	if err := d.cdc.Unmarshal(event, &anyEvent); err != nil {
		return publishEventFile{}, fmt.Errorf("decode publish event any: %w", err)
	}

	if anyEvent.TypeUrl == "" {
		return publishEventFile{}, errors.New("decode publish event any: missing type_url")
	}

	anyJSON, err := sdkcodec.ProtoMarshalJSON(&anyEvent, nil)
	if err != nil {
		return publishEventFile{}, fmt.Errorf("marshal publish event any json %q: %w", anyEvent.TypeUrl, err)
	}

	var value map[string]json.RawMessage
	if err := json.Unmarshal(anyJSON, &value); err != nil {
		return publishEventFile{}, fmt.Errorf("decode publish event json %q: %w", anyEvent.TypeUrl, err)
	}
	delete(value, "@type")

	valueJSON, err := json.Marshal(value)
	if err != nil {
		return publishEventFile{}, fmt.Errorf("marshal publish event json %q: %w", anyEvent.TypeUrl, err)
	}

	return publishEventFile{
		TypeURL: anyEvent.TypeUrl,
		Value:   json.RawMessage(valueJSON),
	}, nil
}
