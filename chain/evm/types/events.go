package types

import (
	"github.com/ethereum/go-ethereum/common"

	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
)

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

func NewEventIBCEVMHookTx( //nolint:revive // all good
	packet ibcexported.PacketI,
	packetSender string,
	receiver,
	contract,
	token,
	from common.Address,
	amount string,
	input []byte,
	gasLimit uint64,
	response *MsgEthereumTxResponse,
) *EventIBCEVMHookTx {
	return &EventIBCEVMHookTx{
		SourcePort:         packet.GetSourcePort(),
		SourceChannel:      packet.GetSourceChannel(),
		DestinationPort:    packet.GetDestPort(),
		DestinationChannel: packet.GetDestChannel(),
		Sequence:           packet.GetSequence(),
		PacketSender:       packetSender,
		Receiver:           receiver.Hex(),
		Contract:           contract.Hex(),
		Token:              token.Hex(),
		Amount:             amount,
		From:               from.Hex(),
		Input:              input,
		GasLimit:           gasLimit,
		Response:           response,
	}
}
