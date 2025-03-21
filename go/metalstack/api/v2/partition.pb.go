// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: metalstack/api/v2/partition.proto

package apiv2

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

// Partition is a failure domain with machines and switches
type Partition struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of this partition
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Meta for this ip
	Meta *Meta `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
	// Descrpartitiontion of this partition
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// PartitionBootConfiguration defines how metal-hammer boots
	BootConfiguration *PartitionBootConfiguration `protobuf:"bytes,4,opt,name=boot_configuration,json=bootConfiguration,proto3" json:"boot_configuration,omitempty"`
	// DNSServers for this partition
	DnsServer []*DNSServer `protobuf:"bytes,5,rep,name=dns_server,json=dnsServer,proto3" json:"dns_server,omitempty"`
	// NTPServers for this partition
	NtpServer []*NTPServer `protobuf:"bytes,6,rep,name=ntp_server,json=ntpServer,proto3" json:"ntp_server,omitempty"`
	// ManagementServiceAddresses defines where the management is reachable
	// should be in the form <ip|host>:<port>
	MgmtServiceAddresses []string `protobuf:"bytes,7,rep,name=mgmt_service_addresses,json=mgmtServiceAddresses,proto3" json:"mgmt_service_addresses,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *Partition) Reset() {
	*x = Partition{}
	mi := &file_metalstack_api_v2_partition_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Partition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Partition) ProtoMessage() {}

func (x *Partition) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_api_v2_partition_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Partition.ProtoReflect.Descriptor instead.
func (*Partition) Descriptor() ([]byte, []int) {
	return file_metalstack_api_v2_partition_proto_rawDescGZIP(), []int{0}
}

func (x *Partition) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Partition) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Partition) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Partition) GetBootConfiguration() *PartitionBootConfiguration {
	if x != nil {
		return x.BootConfiguration
	}
	return nil
}

func (x *Partition) GetDnsServer() []*DNSServer {
	if x != nil {
		return x.DnsServer
	}
	return nil
}

func (x *Partition) GetNtpServer() []*NTPServer {
	if x != nil {
		return x.NtpServer
	}
	return nil
}

func (x *Partition) GetMgmtServiceAddresses() []string {
	if x != nil {
		return x.MgmtServiceAddresses
	}
	return nil
}

// PartitionBootConfiguration defines how metal-hammer boots
type PartitionBootConfiguration struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ImageURL the url to download the initrd for the boot image
	ImageUrl string `protobuf:"bytes,1,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	// KernelURL the url to download the kernel for the boot image
	KernelUrl string `protobuf:"bytes,2,opt,name=kernel_url,json=kernelUrl,proto3" json:"kernel_url,omitempty"`
	// Commandline the cmdline to the kernel for the boot image
	Commandline   string `protobuf:"bytes,3,opt,name=commandline,proto3" json:"commandline,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PartitionBootConfiguration) Reset() {
	*x = PartitionBootConfiguration{}
	mi := &file_metalstack_api_v2_partition_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PartitionBootConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartitionBootConfiguration) ProtoMessage() {}

func (x *PartitionBootConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_api_v2_partition_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartitionBootConfiguration.ProtoReflect.Descriptor instead.
func (*PartitionBootConfiguration) Descriptor() ([]byte, []int) {
	return file_metalstack_api_v2_partition_proto_rawDescGZIP(), []int{1}
}

func (x *PartitionBootConfiguration) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *PartitionBootConfiguration) GetKernelUrl() string {
	if x != nil {
		return x.KernelUrl
	}
	return ""
}

func (x *PartitionBootConfiguration) GetCommandline() string {
	if x != nil {
		return x.Commandline
	}
	return ""
}

// DNSServer
type DNSServer struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// IP address of this dns server
	Ip            string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DNSServer) Reset() {
	*x = DNSServer{}
	mi := &file_metalstack_api_v2_partition_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DNSServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNSServer) ProtoMessage() {}

