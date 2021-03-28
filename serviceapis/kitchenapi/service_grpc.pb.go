// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package kitchenapi

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

// KitchenServiceClient is the client API for KitchenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KitchenServiceClient interface {
	GetRestaurant(ctx context.Context, in *GetRestaurantRequest, opts ...grpc.CallOption) (*GetRestaurantResponse, error)
	AcceptTicket(ctx context.Context, in *AcceptTicketRequest, opts ...grpc.CallOption) (*AcceptTicketResponse, error)
}

type kitchenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKitchenServiceClient(cc grpc.ClientConnInterface) KitchenServiceClient {
	return &kitchenServiceClient{cc}
}

func (c *kitchenServiceClient) GetRestaurant(ctx context.Context, in *GetRestaurantRequest, opts ...grpc.CallOption) (*GetRestaurantResponse, error) {
	out := new(GetRestaurantResponse)
	err := c.cc.Invoke(ctx, "/kitchenapi.KitchenService/GetRestaurant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kitchenServiceClient) AcceptTicket(ctx context.Context, in *AcceptTicketRequest, opts ...grpc.CallOption) (*AcceptTicketResponse, error) {
	out := new(AcceptTicketResponse)
	err := c.cc.Invoke(ctx, "/kitchenapi.KitchenService/AcceptTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KitchenServiceServer is the server API for KitchenService service.
// All implementations must embed UnimplementedKitchenServiceServer
// for forward compatibility
type KitchenServiceServer interface {
	GetRestaurant(context.Context, *GetRestaurantRequest) (*GetRestaurantResponse, error)
	AcceptTicket(context.Context, *AcceptTicketRequest) (*AcceptTicketResponse, error)
	mustEmbedUnimplementedKitchenServiceServer()
}

// UnimplementedKitchenServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKitchenServiceServer struct {
}

func (UnimplementedKitchenServiceServer) GetRestaurant(context.Context, *GetRestaurantRequest) (*GetRestaurantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRestaurant not implemented")
}
func (UnimplementedKitchenServiceServer) AcceptTicket(context.Context, *AcceptTicketRequest) (*AcceptTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptTicket not implemented")
}
func (UnimplementedKitchenServiceServer) mustEmbedUnimplementedKitchenServiceServer() {}

// UnsafeKitchenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KitchenServiceServer will
// result in compilation errors.
type UnsafeKitchenServiceServer interface {
	mustEmbedUnimplementedKitchenServiceServer()
}

func RegisterKitchenServiceServer(s grpc.ServiceRegistrar, srv KitchenServiceServer) {
	s.RegisterService(&KitchenService_ServiceDesc, srv)
}

func _KitchenService_GetRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRestaurantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KitchenServiceServer).GetRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kitchenapi.KitchenService/GetRestaurant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KitchenServiceServer).GetRestaurant(ctx, req.(*GetRestaurantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KitchenService_AcceptTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KitchenServiceServer).AcceptTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kitchenapi.KitchenService/AcceptTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KitchenServiceServer).AcceptTicket(ctx, req.(*AcceptTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KitchenService_ServiceDesc is the grpc.ServiceDesc for KitchenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KitchenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kitchenapi.KitchenService",
	HandlerType: (*KitchenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRestaurant",
			Handler:    _KitchenService_GetRestaurant_Handler,
		},
		{
			MethodName: "AcceptTicket",
			Handler:    _KitchenService_AcceptTicket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kitchenapi/service.proto",
}
