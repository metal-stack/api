from metalstack.api.v2 import common_pb2 as _common_pb2
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
    hardware: MachineHardware
    bios: MachineBIOS
    ipmi: MachineIPMI
    tags: _containers.RepeatedScalarFieldContainer[str]
    metal_hammer_version: str
    partition_id: str
    def __init__(self, uuid: _Optional[str] = ..., hardware: _Optional[_Union[MachineHardware, _Mapping]] = ..., bios: _Optional[_Union[MachineBIOS, _Mapping]] = ..., ipmi: _Optional[_Union[MachineIPMI, _Mapping]] = ..., tags: _Optional[_Iterable[str]] = ..., metal_hammer_version: _Optional[str] = ..., partition_id: _Optional[str] = ...) -> None: ...

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

class MachineHardware(_message.Message):
    __slots__ = ("memory", "disks", "nics", "cpus", "gpus")
    MEMORY_FIELD_NUMBER: _ClassVar[int]
    DISKS_FIELD_NUMBER: _ClassVar[int]
    NICS_FIELD_NUMBER: _ClassVar[int]
    CPUS_FIELD_NUMBER: _ClassVar[int]
    GPUS_FIELD_NUMBER: _ClassVar[int]
    memory: int
    disks: _containers.RepeatedCompositeFieldContainer[MachineBlockDevice]
    nics: _containers.RepeatedCompositeFieldContainer[MachineNic]
    cpus: _containers.RepeatedCompositeFieldContainer[MachineCPU]
    gpus: _containers.RepeatedCompositeFieldContainer[MachineGPU]
    def __init__(self, memory: _Optional[int] = ..., disks: _Optional[_Iterable[_Union[MachineBlockDevice, _Mapping]]] = ..., nics: _Optional[_Iterable[_Union[MachineNic, _Mapping]]] = ..., cpus: _Optional[_Iterable[_Union[MachineCPU, _Mapping]]] = ..., gpus: _Optional[_Iterable[_Union[MachineGPU, _Mapping]]] = ...) -> None: ...

class MachineCPU(_message.Message):
    __slots__ = ("vendor", "model", "cores", "threads")
    VENDOR_FIELD_NUMBER: _ClassVar[int]
    MODEL_FIELD_NUMBER: _ClassVar[int]
    CORES_FIELD_NUMBER: _ClassVar[int]
    THREADS_FIELD_NUMBER: _ClassVar[int]
    vendor: str
    model: str
    cores: int
    threads: int
    def __init__(self, vendor: _Optional[str] = ..., model: _Optional[str] = ..., cores: _Optional[int] = ..., threads: _Optional[int] = ...) -> None: ...

class MachineGPU(_message.Message):
    __slots__ = ("vendor", "model")
    VENDOR_FIELD_NUMBER: _ClassVar[int]
    MODEL_FIELD_NUMBER: _ClassVar[int]
    vendor: str
    model: str
    def __init__(self, vendor: _Optional[str] = ..., model: _Optional[str] = ...) -> None: ...

class MachineNic(_message.Message):
    __slots__ = ("mac", "name", "vendor", "model", "speed", "neighbors", "hostname", "identifier")
    MAC_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    VENDOR_FIELD_NUMBER: _ClassVar[int]
    MODEL_FIELD_NUMBER: _ClassVar[int]
    SPEED_FIELD_NUMBER: _ClassVar[int]
    NEIGHBORS_FIELD_NUMBER: _ClassVar[int]
    HOSTNAME_FIELD_NUMBER: _ClassVar[int]
    IDENTIFIER_FIELD_NUMBER: _ClassVar[int]
    mac: str
    name: str
    vendor: str
    model: str
    speed: int
    neighbors: _containers.RepeatedCompositeFieldContainer[MachineNic]
    hostname: str
    identifier: str
    def __init__(self, mac: _Optional[str] = ..., name: _Optional[str] = ..., vendor: _Optional[str] = ..., model: _Optional[str] = ..., speed: _Optional[int] = ..., neighbors: _Optional[_Iterable[_Union[MachineNic, _Mapping]]] = ..., hostname: _Optional[str] = ..., identifier: _Optional[str] = ...) -> None: ...

class MachineBlockDevice(_message.Message):
    __slots__ = ("name", "size")
    NAME_FIELD_NUMBER: _ClassVar[int]
    SIZE_FIELD_NUMBER: _ClassVar[int]
    name: str
    size: int
    def __init__(self, name: _Optional[str] = ..., size: _Optional[int] = ...) -> None: ...

class MachineBIOS(_message.Message):
    __slots__ = ("version", "vendor", "date")
    VERSION_FIELD_NUMBER: _ClassVar[int]
    VENDOR_FIELD_NUMBER: _ClassVar[int]
    DATE_FIELD_NUMBER: _ClassVar[int]
    version: str
    vendor: str
    date: str
    def __init__(self, version: _Optional[str] = ..., vendor: _Optional[str] = ..., date: _Optional[str] = ...) -> None: ...

