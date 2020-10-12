// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

import (
	context "context"
	model "github.com/binkkatal/proto_user/src/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ActivityServiceClient is the client API for ActivityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActivityServiceClient interface {
	Add(ctx context.Context, in *model.Activity, opts ...grpc.CallOption) (*ActivityResponse, error)
	Get(ctx context.Context, in *GetActivityRequest, opts ...grpc.CallOption) (*ActivityResponse, error)
	Start(ctx context.Context, in *GetActivityRequest, opts ...grpc.CallOption) (*ActivityResponse, error)
	Stop(ctx context.Context, in *GetActivityRequest, opts ...grpc.CallOption) (*ActivityResponse, error)
	IsValid(ctx context.Context, in *GetActivityRequest, opts ...grpc.CallOption) (*BoolResult, error)
	IsDone(ctx context.Context, in *GetActivityRequest, opts ...grpc.CallOption) (*BoolResult, error)
}

type activityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActivityServiceClient(cc grpc.ClientConnInterface) ActivityServiceClient {
	return &activityServiceClient{cc}
}

func (c *activityServiceClient) Add(ctx context.Context, in *model.Activity, opts ...grpc.CallOption) (*ActivityResponse, error) {
	out := new(ActivityResponse)
	err := c.cc.Invoke(ctx, "/service.ActivityService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityServiceClient) Get(ctx context.Context, in *GetActivityRequest, opts ...grpc.CallOption) (*ActivityResponse, error) {
	out := new(ActivityResponse)
	err := c.cc.Invoke(ctx, "/service.ActivityService/get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityServiceClient) Start(ctx context.Context, in *GetActivityRequest, opts ...grpc.CallOption) (*ActivityResponse, error) {
	out := new(ActivityResponse)
	err := c.cc.Invoke(ctx, "/service.ActivityService/start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityServiceClient) Stop(ctx context.Context, in *GetActivityRequest, opts ...grpc.CallOption) (*ActivityResponse, error) {
	out := new(ActivityResponse)
	err := c.cc.Invoke(ctx, "/service.ActivityService/stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityServiceClient) IsValid(ctx context.Context, in *GetActivityRequest, opts ...grpc.CallOption) (*BoolResult, error) {
	out := new(BoolResult)
	err := c.cc.Invoke(ctx, "/service.ActivityService/isValid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityServiceClient) IsDone(ctx context.Context, in *GetActivityRequest, opts ...grpc.CallOption) (*BoolResult, error) {
	out := new(BoolResult)
	err := c.cc.Invoke(ctx, "/service.ActivityService/isDone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivityServiceServer is the server API for ActivityService service.
// All implementations must embed UnimplementedActivityServiceServer
// for forward compatibility
type ActivityServiceServer interface {
	Add(context.Context, *model.Activity) (*ActivityResponse, error)
	Get(context.Context, *GetActivityRequest) (*ActivityResponse, error)
	Start(context.Context, *GetActivityRequest) (*ActivityResponse, error)
	Stop(context.Context, *GetActivityRequest) (*ActivityResponse, error)
	IsValid(context.Context, *GetActivityRequest) (*BoolResult, error)
	IsDone(context.Context, *GetActivityRequest) (*BoolResult, error)
	mustEmbedUnimplementedActivityServiceServer()
}

// UnimplementedActivityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedActivityServiceServer struct {
}

func (UnimplementedActivityServiceServer) Add(context.Context, *model.Activity) (*ActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedActivityServiceServer) Get(context.Context, *GetActivityRequest) (*ActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedActivityServiceServer) Start(context.Context, *GetActivityRequest) (*ActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedActivityServiceServer) Stop(context.Context, *GetActivityRequest) (*ActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedActivityServiceServer) IsValid(context.Context, *GetActivityRequest) (*BoolResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsValid not implemented")
}
func (UnimplementedActivityServiceServer) IsDone(context.Context, *GetActivityRequest) (*BoolResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsDone not implemented")
}
func (UnimplementedActivityServiceServer) mustEmbedUnimplementedActivityServiceServer() {}

// UnsafeActivityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActivityServiceServer will
// result in compilation errors.
type UnsafeActivityServiceServer interface {
	mustEmbedUnimplementedActivityServiceServer()
}

func RegisterActivityServiceServer(s *grpc.Server, srv ActivityServiceServer) {
	s.RegisterService(&_ActivityService_serviceDesc, srv)
}

func _ActivityService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.Activity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ActivityService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).Add(ctx, req.(*model.Activity))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ActivityService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).Get(ctx, req.(*GetActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ActivityService/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).Start(ctx, req.(*GetActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ActivityService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).Stop(ctx, req.(*GetActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityService_IsValid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).IsValid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ActivityService/IsValid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).IsValid(ctx, req.(*GetActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityService_IsDone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).IsDone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ActivityService/IsDone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).IsDone(ctx, req.(*GetActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ActivityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.ActivityService",
	HandlerType: (*ActivityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _ActivityService_Add_Handler,
		},
		{
			MethodName: "get",
			Handler:    _ActivityService_Get_Handler,
		},
		{
			MethodName: "start",
			Handler:    _ActivityService_Start_Handler,
		},
		{
			MethodName: "stop",
			Handler:    _ActivityService_Stop_Handler,
		},
		{
			MethodName: "isValid",
			Handler:    _ActivityService_IsValid_Handler,
		},
		{
			MethodName: "isDone",
			Handler:    _ActivityService_IsDone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/service/activity_service.proto",
}