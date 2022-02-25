BEGIN;

CREATE TABLE find_nodes_rpcs
(
    id                 INT GENERATED ALWAYS AS IDENTITY,
    local_id           INT         NOT NULL,
    remote_id          INT         NOT NULL,
    started_at         TIMESTAMPTZ NOT NULL,
    ended_at           TIMESTAMPTZ NOT NULL,
    closer_peers_count INT,
    error              TEXT,

    CONSTRAINT fk_find_nodes_local_id FOREIGN KEY (local_id) REFERENCES peers (id),
    CONSTRAINT fk_find_nodes_remote_id FOREIGN KEY (remote_id) REFERENCES peers (id),

    PRIMARY KEY (id)
);

CREATE TABLE provides_x_find_nodes_rpcs
(
    provide_id        INT NOT NULL,
    find_nodes_rpc_id INT NOT NULL,

    CONSTRAINT fk_provides_x_find_nodes_rpcs_provide_id FOREIGN KEY (provide_id) REFERENCES provides (id),
    CONSTRAINT fk_provides_x_find_nodes_rpcs_find_nodes_rpc_id FOREIGN KEY (find_nodes_rpc_id) REFERENCES find_nodes_rpcs (id),

    PRIMARY KEY (provide_id, find_nodes_rpc_id)
);

COMMIT;
