import datetime

from buf.validate import validate_pb2 as _validate_pb2
from google.protobuf import duration_pb2 as _duration_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from metalstack.api.v2 import common_pb2 as _common_pb2
from metalstack.api.v2 import predefined_rules_pb2 as _predefined_rules_pb2
from metalstack.api.v2 import token_pb2 as _token_pb2
from metalstack.api.v2 import version_pb2 as _version_pb2
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ComponentType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    COMPONENT_TYPE_UNSPECIFIED: _ClassVar[ComponentType]
    COMPONENT_TYPE_PIXIECORE: _ClassVar[ComponentType]
    COMPONENT_TYPE_METAL_CORE: _ClassVar[ComponentType]
    COMPONENT_TYPE_METAL_BMC: _ClassVar[ComponentType]
    COMPONENT_TYPE_METAL_IMAGE_CACHE_SYNC: _ClassVar[ComponentType]
    COMPONENT_TYPE_METAL_CONSOLE: _ClassVar[ComponentType]
    COMPONENT_TYPE_METAL_METRICS_EXPORTER: _ClassVar[ComponentType]
COMPONENT_TYPE_UNSPECIFIED: ComponentType
COMPONENT_TYPE_PIXIECORE: ComponentType
COMPONENT_TYPE_METAL_CORE: ComponentType
COMPONENT_TYPE_METAL_BMC: ComponentType
COMPONENT_TYPE_METAL_IMAGE_CACHE_SYNC: ComponentType
COMPONENT_TYPE_METAL_CONSOLE: ComponentType
COMPONENT_TYPE_METAL_METRICS_EXPORTER: ComponentType

class Component(_message.Message):
    __slots__ = ("uuid", "type", "identifier", "started_at", "reported_at", "interval", "version", "token")
    UUID_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    IDENTIFIER_FIELD_NUMBER: _ClassVar[int]
    STARTED_AT_FIELD_NUMBER: _ClassVar[int]
    REPORTED_AT_FIELD_NUMBER: _ClassVar[int]
    INTERVAL_FIELD_NUMBER: _ClassVar[int]
    VERSION_FIELD_NUMBER: _ClassVar[int]
    TOKEN_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    type: ComponentType
    identifier: str
    started_at: _timestamp_pb2.Timestamp
    reported_at: _timestamp_pb2.Timestamp
    interval: _duration_pb2.Duration
    version: _version_pb2.Version
    token: _token_pb2.Token
    def __init__(self, uuid: _Optional[str] = ..., type: _Optional[_Union[ComponentType, str]] = ..., identifier: _Optional[str] = ..., started_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ..., reported_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ..., interval: _Optional[_Union[datetime.timedelta, _duration_pb2.Duration, _Mapping]] = ..., version: _Optional[_Union[_version_pb2.Version, _Mapping]] = ..., token: _Optional[_Union[_token_pb2.Token, _Mapping]] = ...) -> None: ...

class ComponentQuery(_message.Message):
    __slots__ = ("uuid", "type", "identifier")
    UUID_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    IDENTIFIER_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    type: ComponentType
    identifier: str
    def __init__(self, uuid: _Optional[str] = ..., type: _Optional[_Union[ComponentType, str]] = ..., identifier: _Optional[str] = ...) -> None: ...
