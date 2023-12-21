package views

import "github.com/flarehotspot/core/sdk/api/http/views"

// should be exported to sdk
type viewData struct {
	contentPath string
	contentData any
	helpers     views.IViewHelpers
}

func (vd *viewData) ContentPath() string {
	return vd.contentPath
}

func (vd *viewData) Helpers() views.IViewHelpers {
	return vd.helpers
}

func (vd *viewData) Data() any {
	return vd.contentData
}
