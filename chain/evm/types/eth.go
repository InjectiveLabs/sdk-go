package types

import (
	"encoding/json"

	errorsmod "cosmossdk.io/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type EthereumTx struct {
	*ethtypes.Transaction
}

func NewEthereumTx(txData ethtypes.TxData) EthereumTx {
	return EthereumTx{ethtypes.NewTx(txData)}
}

func (tx EthereumTx) Size() int {
	if tx.Transaction == nil {
		return 0
	}
	return int(tx.Transaction.Size())
}

func (tx EthereumTx) MarshalTo(dst []byte) (int, error) {
	if tx.Transaction == nil {
		return 0, nil
	}
	bz, err := tx.MarshalBinary()
	if err != nil {
		return 0, err
	}
	copy(dst, bz)
	return len(bz), nil
}

func (tx *EthereumTx) Unmarshal(dst []byte) error {
	if len(dst) == 0 {
		tx.Transaction = nil
		return nil
	}
	if tx.Transaction == nil {
		tx.Transaction = new(ethtypes.Transaction)
	}
	return tx.UnmarshalBinary(dst)
}

func (tx *EthereumTx) UnmarshalJSON(bz []byte) error {
	var data hexutil.Bytes
	if err := json.Unmarshal(bz, &data); err != nil {
		return err
	}
	return tx.Unmarshal(data)
}

func (tx EthereumTx) MarshalJSON() ([]byte, error) {
	bz, err := tx.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return json.Marshal(hexutil.Bytes(bz))
}

func (tx EthereumTx) Validate() error {
	// prevent txs with 0 gas to fill up the mempool
	if tx.Gas() == 0 {
		return errorsmod.Wrap(ErrInvalidGasLimit, "gas limit must not be zero")
	}
	if !IsValidInt256(tx.GasPrice()) {
		return errorsmod.Wrap(ErrInvalidGasPrice, "out of bound")
	}
	if !IsValidInt256(tx.GasFeeCap()) {
		return errorsmod.Wrap(ErrInvalidGasPrice, "out of bound")
	}
	if !IsValidInt256(tx.GasTipCap()) {
		return errorsmod.Wrap(ErrInvalidGasPrice, "out of bound")
	}
	if !IsValidInt256(tx.Cost()) {
		return errorsmod.Wrap(ErrInvalidGasFee, "out of bound")
	}
	return nil
}
