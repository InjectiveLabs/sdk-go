package types

import (
	"github.com/ethereum/go-ethereum/common"
)

// Event type and attribute constants

func (e *EventOrderFail) AddOrderFail(orderHash common.Hash, flag uint32) {
	e.Hashes = append(e.Hashes, orderHash.Bytes())
	e.Flags = append(e.Flags, flag)
}

func (e *EventOrderFail) IsEmpty() bool {
	return len(e.Flags) == 0 && len(e.Hashes) == 0
}
