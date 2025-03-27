// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: metalstack/admin/v2/network.proto

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

// NetworkServiceCreateRequest
type NetworkServiceCreateRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Id of this network
	Id *string `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// Name of this network
	Name *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// Description of this network
	Description *string `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
	// Partition where this network will be created
	Partition *string `protobuf:"bytes,4,opt,name=partition,proto3,oneof" json:"partition,omitempty"`
	// Project where this network belongs to
	Project *string `protobuf:"bytes,5,opt,name=project,proto3,oneof" json:"project,omitempty"`
	// Labels on this network
	Labels *v2.Labels `protobuf:"bytes,6,opt,name=labels,proto3,oneof" json:"labels,omitempty"`
	// Prefixes in this network
	Prefixes []string `protobuf:"bytes,8,rep,name=prefixes,proto3" json:"prefixes,omitempty"`
	// Destination Prefixes in this network
	DestinationPrefixes []string `protobuf:"bytes,9,rep,name=destination_prefixes,json=destinationPrefixes,proto3" json:"destination_prefixes,omitempty"`
	// Default Child Prefix length defines the bitlength of a child network created per addressfamily, of not specified during the allocate request
	DefaultChildPrefixLength []*v2.ChildPrefixLength `protobuf:"bytes,10,rep,name=default_child_prefix_length,json=defaultChildPrefixLength,proto3" json:"default_child_prefix_length,omitempty"`
	// Options of this network
	Options *v2.NetworkOptions `protobuf:"bytes,11,opt,name=options,proto3" json:"options,omitempty"`
	// Vrf ID of this network
	Vrf *uint32 `protobuf:"varint,12,opt,name=vrf,proto3,oneof" json:"vrf,omitempty"`
	// Parent NetworkId points to the id of the parent network if any
	ParentNetworkId *string `protobuf:"bytes,14,opt,name=parent_network_id,json=parentNetworkId,proto3,oneof" json:"parent_network_id,omitempty"`
	// AdditionalAnnounceableCidrs will be added to the allow list on the switch which prefixes might be announced
	AdditionalAnnounceableCidrs []string `protobuf:"bytes,15,rep,name=additional_announceable_cidrs,json=additionalAnnounceableCidrs,proto3" json:"additional_announceable_cidrs,omitempty"`
	// Bitlength per addressfamily
	Length *v2.ChildPrefixLength `protobuf:"bytes,16,opt,name=length,proto3" json:"length,omitempty"`
	// AddressFamily to create, defaults to the same as the parent
	AddressFamily *v2.IPAddressFamily `protobuf:"varint,17,opt,name=address_family,json=addressFamily,proto3,enum=metalstack.api.v2.IPAddressFamily,oneof" json:"address_family,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NetworkServiceCreateRequest) Reset() {
	*x = NetworkServiceCreateRequest{}
	mi := &file_metalstack_admin_v2_network_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetworkServiceCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkServiceCreateRequest) ProtoMessage() {}

func (x *NetworkServiceCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_admin_v2_network_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkServiceCreateRequest.ProtoReflect.Descriptor instead.
func (*NetworkServiceCreateRequest) Descriptor() ([]byte, []int) {
	return file_metalstack_admin_v2_network_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkServiceCreateRequest) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *NetworkServiceCreateRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *NetworkServiceCreateRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *NetworkServiceCreateRequest) GetPartition() string {
	if x != nil && x.Partition != nil {
		return *x.Partition
	}
	return ""
}

func (x *NetworkServiceCreateRequest) GetProject() string {
	if x != nil && x.Project != nil {
		return *x.Project
	}
	return ""
}

func (x *NetworkServiceCreateRequest) GetLabels() *v2.Labels {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *NetworkServiceCreateRequest) GetPrefixes() []string {
	if x != nil {
		return x.Prefixes
	}
	return nil
}

func (x *NetworkServiceCreateRequest) GetDestinationPrefixes() []string {
	if x != nil {
		return x.DestinationPrefixes
	}
	return nil
}

func (x *NetworkServiceCreateRequest) GetDefaultChildPrefixLength() []*v2.ChildPrefixLength {
	if x != nil {
		return x.DefaultChildPrefixLength
	}
	return nil
}

func (x *NetworkServiceCreateRequest) GetOptions() *v2.NetworkOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *NetworkServiceCreateRequest) GetVrf() uint32 {
	if x != nil && x.Vrf != nil {
		return *x.Vrf
	}
	return 0
}

func (x *NetworkServiceCreateRequest) GetParentNetworkId() string {
	if x != nil && x.ParentNetworkId != nil {
		return *x.ParentNetworkId
	}
	return ""
}

func (x *NetworkServiceCreateRequest) GetAdditionalAnnounceableCidrs() []string {
	if x != nil {
		return x.AdditionalAnnounceableCidrs
	}
	return nil
}

func (x *NetworkServiceCreateRequest) GetLength() *v2.ChildPrefixLength {
	if x != nil {
		return x.Length
	}
	return nil
}

func (x *NetworkServiceCreateRequest) GetAddressFamily() v2.IPAddressFamily {
	if x != nil && x.AddressFamily != nil {
		return *x.AddressFamily
	}
	return v2.IPAddressFamily(0)
}

// NetworkServiceUpdateRequest is the request payload for a network update request
type NetworkServiceUpdateRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Network the network
	Network       *v2.Network `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NetworkServiceUpdateRequest) Reset() {
	*x = NetworkServiceUpdateRequest{}
	mi := &file_metalstack_admin_v2_network_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetworkServiceUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkServiceUpdateRequest) ProtoMessage() {}

