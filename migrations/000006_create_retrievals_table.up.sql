-- The `retrievals` table keeps track of all retrieval operations
CREATE TABLE retrievals
(
    -- A unique identifier for this retrieval operation
    id                       INT GENERATED ALWAYS AS IDENTITY,
    -- The peer ID of the retrieval-initiating peer
    retriever_id              INT         NOT NULL,
    -- The content identifier being retrieved
    content_id               TEXT        NOT NULL,
    -- The XOR distance from retriever ID to content ID
    distance                 bytea       NOT NULL,
    -- The state of the routing table when the retrieval operation was started
    initial_routing_table_id INT         NOT NULL,
    -- The state of the routing table when the retrieval operation ended
    final_routing_table_id   INT,
    -- Application level timestamp when this retrieval operation started
    started_at               TIMESTAMPTZ NOT NULL,
    -- Application level timestamp when this retrieval operation ended
    ended_at                 TIMESTAMPTZ,
    -- The returned error of the retrieval operation
    error                    TEXT,
    -- Application level timestamp when this retrieval plus all persistence operations have finished
    done_at                  TIMESTAMPTZ,

    -- database timestamps
    updated_at               TIMESTAMPTZ NOT NULL,
    created_at               TIMESTAMPTZ NOT NULL,

    CONSTRAINT fk_retrievals_retriever_id FOREIGN KEY (retriever_id) REFERENCES peers (id),
    CONSTRAINT fk_retrievals_initial_routing_table_id FOREIGN KEY (initial_routing_table_id) REFERENCES routing_table_snapshots (id),
    CONSTRAINT fk_retrievals_final_routing_table_id FOREIGN KEY (final_routing_table_id) REFERENCES routing_table_snapshots (id),

    PRIMARY KEY (id)
);
