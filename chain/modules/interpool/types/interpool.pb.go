// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: flux/interpool/v1beta1/interpool.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	github_com_cometbft_cometbft_libs_bytes "github.com/cometbft/cometbft/libs/bytes"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type InterPool struct {
	PoolId       github_com_cometbft_cometbft_libs_bytes.HexBytes `protobuf:"bytes,1,opt,name=pool_id,json=poolId,proto3,casttype=github.com/cometbft/cometbft/libs/bytes.HexBytes" json:"pool_id,omitempty" bson:"pool_id"`
	OperatorAddr string                                           `protobuf:"bytes,2,opt,name=operator_addr,json=operatorAddr,proto3" json:"operator_addr,omitempty" bson:"operator_addr"`
	// on-going assets
	InventorySnapshot []github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,3,rep,name=inventory_snapshot,json=inventorySnapshot,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"inventory_snapshot" bson:"inventory_snapshot"`
	// initial assets before any trades from the pool
	BaseCapital        []github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,4,rep,name=base_capital,json=baseCapital,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"base_capital" bson:"base_capital"`
	OperatorCommission uint64                                    `protobuf:"varint,5,opt,name=operator_commission,json=operatorCommission,proto3" json:"operator_commission,omitempty" bson:"operator_commission"`
	// flow control data for cron service
	InputBlob []byte `protobuf:"bytes,6,opt,name=input_blob,json=inputBlob,proto3" json:"input_blob,omitempty" bson:"input_blob"`
	// pool extra state for lp, reward tokens created by cron service
	OutputBlob []byte `protobuf:"bytes,7,opt,name=output_blob,json=outputBlob,proto3" json:"output_blob,omitempty" bson:"output_blob"`
	// cron job controlling the pool
	CronJobId []byte `protobuf:"bytes,8,opt,name=cron_job_id,json=cronJobId,proto3" json:"cron_job_id,omitempty" bson:"cron_job_id"`
}

func (m *InterPool) Reset()         { *m = InterPool{} }
func (m *InterPool) String() string { return proto.CompactTextString(m) }
func (*InterPool) ProtoMessage()    {}
func (*InterPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b6c51933d1d5090, []int{0}
}
func (m *InterPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InterPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InterPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InterPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterPool.Merge(m, src)
}
func (m *InterPool) XXX_Size() int {
	return m.Size()
}
func (m *InterPool) XXX_DiscardUnknown() {
	xxx_messageInfo_InterPool.DiscardUnknown(m)
}

var xxx_messageInfo_InterPool proto.InternalMessageInfo

func (m *InterPool) GetPoolId() github_com_cometbft_cometbft_libs_bytes.HexBytes {
	if m != nil {
		return m.PoolId
	}
	return nil
}

func (m *InterPool) GetOperatorAddr() string {
	if m != nil {
		return m.OperatorAddr
	}
	return ""
}

func (m *InterPool) GetOperatorCommission() uint64 {
	if m != nil {
		return m.OperatorCommission
	}
	return 0
}

func (m *InterPool) GetInputBlob() []byte {
	if m != nil {
		return m.InputBlob
	}
	return nil
}

func (m *InterPool) GetOutputBlob() []byte {
	if m != nil {
		return m.OutputBlob
	}
	return nil
}

func (m *InterPool) GetCronJobId() []byte {
	if m != nil {
		return m.CronJobId
	}
	return nil
}

type PoolShare struct {
	PoolId                github_com_cometbft_cometbft_libs_bytes.HexBytes `protobuf:"bytes,1,opt,name=pool_id,json=poolId,proto3,casttype=github.com/cometbft/cometbft/libs/bytes.HexBytes" json:"pool_id,omitempty" bson:"pool_id"`
	LiquidityProviderAddr string                                           `protobuf:"bytes,2,opt,name=liquidity_provider_addr,json=liquidityProviderAddr,proto3" json:"liquidity_provider_addr,omitempty" bson:"liquidity_provider_addr"`
	// on-going assets of users
	AssetSnapshot []github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,3,rep,name=asset_snapshot,json=assetSnapshot,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"asset_snapshot" bson:"asset_snapshot"`
	// initial assets of users
	BaseCapital []github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,4,rep,name=base_capital,json=baseCapital,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"base_capital" bson:"base_capital"`
	// ownership percentage
	Shares cosmossdk_io_math.Int `protobuf:"bytes,5,opt,name=shares,proto3,customtype=cosmossdk.io/math.Int" json:"shares" bson:"shares"`
}

