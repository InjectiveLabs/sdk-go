package types

import (
	"encoding/json"
)

func NewBeginBlockerExecMsg() ([]byte, error) {
	// Construct Exec message
	beginBlocker := CWBeginBlockerExecMsg{BeginBlockerMsg: &BeginBlockerMsg{}}

	// nolint:all
	// execMsg := []byte(`{"begin_blocker":{}}`)
	execMsg, err := json.Marshal(beginBlocker)
	if err != nil {
		return nil, err
	}

	return execMsg, nil
}

type CWBeginBlockerExecMsg struct {
	BeginBlockerMsg *BeginBlockerMsg `json:"begin_blocker,omitempty"`
}

type BeginBlockerMsg struct {
}
