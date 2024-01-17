package plugins

import (
	"net/http"
	"path/filepath"

	"github.com/flarehotspot/core/sdk/utils/paths"
	resp "github.com/flarehotspot/core/web/response"
)

type HttpResponse struct {
	api      *PluginApi
	viewroot string
}

func NewHttpResponse(api *PluginApi) *HttpResponse {
	viewroot := paths.Strip(api.Resource("views"))
	return &HttpResponse{api, viewroot}
}

func (self *HttpResponse) AdminView(w http.ResponseWriter, r *http.Request, view string, data any) {
	if data == nil {
		data = map[string]any{}
	}

	helpers := NewViewHelpers(self.api, w, r)
	viewsDir := self.api.Resource("views/web-admin")
	layoutFile := filepath.Join(viewsDir, "captive-portal/layout.html")
	viewFile := filepath.Join(viewsDir, view)
	resp.ViewWithLayout(w, &layoutFile, viewFile, helpers, data)
}

func (self *HttpResponse) PortalView(w http.ResponseWriter, r *http.Request, view string, data any) {
	if data == nil {
		data = map[string]any{}
	}

	helpers := NewViewHelpers(self.api, w, r)
	viewsDir := self.api.Resource("views/captive-portal")
	layoutFile := filepath.Join(viewsDir, "captive-portal/layout.html")
	viewFile := filepath.Join(viewsDir, view)
	resp.ViewWithLayout(w, &layoutFile, viewFile, helpers, data)
}

func (self *HttpResponse) View(w http.ResponseWriter, r *http.Request, view string, data any) {
	if data == nil {
		data = map[string]any{}
	}

	helpers := NewViewHelpers(self.api, w, r)
	vdir := self.api.Resource("views")
	viewfile := filepath.Join(vdir, view)

	resp.View(w, viewfile, helpers, data)
}

func (self *HttpResponse) Script(w http.ResponseWriter, r *http.Request, file string, data any) {
	if data == nil {
		data = map[string]any{}
	}

	helpers := NewViewHelpers(self.api, w, r)
	file = self.api.Resource(file)

	w.Header().Set("Content-Type", "text/javascript")
	resp.Text(w, file, helpers, data)
}

func (res *HttpResponse) Json(w http.ResponseWriter, data any, status int) {
	resp.Json(w, data, status)
}
