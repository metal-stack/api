import datetime

from buf.validate import validate_pb2 as _validate_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class VPNFlavor(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    VPN_FLAVOR_UNSPECIFIED: _ClassVar[VPNFlavor]
    VPN_FLAVOR_TAILSCALE: _ClassVar[VPNFlavor]
VPN_FLAVOR_UNSPECIFIED: VPNFlavor
VPN_FLAVOR_TAILSCALE: VPNFlavor

class VPNNode(_message.Message):
    __slots__ = ("id", "name", "project", "ip_addresses", "last_seen", "online", "flavor")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    IP_ADDRESSES_FIELD_NUMBER: _ClassVar[int]
    LAST_SEEN_FIELD_NUMBER: _ClassVar[int]
    ONLINE_FIELD_NUMBER: _ClassVar[int]
    FLAVOR_FIELD_NUMBER: _ClassVar[int]
    id: int
    name: str
    project: str
    ip_addresses: _containers.RepeatedScalarFieldContainer[str]
    last_seen: _timestamp_pb2.Timestamp
    online: bool
    flavor: VPNFlavor
    def __init__(self, id: _Optional[int] = ..., name: _Optional[str] = ..., project: _Optional[str] = ..., ip_addresses: _Optional[_Iterable[str]] = ..., last_seen: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ..., online: _Optional[bool] = ..., flavor: _Optional[_Union[VPNFlavor, str]] = ...) -> None: ...
