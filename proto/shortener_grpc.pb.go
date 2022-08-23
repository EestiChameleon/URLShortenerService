// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: proto/shortener.proto

package proto

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

// ShortenerClient is the client API for Shortener service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShortenerClient interface {
	GetOrigURL(ctx context.Context, in *GetOrigURLRequest, opts ...grpc.CallOption) (*GetOrigURLResponse, error)
	GetAllPairs(ctx context.Context, in *GetGetAllPairsRequest, opts ...grpc.CallOption) (*GetGetAllPairsResponse, error)
	PostProvideShortURL(ctx context.Context, in *PostProvideShortURLRequest, opts ...grpc.CallOption) (*PostProvideShortURLResponse, error)
	PostBatch(ctx context.Context, in *PostBatchRequest, opts ...grpc.CallOption) (*PostBatchResponse, error)
	DeleteBatch(ctx context.Context, in *DelBatchRequest, opts ...grpc.CallOption) (*DelBatchResponse, error)
	GetStat(ctx context.Context, in *GetStatRequest, opts ...grpc.CallOption) (*GetStatResponse, error)
}

type shortenerClient struct {
	cc grpc.ClientConnInterface
}

func NewShortenerClient(cc grpc.ClientConnInterface) ShortenerClient {
	return &shortenerClient{cc}
}

func (c *shortenerClient) GetOrigURL(ctx context.Context, in *GetOrigURLRequest, opts ...grpc.CallOption) (*GetOrigURLResponse, error) {
	out := new(GetOrigURLResponse)
	err := c.cc.Invoke(ctx, "/URLShortenerService.proto.Shortener/GetOrigURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerClient) GetAllPairs(ctx context.Context, in *GetGetAllPairsRequest, opts ...grpc.CallOption) (*GetGetAllPairsResponse, error) {
	out := new(GetGetAllPairsResponse)
	err := c.cc.Invoke(ctx, "/URLShortenerService.proto.Shortener/GetAllPairs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerClient) PostProvideShortURL(ctx context.Context, in *PostProvideShortURLRequest, opts ...grpc.CallOption) (*PostProvideShortURLResponse, error) {
	out := new(PostProvideShortURLResponse)
	err := c.cc.Invoke(ctx, "/URLShortenerService.proto.Shortener/PostProvideShortURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerClient) PostBatch(ctx context.Context, in *PostBatchRequest, opts ...grpc.CallOption) (*PostBatchResponse, error) {
	out := new(PostBatchResponse)
	err := c.cc.Invoke(ctx, "/URLShortenerService.proto.Shortener/PostBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerClient) DeleteBatch(ctx context.Context, in *DelBatchRequest, opts ...grpc.CallOption) (*DelBatchResponse, error) {
	out := new(DelBatchResponse)
	err := c.cc.Invoke(ctx, "/URLShortenerService.proto.Shortener/DeleteBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerClient) GetStat(ctx context.Context, in *GetStatRequest, opts ...grpc.CallOption) (*GetStatResponse, error) {
	out := new(GetStatResponse)
	err := c.cc.Invoke(ctx, "/URLShortenerService.proto.Shortener/GetStat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortenerServer is the server API for Shortener service.
// All implementations must embed UnimplementedShortenerServer
// for forward compatibility
type ShortenerServer interface {
	GetOrigURL(context.Context, *GetOrigURLRequest) (*GetOrigURLResponse, error)
	GetAllPairs(context.Context, *GetGetAllPairsRequest) (*GetGetAllPairsResponse, error)
	PostProvideShortURL(context.Context, *PostProvideShortURLRequest) (*PostProvideShortURLResponse, error)
	PostBatch(context.Context, *PostBatchRequest) (*PostBatchResponse, error)
	DeleteBatch(context.Context, *DelBatchRequest) (*DelBatchResponse, error)
	GetStat(context.Context, *GetStatRequest) (*GetStatResponse, error)
	mustEmbedUnimplementedShortenerServer()
}

// UnimplementedShortenerServer must be embedded to have forward compatible implementations.
type UnimplementedShortenerServer struct {
}

func (UnimplementedShortenerServer) GetOrigURL(context.Context, *GetOrigURLRequest) (*GetOrigURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrigURL not implemented")
}
func (UnimplementedShortenerServer) GetAllPairs(context.Context, *GetGetAllPairsRequest) (*GetGetAllPairsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPairs not implemented")
}
func (UnimplementedShortenerServer) PostProvideShortURL(context.Context, *PostProvideShortURLRequest) (*PostProvideShortURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostProvideShortURL not implemented")
}
func (UnimplementedShortenerServer) PostBatch(context.Context, *PostBatchRequest) (*PostBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostBatch not implemented")
}
func (UnimplementedShortenerServer) DeleteBatch(context.Context, *DelBatchRequest) (*DelBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBatch not implemented")
}
func (UnimplementedShortenerServer) GetStat(context.Context, *GetStatRequest) (*GetStatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStat not implemented")
}
func (UnimplementedShortenerServer) mustEmbedUnimplementedShortenerServer() {}

// UnsafeShortenerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShortenerServer will
// result in compilation errors.
type UnsafeShortenerServer interface {
	mustEmbedUnimplementedShortenerServer()
}

func RegisterShortenerServer(s grpc.ServiceRegistrar, srv ShortenerServer) {
	s.RegisterService(&Shortener_ServiceDesc, srv)
}

func _Shortener_GetOrigURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrigURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).GetOrigURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/URLShortenerService.proto.Shortener/GetOrigURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).GetOrigURL(ctx, req.(*GetOrigURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shortener_GetAllPairs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGetAllPairsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).GetAllPairs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/URLShortenerService.proto.Shortener/GetAllPairs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).GetAllPairs(ctx, req.(*GetGetAllPairsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shortener_PostProvideShortURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostProvideShortURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).PostProvideShortURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/URLShortenerService.proto.Shortener/PostProvideShortURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).PostProvideShortURL(ctx, req.(*PostProvideShortURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shortener_PostBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).PostBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/URLShortenerService.proto.Shortener/PostBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).PostBatch(ctx, req.(*PostBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shortener_DeleteBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).DeleteBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/URLShortenerService.proto.Shortener/DeleteBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).DeleteBatch(ctx, req.(*DelBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shortener_GetStat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).GetStat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/URLShortenerService.proto.Shortener/GetStat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).GetStat(ctx, req.(*GetStatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Shortener_ServiceDesc is the grpc.ServiceDesc for Shortener service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Shortener_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "URLShortenerService.proto.Shortener",
	HandlerType: (*ShortenerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrigURL",
			Handler:    _Shortener_GetOrigURL_Handler,
		},
		{
			MethodName: "GetAllPairs",
			Handler:    _Shortener_GetAllPairs_Handler,
		},
		{
			MethodName: "PostProvideShortURL",
			Handler:    _Shortener_PostProvideShortURL_Handler,
		},
		{
			MethodName: "PostBatch",
			Handler:    _Shortener_PostBatch_Handler,
		},
		{
			MethodName: "DeleteBatch",
			Handler:    _Shortener_DeleteBatch_Handler,
		},
		{
			MethodName: "GetStat",
			Handler:    _Shortener_GetStat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/shortener.proto",
}