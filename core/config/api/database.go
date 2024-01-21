package cfgapi

import (
	"github.com/flarehotspot/core/config/dbcfg"
	"github.com/flarehotspot/core/sdk/api/config"
)

type DbCfgApi struct{}

func (c *DbCfgApi) Read() (*sdkcfg.Database, error) {
	cfg, err := dbcfg.Read()
	if err != nil {
		return nil, err
	}

	return &sdkcfg.Database{
		Host:     cfg.Host,
		Username: cfg.Username,
		Password: cfg.Password,
		Database: cfg.Database,
	}, nil
}

func (c *DbCfgApi) Write(cfg *sdkcfg.Database) error {
	return dbcfg.Write(&dbcfg.DbConfig{
		Host:     cfg.Host,
		Username: cfg.Username,
		Password: cfg.Password,
		Database: cfg.Database,
	})
}

func NewDbCfgApi() *DbCfgApi {
	return &DbCfgApi{}
}
