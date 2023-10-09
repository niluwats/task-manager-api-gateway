// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: task.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_TODO       Status = 0
	Status_INPROGRESS Status = 1
	Status_READY      Status = 2
	Status_COMPLETED  Status = 3
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "TODO",
		1: "INPROGRESS",
		2: "READY",
		3: "COMPLETED",
	}
	Status_value = map[string]int32{
		"TODO":       0,
		"INPROGRESS": 1,
		"READY":      2,
		"COMPLETED":  3,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_task_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_task_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{0}
}

type CommonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,10,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CommonResponse) Reset() {
	*x = CommonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResponse) ProtoMessage() {}

func (x *CommonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResponse.ProtoReflect.Descriptor instead.
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{0}
}

func (x *CommonResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CommonResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type NewProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Creator     int32   `protobuf:"varint,3,opt,name=creator,proto3" json:"creator,omitempty"`
	Assignees   []int32 `protobuf:"varint,4,rep,packed,name=assignees,proto3" json:"assignees,omitempty"`
}

func (x *NewProjectRequest) Reset() {
	*x = NewProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewProjectRequest) ProtoMessage() {}

func (x *NewProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewProjectRequest.ProtoReflect.Descriptor instead.
func (*NewProjectRequest) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{1}
}

func (x *NewProjectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewProjectRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NewProjectRequest) GetCreator() int32 {
	if x != nil {
		return x.Creator
	}
	return 0
}

func (x *NewProjectRequest) GetAssignees() []int32 {
	if x != nil {
		return x.Assignees
	}
	return nil
}

type UpdaterojectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Creator     int32   `protobuf:"varint,4,opt,name=creator,proto3" json:"creator,omitempty"`
	Assignees   []int32 `protobuf:"varint,5,rep,packed,name=assignees,proto3" json:"assignees,omitempty"`
}

func (x *UpdaterojectRequest) Reset() {
	*x = UpdaterojectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdaterojectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdaterojectRequest) ProtoMessage() {}

func (x *UpdaterojectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdaterojectRequest.ProtoReflect.Descriptor instead.
func (*UpdaterojectRequest) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{2}
}

func (x *UpdaterojectRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdaterojectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdaterojectRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdaterojectRequest) GetCreator() int32 {
	if x != nil {
		return x.Creator
	}
	return 0
}

func (x *UpdaterojectRequest) GetAssignees() []int32 {
	if x != nil {
		return x.Assignees
	}
	return nil
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Creator     int32  `protobuf:"varint,3,opt,name=creator,proto3" json:"creator,omitempty"`
	Assignee    int32  `protobuf:"varint,4,opt,name=assignee,proto3" json:"assignee,omitempty"`
	TaskStatus  Status `protobuf:"varint,5,opt,name=task_status,json=taskStatus,proto3,enum=auth.Status" json:"task_status,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{3}
}

func (x *Task) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Task) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Task) GetCreator() int32 {
	if x != nil {
		return x.Creator
	}
	return 0
}

func (x *Task) GetAssignee() int32 {
	if x != nil {
		return x.Assignee
	}
	return 0
}

func (x *Task) GetTaskStatus() Status {
	if x != nil {
		return x.TaskStatus
	}
	return Status_TODO
}

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Creator     int32   `protobuf:"varint,4,opt,name=creator,proto3" json:"creator,omitempty"`
	CreatedOn   string  `protobuf:"bytes,5,opt,name=created_on,json=createdOn,proto3" json:"created_on,omitempty"`
	UpdatedOn   string  `protobuf:"bytes,6,opt,name=updated_on,json=updatedOn,proto3" json:"updated_on,omitempty"`
	Assignees   []int32 `protobuf:"varint,7,rep,packed,name=assignees,proto3" json:"assignees,omitempty"`
	Tasks       []*Task `protobuf:"bytes,8,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{4}
}

func (x *Project) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Project) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Project) GetCreator() int32 {
	if x != nil {
		return x.Creator
	}
	return 0
}

func (x *Project) GetCreatedOn() string {
	if x != nil {
		return x.CreatedOn
	}
	return ""
}

func (x *Project) GetUpdatedOn() string {
	if x != nil {
		return x.UpdatedOn
	}
	return ""
}

func (x *Project) GetAssignees() []int32 {
	if x != nil {
		return x.Assignees
	}
	return nil
}

