from datetime import datetime
from typing import Any
from model_utils import from_str, from_float, from_datetime, to_float


class PeerInfo:
    id: str
    agent_version: str
    distance: str
    rel_discovered_at: float
    discovered_at: datetime
    discovered_from: str

    def __init__(self, id: str, agent_version: str, distance: str, rel_discovered_at: float, discovered_at: datetime,
                 discovered_from: str) -> None:
        self.id = id
        self.agent_version = agent_version
        self.distance = distance
        self.rel_discovered_at = rel_discovered_at
        self.discovered_at = discovered_at
        self.discovered_from = discovered_from

    @staticmethod
    def from_dict(obj: Any) -> 'PeerInfo':
        assert isinstance(obj, dict)
        id = from_str(obj.get("ID"))
        agent_version = from_str(obj.get("AgentVersion"))
        distance = from_str(obj.get("Distance"))
        rel_discovered_at = from_float(obj.get("RelDiscoveredAt"))
        discovered_at = from_datetime(obj.get("DiscoveredAt"))
        discovered_from = from_str(obj.get("DiscoveredFrom"))
        return PeerInfo(id, agent_version, distance, rel_discovered_at, discovered_at, discovered_from)

    def to_dict(self) -> dict:
        result: dict = {}
        result["ID"] = from_str(self.id)
        result["AgentVersion"] = from_str(self.agent_version)
        result["Distance"] = from_str(self.distance)
        result["RelDiscoveredAt"] = to_float(self.rel_discovered_at)
        result["DiscoveredAt"] = self.discovered_at.isoformat()
        result["DiscoveredFrom"] = from_str(self.discovered_from)
        return result

    @property
    def distance_fraction(self):
        return int(self.distance, base=16) / (2 ** 256)

    @property
    def distance_pct(self):
        return self.distance_fraction * 100
