// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: pkg/accommodation/accommodation.proto

package accommodation

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AccommodationServiceClient is the client API for AccommodationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccommodationServiceClient interface {
	CreateAccommodation(ctx context.Context, in *AccommodationRequest, opts ...grpc.CallOption) (*AccommodationResponse, error)
	GetAccommodationById(ctx context.Context, in *AccommodationResponse, opts ...grpc.CallOption) (*AccommodationInfo, error)
	AccommodationSearch(ctx context.Context, in *AccommodationSearchRequest, opts ...grpc.CallOption) (*AccommodationSearchResponse, error)
	DeleteByHost(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type accommodationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccommodationServiceClient(cc grpc.ClientConnInterface) AccommodationServiceClient {
	return &accommodationServiceClient{cc}
}

func (c *accommodationServiceClient) CreateAccommodation(ctx context.Context, in *AccommodationRequest, opts ...grpc.CallOption) (*AccommodationResponse, error) {
	out := new(AccommodationResponse)
	err := c.cc.Invoke(ctx, "/accommodation.AccommodationService/CreateAccommodation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetAccommodationById(ctx context.Context, in *AccommodationResponse, opts ...grpc.CallOption) (*AccommodationInfo, error) {
	out := new(AccommodationInfo)
	err := c.cc.Invoke(ctx, "/accommodation.AccommodationService/GetAccommodationById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) AccommodationSearch(ctx context.Context, in *AccommodationSearchRequest, opts ...grpc.CallOption) (*AccommodationSearchResponse, error) {
	out := new(AccommodationSearchResponse)
	err := c.cc.Invoke(ctx, "/accommodation.AccommodationService/AccommodationSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) DeleteByHost(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/accommodation.AccommodationService/DeleteByHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccommodationServiceServer is the server API for AccommodationService service.
// All implementations must embed UnimplementedAccommodationServiceServer
// for forward compatibility
type AccommodationServiceServer interface {
	CreateAccommodation(context.Context, *AccommodationRequest) (*AccommodationResponse, error)
	GetAccommodationById(context.Context, *AccommodationResponse) (*AccommodationInfo, error)
	AccommodationSearch(context.Context, *AccommodationSearchRequest) (*AccommodationSearchResponse, error)
	DeleteByHost(context.Context, *IdRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedAccommodationServiceServer()
}

// UnimplementedAccommodationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccommodationServiceServer struct {
}

func (UnimplementedAccommodationServiceServer) CreateAccommodation(context.Context, *AccommodationRequest) (*AccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccommodation not implemented")
}
func (UnimplementedAccommodationServiceServer) GetAccommodationById(context.Context, *AccommodationResponse) (*AccommodationInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccommodationById not implemented")
}
func (UnimplementedAccommodationServiceServer) AccommodationSearch(context.Context, *AccommodationSearchRequest) (*AccommodationSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccommodationSearch not implemented")
}
func (UnimplementedAccommodationServiceServer) DeleteByHost(context.Context, *IdRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteByHost not implemented")
}
func (UnimplementedAccommodationServiceServer) mustEmbedUnimplementedAccommodationServiceServer() {}

// UnsafeAccommodationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccommodationServiceServer will
// result in compilation errors.
type UnsafeAccommodationServiceServer interface {
	mustEmbedUnimplementedAccommodationServiceServer()
}

func RegisterAccommodationServiceServer(s grpc.ServiceRegistrar, srv AccommodationServiceServer) {
	s.RegisterService(&AccommodationService_ServiceDesc, srv)
}

func _AccommodationService_CreateAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).CreateAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation.AccommodationService/CreateAccommodation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).CreateAccommodation(ctx, req.(*AccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetAccommodationById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccommodationResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetAccommodationById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation.AccommodationService/GetAccommodationById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAccommodationById(ctx, req.(*AccommodationResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_AccommodationSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccommodationSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).AccommodationSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation.AccommodationService/AccommodationSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).AccommodationSearch(ctx, req.(*AccommodationSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_DeleteByHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).DeleteByHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation.AccommodationService/DeleteByHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).DeleteByHost(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccommodationService_ServiceDesc is the grpc.ServiceDesc for AccommodationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccommodationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accommodation.AccommodationService",
	HandlerType: (*AccommodationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccommodation",
			Handler:    _AccommodationService_CreateAccommodation_Handler,
		},
		{
			MethodName: "GetAccommodationById",
			Handler:    _AccommodationService_GetAccommodationById_Handler,
		},
		{
			MethodName: "AccommodationSearch",
			Handler:    _AccommodationService_AccommodationSearch_Handler,
		},
		{
			MethodName: "DeleteByHost",
			Handler:    _AccommodationService_DeleteByHost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/accommodation/accommodation.proto",
}