class MachineIPMI(_message.Message):
    __slots__ = ("address", "mac", "user", "password", "interface", "fru", "bmc_version", "power_state")
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    MAC_FIELD_NUMBER: _ClassVar[int]
    USER_FIELD_NUMBER: _ClassVar[int]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    INTERFACE_FIELD_NUMBER: _ClassVar[int]
    FRU_FIELD_NUMBER: _ClassVar[int]
    BMC_VERSION_FIELD_NUMBER: _ClassVar[int]
    POWER_STATE_FIELD_NUMBER: _ClassVar[int]
    address: str
    mac: str
    user: str
    password: str
    interface: str
    fru: MachineFRU
    bmc_version: str
    power_state: str
    def __init__(self, address: _Optional[str] = ..., mac: _Optional[str] = ..., user: _Optional[str] = ..., password: _Optional[str] = ..., interface: _Optional[str] = ..., fru: _Optional[_Union[MachineFRU, _Mapping]] = ..., bmc_version: _Optional[str] = ..., power_state: _Optional[str] = ...) -> None: ...

class MachineFRU(_message.Message):
    __slots__ = ("chassis_part_number", "chassis_part_serial", "board_mfg", "board_mfg_serial", "board_part_number", "product_manufacturer", "product_part_number", "product_serial")
    CHASSIS_PART_NUMBER_FIELD_NUMBER: _ClassVar[int]
    CHASSIS_PART_SERIAL_FIELD_NUMBER: _ClassVar[int]
    BOARD_MFG_FIELD_NUMBER: _ClassVar[int]
    BOARD_MFG_SERIAL_FIELD_NUMBER: _ClassVar[int]
    BOARD_PART_NUMBER_FIELD_NUMBER: _ClassVar[int]
    PRODUCT_MANUFACTURER_FIELD_NUMBER: _ClassVar[int]
    PRODUCT_PART_NUMBER_FIELD_NUMBER: _ClassVar[int]
    PRODUCT_SERIAL_FIELD_NUMBER: _ClassVar[int]
    chassis_part_number: str
    chassis_part_serial: str
    board_mfg: str
    board_mfg_serial: str
    board_part_number: str
    product_manufacturer: str
    product_part_number: str
    product_serial: str
    def __init__(self, chassis_part_number: _Optional[str] = ..., chassis_part_serial: _Optional[str] = ..., board_mfg: _Optional[str] = ..., board_mfg_serial: _Optional[str] = ..., board_part_number: _Optional[str] = ..., product_manufacturer: _Optional[str] = ..., product_part_number: _Optional[str] = ..., product_serial: _Optional[str] = ...) -> None: ...

class BootServiceReportRequest(_message.Message):
    __slots__ = ("uuid", "console_password", "boot_info", "success", "message")
    UUID_FIELD_NUMBER: _ClassVar[int]
    CONSOLE_PASSWORD_FIELD_NUMBER: _ClassVar[int]
    BOOT_INFO_FIELD_NUMBER: _ClassVar[int]
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    console_password: str
    boot_info: BootInfo
    success: bool
    message: str
    def __init__(self, uuid: _Optional[str] = ..., console_password: _Optional[str] = ..., boot_info: _Optional[_Union[BootInfo, _Mapping]] = ..., success: bool = ..., message: _Optional[str] = ...) -> None: ...

class BootServiceReportResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class BootInfo(_message.Message):
    __slots__ = ("image_id", "primary_disk", "os_partition", "initrd", "cmdline", "kernel", "bootloader_id")
    IMAGE_ID_FIELD_NUMBER: _ClassVar[int]
    PRIMARY_DISK_FIELD_NUMBER: _ClassVar[int]
    OS_PARTITION_FIELD_NUMBER: _ClassVar[int]
    INITRD_FIELD_NUMBER: _ClassVar[int]
    CMDLINE_FIELD_NUMBER: _ClassVar[int]
    KERNEL_FIELD_NUMBER: _ClassVar[int]
    BOOTLOADER_ID_FIELD_NUMBER: _ClassVar[int]
    image_id: str
    primary_disk: str
    os_partition: str
    initrd: str
    cmdline: str
    kernel: str
    bootloader_id: str
    def __init__(self, image_id: _Optional[str] = ..., primary_disk: _Optional[str] = ..., os_partition: _Optional[str] = ..., initrd: _Optional[str] = ..., cmdline: _Optional[str] = ..., kernel: _Optional[str] = ..., bootloader_id: _Optional[str] = ...) -> None: ...

class BootServiceAbortReinstallRequest(_message.Message):
    __slots__ = ("uuid", "primary_disk_wiped")
    UUID_FIELD_NUMBER: _ClassVar[int]
    PRIMARY_DISK_WIPED_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    primary_disk_wiped: bool
    def __init__(self, uuid: _Optional[str] = ..., primary_disk_wiped: bool = ...) -> None: ...

class BootServiceAbortReinstallResponse(_message.Message):
    __slots__ = ("boot_info",)
    BOOT_INFO_FIELD_NUMBER: _ClassVar[int]
    boot_info: BootInfo
    def __init__(self, boot_info: _Optional[_Union[BootInfo, _Mapping]] = ...) -> None: ...

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
    def __init__(self, feature_disabled: bool = ..., super_user_password: _Optional[str] = ...) -> None: ...
