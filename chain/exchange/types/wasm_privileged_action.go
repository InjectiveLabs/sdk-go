package types

import (
	"bytes"
	"encoding/json"
	"io"

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

// ParseRequestStrict parses a privileged action without accepting unknown fields,
// trailing JSON values, or multiple action variants. The ordinary Wasm privileged
// execution path retains ParseRequest's compatibility behavior; consensus-authenticated
// RFQ liquidation payloads use this stricter boundary instead.
func ParseRequestStrict(data []byte) (InjectiveAction, error) {
	if len(data) == 0 || string(data) == "null" {
		return nil, nil
	}

	request, err := decodePrivilegedActionStrict(data)
	if err != nil {
		return nil, err
	}

	if request.SyntheticTrade != nil {
		if err := request.SyntheticTrade.ValidateBasic(); err != nil {
			return request.SyntheticTrade, errors.Wrap(err, "invalid synthetic trade request")
		}
		return request.SyntheticTrade, nil
	}
	if request.PositionTransfer != nil {
		if err := request.PositionTransfer.ValidateBasic(); err != nil {
			return request.PositionTransfer, errors.Wrap(err, "invalid position transfer request")
		}
		return request.PositionTransfer, nil
	}

	return nil, errors.Wrap(sdkerrors.ErrUnknownRequest, "unknown variant of InjectiveAction")
}

// ParseRFQLiquidationRequestStrict parses the dedicated RFQ liquidation action.
// Unlike ordinary synthetic trades, the isolated provider margin may cover the
// full maximum leg notional; all other strict JSON and scalar bounds remain.
func ParseRFQLiquidationRequestStrict(data []byte) (*SyntheticTradeAction, error) {
	if len(data) == 0 || string(data) == "null" {
		return nil, errors.Wrap(sdkerrors.ErrUnknownRequest, "RFQ liquidation action is empty")
	}

	request, err := decodePrivilegedActionStrict(data)
	if err != nil {
		return nil, err
	}
	if request.SyntheticTrade == nil {
		return nil, errors.Wrap(sdkerrors.ErrUnknownRequest, "RFQ liquidation requires a synthetic trade action")
	}
	if err := request.SyntheticTrade.ValidateRFQLiquidationBasic(); err != nil {
		return request.SyntheticTrade, errors.Wrap(err, "invalid RFQ liquidation synthetic trade request")
	}
	return request.SyntheticTrade, nil
}

func decodePrivilegedActionStrict(data []byte) (PrivilegedAction, error) {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()

	var request PrivilegedAction
	if err := decoder.Decode(&request); err != nil {
		return PrivilegedAction{}, errors.Wrap(err, "failed to parse exchange action request")
	}

	var trailing any
	if err := decoder.Decode(&trailing); err != io.EOF {
		if err == nil {
			return PrivilegedAction{}, errors.Wrap(sdkerrors.ErrUnknownRequest, "multiple exchange action requests")
		}
		return PrivilegedAction{}, errors.Wrap(err, "failed to parse trailing exchange action request data")
	}

	if request.SyntheticTrade != nil && request.PositionTransfer != nil {
		return PrivilegedAction{}, errors.Wrap(sdkerrors.ErrUnknownRequest, "multiple variants of InjectiveAction")
	}
	return request, nil
}
