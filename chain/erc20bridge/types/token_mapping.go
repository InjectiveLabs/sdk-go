package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

// ValidateBasic does basic validation of a TokenMapping
func (v TokenMapping) ValidateBasic() error {
	if err := sdk.ValidateDenom(v.CosmosDenom); err != nil {
		return fmt.Errorf("invalid cosmos denom: %s: %s", v.CosmosDenom, err.Error())
	}

	if !common.IsHexAddress(v.Erc20Address) {
		return fmt.Errorf("invalid child address: %s", v.Erc20Address)
	}

	return nil
}
