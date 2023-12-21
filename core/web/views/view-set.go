//go:build !dev

package views

import (
	"github.com/flarehotspot/core/sdk/libs/jet"
	"github.com/flarehotspot/core/sdk/utils/paths"
)

var set = jet.NewSet(jet.NewOSFileSystemLoader(paths.AppDir))

func GetTemplate(tmpl string) (*jet.Template, error) {
	return set.GetTemplate(tmpl)
}
