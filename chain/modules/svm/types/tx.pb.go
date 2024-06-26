// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: flux/svm/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
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

type MsgLinkSVMAccount struct {
	Sender       string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	SvmPubkey    []byte `protobuf:"bytes,2,opt,name=svm_pubkey,json=svmPubkey,proto3" json:"svm_pubkey,omitempty"`
	SvmSignature []byte `protobuf:"bytes,4,opt,name=svm_signature,json=svmSignature,proto3" json:"svm_signature,omitempty"`
}

func (m *MsgLinkSVMAccount) Reset()         { *m = MsgLinkSVMAccount{} }
func (m *MsgLinkSVMAccount) String() string { return proto.CompactTextString(m) }
func (*MsgLinkSVMAccount) ProtoMessage()    {}
func (*MsgLinkSVMAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_6010395245cd7264, []int{0}
}
func (m *MsgLinkSVMAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgLinkSVMAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgLinkSVMAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgLinkSVMAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLinkSVMAccount.Merge(m, src)
}
func (m *MsgLinkSVMAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgLinkSVMAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLinkSVMAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLinkSVMAccount proto.InternalMessageInfo

type MsgLinkSVMAccountResponse struct {
}

func (m *MsgLinkSVMAccountResponse) Reset()         { *m = MsgLinkSVMAccountResponse{} }
func (m *MsgLinkSVMAccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgLinkSVMAccountResponse) ProtoMessage()    {}
func (*MsgLinkSVMAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6010395245cd7264, []int{1}
}
func (m *MsgLinkSVMAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgLinkSVMAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgLinkSVMAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgLinkSVMAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLinkSVMAccountResponse.Merge(m, src)
}
func (m *MsgLinkSVMAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgLinkSVMAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLinkSVMAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLinkSVMAccountResponse proto.InternalMessageInfo

type MsgTransaction struct {
	// senders are cosmos addresses that signs this message
	Signers       []string       `protobuf:"bytes,1,rep,name=signers,proto3" json:"signers,omitempty"`
	Accounts      []string       `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	Instructions  []*Instruction `protobuf:"bytes,3,rep,name=instructions,proto3" json:"instructions,omitempty"`
	ComputeBudget uint64         `protobuf:"varint,4,opt,name=compute_budget,json=computeBudget,proto3" json:"compute_budget,omitempty"`
}

func (m *MsgTransaction) Reset()         { *m = MsgTransaction{} }
func (m *MsgTransaction) String() string { return proto.CompactTextString(m) }
func (*MsgTransaction) ProtoMessage()    {}
func (*MsgTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_6010395245cd7264, []int{2}
}
func (m *MsgTransaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTransaction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTransaction.Merge(m, src)
}
func (m *MsgTransaction) XXX_Size() int {
	return m.Size()
}
func (m *MsgTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTransaction proto.InternalMessageInfo

type MsgTransactionResponse struct {
	UnitConsumed uint64 `protobuf:"varint,1,opt,name=unit_consumed,json=unitConsumed,proto3" json:"unit_consumed,omitempty"`
}

func (m *MsgTransactionResponse) Reset()         { *m = MsgTransactionResponse{} }
func (m *MsgTransactionResponse) String() string { return proto.CompactTextString(m) }
func (*MsgTransactionResponse) ProtoMessage()    {}
func (*MsgTransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6010395245cd7264, []int{3}
}
func (m *MsgTransactionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTransactionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTransactionResponse.Merge(m, src)
}
func (m *MsgTransactionResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgTransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTransactionResponse proto.InternalMessageInfo

func (m *MsgTransactionResponse) GetUnitConsumed() uint64 {
	if m != nil {
		return m.UnitConsumed
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgLinkSVMAccount)(nil), "flux.svm.v1beta1.MsgLinkSVMAccount")
	proto.RegisterType((*MsgLinkSVMAccountResponse)(nil), "flux.svm.v1beta1.MsgLinkSVMAccountResponse")
	proto.RegisterType((*MsgTransaction)(nil), "flux.svm.v1beta1.MsgTransaction")
	proto.RegisterType((*MsgTransactionResponse)(nil), "flux.svm.v1beta1.MsgTransactionResponse")
}

func init() { proto.RegisterFile("flux/svm/v1beta1/tx.proto", fileDescriptor_6010395245cd7264) }

var fileDescriptor_6010395245cd7264 = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x8d, 0x9b, 0x7e, 0xfd, 0xda, 0x21, 0x8d, 0xa8, 0x55, 0xc0, 0x31, 0xaa, 0x89, 0x52, 0x21,
	0x42, 0x10, 0x36, 0x09, 0xac, 0x2a, 0xb1, 0x48, 0x90, 0x2a, 0x21, 0x35, 0x15, 0x72, 0xaa, 0x2e,
	0xd8, 0x44, 0xfe, 0x19, 0x5c, 0xab, 0x99, 0x99, 0xc8, 0x77, 0xc6, 0x4a, 0x77, 0x88, 0x15, 0x62,
	0xc5, 0x23, 0xf0, 0x08, 0x59, 0xf4, 0x21, 0xba, 0xac, 0x58, 0xb1, 0x44, 0xc9, 0x22, 0x0f, 0xc0,
	0x0b, 0xa0, 0x71, 0xc6, 0xa1, 0x6d, 0x10, 0xb0, 0xb1, 0x3d, 0xe7, 0x9c, 0x7b, 0xe7, 0xdc, 0x33,
	0x1e, 0x54, 0x79, 0x37, 0x10, 0x23, 0x07, 0x52, 0xe2, 0xa4, 0x4d, 0x1f, 0x73, 0xaf, 0xe9, 0xf0,
	0x91, 0x3d, 0x4c, 0x18, 0x67, 0xfa, 0x6d, 0x49, 0xd9, 0x90, 0x12, 0x5b, 0x51, 0x66, 0x25, 0x60,
	0x40, 0x18, 0xf4, 0x33, 0xde, 0x99, 0x2f, 0xe6, 0x62, 0xf3, 0xde, 0x7c, 0xe5, 0x10, 0x88, 0x9c,
	0xb4, 0x29, 0x5f, 0x8a, 0xd8, 0x8e, 0x58, 0xc4, 0xe6, 0x05, 0xf2, 0x4b, 0xa1, 0x5b, 0x1e, 0x89,
	0x29, 0x73, 0xb2, 0xa7, 0x82, 0x2c, 0xd5, 0xc1, 0xf7, 0x00, 0x2f, 0xcc, 0x04, 0x2c, 0xa6, 0x8a,
	0x37, 0x97, 0x9c, 0x4a, 0x6b, 0x19, 0x57, 0x3b, 0xd7, 0xd0, 0x56, 0x17, 0xa2, 0x83, 0x98, 0x9e,
	0xf6, 0x8e, 0xbb, 0xed, 0x20, 0x60, 0x82, 0x72, 0xfd, 0x19, 0x5a, 0x03, 0x4c, 0x43, 0x9c, 0x18,
	0x5a, 0x55, 0xab, 0x6f, 0x74, 0x8c, 0xaf, 0xe7, 0x4f, 0xb7, 0x95, 0xeb, 0x76, 0x18, 0x26, 0x18,
	0xa0, 0xc7, 0x93, 0x98, 0x46, 0xae, 0xd2, 0xe9, 0x3b, 0x08, 0x41, 0x4a, 0xfa, 0x43, 0xe1, 0x9f,
	0xe2, 0x33, 0x63, 0xa5, 0xaa, 0xd5, 0x4b, 0xee, 0x06, 0xa4, 0xe4, 0x4d, 0x06, 0xe8, 0xbb, 0x68,
	0x53, 0xd2, 0x10, 0x47, 0xd4, 0xe3, 0x22, 0xc1, 0xc6, 0x6a, 0xa6, 0x28, 0x41, 0x4a, 0x7a, 0x39,
	0xb6, 0xf7, 0xf8, 0xe3, 0x97, 0x07, 0x85, 0x0f, 0xb3, 0x71, 0x43, 0x35, 0xfd, 0x34, 0x1b, 0x37,
	0xee, 0x48, 0xdb, 0x4b, 0x06, 0x6b, 0xf7, 0x51, 0x65, 0x09, 0x74, 0x31, 0x0c, 0x19, 0x05, 0x5c,
	0xfb, 0xa1, 0xa1, 0x72, 0x17, 0xa2, 0xa3, 0xc4, 0xa3, 0xe0, 0x05, 0x3c, 0x66, 0x54, 0x6f, 0xa1,
	0xff, 0xe5, 0xde, 0x38, 0x01, 0x43, 0xab, 0x16, 0xff, 0x38, 0x51, 0x2e, 0xd4, 0x4d, 0xb4, 0xee,
	0xcd, 0x3b, 0x83, 0xb1, 0x22, 0x8b, 0xdc, 0xc5, 0x5a, 0x6f, 0xa3, 0x52, 0x4c, 0x81, 0x27, 0x22,
	0x6b, 0x0f, 0x46, 0xb1, 0x5a, 0xac, 0xdf, 0x6a, 0xed, 0xd8, 0x37, 0x0f, 0xde, 0x7e, 0xfd, 0x4b,
	0xe5, 0x5e, 0x2b, 0xd1, 0x1f, 0xa2, 0x72, 0xc0, 0xc8, 0x50, 0x70, 0xdc, 0xf7, 0x45, 0x18, 0x61,
	0x9e, 0x65, 0xb2, 0xea, 0x6e, 0x2a, 0xb4, 0x93, 0x81, 0x7b, 0x8f, 0xf2, 0x50, 0x72, 0x5f, 0x32,
	0x15, 0x5d, 0xa5, 0x72, 0x65, 0xc4, 0xda, 0x4b, 0x74, 0xf7, 0x3a, 0x92, 0xe7, 0x21, 0xc3, 0x17,
	0x34, 0xe6, 0xfd, 0x80, 0x51, 0x10, 0x04, 0x87, 0xd9, 0xa1, 0xae, 0xba, 0x25, 0x09, 0xbe, 0x52,
	0x58, 0xeb, 0x42, 0x43, 0xc5, 0x2e, 0x44, 0xba, 0x8f, 0xca, 0x37, 0x7e, 0x86, 0xdd, 0xe5, 0xa9,
	0x96, 0xb2, 0x37, 0x9f, 0xfc, 0x83, 0x68, 0x61, 0xe8, 0x18, 0xad, 0xe7, 0x3e, 0xf5, 0xea, 0x6f,
	0x0b, 0xaf, 0x8c, 0x61, 0xd6, 0xff, 0xa6, 0xc8, 0xfb, 0x9a, 0xff, 0xbd, 0x9f, 0x8d, 0x1b, 0x5a,
	0xe7, 0xf0, 0x62, 0x62, 0x69, 0x97, 0x13, 0x4b, 0xfb, 0x3e, 0xb1, 0xb4, 0xcf, 0x53, 0xab, 0x70,
	0x39, 0xb5, 0x0a, 0xdf, 0xa6, 0x56, 0xe1, 0xed, 0x8b, 0x28, 0xe6, 0x27, 0xc2, 0xb7, 0x03, 0x46,
	0x9c, 0xfd, 0x81, 0x18, 0x1d, 0xee, 0x1f, 0x1d, 0x78, 0x3e, 0x38, 0x72, 0x83, 0xd0, 0x09, 0x4e,
	0xbc, 0x98, 0x3a, 0x84, 0x85, 0x62, 0x80, 0x21, 0xbb, 0x2f, 0xfc, 0x6c, 0x88, 0xc1, 0x5f, 0xcb,
	0xae, 0xca, 0xf3, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xff, 0xf6, 0x10, 0x7a, 0xf2, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	LinkSVMAccount(ctx context.Context, in *MsgLinkSVMAccount, opts ...grpc.CallOption) (*MsgLinkSVMAccountResponse, error)
	Transact(ctx context.Context, in *MsgTransaction, opts ...grpc.CallOption) (*MsgTransactionResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) LinkSVMAccount(ctx context.Context, in *MsgLinkSVMAccount, opts ...grpc.CallOption) (*MsgLinkSVMAccountResponse, error) {
	out := new(MsgLinkSVMAccountResponse)
	err := c.cc.Invoke(ctx, "/flux.svm.v1beta1.Msg/LinkSVMAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Transact(ctx context.Context, in *MsgTransaction, opts ...grpc.CallOption) (*MsgTransactionResponse, error) {
	out := new(MsgTransactionResponse)
	err := c.cc.Invoke(ctx, "/flux.svm.v1beta1.Msg/Transact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	LinkSVMAccount(context.Context, *MsgLinkSVMAccount) (*MsgLinkSVMAccountResponse, error)
	Transact(context.Context, *MsgTransaction) (*MsgTransactionResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) LinkSVMAccount(ctx context.Context, req *MsgLinkSVMAccount) (*MsgLinkSVMAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkSVMAccount not implemented")
}
func (*UnimplementedMsgServer) Transact(ctx context.Context, req *MsgTransaction) (*MsgTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transact not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_LinkSVMAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgLinkSVMAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).LinkSVMAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flux.svm.v1beta1.Msg/LinkSVMAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).LinkSVMAccount(ctx, req.(*MsgLinkSVMAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Transact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Transact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flux.svm.v1beta1.Msg/Transact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Transact(ctx, req.(*MsgTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "flux.svm.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LinkSVMAccount",
			Handler:    _Msg_LinkSVMAccount_Handler,
		},
		{
			MethodName: "Transact",
			Handler:    _Msg_Transact_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flux/svm/v1beta1/tx.proto",
}

func (m *MsgLinkSVMAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLinkSVMAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgLinkSVMAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SvmSignature) > 0 {
		i -= len(m.SvmSignature)
		copy(dAtA[i:], m.SvmSignature)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SvmSignature)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SvmPubkey) > 0 {
		i -= len(m.SvmPubkey)
		copy(dAtA[i:], m.SvmPubkey)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SvmPubkey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgLinkSVMAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLinkSVMAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgLinkSVMAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgTransaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTransaction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTransaction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ComputeBudget != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ComputeBudget))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Instructions) > 0 {
		for iNdEx := len(m.Instructions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Instructions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Accounts) > 0 {
		for iNdEx := len(m.Accounts) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Accounts[iNdEx])
			copy(dAtA[i:], m.Accounts[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Accounts[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signers[iNdEx])
			copy(dAtA[i:], m.Signers[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Signers[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgTransactionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTransactionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTransactionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UnitConsumed != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.UnitConsumed))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgLinkSVMAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SvmPubkey)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SvmSignature)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgLinkSVMAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgTransaction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Signers) > 0 {
		for _, s := range m.Signers {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if len(m.Accounts) > 0 {
		for _, s := range m.Accounts {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if len(m.Instructions) > 0 {
		for _, e := range m.Instructions {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if m.ComputeBudget != 0 {
		n += 1 + sovTx(uint64(m.ComputeBudget))
	}
	return n
}

func (m *MsgTransactionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UnitConsumed != 0 {
		n += 1 + sovTx(uint64(m.UnitConsumed))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgLinkSVMAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgLinkSVMAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLinkSVMAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SvmPubkey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SvmPubkey = append(m.SvmPubkey[:0], dAtA[iNdEx:postIndex]...)
			if m.SvmPubkey == nil {
				m.SvmPubkey = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SvmSignature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SvmSignature = append(m.SvmSignature[:0], dAtA[iNdEx:postIndex]...)
			if m.SvmSignature == nil {
				m.SvmSignature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgLinkSVMAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgLinkSVMAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLinkSVMAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgTransaction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgTransaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTransaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accounts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Accounts = append(m.Accounts, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Instructions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Instructions = append(m.Instructions, &Instruction{})
			if err := m.Instructions[len(m.Instructions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ComputeBudget", wireType)
			}
			m.ComputeBudget = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ComputeBudget |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgTransactionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgTransactionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTransactionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnitConsumed", wireType)
			}
			m.UnitConsumed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnitConsumed |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
