package types

import "github.com/ethereum/go-ethereum/common"

type NullableString string

// IsNil returns true if NullableString is uninitialized
func (s *NullableString) IsNil() bool {
	return s == nil
}

var emptyBytes = []byte{}

// Marshal implements the gogo proto custom type interface.
func (s *NullableString) Marshal() ([]byte, error) {
	if s == nil {
		return emptyBytes, nil
	}

	return []byte(*s), nil
}

// MarshalTo implements the gogo proto custom type interface.
func (s *NullableString) MarshalTo(data []byte) (n int, err error) {
	if s == nil {
		return 0, nil
	}

	if len(*s) == 0 {
		copy(data, []byte{0x30})
		return 1, nil
	}

	bz, err := s.Marshal()
	if err != nil {
		return 0, err
	}

	copy(data, bz)
	return len(bz), nil
}

// Unmarshal implements the gogo proto custom type interface.
func (s *NullableString) Unmarshal(data []byte) error {
	if len(data) == 0 {
		s = nil
		return nil
	}

	if s == nil {
		s = new(NullableString)
	}

	*s = NullableString(data)
	return nil
}

// Size implements the gogo proto custom type interface.
func (s *NullableString) Size() int {
	if s == nil {
		return 0
	}

	return len(*s)
}

func (s *NullableString) String() string {
	if s == nil {
		return ""
	}

	return common.Bytes2Hex([]byte(*s))
}
