from buf.validate import validate_pb2 as _validate_pb2
from metalstack.api.v2 import common_pb2 as _common_pb2
from metalstack.api.v2 import predefined_rules_pb2 as _predefined_rules_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SizeReservationServiceGetRequest(_message.Message):
    __slots__ = ("id", "project")
    ID_FIELD_NUMBER: _ClassVar[int]
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    id: str
    project: str
    def __init__(self, id: _Optional[str] = ..., project: _Optional[str] = ...) -> None: ...

class SizeReservationServiceListRequest(_message.Message):
    __slots__ = ("project", "query")
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    QUERY_FIELD_NUMBER: _ClassVar[int]
    project: str
    query: SizeReservationQuery
    def __init__(self, project: _Optional[str] = ..., query: _Optional[_Union[SizeReservationQuery, _Mapping]] = ...) -> None: ...

class SizeReservationServiceGetResponse(_message.Message):
    __slots__ = ("size_reservation",)
    SIZE_RESERVATION_FIELD_NUMBER: _ClassVar[int]
    size_reservation: SizeReservation
    def __init__(self, size_reservation: _Optional[_Union[SizeReservation, _Mapping]] = ...) -> None: ...

class SizeReservationServiceListResponse(_message.Message):
    __slots__ = ("size_reservations",)
    SIZE_RESERVATIONS_FIELD_NUMBER: _ClassVar[int]
    size_reservations: _containers.RepeatedCompositeFieldContainer[SizeReservation]
    def __init__(self, size_reservations: _Optional[_Iterable[_Union[SizeReservation, _Mapping]]] = ...) -> None: ...

class SizeReservation(_message.Message):
    __slots__ = ("id", "meta", "name", "description", "project", "size", "partitions", "amount")
    ID_FIELD_NUMBER: _ClassVar[int]
    META_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    SIZE_FIELD_NUMBER: _ClassVar[int]
    PARTITIONS_FIELD_NUMBER: _ClassVar[int]
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    id: str
    meta: _common_pb2.Meta
    name: str
    description: str
    project: str
    size: str
    partitions: _containers.RepeatedScalarFieldContainer[str]
    amount: int
    def __init__(self, id: _Optional[str] = ..., meta: _Optional[_Union[_common_pb2.Meta, _Mapping]] = ..., name: _Optional[str] = ..., description: _Optional[str] = ..., project: _Optional[str] = ..., size: _Optional[str] = ..., partitions: _Optional[_Iterable[str]] = ..., amount: _Optional[int] = ...) -> None: ...

class SizeReservationQuery(_message.Message):
    __slots__ = ("id", "name", "description", "size", "project", "partition", "labels")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    SIZE_FIELD_NUMBER: _ClassVar[int]
    PROJECT_FIELD_NUMBER: _ClassVar[int]
    PARTITION_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    description: str
    size: str
    project: str
    partition: str
    labels: _common_pb2.Labels
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., description: _Optional[str] = ..., size: _Optional[str] = ..., project: _Optional[str] = ..., partition: _Optional[str] = ..., labels: _Optional[_Union[_common_pb2.Labels, _Mapping]] = ...) -> None: ...
