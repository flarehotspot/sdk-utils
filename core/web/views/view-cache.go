package views

import (
	"bytes"

	"github.com/flarehotspot/core/sdk/api/http/views"
	"github.com/flarehotspot/core/sdk/libs/jet"
	"github.com/flarehotspot/core/sdk/utils/paths"
)

var viewset = jet.NewSet(jet.NewOSFileSystemLoader(paths.AppDir))

// should be exported to sdk
type viewData struct {
	ContentPath string
	Data        any
	Helpers     views.IViewHelpers
}

type viewCache struct {
	hash    string
	layout  string
	content *string
	helpers views.IViewHelpers
	scripts []string
	styles  []string
}

func (vc *viewCache) RenderHTML(data any) (html string, err error) {
	vdata := &viewData{
		Helpers: vc.helpers,
		Data:    data,
	}

	if vc.content != nil {
		contpath := paths.RelativeFromTo(vc.layout, *vc.content)
		vdata.ContentPath = contpath
	}

	tmpl, err := viewset.GetTemplate(vc.layout)
	if err != nil {
		return "", err
	}

	var buff bytes.Buffer
	if err := tmpl.Execute(&buff, nil, vdata); err != nil {
		return "", err
	}

    html = buff.String()
    assets := []AssetBundle{}
}
