package types

import (
	"encoding/json"

	"cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type PrivilegedAction struct {
	SyntheticTrade   *SyntheticTradeAction `json:"synthetic_trade"`
	PositionTransfer *PositionTransfer     `json:"position_transfer"`
}

type InjectiveAction interface {
	// ValidateBasic does a simple validation check that
	// doesn't require access to any other information.
	ValidateBasic() error
}

func ParseRequest(data []byte) (InjectiveAction, error) {
	if len(data) == 0 || string(data) == "null" {
		return nil, nil
	}

	var request PrivilegedAction
	err := json.Unmarshal(data, &request)

	if err != nil {
		return nil, errors.Wrap(err, "failed to parse exchange action request")
	}

	if request.SyntheticTrade != nil {
		err = request.SyntheticTrade.ValidateBasic()
		if err != nil {
			return request.SyntheticTrade, errors.Wrap(err, "invalid synthetic trade request")
		}

		return request.SyntheticTrade, nil
	}

	if request.PositionTransfer != nil {
		err = request.PositionTransfer.ValidateBasic()
		if err != nil {
			return request.PositionTransfer, errors.Wrap(err, "invalid position transfer request")
		}

		return request.PositionTransfer, nil
	}

	return nil, errors.Wrap(sdkerrors.ErrUnknownRequest, "unknown variant of InjectiveAction")
}
