// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: metalstack/admin/v2/partition.proto

package adminv2

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	v2 "github.com/metal-stack/api/go/metalstack/api/v2"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// PartitionServiceCreateRequest is the request payload for a partition create request
type PartitionServiceCreateRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Partition the partition
	Partition     *v2.Partition `protobuf:"bytes,1,opt,name=partition,proto3" json:"partition,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PartitionServiceCreateRequest) Reset() {
	*x = PartitionServiceCreateRequest{}
	mi := &file_metalstack_admin_v2_partition_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PartitionServiceCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartitionServiceCreateRequest) ProtoMessage() {}

func (x *PartitionServiceCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_admin_v2_partition_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartitionServiceCreateRequest.ProtoReflect.Descriptor instead.
func (*PartitionServiceCreateRequest) Descriptor() ([]byte, []int) {
	return file_metalstack_admin_v2_partition_proto_rawDescGZIP(), []int{0}
}

func (x *PartitionServiceCreateRequest) GetPartition() *v2.Partition {
	if x != nil {
		return x.Partition
	}
	return nil
}

// PartitionServiceUpdateRequest is the request payload for a partition update request
type PartitionServiceUpdateRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Partition the partition
	Partition     *v2.Partition `protobuf:"bytes,1,opt,name=partition,proto3" json:"partition,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PartitionServiceUpdateRequest) Reset() {
	*x = PartitionServiceUpdateRequest{}
	mi := &file_metalstack_admin_v2_partition_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PartitionServiceUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartitionServiceUpdateRequest) ProtoMessage() {}

func (x *PartitionServiceUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_admin_v2_partition_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartitionServiceUpdateRequest.ProtoReflect.Descriptor instead.
func (*PartitionServiceUpdateRequest) Descriptor() ([]byte, []int) {
	return file_metalstack_admin_v2_partition_proto_rawDescGZIP(), []int{1}
}

func (x *PartitionServiceUpdateRequest) GetPartition() *v2.Partition {
	if x != nil {
		return x.Partition
	}
	return nil
}

// PartitionServiceDeleteRequest is the request payload for a partition delete request
type PartitionServiceDeleteRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the partition to get
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PartitionServiceDeleteRequest) Reset() {
	*x = PartitionServiceDeleteRequest{}
	mi := &file_metalstack_admin_v2_partition_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PartitionServiceDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartitionServiceDeleteRequest) ProtoMessage() {}