func (x *DNSServer) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_api_v2_partition_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNSServer.ProtoReflect.Descriptor instead.
func (*DNSServer) Descriptor() ([]byte, []int) {
	return file_metalstack_api_v2_partition_proto_rawDescGZIP(), []int{2}
}

func (x *DNSServer) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

// NTPServer
type NTPServer struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Address either as ip or hostname
	Address       string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NTPServer) Reset() {
	*x = NTPServer{}
	mi := &file_metalstack_api_v2_partition_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NTPServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NTPServer) ProtoMessage() {}

func (x *NTPServer) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_api_v2_partition_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NTPServer.ProtoReflect.Descriptor instead.
func (*NTPServer) Descriptor() ([]byte, []int) {
	return file_metalstack_api_v2_partition_proto_rawDescGZIP(), []int{3}
}

func (x *NTPServer) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

// PartitionServiceGetRequest is the request payload for a partition get request
type PartitionServiceGetRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the partition to get
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PartitionServiceGetRequest) Reset() {
	*x = PartitionServiceGetRequest{}
	mi := &file_metalstack_api_v2_partition_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PartitionServiceGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartitionServiceGetRequest) ProtoMessage() {}

func (x *PartitionServiceGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_api_v2_partition_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartitionServiceGetRequest.ProtoReflect.Descriptor instead.
func (*PartitionServiceGetRequest) Descriptor() ([]byte, []int) {
	return file_metalstack_api_v2_partition_proto_rawDescGZIP(), []int{4}
}

func (x *PartitionServiceGetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// PartitionServiceListRequest is the request payload for a partition list request
type PartitionServiceListRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the partition to get
	Id            *string `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PartitionServiceListRequest) Reset() {
	*x = PartitionServiceListRequest{}
	mi := &file_metalstack_api_v2_partition_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PartitionServiceListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartitionServiceListRequest) ProtoMessage() {}

func (x *PartitionServiceListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_api_v2_partition_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartitionServiceListRequest.ProtoReflect.Descriptor instead.
func (*PartitionServiceListRequest) Descriptor() ([]byte, []int) {
	return file_metalstack_api_v2_partition_proto_rawDescGZIP(), []int{5}
}

func (x *PartitionServiceListRequest) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

// PartitionServiceGetResponse is the response payload for a partition get request
type PartitionServiceGetResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Ip the partition
	Partition     *Partition `protobuf:"bytes,1,opt,name=partition,proto3" json:"partition,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PartitionServiceGetResponse) Reset() {
	*x = PartitionServiceGetResponse{}
	mi := &file_metalstack_api_v2_partition_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PartitionServiceGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartitionServiceGetResponse) ProtoMessage() {}

func (x *PartitionServiceGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_api_v2_partition_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartitionServiceGetResponse.ProtoReflect.Descriptor instead.
func (*PartitionServiceGetResponse) Descriptor() ([]byte, []int) {
	return file_metalstack_api_v2_partition_proto_rawDescGZIP(), []int{6}
}

func (x *PartitionServiceGetResponse) GetPartition() *Partition {
	if x != nil {
		return x.Partition
	}
	return nil
}

// PartitionServiceListResponse is the response payload for a partition list request
type PartitionServiceListResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Ips the partitions
	Partitions    []*Partition `protobuf:"bytes,1,rep,name=partitions,proto3" json:"partitions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PartitionServiceListResponse) Reset() {
	*x = PartitionServiceListResponse{}
	mi := &file_metalstack_api_v2_partition_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PartitionServiceListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartitionServiceListResponse) ProtoMessage() {}

func (x *PartitionServiceListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_api_v2_partition_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartitionServiceListResponse.ProtoReflect.Descriptor instead.
func (*PartitionServiceListResponse) Descriptor() ([]byte, []int) {
	return file_metalstack_api_v2_partition_proto_rawDescGZIP(), []int{7}
}

func (x *PartitionServiceListResponse) GetPartitions() []*Partition {
	if x != nil {
		return x.Partitions
	}
	return nil
}

var File_metalstack_api_v2_partition_proto protoreflect.FileDescriptor

var file_metalstack_api_v2_partition_proto_rawDesc = string([]byte{
	0x0a, 0x21, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x32, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x03, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba,
	0x48, 0x07, 0x72, 0x05, 0x10, 0x02, 0x18, 0x80, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2b, 0x0a,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0x18, 0x80, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5c, 0x0a, 0x12, 0x62, 0x6f, 0x6f, 0x74, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x6f, 0x6f, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x11, 0x62, 0x6f, 0x6f, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x0a, 0x64, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x4e, 0x53,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x08, 0xba, 0x48, 0x05, 0x92, 0x01, 0x02, 0x10, 0x03,
	0x52, 0x09, 0x64, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x0a, 0x6e,
	0x74, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x32, 0x2e, 0x4e, 0x54, 0x50, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x08, 0xba,
	0x48, 0x05, 0x92, 0x01, 0x02, 0x10, 0x0a, 0x52, 0x09, 0x6e, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x12, 0x34, 0x0a, 0x16, 0x6d, 0x67, 0x6d, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x14, 0x6d, 0x67, 0x6d, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0xee, 0x01, 0x0a, 0x1a, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x6f, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x55, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x38, 0xba, 0x48, 0x35, 0xba,
	0x01, 0x32, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x12, 0x17, 0x75,
	0x72, 0x6c, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20, 0x61, 0x20, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x20, 0x55, 0x52, 0x49, 0x1a, 0x0c, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x69, 0x73, 0x55,
	0x72, 0x69, 0x28, 0x29, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x57,
	0x0a, 0x0a, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x38, 0xba, 0x48, 0x35, 0xba, 0x01, 0x32, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x5f, 0x75, 0x72, 0x6c, 0x12, 0x17, 0x75, 0x72, 0x6c, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20,
	0x62, 0x65, 0x20, 0x61, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x55, 0x52, 0x49, 0x1a, 0x0c,
	0x74, 0x68, 0x69, 0x73, 0x2e, 0x69, 0x73, 0x55, 0x72, 0x69, 0x28, 0x29, 0x52, 0x09, 0x6b, 0x65,
	0x72, 0x6e, 0x65, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x24, 0x0a, 0x09, 0x44, 0x4e, 0x53,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x70, 0x01, 0x52, 0x02, 0x69, 0x70, 0x22,
	0x2f, 0x0a, 0x09, 0x4e, 0x54, 0x50, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba,
	0x48, 0x05, 0x72, 0x03, 0x18, 0x80, 0x02, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0x38, 0x0a, 0x1a, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0x48, 0x07, 0x72,
	0x05, 0x10, 0x02, 0x18, 0x80, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x45, 0x0a, 0x1b, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0x48, 0x07, 0x72, 0x05, 0x10, 0x02, 0x18, 0x80,
	0x01, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69,
	0x64, 0x22, 0x59, 0x0a, 0x1b, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3a, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5c, 0x0a, 0x1c,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0a,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xf5, 0x01, 0x0a, 0x10, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x6e, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x2d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x08, 0xd8, 0xf3, 0x18, 0x03, 0xe0, 0xf3, 0x18, 0x02, 0x12,
	0x71, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x08, 0xd8, 0xf3, 0x18, 0x03, 0xe0, 0xf3,
	0x18, 0x02, 0x42, 0xc4, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x42, 0x0e, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x6c,
	0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x3b,
	0x61, 0x70, 0x69, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x4d, 0x41, 0x58, 0xaa, 0x02, 0x11, 0x4d, 0x65,
	0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x32, 0xca,
	0x02, 0x11, 0x4d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x41, 0x70, 0x69,
	0x5c, 0x56, 0x32, 0xe2, 0x02, 0x1d, 0x4d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x4d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_metalstack_api_v2_partition_proto_rawDescOnce sync.Once
	file_metalstack_api_v2_partition_proto_rawDescData []byte
)

func file_metalstack_api_v2_partition_proto_rawDescGZIP() []byte {
	file_metalstack_api_v2_partition_proto_rawDescOnce.Do(func() {
		file_metalstack_api_v2_partition_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_metalstack_api_v2_partition_proto_rawDesc), len(file_metalstack_api_v2_partition_proto_rawDesc)))
	})
	return file_metalstack_api_v2_partition_proto_rawDescData
}

var file_metalstack_api_v2_partition_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_metalstack_api_v2_partition_proto_goTypes = []any{
	(*Partition)(nil),                    // 0: metalstack.api.v2.Partition
	(*PartitionBootConfiguration)(nil),   // 1: metalstack.api.v2.PartitionBootConfiguration
	(*DNSServer)(nil),                    // 2: metalstack.api.v2.DNSServer
	(*NTPServer)(nil),                    // 3: metalstack.api.v2.NTPServer
	(*PartitionServiceGetRequest)(nil),   // 4: metalstack.api.v2.PartitionServiceGetRequest
	(*PartitionServiceListRequest)(nil),  // 5: metalstack.api.v2.PartitionServiceListRequest
	(*PartitionServiceGetResponse)(nil),  // 6: metalstack.api.v2.PartitionServiceGetResponse
	(*PartitionServiceListResponse)(nil), // 7: metalstack.api.v2.PartitionServiceListResponse
	(*Meta)(nil),                         // 8: metalstack.api.v2.Meta
}
var file_metalstack_api_v2_partition_proto_depIdxs = []int32{
	8, // 0: metalstack.api.v2.Partition.meta:type_name -> metalstack.api.v2.Meta
	1, // 1: metalstack.api.v2.Partition.boot_configuration:type_name -> metalstack.api.v2.PartitionBootConfiguration
	2, // 2: metalstack.api.v2.Partition.dns_server:type_name -> metalstack.api.v2.DNSServer
	3, // 3: metalstack.api.v2.Partition.ntp_server:type_name -> metalstack.api.v2.NTPServer
	0, // 4: metalstack.api.v2.PartitionServiceGetResponse.partition:type_name -> metalstack.api.v2.Partition
	0, // 5: metalstack.api.v2.PartitionServiceListResponse.partitions:type_name -> metalstack.api.v2.Partition
	4, // 6: metalstack.api.v2.PartitionService.Get:input_type -> metalstack.api.v2.PartitionServiceGetRequest
	5, // 7: metalstack.api.v2.PartitionService.List:input_type -> metalstack.api.v2.PartitionServiceListRequest
	6, // 8: metalstack.api.v2.PartitionService.Get:output_type -> metalstack.api.v2.PartitionServiceGetResponse
	7, // 9: metalstack.api.v2.PartitionService.List:output_type -> metalstack.api.v2.PartitionServiceListResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_metalstack_api_v2_partition_proto_init() }
func file_metalstack_api_v2_partition_proto_init() {
	if File_metalstack_api_v2_partition_proto != nil {
		return
	}
	file_metalstack_api_v2_common_proto_init()
	file_metalstack_api_v2_partition_proto_msgTypes[5].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_metalstack_api_v2_partition_proto_rawDesc), len(file_metalstack_api_v2_partition_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metalstack_api_v2_partition_proto_goTypes,
		DependencyIndexes: file_metalstack_api_v2_partition_proto_depIdxs,
		MessageInfos:      file_metalstack_api_v2_partition_proto_msgTypes,
	}.Build()
	File_metalstack_api_v2_partition_proto = out.File
	file_metalstack_api_v2_partition_proto_goTypes = nil
	file_metalstack_api_v2_partition_proto_depIdxs = nil
}
