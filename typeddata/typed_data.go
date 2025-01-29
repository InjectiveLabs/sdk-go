package typeddata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"regexp"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

type SigFormat struct {
	Mime        string
	ByteVersion byte
}

var (
	IntendedValidator = SigFormat{
		accounts.MimetypeDataWithValidator,
		0x00,
	}
	DataTyped = SigFormat{
		accounts.MimetypeTypedData,
		0x01,
	}
	ApplicationClique = SigFormat{
		accounts.MimetypeClique,
		0x02,
	}
	TextPlain = SigFormat{
		accounts.MimetypeTextPlain,
		0x45,
	}
)

type ValidatorData struct {
	Address common.Address
	Message hexutil.Bytes
}

type TypedData struct {
	Types       Types            `json:"types"`
	PrimaryType string           `json:"primaryType"`
	Domain      TypedDataDomain  `json:"domain"`
	Message     TypedDataMessage `json:"message"`
}

type Type struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (t *Type) isArray() bool {
	return strings.HasSuffix(t.Type, "[]")
}

// typeName returns the canonical name of the type. If the type is 'Person[]', then
// this method returns 'Person'
func (t *Type) typeName() string {
	if strings.HasSuffix(t.Type, "[]") {
		return strings.TrimSuffix(t.Type, "[]")
	}
	return t.Type
}

func (t *Type) isReferenceType() bool {
	if t.Type == "" {
		return false
	}
	// Reference types must have a leading uppercase character
	firstLetter, _ := utf8.DecodeRuneInString(t.Type)
	return unicode.IsUpper(firstLetter)
}

type Types map[string][]Type

type TypePriority struct {
	Type  string
	Value uint
}

type TypedDataMessage = map[string]interface{}

type TypedDataDomain struct {
	Name              string                `json:"name"`
	Version           string                `json:"version"`
	ChainId           *math.HexOrDecimal256 `json:"chainId"`
	VerifyingContract string                `json:"verifyingContract"`
	Salt              string                `json:"salt"`
}

var typedDataReferenceTypeRegexp = regexp.MustCompile(`^[A-Z](\w*)(\[\])?$`)

// SignTextWithValidator signs the given message which can be further recovered
// with the given validator.
// hash = keccak256("\x19\x00"${address}${data}).
func SignTextValidator(validatorData ValidatorData) (signature hexutil.Bytes, message string) {
	msg := fmt.Sprintf("\x19\x00%s%s", string(validatorData.Address.Bytes()), string(validatorData.Message))
	return crypto.Keccak256([]byte(msg)), msg
}

// ComputeTypedDataAndHash computes the typed data and its keccak hash for signing
func ComputeTypedDataAndHash(typedData TypedData) (hash, data []byte, err error) {
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to pack and hash typedData EIP712Domain")
	}

	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to pack and hash typedData EIP712Domain")
	}

	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	return crypto.Keccak256(rawData), rawData, nil
}

// HashStruct generates a keccak256 hash of the encoding of the provided data
func (typedData *TypedData) HashStruct(primaryType string, data TypedDataMessage) (hexutil.Bytes, error) {
	encodedData, err := typedData.EncodeData(primaryType, data, 1)
	if err != nil {
		return nil, err
	}
	return crypto.Keccak256(encodedData), nil
}

// Dependencies returns an array of custom types ordered by their hierarchical reference tree
func (typedData *TypedData) Dependencies(primaryType string, found []string) []string {
	// BUGFIX (Geth): cut array suffixes so we can lookup the type
	primaryType = strings.TrimSuffix(primaryType, "[]")

	includes := func(arr []string, str string) bool {
		for _, obj := range arr {
			if obj == str {
				return true
			}
		}
		return false
	}

	if includes(found, primaryType) {
		return found
	}
	if typedData.Types[primaryType] == nil {
		return found
	}
	found = append(found, primaryType)
	for _, field := range typedData.Types[primaryType] {
		for _, dep := range typedData.Dependencies(field.Type, found) {
			if !includes(found, dep) {
				found = append(found, dep)
			}
		}
	}
	return found
}

