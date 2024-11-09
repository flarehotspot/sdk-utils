CREATE TABLE IF NOT EXISTS wallet_transactions (
    id SERIAL PRIMARY KEY,
    wallet_id INT NOT NULL,
    amount DECIMAL(8, 2) NOT NULL,
    new_balance DECIMAL(8, 2) NOT NULL,
    description VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (wallet_id) REFERENCES wallets (id) ON DELETE CASCADE
);
