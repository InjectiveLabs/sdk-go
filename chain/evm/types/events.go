package types

import "github.com/ethereum/go-ethereum/common"

// Evm module events
const (
	EventTypeEthereumTx = TypeMsgEthereumTx
	EventTypeBlockBloom = "block_bloom"

	AttributeKeyContractAddress = "contract"
	AttributeKeyRecipient       = "recipient"
	AttributeKeyTxHash          = "txHash"
	AttributeKeyEthereumTxHash  = "ethereumTxHash"
	AttributeKeyTxIndex         = "txIndex"
	AttributeKeyTxGasUsed       = "txGasUsed"
	AttributeKeyTxType          = "txType"

	// tx failed in eth vm execution
	AttributeKeyEthereumTxFailed = "ethereumTxFailed"
	AttributeValueCategory       = ModuleName
	AttributeKeyEthereumBloom    = "bloom"

	MetricKeyTransitionDB = "transition_db"
	MetricKeyStaticCall   = "static_call"
)

// IBCEVMHookAcknowledgement wraps the underlying IBC acknowledgement with the EVM hook contract result.
type IBCEVMHookAcknowledgement struct {
	ContractResult []byte `json:"contract_result"`
	IBCAck         []byte `json:"ibc_ack"`
}

func NewEventIBCHookCall( //nolint:revive // all good
	destPort, destCh string,
	sequence uint64,
	contract,
	input []byte,
	response *MsgEthereumTxResponse,
) *EventIBCHookCall {
	event := &EventIBCHookCall{
		DestinationPort:    destPort,
		DestinationChannel: destCh,
		Sequence:           sequence,
		Contract:           common.BytesToAddress(contract).Hex(),
		Input:              input,
	}

	if response == nil {
		return event
	}

	event.Success = !response.Failed()
	event.ReturnData = response.Ret
	event.Error = response.VmError
	event.GasUsed = response.GasUsed
	if event.Success {
		event.Logs = response.Logs
	}

	return event
}
