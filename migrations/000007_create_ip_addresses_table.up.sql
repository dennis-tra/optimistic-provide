BEGIN;

CREATE TABLE ip_addresses
(
    id         INT GENERATED ALWAYS AS IDENTITY,
    address    INET        NOT NULL,
    country    VARCHAR(2),
    continent  VARCHAR(2),
    asn        INT,

    updated_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,

    CONSTRAINT uq_ip_addresses_address UNIQUE (address),

    PRIMARY KEY (id)
);

CREATE OR REPLACE FUNCTION upsert_ip_address(
    new_address INET,
    new_country VARCHAR(2),
    new_continent VARCHAR(2),
    new_asn INT
) RETURNS INT AS
$upsert_ip_address$
DECLARE
    ip_address_id INT;
    ip_address    ip_addresses%rowtype;
BEGIN
    SELECT *
    FROM ip_addresses ia
    WHERE ia.address = new_address
    INTO ip_address;

    IF ip_address IS NULL THEN
        INSERT INTO ip_addresses (address, country, continent, asn, updated_at, created_at)
        VALUES (new_address, new_country, new_continent, new_asn, NOW(), NOW())
        RETURNING id INTO ip_address_id;

        RETURN ip_address_id;
    END IF;

    IF new_country IS NOT NULL AND ip_address.country != new_country THEN
        UPDATE ip_addresses
        SET country    = new_country,
            updated_at = NOW()
        WHERE id = ip_address.id;
    END IF;

    IF new_continent IS NOT NULL AND ip_address.continent != new_continent THEN
        UPDATE ip_addresses
        SET continent  = new_continent,
            updated_at = NOW()
        WHERE id = ip_address.id;
    END IF;

    IF new_asn IS NOT NULL AND ip_address.asn != new_asn THEN
        UPDATE ip_addresses
        SET asn  = new_asn,
            updated_at = NOW()
        WHERE id = ip_address.id;
    END IF;

    RETURN ip_address.id;
END;
$upsert_ip_address$ LANGUAGE plpgsql;

COMMIT;

