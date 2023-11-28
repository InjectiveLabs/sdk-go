package types

import (
	"fmt"

	bandobi "github.com/bandprotocol/bandchain-packet/obi"
	bandPacket "github.com/bandprotocol/bandchain-packet/packet"

	bandprice "github.com/InjectiveLabs/sdk-go/chain/oracle/bandchain/hooks/price"
)

func NewOracleRequestPacketData(clientID string, calldata []byte, r *BandOracleRequest) bandPacket.OracleRequestPacketData {
	return bandPacket.OracleRequestPacketData{
		ClientID:       clientID,
		OracleScriptID: uint64(r.OracleScriptId),
		Calldata:       calldata,
		AskCount:       r.AskCount,
		MinCount:       r.MinCount,
		FeeLimit:       r.FeeLimit,
		PrepareGas:     r.PrepareGas,
		ExecuteGas:     r.ExecuteGas,
	}
}

// GetCalldata gets the Band IBC request call data based on the symbols and multiplier.
func (r *BandOracleRequest) GetCalldata(legacyScheme bool) []byte {
	if legacyScheme {
		return bandobi.MustEncode(bandprice.Input{
			Symbols:    r.Symbols,
			Multiplier: BandPriceMultiplier,
		})
	}

	return bandobi.MustEncode(bandprice.SymbolInput{
		Symbols:            r.Symbols,
		MinimumSourceCount: uint8(r.MinSourceCount),
	})
}

func DecodeOracleInput(data []byte) (OracleInput, error) {
	var (
		legacyInput LegacyBandInput
		newInput    BandInput
		err         error
	)

	if err = bandobi.Decode(data, &legacyInput); err == nil {
		return legacyInput, nil
	}

	if err = bandobi.Decode(data, &newInput); err == nil {
		return newInput, nil
	}

	return nil, fmt.Errorf("failed to decode oracle input: %w", err)
}

func DecodeOracleOutput(data []byte) (OracleOutput, error) {
	var (
		legacyOutput LegacyBandOutput
		newOutput    BandOutput
		err          error
	)

	if err = bandobi.Decode(data, &legacyOutput); err == nil {
		return legacyOutput, nil
	}

	if err = bandobi.Decode(data, &newOutput); err == nil {
		return newOutput, nil
	}

	return nil, fmt.Errorf("failed to decode oracle output: %w", err)
}

// it is assumed that the id of a symbol
// within OracleInput exists within OracleOutput

type OracleInput interface {
	PriceSymbols() []string
	PriceMultiplier() uint64
}

type (
	LegacyBandInput bandprice.Input
	BandInput       bandprice.SymbolInput
)

func (in LegacyBandInput) PriceSymbols() []string {
	return in.Symbols
}

func (in LegacyBandInput) PriceMultiplier() uint64 {
	return in.Multiplier
}

func (in BandInput) PriceSymbols() []string {
	return in.Symbols
}

func (in BandInput) PriceMultiplier() uint64 {
	return BandPriceMultiplier
}

type OracleOutput interface {
	Rate(id int) uint64
	Valid(id int) bool
}

type (
	LegacyBandOutput bandprice.Output
	BandOutput       bandprice.SymbolOutput
)

func (out LegacyBandOutput) Rate(id int) uint64 {
	return out.Pxs[id]
}

func (out LegacyBandOutput) Valid(id int) bool {
	return true
}

func (out BandOutput) Rate(id int) uint64 {
	return out.Responses[id].Rate
}

func (out BandOutput) Valid(id int) bool {
	return out.Responses[id].ResponseCode == 0
}