// EncodeType generates the following encoding:
// `name ‖ "(" ‖ member₁ ‖ "," ‖ member₂ ‖ "," ‖ … ‖ memberₙ ")"`
//
// each member is written as `type ‖ " " ‖ name` encodings cascade down and are sorted by name
func (typedData *TypedData) EncodeType(primaryType string) hexutil.Bytes {
	// Get dependencies primary first, then alphabetical
	deps := typedData.Dependencies(primaryType, []string{})
	if len(deps) > 0 {
		slicedDeps := deps[1:]
		sort.Strings(slicedDeps)
		deps = append([]string{primaryType}, slicedDeps...)
	}

	// Format as a string with fields
	var buffer bytes.Buffer
	for _, dep := range deps {
		buffer.WriteString(dep)
		buffer.WriteString("(")
		for _, obj := range typedData.Types[dep] {
			buffer.WriteString(obj.Type)
			buffer.WriteString(" ")
			buffer.WriteString(obj.Name)
			buffer.WriteString(",")
		}
		buffer.Truncate(buffer.Len() - 1)
		buffer.WriteString(")")
	}

	return buffer.Bytes()
}

// TypeHash creates the keccak256 hash  of the data
func (typedData *TypedData) TypeHash(primaryType string) hexutil.Bytes {
	return crypto.Keccak256(typedData.EncodeType(primaryType))
}

// EncodeData generates the following encoding:
// `enc(value₁) ‖ enc(value₂) ‖ … ‖ enc(valueₙ)`
//
// each encoded member is 32-byte long
func (typedData *TypedData) EncodeData(primaryType string, data map[string]interface{}, depth int) (hexutil.Bytes, error) {
	if err := typedData.validate(); err != nil {
		return nil, err
	}

	buffer := bytes.Buffer{}

	// Verify extra data
	if exp, got := len(typedData.Types[primaryType]), len(data); exp < got {
		return nil, fmt.Errorf("there is extra data provided in the message (%d < %d)", exp, got)
	}

	// Add typehash
	buffer.Write(typedData.TypeHash(primaryType))

	// Add field contents. Structs and arrays have special handlers.
	for _, field := range typedData.Types[primaryType] {
		encType := field.Type
		encValue := data[field.Name]
		switch {
		case encType[len(encType)-1:] == "]":
			arrayValue, ok := encValue.([]interface{})
			if !ok {
				return nil, dataMismatchError(encType, encValue)
			}

			arrayBuffer := bytes.Buffer{}
			parsedType := strings.Split(encType, "[")[0]
			for _, item := range arrayValue {
				if typedData.Types[parsedType] != nil {
					mapValue, ok := item.(map[string]interface{})
					if !ok {
						return nil, dataMismatchError(parsedType, item)
					}
					encodedData, err := typedData.EncodeData(parsedType, mapValue, depth+1)
					if err != nil {
						return nil, err
					}

					// THIS PART IS CHANGED FROM Geth's
					//
					// Metamask [signTypedData_v4]: The behavior instead was to encode array values as the keccak256 of the
					// concatenated keccak256 of the values.
					//
					// See https://github.com/MetaMask/eth-sig-util/issues/106
					// And https://github.com/MetaMask/eth-sig-util/pull/107/files
					//
					// N.B.: Can be reverted to original, once the fix is rolled for everyone
					arrayBuffer.Write(crypto.Keccak256(encodedData))
				} else {
					bytesValue, err := typedData.EncodePrimitiveValue(parsedType, item, depth)
					if err != nil {
						return nil, err
					}
					arrayBuffer.Write(bytesValue)
				}
			}

			buffer.Write(crypto.Keccak256(arrayBuffer.Bytes()))
		case typedData.Types[field.Type] != nil:
			mapValue, ok := encValue.(map[string]interface{})
			if !ok {
				return nil, dataMismatchError(encType, encValue)
			}
			encodedData, err := typedData.EncodeData(field.Type, mapValue, depth+1)
			if err != nil {
				return nil, err
			}
			buffer.Write(crypto.Keccak256(encodedData))
		default:
			byteValue, err := typedData.EncodePrimitiveValue(encType, encValue, depth)
			if err != nil {
				return nil, err
			}
			buffer.Write(byteValue)
		}
	}

	return buffer.Bytes(), nil
}

