// Code generated by mojo. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: mojo/rpc/longrunning/operations.proto

package longrunning

import (
	_ "github.com/mojo-lang/core/go/pkg/mojo"
	core "github.com/mojo-lang/core/go/pkg/mojo/core"
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

type ListOperationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parent    string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	Filter    string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	PageSize  int32  `protobuf:"varint,2000,opt,name=page_size,json=pageSize,proto3" json:"pageSize,omitempty"`
	PageToken string `protobuf:"bytes,2001,opt,name=page_token,json=pageToken,proto3" json:"pageToken,omitempty"`
	Skip      int32  `protobuf:"varint,2002,opt,name=skip,proto3" json:"skip,omitempty"`
}

func (x *ListOperationsRequest) Reset() {
	*x = ListOperationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_rpc_longrunning_operations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOperationsRequest) ToText() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOperationsRequest) ProtoMessage() {}

func (x *ListOperationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_rpc_longrunning_operations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOperationsRequest.ProtoReflect.Descriptor instead.
func (*ListOperationsRequest) Descriptor() ([]byte, []int) {
	return file_mojo_rpc_longrunning_operations_proto_rawDescGZIP(), []int{0}
}

func (x *ListOperationsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListOperationsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *ListOperationsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListOperationsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListOperationsRequest) GetSkip() int32 {
	if x != nil {
		return x.Skip
	}
	return 0
}

type ListOperationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operations    []*Operation `protobuf:"bytes,1,rep,name=operations,proto3" json:"operations,omitempty"`
	TotalCount    int32        `protobuf:"varint,2000,opt,name=total_count,json=totalCount,proto3" json:"totalCount,omitempty"`
	NextPageToken string       `protobuf:"bytes,2001,opt,name=next_page_token,json=nextPageToken,proto3" json:"nextPageToken,omitempty"`
}

func (x *ListOperationsResponse) Reset() {
	*x = ListOperationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_rpc_longrunning_operations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOperationsResponse) ToText() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOperationsResponse) ProtoMessage() {}

func (x *ListOperationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_rpc_longrunning_operations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOperationsResponse.ProtoReflect.Descriptor instead.
func (*ListOperationsResponse) Descriptor() ([]byte, []int) {
	return file_mojo_rpc_longrunning_operations_proto_rawDescGZIP(), []int{1}
}

func (x *ListOperationsResponse) GetOperations() []*Operation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *ListOperationsResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *ListOperationsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type GetOperationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetOperationRequest) Reset() {
	*x = GetOperationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_rpc_longrunning_operations_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOperationRequest) ToText() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOperationRequest) ProtoMessage() {}

func (x *GetOperationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_rpc_longrunning_operations_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOperationRequest.ProtoReflect.Descriptor instead.
func (*GetOperationRequest) Descriptor() ([]byte, []int) {
	return file_mojo_rpc_longrunning_operations_proto_rawDescGZIP(), []int{2}
}

func (x *GetOperationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteOperationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteOperationRequest) Reset() {
	*x = DeleteOperationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_rpc_longrunning_operations_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOperationRequest) ToText() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOperationRequest) ProtoMessage() {}

func (x *DeleteOperationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_rpc_longrunning_operations_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOperationRequest.ProtoReflect.Descriptor instead.
func (*DeleteOperationRequest) Descriptor() ([]byte, []int) {
	return file_mojo_rpc_longrunning_operations_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteOperationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CancelOperationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CancelOperationRequest) Reset() {
	*x = CancelOperationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_rpc_longrunning_operations_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelOperationRequest) ToText() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelOperationRequest) ProtoMessage() {}

