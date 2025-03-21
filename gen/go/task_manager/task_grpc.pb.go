// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: task_manager/task.proto

package task_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Task_CreateTask_FullMethodName = "/task.Task/CreateTask"
	Task_ReadTask_FullMethodName   = "/task.Task/ReadTask"
	Task_UpdateTask_FullMethodName = "/task.Task/UpdateTask"
	Task_DeleteTask_FullMethodName = "/task.Task/DeleteTask"
)

// TaskClient is the client API for Task service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Сервис для управления задачами
type TaskClient interface {
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	ReadTask(ctx context.Context, in *ReadTaskRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type taskClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskClient(cc grpc.ClientConnInterface) TaskClient {
	return &taskClient{cc}
}

func (c *taskClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, Task_CreateTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) ReadTask(ctx context.Context, in *ReadTaskRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, Task_ReadTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, Task_UpdateTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Task_DeleteTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServer is the server API for Task service.
// All implementations must embed UnimplementedTaskServer
// for forward compatibility.
//
// Сервис для управления задачами
type TaskServer interface {
	CreateTask(context.Context, *CreateTaskRequest) (*TaskResponse, error)
	ReadTask(context.Context, *ReadTaskRequest) (*TaskResponse, error)
	UpdateTask(context.Context, *UpdateTaskRequest) (*TaskResponse, error)
	DeleteTask(context.Context, *DeleteTaskRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedTaskServer()
}

// UnimplementedTaskServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTaskServer struct{}

func (UnimplementedTaskServer) CreateTask(context.Context, *CreateTaskRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedTaskServer) ReadTask(context.Context, *ReadTaskRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadTask not implemented")
}
func (UnimplementedTaskServer) UpdateTask(context.Context, *UpdateTaskRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
func (UnimplementedTaskServer) DeleteTask(context.Context, *DeleteTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (UnimplementedTaskServer) mustEmbedUnimplementedTaskServer() {}
func (UnimplementedTaskServer) testEmbeddedByValue()              {}

// UnsafeTaskServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServer will
// result in compilation errors.
type UnsafeTaskServer interface {
	mustEmbedUnimplementedTaskServer()
}

func RegisterTaskServer(s grpc.ServiceRegistrar, srv TaskServer) {
	// If the following call pancis, it indicates UnimplementedTaskServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Task_ServiceDesc, srv)
}

func _Task_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_CreateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_ReadTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).ReadTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_ReadTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).ReadTask(ctx, req.(*ReadTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_UpdateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).UpdateTask(ctx, req.(*UpdateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_DeleteTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).DeleteTask(ctx, req.(*DeleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Task_ServiceDesc is the grpc.ServiceDesc for Task service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Task_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "task.Task",
	HandlerType: (*TaskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _Task_CreateTask_Handler,
		},
		{
			MethodName: "ReadTask",
			Handler:    _Task_ReadTask_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _Task_UpdateTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _Task_DeleteTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task_manager/task.proto",
}

const (
	TaskCategory_CreateTaskCategory_FullMethodName = "/task.TaskCategory/CreateTaskCategory"
	TaskCategory_ReadTaskCategory_FullMethodName   = "/task.TaskCategory/ReadTaskCategory"
	TaskCategory_UpdateTaskCategory_FullMethodName = "/task.TaskCategory/UpdateTaskCategory"
	TaskCategory_DeleteTaskCategory_FullMethodName = "/task.TaskCategory/DeleteTaskCategory"
)

// TaskCategoryClient is the client API for TaskCategory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Сервис для управления категориями задач
type TaskCategoryClient interface {
	CreateTaskCategory(ctx context.Context, in *CreateTaskCategoryRequest, opts ...grpc.CallOption) (*CreateTaskCategoryResponse, error)
	ReadTaskCategory(ctx context.Context, in *ReadTaskCategoryRequest, opts ...grpc.CallOption) (*TaskCategoryResponse, error)
	UpdateTaskCategory(ctx context.Context, in *UpdateTaskCategoryRequest, opts ...grpc.CallOption) (*TaskCategoryResponse, error)
	DeleteTaskCategory(ctx context.Context, in *DeleteTaskCategoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type taskCategoryClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskCategoryClient(cc grpc.ClientConnInterface) TaskCategoryClient {
	return &taskCategoryClient{cc}
}

func (c *taskCategoryClient) CreateTaskCategory(ctx context.Context, in *CreateTaskCategoryRequest, opts ...grpc.CallOption) (*CreateTaskCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTaskCategoryResponse)
	err := c.cc.Invoke(ctx, TaskCategory_CreateTaskCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskCategoryClient) ReadTaskCategory(ctx context.Context, in *ReadTaskCategoryRequest, opts ...grpc.CallOption) (*TaskCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaskCategoryResponse)
	err := c.cc.Invoke(ctx, TaskCategory_ReadTaskCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskCategoryClient) UpdateTaskCategory(ctx context.Context, in *UpdateTaskCategoryRequest, opts ...grpc.CallOption) (*TaskCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaskCategoryResponse)
	err := c.cc.Invoke(ctx, TaskCategory_UpdateTaskCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskCategoryClient) DeleteTaskCategory(ctx context.Context, in *DeleteTaskCategoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TaskCategory_DeleteTaskCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskCategoryServer is the server API for TaskCategory service.
// All implementations must embed UnimplementedTaskCategoryServer
// for forward compatibility.
//
// Сервис для управления категориями задач
type TaskCategoryServer interface {
	CreateTaskCategory(context.Context, *CreateTaskCategoryRequest) (*CreateTaskCategoryResponse, error)
	ReadTaskCategory(context.Context, *ReadTaskCategoryRequest) (*TaskCategoryResponse, error)
	UpdateTaskCategory(context.Context, *UpdateTaskCategoryRequest) (*TaskCategoryResponse, error)
	DeleteTaskCategory(context.Context, *DeleteTaskCategoryRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedTaskCategoryServer()
}

// UnimplementedTaskCategoryServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTaskCategoryServer struct{}

func (UnimplementedTaskCategoryServer) CreateTaskCategory(context.Context, *CreateTaskCategoryRequest) (*CreateTaskCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTaskCategory not implemented")
}
func (UnimplementedTaskCategoryServer) ReadTaskCategory(context.Context, *ReadTaskCategoryRequest) (*TaskCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadTaskCategory not implemented")
}
func (UnimplementedTaskCategoryServer) UpdateTaskCategory(context.Context, *UpdateTaskCategoryRequest) (*TaskCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTaskCategory not implemented")
}
func (UnimplementedTaskCategoryServer) DeleteTaskCategory(context.Context, *DeleteTaskCategoryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTaskCategory not implemented")
}
func (UnimplementedTaskCategoryServer) mustEmbedUnimplementedTaskCategoryServer() {}
func (UnimplementedTaskCategoryServer) testEmbeddedByValue()                      {}

// UnsafeTaskCategoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskCategoryServer will
// result in compilation errors.
type UnsafeTaskCategoryServer interface {
	mustEmbedUnimplementedTaskCategoryServer()
}

func RegisterTaskCategoryServer(s grpc.ServiceRegistrar, srv TaskCategoryServer) {
	// If the following call pancis, it indicates UnimplementedTaskCategoryServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TaskCategory_ServiceDesc, srv)
}

func _TaskCategory_CreateTaskCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskCategoryServer).CreateTaskCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskCategory_CreateTaskCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskCategoryServer).CreateTaskCategory(ctx, req.(*CreateTaskCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskCategory_ReadTaskCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadTaskCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskCategoryServer).ReadTaskCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskCategory_ReadTaskCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskCategoryServer).ReadTaskCategory(ctx, req.(*ReadTaskCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskCategory_UpdateTaskCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskCategoryServer).UpdateTaskCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskCategory_UpdateTaskCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskCategoryServer).UpdateTaskCategory(ctx, req.(*UpdateTaskCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskCategory_DeleteTaskCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTaskCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskCategoryServer).DeleteTaskCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskCategory_DeleteTaskCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskCategoryServer).DeleteTaskCategory(ctx, req.(*DeleteTaskCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskCategory_ServiceDesc is the grpc.ServiceDesc for TaskCategory service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskCategory_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "task.TaskCategory",
	HandlerType: (*TaskCategoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTaskCategory",
			Handler:    _TaskCategory_CreateTaskCategory_Handler,
		},
		{
			MethodName: "ReadTaskCategory",
			Handler:    _TaskCategory_ReadTaskCategory_Handler,
		},
		{
			MethodName: "UpdateTaskCategory",
			Handler:    _TaskCategory_UpdateTaskCategory_Handler,
		},
		{
			MethodName: "DeleteTaskCategory",
			Handler:    _TaskCategory_DeleteTaskCategory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task_manager/task.proto",
}
