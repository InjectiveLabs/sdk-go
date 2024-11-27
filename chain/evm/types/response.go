package types

import (
	abci "github.com/cometbft/cometbft/abci/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/cosmos/gogoproto/proto"
)

// PatchTxResponses fills the evm tx index and log indexes in the tx result
func PatchTxResponses(input []*abci.ExecTxResult) []*abci.ExecTxResult {
	var (
		txIndex  uint64
		logIndex uint64
	)
	for _, res := range input {
		// assume no error result in msg handler
		if res.Code != 0 {
			continue
		}

		var txMsgData sdk.TxMsgData
		if err := proto.Unmarshal(res.Data, &txMsgData); err != nil {
			panic(err)
		}

		var (
			// if the response data is modified and need to be marshaled back
			dataDirty bool
		)

		for i, rsp := range txMsgData.MsgResponses {
			var response MsgEthereumTxResponse
			if rsp.TypeUrl != "/"+proto.MessageName(&response) {
				continue
			}

			if err := proto.Unmarshal(rsp.Value, &response); err != nil {
				panic(err)
			}

			if len(response.Logs) > 0 {
				for _, log := range response.Logs {
					log.TxIndex = txIndex
					log.Index = logIndex
					logIndex++
				}

				anyRsp, err := codectypes.NewAnyWithValue(&response)
				if err != nil {
					panic(err)
				}
				txMsgData.MsgResponses[i] = anyRsp

				dataDirty = true
			}

			txIndex++
		}

		if dataDirty {
			data, err := proto.Marshal(&txMsgData)
			if err != nil {
				panic(err)
			}

			res.Data = data
		}
	}

	return input
}
