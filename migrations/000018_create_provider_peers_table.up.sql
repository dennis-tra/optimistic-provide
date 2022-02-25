BEGIN;

CREATE TABLE provider_peers
(
    id                INT GENERATED ALWAYS AS IDENTITY,
    get_providers_id  INT   NOT NULL,
    provider_id       INT   NOT NULL,
    multi_address_ids INT[] NOT NULL,

    CONSTRAINT fk_get_providers_get_providers_id FOREIGN KEY (get_providers_id) REFERENCES get_providers (id),
    CONSTRAINT fk_get_providers_provider_id FOREIGN KEY (provider_id) REFERENCES peers (id),

    PRIMARY KEY (id)
);

COMMIT;
