CREATE TABLE spans
(
    id             INT GENERATED ALWAYS AS IDENTITY,

    local_peer_id  INT         NOT NULL,
    remote_peer_id INT         NOT NULL,
    measurement_id INT         NOT NULL,

    error          TEXT,

    started_at     TIMESTAMPTZ NOT NULL,
    ended_at       TIMESTAMPTZ NOT NULL,

    CONSTRAINT fk_spans_measurement_id FOREIGN KEY (measurement_id) REFERENCES measurements (id),
    CONSTRAINT fk_spans_local_peer_id FOREIGN KEY (local_peer_id) REFERENCES peers (id),
    CONSTRAINT fk_spans_remote_peer_id FOREIGN KEY (remote_peer_id) REFERENCES peers (id),

    PRIMARY KEY (id)
);
