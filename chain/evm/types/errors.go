package types

import (
	"encoding/json"
	"errors"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/vm"
)

const (
	codeErrInvalidState      = uint32(iota) + 2 // NOTE: code 1 is reserved for internal errors
	codeErrExecutionReverted                    // IMPORTANT: Do not move this error as it complies with the JSON-RPC error standard
	codeErrChainConfigNotFound
	codeErrInvalidChainConfig
	codeErrZeroAddress
	codeErrEmptyHash
	codeErrBloomNotFound
	codeErrTxReceiptNotFound
	codeErrCreateDisabled
	codeErrCallDisabled
	codeErrInvalidAmount
	codeErrInvalidGasPrice
	codeErrInvalidGasFee
	codeErrVMExecution
	codeErrInvalidRefund
	codeErrInconsistentGas
	codeErrInvalidGasCap
	codeErrInvalidBaseFee
	codeErrGasOverflow
	codeErrInvalidAccount
	codeErrInvalidGasLimit
	codeErrConfigOverrides
)

var ErrPostTxProcessing = errors.New("failed to execute post processing")

var (
	// ErrInvalidState returns an error resulting from an invalid Storage State.
	ErrInvalidState = errorsmod.Register(ModuleName, codeErrInvalidState, "invalid storage state")

	// ErrExecutionReverted returns an error resulting from an error in EVM execution.
	ErrExecutionReverted = errorsmod.Register(ModuleName, codeErrExecutionReverted, vm.ErrExecutionReverted.Error())

	// ErrChainConfigNotFound returns an error if the chain config cannot be found on the store.
	ErrChainConfigNotFound = errorsmod.Register(ModuleName, codeErrChainConfigNotFound, "chain configuration not found")

	// ErrInvalidChainConfig returns an error resulting from an invalid ChainConfig.
	ErrInvalidChainConfig = errorsmod.Register(ModuleName, codeErrInvalidChainConfig, "invalid chain configuration")

	// ErrZeroAddress returns an error resulting from an zero (empty) ethereum Address.
	ErrZeroAddress = errorsmod.Register(ModuleName, codeErrZeroAddress, "invalid zero address")

	// ErrEmptyHash returns an error resulting from an empty ethereum Hash.
	ErrEmptyHash = errorsmod.Register(ModuleName, codeErrEmptyHash, "empty hash")

	// ErrBloomNotFound returns an error if the block bloom cannot be found on the store.
	ErrBloomNotFound = errorsmod.Register(ModuleName, codeErrBloomNotFound, "block bloom not found")

	// ErrTxReceiptNotFound returns an error if the transaction receipt could not be found
	ErrTxReceiptNotFound = errorsmod.Register(ModuleName, codeErrTxReceiptNotFound, "transaction receipt not found")

	// ErrCreateDisabled returns an error if the EnableCreate parameter is false.
	ErrCreateDisabled = errorsmod.Register(ModuleName, codeErrCreateDisabled, "EVM Create operation is disabled")

	// ErrCallDisabled returns an error if the EnableCall parameter is false.
	ErrCallDisabled = errorsmod.Register(ModuleName, codeErrCallDisabled, "EVM Call operation is disabled")

	// ErrInvalidAmount returns an error if a tx contains an invalid amount.
	ErrInvalidAmount = errorsmod.Register(ModuleName, codeErrInvalidAmount, "invalid transaction amount")

	// ErrInvalidGasPrice returns an error if an invalid gas price is provided to the tx.
	ErrInvalidGasPrice = errorsmod.Register(ModuleName, codeErrInvalidGasPrice, "invalid gas price")

	// ErrInvalidGasFee returns an error if the tx gas fee is out of bound.
	ErrInvalidGasFee = errorsmod.Register(ModuleName, codeErrInvalidGasFee, "invalid gas fee")

	// ErrVMExecution returns an error resulting from an error in EVM execution.
	ErrVMExecution = errorsmod.Register(ModuleName, codeErrVMExecution, "evm transaction execution failed")

	// ErrInvalidRefund returns an error if a the gas refund value is invalid.
	ErrInvalidRefund = errorsmod.Register(ModuleName, codeErrInvalidRefund, "invalid gas refund amount")

	// ErrInconsistentGas returns an error if a the gas differs from the expected one.
	ErrInconsistentGas = errorsmod.Register(ModuleName, codeErrInconsistentGas, "inconsistent gas")

	// ErrInvalidGasCap returns an error if a the gas cap value is negative or invalid
	ErrInvalidGasCap = errorsmod.Register(ModuleName, codeErrInvalidGasCap, "invalid gas cap")

	// ErrInvalidBaseFee returns an error if a the base fee cap value is invalid
	ErrInvalidBaseFee = errorsmod.Register(ModuleName, codeErrInvalidBaseFee, "invalid base fee")

	// ErrGasOverflow returns an error if gas computation overlow/underflow
	ErrGasOverflow = errorsmod.Register(ModuleName, codeErrGasOverflow, "gas computation overflow/underflow")

	// ErrInvalidAccount returns an error if the account is not an EVM compatible account
	ErrInvalidAccount = errorsmod.Register(ModuleName, codeErrInvalidAccount, "account type is not a valid ethereum account")

	// ErrInvalidGasLimit returns an error if gas limit value is invalid
	ErrInvalidGasLimit = errorsmod.Register(ModuleName, codeErrInvalidGasLimit, "invalid gas limit")

	ErrConfigOverrides = errorsmod.Register(ModuleName, codeErrConfigOverrides, "failed to apply state override")
)

