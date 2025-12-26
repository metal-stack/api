from buf.validate import validate_pb2 as _validate_pb2
from metalstack.api.v2 import common_pb2 as _common_pb2
from metalstack.api.v2 import machine_pb2 as _machine_pb2
from metalstack.api.v2 import predefined_rules_pb2 as _predefined_rules_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class UpdateBMCInfoRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class UpdateBMCInfoResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class WaitForMachineEventRequest(_message.Message):
    __slots__ = ("partition",)
    PARTITION_FIELD_NUMBER: _ClassVar[int]
    partition: str
    def __init__(self, partition: _Optional[str] = ...) -> None: ...

class WaitForMachineEventResponse(_message.Message):
    __slots__ = ("bmc_command", "machine_ipmi")
    BMC_COMMAND_FIELD_NUMBER: _ClassVar[int]
    MACHINE_IPMI_FIELD_NUMBER: _ClassVar[int]
    bmc_command: _machine_pb2.MachineBMCCommand
    machine_ipmi: _machine_pb2.MachineIPMI
    def __init__(self, bmc_command: _Optional[_Union[_machine_pb2.MachineBMCCommand, str]] = ..., machine_ipmi: _Optional[_Union[_machine_pb2.MachineIPMI, _Mapping]] = ...) -> None: ...
