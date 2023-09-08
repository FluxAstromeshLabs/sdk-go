// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: flux/stream/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/FluxNFTLabs/sdk-go/chain/modules/bazaar/types"
	_ "github.com/FluxNFTLabs/sdk-go/chain/modules/fnft/types"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Op int32

const (
	Op_Update Op = 0
	Op_Delete Op = 1
)

var Op_name = map[int32]string{
	0: "Update",
	1: "Delete",
}

var Op_value = map[string]int32{
	"Update": 0,
	"Delete": 1,
}

func (x Op) String() string {
	return proto.EnumName(Op_name, int32(x))
}

func (Op) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a82db9e8fada1d29, []int{0}
}

type EventsRequest struct {
	Height    uint64   `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Modules   []string `protobuf:"bytes,2,rep,name=modules,proto3" json:"modules,omitempty"`
	TmQueries []string `protobuf:"bytes,3,rep,name=tm_queries,json=tmQueries,proto3" json:"tm_queries,omitempty"`
}

func (m *EventsRequest) Reset()         { *m = EventsRequest{} }
func (m *EventsRequest) String() string { return proto.CompactTextString(m) }
func (*EventsRequest) ProtoMessage()    {}
func (*EventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a82db9e8fada1d29, []int{0}
}
func (m *EventsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventsRequest.Merge(m, src)
}
func (m *EventsRequest) XXX_Size() int {
	return m.Size()
}
func (m *EventsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventsRequest proto.InternalMessageInfo

func (m *EventsRequest) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *EventsRequest) GetModules() []string {
	if m != nil {
		return m.Modules
	}
	return nil
}

func (m *EventsRequest) GetTmQueries() []string {
	if m != nil {
		return m.TmQueries
	}
	return nil
}

type EventsResponse struct {
	Height    uint64    `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Time      uint64    `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	Modules   []string  `protobuf:"bytes,3,rep,name=modules,proto3" json:"modules,omitempty"`
	Events    []*Events `protobuf:"bytes,4,rep,name=events,proto3" json:"events,omitempty"`
	TmQueries []string  `protobuf:"bytes,5,rep,name=tm_queries,json=tmQueries,proto3" json:"tm_queries,omitempty"`
	TmData    []string  `protobuf:"bytes,6,rep,name=tm_data,json=tmData,proto3" json:"tm_data,omitempty"`
}

func (m *EventsResponse) Reset()         { *m = EventsResponse{} }
func (m *EventsResponse) String() string { return proto.CompactTextString(m) }
func (*EventsResponse) ProtoMessage()    {}
func (*EventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a82db9e8fada1d29, []int{1}
}
func (m *EventsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventsResponse.Merge(m, src)
}
func (m *EventsResponse) XXX_Size() int {
	return m.Size()
}
func (m *EventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventsResponse proto.InternalMessageInfo

func (m *EventsResponse) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *EventsResponse) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *EventsResponse) GetModules() []string {
	if m != nil {
		return m.Modules
	}
	return nil
}

func (m *EventsResponse) GetEvents() []*Events {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *EventsResponse) GetTmQueries() []string {
	if m != nil {
		return m.TmQueries
	}
	return nil
}

func (m *EventsResponse) GetTmData() []string {
	if m != nil {
		return m.TmData
	}
	return nil
}

type Events struct {
	EventOps []*EventOp `protobuf:"bytes,1,rep,name=event_ops,json=eventOps,proto3" json:"event_ops,omitempty"`
}

func (m *Events) Reset()         { *m = Events{} }
func (m *Events) String() string { return proto.CompactTextString(m) }
func (*Events) ProtoMessage()    {}
func (*Events) Descriptor() ([]byte, []int) {
	return fileDescriptor_a82db9e8fada1d29, []int{2}
}
func (m *Events) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Events) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Events.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Events) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Events.Merge(m, src)
}
func (m *Events) XXX_Size() int {
	return m.Size()
}
func (m *Events) XXX_DiscardUnknown() {
	xxx_messageInfo_Events.DiscardUnknown(m)
}

var xxx_messageInfo_Events proto.InternalMessageInfo

func (m *Events) GetEventOps() []*EventOp {
	if m != nil {
		return m.EventOps
	}
	return nil
}

type EventOp struct {
	Op   Op         `protobuf:"varint,1,opt,name=op,proto3,enum=flux.stream.v1beta1.Op" json:"op,omitempty"`
	Data *types.Any `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *EventOp) Reset()         { *m = EventOp{} }
