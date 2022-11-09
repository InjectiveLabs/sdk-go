package types

func NewRegistryDeregisterCallbackMsg() RegistryDeregisterCallbackMsg {
	return RegistryDeregisterCallbackMsg{
		DeregisterCallback: &DeregisterCallbackMsg{},
	}
}

type RegistryDeregisterCallbackMsg struct {
	DeregisterCallback *DeregisterCallbackMsg `json:"deregister,omitempty"`
}

type DeregisterCallbackMsg struct{}

func NewRegistryDeactivateCallbackMsg() RegistryDeactivateCallbackMsg {
	return RegistryDeactivateCallbackMsg{
		DeactivateCallback: &DeactivateCallbackMsg{},
	}
}

type RegistryDeactivateCallbackMsg struct {
	DeactivateCallback *DeactivateCallbackMsg `json:"deactivate,omitempty"`
}

type DeactivateCallbackMsg struct{}
