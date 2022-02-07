import numpy as np

import seaborn as sns
import matplotlib.pyplot as plt

sns.set_theme()

NETWORK_SIZE = 10_000

distances = []
peer_ids = np.random.uniform(0, 1, NETWORK_SIZE)
peer_ids = np.sort(peer_ids)

for i in range(len(peer_ids)):
    if i + 1 == len(peer_ids):
        break
    peer_1 = peer_ids[i]
    peer_2 = peer_ids[i + 1]

    distances += [peer_2 - peer_1]

fig, ax = plt.subplots(figsize=(15, 5))


def exponential(x, a, b, c=0):
    return a * np.exp(- x * b) + c


bins = np.arange(0, 10 / NETWORK_SIZE, 0.1/NETWORK_SIZE)
sns.histplot(ax=ax, x=distances, bins=bins)
sns.lineplot(ax=ax, x=bins, y=exponential(bins, NETWORK_SIZE / 10, NETWORK_SIZE))
ax.set_ylabel("Count")
ax.set_xlabel("Normed XOR Distance in %")

print("MEAN", np.mean(distances))
print("MEDIAN", np.median(distances))
ax.axvline(x=np.mean(distances))
ax.axvline(x=np.median(distances), color='r')

plt.title(f"Selected Peers by XOR Target Distance (simulated)")
plt.tight_layout()
plt.show()
