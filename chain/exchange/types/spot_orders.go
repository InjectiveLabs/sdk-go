package types

import (
	"github.com/ethereum/go-ethereum/common"
)

func (o *SpotOrder) SubaccountID() common.Hash {
	return o.OrderInfo.SubaccountID()
}

func (o *SpotLimitOrder) SubaccountID() common.Hash {
	return o.OrderInfo.SubaccountID()
}

func (o *SpotMarketOrder) SubaccountID() common.Hash {
	return o.OrderInfo.SubaccountID()
}
