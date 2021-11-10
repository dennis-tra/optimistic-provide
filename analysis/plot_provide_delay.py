import seaborn as sns
import numpy as np
import matplotlib.pyplot as plt
from typing import Tuple, List
from matplotlib import ticker

from model_loader import ModelLoader


def plot():
    sns.set_theme()

    add_provider_delay = []

    measurements = ModelLoader.open("../data")
    for measurement in measurements:

        selected_peers: List[Tuple[str, float]] = []

        first_provide = None
        for span in measurement.provider.spans:
            if span.type != "ADD_PROVIDER":
                continue
            selected_peers += [(span.peer_id, span.rel_start)]

        for peer in selected_peers:
            peer_info = measurement.provider.peer_infos[peer[0]]
            if peer_info.discovered_from == "":
                add_provider_delay += [peer[1]]
            else:
                add_provider_delay += [peer[1] - peer_info.rel_discovered_at]

    fig, ax = plt.subplots(figsize=(15, 6))

    p50 = "%.2f" % np.percentile(add_provider_delay, 50)
    p90 = "%.2f" % np.percentile(add_provider_delay, 90)
    p95 = "%.2f" % np.percentile(add_provider_delay, 95)

    sns.histplot(ax=ax, x=add_provider_delay, bins=np.arange(0, 100))
    ax.set_xlabel("Time in s")
    ax.set_ylabel("Count")
    ax.get_yaxis().set_major_formatter(ticker.FuncFormatter(lambda x, p: "%.1fk" % (x / 1000)))

    ax.title.set_text(
        f"Peer Discovery to ADD_PROVIDER RPC Delay (Sample Size {len(add_provider_delay)}), {p50}s (p50), {p90}s (p90), {p95}s (p95)")

    plt.tight_layout()
    plt.savefig("../plots/provide_delay.png")
    plt.show()


if __name__ == '__main__':
    plot()
