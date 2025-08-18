from buf.validate import validate_pb2 as _validate_pb2
from metalstack.api.v2 import common_pb2 as _common_pb2
from metalstack.api.v2 import switch_pb2 as _switch_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SwitchServiceUpdateRequest(_message.Message):
    __slots__ = ("id", "description", "rack_id", "replace_mode", "management_ip", "management_user", "console_command", "nics", "os")
    ID_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    RACK_ID_FIELD_NUMBER: _ClassVar[int]
    REPLACE_MODE_FIELD_NUMBER: _ClassVar[int]
    MANAGEMENT_IP_FIELD_NUMBER: _ClassVar[int]
    MANAGEMENT_USER_FIELD_NUMBER: _ClassVar[int]
    CONSOLE_COMMAND_FIELD_NUMBER: _ClassVar[int]
    NICS_FIELD_NUMBER: _ClassVar[int]
    OS_FIELD_NUMBER: _ClassVar[int]
    id: str
    description: str
    rack_id: str
    replace_mode: _switch_pb2.SwitchReplaceMode
    management_ip: str
    management_user: str
    console_command: str
    nics: _containers.RepeatedCompositeFieldContainer[_switch_pb2.Nic]
    os: _switch_pb2.SwitchOS
    def __init__(self, id: _Optional[str] = ..., description: _Optional[str] = ..., rack_id: _Optional[str] = ..., replace_mode: _Optional[_Union[_switch_pb2.SwitchReplaceMode, str]] = ..., management_ip: _Optional[str] = ..., management_user: _Optional[str] = ..., console_command: _Optional[str] = ..., nics: _Optional[_Iterable[_Union[_switch_pb2.Nic, _Mapping]]] = ..., os: _Optional[_Union[_switch_pb2.SwitchOS, _Mapping]] = ...) -> None: ...

class SwitchServiceUpdateResponse(_message.Message):
    __slots__ = ("switch",)
    SWITCH_FIELD_NUMBER: _ClassVar[int]
    switch: _switch_pb2.Switch
    def __init__(self, switch: _Optional[_Union[_switch_pb2.Switch, _Mapping]] = ...) -> None: ...
