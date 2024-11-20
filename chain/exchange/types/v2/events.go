package v2

import "github.com/ethereum/go-ethereum/common"

func NewEventOrderCancelFail(
	marketID,
	subaccountID common.Hash,
	orderHash,
	cid string,
	err error,
) *EventOrderCancelFail {
	ev := &EventOrderCancelFail{
		MarketId:     marketID.Hex(),
		SubaccountId: subaccountID.Hex(),
		OrderHash:    orderHash,
		Cid:          cid,
	}

	if err != nil {
		ev.Description = err.Error()
	}

	return ev
}

func (e *EventOrderFail) AddOrderFail(orderHash common.Hash, cid string, flag uint32) {
	e.Hashes = append(e.Hashes, orderHash.Bytes())
	e.Flags = append(e.Flags, flag)

	if cid != "" {
		e.Cids = append(e.Cids, cid)
	}
}

func (e *EventOrderFail) IsEmpty() bool {
	return len(e.Flags) == 0 && len(e.Hashes) == 0 && len(e.Cids) == 0
}
