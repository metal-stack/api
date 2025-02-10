// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: metalstack/admin/v2/tenant.proto

package adminv2

import (
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

// TenantServiceListRequest is the request payload for a tenant list request
type TenantServiceListRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Login of the tenant to list
	Login *string `protobuf:"bytes,1,opt,name=login,proto3,oneof" json:"login,omitempty"`
	// Name of the tenant to list
	Name *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// Email of the tenant to list
	Email *string `protobuf:"bytes,3,opt,name=email,proto3,oneof" json:"email,omitempty"`
	// Paging details for the list request
	Paging        *v2.Paging `protobuf:"bytes,7,opt,name=paging,proto3" json:"paging,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TenantServiceListRequest) Reset() {
	*x = TenantServiceListRequest{}
	mi := &file_metalstack_admin_v2_tenant_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TenantServiceListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantServiceListRequest) ProtoMessage() {}

func (x *TenantServiceListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_admin_v2_tenant_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantServiceListRequest.ProtoReflect.Descriptor instead.
func (*TenantServiceListRequest) Descriptor() ([]byte, []int) {
	return file_metalstack_admin_v2_tenant_proto_rawDescGZIP(), []int{0}
}

func (x *TenantServiceListRequest) GetLogin() string {
	if x != nil && x.Login != nil {
		return *x.Login
	}
	return ""
}

func (x *TenantServiceListRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *TenantServiceListRequest) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *TenantServiceListRequest) GetPaging() *v2.Paging {
	if x != nil {
		return x.Paging
	}
	return nil
}

// TenantServiceListResponse is the response payload for a tenant list request
type TenantServiceListResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Tenants are the list of tenants
	Tenants []*v2.Tenant `protobuf:"bytes,1,rep,name=tenants,proto3" json:"tenants,omitempty"`
	// NextPage is used for pagination, returns the next page to be fetched and must then be provided in the list request.
	NextPage      *uint64 `protobuf:"varint,2,opt,name=next_page,json=nextPage,proto3,oneof" json:"next_page,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TenantServiceListResponse) Reset() {
	*x = TenantServiceListResponse{}
	mi := &file_metalstack_admin_v2_tenant_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TenantServiceListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantServiceListResponse) ProtoMessage() {}

func (x *TenantServiceListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_admin_v2_tenant_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantServiceListResponse.ProtoReflect.Descriptor instead.
func (*TenantServiceListResponse) Descriptor() ([]byte, []int) {
	return file_metalstack_admin_v2_tenant_proto_rawDescGZIP(), []int{1}
}

func (x *TenantServiceListResponse) GetTenants() []*v2.Tenant {
	if x != nil {
		return x.Tenants
	}
	return nil
}

func (x *TenantServiceListResponse) GetNextPage() uint64 {
	if x != nil && x.NextPage != nil {
		return *x.NextPage
	}
	return 0
}

var File_metalstack_admin_v2_tenant_proto protoreflect.FileDescriptor

var file_metalstack_admin_v2_tenant_proto_rawDesc = string([]byte{
	0x0a, 0x20, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x32, 0x1a, 0x1e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x18, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x88, 0x01, 0x01, 0x12,
	0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x06,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x22, 0x80, 0x01, 0x0a, 0x19, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x07, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x08, 0x6e, 0x65, 0x78,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x65, 0x78,
	0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x32, 0x7e, 0x0a, 0x0d, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x2d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06,
	0xd2, 0xf3, 0x18, 0x02, 0x01, 0x02, 0x42, 0xcf, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x76, 0x32, 0x42, 0x0b, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x6c, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2f, 0x76, 0x32, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x4d,
	0x41, 0x58, 0xaa, 0x02, 0x13, 0x4d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x13, 0x4d, 0x65, 0x74, 0x61, 0x6c,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x56, 0x32, 0xe2, 0x02,
	0x1f, 0x4d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x5c, 0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x15, 0x4d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x3a, 0x3a, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_metalstack_admin_v2_tenant_proto_rawDescOnce sync.Once
	file_metalstack_admin_v2_tenant_proto_rawDescData []byte
)

func file_metalstack_admin_v2_tenant_proto_rawDescGZIP() []byte {
	file_metalstack_admin_v2_tenant_proto_rawDescOnce.Do(func() {
		file_metalstack_admin_v2_tenant_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_metalstack_admin_v2_tenant_proto_rawDesc), len(file_metalstack_admin_v2_tenant_proto_rawDesc)))
	})
	return file_metalstack_admin_v2_tenant_proto_rawDescData
}

var file_metalstack_admin_v2_tenant_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_metalstack_admin_v2_tenant_proto_goTypes = []any{
	(*TenantServiceListRequest)(nil),  // 0: metalstack.admin.v2.TenantServiceListRequest
	(*TenantServiceListResponse)(nil), // 1: metalstack.admin.v2.TenantServiceListResponse
	(*v2.Paging)(nil),                 // 2: metalstack.api.v2.Paging
	(*v2.Tenant)(nil),                 // 3: metalstack.api.v2.Tenant
}
var file_metalstack_admin_v2_tenant_proto_depIdxs = []int32{
	2, // 0: metalstack.admin.v2.TenantServiceListRequest.paging:type_name -> metalstack.api.v2.Paging
	3, // 1: metalstack.admin.v2.TenantServiceListResponse.tenants:type_name -> metalstack.api.v2.Tenant
	0, // 2: metalstack.admin.v2.TenantService.List:input_type -> metalstack.admin.v2.TenantServiceListRequest
	1, // 3: metalstack.admin.v2.TenantService.List:output_type -> metalstack.admin.v2.TenantServiceListResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_metalstack_admin_v2_tenant_proto_init() }
func file_metalstack_admin_v2_tenant_proto_init() {
	if File_metalstack_admin_v2_tenant_proto != nil {
		return
	}
	file_metalstack_admin_v2_tenant_proto_msgTypes[0].OneofWrappers = []any{}
	file_metalstack_admin_v2_tenant_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_metalstack_admin_v2_tenant_proto_rawDesc), len(file_metalstack_admin_v2_tenant_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metalstack_admin_v2_tenant_proto_goTypes,
		DependencyIndexes: file_metalstack_admin_v2_tenant_proto_depIdxs,
		MessageInfos:      file_metalstack_admin_v2_tenant_proto_msgTypes,
	}.Build()
	File_metalstack_admin_v2_tenant_proto = out.File
	file_metalstack_admin_v2_tenant_proto_goTypes = nil
	file_metalstack_admin_v2_tenant_proto_depIdxs = nil
}
