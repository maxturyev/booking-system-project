// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: src/grpc/hotel.proto

package grpc

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
	HotelService_GetHotels_FullMethodName         = "/grpc.HotelService/GetHotels"
	HotelService_GetHotelPriceByID_FullMethodName = "/grpc.HotelService/GetHotelPriceByID"
)

// HotelServiceClient is the client API for HotelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HotelServiceClient interface {
	GetHotels(ctx context.Context, in *GetHotelsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetHotelsResponse], error)
	GetHotelPriceByID(ctx context.Context, in *GetHotelPriceByIDRequest, opts ...grpc.CallOption) (*GetHotelPriceByIDResponse, error)
}

type hotelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHotelServiceClient(cc grpc.ClientConnInterface) HotelServiceClient {
	return &hotelServiceClient{cc}
}

func (c *hotelServiceClient) GetHotels(ctx context.Context, in *GetHotelsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetHotelsResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &HotelService_ServiceDesc.Streams[0], HotelService_GetHotels_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetHotelsRequest, GetHotelsResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type HotelService_GetHotelsClient = grpc.ServerStreamingClient[GetHotelsResponse]

func (c *hotelServiceClient) GetHotelPriceByID(ctx context.Context, in *GetHotelPriceByIDRequest, opts ...grpc.CallOption) (*GetHotelPriceByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetHotelPriceByIDResponse)
	err := c.cc.Invoke(ctx, HotelService_GetHotelPriceByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HotelServiceServer is the server API for HotelService service.
// All implementations must embed UnimplementedHotelServiceServer
// for forward compatibility.
type HotelServiceServer interface {
	GetHotels(*GetHotelsRequest, grpc.ServerStreamingServer[GetHotelsResponse]) error
	GetHotelPriceByID(context.Context, *GetHotelPriceByIDRequest) (*GetHotelPriceByIDResponse, error)
	mustEmbedUnimplementedHotelServiceServer()
}

// UnimplementedHotelServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHotelServiceServer struct{}

func (UnimplementedHotelServiceServer) GetHotels(*GetHotelsRequest, grpc.ServerStreamingServer[GetHotelsResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetHotels not implemented")
}
func (UnimplementedHotelServiceServer) GetHotelPriceByID(context.Context, *GetHotelPriceByIDRequest) (*GetHotelPriceByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHotelPriceByID not implemented")
}
func (UnimplementedHotelServiceServer) mustEmbedUnimplementedHotelServiceServer() {}
func (UnimplementedHotelServiceServer) testEmbeddedByValue()                      {}

// UnsafeHotelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HotelServiceServer will
// result in compilation errors.
type UnsafeHotelServiceServer interface {
	mustEmbedUnimplementedHotelServiceServer()
}

func RegisterHotelServiceServer(s grpc.ServiceRegistrar, srv HotelServiceServer) {
	// If the following call pancis, it indicates UnimplementedHotelServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HotelService_ServiceDesc, srv)
}

func _HotelService_GetHotels_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetHotelsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HotelServiceServer).GetHotels(m, &grpc.GenericServerStream[GetHotelsRequest, GetHotelsResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type HotelService_GetHotelsServer = grpc.ServerStreamingServer[GetHotelsResponse]

func _HotelService_GetHotelPriceByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHotelPriceByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).GetHotelPriceByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_GetHotelPriceByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).GetHotelPriceByID(ctx, req.(*GetHotelPriceByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HotelService_ServiceDesc is the grpc.ServiceDesc for HotelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HotelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.HotelService",
	HandlerType: (*HotelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHotelPriceByID",
			Handler:    _HotelService_GetHotelPriceByID_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetHotels",
			Handler:       _HotelService_GetHotels_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "src/grpc/hotel.proto",
}
