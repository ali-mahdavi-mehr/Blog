// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/post.proto

package compiles

import (
	context "context"
	compiles "github.com/alima12/Blog-Go/service/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostServiceClient interface {
	GetSinglePost(ctx context.Context, in *compiles.RetrievePost, opts ...grpc.CallOption) (*compiles.SinglePostResponse, error)
}

type postServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostServiceClient(cc grpc.ClientConnInterface) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) GetSinglePost(ctx context.Context, in *compiles.RetrievePost, opts ...grpc.CallOption) (*compiles.SinglePostResponse, error) {
	out := new(compiles.SinglePostResponse)
	err := c.cc.Invoke(ctx, "/Post.PostService/GetSinglePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServiceServer is the server API for PostService service.
// All implementations must embed UnimplementedPostServiceServer
// for forward compatibility
type PostServiceServer interface {
	GetSinglePost(context.Context, *compiles.RetrievePost) (*compiles.SinglePostResponse, error)
	mustEmbedUnimplementedPostServiceServer()
}

// UnimplementedPostServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPostServiceServer struct {
}

func (UnimplementedPostServiceServer) GetSinglePost(context.Context, *compiles.RetrievePost) (*compiles.SinglePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSinglePost not implemented")
}
func (UnimplementedPostServiceServer) mustEmbedUnimplementedPostServiceServer() {}

// UnsafePostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostServiceServer will
// result in compilation errors.
type UnsafePostServiceServer interface {
	mustEmbedUnimplementedPostServiceServer()
}

func RegisterPostServiceServer(s grpc.ServiceRegistrar, srv PostServiceServer) {
	s.RegisterService(&PostService_ServiceDesc, srv)
}

func _PostService_GetSinglePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(compiles.RetrievePost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetSinglePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Post.PostService/GetSinglePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetSinglePost(ctx, req.(*compiles.RetrievePost))
	}
	return interceptor(ctx, in, info, handler)
}

// PostService_ServiceDesc is the grpc.ServiceDesc for PostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Post.PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSinglePost",
			Handler:    _PostService_GetSinglePost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/post.proto",
}
