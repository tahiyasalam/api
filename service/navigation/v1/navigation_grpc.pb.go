// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: service/navigation/v1/navigation.proto

package v1

import (
	context "context"
	v1 "go.viam.com/api/common/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NavigationServiceClient is the client API for NavigationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NavigationServiceClient interface {
	GetMode(ctx context.Context, in *GetModeRequest, opts ...grpc.CallOption) (*GetModeResponse, error)
	SetMode(ctx context.Context, in *SetModeRequest, opts ...grpc.CallOption) (*SetModeResponse, error)
	GetLocation(ctx context.Context, in *GetLocationRequest, opts ...grpc.CallOption) (*GetLocationResponse, error)
	GetWaypoints(ctx context.Context, in *GetWaypointsRequest, opts ...grpc.CallOption) (*GetWaypointsResponse, error)
	AddWaypoint(ctx context.Context, in *AddWaypointRequest, opts ...grpc.CallOption) (*AddWaypointResponse, error)
	RemoveWaypoint(ctx context.Context, in *RemoveWaypointRequest, opts ...grpc.CallOption) (*RemoveWaypointResponse, error)
	GetObstacles(ctx context.Context, in *GetObstaclesRequest, opts ...grpc.CallOption) (*GetObstaclesResponse, error)
	// DoCommand sends/receives arbitrary commands
	DoCommand(ctx context.Context, in *v1.DoCommandRequest, opts ...grpc.CallOption) (*v1.DoCommandResponse, error)
}

type navigationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNavigationServiceClient(cc grpc.ClientConnInterface) NavigationServiceClient {
	return &navigationServiceClient{cc}
}