// Attempt to parse bytes in different formats: byte array, hex string, hexutil.Bytes.
func parseBytes(encType interface{}) ([]byte, bool) {
	switch v := encType.(type) {
	case []byte:
		return v, true
	case hexutil.Bytes:
		return v, true
	case string:
		b, err := hexutil.Decode(v)
		if err != nil {
			return nil, false
		}
		return b, true
	default:
		return nil, false
	}
}

func parseInteger(encType string, encValue interface{}) (*big.Int, error) {
	var (
		length int
		signed = strings.HasPrefix(encType, "int")
		b      *big.Int
	)
	if encType == "int" || encType == "uint" {
		length = 256
	} else {
		lengthStr := ""
		if strings.HasPrefix(encType, "uint") {
			lengthStr = strings.TrimPrefix(encType, "uint")
		} else {
			lengthStr = strings.TrimPrefix(encType, "int")
		}
		atoiSize, err := strconv.Atoi(lengthStr)
		if err != nil {
			return nil, fmt.Errorf("invalid size on integer: %v", lengthStr)
		}
		length = atoiSize
	}
	switch v := encValue.(type) {
	case *math.HexOrDecimal256:
		b = (*big.Int)(v)
	case string:
		var hexIntValue math.HexOrDecimal256
		if err := hexIntValue.UnmarshalText([]byte(v)); err != nil {
			return nil, err
		}
		b = (*big.Int)(&hexIntValue)
	case float64:
		// JSON parses non-strings as float64. Fail if we cannot
		// convert it losslessly
		if float64(int64(v)) == v {
			b = big.NewInt(int64(v))
		} else {
			return nil, fmt.Errorf("invalid float value %v for type %v", v, encType)
		}
	}
	if b == nil {
		return nil, fmt.Errorf("invalid integer value %v/%v for type %v", encValue, reflect.TypeOf(encValue), encType)
	}
	if b.BitLen() > length {
		return nil, fmt.Errorf("integer larger than '%v'", encType)
	}
	if !signed && b.Sign() == -1 {
		return nil, fmt.Errorf("invalid negative value for unsigned type %v", encType)
	}
	return b, nil
}

// EncodePrimitiveValue deals with the primitive values found
// while searching through the typed data
func (typedData *TypedData) EncodePrimitiveValue(encType string, encValue interface{}, depth int) ([]byte, error) {
	switch encType {
	case "address":
		stringValue, ok := encValue.(string)
		if !ok || !common.IsHexAddress(stringValue) {
			return nil, dataMismatchError(encType, encValue)
		}
		retval := make([]byte, 32)
		copy(retval[12:], common.HexToAddress(stringValue).Bytes())
		return retval, nil
	case "bool":
		boolValue, ok := encValue.(bool)
		if !ok {
			return nil, dataMismatchError(encType, encValue)
		}
		if boolValue {
			return math.PaddedBigBytes(common.Big1, 32), nil
		}
		return math.PaddedBigBytes(common.Big0, 32), nil
	case "string":
		strVal, ok := encValue.(string)
		if !ok {
			return nil, dataMismatchError(encType, encValue)
		}
		return crypto.Keccak256([]byte(strVal)), nil
	case "bytes":
		bytesValue, ok := parseBytes(encValue)
		if !ok {
			return nil, dataMismatchError(encType, encValue)
		}
		return crypto.Keccak256(bytesValue), nil
	}
	if strings.HasPrefix(encType, "bytes") {
		lengthStr := strings.TrimPrefix(encType, "bytes")
		length, err := strconv.Atoi(lengthStr)
		if err != nil {
			return nil, fmt.Errorf("invalid size on bytes: %v", lengthStr)
		}
		if length < 0 || length > 32 {
			return nil, fmt.Errorf("invalid size on bytes: %d", length)
		}
		if byteValue, ok := parseBytes(encValue); !ok || len(byteValue) != length {
			return nil, dataMismatchError(encType, encValue)
		} else {
			// Right-pad the bits
			dst := make([]byte, 32)
			copy(dst, byteValue)
			return dst, nil
		}
	}
	if strings.HasPrefix(encType, "int") || strings.HasPrefix(encType, "uint") {
		b, err := parseInteger(encType, encValue)
		if err != nil {
			return nil, err
		}
		return math.U256Bytes(b), nil
	}
	return nil, fmt.Errorf("unrecognized type '%s'", encType)

}

