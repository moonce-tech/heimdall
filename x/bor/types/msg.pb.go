// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: heimdall/bor/v1beta1/msg.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgProposeSpan struct {
	SpanId     uint64 `protobuf:"varint,1,opt,name=span_id,json=spanId,proto3" json:"span_id" yaml:"span_id"`
	Proposer   string `protobuf:"bytes,2,opt,name=proposer,proto3" json:"proposer" yaml:"proposer"`
	StartBlock uint64 `protobuf:"varint,3,opt,name=start_block,json=startBlock,proto3" json:"start_block" yaml:"start_block"`
	EndBlock   uint64 `protobuf:"varint,4,opt,name=end_block,json=endBlock,proto3" json:"end_block" yaml:"end_block"`
	ChainId    string `protobuf:"bytes,5,opt,name=chain_id,json=chainId,proto3" json:"chain_id" yaml:"chain_id"`
	Seed       string `protobuf:"bytes,6,opt,name=seed,proto3" json:"seed" yaml:"seed"`
}

func (m *MsgProposeSpan) Reset()         { *m = MsgProposeSpan{} }
func (m *MsgProposeSpan) String() string { return proto.CompactTextString(m) }
func (*MsgProposeSpan) ProtoMessage()    {}
func (*MsgProposeSpan) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d99edf48c57200b, []int{0}
}
func (m *MsgProposeSpan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgProposeSpan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgProposeSpan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgProposeSpan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgProposeSpan.Merge(m, src)
}
func (m *MsgProposeSpan) XXX_Size() int {
	return m.Size()
}
func (m *MsgProposeSpan) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgProposeSpan.DiscardUnknown(m)
}

var xxx_messageInfo_MsgProposeSpan proto.InternalMessageInfo

func (m *MsgProposeSpan) GetSpanId() uint64 {
	if m != nil {
		return m.SpanId
	}
	return 0
}

func (m *MsgProposeSpan) GetProposer() string {
	if m != nil {
		return m.Proposer
	}
	return ""
}

func (m *MsgProposeSpan) GetStartBlock() uint64 {
	if m != nil {
		return m.StartBlock
	}
	return 0
}

func (m *MsgProposeSpan) GetEndBlock() uint64 {
	if m != nil {
		return m.EndBlock
	}
	return 0
}

func (m *MsgProposeSpan) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *MsgProposeSpan) GetSeed() string {
	if m != nil {
		return m.Seed
	}
	return ""
}

// MsgProposeSpanResponse defines the Msg/MsgProposeSpan response type.
type MsgProposeSpanResponse struct {
}

func (m *MsgProposeSpanResponse) Reset()         { *m = MsgProposeSpanResponse{} }
func (m *MsgProposeSpanResponse) String() string { return proto.CompactTextString(m) }
func (*MsgProposeSpanResponse) ProtoMessage()    {}
func (*MsgProposeSpanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d99edf48c57200b, []int{1}
}
func (m *MsgProposeSpanResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgProposeSpanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgProposeSpanResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgProposeSpanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgProposeSpanResponse.Merge(m, src)
}
func (m *MsgProposeSpanResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgProposeSpanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgProposeSpanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgProposeSpanResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgProposeSpan)(nil), "heimdall.bor.v1beta1.MsgProposeSpan")
	proto.RegisterType((*MsgProposeSpanResponse)(nil), "heimdall.bor.v1beta1.MsgProposeSpanResponse")
}

func init() { proto.RegisterFile("heimdall/bor/v1beta1/msg.proto", fileDescriptor_7d99edf48c57200b) }

