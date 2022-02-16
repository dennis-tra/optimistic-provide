BEGIN;

CREATE TABLE multi_addresses
(
    id               INT GENERATED ALWAYS AS IDENTITY,
    maddr            TEXT        NOT NULL,
    country          VARCHAR(2),
    continent        VARCHAR(2),
    asn              INT,
    is_public        BOOL,
    ip_address_count INT,

    updated_at       TIMESTAMPTZ NOT NULL,
    created_at       TIMESTAMPTZ NOT NULL,

    CONSTRAINT uq_multi_addresses_maddr UNIQUE (maddr),

    PRIMARY KEY (id)
);

CREATE OR REPLACE FUNCTION upsert_multi_address(
    new_multi_address TEXT,
    new_country VARCHAR(2),
    new_continent VARCHAR(2),
    new_asn INT,
    new_is_public BOOL,
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
        INSERT INTO multi_addresses (maddr, country, continent, asn, is_public, ip_address_count, updated_at,
                                     created_at)
        VALUES (new_multi_address, new_country, new_continent, new_asn, new_is_public,
                coalesce(array_length(new_ip_address_ids, 1), 0), NOW(), NOW())
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
