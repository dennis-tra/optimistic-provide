import numpy as np
import seaborn as sns
from model_measurement import Measurement
import matplotlib.pyplot as plt


def plot():
    sns.set_theme()

    measurement = Measurement.from_file("../out/2021-11-04T16:20_measurement_002.json")

    distances_success = []
    distances_error = []
    for span in measurement.provider.spans:
        if span.type != "ADD_PROVIDER":
            continue

        distance = int(measurement.provider.peer_infos[span.peer_id].distance, base=16) / (2 ** 256) * 100
        if span.error == "":
            distances_success += [distance]
        else:
            distances_error += [distance]

    fig, ax = plt.subplots(figsize=(15, 5))

    sns.histplot(ax=ax, x=distances_success, bins=np.arange(50) / 100)
    sns.histplot(ax=ax, x=distances_error, bins=np.arange(50) / 100)
    ax.set_ylabel("Count")
    ax.set_xlabel("Normed XOR Distance")

    plt.title("Selected Peers by XOR Target Distance")
    plt.tight_layout()
    plt.show()


if __name__ == '__main__':
    plot()
