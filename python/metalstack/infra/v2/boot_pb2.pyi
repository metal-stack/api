from buf.validate import validate_pb2 as _validate_pb2
from metalstack.api.v2 import common_pb2 as _common_pb2
from metalstack.api.v2 import machine_pb2 as _machine_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class BootServiceDhcpRequest(_message.Message):
    __slots__ = ("uuid",)
    UUID_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    def __init__(self, uuid: _Optional[str] = ...) -> None: ...

class BootServiceDhcpResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class BootServiceBootRequest(_message.Message):
    __slots__ = ("mac", "partition_id")
    MAC_FIELD_NUMBER: _ClassVar[int]
    PARTITION_ID_FIELD_NUMBER: _ClassVar[int]
    mac: str
    partition_id: str
    def __init__(self, mac: _Optional[str] = ..., partition_id: _Optional[str] = ...) -> None: ...

class BootServiceBootResponse(_message.Message):
    __slots__ = ("kernel", "init_ram_disks", "cmdline")
    KERNEL_FIELD_NUMBER: _ClassVar[int]
    INIT_RAM_DISKS_FIELD_NUMBER: _ClassVar[int]
    CMDLINE_FIELD_NUMBER: _ClassVar[int]
    kernel: str
    init_ram_disks: _containers.RepeatedScalarFieldContainer[str]
    cmdline: str
    def __init__(self, kernel: _Optional[str] = ..., init_ram_disks: _Optional[_Iterable[str]] = ..., cmdline: _Optional[str] = ...) -> None: ...

class BootServiceRegisterRequest(_message.Message):
    __slots__ = ("uuid", "hardware", "bios", "ipmi", "tags", "metal_hammer_version", "partition_id")
    UUID_FIELD_NUMBER: _ClassVar[int]
    HARDWARE_FIELD_NUMBER: _ClassVar[int]
    BIOS_FIELD_NUMBER: _ClassVar[int]
    IPMI_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    METAL_HAMMER_VERSION_FIELD_NUMBER: _ClassVar[int]
    PARTITION_ID_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    hardware: _machine_pb2.MachineHardware
    bios: _machine_pb2.MachineBios
    ipmi: _machine_pb2.MachineIPMI
    tags: _containers.RepeatedScalarFieldContainer[str]
    metal_hammer_version: str
    partition_id: str
    def __init__(self, uuid: _Optional[str] = ..., hardware: _Optional[_Union[_machine_pb2.MachineHardware, _Mapping]] = ..., bios: _Optional[_Union[_machine_pb2.MachineBios, _Mapping]] = ..., ipmi: _Optional[_Union[_machine_pb2.MachineIPMI, _Mapping]] = ..., tags: _Optional[_Iterable[str]] = ..., metal_hammer_version: _Optional[str] = ..., partition_id: _Optional[str] = ...) -> None: ...

class BootServiceRegisterResponse(_message.Message):
    __slots__ = ("uuid", "size", "partition_id")
    UUID_FIELD_NUMBER: _ClassVar[int]
    SIZE_FIELD_NUMBER: _ClassVar[int]
    PARTITION_ID_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    size: str
    partition_id: str
    def __init__(self, uuid: _Optional[str] = ..., size: _Optional[str] = ..., partition_id: _Optional[str] = ...) -> None: ...

class BootServiceWaitRequest(_message.Message):
    __slots__ = ("uuid",)
    UUID_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    def __init__(self, uuid: _Optional[str] = ...) -> None: ...

class BootServiceWaitResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class BootServiceReportRequest(_message.Message):
    __slots__ = ("uuid", "console_password", "success", "message")
    UUID_FIELD_NUMBER: _ClassVar[int]
    CONSOLE_PASSWORD_FIELD_NUMBER: _ClassVar[int]
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    console_password: str
    success: bool
    message: str
    def __init__(self, uuid: _Optional[str] = ..., console_password: _Optional[str] = ..., success: _Optional[bool] = ..., message: _Optional[str] = ...) -> None: ...

class BootServiceReportResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class BootServiceSuperUserPasswordRequest(_message.Message):
    __slots__ = ("uuid",)
    UUID_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    def __init__(self, uuid: _Optional[str] = ...) -> None: ...

class BootServiceSuperUserPasswordResponse(_message.Message):
    __slots__ = ("feature_disabled", "super_user_password")
    FEATURE_DISABLED_FIELD_NUMBER: _ClassVar[int]
    SUPER_USER_PASSWORD_FIELD_NUMBER: _ClassVar[int]
    feature_disabled: bool
    super_user_password: str
    def __init__(self, feature_disabled: _Optional[bool] = ..., super_user_password: _Optional[str] = ...) -> None: ...
