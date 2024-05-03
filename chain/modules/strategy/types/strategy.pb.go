// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: flux/strategy/v1beta1/strategy.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
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

type Strategy struct {
	Id    []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Data  []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Strategy) Reset()         { *m = Strategy{} }
func (m *Strategy) String() string { return proto.CompactTextString(m) }
func (*Strategy) ProtoMessage()    {}
func (*Strategy) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7aed9d1e3bd608e, []int{0}
}
func (m *Strategy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Strategy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Strategy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Strategy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Strategy.Merge(m, src)
}
func (m *Strategy) XXX_Size() int {
	return m.Size()
}
func (m *Strategy) XXX_DiscardUnknown() {
	xxx_messageInfo_Strategy.DiscardUnknown(m)
}

var xxx_messageInfo_Strategy proto.InternalMessageInfo

func (m *Strategy) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Strategy) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Strategy) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Strategy)(nil), "flux.strategy.v1beta1.Strategy")
}

func init() {
	proto.RegisterFile("flux/strategy/v1beta1/strategy.proto", fileDescriptor_a7aed9d1e3bd608e)
}

var fileDescriptor_a7aed9d1e3bd608e = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xbb, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0xeb, 0x70, 0x11, 0x44, 0x88, 0x21, 0x2a, 0x52, 0xda, 0xc1, 0xaa, 0x10, 0x43, 0xa7,
	0x1c, 0x55, 0x4c, 0xac, 0x08, 0x75, 0x42, 0x0c, 0x2d, 0x13, 0x0b, 0xb2, 0x13, 0xd7, 0xb5, 0x94,
	0xe4, 0x54, 0xb1, 0x0d, 0xed, 0x5b, 0xf0, 0x58, 0x8c, 0x1d, 0x19, 0x51, 0xf2, 0x22, 0x28, 0xb6,
	0x15, 0xb6, 0xff, 0xf2, 0x59, 0xfe, 0xed, 0xf8, 0x6e, 0x53, 0xda, 0x3d, 0x68, 0xd3, 0x30, 0x23,
	0xe4, 0x01, 0x3e, 0x16, 0x5c, 0x18, 0xb6, 0x18, 0x82, 0x6c, 0xd7, 0xa0, 0xc1, 0xe4, 0xa6, 0xa7,
	0xb2, 0x21, 0x0c, 0xd4, 0x74, 0x22, 0x11, 0x65, 0x29, 0xc0, 0x41, 0xdc, 0x6e, 0x80, 0xd5, 0xe1,
	0xc4, 0x74, 0x2c, 0x51, 0xa2, 0x93, 0xd0, 0xab, 0x90, 0x4e, 0x72, 0xd4, 0x15, 0xea, 0x77, 0x5f,
	0x78, 0x13, 0x2a, 0xea, 0x1d, 0x70, 0xa6, 0xc5, 0x30, 0x23, 0x47, 0x55, 0xfb, 0xfe, 0xf6, 0x29,
	0xbe, 0x58, 0x87, 0xfb, 0x93, 0xeb, 0x38, 0x52, 0x45, 0x4a, 0x66, 0x64, 0x7e, 0xb5, 0x8a, 0x54,
	0x91, 0x8c, 0xe3, 0x33, 0xfc, 0xac, 0x45, 0x93, 0x46, 0x33, 0x32, 0xbf, 0x5c, 0x79, 0x93, 0x24,
	0xf1, 0x69, 0xc1, 0x0c, 0x4b, 0x4f, 0x1c, 0xe7, 0xf4, 0xe3, 0xfa, 0xbb, 0xa5, 0xe4, 0xd8, 0x52,
	0xf2, 0xdb, 0x52, 0xf2, 0xd5, 0xd1, 0xd1, 0xb1, 0xa3, 0xa3, 0x9f, 0x8e, 0x8e, 0xde, 0x1e, 0xa4,
	0x32, 0x5b, 0xcb, 0xb3, 0x1c, 0x2b, 0x58, 0x96, 0x76, 0xff, 0xb2, 0x7c, 0x7d, 0x66, 0x5c, 0x43,
	0xff, 0xf2, 0x02, 0xf2, 0x2d, 0x53, 0x35, 0x54, 0x58, 0xd8, 0x52, 0xe8, 0xff, 0xef, 0x32, 0x87,
	0x9d, 0xd0, 0xfc, 0xdc, 0x2d, 0xbc, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x39, 0x47, 0x56, 0x5a,
	0x4c, 0x01, 0x00, 0x00,
}

func (m *Strategy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Strategy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Strategy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintStrategy(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintStrategy(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintStrategy(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStrategy(dAtA []byte, offset int, v uint64) int {
	offset -= sovStrategy(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Strategy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovStrategy(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovStrategy(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovStrategy(uint64(l))
	}
	return n
}

func sovStrategy(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStrategy(x uint64) (n int) {
	return sovStrategy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Strategy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStrategy
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
			return fmt.Errorf("proto: Strategy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Strategy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStrategy
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
				return ErrInvalidLengthStrategy
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStrategy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStrategy
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
				return ErrInvalidLengthStrategy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStrategy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStrategy
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
				return ErrInvalidLengthStrategy
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStrategy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStrategy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStrategy
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
func skipStrategy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStrategy
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
					return 0, ErrIntOverflowStrategy
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
					return 0, ErrIntOverflowStrategy
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
				return 0, ErrInvalidLengthStrategy
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStrategy
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStrategy
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStrategy        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStrategy          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStrategy = fmt.Errorf("proto: unexpected end of group")
)