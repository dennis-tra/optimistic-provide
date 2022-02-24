BEGIN;

CREATE TABLE providers
(
    id                INT GENERATED ALWAYS AS IDENTITY,
    retrieval_id      INT         NOT NULL,
    remote_id         INT         NOT NULL,
    multi_address_ids INT[]       NOT NULL,
    found_at          TIMESTAMPTZ NOT NULL,

    CONSTRAINT fk_get_providers_retrieval_id FOREIGN KEY (retrieval_id) REFERENCES retrievals (id),
    CONSTRAINT fk_get_providers_remote_id FOREIGN KEY (remote_id) REFERENCES peers (id),

    PRIMARY KEY (id)
);

COMMIT;
