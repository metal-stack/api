from buf.validate import validate_pb2 as _validate_pb2
from metalstack.api.v2 import common_pb2 as _common_pb2
from metalstack.api.v2 import switch_pb2 as _switch_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SwitchServiceRegisterRequest(_message.Message):
    __slots__ = ("switch",)
    SWITCH_FIELD_NUMBER: _ClassVar[int]
    switch: _switch_pb2.Switch
    def __init__(self, switch: _Optional[_Union[_switch_pb2.Switch, _Mapping]] = ...) -> None: ...

class SwitchServiceRegisterResponse(_message.Message):
    __slots__ = ("switch",)
    SWITCH_FIELD_NUMBER: _ClassVar[int]
    switch: _switch_pb2.Switch
    def __init__(self, switch: _Optional[_Union[_switch_pb2.Switch, _Mapping]] = ...) -> None: ...

class SwitchServiceReportBGPRoutesRequest(_message.Message):
    __slots__ = ("switch_id", "bgp_routes")
    SWITCH_ID_FIELD_NUMBER: _ClassVar[int]
    BGP_ROUTES_FIELD_NUMBER: _ClassVar[int]
    switch_id: str
    bgp_routes: _containers.RepeatedCompositeFieldContainer[BGPRoute]
    def __init__(self, switch_id: _Optional[str] = ..., bgp_routes: _Optional[_Iterable[_Union[BGPRoute, _Mapping]]] = ...) -> None: ...

class SwitchServiceReportBGPRoutesResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class BGPRoute(_message.Message):
    __slots__ = ("cidr",)
    CIDR_FIELD_NUMBER: _ClassVar[int]
    cidr: str
    def __init__(self, cidr: _Optional[str] = ...) -> None: ...
