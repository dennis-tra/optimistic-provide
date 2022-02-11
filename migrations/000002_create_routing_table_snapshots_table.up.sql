-- routing_tables captures the routing table of a particular peer at one specific point in time
CREATE TABLE routing_table_snapshots
(
    id          INT GENERATED ALWAYS AS IDENTITY,
    peer_id     INT         NOT NULL,
    bucket_size INT         NOT NULL,
    entry_count INT         NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL,

    CONSTRAINT fk_routing_table_snapshots_peer_id FOREIGN KEY (peer_id) REFERENCES peers (id),
    PRIMARY KEY (id)
);