// dataMismatchError generates an error for a mismatch between
// the provided type and data
func dataMismatchError(encType string, encValue interface{}) error {
	return fmt.Errorf("provided data '%v' doesn't match type '%s'", encValue, encType)
}

// UnmarshalValidatorData converts the bytes input to typed data
func UnmarshalValidatorData(data interface{}) (ValidatorData, error) {
	raw, ok := data.(map[string]interface{})
	if !ok {
		return ValidatorData{}, errors.New("validator input is not a map[string]interface{}")
	}
	addr, ok := raw["address"].(string)
	if !ok {
		return ValidatorData{}, errors.New("validator address is not sent as a string")
	}
	addrBytes, err := hexutil.Decode(addr)
	if err != nil {
		return ValidatorData{}, err
	}
	if !ok || len(addrBytes) == 0 {
		return ValidatorData{}, errors.New("validator address is undefined")
	}

	message, ok := raw["message"].(string)
	if !ok {
		return ValidatorData{}, errors.New("message is not sent as a string")
	}
	messageBytes, err := hexutil.Decode(message)
	if err != nil {
		return ValidatorData{}, err
	}
	if !ok || len(messageBytes) == 0 {
		return ValidatorData{}, errors.New("message is undefined")
	}

	return ValidatorData{
		Address: common.BytesToAddress(addrBytes),
		Message: messageBytes,
	}, nil
}

// validate makes sure the types are sound
func (typedData *TypedData) validate() error {
	if err := typedData.Types.validate(); err != nil {
		return err
	}
	if err := typedData.Domain.validate(); err != nil {
		return err
	}
	return nil
}

// Map generates a map version of the typed data
func (typedData *TypedData) Map() map[string]interface{} {
	dataMap := map[string]interface{}{
		"types":       typedData.Types,
		"domain":      typedData.Domain.Map(),
		"primaryType": typedData.PrimaryType,
		"message":     typedData.Message,
	}
	return dataMap
}

// Format returns a representation of typedData, which can be easily displayed by a user-interface
// without in-depth knowledge about 712 rules
func (typedData *TypedData) Format() ([]*NameValueType, error) {
	domain, err := typedData.formatData("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return nil, err
	}
	ptype, err := typedData.formatData(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return nil, err
	}
	var nvts []*NameValueType
	nvts = append(nvts, &NameValueType{
		Name:  "EIP712Domain",
		Value: domain,
		Typ:   "domain",
	}, &NameValueType{
		Name:  typedData.PrimaryType,
		Value: ptype,
		Typ:   "primary type",
	})
	return nvts, nil
}

func (typedData *TypedData) formatData(primaryType string, data map[string]interface{}) ([]*NameValueType, error) {
	output := []*NameValueType{}

	// Add field contents. Structs and arrays have special handlers.
	for _, field := range typedData.Types[primaryType] {
		encName := field.Name
		encValue := data[encName]
		item := &NameValueType{
			Name: encName,
			Typ:  field.Type,
		}
		switch {
		case field.isArray():
			arrayValue, _ := encValue.([]interface{})
			parsedType := field.typeName()
			for _, v := range arrayValue {
				if typedData.Types[parsedType] != nil {
					mapValue, _ := v.(map[string]interface{})
					mapOutput, err := typedData.formatData(parsedType, mapValue)
					if err != nil {
						return nil, err
					}
					item.Value = mapOutput
				} else {
					primitiveOutput, err := formatPrimitiveValue(field.Type, encValue)
					if err != nil {
						return nil, err
					}
					item.Value = primitiveOutput
				}
			}
		case typedData.Types[field.Type] != nil:
			if mapValue, ok := encValue.(map[string]interface{}); ok {
				mapOutput, err := typedData.formatData(field.Type, mapValue)
				if err != nil {
					return nil, err
				}
				item.Value = mapOutput
			} else {
				item.Value = "<nil>"
			}
		default:
			primitiveOutput, err := formatPrimitiveValue(field.Type, encValue)
			if err != nil {
				return nil, err
			}
			item.Value = primitiveOutput
		}
		output = append(output, item)
	}
	return output, nil
}

