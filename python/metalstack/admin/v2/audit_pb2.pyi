from buf.validate import validate_pb2 as _validate_pb2
from metalstack.api.v2 import audit_pb2 as _audit_pb2
from metalstack.api.v2 import common_pb2 as _common_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AuditServiceListRequest(_message.Message):
    __slots__ = ("query",)
    QUERY_FIELD_NUMBER: _ClassVar[int]
    query: _audit_pb2.AuditQuery
    def __init__(self, query: _Optional[_Union[_audit_pb2.AuditQuery, _Mapping]] = ...) -> None: ...

class AuditServiceListResponse(_message.Message):
    __slots__ = ("traces",)
    TRACES_FIELD_NUMBER: _ClassVar[int]
    traces: _containers.RepeatedCompositeFieldContainer[_audit_pb2.AuditTrace]
    def __init__(self, traces: _Optional[_Iterable[_Union[_audit_pb2.AuditTrace, _Mapping]]] = ...) -> None: ...

class AuditServiceGetRequest(_message.Message):
    __slots__ = ("uuid", "phase")
    UUID_FIELD_NUMBER: _ClassVar[int]
    PHASE_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    phase: _audit_pb2.AuditPhase
    def __init__(self, uuid: _Optional[str] = ..., phase: _Optional[_Union[_audit_pb2.AuditPhase, str]] = ...) -> None: ...

class AuditServiceGetResponse(_message.Message):
    __slots__ = ("trace",)
    TRACE_FIELD_NUMBER: _ClassVar[int]
    trace: _audit_pb2.AuditTrace
    def __init__(self, trace: _Optional[_Union[_audit_pb2.AuditTrace, _Mapping]] = ...) -> None: ...
