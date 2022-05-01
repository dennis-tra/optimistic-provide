BEGIN;

CREATE TABLE network_size_estimates
(
    id               INT GENERATED ALWAYS AS IDENTITY,
    peer_id          INT         NOT NULL,
    network_size     FLOAT8      NOT NULL,
    network_size_err FLOAT8      NOT NULL,
    r_squared        FLOAT8      NOT NULL,
    extra            TEXT,
    created_at       TIMESTAMPTZ NOT NULL,

    CONSTRAINT fk_network_size_estimates_peer_id FOREIGN KEY (peer_id) REFERENCES peers (id),

    PRIMARY KEY (id)
);

COMMIT;
