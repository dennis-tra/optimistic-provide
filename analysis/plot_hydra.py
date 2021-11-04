from model_measurement import Measurement


def main():
    all_peers = 0
    hydra_peers = 0
    seen = {}

    measurement = Measurement.from_file("../out/2021-11-04T16:20_measurement_002.json")
    for peer_info in measurement.provider.peer_infos.values():
        if peer_info.id in seen:
            continue
        seen[peer_info.id] = True
        all_peers += 1
        if "hydra" in peer_info.agent_version:
            hydra_peers += 1

    print("Hydra Peers", hydra_peers)
    print("All Peers", all_peers)
    print("Hydra Peers Percent", hydra_peers / all_peers * 100, "%")


if __name__ == '__main__':
    main()
