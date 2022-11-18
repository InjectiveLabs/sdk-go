package types

import (
	"encoding/json"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewRegistryRegisterMsg(req *ContractRegistrationRequest) RegistryRegisterMsg {
	return RegistryRegisterMsg{
		Register: &RegisterMsg{
			GasLimit:        req.GasLimit,
			ContractAddress: req.ContractAddress,
			GasPrice:        req.GasPrice,
			IsExecutable:    true,
		},
	}
}

type RegistryRegisterMsg struct {
	Register *RegisterMsg `json:"register,omitempty"`
}

type RegisterMsg struct {
	GasLimit        uint64 `json:"gas_limit"`
	ContractAddress string `json:"contract_address"`
	GasPrice        uint64 `json:"gas_price"`
	IsExecutable    bool   `json:"is_executable"`
}

func NewRegistryDeregisterMsg(contract sdk.AccAddress) RegistryDeregisterMsg {
	return RegistryDeregisterMsg{
		Deregister: &DeregisterMsg{
			ContractAddress: contract.String(),
		},
	}
}

type RegistryDeregisterMsg struct {
	Deregister *DeregisterMsg `json:"deregister,omitempty"`
}

type DeregisterMsg struct {
	ContractAddress string `json:"contract_address"`
}

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

func NewRegistryDeactivateMsg(contractAddress string) ([]byte, error) {
	// Construct Exec message
	deactivateMsg := RegistryDeactivateMsg{RegistryDeactivate: &RegistryDeactivate{ContractAddress: contractAddress}}

	// nolint:all
	// execMsg := []byte('{"deactivate":{"contract_address":"inj1nc5tatafv6eyq7llkr2gv50ff9e22mnfhg8yh3"}}')
	execMsg, err := json.Marshal(deactivateMsg)
	if err != nil {
		return nil, err
	}

	return execMsg, nil
}

type RegistryDeactivateMsg struct {
	RegistryDeactivate *RegistryDeactivate `json:"deactivate,omitempty"`
}

type RegistryDeactivate struct {
	ContractAddress string `json:"contract_address"`
}

// NewRegistryContractQuery constructs the registry Exec message
func NewRegistryContractQuery() ([]byte, error) {
	contractQuery := RegistryContractQueryMsg{QueryContractsMsg: &QueryContractsMsg{}}

	queryMsg, err := json.Marshal(contractQuery)
	if err != nil {
		return nil, err
	}

	return queryMsg, nil
}

type RegistryContractQueryMsg struct {
	QueryContractsMsg *QueryContractsMsg `json:"get_contracts,omitempty"`
}

type QueryContractsMsg struct {
}

// NewRegistryActiveContractQuery constructs the registry active contracts query message
func NewRegistryActiveContractQuery() ([]byte, error) {
	contractQuery := RegistryActiveContractQueryMsg{QueryActiveContractsMsg: &QueryActiveContractsMsg{}}
	// nolint:all
	// queryData := []byte("{\"get_active_contracts\": {}}")
	queryMsg, err := json.Marshal(contractQuery)
	if err != nil {
		return nil, err
	}

	return queryMsg, nil
}

type RegistryActiveContractQueryMsg struct {
	QueryActiveContractsMsg *QueryActiveContractsMsg `json:"get_active_contracts,omitempty"`
}

type QueryActiveContractsMsg struct {
}

type RawContractExecutionParams struct {
	Address      string `json:"address"`
	GasLimit     uint64 `json:"gas_limit"`
	GasPrice     uint64 `json:"gas_price"`
	IsExecutable bool   `json:"is_executable"`
}

func (r *RawContractExecutionParams) ToContractExecutionParams() (p *ContractExecutionParams, err error) {
	addr, err := sdk.AccAddressFromBech32(r.Address)
	if err != nil {
		return nil, err
	}

	return &ContractExecutionParams{
		Address:  addr,
		GasLimit: r.GasLimit,
		GasPrice: r.GasPrice,
	}, nil
}

type ContractExecutionParams struct {
	Address      sdk.AccAddress
	GasLimit     uint64
	GasPrice     uint64
	IsExecutable bool
}

// GetSortedContractExecutionParams returns the ContractExecutionParams sorted by descending order of gas price
func GetSortedContractExecutionParams(contractExecutionList []RawContractExecutionParams) ([]*ContractExecutionParams, error) {
	paramList := make([]*ContractExecutionParams, len(contractExecutionList))
	for idx, elem := range contractExecutionList {
		if v, err := elem.ToContractExecutionParams(); err != nil {
			return nil, err
		} else {
			paramList[idx] = v
		}
	}

	sort.SliceStable(paramList, func(i, j int) bool {
		return paramList[i].GasPrice < paramList[j].GasPrice
	})

	return paramList, nil
}
