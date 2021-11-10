import numpy as np
import seaborn as sns

from model_loader import ModelLoader
import matplotlib.pyplot as plt


def plot():
    sns.set_theme()

    provide_latencies = [m.provider.duration.total_seconds() for m in ModelLoader.open("../data")]

    fig, ax = plt.subplots(figsize=(15, 6))

    p50 = "%.2f" % np.percentile(provide_latencies, 50)
    p90 = "%.2f" % np.percentile(provide_latencies, 90)
    p95 = "%.2f" % np.percentile(provide_latencies, 95)

    sns.histplot(ax=ax, x=provide_latencies, bins=np.arange(0, 200, 2))
    ax.set_xlabel("Time in s")
    ax.set_ylabel("Count")

    ax.title.set_text(f"Provide Latency Distribution (Sample Size {len(provide_latencies)}), {p50}s (p50), {p90}s (p90), {p95}s (p95)")

    plt.tight_layout()
    plt.savefig("../plots/provide_latencies.png")
    plt.show()


if __name__ == '__main__':
    plot()
