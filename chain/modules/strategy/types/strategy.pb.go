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
	Id    []byte                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner string                 `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Query *types.FISQueryRequest `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
	// query hash stores hash(query), so that msg server don't need to calculate all the time
	QueryHash []byte `protobuf:"bytes,4,opt,name=query_hash,json=queryHash,proto3" json:"query_hash,omitempty"`
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

func (m *Strategy) GetQuery() *types.FISQueryRequest {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *Strategy) GetQueryHash() []byte {
	if m != nil {
		return m.QueryHash
	}
	return nil
}

type StrategyInput struct {
	Msg      []byte   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	FisInput [][]byte `protobuf:"bytes,2,rep,name=fis_input,json=fisInput,proto3" json:"fis_input,omitempty"`
}

func (m *StrategyInput) Reset()         { *m = StrategyInput{} }
func (m *StrategyInput) String() string { return proto.CompactTextString(m) }
func (*StrategyInput) ProtoMessage()    {}
func (*StrategyInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7aed9d1e3bd608e, []int{1}
}
func (m *StrategyInput) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StrategyInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StrategyInput.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StrategyInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StrategyInput.Merge(m, src)
}
func (m *StrategyInput) XXX_Size() int {
	return m.Size()
}
func (m *StrategyInput) XXX_DiscardUnknown() {
	xxx_messageInfo_StrategyInput.DiscardUnknown(m)
}

var xxx_messageInfo_StrategyInput proto.InternalMessageInfo

func (m *StrategyInput) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *StrategyInput) GetFisInput() [][]byte {
	if m != nil {
		return m.FisInput
	}
	return nil
}

type StrategyOutput struct {
	Instructions []*types.FISInstruction `protobuf:"bytes,1,rep,name=instructions,proto3" json:"instructions,omitempty"`
}

func (m *StrategyOutput) Reset()         { *m = StrategyOutput{} }
func (m *StrategyOutput) String() string { return proto.CompactTextString(m) }
func (*StrategyOutput) ProtoMessage()    {}
func (*StrategyOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7aed9d1e3bd608e, []int{2}
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
	proto.RegisterType((*StrategyInput)(nil), "flux.strategy.v1beta1.StrategyInput")
	proto.RegisterType((*StrategyOutput)(nil), "flux.strategy.v1beta1.StrategyOutput")
}

func init() {
	proto.RegisterFile("flux/strategy/v1beta1/strategy.proto", fileDescriptor_a7aed9d1e3bd608e)
}

