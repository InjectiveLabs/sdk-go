package types

import (
	"fmt"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/ethereum/go-ethereum/common"
)

// constants
const (
	//ProposalTypeRegisterTokenMapping string = "RegisterTokenMapping"
	ProposalTypeRegisterExchange string = "RegisterExchange"
)

func init() {
	//gov.RegisterProposalType(ProposalTypeRegisterTokenMapping)
	gov.RegisterProposalType(ProposalTypeRegisterExchange)
	gov.RegisterProposalTypeCodec(&RegisterExchangeProposal{}, "cosmos-sdk/RegisterExchangeProposal")
	//gov.RegisterProposalTypeCodec(&ResetHubProposal{}, "cosmos-sdk/ResetHubProposal")
}

//// NewTokenMapping returns an instance of TokenMapping
//func NewTokenMapping(name string, erc20Address string, cosmosDenom string, scalingFactor int64, enabled bool) TokenMapping {
//	return TokenMapping{
//		Name:          name,
//		Erc20Address:  erc20Address,
//		CosmosDenom:   cosmosDenom,
//		ScalingFactor: scalingFactor,
//		Enabled:       true,
//	}
//}
//
//// NewRegisterTokenMappingProposal returns new instance of TokenMappingProposal
//func NewRegisterTokenMappingProposal(title, description string, mapping TokenMapping) gov.Content {
//	return &RegisterTokenMappingProposal{title, description, mapping}
//}
//
//// Implements Proposal Interface
//var _ gov.Content = &RegisterTokenMappingProposal{}
//
//// ProposalRoute returns router key for this proposal
//func (sup *RegisterTokenMappingProposal) ProposalRoute() string { return RouterKey }
//
//// ProposalType returns proposal type for this proposal
//func (sup *RegisterTokenMappingProposal) ProposalType() string {
//	return ProposalTypeRegisterTokenMapping
//}
//
//// ValidateBasic returns ValidateBasic result for this proposal
//func (sup *RegisterTokenMappingProposal) ValidateBasic() error {
//	if err := sup.Mapping.ValidateBasic(); err != nil {
//		return err
//	}
//	return gov.ValidateAbstract(sup)
//}

// NewRegisterExchangeProposal returns new instance of RegisterExchangeProposal
func NewRegisterExchangeProposal(title, description string, exchange_address string) gov.Content {
	return &RegisterExchangeProposal{title, description, exchange_address}
}

// Implements Proposal Interface
var _ gov.Content = &RegisterExchangeProposal{}

// ProposalRoute returns router key for this proposal
func (rh *RegisterExchangeProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type for this proposal
func (rh *RegisterExchangeProposal) ProposalType() string {
	return ProposalTypeRegisterExchange
}

// ValidateBasic returns ValidateBasic result for this proposal
func (rh *RegisterExchangeProposal) ValidateBasic() error {
	if !common.IsHexAddress(rh.ExchangeAddress) {
		return fmt.Errorf("invalid hub address: %s", rh.ExchangeAddress)
	}
	return gov.ValidateAbstract(rh)
}