func (x *Project) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type ProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project        *Project        `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	CommonResponse *CommonResponse `protobuf:"bytes,2,opt,name=common_response,json=commonResponse,proto3" json:"common_response,omitempty"`
}

func (x *ProjectResponse) Reset() {
	*x = ProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectResponse) ProtoMessage() {}

func (x *ProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectResponse.ProtoReflect.Descriptor instead.
func (*ProjectResponse) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{5}
}

func (x *ProjectResponse) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

func (x *ProjectResponse) GetCommonResponse() *CommonResponse {
	if x != nil {
		return x.CommonResponse
	}
	return nil
}

type ProjectsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project        []*Project      `protobuf:"bytes,1,rep,name=project,proto3" json:"project,omitempty"`
	CommonResponse *CommonResponse `protobuf:"bytes,2,opt,name=common_response,json=commonResponse,proto3" json:"common_response,omitempty"`
}

func (x *ProjectsResponse) Reset() {
	*x = ProjectsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectsResponse) ProtoMessage() {}

func (x *ProjectsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectsResponse.ProtoReflect.Descriptor instead.
func (*ProjectsResponse) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{6}
}

func (x *ProjectsResponse) GetProject() []*Project {
	if x != nil {
		return x.Project
	}
	return nil
}

func (x *ProjectsResponse) GetCommonResponse() *CommonResponse {
	if x != nil {
		return x.CommonResponse
	}
	return nil
}

type ViewProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ViewProjectRequest) Reset() {
	*x = ViewProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewProjectRequest) ProtoMessage() {}

func (x *ViewProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewProjectRequest.ProtoReflect.Descriptor instead.
func (*ViewProjectRequest) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{7}
}

func (x *ViewProjectRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RemovedProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonResponse *CommonResponse `protobuf:"bytes,1,opt,name=common_response,json=commonResponse,proto3" json:"common_response,omitempty"`
}

func (x *RemovedProjectResponse) Reset() {
	*x = RemovedProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovedProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovedProjectResponse) ProtoMessage() {}

func (x *RemovedProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovedProjectResponse.ProtoReflect.Descriptor instead.
func (*RemovedProjectResponse) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{8}
}

func (x *RemovedProjectResponse) GetCommonResponse() *CommonResponse {
	if x != nil {
		return x.CommonResponse
	}
	return nil
}

type EmptyParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyParams) Reset() {
	*x = EmptyParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyParams) ProtoMessage() {}

func (x *EmptyParams) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyParams.ProtoReflect.Descriptor instead.
func (*EmptyParams) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{9}
}

var File_task_proto protoreflect.FileDescriptor

var file_task_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75,
	0x74, 0x68, 0x22, 0x42, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x11, 0x4e, 0x65, 0x77, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x09, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x73, 0x22, 0x93, 0x01, 0x0a, 0x13, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x73,
	0x22, 0x9d, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x65, 0x12, 0x2d, 0x0a, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0xe7, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x61,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x79, 0x0a, 0x0f, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x3d, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7a, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x3d, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x24, 0x0a, 0x12, 0x56, 0x69, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x57, 0x0a, 0x16, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x64, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3d, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x0d, 0x0a, 0x0b, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2a,
	0x3c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x4f, 0x44,
	0x4f, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4e, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53,
	0x53, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x02, 0x12, 0x0d,
	0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x03, 0x32, 0xd1, 0x02,
	0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x17,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4e, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x56, 0x69, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3c, 0x0a, 0x0f, 0x56, 0x69, 0x65, 0x77, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x12, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x40, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x56, 0x69, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x64, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_task_proto_rawDescOnce sync.Once
	file_task_proto_rawDescData = file_task_proto_rawDesc
)

func file_task_proto_rawDescGZIP() []byte {
	file_task_proto_rawDescOnce.Do(func() {
		file_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_task_proto_rawDescData)
	})
	return file_task_proto_rawDescData
}

var file_task_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_task_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_task_proto_goTypes = []interface{}{
	(Status)(0),                    // 0: auth.Status
	(*CommonResponse)(nil),         // 1: auth.CommonResponse
	(*NewProjectRequest)(nil),      // 2: auth.NewProjectRequest
	(*UpdaterojectRequest)(nil),    // 3: auth.UpdaterojectRequest
	(*Task)(nil),                   // 4: auth.Task
	(*Project)(nil),                // 5: auth.Project
	(*ProjectResponse)(nil),        // 6: auth.ProjectResponse
	(*ProjectsResponse)(nil),       // 7: auth.ProjectsResponse
	(*ViewProjectRequest)(nil),     // 8: auth.ViewProjectRequest
	(*RemovedProjectResponse)(nil), // 9: auth.RemovedProjectResponse
	(*EmptyParams)(nil),            // 10: auth.EmptyParams
}
var file_task_proto_depIdxs = []int32{
	0,  // 0: auth.Task.task_status:type_name -> auth.Status
	4,  // 1: auth.Project.tasks:type_name -> auth.Task
	5,  // 2: auth.ProjectResponse.project:type_name -> auth.Project
	1,  // 3: auth.ProjectResponse.common_response:type_name -> auth.CommonResponse
	5,  // 4: auth.ProjectsResponse.project:type_name -> auth.Project
	1,  // 5: auth.ProjectsResponse.common_response:type_name -> auth.CommonResponse
	1,  // 6: auth.RemovedProjectResponse.common_response:type_name -> auth.CommonResponse
	2,  // 7: auth.TaskService.CreateProject:input_type -> auth.NewProjectRequest
	3,  // 8: auth.TaskService.UpdateProject:input_type -> auth.UpdaterojectRequest
	8,  // 9: auth.TaskService.ViewProject:input_type -> auth.ViewProjectRequest
	10, // 10: auth.TaskService.ViewAllProjects:input_type -> auth.EmptyParams
	8,  // 11: auth.TaskService.Remove:input_type -> auth.ViewProjectRequest
	6,  // 12: auth.TaskService.CreateProject:output_type -> auth.ProjectResponse
	6,  // 13: auth.TaskService.UpdateProject:output_type -> auth.ProjectResponse
	6,  // 14: auth.TaskService.ViewProject:output_type -> auth.ProjectResponse
	7,  // 15: auth.TaskService.ViewAllProjects:output_type -> auth.ProjectsResponse
	9,  // 16: auth.TaskService.Remove:output_type -> auth.RemovedProjectResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_task_proto_init() }
func file_task_proto_init() {
	if File_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewProjectRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdaterojectRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_task_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_task_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_task_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_task_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewProjectRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_task_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovedProjectResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_task_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_task_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_task_proto_goTypes,
		DependencyIndexes: file_task_proto_depIdxs,
		EnumInfos:         file_task_proto_enumTypes,
		MessageInfos:      file_task_proto_msgTypes,
	}.Build()
	File_task_proto = out.File
	file_task_proto_rawDesc = nil
	file_task_proto_goTypes = nil
	file_task_proto_depIdxs = nil
}