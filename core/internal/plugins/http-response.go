package plugins

import (
	"net/http"
	"path/filepath"

	"github.com/flarehotspot/core/internal/web/response"
	resp "github.com/flarehotspot/core/internal/web/response"
	paths "github.com/flarehotspot/sdk/utils/paths"
)

type HttpResponse struct {
	api      *PluginApi
	viewroot string
}

func NewHttpResponse(api *PluginApi) *HttpResponse {
	viewroot := paths.Strip(api.Utl.Resource("views"))
	return &HttpResponse{api, viewroot}
}

func (self *HttpResponse) AdminView(w http.ResponseWriter, r *http.Request, view string, data interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}

	helpers := NewHttpHelpers(self.api)
	viewsDir := self.api.Utl.Resource("views/admin")
	layoutFile := filepath.Join(viewsDir, "layout.html")
	viewFile := filepath.Join(viewsDir, view)
	resp.ViewWithLayout(w, layoutFile, viewFile, helpers, data)
}

func (self *HttpResponse) PortalView(w http.ResponseWriter, r *http.Request, view string, data interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}

	helpers := NewHttpHelpers(self.api)
	viewsDir := self.api.Utl.Resource("views/portal")
	layoutFile := filepath.Join(viewsDir, "layout.html")
	viewFile := filepath.Join(viewsDir, view)
	resp.ViewWithLayout(w, layoutFile, viewFile, helpers, data)
}

func (self *HttpResponse) View(w http.ResponseWriter, r *http.Request, view string, data interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}

	helpers := NewHttpHelpers(self.api)
	vdir := self.api.Utl.Resource("views")
	viewfile := filepath.Join(vdir, view)

	resp.View(w, viewfile, helpers, data)
}

func (self *HttpResponse) File(w http.ResponseWriter, r *http.Request, file string, data interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}

	helpers := NewHttpHelpers(self.api)
	file = self.api.Utl.Resource(file)

	response.File(w, file, helpers, data)
}

func (res *HttpResponse) Json(w http.ResponseWriter, data interface{}, status int) {
	resp.Json(w, data, status)
}
