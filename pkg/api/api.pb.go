// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/api.proto

package api

import (
	context "context"
	fmt "fmt"
	math "math"

	msg "github.com/btwiuse/wetty/pkg/msg"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ChunkRequest struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Request              []byte   `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChunkRequest) Reset()         { *m = ChunkRequest{} }
func (m *ChunkRequest) String() string { return proto.CompactTextString(m) }
func (*ChunkRequest) ProtoMessage()    {}
func (*ChunkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{0}
}

func (m *ChunkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChunkRequest.Unmarshal(m, b)
}
func (m *ChunkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChunkRequest.Marshal(b, m, deterministic)
}
func (m *ChunkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChunkRequest.Merge(m, src)
}
func (m *ChunkRequest) XXX_Size() int {
	return xxx_messageInfo_ChunkRequest.Size(m)
}
func (m *ChunkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChunkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChunkRequest proto.InternalMessageInfo

func (m *ChunkRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ChunkRequest) GetRequest() []byte {
	if m != nil {
		return m.Request
	}
	return nil
}

type Chunk struct {
	Chunk                []byte   `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chunk) Reset()         { *m = Chunk{} }
func (m *Chunk) String() string { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()    {}
func (*Chunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{1}
}

func (m *Chunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chunk.Unmarshal(m, b)
}
func (m *Chunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chunk.Marshal(b, m, deterministic)
}
func (m *Chunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chunk.Merge(m, src)
}
func (m *Chunk) XXX_Size() int {
	return xxx_messageInfo_Chunk.Size(m)
}
func (m *Chunk) XXX_DiscardUnknown() {
	xxx_messageInfo_Chunk.DiscardUnknown(m)
}

var xxx_messageInfo_Chunk proto.InternalMessageInfo

func (m *Chunk) GetChunk() []byte {
	if m != nil {
		return m.Chunk
	}
	return nil
}

type ForwardMessage struct {
	Head                 string   `protobuf:"bytes,1,opt,name=head,proto3" json:"head,omitempty"`
	Body                 []byte   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForwardMessage) Reset()         { *m = ForwardMessage{} }
func (m *ForwardMessage) String() string { return proto.CompactTextString(m) }
func (*ForwardMessage) ProtoMessage()    {}
func (*ForwardMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{2}
}

func (m *ForwardMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForwardMessage.Unmarshal(m, b)
}
func (m *ForwardMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForwardMessage.Marshal(b, m, deterministic)
}
func (m *ForwardMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForwardMessage.Merge(m, src)
}
func (m *ForwardMessage) XXX_Size() int {
	return xxx_messageInfo_ForwardMessage.Size(m)
}
func (m *ForwardMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ForwardMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ForwardMessage proto.InternalMessageInfo

func (m *ForwardMessage) GetHead() string {
	if m != nil {
		return m.Head
	}
	return ""
}

func (m *ForwardMessage) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Message struct {
	Type                 msg.Type `protobuf:"varint,1,opt,name=type,proto3,enum=msg.Type" json:"type,omitempty"`
	Body                 []byte   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{3}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetType() msg.Type {
	if m != nil {
		return m.Type
	}
	return msg.Type_CLIENT_INPUT
}

func (m *Message) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*ChunkRequest)(nil), "api.ChunkRequest")
	proto.RegisterType((*Chunk)(nil), "api.Chunk")
	proto.RegisterType((*ForwardMessage)(nil), "api.ForwardMessage")
	proto.RegisterType((*Message)(nil), "api.Message")
}

func init() { proto.RegisterFile("api/api.proto", fileDescriptor_1b40cafcd4234784) }

var fileDescriptor_1b40cafcd4234784 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0xc1, 0x4a, 0xc3, 0x40,
	0x14, 0xec, 0x6a, 0x6a, 0xe8, 0x23, 0x56, 0x7c, 0x7a, 0x08, 0x81, 0x42, 0xd9, 0x53, 0x0e, 0x92,
	0x96, 0x7a, 0x51, 0xe8, 0x4d, 0xf0, 0xe6, 0x25, 0xf5, 0x07, 0xb6, 0xe6, 0x91, 0x04, 0x69, 0x76,
	0xdd, 0x4d, 0x91, 0xfc, 0x8b, 0x1f, 0x2b, 0xfb, 0x92, 0xa0, 0x85, 0xde, 0xe6, 0xcd, 0x30, 0xb3,
	0x3b, 0x03, 0xd7, 0xca, 0xd4, 0x2b, 0x65, 0xea, 0xcc, 0x58, 0xdd, 0x6a, 0xbc, 0x54, 0xa6, 0x4e,
	0x6e, 0x0e, 0xae, 0x5c, 0xb5, 0x9d, 0x21, 0xd7, 0xb3, 0x72, 0x0b, 0xd1, 0x4b, 0x75, 0x6c, 0x3e,
	0x73, 0xfa, 0x3a, 0x92, 0x6b, 0x11, 0x21, 0x30, 0xaa, 0xad, 0x62, 0xb1, 0x14, 0xe9, 0x2c, 0x67,
	0x8c, 0x31, 0x84, 0xb6, 0x97, 0xe3, 0x8b, 0xa5, 0x48, 0xa3, 0x7c, 0x3c, 0xe5, 0x02, 0xa6, 0xec,
	0xc6, 0x7b, 0x98, 0x7e, 0x78, 0xc0, 0xbe, 0x28, 0xef, 0x0f, 0xf9, 0x04, 0xf3, 0x57, 0x6d, 0xbf,
	0x95, 0x2d, 0xde, 0xc8, 0x39, 0x55, 0x92, 0x8f, 0xaf, 0x48, 0x15, 0x63, 0xbc, 0xc7, 0x9e, 0xdb,
	0xeb, 0xa2, 0x1b, 0xb2, 0x19, 0xcb, 0x2d, 0x84, 0xa3, 0x65, 0x01, 0x81, 0xff, 0x30, 0x5b, 0xe6,
	0x9b, 0x59, 0x76, 0x70, 0x65, 0xf6, 0xde, 0x19, 0xca, 0x99, 0x3e, 0xe7, 0xde, 0xfc, 0x08, 0x08,
	0x77, 0xe4, 0x5c, 0xad, 0x1b, 0x4c, 0x21, 0xd8, 0x51, 0x53, 0x60, 0x94, 0xf9, 0x29, 0x86, 0xd0,
	0xe4, 0xe4, 0x92, 0x93, 0x54, 0xac, 0x05, 0x3e, 0x40, 0xc8, 0x65, 0xc8, 0xe2, 0x2d, 0xcb, 0xff,
	0x87, 0x49, 0xe0, 0x8f, 0x92, 0x93, 0xb5, 0xc0, 0x67, 0x08, 0x87, 0x6e, 0x78, 0xc7, 0xd2, 0x69,
	0xd3, 0xe4, 0x1c, 0xd9, 0x3f, 0xb4, 0xbf, 0xe2, 0xe9, 0x1f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x92, 0x94, 0xf8, 0x76, 0xa1, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SessionClient is the client API for Session service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SessionClient interface {
	Send(ctx context.Context, opts ...grpc.CallOption) (Session_SendClient, error)
	Chunker(ctx context.Context, in *ChunkRequest, opts ...grpc.CallOption) (Session_ChunkerClient, error)
	Forward(ctx context.Context, opts ...grpc.CallOption) (Session_ForwardClient, error)
}

