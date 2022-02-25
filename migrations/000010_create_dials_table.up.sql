BEGIN;

CREATE TYPE dial_transport AS ENUM (
    'tcp',
    'ws',
    'quic',
    'udp'
    );

CREATE TABLE dials
(
    id               INT GENERATED ALWAYS AS IDENTITY,
    provide_id       INT,
    retrieval_id     INT,
    local_id         INT            NOT NULL,
    remote_id        INT            NOT NULL,
    transport        dial_transport NOT NULL,
    multi_address_id INT            NOT NULL,
    started_at       TIMESTAMPTZ    NOT NULL,
    ended_at         TIMESTAMPTZ    NOT NULL,
    error            TEXT,

    CONSTRAINT fk_dials_provide_id FOREIGN KEY (provide_id) REFERENCES provides (id),
    CONSTRAINT fk_dials_retrieval_id FOREIGN KEY (retrieval_id) REFERENCES retrievals (id),
    CONSTRAINT fk_dials_local_id FOREIGN KEY (local_id) REFERENCES peers (id),
    CONSTRAINT fk_dials_remote_id FOREIGN KEY (remote_id) REFERENCES peers (id),
    CONSTRAINT fk_dials_multi_address_id FOREIGN KEY (multi_address_id) REFERENCES multi_addresses (id),

    CHECK ((provide_id IS NULL) != (retrieval_id IS NULL)),

    PRIMARY KEY (id)
);

CREATE TABLE provides_x_dials
(
    provide_id INT NOT NULL,
    dial_id    INT NOT NULL,

    CONSTRAINT fk_provides_x_dials_provide_id FOREIGN KEY (provide_id) REFERENCES provides (id),
    CONSTRAINT fk_provides_x_dials_dial_id FOREIGN KEY (dial_id) REFERENCES dials (id),

    PRIMARY KEY (provide_id, dial_id)
);

CREATE TABLE retrievals_x_dials
(
    retrieval_id INT NOT NULL,
    dial_id      INT NOT NULL,

    CONSTRAINT fk_retrievals_x_dials_retrieval_id FOREIGN KEY (retrieval_id) REFERENCES retrievals (id),
    CONSTRAINT fk_retrievals_x_dials_dial_id FOREIGN KEY (dial_id) REFERENCES dials (id),

    PRIMARY KEY (retrieval_id, dial_id)
);


COMMIT;
