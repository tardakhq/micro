// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test/dep-test/dep-test-service/proto/dep/dep.proto

package service_dep

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Message struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5d3178ded6195e0, []int{0}
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

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5d3178ded6195e0, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5d3178ded6195e0, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type StreamingRequest struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingRequest) Reset()         { *m = StreamingRequest{} }
func (m *StreamingRequest) String() string { return proto.CompactTextString(m) }
func (*StreamingRequest) ProtoMessage()    {}
func (*StreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5d3178ded6195e0, []int{3}
}

func (m *StreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingRequest.Unmarshal(m, b)
}
func (m *StreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingRequest.Marshal(b, m, deterministic)
}
func (m *StreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingRequest.Merge(m, src)
}
func (m *StreamingRequest) XXX_Size() int {
	return xxx_messageInfo_StreamingRequest.Size(m)
}
func (m *StreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingRequest proto.InternalMessageInfo

func (m *StreamingRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type StreamingResponse struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingResponse) Reset()         { *m = StreamingResponse{} }
func (m *StreamingResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingResponse) ProtoMessage()    {}
func (*StreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5d3178ded6195e0, []int{4}
}

func (m *StreamingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingResponse.Unmarshal(m, b)
}
func (m *StreamingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingResponse.Marshal(b, m, deterministic)
}
func (m *StreamingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingResponse.Merge(m, src)
}
func (m *StreamingResponse) XXX_Size() int {
	return xxx_messageInfo_StreamingResponse.Size(m)
}
func (m *StreamingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingResponse proto.InternalMessageInfo

func (m *StreamingResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Ping struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5d3178ded6195e0, []int{5}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

type Pong struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5d3178ded6195e0, []int{6}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "service.dep.Message")
	proto.RegisterType((*Request)(nil), "service.dep.Request")
	proto.RegisterType((*Response)(nil), "service.dep.Response")
	proto.RegisterType((*StreamingRequest)(nil), "service.dep.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "service.dep.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "service.dep.Ping")
	proto.RegisterType((*Pong)(nil), "service.dep.Pong")
}

func init() {
	proto.RegisterFile("test/dep-test/dep-test-service/proto/dep/dep.proto", fileDescriptor_b5d3178ded6195e0)
}

var fileDescriptor_b5d3178ded6195e0 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xbb, 0x24, 0xa6, 0x75, 0xbc, 0xb4, 0x4b, 0x15, 0x89, 0xb6, 0xc8, 0x9e, 0xe2, 0xa1,
	0x69, 0xa9, 0xe8, 0x0b, 0xe8, 0x4d, 0x04, 0x89, 0x4f, 0x10, 0xdb, 0x61, 0x09, 0x36, 0xbb, 0x31,
	0xb3, 0x15, 0x7c, 0x3c, 0xdf, 0x4c, 0x76, 0xb3, 0x81, 0x26, 0xb4, 0x87, 0xc0, 0xcc, 0xff, 0xfd,
	0x33, 0xfc, 0x93, 0x85, 0xb5, 0x41, 0x32, 0xcb, 0x2d, 0x56, 0x8b, 0x4e, 0xb1, 0x20, 0xac, 0x7f,
	0x8a, 0x0d, 0x2e, 0xab, 0x5a, 0x1b, 0x6d, 0x65, 0xfb, 0xa5, 0xae, 0xe3, 0x17, 0x1e, 0xa6, 0x5b,
	0xac, 0xc4, 0x0d, 0x0c, 0xdf, 0x90, 0x28, 0x97, 0xc8, 0xc7, 0x10, 0x50, 0xfe, 0x7b, 0xcd, 0xee,
	0x58, 0x72, 0x9e, 0xd9, 0x52, 0xcc, 0x60, 0x98, 0xe1, 0xf7, 0x1e, 0xc9, 0x70, 0x0e, 0xa1, 0xca,
	0x4b, 0xf4, 0xd4, 0xd5, 0xe2, 0x16, 0x46, 0x19, 0x52, 0xa5, 0x15, 0xb9, 0xe1, 0x92, 0x64, 0x3b,
	0x5c, 0x92, 0x14, 0x09, 0x8c, 0x3f, 0x4c, 0x8d, 0x79, 0x59, 0x28, 0xd9, 0x6e, 0x99, 0xc2, 0xd9,
	0x46, 0xef, 0x95, 0x71, 0xbe, 0x20, 0x6b, 0x1a, 0x71, 0x0f, 0x93, 0x03, 0xa7, 0x5f, 0x78, 0xdc,
	0x3a, 0x87, 0xf0, 0xbd, 0x50, 0x92, 0x5f, 0x41, 0x44, 0xa6, 0xd6, 0x5f, 0xe8, 0xb1, 0xef, 0x1c,
	0xd7, 0xa7, 0xf9, 0xfa, 0x8f, 0x41, 0xf0, 0x82, 0x15, 0x7f, 0x84, 0xf0, 0x39, 0xdf, 0xed, 0xf8,
	0x34, 0x3d, 0xf8, 0x19, 0xa9, 0x8f, 0x19, 0x5f, 0xf6, 0xd4, 0x26, 0x92, 0x18, 0xf0, 0x57, 0x88,
	0x9a, 0xa4, 0x7c, 0xd6, 0xb1, 0xf4, 0x0f, 0x8d, 0xe7, 0xa7, 0x70, 0xbb, 0x6a, 0xc5, 0xf8, 0x13,
	0x8c, 0xec, 0x2d, 0x2e, 0xef, 0xa4, 0xe3, 0xb7, 0x72, 0xdc, 0x93, 0xb4, 0x92, 0x62, 0x90, 0xb0,
	0x15, 0xfb, 0x8c, 0xdc, 0x33, 0x3e, 0xfc, 0x07, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x43, 0x0b, 0x86,
	0xfc, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DepClient is the client API for Dep service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DepClient interface {
	Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...grpc.CallOption) (Dep_StreamClient, error)
	PingPong(ctx context.Context, opts ...grpc.CallOption) (Dep_PingPongClient, error)
}

