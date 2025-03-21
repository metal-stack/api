// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: metalstack/api/v2/image.proto

package apiv2

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// ImageFeature
type ImageFeature int32

const (
	// IMAGE_FEATURE_UNSPECIFIED is not specified
	ImageFeature_IMAGE_FEATURE_UNSPECIFIED ImageFeature = 0
	// IMAGE_FEATURE_MACHINE indicates this image is usable for a machine
	ImageFeature_IMAGE_FEATURE_MACHINE ImageFeature = 1
	// IMAGE_FEATURE_FIREWALL indicates this image is usable for a firewall
	ImageFeature_IMAGE_FEATURE_FIREWALL ImageFeature = 2
)

// Enum value maps for ImageFeature.
var (
	ImageFeature_name = map[int32]string{
		0: "IMAGE_FEATURE_UNSPECIFIED",
		1: "IMAGE_FEATURE_MACHINE",
		2: "IMAGE_FEATURE_FIREWALL",
	}
	ImageFeature_value = map[string]int32{
		"IMAGE_FEATURE_UNSPECIFIED": 0,
		"IMAGE_FEATURE_MACHINE":     1,
		"IMAGE_FEATURE_FIREWALL":    2,
	}
)

func (x ImageFeature) Enum() *ImageFeature {
	p := new(ImageFeature)
	*p = x
	return p
}

func (x ImageFeature) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ImageFeature) Descriptor() protoreflect.EnumDescriptor {
	return file_metalstack_api_v2_image_proto_enumTypes[0].Descriptor()
}

func (ImageFeature) Type() protoreflect.EnumType {
	return &file_metalstack_api_v2_image_proto_enumTypes[0]
}

func (x ImageFeature) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ImageFeature.Descriptor instead.
func (ImageFeature) EnumDescriptor() ([]byte, []int) {
	return file_metalstack_api_v2_image_proto_rawDescGZIP(), []int{0}
}

// Image
type ImageClassification int32

const (
	// IMAGE_CLASSIFICATION_UNSPECIFIED is not specified
	ImageClassification_IMAGE_CLASSIFICATION_UNSPECIFIED ImageClassification = 0
	// IMAGE_CLASSIFICATION_PREVIEW indicates that this image is in preview
	ImageClassification_IMAGE_CLASSIFICATION_PREVIEW ImageClassification = 1
	// IMAGE_CLASSIFICATION_SUPPORTED indicates that this image is supported
	ImageClassification_IMAGE_CLASSIFICATION_SUPPORTED ImageClassification = 2
	// IMAGE_CLASSIFICATION_DEPRECATED indicates that this image is deprecated
	ImageClassification_IMAGE_CLASSIFICATION_DEPRECATED ImageClassification = 3
)

// Enum value maps for ImageClassification.
var (
	ImageClassification_name = map[int32]string{
		0: "IMAGE_CLASSIFICATION_UNSPECIFIED",
		1: "IMAGE_CLASSIFICATION_PREVIEW",
		2: "IMAGE_CLASSIFICATION_SUPPORTED",
		3: "IMAGE_CLASSIFICATION_DEPRECATED",
	}
	ImageClassification_value = map[string]int32{
		"IMAGE_CLASSIFICATION_UNSPECIFIED": 0,
		"IMAGE_CLASSIFICATION_PREVIEW":     1,
		"IMAGE_CLASSIFICATION_SUPPORTED":   2,
		"IMAGE_CLASSIFICATION_DEPRECATED":  3,
	}
)

func (x ImageClassification) Enum() *ImageClassification {
	p := new(ImageClassification)
	*p = x
	return p
}

func (x ImageClassification) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ImageClassification) Descriptor() protoreflect.EnumDescriptor {
	return file_metalstack_api_v2_image_proto_enumTypes[1].Descriptor()
}

func (ImageClassification) Type() protoreflect.EnumType {
	return &file_metalstack_api_v2_image_proto_enumTypes[1]
}

