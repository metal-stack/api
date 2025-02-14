// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: metalstack/api/v2/common.proto

package apiv2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
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

// TenantRole specifies what role a logged in user needs to call this tenant scoped service
type TenantRole int32

const (
	// TENANT_ROLE_UNSPECIFIED is not specified
	TenantRole_TENANT_ROLE_UNSPECIFIED TenantRole = 0
	// TENANT_ROLE_OWNER the logged in user needs at least owner role to call this method
	TenantRole_TENANT_ROLE_OWNER TenantRole = 1
	// TENANT_ROLE_EDITOR the logged in user needs at least editor role to call this method
	TenantRole_TENANT_ROLE_EDITOR TenantRole = 2
	// TENANT_ROLE_VIEWER the logged in user needs at least viewer role to call this method
	TenantRole_TENANT_ROLE_VIEWER TenantRole = 3
	// TENANT_ROLE_GUEST the logged in user needs at least guest role to call this method
	// The guest role is assumed by users who are invited to a tenant's project without them
	// having a direct membership within the tenant.
	TenantRole_TENANT_ROLE_GUEST TenantRole = 4
)

// Enum value maps for TenantRole.
var (
	TenantRole_name = map[int32]string{
		0: "TENANT_ROLE_UNSPECIFIED",
		1: "TENANT_ROLE_OWNER",
		2: "TENANT_ROLE_EDITOR",
		3: "TENANT_ROLE_VIEWER",
		4: "TENANT_ROLE_GUEST",
	}
	TenantRole_value = map[string]int32{
		"TENANT_ROLE_UNSPECIFIED": 0,
		"TENANT_ROLE_OWNER":       1,
		"TENANT_ROLE_EDITOR":      2,
		"TENANT_ROLE_VIEWER":      3,
		"TENANT_ROLE_GUEST":       4,
	}
)

func (x TenantRole) Enum() *TenantRole {
	p := new(TenantRole)
	*p = x
	return p
}

func (x TenantRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TenantRole) Descriptor() protoreflect.EnumDescriptor {
	return file_metalstack_api_v2_common_proto_enumTypes[0].Descriptor()
}

func (TenantRole) Type() protoreflect.EnumType {
	return &file_metalstack_api_v2_common_proto_enumTypes[0]
}

