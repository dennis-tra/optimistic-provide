import numpy as np
import scipy as sc

import seaborn as sns
import matplotlib.pyplot as plt

sns.set_theme()

# CONSTANTS
NETWORK_SIZE = 10_000
LOOKUPS = 100
DEPTH = 8
BIN_WIDTH = LOOKUPS / 10
k = DEPTH


def avg_est(D_i):
    return sum(i / d - 1 for i, d in enumerate(D_i, 1)) / k


def lsq_est(D_i):
    """
    Least Squares
    """
    LSQ_CONST = k * (k + 1) * (2 * k + 1) / 6
    return LSQ_CONST / sum(i * d for i, d in enumerate(D_i, 1)) - 1


dists = {}
for i in range(DEPTH):
    dists[i] = []

# Generate random peer IDs in the range from 0 to 1
peer_ids = np.random.uniform(0, 1, NETWORK_SIZE)

for i in range(LOOKUPS):
    # Pick random point in the keyspace
    random_point = np.random.random()

    # Calculate distances of peer IDs to that point. Use abs due to XOR.
    distances = np.abs(peer_ids - random_point)

    # Sort the distances increasing
    dist_sorted = np.sort(distances)

    for j in range(DEPTH):
        dists[j] += [dist_sorted[j]]

distances = dists[DEPTH-1]
fig, ax = plt.subplots()
x = np.arange(0, 0.003, 1 / 100000)
sns.histplot(ax=ax, x=distances, bins=x)

rv = sc.stats.beta(DEPTH, NETWORK_SIZE - DEPTH + 1)
sns.lineplot(ax=ax, x=x, y=rv.pdf(x) / (NETWORK_SIZE / BIN_WIDTH))
#
# mean = np.mean(distances[:LOOKUPS])
#
# print(k / mean - 1)
#
# ax.axvline(x=mean)
plt.tight_layout()
plt.show()
