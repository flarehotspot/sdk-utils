package utils

import "os"

func EmptyDir(dir string) error {
	if _, err := os.Stat(dir); err != nil && !os.IsNotExist(err) {
		return err
	}

    os.RemoveAll(dir)
	return os.MkdirAll(dir, 0755)
}
