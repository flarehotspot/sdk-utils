package sdkfs

import "os"

func EnsureDir(dir string) error {
	if !Exists(dir) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
