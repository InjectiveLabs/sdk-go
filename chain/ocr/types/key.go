package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is the name of the module
	ModuleName = "chainlink"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// TStoreKey to be used when creating the Transient KVStore
	TStoreKey = "transient_chainlink"

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querier msgs
	QuerierRoute = ModuleName
)

var (
	FeedConfigPrefix           = []byte{0x01}
	FeedConfigInfoPrefix       = []byte{0x02}
	LatestEpochAndRoundPrefix  = []byte{0x03}
	TransmissionPrefix         = []byte{0x04}
	AggregatorRoundIDPrefix    = []byte{0x05}
	FeedPoolPrefix             = []byte{0x06}
	ObservationsCountPrefix    = []byte{0x07}
	TransmissionsCountPrefix   = []byte{0x08}
	PayeePrefix                = []byte{0x09}
	PendingPayeeTransferPrefix = []byte{0x10}
)

func GetFeedConfigKey(feedId string) []byte {
	feedIdBz := getPaddedFeedIdBz(feedId)
	buf := make([]byte, 0, len(FeedConfigPrefix)+len(feedIdBz))
	buf = append(buf, FeedConfigPrefix...)
	buf = append(buf, feedIdBz...)
	return buf
}

func GetFeedConfigInfoKey(feedId string) []byte {
	feedIdBz := getPaddedFeedIdBz(feedId)

	buf := make([]byte, 0, len(FeedConfigInfoPrefix)+len(feedIdBz))
	buf = append(buf, FeedConfigInfoPrefix...)
	buf = append(buf, feedIdBz...)

	return buf
}

func GetLatestEpochAndRoundKey(feedId string) []byte {
	feedIdBz := getPaddedFeedIdBz(feedId)

	buf := make([]byte, 0, len(LatestEpochAndRoundPrefix)+len(feedIdBz))
	buf = append(buf, LatestEpochAndRoundPrefix...)
	buf = append(buf, feedIdBz...)

	return buf
}

func GetTransmissionKey(feedId string) []byte {
	feedIdBz := getPaddedFeedIdBz(feedId)

	buf := make([]byte, 0, len(TransmissionPrefix)+len(feedIdBz))
	buf = append(buf, TransmissionPrefix...)
	buf = append(buf, feedIdBz...)
	return buf
}

func GetAggregatorRoundIDKey(feedId string) []byte {
	feedIdBz := getPaddedFeedIdBz(feedId)

	buf := make([]byte, 0, len(AggregatorRoundIDPrefix)+len(feedIdBz))
	buf = append(buf, AggregatorRoundIDPrefix...)
	buf = append(buf, feedIdBz...)

	return buf
}

func GetFeedPoolKey(feedId string) []byte {
	feedIdBz := getPaddedFeedIdBz(feedId)

	buf := make([]byte, 0, len(FeedPoolPrefix)+len(feedIdBz))
	buf = append(buf, FeedPoolPrefix...)
	buf = append(buf, feedIdBz...)

	return buf
}

func GetFeedObservationsKey(feedId string, addr sdk.AccAddress) []byte {
	feedIdBz := getPaddedFeedIdBz(feedId)
	addrBz := addr.Bytes()

	buf := make([]byte, 0, len(ObservationsCountPrefix)+len(feedIdBz)+len(addrBz))
	buf = append(buf, ObservationsCountPrefix...)
	buf = append(buf, feedIdBz...)
	buf = append(buf, addrBz...)
	return buf
}

func GetFeedObservationsPrefix(feedId string) []byte {
	feedIdBz := getPaddedFeedIdBz(feedId)

	buf := make([]byte, 0, len(ObservationsCountPrefix)+len(feedIdBz))
	buf = append(buf, ObservationsCountPrefix...)
	buf = append(buf, feedIdBz...)
	return buf
}

func GetFeedTransmissionsKey(feedId string, addr sdk.AccAddress) []byte {
	feedIdBz := getPaddedFeedIdBz(feedId)
	addrBz := addr.Bytes()

	buf := make([]byte, 0, len(TransmissionsCountPrefix)+len(feedIdBz)+len(addrBz))
	buf = append(buf, TransmissionsCountPrefix...)
	buf = append(buf, feedIdBz...)
	buf = append(buf, addrBz...)
	return buf
}

func GetFeedTransmissionsPrefix(feedId string) []byte {
	feedIdBz := getPaddedFeedIdBz(feedId)

	buf := make([]byte, 0, len(TransmissionsCountPrefix)+len(feedIdBz))
	buf = append(buf, TransmissionsCountPrefix...)
	buf = append(buf, feedIdBz...)
	return buf
}

func GetPayeePrefix(feedId string, transmitter sdk.AccAddress) []byte {
	feedIdBz := getPaddedFeedIdBz(feedId)

	transmitterBz := transmitter.Bytes()
	buf := make([]byte, 0, len(PayeePrefix)+len(feedIdBz)+len(transmitterBz))
	buf = append(buf, PayeePrefix...)
	buf = append(buf, feedIdBz...)
	buf = append(buf, transmitterBz...)
	return buf
}

func GetPendingPayeeshipTransferPrefix(feedId string, transmitter sdk.AccAddress) []byte {
	feedIdBz := getPaddedFeedIdBz(feedId)

	transmitterBz := transmitter.Bytes()
	buf := make([]byte, 0, len(PendingPayeeTransferPrefix)+len(feedIdBz)+len(transmitterBz))
	buf = append(buf, PendingPayeeTransferPrefix...)
	buf = append(buf, feedIdBz...)
	buf = append(buf, transmitterBz...)
	return buf
}

func getPaddedFeedIdBz(feedId string) string {
	return fmt.Sprintf("%20s", feedId)
}

func GetFeedIdFromPaddedFeedIdBz(feedIdBz []byte) string {
	return strings.TrimSpace(string(feedIdBz))
}
