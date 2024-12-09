// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: parser.proto

package parser

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ParserService_ParserModel_FullMethodName = "/ParserService/ParserModel"
)

// ParserServiceClient is the client API for ParserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ParserServiceClient interface {
	ParserModel(ctx context.Context, in *ParserRequest, opts ...grpc.CallOption) (*ParserResponse, error)
}

type parserServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewParserServiceClient(cc grpc.ClientConnInterface) ParserServiceClient {
	return &parserServiceClient{cc}
}

func (c *parserServiceClient) ParserModel(ctx context.Context, in *ParserRequest, opts ...grpc.CallOption) (*ParserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ParserResponse)
	err := c.cc.Invoke(ctx, ParserService_ParserModel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParserServiceServer is the server API for ParserService service.
// All implementations must embed UnimplementedParserServiceServer
// for forward compatibility.
type ParserServiceServer interface {
	ParserModel(context.Context, *ParserRequest) (*ParserResponse, error)
	mustEmbedUnimplementedParserServiceServer()
}

// UnimplementedParserServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedParserServiceServer struct{}

func (UnimplementedParserServiceServer) ParserModel(context.Context, *ParserRequest) (*ParserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParserModel not implemented")
}
func (UnimplementedParserServiceServer) mustEmbedUnimplementedParserServiceServer() {}
func (UnimplementedParserServiceServer) testEmbeddedByValue()                       {}

// UnsafeParserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ParserServiceServer will
// result in compilation errors.
type UnsafeParserServiceServer interface {
	mustEmbedUnimplementedParserServiceServer()
}

func RegisterParserServiceServer(s grpc.ServiceRegistrar, srv ParserServiceServer) {
	// If the following call pancis, it indicates UnimplementedParserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ParserService_ServiceDesc, srv)
}

func _ParserService_ParserModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParserServiceServer).ParserModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ParserService_ParserModel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParserServiceServer).ParserModel(ctx, req.(*ParserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ParserService_ServiceDesc is the grpc.ServiceDesc for ParserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ParserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ParserService",
	HandlerType: (*ParserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ParserModel",
			Handler:    _ParserService_ParserModel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "parser.proto",
}