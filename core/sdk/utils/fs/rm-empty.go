package fs

import (
	"fmt"
	"os"
	"path/filepath"
)

func RmEmpty(dirPath string) error {
	emptyDirs := make([]string, 0)
	err := findEmptyDirs(dirPath, &emptyDirs)
	if err != nil {
		return err
	}

	// Remove empty directories.
	for _, dir := range emptyDirs {
		removeErr := os.Remove(dir)
		if removeErr != nil {
			fmt.Println("Error removing directory:", removeErr)
		}

		// Remove empty parent directories.
		parentDir := filepath.Dir(dir)
		if isEmpty, err := isEmptyDir(parentDir); err == nil && isEmpty {
			removeErr := os.Remove(parentDir)
			if removeErr != nil {
				fmt.Println("Error removing directory:", removeErr)
			}
		}
	}

	return nil
}

func findEmptyDirs(dirPath string, emptyDirs *[]string) error {
	dir, err := os.Open(dirPath)
	if err != nil {
		return err
	}
	defer dir.Close()

	entries, err := dir.Readdir(-1)
	if err != nil {
		return err
	}

	if len(entries) == 0 {
		*emptyDirs = append(*emptyDirs, dirPath)
		return nil
	}

	for _, entry := range entries {
		if entry.IsDir() {
			subDirPath := filepath.Join(dirPath, entry.Name())
			err := findEmptyDirs(subDirPath, emptyDirs)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func isEmptyDir(dirPath string) (bool, error) {
	dir, err := os.Open(dirPath)
	if err != nil {
		return false, err
	}
	defer dir.Close()

	entries, err := dir.Readdir(-1)
	if err != nil {
		return false, err
	}

	return len(entries) == 0, nil
}
