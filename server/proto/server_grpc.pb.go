// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: server/proto/server.proto

package server

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ServerProtoClient is the client API for ServerProto service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerProtoClient interface {
	SendServer(ctx context.Context, in *ServerRequest, opts ...grpc.CallOption) (*ServerReply, error)
}

type serverProtoClient struct {
	cc grpc.ClientConnInterface
}

func NewServerProtoClient(cc grpc.ClientConnInterface) ServerProtoClient {
	return &serverProtoClient{cc}
}

func (c *serverProtoClient) SendServer(ctx context.Context, in *ServerRequest, opts ...grpc.CallOption) (*ServerReply, error) {
	out := new(ServerReply)
	err := c.cc.Invoke(ctx, "/metaltune.ServerProto/SendServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerProtoServer is the server API for ServerProto service.
// All implementations must embed UnimplementedServerProtoServer
// for forward compatibility
type ServerProtoServer interface {
	SendServer(context.Context, *ServerRequest) (*ServerReply, error)
	mustEmbedUnimplementedServerProtoServer()
}

// UnimplementedServerProtoServer must be embedded to have forward compatible implementations.
type UnimplementedServerProtoServer struct {
}

func (UnimplementedServerProtoServer) SendServer(context.Context, *ServerRequest) (*ServerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendServer not implemented")
}
func (UnimplementedServerProtoServer) mustEmbedUnimplementedServerProtoServer() {}

// UnsafeServerProtoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerProtoServer will
// result in compilation errors.
type UnsafeServerProtoServer interface {
	mustEmbedUnimplementedServerProtoServer()
}

func RegisterServerProtoServer(s grpc.ServiceRegistrar, srv ServerProtoServer) {
	s.RegisterService(&ServerProto_ServiceDesc, srv)
}

func _ServerProto_SendServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerProtoServer).SendServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metaltune.ServerProto/SendServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerProtoServer).SendServer(ctx, req.(*ServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServerProto_ServiceDesc is the grpc.ServiceDesc for ServerProto service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerProto_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "metaltune.ServerProto",
	HandlerType: (*ServerProtoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendServer",
			Handler:    _ServerProto_SendServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/proto/server.proto",
}
