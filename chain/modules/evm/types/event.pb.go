// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: flux/evm/v1beta1/event.proto

package types

import (
	fmt "fmt"
	types "github.com/FluxNFTLabs/sdk-go/chain/eventstream/types"
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

type DeployEvent struct {
	Op       types.Op      `protobuf:"varint,1,opt,name=op,proto3,enum=flux.eventstream.v1beta1.Op" json:"op,omitempty"`
	Contract *ContractInfo `protobuf:"bytes,2,opt,name=contract,proto3" json:"contract,omitempty"`
}

func (m *DeployEvent) Reset()         { *m = DeployEvent{} }
func (m *DeployEvent) String() string { return proto.CompactTextString(m) }
func (*DeployEvent) ProtoMessage()    {}
func (*DeployEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_a86c064bb23b7bf3, []int{0}
}
func (m *DeployEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeployEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeployEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeployEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployEvent.Merge(m, src)
}
func (m *DeployEvent) XXX_Size() int {
	return m.Size()
}
func (m *DeployEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployEvent.DiscardUnknown(m)
}

var xxx_messageInfo_DeployEvent proto.InternalMessageInfo

func (m *DeployEvent) GetOp() types.Op {
	if m != nil {
		return m.Op
	}
	return types.Op_COSMOS_SET
}

func (m *DeployEvent) GetContract() *ContractInfo {
	if m != nil {
		return m.Contract
	}
	return nil
}

type ExecuteEvent struct {
	Op      types.Op `protobuf:"varint,1,opt,name=op,proto3,enum=flux.eventstream.v1beta1.Op" json:"op,omitempty"`
	Address string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *ExecuteEvent) Reset()         { *m = ExecuteEvent{} }
func (m *ExecuteEvent) String() string { return proto.CompactTextString(m) }
func (*ExecuteEvent) ProtoMessage()    {}
func (*ExecuteEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_a86c064bb23b7bf3, []int{1}
}
func (m *ExecuteEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExecuteEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExecuteEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExecuteEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteEvent.Merge(m, src)
}
func (m *ExecuteEvent) XXX_Size() int {
	return m.Size()
}
func (m *ExecuteEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteEvent proto.InternalMessageInfo

func (m *ExecuteEvent) GetOp() types.Op {
	if m != nil {
		return m.Op
	}
	return types.Op_COSMOS_SET
}

func (m *ExecuteEvent) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type EmitLogEvent struct {
	Op      types.Op `protobuf:"varint,1,opt,name=op,proto3,enum=flux.eventstream.v1beta1.Op" json:"op,omitempty"`
	Address string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Topics  [][]byte `protobuf:"bytes,3,rep,name=topics,proto3" json:"topics,omitempty"`
	Data    []byte   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *EmitLogEvent) Reset()         { *m = EmitLogEvent{} }
func (m *EmitLogEvent) String() string { return proto.CompactTextString(m) }
func (*EmitLogEvent) ProtoMessage()    {}
func (*EmitLogEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_a86c064bb23b7bf3, []int{2}
}
func (m *EmitLogEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EmitLogEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EmitLogEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EmitLogEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmitLogEvent.Merge(m, src)
}
func (m *EmitLogEvent) XXX_Size() int {
	return m.Size()
}
func (m *EmitLogEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_EmitLogEvent.DiscardUnknown(m)
}

var xxx_messageInfo_EmitLogEvent proto.InternalMessageInfo

func (m *EmitLogEvent) GetOp() types.Op {
	if m != nil {
		return m.Op
	}
	return types.Op_COSMOS_SET
}

func (m *EmitLogEvent) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *EmitLogEvent) GetTopics() [][]byte {
	if m != nil {
		return m.Topics
	}
	return nil
}

func (m *EmitLogEvent) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*DeployEvent)(nil), "flux.evm.v1beta1.DeployEvent")
	proto.RegisterType((*ExecuteEvent)(nil), "flux.evm.v1beta1.ExecuteEvent")
	proto.RegisterType((*EmitLogEvent)(nil), "flux.evm.v1beta1.EmitLogEvent")
}

func init() { proto.RegisterFile("flux/evm/v1beta1/event.proto", fileDescriptor_a86c064bb23b7bf3) }