func formatPrimitiveValue(encType string, encValue interface{}) (string, error) {
	switch encType {
	case "address":
		if stringValue, ok := encValue.(string); !ok {
			return "", fmt.Errorf("could not format value %v as address", encValue)
		} else {
			return common.HexToAddress(stringValue).String(), nil
		}
	case "bool":
		if boolValue, ok := encValue.(bool); !ok {
			return "", fmt.Errorf("could not format value %v as bool", encValue)
		} else {
			return fmt.Sprintf("%t", boolValue), nil
		}
	case "bytes", "string":
		return fmt.Sprintf("%s", encValue), nil
	}
	if strings.HasPrefix(encType, "bytes") {
		return fmt.Sprintf("%s", encValue), nil

	}
	if strings.HasPrefix(encType, "uint") || strings.HasPrefix(encType, "int") {
		if b, err := parseInteger(encType, encValue); err != nil {
			return "", err
		} else {
			return fmt.Sprintf("%d (0x%x)", b, b), nil
		}
	}
	return "", fmt.Errorf("unhandled type %v", encType)
}

// NameValueType is a very simple struct with Name, Value and Type. It's meant for simple
// json structures used to communicate signing-info about typed data with the UI
type NameValueType struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
	Typ   string      `json:"type"`
}

// Pprint returns a pretty-printed version of nvt
func (nvt *NameValueType) Pprint(depth int) string {
	output := bytes.Buffer{}
	output.WriteString(strings.Repeat("\u00a0", depth*2))
	output.WriteString(fmt.Sprintf("%s [%s]: ", nvt.Name, nvt.Typ))
	if nvts, ok := nvt.Value.([]*NameValueType); ok {
		output.WriteString("\n")
		for _, next := range nvts {
			sublevel := next.Pprint(depth + 1)
			output.WriteString(sublevel)
		}
	} else {
		if nvt.Value != nil {
			output.WriteString(fmt.Sprintf("%q\n", nvt.Value))
		} else {
			output.WriteString("\n")
		}
	}
	return output.String()
}

// Validate checks if the types object is conformant to the specs
func (t Types) validate() error {
	for typeKey, typeArr := range t {
		if typeKey == "" {
			return fmt.Errorf("empty type key")
		}
		for i, typeObj := range typeArr {
			if typeObj.Type == "" {
				return fmt.Errorf("type %q:%d: empty Type", typeKey, i)
			}
			if typeObj.Name == "" {
				return fmt.Errorf("type %q:%d: empty Name", typeKey, i)
			}
			if typeKey == typeObj.Type {
				return fmt.Errorf("type %q cannot reference itself", typeObj.Type)
			}
			if typeObj.isReferenceType() {
				if _, exist := t[typeObj.typeName()]; !exist {
					return fmt.Errorf("reference type %q is undefined", typeObj.Type)
				}
				if !typedDataReferenceTypeRegexp.MatchString(typeObj.Type) {
					return fmt.Errorf("unknown reference type %q", typeObj.Type)
				}
			} else if !isPrimitiveTypeValid(typeObj.Type) {
				return fmt.Errorf("unknown type %q", typeObj.Type)
			}
		}
	}
	return nil
}

