from metalstack.api.v2 import common_pb2 as _common_pb2
from metalstack.api.v2 import machine_pb2 as _machine_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class EventServiceSendRequest(_message.Message):
    __slots__ = ("events",)
    class EventsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: _machine_pb2.MachineProvisioningEvent
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[_machine_pb2.MachineProvisioningEvent, _Mapping]] = ...) -> None: ...
    EVENTS_FIELD_NUMBER: _ClassVar[int]
    events: _containers.MessageMap[str, _machine_pb2.MachineProvisioningEvent]
    def __init__(self, events: _Optional[_Mapping[str, _machine_pb2.MachineProvisioningEvent]] = ...) -> None: ...

class EventServiceSendResponse(_message.Message):
    __slots__ = ("events", "failed")
    EVENTS_FIELD_NUMBER: _ClassVar[int]
    FAILED_FIELD_NUMBER: _ClassVar[int]
    events: int
    failed: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, events: _Optional[int] = ..., failed: _Optional[_Iterable[str]] = ...) -> None: ...
