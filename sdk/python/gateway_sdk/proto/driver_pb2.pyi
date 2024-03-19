from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class RequestArgs(_message.Message):
    __slots__ = ["brokerid", "request"]
    BROKERID_FIELD_NUMBER: _ClassVar[int]
    REQUEST_FIELD_NUMBER: _ClassVar[int]
    brokerid: int
    request: str
    def __init__(self, brokerid: _Optional[int] = ..., request: _Optional[str] = ...) -> None: ...

class ResponseResult(_message.Message):
    __slots__ = ["data"]
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: str
    def __init__(self, data: _Optional[str] = ...) -> None: ...
