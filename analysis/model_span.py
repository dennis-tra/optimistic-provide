from datetime import datetime
from typing import Any
from model_utils import from_str, from_float, from_datetime, to_float


class Span:
    rel_start: float
    duration_s: float
    start: datetime
    end: datetime
    peer_id: str
    operation: str
    type: str
    error: str

    def __init__(self, rel_start: float, duration_s: float, start: datetime, end: datetime, peer_id: str,
                 operation: str, type: str, error: str) -> None:
        self.rel_start = rel_start
        self.duration_s = duration_s
        self.start = start
        self.end = end
        self.peer_id = peer_id
        self.operation = operation
        self.type = type
        self.error = error

    @staticmethod
    def from_dict(obj: Any) -> 'Span':
        assert isinstance(obj, dict)
        rel_start = from_float(obj.get("RelStart"))
        duration_s = from_float(obj.get("DurationS"))
        start = from_datetime(obj.get("Start"))
        end = from_datetime(obj.get("End"))
        peer_id = from_str(obj.get("PeerID"))
        operation = from_str(obj.get("Operation"))
        type = from_str(obj.get("Type"))
        error = from_str(obj.get("Error"))
        return Span(rel_start, duration_s, start, end, peer_id, operation, type, error)

    def to_dict(self) -> dict:
        result: dict = {}
        result["RelStart"] = to_float(self.rel_start)
        result["DurationS"] = to_float(self.duration_s)
        result["Start"] = self.start.isoformat()
        result["End"] = self.end.isoformat()
        result["PeerID"] = from_str(self.peer_id)
        result["Operation"] = from_str(self.operation)
        result["Type"] = from_str(self.type)
        result["Error"] = from_str(self.error)
        return result
