CREATE TABLE closer_peers
(
    provide_id   INT NOT NULL,
    find_node_id INT NOT NULL,
    peer_id      INT NOT NULL,

    CONSTRAINT fk_closer_peers_provide_id FOREIGN KEY (provide_id) REFERENCES provides (id),
    CONSTRAINT fk_closer_peers_find_node_id FOREIGN KEY (find_node_id) REFERENCES find_nodes (id),
    CONSTRAINT fk_closer_peers_peer_id FOREIGN KEY (peer_id) REFERENCES peers (id),

    PRIMARY KEY (provide_id, find_node_id, peer_id)
);
