from buf.validate import validate_pb2 as _validate_pb2
from metalstack.api.v2 import common_pb2 as _common_pb2
from metalstack.api.v2 import ip_pb2 as _ip_pb2
from metalstack.api.v2 import machine_pb2 as _machine_pb2
from metalstack.api.v2 import network_pb2 as _network_pb2
from metalstack.api.v2 import vpn_pb2 as _vpn_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class HousekeepingServiceMachineLivelinessRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class HousekeepingServiceMachineLivelinessResponse(_message.Message):
    __slots__ = ("modified_machines",)
    MODIFIED_MACHINES_FIELD_NUMBER: _ClassVar[int]
    modified_machines: _containers.RepeatedCompositeFieldContainer[MachineWithLiveliness]
    def __init__(self, modified_machines: _Optional[_Iterable[_Union[MachineWithLiveliness, _Mapping]]] = ...) -> None: ...

class MachineWithLiveliness(_message.Message):
    __slots__ = ("uuid", "liveliness")
    UUID_FIELD_NUMBER: _ClassVar[int]
    LIVELINESS_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    liveliness: _machine_pb2.MachineLiveliness
    def __init__(self, uuid: _Optional[str] = ..., liveliness: _Optional[_Union[_machine_pb2.MachineLiveliness, str]] = ...) -> None: ...

class HousekeepingServiceMachineVPNConnectedRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class HousekeepingServiceMachineVPNConnectedResponse(_message.Message):
    __slots__ = ("modified_machines",)
    MODIFIED_MACHINES_FIELD_NUMBER: _ClassVar[int]
    modified_machines: _containers.RepeatedCompositeFieldContainer[_vpn_pb2.VPNNode]
    def __init__(self, modified_machines: _Optional[_Iterable[_Union[_vpn_pb2.VPNNode, _Mapping]]] = ...) -> None: ...

class HousekeepingServiceMachineResurrectRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class HousekeepingServiceMachineResurrectResponse(_message.Message):
    __slots__ = ("modified_machines",)
    MODIFIED_MACHINES_FIELD_NUMBER: _ClassVar[int]
    modified_machines: _containers.RepeatedCompositeFieldContainer[_machine_pb2.Machine]
    def __init__(self, modified_machines: _Optional[_Iterable[_Union[_machine_pb2.Machine, _Mapping]]] = ...) -> None: ...

class HousekeepingServiceNetworksConsistentRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class HousekeepingServiceNetworksConsistentResponse(_message.Message):
    __slots__ = ("networks_missing_in_ipam", "networks_missing_in_metal", "ips_missing_in_ipam", "ips_missing_in_metal")
    NETWORKS_MISSING_IN_IPAM_FIELD_NUMBER: _ClassVar[int]
    NETWORKS_MISSING_IN_METAL_FIELD_NUMBER: _ClassVar[int]
    IPS_MISSING_IN_IPAM_FIELD_NUMBER: _ClassVar[int]
    IPS_MISSING_IN_METAL_FIELD_NUMBER: _ClassVar[int]
    networks_missing_in_ipam: _containers.RepeatedCompositeFieldContainer[_network_pb2.Network]
    networks_missing_in_metal: _containers.RepeatedCompositeFieldContainer[_network_pb2.Network]
    ips_missing_in_ipam: _containers.RepeatedCompositeFieldContainer[_ip_pb2.IP]
    ips_missing_in_metal: _containers.RepeatedCompositeFieldContainer[_ip_pb2.IP]
    def __init__(self, networks_missing_in_ipam: _Optional[_Iterable[_Union[_network_pb2.Network, _Mapping]]] = ..., networks_missing_in_metal: _Optional[_Iterable[_Union[_network_pb2.Network, _Mapping]]] = ..., ips_missing_in_ipam: _Optional[_Iterable[_Union[_ip_pb2.IP, _Mapping]]] = ..., ips_missing_in_metal: _Optional[_Iterable[_Union[_ip_pb2.IP, _Mapping]]] = ...) -> None: ...
