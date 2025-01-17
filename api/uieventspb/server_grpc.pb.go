// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: uieventspb/server.proto

package uieventspb

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

const (
	UIEvents_Store_FullMethodName = "/uievents.UIEvents/Store"
)

// UIEventsClient is the client API for UIEvents service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UIEventsClient interface {
	// Store persists received UI events for further processing.
	Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error)
}

type uIEventsClient struct {
	cc grpc.ClientConnInterface
}

func NewUIEventsClient(cc grpc.ClientConnInterface) UIEventsClient {
	return &uIEventsClient{cc}
}

func (c *uIEventsClient) Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error) {
	out := new(StoreResponse)
	err := c.cc.Invoke(ctx, UIEvents_Store_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UIEventsServer is the server API for UIEvents service.
// All implementations must embed UnimplementedUIEventsServer
// for forward compatibility
type UIEventsServer interface {
	// Store persists received UI events for further processing.
	Store(context.Context, *StoreRequest) (*StoreResponse, error)
	mustEmbedUnimplementedUIEventsServer()
}

// UnimplementedUIEventsServer must be embedded to have forward compatible implementations.
type UnimplementedUIEventsServer struct{}

func (UnimplementedUIEventsServer) Store(context.Context, *StoreRequest) (*StoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (UnimplementedUIEventsServer) mustEmbedUnimplementedUIEventsServer() {}

// UnsafeUIEventsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UIEventsServer will
// result in compilation errors.
type UnsafeUIEventsServer interface {
	mustEmbedUnimplementedUIEventsServer()
}

func RegisterUIEventsServer(s grpc.ServiceRegistrar, srv UIEventsServer) {
	s.RegisterService(&UIEvents_ServiceDesc, srv)
}

func _UIEvents_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UIEventsServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UIEvents_Store_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UIEventsServer).Store(ctx, req.(*StoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UIEvents_ServiceDesc is the grpc.ServiceDesc for UIEvents service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UIEvents_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "uievents.UIEvents",
	HandlerType: (*UIEventsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Store",
			Handler:    _UIEvents_Store_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "uieventspb/server.proto",
}
