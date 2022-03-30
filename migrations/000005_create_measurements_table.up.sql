BEGIN;

CREATE TABLE measurements
(
    id            INT GENERATED ALWAYS AS IDENTITY,
    host_id       INT         NOT NULL,
    started_at    TIMESTAMPTZ NOT NULL,
    ended_at      TIMESTAMPTZ,
    configuration JSON        NOT NULL,
    updated_at    TIMESTAMPTZ NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL,

    CONSTRAINT fk_measurements_host_id FOREIGN KEY (host_id) REFERENCES peers (id),

    PRIMARY KEY (id)
);

COMMIT;