// Checks if the primitive value is valid
func isPrimitiveTypeValid(primitiveType string) bool {
	if primitiveType == "address" ||
		primitiveType == "address[]" ||
		primitiveType == "bool" ||
		primitiveType == "bool[]" ||
		primitiveType == "string" ||
		primitiveType == "string[]" {
		return true
	}
	if primitiveType == "bytes" ||
		primitiveType == "bytes[]" ||
		primitiveType == "bytes1" ||
		primitiveType == "bytes1[]" ||
		primitiveType == "bytes2" ||
		primitiveType == "bytes2[]" ||
		primitiveType == "bytes3" ||
		primitiveType == "bytes3[]" ||
		primitiveType == "bytes4" ||
		primitiveType == "bytes4[]" ||
		primitiveType == "bytes5" ||
		primitiveType == "bytes5[]" ||
		primitiveType == "bytes6" ||
		primitiveType == "bytes6[]" ||
		primitiveType == "bytes7" ||
		primitiveType == "bytes7[]" ||
		primitiveType == "bytes8" ||
		primitiveType == "bytes8[]" ||
		primitiveType == "bytes9" ||
		primitiveType == "bytes9[]" ||
		primitiveType == "bytes10" ||
		primitiveType == "bytes10[]" ||
		primitiveType == "bytes11" ||
		primitiveType == "bytes11[]" ||
		primitiveType == "bytes12" ||
		primitiveType == "bytes12[]" ||
		primitiveType == "bytes13" ||
		primitiveType == "bytes13[]" ||
		primitiveType == "bytes14" ||
		primitiveType == "bytes14[]" ||
		primitiveType == "bytes15" ||
		primitiveType == "bytes15[]" ||
		primitiveType == "bytes16" ||
		primitiveType == "bytes16[]" ||
		primitiveType == "bytes17" ||
		primitiveType == "bytes17[]" ||
		primitiveType == "bytes18" ||
		primitiveType == "bytes18[]" ||
		primitiveType == "bytes19" ||
		primitiveType == "bytes19[]" ||
		primitiveType == "bytes20" ||
		primitiveType == "bytes20[]" ||
		primitiveType == "bytes21" ||
		primitiveType == "bytes21[]" ||
		primitiveType == "bytes22" ||
		primitiveType == "bytes22[]" ||
		primitiveType == "bytes23" ||
		primitiveType == "bytes23[]" ||
		primitiveType == "bytes24" ||
		primitiveType == "bytes24[]" ||
		primitiveType == "bytes25" ||
		primitiveType == "bytes25[]" ||
		primitiveType == "bytes26" ||
		primitiveType == "bytes26[]" ||
		primitiveType == "bytes27" ||
		primitiveType == "bytes27[]" ||
		primitiveType == "bytes28" ||
		primitiveType == "bytes28[]" ||
		primitiveType == "bytes29" ||
		primitiveType == "bytes29[]" ||
		primitiveType == "bytes30" ||
		primitiveType == "bytes30[]" ||
		primitiveType == "bytes31" ||
		primitiveType == "bytes31[]" ||
		primitiveType == "bytes32" ||
		primitiveType == "bytes32[]" {
		return true
	}
	if primitiveType == "int" ||
		primitiveType == "int[]" ||
		primitiveType == "int8" ||
		primitiveType == "int8[]" ||
		primitiveType == "int16" ||
		primitiveType == "int16[]" ||
		primitiveType == "int32" ||
		primitiveType == "int32[]" ||
		primitiveType == "int64" ||
		primitiveType == "int64[]" ||
		primitiveType == "int128" ||
		primitiveType == "int128[]" ||
		primitiveType == "int256" ||
		primitiveType == "int256[]" {
		return true
	}
	if primitiveType == "uint" ||
		primitiveType == "uint[]" ||
		primitiveType == "uint8" ||
		primitiveType == "uint8[]" ||
		primitiveType == "uint16" ||
		primitiveType == "uint16[]" ||
		primitiveType == "uint32" ||
		primitiveType == "uint32[]" ||
		primitiveType == "uint64" ||
		primitiveType == "uint64[]" ||
		primitiveType == "uint128" ||
		primitiveType == "uint128[]" ||
		primitiveType == "uint256" ||
		primitiveType == "uint256[]" {
		return true
	}
	return false
}

// validate checks if the given domain is valid, i.e. contains at least
// the minimum viable keys and values
func (domain *TypedDataDomain) validate() error {
	if domain.ChainId == nil && domain.Name == "" && domain.Version == "" && domain.VerifyingContract == "" && domain.Salt == "" {
		return errors.New("domain is undefined")
	}

	return nil
}

// Map is a helper function to generate a map version of the domain
func (domain *TypedDataDomain) Map() map[string]interface{} {
	dataMap := map[string]interface{}{}

	if domain.ChainId != nil {
		dataMap["chainId"] = domain.ChainId
	}

	if domain.Name != "" {
		dataMap["name"] = domain.Name
	}

	if domain.Version != "" {
		dataMap["version"] = domain.Version
	}

	if domain.VerifyingContract != "" {
		dataMap["verifyingContract"] = domain.VerifyingContract
	}

	if domain.Salt != "" {
		dataMap["salt"] = domain.Salt
	}
	return dataMap
}

var (
	LegacyAminoCodec *codec.LegacyAmino
	ProtoCodec       *codec.ProtoCodec
)