func (x TenantRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TenantRole.Descriptor instead.
func (TenantRole) EnumDescriptor() ([]byte, []int) {
	return file_metalstack_api_v2_common_proto_rawDescGZIP(), []int{0}
}

// ProjectRole specifies what role a logged in user needs to call this project scoped service
type ProjectRole int32

const (
	// PROJECT_ROLE_UNSPECIFIED is not specified
	ProjectRole_PROJECT_ROLE_UNSPECIFIED ProjectRole = 0
	// PROJECT_ROLE_OWNER the logged in user needs at least owner role to call this method
	ProjectRole_PROJECT_ROLE_OWNER ProjectRole = 1
	// PROJECT_ROLE_EDITOR the logged in user needs at least editor role to call this method
	ProjectRole_PROJECT_ROLE_EDITOR ProjectRole = 2
	// PROJECT_ROLE_VIEWER the logged in user needs at least viewer role to call this method
	ProjectRole_PROJECT_ROLE_VIEWER ProjectRole = 3
)

// Enum value maps for ProjectRole.
var (
	ProjectRole_name = map[int32]string{
		0: "PROJECT_ROLE_UNSPECIFIED",
		1: "PROJECT_ROLE_OWNER",
		2: "PROJECT_ROLE_EDITOR",
		3: "PROJECT_ROLE_VIEWER",
	}
	ProjectRole_value = map[string]int32{
		"PROJECT_ROLE_UNSPECIFIED": 0,
		"PROJECT_ROLE_OWNER":       1,
		"PROJECT_ROLE_EDITOR":      2,
		"PROJECT_ROLE_VIEWER":      3,
	}
)

func (x ProjectRole) Enum() *ProjectRole {
	p := new(ProjectRole)
	*p = x
	return p
}

func (x ProjectRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProjectRole) Descriptor() protoreflect.EnumDescriptor {
	return file_metalstack_api_v2_common_proto_enumTypes[1].Descriptor()
}

func (ProjectRole) Type() protoreflect.EnumType {
	return &file_metalstack_api_v2_common_proto_enumTypes[1]
}

func (x ProjectRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProjectRole.Descriptor instead.
func (ProjectRole) EnumDescriptor() ([]byte, []int) {
	return file_metalstack_api_v2_common_proto_rawDescGZIP(), []int{1}
}

// AdminRole specifies what role a logged in user needs to call this admin service
type AdminRole int32

const (
	// ADMIN_ROLE_UNSPECIFIED is not specified
	AdminRole_ADMIN_ROLE_UNSPECIFIED AdminRole = 0
	// ADMIN_ROLE_EDITOR the logged in user needs at least editor role to call this method
	AdminRole_ADMIN_ROLE_EDITOR AdminRole = 1
	// ADMIN_ROLE_VIEWER the logged in user needs at least viewer role to call this method
	AdminRole_ADMIN_ROLE_VIEWER AdminRole = 2
)

// Enum value maps for AdminRole.
var (
	AdminRole_name = map[int32]string{
		0: "ADMIN_ROLE_UNSPECIFIED",
		1: "ADMIN_ROLE_EDITOR",
		2: "ADMIN_ROLE_VIEWER",
	}
	AdminRole_value = map[string]int32{
		"ADMIN_ROLE_UNSPECIFIED": 0,
		"ADMIN_ROLE_EDITOR":      1,
		"ADMIN_ROLE_VIEWER":      2,
	}
)

func (x AdminRole) Enum() *AdminRole {
	p := new(AdminRole)
	*p = x
	return p
}

func (x AdminRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdminRole) Descriptor() protoreflect.EnumDescriptor {
	return file_metalstack_api_v2_common_proto_enumTypes[2].Descriptor()
}

func (AdminRole) Type() protoreflect.EnumType {
	return &file_metalstack_api_v2_common_proto_enumTypes[2]
}

func (x AdminRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdminRole.Descriptor instead.
func (AdminRole) EnumDescriptor() ([]byte, []int) {
	return file_metalstack_api_v2_common_proto_rawDescGZIP(), []int{2}
}

// Visibility of a method
type Visibility int32

const (
	// VISIBILITY_UNSPECIFIED is not defined
	Visibility_VISIBILITY_UNSPECIFIED Visibility = 0
	// VISIBILITY_PUBLIC specifies that this service is accessible without authentication
	Visibility_VISIBILITY_PUBLIC Visibility = 1
	// VISIBILITY_SELF enable call this endpoint from the authenticated user only
	Visibility_VISIBILITY_SELF Visibility = 3
)

// Enum value maps for Visibility.
var (
	Visibility_name = map[int32]string{
		0: "VISIBILITY_UNSPECIFIED",
		1: "VISIBILITY_PUBLIC",
		3: "VISIBILITY_SELF",
	}
	Visibility_value = map[string]int32{
		"VISIBILITY_UNSPECIFIED": 0,
		"VISIBILITY_PUBLIC":      1,
		"VISIBILITY_SELF":        3,
	}
)

func (x Visibility) Enum() *Visibility {
	p := new(Visibility)
	*p = x
	return p
}

func (x Visibility) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Visibility) Descriptor() protoreflect.EnumDescriptor {
	return file_metalstack_api_v2_common_proto_enumTypes[3].Descriptor()
}

func (Visibility) Type() protoreflect.EnumType {
	return &file_metalstack_api_v2_common_proto_enumTypes[3]
}

func (x Visibility) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Visibility.Descriptor instead.
func (Visibility) EnumDescriptor() ([]byte, []int) {
	return file_metalstack_api_v2_common_proto_rawDescGZIP(), []int{3}
}

// Auditing option specified per service method
// by default all service methods are included
// add the auditing option if you want to exclude a method in auditing
type Auditing int32

const (
	// AUDITING_UNSPECIFIED is not specified
	Auditing_AUDITING_UNSPECIFIED Auditing = 0
	// AUDITING_INCLUDED if a method is annotated with this, all calls are audited
	Auditing_AUDITING_INCLUDED Auditing = 1
	// AUDITING_EXCLUDED if a method is annotated with this, no calls are audited
	Auditing_AUDITING_EXCLUDED Auditing = 2
)

// Enum value maps for Auditing.
var (
	Auditing_name = map[int32]string{
		0: "AUDITING_UNSPECIFIED",
		1: "AUDITING_INCLUDED",
		2: "AUDITING_EXCLUDED",
	}
	Auditing_value = map[string]int32{
		"AUDITING_UNSPECIFIED": 0,
		"AUDITING_INCLUDED":    1,
		"AUDITING_EXCLUDED":    2,
	}
)

func (x Auditing) Enum() *Auditing {
	p := new(Auditing)
	*p = x
	return p
}

func (x Auditing) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Auditing) Descriptor() protoreflect.EnumDescriptor {
	return file_metalstack_api_v2_common_proto_enumTypes[4].Descriptor()
}

func (Auditing) Type() protoreflect.EnumType {
	return &file_metalstack_api_v2_common_proto_enumTypes[4]
}

func (x Auditing) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Auditing.Descriptor instead.
func (Auditing) EnumDescriptor() ([]byte, []int) {
	return file_metalstack_api_v2_common_proto_rawDescGZIP(), []int{4}
}

// Paging defines paging for methods with a lot of results
type Paging struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Page is used for pagination, if unset only the first page is returned,
	// the list response contains then the page number for the next page.
	Page *uint64 `protobuf:"varint,1,opt,name=page,proto3,oneof" json:"page,omitempty"`
	// Count is the number of results returned per page, if not given server side defaults apply
	Count         *uint64 `protobuf:"varint,2,opt,name=count,proto3,oneof" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Paging) Reset() {
	*x = Paging{}
	mi := &file_metalstack_api_v2_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Paging) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Paging) ProtoMessage() {}

func (x *Paging) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_api_v2_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Paging.ProtoReflect.Descriptor instead.
func (*Paging) Descriptor() ([]byte, []int) {
	return file_metalstack_api_v2_common_proto_rawDescGZIP(), []int{0}
}

func (x *Paging) GetPage() uint64 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *Paging) GetCount() uint64 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

var file_metalstack_api_v2_common_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: ([]TenantRole)(nil),
		Field:         51000,
		Name:          "metalstack.api.v2.tenant_roles",
		Tag:           "varint,51000,rep,packed,name=tenant_roles,enum=metalstack.api.v2.TenantRole",
		Filename:      "metalstack/api/v2/common.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: ([]ProjectRole)(nil),
		Field:         51001,
		Name:          "metalstack.api.v2.project_roles",
		Tag:           "varint,51001,rep,packed,name=project_roles,enum=metalstack.api.v2.ProjectRole",
		Filename:      "metalstack/api/v2/common.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: ([]AdminRole)(nil),
		Field:         51002,
		Name:          "metalstack.api.v2.admin_roles",
		Tag:           "varint,51002,rep,packed,name=admin_roles,enum=metalstack.api.v2.AdminRole",
		Filename:      "metalstack/api/v2/common.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*Visibility)(nil),
		Field:         51003,
		Name:          "metalstack.api.v2.visibility",
		Tag:           "varint,51003,opt,name=visibility,enum=metalstack.api.v2.Visibility",
		Filename:      "metalstack/api/v2/common.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*Auditing)(nil),
		Field:         51004,
		Name:          "metalstack.api.v2.auditing",
		Tag:           "varint,51004,opt,name=auditing,enum=metalstack.api.v2.Auditing",
		Filename:      "metalstack/api/v2/common.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         52000,
		Name:          "metalstack.api.v2.enum_string_value",
		Tag:           "bytes,52000,opt,name=enum_string_value",
		Filename:      "metalstack/api/v2/common.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// TenantRoles are used to define which tenant role a logged in user must provide to call this method
	//
	// repeated metalstack.api.v2.TenantRole tenant_roles = 51000;
	E_TenantRoles = &file_metalstack_api_v2_common_proto_extTypes[0]
	// ProjectRoles are used to define which project role a logged in user must provide to call this method
	//
	// repeated metalstack.api.v2.ProjectRole project_roles = 51001;
	E_ProjectRoles = &file_metalstack_api_v2_common_proto_extTypes[1]
	// AdminRoles are used to define which admin role a logged in user must provide to call this method
	//
	// repeated metalstack.api.v2.AdminRole admin_roles = 51002;
	E_AdminRoles = &file_metalstack_api_v2_common_proto_extTypes[2]
	// Visibility defines the visibility of this method, this is used to have public or self visible methods
	//
	// optional metalstack.api.v2.Visibility visibility = 51003;
	E_Visibility = &file_metalstack_api_v2_common_proto_extTypes[3]
	// Auditing defines if calls to this method should be audited or not
	//
	// optional metalstack.api.v2.Auditing auditing = 51004;
	E_Auditing = &file_metalstack_api_v2_common_proto_extTypes[4]
)

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// StringValue which can be set to a enum
	//
	// optional string enum_string_value = 52000;
	E_EnumStringValue = &file_metalstack_api_v2_common_proto_extTypes[5]
)

var File_metalstack_api_v2_common_proto protoreflect.FileDescriptor

var file_metalstack_api_v2_common_proto_rawDesc = string([]byte{
	0x0a, 0x1e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x32, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x06, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x12,
	0x17, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x48, 0x01, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2a, 0x87, 0x01, 0x0a, 0x0a, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x45, 0x4e, 0x41, 0x4e, 0x54, 0x5f,
	0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x45, 0x4e, 0x41, 0x4e, 0x54, 0x5f, 0x52, 0x4f, 0x4c,
	0x45, 0x5f, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x45, 0x4e,
	0x41, 0x4e, 0x54, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x45, 0x44, 0x49, 0x54, 0x4f, 0x52, 0x10,
	0x02, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x45, 0x4e, 0x41, 0x4e, 0x54, 0x5f, 0x52, 0x4f, 0x4c, 0x45,
	0x5f, 0x56, 0x49, 0x45, 0x57, 0x45, 0x52, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x45, 0x4e,
	0x41, 0x4e, 0x54, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x47, 0x55, 0x45, 0x53, 0x54, 0x10, 0x04,
	0x2a, 0x75, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x12,
	0x1c, 0x0a, 0x18, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a,
	0x12, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x4f, 0x57,
	0x4e, 0x45, 0x52, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54,
	0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x45, 0x44, 0x49, 0x54, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x17,
	0x0a, 0x13, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x56,
	0x49, 0x45, 0x57, 0x45, 0x52, 0x10, 0x03, 0x2a, 0x55, 0x0a, 0x09, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x5f, 0x52, 0x4f,
	0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x15, 0x0a, 0x11, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x45,
	0x44, 0x49, 0x54, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x44, 0x4d, 0x49, 0x4e,
	0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x45, 0x52, 0x10, 0x02, 0x2a, 0x54,
	0x0a, 0x0a, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x16,
	0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x56, 0x49, 0x53, 0x49,
	0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x10, 0x01, 0x12,
	0x13, 0x0a, 0x0f, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x53, 0x45,
	0x4c, 0x46, 0x10, 0x03, 0x2a, 0x52, 0x0a, 0x08, 0x41, 0x75, 0x64, 0x69, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0x18, 0x0a, 0x14, 0x41, 0x55, 0x44, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x55,
	0x44, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x49, 0x4e, 0x43, 0x4c, 0x55, 0x44, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x55, 0x44, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x45, 0x58,
	0x43, 0x4c, 0x55, 0x44, 0x45, 0x44, 0x10, 0x02, 0x3a, 0x62, 0x0a, 0x0c, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb8, 0x8e, 0x03, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x1d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x0b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x3a, 0x65, 0x0a, 0x0d,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x1e, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb9, 0x8e,
	0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x6f,
	0x6c, 0x65, 0x73, 0x3a, 0x5f, 0x0a, 0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x72, 0x6f, 0x6c,
	0x65, 0x73, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xba, 0x8e, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x0a, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x52,
	0x6f, 0x6c, 0x65, 0x73, 0x3a, 0x5f, 0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xbb, 0x8e, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x56,
	0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x3a, 0x59, 0x0a, 0x08, 0x61, 0x75, 0x64, 0x69, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xbc, 0x8e, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x61, 0x75, 0x64, 0x69, 0x74, 0x69, 0x6e, 0x67,
	0x3a, 0x4f, 0x0a, 0x11, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa0, 0x96, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x65, 0x6e, 0x75, 0x6d, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0xc1, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x42, 0x0b, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2d, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x3b, 0x61, 0x70, 0x69, 0x76,
	0x32, 0xa2, 0x02, 0x03, 0x4d, 0x41, 0x58, 0xaa, 0x02, 0x11, 0x4d, 0x65, 0x74, 0x61, 0x6c, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x11, 0x4d, 0x65,
	0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x32, 0xe2,
	0x02, 0x1d, 0x4d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x41, 0x70, 0x69,
	0x5c, 0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x13, 0x4d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x3a, 0x3a, 0x41, 0x70,
	0x69, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_metalstack_api_v2_common_proto_rawDescOnce sync.Once
	file_metalstack_api_v2_common_proto_rawDescData []byte
)

func file_metalstack_api_v2_common_proto_rawDescGZIP() []byte {
	file_metalstack_api_v2_common_proto_rawDescOnce.Do(func() {
		file_metalstack_api_v2_common_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_metalstack_api_v2_common_proto_rawDesc), len(file_metalstack_api_v2_common_proto_rawDesc)))
	})
	return file_metalstack_api_v2_common_proto_rawDescData
}

var file_metalstack_api_v2_common_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_metalstack_api_v2_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_metalstack_api_v2_common_proto_goTypes = []any{
	(TenantRole)(0),                       // 0: metalstack.api.v2.TenantRole
	(ProjectRole)(0),                      // 1: metalstack.api.v2.ProjectRole
	(AdminRole)(0),                        // 2: metalstack.api.v2.AdminRole
	(Visibility)(0),                       // 3: metalstack.api.v2.Visibility
	(Auditing)(0),                         // 4: metalstack.api.v2.Auditing
	(*Paging)(nil),                        // 5: metalstack.api.v2.Paging
	(*descriptorpb.MethodOptions)(nil),    // 6: google.protobuf.MethodOptions
	(*descriptorpb.EnumValueOptions)(nil), // 7: google.protobuf.EnumValueOptions
}
var file_metalstack_api_v2_common_proto_depIdxs = []int32{
	6,  // 0: metalstack.api.v2.tenant_roles:extendee -> google.protobuf.MethodOptions
	6,  // 1: metalstack.api.v2.project_roles:extendee -> google.protobuf.MethodOptions
	6,  // 2: metalstack.api.v2.admin_roles:extendee -> google.protobuf.MethodOptions
	6,  // 3: metalstack.api.v2.visibility:extendee -> google.protobuf.MethodOptions
	6,  // 4: metalstack.api.v2.auditing:extendee -> google.protobuf.MethodOptions
	7,  // 5: metalstack.api.v2.enum_string_value:extendee -> google.protobuf.EnumValueOptions
	0,  // 6: metalstack.api.v2.tenant_roles:type_name -> metalstack.api.v2.TenantRole
	1,  // 7: metalstack.api.v2.project_roles:type_name -> metalstack.api.v2.ProjectRole
	2,  // 8: metalstack.api.v2.admin_roles:type_name -> metalstack.api.v2.AdminRole
	3,  // 9: metalstack.api.v2.visibility:type_name -> metalstack.api.v2.Visibility
	4,  // 10: metalstack.api.v2.auditing:type_name -> metalstack.api.v2.Auditing
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	6,  // [6:11] is the sub-list for extension type_name
	0,  // [0:6] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_metalstack_api_v2_common_proto_init() }
func file_metalstack_api_v2_common_proto_init() {
	if File_metalstack_api_v2_common_proto != nil {
		return
	}
	file_metalstack_api_v2_common_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_metalstack_api_v2_common_proto_rawDesc), len(file_metalstack_api_v2_common_proto_rawDesc)),
			NumEnums:      5,
			NumMessages:   1,
			NumExtensions: 6,
			NumServices:   0,
		},
		GoTypes:           file_metalstack_api_v2_common_proto_goTypes,
		DependencyIndexes: file_metalstack_api_v2_common_proto_depIdxs,
		EnumInfos:         file_metalstack_api_v2_common_proto_enumTypes,
		MessageInfos:      file_metalstack_api_v2_common_proto_msgTypes,
		ExtensionInfos:    file_metalstack_api_v2_common_proto_extTypes,
	}.Build()
	File_metalstack_api_v2_common_proto = out.File
	file_metalstack_api_v2_common_proto_goTypes = nil
	file_metalstack_api_v2_common_proto_depIdxs = nil
}
