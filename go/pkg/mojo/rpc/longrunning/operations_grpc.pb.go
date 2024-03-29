// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: mojo/rpc/longrunning/operations.proto

package longrunning

import (
	context "context"
	core "github.com/mojo-lang/core/go/pkg/mojo/core"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OperationsClient is the client API for Operations service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OperationsClient interface {
	ListOperations(ctx context.Context, in *ListOperationsRequest, opts ...grpc.CallOption) (*ListOperationsResponse, error)
	GetOperation(ctx context.Context, in *GetOperationRequest, opts ...grpc.CallOption) (*Operation, error)
	DeleteOperation(ctx context.Context, in *DeleteOperationRequest, opts ...grpc.CallOption) (*core.Null, error)
	CancelOperation(ctx context.Context, in *CancelOperationRequest, opts ...grpc.CallOption) (*core.Null, error)
	WaitOperation(ctx context.Context, in *WaitOperationRequest, opts ...grpc.CallOption) (*Operation, error)
}

type operationsClient struct {
	cc grpc.ClientConnInterface
}

func NewOperationsClient(cc grpc.ClientConnInterface) OperationsClient {
	return &operationsClient{cc}
}

func (c *operationsClient) ListOperations(ctx context.Context, in *ListOperationsRequest, opts ...grpc.CallOption) (*ListOperationsResponse, error) {
	out := new(ListOperationsResponse)
	err := c.cc.Invoke(ctx, "/mojo.rpc.longrunning.Operations/list_operations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsClient) GetOperation(ctx context.Context, in *GetOperationRequest, opts ...grpc.CallOption) (*Operation, error) {
	out := new(Operation)
	err := c.cc.Invoke(ctx, "/mojo.rpc.longrunning.Operations/get_operation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsClient) DeleteOperation(ctx context.Context, in *DeleteOperationRequest, opts ...grpc.CallOption) (*core.Null, error) {
	out := new(core.Null)
	err := c.cc.Invoke(ctx, "/mojo.rpc.longrunning.Operations/delete_operation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsClient) CancelOperation(ctx context.Context, in *CancelOperationRequest, opts ...grpc.CallOption) (*core.Null, error) {
	out := new(core.Null)
	err := c.cc.Invoke(ctx, "/mojo.rpc.longrunning.Operations/cancel_operation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsClient) WaitOperation(ctx context.Context, in *WaitOperationRequest, opts ...grpc.CallOption) (*Operation, error) {
	out := new(Operation)
	err := c.cc.Invoke(ctx, "/mojo.rpc.longrunning.Operations/wait_operation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperationsServer is the server API for Operations service.
// All implementations must embed UnimplementedOperationsServer
// for forward compatibility
type OperationsServer interface {
	ListOperations(context.Context, *ListOperationsRequest) (*ListOperationsResponse, error)
	GetOperation(context.Context, *GetOperationRequest) (*Operation, error)
	DeleteOperation(context.Context, *DeleteOperationRequest) (*core.Null, error)
	CancelOperation(context.Context, *CancelOperationRequest) (*core.Null, error)
	WaitOperation(context.Context, *WaitOperationRequest) (*Operation, error)
	mustEmbedUnimplementedOperationsServer()
}

// UnimplementedOperationsServer must be embedded to have forward compatible implementations.
type UnimplementedOperationsServer struct {
}

func (UnimplementedOperationsServer) ListOperations(context.Context, *ListOperationsRequest) (*ListOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedOperationsServer) GetOperation(context.Context, *GetOperationRequest) (*Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOperation not implemented")
}
func (UnimplementedOperationsServer) DeleteOperation(context.Context, *DeleteOperationRequest) (*core.Null, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOperation not implemented")
}
func (UnimplementedOperationsServer) CancelOperation(context.Context, *CancelOperationRequest) (*core.Null, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOperation not implemented")
}
func (UnimplementedOperationsServer) WaitOperation(context.Context, *WaitOperationRequest) (*Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitOperation not implemented")
}
func (UnimplementedOperationsServer) mustEmbedUnimplementedOperationsServer() {}

// UnsafeOperationsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OperationsServer will
// result in compilation errors.
type UnsafeOperationsServer interface {
	mustEmbedUnimplementedOperationsServer()
}

func RegisterOperationsServer(s grpc.ServiceRegistrar, srv OperationsServer) {
	s.RegisterService(&Operations_ServiceDesc, srv)
}

func _Operations_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mojo.rpc.longrunning.Operations/list_operations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServer).ListOperations(ctx, req.(*ListOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operations_GetOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServer).GetOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mojo.rpc.longrunning.Operations/get_operation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServer).GetOperation(ctx, req.(*GetOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operations_DeleteOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServer).DeleteOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mojo.rpc.longrunning.Operations/delete_operation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServer).DeleteOperation(ctx, req.(*DeleteOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operations_CancelOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServer).CancelOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mojo.rpc.longrunning.Operations/cancel_operation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServer).CancelOperation(ctx, req.(*CancelOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operations_WaitOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WaitOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServer).WaitOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mojo.rpc.longrunning.Operations/wait_operation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServer).WaitOperation(ctx, req.(*WaitOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Operations_ServiceDesc is the grpc.ServiceDesc for Operations service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Operations_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mojo.rpc.longrunning.Operations",
	HandlerType: (*OperationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "list_operations",
			Handler:    _Operations_ListOperations_Handler,
		},
		{
			MethodName: "get_operation",
			Handler:    _Operations_GetOperation_Handler,
		},
		{
			MethodName: "delete_operation",
			Handler:    _Operations_DeleteOperation_Handler,
		},
		{
			MethodName: "cancel_operation",
			Handler:    _Operations_CancelOperation_Handler,
		},
		{
			MethodName: "wait_operation",
			Handler:    _Operations_WaitOperation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mojo/rpc/longrunning/operations.proto",
}
