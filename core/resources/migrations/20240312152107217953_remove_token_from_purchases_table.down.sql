ALTER TABLE purchases ADD COLUMN token VARCHAR(255) NOT NULL;
CREATE UNIQUE INDEX IF NOT EXISTS index_purchase_token ON purchases(token);
