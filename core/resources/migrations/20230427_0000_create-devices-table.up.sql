CREATE TABLE IF NOT EXISTS devices (
    id SERIAL PRIMARY KEY,
    ip_address CHAR(15) NOT NULL,
    mac_address CHAR(17) NOT NULL,
    hostname VARCHAR(64),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
