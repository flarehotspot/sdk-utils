//go:build !dev

package plugins

import (
	"io"
)

func OverrideLocalVersion(w io.Writer, pkg string) (ok bool) {
	// always return false if not in dev mode
	return false
}
