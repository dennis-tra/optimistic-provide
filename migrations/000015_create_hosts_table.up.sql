BEGIN;

CREATE TABLE hosts
(
    id          INT GENERATED ALWAYS AS IDENTITY,
    host_id     INT   NOT NULL,
    private_key bytea NOT NULL,

    CONSTRAINT fk_hosts_host_id FOREIGN KEY (host_id) REFERENCES peers (id),

    PRIMARY KEY (id)
);

COMMIT;
