BEGIN;

ALTER TABLE hosts
    DROP COLUMN network;

DROP TYPE network_type;

COMMIT;
