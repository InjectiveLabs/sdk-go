package types

import (
	bandprice "github.com/InjectiveLabs/sdk-go/bandchain/hooks/price"
	"github.com/InjectiveLabs/sdk-go/bandchain/obi"
	bandoracle "github.com/InjectiveLabs/sdk-go/bandchain/oracle/types"
)

func NewOracleRequestPacketData(clientID string, calldata []byte, r *BandOracleRequest) bandoracle.OracleRequestPacketData {
	return bandoracle.OracleRequestPacketData{
		ClientID:       clientID,
		OracleScriptID: bandoracle.OracleScriptID(r.OracleScriptId),
		Calldata:       calldata,
		AskCount:       r.AskCount,
		MinCount:       r.MinCount,
		FeeLimit:       r.FeeLimit,
		RequestKey:     r.RequestKey,
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

	callData := obi.MustEncode(requestCallData)

	return callData
}