func (x *PartitionServiceDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_admin_v2_partition_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartitionServiceDeleteRequest.ProtoReflect.Descriptor instead.
func (*PartitionServiceDeleteRequest) Descriptor() ([]byte, []int) {
	return file_metalstack_admin_v2_partition_proto_rawDescGZIP(), []int{2}
}

func (x *PartitionServiceDeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// PartitionServiceCreateResponse is the response payload for a partition create request
type PartitionServiceCreateResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Partition the partition
	Partition     *v2.Partition `protobuf:"bytes,1,opt,name=partition,proto3" json:"partition,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PartitionServiceCreateResponse) Reset() {
	*x = PartitionServiceCreateResponse{}
	mi := &file_metalstack_admin_v2_partition_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PartitionServiceCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartitionServiceCreateResponse) ProtoMessage() {}

func (x *PartitionServiceCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_admin_v2_partition_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartitionServiceCreateResponse.ProtoReflect.Descriptor instead.
func (*PartitionServiceCreateResponse) Descriptor() ([]byte, []int) {
	return file_metalstack_admin_v2_partition_proto_rawDescGZIP(), []int{3}
}

func (x *PartitionServiceCreateResponse) GetPartition() *v2.Partition {
	if x != nil {
		return x.Partition
	}
	return nil
}

// PartitionServiceUpdateResponse is the response payload for a partition update request
type PartitionServiceUpdateResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Partition the partition
	Partition     *v2.Partition `protobuf:"bytes,1,opt,name=partition,proto3" json:"partition,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PartitionServiceUpdateResponse) Reset() {
	*x = PartitionServiceUpdateResponse{}
	mi := &file_metalstack_admin_v2_partition_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PartitionServiceUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartitionServiceUpdateResponse) ProtoMessage() {}

func (x *PartitionServiceUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_admin_v2_partition_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartitionServiceUpdateResponse.ProtoReflect.Descriptor instead.
func (*PartitionServiceUpdateResponse) Descriptor() ([]byte, []int) {
	return file_metalstack_admin_v2_partition_proto_rawDescGZIP(), []int{4}
}

func (x *PartitionServiceUpdateResponse) GetPartition() *v2.Partition {
	if x != nil {
		return x.Partition
	}
	return nil
}

// PartitionServiceCapacityResponse is the response payload for a partition delete request
type PartitionServiceDeleteResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Partition the partition
	Partition     *v2.Partition `protobuf:"bytes,1,opt,name=partition,proto3" json:"partition,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PartitionServiceDeleteResponse) Reset() {
	*x = PartitionServiceDeleteResponse{}
	mi := &file_metalstack_admin_v2_partition_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PartitionServiceDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartitionServiceDeleteResponse) ProtoMessage() {}

func (x *PartitionServiceDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_admin_v2_partition_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartitionServiceDeleteResponse.ProtoReflect.Descriptor instead.
func (*PartitionServiceDeleteResponse) Descriptor() ([]byte, []int) {
	return file_metalstack_admin_v2_partition_proto_rawDescGZIP(), []int{5}
}

func (x *PartitionServiceDeleteResponse) GetPartition() *v2.Partition {
	if x != nil {
		return x.Partition
	}
	return nil
}

// PartitionServiceListRequest is the request payload for a partition capacity request
type PartitionServiceCapacityRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the partition to get
	Id *string `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// Size of machines to show the capacity
	Size *string `protobuf:"bytes,2,opt,name=size,proto3,oneof" json:"size,omitempty"`
	// Project of machines to show the capacity
	Project       *string `protobuf:"bytes,3,opt,name=project,proto3,oneof" json:"project,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PartitionServiceCapacityRequest) Reset() {
	*x = PartitionServiceCapacityRequest{}
	mi := &file_metalstack_admin_v2_partition_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PartitionServiceCapacityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartitionServiceCapacityRequest) ProtoMessage() {}

func (x *PartitionServiceCapacityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_admin_v2_partition_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartitionServiceCapacityRequest.ProtoReflect.Descriptor instead.
func (*PartitionServiceCapacityRequest) Descriptor() ([]byte, []int) {
	return file_metalstack_admin_v2_partition_proto_rawDescGZIP(), []int{6}
}

func (x *PartitionServiceCapacityRequest) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *PartitionServiceCapacityRequest) GetSize() string {
	if x != nil && x.Size != nil {
		return *x.Size
	}
	return ""
}

func (x *PartitionServiceCapacityRequest) GetProject() string {
	if x != nil && x.Project != nil {
		return *x.Project
	}
	return ""
}

// PartitionServiceCapacityResponse is the response payload for a partition capacity request
type PartitionServiceCapacityResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Size is the size id correlating to all counts in this server capacity.
	Size string `protobuf:"bytes,1,opt,name=size,proto3" json:"size,omitempty"`
	// Total is the total amount of machines for this size.
	Total int64 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	// PhonedHome is the amount of machines that are currently in the provisioning state "phoned home".
	PhonedHome int64 `protobuf:"varint,3,opt,name=phoned_home,json=phonedHome,proto3" json:"phoned_home,omitempty"`
	// Waiting is the amount of machines that are currently in the provisioning state "waiting".
	Waiting int64 `protobuf:"varint,4,opt,name=waiting,proto3" json:"waiting,omitempty"`
	// Other is the amount of machines that are neither in the provisioning state waiting nor in phoned home but in another provisioning state.
	Other int64 `protobuf:"varint,5,opt,name=other,proto3" json:"other,omitempty"`
	// OtherMachines contains the machine IDs for machines that were classified into "Other".
	OtherMachines []string `protobuf:"bytes,6,rep,name=other_machines,json=otherMachines,proto3" json:"other_machines,omitempty"`
	// Allocated is the amount of machines that are currently allocated.
	Allocated int64 `protobuf:"varint,7,opt,name=allocated,proto3" json:"allocated,omitempty"`
	// Allocatable is the amount of machines in a partition is the amount of machines that can be allocated.
	// Effectively this is the amount of waiting machines minus the machines that are unavailable due to machine state or un-allocatable. Size reservations are not considered in this count.
	Allocatable int64 `protobuf:"varint,8,opt,name=allocatable,proto3" json:"allocatable,omitempty"`
	// Free is the amount of machines in a partition that can be freely allocated at any given moment by a project.
	// Effectively this is the amount of waiting machines minus the machines that are unavailable due to machine state or un-allocatable due to size reservations.
	Free int64 `protobuf:"varint,9,opt,name=free,proto3" json:"free,omitempty"`
	// Unavailable is the amount of machine in a partition that are currently not allocatable because they are not waiting or
	// not in the machine state "available", e.g. locked or reserved.
	Unavailable int64 `protobuf:"varint,10,opt,name=unavailable,proto3" json:"unavailable,omitempty"`
	// Faulty is the amount of machines that are neither allocated nor in the pool of available machines because they report an error.
	Faulty int64 `protobuf:"varint,11,opt,name=faulty,proto3" json:"faulty,omitempty"`
	// FaultyMachines contains the machine IDs for machines that were classified into "Faulty".
	FaultyMachines []string `protobuf:"bytes,12,rep,name=faulty_machines,json=faultyMachines,proto3" json:"faulty_machines,omitempty"`
	// Reservations is the amount of reservations made for this size.
	Reservations int64 `protobuf:"varint,13,opt,name=reservations,proto3" json:"reservations,omitempty"`
	// UsedReservations is the amount of reservations already used up for this size.
	UsedReservations int64 `protobuf:"varint,14,opt,name=used_reservations,json=usedReservations,proto3" json:"used_reservations,omitempty"`
	// RemainingReservations is the amount of reservations remaining for this size.
	RemainingReservations int64 `protobuf:"varint,15,opt,name=remaining_reservations,json=remainingReservations,proto3" json:"remaining_reservations,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *PartitionServiceCapacityResponse) Reset() {
	*x = PartitionServiceCapacityResponse{}
	mi := &file_metalstack_admin_v2_partition_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PartitionServiceCapacityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartitionServiceCapacityResponse) ProtoMessage() {}

func (x *PartitionServiceCapacityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_admin_v2_partition_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartitionServiceCapacityResponse.ProtoReflect.Descriptor instead.
func (*PartitionServiceCapacityResponse) Descriptor() ([]byte, []int) {
	return file_metalstack_admin_v2_partition_proto_rawDescGZIP(), []int{7}
}

func (x *PartitionServiceCapacityResponse) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *PartitionServiceCapacityResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PartitionServiceCapacityResponse) GetPhonedHome() int64 {
	if x != nil {
		return x.PhonedHome
	}
	return 0
}

func (x *PartitionServiceCapacityResponse) GetWaiting() int64 {
	if x != nil {
		return x.Waiting
	}
	return 0
}

func (x *PartitionServiceCapacityResponse) GetOther() int64 {
	if x != nil {
		return x.Other
	}
	return 0
}

func (x *PartitionServiceCapacityResponse) GetOtherMachines() []string {
	if x != nil {
		return x.OtherMachines
	}
	return nil
}

func (x *PartitionServiceCapacityResponse) GetAllocated() int64 {
	if x != nil {
		return x.Allocated
	}
	return 0
}

func (x *PartitionServiceCapacityResponse) GetAllocatable() int64 {
	if x != nil {
		return x.Allocatable
	}
	return 0
}

func (x *PartitionServiceCapacityResponse) GetFree() int64 {
	if x != nil {
		return x.Free
	}
	return 0
}

func (x *PartitionServiceCapacityResponse) GetUnavailable() int64 {
	if x != nil {
		return x.Unavailable
	}
	return 0
}

func (x *PartitionServiceCapacityResponse) GetFaulty() int64 {
	if x != nil {
		return x.Faulty
	}
	return 0
}

func (x *PartitionServiceCapacityResponse) GetFaultyMachines() []string {
	if x != nil {
		return x.FaultyMachines
	}
	return nil
}

func (x *PartitionServiceCapacityResponse) GetReservations() int64 {
	if x != nil {
		return x.Reservations
	}
	return 0
}

func (x *PartitionServiceCapacityResponse) GetUsedReservations() int64 {
	if x != nil {
		return x.UsedReservations
	}
	return 0
}

func (x *PartitionServiceCapacityResponse) GetRemainingReservations() int64 {
	if x != nil {
		return x.RemainingReservations
	}
	return 0
}

var File_metalstack_admin_v2_partition_proto protoreflect.FileDescriptor

var file_metalstack_admin_v2_partition_proto_rawDesc = string([]byte{
	0x0a, 0x23, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x32, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x1d, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x09, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5b, 0x0a, 0x1d, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3b, 0x0a, 0x1d, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0a, 0xba, 0x48, 0x07, 0x72, 0x05, 0x10, 0x02, 0x18, 0x80, 0x01, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x5c, 0x0a, 0x1e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x5c, 0x0a, 0x1e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3a, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5c, 0x0a,
	0x1e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3a, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xac, 0x01, 0x0a, 0x1f,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0x48, 0x07,
	0x72, 0x05, 0x10, 0x02, 0x18, 0x80, 0x01, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x23, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0xba, 0x48, 0x07, 0x72, 0x05, 0x10, 0x02, 0x18, 0x80, 0x01, 0x48, 0x01, 0x52, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01,
	0x48, 0x02, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x88, 0x01, 0x01, 0x42, 0x05,
	0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x83, 0x04, 0x0a, 0x20, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43,
	0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x64, 0x5f, 0x68, 0x6f, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x64, 0x48, 0x6f, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x61,
	0x69, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x77, 0x61, 0x69,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0d, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x65, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x66, 0x72, 0x65, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x75, 0x6e, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x79, 0x12,
	0x27, 0x0a, 0x0f, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x79, 0x5f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x79,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x11,
	0x75, 0x73, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x75, 0x73, 0x65, 0x64, 0x52, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x35, 0x0a, 0x16, 0x72, 0x65, 0x6d,
	0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x15, 0x72, 0x65, 0x6d, 0x61, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x32, 0x92, 0x04, 0x0a, 0x10, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x32, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x09, 0xd2, 0xf3, 0x18, 0x01, 0x01, 0xe0,
	0xf3, 0x18, 0x01, 0x12, 0x7c, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x32, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x76, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x33, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x09, 0xd2, 0xf3, 0x18, 0x01, 0x01, 0xe0, 0xf3, 0x18,
	0x01, 0x12, 0x7c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x32, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76,
	0x32, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x33, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x09, 0xd2, 0xf3, 0x18, 0x01, 0x01, 0xe0, 0xf3, 0x18, 0x01, 0x12,
	0x83, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x34, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x76, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x35, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0a, 0xd2, 0xf3, 0x18, 0x02, 0x02,
	0x01, 0xe0, 0xf3, 0x18, 0x02, 0x42, 0xd2, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76,
	0x32, 0x42, 0x0e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x6f, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x32, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x76, 0x32, 0xa2, 0x02,
	0x03, 0x4d, 0x41, 0x58, 0xaa, 0x02, 0x13, 0x4d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x13, 0x4d, 0x65, 0x74,
	0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x56, 0x32,
	0xe2, 0x02, 0x1f, 0x4d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x5c, 0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x15, 0x4d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x3a,
	0x3a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_metalstack_admin_v2_partition_proto_rawDescOnce sync.Once
	file_metalstack_admin_v2_partition_proto_rawDescData []byte
)

func file_metalstack_admin_v2_partition_proto_rawDescGZIP() []byte {
	file_metalstack_admin_v2_partition_proto_rawDescOnce.Do(func() {
		file_metalstack_admin_v2_partition_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_metalstack_admin_v2_partition_proto_rawDesc), len(file_metalstack_admin_v2_partition_proto_rawDesc)))
	})
	return file_metalstack_admin_v2_partition_proto_rawDescData
}

