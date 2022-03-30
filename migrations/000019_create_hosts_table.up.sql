BEGIN;

CREATE TABLE hosts
(
    id          INT GENERATED ALWAYS AS IDENTITY,
    peer_id     INT         NOT NULL,
    name        TEXT        NOT NULL,
    private_key bytea       NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL,
    updated_at  TIMESTAMPTZ NOT NULL,
    archived_at TIMESTAMPTZ,

    CONSTRAINT fk_hosts_peer_id FOREIGN KEY (peer_id) REFERENCES peers (id),

    PRIMARY KEY (id)
);

COMMIT;