// VmError is an interface that represents a reverted or failed EVM execution.
// The Ret() method returns the revert reason bytes associated with the error, if any.
// The Cause() method returns the unwrapped error. For ABCIInfo integration.
type VmError interface {
	String() string
	Cause() error
	Error() string
	VmError() string
	Ret() []byte
	Reason() string
}

type vmErrorWithRet struct {
	cause   error
	vmErr   string
	ret     []byte
	hash    string
	gasUsed uint64

	// derived data

	err    error
	reason string
}

// Ret returns the revert reason bytes associated with the VmError.
func (e *vmErrorWithRet) Ret() []byte {
	if e == nil {
		return nil
	}

	return e.ret
}

// VmError returns the VM error string associated with the VmError.
func (e *vmErrorWithRet) VmError() string {
	if e == nil {
		return ""
	}

	return e.vmErr
}

// Reason returns the reason string associated with the VmError.
func (e *vmErrorWithRet) Reason() string {
	if e == nil {
		return ""
	}

	return e.reason
}

// Cause returns the module-level error that can be used for ABCIInfo integration.
func (e *vmErrorWithRet) Cause() error {
	if e == nil {
		return nil
	}

	return e.cause
}

type abciLogVmError struct {
	Hash    string `json:"tx_hash"`
	GasUsed uint64 `json:"gas_used"`
	Reason  string `json:"reason,omitempty"`
	VmError string `json:"vm_error"`
	Ret     []byte `json:"ret,omitempty"`
}

// String returns a human-readable string of the error. Includes revert reason if available.
func (e *vmErrorWithRet) String() string {
	if e == nil || e.err == nil {
		return ""
	}

	return e.err.Error()
}

// Error returns a JSON-encoded string representation of the VmErrorWithRet.
// This includes the transaction hash, gas used, revert reason (if available),
// and the VM error that occurred. For integration with ABCIInfo.
func (e *vmErrorWithRet) Error() string {
	if e == nil {
		return ""
	}

	return e.toJSON()
}

func (e *vmErrorWithRet) toJSON() string {
	logData := &abciLogVmError{
		Hash:    e.hash,
		GasUsed: e.gasUsed,
		Reason:  e.reason,
		VmError: e.vmErr,
		Ret:     e.ret,
	}

	logStr, marshalError := json.Marshal(logData)
	if marshalError != nil {
		// fallback to the original error as text
		// we cannot handle error in this flow, nor panic
		return e.err.Error()
	}

	return string(logStr)
}

// NewVmErrorWithRet creates a new VmError augmented with the revert reason bytes.
// cause is the module-level error that can be used for ABCIInfo integration.
// vmErr is the VM error string associated with the VmError.
// ret is the revert reason bytes associated with the VmError (only if the VM error is ErrExecutionReverted).
func NewVmErrorWithRet(
	vmErr string,
	ret []byte,
	hash string,
	gasUsed uint64,
) VmError {
	e := &vmErrorWithRet{
		vmErr:   vmErr,
		hash:    hash,
		gasUsed: gasUsed,
	}

	if e.vmErr == vm.ErrExecutionReverted.Error() {
		e.err = vm.ErrExecutionReverted

		// store only if the VM error is ErrExecutionReverted
		e.ret = common.CopyBytes(ret)

		reason, errUnpack := abi.UnpackRevert(e.ret)
		if errUnpack == nil {
			e.err = fmt.Errorf("%s: %s", e.err.Error(), reason)
			e.reason = reason
			e.cause = errorsmod.Wrap(ErrExecutionReverted, e.toJSON())
		}
	} else {
		e.err = errors.New(e.vmErr)
		e.cause = errorsmod.Wrap(ErrVMExecution, e.toJSON())
	}

	return e
}

// NewExecErrorWithReason unpacks the revert return bytes and returns a wrapped error
// with the return reason.
func NewExecErrorWithReason(revertReason []byte) *RevertError {
	result := common.CopyBytes(revertReason)
	reason, errUnpack := abi.UnpackRevert(result)

	err := vm.ErrExecutionReverted
	if errUnpack == nil {
		err = fmt.Errorf("%s: %v", err.Error(), reason)
	}

	return &RevertError{
		error:  err,
		reason: hexutil.Encode(result),
	}
}

// RevertError is an API error that encompass an EVM revert with JSON error
// code and a binary data blob.
type RevertError struct {
	error
	reason string // revert reason hex encoded
}

// ErrorCode returns the JSON error code for a revert.
// See: https://github.com/ethereum/wiki/wiki/JSON-RPC-Error-Codes-Improvement-Proposal
func (e *RevertError) ErrorCode() int {
	return 3
}

// ErrorData returns the hex encoded revert reason.
func (e *RevertError) ErrorData() interface{} {
	return e.reason
}
