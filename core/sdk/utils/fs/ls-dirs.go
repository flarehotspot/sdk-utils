package fs

import (
	"io/fs"
	"io/ioutil"
	"path/filepath"
)

// LsDirs returns directories inside dir. Directory paths are prepended with dir
func LsDirs(dir string, recursive bool) ([]string, error) {
	darr := []string{}

	if !recursive {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			return darr, err
		}
		for _, f := range files {
			if f.IsDir() {
				darr = append(darr, filepath.Join(dir, f.Name()))
			}
		}
		return darr, nil
	}

	err := filepath.WalkDir(dir, func(s string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			darr = append(darr, s)
		}
		return nil
	})

  if err != nil {
    return darr, err
  }

	return darr, nil
}
