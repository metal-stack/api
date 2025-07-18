// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: metalstack/admin/v2/size.proto

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

// SizeServiceCreateRequest is the request payload for a size create request
type SizeServiceCreateRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Size is the size to create
	Size          *v2.Size `protobuf:"bytes,1,opt,name=size,proto3" json:"size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SizeServiceCreateRequest) Reset() {
	*x = SizeServiceCreateRequest{}
	mi := &file_metalstack_admin_v2_size_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SizeServiceCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SizeServiceCreateRequest) ProtoMessage() {}

func (x *SizeServiceCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_admin_v2_size_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SizeServiceCreateRequest.ProtoReflect.Descriptor instead.
func (*SizeServiceCreateRequest) Descriptor() ([]byte, []int) {
	return file_metalstack_admin_v2_size_proto_rawDescGZIP(), []int{0}
}

func (x *SizeServiceCreateRequest) GetSize() *v2.Size {
	if x != nil {
		return x.Size
	}
	return nil
}

// SizeServiceGetResponse is the response payload for a size create request
type SizeServiceCreateResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Size the size
	Size          *v2.Size `protobuf:"bytes,1,opt,name=size,proto3" json:"size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SizeServiceCreateResponse) Reset() {
	*x = SizeServiceCreateResponse{}
	mi := &file_metalstack_admin_v2_size_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SizeServiceCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SizeServiceCreateResponse) ProtoMessage() {}

func (x *SizeServiceCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_admin_v2_size_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SizeServiceCreateResponse.ProtoReflect.Descriptor instead.
func (*SizeServiceCreateResponse) Descriptor() ([]byte, []int) {
	return file_metalstack_admin_v2_size_proto_rawDescGZIP(), []int{1}
}

func (x *SizeServiceCreateResponse) GetSize() *v2.Size {
	if x != nil {
		return x.Size
	}
	return nil
}

// SizeServiceUpdateRequest is the request payload for a size update request
type SizeServiceUpdateRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Id of this size
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of this size
	Name *string `protobuf:"bytes,4,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// Description of this size
	Description *string `protobuf:"bytes,5,opt,name=description,proto3,oneof" json:"description,omitempty"`
	// Constraints which must match that a specific machine is considered of this size
	Constraints []*v2.SizeConstraint `protobuf:"bytes,6,rep,name=constraints,proto3" json:"constraints,omitempty"`
	// Labels to update on this size
	Labels        *v2.UpdateLabels `protobuf:"bytes,7,opt,name=labels,proto3,oneof" json:"labels,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SizeServiceUpdateRequest) Reset() {
	*x = SizeServiceUpdateRequest{}
	mi := &file_metalstack_admin_v2_size_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SizeServiceUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SizeServiceUpdateRequest) ProtoMessage() {}

func (x *SizeServiceUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_admin_v2_size_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SizeServiceUpdateRequest.ProtoReflect.Descriptor instead.
func (*SizeServiceUpdateRequest) Descriptor() ([]byte, []int) {
	return file_metalstack_admin_v2_size_proto_rawDescGZIP(), []int{2}
}

func (x *SizeServiceUpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SizeServiceUpdateRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *SizeServiceUpdateRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *SizeServiceUpdateRequest) GetConstraints() []*v2.SizeConstraint {
	if x != nil {
		return x.Constraints
	}
	return nil
}

func (x *SizeServiceUpdateRequest) GetLabels() *v2.UpdateLabels {
	if x != nil {
		return x.Labels
	}
	return nil
}

// SizeServiceUpdateResponse is the response payload for a size update request
type SizeServiceUpdateResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Size the size
	Size          *v2.Size `protobuf:"bytes,1,opt,name=size,proto3" json:"size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SizeServiceUpdateResponse) Reset() {
	*x = SizeServiceUpdateResponse{}
	mi := &file_metalstack_admin_v2_size_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SizeServiceUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SizeServiceUpdateResponse) ProtoMessage() {}

func (x *SizeServiceUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_admin_v2_size_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SizeServiceUpdateResponse.ProtoReflect.Descriptor instead.
func (*SizeServiceUpdateResponse) Descriptor() ([]byte, []int) {
	return file_metalstack_admin_v2_size_proto_rawDescGZIP(), []int{3}
}

func (x *SizeServiceUpdateResponse) GetSize() *v2.Size {
	if x != nil {
		return x.Size
	}
	return nil
}

// SizeServiceDeleteRequest is the request payload for a size delete request
type SizeServiceDeleteRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the size to delete
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SizeServiceDeleteRequest) Reset() {
	*x = SizeServiceDeleteRequest{}
	mi := &file_metalstack_admin_v2_size_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SizeServiceDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SizeServiceDeleteRequest) ProtoMessage() {}

func (x *SizeServiceDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_admin_v2_size_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SizeServiceDeleteRequest.ProtoReflect.Descriptor instead.
func (*SizeServiceDeleteRequest) Descriptor() ([]byte, []int) {
	return file_metalstack_admin_v2_size_proto_rawDescGZIP(), []int{4}
}

func (x *SizeServiceDeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// SizeServiceDeleteResponse is the response payload for a size delete request
type SizeServiceDeleteResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Size the size
	Size          *v2.Size `protobuf:"bytes,1,opt,name=size,proto3" json:"size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SizeServiceDeleteResponse) Reset() {
	*x = SizeServiceDeleteResponse{}
	mi := &file_metalstack_admin_v2_size_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SizeServiceDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SizeServiceDeleteResponse) ProtoMessage() {}