func (x *NetworkServiceUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_admin_v2_network_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkServiceUpdateRequest.ProtoReflect.Descriptor instead.
func (*NetworkServiceUpdateRequest) Descriptor() ([]byte, []int) {
	return file_metalstack_admin_v2_network_proto_rawDescGZIP(), []int{1}
}

func (x *NetworkServiceUpdateRequest) GetNetwork() *v2.Network {
	if x != nil {
		return x.Network
	}
	return nil
}

// NetworkServiceDeleteRequest is the request payload for a network delete request
type NetworkServiceDeleteRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the network to get
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NetworkServiceDeleteRequest) Reset() {
	*x = NetworkServiceDeleteRequest{}
	mi := &file_metalstack_admin_v2_network_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetworkServiceDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkServiceDeleteRequest) ProtoMessage() {}

func (x *NetworkServiceDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_admin_v2_network_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkServiceDeleteRequest.ProtoReflect.Descriptor instead.
func (*NetworkServiceDeleteRequest) Descriptor() ([]byte, []int) {
	return file_metalstack_admin_v2_network_proto_rawDescGZIP(), []int{2}
}

func (x *NetworkServiceDeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// NetworkServiceListRequest
type NetworkServiceListRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Query which specifies which networks to return
	Query         *v2.NetworkQuery `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NetworkServiceListRequest) Reset() {
	*x = NetworkServiceListRequest{}
	mi := &file_metalstack_admin_v2_network_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetworkServiceListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkServiceListRequest) ProtoMessage() {}

func (x *NetworkServiceListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_admin_v2_network_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkServiceListRequest.ProtoReflect.Descriptor instead.
func (*NetworkServiceListRequest) Descriptor() ([]byte, []int) {
	return file_metalstack_admin_v2_network_proto_rawDescGZIP(), []int{3}
}

func (x *NetworkServiceListRequest) GetQuery() *v2.NetworkQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

// NetworkServiceCreateResponse is the response payload for a network create request
type NetworkServiceCreateResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Network the network
	Network       *v2.Network `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NetworkServiceCreateResponse) Reset() {
	*x = NetworkServiceCreateResponse{}
	mi := &file_metalstack_admin_v2_network_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetworkServiceCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkServiceCreateResponse) ProtoMessage() {}

func (x *NetworkServiceCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_admin_v2_network_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkServiceCreateResponse.ProtoReflect.Descriptor instead.
func (*NetworkServiceCreateResponse) Descriptor() ([]byte, []int) {
	return file_metalstack_admin_v2_network_proto_rawDescGZIP(), []int{4}
}

func (x *NetworkServiceCreateResponse) GetNetwork() *v2.Network {
	if x != nil {
		return x.Network
	}
	return nil
}

// NetworkServiceUpdateResponse is the response payload for a network update request
type NetworkServiceUpdateResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Network the network
	Network       *v2.Network `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NetworkServiceUpdateResponse) Reset() {
	*x = NetworkServiceUpdateResponse{}
	mi := &file_metalstack_admin_v2_network_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetworkServiceUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkServiceUpdateResponse) ProtoMessage() {}

func (x *NetworkServiceUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_admin_v2_network_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkServiceUpdateResponse.ProtoReflect.Descriptor instead.
func (*NetworkServiceUpdateResponse) Descriptor() ([]byte, []int) {
	return file_metalstack_admin_v2_network_proto_rawDescGZIP(), []int{5}
}

func (x *NetworkServiceUpdateResponse) GetNetwork() *v2.Network {
	if x != nil {
		return x.Network
	}
	return nil
}

// NetworkServiceCapacityResponse is the response payload for a network delete request
type NetworkServiceDeleteResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Network the network
	Network       *v2.Network `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NetworkServiceDeleteResponse) Reset() {
	*x = NetworkServiceDeleteResponse{}
	mi := &file_metalstack_admin_v2_network_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetworkServiceDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkServiceDeleteResponse) ProtoMessage() {}

func (x *NetworkServiceDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_admin_v2_network_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkServiceDeleteResponse.ProtoReflect.Descriptor instead.
func (*NetworkServiceDeleteResponse) Descriptor() ([]byte, []int) {
	return file_metalstack_admin_v2_network_proto_rawDescGZIP(), []int{6}
}

func (x *NetworkServiceDeleteResponse) GetNetwork() *v2.Network {
	if x != nil {
		return x.Network
	}
	return nil
}

// NetworkServiceListResponse
type NetworkServiceListResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Networks are the requested networks
	Networks      []*v2.Network `protobuf:"bytes,1,rep,name=networks,proto3" json:"networks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NetworkServiceListResponse) Reset() {
	*x = NetworkServiceListResponse{}
	mi := &file_metalstack_admin_v2_network_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetworkServiceListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkServiceListResponse) ProtoMessage() {}

func (x *NetworkServiceListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_admin_v2_network_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkServiceListResponse.ProtoReflect.Descriptor instead.
func (*NetworkServiceListResponse) Descriptor() ([]byte, []int) {
	return file_metalstack_admin_v2_network_proto_rawDescGZIP(), []int{7}
}

func (x *NetworkServiceListResponse) GetNetworks() []*v2.Network {
	if x != nil {
		return x.Networks
	}
	return nil
}

var File_metalstack_admin_v2_network_proto protoreflect.FileDescriptor

const file_metalstack_admin_v2_network_proto_rawDesc = "" +
	"\n" +
	"!metalstack/admin/v2/network.proto\x12\x13metalstack.admin.v2\x1a\x1bbuf/validate/validate.proto\x1a\x1emetalstack/api/v2/common.proto\x1a\x1ametalstack/api/v2/ip.proto\x1a\x1fmetalstack/api/v2/network.proto\"\x98\n" +
	"\n" +
	"\x1bNetworkServiceCreateRequest\x12\x1f\n" +
	"\x02id\x18\x01 \x01(\tB\n" +
	"\xbaH\ar\x05\x10\x02\x18\x80\x01H\x00R\x02id\x88\x01\x01\x12#\n" +
	"\x04name\x18\x02 \x01(\tB\n" +
	"\xbaH\ar\x05\x10\x02\x18\x80\x01H\x01R\x04name\x88\x01\x01\x121\n" +
	"\vdescription\x18\x03 \x01(\tB\n" +
	"\xbaH\ar\x05\x10\x02\x18\x80\x01H\x02R\vdescription\x88\x01\x01\x12-\n" +
	"\tpartition\x18\x04 \x01(\tB\n" +
	"\xbaH\ar\x05\x10\x02\x18\x80\x01H\x03R\tpartition\x88\x01\x01\x12'\n" +
	"\aproject\x18\x05 \x01(\tB\b\xbaH\x05r\x03\xb0\x01\x01H\x04R\aproject\x88\x01\x01\x126\n" +
	"\x06labels\x18\x06 \x01(\v2\x19.metalstack.api.v2.LabelsH\x05R\x06labels\x88\x01\x01\x12\x1a\n" +
	"\bprefixes\x18\b \x03(\tR\bprefixes\x121\n" +
	"\x14destination_prefixes\x18\t \x03(\tR\x13destinationPrefixes\x12c\n" +
	"\x1bdefault_child_prefix_length\x18\n" +
	" \x03(\v2$.metalstack.api.v2.ChildPrefixLengthR\x18defaultChildPrefixLength\x12;\n" +
	"\aoptions\x18\v \x01(\v2!.metalstack.api.v2.NetworkOptionsR\aoptions\x12\x15\n" +
	"\x03vrf\x18\f \x01(\rH\x06R\x03vrf\x88\x01\x01\x12;\n" +
	"\x11parent_network_id\x18\x0e \x01(\tB\n" +
	"\xbaH\ar\x05\x10\x02\x18\x80\x01H\aR\x0fparentNetworkId\x88\x01\x01\x12B\n" +
	"\x1dadditional_announceable_cidrs\x18\x0f \x03(\tR\x1badditionalAnnounceableCidrs\x12<\n" +
	"\x06length\x18\x10 \x01(\v2$.metalstack.api.v2.ChildPrefixLengthR\x06length\x12X\n" +
	"\x0eaddress_family\x18\x11 \x01(\x0e2\".metalstack.api.v2.IPAddressFamilyB\b\xbaH\x05\x82\x01\x02\x10\x01H\bR\raddressFamily\x88\x01\x01:\xd8\x02\xbaH\xd4\x02\x1aN\n" +
	"\bprefixes\x12\x1cgiven prefixes must be valid\x1a$this.prefixes.all(m, m.isIpPrefix())\x1ar\n" +
	"\x14destination_prefixes\x12(given destination_prefixes must be valid\x1a0this.destination_prefixes.all(m, m.isIpPrefix())\x1a\x8d\x01\n" +
	"\x1dadditional_announceable_cidrs\x121given additional_announceable_cidrs must be valid\x1a9this.additional_announceable_cidrs.all(m, m.isIpPrefix())B\x05\n" +
	"\x03_idB\a\n" +
	"\x05_nameB\x0e\n" +
	"\f_descriptionB\f\n" +
	"\n" +
	"_partitionB\n" +
	"\n" +
	"\b_projectB\t\n" +
	"\a_labelsB\x06\n" +
	"\x04_vrfB\x14\n" +
	"\x12_parent_network_idB\x11\n" +
	"\x0f_address_family\"S\n" +
	"\x1bNetworkServiceUpdateRequest\x124\n" +
	"\anetwork\x18\x01 \x01(\v2\x1a.metalstack.api.v2.NetworkR\anetwork\"9\n" +
	"\x1bNetworkServiceDeleteRequest\x12\x1a\n" +
	"\x02id\x18\x01 \x01(\tB\n" +
	"\xbaH\ar\x05\x10\x02\x18\x80\x01R\x02id\"R\n" +
	"\x19NetworkServiceListRequest\x125\n" +
	"\x05query\x18\x02 \x01(\v2\x1f.metalstack.api.v2.NetworkQueryR\x05query\"T\n" +
	"\x1cNetworkServiceCreateResponse\x124\n" +
	"\anetwork\x18\x01 \x01(\v2\x1a.metalstack.api.v2.NetworkR\anetwork\"T\n" +
	"\x1cNetworkServiceUpdateResponse\x124\n" +
	"\anetwork\x18\x01 \x01(\v2\x1a.metalstack.api.v2.NetworkR\anetwork\"T\n" +
	"\x1cNetworkServiceDeleteResponse\x124\n" +
	"\anetwork\x18\x01 \x01(\v2\x1a.metalstack.api.v2.NetworkR\anetwork\"T\n" +
	"\x1aNetworkServiceListResponse\x126\n" +
	"\bnetworks\x18\x01 \x03(\v2\x1a.metalstack.api.v2.NetworkR\bnetworks2\xf3\x03\n" +
	"\x0eNetworkService\x12x\n" +
	"\x06Create\x120.metalstack.admin.v2.NetworkServiceCreateRequest\x1a1.metalstack.admin.v2.NetworkServiceCreateResponse\"\t\xd2\xf3\x18\x01\x01\xe0\xf3\x18\x01\x12x\n" +
	"\x06Update\x120.metalstack.admin.v2.NetworkServiceUpdateRequest\x1a1.metalstack.admin.v2.NetworkServiceUpdateResponse\"\t\xd2\xf3\x18\x01\x01\xe0\xf3\x18\x01\x12x\n" +
	"\x06Delete\x120.metalstack.admin.v2.NetworkServiceDeleteRequest\x1a1.metalstack.admin.v2.NetworkServiceDeleteResponse\"\t\xd2\xf3\x18\x01\x01\xe0\xf3\x18\x01\x12s\n" +
	"\x04List\x12..metalstack.admin.v2.NetworkServiceListRequest\x1a/.metalstack.admin.v2.NetworkServiceListResponse\"\n" +
	"\xd2\xf3\x18\x02\x01\x02\xe0\xf3\x18\x02B\xd0\x01\n" +
	"\x17com.metalstack.admin.v2B\fNetworkProtoP\x01Z9github.com/metal-stack/api/go/metalstack/admin/v2;adminv2\xa2\x02\x03MAX\xaa\x02\x13Metalstack.Admin.V2\xca\x02\x13Metalstack\\Admin\\V2\xe2\x02\x1fMetalstack\\Admin\\V2\\GPBMetadata\xea\x02\x15Metalstack::Admin::V2b\x06proto3"

var (
	file_metalstack_admin_v2_network_proto_rawDescOnce sync.Once
	file_metalstack_admin_v2_network_proto_rawDescData []byte
)

func file_metalstack_admin_v2_network_proto_rawDescGZIP() []byte {
	file_metalstack_admin_v2_network_proto_rawDescOnce.Do(func() {
		file_metalstack_admin_v2_network_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_metalstack_admin_v2_network_proto_rawDesc), len(file_metalstack_admin_v2_network_proto_rawDesc)))
	})
	return file_metalstack_admin_v2_network_proto_rawDescData
}

var file_metalstack_admin_v2_network_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_metalstack_admin_v2_network_proto_goTypes = []any{
	(*NetworkServiceCreateRequest)(nil),  // 0: metalstack.admin.v2.NetworkServiceCreateRequest
	(*NetworkServiceUpdateRequest)(nil),  // 1: metalstack.admin.v2.NetworkServiceUpdateRequest
	(*NetworkServiceDeleteRequest)(nil),  // 2: metalstack.admin.v2.NetworkServiceDeleteRequest
	(*NetworkServiceListRequest)(nil),    // 3: metalstack.admin.v2.NetworkServiceListRequest
	(*NetworkServiceCreateResponse)(nil), // 4: metalstack.admin.v2.NetworkServiceCreateResponse
	(*NetworkServiceUpdateResponse)(nil), // 5: metalstack.admin.v2.NetworkServiceUpdateResponse
	(*NetworkServiceDeleteResponse)(nil), // 6: metalstack.admin.v2.NetworkServiceDeleteResponse
	(*NetworkServiceListResponse)(nil),   // 7: metalstack.admin.v2.NetworkServiceListResponse
	(*v2.Labels)(nil),                    // 8: metalstack.api.v2.Labels
	(*v2.ChildPrefixLength)(nil),         // 9: metalstack.api.v2.ChildPrefixLength
	(*v2.NetworkOptions)(nil),            // 10: metalstack.api.v2.NetworkOptions
	(v2.IPAddressFamily)(0),              // 11: metalstack.api.v2.IPAddressFamily
	(*v2.Network)(nil),                   // 12: metalstack.api.v2.Network
	(*v2.NetworkQuery)(nil),              // 13: metalstack.api.v2.NetworkQuery
}
var file_metalstack_admin_v2_network_proto_depIdxs = []int32{
	8,  // 0: metalstack.admin.v2.NetworkServiceCreateRequest.labels:type_name -> metalstack.api.v2.Labels
	9,  // 1: metalstack.admin.v2.NetworkServiceCreateRequest.default_child_prefix_length:type_name -> metalstack.api.v2.ChildPrefixLength
	10, // 2: metalstack.admin.v2.NetworkServiceCreateRequest.options:type_name -> metalstack.api.v2.NetworkOptions
	9,  // 3: metalstack.admin.v2.NetworkServiceCreateRequest.length:type_name -> metalstack.api.v2.ChildPrefixLength
	11, // 4: metalstack.admin.v2.NetworkServiceCreateRequest.address_family:type_name -> metalstack.api.v2.IPAddressFamily
	12, // 5: metalstack.admin.v2.NetworkServiceUpdateRequest.network:type_name -> metalstack.api.v2.Network
	13, // 6: metalstack.admin.v2.NetworkServiceListRequest.query:type_name -> metalstack.api.v2.NetworkQuery
	12, // 7: metalstack.admin.v2.NetworkServiceCreateResponse.network:type_name -> metalstack.api.v2.Network
	12, // 8: metalstack.admin.v2.NetworkServiceUpdateResponse.network:type_name -> metalstack.api.v2.Network
	12, // 9: metalstack.admin.v2.NetworkServiceDeleteResponse.network:type_name -> metalstack.api.v2.Network
	12, // 10: metalstack.admin.v2.NetworkServiceListResponse.networks:type_name -> metalstack.api.v2.Network
	0,  // 11: metalstack.admin.v2.NetworkService.Create:input_type -> metalstack.admin.v2.NetworkServiceCreateRequest
	1,  // 12: metalstack.admin.v2.NetworkService.Update:input_type -> metalstack.admin.v2.NetworkServiceUpdateRequest
	2,  // 13: metalstack.admin.v2.NetworkService.Delete:input_type -> metalstack.admin.v2.NetworkServiceDeleteRequest
	3,  // 14: metalstack.admin.v2.NetworkService.List:input_type -> metalstack.admin.v2.NetworkServiceListRequest
	4,  // 15: metalstack.admin.v2.NetworkService.Create:output_type -> metalstack.admin.v2.NetworkServiceCreateResponse
	5,  // 16: metalstack.admin.v2.NetworkService.Update:output_type -> metalstack.admin.v2.NetworkServiceUpdateResponse
	6,  // 17: metalstack.admin.v2.NetworkService.Delete:output_type -> metalstack.admin.v2.NetworkServiceDeleteResponse
	7,  // 18: metalstack.admin.v2.NetworkService.List:output_type -> metalstack.admin.v2.NetworkServiceListResponse
	15, // [15:19] is the sub-list for method output_type
	11, // [11:15] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_metalstack_admin_v2_network_proto_init() }
func file_metalstack_admin_v2_network_proto_init() {
	if File_metalstack_admin_v2_network_proto != nil {
		return
	}
	file_metalstack_admin_v2_network_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_metalstack_admin_v2_network_proto_rawDesc), len(file_metalstack_admin_v2_network_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metalstack_admin_v2_network_proto_goTypes,
		DependencyIndexes: file_metalstack_admin_v2_network_proto_depIdxs,
		MessageInfos:      file_metalstack_admin_v2_network_proto_msgTypes,
	}.Build()
	File_metalstack_admin_v2_network_proto = out.File
	file_metalstack_admin_v2_network_proto_goTypes = nil
	file_metalstack_admin_v2_network_proto_depIdxs = nil
}
