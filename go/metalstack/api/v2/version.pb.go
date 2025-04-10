// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
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

const file_metalstack_api_v2_version_proto_rawDesc = "" +
	"\n" +
	"\x1fmetalstack/api/v2/version.proto\x12\x11metalstack.api.v2\x1a\x1emetalstack/api/v2/common.proto\"y\n" +
	"\aVersion\x12\x18\n" +
	"\aversion\x18\x01 \x01(\tR\aversion\x12\x1a\n" +
	"\brevision\x18\x02 \x01(\tR\brevision\x12\x19\n" +
	"\bgit_sha1\x18\x03 \x01(\tR\agitSha1\x12\x1d\n" +
	"\n" +
	"build_date\x18\x04 \x01(\tR\tbuildDate\"\x1a\n" +
	"\x18VersionServiceGetRequest\"Q\n" +
	"\x19VersionServiceGetResponse\x124\n" +
	"\aversion\x18\x01 \x01(\v2\x1a.metalstack.api.v2.VersionR\aversion2|\n" +
	"\x0eVersionService\x12j\n" +
	"\x03Get\x12+.metalstack.api.v2.VersionServiceGetRequest\x1a,.metalstack.api.v2.VersionServiceGetResponse\"\b\xd8\xf3\x18\x01\xe0\xf3\x18\x02B\xc2\x01\n" +
	"\x15com.metalstack.api.v2B\fVersionProtoP\x01Z5github.com/metal-stack/api/go/metalstack/api/v2;apiv2\xa2\x02\x03MAX\xaa\x02\x11Metalstack.Api.V2\xca\x02\x11Metalstack\\Api\\V2\xe2\x02\x1dMetalstack\\Api\\V2\\GPBMetadata\xea\x02\x13Metalstack::Api::V2b\x06proto3"

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
