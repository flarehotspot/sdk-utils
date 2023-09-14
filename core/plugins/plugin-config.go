package plugins

import (
	"github.com/flarehotspot/core/sdk/utils/fs"
	"github.com/flarehotspot/core/sdk/utils/paths"
	"os"
	"path/filepath"
)

type PluginConfig struct {
	api *PluginApi
}

func (c *PluginConfig) configPath() string {
	return filepath.Join(paths.ConfigDir, "plugins", c.api.Pkg()+".yml")
}

func (c *PluginConfig) Write(b []byte) error {
	dir := filepath.Join(paths.ConfigDir, "plugins")
	err := fs.EnsureDir(dir)
	if err != nil {
		return err
	}
	err = os.WriteFile(c.configPath(), b, 0644)
	return err
}

func (c *PluginConfig) Read() ([]byte, error) {
	bytes, err := os.ReadFile(c.configPath())
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func NewPLuginConfig(api *PluginApi) *PluginConfig {
	return &PluginConfig{api}
}
