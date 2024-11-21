// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: flux/interpool/v1beta1/event.proto

package types

import (
	fmt "fmt"
	types "github.com/FluxNFTLabs/sdk-go/chain/eventstream/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// Event to capture the creation, update, or deletion of a pool
type InterPoolEvent struct {
	Op   types.Op   `protobuf:"varint,1,opt,name=op,proto3,enum=flux.eventstream.v1beta1.Op" json:"op,omitempty"`
	Pool *InterPool `protobuf:"bytes,2,opt,name=pool,proto3" json:"pool,omitempty"`
}

func (m *InterPoolEvent) Reset()         { *m = InterPoolEvent{} }
func (m *InterPoolEvent) String() string { return proto.CompactTextString(m) }
func (*InterPoolEvent) ProtoMessage()    {}
func (*InterPoolEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_74f01531e47c63f0, []int{0}
}
func (m *InterPoolEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InterPoolEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InterPoolEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InterPoolEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterPoolEvent.Merge(m, src)
}
func (m *InterPoolEvent) XXX_Size() int {
	return m.Size()
}
func (m *InterPoolEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_InterPoolEvent.DiscardUnknown(m)
}

var xxx_messageInfo_InterPoolEvent proto.InternalMessageInfo

func (m *InterPoolEvent) GetOp() types.Op {
	if m != nil {
		return m.Op
	}
	return types.Op_COSMOS_SET
}

func (m *InterPoolEvent) GetPool() *InterPool {
	if m != nil {
		return m.Pool
	}
	return nil
}

// Event to capture liquidity changes (deposit, withdrawal, etc.)
type LiquidityEvent struct {
	Op types.Op `protobuf:"varint,1,opt,name=op,proto3,enum=flux.eventstream.v1beta1.Op" json:"op,omitempty"`
	// The pool information and its updates
	Pool *InterPool `protobuf:"bytes,2,opt,name=pool,proto3" json:"pool,omitempty"`
	// The specific pool shares that were updated as part of the liquidity event
	PoolShare *PoolShare `protobuf:"bytes,3,opt,name=pool_share,json=poolShare,proto3" json:"pool_share,omitempty"`
	// The capital that has changed as part of the event (amounts added or
	// removed)
	ChangedCapital []types1.Coin `protobuf:"bytes,4,rep,name=changed_capital,json=changedCapital,proto3" json:"changed_capital"`
}

func (m *LiquidityEvent) Reset()         { *m = LiquidityEvent{} }
func (m *LiquidityEvent) String() string { return proto.CompactTextString(m) }
func (*LiquidityEvent) ProtoMessage()    {}
func (*LiquidityEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_74f01531e47c63f0, []int{1}
}
func (m *LiquidityEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LiquidityEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LiquidityEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LiquidityEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiquidityEvent.Merge(m, src)
}
func (m *LiquidityEvent) XXX_Size() int {
	return m.Size()
}
func (m *LiquidityEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_LiquidityEvent.DiscardUnknown(m)
}

var xxx_messageInfo_LiquidityEvent proto.InternalMessageInfo

func (m *LiquidityEvent) GetOp() types.Op {
	if m != nil {
		return m.Op
	}
	return types.Op_COSMOS_SET
}

func (m *LiquidityEvent) GetPool() *InterPool {
	if m != nil {
		return m.Pool
	}
	return nil
}

func (m *LiquidityEvent) GetPoolShare() *PoolShare {
	if m != nil {
		return m.PoolShare
	}
	return nil
}

func (m *LiquidityEvent) GetChangedCapital() []types1.Coin {
	if m != nil {
		return m.ChangedCapital
	}
	return nil
}

func init() {
	proto.RegisterType((*InterPoolEvent)(nil), "flux.interpool.v1beta1.InterPoolEvent")
	proto.RegisterType((*LiquidityEvent)(nil), "flux.interpool.v1beta1.LiquidityEvent")
}

func init() {
	proto.RegisterFile("flux/interpool/v1beta1/event.proto", fileDescriptor_74f01531e47c63f0)
}

var fileDescriptor_74f01531e47c63f0 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4d, 0x8b, 0xd3, 0x40,
	0x18, 0xc7, 0x93, 0x6c, 0x11, 0x9c, 0x85, 0x08, 0x41, 0x24, 0x2e, 0x92, 0xad, 0x45, 0xa4, 0x07,
	0x99, 0x61, 0x2b, 0x5e, 0xf4, 0xb6, 0xc5, 0xa2, 0x50, 0x54, 0x6a, 0x4f, 0x5e, 0xca, 0x24, 0x99,
	0x26, 0x03, 0x49, 0x9e, 0x69, 0x66, 0xa6, 0xb4, 0x1f, 0x42, 0xf0, 0x63, 0xf5, 0xd8, 0xa3, 0x27,
	0x91, 0xf6, 0x8b, 0xc8, 0x4c, 0xd2, 0xd8, 0x83, 0x0a, 0xee, 0x29, 0xc3, 0xe4, 0xf7, 0x7f, 0x79,
	0xf2, 0x04, 0x0d, 0x96, 0x85, 0xde, 0x10, 0x5e, 0x29, 0x56, 0x0b, 0x80, 0x82, 0xac, 0x6f, 0x62,
	0xa6, 0xe8, 0x0d, 0x61, 0x6b, 0x56, 0x29, 0x2c, 0x6a, 0x50, 0x10, 0x3c, 0x32, 0x0c, 0xee, 0x18,
	0xdc, 0x32, 0x57, 0xd7, 0x19, 0x40, 0x56, 0x30, 0x62, 0xa9, 0x58, 0x2f, 0x89, 0xe2, 0x25, 0x93,
	0x8a, 0x96, 0xa2, 0x11, 0x5e, 0x3d, 0xcc, 0x20, 0x03, 0x7b, 0x24, 0xe6, 0xd4, 0xde, 0x46, 0x09,
	0xc8, 0x12, 0x24, 0x89, 0xa9, 0x64, 0x5d, 0x5e, 0x02, 0xbc, 0x6a, 0xdf, 0x3f, 0xff, 0x4b, 0xa5,
	0xdf, 0x05, 0x1a, 0xee, 0x99, 0xe5, 0x6c, 0x51, 0xa9, 0x6a, 0x46, 0xcb, 0x8e, 0x5c, 0x69, 0x56,
	0x6f, 0x1b, 0x6a, 0xa0, 0x91, 0xff, 0xde, 0x08, 0x3f, 0x01, 0x14, 0x6f, 0x0d, 0x1b, 0xbc, 0x40,
	0x1e, 0x88, 0xd0, 0xed, 0xbb, 0x43, 0x7f, 0xf4, 0x04, 0xdb, 0xd9, 0xce, 0x4c, 0x4e, 0xd3, 0xe1,
	0x8f, 0x62, 0xe6, 0x81, 0x08, 0x5e, 0xa1, 0x9e, 0xc9, 0x0c, 0xbd, 0xbe, 0x3b, 0xbc, 0x1c, 0x3d,
	0xc5, 0x7f, 0xfe, 0x16, 0xb8, 0xcb, 0x98, 0x59, 0x7c, 0xf0, 0xd5, 0x43, 0xfe, 0x94, 0xaf, 0x34,
	0x4f, 0xb9, 0xda, 0xde, 0x25, 0xf7, 0xcd, 0x7f, 0xe6, 0xde, 0xf6, 0x76, 0x3f, 0xae, 0xdd, 0x26,
	0x3d, 0x98, 0x20, 0x64, 0x9e, 0x0b, 0x99, 0xd3, 0x9a, 0x85, 0x17, 0xff, 0xb6, 0x30, 0xea, 0xcf,
	0x06, 0x6c, 0x2d, 0xee, 0x8b, 0xd3, 0x45, 0xf0, 0x0e, 0x3d, 0x48, 0x72, 0x5a, 0x65, 0x2c, 0x5d,
	0x24, 0x54, 0x70, 0x45, 0x8b, 0xb0, 0xd7, 0xbf, 0x18, 0x5e, 0x8e, 0x1e, 0xe3, 0x66, 0x89, 0xd8,
	0x2c, 0xb1, 0x73, 0x1a, 0x03, 0xaf, 0xac, 0x89, 0x33, 0xf3, 0x5b, 0xdd, 0xb8, 0x91, 0xdd, 0xce,
	0x77, 0x87, 0xc8, 0xdd, 0x1f, 0x22, 0xf7, 0xe7, 0x21, 0x72, 0xbf, 0x1d, 0x23, 0x67, 0x7f, 0x8c,
	0x9c, 0xef, 0xc7, 0xc8, 0xf9, 0xf2, 0x3a, 0xe3, 0x2a, 0xd7, 0x31, 0x4e, 0xa0, 0x24, 0x93, 0x42,
	0x6f, 0x3e, 0x4c, 0xe6, 0x53, 0x1a, 0x4b, 0x62, 0xda, 0xa6, 0x24, 0xc9, 0x29, 0xaf, 0x48, 0x09,
	0xa9, 0x2e, 0x98, 0x3c, 0xfb, 0x29, 0xd4, 0x56, 0x30, 0x19, 0xdf, 0xb3, 0x3b, 0x7e, 0xf9, 0x2b,
	0x00, 0x00, 0xff, 0xff, 0x30, 0x3b, 0x64, 0x2f, 0xc6, 0x02, 0x00, 0x00,
}