func (x ImageClassification) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ImageClassification.Descriptor instead.
func (ImageClassification) EnumDescriptor() ([]byte, []int) {
	return file_metalstack_api_v2_image_proto_rawDescGZIP(), []int{1}
}

// ImageServiceGetRequest is the request payload for a image get request
type ImageServiceGetRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the image to get
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImageServiceGetRequest) Reset() {
	*x = ImageServiceGetRequest{}
	mi := &file_metalstack_api_v2_image_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageServiceGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageServiceGetRequest) ProtoMessage() {}

func (x *ImageServiceGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_api_v2_image_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageServiceGetRequest.ProtoReflect.Descriptor instead.
func (*ImageServiceGetRequest) Descriptor() ([]byte, []int) {
	return file_metalstack_api_v2_image_proto_rawDescGZIP(), []int{0}
}

func (x *ImageServiceGetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// ImageServiceListRequest is the request payload for a image list request
type ImageServiceListRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Query for images
	Query         *ImageQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImageServiceListRequest) Reset() {
	*x = ImageServiceListRequest{}
	mi := &file_metalstack_api_v2_image_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageServiceListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageServiceListRequest) ProtoMessage() {}

func (x *ImageServiceListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_api_v2_image_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageServiceListRequest.ProtoReflect.Descriptor instead.
func (*ImageServiceListRequest) Descriptor() ([]byte, []int) {
	return file_metalstack_api_v2_image_proto_rawDescGZIP(), []int{1}
}

func (x *ImageServiceListRequest) GetQuery() *ImageQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

// ImageServiceLatestRequest is the request payload for a image latest request
type ImageServiceLatestRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// OS for which the latest image should be fetched
	// should contain os and major.minor then latest patch version of this os is returned
	Os            string `protobuf:"bytes,1,opt,name=os,proto3" json:"os,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImageServiceLatestRequest) Reset() {
	*x = ImageServiceLatestRequest{}
	mi := &file_metalstack_api_v2_image_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageServiceLatestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageServiceLatestRequest) ProtoMessage() {}

func (x *ImageServiceLatestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_api_v2_image_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageServiceLatestRequest.ProtoReflect.Descriptor instead.
func (*ImageServiceLatestRequest) Descriptor() ([]byte, []int) {
	return file_metalstack_api_v2_image_proto_rawDescGZIP(), []int{2}
}

func (x *ImageServiceLatestRequest) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

// ImageServiceGetResponse is the response payload for a image get request
type ImageServiceGetResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Image the image
	Image         *Image `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImageServiceGetResponse) Reset() {
	*x = ImageServiceGetResponse{}
	mi := &file_metalstack_api_v2_image_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageServiceGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageServiceGetResponse) ProtoMessage() {}

func (x *ImageServiceGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_api_v2_image_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageServiceGetResponse.ProtoReflect.Descriptor instead.
func (*ImageServiceGetResponse) Descriptor() ([]byte, []int) {
	return file_metalstack_api_v2_image_proto_rawDescGZIP(), []int{3}
}

func (x *ImageServiceGetResponse) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

// ImageServiceListResponse is the response payload for a image list request
type ImageServiceListResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Images the images
	Images        []*Image `protobuf:"bytes,1,rep,name=images,proto3" json:"images,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImageServiceListResponse) Reset() {
	*x = ImageServiceListResponse{}
	mi := &file_metalstack_api_v2_image_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageServiceListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageServiceListResponse) ProtoMessage() {}

func (x *ImageServiceListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_api_v2_image_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageServiceListResponse.ProtoReflect.Descriptor instead.
func (*ImageServiceListResponse) Descriptor() ([]byte, []int) {
	return file_metalstack_api_v2_image_proto_rawDescGZIP(), []int{4}
}

func (x *ImageServiceListResponse) GetImages() []*Image {
	if x != nil {
		return x.Images
	}
	return nil
}

// ImageServiceLatestResponse is the response payload for a image latest request
type ImageServiceLatestResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Image which is the latest for one os
	Image         *Image `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImageServiceLatestResponse) Reset() {
	*x = ImageServiceLatestResponse{}
	mi := &file_metalstack_api_v2_image_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageServiceLatestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageServiceLatestResponse) ProtoMessage() {}

func (x *ImageServiceLatestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_api_v2_image_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageServiceLatestResponse.ProtoReflect.Descriptor instead.
func (*ImageServiceLatestResponse) Descriptor() ([]byte, []int) {
	return file_metalstack_api_v2_image_proto_rawDescGZIP(), []int{5}
}

func (x *ImageServiceLatestResponse) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

// Image
type Image struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Id of this imageLayout
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Meta for this ip
	Meta *Meta `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
	// URL where this image is located
	Url string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	// Name of this imageLayout
	Name *string `protobuf:"bytes,4,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// Description of this imageLayout
	Description *string `protobuf:"bytes,5,opt,name=description,proto3,oneof" json:"description,omitempty"`
	// Features of this image
	Features []ImageFeature `protobuf:"varint,6,rep,packed,name=features,proto3,enum=metalstack.api.v2.ImageFeature" json:"features,omitempty"`
	// Classification of this image
	Classification ImageClassification `protobuf:"varint,7,opt,name=classification,proto3,enum=metalstack.api.v2.ImageClassification" json:"classification,omitempty"`
	// ExpiresAt usage is not possible after this date
	ExpiresAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Image) Reset() {
	*x = Image{}
	mi := &file_metalstack_api_v2_image_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_api_v2_image_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_metalstack_api_v2_image_proto_rawDescGZIP(), []int{6}
}

func (x *Image) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Image) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Image) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Image) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Image) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *Image) GetFeatures() []ImageFeature {
	if x != nil {
		return x.Features
	}
	return nil
}

