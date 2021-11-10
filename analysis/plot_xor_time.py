import seaborn as sns
import matplotlib.pyplot as plt

from model_loader import ModelLoader


def plot():
    sns.set_theme()

    distances = []
    provide_times = []

    measurements = ModelLoader.open("../data")
    for measurement in measurements:
        first_provide = None
        for span in measurement.provider.spans:
            if span.type != "ADD_PROVIDER":
                continue

            if first_provide is None or span.rel_start < first_provide:
                first_provide = span.rel_start

        provide_times += [first_provide]
        distances += [int(measurement.provider.distance, base=16) / (2 ** 256) * 100]

    fig, ax = plt.subplots(figsize=(15, 6))

    ax.scatter(distances, provide_times)
    ax.set_ylabel("Time in s")
    ax.set_xlabel("Normed XOR Distance")

    ax.title.set_text(f"Time to First ADD_PROVIDER RPC by Distance of Providing Peer to CID (Sample Size {len(distances)})")

    plt.tight_layout()
    plt.show()


if __name__ == '__main__':
    plot()
