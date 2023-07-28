// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: flux/indexer/web3gw/query.proto

package web3gw

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SignRequest struct {
	TypedData string `protobuf:"bytes,1,opt,name=typed_data,json=typedData,proto3" json:"typed_data,omitempty"`
}

func (m *SignRequest) Reset()         { *m = SignRequest{} }
func (m *SignRequest) String() string { return proto.CompactTextString(m) }
func (*SignRequest) ProtoMessage()    {}
func (*SignRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb8009f4ae0c1d99, []int{0}
}
func (m *SignRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignRequest.Merge(m, src)
}
func (m *SignRequest) XXX_Size() int {
	return m.Size()
}
func (m *SignRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignRequest proto.InternalMessageInfo

func (m *SignRequest) GetTypedData() string {
	if m != nil {
		return m.TypedData
	}
	return ""
}

type SignResponse struct {
	Hash      string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Signature string `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *SignResponse) Reset()         { *m = SignResponse{} }
func (m *SignResponse) String() string { return proto.CompactTextString(m) }
func (*SignResponse) ProtoMessage()    {}
func (*SignResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb8009f4ae0c1d99, []int{1}
}
func (m *SignResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignResponse.Merge(m, src)
}
func (m *SignResponse) XXX_Size() int {
	return m.Size()
}
func (m *SignResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignResponse proto.InternalMessageInfo

func (m *SignResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *SignResponse) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func init() {
	proto.RegisterType((*SignRequest)(nil), "flux.indexer.web3gw.SignRequest")
	proto.RegisterType((*SignResponse)(nil), "flux.indexer.web3gw.SignResponse")
}

func init() { proto.RegisterFile("flux/indexer/web3gw/query.proto", fileDescriptor_eb8009f4ae0c1d99) }

var fileDescriptor_eb8009f4ae0c1d99 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4a, 0x03, 0x31,
	0x14, 0x86, 0x3b, 0x52, 0x0a, 0x8d, 0x6e, 0x8c, 0x2e, 0x4a, 0xd1, 0x58, 0xbb, 0x72, 0x21, 0x09,
	0xb5, 0x17, 0x10, 0x91, 0xba, 0xa8, 0xb8, 0xa8, 0x4a, 0xc1, 0x8d, 0x24, 0x33, 0xcf, 0x4c, 0xa0,
	0x4d, 0xda, 0x49, 0x62, 0xdb, 0x5b, 0x78, 0x2c, 0x97, 0x5d, 0xba, 0x94, 0x99, 0x8b, 0xc8, 0x4c,
	0x06, 0x11, 0x11, 0x37, 0xe1, 0xe7, 0xe3, 0xcb, 0xcf, 0x7b, 0x0f, 0x9d, 0xbc, 0xcc, 0xfc, 0x9a,
	0x29, 0x9d, 0xc0, 0x1a, 0x32, 0xb6, 0x02, 0x31, 0x94, 0x2b, 0xb6, 0xf4, 0x90, 0x6d, 0xe8, 0x22,
	0x33, 0xce, 0xe0, 0x83, 0x52, 0xa0, 0xb5, 0x40, 0x83, 0xd0, 0x25, 0xb1, 0xb1, 0x73, 0x63, 0x99,
	0xe0, 0x16, 0xd8, 0xeb, 0x40, 0x80, 0xe3, 0x03, 0x16, 0x1b, 0xa5, 0xc3, 0xa7, 0xee, 0xa1, 0x34,
	0xd2, 0x54, 0x91, 0x95, 0xa9, 0xa6, 0xfb, 0x7c, 0xae, 0xb4, 0x61, 0xd5, 0x1b, 0x50, 0xff, 0x1c,
	0xed, 0xde, 0x2b, 0xa9, 0x27, 0xb0, 0xf4, 0x60, 0x1d, 0x3e, 0x46, 0xc8, 0x6d, 0x16, 0x90, 0x3c,
	0x27, 0xdc, 0xf1, 0x4e, 0xd4, 0x8b, 0xce, 0xda, 0x93, 0x76, 0x45, 0xae, 0xb9, 0xe3, 0xfd, 0x4b,
	0xb4, 0x17, 0x6c, 0xbb, 0x30, 0xda, 0x02, 0xc6, 0xa8, 0x99, 0x72, 0x9b, 0xd6, 0x62, 0x95, 0xf1,
	0x11, 0x6a, 0x5b, 0x25, 0x35, 0x77, 0x3e, 0x83, 0xce, 0x4e, 0x68, 0xf8, 0x06, 0x17, 0x8f, 0xa8,
	0x35, 0x05, 0x31, 0xbc, 0x99, 0xe2, 0x31, 0x6a, 0x96, 0x5d, 0xb8, 0x47, 0xff, 0x58, 0x90, 0xfe,
	0x18, 0xaa, 0x7b, 0xfa, 0x8f, 0x11, 0x06, 0xb9, 0x1a, 0xbf, 0xe7, 0x24, 0xda, 0xe6, 0x24, 0xfa,
	0xcc, 0x49, 0xf4, 0x56, 0x90, 0xc6, 0xb6, 0x20, 0x8d, 0x8f, 0x82, 0x34, 0x9e, 0x06, 0x52, 0xb9,
	0xd4, 0x0b, 0x1a, 0x9b, 0x39, 0x1b, 0xcd, 0xfc, 0xfa, 0x6e, 0xf4, 0x70, 0xcb, 0x85, 0x65, 0x65,
	0x65, 0xc2, 0xe2, 0x94, 0x2b, 0xfd, 0xeb, 0xfa, 0xa2, 0x55, 0x9d, 0x66, 0xf8, 0x15, 0x00, 0x00,
	0xff, 0xff, 0x89, 0x1d, 0x6c, 0xb2, 0x9b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Web3GWClient is the client API for Web3GW service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Web3GWClient interface {
	Sign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error)
}

type web3GWClient struct {
	cc grpc1.ClientConn
}

func NewWeb3GWClient(cc grpc1.ClientConn) Web3GWClient {
	return &web3GWClient{cc}
}

func (c *web3GWClient) Sign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error) {
	out := new(SignResponse)
	err := c.cc.Invoke(ctx, "/flux.indexer.web3gw.Web3GW/Sign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Web3GWServer is the server API for Web3GW service.
type Web3GWServer interface {
	Sign(context.Context, *SignRequest) (*SignResponse, error)
}

// UnimplementedWeb3GWServer can be embedded to have forward compatible implementations.
type UnimplementedWeb3GWServer struct {
}

func (*UnimplementedWeb3GWServer) Sign(ctx context.Context, req *SignRequest) (*SignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sign not implemented")
}

func RegisterWeb3GWServer(s grpc1.Server, srv Web3GWServer) {
	s.RegisterService(&_Web3GW_serviceDesc, srv)
}

func _Web3GW_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Web3GWServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flux.indexer.web3gw.Web3GW/Sign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Web3GWServer).Sign(ctx, req.(*SignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Web3GW_serviceDesc = grpc.ServiceDesc{
	ServiceName: "flux.indexer.web3gw.Web3GW",
	HandlerType: (*Web3GWServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sign",
			Handler:    _Web3GW_Sign_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flux/indexer/web3gw/query.proto",
}

func (m *SignRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TypedData) > 0 {
		i -= len(m.TypedData)
		copy(dAtA[i:], m.TypedData)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.TypedData)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SignResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SignRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TypedData)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *SignResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SignRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: SignRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypedData", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TypedData = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *SignResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: SignResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
