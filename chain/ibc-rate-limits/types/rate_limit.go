package types

import (
	"encoding/json"
	"strings"

	"cosmossdk.io/errors"
	oracletypes "github.com/InjectiveLabs/sdk-go/chain/oracle/types"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	"github.com/ethereum/go-ethereum/common"
)

const (
	ModuleName = "rate-limits"

	StoreKey = ModuleName

	// RouterKey is the message route for IBC fee module
	RouterKey = ModuleName

	// QuerierRoute is the querier route for IBC fee module
	QuerierRoute = ModuleName
)

// NewTokenID returns a unique hash representation of either a native token or an ibc token
func NewTokenID(denomTrace string) []byte {
	dt := transfertypes.ParseDenomTrace(denomTrace)
	ibcDenom := dt.IBCDenom() // this is either a regular string (native) or an ibc/hash format
	if dt.IsNativeDenom() {
		return common.BytesToHash([]byte(ibcDenom)).Bytes() // just hash the string to something unique
	}

	// return the original ibc hash
	ibcDenom = strings.TrimPrefix(ibcDenom, "ibc/")
	return common.HexToHash(ibcDenom).Bytes()
}

func ParseFungibleTokenData(packet channeltypes.Packet) (*transfertypes.FungibleTokenPacketData, error) {
	var ftpd transfertypes.FungibleTokenPacketData
	if err := json.Unmarshal(packet.GetData(), &ftpd); err != nil {
		return nil, err
	}

	return &ftpd, nil
}

// IsAckError checks an IBC acknowledgement to see if it's an error.
// This is a replacement for ack.Success() which is currently not working on some circumstances
func IsAckError(acknowledgement []byte) bool {
	var ackErr channeltypes.Acknowledgement_Error
	if err := json.Unmarshal(acknowledgement, &ackErr); err == nil && ackErr.Error != "" {
		return true
	}

	return false
}

func DefaultGenesis() *GenesisState {
	return &GenesisState{RateLimits: []*RateLimit{}}
}

func (g *GenesisState) ValidateBasic() error {
	for _, limit := range g.RateLimits {
		if err := limit.ValidateBasic(); err != nil {
			return err
		}
	}

	return nil
}

func (rl *RateLimit) TokenID() []byte {
	return NewTokenID(rl.Token)
}

func (rl *RateLimit) ValidateBasic() error {
	if rl.Token == "" {
		return errors.Wrap(ErrInvalidRateLimit, "token cannot be empty")
	}

	if rl.OracleType == oracletypes.OracleType_Unspecified {
		return errors.Wrap(ErrInvalidRateLimit, "oracle_type cannot be unspecified")
	}

	if rl.OracleId == "" {
		return errors.Wrap(ErrInvalidRateLimit, "oracle_id cannot be empty")
	}

	if !rl.MaxOutgoingLimitUsd.IsPositive() {
		return errors.Wrap(ErrInvalidRateLimit, "max_outgoing_limit usd must be positive")
	}

	if rl.SlidingWindowSize == 0 {
		return errors.Wrap(ErrInvalidRateLimit, "sliding_window_size cannot be empty")
	}

	if rl.TokenDecimals == 0 {
		return errors.Wrap(ErrInvalidRateLimit, "token_decimals cannot be 0")
	}

	return nil
}
