import requests

from time import sleep
from typing import Optional
from plans.models.host import Host

OPTIMISTIC_PROVIDE_HOST = "localhost"
OPTIMISTIC_PROVIDE_PORT = "7000"
BASE_URL = f"http://{OPTIMISTIC_PROVIDE_HOST}:{OPTIMISTIC_PROVIDE_PORT}"
ITERATIONS = 1000


def run(host_id: Optional[str]):
    host: Host

    if host_id is None or host_id == "":
        print("Creating new host...")
        r = requests.post(f"{BASE_URL}/hosts", json={"name": "Provider"})
        host = Host.from_dict(r.json())
    else:
        r = requests.get(f"{BASE_URL}/hosts/{host_id}")
        host = Host.from_dict(r.json())

    if host.bootstrapped_at is None:
        print("Bootstrapping host...")
        requests.post(f"{BASE_URL}/hosts/{host.host_id}/bootstrap")

    while True:
        print("Checking if routing table is sufficiently filled")
        r = requests.get(f"{BASE_URL}/hosts/{host.host_id}/routing-table")
        if len(r.json()) < 180:
            print(f"  Filled by {len(r.json())} peers. Checking again in 5s")
            sleep(5)  # 5s
            continue
        break

    for i in range(ITERATIONS):
        print(f"Providing content SINGLE_QUERY #{i}...")
        r = requests.post(f"{BASE_URL}/hosts/{host.host_id}/provides", json={"type": "SINGLE_QUERY"})
        provide = r.json()
        while True:
            r = requests.get(f"{BASE_URL}/hosts/{host.host_id}/provides/{provide['provideId']}")
            provide_details = r.json()
            if provide_details['endedAt'] is None:
                print(f"Provide operation {provide['provideId']} is still in progress")
                sleep(1)  # 5s
                continue
            break
        print(f"Provide operation {provide['provideId']} finished!")


if __name__ == "__main__":
    run("16Uiu2HAmQUdLitoNuEyNePb3LZtQV16adu7tjgHzk2vT8kfsv6V9")
