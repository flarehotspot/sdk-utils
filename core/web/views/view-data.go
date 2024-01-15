package views

import "github.com/flarehotspot/core/sdk/api/http/views"

// should be exported to sdk
type ViewData struct {
	contentPath string
	contentData any
	helpers     views.IViewHelpers
}

func (vd *ViewData) ContentPath() string {
	return vd.contentPath
}

func (vd *ViewData) Helpers() views.IViewHelpers {
	return vd.helpers
}

func (vd *ViewData) Data() any {
	return vd.contentData
}
