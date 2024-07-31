// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: flux/astromesh/v1beta1/event.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	types "github.com/FluxNFTLabs/sdk-go/chain/eventstream/types"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
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

type AccountBalance struct {
	Acc     []byte                `protobuf:"bytes,1,opt,name=acc,proto3" json:"acc,omitempty"`
	Balance cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=balance,proto3,customtype=cosmossdk.io/math.Int" json:"balance"`
}

func (m *AccountBalance) Reset()         { *m = AccountBalance{} }
func (m *AccountBalance) String() string { return proto.CompactTextString(m) }
func (*AccountBalance) ProtoMessage()    {}
func (*AccountBalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_f84d20cff47df4e1, []int{0}
}
func (m *AccountBalance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountBalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccountBalance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccountBalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountBalance.Merge(m, src)
}
func (m *AccountBalance) XXX_Size() int {
	return m.Size()
}
func (m *AccountBalance) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountBalance.DiscardUnknown(m)
}

var xxx_messageInfo_AccountBalance proto.InternalMessageInfo

func (m *AccountBalance) GetAcc() []byte {
	if m != nil {
		return m.Acc
	}
	return nil
}

type DenomBalanceUpdate struct {
	Denom    string            `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Balances []*AccountBalance `protobuf:"bytes,2,rep,name=balances,proto3" json:"balances,omitempty"`
}

func (m *DenomBalanceUpdate) Reset()         { *m = DenomBalanceUpdate{} }
func (m *DenomBalanceUpdate) String() string { return proto.CompactTextString(m) }
func (*DenomBalanceUpdate) ProtoMessage()    {}
func (*DenomBalanceUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_f84d20cff47df4e1, []int{1}
}
func (m *DenomBalanceUpdate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DenomBalanceUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DenomBalanceUpdate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DenomBalanceUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DenomBalanceUpdate.Merge(m, src)
}
func (m *DenomBalanceUpdate) XXX_Size() int {
	return m.Size()
}
func (m *DenomBalanceUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_DenomBalanceUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_DenomBalanceUpdate proto.InternalMessageInfo

func (m *DenomBalanceUpdate) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *DenomBalanceUpdate) GetBalances() []*AccountBalance {
	if m != nil {
		return m.Balances
	}
	return nil
}

type BalanceUpdateEvent struct {
	Op    types.Op              `protobuf:"varint,1,opt,name=op,proto3,enum=flux.eventstream.v1beta1.Op" json:"op,omitempty"`
	Upd   []*DenomBalanceUpdate `protobuf:"bytes,2,rep,name=upd,proto3" json:"upd,omitempty"`
	Plane Plane                 `protobuf:"varint,3,opt,name=plane,proto3,enum=flux.astromesh.v1beta1.Plane" json:"plane,omitempty"`
}

func (m *BalanceUpdateEvent) Reset()         { *m = BalanceUpdateEvent{} }
func (m *BalanceUpdateEvent) String() string { return proto.CompactTextString(m) }
func (*BalanceUpdateEvent) ProtoMessage()    {}
func (*BalanceUpdateEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_f84d20cff47df4e1, []int{2}
}
func (m *BalanceUpdateEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BalanceUpdateEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BalanceUpdateEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BalanceUpdateEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceUpdateEvent.Merge(m, src)
}
func (m *BalanceUpdateEvent) XXX_Size() int {
	return m.Size()
}
func (m *BalanceUpdateEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceUpdateEvent.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceUpdateEvent proto.InternalMessageInfo

func (m *BalanceUpdateEvent) GetOp() types.Op {
	if m != nil {
		return m.Op
	}
	return types.Op_COSMOS_SET
}

func (m *BalanceUpdateEvent) GetUpd() []*DenomBalanceUpdate {
	if m != nil {
		return m.Upd
	}
	return nil
}

func (m *BalanceUpdateEvent) GetPlane() Plane {
	if m != nil {
		return m.Plane
	}
	return Plane_COSMOS
}

func init() {
	proto.RegisterType((*AccountBalance)(nil), "flux.astromesh.v1beta1.AccountBalance")
	proto.RegisterType((*DenomBalanceUpdate)(nil), "flux.astromesh.v1beta1.DenomBalanceUpdate")
	proto.RegisterType((*BalanceUpdateEvent)(nil), "flux.astromesh.v1beta1.BalanceUpdateEvent")
}

func init() {
	proto.RegisterFile("flux/astromesh/v1beta1/event.proto", fileDescriptor_f84d20cff47df4e1)
}

var fileDescriptor_f84d20cff47df4e1 = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0xb3, 0x89, 0x0a, 0xd4, 0xa0, 0x0a, 0xac, 0x82, 0x42, 0x05, 0x9b, 0x28, 0x42, 0x28,
	0xaa, 0xc0, 0xa6, 0xed, 0x0d, 0x71, 0x21, 0x82, 0x4a, 0x45, 0x08, 0xd0, 0xaa, 0x5c, 0xb8, 0x20,
	0xaf, 0xd7, 0x24, 0x2b, 0xd6, 0x1e, 0x13, 0x7b, 0xab, 0xf4, 0x2d, 0x78, 0x0c, 0x8e, 0x1c, 0xe0,
	0x1d, 0x7a, 0xac, 0x38, 0x21, 0x0e, 0x15, 0x4a, 0x0e, 0xbc, 0x06, 0xf2, 0x1f, 0x56, 0x8d, 0x68,
	0x2e, 0x2b, 0x8f, 0xe7, 0x37, 0xf3, 0x7d, 0x33, 0x6b, 0x34, 0xf8, 0x50, 0xd5, 0x33, 0xca, 0x8c,
	0x9d, 0x82, 0x14, 0x66, 0x42, 0x8f, 0x76, 0x72, 0x61, 0xd9, 0x0e, 0x15, 0x47, 0x42, 0x59, 0xa2,
	0xa7, 0x60, 0x01, 0xdf, 0x72, 0x0c, 0x69, 0x18, 0x12, 0x99, 0xad, 0xcd, 0x31, 0x8c, 0xc1, 0x23,
	0xd4, 0x9d, 0x02, 0xbd, 0x75, 0x9b, 0x83, 0x91, 0x60, 0xde, 0x87, 0x44, 0x08, 0x62, 0x2a, 0x0d,
	0x11, 0xcd, 0x99, 0x11, 0x8d, 0x12, 0x87, 0x52, 0xc5, 0xfc, 0x0d, 0x26, 0x4b, 0x05, 0xd4, 0x7f,
	0xe3, 0x55, 0x6f, 0x85, 0x3f, 0x3b, 0x8b, 0xc0, 0x3d, 0x0f, 0x78, 0xbb, 0xc6, 0x4e, 0x05, 0x93,
	0x0d, 0xf2, 0xa9, 0x16, 0xd3, 0xe3, 0x40, 0x0d, 0x14, 0xda, 0x78, 0xca, 0x39, 0xd4, 0xca, 0x8e,
	0x58, 0xc5, 0x14, 0x17, 0xf8, 0x3a, 0xea, 0x30, 0xce, 0xbb, 0x49, 0x3f, 0x19, 0x5e, 0xcb, 0xdc,
	0x11, 0xbf, 0x40, 0x97, 0xf3, 0x90, 0xec, 0xb6, 0xfb, 0xc9, 0x70, 0x7d, 0xf4, 0xe8, 0xe4, 0xac,
	0xd7, 0xfa, 0x75, 0xd6, 0xbb, 0x19, 0x6c, 0x9b, 0xe2, 0x23, 0x29, 0x81, 0x4a, 0x66, 0x27, 0xe4,
	0x40, 0xd9, 0x1f, 0xdf, 0x1e, 0xa2, 0x38, 0xdd, 0x81, 0xb2, 0x5f, 0xfe, 0x7c, 0xdd, 0x4e, 0xb2,
	0x7f, 0x0d, 0x06, 0x0a, 0xe1, 0x67, 0x42, 0x81, 0x8c, 0x6a, 0x6f, 0x75, 0xc1, 0xac, 0xc0, 0x9b,
	0x68, 0xad, 0x70, 0xb7, 0x5e, 0x75, 0x3d, 0x0b, 0x01, 0x1e, 0xa1, 0x2b, 0xb1, 0xcc, 0x74, 0xdb,
	0xfd, 0xce, 0xf0, 0xea, 0xee, 0x7d, 0x72, 0xf1, 0xc6, 0xc9, 0xf2, 0x0c, 0x59, 0x53, 0x37, 0xf8,
	0x9e, 0x20, 0xbc, 0xa4, 0xf5, 0xdc, 0x2d, 0x04, 0x3f, 0x40, 0x6d, 0xd0, 0x5e, 0x6d, 0x63, 0xf7,
	0x4e, 0x68, 0x7a, 0x6e, 0x53, 0x4d, 0xdb, 0xd7, 0x3a, 0x6b, 0x83, 0xc6, 0x4f, 0x50, 0xa7, 0xd6,
	0x45, 0xf4, 0xb0, 0xbd, 0xca, 0xc3, 0xff, 0x73, 0x65, 0xae, 0x0c, 0xef, 0xa1, 0x35, 0x5d, 0x31,
	0x25, 0xba, 0x1d, 0x2f, 0x77, 0x77, 0x55, 0xfd, 0x1b, 0x07, 0x65, 0x81, 0x1d, 0x1d, 0x9e, 0xcc,
	0xd3, 0xe4, 0x74, 0x9e, 0x26, 0xbf, 0xe7, 0x69, 0xf2, 0x79, 0x91, 0xb6, 0x4e, 0x17, 0x69, 0xeb,
	0xe7, 0x22, 0x6d, 0xbd, 0x7b, 0x3c, 0x2e, 0xed, 0xa4, 0xce, 0x09, 0x07, 0x49, 0xf7, 0xab, 0x7a,
	0xf6, 0x6a, 0xff, 0xf0, 0x25, 0xcb, 0x0d, 0x75, 0x5d, 0x0b, 0xca, 0x27, 0xac, 0x54, 0x54, 0x42,
	0x51, 0x57, 0xc2, 0x9c, 0x7b, 0x1e, 0xf6, 0x58, 0x0b, 0x93, 0x5f, 0xf2, 0x3f, 0x7d, 0xef, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xdb, 0xaa, 0x9c, 0x1c, 0xdd, 0x02, 0x00, 0x00,
}

func (m *AccountBalance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountBalance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccountBalance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Balance.Size()
		i -= size
		if _, err := m.Balance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Acc) > 0 {
		i -= len(m.Acc)
		copy(dAtA[i:], m.Acc)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Acc)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DenomBalanceUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DenomBalanceUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DenomBalanceUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Balances) > 0 {
		for iNdEx := len(m.Balances) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Balances[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvent(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BalanceUpdateEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BalanceUpdateEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BalanceUpdateEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Plane != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.Plane))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Upd) > 0 {
		for iNdEx := len(m.Upd) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Upd[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvent(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Op != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.Op))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AccountBalance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Acc)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = m.Balance.Size()
	n += 1 + l + sovEvent(uint64(l))
	return n
}

func (m *DenomBalanceUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if len(m.Balances) > 0 {
		for _, e := range m.Balances {
			l = e.Size()
			n += 1 + l + sovEvent(uint64(l))
		}
	}
	return n
}

func (m *BalanceUpdateEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Op != 0 {
		n += 1 + sovEvent(uint64(m.Op))
	}
	if len(m.Upd) > 0 {
		for _, e := range m.Upd {
			l = e.Size()
			n += 1 + l + sovEvent(uint64(l))
		}
	}
	if m.Plane != 0 {
		n += 1 + sovEvent(uint64(m.Plane))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccountBalance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: AccountBalance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountBalance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Acc", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Acc = append(m.Acc[:0], dAtA[iNdEx:postIndex]...)
			if m.Acc == nil {
				m.Acc = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Balance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *DenomBalanceUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: DenomBalanceUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DenomBalanceUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balances", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Balances = append(m.Balances, &AccountBalance{})
			if err := m.Balances[len(m.Balances)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *BalanceUpdateEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: BalanceUpdateEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BalanceUpdateEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Op", wireType)
			}
			m.Op = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Op |= types.Op(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Upd", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Upd = append(m.Upd, &DenomBalanceUpdate{})
			if err := m.Upd[len(m.Upd)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Plane", wireType)
			}
			m.Plane = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Plane |= Plane(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)