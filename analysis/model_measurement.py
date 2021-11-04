import json
from datetime import datetime
from typing import Any, Dict
from model_run import Run
from model_utils import from_datetime, from_str, from_dict, from_bool


class Measurement:
    started_at: datetime
    ended_at: datetime
    content_id: str
    provider: Run
    requesters: Dict[str, Run]
    init_rt: bool

    def __init__(self, started_at: datetime, ended_at: datetime, content_id: str, provider: Run,
                 requesters: Dict[str, Run], init_rt: bool) -> None:
        self.started_at = started_at
        self.ended_at = ended_at
        self.content_id = content_id
        self.provider = provider
        self.requesters = requesters
        self.init_rt = init_rt

    @staticmethod
    def from_dict(obj: Any) -> 'Measurement':
        assert isinstance(obj, dict)
        started_at = from_datetime(obj.get("StartedAt"))
        ended_at = from_datetime(obj.get("EndedAt"))
        content_id = from_str(obj.get("ContentID"))
        provider = Run.from_dict(obj.get("Provider"))
        requesters = from_dict(Run.from_dict, obj.get("Requesters"))
        init_rt = from_bool(obj.get("InitRT"))
        return Measurement(started_at, ended_at, content_id, provider, requesters, init_rt)

    @staticmethod
    def from_file(filepath: str) -> 'Measurement':
        with open(filepath) as f:
            return Measurement.from_dict(json.load(f))
