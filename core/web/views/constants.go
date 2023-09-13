package views

import (
	"path/filepath"

	"github.com/flarehotspot/core/sdk/utils/paths"
)

const (
	PublicPrefix  = "/public"
	TagTypeScript = "script"
	TagTypeStyle  = "style"
)

var (
	LayoutsDir = filepath.Join(paths.CoreDir, "web/views/layouts")
)