func (m *PoolShare) Reset()         { *m = PoolShare{} }
func (m *PoolShare) String() string { return proto.CompactTextString(m) }
func (*PoolShare) ProtoMessage()    {}
func (*PoolShare) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b6c51933d1d5090, []int{1}
}
func (m *PoolShare) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolShare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolShare.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolShare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolShare.Merge(m, src)
}
func (m *PoolShare) XXX_Size() int {
	return m.Size()
}
func (m *PoolShare) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolShare.DiscardUnknown(m)
}

var xxx_messageInfo_PoolShare proto.InternalMessageInfo

func (m *PoolShare) GetPoolId() github_com_cometbft_cometbft_libs_bytes.HexBytes {
	if m != nil {
		return m.PoolId
	}
	return nil
}

func (m *PoolShare) GetLiquidityProviderAddr() string {
	if m != nil {
		return m.LiquidityProviderAddr
	}
	return ""
}

func init() {
	proto.RegisterType((*InterPool)(nil), "flux.interpool.v1beta1.InterPool")
	proto.RegisterType((*PoolShare)(nil), "flux.interpool.v1beta1.PoolShare")
}

func init() {
	proto.RegisterFile("flux/interpool/v1beta1/interpool.proto", fileDescriptor_2b6c51933d1d5090)
}

