# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: metalstack/admin/v2/size.proto
# Protobuf Python Version: 6.31.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    6,
    31,
    1,
    '',
    'metalstack/admin/v2/size.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from buf.validate import validate_pb2 as buf_dot_validate_dot_validate__pb2
from metalstack.api.v2 import common_pb2 as metalstack_dot_api_dot_v2_dot_common__pb2
from metalstack.api.v2 import size_pb2 as metalstack_dot_api_dot_v2_dot_size__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1emetalstack/admin/v2/size.proto\x12\x13metalstack.admin.v2\x1a\x1b\x62uf/validate/validate.proto\x1a\x1emetalstack/api/v2/common.proto\x1a\x1cmetalstack/api/v2/size.proto\"G\n\x18SizeServiceCreateRequest\x12+\n\x04size\x18\x01 \x01(\x0b\x32\x17.metalstack.api.v2.SizeR\x04size\"H\n\x19SizeServiceCreateResponse\x12+\n\x04size\x18\x01 \x01(\x0b\x32\x17.metalstack.api.v2.SizeR\x04size\"\xb5\x02\n\x18SizeServiceUpdateRequest\x12\x1a\n\x02id\x18\x01 \x01(\tB\n\xbaH\x07r\x05\x10\x02\x18\x80\x01R\x02id\x12#\n\x04name\x18\x04 \x01(\tB\n\xbaH\x07r\x05\x10\x02\x18\x80\x01H\x00R\x04name\x88\x01\x01\x12\x31\n\x0b\x64\x65scription\x18\x05 \x01(\tB\n\xbaH\x07r\x05\x10\x02\x18\x80\x01H\x01R\x0b\x64\x65scription\x88\x01\x01\x12\x43\n\x0b\x63onstraints\x18\x06 \x03(\x0b\x32!.metalstack.api.v2.SizeConstraintR\x0b\x63onstraints\x12<\n\x06labels\x18\x07 \x01(\x0b\x32\x1f.metalstack.api.v2.UpdateLabelsH\x02R\x06labels\x88\x01\x01\x42\x07\n\x05_nameB\x0e\n\x0c_descriptionB\t\n\x07_labels\"H\n\x19SizeServiceUpdateResponse\x12+\n\x04size\x18\x01 \x01(\x0b\x32\x17.metalstack.api.v2.SizeR\x04size\"6\n\x18SizeServiceDeleteRequest\x12\x1a\n\x02id\x18\x01 \x01(\tB\n\xbaH\x07r\x05\x10\x02\x18\x80\x01R\x02id\"H\n\x19SizeServiceDeleteResponse\x12+\n\x04size\x18\x01 \x01(\x0b\x32\x17.metalstack.api.v2.SizeR\x04size2\xe0\x02\n\x0bSizeService\x12o\n\x06\x43reate\x12-.metalstack.admin.v2.SizeServiceCreateRequest\x1a..metalstack.admin.v2.SizeServiceCreateResponse\"\x06\xd2\xf3\x18\x02\x01\x02\x12o\n\x06Update\x12-.metalstack.admin.v2.SizeServiceUpdateRequest\x1a..metalstack.admin.v2.SizeServiceUpdateResponse\"\x06\xd2\xf3\x18\x02\x01\x02\x12o\n\x06\x44\x65lete\x12-.metalstack.admin.v2.SizeServiceDeleteRequest\x1a..metalstack.admin.v2.SizeServiceDeleteResponse\"\x06\xd2\xf3\x18\x02\x01\x02\x42\xcd\x01\n\x17\x63om.metalstack.admin.v2B\tSizeProtoP\x01Z9github.com/metal-stack/api/go/metalstack/admin/v2;adminv2\xa2\x02\x03MAX\xaa\x02\x13Metalstack.Admin.V2\xca\x02\x13Metalstack\\Admin\\V2\xe2\x02\x1fMetalstack\\Admin\\V2\\GPBMetadata\xea\x02\x15Metalstack::Admin::V2b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'metalstack.admin.v2.size_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\027com.metalstack.admin.v2B\tSizeProtoP\001Z9github.com/metal-stack/api/go/metalstack/admin/v2;adminv2\242\002\003MAX\252\002\023Metalstack.Admin.V2\312\002\023Metalstack\\Admin\\V2\342\002\037Metalstack\\Admin\\V2\\GPBMetadata\352\002\025Metalstack::Admin::V2'
  _globals['_SIZESERVICEUPDATEREQUEST'].fields_by_name['id']._loaded_options = None
  _globals['_SIZESERVICEUPDATEREQUEST'].fields_by_name['id']._serialized_options = b'\272H\007r\005\020\002\030\200\001'
  _globals['_SIZESERVICEUPDATEREQUEST'].fields_by_name['name']._loaded_options = None
  _globals['_SIZESERVICEUPDATEREQUEST'].fields_by_name['name']._serialized_options = b'\272H\007r\005\020\002\030\200\001'
  _globals['_SIZESERVICEUPDATEREQUEST'].fields_by_name['description']._loaded_options = None
  _globals['_SIZESERVICEUPDATEREQUEST'].fields_by_name['description']._serialized_options = b'\272H\007r\005\020\002\030\200\001'
  _globals['_SIZESERVICEDELETEREQUEST'].fields_by_name['id']._loaded_options = None
  _globals['_SIZESERVICEDELETEREQUEST'].fields_by_name['id']._serialized_options = b'\272H\007r\005\020\002\030\200\001'
  _globals['_SIZESERVICE'].methods_by_name['Create']._loaded_options = None
  _globals['_SIZESERVICE'].methods_by_name['Create']._serialized_options = b'\322\363\030\002\001\002'
  _globals['_SIZESERVICE'].methods_by_name['Update']._loaded_options = None
  _globals['_SIZESERVICE'].methods_by_name['Update']._serialized_options = b'\322\363\030\002\001\002'
  _globals['_SIZESERVICE'].methods_by_name['Delete']._loaded_options = None
  _globals['_SIZESERVICE'].methods_by_name['Delete']._serialized_options = b'\322\363\030\002\001\002'
  _globals['_SIZESERVICECREATEREQUEST']._serialized_start=146
  _globals['_SIZESERVICECREATEREQUEST']._serialized_end=217
  _globals['_SIZESERVICECREATERESPONSE']._serialized_start=219
  _globals['_SIZESERVICECREATERESPONSE']._serialized_end=291
  _globals['_SIZESERVICEUPDATEREQUEST']._serialized_start=294
  _globals['_SIZESERVICEUPDATEREQUEST']._serialized_end=603
  _globals['_SIZESERVICEUPDATERESPONSE']._serialized_start=605
  _globals['_SIZESERVICEUPDATERESPONSE']._serialized_end=677
  _globals['_SIZESERVICEDELETEREQUEST']._serialized_start=679
  _globals['_SIZESERVICEDELETEREQUEST']._serialized_end=733
  _globals['_SIZESERVICEDELETERESPONSE']._serialized_start=735
  _globals['_SIZESERVICEDELETERESPONSE']._serialized_end=807
  _globals['_SIZESERVICE']._serialized_start=810
  _globals['_SIZESERVICE']._serialized_end=1162
# @@protoc_insertion_point(module_scope)
