package sdk

import (
	"bytes"
	"encoding/json"
	"math/big"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/pkg/errors"
	"github.com/xlab/structwalk"

	"github.com/InjectiveLabs/sdk-go/typeddata"
)

// WrapTxToEIP712 is an ultimate method that wraps Amino-encoded Cosmos Tx JSON data
// into an EIP712-compatible request. All messages must be of the same type.
func WrapTxToEIP712(chainID uint64, data []byte) (typeddata.TypedData, error) {
	message := make(map[string]interface{})
	if err := json.Unmarshal(data, &message); err != nil {
		err = errors.Wrap(err, "failed to unmarshal data provided into WrapTxToEIP712")
		return typeddata.TypedData{}, err
	}

	domain := typeddata.TypedDataDomain{
		Name:              "Injective Web3",
		Version:           "1.0.0",
		ChainId:           math.NewHexOrDecimal256(int64(chainID)),
		VerifyingContract: "cosmos",
		Salt:              "0",
	}

	var typedData = typeddata.TypedData{
		Types:       extractTypes("Tx", message),
		PrimaryType: "Tx",
		Domain:      domain,
		Message:     message,
	}

	return typedData, nil
}

func extractTypes(rootType string, v map[string]interface{}) typeddata.Types {
	rootTypes := typeddata.Types{
		"EIP712Domain": {
			{
				Name: "name",
				Type: "string",
			},
			{
				Name: "version",
				Type: "string",
			},
			{
				Name: "chainId",
				Type: "uint256",
			},
			{
				Name: "verifyingContract",
				Type: "string",
			},
			{
				Name: "salt",
				Type: "string",
			},
		},
		"Tx": {
			{Name: "account_number", Type: "string"},
			{Name: "chain_id", Type: "string"},
			{Name: "fee", Type: "Fee"},
			{Name: "memo", Type: "string"},
			{Name: "msgs", Type: "Msg[]"},
			{Name: "sequence", Type: "string"},
		},
		"Fee": {
			{Name: "amount", Type: "Coin[]"},
			{Name: "gas", Type: "string"},
		},
		"Coin": {
			{Name: "denom", Type: "string"},
			{Name: "amount", Type: "string"},
		},
		"Msg": {
			{Name: "type", Type: "string"},
			{Name: "value", Type: "MsgValue"},
		},
		"MsgValue": {},
	}

	// get first Msg only, all Msgs must be of the same type though
	msgValue := v["msgs"].([]interface{})[0].(map[string]interface{})["value"]

	const typeDefPrefix = "_."
	typeDefs := map[string]string{}

	fields := structwalk.FieldList(msgValue)
	for _, field := range fields {
		parent, childName, childType, ok := childTypeInParent(field, msgValue)
		if !ok {
			continue
		}

		if len(parent) == 0 {
			rootTypes["MsgValue"] = append(rootTypes["MsgValue"], typeddata.Type{
				Name: childName,
				Type: childType,
			})

			continue
		}

		typeDef := sanitizeTypedef(typeDefPrefix + parent)
		rootTypes[typeDef] = append(rootTypes[typeDef], typeddata.Type{
			Name: childName,
			Type: childType,
		})

		typeDefs[parent] = parentName(parent)
	}

	for parent, subParent := range typeDefs {
		if len(subParent) == 0 {
			rootTypes["MsgValue"] = append(rootTypes["MsgValue"], typeddata.Type{
				Name: parent,
				Type: sanitizeTypedef(typeDefPrefix + parent),
			})

			continue
		}

		subParent = sanitizeTypedef(typeDefPrefix + subParent)
		rootTypes[subParent] = append(rootTypes[subParent], typeddata.Type{
			Name: parent,
			Type: sanitizeTypedef(typeDefPrefix + parent),
		})
	}

	return rootTypes
}

// _.foo_bar.baz -> TypeFooBarBaz
//
// this is needed for Geth's own signing code which doesn't
// tolerate complex type names
func sanitizeTypedef(str string) string {
	buf := new(bytes.Buffer)
	parts := strings.Split(str, ".")

	for _, part := range parts {
		if part == "_" {
			buf.WriteString("Type")
			continue
		}

		subparts := strings.Split(part, "_")
		for _, subpart := range subparts {
			buf.WriteString(strings.Title(subpart))
		}
	}

	return buf.String()
}

// foo.bar.baz = "5" -> FooBar: { Baz: "string" }
func childTypeInParent(fullPath string, obj interface{}) (parentName, childName, childType string, ok bool) {
	parts := strings.Split(fullPath, ".")
	v, _ := structwalk.FieldValue(fullPath, obj)

	childType = typToEth(reflect.TypeOf(v))
	if len(childType) == 0 {
		// not a basic type
		return
	}

	if len(parts) == 1 {
		// only root field
		childName = parts[0]
		ok = true
		return
	}

	parentName = strings.Join(parts[:len(parts)-1], ".")
	childName = parts[len(parts)-1]
	ok = true
	return
}

func parentName(fullPath string) string {
	parts := strings.Split(fullPath, ".")
	if len(parts) < 2 {
		return ""
	}

	return strings.Join(parts[:len(parts)-1], ".")
}

var (
	hashType    = reflect.TypeOf(common.Hash{})
	addressType = reflect.TypeOf(common.Address{})
	bigType     = reflect.TypeOf(big.Int{})
)

// typToEth supports only basic types and arrays of basic types.
// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-712.md
func typToEth(typ reflect.Type) string {
	switch typ.Kind() {
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "bool"
	case reflect.Int:
		return "int64"
	case reflect.Int8:
		return "int8"
	case reflect.Int16:
		return "int16"
	case reflect.Int32:
		return "int32"
	case reflect.Int64:
		return "int64"
	case reflect.Uint:
		return "uint64"
	case reflect.Uint8:
		return "uint8"
	case reflect.Uint16:
		return "uint16"
	case reflect.Uint32:
		return "uint32"
	case reflect.Uint64:
		return "uint64"
	case reflect.Slice:
		ethName := typToEth(typ.Elem())
		if len(ethName) > 0 {
			return ethName + "[]"
		}
	case reflect.Ptr:
		if typ.Elem().ConvertibleTo(bigType) {
			return "string"
		}
	case reflect.Struct:
		if typ.ConvertibleTo(hashType) {
			return "string"
		} else if typ.ConvertibleTo(addressType) {
			return "string"
		}
	}

	return ""
}
