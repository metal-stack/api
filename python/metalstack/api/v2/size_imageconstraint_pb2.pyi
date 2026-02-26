from buf.validate import validate_pb2 as _validate_pb2
from metalstack.api.v2 import common_pb2 as _common_pb2
from metalstack.api.v2 import predefined_rules_pb2 as _predefined_rules_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SizeImageConstraintServiceTryRequest(_message.Message):
    __slots__ = ("size", "image")
    SIZE_FIELD_NUMBER: _ClassVar[int]
    IMAGE_FIELD_NUMBER: _ClassVar[int]
    size: str
    image: str
    def __init__(self, size: _Optional[str] = ..., image: _Optional[str] = ...) -> None: ...

class SizeImageConstraintServiceTryResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class SizeImageConstraint(_message.Message):
    __slots__ = ("size", "image_constraints", "meta", "name", "description")
    SIZE_FIELD_NUMBER: _ClassVar[int]
    IMAGE_CONSTRAINTS_FIELD_NUMBER: _ClassVar[int]
    META_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    size: str
    image_constraints: _containers.RepeatedCompositeFieldContainer[ImageConstraint]
    meta: _common_pb2.Meta
    name: str
    description: str
    def __init__(self, size: _Optional[str] = ..., image_constraints: _Optional[_Iterable[_Union[ImageConstraint, _Mapping]]] = ..., meta: _Optional[_Union[_common_pb2.Meta, _Mapping]] = ..., name: _Optional[str] = ..., description: _Optional[str] = ...) -> None: ...

class ImageConstraint(_message.Message):
    __slots__ = ("image", "semver_match")
    IMAGE_FIELD_NUMBER: _ClassVar[int]
    SEMVER_MATCH_FIELD_NUMBER: _ClassVar[int]
    image: str
    semver_match: str
    def __init__(self, image: _Optional[str] = ..., semver_match: _Optional[str] = ...) -> None: ...

class SizeImageConstraintQuery(_message.Message):
    __slots__ = ("size", "name", "description")
    SIZE_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    size: str
    name: str
    description: str
    def __init__(self, size: _Optional[str] = ..., name: _Optional[str] = ..., description: _Optional[str] = ...) -> None: ...
