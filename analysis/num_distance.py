import numpy as np

import seaborn as sns
import matplotlib.pyplot as plt

sns.set_theme()

NETWORK_SIZE = 10_000
MEASUREMENTS = 1269
BETA = 20

distances_all = []
peer_ids = np.random.uniform(0, 1, NETWORK_SIZE)
for i in range(MEASUREMENTS):
    cid = np.random.random()
    distances = np.sort(np.abs(peer_ids - cid))

    for dist in distances[:BETA]:
        distances_all += [dist * 100]

fig, ax = plt.subplots(figsize=(15, 5))

sns.histplot(ax=ax, x=distances_all, bins=np.arange(0, 50) / 100)
ax.set_ylabel("Count")
ax.set_xlabel("Normed XOR Distance in %")

plt.title(f"Selected Peers by XOR Target Distance (simulated)")
plt.tight_layout()
plt.show()