var fileDescriptor_7d99edf48c57200b = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x73, 0x69, 0x48, 0xd3, 0xab, 0x54, 0xd0, 0xa9, 0x80, 0x15, 0x81, 0xaf, 0x9c, 0x40,
	0x54, 0xd0, 0xda, 0x2a, 0x48, 0x0c, 0x41, 0x62, 0xc8, 0x50, 0xa9, 0x43, 0xa5, 0xca, 0x65, 0x62,
	0xa9, 0xce, 0xb9, 0x93, 0x63, 0xd5, 0xbe, 0xb3, 0x7c, 0x07, 0xb4, 0x2b, 0x9f, 0x00, 0x89, 0x8d,
	0x91, 0x4f, 0xc3, 0x58, 0x89, 0x85, 0xe9, 0x84, 0x12, 0x26, 0x8f, 0x1e, 0x98, 0x51, 0xce, 0x76,
	0xdc, 0x4a, 0x19, 0xba, 0xbd, 0xf7, 0x7b, 0xff, 0xf7, 0x7f, 0x7e, 0x7e, 0x07, 0xdd, 0x29, 0x8f,
	0x53, 0x46, 0x93, 0xc4, 0x0f, 0x65, 0xee, 0x7f, 0x3a, 0x08, 0xb9, 0xa6, 0x07, 0x7e, 0xaa, 0x22,
	0x2f, 0xcb, 0xa5, 0x96, 0x68, 0xbb, 0xa9, 0x7b, 0xa1, 0xcc, 0xbd, 0xba, 0x3e, 0xdc, 0x8e, 0x64,
	0x24, 0xad, 0xc0, 0x5f, 0x44, 0x95, 0x76, 0xf8, 0x28, 0x92, 0x32, 0x4a, 0xb8, 0x4f, 0xb3, 0xd8,
	0xa7, 0x42, 0x48, 0x4d, 0x75, 0x2c, 0x85, 0xaa, 0xaa, 0xe4, 0x5f, 0x17, 0x6e, 0x1d, 0xab, 0xe8,
	0x24, 0x97, 0x99, 0x54, 0xfc, 0x34, 0xa3, 0x02, 0xbd, 0x81, 0xeb, 0x2a, 0xa3, 0xe2, 0x2c, 0x66,
	0x0e, 0xd8, 0x01, 0xbb, 0xbd, 0xf1, 0xe3, 0xc2, 0xe0, 0x06, 0x95, 0x06, 0x6f, 0x5d, 0xd2, 0x34,
	0x19, 0x91, 0x1a, 0x90, 0xa0, 0xbf, 0x88, 0x8e, 0x18, 0x7a, 0x0b, 0x07, 0x59, 0x65, 0x93, 0x3b,
	0xdd, 0x1d, 0xb0, 0xbb, 0x31, 0xc6, 0x85, 0xc1, 0x4b, 0x56, 0x1a, 0x7c, 0xb7, 0xea, 0x6c, 0x08,
	0x09, 0x96, 0x45, 0x74, 0x08, 0x37, 0x95, 0xa6, 0xb9, 0x3e, 0x0b, 0x13, 0x39, 0x39, 0x77, 0xd6,
	0xec, 0xe0, 0x67, 0x85, 0xc1, 0xd7, 0x71, 0x69, 0x30, 0xaa, 0x87, 0xb7, 0x90, 0x04, 0xd0, 0x66,
	0xe3, 0x45, 0x82, 0xde, 0xc1, 0x0d, 0x2e, 0x58, 0xed, 0xd2, 0xb3, 0x2e, 0x4f, 0x0a, 0x83, 0x5b,
	0x58, 0x1a, 0x7c, 0xaf, 0xf2, 0x58, 0x22, 0x12, 0x0c, 0xb8, 0x60, 0x55, 0xff, 0x08, 0x0e, 0x26,
	0x53, 0x1a, 0xdb, 0xed, 0xef, 0xb4, 0x4b, 0x34, 0xac, 0x5d, 0xa2, 0x21, 0x24, 0x58, 0xb7, 0xe1,
	0x11, 0x43, 0x2f, 0x61, 0x4f, 0x71, 0xce, 0x9c, 0xbe, 0xed, 0x7b, 0x58, 0x18, 0x6c, 0xf3, 0xd2,
	0xe0, 0xcd, 0xfa, 0xab, 0x39, 0x67, 0x24, 0xb0, 0x90, 0x38, 0xf0, 0xc1, 0xcd, 0xff, 0x1e, 0x70,
	0x95, 0x49, 0xa1, 0xf8, 0xab, 0x1f, 0x00, 0xae, 0x1d, 0xab, 0x08, 0x7d, 0x07, 0xf0, 0xfe, 0x89,
	0x54, 0xfa, 0x94, 0x0b, 0x76, 0x4d, 0xf7, 0xfe, 0x02, 0x3d, 0xf5, 0x56, 0xdd, 0xdf, 0xbb, 0xe9,
	0x37, 0xdc, 0xbb, 0x8d, 0xaa, 0x99, 0x4a, 0xf6, 0xbf, 0xfc, 0xfa, 0xfb, 0xad, 0xfb, 0x9c, 0x10,
	0x7f, 0xe5, 0xdb, 0xab, 0x0f, 0xb5, 0xbf, 0xb8, 0xf5, 0x08, 0xbc, 0x18, 0x1f, 0xfe, 0x9c, 0xb9,
	0xe0, 0x6a, 0xe6, 0x82, 0x3f, 0x33, 0x17, 0x7c, 0x9d, 0xbb, 0x9d, 0xab, 0xb9, 0xdb, 0xf9, 0x3d,
	0x77, 0x3b, 0x1f, 0xf6, 0xa2, 0x58, 0x4f, 0x3f, 0x86, 0xde, 0x44, 0xa6, 0x7e, 0x4a, 0x75, 0x3c,
	0x11, 0x5c, 0x7f, 0x96, 0xf9, 0x79, 0xeb, 0x7b, 0x61, 0x9d, 0xf5, 0x65, 0xc6, 0x55, 0xd8, 0xb7,
	0xcf, 0xf0, 0xf5, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x41, 0x4c, 0xea, 0x48, 0xf2, 0x02, 0x00,
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
	PostSendProposeSpanTx(ctx context.Context, in *MsgProposeSpan, opts ...grpc.CallOption) (*MsgProposeSpanResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) PostSendProposeSpanTx(ctx context.Context, in *MsgProposeSpan, opts ...grpc.CallOption) (*MsgProposeSpanResponse, error) {
	out := new(MsgProposeSpanResponse)
	err := c.cc.Invoke(ctx, "/heimdall.bor.v1beta1.Msg/PostSendProposeSpanTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	PostSendProposeSpanTx(context.Context, *MsgProposeSpan) (*MsgProposeSpanResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) PostSendProposeSpanTx(ctx context.Context, req *MsgProposeSpan) (*MsgProposeSpanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostSendProposeSpanTx not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_PostSendProposeSpanTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgProposeSpan)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PostSendProposeSpanTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heimdall.bor.v1beta1.Msg/PostSendProposeSpanTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PostSendProposeSpanTx(ctx, req.(*MsgProposeSpan))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "heimdall.bor.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostSendProposeSpanTx",
			Handler:    _Msg_PostSendProposeSpanTx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "heimdall/bor/v1beta1/msg.proto",
}

func (m *MsgProposeSpan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgProposeSpan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgProposeSpan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Seed) > 0 {
		i -= len(m.Seed)
		copy(dAtA[i:], m.Seed)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Seed)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x2a
	}
	if m.EndBlock != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.EndBlock))
		i--
		dAtA[i] = 0x20
	}
	if m.StartBlock != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.StartBlock))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Proposer) > 0 {
		i -= len(m.Proposer)
		copy(dAtA[i:], m.Proposer)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Proposer)))
		i--
		dAtA[i] = 0x12
	}
	if m.SpanId != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.SpanId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgProposeSpanResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgProposeSpanResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgProposeSpanResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgProposeSpan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SpanId != 0 {
		n += 1 + sovMsg(uint64(m.SpanId))
	}
	l = len(m.Proposer)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.StartBlock != 0 {
		n += 1 + sovMsg(uint64(m.StartBlock))
	}
	if m.EndBlock != 0 {
		n += 1 + sovMsg(uint64(m.EndBlock))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = len(m.Seed)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	return n
}

func (m *MsgProposeSpanResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsg(x uint64) (n int) {
	return sovMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgProposeSpan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgProposeSpan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgProposeSpan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpanId", wireType)
			}
			m.SpanId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SpanId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartBlock", wireType)
			}
			m.StartBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndBlock", wireType)
			}
			m.EndBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Seed = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *MsgProposeSpanResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgProposeSpanResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgProposeSpanResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsg
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
func skipMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
				return 0, ErrInvalidLengthMsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsg = fmt.Errorf("proto: unexpected end of group")
)