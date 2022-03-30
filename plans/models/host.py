from dataclasses import dataclass
from typing import Any, TypeVar, Type, cast

T = TypeVar("T")


def from_str(x: Any) -> str:
    assert isinstance(x, str)
    return x


def from_int(x: Any) -> int:
    assert isinstance(x, int) and not isinstance(x, bool)
    return x


def to_class(c: Type[T], x: Any) -> dict:
    assert isinstance(x, c)
    return cast(Any, x).to_dict()


@dataclass
class Host:
    host_id: str
    name: str
    started_at: str
    created_at: str
    bootstrapped_at: str
    running_provides_count: int

    @staticmethod
    def from_dict(obj: Any) -> 'Host':
        assert isinstance(obj, dict)
        host_id = from_str(obj.get("hostId"))
        name = from_str(obj.get("name"))
        started_at = from_str(obj.get("startedAt"))
        created_at = from_str(obj.get("createdAt"))
        bootstrapped_at = from_str(obj.get("bootstrappedAt"))
        running_provides_count = from_int(obj.get("runningProvidesCount"))
        return Host(host_id, name, started_at, created_at, bootstrapped_at, running_provides_count)

    def to_dict(self) -> dict:
        result: dict = {
            "hostId": from_str(self.host_id),
            "name": from_str(self.name),
            "startedAt": from_str(self.started_at),
            "createdAt": from_str(self.created_at),
            "bootstrappedAt": from_str(self.bootstrapped_at),
            "runningProvidesCount": from_int(self.running_provides_count)
        }
        return result
