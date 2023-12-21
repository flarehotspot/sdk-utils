//go:build dev

package views

import (
	"github.com/flarehotspot/core/sdk/libs/jet"
	"github.com/flarehotspot/core/sdk/utils/paths"
)

// Return new set everytime to avoid caching in dev
func GetTemplate(tmpl string) (*jet.Template, error) {
	set := jet.NewSet(jet.NewOSFileSystemLoader(paths.AppDir))
	return set.GetTemplate(tmpl)
}
