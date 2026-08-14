package types

import (
	"fmt"
	"math"
	"math/big"
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

const (
	// IBCHookCallerAddressHex is the trusted system caller used for inbound IBC-triggered EVM calls.
	IBCHookCallerAddressHex = "0x0000000000000000000000000000000000000068"
	onIBCTransferMethod     = "onIBCTransfer"
)

var onIBCTransferABI = mustParseABI(`[{
		"type": "function",
		"name": "onIBCTransfer",
		"inputs": [
			{"name": "destChannelId", "type": "string"},
			{"name": "packetSender", "type": "string"},
			{"name": "token", "type": "address"},
			{"name": "amount", "type": "uint256"},
			{"name": "receiver", "type": "address"},
			{"name": "payload", "type": "bytes"}
		],
		"outputs": [{"name": "", "type": "bytes"}]
	}]`)

// PackIBCHookCall ABI-encodes an onIBCTransfer call
func PackIBCHookCall(
	destChannelID,
	packetSender,
	hookContract,
	token string,
	amount sdkmath.Int,
	receiver sdk.AccAddress,
	payload []byte,
	gasLimit uint64,
) (*IBCHookCall, error) {
	if !common.IsHexAddress(hookContract) {
		return nil, fmt.Errorf("invalid hook contract address")
	}

	input, err := onIBCTransferABI.Pack(
		onIBCTransferMethod,
		destChannelID,
		packetSender,
		common.HexToAddress(token).Bytes(),
		amount.BigInt(),
		common.BytesToAddress(receiver),
		payload,
	)

	if err != nil {
		return nil, err
	}

	call := &IBCHookCall{
		Contract: common.FromHex(hookContract),
		Input:    input,
		GasLimit: gasLimit,
	}

	return call, nil
}

// GetTxPriority returns the priority of a given Ethereum tx. It relies of the
// priority reduction global variable to calculate the tx priority given the tx
// gas price:
//
//	tx_priority = gas_price / priority_reduction
func GetTxPriority(msg *MsgEthereumTx) (priority int64) {
	// calculate priority based on gas price
	gasPrice := msg.AsTransaction().GasPrice()

	priority = math.MaxInt64
	priorityBig := new(big.Int).Quo(gasPrice, DefaultPriorityReduction.BigInt())

	// safety check
	if priorityBig.IsInt64() {
		priority = priorityBig.Int64()
	}

	return priority
}

// Failed returns if the contract execution failed in vm errors
func (m *MsgEthereumTxResponse) Failed() bool {
	return m.VmError != ""
}

// Return is a helper function to help caller distinguish between revert reason
// and function return. Return returns the data after execution if no error occurs.
func (m *MsgEthereumTxResponse) Return() []byte {
	if m.Failed() {
		return nil
	}
	return common.CopyBytes(m.Ret)
}

// Revert returns the concrete revert reason if the execution is aborted by `REVERT`
// opcode. Note the reason can be nil if no data supplied with revert opcode.
func (m *MsgEthereumTxResponse) Revert() []byte {
	if m.VmError != vm.ErrExecutionReverted.Error() {
		return nil
	}
	return common.CopyBytes(m.Ret)
}

func mustParseABI(raw string) abi.ABI {
	parsed, err := abi.JSON(strings.NewReader(raw))
	if err != nil {
		panic(errorsmod.Wrap(err, "failed to parse IBC EVM hook ABI"))
	}
	return parsed
}