func (m *InterPoolEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InterPoolEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InterPoolEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pool != nil {
		{
			size, err := m.Pool.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Op != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.Op))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LiquidityEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LiquidityEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LiquidityEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChangedCapital) > 0 {
		for iNdEx := len(m.ChangedCapital) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChangedCapital[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvent(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.PoolShare != nil {
		{
			size, err := m.PoolShare.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Pool != nil {
		{
			size, err := m.Pool.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
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
func (m *InterPoolEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Op != 0 {
		n += 1 + sovEvent(uint64(m.Op))
	}
	if m.Pool != nil {
		l = m.Pool.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *LiquidityEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Op != 0 {
		n += 1 + sovEvent(uint64(m.Op))
	}
	if m.Pool != nil {
		l = m.Pool.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.PoolShare != nil {
		l = m.PoolShare.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	if len(m.ChangedCapital) > 0 {
		for _, e := range m.ChangedCapital {
			l = e.Size()
			n += 1 + l + sovEvent(uint64(l))
		}
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InterPoolEvent) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: InterPoolEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InterPoolEvent: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Pool", wireType)
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
			if m.Pool == nil {
				m.Pool = &InterPool{}
			}
			if err := m.Pool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *LiquidityEvent) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: LiquidityEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LiquidityEvent: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Pool", wireType)
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
			if m.Pool == nil {
				m.Pool = &InterPool{}
			}
			if err := m.Pool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolShare", wireType)
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
			if m.PoolShare == nil {
				m.PoolShare = &PoolShare{}
			}
			if err := m.PoolShare.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChangedCapital", wireType)
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
			m.ChangedCapital = append(m.ChangedCapital, types1.Coin{})
			if err := m.ChangedCapital[len(m.ChangedCapital)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
