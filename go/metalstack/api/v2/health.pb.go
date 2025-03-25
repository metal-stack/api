// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: metalstack/api/v2/health.proto

package apiv2

import (
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

// ServiceStatus defines the status of a service
type ServiceStatus int32

const (
	// SERVICE_STATUS_UNSPECIFIED service status is not known or unspecified
	ServiceStatus_SERVICE_STATUS_UNSPECIFIED ServiceStatus = 0
	// SERVICE_STATUS_DEGRADED the service is in degraded status, not the whole functionality is available
	ServiceStatus_SERVICE_STATUS_DEGRADED ServiceStatus = 1
	// SERVICE_STATUS_UNHEALTHY the service is in unhealthy status, serious impact is expected
	ServiceStatus_SERVICE_STATUS_UNHEALTHY ServiceStatus = 2
	// SERVICE_STATUS_HEALTHY the service is in healthy status e.g. fully functional
	ServiceStatus_SERVICE_STATUS_HEALTHY ServiceStatus = 3
)

// Enum value maps for ServiceStatus.
var (
	ServiceStatus_name = map[int32]string{
		0: "SERVICE_STATUS_UNSPECIFIED",
		1: "SERVICE_STATUS_DEGRADED",
		2: "SERVICE_STATUS_UNHEALTHY",
		3: "SERVICE_STATUS_HEALTHY",
	}
	ServiceStatus_value = map[string]int32{
		"SERVICE_STATUS_UNSPECIFIED": 0,
		"SERVICE_STATUS_DEGRADED":    1,
		"SERVICE_STATUS_UNHEALTHY":   2,
		"SERVICE_STATUS_HEALTHY":     3,
	}
)

func (x ServiceStatus) Enum() *ServiceStatus {
	p := new(ServiceStatus)
	*p = x
	return p
}

func (x ServiceStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_metalstack_api_v2_health_proto_enumTypes[0].Descriptor()
}

func (ServiceStatus) Type() protoreflect.EnumType {
	return &file_metalstack_api_v2_health_proto_enumTypes[0]
}

func (x ServiceStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceStatus.Descriptor instead.
func (ServiceStatus) EnumDescriptor() ([]byte, []int) {
	return file_metalstack_api_v2_health_proto_rawDescGZIP(), []int{0}
}

// Service defines the service for which the healtyness is reported
type Service int32

const (
	// SERVICE_UNSPECIFIED is a unknown service
	Service_SERVICE_UNSPECIFIED Service = 0
	// SERVICE_IPAM the ipam service
	Service_SERVICE_IPAM Service = 1
	// SERVICE_RETHINK the rethinkdb
	Service_SERVICE_RETHINK Service = 2
	// SERVICE_MACHINES the machine service
	Service_SERVICE_MACHINES Service = 3
	// SERVICE_AUDIT the auditing
	Service_SERVICE_AUDIT Service = 4
	// SERVICE_VPN the vpn service
	Service_SERVICE_VPN Service = 5
)

// Enum value maps for Service.
var (
	Service_name = map[int32]string{
		0: "SERVICE_UNSPECIFIED",
		1: "SERVICE_IPAM",
		2: "SERVICE_RETHINK",
		3: "SERVICE_MACHINES",
		4: "SERVICE_AUDIT",
		5: "SERVICE_VPN",
	}
	Service_value = map[string]int32{
		"SERVICE_UNSPECIFIED": 0,
		"SERVICE_IPAM":        1,
		"SERVICE_RETHINK":     2,
		"SERVICE_MACHINES":    3,
		"SERVICE_AUDIT":       4,
		"SERVICE_VPN":         5,
	}
)

func (x Service) Enum() *Service {
	p := new(Service)
	*p = x
	return p
}

func (x Service) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Service) Descriptor() protoreflect.EnumDescriptor {
	return file_metalstack_api_v2_health_proto_enumTypes[1].Descriptor()
}

func (Service) Type() protoreflect.EnumType {
	return &file_metalstack_api_v2_health_proto_enumTypes[1]
}

func (x Service) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Service.Descriptor instead.
func (Service) EnumDescriptor() ([]byte, []int) {
	return file_metalstack_api_v2_health_proto_rawDescGZIP(), []int{1}
}

// Health reports the health status of all services
type Health struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Services the health of all individual services
	Services      []*HealthStatus `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Health) Reset() {
	*x = Health{}
	mi := &file_metalstack_api_v2_health_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Health) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Health) ProtoMessage() {}

func (x *Health) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_api_v2_health_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Health.ProtoReflect.Descriptor instead.
func (*Health) Descriptor() ([]byte, []int) {
	return file_metalstack_api_v2_health_proto_rawDescGZIP(), []int{0}
}

func (x *Health) GetServices() []*HealthStatus {
	if x != nil {
		return x.Services
	}
	return nil
}

// HealthStatus the health of one service
type HealthStatus struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name the name of the service
	Name Service `protobuf:"varint,1,opt,name=name,proto3,enum=metalstack.api.v2.Service" json:"name,omitempty"`
	// Status the status of this service
	Status ServiceStatus `protobuf:"varint,2,opt,name=status,proto3,enum=metalstack.api.v2.ServiceStatus" json:"status,omitempty"`
	// Message describes the reason for the unhealthy status if possible
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// Partitions describes the health of the service by partition
	Partitions    map[string]*PartitionHealth `protobuf:"bytes,4,rep,name=partitions,proto3" json:"partitions,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HealthStatus) Reset() {
	*x = HealthStatus{}
	mi := &file_metalstack_api_v2_health_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HealthStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthStatus) ProtoMessage() {}

