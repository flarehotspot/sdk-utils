package views

import (
	"html/template"

	"github.com/flarehotspot/core/sdk/api/http"
)

// should be exported to sdk
type ViewData struct {
	contentHtml template.HTML
	contentData any
	helpers     http.IHelpers
}

func (vd *ViewData) ContentHtml() template.HTML {
	return vd.contentHtml
}

func (vd *ViewData) Helpers() http.IHelpers {
	return vd.helpers
}

func (vd *ViewData) Data() any {
	return vd.contentData
}
