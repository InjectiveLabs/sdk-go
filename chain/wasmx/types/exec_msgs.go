package types

import (
	"encoding/json"
	fmt "fmt"
)

const DefaultGasLimitContractRegistration = uint64(8987600)

type CWRegisterExecMsg struct {
	Register *RegisterMsg `json:"register,omitempty"`
}

type RegisterMsg struct {
	GasLimit        uint64 `json:"gas_limit"`
	ContractAddress string `json:"contract_address"`
	GasPrice        string `json:"gas_price"`
}

type CWBeginBlockerExecMsg struct {
	BeginBlockerMsg *BeginBlockerMsg `json:"begin_blocker,omitempty"`
}

type BeginBlockerMsg struct {
}

func BeginBlockerExecMsg() ([]byte, error) {
	// Construct Exec message
	beginBlocker := CWBeginBlockerExecMsg{BeginBlockerMsg: &BeginBlockerMsg{}}

	// execMsg := []byte("{\"begin_blocker\":{}}")
	execMsg, err := json.Marshal(beginBlocker)
	if err != nil {
		fmt.Println("Register marshal failed")
		return nil, err
	}

	return execMsg, nil
}

type CWGetContractsQueryMsg struct {
	QueryContractsMsg *QueryContractsMsg `json:"get_contracts,omitempty"`
}

type QueryContractsMsg struct {
}

func ContractsQueryMsg() ([]byte, error) {
	// Construct Exec message
	contractQuery := CWGetContractsQueryMsg{QueryContractsMsg: &QueryContractsMsg{}}

	// queryData := []byte("{\"get_contracts\": {}}")
	queryMsg, err := json.Marshal(contractQuery)
	if err != nil {
		fmt.Println("Register marshal failed")
		return nil, err
	}

	return queryMsg, nil
}

type Contract struct {
	Address  string `json:"address"`
	GasLimit uint64 `json:"gas_limit"`
	GasPrice string `json:"gas_price"`
}
