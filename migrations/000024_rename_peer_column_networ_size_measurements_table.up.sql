BEGIN;

ALTER TABLE network_size_estimates RENAME COLUMN peer_id TO host_id;
ALTER TABLE network_size_estimates DROP CONSTRAINT fk_network_size_estimates_peer_id;
ALTER TABLE network_size_estimates ADD CONSTRAINT fk_network_size_estimates_host_id FOREIGN KEY (host_id) REFERENCES hosts (id);

COMMIT;
