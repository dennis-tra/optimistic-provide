BEGIN;

CREATE TYPE measurement_state AS ENUM (
    'started',
    'completed'
    );

CREATE TABLE measurements
(
    id         INT GENERATED ALWAYS AS IDENTITY,

    state      measurement_state NOT NULL,
    cid        TEXT              NOT NULL,
    started_at TIMESTAMPTZ       NOT NULL,
    ended_at   TIMESTAMPTZ,

    updated_at TIMESTAMPTZ       NOT NULL,
    created_at TIMESTAMPTZ       NOT NULL,

    PRIMARY KEY (id)
);

END;
