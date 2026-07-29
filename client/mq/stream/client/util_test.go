package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	abcitypes "github.com/cometbft/cometbft/api/cometbft/abci/v1"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	mqtypes "github.com/InjectiveLabs/sdk-go/chain/mq/types"
)

func TestDecodePublishEventRendersAnyAsJSON(t *testing.T) {
	t.Parallel()

	event := &exchangetypes.EventSubaccountWithdraw{
		Amount: sdk.NewCoin("inj", sdkmath.NewInt(123123)),
	}
	eventBz := packPublishEventForClientTest(t, event)

	decoded, err := decodePublishEvent(eventBz)

	require.NoError(t, err)
	require.Equal(t, codectypes.MsgTypeURL(event), decoded.TypeURL)

	var value map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(decoded.Value, &value))
	require.JSONEq(t, `{"denom":"inj","amount":"123123"}`, string(value["amount"]))
}

func TestShortHash(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name:     "empty",
			expected: "-",
		},
		{
			name:     "short",
			input:    []byte{0x01, 0x02, 0x0f},
			expected: "01020f",
		},
		{
			name:     "truncated",
			input:    []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09},
			expected: "0102030405060708",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, shortHash(tc.input))
		})
	}
}

func TestEncodeBytesIncludesUTF8OnlyForValidData(t *testing.T) {
	t.Parallel()

	require.Equal(t, encodedBytesFile{
		Base64: "aGVsbG8=",
		UTF8:   "hello",
	}, encodeBytes([]byte("hello")))
	require.Equal(t, encodedBytesFile{
		Base64: "//4=",
	}, encodeBytes([]byte{0xff, 0xfe}))
}

func TestTransformEventSet(t *testing.T) {
	t.Parallel()

	event := &exchangetypes.EventSubaccountWithdraw{
		Amount: sdk.NewCoin("inj", sdkmath.NewInt(123123)),
	}
	abciEvent := abcitypes.Event{
		Type: "block.event",
		Attributes: []abcitypes.EventAttribute{
			{Key: "scope", Value: "block", Index: true},
		},
	}

	transformed, err := transformEventSet(mqtypes.EventSet{
		PublishedEvents: [][]byte{packPublishEventForClientTest(t, event)},
		TrueOrders: []mqtypes.EventType{
			mqtypes.EventType_ABCI,
			mqtypes.EventType_PUBLISH,
		},
		AbciEvents: []abcitypes.Event{abciEvent},
	})

	require.NoError(t, err)
	require.Equal(t, []string{"ABCI", "PUBLISH"}, transformed.TrueOrders)
	require.Equal(t, []abcitypes.Event{abciEvent}, transformed.ABCIEvents)
	require.Len(t, transformed.PublishEvents, 1)
	require.Equal(t, codectypes.MsgTypeURL(event), transformed.PublishEvents[0].TypeURL)

	var value map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(transformed.PublishEvents[0].Value, &value))
	require.JSONEq(t, `{"denom":"inj","amount":"123123"}`, string(value["amount"]))
}

func TestTransformEventSetReturnsPublishDecodeError(t *testing.T) {
	t.Parallel()

	_, err := transformEventSet(mqtypes.EventSet{
		PublishedEvents: [][]byte{{0xff}},
	})

	require.ErrorContains(t, err, "decode publish event 0")
}

func TestTransformAdditionalData(t *testing.T) {
	t.Parallel()

	require.Equal(t, []additionalDataFile{
		{
			Type:   "EVM",
			Base64: "ZXZtLWRhdGE=",
			UTF8:   "evm-data",
		},
		{
			Type:   "EVM",
			Base64: "//4=",
		},
	}, transformAdditionalData([]*mqtypes.AdditionalDataEntry{
		{
			Type: mqtypes.AdditionalDataTypes_EVM,
			Data: []byte("evm-data"),
		},
		{
			Type: mqtypes.AdditionalDataTypes_EVM,
			Data: []byte{0xff, 0xfe},
		},
	}))
}

func TestDecodePublishEventRejectsRawBytes(t *testing.T) {
	t.Parallel()

	event := &exchangetypes.EventSubaccountWithdraw{
		Amount: sdk.NewCoin("inj", sdkmath.NewInt(123123)),
	}
	eventBz, err := proto.Marshal(event)
	require.NoError(t, err)

	_, err = decodePublishEvent(eventBz)

	require.ErrorContains(t, err, "missing type_url")
}

