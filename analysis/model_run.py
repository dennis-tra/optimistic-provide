from datetime import datetime, timedelta
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

    def plot_y_position(self, peer_id) -> int:
        return len(self.peer_order) - self.peer_order.index(peer_id)

    @property
    def duration(self) -> timedelta:
        return self.ended_at - self.started_at

    @property
    def providers(self) -> set[str]:
        providers: set[str] = set()
        for span in self.spans:
            if span.type == "ADD_PROVIDER":
                providers.add(span.peer_id)
        return providers
