import os

import numpy as np
import seaborn as sns

from model_loader import ModelLoader
from model_measurement import Measurement
import matplotlib.pyplot as plt

folder = "../results"


def plot():
    sns.set_theme()

    measurements = ModelLoader.open("../measurements")

    distances_success = []
    distances_error = []
    for measurement in measurements:
        for span in measurement.provider.spans:
            if span.type != "ADD_PROVIDER":
                continue

            peer_info = measurement.provider.peer_infos[span.peer_id]
            if span.has_error:
                distances_success += [peer_info.distance_pct]
            else:
                distances_error += [peer_info.distance_pct]

    fig, ax = plt.subplots(figsize=(15, 5))

    sns.histplot(ax=ax, x=distances_success, bins=np.arange(50) / 100)
    sns.histplot(ax=ax, x=distances_error, bins=np.arange(50) / 100)
    ax.set_ylabel("Count")
    ax.set_xlabel("Normed XOR Distance in %")

    plt.title("Selected Peers by XOR Target Distance")
    plt.tight_layout()
    plt.show()


if __name__ == '__main__':
    plot()
