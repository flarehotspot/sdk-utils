package views

import (
	"github.com/flarehotspot/core/sdk/api/http"
)

// should be exported to sdk
type ViewData struct {
	contentPath string
	contentData any
	helpers     http.IHelpers
}

func (vd *ViewData) ContentPath() string {
	return vd.contentPath
}

func (vd *ViewData) Helpers() http.IHelpers {
	return vd.helpers
}

func (vd *ViewData) Data() any {
	return vd.contentData
}
