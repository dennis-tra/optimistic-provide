BEGIN;

CREATE TABLE add_providers
(
    id                INT GENERATED ALWAYS AS IDENTITY,
    provide_id        INT         NOT NULL,
    local_id          INT         NOT NULL,
    remote_id         INT         NOT NULL,
    distance          bytea       NOT NULL,
    multi_address_ids INT[]       NOT NULL,
    started_at        TIMESTAMPTZ NOT NULL,
    ended_at          TIMESTAMPTZ NOT NULL,
    error             TEXT,

    CONSTRAINT fk_add_providers_provide_id FOREIGN KEY (provide_id) REFERENCES provides (id),
    CONSTRAINT fk_add_providers_local_id FOREIGN KEY (local_id) REFERENCES peers (id),
    CONSTRAINT fk_add_providers_remote_id FOREIGN KEY (remote_id) REFERENCES peers (id),

    PRIMARY KEY (id)
);

COMMIT;