func (x *HealthStatus) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_api_v2_health_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthStatus.ProtoReflect.Descriptor instead.
func (*HealthStatus) Descriptor() ([]byte, []int) {
	return file_metalstack_api_v2_health_proto_rawDescGZIP(), []int{1}
}

func (x *HealthStatus) GetName() Service {
	if x != nil {
		return x.Name
	}
	return Service_SERVICE_UNSPECIFIED
}

func (x *HealthStatus) GetStatus() ServiceStatus {
	if x != nil {
		return x.Status
	}
	return ServiceStatus_SERVICE_STATUS_UNSPECIFIED
}

func (x *HealthStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *HealthStatus) GetPartitions() map[string]*PartitionHealth {
	if x != nil {
		return x.Partitions
	}
	return nil
}

// PartitionHealth the status of a specific service in this partition
type PartitionHealth struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Status the health status of the service in this partition
	Status ServiceStatus `protobuf:"varint,1,opt,name=status,proto3,enum=metalstack.api.v2.ServiceStatus" json:"status,omitempty"`
	// Message describes the reason for the unhealthy status if possible
	Message       string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PartitionHealth) Reset() {
	*x = PartitionHealth{}
	mi := &file_metalstack_api_v2_health_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PartitionHealth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartitionHealth) ProtoMessage() {}

func (x *PartitionHealth) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_api_v2_health_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartitionHealth.ProtoReflect.Descriptor instead.
func (*PartitionHealth) Descriptor() ([]byte, []int) {
	return file_metalstack_api_v2_health_proto_rawDescGZIP(), []int{2}
}

func (x *PartitionHealth) GetStatus() ServiceStatus {
	if x != nil {
		return x.Status
	}
	return ServiceStatus_SERVICE_STATUS_UNSPECIFIED
}

func (x *PartitionHealth) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// HealthServiceGetRequest is request payload to get the health of the system
type HealthServiceGetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HealthServiceGetRequest) Reset() {
	*x = HealthServiceGetRequest{}
	mi := &file_metalstack_api_v2_health_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HealthServiceGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthServiceGetRequest) ProtoMessage() {}

