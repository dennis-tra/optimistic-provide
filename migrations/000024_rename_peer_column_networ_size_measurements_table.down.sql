BEGIN;

ALTER TABLE network_size_estimates RENAME COLUMN host_id TO peer_id;
ALTER TABLE network_size_estimates DROP CONSTRAINT fk_network_size_estimates_host_id;
ALTER TABLE network_size_estimates ADD CONSTRAINT fk_network_size_estimates_peer_id FOREIGN KEY (peer_id) REFERENCES peers (id);

COMMIT;