func (c *navigationServiceClient) GetMode(ctx context.Context, in *GetModeRequest, opts ...grpc.CallOption) (*GetModeResponse, error) {
	out := new(GetModeResponse)
	err := c.cc.Invoke(ctx, "/viam.service.navigation.v1.NavigationService/GetMode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *navigationServiceClient) SetMode(ctx context.Context, in *SetModeRequest, opts ...grpc.CallOption) (*SetModeResponse, error) {
	out := new(SetModeResponse)
	err := c.cc.Invoke(ctx, "/viam.service.navigation.v1.NavigationService/SetMode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *navigationServiceClient) GetLocation(ctx context.Context, in *GetLocationRequest, opts ...grpc.CallOption) (*GetLocationResponse, error) {
	out := new(GetLocationResponse)
	err := c.cc.Invoke(ctx, "/viam.service.navigation.v1.NavigationService/GetLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *navigationServiceClient) GetWaypoints(ctx context.Context, in *GetWaypointsRequest, opts ...grpc.CallOption) (*GetWaypointsResponse, error) {
	out := new(GetWaypointsResponse)
	err := c.cc.Invoke(ctx, "/viam.service.navigation.v1.NavigationService/GetWaypoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *navigationServiceClient) AddWaypoint(ctx context.Context, in *AddWaypointRequest, opts ...grpc.CallOption) (*AddWaypointResponse, error) {
	out := new(AddWaypointResponse)
	err := c.cc.Invoke(ctx, "/viam.service.navigation.v1.NavigationService/AddWaypoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *navigationServiceClient) RemoveWaypoint(ctx context.Context, in *RemoveWaypointRequest, opts ...grpc.CallOption) (*RemoveWaypointResponse, error) {
	out := new(RemoveWaypointResponse)
	err := c.cc.Invoke(ctx, "/viam.service.navigation.v1.NavigationService/RemoveWaypoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *navigationServiceClient) GetObstacles(ctx context.Context, in *GetObstaclesRequest, opts ...grpc.CallOption) (*GetObstaclesResponse, error) {
	out := new(GetObstaclesResponse)
	err := c.cc.Invoke(ctx, "/viam.service.navigation.v1.NavigationService/GetObstacles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *navigationServiceClient) DoCommand(ctx context.Context, in *v1.DoCommandRequest, opts ...grpc.CallOption) (*v1.DoCommandResponse, error) {
	out := new(v1.DoCommandResponse)
	err := c.cc.Invoke(ctx, "/viam.service.navigation.v1.NavigationService/DoCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NavigationServiceServer is the server API for NavigationService service.
// All implementations must embed UnimplementedNavigationServiceServer
// for forward compatibility
type NavigationServiceServer interface {
	GetMode(context.Context, *GetModeRequest) (*GetModeResponse, error)
	SetMode(context.Context, *SetModeRequest) (*SetModeResponse, error)
	GetLocation(context.Context, *GetLocationRequest) (*GetLocationResponse, error)
	GetWaypoints(context.Context, *GetWaypointsRequest) (*GetWaypointsResponse, error)
	AddWaypoint(context.Context, *AddWaypointRequest) (*AddWaypointResponse, error)
	RemoveWaypoint(context.Context, *RemoveWaypointRequest) (*RemoveWaypointResponse, error)
	GetObstacles(context.Context, *GetObstaclesRequest) (*GetObstaclesResponse, error)
	// DoCommand sends/receives arbitrary commands
	DoCommand(context.Context, *v1.DoCommandRequest) (*v1.DoCommandResponse, error)
	mustEmbedUnimplementedNavigationServiceServer()
}

// UnimplementedNavigationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNavigationServiceServer struct {
}

func (UnimplementedNavigationServiceServer) GetMode(context.Context, *GetModeRequest) (*GetModeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMode not implemented")
}
func (UnimplementedNavigationServiceServer) SetMode(context.Context, *SetModeRequest) (*SetModeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMode not implemented")
}
func (UnimplementedNavigationServiceServer) GetLocation(context.Context, *GetLocationRequest) (*GetLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocation not implemented")
}
func (UnimplementedNavigationServiceServer) GetWaypoints(context.Context, *GetWaypointsRequest) (*GetWaypointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWaypoints not implemented")
}
func (UnimplementedNavigationServiceServer) AddWaypoint(context.Context, *AddWaypointRequest) (*AddWaypointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddWaypoint not implemented")
}
func (UnimplementedNavigationServiceServer) RemoveWaypoint(context.Context, *RemoveWaypointRequest) (*RemoveWaypointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveWaypoint not implemented")
}
func (UnimplementedNavigationServiceServer) GetObstacles(context.Context, *GetObstaclesRequest) (*GetObstaclesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObstacles not implemented")
}
func (UnimplementedNavigationServiceServer) DoCommand(context.Context, *v1.DoCommandRequest) (*v1.DoCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoCommand not implemented")
}
func (UnimplementedNavigationServiceServer) mustEmbedUnimplementedNavigationServiceServer() {}

// UnsafeNavigationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NavigationServiceServer will
// result in compilation errors.
type UnsafeNavigationServiceServer interface {
	mustEmbedUnimplementedNavigationServiceServer()
}

func RegisterNavigationServiceServer(s grpc.ServiceRegistrar, srv NavigationServiceServer) {
	s.RegisterService(&NavigationService_ServiceDesc, srv)
}

func _NavigationService_GetMode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NavigationServiceServer).GetMode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.service.navigation.v1.NavigationService/GetMode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NavigationServiceServer).GetMode(ctx, req.(*GetModeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NavigationService_SetMode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetModeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NavigationServiceServer).SetMode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.service.navigation.v1.NavigationService/SetMode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NavigationServiceServer).SetMode(ctx, req.(*SetModeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NavigationService_GetLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NavigationServiceServer).GetLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.service.navigation.v1.NavigationService/GetLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NavigationServiceServer).GetLocation(ctx, req.(*GetLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NavigationService_GetWaypoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWaypointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NavigationServiceServer).GetWaypoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.service.navigation.v1.NavigationService/GetWaypoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NavigationServiceServer).GetWaypoints(ctx, req.(*GetWaypointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NavigationService_AddWaypoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddWaypointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NavigationServiceServer).AddWaypoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.service.navigation.v1.NavigationService/AddWaypoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NavigationServiceServer).AddWaypoint(ctx, req.(*AddWaypointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NavigationService_RemoveWaypoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveWaypointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NavigationServiceServer).RemoveWaypoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.service.navigation.v1.NavigationService/RemoveWaypoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NavigationServiceServer).RemoveWaypoint(ctx, req.(*RemoveWaypointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NavigationService_GetObstacles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObstaclesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NavigationServiceServer).GetObstacles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.service.navigation.v1.NavigationService/GetObstacles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NavigationServiceServer).GetObstacles(ctx, req.(*GetObstaclesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NavigationService_DoCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.DoCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NavigationServiceServer).DoCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.service.navigation.v1.NavigationService/DoCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NavigationServiceServer).DoCommand(ctx, req.(*v1.DoCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NavigationService_ServiceDesc is the grpc.ServiceDesc for NavigationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NavigationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "viam.service.navigation.v1.NavigationService",
	HandlerType: (*NavigationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMode",
			Handler:    _NavigationService_GetMode_Handler,
		},
		{
			MethodName: "SetMode",
			Handler:    _NavigationService_SetMode_Handler,
		},
		{
			MethodName: "GetLocation",
			Handler:    _NavigationService_GetLocation_Handler,
		},
		{
			MethodName: "GetWaypoints",
			Handler:    _NavigationService_GetWaypoints_Handler,
		},
		{
			MethodName: "AddWaypoint",
			Handler:    _NavigationService_AddWaypoint_Handler,
		},
		{
			MethodName: "RemoveWaypoint",
			Handler:    _NavigationService_RemoveWaypoint_Handler,
		},
		{
			MethodName: "GetObstacles",
			Handler:    _NavigationService_GetObstacles_Handler,
		},
		{
			MethodName: "DoCommand",
			Handler:    _NavigationService_DoCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/navigation/v1/navigation.proto",
}
