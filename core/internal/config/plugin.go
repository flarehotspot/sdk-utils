package config

import (
	"os"
	"path/filepath"

	sdkutils "github.com/flarehotspot/sdk-utils"
)

func NewPluginCfgApi(pkg string) *PluginCfgApi {
	dirPath := filepath.Join(sdkutils.PathConfigDir, "plugins", pkg)
	return &PluginCfgApi{
		Pkg:     pkg,
		DirPath: dirPath,
	}
}

type PluginCfgApi struct {
	Pkg     string
	DirPath string
}

func (p *PluginCfgApi) Read(key string) ([]byte, error) {
	file := filepath.Join(p.DirPath, key)
	return os.ReadFile(file)
}

func (p *PluginCfgApi) Write(key string, data []byte) error {
	file := filepath.Join(p.DirPath, key)
	return sdkutils.FsWriteFile(file, data)
}
