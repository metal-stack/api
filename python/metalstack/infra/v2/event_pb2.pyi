import datetime

from buf.validate import validate_pb2 as _validate_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from metalstack.api.v2 import common_pb2 as _common_pb2
from metalstack.api.v2 import machine_pb2 as _machine_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class EventServiceSendRequest(_message.Message):
    __slots__ = ("uuid", "event")
    UUID_FIELD_NUMBER: _ClassVar[int]
    EVENT_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    event: MachineProvisioningEvent
    def __init__(self, uuid: _Optional[str] = ..., event: _Optional[_Union[MachineProvisioningEvent, _Mapping]] = ...) -> None: ...

class EventServiceSendResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class EventServiceSendMultiRequest(_message.Message):
    __slots__ = ("events",)
    class EventsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: MachineProvisioningEvent
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[MachineProvisioningEvent, _Mapping]] = ...) -> None: ...
    EVENTS_FIELD_NUMBER: _ClassVar[int]
    events: _containers.MessageMap[str, MachineProvisioningEvent]
    def __init__(self, events: _Optional[_Mapping[str, MachineProvisioningEvent]] = ...) -> None: ...

class EventServiceSendMultiResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class MachineProvisioningEvent(_message.Message):
    __slots__ = ("time", "event", "message")
    TIME_FIELD_NUMBER: _ClassVar[int]
    EVENT_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    time: _timestamp_pb2.Timestamp
    event: _machine_pb2.MachineProvisioningEventType
    message: str
    def __init__(self, time: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ..., event: _Optional[_Union[_machine_pb2.MachineProvisioningEventType, str]] = ..., message: _Optional[str] = ...) -> None: ...
