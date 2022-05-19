BEGIN;

CREATE TYPE network_type AS ENUM (
    'IPFS',
    'FILECOIN',
    'POLKADOT',
    'KUSAMA'
    );

ALTER TABLE hosts
    ADD COLUMN network network_type NOT NULL DEFAULT 'IPFS';

COMMIT;
