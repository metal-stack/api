from buf.validate import validate_pb2 as _validate_pb2
from metalstack.api.v2 import common_pb2 as _common_pb2
from metalstack.api.v2 import component_pb2 as _component_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ComponentServiceListRequest(_message.Message):
    __slots__ = ("query",)
    QUERY_FIELD_NUMBER: _ClassVar[int]
    query: _component_pb2.ComponentQuery
    def __init__(self, query: _Optional[_Union[_component_pb2.ComponentQuery, _Mapping]] = ...) -> None: ...

class ComponentServiceListResponse(_message.Message):
    __slots__ = ("components",)
    COMPONENTS_FIELD_NUMBER: _ClassVar[int]
    components: _containers.RepeatedCompositeFieldContainer[_component_pb2.Component]
    def __init__(self, components: _Optional[_Iterable[_Union[_component_pb2.Component, _Mapping]]] = ...) -> None: ...

class ComponentServiceGetRequest(_message.Message):
    __slots__ = ("uuid",)
    UUID_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    def __init__(self, uuid: _Optional[str] = ...) -> None: ...

class ComponentServiceGetResponse(_message.Message):
    __slots__ = ("component",)
    COMPONENT_FIELD_NUMBER: _ClassVar[int]
    component: _component_pb2.Component
    def __init__(self, component: _Optional[_Union[_component_pb2.Component, _Mapping]] = ...) -> None: ...

class ComponentServiceDeleteRequest(_message.Message):
    __slots__ = ("uuid",)
    UUID_FIELD_NUMBER: _ClassVar[int]
    uuid: str
    def __init__(self, uuid: _Optional[str] = ...) -> None: ...

class ComponentServiceDeleteResponse(_message.Message):
    __slots__ = ("component",)
    COMPONENT_FIELD_NUMBER: _ClassVar[int]
    component: _component_pb2.Component
    def __init__(self, component: _Optional[_Union[_component_pb2.Component, _Mapping]] = ...) -> None: ...
