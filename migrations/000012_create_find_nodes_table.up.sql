BEGIN;

CREATE TABLE find_nodes
(
    id                 INT GENERATED ALWAYS AS IDENTITY,
    provide_id         INT         NOT NULL,
    local_id           INT         NOT NULL,
    remote_id          INT         NOT NULL,
    started_at         TIMESTAMPTZ NOT NULL,
    ended_at           TIMESTAMPTZ NOT NULL,
    closer_peers_count INT,
    error              TEXT,

    CONSTRAINT fk_dials_provide_id FOREIGN KEY (provide_id) REFERENCES provides (id),
    CONSTRAINT fk_dials_local_id FOREIGN KEY (local_id) REFERENCES peers (id),
    CONSTRAINT fk_dials_remote_id FOREIGN KEY (remote_id) REFERENCES peers (id),

    PRIMARY KEY (id)
);

COMMIT;
