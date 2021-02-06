package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"strings"

	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/pkg/errors"

	"github.com/InjectiveLabs/sdk-go/typeddata"
)

// WrapTxToEIP712 is an ultimate method that wraps Amino-encoded Cosmos Tx JSON data
// into an EIP712-compatible request. All messages must be of the same type.
func WrapTxToEIP712(chainID uint64, msg cosmtypes.Msg, data []byte) (typeddata.TypedData, error) {
	txData := make(map[string]interface{})
	if err := json.Unmarshal(data, &txData); err != nil {
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
		Types:       extractMsgTypes("MsgValue", msg),
		PrimaryType: "Tx",
		Domain:      domain,
		Message:     txData,
	}

	return typedData, nil
}

func extractMsgTypes(msgTypeName string, msg cosmtypes.Msg) typeddata.Types {
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
			{Name: "value", Type: msgTypeName},
		},
		msgTypeName: {},
	}

	walkFields(rootTypes, msgTypeName, msg)

	return rootTypes
}

const typeDefPrefix = "_"

func walkFields(typeMap typeddata.Types, rootType string, in interface{}) {
	defer func() {
		if x := recover(); x != nil {
			return
		}
	}()

	t := reflect.TypeOf(in)
	v := reflect.ValueOf(in)

	for {
		if t.Kind() == reflect.Ptr ||
			t.Kind() == reflect.Interface {
			t = t.Elem()
			v = v.Elem()

			continue
		}

		break
	}

	traverseFields(typeMap, rootType, typeDefPrefix, t, v)
}

func traverseFields(typeMap typeddata.Types, rootType string, prefix string, t reflect.Type, v reflect.Value) {
	n := t.NumField()

	if prefix == typeDefPrefix {
		if len(typeMap[rootType]) == n {
			return
		}
	} else {
		typeDef := sanitizeTypedef(prefix)
		if len(typeMap[typeDef]) == n {
			return
		}
	}

	for i := 0; i < n; i++ {
		var field reflect.Value
		if v.IsValid() {
			field = v.Field(i)
		}

		fieldType := t.Field(i).Type

		for {
			if fieldType.Kind() == reflect.Ptr || fieldType.Kind() == reflect.Interface {
				fieldType = fieldType.Elem()

				if field.IsValid() {
					field = field.Elem()
				}
				continue
			}
			break
		}

		if fieldType.Kind() == reflect.Array {
			fieldType = fieldType.Elem()
			field = field.Index(0)
		}

		for {
			if fieldType.Kind() == reflect.Ptr || fieldType.Kind() == reflect.Interface {
				fieldType = fieldType.Elem()

				if field.IsValid() {
					field = field.Elem()
				}
				continue
			}
			break
		}

		fieldName := jsonNameFromTag(t.Field(i).Tag)
		fieldPrefix := fmt.Sprintf("%s.%s", prefix, fieldName)

		ethTyp := typToEth(fieldType)
		if len(ethTyp) > 0 {
			if prefix == typeDefPrefix {
				typeMap[rootType] = append(typeMap[rootType], typeddata.Type{
					Name: fieldName,
					Type: ethTyp,
				})
			} else {
				typeDef := sanitizeTypedef(prefix)
				typeMap[typeDef] = append(typeMap[typeDef], typeddata.Type{
					Name: fieldName,
					Type: ethTyp,
				})
			}

			continue
		}

		if fieldType.Kind() == reflect.Struct {
			if prefix == typeDefPrefix {
				typeMap[rootType] = append(typeMap[rootType], typeddata.Type{
					Name: fieldName,
					Type: sanitizeTypedef(fieldPrefix),
				})
			} else {
				typeDef := sanitizeTypedef(prefix)
				typeMap[typeDef] = append(typeMap[typeDef], typeddata.Type{
					Name: fieldName,
					Type: sanitizeTypedef(fieldPrefix),
				})
			}

			traverseFields(typeMap, rootType, fieldPrefix, fieldType, field)
			continue
		}
	}
}

func jsonNameFromTag(tag reflect.StructTag) string {
	jsonTags := tag.Get("json")
	parts := strings.Split(jsonTags, ",")
	return parts[0]
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

var (
	hashType    = reflect.TypeOf(common.Hash{})
	addressType = reflect.TypeOf(common.Address{})
	bigIntType  = reflect.TypeOf(big.Int{})
	cosmIntType = reflect.TypeOf(cosmtypes.Int{})
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
		if typ.Elem().ConvertibleTo(bigIntType) ||
			typ.Elem().ConvertibleTo(cosmIntType) {
			return "string"
		}
	case reflect.Struct:
		if typ.ConvertibleTo(hashType) ||
			typ.ConvertibleTo(addressType) ||
			typ.ConvertibleTo(bigIntType) ||
			typ.ConvertibleTo(cosmIntType) {
			return "string"
		}
	}

	return ""
}
