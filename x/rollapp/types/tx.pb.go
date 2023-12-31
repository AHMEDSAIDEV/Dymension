// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymension/rollapp/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// ===================== MsgCreateRollapp
type MsgCreateRollapp struct {
	// creator is the bech32-encoded address of the rollapp creator
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// rollappId is the unique identifier of the rollapp chain.
	// The rollappId follows the same standard as cosmos chain_id
	RollappId string `protobuf:"bytes,2,opt,name=rollappId,proto3" json:"rollappId,omitempty"`
	// codeStamp is the description of the genesis file location on the DA
	CodeStamp string `protobuf:"bytes,3,opt,name=codeStamp,proto3" json:"codeStamp,omitempty"` // Deprecated: Do not use.
	// genesisPath is the description of the genesis file location on the DA
	GenesisPath string `protobuf:"bytes,4,opt,name=genesisPath,proto3" json:"genesisPath,omitempty"` // Deprecated: Do not use.
	// maxWithholdingBlocks is the maximum number of blocks for
	// an active sequencer to send a state update (MsgUpdateState)
	MaxWithholdingBlocks uint64 `protobuf:"varint,5,opt,name=maxWithholdingBlocks,proto3" json:"maxWithholdingBlocks,omitempty"` // Deprecated: Do not use.
	// maxSequencers is the maximum number of sequencers
	MaxSequencers uint64 `protobuf:"varint,6,opt,name=maxSequencers,proto3" json:"maxSequencers,omitempty"`
	// permissionedAddresses is a bech32-encoded address list of the
	// sequencers that are allowed to serve this rollappId.
	// In the case of an empty list, the rollapp is considered permissionless
	PermissionedAddresses []string `protobuf:"bytes,7,rep,name=permissionedAddresses,proto3" json:"permissionedAddresses,omitempty"`
	// metadata provides the client information for all the registered tokens.
	Metadatas []TokenMetadata `protobuf:"bytes,8,rep,name=metadatas,proto3" json:"metadatas"`
}

func (m *MsgCreateRollapp) Reset()         { *m = MsgCreateRollapp{} }
func (m *MsgCreateRollapp) String() string { return proto.CompactTextString(m) }
func (*MsgCreateRollapp) ProtoMessage()    {}
func (*MsgCreateRollapp) Descriptor() ([]byte, []int) {
	return fileDescriptor_935cc363af28220c, []int{0}
}
func (m *MsgCreateRollapp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateRollapp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateRollapp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateRollapp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateRollapp.Merge(m, src)
}
func (m *MsgCreateRollapp) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateRollapp) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateRollapp.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateRollapp proto.InternalMessageInfo

func (m *MsgCreateRollapp) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateRollapp) GetRollappId() string {
	if m != nil {
		return m.RollappId
	}
	return ""
}

// Deprecated: Do not use.
func (m *MsgCreateRollapp) GetCodeStamp() string {
	if m != nil {
		return m.CodeStamp
	}
	return ""
}

// Deprecated: Do not use.
func (m *MsgCreateRollapp) GetGenesisPath() string {
	if m != nil {
		return m.GenesisPath
	}
	return ""
}

// Deprecated: Do not use.
func (m *MsgCreateRollapp) GetMaxWithholdingBlocks() uint64 {
	if m != nil {
		return m.MaxWithholdingBlocks
	}
	return 0
}

func (m *MsgCreateRollapp) GetMaxSequencers() uint64 {
	if m != nil {
		return m.MaxSequencers
	}
	return 0
}

func (m *MsgCreateRollapp) GetPermissionedAddresses() []string {
	if m != nil {
		return m.PermissionedAddresses
	}
	return nil
}

func (m *MsgCreateRollapp) GetMetadatas() []TokenMetadata {
	if m != nil {
		return m.Metadatas
	}
	return nil
}

type MsgCreateRollappResponse struct {
}

