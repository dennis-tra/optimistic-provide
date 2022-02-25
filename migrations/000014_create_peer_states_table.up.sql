BEGIN;

CREATE TYPE peer_state AS ENUM (
    'HEARD',
    'WAITING',
    'QUERIED',
    'UNREACHABLE'
    );

CREATE TABLE peer_states
(
    id          INT GENERATED ALWAYS AS IDENTITY,
    peer_id     INT        NOT NULL,
    referrer_id INT        NOT NULL,
    state       peer_state NOT NULL,
    distance    bytea      NOT NULL,

    CONSTRAINT fk_peer_states_peer_id FOREIGN KEY (peer_id) REFERENCES peers (id),
    CONSTRAINT fk_peer_states_referrer_id FOREIGN KEY (referrer_id) REFERENCES peers (id),

    PRIMARY KEY (id)
);

CREATE TABLE provides_x_peer_states
(
    provide_id    INT NOT NULL,
    peer_state_id INT NOT NULL,

    CONSTRAINT fk_provides_x_peer_states_provide_id FOREIGN KEY (provide_id) REFERENCES provides (id),
    CONSTRAINT fk_provides_x_peer_states_peer_state_id FOREIGN KEY (peer_state_id) REFERENCES peer_states (id),

    PRIMARY KEY (provide_id, peer_state_id)
);

CREATE TABLE retrievals_x_peer_states
(
    retrieval_id  INT NOT NULL,
    peer_state_id INT NOT NULL,

    CONSTRAINT fk_retrievals_x_peer_states_retrieval_id FOREIGN KEY (retrieval_id) REFERENCES retrievals (id),
    CONSTRAINT fk_retrievals_x_peer_states_peer_state_id FOREIGN KEY (peer_state_id) REFERENCES peer_states (id),

    PRIMARY KEY (retrieval_id, peer_state_id)
);

COMMIT;
