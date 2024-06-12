package utils

import (
	"launcher/config"
	"os"
	"path/filepath"
)

func CoreExists() bool {
	stat, err := os.Stat(filepath.Join(config.AppPath, "core/plugin.so"))
	if err != nil {
		return false
	}

	return !stat.IsDir()
}
