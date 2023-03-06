from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Account(_message.Message):
    __slots__ = ["id", "telegram_id", "username", "fullname", "created_at"]
    ID_FIELD_NUMBER: _ClassVar[int]
    TELEGRAM_ID_FIELD_NUMBER: _ClassVar[int]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    FULLNAME_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    id: int
    telegram_id: int
    username: str
    fullname: str
    created_at: _timestamp_pb2.Timestamp
    def __init__(self, id: _Optional[int] = ..., telegram_id: _Optional[int] = ..., username: _Optional[str] = ..., fullname: _Optional[str] = ..., created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class CreateRequest(_message.Message):
    __slots__ = ["telegram_id", "username", "fullname"]
    TELEGRAM_ID_FIELD_NUMBER: _ClassVar[int]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    FULLNAME_FIELD_NUMBER: _ClassVar[int]
    telegram_id: int
    username: str
    fullname: str
    def __init__(self, telegram_id: _Optional[int] = ..., username: _Optional[str] = ..., fullname: _Optional[str] = ...) -> None: ...

class CreateReply(_message.Message):
    __slots__ = ["id"]
    ID_FIELD_NUMBER: _ClassVar[int]
    id: int
    def __init__(self, id: _Optional[int] = ...) -> None: ...