func SetCodec(amino *codec.LegacyAmino, proto *codec.ProtoCodec) {
	LegacyAminoCodec, ProtoCodec = amino, proto
}

func GetEIP712TypedDataForMsg(signDocBytes []byte) (TypedData, error) {
	txData := make(map[string]interface{})
	if err := json.Unmarshal(signDocBytes, &txData); err != nil {
		return TypedData{}, errors.Wrap(err, "failed to unmarshal signDocBytes")
	}

	// parse txData to StdSignDoc in order to parse the msg for ExtractMsgTypes
	var signDoc legacytx.StdSignDoc
	if err := LegacyAminoCodec.UnmarshalJSON(signDocBytes, &signDoc); err != nil {
		return TypedData{}, err
	}

	var msg cosmtypes.Msg
	if err := LegacyAminoCodec.UnmarshalJSON(signDoc.Msgs[0], &msg); err != nil {
		return TypedData{}, errors.Wrap(err, "failed to unmarshal msg")
	}

	// For some reason this type cast does not work during UnmarshalJSON above
	// If the underlying msg does implement the UnpackInterfacesMessage interface (MsgGrant, MsgExec...),
	// we explicitly call the method here to ensure potential Any fields within the message are correctly parsed
	if unpacker, ok := msg.(codectypes.UnpackInterfacesMessage); ok {
		if err := unpacker.UnpackInterfaces(codectypes.AminoJSONUnpacker{Cdc: LegacyAminoCodec.Amino}); err != nil {
			return TypedData{}, errors.Wrap(err, "failed to unpack msg")
		}
	}

	msgTypes, err := ExtractMsgTypes(ProtoCodec, "MsgValue", msg)
	if err != nil {
		return TypedData{}, errors.Wrap(err, "failed to extract msg types")
	}

	if txData["fee"] != nil {
		msgTypes["Fee"] = []Type{
			{Name: "amount", Type: "Coin[]"},
			{Name: "gas", Type: "string"},
		}
	}

	// set timeout_height to 0 in case the user forgot to provide the flag
	if txData["timeout_height"] == nil {
		txData["timeout_height"] = "0"
	}

	chainID, err := ParseCosmosChainID(txData["chain_id"].(string))
	if err != nil {
		return TypedData{}, err
	}

	// see VerifySignatureEIP712 func and its handling of chain id
	switch chainID.Int64() {
	case 777, 888:
		chainID = big.NewInt(11155111) // Sepolia
	default:
		chainID = big.NewInt(1)
	}

	domain := TypedDataDomain{
		Name:              "Injective Web3",
		Version:           "1.0.0",
		ChainId:           math.NewHexOrDecimal256(chainID.Int64()),
		VerifyingContract: "cosmos",
		Salt:              "0",
	}

	td := TypedData{
		Types:       msgTypes,
		PrimaryType: "Tx",
		Domain:      domain,
		Message:     txData,
	}

	return td, nil
}