var fileDescriptor_2b6c51933d1d5090 = []byte{
	// 673 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0xcd, 0x4e, 0xdb, 0x4c,
	0x14, 0x4d, 0x3e, 0x20, 0x7c, 0x19, 0x7e, 0x24, 0x0c, 0xa1, 0x81, 0x85, 0x1d, 0x79, 0xd1, 0x66,
	0x83, 0x0d, 0x6a, 0xd5, 0x4a, 0xec, 0x6a, 0xa4, 0x88, 0x54, 0xa8, 0x45, 0x86, 0x6e, 0xba, 0xb1,
	0x3c, 0xf6, 0x90, 0x4c, 0x71, 0xe6, 0xba, 0x9e, 0x31, 0x22, 0xea, 0xb6, 0x0f, 0xd0, 0x77, 0xe8,
	0x2b, 0xf4, 0x21, 0x58, 0xa2, 0xae, 0xaa, 0x2e, 0xac, 0x0a, 0xde, 0xc0, 0xcb, 0xae, 0xaa, 0xf1,
	0x38, 0x7f, 0x54, 0x2c, 0x58, 0x54, 0x5d, 0x65, 0xee, 0x39, 0xf7, 0xdc, 0x7b, 0x7d, 0xe7, 0x64,
	0xd0, 0xe3, 0xb3, 0x28, 0xbd, 0xb4, 0x29, 0x13, 0x24, 0x89, 0x01, 0x22, 0xfb, 0x62, 0x0f, 0x13,
	0xe1, 0xef, 0x4d, 0x10, 0x2b, 0x4e, 0x40, 0x80, 0xb6, 0x29, 0xf3, 0xac, 0x09, 0x5a, 0xe6, 0x6d,
	0x6f, 0xf5, 0x00, 0x7a, 0x11, 0xb1, 0x8b, 0x2c, 0x9c, 0x9e, 0xd9, 0x3e, 0x1b, 0x2a, 0xc9, 0xf6,
	0x46, 0x0f, 0x7a, 0x50, 0x1c, 0x6d, 0x79, 0x2a, 0xd1, 0xad, 0x00, 0xf8, 0x00, 0xb8, 0xa7, 0x08,
	0x15, 0x94, 0x94, 0xae, 0x22, 0x1b, 0xfb, 0x9c, 0x8c, 0x07, 0x09, 0x80, 0x32, 0xc5, 0x9b, 0x5f,
	0x16, 0x50, 0xbd, 0x2b, 0x27, 0x38, 0x06, 0x88, 0x34, 0x0f, 0x2d, 0xca, 0x49, 0x3c, 0x1a, 0x36,
	0xab, 0xad, 0x6a, 0x7b, 0xd9, 0xe9, 0xe4, 0x99, 0xb1, 0x8a, 0x39, 0xb0, 0x7d, 0xb3, 0x24, 0xcc,
	0x5f, 0x99, 0xb1, 0xdb, 0xa3, 0xa2, 0x9f, 0x62, 0x2b, 0x80, 0x81, 0x1d, 0xc0, 0x80, 0x08, 0x7c,
	0x26, 0x26, 0x87, 0x88, 0x62, 0x6e, 0xe3, 0xa1, 0x20, 0xdc, 0x3a, 0x24, 0x97, 0x8e, 0x3c, 0xb8,
	0x35, 0xa9, 0xee, 0x86, 0xda, 0x5b, 0xb4, 0x02, 0x31, 0x49, 0x7c, 0x01, 0x89, 0xe7, 0x87, 0x61,
	0xd2, 0xfc, 0xaf, 0x55, 0x6d, 0xd7, 0x9d, 0xdd, 0x3c, 0x33, 0x36, 0x54, 0x9b, 0x19, 0xda, 0xfc,
	0xf6, 0x75, 0x67, 0xa3, 0xfc, 0x9e, 0x97, 0x61, 0x98, 0x10, 0xce, 0x4f, 0x44, 0x42, 0x59, 0xcf,
	0x5d, 0x1e, 0xe5, 0x49, 0x58, 0xfb, 0x88, 0x34, 0xca, 0x2e, 0x08, 0x13, 0x90, 0x0c, 0x3d, 0xce,
	0xfc, 0x98, 0xf7, 0x41, 0x34, 0xe7, 0x5a, 0x73, 0xed, 0xba, 0x73, 0x74, 0x95, 0x19, 0x95, 0x1f,
	0x99, 0xf1, 0x64, 0x66, 0xe8, 0x62, 0x29, 0xea, 0x67, 0x87, 0x87, 0xe7, 0xb6, 0x18, 0xc6, 0x84,
	0x5b, 0x07, 0x40, 0x59, 0x9e, 0x19, 0x5b, 0x6a, 0x94, 0x3f, 0x4b, 0x9a, 0xee, 0xda, 0x18, 0x3c,
	0x29, 0x31, 0x8d, 0xa2, 0x65, 0xb9, 0x5d, 0x2f, 0xf0, 0x63, 0x2a, 0xfc, 0xa8, 0x39, 0x5f, 0xb4,
	0xed, 0x3c, 0xbc, 0xed, 0xba, 0x6a, 0x3b, 0x5d, 0xcc, 0x74, 0x97, 0x64, 0x78, 0xa0, 0x22, 0xed,
	0x0d, 0x5a, 0x1f, 0xef, 0x27, 0x80, 0xc1, 0x80, 0x72, 0x4e, 0x81, 0x35, 0x17, 0x5a, 0xd5, 0xf6,
	0xbc, 0xa3, 0xe7, 0x99, 0xb1, 0x7d, 0x67, 0x89, 0x93, 0x24, 0xd3, 0xd5, 0x46, 0xe8, 0xc1, 0x18,
	0xd4, 0x9e, 0x21, 0x44, 0x59, 0x9c, 0x0a, 0x0f, 0x47, 0x80, 0x9b, 0xb5, 0xe2, 0xce, 0x1b, 0x79,
	0x66, 0xac, 0x8d, 0x36, 0x30, 0xe2, 0x4c, 0xb7, 0x5e, 0x04, 0x4e, 0x04, 0x58, 0x7b, 0x81, 0x96,
	0x20, 0x15, 0x63, 0xd9, 0x62, 0x21, 0xdb, 0xcc, 0x33, 0x43, 0x2b, 0xdb, 0x4f, 0x48, 0xd3, 0x45,
	0x2a, 0x2a, 0x84, 0xcf, 0xd1, 0x52, 0x90, 0x00, 0xf3, 0xde, 0x03, 0x96, 0x1e, 0xfb, 0xff, 0xae,
	0x70, 0x8a, 0x34, 0xdd, 0xba, 0x8c, 0x5e, 0x01, 0xee, 0x86, 0xe6, 0xa7, 0x79, 0x54, 0x97, 0x06,
	0x3d, 0xe9, 0xfb, 0x09, 0xf9, 0xfb, 0x2e, 0x4d, 0xd0, 0xa3, 0x88, 0x7e, 0x48, 0x69, 0x48, 0xc5,
	0x50, 0xfe, 0xa9, 0x2e, 0x68, 0x48, 0x66, 0xfc, 0xba, 0x9f, 0x67, 0x86, 0xae, 0x1a, 0xde, 0x93,
	0x78, 0xbf, 0x73, 0x1b, 0x63, 0xc5, 0x71, 0x29, 0x28, 0x2c, 0x0c, 0x68, 0xd5, 0xe7, 0x9c, 0x88,
	0xbb, 0xf6, 0x3d, 0x7c, 0xb8, 0x8f, 0x1a, 0x6a, 0xb2, 0xd9, 0x72, 0xa6, 0xbb, 0x52, 0x00, 0xff,
	0xc2, 0xb6, 0x1d, 0x54, 0xe3, 0xf2, 0xe6, 0x78, 0xe1, 0xd4, 0xba, 0x63, 0x95, 0x4d, 0x1a, 0xaa,
	0x24, 0x0f, 0xcf, 0x2d, 0x0a, 0xf6, 0xc0, 0x17, 0x7d, 0xab, 0xcb, 0x44, 0x9e, 0x19, 0x2b, 0xaa,
	0xa4, 0x12, 0x99, 0x6e, 0xa9, 0x76, 0x4e, 0xaf, 0x6e, 0xf4, 0xea, 0xf5, 0x8d, 0x5e, 0xfd, 0x79,
	0xa3, 0x57, 0x3f, 0xdf, 0xea, 0x95, 0xeb, 0x5b, 0xbd, 0xf2, 0xfd, 0x56, 0xaf, 0xbc, 0xdb, 0x9f,
	0x1a, 0xb7, 0x13, 0xa5, 0x97, 0xaf, 0x3b, 0xa7, 0x47, 0x3e, 0xe6, 0xb6, 0x7c, 0x61, 0x43, 0x3b,
	0xe8, 0xfb, 0x94, 0xd9, 0x03, 0x08, 0xd3, 0x88, 0xf0, 0xa9, 0x87, 0xb9, 0xf8, 0x0c, 0x5c, 0x2b,
	0x5e, 0xc2, 0xa7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x80, 0x55, 0xf6, 0xaf, 0xb7, 0x05, 0x00,
	0x00,
}

