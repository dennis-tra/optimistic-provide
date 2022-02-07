-- The `peers` table keeps track of all peers ever found in the DHT
CREATE TABLE peers
(
    id         INT GENERATED ALWAYS AS IDENTITY,
    multi_hash TEXT        NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    PRIMARY KEY (id)
);
