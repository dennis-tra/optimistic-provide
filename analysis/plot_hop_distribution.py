import numpy as np
import seaborn as sns

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

    hop_count_distribution = []

    measurements = ModelLoader.open("../data")
    for measurement in measurements:
        for span in measurement.provider.spans:
            if span.type != "ADD_PROVIDER":
                continue
            hop_count = calc_hop_count(span.peer_id, measurement.provider.peer_infos, 0)
            hop_count_distribution += [hop_count]

    fig, ax = plt.subplots(figsize=(15, 6))

    sns.histplot(ax=ax, x=hop_count_distribution, bins=np.arange(0, 10))

    ax.set_xticks(np.arange(0, 10))
    ax.set_xlabel("Number of Hops")
    ax.set_ylabel("Count (log scale)")
    ax.set_yscale('log')

    ax.title.set_text(
        f"Number of Hops to Discover a Peer that was Selected to Store a Provider Record (Sample Size {len(hop_count_distribution)})")

    plt.tight_layout()

    show_values_on_bars(ax, len(hop_count_distribution))

    plt.savefig("../plots/hop_distribution.png")
    plt.show()


if __name__ == '__main__':
    plot()
