from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class PlacementType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
    OFFER: _ClassVar[PlacementType]
    ORDER: _ClassVar[PlacementType]
OFFER: PlacementType
ORDER: PlacementType

class Placement(_message.Message):
    __slots__ = ["id", "telegram_id", "username", "typ", "title", "content", "created_at"]
    ID_FIELD_NUMBER: _ClassVar[int]
    TELEGRAM_ID_FIELD_NUMBER: _ClassVar[int]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    TYP_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    CONTENT_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    id: int
    telegram_id: int
    username: str
    typ: PlacementType
    title: str
    content: str
    created_at: _timestamp_pb2.Timestamp
    def __init__(self, id: _Optional[int] = ..., telegram_id: _Optional[int] = ..., username: _Optional[str] = ..., typ: _Optional[_Union[PlacementType, str]] = ..., title: _Optional[str] = ..., content: _Optional[str] = ..., created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class CreateRequest(_message.Message):
    __slots__ = ["telegram_id", "username", "typ", "title", "content"]
    TELEGRAM_ID_FIELD_NUMBER: _ClassVar[int]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    TYP_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    CONTENT_FIELD_NUMBER: _ClassVar[int]
    telegram_id: int
    username: str
    typ: PlacementType
    title: str
    content: str
    def __init__(self, telegram_id: _Optional[int] = ..., username: _Optional[str] = ..., typ: _Optional[_Union[PlacementType, str]] = ..., title: _Optional[str] = ..., content: _Optional[str] = ...) -> None: ...

class CreateReply(_message.Message):
    __slots__ = ["id"]
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...
