// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: injective/evm/v1/tx_result.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// TxResult stores results of Tx execution.
type TxResult struct {
	// contract_address contains the ethereum address of the created contract (if
	// any). If the state transition is an evm.Call, the contract address will be
	// empty.
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty" yaml:"contract_address"`
	// bloom represents the bloom filter bytes
	Bloom []byte `protobuf:"bytes,2,opt,name=bloom,proto3" json:"bloom,omitempty"`
	// tx_logs contains the transaction hash and the proto-compatible ethereum
	// logs.
	TxLogs TransactionLogs `protobuf:"bytes,3,opt,name=tx_logs,json=txLogs,proto3" json:"tx_logs" yaml:"tx_logs"`
	// ret defines the bytes from the execution.
	Ret []byte `protobuf:"bytes,4,opt,name=ret,proto3" json:"ret,omitempty"`
	// reverted flag is set to true when the call has been reverted
	Reverted bool `protobuf:"varint,5,opt,name=reverted,proto3" json:"reverted,omitempty"`
	// gas_used notes the amount of gas consumed while execution
	GasUsed uint64 `protobuf:"varint,6,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
}

func (m *TxResult) Reset()         { *m = TxResult{} }
func (m *TxResult) String() string { return proto.CompactTextString(m) }
func (*TxResult) ProtoMessage()    {}
func (*TxResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_6868197df3a51cef, []int{0}
}
func (m *TxResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxResult.Merge(m, src)
}
func (m *TxResult) XXX_Size() int {
	return m.Size()
}
func (m *TxResult) XXX_DiscardUnknown() {
	xxx_messageInfo_TxResult.DiscardUnknown(m)
}

var xxx_messageInfo_TxResult proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TxResult)(nil), "injective.evm.v1.TxResult")
}

func init() { proto.RegisterFile("injective/evm/v1/tx_result.proto", fileDescriptor_6868197df3a51cef) }

var fileDescriptor_6868197df3a51cef = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x3f, 0x6e, 0xa3, 0x40,
	0x14, 0xc6, 0x19, 0xff, 0x65, 0x67, 0x57, 0xbb, 0x16, 0xb2, 0x76, 0x59, 0xaf, 0x04, 0x2c, 0xcd,
	0xd2, 0x2c, 0xc8, 0x49, 0xe7, 0x2e, 0x2e, 0x22, 0x59, 0x72, 0x85, 0x9c, 0x26, 0x0d, 0x1a, 0xe0,
	0x09, 0x13, 0x01, 0x63, 0xcd, 0x0c, 0x08, 0xdf, 0x20, 0x52, 0x9a, 0x1c, 0x21, 0xc7, 0x71, 0xe9,
	0x32, 0x95, 0x15, 0xd9, 0x37, 0xf0, 0x09, 0x22, 0x43, 0xec, 0x44, 0x4e, 0xf7, 0xbe, 0xf7, 0xbe,
	0x99, 0xdf, 0xfb, 0x66, 0xb0, 0x11, 0x67, 0x77, 0x10, 0x88, 0xb8, 0x00, 0x07, 0x8a, 0xd4, 0x29,
	0x86, 0x8e, 0x28, 0x3d, 0x06, 0x3c, 0x4f, 0x84, 0xbd, 0x60, 0x54, 0x50, 0xa5, 0x77, 0x72, 0xd8,
	0x50, 0xa4, 0x76, 0x31, 0x1c, 0xf4, 0x23, 0x1a, 0xd1, 0x6a, 0xe8, 0x1c, 0xaa, 0xda, 0x37, 0xf8,
	0xf7, 0xf9, 0x26, 0x46, 0x32, 0x4e, 0x02, 0x11, 0xd3, 0xcc, 0x4b, 0x68, 0xc4, 0x6b, 0xa3, 0xf9,
	0xd0, 0xc0, 0xf2, 0xac, 0x74, 0x2b, 0x86, 0x72, 0x8d, 0x7b, 0x01, 0xcd, 0x04, 0x23, 0x81, 0xf0,
	0x48, 0x18, 0x32, 0xe0, 0x5c, 0x45, 0x06, 0xb2, 0xbe, 0x8c, 0xff, 0xec, 0x37, 0xfa, 0xaf, 0x25,
	0x49, 0x93, 0x91, 0x79, 0xee, 0x30, 0xdd, 0x1f, 0xc7, 0xd6, 0x55, 0xdd, 0x51, 0xfa, 0xb8, 0xed,
	0x27, 0x94, 0xa6, 0x6a, 0xc3, 0x40, 0xd6, 0x37, 0xb7, 0x16, 0x8a, 0x8b, 0xbb, 0xa2, 0xac, 0xd8,
	0x6a, 0xd3, 0x40, 0xd6, 0xd7, 0x8b, 0xbf, 0xf6, 0x79, 0x1a, 0x7b, 0xf6, 0xbe, 0xe5, 0x94, 0x46,
	0x7c, 0xfc, 0x73, 0xb5, 0xd1, 0xa5, 0xfd, 0x46, 0xff, 0x5e, 0xb3, 0xdf, 0xce, 0x9b, 0x6e, 0x47,
	0x94, 0x87, 0xb9, 0xd2, 0xc3, 0x4d, 0x06, 0x42, 0x6d, 0x55, 0x9c, 0x43, 0xa9, 0x0c, 0xb0, 0xcc,
	0xa0, 0x00, 0x26, 0x20, 0x54, 0xdb, 0x06, 0xb2, 0x64, 0xf7, 0xa4, 0x95, 0xdf, 0x58, 0x8e, 0x08,
	0xf7, 0x72, 0x0e, 0xa1, 0xda, 0x31, 0x90, 0xd5, 0x72, 0xbb, 0x11, 0xe1, 0x37, 0x1c, 0xc2, 0x51,
	0xeb, 0xfe, 0x49, 0x97, 0xc6, 0xc1, 0x6a, 0xab, 0xa1, 0xf5, 0x56, 0x43, 0x2f, 0x5b, 0x0d, 0x3d,
	0xee, 0x34, 0x69, 0xbd, 0xd3, 0xa4, 0xe7, 0x9d, 0x26, 0xdd, 0x4e, 0xa2, 0x58, 0xcc, 0x73, 0xdf,
	0x0e, 0x68, 0xea, 0x4c, 0x8e, 0x5b, 0x4f, 0x89, 0xcf, 0x9d, 0x53, 0x86, 0xff, 0x01, 0x65, 0xf0,
	0x51, 0xce, 0x49, 0x9c, 0x39, 0x29, 0x0d, 0xf3, 0x04, 0x78, 0xf5, 0x0d, 0x62, 0xb9, 0x00, 0xee,
	0x77, 0xaa, 0x97, 0xbf, 0x7c, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x68, 0xf8, 0x65, 0xee, 0x01,
	0x00, 0x00,
}

func (m *TxResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GasUsed != 0 {
		i = encodeVarintTxResult(dAtA, i, uint64(m.GasUsed))
		i--
		dAtA[i] = 0x30
	}
	if m.Reverted {
		i--
		if m.Reverted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.Ret) > 0 {
		i -= len(m.Ret)
		copy(dAtA[i:], m.Ret)
		i = encodeVarintTxResult(dAtA, i, uint64(len(m.Ret)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.TxLogs.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTxResult(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Bloom) > 0 {
		i -= len(m.Bloom)
		copy(dAtA[i:], m.Bloom)
		i = encodeVarintTxResult(dAtA, i, uint64(len(m.Bloom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintTxResult(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTxResult(dAtA []byte, offset int, v uint64) int {
	offset -= sovTxResult(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TxResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovTxResult(uint64(l))
	}
	l = len(m.Bloom)
	if l > 0 {
		n += 1 + l + sovTxResult(uint64(l))
	}
	l = m.TxLogs.Size()
	n += 1 + l + sovTxResult(uint64(l))
	l = len(m.Ret)
	if l > 0 {
		n += 1 + l + sovTxResult(uint64(l))
	}
	if m.Reverted {
		n += 2
	}
	if m.GasUsed != 0 {
		n += 1 + sovTxResult(uint64(m.GasUsed))
	}
	return n
}

func sovTxResult(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTxResult(x uint64) (n int) {
	return sovTxResult(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TxResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTxResult
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TxResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTxResult
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTxResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bloom", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTxResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTxResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bloom = append(m.Bloom[:0], dAtA[iNdEx:postIndex]...)
			if m.Bloom == nil {
				m.Bloom = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxLogs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTxResult
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTxResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TxLogs.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ret", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTxResult
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTxResult
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ret = append(m.Ret[:0], dAtA[iNdEx:postIndex]...)
			if m.Ret == nil {
				m.Ret = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reverted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Reverted = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasUsed", wireType)
			}
			m.GasUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxResult
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasUsed |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTxResult(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTxResult
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTxResult(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTxResult
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTxResult
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTxResult
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTxResult
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTxResult
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTxResult
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTxResult        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTxResult          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTxResult = fmt.Errorf("proto: unexpected end of group")
)