func TestDecodePublishEventRejectsUnknownTypeURL(t *testing.T) {
	t.Parallel()

	eventBz, err := proto.Marshal(&codectypes.Any{
		TypeUrl: "/injective.unknown.Event",
		Value:   []byte{0x08, 0x01},
	})
	require.NoError(t, err)

	_, err = decodePublishEvent(eventBz)

	require.ErrorContains(t, err, `unknown publish event type "/injective.unknown.Event"`)
}

func TestDecodePublishEventRejectsMalformedAnyValue(t *testing.T) {
	t.Parallel()

	event := &exchangetypes.EventSubaccountWithdraw{}
	eventBz, err := proto.Marshal(&codectypes.Any{
		TypeUrl: codectypes.MsgTypeURL(event),
		Value:   []byte{0xff},
	})
	require.NoError(t, err)

	_, err = decodePublishEvent(eventBz)

	require.ErrorContains(t, err, "unmarshal publish event")
	require.ErrorContains(t, err, codectypes.MsgTypeURL(event))
}

func TestMessageNameFromTypeURL(t *testing.T) {
	t.Parallel()

	require.Equal(
		t,
		"injective.exchange.v1beta1.EventSubaccountWithdraw",
		messageNameFromTypeURL("/injective.exchange.v1beta1.EventSubaccountWithdraw"),
	)
	require.Equal(
		t,
		"injective.exchange.v1beta1.EventSubaccountWithdraw",
		messageNameFromTypeURL("type.googleapis.com/injective.exchange.v1beta1.EventSubaccountWithdraw"),
	)
}

func TestWriteEventsFileCreatesDirectoryWithRestrictedPermissions(t *testing.T) {
	t.Parallel()

	eventsDir := filepath.Join(t.TempDir(), "events")

	err := writeEventsFile(eventsDir, &mqtypes.EventStreamResponse{BlockHeight: 7})
	require.NoError(t, err)

	info, err := os.Stat(eventsDir)
	require.NoError(t, err)
	require.True(t, info.IsDir())
	require.Zero(t, info.Mode().Perm()&^os.FileMode(0o750))

	fileInfo, err := os.Stat(filepath.Join(eventsDir, "block-000000000007.json"))
	require.NoError(t, err)
	require.Zero(t, fileInfo.Mode().Perm()&^os.FileMode(0o600))
}

func TestWriteEventsFileSkipsWhenDirectoryUnset(t *testing.T) {
	t.Parallel()

	err := writeEventsFile("", &mqtypes.EventStreamResponse{
		BlockEvents: mqtypes.EventSet{
			PublishedEvents: [][]byte{{0xff}},
		},
	})

	require.NoError(t, err)
}

func TestResolveConsumerIDUsesFlagValue(t *testing.T) {
	t.Parallel()

	filePath := filepath.Join(t.TempDir(), "consumer-id")

	consumerID, err := resolveConsumerID(" from-flag ", filePath)

	require.NoError(t, err)
	require.Equal(t, "from-flag", consumerID)
	require.NoFileExists(t, filePath)
}

func TestResolveConsumerIDReadsExistingFile(t *testing.T) {
	t.Parallel()

	filePath := filepath.Join(t.TempDir(), "consumer-id")
	require.NoError(t, os.WriteFile(filePath, []byte("from-file\n"), 0o600))

	consumerID, err := resolveConsumerID("", filePath)

	require.NoError(t, err)
	require.Equal(t, "from-file", consumerID)
}

func TestResolveConsumerIDGeneratesAndWritesMissingFile(t *testing.T) {
	t.Parallel()

	filePath := filepath.Join(t.TempDir(), "nested", "consumer-id")

	consumerID, err := resolveConsumerID("", filePath)

	require.NoError(t, err)
	require.NotEmpty(t, consumerID)
	_, err = uuid.Parse(consumerID)
	require.NoError(t, err)

	bz, err := os.ReadFile(filePath)
	require.NoError(t, err)
	require.Equal(t, consumerID+"\n", string(bz))
}

