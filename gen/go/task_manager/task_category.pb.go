// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: task_manager/task_category.proto

package task_category_v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Запрос на создание категории задач
type CreateTaskCategoryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"` // Заголовок от 1 до 100 символов
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTaskCategoryRequest) Reset() {
	*x = CreateTaskCategoryRequest{}
	mi := &file_task_manager_task_category_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTaskCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskCategoryRequest) ProtoMessage() {}

func (x *CreateTaskCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_manager_task_category_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskCategoryRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskCategoryRequest) Descriptor() ([]byte, []int) {
	return file_task_manager_task_category_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTaskCategoryRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

// Ответ на создание категории задач
type CreateTaskCategoryResponse struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	TaskCategoryId int64                  `protobuf:"varint,1,opt,name=task_category_id,json=taskCategoryId,proto3" json:"task_category_id,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *CreateTaskCategoryResponse) Reset() {
	*x = CreateTaskCategoryResponse{}
	mi := &file_task_manager_task_category_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTaskCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskCategoryResponse) ProtoMessage() {}

func (x *CreateTaskCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_manager_task_category_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskCategoryResponse.ProtoReflect.Descriptor instead.
func (*CreateTaskCategoryResponse) Descriptor() ([]byte, []int) {
	return file_task_manager_task_category_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTaskCategoryResponse) GetTaskCategoryId() int64 {
	if x != nil {
		return x.TaskCategoryId
	}
	return 0
}

// Запрос на чтение категории задач
type ReadTaskCategoryRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	TaskCategoryId int64                  `protobuf:"varint,1,opt,name=task_category_id,json=taskCategoryId,proto3" json:"task_category_id,omitempty"` // ID категории должен быть положительным
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ReadTaskCategoryRequest) Reset() {
	*x = ReadTaskCategoryRequest{}
	mi := &file_task_manager_task_category_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadTaskCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadTaskCategoryRequest) ProtoMessage() {}

func (x *ReadTaskCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_manager_task_category_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadTaskCategoryRequest.ProtoReflect.Descriptor instead.
func (*ReadTaskCategoryRequest) Descriptor() ([]byte, []int) {
	return file_task_manager_task_category_proto_rawDescGZIP(), []int{2}
}

func (x *ReadTaskCategoryRequest) GetTaskCategoryId() int64 {
	if x != nil {
		return x.TaskCategoryId
	}
	return 0
}

// Ответ с данными категории задач
type TaskCategoryResponse struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	TaskCategoryId int64                  `protobuf:"varint,1,opt,name=task_category_id,json=taskCategoryId,proto3" json:"task_category_id,omitempty"`
	Title          string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *TaskCategoryResponse) Reset() {
	*x = TaskCategoryResponse{}
	mi := &file_task_manager_task_category_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskCategoryResponse) ProtoMessage() {}

func (x *TaskCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_manager_task_category_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskCategoryResponse.ProtoReflect.Descriptor instead.
func (*TaskCategoryResponse) Descriptor() ([]byte, []int) {
	return file_task_manager_task_category_proto_rawDescGZIP(), []int{3}
}

func (x *TaskCategoryResponse) GetTaskCategoryId() int64 {
	if x != nil {
		return x.TaskCategoryId
	}
	return 0
}

func (x *TaskCategoryResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

// Запрос на обновление категории задач
type UpdateTaskCategoryRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	TaskCategoryId int64                  `protobuf:"varint,1,opt,name=task_category_id,json=taskCategoryId,proto3" json:"task_category_id,omitempty"` // ID категории должен быть положительным
	Title          string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`                                            // Заголовок от 1 до 100 символов
	UpdateMask     *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *UpdateTaskCategoryRequest) Reset() {
	*x = UpdateTaskCategoryRequest{}
	mi := &file_task_manager_task_category_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTaskCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskCategoryRequest) ProtoMessage() {}

