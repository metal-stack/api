# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: metalstack/admin/v2/partition.proto
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
    'metalstack/admin/v2/partition.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from buf.validate import validate_pb2 as buf_dot_validate_dot_validate__pb2
from metalstack.api.v2 import common_pb2 as metalstack_dot_api_dot_v2_dot_common__pb2
from metalstack.api.v2 import partition_pb2 as metalstack_dot_api_dot_v2_dot_partition__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n#metalstack/admin/v2/partition.proto\x12\x13metalstack.admin.v2\x1a\x1b\x62uf/validate/validate.proto\x1a\x1emetalstack/api/v2/common.proto\x1a!metalstack/api/v2/partition.proto\"[\n\x1dPartitionServiceCreateRequest\x12:\n\tpartition\x18\x01 \x01(\x0b\x32\x1c.metalstack.api.v2.PartitionR\tpartition\"[\n\x1dPartitionServiceUpdateRequest\x12:\n\tpartition\x18\x01 \x01(\x0b\x32\x1c.metalstack.api.v2.PartitionR\tpartition\";\n\x1dPartitionServiceDeleteRequest\x12\x1a\n\x02id\x18\x01 \x01(\tB\n\xbaH\x07r\x05\x10\x02\x18\x80\x01R\x02id\"\\\n\x1ePartitionServiceCreateResponse\x12:\n\tpartition\x18\x01 \x01(\x0b\x32\x1c.metalstack.api.v2.PartitionR\tpartition\"\\\n\x1ePartitionServiceUpdateResponse\x12:\n\tpartition\x18\x01 \x01(\x0b\x32\x1c.metalstack.api.v2.PartitionR\tpartition\"\\\n\x1ePartitionServiceDeleteResponse\x12:\n\tpartition\x18\x01 \x01(\x0b\x32\x1c.metalstack.api.v2.PartitionR\tpartition\"\xac\x01\n\x1fPartitionServiceCapacityRequest\x12\x1f\n\x02id\x18\x01 \x01(\tB\n\xbaH\x07r\x05\x10\x02\x18\x80\x01H\x00R\x02id\x88\x01\x01\x12#\n\x04size\x18\x02 \x01(\tB\n\xbaH\x07r\x05\x10\x02\x18\x80\x01H\x01R\x04size\x88\x01\x01\x12\'\n\x07project\x18\x03 \x01(\tB\x08\xbaH\x05r\x03\xb0\x01\x01H\x02R\x07project\x88\x01\x01\x42\x05\n\x03_idB\x07\n\x05_sizeB\n\n\x08_project\"\x83\x04\n PartitionServiceCapacityResponse\x12\x12\n\x04size\x18\x01 \x01(\tR\x04size\x12\x14\n\x05total\x18\x02 \x01(\x03R\x05total\x12\x1f\n\x0bphoned_home\x18\x03 \x01(\x03R\nphonedHome\x12\x18\n\x07waiting\x18\x04 \x01(\x03R\x07waiting\x12\x14\n\x05other\x18\x05 \x01(\x03R\x05other\x12%\n\x0eother_machines\x18\x06 \x03(\tR\rotherMachines\x12\x1c\n\tallocated\x18\x07 \x01(\x03R\tallocated\x12 \n\x0b\x61llocatable\x18\x08 \x01(\x03R\x0b\x61llocatable\x12\x12\n\x04\x66ree\x18\t \x01(\x03R\x04\x66ree\x12 \n\x0bunavailable\x18\n \x01(\x03R\x0bunavailable\x12\x16\n\x06\x66\x61ulty\x18\x0b \x01(\x03R\x06\x66\x61ulty\x12\'\n\x0f\x66\x61ulty_machines\x18\x0c \x03(\tR\x0e\x66\x61ultyMachines\x12\"\n\x0creservations\x18\r \x01(\x03R\x0creservations\x12+\n\x11used_reservations\x18\x0e \x01(\x03R\x10usedReservations\x12\x35\n\x16remaining_reservations\x18\x0f \x01(\x03R\x15remainingReservations2\x92\x04\n\x10PartitionService\x12|\n\x06\x43reate\x12\x32.metalstack.admin.v2.PartitionServiceCreateRequest\x1a\x33.metalstack.admin.v2.PartitionServiceCreateResponse\"\t\xd2\xf3\x18\x01\x01\xe0\xf3\x18\x01\x12|\n\x06Update\x12\x32.metalstack.admin.v2.PartitionServiceUpdateRequest\x1a\x33.metalstack.admin.v2.PartitionServiceUpdateResponse\"\t\xd2\xf3\x18\x01\x01\xe0\xf3\x18\x01\x12|\n\x06\x44\x65lete\x12\x32.metalstack.admin.v2.PartitionServiceDeleteRequest\x1a\x33.metalstack.admin.v2.PartitionServiceDeleteResponse\"\t\xd2\xf3\x18\x01\x01\xe0\xf3\x18\x01\x12\x83\x01\n\x08\x43\x61pacity\x12\x34.metalstack.admin.v2.PartitionServiceCapacityRequest\x1a\x35.metalstack.admin.v2.PartitionServiceCapacityResponse\"\n\xd2\xf3\x18\x02\x02\x01\xe0\xf3\x18\x02\x42\xd2\x01\n\x17\x63om.metalstack.admin.v2B\x0ePartitionProtoP\x01Z9github.com/metal-stack/api/go/metalstack/admin/v2;adminv2\xa2\x02\x03MAX\xaa\x02\x13Metalstack.Admin.V2\xca\x02\x13Metalstack\\Admin\\V2\xe2\x02\x1fMetalstack\\Admin\\V2\\GPBMetadata\xea\x02\x15Metalstack::Admin::V2b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'metalstack.admin.v2.partition_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\027com.metalstack.admin.v2B\016PartitionProtoP\001Z9github.com/metal-stack/api/go/metalstack/admin/v2;adminv2\242\002\003MAX\252\002\023Metalstack.Admin.V2\312\002\023Metalstack\\Admin\\V2\342\002\037Metalstack\\Admin\\V2\\GPBMetadata\352\002\025Metalstack::Admin::V2'
  _globals['_PARTITIONSERVICEDELETEREQUEST'].fields_by_name['id']._loaded_options = None
  _globals['_PARTITIONSERVICEDELETEREQUEST'].fields_by_name['id']._serialized_options = b'\272H\007r\005\020\002\030\200\001'
  _globals['_PARTITIONSERVICECAPACITYREQUEST'].fields_by_name['id']._loaded_options = None
  _globals['_PARTITIONSERVICECAPACITYREQUEST'].fields_by_name['id']._serialized_options = b'\272H\007r\005\020\002\030\200\001'
  _globals['_PARTITIONSERVICECAPACITYREQUEST'].fields_by_name['size']._loaded_options = None
  _globals['_PARTITIONSERVICECAPACITYREQUEST'].fields_by_name['size']._serialized_options = b'\272H\007r\005\020\002\030\200\001'
  _globals['_PARTITIONSERVICECAPACITYREQUEST'].fields_by_name['project']._loaded_options = None
  _globals['_PARTITIONSERVICECAPACITYREQUEST'].fields_by_name['project']._serialized_options = b'\272H\005r\003\260\001\001'
  _globals['_PARTITIONSERVICE'].methods_by_name['Create']._loaded_options = None
  _globals['_PARTITIONSERVICE'].methods_by_name['Create']._serialized_options = b'\322\363\030\001\001\340\363\030\001'
  _globals['_PARTITIONSERVICE'].methods_by_name['Update']._loaded_options = None
  _globals['_PARTITIONSERVICE'].methods_by_name['Update']._serialized_options = b'\322\363\030\001\001\340\363\030\001'
  _globals['_PARTITIONSERVICE'].methods_by_name['Delete']._loaded_options = None
  _globals['_PARTITIONSERVICE'].methods_by_name['Delete']._serialized_options = b'\322\363\030\001\001\340\363\030\001'
  _globals['_PARTITIONSERVICE'].methods_by_name['Capacity']._loaded_options = None
  _globals['_PARTITIONSERVICE'].methods_by_name['Capacity']._serialized_options = b'\322\363\030\002\002\001\340\363\030\002'
  _globals['_PARTITIONSERVICECREATEREQUEST']._serialized_start=156
  _globals['_PARTITIONSERVICECREATEREQUEST']._serialized_end=247
  _globals['_PARTITIONSERVICEUPDATEREQUEST']._serialized_start=249
  _globals['_PARTITIONSERVICEUPDATEREQUEST']._serialized_end=340
  _globals['_PARTITIONSERVICEDELETEREQUEST']._serialized_start=342
  _globals['_PARTITIONSERVICEDELETEREQUEST']._serialized_end=401
  _globals['_PARTITIONSERVICECREATERESPONSE']._serialized_start=403
  _globals['_PARTITIONSERVICECREATERESPONSE']._serialized_end=495
  _globals['_PARTITIONSERVICEUPDATERESPONSE']._serialized_start=497
  _globals['_PARTITIONSERVICEUPDATERESPONSE']._serialized_end=589
  _globals['_PARTITIONSERVICEDELETERESPONSE']._serialized_start=591
  _globals['_PARTITIONSERVICEDELETERESPONSE']._serialized_end=683
  _globals['_PARTITIONSERVICECAPACITYREQUEST']._serialized_start=686
  _globals['_PARTITIONSERVICECAPACITYREQUEST']._serialized_end=858
  _globals['_PARTITIONSERVICECAPACITYRESPONSE']._serialized_start=861
  _globals['_PARTITIONSERVICECAPACITYRESPONSE']._serialized_end=1376
  _globals['_PARTITIONSERVICE']._serialized_start=1379
  _globals['_PARTITIONSERVICE']._serialized_end=1909
# @@protoc_insertion_point(module_scope)
