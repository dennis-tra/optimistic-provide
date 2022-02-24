BEGIN;

CREATE TYPE peer_state AS ENUM (
    'HEARD',
    'WAITING',
    'QUERIED',
    'UNREACHABLE'
    );

CREATE TABLE peer_states
(
    id           INT GENERATED ALWAYS AS IDENTITY,
    provide_id   INT,
    retrieval_id INT,
    peer_id      INT        NOT NULL,
    referrer_id  INT        NOT NULL,
    state        peer_state NOT NULL,
    distance     bytea      NOT NULL,

    CONSTRAINT fk_peer_states_provide_id FOREIGN KEY (provide_id) REFERENCES provides (id),
    CONSTRAINT fk_peer_states_retrieval_id FOREIGN KEY (retrieval_id) REFERENCES retrievals (id),
    CONSTRAINT fk_peer_states_peer_id FOREIGN KEY (peer_id) REFERENCES peers (id),
    CONSTRAINT fk_peer_states_referrer_id FOREIGN KEY (referrer_id) REFERENCES peers (id),

    CHECK ((provide_id IS NULL) != (retrieval_id IS NULL)),

    PRIMARY KEY (id)
);

COMMIT;
