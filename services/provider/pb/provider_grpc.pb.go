// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package providerpb

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

// OCSProviderClient is the client API for OCSProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OCSProviderClient interface {
	// OnboardConsumer RPC call to validate the consumer and create StorageConsumer
	// resource on the StorageProvider cluster
	OnboardConsumer(ctx context.Context, in *OnboardConsumerRequest, opts ...grpc.CallOption) (*OnboardConsumerResponse, error)
	// GetStorageConfig RPC call to generate the json config for connecting to storage provider cluster
	GetStorageConfig(ctx context.Context, in *StorageConfigRequest, opts ...grpc.CallOption) (*StorageConfigResponse, error)
	// OffboardConsumer RPC call to delete StorageConsumer CR on the storage provider cluster.
	OffboardConsumer(ctx context.Context, in *OffboardConsumerRequest, opts ...grpc.CallOption) (*OffboardConsumerResponse, error)
	// UpdateCapacity PRC call to increase or decrease the block pool size
	UpdateCapacity(ctx context.Context, in *UpdateCapacityRequest, opts ...grpc.CallOption) (*UpdateCapacityResponse, error)
	// AcknowledgeOnboarding RPC call acknowledge the onboarding
	AcknowledgeOnboarding(ctx context.Context, in *AcknowledgeOnboardingRequest, opts ...grpc.CallOption) (*AcknowledgeOnboardingResponse, error)
}

type oCSProviderClient struct {
	cc grpc.ClientConnInterface
}

func NewOCSProviderClient(cc grpc.ClientConnInterface) OCSProviderClient {
	return &oCSProviderClient{cc}
}

