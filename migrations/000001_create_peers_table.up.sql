BEGIN;

-- The `peers` table keeps track of all peers ever seen
CREATE TABLE peers
(
    id            INT GENERATED ALWAYS AS IDENTITY,
    multi_hash    TEXT        NOT NULL,
    agent_version TEXT,
    protocols     TEXT[],
    updated_at    TIMESTAMPTZ NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL,

    UNIQUE (multi_hash),

    PRIMARY KEY (id)
);

CREATE OR REPLACE FUNCTION upsert_peer(
    new_multi_hash TEXT,
    new_agent_version TEXT,
    new_protocols TEXT[]
) RETURNS INT AS
$upsert_peer$
DECLARE
    peer_id INT;
    peer    peers%rowtype;
BEGIN
    SELECT *
    FROM peers p
    WHERE p.multi_hash = new_multi_hash
    INTO peer;

    IF peer IS NULL THEN
        INSERT INTO peers (multi_hash, agent_version, protocols, updated_at, created_at)
        VALUES (new_multi_hash, new_agent_version, new_protocols, NOW(), NOW())
        RETURNING id INTO peer_id;

        RETURN peer_id;
    END IF;

    IF new_agent_version IS NOT NULL AND peer.agent_version != new_agent_version THEN
        UPDATE peers
        SET agent_version = new_agent_version,
            updated_at    = NOW()
        WHERE id = peer.id;
    END IF;

    IF new_protocols IS NOT NULL AND array_length(new_protocols, 1) != 0 AND peer.protocols != new_protocols THEN
        UPDATE peers
        SET protocols  = new_protocols,
            updated_at = NOW()
        WHERE id = peer.id;
    END IF;

    RETURN peer.id;
END;
$upsert_peer$ LANGUAGE plpgsql;

COMMIT;
