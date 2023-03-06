from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class Message(_message.Message):
    __slots__ = ["id", "coords", "text"]
    ID_FIELD_NUMBER: _ClassVar[int]
    COORDS_FIELD_NUMBER: _ClassVar[int]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    id: int
    coords: _containers.RepeatedScalarFieldContainer[float]
    text: str
    def __init__(self, id: _Optional[int] = ..., coords: _Optional[_Iterable[float]] = ..., text: _Optional[str] = ...) -> None: ...

class CreateRequest(_message.Message):
    __slots__ = ["coords", "text"]
    COORDS_FIELD_NUMBER: _ClassVar[int]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    coords: _containers.RepeatedScalarFieldContainer[float]
    text: str
    def __init__(self, coords: _Optional[_Iterable[float]] = ..., text: _Optional[str] = ...) -> None: ...

class CreateReply(_message.Message):
    __slots__ = ["id"]
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...
