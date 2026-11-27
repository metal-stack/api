from buf.validate import validate_pb2 as _validate_pb2
from metalstack.api.v2 import common_pb2 as _common_pb2
from metalstack.api.v2 import predefined_rules_pb2 as _predefined_rules_pb2
from metalstack.api.v2 import quota_pb2 as _quota_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class QuotaServiceListRequest(_message.Message):
    __slots__ = ("login", "resource")
    LOGIN_FIELD_NUMBER: _ClassVar[int]
    RESOURCE_FIELD_NUMBER: _ClassVar[int]
    login: str
    resource: _quota_pb2.QuotaResource
    def __init__(self, login: _Optional[str] = ..., resource: _Optional[_Union[_quota_pb2.QuotaResource, str]] = ...) -> None: ...

class QuotaServiceListResponse(_message.Message):
    __slots__ = ("quotas",)
    QUOTAS_FIELD_NUMBER: _ClassVar[int]
    quotas: _containers.RepeatedCompositeFieldContainer[_quota_pb2.TenantQuota]
    def __init__(self, quotas: _Optional[_Iterable[_Union[_quota_pb2.TenantQuota, _Mapping]]] = ...) -> None: ...

class QuotaServiceListRequestsRequest(_message.Message):
    __slots__ = ("login", "resource", "state")
    LOGIN_FIELD_NUMBER: _ClassVar[int]
    RESOURCE_FIELD_NUMBER: _ClassVar[int]
    STATE_FIELD_NUMBER: _ClassVar[int]
    login: str
    resource: _quota_pb2.QuotaResource
    state: _quota_pb2.QuotaRequestState
    def __init__(self, login: _Optional[str] = ..., resource: _Optional[_Union[_quota_pb2.QuotaResource, str]] = ..., state: _Optional[_Union[_quota_pb2.QuotaRequestState, str]] = ...) -> None: ...

class QuotaServiceListRequestsResponse(_message.Message):
    __slots__ = ("requests",)
    REQUESTS_FIELD_NUMBER: _ClassVar[int]
    requests: _containers.RepeatedCompositeFieldContainer[_quota_pb2.QuotaRequest]
    def __init__(self, requests: _Optional[_Iterable[_Union[_quota_pb2.QuotaRequest, _Mapping]]] = ...) -> None: ...

class QuotaServiceAdmitRequest(_message.Message):
    __slots__ = ("id", "max", "comment")
    ID_FIELD_NUMBER: _ClassVar[int]
    MAX_FIELD_NUMBER: _ClassVar[int]
    COMMENT_FIELD_NUMBER: _ClassVar[int]
    id: str
    max: int
    comment: str
    def __init__(self, id: _Optional[str] = ..., max: _Optional[int] = ..., comment: _Optional[str] = ...) -> None: ...

class QuotaServiceAdmitResponse(_message.Message):
    __slots__ = ("request", "quota")
    REQUEST_FIELD_NUMBER: _ClassVar[int]
    QUOTA_FIELD_NUMBER: _ClassVar[int]
    request: _quota_pb2.QuotaRequest
    quota: _quota_pb2.TenantQuota
    def __init__(self, request: _Optional[_Union[_quota_pb2.QuotaRequest, _Mapping]] = ..., quota: _Optional[_Union[_quota_pb2.TenantQuota, _Mapping]] = ...) -> None: ...

class QuotaServiceRejectRequest(_message.Message):
    __slots__ = ("id", "comment")
    ID_FIELD_NUMBER: _ClassVar[int]
    COMMENT_FIELD_NUMBER: _ClassVar[int]
    id: str
    comment: str
    def __init__(self, id: _Optional[str] = ..., comment: _Optional[str] = ...) -> None: ...

class QuotaServiceRejectResponse(_message.Message):
    __slots__ = ("request",)
    REQUEST_FIELD_NUMBER: _ClassVar[int]
    request: _quota_pb2.QuotaRequest
    def __init__(self, request: _Optional[_Union[_quota_pb2.QuotaRequest, _Mapping]] = ...) -> None: ...
