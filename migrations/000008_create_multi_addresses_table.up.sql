BEGIN;

CREATE TABLE multi_addresses
(
    id         INT GENERATED ALWAYS AS IDENTITY,
    maddr      TEXT        NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,

    CONSTRAINT uq_multi_addresses_maddr UNIQUE (maddr),

    PRIMARY KEY (id)
);

CREATE OR REPLACE FUNCTION upsert_multi_address(
    new_multi_address TEXT,
    new_ip_address_ids INT[]
) RETURNS INT AS
$upsert_multi_address$
DECLARE
    multi_address_id INT;
    multi_address    multi_addresses%rowtype;
BEGIN
    SELECT *
    FROM multi_addresses ma
    WHERE ma.maddr = new_multi_address
    INTO multi_address;

    IF multi_address IS NULL THEN
        INSERT INTO multi_addresses (maddr, updated_at, created_at)
        VALUES (new_multi_address, NOW(), NOW())
        RETURNING id INTO multi_address_id;

        INSERT INTO multi_addresses_x_ip_addresses
        SELECT multi_address_id, unnest(new_ip_address_ids)
        ON CONFLICT DO NOTHING;

        RETURN multi_address_id;
    END IF;

    INSERT INTO multi_addresses_x_ip_addresses
    SELECT multi_address.id, unnest(new_ip_address_ids)
    ON CONFLICT DO NOTHING;

    RETURN multi_address.id;
END;
$upsert_multi_address$ LANGUAGE plpgsql;

COMMIT;