func (m *MsgCreateRollappResponse) Reset()         { *m = MsgCreateRollappResponse{} }
func (m *MsgCreateRollappResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateRollappResponse) ProtoMessage()    {}
func (*MsgCreateRollappResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_935cc363af28220c, []int{1}
}
func (m *MsgCreateRollappResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateRollappResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateRollappResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateRollappResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateRollappResponse.Merge(m, src)
}
func (m *MsgCreateRollappResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateRollappResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateRollappResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateRollappResponse proto.InternalMessageInfo

// ===================== MsgUpdateState
// Updating a rollapp state with a block batch
// a block batch is a list of ordered blocks (by height)
type MsgUpdateState struct {
	// creator is the bech32-encoded address of the sequencer sending the update
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// rollappId is the rollapp that the sequencer belongs to and asking to update
	// The rollappId follows the same standard as cosmos chain_id
	RollappId string `protobuf:"bytes,2,opt,name=rollappId,proto3" json:"rollappId,omitempty"`
	// startHeight is the block height of the first block in the batch
	StartHeight uint64 `protobuf:"varint,3,opt,name=startHeight,proto3" json:"startHeight,omitempty"`
	// numBlocks is the number of blocks included in this batch update
	NumBlocks uint64 `protobuf:"varint,4,opt,name=numBlocks,proto3" json:"numBlocks,omitempty"`
	// DAPath is the description of the location on the DA layer
	DAPath string `protobuf:"bytes,5,opt,name=DAPath,proto3" json:"DAPath,omitempty"`
	// version is the version of the rollapp
	Version uint64 `protobuf:"varint,6,opt,name=version,proto3" json:"version,omitempty"`
	// BDs is a list of block description objects (one per block)
	// the list must be ordered by height, starting from startHeight to startHeight+numBlocks-1
	BDs BlockDescriptors `protobuf:"bytes,7,opt,name=BDs,proto3" json:"BDs"`
}

func (m *MsgUpdateState) Reset()         { *m = MsgUpdateState{} }
func (m *MsgUpdateState) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateState) ProtoMessage()    {}
func (*MsgUpdateState) Descriptor() ([]byte, []int) {
	return fileDescriptor_935cc363af28220c, []int{2}
}
func (m *MsgUpdateState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateState.Merge(m, src)
}
func (m *MsgUpdateState) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateState) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateState.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateState proto.InternalMessageInfo

func (m *MsgUpdateState) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgUpdateState) GetRollappId() string {
	if m != nil {
		return m.RollappId
	}
	return ""
}

func (m *MsgUpdateState) GetStartHeight() uint64 {
	if m != nil {
		return m.StartHeight
	}
	return 0
}

func (m *MsgUpdateState) GetNumBlocks() uint64 {
	if m != nil {
		return m.NumBlocks
	}
	return 0
}

func (m *MsgUpdateState) GetDAPath() string {
	if m != nil {
		return m.DAPath
	}
	return ""
}

func (m *MsgUpdateState) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *MsgUpdateState) GetBDs() BlockDescriptors {
	if m != nil {
		return m.BDs
	}
	return BlockDescriptors{}
}

type MsgUpdateStateResponse struct {
}

func (m *MsgUpdateStateResponse) Reset()         { *m = MsgUpdateStateResponse{} }
func (m *MsgUpdateStateResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateStateResponse) ProtoMessage()    {}
func (*MsgUpdateStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_935cc363af28220c, []int{3}
}
func (m *MsgUpdateStateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateStateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateStateResponse.Merge(m, src)
}
func (m *MsgUpdateStateResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateStateResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateRollapp)(nil), "dymensionxyz.dymension.rollapp.MsgCreateRollapp")
	proto.RegisterType((*MsgCreateRollappResponse)(nil), "dymensionxyz.dymension.rollapp.MsgCreateRollappResponse")
	proto.RegisterType((*MsgUpdateState)(nil), "dymensionxyz.dymension.rollapp.MsgUpdateState")
	proto.RegisterType((*MsgUpdateStateResponse)(nil), "dymensionxyz.dymension.rollapp.MsgUpdateStateResponse")
}

func init() { proto.RegisterFile("dymension/rollapp/tx.proto", fileDescriptor_935cc363af28220c) }

