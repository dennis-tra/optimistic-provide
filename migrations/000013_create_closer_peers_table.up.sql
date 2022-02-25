CREATE TABLE closer_peers
(
    id                INT GENERATED ALWAYS AS IDENTITY,
    find_node_rpc_id  INT   NOT NULL,
    peer_id           INT   NOT NULL,
    multi_address_ids INT[] NOT NULL,

    CONSTRAINT fk_closer_peers_find_node_rpc_id FOREIGN KEY (find_node_rpc_id) REFERENCES find_nodes_rpcs (id),
    CONSTRAINT fk_closer_peers_peer_id FOREIGN KEY (peer_id) REFERENCES peers (id),

    PRIMARY KEY (find_node_rpc_id, peer_id)
);
