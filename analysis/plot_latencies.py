import os

import numpy as np
import seaborn as sns
from model_measurement import Measurement
import matplotlib.pyplot as plt

folder = "../results"


def plot():
    sns.set_theme()

    provide_latencies = []
    request_latencies = []
    discover_provider_latencies = []
    for fn in os.listdir(folder):
        filepath = os.path.join(folder, fn)
        if not os.path.isfile(filepath):
            continue

        print("Loading", filepath)
        measurement = Measurement.from_file(filepath)

        provide_latencies += [measurement.provider.duration.total_seconds()]

        for requester_id in measurement.requesters:
            request_run = measurement.requesters[requester_id]
            request_latencies += [request_run.duration.total_seconds()]

        for provider_id in measurement.provider.providers:
            peer_info = measurement.provider.peer_infos[provider_id]
            if peer_info.discovered_from == "":
                continue
            discover_provider_latencies += [peer_info.rel_discovered_at]

    fig, ((ax1, ax2), (ax3, ax4)) = plt.subplots(2, 2, figsize=(15, 10))

    sns.histplot(ax=ax1, x=provide_latencies, bins=np.arange(0, 120, 5))
    ax1.set_xlabel("Time in s")
    ax1.set_ylabel("Count")
    ax1.title.set_text(f"Provide Latency Distribution (Sample Size {len(provide_latencies)})")

    sns.histplot(ax=ax2, x=request_latencies, bins=np.arange(0, 3, 0.1))
    ax2.set_xlabel("Time in s")
    ax2.set_ylabel("Count")
    ax2.title.set_text(f"Find Provider Latency Distribution (Sample Size {len(request_latencies)})")

    sns.histplot(ax=ax3, x=discover_provider_latencies, bins=np.arange(0, 3, 0.1))
    ax3.set_xlabel("Time in s")
    ax3.set_ylabel("Count")
    ax3.title.set_text(f"Find Provider Latency Distribution (Sample Size {len(discover_provider_latencies)})")

    plt.tight_layout()
    plt.show()


if __name__ == '__main__':
    plot()
