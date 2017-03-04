// Code generated by protoc-gen-go.
// source: rlistener.proto
// DO NOT EDIT!

/*
Package rlistener is a generated protocol buffer package.

It is generated from these files:
	rlistener.proto

It has these top-level messages:
	Buffer
*/
package rlistener

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Buffer struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Buffer) Reset()                    { *m = Buffer{} }
func (m *Buffer) String() string            { return proto.CompactTextString(m) }
func (*Buffer) ProtoMessage()               {}
func (*Buffer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Buffer) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Buffer)(nil), "rlistener.Buffer")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RemoteListener service

type RemoteListenerClient interface {
	Accept(ctx context.Context, opts ...grpc.CallOption) (RemoteListener_AcceptClient, error)
}

type remoteListenerClient struct {
	cc *grpc.ClientConn
}

func NewRemoteListenerClient(cc *grpc.ClientConn) RemoteListenerClient {
	return &remoteListenerClient{cc}
}

func (c *remoteListenerClient) Accept(ctx context.Context, opts ...grpc.CallOption) (RemoteListener_AcceptClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RemoteListener_serviceDesc.Streams[0], c.cc, "/rlistener.RemoteListener/Accept", opts...)
	if err != nil {
		return nil, err
	}
	x := &remoteListenerAcceptClient{stream}
	return x, nil
}

type RemoteListener_AcceptClient interface {
	Send(*Buffer) error
	Recv() (*Buffer, error)
	grpc.ClientStream
}

type remoteListenerAcceptClient struct {
	grpc.ClientStream
}

func (x *remoteListenerAcceptClient) Send(m *Buffer) error {
	return x.ClientStream.SendMsg(m)
}

func (x *remoteListenerAcceptClient) Recv() (*Buffer, error) {
	m := new(Buffer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for RemoteListener service

type RemoteListenerServer interface {
	Accept(RemoteListener_AcceptServer) error
}

func RegisterRemoteListenerServer(s *grpc.Server, srv RemoteListenerServer) {
	s.RegisterService(&_RemoteListener_serviceDesc, srv)
}

func _RemoteListener_Accept_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RemoteListenerServer).Accept(&remoteListenerAcceptServer{stream})
}

type RemoteListener_AcceptServer interface {
	Send(*Buffer) error
	Recv() (*Buffer, error)
	grpc.ServerStream
}

type remoteListenerAcceptServer struct {
	grpc.ServerStream
}

func (x *remoteListenerAcceptServer) Send(m *Buffer) error {
	return x.ServerStream.SendMsg(m)
}

func (x *remoteListenerAcceptServer) Recv() (*Buffer, error) {
	m := new(Buffer)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _RemoteListener_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rlistener.RemoteListener",
	HandlerType: (*RemoteListenerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Accept",
			Handler:       _RemoteListener_Accept_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "rlistener.proto",
}

func init() { proto.RegisterFile("rlistener.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xca, 0xc9, 0x2c,
	0x2e, 0x49, 0xcd, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0xc9, 0x70, 0xb1, 0x39, 0x95, 0xa6, 0xa5, 0xa5, 0x16, 0x09, 0x09, 0x71, 0xb1, 0xa4, 0x24, 0x96,
	0x24, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x04, 0x81, 0xd9, 0x46, 0x2e, 0x5c, 0x7c, 0x41, 0xa9,
	0xb9, 0xf9, 0x25, 0xa9, 0x3e, 0x50, 0xf5, 0x42, 0x46, 0x5c, 0x6c, 0x8e, 0xc9, 0xc9, 0xa9, 0x05,
	0x25, 0x42, 0x82, 0x7a, 0x08, 0x63, 0x21, 0x46, 0x48, 0x61, 0x0a, 0x69, 0x30, 0x1a, 0x30, 0x3a,
	0xe9, 0x45, 0xe9, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x97, 0x64,
	0x64, 0xe6, 0x55, 0xa4, 0x16, 0xe9, 0xc3, 0x95, 0xea, 0x83, 0x1d, 0x65, 0x0d, 0xe7, 0x27, 0xb1,
	0x81, 0x05, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x85, 0x60, 0x1c, 0x6f, 0xb8, 0x00, 0x00,
	0x00,
}
