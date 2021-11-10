import numpy as np
import seaborn as sns
from matplotlib import ticker

from model_loader import ModelLoader
import matplotlib.pyplot as plt

folder = "../results"


def plot():
    sns.set_theme()

    discover_latencies = []

    measurements = ModelLoader.open("../data")
    for measurement in measurements:
        for provider_id in measurement.provider.providers:
            peer_info = measurement.provider.peer_infos[provider_id]
            discover_latencies += [peer_info.rel_discovered_at]

    fig, ax = plt.subplots(figsize=(15, 6))

    p50 = "%.2f" % np.percentile(discover_latencies, 50)
    p90 = "%.2f" % np.percentile(discover_latencies, 90)
    p95 = "%.2f" % np.percentile(discover_latencies, 95)

    sns.histplot(ax=ax, x=discover_latencies, bins=np.arange(0, 2.5, 0.05))
    ax.set_xlabel("Time in s")
    ax.set_ylabel("Count")

    ax.get_yaxis().set_major_formatter(ticker.FuncFormatter(lambda x, p: "%.1fk" % (x / 1000)))
    ax.title.set_text(f"Discover Provider Latency Distribution (Sample Size {len(discover_latencies)}), {p50}s (p50), {p90}s (p90), {p95}s (p95)")

    plt.tight_layout()
    plt.savefig("../plots/discover_latencies.png")
    plt.show()


if __name__ == '__main__':
    plot()
