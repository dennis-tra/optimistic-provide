BEGIN;

CREATE TABLE get_providers
(
    id                 INT GENERATED ALWAYS AS IDENTITY,
    retrieval_id       INT         NOT NULL,
    local_id           INT         NOT NULL,
    remote_id          INT         NOT NULL,
    started_at         TIMESTAMPTZ NOT NULL,
    ended_at           TIMESTAMPTZ NOT NULL,
    closer_peers_count INT,
    error              TEXT,

    CONSTRAINT fk_get_providers_retrieval_id FOREIGN KEY (retrieval_id) REFERENCES retrievals (id),
    CONSTRAINT fk_get_providers_local_id FOREIGN KEY (local_id) REFERENCES peers (id),
    CONSTRAINT fk_get_providers_remote_id FOREIGN KEY (remote_id) REFERENCES peers (id),

    PRIMARY KEY (id)
);

COMMIT;