func (x *UpdateTaskCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_manager_task_category_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskCategoryRequest.ProtoReflect.Descriptor instead.
func (*UpdateTaskCategoryRequest) Descriptor() ([]byte, []int) {
	return file_task_manager_task_category_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTaskCategoryRequest) GetTaskCategoryId() int64 {
	if x != nil {
		return x.TaskCategoryId
	}
	return 0
}

func (x *UpdateTaskCategoryRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateTaskCategoryRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// Запрос на удаление категории задач
type DeleteTaskCategoryRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	TaskCategoryId int64                  `protobuf:"varint,1,opt,name=task_category_id,json=taskCategoryId,proto3" json:"task_category_id,omitempty"` // ID категории должен быть положительным
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *DeleteTaskCategoryRequest) Reset() {
	*x = DeleteTaskCategoryRequest{}
	mi := &file_task_manager_task_category_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTaskCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskCategoryRequest) ProtoMessage() {}

func (x *DeleteTaskCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_manager_task_category_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskCategoryRequest.ProtoReflect.Descriptor instead.
func (*DeleteTaskCategoryRequest) Descriptor() ([]byte, []int) {
	return file_task_manager_task_category_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTaskCategoryRequest) GetTaskCategoryId() int64 {
	if x != nil {
		return x.TaskCategoryId
	}
	return 0
}

var File_task_manager_task_category_proto protoreflect.FileDescriptor

var file_task_manager_task_category_proto_rawDesc = string([]byte{
	0x0a, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x19, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x64,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x46, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0e, 0x74, 0x61, 0x73, 0x6b, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22,
	0x4c, 0x0a, 0x17, 0x52, 0x65, 0x61, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x10, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x0e, 0x74,
	0x61, 0x73, 0x6b, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0x56, 0x0a,
	0x14, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0e, 0x74, 0x61, 0x73, 0x6b, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0xac, 0x01, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x10, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x0e, 0x74, 0x61, 0x73, 0x6b, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x64,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x61, 0x73, 0x6b, 0x22, 0x4e, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x31, 0x0a, 0x10, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x22, 0x02, 0x20, 0x00, 0x52, 0x0e, 0x74, 0x61, 0x73, 0x6b, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x49, 0x64, 0x32, 0x97, 0x03, 0x0a, 0x0c, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x69, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x28, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x5f, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x26, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x63, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x28, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x28, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x38,
	0x5a, 0x36, 0x64, 0x6f, 0x76, 0x67, 0x61, 0x6c, 0x62, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x3b, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_task_manager_task_category_proto_rawDescOnce sync.Once
	file_task_manager_task_category_proto_rawDescData []byte
)

func file_task_manager_task_category_proto_rawDescGZIP() []byte {
	file_task_manager_task_category_proto_rawDescOnce.Do(func() {
		file_task_manager_task_category_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_task_manager_task_category_proto_rawDesc), len(file_task_manager_task_category_proto_rawDesc)))
	})
	return file_task_manager_task_category_proto_rawDescData
}

var file_task_manager_task_category_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_task_manager_task_category_proto_goTypes = []any{
	(*CreateTaskCategoryRequest)(nil),  // 0: task_category.CreateTaskCategoryRequest
	(*CreateTaskCategoryResponse)(nil), // 1: task_category.CreateTaskCategoryResponse
	(*ReadTaskCategoryRequest)(nil),    // 2: task_category.ReadTaskCategoryRequest
	(*TaskCategoryResponse)(nil),       // 3: task_category.TaskCategoryResponse
	(*UpdateTaskCategoryRequest)(nil),  // 4: task_category.UpdateTaskCategoryRequest
	(*DeleteTaskCategoryRequest)(nil),  // 5: task_category.DeleteTaskCategoryRequest
	(*fieldmaskpb.FieldMask)(nil),      // 6: google.protobuf.FieldMask
	(*emptypb.Empty)(nil),              // 7: google.protobuf.Empty
}
var file_task_manager_task_category_proto_depIdxs = []int32{
	6, // 0: task_category.UpdateTaskCategoryRequest.update_mask:type_name -> google.protobuf.FieldMask
	0, // 1: task_category.TaskCategory.CreateTaskCategory:input_type -> task_category.CreateTaskCategoryRequest
	2, // 2: task_category.TaskCategory.ReadTaskCategory:input_type -> task_category.ReadTaskCategoryRequest
	4, // 3: task_category.TaskCategory.UpdateTaskCategory:input_type -> task_category.UpdateTaskCategoryRequest
	5, // 4: task_category.TaskCategory.DeleteTaskCategory:input_type -> task_category.DeleteTaskCategoryRequest
	1, // 5: task_category.TaskCategory.CreateTaskCategory:output_type -> task_category.CreateTaskCategoryResponse
	3, // 6: task_category.TaskCategory.ReadTaskCategory:output_type -> task_category.TaskCategoryResponse
	3, // 7: task_category.TaskCategory.UpdateTaskCategory:output_type -> task_category.TaskCategoryResponse
	7, // 8: task_category.TaskCategory.DeleteTaskCategory:output_type -> google.protobuf.Empty
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_task_manager_task_category_proto_init() }
func file_task_manager_task_category_proto_init() {
	if File_task_manager_task_category_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_task_manager_task_category_proto_rawDesc), len(file_task_manager_task_category_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_task_manager_task_category_proto_goTypes,
		DependencyIndexes: file_task_manager_task_category_proto_depIdxs,
		MessageInfos:      file_task_manager_task_category_proto_msgTypes,
	}.Build()
	File_task_manager_task_category_proto = out.File
	file_task_manager_task_category_proto_goTypes = nil
	file_task_manager_task_category_proto_depIdxs = nil
}
