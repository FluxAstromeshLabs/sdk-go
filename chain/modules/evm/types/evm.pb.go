// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: flux/evm/v1beta1/evm.proto

package types

import (
	fmt "fmt"
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

type Account struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Balance []byte `protobuf:"bytes,2,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_b155af604ae5b1fb, []int{0}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Account.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return m.Size()
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Account) GetBalance() []byte {
	if m != nil {
		return m.Balance
	}
	return nil
}

type Code struct {
	Address  string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Bytecode []byte `protobuf:"bytes,2,opt,name=bytecode,proto3" json:"bytecode,omitempty"`
}

func (m *Code) Reset()         { *m = Code{} }
func (m *Code) String() string { return proto.CompactTextString(m) }
func (*Code) ProtoMessage()    {}
func (*Code) Descriptor() ([]byte, []int) {
	return fileDescriptor_b155af604ae5b1fb, []int{1}
}
func (m *Code) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Code) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Code.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Code) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Code.Merge(m, src)
}
func (m *Code) XXX_Size() int {
	return m.Size()
}
func (m *Code) XXX_DiscardUnknown() {
	xxx_messageInfo_Code.DiscardUnknown(m)
}

var xxx_messageInfo_Code proto.InternalMessageInfo

func (m *Code) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Code) GetBytecode() []byte {
	if m != nil {
		return m.Bytecode
	}
	return nil
}

type ContractStorage struct {
	Address  string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Storages []*Storage `protobuf:"bytes,2,rep,name=storages,proto3" json:"storages,omitempty"`
}

func (m *ContractStorage) Reset()         { *m = ContractStorage{} }
func (m *ContractStorage) String() string { return proto.CompactTextString(m) }
func (*ContractStorage) ProtoMessage()    {}
func (*ContractStorage) Descriptor() ([]byte, []int) {
	return fileDescriptor_b155af604ae5b1fb, []int{2}
}
func (m *ContractStorage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractStorage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractStorage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractStorage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractStorage.Merge(m, src)
}
func (m *ContractStorage) XXX_Size() int {
	return m.Size()
}
func (m *ContractStorage) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractStorage.DiscardUnknown(m)
}

var xxx_messageInfo_ContractStorage proto.InternalMessageInfo

func (m *ContractStorage) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ContractStorage) GetStorages() []*Storage {
	if m != nil {
		return m.Storages
	}
	return nil
}

type Storage struct {
	Key   []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Storage) Reset()         { *m = Storage{} }
func (m *Storage) String() string { return proto.CompactTextString(m) }
func (*Storage) ProtoMessage()    {}
func (*Storage) Descriptor() ([]byte, []int) {
	return fileDescriptor_b155af604ae5b1fb, []int{3}
}
func (m *Storage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Storage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Storage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Storage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Storage.Merge(m, src)
}
func (m *Storage) XXX_Size() int {
	return m.Size()
}
func (m *Storage) XXX_DiscardUnknown() {
	xxx_messageInfo_Storage.DiscardUnknown(m)
}

var xxx_messageInfo_Storage proto.InternalMessageInfo

func (m *Storage) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Storage) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type ContractInfo struct {
	Address  string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Bytecode []byte `protobuf:"bytes,2,opt,name=bytecode,proto3" json:"bytecode,omitempty"`
	Hash     []byte `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	Sender   string `protobuf:"bytes,4,opt,name=sender,proto3" json:"sender,omitempty"`
	Calldata []byte `protobuf:"bytes,5,opt,name=calldata,proto3" json:"calldata,omitempty"`
	Value    []byte `protobuf:"bytes,6,opt,name=value,proto3" json:"value,omitempty"`
	Number   int64  `protobuf:"varint,7,opt,name=number,proto3" json:"number,omitempty"`
}

func (m *ContractInfo) Reset()         { *m = ContractInfo{} }
func (m *ContractInfo) String() string { return proto.CompactTextString(m) }
func (*ContractInfo) ProtoMessage()    {}
func (*ContractInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b155af604ae5b1fb, []int{4}
}
func (m *ContractInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractInfo.Merge(m, src)
}
func (m *ContractInfo) XXX_Size() int {
	return m.Size()
}
func (m *ContractInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ContractInfo proto.InternalMessageInfo

func (m *ContractInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ContractInfo) GetBytecode() []byte {
	if m != nil {
		return m.Bytecode
	}
	return nil
}

func (m *ContractInfo) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *ContractInfo) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *ContractInfo) GetCalldata() []byte {
	if m != nil {
		return m.Calldata
	}
	return nil
}

