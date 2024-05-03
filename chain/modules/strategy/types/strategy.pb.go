// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: flux/strategy/v1beta1/strategy.proto

package types

import (
	fmt "fmt"
	types "github.com/FluxNFTLabs/sdk-go/chain/modules/astromesh/types"
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

type StrategyOutput struct {
	Instructions []*types.FISInstruction `protobuf:"bytes,1,rep,name=instructions,proto3" json:"instructions,omitempty"`
}

func (m *StrategyOutput) Reset()         { *m = StrategyOutput{} }
func (m *StrategyOutput) String() string { return proto.CompactTextString(m) }
func (*StrategyOutput) ProtoMessage()    {}
func (*StrategyOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7aed9d1e3bd608e, []int{1}
}
func (m *StrategyOutput) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StrategyOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StrategyOutput.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StrategyOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StrategyOutput.Merge(m, src)
}
func (m *StrategyOutput) XXX_Size() int {
	return m.Size()
}
func (m *StrategyOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_StrategyOutput.DiscardUnknown(m)
}

var xxx_messageInfo_StrategyOutput proto.InternalMessageInfo

func (m *StrategyOutput) GetInstructions() []*types.FISInstruction {
	if m != nil {
		return m.Instructions
	}
	return nil
}

func init() {
	proto.RegisterType((*Strategy)(nil), "flux.strategy.v1beta1.Strategy")
	proto.RegisterType((*StrategyOutput)(nil), "flux.strategy.v1beta1.StrategyOutput")
}

func init() {
	proto.RegisterFile("flux/strategy/v1beta1/strategy.proto", fileDescriptor_a7aed9d1e3bd608e)
}

var fileDescriptor_a7aed9d1e3bd608e = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xc1, 0x4a, 0x33, 0x31,
	0x14, 0x85, 0x9b, 0xfe, 0xfc, 0xa2, 0x63, 0xe9, 0x62, 0xa8, 0xd0, 0x76, 0x11, 0x4b, 0x11, 0xe9,
	0x6a, 0x62, 0x75, 0xe5, 0xd6, 0x45, 0xa1, 0x22, 0x0a, 0xad, 0x2b, 0x11, 0x24, 0x99, 0x49, 0xd3,
	0xc0, 0x4c, 0x6e, 0x99, 0xdc, 0x68, 0xfb, 0x16, 0x3e, 0x96, 0xcb, 0x2e, 0x5d, 0x4a, 0xfb, 0x22,
	0xd2, 0x99, 0x74, 0x8a, 0xbb, 0xdc, 0x7b, 0xbe, 0xc3, 0x3d, 0x87, 0x04, 0x17, 0xb3, 0xd4, 0x2d,
	0x99, 0xc5, 0x9c, 0xa3, 0x54, 0x2b, 0xf6, 0x3e, 0x14, 0x12, 0xf9, 0xb0, 0x5a, 0x44, 0x8b, 0x1c,
	0x10, 0xc2, 0xb3, 0x1d, 0x15, 0x55, 0x4b, 0x4f, 0x75, 0x3b, 0x0a, 0x40, 0xa5, 0x92, 0x15, 0x90,
	0x70, 0x33, 0xc6, 0x8d, 0x77, 0x74, 0x5b, 0x0a, 0x14, 0x14, 0x4f, 0xb6, 0x7b, 0xf9, 0x6d, 0x27,
	0x06, 0x9b, 0x81, 0x7d, 0x2b, 0x85, 0x72, 0xf0, 0x12, 0x2d, 0x27, 0x26, 0xb8, 0x95, 0x55, 0x8c,
	0x18, 0xb4, 0xf1, 0xfa, 0x79, 0x11, 0x94, 0x5b, 0xcc, 0x21, 0x93, 0x76, 0x5e, 0x21, 0xb8, 0x2c,
	0x81, 0xfe, 0x55, 0x70, 0x3c, 0xf5, 0x01, 0xc3, 0x66, 0x50, 0xd7, 0x49, 0x9b, 0xf4, 0xc8, 0xa0,
	0x31, 0xa9, 0xeb, 0x24, 0x6c, 0x05, 0xff, 0xe1, 0xc3, 0xc8, 0xbc, 0x5d, 0xef, 0x91, 0xc1, 0xc9,
	0xa4, 0x1c, 0xfa, 0xaf, 0x41, 0x73, 0xef, 0x78, 0x72, 0xb8, 0x70, 0x18, 0xde, 0x07, 0x0d, 0x6d,
	0x2c, 0xe6, 0x2e, 0x46, 0x0d, 0xc6, 0xb6, 0x49, 0xef, 0xdf, 0xe0, 0xf4, 0xfa, 0x32, 0x2a, 0xea,
	0x57, 0xb7, 0xf7, 0xfd, 0xa3, 0xd1, 0x78, 0x3a, 0x3e, 0xe0, 0x93, 0x3f, 0xde, 0xbb, 0xe9, 0xd7,
	0x86, 0x92, 0xf5, 0x86, 0x92, 0x9f, 0x0d, 0x25, 0x9f, 0x5b, 0x5a, 0x5b, 0x6f, 0x69, 0xed, 0x7b,
	0x4b, 0x6b, 0x2f, 0xb7, 0x4a, 0xe3, 0xdc, 0x89, 0x28, 0x86, 0x8c, 0x8d, 0x52, 0xb7, 0x7c, 0x1c,
	0x3d, 0x3f, 0x70, 0x61, 0xd9, 0xee, 0x4a, 0xc2, 0xe2, 0x39, 0xd7, 0x86, 0x65, 0x90, 0xb8, 0x54,
	0xda, 0xc3, 0xcf, 0xe0, 0x6a, 0x21, 0xad, 0x38, 0x2a, 0xba, 0xde, 0xfc, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x8c, 0x88, 0x56, 0x1b, 0xb7, 0x01, 0x00, 0x00,
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

func (m *StrategyOutput) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StrategyOutput) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StrategyOutput) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Instructions) > 0 {
		for iNdEx := len(m.Instructions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Instructions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStrategy(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
	return n
}

func (m *StrategyOutput) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Instructions) > 0 {
		for _, e := range m.Instructions {
			l = e.Size()
			n += 1 + l + sovStrategy(uint64(l))
		}
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
func (m *StrategyOutput) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: StrategyOutput: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StrategyOutput: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Instructions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStrategy
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
				return ErrInvalidLengthStrategy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStrategy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Instructions = append(m.Instructions, &types.FISInstruction{})
			if err := m.Instructions[len(m.Instructions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
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
