import datetime

from buf.validate import validate_pb2 as _validate_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from metalstack.api.v2 import common_pb2 as _common_pb2
from metalstack.api.v2 import predefined_rules_pb2 as _predefined_rules_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class TenantMember(_message.Message):
    __slots__ = ("member", "role", "projects", "created_at", "meta", "tenant")
    MEMBER_FIELD_NUMBER: _ClassVar[int]
    ROLE_FIELD_NUMBER: _ClassVar[int]
    PROJECTS_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    META_FIELD_NUMBER: _ClassVar[int]
    TENANT_FIELD_NUMBER: _ClassVar[int]
    member: str
    role: _common_pb2.TenantRole
    projects: _containers.RepeatedScalarFieldContainer[str]
    created_at: _timestamp_pb2.Timestamp
    meta: _common_pb2.Meta
    tenant: str
    def __init__(self, member: _Optional[str] = ..., role: _Optional[_Union[_common_pb2.TenantRole, str]] = ..., projects: _Optional[_Iterable[str]] = ..., created_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ..., meta: _Optional[_Union[_common_pb2.Meta, _Mapping]] = ..., tenant: _Optional[str] = ...) -> None: ...

class TenantMemberServiceListRequest(_message.Message):
    __slots__ = ("login", "query")
    LOGIN_FIELD_NUMBER: _ClassVar[int]
    QUERY_FIELD_NUMBER: _ClassVar[int]
    login: str
    query: TenantMemberQuery
    def __init__(self, login: _Optional[str] = ..., query: _Optional[_Union[TenantMemberQuery, _Mapping]] = ...) -> None: ...

class TenantMemberQuery(_message.Message):
    __slots__ = ("member", "role", "projects")
    MEMBER_FIELD_NUMBER: _ClassVar[int]
    ROLE_FIELD_NUMBER: _ClassVar[int]
    PROJECTS_FIELD_NUMBER: _ClassVar[int]
    member: str
    role: _common_pb2.TenantRole
    projects: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, member: _Optional[str] = ..., role: _Optional[_Union[_common_pb2.TenantRole, str]] = ..., projects: _Optional[_Iterable[str]] = ...) -> None: ...

class TenantMemberServiceGetRequest(_message.Message):
    __slots__ = ("login", "member")
    LOGIN_FIELD_NUMBER: _ClassVar[int]
    MEMBER_FIELD_NUMBER: _ClassVar[int]
    login: str
    member: str
    def __init__(self, login: _Optional[str] = ..., member: _Optional[str] = ...) -> None: ...

class TenantMemberServiceDeleteRequest(_message.Message):
    __slots__ = ("login", "member")
    LOGIN_FIELD_NUMBER: _ClassVar[int]
    MEMBER_FIELD_NUMBER: _ClassVar[int]
    login: str
    member: str
    def __init__(self, login: _Optional[str] = ..., member: _Optional[str] = ...) -> None: ...

class TenantMemberServiceCreateRequest(_message.Message):
    __slots__ = ("login", "member", "role")
    LOGIN_FIELD_NUMBER: _ClassVar[int]
    MEMBER_FIELD_NUMBER: _ClassVar[int]
    ROLE_FIELD_NUMBER: _ClassVar[int]
    login: str
    member: str
    role: _common_pb2.TenantRole
    def __init__(self, login: _Optional[str] = ..., member: _Optional[str] = ..., role: _Optional[_Union[_common_pb2.TenantRole, str]] = ...) -> None: ...

class TenantMemberServiceUpdateRequest(_message.Message):
    __slots__ = ("login", "member", "role", "update_meta")
    LOGIN_FIELD_NUMBER: _ClassVar[int]
    MEMBER_FIELD_NUMBER: _ClassVar[int]
    ROLE_FIELD_NUMBER: _ClassVar[int]
    UPDATE_META_FIELD_NUMBER: _ClassVar[int]
    login: str
    member: str
    role: _common_pb2.TenantRole
    update_meta: _common_pb2.UpdateMeta
    def __init__(self, login: _Optional[str] = ..., member: _Optional[str] = ..., role: _Optional[_Union[_common_pb2.TenantRole, str]] = ..., update_meta: _Optional[_Union[_common_pb2.UpdateMeta, _Mapping]] = ...) -> None: ...

class TenantMemberServiceGetResponse(_message.Message):
    __slots__ = ("member",)
    MEMBER_FIELD_NUMBER: _ClassVar[int]
    member: TenantMember
    def __init__(self, member: _Optional[_Union[TenantMember, _Mapping]] = ...) -> None: ...

class TenantMemberServiceListResponse(_message.Message):
    __slots__ = ("members",)
    MEMBERS_FIELD_NUMBER: _ClassVar[int]
    members: _containers.RepeatedCompositeFieldContainer[TenantMember]
    def __init__(self, members: _Optional[_Iterable[_Union[TenantMember, _Mapping]]] = ...) -> None: ...

class TenantMemberServiceCreateResponse(_message.Message):
    __slots__ = ("member",)
    MEMBER_FIELD_NUMBER: _ClassVar[int]
    member: TenantMember
    def __init__(self, member: _Optional[_Union[TenantMember, _Mapping]] = ...) -> None: ...

class TenantMemberServiceUpdateResponse(_message.Message):
    __slots__ = ("member",)
    MEMBER_FIELD_NUMBER: _ClassVar[int]
    member: TenantMember
    def __init__(self, member: _Optional[_Union[TenantMember, _Mapping]] = ...) -> None: ...

class TenantMemberServiceDeleteResponse(_message.Message):
    __slots__ = ("member",)
    MEMBER_FIELD_NUMBER: _ClassVar[int]
    member: TenantMember
    def __init__(self, member: _Optional[_Union[TenantMember, _Mapping]] = ...) -> None: ...

class TenantMemberServiceLeaveRequest(_message.Message):
    __slots__ = ("login",)
    LOGIN_FIELD_NUMBER: _ClassVar[int]
    login: str
    def __init__(self, login: _Optional[str] = ...) -> None: ...

class TenantMemberServiceLeaveResponse(_message.Message):
    __slots__ = ("member",)
    MEMBER_FIELD_NUMBER: _ClassVar[int]
    member: TenantMember
    def __init__(self, member: _Optional[_Union[TenantMember, _Mapping]] = ...) -> None: ...
