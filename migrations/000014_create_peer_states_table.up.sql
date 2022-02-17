BEGIN;

CREATE TYPE peer_state AS ENUM (
    'HEARD',
    'WAITING',
    'QUERIED',
    'UNREACHABLE'
    );

CREATE TABLE peer_states
(
    provide_id  INT        NOT NULL,
    peer_id     INT        NOT NULL,
    referrer_id INT        NOT NULL,
    state       peer_state NOT NULL,
    distance    bytea      NOT NULL,

    CONSTRAINT fk_dials_provide_id FOREIGN KEY (provide_id) REFERENCES provides (id),
    CONSTRAINT fk_dials_peer_id FOREIGN KEY (peer_id) REFERENCES peers (id),
    CONSTRAINT fk_dials_referrer_id FOREIGN KEY (referrer_id) REFERENCES peers (id),

    PRIMARY KEY (provide_id, peer_id)
);

COMMIT;
