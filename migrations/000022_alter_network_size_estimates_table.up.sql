BEGIN;

ALTER TABLE network_size_estimates
    ADD COLUMN cpl INT NOT NULL DEFAULT -1;
ALTER TABLE network_size_estimates
    ADD COLUMN distances FLOAT[] NOT NULL DEFAULT '{}';
ALTER TABLE network_size_estimates
    ADD COLUMN key TEXT NOT NULL DEFAULT '';

COMMIT;
