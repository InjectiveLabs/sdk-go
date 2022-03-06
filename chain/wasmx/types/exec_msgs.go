package types

type CWRegisterExecMsg struct {
	Register *RegisterMsg `json:"register,omitempty"`
}

type RegisterMsg struct {
	GasLimit        uint64 `json:"gas_limit"`
	ContractAddress string `json:"contract_address"`
}

type CWBeginBlockerExecMsg struct {
	BeginBlockerMsg *BeginBlockerMsg `json:"begin_blocker,omitempty"`
}

type BeginBlockerMsg struct {
}

type Contract struct {
	Address  string `json:"address"`
	GasLimit uint64 `json:"gas_limit"`
}
