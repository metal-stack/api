from buf.validate import validate_pb2 as _validate_pb2
from metalstack.api.v2 import common_pb2 as _common_pb2
from metalstack.api.v2 import predefined_rules_pb2 as _predefined_rules_pb2
from metalstack.api.v2 import project_pb2 as _project_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ProjectServiceListRequest(_message.Message):
    __slots__ = ("query",)
    QUERY_FIELD_NUMBER: _ClassVar[int]
    query: _project_pb2.ProjectQuery
    def __init__(self, query: _Optional[_Union[_project_pb2.ProjectQuery, _Mapping]] = ...) -> None: ...

class ProjectServiceListResponse(_message.Message):
    __slots__ = ("projects",)
    PROJECTS_FIELD_NUMBER: _ClassVar[int]
    projects: _containers.RepeatedCompositeFieldContainer[_project_pb2.Project]
    def __init__(self, projects: _Optional[_Iterable[_Union[_project_pb2.Project, _Mapping]]] = ...) -> None: ...

class ProjectServiceCreateRequest(_message.Message):
    __slots__ = ("login", "name", "description", "avatar_url", "labels", "project")
    LOGIN_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    AVATAR_URL_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    login: str
    name: str
    description: str
    avatar_url: str
    labels: _common_pb2.Labels
    project: str
    def __init__(self, login: _Optional[str] = ..., name: _Optional[str] = ..., description: _Optional[str] = ..., avatar_url: _Optional[str] = ..., labels: _Optional[_Union[_common_pb2.Labels, _Mapping]] = ..., project: _Optional[str] = ...) -> None: ...

class ProjectServiceCreateResponse(_message.Message):
    __slots__ = ("project",)
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    project: _project_pb2.Project
    def __init__(self, project: _Optional[_Union[_project_pb2.Project, _Mapping]] = ...) -> None: ...
