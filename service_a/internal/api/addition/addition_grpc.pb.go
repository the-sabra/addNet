// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: addition.proto

package addition

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
	AdditionService_Add_FullMethodName = "/addition.AdditionService/Add"
)

// AdditionServiceClient is the client API for AdditionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdditionServiceClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
}

type additionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdditionServiceClient(cc grpc.ClientConnInterface) AdditionServiceClient {
	return &additionServiceClient{cc}
}

func (c *additionServiceClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, AdditionService_Add_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdditionServiceServer is the server API for AdditionService service.
// All implementations must embed UnimplementedAdditionServiceServer
// for forward compatibility.
type AdditionServiceServer interface {
	Add(context.Context, *AddRequest) (*AddResponse, error)
	mustEmbedUnimplementedAdditionServiceServer()
}

// UnimplementedAdditionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAdditionServiceServer struct{}

func (UnimplementedAdditionServiceServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedAdditionServiceServer) mustEmbedUnimplementedAdditionServiceServer() {}
func (UnimplementedAdditionServiceServer) testEmbeddedByValue()                         {}

// UnsafeAdditionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdditionServiceServer will
// result in compilation errors.
type UnsafeAdditionServiceServer interface {
	mustEmbedUnimplementedAdditionServiceServer()
}

func RegisterAdditionServiceServer(s grpc.ServiceRegistrar, srv AdditionServiceServer) {
	// If the following call pancis, it indicates UnimplementedAdditionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AdditionService_ServiceDesc, srv)
}

func _AdditionService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdditionServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdditionService_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdditionServiceServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdditionService_ServiceDesc is the grpc.ServiceDesc for AdditionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdditionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "addition.AdditionService",
	HandlerType: (*AdditionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _AdditionService_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "addition.proto",
}