func (m *ContractInfo) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ContractInfo) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func init() {
	proto.RegisterType((*Account)(nil), "flux.evm.v1beta1.Account")
	proto.RegisterType((*Code)(nil), "flux.evm.v1beta1.Code")
	proto.RegisterType((*ContractStorage)(nil), "flux.evm.v1beta1.ContractStorage")
	proto.RegisterType((*Storage)(nil), "flux.evm.v1beta1.Storage")
	proto.RegisterType((*ContractInfo)(nil), "flux.evm.v1beta1.ContractInfo")
}

func init() { proto.RegisterFile("flux/evm/v1beta1/evm.proto", fileDescriptor_b155af604ae5b1fb) }

var fileDescriptor_b155af604ae5b1fb = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xbd, 0x6e, 0xdb, 0x30,
	0x14, 0x85, 0xcd, 0xd8, 0xb1, 0x53, 0x36, 0x40, 0x03, 0x21, 0x28, 0x68, 0x0f, 0x82, 0xa1, 0xc9,
	0x93, 0x08, 0xf7, 0x67, 0x6b, 0x87, 0x36, 0x40, 0x80, 0x02, 0x45, 0x06, 0xb7, 0x53, 0x97, 0x82,
	0x7f, 0x96, 0x8d, 0x52, 0xbc, 0x81, 0x48, 0x1a, 0xd1, 0x5b, 0xf4, 0x79, 0xfa, 0x04, 0x1d, 0x33,
	0x76, 0x2c, 0xec, 0x17, 0x29, 0x28, 0x52, 0x6e, 0xd1, 0x21, 0x43, 0xb6, 0xf3, 0xf1, 0x5c, 0x1e,
	0x1d, 0x11, 0x17, 0xcf, 0xd6, 0xda, 0xdf, 0x51, 0xb5, 0xab, 0xe9, 0x6e, 0xc9, 0x95, 0x63, 0xcb,
	0xa0, 0xcb, 0xdb, 0x06, 0x1c, 0x64, 0x17, 0xc1, 0x2b, 0x03, 0x27, 0x6f, 0x36, 0xad, 0x00, 0x2a,
	0xad, 0x68, 0xe7, 0x73, 0xbf, 0xa6, 0xcc, 0xb4, 0x71, 0x78, 0x76, 0x59, 0x41, 0x05, 0x9d, 0xa4,
	0x41, 0xa5, 0xd3, 0xa9, 0x00, 0x5b, 0x83, 0xfd, 0x1a, 0x8d, 0x08, 0xc9, 0xca, 0x23, 0x51, 0xce,
	0xac, 0x3a, 0x7e, 0x5c, 0xc0, 0xd6, 0x44, 0xbf, 0x78, 0x8b, 0x27, 0xef, 0x84, 0x00, 0x6f, 0x5c,
	0x46, 0xf0, 0x84, 0x49, 0xd9, 0x28, 0x6b, 0x09, 0x9a, 0xa3, 0xc5, 0x93, 0x55, 0x8f, 0xc1, 0xe1,
	0x4c, 0x33, 0x23, 0x14, 0x39, 0x99, 0xa3, 0xc5, 0xf9, 0xaa, 0xc7, 0xe2, 0x0d, 0x1e, 0x5d, 0x81,
	0x54, 0x0f, 0xdc, 0x9d, 0xe1, 0x33, 0xde, 0x3a, 0x25, 0x40, 0xf6, 0x97, 0x8f, 0x5c, 0x70, 0xfc,
	0xec, 0x0a, 0x8c, 0x6b, 0x98, 0x70, 0x9f, 0x1c, 0x34, 0xac, 0x7a, 0x28, 0xe8, 0x35, 0x3e, 0xb3,
	0x71, 0xc8, 0x92, 0x93, 0xf9, 0x70, 0xf1, 0xf4, 0xc5, 0xb4, 0xfc, 0xff, 0xe9, 0xca, 0x14, 0xb3,
	0x3a, 0x8e, 0x16, 0x4b, 0x3c, 0xe9, 0xb3, 0x2f, 0xf0, 0xf0, 0x9b, 0x6a, 0x53, 0x8b, 0x20, 0xb3,
	0x4b, 0x7c, 0xba, 0x63, 0xda, 0x2b, 0x32, 0xec, 0xce, 0x22, 0x14, 0x3f, 0x10, 0x3e, 0xef, 0x7b,
	0x7d, 0x30, 0x6b, 0x78, 0xdc, 0xdf, 0x65, 0x19, 0x1e, 0x6d, 0x98, 0xdd, 0xa4, 0xec, 0x4e, 0x67,
	0xcf, 0xf1, 0xd8, 0x2a, 0x23, 0x55, 0x43, 0x46, 0x5d, 0x50, 0xa2, 0x90, 0x23, 0x98, 0xd6, 0x92,
	0x39, 0x46, 0x4e, 0x63, 0x4e, 0xcf, 0x7f, 0x4b, 0x8e, 0xff, 0x29, 0x19, 0x92, 0x8c, 0xaf, 0xb9,
	0x6a, 0xc8, 0x64, 0x8e, 0x16, 0xc3, 0x55, 0xa2, 0xf7, 0x37, 0x3f, 0xf7, 0x39, 0xba, 0xdf, 0xe7,
	0xe8, 0xf7, 0x3e, 0x47, 0xdf, 0x0f, 0xf9, 0xe0, 0xfe, 0x90, 0x0f, 0x7e, 0x1d, 0xf2, 0xc1, 0x97,
	0x57, 0xd5, 0xd6, 0x6d, 0x3c, 0x2f, 0x05, 0xd4, 0xf4, 0x5a, 0xfb, 0xbb, 0x9b, 0xeb, 0xcf, 0x1f,
	0x19, 0xb7, 0x34, 0x3c, 0xa2, 0xa4, 0x62, 0xc3, 0xb6, 0x86, 0xd6, 0x20, 0xbd, 0x56, 0xb6, 0x5b,
	0x55, 0xd7, 0xde, 0x2a, 0xcb, 0xc7, 0xdd, 0x9e, 0xbc, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0xa3,
	0x93, 0x21, 0x59, 0xc3, 0x02, 0x00, 0x00,
}

