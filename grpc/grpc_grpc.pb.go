// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: grpc/grpc.proto

package __

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

// ApiCallerScaleClient is the client API for ApiCallerScale service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiCallerScaleClient interface {
	ScalesMessageOutChannel(ctx context.Context, opts ...grpc.CallOption) (ApiCallerScale_ScalesMessageOutChannelClient, error)
	SetTare(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResponseSetScale, error)
	SetTareValue(ctx context.Context, in *RequestTareValue, opts ...grpc.CallOption) (*ResponseSetScale, error)
	SetZero(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResponseSetScale, error)
	GetInstantWeight(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResponseInstantWeight, error)
	GetState(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResponseScale, error)
}

type apiCallerScaleClient struct {
	cc grpc.ClientConnInterface
}

func NewApiCallerScaleClient(cc grpc.ClientConnInterface) ApiCallerScaleClient {
	return &apiCallerScaleClient{cc}
}

func (c *apiCallerScaleClient) ScalesMessageOutChannel(ctx context.Context, opts ...grpc.CallOption) (ApiCallerScale_ScalesMessageOutChannelClient, error) {
	stream, err := c.cc.NewStream(ctx, &ApiCallerScale_ServiceDesc.Streams[0], "/stream.ApiCallerScale/ScalesMessageOutChannel", opts...)
	if err != nil {
		return nil, err
	}
	x := &apiCallerScaleScalesMessageOutChannelClient{stream}
	return x, nil
}

type ApiCallerScale_ScalesMessageOutChannelClient interface {
	Send(*RequestScale) error
	Recv() (*ResponseScale, error)
	grpc.ClientStream
}

type apiCallerScaleScalesMessageOutChannelClient struct {
	grpc.ClientStream
}

func (x *apiCallerScaleScalesMessageOutChannelClient) Send(m *RequestScale) error {
	return x.ClientStream.SendMsg(m)
}

