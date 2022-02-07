BEGIN;

CREATE TABLE peer_infos
(
    id               INT GENERATED ALWAYS AS IDENTITY,

    agent_version_id INT         NOT NULL,
    peer_id          INT         NOT NULL,
    measurement_id   INT         NOT NULL,


    created_at       TIMESTAMPTZ NOT NULL,

    PRIMARY KEY (id)
);

END;
