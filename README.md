# Tenma Probe

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg)](https://github.com/RichardLitt/standard-readme)
[![readme tenma](https://img.shields.io/badge/readme-Tenma-blueviolet)](README.md)
[![GitHub license](https://img.shields.io/github/license/dennis-tra/tenma-probe)](https://github.com/dennis-tra/tenma-probe/blob/main/LICENSE)

[Tenma](https://en.wikipedia.org/wiki/Tenma) is a libp2p DHT performance measurement tool. As of now, it primarily measures the performance of providing content in the network.

## Table of Contents

- [Project Status](#project-status)
- [Usage](#usage)
- [Install](#install)
    - [Release download](#release-download) | [From source](#from-source)
- [Development](#development)
- [Deployment](#deployment)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [Support](#support)
- [Other Projects](#other-projects)
- [License](#license)
- [Results](#results)

## Project Status

TODO

## Usage

TODO

## Install

### Release download

There is no release yet.

### From source

TODO

## Development

TODO

## Deployment

TODO

## Maintainers

[@dennis-tra](https://github.com/dennis-tra).

## Contributing

Feel free to dive in! [Open an issue](https://github.com/dennis-tra/tenma-probe/issues/new) or submit PRs.

## Support

It would really make my day if you supported this project through [Buy Me A Coffee](https://www.buymeacoffee.com/dennistra).

## Other Projects

You may be interested in one of my other projects:

- [`pcp`](https://github.com/dennis-tra/pcp) - Command line peer-to-peer data transfer tool based on [libp2p](https://github.com/libp2p/go-libp2p).
- [`nebula-crawler`](https://github.com/dennis-tra/nebula-crawler) - A libp2p DHT crawler, monitor, and measurement tool that exposes timely information about DHT networks.
- [`image-stego`](https://github.com/dennis-tra/image-stego) - A novel way to image manipulation detection. Steganography-based image integrity - Merkle tree nodes embedded into image chunks so that each chunk's integrity can be verified on its own.

## License

[Apache License Version 2.0](LICENSE) © Dennis Trautwein

---

## Optimistic Provide

## Abstract

The lifecycle of content in the IPFS network can be divided in three stages: publication, discovery and retrieval.
In the past much work has focussed on improving content discovery while efficiently storing provider records at appropriate peers is similarly important as it needs to be repeated periodically.
This document proposes an optimistic approach to storing provider records in the libp2p Kademlia DHT to significantly speed up the process 
based on a priori information about the network size and accepting minimal too many provider records.

## Motivation

When IPFS attempts to store a provider record in the DHT it tries to find the _beta_ (20) closest peers to the corresponding `CID` (XOR distance).
To find these peers IPFS sends `FIND_NODES` RPCs to the closest peers it has in its routing table and then repeats the process for the returned set of peers.
There are two termination conditions for this process:

1. Termination: The current _beta_ closest peers were queried for even closer peers but didn't yield closer ones.
2. Starvation: All peers in the network were queried (if I interpret this condition correctly: `q.queryPeers.NumHeard() == 0 && q.queryPeers.NumWaiting() == 0`)

This can lead to huge delays if some of the 20 closest peers don't respond timely or are straight out not reachable.
The following graph shows the provide latency distribution.

![](./plots/provide_latencies.png)

In other words, it shows the distribution of how long it takes for the `dht.Provide(ctx, content.cid, true)` call to return.
At the top of the graph you can find the percentiles and total sample size. There is a huge spike at around 10s which is probably related to an exceeded context deadline - not sure though.

If we on the other hand look at how long it took to find the peers that we eventually stored the provider records at, we see that it takes less than 1.6s for the vast majority of cases.

![](./plots/find_latencies.png)

The sample size corresponds to roughly `623 * 20` as in every `Provide`-run we attempt to save the provider record at _beta_ (20) peers. I'm not sure why the Sample Size is not exactly `623 * 20 = 12460`.

This repository also contains code to visualize the provide process. Here is an example:

![](./plots/provide_process.png)

## Methodology

### Measurement Setup

The measurements were conducted on the following machine:

- `vCPU` - `2`
- `RAM` - `4GB`
- `Disk` - `40GB`
- `Datacenter` - `nbg1-dc3`
- `Country` - `Germany`
- `City` - `Nuremberg`

The following results show measurement data that was collected from 2021-11-05 to 2021-11-07.

- Number of measurements `621`


### Normed XOR Distance

TODO