func (m *Account) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Account) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Balance) > 0 {
		i -= len(m.Balance)
		copy(dAtA[i:], m.Balance)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Balance)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Code) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Code) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Code) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Bytecode) > 0 {
		i -= len(m.Bytecode)
		copy(dAtA[i:], m.Bytecode)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Bytecode)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ContractStorage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractStorage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractStorage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Storages) > 0 {
		for iNdEx := len(m.Storages) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Storages[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvm(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Storage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Storage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Storage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *ContractInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Number != 0 {
		i = encodeVarintEvm(dAtA, i, uint64(m.Number))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Calldata) > 0 {
		i -= len(m.Calldata)
		copy(dAtA[i:], m.Calldata)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Calldata)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Bytecode) > 0 {
		i -= len(m.Bytecode)
		copy(dAtA[i:], m.Bytecode)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Bytecode)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvm(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvm(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Account) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	l = len(m.Balance)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	return n
}

func (m *Code) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	l = len(m.Bytecode)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	return n
}

func (m *ContractStorage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	if len(m.Storages) > 0 {
		for _, e := range m.Storages {
			l = e.Size()
			n += 1 + l + sovEvm(uint64(l))
		}
	}
	return n
}

func (m *Storage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	return n
}

func (m *ContractInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	l = len(m.Bytecode)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	l = len(m.Calldata)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.Number != 0 {
		n += 1 + sovEvm(uint64(m.Number))
	}
	return n
}

func sovEvm(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvm(x uint64) (n int) {
	return sovEvm(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Account) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvm
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
			return fmt.Errorf("proto: Account: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Account: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Balance = append(m.Balance[:0], dAtA[iNdEx:postIndex]...)
			if m.Balance == nil {
				m.Balance = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvm
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
func (m *Code) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvm
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
			return fmt.Errorf("proto: Code: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Code: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bytecode", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bytecode = append(m.Bytecode[:0], dAtA[iNdEx:postIndex]...)
			if m.Bytecode == nil {
				m.Bytecode = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvm
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
func (m *ContractStorage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvm
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
			return fmt.Errorf("proto: ContractStorage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractStorage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Storages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Storages = append(m.Storages, &Storage{})
			if err := m.Storages[len(m.Storages)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvm
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
func (m *Storage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvm
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
			return fmt.Errorf("proto: Storage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Storage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvm
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
func (m *ContractInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvm
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
			return fmt.Errorf("proto: ContractInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bytecode", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bytecode = append(m.Bytecode[:0], dAtA[iNdEx:postIndex]...)
			if m.Bytecode == nil {
				m.Bytecode = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Calldata", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Calldata = append(m.Calldata[:0], dAtA[iNdEx:postIndex]...)
			if m.Calldata == nil {
				m.Calldata = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			m.Number = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Number |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvm
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
func skipEvm(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvm
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
					return 0, ErrIntOverflowEvm
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
					return 0, ErrIntOverflowEvm
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
				return 0, ErrInvalidLengthEvm
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvm
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvm
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvm        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvm          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvm = fmt.Errorf("proto: unexpected end of group")
)