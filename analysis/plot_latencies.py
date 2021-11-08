import numpy as np
import seaborn as sns
import numpy as np

from model_loader import ModelLoader
import matplotlib.pyplot as plt

folder = "../results"


def plot():
    sns.set_theme()

    provide_latencies = []
    request_latencies = []
    discover_provider_latencies = []

    measurements = ModelLoader.open("../measurements")
    for measurement in measurements:
        provide_latencies += [measurement.provider.duration.total_seconds()]

        for requester_id in measurement.requesters:
            request_run = measurement.requesters[requester_id]
            request_latencies += [request_run.duration.total_seconds()]

        for provider_id in measurement.provider.providers:
            peer_info = measurement.provider.peer_infos[provider_id]
            if peer_info.discovered_from == "":
                continue
            discover_provider_latencies += [peer_info.rel_discovered_at]

    print("discover_provider_latencies", np.percentile(discover_provider_latencies, 50))
    print("discover_provider_latencies", np.percentile(discover_provider_latencies, 90))
    print("discover_provider_latencies", np.percentile(discover_provider_latencies, 95))
    fig, (ax) = plt.subplots(figsize=(15, 5))
    sns.histplot(ax=ax, x=discover_provider_latencies, bins=np.arange(0, 3, 0.05))
    ax.set_xlabel("Time in s")
    ax.set_ylabel("Count")
    p50 = "%.2f" % np.percentile(discover_provider_latencies, 50)
    p90 = "%.2f" % np.percentile(discover_provider_latencies, 90)
    p95 = "%.2f" % np.percentile(discover_provider_latencies, 95)
    ax.title.set_text(f"Find Provider Latency Distribution (Sample Size {len(discover_provider_latencies)}) - Percentiles: {p50}s (p50), {p90}s (p90), {p95}s (p95)")
    plt.tight_layout()
    plt.savefig("../plots/find_latencies.png")

    fig, ((ax1, ax2), (ax3, ax4)) = plt.subplots(2, 2, figsize=(15, 10))

    print("provide_latencies", np.percentile(provide_latencies, 50))
    print("provide_latencies", np.percentile(provide_latencies, 90))
    print("provide_latencies", np.percentile(provide_latencies, 95))
    sns.histplot(ax=ax1, x=provide_latencies, bins=np.arange(0, 120, 5))
    ax1.set_xlabel("Time in s")
    ax1.set_ylabel("Count")
    ax1.title.set_text(f"Provide Latency Distribution (Sample Size {len(provide_latencies)})")

    print("request_latencies", np.percentile(request_latencies, 50))
    print("request_latencies", np.percentile(request_latencies, 90))
    print("request_latencies", np.percentile(request_latencies, 95))
    sns.histplot(ax=ax2, x=request_latencies, bins=np.arange(0, 3, 0.1))
    ax2.set_xlabel("Time in s")
    ax2.set_ylabel("Count")
    ax2.title.set_text(f"Find Provider Latency Distribution (Sample Size {len(request_latencies)})")

    print("discover_provider_latencies", np.percentile(discover_provider_latencies, 50))
    print("discover_provider_latencies", np.percentile(discover_provider_latencies, 90))
    print("discover_provider_latencies", np.percentile(discover_provider_latencies, 95))
    sns.histplot(ax=ax3, x=discover_provider_latencies, bins=np.arange(0, 3, 0.1))
    ax3.set_xlabel("Time in s")
    ax3.set_ylabel("Count")
    ax3.title.set_text(f"Find Provider Latency Distribution (Sample Size {len(discover_provider_latencies)})")

    ax4.axis('off')

    plt.tight_layout()
    plt.show()


if __name__ == '__main__':
    plot()
