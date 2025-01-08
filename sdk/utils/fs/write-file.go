package sdkfs

import (
	"os"
	"path/filepath"
)

func WriteFile(path string, data []byte) error {
	if err := EnsureDir(filepath.Dir(path)); err != nil {
		return err
	}

	if err := os.WriteFile(path, data, PermFile); err != nil {
		return err
	}

	return nil
}
