BEGIN;

CREATE TABLE connections
(
    id               INT GENERATED ALWAYS AS IDENTITY,
    local_id         INT         NOT NULL,
    remote_id        INT         NOT NULL,
    multi_address_id INT         NOT NULL,
    started_at       TIMESTAMPTZ NOT NULL,
    ended_at         TIMESTAMPTZ NOT NULL,

    CONSTRAINT fk_connections_local_id FOREIGN KEY (local_id) REFERENCES peers (id),
    CONSTRAINT fk_connections_remote_id FOREIGN KEY (remote_id) REFERENCES peers (id),
    CONSTRAINT fk_connections_multi_address_id FOREIGN KEY (multi_address_id) REFERENCES multi_addresses (id),

    PRIMARY KEY (id)
);

CREATE TABLE provides_x_connections
(
    provide_id    INT NOT NULL,
    connection_id INT NOT NULL,

    CONSTRAINT fk_provides_x_connections_provide_id FOREIGN KEY (provide_id) REFERENCES provides (id),
    CONSTRAINT fk_provides_x_connections_connection_id FOREIGN KEY (connection_id) REFERENCES connections (id),

    PRIMARY KEY (provide_id, connection_id)
);

CREATE TABLE retrievals_x_connections
(
    retrieval_id  INT NOT NULL,
    connection_id INT NOT NULL,

    CONSTRAINT fk_retrievals_x_connections_retrieval_id FOREIGN KEY (retrieval_id) REFERENCES retrievals (id),
    CONSTRAINT fk_retrievals_x_connections_connection_id FOREIGN KEY (connection_id) REFERENCES connections (id),

    PRIMARY KEY (retrieval_id, connection_id)
);

COMMIT;
