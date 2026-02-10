ALTER TABLE contacts
DROP CONSTRAINT IF EXISTS contacts_username_fkey;

DROP INDEX IF EXISTS uniq_contacts_email_active;

DROP TABLE IF EXISTS contacts;