from buf.validate import validate_pb2 as _validate_pb2
from metalstack.api.v2 import common_pb2 as _common_pb2
from metalstack.api.v2 import predefined_rules_pb2 as _predefined_rules_pb2
from metalstack.api.v2 import size_imageconstraint_pb2 as _size_imageconstraint_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SizeImageConstraintServiceCreateRequest(_message.Message):
    __slots__ = ("size", "image_constraints", "meta", "name", "description")
    SIZE_FIELD_NUMBER: _ClassVar[int]
    IMAGE_CONSTRAINTS_FIELD_NUMBER: _ClassVar[int]
    META_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    size: str
    image_constraints: _containers.RepeatedCompositeFieldContainer[_size_imageconstraint_pb2.ImageConstraint]
    meta: _common_pb2.Meta
    name: str
    description: str
    def __init__(self, size: _Optional[str] = ..., image_constraints: _Optional[_Iterable[_Union[_size_imageconstraint_pb2.ImageConstraint, _Mapping]]] = ..., meta: _Optional[_Union[_common_pb2.Meta, _Mapping]] = ..., name: _Optional[str] = ..., description: _Optional[str] = ...) -> None: ...

class SizeImageConstraintServiceCreateResponse(_message.Message):
    __slots__ = ("size_image_constraint",)
    SIZE_IMAGE_CONSTRAINT_FIELD_NUMBER: _ClassVar[int]
    size_image_constraint: _size_imageconstraint_pb2.SizeImageConstraint
    def __init__(self, size_image_constraint: _Optional[_Union[_size_imageconstraint_pb2.SizeImageConstraint, _Mapping]] = ...) -> None: ...

class SizeImageConstraintServiceUpdateRequest(_message.Message):
    __slots__ = ("size", "update_meta", "image_constraints", "name", "description")
    SIZE_FIELD_NUMBER: _ClassVar[int]
    UPDATE_META_FIELD_NUMBER: _ClassVar[int]
    IMAGE_CONSTRAINTS_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    size: str
    update_meta: _common_pb2.UpdateMeta
    image_constraints: _containers.RepeatedCompositeFieldContainer[_size_imageconstraint_pb2.ImageConstraint]
    name: str
    description: str
    def __init__(self, size: _Optional[str] = ..., update_meta: _Optional[_Union[_common_pb2.UpdateMeta, _Mapping]] = ..., image_constraints: _Optional[_Iterable[_Union[_size_imageconstraint_pb2.ImageConstraint, _Mapping]]] = ..., name: _Optional[str] = ..., description: _Optional[str] = ...) -> None: ...

class SizeImageConstraintServiceUpdateResponse(_message.Message):
    __slots__ = ("size_image_constraint",)
    SIZE_IMAGE_CONSTRAINT_FIELD_NUMBER: _ClassVar[int]
    size_image_constraint: _size_imageconstraint_pb2.SizeImageConstraint
    def __init__(self, size_image_constraint: _Optional[_Union[_size_imageconstraint_pb2.SizeImageConstraint, _Mapping]] = ...) -> None: ...

class SizeImageConstraintServiceDeleteRequest(_message.Message):
    __slots__ = ("size",)
    SIZE_FIELD_NUMBER: _ClassVar[int]
    size: str
    def __init__(self, size: _Optional[str] = ...) -> None: ...

class SizeImageConstraintServiceDeleteResponse(_message.Message):
    __slots__ = ("size_image_constraint",)
    SIZE_IMAGE_CONSTRAINT_FIELD_NUMBER: _ClassVar[int]
    size_image_constraint: _size_imageconstraint_pb2.SizeImageConstraint
    def __init__(self, size_image_constraint: _Optional[_Union[_size_imageconstraint_pb2.SizeImageConstraint, _Mapping]] = ...) -> None: ...

class SizeImageConstraintServiceGetRequest(_message.Message):
    __slots__ = ("size",)
    SIZE_FIELD_NUMBER: _ClassVar[int]
    size: str
    def __init__(self, size: _Optional[str] = ...) -> None: ...

class SizeImageConstraintServiceGetResponse(_message.Message):
    __slots__ = ("size_image_constraint",)
    SIZE_IMAGE_CONSTRAINT_FIELD_NUMBER: _ClassVar[int]
    size_image_constraint: _size_imageconstraint_pb2.SizeImageConstraint
    def __init__(self, size_image_constraint: _Optional[_Union[_size_imageconstraint_pb2.SizeImageConstraint, _Mapping]] = ...) -> None: ...

class SizeImageConstraintServiceListRequest(_message.Message):
    __slots__ = ("query",)
    QUERY_FIELD_NUMBER: _ClassVar[int]
    query: _size_imageconstraint_pb2.SizeImageConstraintQuery
    def __init__(self, query: _Optional[_Union[_size_imageconstraint_pb2.SizeImageConstraintQuery, _Mapping]] = ...) -> None: ...

class SizeImageConstraintServiceListResponse(_message.Message):
    __slots__ = ("size_image_constraints",)
    SIZE_IMAGE_CONSTRAINTS_FIELD_NUMBER: _ClassVar[int]
    size_image_constraints: _containers.RepeatedCompositeFieldContainer[_size_imageconstraint_pb2.SizeImageConstraint]
    def __init__(self, size_image_constraints: _Optional[_Iterable[_Union[_size_imageconstraint_pb2.SizeImageConstraint, _Mapping]]] = ...) -> None: ...