func (x *SizeServiceDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_admin_v2_size_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SizeServiceDeleteResponse.ProtoReflect.Descriptor instead.
func (*SizeServiceDeleteResponse) Descriptor() ([]byte, []int) {
	return file_metalstack_admin_v2_size_proto_rawDescGZIP(), []int{5}
}

func (x *SizeServiceDeleteResponse) GetSize() *v2.Size {
	if x != nil {
		return x.Size
	}
	return nil
}

var File_metalstack_admin_v2_size_proto protoreflect.FileDescriptor

const file_metalstack_admin_v2_size_proto_rawDesc = "" +
	"\n" +
	"\x1emetalstack/admin/v2/size.proto\x12\x13metalstack.admin.v2\x1a\x1bbuf/validate/validate.proto\x1a\x1emetalstack/api/v2/common.proto\x1a\x1cmetalstack/api/v2/size.proto\"G\n" +
	"\x18SizeServiceCreateRequest\x12+\n" +
	"\x04size\x18\x01 \x01(\v2\x17.metalstack.api.v2.SizeR\x04size\"H\n" +
	"\x19SizeServiceCreateResponse\x12+\n" +
	"\x04size\x18\x01 \x01(\v2\x17.metalstack.api.v2.SizeR\x04size\"\xb5\x02\n" +
	"\x18SizeServiceUpdateRequest\x12\x1a\n" +
	"\x02id\x18\x01 \x01(\tB\n" +
	"\xbaH\ar\x05\x10\x02\x18\x80\x01R\x02id\x12#\n" +
	"\x04name\x18\x04 \x01(\tB\n" +
	"\xbaH\ar\x05\x10\x02\x18\x80\x01H\x00R\x04name\x88\x01\x01\x121\n" +
	"\vdescription\x18\x05 \x01(\tB\n" +
	"\xbaH\ar\x05\x10\x02\x18\x80\x01H\x01R\vdescription\x88\x01\x01\x12C\n" +
	"\vconstraints\x18\x06 \x03(\v2!.metalstack.api.v2.SizeConstraintR\vconstraints\x12<\n" +
	"\x06labels\x18\a \x01(\v2\x1f.metalstack.api.v2.UpdateLabelsH\x02R\x06labels\x88\x01\x01B\a\n" +
	"\x05_nameB\x0e\n" +
	"\f_descriptionB\t\n" +
	"\a_labels\"H\n" +
	"\x19SizeServiceUpdateResponse\x12+\n" +
	"\x04size\x18\x01 \x01(\v2\x17.metalstack.api.v2.SizeR\x04size\"6\n" +
	"\x18SizeServiceDeleteRequest\x12\x1a\n" +
	"\x02id\x18\x01 \x01(\tB\n" +
	"\xbaH\ar\x05\x10\x02\x18\x80\x01R\x02id\"H\n" +
	"\x19SizeServiceDeleteResponse\x12+\n" +
	"\x04size\x18\x01 \x01(\v2\x17.metalstack.api.v2.SizeR\x04size2\xe0\x02\n" +
	"\vSizeService\x12o\n" +
	"\x06Create\x12-.metalstack.admin.v2.SizeServiceCreateRequest\x1a..metalstack.admin.v2.SizeServiceCreateResponse\"\x06\xd2\xf3\x18\x02\x01\x02\x12o\n" +
	"\x06Update\x12-.metalstack.admin.v2.SizeServiceUpdateRequest\x1a..metalstack.admin.v2.SizeServiceUpdateResponse\"\x06\xd2\xf3\x18\x02\x01\x02\x12o\n" +
	"\x06Delete\x12-.metalstack.admin.v2.SizeServiceDeleteRequest\x1a..metalstack.admin.v2.SizeServiceDeleteResponse\"\x06\xd2\xf3\x18\x02\x01\x02B\xcd\x01\n" +
	"\x17com.metalstack.admin.v2B\tSizeProtoP\x01Z9github.com/metal-stack/api/go/metalstack/admin/v2;adminv2\xa2\x02\x03MAX\xaa\x02\x13Metalstack.Admin.V2\xca\x02\x13Metalstack\\Admin\\V2\xe2\x02\x1fMetalstack\\Admin\\V2\\GPBMetadata\xea\x02\x15Metalstack::Admin::V2b\x06proto3"

var (
	file_metalstack_admin_v2_size_proto_rawDescOnce sync.Once
	file_metalstack_admin_v2_size_proto_rawDescData []byte
)

func file_metalstack_admin_v2_size_proto_rawDescGZIP() []byte {
	file_metalstack_admin_v2_size_proto_rawDescOnce.Do(func() {
		file_metalstack_admin_v2_size_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_metalstack_admin_v2_size_proto_rawDesc), len(file_metalstack_admin_v2_size_proto_rawDesc)))
	})
	return file_metalstack_admin_v2_size_proto_rawDescData
}

