BEGIN;

ALTER TABLE network_size_estimates
    DROP COLUMN cpl;
ALTER TABLE network_size_estimates
    DROP COLUMN distances;
ALTER TABLE network_size_estimates
    DROP COLUMN key;

COMMIT;
