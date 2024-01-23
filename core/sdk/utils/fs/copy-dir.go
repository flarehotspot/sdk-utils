package sdkfs

import (
	"os"
	"path/filepath"
)

type CopyOpts struct {
	Override  bool
	Recursive bool
}

func CopyDir(srcDir, destDir string, opts CopyOpts) error {
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		sourcePath := filepath.Join(srcDir, entry.Name())
		destPath := filepath.Join(destDir, entry.Name())

		if !opts.Override {
			if _, err := os.Stat(destPath); err == nil {
				return nil
			}
		}

		fileInfo, err := os.Stat(sourcePath)
		if err != nil {
			return err
		}

		dir := filepath.Dir(destPath)
		if err := EnsureDir(dir); err != nil {
			return err
		}

		if entry.IsDir() {
			if err := CopyDir(sourcePath, destPath, opts); err != nil {
				return err
			}
		} else if fileInfo.Mode()&os.ModeSymlink == os.ModeSymlink {
			if err := CopySymLink(sourcePath, destPath); err != nil {
				return err
			}
		} else {
			if err := CopyFile(sourcePath, destPath); err != nil {
				return err
			}
		}
	}
	return nil
}