func (x *CancelOperationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_rpc_longrunning_operations_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelOperationRequest.ProtoReflect.Descriptor instead.
func (*CancelOperationRequest) Descriptor() ([]byte, []int) {
	return file_mojo_rpc_longrunning_operations_proto_rawDescGZIP(), []int{4}
}

func (x *CancelOperationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type WaitOperationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Timeout *core.Duration `protobuf:"bytes,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *WaitOperationRequest) Reset() {
	*x = WaitOperationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mojo_rpc_longrunning_operations_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WaitOperationRequest) ToText() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WaitOperationRequest) ProtoMessage() {}

func (x *WaitOperationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mojo_rpc_longrunning_operations_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WaitOperationRequest.ProtoReflect.Descriptor instead.
func (*WaitOperationRequest) Descriptor() ([]byte, []int) {
	return file_mojo_rpc_longrunning_operations_proto_rawDescGZIP(), []int{5}
}

func (x *WaitOperationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WaitOperationRequest) GetTimeout() *core.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

var File_mojo_rpc_longrunning_operations_proto protoreflect.FileDescriptor

var file_mojo_rpc_longrunning_operations_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x72,
	0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x1a, 0x14, 0x6d,
	0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6e, 0x75, 0x6c, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x6d, 0x6f, 0x6a, 0x6f, 0x2f,
	0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x6d, 0x6f, 0x6a, 0x6f,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x9a, 0x01, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0xd0, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0xd1, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x13, 0x0a, 0x04, 0x73, 0x6b, 0x69, 0x70,
	0x18, 0xd2, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x22, 0xa4, 0x01,
	0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6d,
	0x6f, 0x6a, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0xd0, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0xd1,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x29, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x2c, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2c, 0x0a,
	0x16, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x59, 0x0a, 0x14, 0x57,
	0x61, 0x69, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x32, 0xdc, 0x03, 0x0a, 0x0a, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x6c, 0x0a, 0x0f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2b, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x0d, 0x67, 0x65, 0x74, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72,
	0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x51, 0x0a, 0x10, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4e,
	0x75, 0x6c, 0x6c, 0x12, 0x51, 0x0a, 0x10, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x5f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x12, 0x5d, 0x0a, 0x0e, 0x77, 0x61, 0x69, 0x74, 0x5f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e,
	0x57, 0x61, 0x69, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x0a, 0x22, 0x6f, 0x72, 0x67, 0x2e, 0x6d, 0x6f, 0x6a,
	0x6f, 0x2d, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6d, 0x6f, 0x6a, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x42, 0x0f, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x6a, 0x6f, 0x2d,
	0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x6d, 0x6f, 0x6a, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x3b, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mojo_rpc_longrunning_operations_proto_rawDescOnce sync.Once
	file_mojo_rpc_longrunning_operations_proto_rawDescData = file_mojo_rpc_longrunning_operations_proto_rawDesc
)

func file_mojo_rpc_longrunning_operations_proto_rawDescGZIP() []byte {
	file_mojo_rpc_longrunning_operations_proto_rawDescOnce.Do(func() {
		file_mojo_rpc_longrunning_operations_proto_rawDescData = protoimpl.X.CompressGZIP(file_mojo_rpc_longrunning_operations_proto_rawDescData)
	})
	return file_mojo_rpc_longrunning_operations_proto_rawDescData
}

var file_mojo_rpc_longrunning_operations_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_mojo_rpc_longrunning_operations_proto_goTypes = []interface{}{
	(*ListOperationsRequest)(nil),  // 0: mojo.rpc.longrunning.ListOperationsRequest
	(*ListOperationsResponse)(nil), // 1: mojo.rpc.longrunning.ListOperationsResponse
	(*GetOperationRequest)(nil),    // 2: mojo.rpc.longrunning.GetOperationRequest
	(*DeleteOperationRequest)(nil), // 3: mojo.rpc.longrunning.DeleteOperationRequest
	(*CancelOperationRequest)(nil), // 4: mojo.rpc.longrunning.CancelOperationRequest
	(*WaitOperationRequest)(nil),   // 5: mojo.rpc.longrunning.WaitOperationRequest
	(*Operation)(nil),              // 6: mojo.rpc.longrunning.Operation
	(*core.Duration)(nil),          // 7: mojo.core.Duration
	(*core.Null)(nil),              // 8: mojo.core.Null
}
var file_mojo_rpc_longrunning_operations_proto_depIdxs = []int32{
	6, // 0: mojo.rpc.longrunning.ListOperationsResponse.operations:type_name -> mojo.rpc.longrunning.Operation
	7, // 1: mojo.rpc.longrunning.WaitOperationRequest.timeout:type_name -> mojo.core.Duration
	0, // 2: mojo.rpc.longrunning.Operations.list_operations:input_type -> mojo.rpc.longrunning.ListOperationsRequest
	2, // 3: mojo.rpc.longrunning.Operations.get_operation:input_type -> mojo.rpc.longrunning.GetOperationRequest
	3, // 4: mojo.rpc.longrunning.Operations.delete_operation:input_type -> mojo.rpc.longrunning.DeleteOperationRequest
	4, // 5: mojo.rpc.longrunning.Operations.cancel_operation:input_type -> mojo.rpc.longrunning.CancelOperationRequest
	5, // 6: mojo.rpc.longrunning.Operations.wait_operation:input_type -> mojo.rpc.longrunning.WaitOperationRequest
	1, // 7: mojo.rpc.longrunning.Operations.list_operations:output_type -> mojo.rpc.longrunning.ListOperationsResponse
	6, // 8: mojo.rpc.longrunning.Operations.get_operation:output_type -> mojo.rpc.longrunning.Operation
	8, // 9: mojo.rpc.longrunning.Operations.delete_operation:output_type -> mojo.core.Null
	8, // 10: mojo.rpc.longrunning.Operations.cancel_operation:output_type -> mojo.core.Null
	6, // 11: mojo.rpc.longrunning.Operations.wait_operation:output_type -> mojo.rpc.longrunning.Operation
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_mojo_rpc_longrunning_operations_proto_init() }
func file_mojo_rpc_longrunning_operations_proto_init() {
	if File_mojo_rpc_longrunning_operations_proto != nil {
		return
	}
	file_mojo_rpc_longrunning_operation_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mojo_rpc_longrunning_operations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOperationsRequest); i {
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
		file_mojo_rpc_longrunning_operations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOperationsResponse); i {
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
		file_mojo_rpc_longrunning_operations_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOperationRequest); i {
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
		file_mojo_rpc_longrunning_operations_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteOperationRequest); i {
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
		file_mojo_rpc_longrunning_operations_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelOperationRequest); i {
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
		file_mojo_rpc_longrunning_operations_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WaitOperationRequest); i {
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
			RawDescriptor: file_mojo_rpc_longrunning_operations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mojo_rpc_longrunning_operations_proto_goTypes,
		DependencyIndexes: file_mojo_rpc_longrunning_operations_proto_depIdxs,
		MessageInfos:      file_mojo_rpc_longrunning_operations_proto_msgTypes,
	}.Build()
	File_mojo_rpc_longrunning_operations_proto = out.File
	file_mojo_rpc_longrunning_operations_proto_rawDesc = nil
	file_mojo_rpc_longrunning_operations_proto_goTypes = nil
	file_mojo_rpc_longrunning_operations_proto_depIdxs = nil
}
