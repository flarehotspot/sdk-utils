package plugins

import (
	"net/http"
	"path/filepath"

	"github.com/flarehotspot/core/sdk/api/http/response"
	"github.com/flarehotspot/core/sdk/api/http/router"
	"github.com/flarehotspot/core/sdk/utils/flash"
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

func (self *HttpResponse) SetFlashMsg(w http.ResponseWriter, t flash.FlashType, msg string) {
	flash.SetFlashMsg(w, t, msg)
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

func (self *HttpResponse) Text(w http.ResponseWriter, r *http.Request, file string, data any) {
	if data == nil {
		data = map[string]any{}
	}

	helpers := NewViewHelpers(self.api, w, r)
	file = self.api.Resource(file)
	resp.Text(w, file, helpers, data)
}

func (res *HttpResponse) Json(w http.ResponseWriter, data any, status int) {
	resp.Json(w, data, status)
}

func (res *HttpResponse) NewErrRoute(route router.PluginRouteName, pairs ...string) response.IErrorRedirect {
	muxroute := res.api.HttpApi().HttpRouter().MuxRouteName(route)
	return resp.NewErrRoute(muxroute, pairs...)
}

func (res *HttpResponse) NewErrUrl(url string) response.IErrorRedirect {
	return resp.NewErrUrl(url)
}