func (x *apiCallerScaleScalesMessageOutChannelClient) Recv() (*ResponseScale, error) {
	m := new(ResponseScale)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *apiCallerScaleClient) SetTare(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResponseSetScale, error) {
	out := new(ResponseSetScale)
	err := c.cc.Invoke(ctx, "/stream.ApiCallerScale/SetTare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiCallerScaleClient) SetTareValue(ctx context.Context, in *RequestTareValue, opts ...grpc.CallOption) (*ResponseSetScale, error) {
	out := new(ResponseSetScale)
	err := c.cc.Invoke(ctx, "/stream.ApiCallerScale/SetTareValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiCallerScaleClient) SetZero(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResponseSetScale, error) {
	out := new(ResponseSetScale)
	err := c.cc.Invoke(ctx, "/stream.ApiCallerScale/SetZero", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiCallerScaleClient) GetInstantWeight(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResponseInstantWeight, error) {
	out := new(ResponseInstantWeight)
	err := c.cc.Invoke(ctx, "/stream.ApiCallerScale/GetInstantWeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiCallerScaleClient) GetState(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResponseScale, error) {
	out := new(ResponseScale)
	err := c.cc.Invoke(ctx, "/stream.ApiCallerScale/GetState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiCallerScaleServer is the server API for ApiCallerScale service.
// All implementations must embed UnimplementedApiCallerScaleServer
// for forward compatibility
type ApiCallerScaleServer interface {
	ScalesMessageOutChannel(ApiCallerScale_ScalesMessageOutChannelServer) error
	SetTare(context.Context, *Empty) (*ResponseSetScale, error)
	SetTareValue(context.Context, *RequestTareValue) (*ResponseSetScale, error)
	SetZero(context.Context, *Empty) (*ResponseSetScale, error)
	GetInstantWeight(context.Context, *Empty) (*ResponseInstantWeight, error)
	GetState(context.Context, *Empty) (*ResponseScale, error)
	mustEmbedUnimplementedApiCallerScaleServer()
}

// UnimplementedApiCallerScaleServer must be embedded to have forward compatible implementations.
type UnimplementedApiCallerScaleServer struct {
}

func (UnimplementedApiCallerScaleServer) ScalesMessageOutChannel(ApiCallerScale_ScalesMessageOutChannelServer) error {
	return status.Errorf(codes.Unimplemented, "method ScalesMessageOutChannel not implemented")
}
func (UnimplementedApiCallerScaleServer) SetTare(context.Context, *Empty) (*ResponseSetScale, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTare not implemented")
}
func (UnimplementedApiCallerScaleServer) SetTareValue(context.Context, *RequestTareValue) (*ResponseSetScale, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTareValue not implemented")
}
func (UnimplementedApiCallerScaleServer) SetZero(context.Context, *Empty) (*ResponseSetScale, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetZero not implemented")
}
func (UnimplementedApiCallerScaleServer) GetInstantWeight(context.Context, *Empty) (*ResponseInstantWeight, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstantWeight not implemented")
}
func (UnimplementedApiCallerScaleServer) GetState(context.Context, *Empty) (*ResponseScale, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetState not implemented")
}
func (UnimplementedApiCallerScaleServer) mustEmbedUnimplementedApiCallerScaleServer() {}

// UnsafeApiCallerScaleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiCallerScaleServer will
// result in compilation errors.
type UnsafeApiCallerScaleServer interface {
	mustEmbedUnimplementedApiCallerScaleServer()
}

func RegisterApiCallerScaleServer(s grpc.ServiceRegistrar, srv ApiCallerScaleServer) {
	s.RegisterService(&ApiCallerScale_ServiceDesc, srv)
}

func _ApiCallerScale_ScalesMessageOutChannel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ApiCallerScaleServer).ScalesMessageOutChannel(&apiCallerScaleScalesMessageOutChannelServer{stream})
}

type ApiCallerScale_ScalesMessageOutChannelServer interface {
	Send(*ResponseScale) error
	Recv() (*RequestScale, error)
	grpc.ServerStream
}

type apiCallerScaleScalesMessageOutChannelServer struct {
	grpc.ServerStream
}

func (x *apiCallerScaleScalesMessageOutChannelServer) Send(m *ResponseScale) error {
	return x.ServerStream.SendMsg(m)
}

func (x *apiCallerScaleScalesMessageOutChannelServer) Recv() (*RequestScale, error) {
	m := new(RequestScale)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ApiCallerScale_SetTare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiCallerScaleServer).SetTare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stream.ApiCallerScale/SetTare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiCallerScaleServer).SetTare(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiCallerScale_SetTareValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestTareValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiCallerScaleServer).SetTareValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stream.ApiCallerScale/SetTareValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiCallerScaleServer).SetTareValue(ctx, req.(*RequestTareValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiCallerScale_SetZero_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiCallerScaleServer).SetZero(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stream.ApiCallerScale/SetZero",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiCallerScaleServer).SetZero(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiCallerScale_GetInstantWeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiCallerScaleServer).GetInstantWeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stream.ApiCallerScale/GetInstantWeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiCallerScaleServer).GetInstantWeight(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiCallerScale_GetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiCallerScaleServer).GetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stream.ApiCallerScale/GetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiCallerScaleServer).GetState(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiCallerScale_ServiceDesc is the grpc.ServiceDesc for ApiCallerScale service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiCallerScale_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stream.ApiCallerScale",
	HandlerType: (*ApiCallerScaleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetTare",
			Handler:    _ApiCallerScale_SetTare_Handler,
		},
		{
			MethodName: "SetTareValue",
			Handler:    _ApiCallerScale_SetTareValue_Handler,
		},
		{
			MethodName: "SetZero",
			Handler:    _ApiCallerScale_SetZero_Handler,
		},
		{
			MethodName: "GetInstantWeight",
			Handler:    _ApiCallerScale_GetInstantWeight_Handler,
		},
		{
			MethodName: "GetState",
			Handler:    _ApiCallerScale_GetState_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ScalesMessageOutChannel",
			Handler:       _ApiCallerScale_ScalesMessageOutChannel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc/grpc.proto",
}
