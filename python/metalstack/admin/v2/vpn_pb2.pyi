import datetime

from buf.validate import validate_pb2 as _validate_pb2
from google.protobuf import duration_pb2 as _duration_pb2
from metalstack.api.v2 import common_pb2 as _common_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class VPNServiceAuthkeyRequest(_message.Message):
    __slots__ = ("project", "ephemeral", "expires")
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    EPHEMERAL_FIELD_NUMBER: _ClassVar[int]
    EXPIRES_FIELD_NUMBER: _ClassVar[int]
    project: str
    ephemeral: bool
    expires: _duration_pb2.Duration
    def __init__(self, project: _Optional[str] = ..., ephemeral: _Optional[bool] = ..., expires: _Optional[_Union[datetime.timedelta, _duration_pb2.Duration, _Mapping]] = ...) -> None: ...

class VPNServiceAuthkeyResponse(_message.Message):
    __slots__ = ("address", "authkey")
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    AUTHKEY_FIELD_NUMBER: _ClassVar[int]
    address: str
    authkey: str
    def __init__(self, address: _Optional[str] = ..., authkey: _Optional[str] = ...) -> None: ...
