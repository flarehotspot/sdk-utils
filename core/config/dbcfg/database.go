package dbcfg

import (
	"fmt"
	"github.com/flarehotspot/core/sdk/libs/yaml-3"
	"github.com/flarehotspot/core/sdk/utils/paths"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"path/filepath"
)

var configPath = filepath.Join(paths.ConfigDir, "database.yml")

type DbConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	SslMode  string `yaml:"sslmode"`
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

	err = yaml.Unmarshal(dbBytes, &cfg)
	if err != nil {
		return nil, err
	}

	if cfg.Host == "" {
		cfg.Host = "localhost"
	}

	return &cfg, nil
}

func Write(cfg *DbConfig) error {
	b, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, b, 0644)
}