func (c *oCSProviderClient) OnboardConsumer(ctx context.Context, in *OnboardConsumerRequest, opts ...grpc.CallOption) (*OnboardConsumerResponse, error) {
	out := new(OnboardConsumerResponse)
	err := c.cc.Invoke(ctx, "/provider.OCSProvider/OnboardConsumer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oCSProviderClient) GetStorageConfig(ctx context.Context, in *StorageConfigRequest, opts ...grpc.CallOption) (*StorageConfigResponse, error) {
	out := new(StorageConfigResponse)
	err := c.cc.Invoke(ctx, "/provider.OCSProvider/GetStorageConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oCSProviderClient) OffboardConsumer(ctx context.Context, in *OffboardConsumerRequest, opts ...grpc.CallOption) (*OffboardConsumerResponse, error) {
	out := new(OffboardConsumerResponse)
	err := c.cc.Invoke(ctx, "/provider.OCSProvider/OffboardConsumer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oCSProviderClient) UpdateCapacity(ctx context.Context, in *UpdateCapacityRequest, opts ...grpc.CallOption) (*UpdateCapacityResponse, error) {
	out := new(UpdateCapacityResponse)
	err := c.cc.Invoke(ctx, "/provider.OCSProvider/UpdateCapacity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oCSProviderClient) AcknowledgeOnboarding(ctx context.Context, in *AcknowledgeOnboardingRequest, opts ...grpc.CallOption) (*AcknowledgeOnboardingResponse, error) {
	out := new(AcknowledgeOnboardingResponse)
	err := c.cc.Invoke(ctx, "/provider.OCSProvider/AcknowledgeOnboarding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OCSProviderServer is the server API for OCSProvider service.
// All implementations must embed UnimplementedOCSProviderServer
// for forward compatibility
type OCSProviderServer interface {
	// OnboardConsumer RPC call to validate the consumer and create StorageConsumer
	// resource on the StorageProvider cluster
	OnboardConsumer(context.Context, *OnboardConsumerRequest) (*OnboardConsumerResponse, error)
	// GetStorageConfig RPC call to generate the json config for connecting to storage provider cluster
	GetStorageConfig(context.Context, *StorageConfigRequest) (*StorageConfigResponse, error)
	// OffboardConsumer RPC call to delete StorageConsumer CR on the storage provider cluster.
	OffboardConsumer(context.Context, *OffboardConsumerRequest) (*OffboardConsumerResponse, error)
	// UpdateCapacity PRC call to increase or decrease the block pool size
	UpdateCapacity(context.Context, *UpdateCapacityRequest) (*UpdateCapacityResponse, error)
	// AcknowledgeOnboarding RPC call acknowledge the onboarding
	AcknowledgeOnboarding(context.Context, *AcknowledgeOnboardingRequest) (*AcknowledgeOnboardingResponse, error)
	mustEmbedUnimplementedOCSProviderServer()
}

// UnimplementedOCSProviderServer must be embedded to have forward compatible implementations.
type UnimplementedOCSProviderServer struct {
}

func (UnimplementedOCSProviderServer) OnboardConsumer(context.Context, *OnboardConsumerRequest) (*OnboardConsumerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnboardConsumer not implemented")
}
func (UnimplementedOCSProviderServer) GetStorageConfig(context.Context, *StorageConfigRequest) (*StorageConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStorageConfig not implemented")
}
func (UnimplementedOCSProviderServer) OffboardConsumer(context.Context, *OffboardConsumerRequest) (*OffboardConsumerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OffboardConsumer not implemented")
}
func (UnimplementedOCSProviderServer) UpdateCapacity(context.Context, *UpdateCapacityRequest) (*UpdateCapacityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCapacity not implemented")
}
func (UnimplementedOCSProviderServer) AcknowledgeOnboarding(context.Context, *AcknowledgeOnboardingRequest) (*AcknowledgeOnboardingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcknowledgeOnboarding not implemented")
}
func (UnimplementedOCSProviderServer) mustEmbedUnimplementedOCSProviderServer() {}

// UnsafeOCSProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OCSProviderServer will
// result in compilation errors.
type UnsafeOCSProviderServer interface {
	mustEmbedUnimplementedOCSProviderServer()
}

func RegisterOCSProviderServer(s grpc.ServiceRegistrar, srv OCSProviderServer) {
	s.RegisterService(&OCSProvider_ServiceDesc, srv)
}

func _OCSProvider_OnboardConsumer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnboardConsumerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OCSProviderServer).OnboardConsumer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/provider.OCSProvider/OnboardConsumer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OCSProviderServer).OnboardConsumer(ctx, req.(*OnboardConsumerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OCSProvider_GetStorageConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OCSProviderServer).GetStorageConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/provider.OCSProvider/GetStorageConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OCSProviderServer).GetStorageConfig(ctx, req.(*StorageConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OCSProvider_OffboardConsumer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OffboardConsumerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OCSProviderServer).OffboardConsumer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/provider.OCSProvider/OffboardConsumer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OCSProviderServer).OffboardConsumer(ctx, req.(*OffboardConsumerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OCSProvider_UpdateCapacity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCapacityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OCSProviderServer).UpdateCapacity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/provider.OCSProvider/UpdateCapacity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OCSProviderServer).UpdateCapacity(ctx, req.(*UpdateCapacityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OCSProvider_AcknowledgeOnboarding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcknowledgeOnboardingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OCSProviderServer).AcknowledgeOnboarding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/provider.OCSProvider/AcknowledgeOnboarding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OCSProviderServer).AcknowledgeOnboarding(ctx, req.(*AcknowledgeOnboardingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OCSProvider_ServiceDesc is the grpc.ServiceDesc for OCSProvider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OCSProvider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "provider.OCSProvider",
	HandlerType: (*OCSProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnboardConsumer",
			Handler:    _OCSProvider_OnboardConsumer_Handler,
		},
		{
			MethodName: "GetStorageConfig",
			Handler:    _OCSProvider_GetStorageConfig_Handler,
		},
		{
			MethodName: "OffboardConsumer",
			Handler:    _OCSProvider_OffboardConsumer_Handler,
		},
		{
			MethodName: "UpdateCapacity",
			Handler:    _OCSProvider_UpdateCapacity_Handler,
		},
		{
			MethodName: "AcknowledgeOnboarding",
			Handler:    _OCSProvider_AcknowledgeOnboarding_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provider.proto",
}
