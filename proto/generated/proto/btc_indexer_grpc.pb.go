// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: btc_indexer.proto

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

// BtcIndexerClient is the client API for BtcIndexer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BtcIndexerClient interface {
	AddressBalance(ctx context.Context, in *AddressBalanceRequest, opts ...grpc.CallOption) (*AddressBalanceReply, error)
}

type btcIndexerClient struct {
	cc grpc.ClientConnInterface
}

func NewBtcIndexerClient(cc grpc.ClientConnInterface) BtcIndexerClient {
	return &btcIndexerClient{cc}
}

func (c *btcIndexerClient) AddressBalance(ctx context.Context, in *AddressBalanceRequest, opts ...grpc.CallOption) (*AddressBalanceReply, error) {
	out := new(AddressBalanceReply)
	err := c.cc.Invoke(ctx, "/unindexer.btcindexer.BtcIndexer/AddressBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BtcIndexerServer is the server API for BtcIndexer service.
// All implementations must embed UnimplementedBtcIndexerServer
// for forward compatibility
type BtcIndexerServer interface {
	AddressBalance(context.Context, *AddressBalanceRequest) (*AddressBalanceReply, error)
	mustEmbedUnimplementedBtcIndexerServer()
}

// UnimplementedBtcIndexerServer must be embedded to have forward compatible implementations.
type UnimplementedBtcIndexerServer struct {
}

func (UnimplementedBtcIndexerServer) AddressBalance(context.Context, *AddressBalanceRequest) (*AddressBalanceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddressBalance not implemented")
}
func (UnimplementedBtcIndexerServer) mustEmbedUnimplementedBtcIndexerServer() {}

// UnsafeBtcIndexerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BtcIndexerServer will
// result in compilation errors.
type UnsafeBtcIndexerServer interface {
	mustEmbedUnimplementedBtcIndexerServer()
}

func RegisterBtcIndexerServer(s grpc.ServiceRegistrar, srv BtcIndexerServer) {
	s.RegisterService(&BtcIndexer_ServiceDesc, srv)
}

func _BtcIndexer_AddressBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BtcIndexerServer).AddressBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/unindexer.btcindexer.BtcIndexer/AddressBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BtcIndexerServer).AddressBalance(ctx, req.(*AddressBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BtcIndexer_ServiceDesc is the grpc.ServiceDesc for BtcIndexer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BtcIndexer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "unindexer.btcindexer.BtcIndexer",
	HandlerType: (*BtcIndexerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddressBalance",
			Handler:    _BtcIndexer_AddressBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "btc_indexer.proto",
}