var fileDescriptor_a86c064bb23b7bf3 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0x31, 0x4b, 0xc3, 0x40,
	0x14, 0xc7, 0x7b, 0x6d, 0xa9, 0x7a, 0x0d, 0x22, 0x19, 0x24, 0x94, 0x72, 0x84, 0xe2, 0x90, 0x41,
	0xee, 0x68, 0x75, 0x72, 0x54, 0x5b, 0x10, 0x4a, 0x85, 0x20, 0x0e, 0x6e, 0x97, 0xe4, 0xb5, 0x0d,
	0x24, 0xb9, 0x33, 0x77, 0x89, 0xed, 0xea, 0x27, 0xf0, 0x63, 0x39, 0x76, 0x74, 0x94, 0xf6, 0x8b,
	0x48, 0xd2, 0xb4, 0x82, 0x38, 0x89, 0xdb, 0x7b, 0xfc, 0x7f, 0xbc, 0xdf, 0x83, 0x3f, 0xee, 0x4e,
	0xa3, 0x6c, 0xc1, 0x20, 0x8f, 0x59, 0xde, 0xf7, 0x40, 0xf3, 0x3e, 0x83, 0x1c, 0x12, 0x4d, 0x65,
	0x2a, 0xb4, 0x30, 0x4f, 0x8a, 0x94, 0x42, 0x1e, 0xd3, 0x2a, 0xed, 0x74, 0x7e, 0xe1, 0xe3, 0x2d,
	0xdd, 0x39, 0xab, 0x32, 0x48, 0xb4, 0xd2, 0x29, 0xf0, 0x6f, 0xe6, 0x39, 0x83, 0x74, 0xb9, 0xa5,
	0x7a, 0x2f, 0xb8, 0x7d, 0x0b, 0x32, 0x12, 0xcb, 0x61, 0x01, 0x9a, 0xe7, 0xb8, 0x2e, 0xa4, 0x85,
	0x6c, 0xe4, 0x1c, 0x0f, 0xba, 0xb4, 0xf2, 0xed, 0x2f, 0xec, 0xbc, 0xf4, 0x5e, 0xba, 0x75, 0x21,
	0xcd, 0x2b, 0x7c, 0xe8, 0x8b, 0x44, 0xa7, 0xdc, 0xd7, 0x56, 0xdd, 0x46, 0x4e, 0x7b, 0x40, 0xe8,
	0xcf, 0x1f, 0xe9, 0x4d, 0x45, 0xdc, 0x25, 0x53, 0xe1, 0xee, 0xf9, 0xde, 0x23, 0x36, 0x86, 0x0b,
	0xf0, 0x33, 0x0d, 0x7f, 0x31, 0x5b, 0xf8, 0x80, 0x07, 0x41, 0x0a, 0x4a, 0x95, 0xe2, 0x23, 0x77,
	0xb7, 0xf6, 0x5e, 0x11, 0x36, 0x86, 0x71, 0xa8, 0xc7, 0x62, 0xf6, 0xaf, 0x87, 0xcd, 0x53, 0xdc,
	0xd2, 0x42, 0x86, 0xbe, 0xb2, 0x1a, 0x76, 0xc3, 0x31, 0xdc, 0x6a, 0x33, 0x4d, 0xdc, 0x0c, 0xb8,
	0xe6, 0x56, 0xd3, 0x46, 0x8e, 0xe1, 0x96, 0xf3, 0xf5, 0xe4, 0x7d, 0x4d, 0xd0, 0x6a, 0x4d, 0xd0,
	0xe7, 0x9a, 0xa0, 0xb7, 0x0d, 0xa9, 0xad, 0x36, 0xa4, 0xf6, 0xb1, 0x21, 0xb5, 0xa7, 0xcb, 0x59,
	0xa8, 0xe7, 0x99, 0x47, 0x7d, 0x11, 0xb3, 0x51, 0x94, 0x2d, 0x26, 0xa3, 0x87, 0x31, 0xf7, 0x14,
	0x2b, 0xfe, 0x0a, 0x98, 0x3f, 0xe7, 0x61, 0xc2, 0x62, 0x11, 0x64, 0x11, 0xa8, 0xb2, 0x57, 0xbd,
	0x94, 0xa0, 0xbc, 0x56, 0x59, 0xd6, 0xc5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x36, 0x53, 0xb3,
	0xc4, 0x20, 0x02, 0x00, 0x00,
}

func (m *DeployEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeployEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeployEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Contract != nil {
		{
			size, err := m.Contract.MarshalToSizedBuffer(dAtA[:i])
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

func (m *ExecuteEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExecuteEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExecuteEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Address)))
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

func (m *EmitLogEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EmitLogEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EmitLogEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Topics) > 0 {
		for iNdEx := len(m.Topics) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Topics[iNdEx])
			copy(dAtA[i:], m.Topics[iNdEx])
			i = encodeVarintEvent(dAtA, i, uint64(len(m.Topics[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Address)))
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
func (m *DeployEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Op != 0 {
		n += 1 + sovEvent(uint64(m.Op))
	}
	if m.Contract != nil {
		l = m.Contract.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *ExecuteEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Op != 0 {
		n += 1 + sovEvent(uint64(m.Op))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *EmitLogEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Op != 0 {
		n += 1 + sovEvent(uint64(m.Op))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if len(m.Topics) > 0 {
		for _, b := range m.Topics {
			l = len(b)
			n += 1 + l + sovEvent(uint64(l))
		}
	}
	l = len(m.Data)
	if l > 0 {
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
func (m *DeployEvent) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: DeployEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeployEvent: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
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
			if m.Contract == nil {
				m.Contract = &ContractInfo{}
			}
			if err := m.Contract.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *ExecuteEvent) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ExecuteEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExecuteEvent: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
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
func (m *EmitLogEvent) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EmitLogEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EmitLogEvent: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topics", wireType)
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
			m.Topics = append(m.Topics, make([]byte, postIndex-iNdEx))
			copy(m.Topics[len(m.Topics)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
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
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
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
