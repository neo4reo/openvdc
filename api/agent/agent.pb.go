// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent.proto

/*
Package agent is a generated protocol buffer package.

It is generated from these files:
	agent.proto

It has these top-level messages:
*/
package agent

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import model "github.com/axsh/openvdc/model"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ResourceCollector service

type ResourceCollectorClient interface {
	GetResources(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*model.ComputingResources, error)
}

type resourceCollectorClient struct {
	cc *grpc.ClientConn
}

func NewResourceCollectorClient(cc *grpc.ClientConn) ResourceCollectorClient {
	return &resourceCollectorClient{cc}
}

func (c *resourceCollectorClient) GetResources(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*model.ComputingResources, error) {
	out := new(model.ComputingResources)
	err := grpc.Invoke(ctx, "/agent.ResourceCollector/GetResources", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ResourceCollector service

type ResourceCollectorServer interface {
	GetResources(context.Context, *google_protobuf.Empty) (*model.ComputingResources, error)
}

func RegisterResourceCollectorServer(s *grpc.Server, srv ResourceCollectorServer) {
	s.RegisterService(&_ResourceCollector_serviceDesc, srv)
}

func _ResourceCollector_GetResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceCollectorServer).GetResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.ResourceCollector/GetResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceCollectorServer).GetResources(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceCollector_serviceDesc = grpc.ServiceDesc{
	ServiceName: "agent.ResourceCollector",
	HandlerType: (*ResourceCollectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetResources",
			Handler:    _ResourceCollector_GetResources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent.proto",
}

func init() { proto.RegisterFile("agent.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x31, 0xab, 0xc2, 0x30,
	0x10, 0xc7, 0xdf, 0xf2, 0xde, 0xd0, 0x27, 0x82, 0x1d, 0x04, 0xbb, 0xa9, 0x7b, 0x02, 0xf6, 0x1b,
	0x58, 0xc4, 0xdd, 0x49, 0xdc, 0x92, 0xf4, 0x9a, 0x16, 0x92, 0x5c, 0x48, 0x2e, 0x62, 0xbf, 0xbd,
	0x98, 0x56, 0x71, 0x39, 0xb8, 0xff, 0xf1, 0xfb, 0xdd, 0x5d, 0xf1, 0x2f, 0x34, 0x38, 0x62, 0x3e,
	0x20, 0x61, 0xf9, 0x9b, 0x9b, 0x6a, 0x19, 0x20, 0x62, 0x0a, 0x0a, 0xa6, 0xb8, 0xaa, 0xf5, 0x40,
	0x7d, 0x92, 0x4c, 0xa1, 0xe5, 0x1a, 0x8d, 0x70, 0x9a, 0xe7, 0x81, 0x4c, 0x1d, 0xf7, 0x34, 0x7a,
	0x88, 0x1c, 0xac, 0xa7, 0x71, 0xaa, 0x13, 0x74, 0xb8, 0x16, 0xab, 0xcb, 0xac, 0x69, 0xd0, 0x18,
	0x50, 0x84, 0xa1, 0x6c, 0x8a, 0xc5, 0x19, 0xe8, 0x9d, 0xc7, 0x72, 0xcd, 0x34, 0xa2, 0x36, 0xf3,
	0x22, 0x99, 0x3a, 0x76, 0x7a, 0x29, 0xaa, 0x0d, 0xb3, 0xd8, 0x82, 0x61, 0x0d, 0x5a, 0x9f, 0x68,
	0x70, 0xfa, 0x83, 0xec, 0x7e, 0x8e, 0xfb, 0xdb, 0xf6, 0xeb, 0x20, 0xf1, 0x88, 0x3d, 0x47, 0x0f,
	0xee, 0xde, 0x2a, 0x2e, 0xfc, 0xc0, 0xf3, 0x0f, 0xf2, 0x2f, 0x1b, 0xeb, 0x67, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xce, 0x9c, 0x37, 0x83, 0xe0, 0x00, 0x00, 0x00,
}
