from buf.validate import validate_pb2 as _validate_pb2
from metalstack.api.v2 import common_pb2 as _common_pb2
from metalstack.api.v2 import predefined_rules_pb2 as _predefined_rules_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ProjectMember(_message.Message):
    __slots__ = ("member", "role", "inherited_membership", "meta", "project")
    MEMBER_FIELD_NUMBER: _ClassVar[int]
    ROLE_FIELD_NUMBER: _ClassVar[int]
    INHERITED_MEMBERSHIP_FIELD_NUMBER: _ClassVar[int]
    META_FIELD_NUMBER: _ClassVar[int]
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    member: str
    role: _common_pb2.ProjectRole
    inherited_membership: bool
    meta: _common_pb2.Meta
    project: str
    def __init__(self, member: _Optional[str] = ..., role: _Optional[_Union[_common_pb2.ProjectRole, str]] = ..., inherited_membership: _Optional[bool] = ..., meta: _Optional[_Union[_common_pb2.Meta, _Mapping]] = ..., project: _Optional[str] = ...) -> None: ...

class ProjectMemberQuery(_message.Message):
    __slots__ = ("member", "role", "labels")
    MEMBER_FIELD_NUMBER: _ClassVar[int]
    ROLE_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    member: str
    role: _common_pb2.ProjectRole
    labels: _common_pb2.Labels
    def __init__(self, member: _Optional[str] = ..., role: _Optional[_Union[_common_pb2.ProjectRole, str]] = ..., labels: _Optional[_Union[_common_pb2.Labels, _Mapping]] = ...) -> None: ...

class ProjectMemberServiceListRequest(_message.Message):
    __slots__ = ("project", "query")
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    QUERY_FIELD_NUMBER: _ClassVar[int]
    project: str
    query: ProjectMemberQuery
    def __init__(self, project: _Optional[str] = ..., query: _Optional[_Union[ProjectMemberQuery, _Mapping]] = ...) -> None: ...

class ProjectMemberServiceGetRequest(_message.Message):
    __slots__ = ("project", "member")
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    MEMBER_FIELD_NUMBER: _ClassVar[int]
    project: str
    member: str
    def __init__(self, project: _Optional[str] = ..., member: _Optional[str] = ...) -> None: ...

class ProjectMemberServiceCreateRequest(_message.Message):
    __slots__ = ("project", "member", "role", "labels")
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    MEMBER_FIELD_NUMBER: _ClassVar[int]
    ROLE_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    project: str
    member: str
    role: _common_pb2.ProjectRole
    labels: _common_pb2.Labels
    def __init__(self, project: _Optional[str] = ..., member: _Optional[str] = ..., role: _Optional[_Union[_common_pb2.ProjectRole, str]] = ..., labels: _Optional[_Union[_common_pb2.Labels, _Mapping]] = ...) -> None: ...

class ProjectMemberServiceUpdateRequest(_message.Message):
    __slots__ = ("project", "member", "role", "update_meta", "labels")
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    MEMBER_FIELD_NUMBER: _ClassVar[int]
    ROLE_FIELD_NUMBER: _ClassVar[int]
    UPDATE_META_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    project: str
    member: str
    role: _common_pb2.ProjectRole
    update_meta: _common_pb2.UpdateMeta
    labels: _common_pb2.UpdateLabels
    def __init__(self, project: _Optional[str] = ..., member: _Optional[str] = ..., role: _Optional[_Union[_common_pb2.ProjectRole, str]] = ..., update_meta: _Optional[_Union[_common_pb2.UpdateMeta, _Mapping]] = ..., labels: _Optional[_Union[_common_pb2.UpdateLabels, _Mapping]] = ...) -> None: ...

class ProjectMemberServiceDeleteRequest(_message.Message):
    __slots__ = ("project", "member")
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    MEMBER_FIELD_NUMBER: _ClassVar[int]
    project: str
    member: str
    def __init__(self, project: _Optional[str] = ..., member: _Optional[str] = ...) -> None: ...

class ProjectMemberServiceListResponse(_message.Message):
    __slots__ = ("members",)
    MEMBERS_FIELD_NUMBER: _ClassVar[int]
    members: _containers.RepeatedCompositeFieldContainer[ProjectMember]
    def __init__(self, members: _Optional[_Iterable[_Union[ProjectMember, _Mapping]]] = ...) -> None: ...

class ProjectMemberServiceGetResponse(_message.Message):
    __slots__ = ("member",)
    MEMBER_FIELD_NUMBER: _ClassVar[int]
    member: ProjectMember
    def __init__(self, member: _Optional[_Union[ProjectMember, _Mapping]] = ...) -> None: ...

class ProjectMemberServiceCreateResponse(_message.Message):
    __slots__ = ("member",)
    MEMBER_FIELD_NUMBER: _ClassVar[int]
    member: ProjectMember
    def __init__(self, member: _Optional[_Union[ProjectMember, _Mapping]] = ...) -> None: ...

class ProjectMemberServiceDeleteResponse(_message.Message):
    __slots__ = ("member",)
    MEMBER_FIELD_NUMBER: _ClassVar[int]
    member: ProjectMember
    def __init__(self, member: _Optional[_Union[ProjectMember, _Mapping]] = ...) -> None: ...

class ProjectMemberServiceUpdateResponse(_message.Message):
    __slots__ = ("member",)
    MEMBER_FIELD_NUMBER: _ClassVar[int]
    member: ProjectMember
    def __init__(self, member: _Optional[_Union[ProjectMember, _Mapping]] = ...) -> None: ...

class ProjectMemberServiceLeaveRequest(_message.Message):
    __slots__ = ("project",)
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    project: str
    def __init__(self, project: _Optional[str] = ...) -> None: ...

class ProjectMemberServiceLeaveResponse(_message.Message):
    __slots__ = ("member",)
    MEMBER_FIELD_NUMBER: _ClassVar[int]
    member: ProjectMember
    def __init__(self, member: _Optional[_Union[ProjectMember, _Mapping]] = ...) -> None: ...
