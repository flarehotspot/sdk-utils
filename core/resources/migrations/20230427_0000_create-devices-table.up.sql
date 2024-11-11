CREATE  EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS devices (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    ip_address CHAR(15) NOT NULL,
    mac_address CHAR(17) NOT NULL,
    hostname VARCHAR(64),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
