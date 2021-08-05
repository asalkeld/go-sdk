// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// SecretServiceClient is the client API for SecretService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecretServiceClient interface {
	// Updates a secret, creating a new one if it doesn't already exist
	Put(ctx context.Context, in *SecretPutRequest, opts ...grpc.CallOption) (*SecretPutResponse, error)
	// Gets a secret from a Secret Store
	Access(ctx context.Context, in *SecretAccessRequest, opts ...grpc.CallOption) (*SecretAccessResponse, error)
}

type secretServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecretServiceClient(cc grpc.ClientConnInterface) SecretServiceClient {
	return &secretServiceClient{cc}
}

func (c *secretServiceClient) Put(ctx context.Context, in *SecretPutRequest, opts ...grpc.CallOption) (*SecretPutResponse, error) {
	out := new(SecretPutResponse)
	err := c.cc.Invoke(ctx, "/nitric.secret.v1.SecretService/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretServiceClient) Access(ctx context.Context, in *SecretAccessRequest, opts ...grpc.CallOption) (*SecretAccessResponse, error) {
	out := new(SecretAccessResponse)
	err := c.cc.Invoke(ctx, "/nitric.secret.v1.SecretService/Access", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretServiceServer is the server API for SecretService service.
// All implementations should embed UnimplementedSecretServiceServer
// for forward compatibility
type SecretServiceServer interface {
	// Updates a secret, creating a new one if it doesn't already exist
	Put(context.Context, *SecretPutRequest) (*SecretPutResponse, error)
	// Gets a secret from a Secret Store
	Access(context.Context, *SecretAccessRequest) (*SecretAccessResponse, error)
}

// UnimplementedSecretServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSecretServiceServer struct {
}

func (UnimplementedSecretServiceServer) Put(context.Context, *SecretPutRequest) (*SecretPutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedSecretServiceServer) Access(context.Context, *SecretAccessRequest) (*SecretAccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Access not implemented")
}

// UnsafeSecretServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecretServiceServer will
// result in compilation errors.
type UnsafeSecretServiceServer interface {
	mustEmbedUnimplementedSecretServiceServer()
}

func RegisterSecretServiceServer(s grpc.ServiceRegistrar, srv SecretServiceServer) {
	s.RegisterService(&SecretService_ServiceDesc, srv)
}

func _SecretService_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecretPutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nitric.secret.v1.SecretService/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).Put(ctx, req.(*SecretPutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretService_Access_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecretAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).Access(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nitric.secret.v1.SecretService/Access",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).Access(ctx, req.(*SecretAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SecretService_ServiceDesc is the grpc.ServiceDesc for SecretService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecretService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nitric.secret.v1.SecretService",
	HandlerType: (*SecretServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _SecretService_Put_Handler,
		},
		{
			MethodName: "Access",
			Handler:    _SecretService_Access_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "secret/v1/secret.proto",
}