func (x *HealthServiceGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_api_v2_health_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthServiceGetRequest.ProtoReflect.Descriptor instead.
func (*HealthServiceGetRequest) Descriptor() ([]byte, []int) {
	return file_metalstack_api_v2_health_proto_rawDescGZIP(), []int{3}
}

// HealthServiceGetRequest is the response payload with the health of the system
type HealthServiceGetResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Health is the overall health of the system
	Health        *Health `protobuf:"bytes,1,opt,name=health,proto3" json:"health,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HealthServiceGetResponse) Reset() {
	*x = HealthServiceGetResponse{}
	mi := &file_metalstack_api_v2_health_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HealthServiceGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthServiceGetResponse) ProtoMessage() {}

func (x *HealthServiceGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_api_v2_health_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthServiceGetResponse.ProtoReflect.Descriptor instead.
func (*HealthServiceGetResponse) Descriptor() ([]byte, []int) {
	return file_metalstack_api_v2_health_proto_rawDescGZIP(), []int{4}
}

func (x *HealthServiceGetResponse) GetHealth() *Health {
	if x != nil {
		return x.Health
	}
	return nil
}

var File_metalstack_api_v2_health_proto protoreflect.FileDescriptor

const file_metalstack_api_v2_health_proto_rawDesc = "" +
	"\n" +
	"\x1emetalstack/api/v2/health.proto\x12\x11metalstack.api.v2\x1a\x1emetalstack/api/v2/common.proto\"E\n" +
	"\x06Health\x12;\n" +
	"\bservices\x18\x01 \x03(\v2\x1f.metalstack.api.v2.HealthStatusR\bservices\"\xc6\x02\n" +
	"\fHealthStatus\x12.\n" +
	"\x04name\x18\x01 \x01(\x0e2\x1a.metalstack.api.v2.ServiceR\x04name\x128\n" +
	"\x06status\x18\x02 \x01(\x0e2 .metalstack.api.v2.ServiceStatusR\x06status\x12\x18\n" +
	"\amessage\x18\x03 \x01(\tR\amessage\x12O\n" +
	"\n" +
	"partitions\x18\x04 \x03(\v2/.metalstack.api.v2.HealthStatus.PartitionsEntryR\n" +
	"partitions\x1aa\n" +
	"\x0fPartitionsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x128\n" +
	"\x05value\x18\x02 \x01(\v2\".metalstack.api.v2.PartitionHealthR\x05value:\x028\x01\"e\n" +
	"\x0fPartitionHealth\x128\n" +
	"\x06status\x18\x01 \x01(\x0e2 .metalstack.api.v2.ServiceStatusR\x06status\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\"\x19\n" +
	"\x17HealthServiceGetRequest\"M\n" +
	"\x18HealthServiceGetResponse\x121\n" +
	"\x06health\x18\x01 \x01(\v2\x19.metalstack.api.v2.HealthR\x06health*\x86\x01\n" +
	"\rServiceStatus\x12\x1e\n" +
	"\x1aSERVICE_STATUS_UNSPECIFIED\x10\x00\x12\x1b\n" +
	"\x17SERVICE_STATUS_DEGRADED\x10\x01\x12\x1c\n" +
	"\x18SERVICE_STATUS_UNHEALTHY\x10\x02\x12\x1a\n" +
	"\x16SERVICE_STATUS_HEALTHY\x10\x03*\x83\x01\n" +
	"\aService\x12\x17\n" +
	"\x13SERVICE_UNSPECIFIED\x10\x00\x12\x10\n" +
	"\fSERVICE_IPAM\x10\x01\x12\x13\n" +
	"\x0fSERVICE_RETHINK\x10\x02\x12\x14\n" +
	"\x10SERVICE_MACHINES\x10\x03\x12\x11\n" +
	"\rSERVICE_AUDIT\x10\x04\x12\x0f\n" +
	"\vSERVICE_VPN\x10\x052y\n" +
	"\rHealthService\x12h\n" +
	"\x03Get\x12*.metalstack.api.v2.HealthServiceGetRequest\x1a+.metalstack.api.v2.HealthServiceGetResponse\"\b\xd8\xf3\x18\x01\xe0\xf3\x18\x02B\xc1\x01\n" +
	"\x15com.metalstack.api.v2B\vHealthProtoP\x01Z5github.com/metal-stack/api/go/metalstack/api/v2;apiv2\xa2\x02\x03MAX\xaa\x02\x11Metalstack.Api.V2\xca\x02\x11Metalstack\\Api\\V2\xe2\x02\x1dMetalstack\\Api\\V2\\GPBMetadata\xea\x02\x13Metalstack::Api::V2b\x06proto3"

var (
	file_metalstack_api_v2_health_proto_rawDescOnce sync.Once
	file_metalstack_api_v2_health_proto_rawDescData []byte
)

func file_metalstack_api_v2_health_proto_rawDescGZIP() []byte {
	file_metalstack_api_v2_health_proto_rawDescOnce.Do(func() {
		file_metalstack_api_v2_health_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_metalstack_api_v2_health_proto_rawDesc), len(file_metalstack_api_v2_health_proto_rawDesc)))
	})
	return file_metalstack_api_v2_health_proto_rawDescData
}

var file_metalstack_api_v2_health_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_metalstack_api_v2_health_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_metalstack_api_v2_health_proto_goTypes = []any{
	(ServiceStatus)(0),               // 0: metalstack.api.v2.ServiceStatus
	(Service)(0),                     // 1: metalstack.api.v2.Service
	(*Health)(nil),                   // 2: metalstack.api.v2.Health
	(*HealthStatus)(nil),             // 3: metalstack.api.v2.HealthStatus
	(*PartitionHealth)(nil),          // 4: metalstack.api.v2.PartitionHealth
	(*HealthServiceGetRequest)(nil),  // 5: metalstack.api.v2.HealthServiceGetRequest
	(*HealthServiceGetResponse)(nil), // 6: metalstack.api.v2.HealthServiceGetResponse
	nil,                              // 7: metalstack.api.v2.HealthStatus.PartitionsEntry
}
var file_metalstack_api_v2_health_proto_depIdxs = []int32{
	3, // 0: metalstack.api.v2.Health.services:type_name -> metalstack.api.v2.HealthStatus
	1, // 1: metalstack.api.v2.HealthStatus.name:type_name -> metalstack.api.v2.Service
	0, // 2: metalstack.api.v2.HealthStatus.status:type_name -> metalstack.api.v2.ServiceStatus
	7, // 3: metalstack.api.v2.HealthStatus.partitions:type_name -> metalstack.api.v2.HealthStatus.PartitionsEntry
	0, // 4: metalstack.api.v2.PartitionHealth.status:type_name -> metalstack.api.v2.ServiceStatus
	2, // 5: metalstack.api.v2.HealthServiceGetResponse.health:type_name -> metalstack.api.v2.Health
	4, // 6: metalstack.api.v2.HealthStatus.PartitionsEntry.value:type_name -> metalstack.api.v2.PartitionHealth
	5, // 7: metalstack.api.v2.HealthService.Get:input_type -> metalstack.api.v2.HealthServiceGetRequest
	6, // 8: metalstack.api.v2.HealthService.Get:output_type -> metalstack.api.v2.HealthServiceGetResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_metalstack_api_v2_health_proto_init() }
func file_metalstack_api_v2_health_proto_init() {
	if File_metalstack_api_v2_health_proto != nil {
		return
	}
	file_metalstack_api_v2_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_metalstack_api_v2_health_proto_rawDesc), len(file_metalstack_api_v2_health_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metalstack_api_v2_health_proto_goTypes,
		DependencyIndexes: file_metalstack_api_v2_health_proto_depIdxs,
		EnumInfos:         file_metalstack_api_v2_health_proto_enumTypes,
		MessageInfos:      file_metalstack_api_v2_health_proto_msgTypes,
	}.Build()
	File_metalstack_api_v2_health_proto = out.File
	file_metalstack_api_v2_health_proto_goTypes = nil
	file_metalstack_api_v2_health_proto_depIdxs = nil
}