func (m *InterPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InterPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InterPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CronJobId) > 0 {
		i -= len(m.CronJobId)
		copy(dAtA[i:], m.CronJobId)
		i = encodeVarintInterpool(dAtA, i, uint64(len(m.CronJobId)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.OutputBlob) > 0 {
		i -= len(m.OutputBlob)
		copy(dAtA[i:], m.OutputBlob)
		i = encodeVarintInterpool(dAtA, i, uint64(len(m.OutputBlob)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.InputBlob) > 0 {
		i -= len(m.InputBlob)
		copy(dAtA[i:], m.InputBlob)
		i = encodeVarintInterpool(dAtA, i, uint64(len(m.InputBlob)))
		i--
		dAtA[i] = 0x32
	}
	if m.OperatorCommission != 0 {
		i = encodeVarintInterpool(dAtA, i, uint64(m.OperatorCommission))
		i--
		dAtA[i] = 0x28
	}
	if len(m.BaseCapital) > 0 {
		for iNdEx := len(m.BaseCapital) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.BaseCapital[iNdEx].Size()
				i -= size
				if _, err := m.BaseCapital[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintInterpool(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.InventorySnapshot) > 0 {
		for iNdEx := len(m.InventorySnapshot) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.InventorySnapshot[iNdEx].Size()
				i -= size
				if _, err := m.InventorySnapshot[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintInterpool(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.OperatorAddr) > 0 {
		i -= len(m.OperatorAddr)
		copy(dAtA[i:], m.OperatorAddr)
		i = encodeVarintInterpool(dAtA, i, uint64(len(m.OperatorAddr)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PoolId) > 0 {
		i -= len(m.PoolId)
		copy(dAtA[i:], m.PoolId)
		i = encodeVarintInterpool(dAtA, i, uint64(len(m.PoolId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PoolShare) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolShare) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolShare) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Shares.Size()
		i -= size
		if _, err := m.Shares.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInterpool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.BaseCapital) > 0 {
		for iNdEx := len(m.BaseCapital) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.BaseCapital[iNdEx].Size()
				i -= size
				if _, err := m.BaseCapital[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintInterpool(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.AssetSnapshot) > 0 {
		for iNdEx := len(m.AssetSnapshot) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.AssetSnapshot[iNdEx].Size()
				i -= size
				if _, err := m.AssetSnapshot[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintInterpool(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.LiquidityProviderAddr) > 0 {
		i -= len(m.LiquidityProviderAddr)
		copy(dAtA[i:], m.LiquidityProviderAddr)
		i = encodeVarintInterpool(dAtA, i, uint64(len(m.LiquidityProviderAddr)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PoolId) > 0 {
		i -= len(m.PoolId)
		copy(dAtA[i:], m.PoolId)
		i = encodeVarintInterpool(dAtA, i, uint64(len(m.PoolId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintInterpool(dAtA []byte, offset int, v uint64) int {
	offset -= sovInterpool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *InterPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PoolId)
	if l > 0 {
		n += 1 + l + sovInterpool(uint64(l))
	}
	l = len(m.OperatorAddr)
	if l > 0 {
		n += 1 + l + sovInterpool(uint64(l))
	}
	if len(m.InventorySnapshot) > 0 {
		for _, e := range m.InventorySnapshot {
			l = e.Size()
			n += 1 + l + sovInterpool(uint64(l))
		}
	}
	if len(m.BaseCapital) > 0 {
		for _, e := range m.BaseCapital {
			l = e.Size()
			n += 1 + l + sovInterpool(uint64(l))
		}
	}
	if m.OperatorCommission != 0 {
		n += 1 + sovInterpool(uint64(m.OperatorCommission))
	}
	l = len(m.InputBlob)
	if l > 0 {
		n += 1 + l + sovInterpool(uint64(l))
	}
	l = len(m.OutputBlob)
	if l > 0 {
		n += 1 + l + sovInterpool(uint64(l))
	}
	l = len(m.CronJobId)
	if l > 0 {
		n += 1 + l + sovInterpool(uint64(l))
	}
	return n
}

func (m *PoolShare) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PoolId)
	if l > 0 {
		n += 1 + l + sovInterpool(uint64(l))
	}
	l = len(m.LiquidityProviderAddr)
	if l > 0 {
		n += 1 + l + sovInterpool(uint64(l))
	}
	if len(m.AssetSnapshot) > 0 {
		for _, e := range m.AssetSnapshot {
			l = e.Size()
			n += 1 + l + sovInterpool(uint64(l))
		}
	}
	if len(m.BaseCapital) > 0 {
		for _, e := range m.BaseCapital {
			l = e.Size()
			n += 1 + l + sovInterpool(uint64(l))
		}
	}
	l = m.Shares.Size()
	n += 1 + l + sovInterpool(uint64(l))
	return n
}

func sovInterpool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInterpool(x uint64) (n int) {
	return sovInterpool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InterPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInterpool
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
			return fmt.Errorf("proto: InterPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InterPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterpool
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
				return ErrInvalidLengthInterpool
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthInterpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolId = append(m.PoolId[:0], dAtA[iNdEx:postIndex]...)
			if m.PoolId == nil {
				m.PoolId = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterpool
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
				return ErrInvalidLengthInterpool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OperatorAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InventorySnapshot", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterpool
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
				return ErrInvalidLengthInterpool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Coin
			m.InventorySnapshot = append(m.InventorySnapshot, v)
			if err := m.InventorySnapshot[len(m.InventorySnapshot)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseCapital", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterpool
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
				return ErrInvalidLengthInterpool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Coin
			m.BaseCapital = append(m.BaseCapital, v)
			if err := m.BaseCapital[len(m.BaseCapital)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorCommission", wireType)
			}
			m.OperatorCommission = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterpool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OperatorCommission |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InputBlob", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterpool
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
				return ErrInvalidLengthInterpool
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthInterpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InputBlob = append(m.InputBlob[:0], dAtA[iNdEx:postIndex]...)
			if m.InputBlob == nil {
				m.InputBlob = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutputBlob", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterpool
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
				return ErrInvalidLengthInterpool
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthInterpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OutputBlob = append(m.OutputBlob[:0], dAtA[iNdEx:postIndex]...)
			if m.OutputBlob == nil {
				m.OutputBlob = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CronJobId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterpool
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
				return ErrInvalidLengthInterpool
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthInterpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CronJobId = append(m.CronJobId[:0], dAtA[iNdEx:postIndex]...)
			if m.CronJobId == nil {
				m.CronJobId = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInterpool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInterpool
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
func (m *PoolShare) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInterpool
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
			return fmt.Errorf("proto: PoolShare: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolShare: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterpool
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
				return ErrInvalidLengthInterpool
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthInterpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolId = append(m.PoolId[:0], dAtA[iNdEx:postIndex]...)
			if m.PoolId == nil {
				m.PoolId = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidityProviderAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterpool
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
				return ErrInvalidLengthInterpool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LiquidityProviderAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetSnapshot", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterpool
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
				return ErrInvalidLengthInterpool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Coin
			m.AssetSnapshot = append(m.AssetSnapshot, v)
			if err := m.AssetSnapshot[len(m.AssetSnapshot)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseCapital", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterpool
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
				return ErrInvalidLengthInterpool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Coin
			m.BaseCapital = append(m.BaseCapital, v)
			if err := m.BaseCapital[len(m.BaseCapital)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shares", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterpool
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
				return ErrInvalidLengthInterpool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterpool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Shares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInterpool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInterpool
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
func skipInterpool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInterpool
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
					return 0, ErrIntOverflowInterpool
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
					return 0, ErrIntOverflowInterpool
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
				return 0, ErrInvalidLengthInterpool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInterpool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInterpool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInterpool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInterpool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInterpool = fmt.Errorf("proto: unexpected end of group")
)
