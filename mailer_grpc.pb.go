// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package Project_2

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

// MailerClient is the client API for Mailer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MailerClient interface {
	SendPass(ctx context.Context, in *MsgRequest, opts ...grpc.CallOption) (*MsgReply, error)
	RetrievePass(ctx context.Context, in *MsgRequest, opts ...grpc.CallOption) (*MsgReply, error)
}

type mailerClient struct {
	cc grpc.ClientConnInterface
}

func NewMailerClient(cc grpc.ClientConnInterface) MailerClient {
	return &mailerClient{cc}
}

func (c *mailerClient) SendPass(ctx context.Context, in *MsgRequest, opts ...grpc.CallOption) (*MsgReply, error) {
	out := new(MsgReply)
	err := c.cc.Invoke(ctx, "/Mailer/SendPass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailerClient) RetrievePass(ctx context.Context, in *MsgRequest, opts ...grpc.CallOption) (*MsgReply, error) {
	out := new(MsgReply)
	err := c.cc.Invoke(ctx, "/Mailer/RetrievePass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MailerServer is the server API for Mailer service.
// All implementations must embed UnimplementedMailerServer
// for forward compatibility
type MailerServer interface {
	SendPass(context.Context, *MsgRequest) (*MsgReply, error)
	RetrievePass(context.Context, *MsgRequest) (*MsgReply, error)
	mustEmbedUnimplementedMailerServer()
}

// UnimplementedMailerServer must be embedded to have forward compatible implementations.
type UnimplementedMailerServer struct {
}

func (UnimplementedMailerServer) SendPass(context.Context, *MsgRequest) (*MsgReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPass not implemented")
}
func (UnimplementedMailerServer) RetrievePass(context.Context, *MsgRequest) (*MsgReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrievePass not implemented")
}
func (UnimplementedMailerServer) mustEmbedUnimplementedMailerServer() {}

// UnsafeMailerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MailerServer will
// result in compilation errors.
type UnsafeMailerServer interface {
	mustEmbedUnimplementedMailerServer()
}

func RegisterMailerServer(s grpc.ServiceRegistrar, srv MailerServer) {
	s.RegisterService(&Mailer_ServiceDesc, srv)
}

func _Mailer_SendPass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailerServer).SendPass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Mailer/SendPass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailerServer).SendPass(ctx, req.(*MsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mailer_RetrievePass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailerServer).RetrievePass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Mailer/RetrievePass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailerServer).RetrievePass(ctx, req.(*MsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Mailer_ServiceDesc is the grpc.ServiceDesc for Mailer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mailer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Mailer",
	HandlerType: (*MailerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendPass",
			Handler:    _Mailer_SendPass_Handler,
		},
		{
			MethodName: "RetrievePass",
			Handler:    _Mailer_RetrievePass_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mailer.proto",
}
