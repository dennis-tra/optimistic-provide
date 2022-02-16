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
    provide_id       INT            NOT NULL,
    local_id         INT            NOT NULL,
    remote_id        INT            NOT NULL,
    transport        dial_transport NOT NULL,
    multi_address_id INT            NOT NULL,
    started_at       TIMESTAMPTZ    NOT NULL,
    ended_at         TIMESTAMPTZ    NOT NULL,
    error            TEXT,

    CONSTRAINT fk_dials_provide_id FOREIGN KEY (provide_id) REFERENCES provides (id),
    CONSTRAINT fk_dials_local_id FOREIGN KEY (local_id) REFERENCES peers (id),
    CONSTRAINT fk_dials_remote_id FOREIGN KEY (remote_id) REFERENCES peers (id),
    CONSTRAINT fk_dials_multi_address_id FOREIGN KEY (multi_address_id) REFERENCES multi_addresses (id),

    PRIMARY KEY (id)
);

COMMIT;
