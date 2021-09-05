package types

import (
	bandprice "github.com/InjectiveLabs/sdk-go/bandchain/hooks/price"
	bandobi "github.com/bandprotocol/bandchain-packet/obi"
	bandPacket "github.com/bandprotocol/bandchain-packet/packet"
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
func (r *BandOracleRequest) GetCalldata() []byte {

	requestCallData := bandprice.Input{
		Symbols:    r.Symbols,
		Multiplier: BandPriceMultiplier,
	}

	callData := bandobi.MustEncode(requestCallData)

	return callData
}
