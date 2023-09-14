package fs

import (
	"path/filepath"
	"strings"
)

// baseName Returns the filename without the file extension
func BaseName(f string) string {
	ext := filepath.Ext(f)
	return strings.Replace(filepath.Base(f), ext, "", 1)
}