func (m *EventOp) String() string { return proto.CompactTextString(m) }
func (*EventOp) ProtoMessage()    {}
func (*EventOp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a82db9e8fada1d29, []int{3}
}
func (m *EventOp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventOp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventOp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventOp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventOp.Merge(m, src)
}
func (m *EventOp) XXX_Size() int {
	return m.Size()
}
func (m *EventOp) XXX_DiscardUnknown() {
	xxx_messageInfo_EventOp.DiscardUnknown(m)
}

var xxx_messageInfo_EventOp proto.InternalMessageInfo

func (m *EventOp) GetOp() Op {
	if m != nil {
		return m.Op
	}
	return Op_Update
}

func (m *EventOp) GetData() *types.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("flux.stream.v1beta1.Op", Op_name, Op_value)
	proto.RegisterType((*EventsRequest)(nil), "flux.stream.v1beta1.EventsRequest")
	proto.RegisterType((*EventsResponse)(nil), "flux.stream.v1beta1.EventsResponse")
	proto.RegisterType((*Events)(nil), "flux.stream.v1beta1.Events")
	proto.RegisterType((*EventOp)(nil), "flux.stream.v1beta1.EventOp")
}

func init() { proto.RegisterFile("flux/stream/v1beta1/query.proto", fileDescriptor_a82db9e8fada1d29) }

