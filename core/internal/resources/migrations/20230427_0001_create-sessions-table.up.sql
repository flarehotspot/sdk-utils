CREATE TABLE IF NOT EXISTS sessions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    device_id INT NOT NULL,
    session_type TINYINT NOT NULL,
    time_secs INT DEFAULT 0,
    data_mbytes DECIMAL(18, 9) DEFAULT 0.0,
    consumption_secs INT DEFAULT 0,
    consumption_mb DECIMAL(18, 9) DEFAULT 0.0,
    started_at TIMESTAMP,
    exp_days INT DEFAULT NULL,
    down_mbits INT NOT NULL DEFAULT 0,
    up_mbits INT NOT NULL DEFAULT 0,
    use_global TINYINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (device_id) REFERENCES devices (id) ON DELETE CASCADE
);
