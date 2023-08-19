package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"runtime/debug"
	"strings"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/pkg/errors"

	"github.com/InjectiveLabs/sdk-go/typeddata"
)

// WrapTxToEIP712 is an ultimate method that wraps Amino-encoded Cosmos Tx JSON data
// into an EIP712-compatible request. All messages must be of the same type.
func WrapTxToEIP712(
	cdc codectypes.AnyUnpacker,
	chainID uint64,
	msg cosmtypes.Msg,
	data []byte,
	feeDelegation *FeeDelegationOptions,
) (typeddata.TypedData, error) {
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

	msgTypes, err := extractMsgTypes(cdc, "MsgValue", msg)
	if err != nil {
		return typeddata.TypedData{}, err
	}

	if feeDelegation != nil {
		feeInfo := txData["fee"].(map[string]interface{})
		feeInfo["feePayer"] = feeDelegation.FeePayer.String()

		// also patching msgTypes to include feePayer
		msgTypes["Fee"] = []typeddata.Type{
			{Name: "feePayer", Type: "string"},
			{Name: "amount", Type: "Coin[]"},
			{Name: "gas", Type: "string"},
		}
	}

	var typedData = typeddata.TypedData{
		Types:       msgTypes,
		PrimaryType: "Tx",
		Domain:      domain,
		Message:     txData,
	}

	return typedData, nil
}

type FeeDelegationOptions struct {
	FeePayer cosmtypes.AccAddress
}

func extractMsgTypes(cdc codectypes.AnyUnpacker, msgTypeName string, msg cosmtypes.Msg) (typeddata.Types, error) {
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
			{Name: "timeout_height", Type: "string"},
			{Name: "tip", Type: "Tip"},
		},
		"Tip": {
			{Name: "amount", Type: "Coin"},
			{Name: "tipper", Type: "string"},
		},
		"Fee": {
			{Name: "amount", Type: "Coin[]"},
			{Name: "gas", Type: "string"},
			{Name: "granter", Type: "string"},
			{Name: "payer", Type: "string"},
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

	err := walkFields(cdc, rootTypes, msgTypeName, msg)
	if err != nil {
		return nil, err
	}

	return rootTypes, nil
}

const typeDefPrefix = "_"

func walkFields(cdc codectypes.AnyUnpacker, typeMap typeddata.Types, rootType string, in interface{}) (err error) {
	defer doRecover(&err)

	t := reflect.TypeOf(in)
	v := reflect.ValueOf(in)

	for {
		if t.Kind() == reflect.Ptr || t.Kind() == reflect.Interface {
			t = t.Elem()
			v = v.Elem()
			continue
		}
		break
	}

	err = traverseFields(cdc, typeMap, rootType, typeDefPrefix, t, v)
	return
}

type cosmosAnyWrapper struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

func traverseFields(
	cdc codectypes.AnyUnpacker,
	typeMap typeddata.Types,
	rootType string,
	prefix string,
	t reflect.Type,
	v reflect.Value,
) (err error) {
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
		fieldName := jsonNameFromTag(t.Field(i).Tag)

		if fieldType == cosmosAnyType {
			any := field.Interface().(*codectypes.Any)
			anyWrapper := &cosmosAnyWrapper{
				Type: any.TypeUrl,
			}

			err = cdc.UnpackAny(any, &anyWrapper.Value)
			if err != nil {
				err = errors.Wrap(err, "failed to unpack Any in msg struct")
				return
			}

			fieldType = reflect.TypeOf(anyWrapper)
			field = reflect.ValueOf(anyWrapper)

			// then continue as normal
		}

		for {
			if fieldType.Kind() == reflect.Ptr {
				fieldType = fieldType.Elem()

				if field.IsValid() {
					field = field.Elem()
				}

				continue
			}

			if fieldType.Kind() == reflect.Interface {
				fieldType = reflect.TypeOf(field.Interface())
				continue
			}

			if field.Kind() == reflect.Ptr {
				field = field.Elem()
				continue
			}

			break
		}

		var isCollection bool
		if fieldType.Kind() == reflect.Array || fieldType.Kind() == reflect.Slice {
			// if field.Len() == 0 {
			// 	// skip empty collections from type mapping
			// 	continue
			// }

			fieldType = fieldType.Elem()
			field = field.Index(0)
			isCollection = true
		}

		for {
			if fieldType.Kind() == reflect.Ptr {
				fieldType = fieldType.Elem()

				if field.IsValid() {
					field = field.Elem()
				}

				continue
			}

			if fieldType.Kind() == reflect.Interface {
				fieldType = reflect.TypeOf(field.Interface())
				continue
			}

			if field.Kind() == reflect.Ptr {
				field = field.Elem()
				continue
			}

			break
		}

		fieldPrefix := fmt.Sprintf("%s.%s", prefix, fieldName)
		ethTyp := typToEth(fieldType)
		if len(ethTyp) > 0 {
			if isCollection {
				ethTyp += "[]"
			}
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
			var fieldTypedef string
			if isCollection {
				fieldTypedef = sanitizeTypedef(fieldPrefix) + "[]"
			} else {
				fieldTypedef = sanitizeTypedef(fieldPrefix)
			}

			if prefix == typeDefPrefix {
				typeMap[rootType] = append(typeMap[rootType], typeddata.Type{
					Name: fieldName,
					Type: fieldTypedef,
				})
			} else {
				typeDef := sanitizeTypedef(prefix)
				typeMap[typeDef] = append(typeMap[typeDef], typeddata.Type{
					Name: fieldName,
					Type: fieldTypedef,
				})
			}

			err = traverseFields(cdc, typeMap, rootType, fieldPrefix, fieldType, field)
			if err != nil {
				return
			}

			continue
		}
	}

	return nil
}

func jsonNameFromTag(tag reflect.StructTag) string {
	jsonTags := tag.Get("json")
	parts := strings.Split(jsonTags, ",")
	return parts[0]
}

// _.foo_bar.baz -> TypeFooBarBaz
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
	hashType      = reflect.TypeOf(common.Hash{})
	addressType   = reflect.TypeOf(common.Address{})
	bigIntType    = reflect.TypeOf(big.Int{})
	cosmIntType   = reflect.TypeOf(cosmtypes.Int{})
	cosmosAnyType = reflect.TypeOf(&codectypes.Any{})
	timeType      = reflect.TypeOf(time.Time{})
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
	case reflect.Array:
		ethName := typToEth(typ.Elem())
		if len(ethName) > 0 {
			return ethName + "[]"
		}
	case reflect.Ptr:
		if typ.Elem().ConvertibleTo(bigIntType) ||
			typ.Elem().ConvertibleTo(timeType) ||
			typ.Elem().ConvertibleTo(cosmIntType) {
			return "string"
		}
	case reflect.Struct:
		if typ.ConvertibleTo(hashType) ||
			typ.ConvertibleTo(addressType) ||
			typ.ConvertibleTo(bigIntType) ||
			typ.ConvertibleTo(timeType) ||
			typ.ConvertibleTo(cosmIntType) {
			return "string"
		}
	}

	return ""
}

func doRecover(err *error) {
	if r := recover(); r != nil {
		debug.PrintStack()

		if e, ok := r.(error); ok {
			e = errors.Wrap(e, "panicked with error")
			*err = e
			return
		}

		*err = errors.Errorf("%v", r)
	}
}