func (x *Image) GetClassification() ImageClassification {
	if x != nil {
		return x.Classification
	}
	return ImageClassification_IMAGE_CLASSIFICATION_UNSPECIFIED
}

func (x *Image) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

// ImageUsage reports which machines/firewalls actually use this image
type ImageUsage struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Image with usage
	Image *Image `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	// UsedBy the following machines/firewalls
	UsedBy        []string `protobuf:"bytes,2,rep,name=used_by,json=usedBy,proto3" json:"used_by,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImageUsage) Reset() {
	*x = ImageUsage{}
	mi := &file_metalstack_api_v2_image_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageUsage) ProtoMessage() {}

func (x *ImageUsage) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_api_v2_image_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageUsage.ProtoReflect.Descriptor instead.
func (*ImageUsage) Descriptor() ([]byte, []int) {
	return file_metalstack_api_v2_image_proto_rawDescGZIP(), []int{7}
}

func (x *ImageUsage) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *ImageUsage) GetUsedBy() []string {
	if x != nil {
		return x.UsedBy
	}
	return nil
}

// ImageQuery is used to search images
type ImageQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the image to get
	Id *string `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// OS of the image
	Os *string `protobuf:"bytes,2,opt,name=os,proto3,oneof" json:"os,omitempty"`
	// Version of the Image
	Version *string `protobuf:"bytes,3,opt,name=version,proto3,oneof" json:"version,omitempty"`
	// Name of the image to query
	Name *string `protobuf:"bytes,4,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// Description of the image to query
	Description *string `protobuf:"bytes,5,opt,name=description,proto3,oneof" json:"description,omitempty"`
	// Feature of the image to query
	Feature *ImageFeature `protobuf:"varint,6,opt,name=feature,proto3,enum=metalstack.api.v2.ImageFeature,oneof" json:"feature,omitempty"`
	// Classification of the image to query
	Classification *ImageClassification `protobuf:"varint,7,opt,name=classification,proto3,enum=metalstack.api.v2.ImageClassification,oneof" json:"classification,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ImageQuery) Reset() {
	*x = ImageQuery{}
	mi := &file_metalstack_api_v2_image_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageQuery) ProtoMessage() {}

func (x *ImageQuery) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_api_v2_image_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageQuery.ProtoReflect.Descriptor instead.
func (*ImageQuery) Descriptor() ([]byte, []int) {
	return file_metalstack_api_v2_image_proto_rawDescGZIP(), []int{8}
}

func (x *ImageQuery) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *ImageQuery) GetOs() string {
	if x != nil && x.Os != nil {
		return *x.Os
	}
	return ""
}

func (x *ImageQuery) GetVersion() string {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return ""
}

func (x *ImageQuery) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ImageQuery) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *ImageQuery) GetFeature() ImageFeature {
	if x != nil && x.Feature != nil {
		return *x.Feature
	}
	return ImageFeature_IMAGE_FEATURE_UNSPECIFIED
}

func (x *ImageQuery) GetClassification() ImageClassification {
	if x != nil && x.Classification != nil {
		return *x.Classification
	}
	return ImageClassification_IMAGE_CLASSIFICATION_UNSPECIFIED
}

var File_metalstack_api_v2_image_proto protoreflect.FileDescriptor

var file_metalstack_api_v2_image_proto_rawDesc = string([]byte{
	0x0a, 0x1d, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x32, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x34, 0x0a, 0x16, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0x48, 0x07, 0x72, 0x05, 0x10, 0x02, 0x18,
	0x80, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4e, 0x0a, 0x17, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x33, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x2b, 0x0a, 0x19, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x6f, 0x73, 0x22, 0x49, 0x0a, 0x17, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e,
	0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x32, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x4c,
	0x0a, 0x18, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0x4c, 0x0a, 0x1a,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x61, 0x74, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0xa8, 0x04, 0x0a, 0x05, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0a, 0xba, 0x48, 0x07, 0x72, 0x05, 0x10, 0x02, 0x18, 0x80, 0x01, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x2b, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x32, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x4a, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x38, 0xba, 0x48, 0x35, 0xba,
	0x01, 0x32, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x12, 0x17, 0x75,
	0x72, 0x6c, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20, 0x61, 0x20, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x20, 0x55, 0x52, 0x49, 0x1a, 0x0c, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x69, 0x73, 0x55,
	0x72, 0x69, 0x28, 0x29, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x23, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0x48, 0x07, 0x72, 0x05, 0x10, 0x02,
	0x18, 0x80, 0x01, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x31,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0x48, 0x07, 0x72, 0x05, 0x10, 0x02, 0x18, 0x80, 0x01, 0x48,
	0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01,
	0x01, 0x12, 0x83, 0x01, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x46, 0xba, 0x48, 0x43, 0x92, 0x01, 0x40, 0x08, 0x01, 0x22,
	0x3c, 0xba, 0x01, 0x39, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x15,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x1a, 0x16, 0x74, 0x68, 0x69, 0x73, 0x20, 0x3e, 0x3d, 0x20, 0x30,
	0x20, 0x26, 0x26, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x3c, 0x3d, 0x20, 0x32, 0x52, 0x08, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x58, 0x0a, 0x0e, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x26, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xba, 0x48, 0x05, 0x82, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x0e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x55, 0x0a, 0x0a, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x64, 0x42, 0x79, 0x22, 0xcc, 0x03, 0x0a,
	0x0a, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0x48, 0x07, 0x72, 0x05, 0x10, 0x02,
	0x18, 0x80, 0x01, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x02,
	0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0x48, 0x07, 0x72, 0x05, 0x10,
	0x02, 0x18, 0x80, 0x01, 0x48, 0x01, 0x52, 0x02, 0x6f, 0x73, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0xba, 0x48, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0x80, 0x01, 0x48, 0x02, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0x48, 0x07, 0x72, 0x05, 0x10, 0x02, 0x18,
	0x80, 0x01, 0x48, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0a, 0xba, 0x48, 0x07, 0x72, 0x05, 0x10, 0x02, 0x18, 0x80, 0x01, 0x48, 0x04,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x48, 0x0a, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1f, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x42, 0x08, 0xba, 0x48, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x48, 0x05, 0x52, 0x07,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x88, 0x01, 0x01, 0x12, 0x5d, 0x0a, 0x0e, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x26, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xba, 0x48, 0x05, 0x82,
	0x01, 0x02, 0x10, 0x01, 0x48, 0x06, 0x52, 0x0e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64,
	0x42, 0x05, 0x0a, 0x03, 0x5f, 0x6f, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x7f, 0x0a, 0x0c, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x49,
	0x4d, 0x41, 0x47, 0x45, 0x5f, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x26, 0x0a, 0x15, 0x49, 0x4d,
	0x41, 0x47, 0x45, 0x5f, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x4d, 0x41, 0x43, 0x48,
	0x49, 0x4e, 0x45, 0x10, 0x01, 0x1a, 0x0b, 0x82, 0xb2, 0x19, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x12, 0x28, 0x0a, 0x16, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x46, 0x45, 0x41, 0x54,
	0x55, 0x52, 0x45, 0x5f, 0x46, 0x49, 0x52, 0x45, 0x57, 0x41, 0x4c, 0x4c, 0x10, 0x02, 0x1a, 0x0c,
	0x82, 0xb2, 0x19, 0x08, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x2a, 0xd2, 0x01, 0x0a,
	0x13, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x20, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x4c,
	0x41, 0x53, 0x53, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x2d, 0x0a, 0x1c, 0x49, 0x4d,
	0x41, 0x47, 0x45, 0x5f, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x50, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x01, 0x1a, 0x0b, 0x82, 0xb2,
	0x19, 0x07, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x31, 0x0a, 0x1e, 0x49, 0x4d, 0x41,
	0x47, 0x45, 0x5f, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x02, 0x1a, 0x0d, 0x82,
	0xb2, 0x19, 0x09, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x33, 0x0a, 0x1f,
	0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x49, 0x46, 0x49, 0x43, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x50, 0x52, 0x45, 0x43, 0x41, 0x54, 0x45, 0x44, 0x10,
	0x03, 0x1a, 0x0e, 0x82, 0xb2, 0x19, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65,
	0x64, 0x32, 0xd2, 0x02, 0x0a, 0x0c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x66, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x29, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x08, 0xd8, 0xf3, 0x18, 0x03, 0xe0, 0xf3, 0x18, 0x02, 0x12, 0x69, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x2a, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x32, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x08, 0xd8, 0xf3, 0x18,
	0x03, 0xe0, 0xf3, 0x18, 0x02, 0x12, 0x6f, 0x0a, 0x06, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x12,
	0x2c, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x32, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x61,
	0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x08, 0xd8, 0xf3,
	0x18, 0x03, 0xe0, 0xf3, 0x18, 0x02, 0x42, 0xc0, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32,
	0x42, 0x0a, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35,
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
	file_metalstack_api_v2_image_proto_rawDescOnce sync.Once
	file_metalstack_api_v2_image_proto_rawDescData []byte
)

func file_metalstack_api_v2_image_proto_rawDescGZIP() []byte {
	file_metalstack_api_v2_image_proto_rawDescOnce.Do(func() {
		file_metalstack_api_v2_image_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_metalstack_api_v2_image_proto_rawDesc), len(file_metalstack_api_v2_image_proto_rawDesc)))
	})
	return file_metalstack_api_v2_image_proto_rawDescData
}

var file_metalstack_api_v2_image_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_metalstack_api_v2_image_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_metalstack_api_v2_image_proto_goTypes = []any{
	(ImageFeature)(0),                  // 0: metalstack.api.v2.ImageFeature
	(ImageClassification)(0),           // 1: metalstack.api.v2.ImageClassification
	(*ImageServiceGetRequest)(nil),     // 2: metalstack.api.v2.ImageServiceGetRequest
	(*ImageServiceListRequest)(nil),    // 3: metalstack.api.v2.ImageServiceListRequest
	(*ImageServiceLatestRequest)(nil),  // 4: metalstack.api.v2.ImageServiceLatestRequest
	(*ImageServiceGetResponse)(nil),    // 5: metalstack.api.v2.ImageServiceGetResponse
	(*ImageServiceListResponse)(nil),   // 6: metalstack.api.v2.ImageServiceListResponse
	(*ImageServiceLatestResponse)(nil), // 7: metalstack.api.v2.ImageServiceLatestResponse
	(*Image)(nil),                      // 8: metalstack.api.v2.Image
	(*ImageUsage)(nil),                 // 9: metalstack.api.v2.ImageUsage
	(*ImageQuery)(nil),                 // 10: metalstack.api.v2.ImageQuery
	(*Meta)(nil),                       // 11: metalstack.api.v2.Meta
	(*timestamppb.Timestamp)(nil),      // 12: google.protobuf.Timestamp
}
var file_metalstack_api_v2_image_proto_depIdxs = []int32{
	10, // 0: metalstack.api.v2.ImageServiceListRequest.query:type_name -> metalstack.api.v2.ImageQuery
	8,  // 1: metalstack.api.v2.ImageServiceGetResponse.image:type_name -> metalstack.api.v2.Image
	8,  // 2: metalstack.api.v2.ImageServiceListResponse.images:type_name -> metalstack.api.v2.Image
	8,  // 3: metalstack.api.v2.ImageServiceLatestResponse.image:type_name -> metalstack.api.v2.Image
	11, // 4: metalstack.api.v2.Image.meta:type_name -> metalstack.api.v2.Meta
	0,  // 5: metalstack.api.v2.Image.features:type_name -> metalstack.api.v2.ImageFeature
	1,  // 6: metalstack.api.v2.Image.classification:type_name -> metalstack.api.v2.ImageClassification
	12, // 7: metalstack.api.v2.Image.expires_at:type_name -> google.protobuf.Timestamp
	8,  // 8: metalstack.api.v2.ImageUsage.image:type_name -> metalstack.api.v2.Image
	0,  // 9: metalstack.api.v2.ImageQuery.feature:type_name -> metalstack.api.v2.ImageFeature
	1,  // 10: metalstack.api.v2.ImageQuery.classification:type_name -> metalstack.api.v2.ImageClassification
	2,  // 11: metalstack.api.v2.ImageService.Get:input_type -> metalstack.api.v2.ImageServiceGetRequest
	3,  // 12: metalstack.api.v2.ImageService.List:input_type -> metalstack.api.v2.ImageServiceListRequest
	4,  // 13: metalstack.api.v2.ImageService.Latest:input_type -> metalstack.api.v2.ImageServiceLatestRequest
	5,  // 14: metalstack.api.v2.ImageService.Get:output_type -> metalstack.api.v2.ImageServiceGetResponse
	6,  // 15: metalstack.api.v2.ImageService.List:output_type -> metalstack.api.v2.ImageServiceListResponse
	7,  // 16: metalstack.api.v2.ImageService.Latest:output_type -> metalstack.api.v2.ImageServiceLatestResponse
	14, // [14:17] is the sub-list for method output_type
	11, // [11:14] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_metalstack_api_v2_image_proto_init() }
func file_metalstack_api_v2_image_proto_init() {
	if File_metalstack_api_v2_image_proto != nil {
		return
	}
	file_metalstack_api_v2_common_proto_init()
	file_metalstack_api_v2_image_proto_msgTypes[6].OneofWrappers = []any{}
	file_metalstack_api_v2_image_proto_msgTypes[8].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_metalstack_api_v2_image_proto_rawDesc), len(file_metalstack_api_v2_image_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metalstack_api_v2_image_proto_goTypes,
		DependencyIndexes: file_metalstack_api_v2_image_proto_depIdxs,
		EnumInfos:         file_metalstack_api_v2_image_proto_enumTypes,
		MessageInfos:      file_metalstack_api_v2_image_proto_msgTypes,
	}.Build()
	File_metalstack_api_v2_image_proto = out.File
	file_metalstack_api_v2_image_proto_goTypes = nil
	file_metalstack_api_v2_image_proto_depIdxs = nil
}