func TestWriteEventsFileWritesStructuredPayload(t *testing.T) {
	t.Parallel()

	event := &exchangetypes.EventSubaccountWithdraw{
		Amount: sdk.NewCoin("inj", sdkmath.NewInt(123123)),
	}
	blockABCIEvent := abcitypes.Event{
		Type: "block.event",
		Attributes: []abcitypes.EventAttribute{
			{Key: "scope", Value: "block", Index: true},
		},
	}
	eventsDir := t.TempDir()

	err := writeEventsFile(eventsDir, &mqtypes.EventStreamResponse{
		BlockHeight: 11,
		AppHash:     []byte{0x01, 0x02, 0x03},
		LastAppHash: []byte{0x04, 0x05, 0x06},
		BlockEvents: mqtypes.EventSet{
			TrueOrders: []mqtypes.EventType{mqtypes.EventType_ABCI},
			AbciEvents: []abcitypes.Event{blockABCIEvent},
		},
		TxEvents: []mqtypes.EventSet{
			{
				PublishedEvents: [][]byte{packPublishEventForClientTest(t, event)},
				TrueOrders:      []mqtypes.EventType{mqtypes.EventType_PUBLISH},
			},
		},
		AdditionalDataEntries: []*mqtypes.AdditionalDataEntry{
			{
				Type: mqtypes.AdditionalDataTypes_EVM,
				Data: []byte("extra"),
			},
		},
	})

	require.NoError(t, err)

	bz, err := os.ReadFile(filepath.Join(eventsDir, "block-000000000011.json"))
	require.NoError(t, err)

	var written blockEventsFile
	require.NoError(t, json.Unmarshal(bz, &written))
	require.Equal(t, int64(11), written.BlockHeight)
	require.Equal(t, "010203", written.AppHash)
	require.Equal(t, "040506", written.LastAppHash)
	require.Equal(t, []string{"ABCI"}, written.BlockEvents.TrueOrders)
	require.Equal(t, []abcitypes.Event{blockABCIEvent}, written.BlockEvents.ABCIEvents)
	require.Len(t, written.TxEvents, 1)
	require.Equal(t, []string{"PUBLISH"}, written.TxEvents[0].TrueOrders)
	require.Len(t, written.TxEvents[0].PublishEvents, 1)
	require.Equal(t, codectypes.MsgTypeURL(event), written.TxEvents[0].PublishEvents[0].TypeURL)
	require.Equal(t, []additionalDataFile{
		{
			Type:   "EVM",
			Base64: "ZXh0cmE=",
			UTF8:   "extra",
		},
	}, written.AdditionalDataEntries)
	_, err = time.Parse(time.RFC3339, written.CapturedAt)
	require.NoError(t, err)
}

func TestWriteEventsFileReturnsTransformErrors(t *testing.T) {
	t.Parallel()

	eventsDir := t.TempDir()

	err := writeEventsFile(eventsDir, &mqtypes.EventStreamResponse{
		BlockEvents: mqtypes.EventSet{
			PublishedEvents: [][]byte{{0xff}},
		},
	})
	require.ErrorContains(t, err, "transform block event set")
	require.ErrorContains(t, err, "decode publish event 0")

	err = writeEventsFile(eventsDir, &mqtypes.EventStreamResponse{
		TxEvents: []mqtypes.EventSet{
			{
				PublishedEvents: [][]byte{{0xff}},
			},
		},
	})
	require.ErrorContains(t, err, "transform tx event set 0")
	require.ErrorContains(t, err, "decode publish event 0")
}

func TestIsExpectedStreamClose(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "context canceled",
			err:      context.Canceled,
			expected: true,
		},
		{
			name:     "eof",
			err:      io.EOF,
			expected: true,
		},
		{
			name:     "grpc canceled",
			err:      status.Error(codes.Canceled, "stream canceled"),
			expected: true,
		},
		{
			name:     "grpc unavailable eof",
			err:      status.Error(codes.Unavailable, "error reading from server: EOF"),
			expected: true,
		},
		{
			name:     "grpc unavailable transport closing",
			err:      status.Error(codes.Unavailable, "transport is closing"),
			expected: true,
		},
		{
			name:     "grpc unavailable client closing",
			err:      status.Error(codes.Unavailable, "client connection is closing"),
			expected: true,
		},
		{
			name: "non status error",
			err:  errors.New("stream failed"),
		},
		{
			name: "grpc unavailable other",
			err:  status.Error(codes.Unavailable, "connection refused"),
		},
		{
			name: "grpc internal eof",
			err:  status.Error(codes.Internal, "EOF"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, isExpectedStreamClose(tc.err))
		})
	}
}

func packPublishEventForClientTest(t *testing.T, event proto.Message) []byte {
	t.Helper()

	eventAny, err := codectypes.NewAnyWithValue(event)
	require.NoError(t, err)
	eventBz, err := proto.Marshal(eventAny)
	require.NoError(t, err)

	return eventBz
}
