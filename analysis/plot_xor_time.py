import seaborn as sns
import matplotlib.pyplot as plt

from model_measurement import Measurement


def plot():
    sns.set_theme()

    measurement = Measurement.from_file("../out/2021-11-04T16:20_measurement_002.json")

    distances = []
    provide_times = []
    first_provide = None
    for span in measurement.provider.spans:
        if span.type != "send_message":
            continue

        if first_provide is None or span.rel_start < first_provide:
            first_provide = span.rel_start

    provide_times += [first_provide]
    distances += [int(measurement.provider.distance, base=16) / (2 ** 256) * 100]

    fig, ax = plt.subplots(figsize=(15, 5))

    ax.scatter(distances, provide_times)
    ax.set_ylabel("Time in s")
    ax.set_xlabel("Normed XOR Distance")

    plt.tight_layout()
    plt.show()


if __name__ == '__main__':
    plot()
