import seaborn as sns
import matplotlib.pyplot as plt
import matplotlib.patches as mpatches

from model_measurement import Measurement

span_colors = {
    "dial": sns.color_palette()[3],
    "send_request": sns.color_palette()[2],
    "send_message": sns.color_palette()[1],
}


def plot(filename: str):
    sns.set_theme(style="darkgrid")

    measurement = Measurement.from_file(filename)

    fig, ax = plt.subplots(1, figsize=(15, 6))

    # Plot the horizontal spans for each peer
    for span in measurement.provider.spans:
        y = measurement.provider.plot_y_position(span.peer_id)
        xmin = span.rel_start
        xmax = xmin + span.duration_s
        ax.hlines(
            y,
            xmin,
            xmax,
            color=span_colors[span.operation],
            linewidth=3,
            alpha=0.5 if span.has_error else 1.0
        )

        if span.has_error:
            ax.plot([xmax], [y], marker='x', color=span_colors[span.operation], markersize=4)
        if span.operation == 'send_message':
            ax.plot([xmin], [y], marker='.', color=span_colors[span.operation], markersize=4)

    # plot the vertical lines indicating causality
    for peer_id in measurement.provider.peer_infos:
        peer_info = measurement.provider.peer_infos[peer_id]
        if peer_info.discovered_from == "":
            continue

        origin = measurement.provider.plot_y_position(peer_info.discovered_from)
        target = measurement.provider.plot_y_position(peer_info.id)
        ax.arrow(
            peer_info.rel_discovered_at,
            origin,
            0,
            target - origin,
            head_width=0.01,
            head_length=0.5,
            color="#ccc",
            length_includes_head=True
        )

    labels = []
    for peer_id in measurement.provider.peer_order:
        distance = int(measurement.provider.peer_infos[peer_id].distance, base=16)
        distance_norm = distance / (2 ** 256)
        labels += [
            "{:s} | {:.2f} | {:s}".format(measurement.provider.peer_infos[peer_id].agent_version, distance_norm * 100,
                                          peer_id[:16])]
    labels.reverse()

    ax.set_yticklabels(
        labels,
        fontsize=8,
        fontname='Monospace',
        fontweight="bold")
    ax.set_yticks(range(1, len(measurement.provider.peer_order) + 1))

    plt.legend(handles=[
        mpatches.Patch(color=span_colors["dial"], label='Dial'),
        mpatches.Patch(color=span_colors["send_request"], label='Find Node'),
        mpatches.Patch(color=span_colors["send_message"], label='Add Provider'),
        mpatches.Patch(color="#ccc", label='Discovered'),
    ])

    plt.title(
        "Providing content with distance {:.2f}".format(int(measurement.provider.distance, base=16) / (2 ** 256) * 100))

    plt.xlabel("Time in s")
    plt.xlim(0)
    plt.tight_layout()
    plt.savefig("../plots/provide_process.png")
    plt.show()


if __name__ == '__main__':
    plot("../data/2021-11-08T09:45_measurement_001.json")
