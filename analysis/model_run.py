from datetime import datetime
from typing import Any, List, Dict
from model_peer_info import PeerInfo
from model_utils import from_datetime, from_str, from_list, from_dict
from model_span import Span


class Run:
    started_at: datetime
    ended_at: datetime
    local_id: str
    distance: str
    spans: List[Span]
    peer_infos: Dict[str, PeerInfo]
    peer_order: List[str]

    def __init__(self, started_at: datetime, ended_at: datetime, local_id: str, distance: str, spans: List[Span],
                 peer_infos: Dict[str, PeerInfo], peer_order: List[str]) -> None:
        self.started_at = started_at
        self.ended_at = ended_at
        self.local_id = local_id
        self.distance = distance
        self.spans = spans
        self.peer_infos = peer_infos
        self.peer_order = peer_order

    @staticmethod
    def from_dict(obj: Any) -> 'Run':
        assert isinstance(obj, dict)
        started_at = from_datetime(obj.get("StartedAt"))
        ended_at = from_datetime(obj.get("EndedAt"))
        local_id = from_str(obj.get("LocalID"))
        distance = from_str(obj.get("Distance"))
        spans = from_list(Span.from_dict, obj.get("Spans"))
        peer_infos = from_dict(PeerInfo.from_dict, obj.get("PeerInfos"))
        peer_order = from_list(from_str, obj.get("PeerOrder"))
        return Run(started_at, ended_at, local_id, distance, spans, peer_infos, peer_order)
