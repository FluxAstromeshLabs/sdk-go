// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: flux/astromesh/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
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

type GenesisState struct {
	DenomLinks []*DenomLink `protobuf:"bytes,1,rep,name=denom_links,json=denomLinks,proto3" json:"denom_links,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_56a2028c7fe76419, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetDenomLinks() []*DenomLink {
	if m != nil {
		return m.DenomLinks
	}
	return nil
}

type DenomLink struct {
	SrcPlane    Plane  `protobuf:"varint,1,opt,name=src_plane,json=srcPlane,proto3,enum=flux.astromesh.v1beta1.Plane" json:"src_plane,omitempty"`
	DstPlane    Plane  `protobuf:"varint,2,opt,name=dst_plane,json=dstPlane,proto3,enum=flux.astromesh.v1beta1.Plane" json:"dst_plane,omitempty"`
	SrcAddr     string `protobuf:"bytes,3,opt,name=src_addr,json=srcAddr,proto3" json:"src_addr,omitempty"`
	DstAddr     string `protobuf:"bytes,4,opt,name=dst_addr,json=dstAddr,proto3" json:"dst_addr,omitempty"`
	SrcDecimals int32  `protobuf:"varint,5,opt,name=src_decimals,json=srcDecimals,proto3" json:"src_decimals,omitempty"`
	DstDecimals int32  `protobuf:"varint,6,opt,name=dst_decimals,json=dstDecimals,proto3" json:"dst_decimals,omitempty"`
}

func (m *DenomLink) Reset()         { *m = DenomLink{} }
func (m *DenomLink) String() string { return proto.CompactTextString(m) }
func (*DenomLink) ProtoMessage()    {}
func (*DenomLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_56a2028c7fe76419, []int{1}
}
func (m *DenomLink) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DenomLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DenomLink.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DenomLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DenomLink.Merge(m, src)
}
func (m *DenomLink) XXX_Size() int {
	return m.Size()
}
func (m *DenomLink) XXX_DiscardUnknown() {
	xxx_messageInfo_DenomLink.DiscardUnknown(m)
}

var xxx_messageInfo_DenomLink proto.InternalMessageInfo

func (m *DenomLink) GetSrcPlane() Plane {
	if m != nil {
		return m.SrcPlane
	}
	return Plane_COSMOS
}

func (m *DenomLink) GetDstPlane() Plane {
	if m != nil {
		return m.DstPlane
	}
	return Plane_COSMOS
}

func (m *DenomLink) GetSrcAddr() string {
	if m != nil {
		return m.SrcAddr
	}
	return ""
}

func (m *DenomLink) GetDstAddr() string {
	if m != nil {
		return m.DstAddr
	}
	return ""
}

func (m *DenomLink) GetSrcDecimals() int32 {
	if m != nil {
		return m.SrcDecimals
	}
	return 0
}

func (m *DenomLink) GetDstDecimals() int32 {
	if m != nil {
		return m.DstDecimals
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "flux.astromesh.v1beta1.GenesisState")
	proto.RegisterType((*DenomLink)(nil), "flux.astromesh.v1beta1.DenomLink")
}

func init() {
	proto.RegisterFile("flux/astromesh/v1beta1/genesis.proto", fileDescriptor_56a2028c7fe76419)
}

var fileDescriptor_56a2028c7fe76419 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x3f, 0x4f, 0xc2, 0x40,
	0x18, 0xc6, 0x39, 0x11, 0x94, 0x83, 0x38, 0x30, 0x98, 0x6a, 0x62, 0x2d, 0xc4, 0xa1, 0x53, 0x2f,
	0xe0, 0xc6, 0x26, 0x21, 0xb8, 0x10, 0x63, 0x2a, 0x93, 0x0b, 0xb9, 0xde, 0x7b, 0x42, 0x43, 0xff,
	0x90, 0x7b, 0xaf, 0x06, 0xbf, 0x85, 0x1f, 0xcb, 0x91, 0xd1, 0xd1, 0xc0, 0x07, 0xd1, 0x5c, 0x0b,
	0xc4, 0x41, 0x12, 0xb7, 0x37, 0xf7, 0xfc, 0x7e, 0x4f, 0x72, 0x79, 0xe8, 0xcd, 0x4b, 0x94, 0x2d,
	0x19, 0x47, 0xad, 0xd2, 0x58, 0xe2, 0x8c, 0xbd, 0x76, 0x02, 0xa9, 0x79, 0x87, 0x4d, 0x65, 0x22,
	0x31, 0x44, 0x6f, 0xa1, 0x52, 0x9d, 0x36, 0xcf, 0x0d, 0xe5, 0xed, 0x29, 0x6f, 0x4b, 0x5d, 0x5e,
	0x1f, 0xb0, 0xf5, 0xb2, 0x10, 0xdb, 0x3e, 0x6d, 0xdc, 0x17, 0x4d, 0x4f, 0x9a, 0x6b, 0xd9, 0xec,
	0xd3, 0x3a, 0xc8, 0x24, 0x8d, 0x27, 0x51, 0x98, 0xcc, 0xd1, 0x22, 0x4e, 0xd9, 0xad, 0x77, 0x5b,
	0xde, 0xdf, 0xf5, 0xde, 0xc0, 0xa0, 0xa3, 0x30, 0x99, 0xfb, 0x14, 0x76, 0x27, 0xb6, 0xbf, 0x09,
	0xad, 0xed, 0x93, 0x66, 0x8f, 0xd6, 0x50, 0x89, 0xc9, 0x22, 0xe2, 0x89, 0xb4, 0x88, 0x43, 0xdc,
	0xb3, 0xee, 0xd5, 0xa1, 0xbe, 0x47, 0x03, 0xf9, 0xa7, 0xa8, 0x44, 0x7e, 0x19, 0x17, 0x50, 0x6f,
	0xdd, 0xa3, 0x7f, 0xb9, 0x80, 0xba, 0x70, 0x2f, 0xa8, 0xe9, 0x99, 0x70, 0x00, 0x65, 0x95, 0x1d,
	0xe2, 0xd6, 0xfc, 0x13, 0x54, 0xe2, 0x0e, 0x40, 0x99, 0xc8, 0xd4, 0xe6, 0xd1, 0x71, 0x11, 0x01,
	0xea, 0x3c, 0x6a, 0xd1, 0x86, 0xb1, 0x40, 0x8a, 0x30, 0xe6, 0x11, 0x5a, 0x15, 0x87, 0xb8, 0x15,
	0xbf, 0x8e, 0x4a, 0x0c, 0xb6, 0x4f, 0x06, 0x31, 0xf6, 0x1e, 0xa9, 0x16, 0x08, 0xa0, 0xde, 0x21,
	0xfd, 0xf1, 0xc7, 0xda, 0x26, 0xab, 0xb5, 0x4d, 0xbe, 0xd6, 0x36, 0x79, 0xdf, 0xd8, 0xa5, 0xd5,
	0xc6, 0x2e, 0x7d, 0x6e, 0xec, 0xd2, 0x73, 0x6f, 0x1a, 0xea, 0x59, 0x16, 0x78, 0x22, 0x8d, 0xd9,
	0x30, 0xca, 0x96, 0x0f, 0xc3, 0xf1, 0x88, 0x07, 0xc8, 0xcc, 0xa7, 0x80, 0x89, 0x19, 0x0f, 0x13,
	0x16, 0xa7, 0x90, 0x45, 0x12, 0x7f, 0xcd, 0xa6, 0xdf, 0x16, 0x12, 0x83, 0x6a, 0x3e, 0xd9, 0xed,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xba, 0xa3, 0x71, 0xf8, 0x13, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DenomLinks) > 0 {
		for iNdEx := len(m.DenomLinks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DenomLinks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *DenomLink) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DenomLink) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DenomLink) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DstDecimals != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.DstDecimals))
		i--
		dAtA[i] = 0x30
	}
	if m.SrcDecimals != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SrcDecimals))
		i--
		dAtA[i] = 0x28
	}
	if len(m.DstAddr) > 0 {
		i -= len(m.DstAddr)
		copy(dAtA[i:], m.DstAddr)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.DstAddr)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SrcAddr) > 0 {
		i -= len(m.SrcAddr)
		copy(dAtA[i:], m.SrcAddr)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.SrcAddr)))
		i--
		dAtA[i] = 0x1a
	}
	if m.DstPlane != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.DstPlane))
		i--
		dAtA[i] = 0x10
	}
	if m.SrcPlane != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SrcPlane))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DenomLinks) > 0 {
		for _, e := range m.DenomLinks {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *DenomLink) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SrcPlane != 0 {
		n += 1 + sovGenesis(uint64(m.SrcPlane))
	}
	if m.DstPlane != 0 {
		n += 1 + sovGenesis(uint64(m.DstPlane))
	}
	l = len(m.SrcAddr)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.DstAddr)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.SrcDecimals != 0 {
		n += 1 + sovGenesis(uint64(m.SrcDecimals))
	}
	if m.DstDecimals != 0 {
		n += 1 + sovGenesis(uint64(m.DstDecimals))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomLinks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomLinks = append(m.DenomLinks, &DenomLink{})
			if err := m.DenomLinks[len(m.DenomLinks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *DenomLink) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: DenomLink: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DenomLink: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcPlane", wireType)
			}
			m.SrcPlane = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SrcPlane |= Plane(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DstPlane", wireType)
			}
			m.DstPlane = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DstPlane |= Plane(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SrcAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DstAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DstAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcDecimals", wireType)
			}
			m.SrcDecimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SrcDecimals |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DstDecimals", wireType)
			}
			m.DstDecimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DstDecimals |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
