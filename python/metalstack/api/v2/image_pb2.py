# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: metalstack/api/v2/image.proto
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
    'metalstack/api/v2/image.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from buf.validate import validate_pb2 as buf_dot_validate_dot_validate__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from metalstack.api.v2 import common_pb2 as metalstack_dot_api_dot_v2_dot_common__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1dmetalstack/api/v2/image.proto\x12\x11metalstack.api.v2\x1a\x1b\x62uf/validate/validate.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1emetalstack/api/v2/common.proto\"4\n\x16ImageServiceGetRequest\x12\x1a\n\x02id\x18\x01 \x01(\tB\n\xbaH\x07r\x05\x10\x02\x18\x80\x01R\x02id\"N\n\x17ImageServiceListRequest\x12\x33\n\x05query\x18\x01 \x01(\x0b\x32\x1d.metalstack.api.v2.ImageQueryR\x05query\"+\n\x19ImageServiceLatestRequest\x12\x0e\n\x02os\x18\x01 \x01(\tR\x02os\"I\n\x17ImageServiceGetResponse\x12.\n\x05image\x18\x01 \x01(\x0b\x32\x18.metalstack.api.v2.ImageR\x05image\"L\n\x18ImageServiceListResponse\x12\x30\n\x06images\x18\x01 \x03(\x0b\x32\x18.metalstack.api.v2.ImageR\x06images\"L\n\x1aImageServiceLatestResponse\x12.\n\x05image\x18\x01 \x01(\x0b\x32\x18.metalstack.api.v2.ImageR\x05image\"\xa8\x04\n\x05Image\x12\x1a\n\x02id\x18\x01 \x01(\tB\n\xbaH\x07r\x05\x10\x02\x18\x80\x01R\x02id\x12+\n\x04meta\x18\x02 \x01(\x0b\x32\x17.metalstack.api.v2.MetaR\x04meta\x12J\n\x03url\x18\x03 \x01(\tB8\xbaH5\xba\x01\x32\n\tvalid_url\x12\x17url must be a valid URI\x1a\x0cthis.isUri()R\x03url\x12#\n\x04name\x18\x04 \x01(\tB\n\xbaH\x07r\x05\x10\x02\x18\x80\x01H\x00R\x04name\x88\x01\x01\x12\x31\n\x0b\x64\x65scription\x18\x05 \x01(\tB\n\xbaH\x07r\x05\x10\x02\x18\x80\x01H\x01R\x0b\x64\x65scription\x88\x01\x01\x12\x83\x01\n\x08\x66\x65\x61tures\x18\x06 \x03(\x0e\x32\x1f.metalstack.api.v2.ImageFeatureBF\xbaHC\x92\x01@\x08\x01\"<\xba\x01\x39\n\x08\x66\x65\x61tures\x12\x15\x66\x65\x61ture must be valid\x1a\x16this >= 0 && this <= 2R\x08\x66\x65\x61tures\x12X\n\x0e\x63lassification\x18\x07 \x01(\x0e\x32&.metalstack.api.v2.ImageClassificationB\x08\xbaH\x05\x82\x01\x02\x10\x01R\x0e\x63lassification\x12\x39\n\nexpires_at\x18\x08 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\texpiresAtB\x07\n\x05_nameB\x0e\n\x0c_description\"U\n\nImageUsage\x12.\n\x05image\x18\x01 \x01(\x0b\x32\x18.metalstack.api.v2.ImageR\x05image\x12\x17\n\x07used_by\x18\x02 \x03(\tR\x06usedBy\"\xcc\x03\n\nImageQuery\x12\x1f\n\x02id\x18\x01 \x01(\tB\n\xbaH\x07r\x05\x10\x02\x18\x80\x01H\x00R\x02id\x88\x01\x01\x12\x1f\n\x02os\x18\x02 \x01(\tB\n\xbaH\x07r\x05\x10\x02\x18\x80\x01H\x01R\x02os\x88\x01\x01\x12)\n\x07version\x18\x03 \x01(\tB\n\xbaH\x07r\x05\x10\x01\x18\x80\x01H\x02R\x07version\x88\x01\x01\x12#\n\x04name\x18\x04 \x01(\tB\n\xbaH\x07r\x05\x10\x02\x18\x80\x01H\x03R\x04name\x88\x01\x01\x12\x31\n\x0b\x64\x65scription\x18\x05 \x01(\tB\n\xbaH\x07r\x05\x10\x02\x18\x80\x01H\x04R\x0b\x64\x65scription\x88\x01\x01\x12H\n\x07\x66\x65\x61ture\x18\x06 \x01(\x0e\x32\x1f.metalstack.api.v2.ImageFeatureB\x08\xbaH\x05\x82\x01\x02\x10\x01H\x05R\x07\x66\x65\x61ture\x88\x01\x01\x12]\n\x0e\x63lassification\x18\x07 \x01(\x0e\x32&.metalstack.api.v2.ImageClassificationB\x08\xbaH\x05\x82\x01\x02\x10\x01H\x06R\x0e\x63lassification\x88\x01\x01\x42\x05\n\x03_idB\x05\n\x03_osB\n\n\x08_versionB\x07\n\x05_nameB\x0e\n\x0c_descriptionB\n\n\x08_featureB\x11\n\x0f_classification*\x7f\n\x0cImageFeature\x12\x1d\n\x19IMAGE_FEATURE_UNSPECIFIED\x10\x00\x12&\n\x15IMAGE_FEATURE_MACHINE\x10\x01\x1a\x0b\x82\xb2\x19\x07machine\x12(\n\x16IMAGE_FEATURE_FIREWALL\x10\x02\x1a\x0c\x82\xb2\x19\x08\x66irewall*\xd2\x01\n\x13ImageClassification\x12$\n IMAGE_CLASSIFICATION_UNSPECIFIED\x10\x00\x12-\n\x1cIMAGE_CLASSIFICATION_PREVIEW\x10\x01\x1a\x0b\x82\xb2\x19\x07preview\x12\x31\n\x1eIMAGE_CLASSIFICATION_SUPPORTED\x10\x02\x1a\r\x82\xb2\x19\tsupported\x12\x33\n\x1fIMAGE_CLASSIFICATION_DEPRECATED\x10\x03\x1a\x0e\x82\xb2\x19\ndeprecated2\xd2\x02\n\x0cImageService\x12\x66\n\x03Get\x12).metalstack.api.v2.ImageServiceGetRequest\x1a*.metalstack.api.v2.ImageServiceGetResponse\"\x08\xd8\xf3\x18\x03\xe0\xf3\x18\x02\x12i\n\x04List\x12*.metalstack.api.v2.ImageServiceListRequest\x1a+.metalstack.api.v2.ImageServiceListResponse\"\x08\xd8\xf3\x18\x03\xe0\xf3\x18\x02\x12o\n\x06Latest\x12,.metalstack.api.v2.ImageServiceLatestRequest\x1a-.metalstack.api.v2.ImageServiceLatestResponse\"\x08\xd8\xf3\x18\x03\xe0\xf3\x18\x02\x42\xc0\x01\n\x15\x63om.metalstack.api.v2B\nImageProtoP\x01Z5github.com/metal-stack/api/go/metalstack/api/v2;apiv2\xa2\x02\x03MAX\xaa\x02\x11Metalstack.Api.V2\xca\x02\x11Metalstack\\Api\\V2\xe2\x02\x1dMetalstack\\Api\\V2\\GPBMetadata\xea\x02\x13Metalstack::Api::V2b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'metalstack.api.v2.image_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\025com.metalstack.api.v2B\nImageProtoP\001Z5github.com/metal-stack/api/go/metalstack/api/v2;apiv2\242\002\003MAX\252\002\021Metalstack.Api.V2\312\002\021Metalstack\\Api\\V2\342\002\035Metalstack\\Api\\V2\\GPBMetadata\352\002\023Metalstack::Api::V2'
  _globals['_IMAGEFEATURE'].values_by_name["IMAGE_FEATURE_MACHINE"]._loaded_options = None
  _globals['_IMAGEFEATURE'].values_by_name["IMAGE_FEATURE_MACHINE"]._serialized_options = b'\202\262\031\007machine'
  _globals['_IMAGEFEATURE'].values_by_name["IMAGE_FEATURE_FIREWALL"]._loaded_options = None
  _globals['_IMAGEFEATURE'].values_by_name["IMAGE_FEATURE_FIREWALL"]._serialized_options = b'\202\262\031\010firewall'
  _globals['_IMAGECLASSIFICATION'].values_by_name["IMAGE_CLASSIFICATION_PREVIEW"]._loaded_options = None
  _globals['_IMAGECLASSIFICATION'].values_by_name["IMAGE_CLASSIFICATION_PREVIEW"]._serialized_options = b'\202\262\031\007preview'
  _globals['_IMAGECLASSIFICATION'].values_by_name["IMAGE_CLASSIFICATION_SUPPORTED"]._loaded_options = None
  _globals['_IMAGECLASSIFICATION'].values_by_name["IMAGE_CLASSIFICATION_SUPPORTED"]._serialized_options = b'\202\262\031\tsupported'
  _globals['_IMAGECLASSIFICATION'].values_by_name["IMAGE_CLASSIFICATION_DEPRECATED"]._loaded_options = None
  _globals['_IMAGECLASSIFICATION'].values_by_name["IMAGE_CLASSIFICATION_DEPRECATED"]._serialized_options = b'\202\262\031\ndeprecated'
  _globals['_IMAGESERVICEGETREQUEST'].fields_by_name['id']._loaded_options = None
  _globals['_IMAGESERVICEGETREQUEST'].fields_by_name['id']._serialized_options = b'\272H\007r\005\020\002\030\200\001'
  _globals['_IMAGE'].fields_by_name['id']._loaded_options = None
  _globals['_IMAGE'].fields_by_name['id']._serialized_options = b'\272H\007r\005\020\002\030\200\001'
  _globals['_IMAGE'].fields_by_name['url']._loaded_options = None
  _globals['_IMAGE'].fields_by_name['url']._serialized_options = b'\272H5\272\0012\n\tvalid_url\022\027url must be a valid URI\032\014this.isUri()'
  _globals['_IMAGE'].fields_by_name['name']._loaded_options = None
  _globals['_IMAGE'].fields_by_name['name']._serialized_options = b'\272H\007r\005\020\002\030\200\001'
  _globals['_IMAGE'].fields_by_name['description']._loaded_options = None
  _globals['_IMAGE'].fields_by_name['description']._serialized_options = b'\272H\007r\005\020\002\030\200\001'
  _globals['_IMAGE'].fields_by_name['features']._loaded_options = None
  _globals['_IMAGE'].fields_by_name['features']._serialized_options = b'\272HC\222\001@\010\001\"<\272\0019\n\010features\022\025feature must be valid\032\026this >= 0 && this <= 2'
  _globals['_IMAGE'].fields_by_name['classification']._loaded_options = None
  _globals['_IMAGE'].fields_by_name['classification']._serialized_options = b'\272H\005\202\001\002\020\001'
  _globals['_IMAGEQUERY'].fields_by_name['id']._loaded_options = None
  _globals['_IMAGEQUERY'].fields_by_name['id']._serialized_options = b'\272H\007r\005\020\002\030\200\001'
  _globals['_IMAGEQUERY'].fields_by_name['os']._loaded_options = None
  _globals['_IMAGEQUERY'].fields_by_name['os']._serialized_options = b'\272H\007r\005\020\002\030\200\001'
  _globals['_IMAGEQUERY'].fields_by_name['version']._loaded_options = None
  _globals['_IMAGEQUERY'].fields_by_name['version']._serialized_options = b'\272H\007r\005\020\001\030\200\001'
  _globals['_IMAGEQUERY'].fields_by_name['name']._loaded_options = None
  _globals['_IMAGEQUERY'].fields_by_name['name']._serialized_options = b'\272H\007r\005\020\002\030\200\001'
  _globals['_IMAGEQUERY'].fields_by_name['description']._loaded_options = None
  _globals['_IMAGEQUERY'].fields_by_name['description']._serialized_options = b'\272H\007r\005\020\002\030\200\001'
  _globals['_IMAGEQUERY'].fields_by_name['feature']._loaded_options = None
  _globals['_IMAGEQUERY'].fields_by_name['feature']._serialized_options = b'\272H\005\202\001\002\020\001'
  _globals['_IMAGEQUERY'].fields_by_name['classification']._loaded_options = None
  _globals['_IMAGEQUERY'].fields_by_name['classification']._serialized_options = b'\272H\005\202\001\002\020\001'
  _globals['_IMAGESERVICE'].methods_by_name['Get']._loaded_options = None
  _globals['_IMAGESERVICE'].methods_by_name['Get']._serialized_options = b'\330\363\030\003\340\363\030\002'
  _globals['_IMAGESERVICE'].methods_by_name['List']._loaded_options = None
  _globals['_IMAGESERVICE'].methods_by_name['List']._serialized_options = b'\330\363\030\003\340\363\030\002'
  _globals['_IMAGESERVICE'].methods_by_name['Latest']._loaded_options = None
  _globals['_IMAGESERVICE'].methods_by_name['Latest']._serialized_options = b'\330\363\030\003\340\363\030\002'
  _globals['_IMAGEFEATURE']._serialized_start=1661
  _globals['_IMAGEFEATURE']._serialized_end=1788
  _globals['_IMAGECLASSIFICATION']._serialized_start=1791
  _globals['_IMAGECLASSIFICATION']._serialized_end=2001
  _globals['_IMAGESERVICEGETREQUEST']._serialized_start=146
  _globals['_IMAGESERVICEGETREQUEST']._serialized_end=198
  _globals['_IMAGESERVICELISTREQUEST']._serialized_start=200
  _globals['_IMAGESERVICELISTREQUEST']._serialized_end=278
  _globals['_IMAGESERVICELATESTREQUEST']._serialized_start=280
  _globals['_IMAGESERVICELATESTREQUEST']._serialized_end=323
  _globals['_IMAGESERVICEGETRESPONSE']._serialized_start=325
  _globals['_IMAGESERVICEGETRESPONSE']._serialized_end=398
  _globals['_IMAGESERVICELISTRESPONSE']._serialized_start=400
  _globals['_IMAGESERVICELISTRESPONSE']._serialized_end=476
  _globals['_IMAGESERVICELATESTRESPONSE']._serialized_start=478
  _globals['_IMAGESERVICELATESTRESPONSE']._serialized_end=554
  _globals['_IMAGE']._serialized_start=557
  _globals['_IMAGE']._serialized_end=1109
  _globals['_IMAGEUSAGE']._serialized_start=1111
  _globals['_IMAGEUSAGE']._serialized_end=1196
  _globals['_IMAGEQUERY']._serialized_start=1199
  _globals['_IMAGEQUERY']._serialized_end=1659
  _globals['_IMAGESERVICE']._serialized_start=2004
  _globals['_IMAGESERVICE']._serialized_end=2342
# @@protoc_insertion_point(module_scope)
