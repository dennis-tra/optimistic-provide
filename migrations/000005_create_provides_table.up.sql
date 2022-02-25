-- The `provides` table keeps track of all provide operations
CREATE TABLE provides
(
    -- A unique identifier for this provide operation
    id                       INT GENERATED ALWAYS AS IDENTITY,
    -- The peer ID of the provide-initiating peer
    provider_id              INT         NOT NULL,
    -- The content identifier being provided
    content_id               TEXT        NOT NULL,
    -- The XOR distance from peer ID to content ID
    distance                 bytea       NOT NULL,
    -- The state of the routing table when the provide operation was started
    initial_routing_table_id INT         NOT NULL,
    -- The state of the routing table when the provide operation ended
    final_routing_table_id   INT,
    -- Application level timestamp when this provide operation started
    started_at               TIMESTAMPTZ NOT NULL,
    -- Application level timestamp when this provide operation ended
    ended_at                 TIMESTAMPTZ,
    -- The returned error of the provide operation
    error                    TEXT,
    -- Application level timestamp when this provide plus all persistence operations have finished
    done_at                  TIMESTAMPTZ,

    -- database timestamps
    updated_at               TIMESTAMPTZ NOT NULL,
    created_at               TIMESTAMPTZ NOT NULL,

    CONSTRAINT fk_provides_provider_id FOREIGN KEY (provider_id) REFERENCES peers (id),

    PRIMARY KEY (id)
);
