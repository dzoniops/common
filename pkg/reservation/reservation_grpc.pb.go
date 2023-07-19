// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: pkg/reservation/reservation.proto

package reservation

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

// ReservationServiceClient is the client API for ReservationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReservationServiceClient interface {
	// Sends a greeting
	// rpc Reserve(User) returns (RegisterResponse) {}
	// rpc Accept(LoginRequest) returns (LoginResponse) {}
	// rpc Decline(User) returns (Response) {}
	ActivateReservationsGuest(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*ActiveReservationsResponse, error)
	ActivateReservationsHost(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*ActiveReservationsResponse, error)
}

type reservationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReservationServiceClient(cc grpc.ClientConnInterface) ReservationServiceClient {
	return &reservationServiceClient{cc}
}

func (c *reservationServiceClient) ActivateReservationsGuest(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*ActiveReservationsResponse, error) {
	out := new(ActiveReservationsResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/ActivateReservationsGuest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) ActivateReservationsHost(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*ActiveReservationsResponse, error) {
	out := new(ActiveReservationsResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/ActivateReservationsHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReservationServiceServer is the server API for ReservationService service.
// All implementations must embed UnimplementedReservationServiceServer
// for forward compatibility
type ReservationServiceServer interface {
	// Sends a greeting
	// rpc Reserve(User) returns (RegisterResponse) {}
	// rpc Accept(LoginRequest) returns (LoginResponse) {}
	// rpc Decline(User) returns (Response) {}
	ActivateReservationsGuest(context.Context, *IdRequest) (*ActiveReservationsResponse, error)
	ActivateReservationsHost(context.Context, *IdRequest) (*ActiveReservationsResponse, error)
	mustEmbedUnimplementedReservationServiceServer()
}

// UnimplementedReservationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReservationServiceServer struct {
}

func (UnimplementedReservationServiceServer) ActivateReservationsGuest(context.Context, *IdRequest) (*ActiveReservationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateReservationsGuest not implemented")
}
func (UnimplementedReservationServiceServer) ActivateReservationsHost(context.Context, *IdRequest) (*ActiveReservationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateReservationsHost not implemented")
}
func (UnimplementedReservationServiceServer) mustEmbedUnimplementedReservationServiceServer() {}

// UnsafeReservationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReservationServiceServer will
// result in compilation errors.
type UnsafeReservationServiceServer interface {
	mustEmbedUnimplementedReservationServiceServer()
}

func RegisterReservationServiceServer(s grpc.ServiceRegistrar, srv ReservationServiceServer) {
	s.RegisterService(&ReservationService_ServiceDesc, srv)
}

func _ReservationService_ActivateReservationsGuest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).ActivateReservationsGuest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/ActivateReservationsGuest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).ActivateReservationsGuest(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_ActivateReservationsHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).ActivateReservationsHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/ActivateReservationsHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).ActivateReservationsHost(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReservationService_ServiceDesc is the grpc.ServiceDesc for ReservationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReservationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reservation.ReservationService",
	HandlerType: (*ReservationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ActivateReservationsGuest",
			Handler:    _ReservationService_ActivateReservationsGuest_Handler,
		},
		{
			MethodName: "ActivateReservationsHost",
			Handler:    _ReservationService_ActivateReservationsHost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/reservation/reservation.proto",
}
