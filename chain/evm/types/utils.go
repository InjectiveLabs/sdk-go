package types

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/gogoproto/proto"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// DefaultPriorityReduction is the default amount of price values required for 1 unit of priority.
// Because priority is `int64` while price is `big.Int`, it's necessary to scale down the range to keep it more pratical.
// The default value is the same as the `sdk.DefaultPowerReduction`.
var DefaultPriorityReduction = sdk.DefaultPowerReduction

var (
	// EmptyCodeHash is the known hash of the empty EVM bytecode.
	EmptyCodeHash = crypto.Keccak256(nil)

	// EmptyRootHash is the known root hash of an empty merkle trie.
	EmptyRootHash = common.HexToHash("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421")
)

// DecodeTxResponse decodes a protobuf-encoded byte slice into TxResponse
func DecodeTxResponse(in []byte) (*MsgEthereumTxResponse, error) {
	responses, err := DecodeTxResponses(in)
	if err != nil {
		return nil, err
	}
	if len(responses) == 0 {
		return &MsgEthereumTxResponse{}, nil
	}
	return responses[0], nil
}

// DecodeTxResponses decodes a protobuf-encoded byte slice into TxResponses
func DecodeTxResponses(in []byte) ([]*MsgEthereumTxResponse, error) {
	var txMsgData sdk.TxMsgData
	if err := proto.Unmarshal(in, &txMsgData); err != nil {
		return nil, err
	}
	responses := make([]*MsgEthereumTxResponse, 0, len(txMsgData.MsgResponses))
	for _, res := range txMsgData.MsgResponses {
		var response MsgEthereumTxResponse
		if res.TypeUrl != "/"+proto.MessageName(&response) {
			continue
		}
		err := proto.Unmarshal(res.Value, &response)
		if err != nil {
			return nil, errorsmod.Wrap(err, "failed to unmarshal tx response message data")
		}
		responses = append(responses, &response)
	}
	return responses, nil
}

func logsFromTxResponse(dst []*ethtypes.Log, rsp *MsgEthereumTxResponse, blockNumber uint64) []*ethtypes.Log {
	if len(rsp.Logs) == 0 {
		return nil
	}

	if dst == nil {
		dst = make([]*ethtypes.Log, 0, len(rsp.Logs))
	}

	txHash := common.HexToHash(rsp.Hash)
	for _, log := range rsp.Logs {
		// fill in the tx/block informations
		l := log.ToEthereum()
		l.TxHash = txHash
		l.BlockNumber = blockNumber
		if len(rsp.BlockHash) > 0 {
			l.BlockHash = common.BytesToHash(rsp.BlockHash)
		}
		dst = append(dst, l)
	}
	return dst
}

// DecodeMsgLogsFromEvents decodes a protobuf-encoded byte slice into ethereum logs, for a single message.
func DecodeMsgLogsFromEvents(in []byte, events []abci.Event, msgIndex int, blockNumber uint64) ([]*ethtypes.Log, error) {
	txResponses, err := DecodeTxResponses(in)
	if err != nil {
		return nil, err
	}

	var logs []*ethtypes.Log
	if msgIndex < len(txResponses) {
		logs = logsFromTxResponse(nil, txResponses[msgIndex], blockNumber)
	}

	if len(logs) == 0 {
		logs, err = TxLogsFromEvents(events, msgIndex)
	}

	return logs, err
}

// TxLogsFromEvents parses ethereum logs from cosmos events for specific msg index
func TxLogsFromEvents(events []abci.Event, msgIndex int) ([]*ethtypes.Log, error) {
	for _, event := range events {
		if event.Type != EventTypeTxLog {
			continue
		}

		if msgIndex > 0 {
			// not the eth tx we want
			msgIndex--
			continue
		}

		return ParseTxLogsFromEvent(event)
	}

	return []*ethtypes.Log{}, nil
}

// ParseTxLogsFromEvent parse tx logs from one event
func ParseTxLogsFromEvent(event abci.Event) ([]*ethtypes.Log, error) {
	logs := make([]*Log, 0, len(event.Attributes))
	for _, attr := range event.Attributes {
		if attr.Key != AttributeKeyTxLog {
			continue
		}

		var log Log
		if err := json.Unmarshal([]byte(attr.Value), &log); err != nil {
			return nil, err
		}

		logs = append(logs, &log)
	}
	return LogsToEthereum(logs), nil
}

// DecodeTxLogsFromEvents decodes a protobuf-encoded byte slice into ethereum logs
func DecodeTxLogsFromEvents(in []byte, events []abci.Event, blockNumber uint64) ([]*ethtypes.Log, error) {
	txResponses, err := DecodeTxResponses(in)
	if err != nil {
		return nil, err
	}
	var logs []*ethtypes.Log
	for _, response := range txResponses {
		logs = logsFromTxResponse(logs, response, blockNumber)
	}
	if len(logs) == 0 {
		for _, event := range events {
			if event.Type != EventTypeTxLog {
				continue
			}
			txLogs, err := ParseTxLogsFromEvent(event)
			if err != nil {
				return nil, err
			}
			logs = append(logs, txLogs...)
		}
	}
	return logs, nil
}

// EncodeTransactionLogs encodes TransactionLogs slice into a protobuf-encoded byte slice.
func EncodeTransactionLogs(res *TransactionLogs) ([]byte, error) {
	return proto.Marshal(res)
}

// DecodeTransactionLogs decodes an protobuf-encoded byte slice into TransactionLogs
func DecodeTransactionLogs(data []byte) (TransactionLogs, error) {
	var logs TransactionLogs
	err := proto.Unmarshal(data, &logs)
	if err != nil {
		return TransactionLogs{}, err
	}
	return logs, nil
}

// UnwrapEthereumMsg extract MsgEthereumTx from wrapping sdk.Tx
func UnwrapEthereumMsg(tx sdk.Tx, ethHash common.Hash) (*MsgEthereumTx, error) {
	if tx == nil {
		return nil, fmt.Errorf("invalid tx: nil")
	}

	for _, msg := range tx.GetMsgs() {
		ethMsg, ok := msg.(*MsgEthereumTx)
		if !ok {
			return nil, fmt.Errorf("invalid tx type: %T", tx)
		}
		txHash := ethMsg.AsTransaction().Hash()
		if txHash == ethHash {
			return ethMsg, nil
		}
	}

	return nil, fmt.Errorf("eth tx not found: %s", ethHash)
}

// BinSearch execute the binary search and hone in on an executable gas limit
func BinSearch(lo, hi uint64, executable func(uint64) (bool, *MsgEthereumTxResponse, error)) (uint64, error) {
	for lo+1 < hi {
		mid := (hi + lo) / 2
		failed, _, err := executable(mid)
		// If the error is not nil(consensus error), it means the provided message
		// call or transaction will never be accepted no matter how much gas it is
		// assigned. Return the error directly, don't struggle anymore.
		if err != nil {
			return 0, err
		}
		if failed {
			lo = mid
		} else {
			hi = mid
		}
	}
	return hi, nil
}

// EffectiveGasPrice compute the effective gas price based on eip-1159 rules
// `effectiveGasPrice = min(baseFee + tipCap, feeCap)`
func EffectiveGasPrice(baseFee, feeCap, tipCap *big.Int) *big.Int {
	return math.BigMin(new(big.Int).Add(tipCap, baseFee), feeCap)
}

// HexAddress encode ethereum address without checksum, faster to run for state machine
func HexAddress(a []byte) string {
	var buf [common.AddressLength*2 + 2]byte
	copy(buf[:2], "0x")
	hex.Encode(buf[2:], a)
	return string(buf[:])
}
