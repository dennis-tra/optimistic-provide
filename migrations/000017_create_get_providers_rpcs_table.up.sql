BEGIN;

CREATE TABLE get_providers_rpcs
(
    id                   INT GENERATED ALWAYS AS IDENTITY,
    local_id             INT         NOT NULL,
    remote_id            INT         NOT NULL,
    started_at           TIMESTAMPTZ NOT NULL,
    ended_at             TIMESTAMPTZ NOT NULL,
    provider_peers_count INT,
    error                TEXT,

    CONSTRAINT fk_get_providers_rpcs_local_id FOREIGN KEY (local_id) REFERENCES peers (id),
    CONSTRAINT fk_get_providers_rpcs_remote_id FOREIGN KEY (remote_id) REFERENCES peers (id),

    PRIMARY KEY (id)
);

CREATE TABLE retrievals_x_get_providers_rpcs
(
    retrieval_id        INT NOT NULL,
    get_provider_rpc_id INT NOT NULL,

    CONSTRAINT fk_retrievals_x_get_providers_rpcs_retrieval_id FOREIGN KEY (retrieval_id) REFERENCES retrievals (id),
    CONSTRAINT fk_retrievals_x_get_providers_rpcs_get_provider_rpc_id FOREIGN KEY (get_provider_rpc_id) REFERENCES get_providers_rpcs (id),

    PRIMARY KEY (retrieval_id, get_provider_rpc_id)
);

COMMIT;
