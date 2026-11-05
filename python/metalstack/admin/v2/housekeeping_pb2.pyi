from buf.validate import validate_pb2 as _validate_pb2
from metalstack.api.v2 import common_pb2 as _common_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class MachineHousekeepingAction(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    MACHINE_HOUSEKEEPING_ACTION_UNSPECIFIED: _ClassVar[MachineHousekeepingAction]
    MACHINE_HOUSEKEEPING_ACTION_VPN_CONNECTED: _ClassVar[MachineHousekeepingAction]
    MACHINE_HOUSEKEEPING_ACTION_LIVELINESS: _ClassVar[MachineHousekeepingAction]
    MACHINE_HOUSEKEEPING_ACTION_RESURRECT: _ClassVar[MachineHousekeepingAction]
MACHINE_HOUSEKEEPING_ACTION_UNSPECIFIED: MachineHousekeepingAction
MACHINE_HOUSEKEEPING_ACTION_VPN_CONNECTED: MachineHousekeepingAction
MACHINE_HOUSEKEEPING_ACTION_LIVELINESS: MachineHousekeepingAction
MACHINE_HOUSEKEEPING_ACTION_RESURRECT: MachineHousekeepingAction

class HousekeepingServiceMachineRequest(_message.Message):
    __slots__ = ("action",)
    ACTION_FIELD_NUMBER: _ClassVar[int]
    action: MachineHousekeepingAction
    def __init__(self, action: _Optional[_Union[MachineHousekeepingAction, str]] = ...) -> None: ...

class HousekeepingServiceMachineResponse(_message.Message):
    __slots__ = ("action", "modified_machine_ids")
    ACTION_FIELD_NUMBER: _ClassVar[int]
    MODIFIED_MACHINE_IDS_FIELD_NUMBER: _ClassVar[int]
    action: MachineHousekeepingAction
    modified_machine_ids: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, action: _Optional[_Union[MachineHousekeepingAction, str]] = ..., modified_machine_ids: _Optional[_Iterable[str]] = ...) -> None: ...
