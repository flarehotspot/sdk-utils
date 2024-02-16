package plugins

import (
	"encoding/json"
	"os"
	"path/filepath"

	fs "github.com/flarehotspot/flarehotspot/core/sdk/utils/fs"
	paths "github.com/flarehotspot/flarehotspot/core/sdk/utils/paths"
)

func NewPLuginConfig(api *PluginApi) *PluginConfig {
	return &PluginConfig{api}
}

type PluginConfig struct {
	api *PluginApi
}

func (c *PluginConfig) configPath() string {
	return filepath.Join(paths.ConfigDir, "plugins", c.api.Pkg()+".json")
}

func (c *PluginConfig) WriteJson(v any) error {
	dir := filepath.Join(paths.ConfigDir, "plugins")
	err := fs.EnsureDir(dir)
	if err != nil {
		return err
	}
	b, err := json.Marshal(v)
	err = os.WriteFile(c.configPath(), b, 0644)
	return err
}

func (c *PluginConfig) ReadJson(v any) error {
	bytes, err := os.ReadFile(c.configPath())
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, v)
}
