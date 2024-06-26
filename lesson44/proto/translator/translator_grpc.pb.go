// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: translator.proto

package translator

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

// TranslatorServerClient is the client API for TranslatorServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TranslatorServerClient interface {
	Translate(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type translatorServerClient struct {
	cc grpc.ClientConnInterface
}

func NewTranslatorServerClient(cc grpc.ClientConnInterface) TranslatorServerClient {
	return &translatorServerClient{cc}
}

func (c *translatorServerClient) Translate(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/TranslatorServer/Translate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TranslatorServerServer is the server API for TranslatorServer service.
// All implementations must embed UnimplementedTranslatorServerServer
// for forward compatibility
type TranslatorServerServer interface {
	Translate(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedTranslatorServerServer()
}

// UnimplementedTranslatorServerServer must be embedded to have forward compatible implementations.
type UnimplementedTranslatorServerServer struct {
}

func (UnimplementedTranslatorServerServer) Translate(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Translate not implemented")
}
func (UnimplementedTranslatorServerServer) mustEmbedUnimplementedTranslatorServerServer() {}

// UnsafeTranslatorServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TranslatorServerServer will
// result in compilation errors.
type UnsafeTranslatorServerServer interface {
	mustEmbedUnimplementedTranslatorServerServer()
}

func RegisterTranslatorServerServer(s grpc.ServiceRegistrar, srv TranslatorServerServer) {
	s.RegisterService(&TranslatorServer_ServiceDesc, srv)
}

func _TranslatorServer_Translate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslatorServerServer).Translate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TranslatorServer/Translate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslatorServerServer).Translate(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// TranslatorServer_ServiceDesc is the grpc.ServiceDesc for TranslatorServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TranslatorServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TranslatorServer",
	HandlerType: (*TranslatorServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Translate",
			Handler:    _TranslatorServer_Translate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "translator.proto",
}