type sessionClient struct {
	cc *grpc.ClientConn
}

func NewSessionClient(cc *grpc.ClientConn) SessionClient {
	return &sessionClient{cc}
}

func (c *sessionClient) Send(ctx context.Context, opts ...grpc.CallOption) (Session_SendClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Session_serviceDesc.Streams[0], "/api.Session/Send", opts...)
	if err != nil {
		return nil, err
	}
	x := &sessionSendClient{stream}
	return x, nil
}

type Session_SendClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type sessionSendClient struct {
	grpc.ClientStream
}

func (x *sessionSendClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sessionSendClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sessionClient) Chunker(ctx context.Context, in *ChunkRequest, opts ...grpc.CallOption) (Session_ChunkerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Session_serviceDesc.Streams[1], "/api.Session/Chunker", opts...)
	if err != nil {
		return nil, err
	}
	x := &sessionChunkerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Session_ChunkerClient interface {
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type sessionChunkerClient struct {
	grpc.ClientStream
}

func (x *sessionChunkerClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sessionClient) Forward(ctx context.Context, opts ...grpc.CallOption) (Session_ForwardClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Session_serviceDesc.Streams[2], "/api.Session/Forward", opts...)
	if err != nil {
		return nil, err
	}
	x := &sessionForwardClient{stream}
	return x, nil
}

type Session_ForwardClient interface {
	Send(*ForwardMessage) error
	Recv() (*ForwardMessage, error)
	grpc.ClientStream
}

type sessionForwardClient struct {
	grpc.ClientStream
}

func (x *sessionForwardClient) Send(m *ForwardMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sessionForwardClient) Recv() (*ForwardMessage, error) {
	m := new(ForwardMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SessionServer is the server API for Session service.
type SessionServer interface {
	Send(Session_SendServer) error
	Chunker(*ChunkRequest, Session_ChunkerServer) error
	Forward(Session_ForwardServer) error
}

// UnimplementedSessionServer can be embedded to have forward compatible implementations.
type UnimplementedSessionServer struct {
}

func (*UnimplementedSessionServer) Send(srv Session_SendServer) error {
	return status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (*UnimplementedSessionServer) Chunker(req *ChunkRequest, srv Session_ChunkerServer) error {
	return status.Errorf(codes.Unimplemented, "method Chunker not implemented")
}
func (*UnimplementedSessionServer) Forward(srv Session_ForwardServer) error {
	return status.Errorf(codes.Unimplemented, "method Forward not implemented")
}

func RegisterSessionServer(s *grpc.Server, srv SessionServer) {
	s.RegisterService(&_Session_serviceDesc, srv)
}

func _Session_Send_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SessionServer).Send(&sessionSendServer{stream})
}

type Session_SendServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type sessionSendServer struct {
	grpc.ServerStream
}

func (x *sessionSendServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sessionSendServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Session_Chunker_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ChunkRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SessionServer).Chunker(m, &sessionChunkerServer{stream})
}

type Session_ChunkerServer interface {
	Send(*Chunk) error
	grpc.ServerStream
}

type sessionChunkerServer struct {
	grpc.ServerStream
}

func (x *sessionChunkerServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

func _Session_Forward_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SessionServer).Forward(&sessionForwardServer{stream})
}

type Session_ForwardServer interface {
	Send(*ForwardMessage) error
	Recv() (*ForwardMessage, error)
	grpc.ServerStream
}

type sessionForwardServer struct {
	grpc.ServerStream
}

func (x *sessionForwardServer) Send(m *ForwardMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sessionForwardServer) Recv() (*ForwardMessage, error) {
	m := new(ForwardMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Session_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Session",
	HandlerType: (*SessionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Send",
			Handler:       _Session_Send_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Chunker",
			Handler:       _Session_Chunker_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Forward",
			Handler:       _Session_Forward_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/api.proto",
}
