ALTER TABLE IF EXISTS accounts DROP CONSTRAINT accounts_owner_fkey;
ALTER TABLE IF EXISTS accounts DROP CONSTRAINT accounts_owner_currency_key;

DROP TABLE users;
