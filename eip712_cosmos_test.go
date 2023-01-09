package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/InjectiveLabs/sdk-go/typeddata"
	"reflect"
	"testing"
)

var validEIP712Types = map[string]bool{
	// custom struct and array types
	"Value":   true,
	"Value[]": true,
	"string":  true,
	"bool":    true,
	"int8":    true,
	"int16":   true,
	"int32":   true,
	"int64":   true,
	"uint8":   true,
	"uint16":  true,
	"uint32":  true,
	"uint64":  true,
}

var cosmwasmMsgsInputs = []string{
	`{
		"tags": ["a", "b", "c"]
	}`,
	`{
		"recipient": "Bob",
		"amount": "100"
	}`,
	`{
		"mint":	{
			"recipient": "Bob",
			"amount": "100"
		}
	}`,
	`{
		"mint":	{
			"recipient": "Bob",
			"amount": "100",
			"funds": ["10inj", "20usdt"]
		}
	}`,
	`{
		"mint":	{
			"recipient": "Bob",
			"amount": "100",
			"funds": [
				{"denom": "inj", "amount": "1000"},
				{"denom": "usdt", "amount": "2000"}
			]
		}
	}`,
	`{
	 "action": "execute_swap_operations",
	 "msg": {
		"operation": "hey"
	 }
	}`,
	`{
	"action": "execute_swap_operations",
	"msg": {
		"operations": [
		  {
			"terra_swap": {
			  "offer_asset_info": {
				"native_token": {
				  "denom": "inj"
				}
			  },
			  "ask_asset_info": {
				"native_token": {
				  "denom": "peggy0x87aB3B4C8661e07D6372361211B96ed4Dc36B1B5"
				}
			  }
			}
		  }
		]
	}
	}`,
	`{
	"action": "execute_swap_operations",
	"tag": [1, 2, 3, 4],
	"msg": {
		"token": [{"denom": "inj", "amount": 20}],
		"tag": [1, 2, 3],
		"operations": [
		  {
			"terra_swap": {
			  "offer_asset_info": {
				"native_token": {
				  "denom": "inj"
				}
			  },
			  "ask_asset_info": {
				"native_token": {
				  "denom": "peggy0x87aB3B4C8661e07D6372361211B96ed4Dc36B1B5"
				}
			  }
			}
		  }
		]
	}
	}`,
	`{
	  "register_vault": {
		"vault_code_id": 5,
		"vault_label": "Spot Vault",
		"instantiate_vault_msg": {
		  "Spot": {
			"owner": "inj17gkuet8f6pssxd8nycm3qr9d9y699rupv6397z",
			"order_density": 10,
			"reservation_price_sensitivity_ratio": "0.5",
			"reservation_spread_sensitivity_ratio": "0.5",
			"max_active_capital_utilization_ratio": "0.5",
			"head_change_tolerance_ratio": "0.0",
			"min_head_to_tail_deviation_ratio": "0.2",
			"signed_min_head_to_fair_price_deviation_ratio": "0.1",
			"signed_min_head_to_tob_deviation_ratio": "0.1",
			"trade_volatility_group_sec": 1,
			"min_trade_volatility_sample_size": 1,
			"default_mid_price_volatility_ratio": "0.005",
			"min_volatility_ratio": "0.5",
			"master_address": "inj1qg5ega6dykkxc307y25pecuufrjkxkag6xhp6y",
			"redemption_lock_time": 60,
			"market_id": "0xa508cb32923323679f29a032c70342c147c17d0145625922b0ef22e955c844c0",
			"fair_price_tail_deviation_ratio": "0.5",
			"target_base_weight": "1",
			"allowed_subscription_types": -1,
			"allowed_redemption_types": -17
		  }
		}
	  }
	}`,
}

func TestCosmWasmParser(t *testing.T) {
	for _, msgInput := range cosmwasmMsgsInputs {
		var msg map[string]interface{}
		if err := json.Unmarshal([]byte(msgInput), &msg); err != nil {
			panic(err)
		}
		rootTypes := typeddata.Types{}
		ExtractCosmwasmTypes(CosmwasmInnerMsgMarker, rootTypes, reflect.ValueOf(msg))
		for k, t := range rootTypes {
			fmt.Println(k, t)
		}
		fmt.Println("================")
	}
}
