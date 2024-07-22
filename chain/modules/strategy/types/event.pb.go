// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: flux/strategy/v1beta1/event.proto

package types

import (
	fmt "fmt"
	types "github.com/FluxNFTLabs/sdk-go/chain/eventstream/types"
	types2 "github.com/FluxNFTLabs/sdk-go/chain/modules/astromesh/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types1 "github.com/cosmos/gogoproto/types"
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

type StrategyUpdateEvent struct {
	Op                types.Op                `protobuf:"varint,1,opt,name=op,proto3,enum=flux.eventstream.v1beta1.Op" json:"op,omitempty"`
	Id                []byte                  `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	IsEnabled         *types1.BoolValue       `protobuf:"bytes,3,opt,name=is_enabled,json=isEnabled,proto3" json:"is_enabled,omitempty" bson:"is_enabled"`
	TriggerPermission *PermissionConfig       `protobuf:"bytes,4,opt,name=trigger_permission,json=triggerPermission,proto3" json:"trigger_permission,omitempty" bson:"trigger_permission"`
	Query             *types2.FISQueryRequest `protobuf:"bytes,5,opt,name=query,proto3" json:"query,omitempty"`
	QueryHash         []byte                  `protobuf:"bytes,6,opt,name=query_hash,json=queryHash,proto3" json:"query_hash,omitempty" bson:"query_hash"`
	Metadata          *StrategyMetadata       `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *StrategyUpdateEvent) Reset()         { *m = StrategyUpdateEvent{} }
func (m *StrategyUpdateEvent) String() string { return proto.CompactTextString(m) }
func (*StrategyUpdateEvent) ProtoMessage()    {}
func (*StrategyUpdateEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_c72759cbd5b43ae0, []int{0}
}
func (m *StrategyUpdateEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StrategyUpdateEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StrategyUpdateEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StrategyUpdateEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StrategyUpdateEvent.Merge(m, src)
}
func (m *StrategyUpdateEvent) XXX_Size() int {
	return m.Size()
}
func (m *StrategyUpdateEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_StrategyUpdateEvent.DiscardUnknown(m)
}

var xxx_messageInfo_StrategyUpdateEvent proto.InternalMessageInfo

func (m *StrategyUpdateEvent) GetOp() types.Op {
	if m != nil {
		return m.Op
	}
	return types.Op_COSMOS_SET
}

func (m *StrategyUpdateEvent) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *StrategyUpdateEvent) GetIsEnabled() *types1.BoolValue {
	if m != nil {
		return m.IsEnabled
	}
	return nil
}

func (m *StrategyUpdateEvent) GetTriggerPermission() *PermissionConfig {
	if m != nil {
		return m.TriggerPermission
	}
	return nil
}

func (m *StrategyUpdateEvent) GetQuery() *types2.FISQueryRequest {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *StrategyUpdateEvent) GetQueryHash() []byte {
	if m != nil {
		return m.QueryHash
	}
	return nil
}

func (m *StrategyUpdateEvent) GetMetadata() *StrategyMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*StrategyUpdateEvent)(nil), "flux.strategy.v1beta1.StrategyUpdateEvent")
}

func init() { proto.RegisterFile("flux/strategy/v1beta1/event.proto", fileDescriptor_c72759cbd5b43ae0) }

