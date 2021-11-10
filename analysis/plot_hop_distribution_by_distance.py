import math
from typing import List, Tuple

import numpy as np
import seaborn as sns
import pandas as pd

from analysis.model_peer_info import PeerInfo
from model_loader import ModelLoader
import matplotlib.pyplot as plt


def show_values_on_bars(axs, total):
    def _show_on_single_plot(ax):
        for p in ax.patches:
            _x = p.get_x() + p.get_width() / 2
            _y = p.get_y() + p.get_height()
            value = '{:.1f}%'.format(100 * p.get_height() / total)
            ax.text(_x, _y, value, ha="center")

    if isinstance(axs, np.ndarray):
        for idx, ax in np.ndenumerate(axs):
            _show_on_single_plot(ax)
    else:
        _show_on_single_plot(axs)


def plot():
    sns.set_theme()

    def calc_hop_count(peer_id: str, peer_infos: dict[str, PeerInfo], hop_count: int) -> int:
        peer_info = peer_infos[peer_id]
        if peer_info.discovered_from == "":
            return hop_count
        return calc_hop_count(peer_info.discovered_from, peer_infos, hop_count + 1)

    data: List[Tuple[float, int]] = []

    measurements = ModelLoader.open("../data")
    for measurement in measurements:
        for span in measurement.provider.spans:
            if span.type != "ADD_PROVIDER":
                continue
            hop_count = calc_hop_count(span.peer_id, measurement.provider.peer_infos, 0)
            xor_distance = math.floor(int(measurement.provider.distance, base=16) / (2 ** 256) * 10) * 10
            data += [(xor_distance, hop_count)]

    data_df = pd.DataFrame(data, columns=['xor_distances', 'Hop Count'])

    fig, ax = plt.subplots(figsize=(15, 6))

    sns.histplot(
        ax=ax,
        x="xor_distances",
        data=data_df,
        hue="Hop Count",
        multiple="dodge",
        edgecolor='white',
        palette=plt.cm.Accent,
        legend=True,
        bins=np.arange(0, 110, 10),
    )

    ax.set_xticks(np.arange(0, 110, 10))

    ax.set_xlabel("Normed XOR Distance in % and multiples of 10")
    ax.set_ylabel("Count of ADD_PROVIDER RPCs")

    ax.title.set_text(
        f"Number of Hops to Discover a Peer that was Selected to Store a Provider Record Based on the Provider <-> CID Distance (Sample Size {len(data_df)})")

    plt.tight_layout()
    plt.savefig("../plots/hop_distribution_by_distance.png")
    plt.show()


if __name__ == '__main__':
    plot()
