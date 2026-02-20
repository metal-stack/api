import datetime

from buf.validate import validate_pb2 as _validate_pb2
from google.protobuf import duration_pb2 as _duration_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from metalstack.api.v2 import common_pb2 as _common_pb2
from metalstack.api.v2 import component_pb2 as _component_pb2
from metalstack.api.v2 import predefined_rules_pb2 as _predefined_rules_pb2
from metalstack.api.v2 import version_pb2 as _version_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ComponentServicePingRequest(_message.Message):
    __slots__ = ("type", "identifier", "started_at", "interval", "version")
    TYPE_FIELD_NUMBER: _ClassVar[int]
    IDENTIFIER_FIELD_NUMBER: _ClassVar[int]
    STARTED_AT_FIELD_NUMBER: _ClassVar[int]
    INTERVAL_FIELD_NUMBER: _ClassVar[int]
    VERSION_FIELD_NUMBER: _ClassVar[int]
    type: _component_pb2.ComponentType
    identifier: str
    started_at: _timestamp_pb2.Timestamp
    interval: _duration_pb2.Duration
    version: _version_pb2.Version
    def __init__(self, type: _Optional[_Union[_component_pb2.ComponentType, str]] = ..., identifier: _Optional[str] = ..., started_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ..., interval: _Optional[_Union[datetime.timedelta, _duration_pb2.Duration, _Mapping]] = ..., version: _Optional[_Union[_version_pb2.Version, _Mapping]] = ...) -> None: ...

class ComponentServicePingResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
