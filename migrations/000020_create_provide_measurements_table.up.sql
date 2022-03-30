BEGIN;

CREATE TABLE provide_measurements
(
    id           INT GENERATED ALWAYS AS IDENTITY,
    host_id      INT          NOT NULL,
    started_at   TIMESTAMPTZ  NOT NULL,
    ended_at     TIMESTAMPTZ,
    provide_type provide_type NOT NULL,
    iterations   INT          NOT NULL,
    updated_at   TIMESTAMPTZ  NOT NULL,
    created_at   TIMESTAMPTZ  NOT NULL,

    CONSTRAINT fk_provide_measurements_host_id FOREIGN KEY (host_id) REFERENCES peers (id),

    PRIMARY KEY (id)
);

COMMIT;