var fileDescriptor_935cc363af28220c = []byte{
	// 532 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xc1, 0x6f, 0xd3, 0x3e,
	0x18, 0x6d, 0xda, 0xae, 0xfb, 0xd5, 0xd5, 0x7e, 0x42, 0xd6, 0x98, 0xac, 0x68, 0x0a, 0x55, 0xb5,
	0x43, 0x2e, 0xa4, 0x50, 0xd0, 0xc4, 0x75, 0xa5, 0x87, 0x21, 0x54, 0x09, 0x32, 0x10, 0x12, 0x17,
	0xe4, 0xc6, 0x9f, 0x92, 0xa8, 0x4d, 0x1c, 0x6c, 0x17, 0xa5, 0x70, 0xe5, 0x0f, 0xe0, 0xcf, 0xda,
	0x71, 0x47, 0x4e, 0x08, 0xb5, 0x7f, 0xc5, 0x6e, 0xc8, 0x69, 0xd2, 0xa4, 0xa3, 0x62, 0x4c, 0xdc,
	0xfc, 0xbd, 0xf7, 0xbd, 0x67, 0xfb, 0x7d, 0x96, 0x91, 0xc9, 0x16, 0x11, 0xc4, 0x32, 0xe4, 0x71,
	0x5f, 0xf0, 0xd9, 0x8c, 0x26, 0x49, 0x5f, 0xa5, 0x4e, 0x22, 0xb8, 0xe2, 0xd8, 0xda, 0x70, 0xe9,
	0xe2, 0xb3, 0xb3, 0x29, 0x9c, 0xbc, 0xd1, 0xb4, 0x7f, 0xd7, 0x4e, 0x66, 0xdc, 0x9b, 0x7e, 0x60,
	0x20, 0x3d, 0x11, 0x26, 0x8a, 0x8b, 0xb5, 0x93, 0x79, 0xbc, 0xa3, 0x93, 0xc6, 0xd3, 0x9c, 0x3d,
	0xf4, 0xb9, 0xcf, 0xb3, 0x65, 0x5f, 0xaf, 0xd6, 0x68, 0xef, 0xba, 0x8e, 0xee, 0x8d, 0xa5, 0xff,
	0x5c, 0x00, 0x55, 0xe0, 0xae, 0x55, 0x98, 0xa0, 0x7d, 0x4f, 0x03, 0x5c, 0x10, 0xa3, 0x6b, 0xd8,
	0x6d, 0xb7, 0x28, 0xf1, 0x31, 0x6a, 0xe7, 0xd6, 0x2f, 0x18, 0xa9, 0x67, 0x5c, 0x09, 0xe0, 0x2e,
	0x6a, 0x7b, 0x9c, 0xc1, 0x85, 0xa2, 0x51, 0x42, 0x1a, 0x9a, 0x1d, 0xd6, 0x89, 0xe1, 0x96, 0x20,
	0x3e, 0x41, 0x1d, 0x1f, 0x62, 0x90, 0xa1, 0x7c, 0x45, 0x55, 0x40, 0x9a, 0x9b, 0x9e, 0x2a, 0x8c,
	0x4f, 0xd1, 0x61, 0x44, 0xd3, 0x77, 0xa1, 0x0a, 0x02, 0x3e, 0x63, 0x61, 0xec, 0x0f, 0xf5, 0x85,
	0x25, 0xd9, 0xeb, 0x1a, 0x76, 0x33, 0x6b, 0xdf, 0xc9, 0xe3, 0x13, 0x74, 0x10, 0xd1, 0xf4, 0x02,
	0x3e, 0xce, 0x21, 0xf6, 0x40, 0x48, 0xd2, 0xd2, 0x02, 0x77, 0x1b, 0xc4, 0x4f, 0xd1, 0xfd, 0x04,
	0x44, 0x14, 0x4a, 0x9d, 0x14, 0xb0, 0x33, 0xc6, 0x04, 0x48, 0x09, 0x92, 0xec, 0x77, 0x1b, 0x76,
	0xdb, 0xdd, 0x4d, 0xe2, 0xd7, 0xa8, 0x1d, 0x81, 0xa2, 0x8c, 0x2a, 0x2a, 0xc9, 0x7f, 0xdd, 0x86,
	0xdd, 0x19, 0x3c, 0x74, 0xfe, 0x3c, 0x3a, 0xe7, 0x0d, 0x9f, 0x42, 0x3c, 0xce, 0x55, 0xc3, 0xe6,
	0xe5, 0x8f, 0x07, 0x35, 0xb7, 0x74, 0xe9, 0x99, 0x88, 0xdc, 0x8c, 0xde, 0x05, 0x99, 0xf0, 0x58,
	0x42, 0xef, 0x6b, 0x1d, 0xfd, 0x3f, 0x96, 0xfe, 0xdb, 0x84, 0x51, 0xa5, 0xb3, 0x53, 0xf0, 0x0f,
	0x53, 0xe9, 0x48, 0x45, 0x85, 0x3a, 0x87, 0xd0, 0x0f, 0x54, 0x36, 0x97, 0xa6, 0x5b, 0x85, 0xb4,
	0x3e, 0x9e, 0x47, 0x79, 0xc8, 0xcd, 0x8c, 0x2f, 0x01, 0x7c, 0x84, 0x5a, 0xa3, 0xb3, 0x6c, 0x5c,
	0x7b, 0x99, 0x75, 0x5e, 0xe9, 0xf3, 0x7c, 0x02, 0xa1, 0x2f, 0x9c, 0xe7, 0x5c, 0x94, 0xf8, 0x1c,
	0x35, 0x86, 0x23, 0x9d, 0xa7, 0x61, 0x77, 0x06, 0x8f, 0x6e, 0x4b, 0x29, 0xdb, 0x66, 0xb4, 0x79,
	0xcc, 0x32, 0x0f, 0x4a, 0x5b, 0xf4, 0x08, 0x3a, 0xda, 0x4e, 0xa1, 0x08, 0x68, 0x70, 0x6d, 0xa0,
	0xc6, 0x58, 0xfa, 0xf8, 0x0b, 0x3a, 0xd8, 0x7e, 0xbc, 0xb7, 0xee, 0x77, 0x33, 0x73, 0xf3, 0xd9,
	0x5d, 0x15, 0xc5, 0x21, 0xf0, 0x1c, 0x75, 0xaa, 0x13, 0x72, 0xfe, 0xc2, 0xa8, 0xd2, 0x6f, 0x9e,
	0xde, 0xad, 0xbf, 0xd8, 0x76, 0xf8, 0xf2, 0x72, 0x69, 0x19, 0x57, 0x4b, 0xcb, 0xf8, 0xb9, 0xb4,
	0x8c, 0x6f, 0x2b, 0xab, 0x76, 0xb5, 0xb2, 0x6a, 0xdf, 0x57, 0x56, 0xed, 0xfd, 0x63, 0x3f, 0x54,
	0xc1, 0x7c, 0xe2, 0x78, 0x3c, 0xea, 0x57, 0xbd, 0xcb, 0xa2, 0x9f, 0x96, 0x5f, 0xd0, 0x22, 0x01,
	0x39, 0x69, 0x65, 0x1f, 0xc1, 0x93, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x3e, 0x6c, 0x7e,
	0xa4, 0x04, 0x00, 0x00,
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
	CreateRollapp(ctx context.Context, in *MsgCreateRollapp, opts ...grpc.CallOption) (*MsgCreateRollappResponse, error)
	UpdateState(ctx context.Context, in *MsgUpdateState, opts ...grpc.CallOption) (*MsgUpdateStateResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateRollapp(ctx context.Context, in *MsgCreateRollapp, opts ...grpc.CallOption) (*MsgCreateRollappResponse, error) {
	out := new(MsgCreateRollappResponse)
	err := c.cc.Invoke(ctx, "/dymensionxyz.dymension.rollapp.Msg/CreateRollapp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateState(ctx context.Context, in *MsgUpdateState, opts ...grpc.CallOption) (*MsgUpdateStateResponse, error) {
	out := new(MsgUpdateStateResponse)
	err := c.cc.Invoke(ctx, "/dymensionxyz.dymension.rollapp.Msg/UpdateState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateRollapp(context.Context, *MsgCreateRollapp) (*MsgCreateRollappResponse, error)
	UpdateState(context.Context, *MsgUpdateState) (*MsgUpdateStateResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateRollapp(ctx context.Context, req *MsgCreateRollapp) (*MsgCreateRollappResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRollapp not implemented")
}
func (*UnimplementedMsgServer) UpdateState(ctx context.Context, req *MsgUpdateState) (*MsgUpdateStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateState not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateRollapp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateRollapp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateRollapp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dymensionxyz.dymension.rollapp.Msg/CreateRollapp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateRollapp(ctx, req.(*MsgCreateRollapp))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateState)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dymensionxyz.dymension.rollapp.Msg/UpdateState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateState(ctx, req.(*MsgUpdateState))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dymensionxyz.dymension.rollapp.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRollapp",
			Handler:    _Msg_CreateRollapp_Handler,
		},
		{
			MethodName: "UpdateState",
			Handler:    _Msg_UpdateState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dymension/rollapp/tx.proto",
}

func (m *MsgCreateRollapp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateRollapp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateRollapp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Metadatas) > 0 {
		for iNdEx := len(m.Metadatas) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Metadatas[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.PermissionedAddresses) > 0 {
		for iNdEx := len(m.PermissionedAddresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PermissionedAddresses[iNdEx])
			copy(dAtA[i:], m.PermissionedAddresses[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.PermissionedAddresses[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.MaxSequencers != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.MaxSequencers))
		i--
		dAtA[i] = 0x30
	}
	if m.MaxWithholdingBlocks != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.MaxWithholdingBlocks))
		i--
		dAtA[i] = 0x28
	}
	if len(m.GenesisPath) > 0 {
		i -= len(m.GenesisPath)
		copy(dAtA[i:], m.GenesisPath)
		i = encodeVarintTx(dAtA, i, uint64(len(m.GenesisPath)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.CodeStamp) > 0 {
		i -= len(m.CodeStamp)
		copy(dAtA[i:], m.CodeStamp)
		i = encodeVarintTx(dAtA, i, uint64(len(m.CodeStamp)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RollappId) > 0 {
		i -= len(m.RollappId)
		copy(dAtA[i:], m.RollappId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.RollappId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateRollappResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateRollappResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateRollappResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.BDs.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if m.Version != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x30
	}
	if len(m.DAPath) > 0 {
		i -= len(m.DAPath)
		copy(dAtA[i:], m.DAPath)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DAPath)))
		i--
		dAtA[i] = 0x2a
	}
	if m.NumBlocks != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.NumBlocks))
		i--
		dAtA[i] = 0x20
	}
	if m.StartHeight != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.StartHeight))
		i--
		dAtA[i] = 0x18
	}
	if len(m.RollappId) > 0 {
		i -= len(m.RollappId)
		copy(dAtA[i:], m.RollappId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.RollappId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateStateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateStateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateStateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
func (m *MsgCreateRollapp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.RollappId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.CodeStamp)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.GenesisPath)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.MaxWithholdingBlocks != 0 {
		n += 1 + sovTx(uint64(m.MaxWithholdingBlocks))
	}
	if m.MaxSequencers != 0 {
		n += 1 + sovTx(uint64(m.MaxSequencers))
	}
	if len(m.PermissionedAddresses) > 0 {
		for _, s := range m.PermissionedAddresses {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if len(m.Metadatas) > 0 {
		for _, e := range m.Metadatas {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgCreateRollappResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.RollappId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.StartHeight != 0 {
		n += 1 + sovTx(uint64(m.StartHeight))
	}
	if m.NumBlocks != 0 {
		n += 1 + sovTx(uint64(m.NumBlocks))
	}
	l = len(m.DAPath)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovTx(uint64(m.Version))
	}
	l = m.BDs.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateStateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateRollapp) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateRollapp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateRollapp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappId", wireType)
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
			m.RollappId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeStamp", wireType)
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
			m.CodeStamp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisPath", wireType)
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
			m.GenesisPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxWithholdingBlocks", wireType)
			}
			m.MaxWithholdingBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxWithholdingBlocks |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSequencers", wireType)
			}
			m.MaxSequencers = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSequencers |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PermissionedAddresses", wireType)
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
			m.PermissionedAddresses = append(m.PermissionedAddresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadatas", wireType)
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
			m.Metadatas = append(m.Metadatas, TokenMetadata{})
			if err := m.Metadatas[len(m.Metadatas)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
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
func (m *MsgCreateRollappResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateRollappResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateRollappResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgUpdateState) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappId", wireType)
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
			m.RollappId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartHeight", wireType)
			}
			m.StartHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumBlocks", wireType)
			}
			m.NumBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumBlocks |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DAPath", wireType)
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
			m.DAPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BDs", wireType)
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
			if err := m.BDs.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
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
func (m *MsgUpdateStateResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateStateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateStateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
