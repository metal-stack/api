from buf.validate import validate_pb2 as _validate_pb2
from metalstack.api.v2 import predefined_rules_pb2 as _predefined_rules_pb2
from metalstack.api.v2 import tenant_pb2 as _tenant_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class TenantMember(_message.Message):
    __slots__ = ("tenant", "tenant_member")
    TENANT_FIELD_NUMBER: _ClassVar[int]
    TENANT_MEMBER_FIELD_NUMBER: _ClassVar[int]
    tenant: str
    tenant_member: _tenant_pb2.TenantMember
    def __init__(self, tenant: _Optional[str] = ..., tenant_member: _Optional[_Union[_tenant_pb2.TenantMember, _Mapping]] = ...) -> None: ...
