import json
import os

import numpy as np
import matplotlib.pyplot as plt

# Load measurement information
from model_measurement import Measurement
folder = "../results"
discovered_times = []
for fn in os.listdir(folder):
    filepath = os.path.join(folder, fn)
    if not os.path.isfile(filepath):
        continue

    measurement = Measurement.from_file(filepath)

    selected_peers = []
    provided_times = []

    first_provide = None
    for span in measurement.provider.spans:
        if span.operation != "send_message":
            continue
        selected_peers += [span.peer_id]
        provided_times += [span.rel_start]

    for idx, peer in enumerate(selected_peers):
        provided_time = provided_times[idx]
        peer_info = measurement.provider.peer_infos[peer]
        if peer_info.discovered_from == "":
            discovered_times += [provided_time]
        else:
            discovered_times += [provided_time - peer_info.rel_discovered_at]

plt.hist(discovered_times, bins=np.arange(100))

plt.xlabel("Discover Provide Delay in s")
plt.ylabel("Count")

plt.show()
