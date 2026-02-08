from buf.validate import validate_pb2 as _validate_pb2
from metalstack.api.v2 import common_pb2 as _common_pb2
from metalstack.api.v2 import predefined_rules_pb2 as _predefined_rules_pb2
from metalstack.api.v2 import size_reservation_pb2 as _size_reservation_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Iterable as _Iterable, Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SizeReservationServiceCreateRequest(_message.Message):
    __slots__ = ("size_reservation",)
    SIZE_RESERVATION_FIELD_NUMBER: _ClassVar[int]
    size_reservation: _size_reservation_pb2.SizeReservation
    def __init__(self, size_reservation: _Optional[_Union[_size_reservation_pb2.SizeReservation, _Mapping]] = ...) -> None: ...

class SizeReservationServiceCreateResponse(_message.Message):
    __slots__ = ("size_reservation",)
    SIZE_RESERVATION_FIELD_NUMBER: _ClassVar[int]
    size_reservation: _size_reservation_pb2.SizeReservation
    def __init__(self, size_reservation: _Optional[_Union[_size_reservation_pb2.SizeReservation, _Mapping]] = ...) -> None: ...

class SizeReservationServiceUpdateRequest(_message.Message):
    __slots__ = ("id", "name", "description", "update_meta", "partitions", "amount", "labels")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    UPDATE_META_FIELD_NUMBER: _ClassVar[int]
    PARTITIONS_FIELD_NUMBER: _ClassVar[int]
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    LABELS_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    description: str
    update_meta: _common_pb2.UpdateMeta
    partitions: _containers.RepeatedScalarFieldContainer[str]
    amount: int
    labels: _common_pb2.UpdateLabels
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., description: _Optional[str] = ..., update_meta: _Optional[_Union[_common_pb2.UpdateMeta, _Mapping]] = ..., partitions: _Optional[_Iterable[str]] = ..., amount: _Optional[int] = ..., labels: _Optional[_Union[_common_pb2.UpdateLabels, _Mapping]] = ...) -> None: ...

class SizeReservationServiceUpdateResponse(_message.Message):
    __slots__ = ("size_reservation",)
    SIZE_RESERVATION_FIELD_NUMBER: _ClassVar[int]
    size_reservation: _size_reservation_pb2.SizeReservation
    def __init__(self, size_reservation: _Optional[_Union[_size_reservation_pb2.SizeReservation, _Mapping]] = ...) -> None: ...

class SizeReservationServiceDeleteRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class SizeReservationServiceDeleteResponse(_message.Message):
    __slots__ = ("size_reservation",)
    SIZE_RESERVATION_FIELD_NUMBER: _ClassVar[int]
    size_reservation: _size_reservation_pb2.SizeReservation
    def __init__(self, size_reservation: _Optional[_Union[_size_reservation_pb2.SizeReservation, _Mapping]] = ...) -> None: ...

class SizeReservationServiceListRequest(_message.Message):
    __slots__ = ("query",)
    QUERY_FIELD_NUMBER: _ClassVar[int]
    query: _size_reservation_pb2.SizeReservationQuery
    def __init__(self, query: _Optional[_Union[_size_reservation_pb2.SizeReservationQuery, _Mapping]] = ...) -> None: ...

class SizeReservationServiceListResponse(_message.Message):
    __slots__ = ("size_reservations",)
    SIZE_RESERVATIONS_FIELD_NUMBER: _ClassVar[int]
    size_reservations: _containers.RepeatedCompositeFieldContainer[_size_reservation_pb2.SizeReservation]
    def __init__(self, size_reservations: _Optional[_Iterable[_Union[_size_reservation_pb2.SizeReservation, _Mapping]]] = ...) -> None: ...