var file_metalstack_admin_v2_size_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_metalstack_admin_v2_size_proto_goTypes = []any{
	(*SizeServiceCreateRequest)(nil),  // 0: metalstack.admin.v2.SizeServiceCreateRequest
	(*SizeServiceCreateResponse)(nil), // 1: metalstack.admin.v2.SizeServiceCreateResponse
	(*SizeServiceUpdateRequest)(nil),  // 2: metalstack.admin.v2.SizeServiceUpdateRequest
	(*SizeServiceUpdateResponse)(nil), // 3: metalstack.admin.v2.SizeServiceUpdateResponse
	(*SizeServiceDeleteRequest)(nil),  // 4: metalstack.admin.v2.SizeServiceDeleteRequest
	(*SizeServiceDeleteResponse)(nil), // 5: metalstack.admin.v2.SizeServiceDeleteResponse
	(*v2.Size)(nil),                   // 6: metalstack.api.v2.Size
	(*v2.SizeConstraint)(nil),         // 7: metalstack.api.v2.SizeConstraint
	(*v2.UpdateLabels)(nil),           // 8: metalstack.api.v2.UpdateLabels
}
var file_metalstack_admin_v2_size_proto_depIdxs = []int32{
	6, // 0: metalstack.admin.v2.SizeServiceCreateRequest.size:type_name -> metalstack.api.v2.Size
	6, // 1: metalstack.admin.v2.SizeServiceCreateResponse.size:type_name -> metalstack.api.v2.Size
	7, // 2: metalstack.admin.v2.SizeServiceUpdateRequest.constraints:type_name -> metalstack.api.v2.SizeConstraint
	8, // 3: metalstack.admin.v2.SizeServiceUpdateRequest.labels:type_name -> metalstack.api.v2.UpdateLabels
	6, // 4: metalstack.admin.v2.SizeServiceUpdateResponse.size:type_name -> metalstack.api.v2.Size
	6, // 5: metalstack.admin.v2.SizeServiceDeleteResponse.size:type_name -> metalstack.api.v2.Size
	0, // 6: metalstack.admin.v2.SizeService.Create:input_type -> metalstack.admin.v2.SizeServiceCreateRequest
	2, // 7: metalstack.admin.v2.SizeService.Update:input_type -> metalstack.admin.v2.SizeServiceUpdateRequest
	4, // 8: metalstack.admin.v2.SizeService.Delete:input_type -> metalstack.admin.v2.SizeServiceDeleteRequest
	1, // 9: metalstack.admin.v2.SizeService.Create:output_type -> metalstack.admin.v2.SizeServiceCreateResponse
	3, // 10: metalstack.admin.v2.SizeService.Update:output_type -> metalstack.admin.v2.SizeServiceUpdateResponse
	5, // 11: metalstack.admin.v2.SizeService.Delete:output_type -> metalstack.admin.v2.SizeServiceDeleteResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_metalstack_admin_v2_size_proto_init() }
func file_metalstack_admin_v2_size_proto_init() {
	if File_metalstack_admin_v2_size_proto != nil {
		return
	}
	file_metalstack_admin_v2_size_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_metalstack_admin_v2_size_proto_rawDesc), len(file_metalstack_admin_v2_size_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metalstack_admin_v2_size_proto_goTypes,
		DependencyIndexes: file_metalstack_admin_v2_size_proto_depIdxs,
		MessageInfos:      file_metalstack_admin_v2_size_proto_msgTypes,
	}.Build()
	File_metalstack_admin_v2_size_proto = out.File
	file_metalstack_admin_v2_size_proto_goTypes = nil
	file_metalstack_admin_v2_size_proto_depIdxs = nil
}
