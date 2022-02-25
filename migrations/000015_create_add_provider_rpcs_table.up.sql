BEGIN;

CREATE TABLE add_provider_rpcs
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

CREATE TABLE provides_x_add_provider_rpcs
(
    provide_id        INT NOT NULL,
    add_provider_rpc_id INT NOT NULL,

    CONSTRAINT fk_provides_x_add_provider_rpcs_provide_id FOREIGN KEY (provide_id) REFERENCES provides (id),
    CONSTRAINT fk_provides_x_add_provider_rpcs_add_provider_rpc_id FOREIGN KEY (add_provider_rpc_id) REFERENCES add_provider_rpcs (id),

    PRIMARY KEY (provide_id, add_provider_rpc_id)
);

COMMIT;