func ExtractMsgTypes(cdc codec.ProtoCodecMarshaler, msgTypeName string, msg cosmtypes.Msg) (Types, error) {
	rootTypes := Types{
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

	err := walkFields(cdc, rootTypes, msgTypeName, msg)
	if err != nil {
		return nil, err
	}

	return rootTypes, nil
}

func walkFields(cdc codec.ProtoCodecMarshaler, typeMap Types, rootType string, in interface{}) (err error) {
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

//nolint:gocritic // this is a handy way to return err in deferred funcs
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

const typeDefPrefix = "_"

func traverseFields(
	cdc codec.ProtoCodecMarshaler,
	typeMap Types,
	rootType string,
	prefix string,
	t reflect.Type,
	v reflect.Value,
) (err error) {
	n := t.NumField()

	if prefix == typeDefPrefix {
		if len(typeMap[rootType]) == n {
			return nil
		}
	} else {
		typeDef := sanitizeTypedef(prefix)
		if len(typeMap[typeDef]) == n {
			return nil
		}
	}

	for i := 0; i < n; i++ {
		var field reflect.Value
		if v.IsValid() {
			field = v.Field(i)
		}

		fieldType := t.Field(i).Type
		fieldName := jsonNameFromTag(t.Field(i).Tag)
		var isCollection bool
		if fieldType.Kind() == reflect.Array || fieldType.Kind() == reflect.Slice {
			if field.Len() == 0 {
				// skip empty collections from type mapping
				continue
			}

			fieldType = fieldType.Elem()
			field = field.Index(0)
			isCollection = true
		}

		if fieldType == cosmosAnyType {
			any := field.Interface().(*codectypes.Any)
			anyWrapper := &cosmosAnyWrapper{
				Type: any.TypeUrl,
			}

			err = cdc.UnpackAny(any, &anyWrapper.Value)
			if err != nil {
				err = errors.Wrap(err, "failed to unpack Any in msg struct")
				return err
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
		if ethTyp != "" {
			if isCollection {
				ethTyp += "[]"
			}

			if field.Kind() == reflect.String && field.Len() == 0 {
				// skip empty strings from type mapping
				continue
			}

			if prefix == typeDefPrefix {
				typeMap[rootType] = append(typeMap[rootType], Type{
					Name: fieldName,
					Type: ethTyp,
				})
			} else {
				typeDef := sanitizeTypedef(prefix)
				typeMap[typeDef] = append(typeMap[typeDef], Type{
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
				typeMap[rootType] = append(typeMap[rootType], Type{
					Name: fieldName,
					Type: fieldTypedef,
				})
			} else {
				typeDef := sanitizeTypedef(prefix)
				typeMap[typeDef] = append(typeMap[typeDef], Type{
					Name: fieldName,
					Type: fieldTypedef,
				})
			}

			err = traverseFields(cdc, typeMap, rootType, fieldPrefix, fieldType, field)
			if err != nil {
				return errors.WithStack(err)
			}

			continue
		}
	}

	return nil
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
			buf.WriteString(strings.Title(subpart)) //nolint // strings is used for compat
		}
	}

	return buf.String()
}

func jsonNameFromTag(tag reflect.StructTag) string {
	jsonTags := tag.Get("json")
	parts := strings.Split(jsonTags, ",")
	return parts[0]
}

var (
	hashType      = reflect.TypeOf(common.Hash{})
	addressType   = reflect.TypeOf(common.Address{})
	bigIntType    = reflect.TypeOf(big.Int{})
	cosmIntType   = reflect.TypeOf(sdkmath.Int{})
	cosmosAnyType = reflect.TypeOf(&codectypes.Any{})
	timeType      = reflect.TypeOf(time.Time{})
)

type cosmosAnyWrapper struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

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
		if ethName != "" {
			return ethName + "[]"
		}
	case reflect.Array:
		ethName := typToEth(typ.Elem())
		if ethName != "" {
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

func SignableTypes() Types {
	return Types{
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
				Type: "address",
			},
			{
				Name: "salt",
				Type: "string",
			},
		},
		"Tx": {
			{Name: "context", Type: "string"},
			{Name: "msgs", Type: "string"},
		},
	}
}

var (
	ErrInvalidChainID = errors.New("invalid chain-id")

	regexChainID   = `[a-z]*`
	regexSeparator = `-{1}`
	regexEpoch     = `[1-9][0-9]*`

	cosmosChainID = regexp.MustCompile(fmt.Sprintf(`^(%s)%s(%s)$`, regexChainID, regexSeparator, regexEpoch))
)

// ParseCosmosChainID parses a string chain identifier's epoch to an Ethereum-compatible
// chain-id in *big.Int format. The function returns an error if the chain-id has an invalid format
func ParseCosmosChainID(chainID string) (*big.Int, error) {
	chainID = strings.TrimSpace(chainID)
	if len(chainID) > 48 {
		return nil, errors.Wrapf(ErrInvalidChainID, "chain-id '%s' cannot exceed 48 chars", chainID)
	}

	matches := cosmosChainID.FindStringSubmatch(chainID)
	if matches == nil || len(matches) != 3 || matches[1] == "" {
		return nil, errors.Wrap(ErrInvalidChainID, chainID)
	}

	// verify that the chain-id entered is a base 10 integer
	chainIDInt, ok := new(big.Int).SetString(matches[2], 10)
	if !ok {
		return nil, errors.Wrapf(ErrInvalidChainID, "epoch %s must be base-10 integer format", matches[2])
	}

	return chainIDInt, nil
}
