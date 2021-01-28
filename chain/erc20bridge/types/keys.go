package types

// constants
const (
	// module name
	ModuleName = "erc20bridge"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for message routing
	RouterKey = ModuleName

	// events
	EventTypeTemplateCreation = "template_creation"
	EventTypeBurn             = "burn"

	AttributeKeyAddress                = "address"
	AttributeKeyProposedCosmosCoin     = "propsed_cosmos_coin"
	AttributeKeyContractAddress        = "burn_contract_address"
	AttributeKeyBurnEthAddress         = "burn_eth_address"
	AttributeKeyRecipientCosmosAddress = "recipient_cosmos_address"
	AttributeKeyAmount                 = "amount"

	TypeMsgMint = "mint_erc20"
)
