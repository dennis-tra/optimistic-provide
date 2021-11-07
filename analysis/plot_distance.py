import pandas as pd
import numpy as np
import seaborn as sns

from model_loader import ModelLoader
import matplotlib.pyplot as plt

folder = "../results"


def plot():
    sns.set_theme()

    measurements = ModelLoader.open("../measurements")

    # for measurement in measurements:
    #     count = 0
    #     for span in measurement.provider.spans:
    #         if span.type == "ADD_PROVIDER":
    #             count += 1
    #     if count != 20:
    #         print(measurement.started_at, count)

    distances = []
    has_errors = []
    for measurement in measurements:
        for span in measurement.provider.spans:
            if span.type != "ADD_PROVIDER":
                continue

            peer_info = measurement.provider.peer_infos[span.peer_id]
            distances += [peer_info.distance_pct]
            has_errors += [span.has_error]

    combined = pd.DataFrame({
        "distances": distances,
        "error": has_errors
    })
    fig, ax = plt.subplots(figsize=(15, 5))

    sns.histplot(
        ax=ax,
        data=combined,
        x="distances",
        bins=np.arange(50) / 100,
        multiple="stack",
        hue="error",
        legend=True)
    ax.set_ylabel("Count")
    ax.set_xlabel("Normed XOR Distance in %")

    plt.title(
        f"Selected Peers by XOR Target Distance (Measurements {len(measurements)}, ADD_PROVIDER RPCs {len(combined)})")
    plt.tight_layout()
    plt.show()


if __name__ == '__main__':
    plot()
