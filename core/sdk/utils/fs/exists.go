package fs

import (
	"errors"
	"os"
)

func Exists(p string) bool {
	if _, err := os.Stat(p); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