var file_metalstack_admin_v2_partition_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_metalstack_admin_v2_partition_proto_goTypes = []any{
	(*PartitionServiceCreateRequest)(nil),    // 0: metalstack.admin.v2.PartitionServiceCreateRequest
	(*PartitionServiceUpdateRequest)(nil),    // 1: metalstack.admin.v2.PartitionServiceUpdateRequest
	(*PartitionServiceDeleteRequest)(nil),    // 2: metalstack.admin.v2.PartitionServiceDeleteRequest
	(*PartitionServiceCreateResponse)(nil),   // 3: metalstack.admin.v2.PartitionServiceCreateResponse
	(*PartitionServiceUpdateResponse)(nil),   // 4: metalstack.admin.v2.PartitionServiceUpdateResponse
	(*PartitionServiceDeleteResponse)(nil),   // 5: metalstack.admin.v2.PartitionServiceDeleteResponse
	(*PartitionServiceCapacityRequest)(nil),  // 6: metalstack.admin.v2.PartitionServiceCapacityRequest
	(*PartitionServiceCapacityResponse)(nil), // 7: metalstack.admin.v2.PartitionServiceCapacityResponse
	(*v2.Partition)(nil),                     // 8: metalstack.api.v2.Partition
}
var file_metalstack_admin_v2_partition_proto_depIdxs = []int32{
	8, // 0: metalstack.admin.v2.PartitionServiceCreateRequest.partition:type_name -> metalstack.api.v2.Partition
	8, // 1: metalstack.admin.v2.PartitionServiceUpdateRequest.partition:type_name -> metalstack.api.v2.Partition
	8, // 2: metalstack.admin.v2.PartitionServiceCreateResponse.partition:type_name -> metalstack.api.v2.Partition
	8, // 3: metalstack.admin.v2.PartitionServiceUpdateResponse.partition:type_name -> metalstack.api.v2.Partition
	8, // 4: metalstack.admin.v2.PartitionServiceDeleteResponse.partition:type_name -> metalstack.api.v2.Partition
	0, // 5: metalstack.admin.v2.PartitionService.Create:input_type -> metalstack.admin.v2.PartitionServiceCreateRequest
	1, // 6: metalstack.admin.v2.PartitionService.Update:input_type -> metalstack.admin.v2.PartitionServiceUpdateRequest
	2, // 7: metalstack.admin.v2.PartitionService.Delete:input_type -> metalstack.admin.v2.PartitionServiceDeleteRequest
	6, // 8: metalstack.admin.v2.PartitionService.Capacity:input_type -> metalstack.admin.v2.PartitionServiceCapacityRequest
	3, // 9: metalstack.admin.v2.PartitionService.Create:output_type -> metalstack.admin.v2.PartitionServiceCreateResponse
	4, // 10: metalstack.admin.v2.PartitionService.Update:output_type -> metalstack.admin.v2.PartitionServiceUpdateResponse
	5, // 11: metalstack.admin.v2.PartitionService.Delete:output_type -> metalstack.admin.v2.PartitionServiceDeleteResponse
	7, // 12: metalstack.admin.v2.PartitionService.Capacity:output_type -> metalstack.admin.v2.PartitionServiceCapacityResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_metalstack_admin_v2_partition_proto_init() }
func file_metalstack_admin_v2_partition_proto_init() {
	if File_metalstack_admin_v2_partition_proto != nil {
		return
	}
	file_metalstack_admin_v2_partition_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_metalstack_admin_v2_partition_proto_rawDesc), len(file_metalstack_admin_v2_partition_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metalstack_admin_v2_partition_proto_goTypes,
		DependencyIndexes: file_metalstack_admin_v2_partition_proto_depIdxs,
		MessageInfos:      file_metalstack_admin_v2_partition_proto_msgTypes,
	}.Build()
	File_metalstack_admin_v2_partition_proto = out.File
	file_metalstack_admin_v2_partition_proto_goTypes = nil
	file_metalstack_admin_v2_partition_proto_depIdxs = nil
}
