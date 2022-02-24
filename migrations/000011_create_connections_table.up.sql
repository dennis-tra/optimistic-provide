BEGIN;

CREATE TABLE connections
(
    id               INT GENERATED ALWAYS AS IDENTITY,
    provide_id       INT,
    retrieval_id     INT,
    local_id         INT            NOT NULL,
    remote_id        INT            NOT NULL,
    multi_address_id INT            NOT NULL,
    started_at       TIMESTAMPTZ    NOT NULL,
    ended_at         TIMESTAMPTZ    NOT NULL,

    CONSTRAINT fk_connections_provide_id FOREIGN KEY (provide_id) REFERENCES provides (id),
    CONSTRAINT fk_connections_retrieval_id FOREIGN KEY (retrieval_id) REFERENCES retrievals (id),
    CONSTRAINT fk_connections_local_id FOREIGN KEY (local_id) REFERENCES peers (id),
    CONSTRAINT fk_connections_remote_id FOREIGN KEY (remote_id) REFERENCES peers (id),
    CONSTRAINT fk_connections_multi_address_id FOREIGN KEY (multi_address_id) REFERENCES multi_addresses (id),

    PRIMARY KEY (id)
);

COMMIT;
