// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: metalstack/api/v2/version.proto

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

// Version of the application
type Version struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Version of the application
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// Revision of the application
	Revision string `protobuf:"bytes,2,opt,name=revision,proto3" json:"revision,omitempty"`
	// GitSHA1 of the application
	GitSha1 string `protobuf:"bytes,3,opt,name=git_sha1,json=gitSha1,proto3" json:"git_sha1,omitempty"`
	// BuildDate of the application
	BuildDate     string `protobuf:"bytes,4,opt,name=build_date,json=buildDate,proto3" json:"build_date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Version) Reset() {
	*x = Version{}
	mi := &file_metalstack_api_v2_version_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_api_v2_version_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_metalstack_api_v2_version_proto_rawDescGZIP(), []int{0}
}

func (x *Version) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Version) GetRevision() string {
	if x != nil {
		return x.Revision
	}
	return ""
}

func (x *Version) GetGitSha1() string {
	if x != nil {
		return x.GitSha1
	}
	return ""
}

func (x *Version) GetBuildDate() string {
	if x != nil {
		return x.BuildDate
	}
	return ""
}

// VersionServiceGetRequest is the request payload to get the version
type VersionServiceGetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VersionServiceGetRequest) Reset() {
	*x = VersionServiceGetRequest{}
	mi := &file_metalstack_api_v2_version_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VersionServiceGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionServiceGetRequest) ProtoMessage() {}

func (x *VersionServiceGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_api_v2_version_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionServiceGetRequest.ProtoReflect.Descriptor instead.
func (*VersionServiceGetRequest) Descriptor() ([]byte, []int) {
	return file_metalstack_api_v2_version_proto_rawDescGZIP(), []int{1}
}

// VersionServiceGetResponse is the response payload with the version
type VersionServiceGetResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Version of the application
	Version       *Version `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VersionServiceGetResponse) Reset() {
	*x = VersionServiceGetResponse{}
	mi := &file_metalstack_api_v2_version_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VersionServiceGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionServiceGetResponse) ProtoMessage() {}

func (x *VersionServiceGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_api_v2_version_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionServiceGetResponse.ProtoReflect.Descriptor instead.
func (*VersionServiceGetResponse) Descriptor() ([]byte, []int) {
	return file_metalstack_api_v2_version_proto_rawDescGZIP(), []int{2}
}

func (x *VersionServiceGetResponse) GetVersion() *Version {
	if x != nil {
		return x.Version
	}
	return nil
}

var File_metalstack_api_v2_version_proto protoreflect.FileDescriptor

var file_metalstack_api_v2_version_proto_rawDesc = string([]byte{
	0x0a, 0x1f, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x32, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x32, 0x1a, 0x1e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x69, 0x74, 0x5f, 0x73, 0x68, 0x61,
	0x31, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x69, 0x74, 0x53, 0x68, 0x61, 0x31,
	0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22,
	0x1a, 0x0a, 0x18, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x51, 0x0a, 0x19, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x7c,
	0x0a, 0x0e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x6a, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x2b, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x08, 0xd8, 0xf3, 0x18, 0x01, 0xe0, 0xf3, 0x18, 0x02, 0x42, 0xc2, 0x01, 0x0a,
	0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x42, 0x0c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x3b, 0x61, 0x70, 0x69, 0x76, 0x32, 0xa2, 0x02, 0x03,
	0x4d, 0x41, 0x58, 0xaa, 0x02, 0x11, 0x4d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x2e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x11, 0x4d, 0x65, 0x74, 0x61, 0x6c, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x32, 0xe2, 0x02, 0x1d, 0x4d, 0x65,
	0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x32, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x4d, 0x65,
	0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56,
	0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_metalstack_api_v2_version_proto_rawDescOnce sync.Once
	file_metalstack_api_v2_version_proto_rawDescData []byte
)

func file_metalstack_api_v2_version_proto_rawDescGZIP() []byte {
	file_metalstack_api_v2_version_proto_rawDescOnce.Do(func() {
		file_metalstack_api_v2_version_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_metalstack_api_v2_version_proto_rawDesc), len(file_metalstack_api_v2_version_proto_rawDesc)))
	})
	return file_metalstack_api_v2_version_proto_rawDescData
}

var file_metalstack_api_v2_version_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_metalstack_api_v2_version_proto_goTypes = []any{
	(*Version)(nil),                   // 0: metalstack.api.v2.Version
	(*VersionServiceGetRequest)(nil),  // 1: metalstack.api.v2.VersionServiceGetRequest
	(*VersionServiceGetResponse)(nil), // 2: metalstack.api.v2.VersionServiceGetResponse
}
var file_metalstack_api_v2_version_proto_depIdxs = []int32{
	0, // 0: metalstack.api.v2.VersionServiceGetResponse.version:type_name -> metalstack.api.v2.Version
	1, // 1: metalstack.api.v2.VersionService.Get:input_type -> metalstack.api.v2.VersionServiceGetRequest
	2, // 2: metalstack.api.v2.VersionService.Get:output_type -> metalstack.api.v2.VersionServiceGetResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_metalstack_api_v2_version_proto_init() }
func file_metalstack_api_v2_version_proto_init() {
	if File_metalstack_api_v2_version_proto != nil {
		return
	}
	file_metalstack_api_v2_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_metalstack_api_v2_version_proto_rawDesc), len(file_metalstack_api_v2_version_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metalstack_api_v2_version_proto_goTypes,
		DependencyIndexes: file_metalstack_api_v2_version_proto_depIdxs,
		MessageInfos:      file_metalstack_api_v2_version_proto_msgTypes,
	}.Build()
	File_metalstack_api_v2_version_proto = out.File
	file_metalstack_api_v2_version_proto_goTypes = nil
	file_metalstack_api_v2_version_proto_depIdxs = nil
}
