package cfgapi

import (
	"github.com/flarehotspot/core/internal/config"
	"github.com/flarehotspot/core/sdk/api/config"
)

func NewDbCfgApi() *DbCfgApi {
	return &DbCfgApi{}
}

type DbCfgApi struct{}

func (c *DbCfgApi) Read() (sdkcfg.Database, error) {
	cfg, err := config.ReadDatabaseConfig()
	if err != nil {
		return sdkcfg.Database{}, err
	}

	return sdkcfg.Database{
		Host:     cfg.Host,
		Username: cfg.Username,
		Password: cfg.Password,
		Database: cfg.Database,
	}, nil
}

func (c *DbCfgApi) Write(cfg sdkcfg.Database) error {
	dbcfg := config.DbConfig{
		Host:     cfg.Host,
		Username: cfg.Username,
		Password: cfg.Password,
		Database: cfg.Database,
	}
	return config.WriteDatabaseConfig(dbcfg)
}
