package config

import (
	"os"
	"path/filepath"

	"github.com/goccy/go-json"

	sdkcfg "sdk/api/config"

	fs "github.com/flarehotspot/go-utils/fs"
	sdkfs "github.com/flarehotspot/go-utils/fs"
	paths "github.com/flarehotspot/go-utils/paths"
	sdkstr "github.com/flarehotspot/go-utils/strings"
)

func NewCustomConfig(key string, pkg string) *CustomConfig {
	return &CustomConfig{key, pkg}
}

type CustomConfig struct {
	key string
	pkg string
}

func (c *CustomConfig) configPath() string {
	return filepath.Join(paths.ConfigDir, "plugins", c.pkg, sdkstr.Slugify(c.key, "_")+".json")
}

func (c *CustomConfig) Get(v interface{}) error {
	cfgPath := c.configPath()

	if !sdkfs.Exists(cfgPath) {
		return sdkcfg.ErrNoConfig
	}

	bytes, err := os.ReadFile(cfgPath)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, v)
}

func (c *CustomConfig) Save(v interface{}) error {
	dir := filepath.Join(paths.ConfigDir, "plugins", c.pkg)
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
