package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	fs "github.com/flarehotspot/sdk/utils/fs"
	sdkfs "github.com/flarehotspot/sdk/utils/fs"
	paths "github.com/flarehotspot/sdk/utils/paths"
)

func NewPLuginConfig(pkg string) *PluginConfig {
	return &PluginConfig{pkg}
}

type PluginConfig struct {
	pkg string
}

func (c *PluginConfig) configPath() string {
	return filepath.Join(paths.ConfigDir, "plugins", c.pkg+".json")
}

func (c *PluginConfig) Get(v interface{}) error {
	bytes, err := os.ReadFile(c.configPath())
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, v)
}

func (c *PluginConfig) Save(v interface{}) error {
	dir := filepath.Join(paths.ConfigDir, "plugins")
	err := fs.EnsureDir(dir)
	if err != nil {
		return err
	}

	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(c.configPath(), b, sdkfs.PermFile)
	return err
}
