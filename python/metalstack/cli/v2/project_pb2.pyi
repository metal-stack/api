from buf.validate import validate_pb2 as _validate_pb2
from metalstack.api.v2 import project_pb2 as _project_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ProjectMember(_message.Message):
    __slots__ = ("project", "project_member")
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    PROJECT_MEMBER_FIELD_NUMBER: _ClassVar[int]
    project: str
    project_member: _project_pb2.ProjectMember
    def __init__(self, project: _Optional[str] = ..., project_member: _Optional[_Union[_project_pb2.ProjectMember, _Mapping]] = ...) -> None: ...
