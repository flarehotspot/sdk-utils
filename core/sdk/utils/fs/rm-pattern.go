package fs

import (
	"os"
	"path/filepath"
)

func RmPattern(dirPath string, globPattern string) error {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			matched, matchErr := filepath.Match(globPattern, info.Name())
			if matchErr != nil {
				return matchErr
			}

			if matched {
				removeErr := os.Remove(path)
				if removeErr != nil {
					return removeErr
				}
			}
		}
		return nil
	})

	return err
}