var fileDescriptor_c72759cbd5b43ae0 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0xe3, 0x94, 0x16, 0x3a, 0xa0, 0x4a, 0x35, 0x54, 0x72, 0x23, 0x70, 0x82, 0x85, 0x44,
	0x16, 0x68, 0x46, 0x2d, 0x6c, 0x40, 0x62, 0xe3, 0xaa, 0x11, 0x48, 0xfc, 0x3a, 0x85, 0x05, 0x9b,
	0x68, 0x5c, 0xdf, 0xd8, 0x23, 0xd9, 0x1e, 0x77, 0x7e, 0x4a, 0xf3, 0x16, 0x3c, 0x56, 0x97, 0x5d,
	0x22, 0x16, 0x11, 0x4a, 0xde, 0xa0, 0x4f, 0x80, 0x3c, 0x76, 0xec, 0x22, 0x22, 0x76, 0xd7, 0xe3,
	0xef, 0x9c, 0xe3, 0x7b, 0xc6, 0xe8, 0xf1, 0x34, 0xd5, 0x17, 0x44, 0x2a, 0x41, 0x15, 0xc4, 0x33,
	0x72, 0x7e, 0x10, 0x82, 0xa2, 0x07, 0x04, 0xce, 0x21, 0x57, 0xb8, 0x10, 0x5c, 0x71, 0x7b, 0xaf,
	0x44, 0xf0, 0x0a, 0xc1, 0x35, 0xd2, 0x73, 0x63, 0xce, 0xe3, 0x14, 0x88, 0x81, 0x42, 0x3d, 0x25,
	0xdf, 0x05, 0x2d, 0x0a, 0x10, 0xb2, 0x92, 0xf5, 0x1e, 0xc4, 0x3c, 0xe6, 0x66, 0x24, 0xe5, 0x54,
	0x9f, 0x3e, 0x59, 0x9f, 0xd7, 0xb8, 0x57, 0x94, 0x67, 0x28, 0x2a, 0x95, 0xe0, 0x19, 0xc8, 0xa4,
	0xc1, 0xce, 0x34, 0x88, 0xd9, 0x5f, 0x4e, 0xe6, 0x43, 0xa5, 0x12, 0x40, 0xb3, 0x75, 0x94, 0xf7,
	0x6b, 0x03, 0xdd, 0x1f, 0xd7, 0xe6, 0x5f, 0x8a, 0x88, 0x2a, 0x38, 0x2e, 0x15, 0xf6, 0x33, 0xd4,
	0xe5, 0x85, 0x63, 0x0d, 0xac, 0xe1, 0xce, 0xe1, 0x43, 0x6c, 0x36, 0xbc, 0x61, 0xb5, 0x5a, 0x12,
	0x7f, 0x2c, 0x82, 0x2e, 0x2f, 0xec, 0x1d, 0xd4, 0x65, 0x91, 0xd3, 0x1d, 0x58, 0xc3, 0x7b, 0x41,
	0x97, 0x45, 0xf6, 0x09, 0x42, 0x4c, 0x4e, 0x20, 0xa7, 0x61, 0x0a, 0x91, 0xb3, 0x31, 0xb0, 0x86,
	0x77, 0x0f, 0x7b, 0xb8, 0x2a, 0x04, 0xaf, 0x0a, 0xc1, 0x3e, 0xe7, 0xe9, 0x57, 0x9a, 0x6a, 0xf0,
	0xf7, 0x2f, 0xe7, 0x7d, 0xeb, 0x7a, 0xde, 0xdf, 0x0d, 0x25, 0xcf, 0x5f, 0x79, 0xad, 0xd6, 0x0b,
	0xb6, 0x99, 0x3c, 0xae, 0x66, 0x5b, 0x23, 0x5b, 0x09, 0x16, 0xc7, 0x20, 0x26, 0x05, 0x88, 0x8c,
	0x49, 0xc9, 0x78, 0xee, 0xdc, 0x32, 0xee, 0x4f, 0xf1, 0xda, 0x5b, 0xc0, 0x9f, 0x1a, 0xf0, 0x88,
	0xe7, 0x53, 0x16, 0xfb, 0x8f, 0xae, 0xe7, 0xfd, 0xfd, 0x2a, 0xe6, 0x5f, 0x33, 0x2f, 0xd8, 0xad,
	0x0f, 0x5b, 0x9d, 0xfd, 0x1a, 0x6d, 0x9a, 0xc6, 0x9c, 0xcd, 0x9b, 0x49, 0x4d, 0xf9, 0x4d, 0xd4,
	0xe8, 0xed, 0xf8, 0x73, 0xc9, 0x05, 0x70, 0xa6, 0x41, 0xaa, 0xa0, 0x52, 0xd9, 0x2f, 0x10, 0x32,
	0xc3, 0x24, 0xa1, 0x32, 0x71, 0xb6, 0xca, 0x8e, 0xfc, 0xbd, 0x76, 0xd7, 0xf6, 0x9d, 0x17, 0x6c,
	0x9b, 0x87, 0x37, 0x54, 0x26, 0xf6, 0x11, 0xba, 0x93, 0x81, 0xa2, 0x11, 0x55, 0xd4, 0xb9, 0xfd,
	0xdf, 0x0d, 0x57, 0xb7, 0xf7, 0xbe, 0xc6, 0x83, 0x46, 0xe8, 0x8f, 0x2f, 0x17, 0xae, 0x75, 0xb5,
	0x70, 0xad, 0xdf, 0x0b, 0xd7, 0xfa, 0xb1, 0x74, 0x3b, 0x57, 0x4b, 0xb7, 0xf3, 0x73, 0xe9, 0x76,
	0xbe, 0xbd, 0x8c, 0x99, 0x4a, 0x74, 0x88, 0x4f, 0x79, 0x46, 0x46, 0xa9, 0xbe, 0xf8, 0x30, 0x3a,
	0x79, 0x47, 0x43, 0x49, 0xca, 0x88, 0x88, 0x9c, 0x26, 0x94, 0xe5, 0x24, 0xe3, 0x91, 0x4e, 0x41,
	0xb6, 0x3f, 0xa3, 0x9a, 0x15, 0x20, 0xc3, 0x2d, 0x73, 0x7f, 0xcf, 0xff, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x78, 0x21, 0xc5, 0xf7, 0x1a, 0x03, 0x00, 0x00,
}

func (m *StrategyUpdateEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StrategyUpdateEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StrategyUpdateEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if len(m.QueryHash) > 0 {
		i -= len(m.QueryHash)
		copy(dAtA[i:], m.QueryHash)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.QueryHash)))
		i--
		dAtA[i] = 0x32
	}
	if m.Query != nil {
		{
			size, err := m.Query.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.TriggerPermission != nil {
		{
			size, err := m.TriggerPermission.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.IsEnabled != nil {
		{
			size, err := m.IsEnabled.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Id)))
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
func (m *StrategyUpdateEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Op != 0 {
		n += 1 + sovEvent(uint64(m.Op))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.IsEnabled != nil {
		l = m.IsEnabled.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.TriggerPermission != nil {
		l = m.TriggerPermission.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.Query != nil {
		l = m.Query.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.QueryHash)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StrategyUpdateEvent) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: StrategyUpdateEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StrategyUpdateEvent: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsEnabled", wireType)
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
			if m.IsEnabled == nil {
				m.IsEnabled = &types1.BoolValue{}
			}
			if err := m.IsEnabled.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TriggerPermission", wireType)
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
			if m.TriggerPermission == nil {
				m.TriggerPermission = &PermissionConfig{}
			}
			if err := m.TriggerPermission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
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
			if m.Query == nil {
				m.Query = &types2.FISQueryRequest{}
			}
			if err := m.Query.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryHash", wireType)
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
			m.QueryHash = append(m.QueryHash[:0], dAtA[iNdEx:postIndex]...)
			if m.QueryHash == nil {
				m.QueryHash = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
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
			if m.Metadata == nil {
				m.Metadata = &StrategyMetadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
