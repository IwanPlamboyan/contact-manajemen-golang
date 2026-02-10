ALTER TABLE addresses
DROP CONSTRAINT IF EXISTS addresses_contact_id_fkey;

DROP TABLE IF EXISTS addresses;