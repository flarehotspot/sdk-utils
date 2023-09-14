package fs

import (
	"io/fs"
	"os"
	"path/filepath"
)

// LsFiles returns list if files within dir. File paths are prepended with dir.
func LsFiles(dir string, recursive bool) ([]string, error) {
	farr := []string{}

	if !recursive {
		files, err := os.ReadDir(dir)
		if err != nil {
			return farr, err
		}
		for _, f := range files {
			if !f.IsDir() {
				farr = append(farr, filepath.Join(dir, f.Name()))
			}
		}
		return farr, nil
	}

	err := filepath.WalkDir(dir, func(s string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			farr = append(farr, s)
		}
		return nil
	})

	if err != nil {
		return farr, err
	}

	return farr, nil
}
