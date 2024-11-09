DROP INDEX IF EXISTS index_purchase_token;
ALTER TABLE purchases DROP COLUMN token;
