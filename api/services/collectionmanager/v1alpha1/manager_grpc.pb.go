// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: api/services/collectionmanager/v1alpha1/manager.proto

package v1alpha1

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

// CollectionManagerClient is the client API for CollectionManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CollectionManagerClient interface {
	PublishContent(ctx context.Context, in *Publish_Request, opts ...grpc.CallOption) (*Publish_Response, error)
	// RetrieveContent retrieve content based on the Control Message.
	RetrieveContent(ctx context.Context, in *Retrieve_Request, opts ...grpc.CallOption) (*Retrieve_Response, error)
}

type collectionManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewCollectionManagerClient(cc grpc.ClientConnInterface) CollectionManagerClient {
	return &collectionManagerClient{cc}
}

func (c *collectionManagerClient) PublishContent(ctx context.Context, in *Publish_Request, opts ...grpc.CallOption) (*Publish_Response, error) {
	out := new(Publish_Response)
	err := c.cc.Invoke(ctx, "/manager.CollectionManager/PublishContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionManagerClient) RetrieveContent(ctx context.Context, in *Retrieve_Request, opts ...grpc.CallOption) (*Retrieve_Response, error) {
	out := new(Retrieve_Response)
	err := c.cc.Invoke(ctx, "/manager.CollectionManager/RetrieveContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollectionManagerServer is the server API for CollectionManager service.
// All implementations must embed UnimplementedCollectionManagerServer
// for forward compatibility
type CollectionManagerServer interface {
	PublishContent(context.Context, *Publish_Request) (*Publish_Response, error)
	// RetrieveContent retrieve content based on the Control Message.
	RetrieveContent(context.Context, *Retrieve_Request) (*Retrieve_Response, error)
	mustEmbedUnimplementedCollectionManagerServer()
}

// UnimplementedCollectionManagerServer must be embedded to have forward compatible implementations.
type UnimplementedCollectionManagerServer struct {
}

func (UnimplementedCollectionManagerServer) PublishContent(context.Context, *Publish_Request) (*Publish_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishContent not implemented")
}
func (UnimplementedCollectionManagerServer) RetrieveContent(context.Context, *Retrieve_Request) (*Retrieve_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveContent not implemented")
}
func (UnimplementedCollectionManagerServer) mustEmbedUnimplementedCollectionManagerServer() {}

// UnsafeCollectionManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CollectionManagerServer will
// result in compilation errors.
type UnsafeCollectionManagerServer interface {
	mustEmbedUnimplementedCollectionManagerServer()
}

func RegisterCollectionManagerServer(s grpc.ServiceRegistrar, srv CollectionManagerServer) {
	s.RegisterService(&CollectionManager_ServiceDesc, srv)
}

func _CollectionManager_PublishContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Publish_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionManagerServer).PublishContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.CollectionManager/PublishContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionManagerServer).PublishContent(ctx, req.(*Publish_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectionManager_RetrieveContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Retrieve_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionManagerServer).RetrieveContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.CollectionManager/RetrieveContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionManagerServer).RetrieveContent(ctx, req.(*Retrieve_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// CollectionManager_ServiceDesc is the grpc.ServiceDesc for CollectionManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CollectionManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "manager.CollectionManager",
	HandlerType: (*CollectionManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishContent",
			Handler:    _CollectionManager_PublishContent_Handler,
		},
		{
			MethodName: "RetrieveContent",
			Handler:    _CollectionManager_RetrieveContent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/services/collectionmanager/v1alpha1/manager.proto",
}
