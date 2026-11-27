import datetime

from buf.validate import validate_pb2 as _validate_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from metalstack.api.v2 import common_pb2 as _common_pb2
from metalstack.api.v2 import predefined_rules_pb2 as _predefined_rules_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class QuotaResource(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    QUOTA_RESOURCE_UNSPECIFIED: _ClassVar[QuotaResource]
    QUOTA_RESOURCE_MACHINES: _ClassVar[QuotaResource]
    QUOTA_RESOURCE_NETWORKS: _ClassVar[QuotaResource]
    QUOTA_RESOURCE_PROJECTS: _ClassVar[QuotaResource]
    QUOTA_RESOURCE_IPS: _ClassVar[QuotaResource]

class QuotaSource(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    QUOTA_SOURCE_UNSPECIFIED: _ClassVar[QuotaSource]
    QUOTA_SOURCE_DEFAULT: _ClassVar[QuotaSource]
    QUOTA_SOURCE_TENANT: _ClassVar[QuotaSource]

class QuotaRequestState(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    QUOTA_REQUEST_STATE_UNSPECIFIED: _ClassVar[QuotaRequestState]
    QUOTA_REQUEST_STATE_PENDING: _ClassVar[QuotaRequestState]
    QUOTA_REQUEST_STATE_ADMITTED: _ClassVar[QuotaRequestState]
    QUOTA_REQUEST_STATE_REJECTED: _ClassVar[QuotaRequestState]
QUOTA_RESOURCE_UNSPECIFIED: QuotaResource
QUOTA_RESOURCE_MACHINES: QuotaResource
QUOTA_RESOURCE_NETWORKS: QuotaResource
QUOTA_RESOURCE_PROJECTS: QuotaResource
QUOTA_RESOURCE_IPS: QuotaResource
QUOTA_SOURCE_UNSPECIFIED: QuotaSource
QUOTA_SOURCE_DEFAULT: QuotaSource
QUOTA_SOURCE_TENANT: QuotaSource
QUOTA_REQUEST_STATE_UNSPECIFIED: QuotaRequestState
QUOTA_REQUEST_STATE_PENDING: QuotaRequestState
QUOTA_REQUEST_STATE_ADMITTED: QuotaRequestState
QUOTA_REQUEST_STATE_REJECTED: QuotaRequestState

class TenantQuota(_message.Message):
    __slots__ = ("login", "resource", "max", "used", "source", "pending_request")
    LOGIN_FIELD_NUMBER: _ClassVar[int]
    RESOURCE_FIELD_NUMBER: _ClassVar[int]
    MAX_FIELD_NUMBER: _ClassVar[int]
    USED_FIELD_NUMBER: _ClassVar[int]
    SOURCE_FIELD_NUMBER: _ClassVar[int]
    PENDING_REQUEST_FIELD_NUMBER: _ClassVar[int]
    login: str
    resource: QuotaResource
    max: int
    used: int
    source: QuotaSource
    pending_request: QuotaRequest
    def __init__(self, login: _Optional[str] = ..., resource: _Optional[_Union[QuotaResource, str]] = ..., max: _Optional[int] = ..., used: _Optional[int] = ..., source: _Optional[_Union[QuotaSource, str]] = ..., pending_request: _Optional[_Union[QuotaRequest, _Mapping]] = ...) -> None: ...

class QuotaRequest(_message.Message):
    __slots__ = ("id", "meta", "login", "resource", "max", "reason", "state", "requested_by", "decided_at", "comment")
    ID_FIELD_NUMBER: _ClassVar[int]
    META_FIELD_NUMBER: _ClassVar[int]
    LOGIN_FIELD_NUMBER: _ClassVar[int]
    RESOURCE_FIELD_NUMBER: _ClassVar[int]
    MAX_FIELD_NUMBER: _ClassVar[int]
    REASON_FIELD_NUMBER: _ClassVar[int]
    STATE_FIELD_NUMBER: _ClassVar[int]
    REQUESTED_BY_FIELD_NUMBER: _ClassVar[int]
    DECIDED_AT_FIELD_NUMBER: _ClassVar[int]
    COMMENT_FIELD_NUMBER: _ClassVar[int]
    id: str
    meta: _common_pb2.Meta
    login: str
    resource: QuotaResource
    max: int
    reason: str
    state: QuotaRequestState
    requested_by: str
    decided_at: _timestamp_pb2.Timestamp
    comment: str
    def __init__(self, id: _Optional[str] = ..., meta: _Optional[_Union[_common_pb2.Meta, _Mapping]] = ..., login: _Optional[str] = ..., resource: _Optional[_Union[QuotaResource, str]] = ..., max: _Optional[int] = ..., reason: _Optional[str] = ..., state: _Optional[_Union[QuotaRequestState, str]] = ..., requested_by: _Optional[str] = ..., decided_at: _Optional[_Union[datetime.datetime, _timestamp_pb2.Timestamp, _Mapping]] = ..., comment: _Optional[str] = ...) -> None: ...

class QuotaServiceGetRequest(_message.Message):
    __slots__ = ("login", "resource")
    LOGIN_FIELD_NUMBER: _ClassVar[int]
    RESOURCE_FIELD_NUMBER: _ClassVar[int]
    login: str
    resource: QuotaResource
    def __init__(self, login: _Optional[str] = ..., resource: _Optional[_Union[QuotaResource, str]] = ...) -> None: ...

class QuotaServiceGetResponse(_message.Message):
    __slots__ = ("quota",)
    QUOTA_FIELD_NUMBER: _ClassVar[int]
    quota: TenantQuota
    def __init__(self, quota: _Optional[_Union[TenantQuota, _Mapping]] = ...) -> None: ...

class QuotaServiceListRequest(_message.Message):
    __slots__ = ("login",)
    LOGIN_FIELD_NUMBER: _ClassVar[int]
    login: str
    def __init__(self, login: _Optional[str] = ...) -> None: ...

class QuotaServiceListResponse(_message.Message):
    __slots__ = ("quotas",)
    QUOTAS_FIELD_NUMBER: _ClassVar[int]
    quotas: _containers.RepeatedCompositeFieldContainer[TenantQuota]
    def __init__(self, quotas: _Optional[_Iterable[_Union[TenantQuota, _Mapping]]] = ...) -> None: ...

class QuotaServiceChangeRequestRequest(_message.Message):
    __slots__ = ("login", "resource", "max", "reason")
    LOGIN_FIELD_NUMBER: _ClassVar[int]
    RESOURCE_FIELD_NUMBER: _ClassVar[int]
    MAX_FIELD_NUMBER: _ClassVar[int]
    REASON_FIELD_NUMBER: _ClassVar[int]
    login: str
    resource: QuotaResource
    max: int
    reason: str
    def __init__(self, login: _Optional[str] = ..., resource: _Optional[_Union[QuotaResource, str]] = ..., max: _Optional[int] = ..., reason: _Optional[str] = ...) -> None: ...

class QuotaServiceChangeRequestResponse(_message.Message):
    __slots__ = ("request",)
    REQUEST_FIELD_NUMBER: _ClassVar[int]
    request: QuotaRequest
    def __init__(self, request: _Optional[_Union[QuotaRequest, _Mapping]] = ...) -> None: ...

class QuotaServiceDeleteRequestRequest(_message.Message):
    __slots__ = ("login", "id")
    LOGIN_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    login: str
    id: str
    def __init__(self, login: _Optional[str] = ..., id: _Optional[str] = ...) -> None: ...

class QuotaServiceDeleteRequestResponse(_message.Message):
    __slots__ = ("request",)
    REQUEST_FIELD_NUMBER: _ClassVar[int]
    request: QuotaRequest
    def __init__(self, request: _Optional[_Union[QuotaRequest, _Mapping]] = ...) -> None: ...
