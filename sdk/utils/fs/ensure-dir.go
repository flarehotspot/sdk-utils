package sdkfs

import (
	"os"
)

func EnsureDir(dir string) error {
	if !Exists(dir) {
		err := os.MkdirAll(dir, PermDir)
		if err != nil {
			return err
		}
	}
	return nil
}