type depClient struct {
	cc *grpc.ClientConn
}

func NewDepClient(cc *grpc.ClientConn) DepClient {
	return &depClient{cc}
}

func (c *depClient) Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/service.dep.Dep/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *depClient) Stream(ctx context.Context, in *StreamingRequest, opts ...grpc.CallOption) (Dep_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Dep_serviceDesc.Streams[0], "/service.dep.Dep/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &depStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Dep_StreamClient interface {
	Recv() (*StreamingResponse, error)
	grpc.ClientStream
}

type depStreamClient struct {
	grpc.ClientStream
}

func (x *depStreamClient) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *depClient) PingPong(ctx context.Context, opts ...grpc.CallOption) (Dep_PingPongClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Dep_serviceDesc.Streams[1], "/service.dep.Dep/PingPong", opts...)
	if err != nil {
		return nil, err
	}
	x := &depPingPongClient{stream}
	return x, nil
}

type Dep_PingPongClient interface {
	Send(*Ping) error
	Recv() (*Pong, error)
	grpc.ClientStream
}

type depPingPongClient struct {
	grpc.ClientStream
}

func (x *depPingPongClient) Send(m *Ping) error {
	return x.ClientStream.SendMsg(m)
}

func (x *depPingPongClient) Recv() (*Pong, error) {
	m := new(Pong)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DepServer is the server API for Dep service.
type DepServer interface {
	Call(context.Context, *Request) (*Response, error)
	Stream(*StreamingRequest, Dep_StreamServer) error
	PingPong(Dep_PingPongServer) error
}

// UnimplementedDepServer can be embedded to have forward compatible implementations.
type UnimplementedDepServer struct {
}

func (*UnimplementedDepServer) Call(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (*UnimplementedDepServer) Stream(req *StreamingRequest, srv Dep_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (*UnimplementedDepServer) PingPong(srv Dep_PingPongServer) error {
	return status.Errorf(codes.Unimplemented, "method PingPong not implemented")
}

func RegisterDepServer(s *grpc.Server, srv DepServer) {
	s.RegisterService(&_Dep_serviceDesc, srv)
}

func _Dep_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.dep.Dep/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepServer).Call(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dep_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamingRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DepServer).Stream(m, &depStreamServer{stream})
}

type Dep_StreamServer interface {
	Send(*StreamingResponse) error
	grpc.ServerStream
}

type depStreamServer struct {
	grpc.ServerStream
}

func (x *depStreamServer) Send(m *StreamingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Dep_PingPong_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DepServer).PingPong(&depPingPongServer{stream})
}

type Dep_PingPongServer interface {
	Send(*Pong) error
	Recv() (*Ping, error)
	grpc.ServerStream
}

type depPingPongServer struct {
	grpc.ServerStream
}

func (x *depPingPongServer) Send(m *Pong) error {
	return x.ServerStream.SendMsg(m)
}

func (x *depPingPongServer) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Dep_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.dep.Dep",
	HandlerType: (*DepServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _Dep_Call_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Dep_Stream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PingPong",
			Handler:       _Dep_PingPong_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "test/dep-test/dep-test-service/proto/dep/dep.proto",
}
