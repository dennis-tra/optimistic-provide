CREATE TABLE closer_peers
(
    id                   INT GENERATED ALWAYS AS IDENTITY,
    find_node_rpc_id     INT,
    get_providers_rpc_id INT,
    peer_id              INT   NOT NULL,
    multi_address_ids    INT[] NOT NULL,

    CONSTRAINT fk_closer_peers_find_node_rpc_id FOREIGN KEY (find_node_rpc_id) REFERENCES find_nodes_rpcs (id),
    CONSTRAINT fk_closer_peers_get_providers_rpc_id FOREIGN KEY (get_providers_rpc_id) REFERENCES get_providers_rpcs (id),
    CONSTRAINT fk_closer_peers_peer_id FOREIGN KEY (peer_id) REFERENCES peers (id),

    PRIMARY KEY (id)
);
