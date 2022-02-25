CREATE TABLE routing_table_entries
(
    routing_table_snapshot_id         INT         NOT NULL,
    -- The peer that is in this routing table bucket
    peer_id                           INT         NOT NULL,
    -- The bucket that this peer resides in
    bucket                            SMALLINT    NOT NULL,
    -- LastUsefulAt is the time instant at which the peer was last "useful" to us.
    -- Please see the DHT docs for the definition of usefulness.
    last_useful_at                    TIMESTAMPTZ,
    -- LastSuccessfulOutboundQueryAt is the time instant at which we last got a successful query response from the peer.
    last_successful_outbound_query_at TIMESTAMPTZ NOT NULL,
    -- AddedAt is the time this peer was added to the routing table.
    added_at                          TIMESTAMPTZ NOT NULL,
    -- connected_since is the time a connection to this peer was established
    connected_since                   TIMESTAMPTZ,

    CONSTRAINT fk_routing_table_entries_routing_table_snapshot_id FOREIGN KEY (routing_table_snapshot_id) REFERENCES routing_table_snapshots (id),
    CONSTRAINT fk_routing_table_entries_peer_id FOREIGN KEY (peer_id) REFERENCES peers (id),

    PRIMARY KEY (routing_table_snapshot_id, peer_id)
);