var fileDescriptor_a82db9e8fada1d29 = []byte{
	// 548 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xce, 0x39, 0xa9, 0x43, 0xae, 0x50, 0x55, 0x07, 0xa2, 0x26, 0x0d, 0x26, 0x98, 0x81, 0x88,
	0xc1, 0x47, 0xd3, 0x09, 0x21, 0x86, 0x92, 0x52, 0x54, 0x09, 0x11, 0x61, 0x60, 0x41, 0x48, 0xe1,
	0x9c, 0x5c, 0x12, 0x4b, 0xb1, 0xef, 0x9a, 0x7b, 0xae, 0x1a, 0xa2, 0x2c, 0xfc, 0x05, 0x48, 0x4c,
	0xec, 0xfc, 0x27, 0x2c, 0x88, 0xa9, 0x12, 0x0b, 0x23, 0x4a, 0xf8, 0x43, 0x50, 0xce, 0x97, 0x02,
	0x55, 0x5a, 0x31, 0xb0, 0xdd, 0x7b, 0xdf, 0xe7, 0xf7, 0x7d, 0xef, 0x87, 0xf1, 0x8d, 0xee, 0x20,
	0x3d, 0xa2, 0x0a, 0x86, 0x9c, 0xc5, 0xf4, 0x70, 0x2b, 0xe4, 0xc0, 0xb6, 0xe8, 0x41, 0xca, 0x87,
	0x23, 0x5f, 0x0e, 0x05, 0x08, 0x72, 0x79, 0x4e, 0xf0, 0x33, 0x82, 0x6f, 0x08, 0xe5, 0x6b, 0x3d,
	0x21, 0x7a, 0x03, 0x4e, 0x35, 0x25, 0x4c, 0xbb, 0x94, 0x25, 0x86, 0x5f, 0xae, 0x18, 0x88, 0xc9,
	0x88, 0xb2, 0x24, 0x11, 0xc0, 0x20, 0x12, 0x89, 0x32, 0xe8, 0xa6, 0x96, 0xeb, 0x26, 0x5d, 0x38,
	0x11, 0x4b, 0xba, 0x60, 0xc0, 0x9b, 0x1a, 0x0c, 0xd9, 0x5b, 0xc6, 0x86, 0x27, 0xb0, 0x1c, 0x8a,
	0x4e, 0xda, 0x36, 0x14, 0xef, 0x0d, 0xbe, 0xf4, 0xe8, 0x90, 0x27, 0xa0, 0x02, 0x7e, 0x90, 0x72,
	0x05, 0xe4, 0x2a, 0xb6, 0xfb, 0x3c, 0xea, 0xf5, 0xc1, 0x41, 0x55, 0x54, 0x2b, 0x04, 0x26, 0x22,
	0x0e, 0x2e, 0xc6, 0xa2, 0x93, 0x0e, 0xb8, 0x72, 0xac, 0x6a, 0xbe, 0x56, 0x0a, 0x16, 0x21, 0xb9,
	0x8e, 0x31, 0xc4, 0xad, 0x79, 0x8b, 0x11, 0x57, 0x4e, 0x5e, 0x83, 0x25, 0x88, 0x9f, 0x65, 0x09,
	0xef, 0x33, 0xc2, 0x6b, 0x0b, 0x09, 0x25, 0x45, 0xa2, 0xf8, 0x99, 0x1a, 0x04, 0x17, 0x20, 0x8a,
	0xb9, 0x63, 0xe9, 0xac, 0x7e, 0xff, 0xa9, 0x9b, 0xff, 0x5b, 0x77, 0x1b, 0xdb, 0x5c, 0xd7, 0x75,
	0x0a, 0xd5, 0x7c, 0x6d, 0xb5, 0xbe, 0xe9, 0x2f, 0x99, 0xac, 0x6f, 0xa4, 0x0d, 0xf5, 0x94, 0xd9,
	0x95, 0x53, 0x66, 0xc9, 0x06, 0x2e, 0x42, 0xdc, 0xea, 0x30, 0x60, 0x8e, 0xad, 0x31, 0x1b, 0xe2,
	0x5d, 0x06, 0xcc, 0x6b, 0x60, 0x3b, 0xab, 0x44, 0xee, 0xe1, 0x92, 0xae, 0xd5, 0x12, 0x52, 0x39,
	0x48, 0x2b, 0x57, 0xce, 0x56, 0x6e, 0xca, 0xe0, 0x02, 0xcf, 0x1e, 0xca, 0x7b, 0x8d, 0x8b, 0x26,
	0x49, 0x6e, 0x63, 0x4b, 0x48, 0xdd, 0xfe, 0x5a, 0x7d, 0x63, 0xe9, 0xe7, 0x4d, 0x19, 0x58, 0x42,
	0x92, 0x1a, 0x2e, 0x68, 0x3b, 0xf3, 0x99, 0xac, 0xd6, 0xaf, 0xf8, 0xd9, 0x35, 0xf8, 0x8b, 0x43,
	0xf1, 0x77, 0x92, 0x51, 0xa0, 0x19, 0x77, 0x2a, 0xd8, 0x6a, 0x4a, 0x82, 0xb1, 0xfd, 0x52, 0x76,
	0x18, 0xf0, 0xf5, 0xdc, 0xfc, 0xbd, 0xcb, 0x07, 0x1c, 0xf8, 0x3a, 0xaa, 0x7f, 0xb5, 0xf0, 0xca,
	0xbc, 0xcb, 0x11, 0xf9, 0x88, 0x70, 0xe9, 0x31, 0x07, 0xd3, 0x8e, 0x77, 0xde, 0xd4, 0xb2, 0x9b,
	0x28, 0xdf, 0x3a, 0x97, 0x93, 0x2d, 0xd5, 0x6b, 0xbc, 0xfb, 0xf6, 0xf3, 0x83, 0xf5, 0x80, 0xdc,
	0xa7, 0xcb, 0xfe, 0x80, 0x6c, 0xfc, 0x74, 0x9c, 0x6d, 0x7a, 0x42, 0xc7, 0x66, 0x8b, 0x13, 0x3a,
	0xfe, 0xbd, 0x8f, 0x09, 0xf9, 0x84, 0xf0, 0xc5, 0xe7, 0xfa, 0xcb, 0xff, 0x6d, 0x6f, 0x5f, 0xdb,
	0x6b, 0x90, 0x9d, 0xa5, 0xf6, 0xb2, 0xf0, 0xdf, 0x4c, 0xde, 0x45, 0x0f, 0xf7, 0xbf, 0x4c, 0x5d,
	0x74, 0x3c, 0x75, 0xd1, 0x8f, 0xa9, 0x8b, 0xde, 0xcf, 0xdc, 0xdc, 0xf1, 0xcc, 0xcd, 0x7d, 0x9f,
	0xb9, 0xb9, 0x57, 0xb4, 0x17, 0x41, 0x3f, 0x0d, 0xfd, 0xb6, 0x88, 0xe9, 0xde, 0x20, 0x3d, 0x7a,
	0xba, 0xf7, 0xe2, 0x09, 0x0b, 0x95, 0x16, 0xed, 0xd0, 0x76, 0x9f, 0x45, 0xc9, 0x42, 0x1b, 0x46,
	0x92, 0xab, 0xd0, 0xd6, 0x9b, 0xdc, 0xfe, 0x15, 0x00, 0x00, 0xff, 0xff, 0xca, 0x92, 0x7b, 0xcc,
	0x38, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	GetEvents(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (*EventsResponse, error)
	StreamEvents(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (Query_StreamEventsClient, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) GetEvents(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/flux.stream.v1beta1.Query/GetEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) StreamEvents(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (Query_StreamEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Query_serviceDesc.Streams[0], "/flux.stream.v1beta1.Query/StreamEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &queryStreamEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Query_StreamEventsClient interface {
	Recv() (*EventsResponse, error)
	grpc.ClientStream
}

type queryStreamEventsClient struct {
	grpc.ClientStream
}

func (x *queryStreamEventsClient) Recv() (*EventsResponse, error) {
	m := new(EventsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	GetEvents(context.Context, *EventsRequest) (*EventsResponse, error)
	StreamEvents(*EventsRequest, Query_StreamEventsServer) error
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) GetEvents(ctx context.Context, req *EventsRequest) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}
func (*UnimplementedQueryServer) StreamEvents(req *EventsRequest, srv Query_StreamEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamEvents not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_GetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flux.stream.v1beta1.Query/GetEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetEvents(ctx, req.(*EventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_StreamEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServer).StreamEvents(m, &queryStreamEventsServer{stream})
}

type Query_StreamEventsServer interface {
	Send(*EventsResponse) error
	grpc.ServerStream
}

type queryStreamEventsServer struct {
	grpc.ServerStream
}

func (x *queryStreamEventsServer) Send(m *EventsResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "flux.stream.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEvents",
			Handler:    _Query_GetEvents_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEvents",
			Handler:       _Query_StreamEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "flux/stream/v1beta1/query.proto",
}

func (m *EventsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TmQueries) > 0 {
		for iNdEx := len(m.TmQueries) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.TmQueries[iNdEx])
			copy(dAtA[i:], m.TmQueries[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.TmQueries[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Modules) > 0 {
		for iNdEx := len(m.Modules) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Modules[iNdEx])
			copy(dAtA[i:], m.Modules[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.Modules[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Height != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TmData) > 0 {
		for iNdEx := len(m.TmData) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.TmData[iNdEx])
			copy(dAtA[i:], m.TmData[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.TmData[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.TmQueries) > 0 {
		for iNdEx := len(m.TmQueries) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.TmQueries[iNdEx])
			copy(dAtA[i:], m.TmQueries[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.TmQueries[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Events) > 0 {
		for iNdEx := len(m.Events) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Events[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Modules) > 0 {
		for iNdEx := len(m.Modules) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Modules[iNdEx])
			copy(dAtA[i:], m.Modules[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.Modules[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Time != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Time))
		i--
		dAtA[i] = 0x10
	}
	if m.Height != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Events) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Events) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Events) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EventOps) > 0 {
		for iNdEx := len(m.EventOps) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.EventOps[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *EventOp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventOp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventOp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Op != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Op))
		i--
		dAtA[i] = 0x8
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
func (m *EventsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovQuery(uint64(m.Height))
	}
	if len(m.Modules) > 0 {
		for _, s := range m.Modules {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if len(m.TmQueries) > 0 {
		for _, s := range m.TmQueries {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *EventsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovQuery(uint64(m.Height))
	}
	if m.Time != 0 {
		n += 1 + sovQuery(uint64(m.Time))
	}
	if len(m.Modules) > 0 {
		for _, s := range m.Modules {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if len(m.TmQueries) > 0 {
		for _, s := range m.TmQueries {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if len(m.TmData) > 0 {
		for _, s := range m.TmData {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *Events) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.EventOps) > 0 {
		for _, e := range m.EventOps {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *EventOp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Op != 0 {
		n += 1 + sovQuery(uint64(m.Op))
	}
	if m.Data != nil {
		l = m.Data.Size()
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
func (m *EventsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Modules", wireType)
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
			m.Modules = append(m.Modules, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TmQueries", wireType)
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
			m.TmQueries = append(m.TmQueries, string(dAtA[iNdEx:postIndex]))
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
func (m *EventsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Modules", wireType)
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
			m.Modules = append(m.Modules, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, &Events{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TmQueries", wireType)
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
			m.TmQueries = append(m.TmQueries, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TmData", wireType)
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
			m.TmData = append(m.TmData, string(dAtA[iNdEx:postIndex]))
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
func (m *Events) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Events: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Events: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventOps", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EventOps = append(m.EventOps, &EventOp{})
			if err := m.EventOps[len(m.EventOps)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *EventOp) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventOp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventOp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Op", wireType)
			}
			m.Op = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Op |= Op(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &types.Any{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