var fileDescriptor_a7aed9d1e3bd608e = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0xeb, 0x84, 0x45, 0x5b, 0x6f, 0x59, 0xa1, 0x68, 0x91, 0xb2, 0x8b, 0x08, 0x51, 0x84,
	0x20, 0xa7, 0x58, 0x5b, 0x4e, 0x1c, 0xe0, 0xc0, 0xa1, 0xa2, 0x08, 0x81, 0x48, 0x39, 0x21, 0xa4,
	0xca, 0x49, 0x9c, 0xc4, 0x52, 0x12, 0xb7, 0x19, 0x1b, 0xda, 0x97, 0x40, 0x3c, 0x16, 0xc7, 0x1e,
	0x39, 0xa2, 0xf6, 0x45, 0x50, 0x1c, 0x37, 0x15, 0x07, 0xb8, 0xcd, 0xcc, 0xff, 0xfd, 0x9e, 0x3f,
	0x8e, 0xf1, 0x93, 0xbc, 0x52, 0x1b, 0x02, 0xb2, 0xa5, 0x92, 0x15, 0x5b, 0xf2, 0xf5, 0x36, 0x61,
	0x92, 0xde, 0x0e, 0x83, 0x68, 0xd5, 0x0a, 0x29, 0x9c, 0x07, 0x1d, 0x15, 0x0d, 0x43, 0x43, 0xdd,
	0x5c, 0x17, 0x42, 0x14, 0x15, 0x23, 0x1a, 0x4a, 0x54, 0x4e, 0x68, 0x63, 0x1c, 0x37, 0x57, 0x85,
	0x28, 0x84, 0x2e, 0x49, 0x57, 0x99, 0xe9, 0x75, 0x2a, 0xa0, 0x16, 0xb0, 0xec, 0x85, 0xbe, 0x31,
	0x92, 0xd7, 0x77, 0x24, 0xa1, 0xc0, 0x86, 0x18, 0xa9, 0xe0, 0x8d, 0xd1, 0x1f, 0xeb, 0xa0, 0x14,
	0x64, 0x2b, 0x6a, 0x06, 0xe5, 0x80, 0xc8, 0x8d, 0x01, 0x82, 0x7f, 0x00, 0x6b, 0xc5, 0x5a, 0x93,
	0x2a, 0xf8, 0x8e, 0xf0, 0xf9, 0xc2, 0x7c, 0x85, 0x73, 0x89, 0x2d, 0x9e, 0xb9, 0xc8, 0x47, 0xe1,
	0x24, 0xb6, 0x78, 0xe6, 0x5c, 0xe1, 0x33, 0xf1, 0xad, 0x61, 0xad, 0x6b, 0xf9, 0x28, 0x1c, 0xc7,
	0x7d, 0xe3, 0xbc, 0xc4, 0x67, 0xfa, 0x04, 0xd7, 0xf6, 0x51, 0x78, 0x31, 0x7d, 0x16, 0xe9, 0xab,
	0x18, 0xd6, 0x1c, 0xef, 0x22, 0x9a, 0xcd, 0x17, 0x1f, 0x3b, 0x2e, 0x66, 0x6b, 0xc5, 0x40, 0xc6,
	0xbd, 0xcb, 0x79, 0x84, 0xb1, 0x2e, 0x96, 0x25, 0x85, 0xd2, 0xbd, 0xa3, 0x97, 0x8d, 0xf5, 0xe4,
	0x0d, 0x85, 0x32, 0x78, 0x85, 0xef, 0x1d, 0xf3, 0xcc, 0x9b, 0x95, 0x92, 0xce, 0x7d, 0x6c, 0xd7,
	0x50, 0x98, 0x54, 0x5d, 0xe9, 0x3c, 0xc4, 0xe3, 0x9c, 0xc3, 0x92, 0x77, 0xb2, 0x6b, 0xf9, 0x76,
	0x38, 0x89, 0xcf, 0x73, 0x0e, 0x1a, 0x0f, 0xbe, 0xe0, 0xcb, 0xa3, 0xff, 0x83, 0x92, 0xdd, 0x01,
	0x6f, 0xf1, 0x84, 0x37, 0x20, 0x5b, 0x95, 0x4a, 0x2e, 0x1a, 0x70, 0x91, 0x6f, 0x87, 0x17, 0xd3,
	0xa7, 0xff, 0x89, 0x3d, 0x3f, 0xe1, 0xf1, 0x5f, 0xde, 0xd7, 0x8b, 0x9f, 0x7b, 0x0f, 0xed, 0xf6,
	0x1e, 0xfa, 0xbd, 0xf7, 0xd0, 0x8f, 0x83, 0x37, 0xda, 0x1d, 0xbc, 0xd1, 0xaf, 0x83, 0x37, 0xfa,
	0xfc, 0xa2, 0xe0, 0xb2, 0x54, 0x49, 0x94, 0x8a, 0x9a, 0xcc, 0x2a, 0xb5, 0x79, 0x3f, 0xfb, 0xf4,
	0x8e, 0x26, 0x40, 0xba, 0x2d, 0x19, 0x49, 0x4b, 0xca, 0x1b, 0x52, 0x8b, 0x4c, 0x55, 0x0c, 0x4e,
	0x8f, 0x4b, 0x6e, 0x57, 0x0c, 0x92, 0xbb, 0xfa, 0x57, 0x3c, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff,
	0xa4, 0xf8, 0x59, 0x26, 0x7a, 0x02, 0x00, 0x00,
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
	if len(m.QueryHash) > 0 {
		i -= len(m.QueryHash)
		copy(dAtA[i:], m.QueryHash)
		i = encodeVarintStrategy(dAtA, i, uint64(len(m.QueryHash)))
		i--
		dAtA[i] = 0x22
	}
	if m.Query != nil {
		{
			size, err := m.Query.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStrategy(dAtA, i, uint64(size))
		}
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

func (m *StrategyInput) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StrategyInput) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StrategyInput) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FisInput) > 0 {
		for iNdEx := len(m.FisInput) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.FisInput[iNdEx])
			copy(dAtA[i:], m.FisInput[iNdEx])
			i = encodeVarintStrategy(dAtA, i, uint64(len(m.FisInput[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Msg) > 0 {
		i -= len(m.Msg)
		copy(dAtA[i:], m.Msg)
		i = encodeVarintStrategy(dAtA, i, uint64(len(m.Msg)))
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
	if m.Query != nil {
		l = m.Query.Size()
		n += 1 + l + sovStrategy(uint64(l))
	}
	l = len(m.QueryHash)
	if l > 0 {
		n += 1 + l + sovStrategy(uint64(l))
	}
	return n
}

func (m *StrategyInput) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovStrategy(uint64(l))
	}
	if len(m.FisInput) > 0 {
		for _, b := range m.FisInput {
			l = len(b)
			n += 1 + l + sovStrategy(uint64(l))
		}
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
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
			if m.Query == nil {
				m.Query = &types.FISQueryRequest{}
			}
			if err := m.Query.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryHash", wireType)
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
			m.QueryHash = append(m.QueryHash[:0], dAtA[iNdEx:postIndex]...)
			if m.QueryHash == nil {
				m.QueryHash = []byte{}
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
func (m *StrategyInput) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: StrategyInput: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StrategyInput: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
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
			m.Msg = append(m.Msg[:0], dAtA[iNdEx:postIndex]...)
			if m.Msg == nil {
				m.Msg = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FisInput", wireType)
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
			m.FisInput = append(m.FisInput, make([]byte, postIndex-iNdEx))
			copy(m.FisInput[len(m.FisInput)-1], dAtA[iNdEx:postIndex])
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
