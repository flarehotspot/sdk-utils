package dbcfg

import (
	"encoding/json"
	"fmt"
	paths "github.com/flarehotspot/core/sdk/utils/paths"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"path/filepath"
)

var configPath = filepath.Join(paths.ConfigDir, "database.json")

type DbConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
	SslMode  string `json:"sslmode"`
}

func (cfg *DbConfig) UrlString() string {
	return fmt.Sprintf("%s%s?parseTime=true", cfg.BaseConnStr(), cfg.Database)
}

func (cfg *DbConfig) BaseConnStr() string {
	var password string
	if cfg.Password != "" {
		password = ":" + cfg.Password
	} else {
		password = ""
	}

	var port string
	if cfg.Port != 0 {
		port = fmt.Sprintf(":%d", cfg.Port)
	} else {
		port = ""
	}

	return fmt.Sprintf("%s%s@tcp(%s%s)/", cfg.Username, password, cfg.Host, port)
}

func Read() (*DbConfig, error) {
	var cfg DbConfig
	dbBytes, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(dbBytes, &cfg)
	if err != nil {
		return nil, err
	}

	if cfg.Host == "" {
		cfg.Host = "localhost"
	}

	return &cfg, nil
}

func Write(cfg *DbConfig) error {
	b, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, b, 0644)
}
