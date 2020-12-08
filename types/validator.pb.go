// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: heimdall/types/validator.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type ValidatorID int32

const (
	DEFAULT ValidatorID = 0
)

var ValidatorID_name = map[int32]string{
	0: "DEFAULT",
}

var ValidatorID_value = map[string]int32{
	"DEFAULT": 0,
}

func (ValidatorID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_30e71641a391e76e, []int{0}
}

type ValidatorSet struct {
	Validators       []*Validator `protobuf:"bytes,1,rep,name=validators,proto3" json:"validators,omitempty"`
	Proposer         *Validator   `protobuf:"bytes,2,opt,name=proposer,proto3" json:"proposer,omitempty"`
	TotalVotingPower int64        `protobuf:"varint,3,opt,name=total_voting_power,json=totalVotingPower,proto3" json:"total_voting_power,omitempty"`
}

func (m *ValidatorSet) Reset()      { *m = ValidatorSet{} }
func (*ValidatorSet) ProtoMessage() {}
func (*ValidatorSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_30e71641a391e76e, []int{0}
}
func (m *ValidatorSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSet.Merge(m, src)
}
func (m *ValidatorSet) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSet.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSet proto.InternalMessageInfo

type Validator struct {
	ID               ValidatorID `protobuf:"varint,1,opt,name=ID,proto3,enum=heimdall.types.ValidatorID" json:"ID,omitempty"`
	StartEpoch       uint64      `protobuf:"varint,2,opt,name=start_epoch,json=startEpoch,proto3" json:"start_epoch,omitempty"`
	EndEpoch         uint64      `protobuf:"varint,3,opt,name=end_epoch,json=endEpoch,proto3" json:"end_epoch,omitempty"`
	Nonce            uint64      `protobuf:"varint,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	VotingPower      int64       `protobuf:"varint,5,opt,name=voting_power,json=votingPower,proto3" json:"voting_power,omitempty"`
	PubKey           string      `protobuf:"bytes,6,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	Signer           string      `protobuf:"bytes,7,opt,name=signer,proto3" json:"signer,omitempty"`
	LastUpdated      string      `protobuf:"bytes,8,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	Jailed           bool        `protobuf:"varint,9,opt,name=jailed,proto3" json:"jailed,omitempty"`
	ProposerPriority int64       `protobuf:"varint,10,opt,name=proposer_priority,json=proposerPriority,proto3" json:"proposer_priority,omitempty"`
}

func (m *Validator) Reset()      { *m = Validator{} }
func (*Validator) ProtoMessage() {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_30e71641a391e76e, []int{1}
}
func (m *Validator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return m.Size()
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

type MinimalVal struct {
	ID          ValidatorID `protobuf:"varint,1,opt,name=ID,proto3,enum=heimdall.types.ValidatorID" json:"ID,omitempty"`
	VotingPower uint64      `protobuf:"varint,2,opt,name=voting_power,json=votingPower,proto3" json:"voting_power,omitempty"`
	Signer      string      `protobuf:"bytes,3,opt,name=signer,proto3" json:"signer,omitempty"`
}

func (m *MinimalVal) Reset()         { *m = MinimalVal{} }
func (m *MinimalVal) String() string { return proto.CompactTextString(m) }
func (*MinimalVal) ProtoMessage()    {}
func (*MinimalVal) Descriptor() ([]byte, []int) {
	return fileDescriptor_30e71641a391e76e, []int{2}
}
func (m *MinimalVal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MinimalVal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MinimalVal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MinimalVal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MinimalVal.Merge(m, src)
}
func (m *MinimalVal) XXX_Size() int {
	return m.Size()
}
func (m *MinimalVal) XXX_DiscardUnknown() {
	xxx_messageInfo_MinimalVal.DiscardUnknown(m)
}

var xxx_messageInfo_MinimalVal proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("heimdall.types.ValidatorID", ValidatorID_name, ValidatorID_value)
	proto.RegisterType((*ValidatorSet)(nil), "heimdall.types.ValidatorSet")
	proto.RegisterType((*Validator)(nil), "heimdall.types.Validator")
	proto.RegisterType((*MinimalVal)(nil), "heimdall.types.MinimalVal")
}

func init() { proto.RegisterFile("heimdall/types/validator.proto", fileDescriptor_30e71641a391e76e) }

var fileDescriptor_30e71641a391e76e = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x77, 0x92, 0x34, 0xd9, 0xbc, 0x2d, 0xa5, 0x0e, 0x45, 0xc7, 0x16, 0x36, 0x6b, 0x11,
	0x09, 0x56, 0xb2, 0x50, 0xf1, 0x60, 0x4f, 0x2a, 0xa9, 0x10, 0x54, 0x28, 0xab, 0xcd, 0xc1, 0xcb,
	0x32, 0xc9, 0x0e, 0x9b, 0xb1, 0x9b, 0x9d, 0x61, 0x76, 0x92, 0x12, 0x3f, 0x41, 0x8f, 0x1e, 0x3d,
	0x16, 0xf4, 0xe0, 0x47, 0xf0, 0x23, 0x78, 0x2c, 0x9e, 0x3c, 0x4a, 0xf2, 0x45, 0x64, 0x67, 0x93,
	0x34, 0x11, 0x14, 0x7a, 0xdb, 0xf7, 0xff, 0xbd, 0xf7, 0x9f, 0x79, 0x7f, 0x76, 0xc0, 0x1d, 0x30,
	0x3e, 0x8c, 0x68, 0x92, 0xf8, 0x7a, 0x22, 0x59, 0xe6, 0x8f, 0x69, 0xc2, 0x23, 0xaa, 0x85, 0x6a,
	0x49, 0x25, 0xb4, 0xc0, 0x5b, 0x0b, 0xde, 0x32, 0x7c, 0x77, 0x27, 0x16, 0xb1, 0x30, 0xc8, 0xcf,
	0xbf, 0x8a, 0xae, 0xfd, 0xef, 0x08, 0x36, 0xbb, 0x8b, 0xc9, 0xb7, 0x4c, 0xe3, 0xa7, 0x00, 0x4b,
	0xa7, 0x8c, 0x20, 0xaf, 0xdc, 0x74, 0x0e, 0xef, 0xb6, 0xd6, 0xbd, 0x5a, 0xcb, 0x89, 0x60, 0xa5,
	0x19, 0x3f, 0x01, 0x5b, 0x2a, 0x21, 0x45, 0xc6, 0x14, 0x29, 0x79, 0xe8, 0xff, 0x83, 0xcb, 0x56,
	0xfc, 0x08, 0xb0, 0x16, 0x9a, 0x26, 0xe1, 0x58, 0x68, 0x9e, 0xc6, 0xa1, 0x14, 0xe7, 0x4c, 0x91,
	0xb2, 0x87, 0x9a, 0xe5, 0x60, 0xdb, 0x90, 0xae, 0x01, 0x27, 0xb9, 0x7e, 0x64, 0x5f, 0x5c, 0x36,
	0xac, 0xcf, 0x97, 0x0d, 0x6b, 0xff, 0x67, 0x09, 0xea, 0x4b, 0x3f, 0x7c, 0x00, 0xa5, 0x4e, 0x9b,
	0x20, 0x0f, 0x35, 0xb7, 0x0e, 0xf7, 0xfe, 0x79, 0x6c, 0xa7, 0x1d, 0x94, 0x3a, 0x6d, 0xdc, 0x00,
	0x27, 0xd3, 0x54, 0xe9, 0x90, 0x49, 0xd1, 0x1f, 0x98, 0xcb, 0x56, 0x02, 0x30, 0xd2, 0x71, 0xae,
	0xe0, 0x3d, 0xa8, 0xb3, 0x34, 0x9a, 0xe3, 0xb2, 0xc1, 0x36, 0x4b, 0xa3, 0x02, 0xee, 0xc0, 0x46,
	0x2a, 0xd2, 0x3e, 0x23, 0x15, 0x03, 0x8a, 0x02, 0xdf, 0x83, 0xcd, 0xb5, 0x05, 0x36, 0xcc, 0x02,
	0xce, 0xf8, 0xfa, 0xee, 0xf8, 0x0e, 0xd4, 0xe4, 0xa8, 0x17, 0x9e, 0xb1, 0x09, 0xa9, 0x7a, 0xa8,
	0x59, 0x0f, 0xaa, 0x72, 0xd4, 0x7b, 0xc5, 0x26, 0xf8, 0x36, 0x54, 0x33, 0x1e, 0xa7, 0x4c, 0x91,
	0x5a, 0xa1, 0x17, 0x55, 0xee, 0x99, 0xd0, 0x4c, 0x87, 0x23, 0x19, 0x51, 0xcd, 0x22, 0x62, 0x1b,
	0xea, 0xe4, 0xda, 0x69, 0x21, 0xe5, 0xa3, 0x1f, 0x28, 0x4f, 0x58, 0x44, 0xea, 0x1e, 0x6a, 0xda,
	0xc1, 0xbc, 0xc2, 0x07, 0x70, 0x6b, 0x91, 0x70, 0x28, 0x15, 0x17, 0x8a, 0xeb, 0x09, 0x81, 0x22,
	0xd4, 0x05, 0x38, 0x99, 0xeb, 0x2b, 0xa1, 0x7e, 0x04, 0x78, 0xc3, 0x53, 0x3e, 0xa4, 0x49, 0x97,
	0x26, 0x37, 0x0b, 0xf5, 0xef, 0x00, 0x8a, 0x54, 0xd7, 0x02, 0xb8, 0xde, 0xb3, 0xbc, 0xba, 0xe7,
	0x51, 0x25, 0x3f, 0xff, 0xe1, 0x7d, 0x70, 0x56, 0x3c, 0xb1, 0x03, 0xb5, 0xf6, 0xf1, 0xcb, 0xe7,
	0xa7, 0xaf, 0xdf, 0x6d, 0x5b, 0xbb, 0xf6, 0xc5, 0x17, 0xd7, 0xfa, 0xf6, 0xd5, 0xb5, 0x5e, 0x3c,
	0xfb, 0x31, 0x75, 0xd1, 0xd5, 0xd4, 0x45, 0xbf, 0xa7, 0x2e, 0xfa, 0x34, 0x73, 0xad, 0xab, 0x99,
	0x6b, 0xfd, 0x9a, 0xb9, 0xd6, 0xfb, 0x07, 0x31, 0xd7, 0x83, 0x51, 0xaf, 0xd5, 0x17, 0x43, 0x7f,
	0x48, 0x35, 0xef, 0xa7, 0x4c, 0x9f, 0x0b, 0x75, 0xe6, 0xaf, 0xbf, 0x94, 0x5e, 0xd5, 0xfc, 0xfa,
	0x8f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x36, 0x4d, 0xcb, 0xe2, 0x42, 0x03, 0x00, 0x00,
}

func (m *ValidatorSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TotalVotingPower != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.TotalVotingPower))
		i--
		dAtA[i] = 0x18
	}
	if m.Proposer != nil {
		{
			size, err := m.Proposer.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintValidator(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Validators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintValidator(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Validator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Validator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Validator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProposerPriority != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.ProposerPriority))
		i--
		dAtA[i] = 0x50
	}
	if m.Jailed {
		i--
		if m.Jailed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if len(m.LastUpdated) > 0 {
		i -= len(m.LastUpdated)
		copy(dAtA[i:], m.LastUpdated)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.LastUpdated)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x32
	}
	if m.VotingPower != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.VotingPower))
		i--
		dAtA[i] = 0x28
	}
	if m.Nonce != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x20
	}
	if m.EndEpoch != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.EndEpoch))
		i--
		dAtA[i] = 0x18
	}
	if m.StartEpoch != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.StartEpoch))
		i--
		dAtA[i] = 0x10
	}
	if m.ID != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MinimalVal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MinimalVal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MinimalVal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x1a
	}
	if m.VotingPower != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.VotingPower))
		i--
		dAtA[i] = 0x10
	}
	if m.ID != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintValidator(dAtA []byte, offset int, v uint64) int {
	offset -= sovValidator(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ValidatorSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Validators) > 0 {
		for _, e := range m.Validators {
			l = e.Size()
			n += 1 + l + sovValidator(uint64(l))
		}
	}
	if m.Proposer != nil {
		l = m.Proposer.Size()
		n += 1 + l + sovValidator(uint64(l))
	}
	if m.TotalVotingPower != 0 {
		n += 1 + sovValidator(uint64(m.TotalVotingPower))
	}
	return n
}

func (m *Validator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovValidator(uint64(m.ID))
	}
	if m.StartEpoch != 0 {
		n += 1 + sovValidator(uint64(m.StartEpoch))
	}
	if m.EndEpoch != 0 {
		n += 1 + sovValidator(uint64(m.EndEpoch))
	}
	if m.Nonce != 0 {
		n += 1 + sovValidator(uint64(m.Nonce))
	}
	if m.VotingPower != 0 {
		n += 1 + sovValidator(uint64(m.VotingPower))
	}
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	l = len(m.LastUpdated)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	if m.Jailed {
		n += 2
	}
	if m.ProposerPriority != 0 {
		n += 1 + sovValidator(uint64(m.ProposerPriority))
	}
	return n
}

func (m *MinimalVal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovValidator(uint64(m.ID))
	}
	if m.VotingPower != 0 {
		n += 1 + sovValidator(uint64(m.VotingPower))
	}
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	return n
}

func sovValidator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozValidator(x uint64) (n int) {
	return sovValidator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValidatorSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
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
			return fmt.Errorf("proto: ValidatorSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validators = append(m.Validators, &Validator{})
			if err := m.Validators[len(m.Validators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Proposer == nil {
				m.Proposer = &Validator{}
			}
			if err := m.Proposer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalVotingPower", wireType)
			}
			m.TotalVotingPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalVotingPower |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthValidator
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthValidator
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
func (m *Validator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
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
			return fmt.Errorf("proto: Validator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= ValidatorID(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartEpoch", wireType)
			}
			m.StartEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartEpoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndEpoch", wireType)
			}
			m.EndEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndEpoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingPower", wireType)
			}
			m.VotingPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotingPower |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUpdated", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastUpdated = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Jailed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Jailed = bool(v != 0)
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposerPriority", wireType)
			}
			m.ProposerPriority = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposerPriority |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthValidator
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthValidator
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
func (m *MinimalVal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
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
			return fmt.Errorf("proto: MinimalVal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MinimalVal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= ValidatorID(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingPower", wireType)
			}
			m.VotingPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotingPower |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthValidator
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthValidator
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
func skipValidator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowValidator
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
					return 0, ErrIntOverflowValidator
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
					return 0, ErrIntOverflowValidator
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
				return 0, ErrInvalidLengthValidator
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupValidator
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthValidator
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthValidator        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowValidator          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupValidator = fmt.Errorf("proto: unexpected end of group")
)
