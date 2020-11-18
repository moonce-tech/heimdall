// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: auth/v1beta/account.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_maticnetwork_heimdall_types_common "github.com/maticnetwork/heimdall/types/common"
	github_com_tendermint_tendermint_crypto "github.com/tendermint/tendermint/crypto"
	math "math"
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

// GenesisState defines the checkpoint module's genesis state.
type BaseAccount struct {
	Address       github_com_maticnetwork_heimdall_types_common.HeimdallAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/maticnetwork/heimdall/types/common.HeimdallAddress" json:"address,omitempty"`
	Coins         github_com_cosmos_cosmos_sdk_types.Coins                      `protobuf:"bytes,2,opt,name=coins,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins,omitempty"`
	PubKey        github_com_tendermint_tendermint_crypto.PubKey                `protobuf:"bytes,3,opt,name=pub_key,json=pubKey,proto3,casttype=github.com/tendermint/tendermint/crypto.PubKey" json:"pub_key,omitempty"`
	AccountNumber uint64                                                        `protobuf:"varint,4,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	Sequence      uint64                                                        `protobuf:"varint,5,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (m *BaseAccount) Reset()      { *m = BaseAccount{} }
func (*BaseAccount) ProtoMessage() {}
func (*BaseAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_3acecb14b5417769, []int{0}
}
func (m *BaseAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseAccount.Unmarshal(m, b)
}
func (m *BaseAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseAccount.Marshal(b, m, deterministic)
}
func (m *BaseAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseAccount.Merge(m, src)
}
func (m *BaseAccount) XXX_Size() int {
	return xxx_messageInfo_BaseAccount.Size(m)
}
func (m *BaseAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseAccount.DiscardUnknown(m)
}

var xxx_messageInfo_BaseAccount proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BaseAccount)(nil), "heimdall.auth.v1beta1.BaseAccount")
}

func init() { proto.RegisterFile("auth/v1beta/account.proto", fileDescriptor_3acecb14b5417769) }

var fileDescriptor_3acecb14b5417769 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xbd, 0x6e, 0xea, 0x30,
	0x14, 0xc7, 0x1d, 0x2e, 0x1f, 0xf7, 0xe6, 0x7e, 0x0c, 0x11, 0x57, 0xca, 0x65, 0x88, 0xd1, 0x95,
	0x2a, 0x65, 0x28, 0xb6, 0x68, 0xb7, 0x4a, 0x1d, 0x48, 0x87, 0x56, 0x42, 0xaa, 0xaa, 0x8c, 0xed,
	0x80, 0x12, 0xc7, 0x82, 0x08, 0x6c, 0xa7, 0xb1, 0xd3, 0x36, 0x6f, 0xc0, 0xc8, 0xd8, 0x11, 0xde,
	0xa6, 0x23, 0x43, 0x87, 0x4e, 0xa8, 0x82, 0xb7, 0x60, 0xaa, 0x70, 0x42, 0xc5, 0xd6, 0xe9, 0x9c,
	0xfc, 0xf5, 0x3b, 0xbf, 0xc8, 0xe7, 0x98, 0xff, 0x82, 0x4c, 0x8d, 0xf0, 0x43, 0x37, 0xa4, 0x2a,
	0xc0, 0x01, 0x21, 0x22, 0xe3, 0x0a, 0x25, 0xa9, 0x50, 0xc2, 0xfa, 0x3b, 0xa2, 0x31, 0x8b, 0x82,
	0xc9, 0x04, 0xed, 0x18, 0x54, 0x30, 0xdd, 0x56, 0x73, 0x28, 0x86, 0x42, 0x13, 0x78, 0xd7, 0x15,
	0xf0, 0xff, 0xd7, 0x8a, 0xf9, 0xd3, 0x0b, 0x24, 0xed, 0x15, 0x0a, 0xeb, 0xce, 0x6c, 0x04, 0x51,
	0x94, 0x52, 0x29, 0x6d, 0xa3, 0x6d, 0xb8, 0x3f, 0xbc, 0xde, 0x76, 0x05, 0xcf, 0x87, 0xb1, 0x1a,
	0x65, 0x21, 0x22, 0x82, 0x61, 0x16, 0xa8, 0x98, 0x70, 0xaa, 0x1e, 0x45, 0x3a, 0xc6, 0xfb, 0x3f,
	0x61, 0x95, 0x27, 0x54, 0x62, 0x22, 0x18, 0x13, 0x1c, 0x5d, 0x95, 0x69, 0xaf, 0x10, 0xf9, 0x7b,
	0xa3, 0xe5, 0x99, 0x35, 0x22, 0x62, 0x2e, 0xed, 0x8a, 0x56, 0x1f, 0x6f, 0x57, 0xd0, 0x3d, 0x50,
	0x13, 0x21, 0x99, 0x90, 0x65, 0xe9, 0xc8, 0x68, 0x5c, 0x68, 0xd1, 0xc5, 0x6e, 0xc6, 0x2f, 0x46,
	0xad, 0xbe, 0xd9, 0x48, 0xb2, 0x70, 0x30, 0xa6, 0xb9, 0xfd, 0xad, 0x6d, 0xb8, 0xbf, 0xbc, 0x93,
	0xed, 0x0a, 0xa2, 0x03, 0x8b, 0xa2, 0x3c, 0xa2, 0x29, 0x8b, 0xb9, 0x3a, 0x6c, 0x49, 0x9a, 0x27,
	0x4a, 0xa0, 0x9b, 0x2c, 0xec, 0xd3, 0xdc, 0xaf, 0x27, 0xba, 0x5a, 0x47, 0xe6, 0x9f, 0x72, 0x77,
	0x03, 0x9e, 0xb1, 0x90, 0xa6, 0x76, 0xb5, 0x6d, 0xb8, 0x55, 0xff, 0x77, 0x99, 0x5e, 0xeb, 0xd0,
	0x6a, 0x99, 0xdf, 0x25, 0xbd, 0xcf, 0x28, 0x27, 0xd4, 0xae, 0x69, 0xe0, 0xf3, 0xfb, 0xac, 0x39,
	0x9d, 0x43, 0xf0, 0x3c, 0x87, 0x60, 0xba, 0x80, 0x60, 0xb6, 0x80, 0x60, 0xbe, 0x80, 0xc0, 0xbb,
	0x7c, 0x59, 0x3b, 0xc6, 0x72, 0xed, 0x18, 0xef, 0x6b, 0xc7, 0x98, 0x6d, 0x1c, 0xb0, 0xdc, 0x38,
	0xe0, 0x6d, 0xe3, 0x80, 0xdb, 0xce, 0x97, 0xbb, 0x7c, 0xc2, 0xfa, 0xb6, 0xfa, 0xed, 0x61, 0x5d,
	0x9f, 0xe9, 0xf4, 0x23, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x0f, 0xd1, 0x50, 0xf0, 0x01, 0x00, 0x00,
}
