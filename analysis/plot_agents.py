import pandas as pd
import numpy as np
import seaborn as sns

from lib_agent import agent_name, go_ipfs_version, go_ipfs_v08_version
from model_loader import ModelLoader
import matplotlib.pyplot as plt


def plot():
    sns.set_theme()

    measurements = ModelLoader.open("../data")

    agent_versions_all = []
    agent_version_providers_success = []
    for measurement in measurements:
        for peer_id in measurement.provider.peer_infos:
            peer_info = measurement.provider.peer_infos[peer_id]
            if peer_info.agent_version != "":
                agent_versions_all += [peer_info.agent_version]

            for span in measurement.provider.spans:
                if span.peer_id == peer_id and span.type == "ADD_PROVIDER" and not span.has_error:
                    agent_version_providers_success += [peer_info.agent_version]
                    break

    combined = pd.DataFrame({"agent_versions": agent_versions_all}).assign(
        agent_name=lambda df: df.agent_versions.apply(agent_name),
        go_ipfs_version=lambda df: df.agent_versions.apply(go_ipfs_version),
        go_ipfs_v08_version=lambda df: df.agent_versions.apply(go_ipfs_v08_version)
    )

    fig, ((ax1, ax2, ax3), (ax4, ax5, ax6)) = plt.subplots(2, 3, figsize=(15, 10))

    sns.histplot(
        ax=ax1,
        x="agent_name",
        data=combined,
        hue="agent_name",
        stat="percent",
        legend=False
    )
    ax1.set_xlabel("")
    ax1.title.set_text(f"{len(combined['agent_name'])}")


    sns.histplot(
        ax=ax2,
        x="go_ipfs_version",
        data=combined,
        hue="go_ipfs_version",
        stat="percent",
        legend=False
    )
    ax2.set_xlabel("")
    ax2.title.set_text(f"{len(combined['go_ipfs_version'])}")

    sns.histplot(
        ax=ax3,
        x="go_ipfs_v08_version",
        data=combined,
        hue="go_ipfs_v08_version",
        stat="percent",
        legend=False
    )
    ax3.set_xlabel("")
    ax3.title.set_text(f"{len(combined['go_ipfs_v08_version'])}")

    combined = pd.DataFrame({"agent_versions": agent_version_providers_success}).assign(
        agent_name=lambda df: df.agent_versions.apply(agent_name),
        go_ipfs_version=lambda df: df.agent_versions.apply(go_ipfs_version),
        go_ipfs_v08_version=lambda df: df.agent_versions.apply(go_ipfs_v08_version)
    )

    sns.histplot(
        ax=ax4,
        x="agent_name",
        data=combined,
        hue="agent_name",
        stat="percent",
        legend=False
    )
    ax4.set_xlabel("")
    ax4.title.set_text(f"{len(combined['agent_name'])}")

    sns.histplot(
        ax=ax5,
        x="go_ipfs_version",
        data=combined,
        hue="go_ipfs_version",
        stat="percent",
        legend=False
    )
    ax5.set_xlabel("")
    ax5.title.set_text(f"{len(combined['go_ipfs_version'])}")

    sns.histplot(
        ax=ax6,
        x="go_ipfs_v08_version",
        data=combined,
        hue="go_ipfs_v08_version",
        stat="percent",
        legend=False
    )
    ax6.set_xlabel("")
    ax6.title.set_text(f"{len(combined['go_ipfs_v08_version'])}")

    # stats_str = f"Measurements {len(measurements)}; ADD_PROVIDER RPCs {len(combined)}"
    # plt.title(f"Selected Peers by XOR Target Distance to Provided CID ({stats_str})")
    plt.tight_layout()
    plt.savefig("../plots/agents.png")
    plt.show()


if __name__ == '__main__':
    plot()
