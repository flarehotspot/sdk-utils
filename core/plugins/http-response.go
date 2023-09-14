package plugins

import (
	"net/http"
	"path/filepath"

	"github.com/flarehotspot/core/themes"
	resp "github.com/flarehotspot/core/web/response"
	v "github.com/flarehotspot/core/web/views"
	"github.com/flarehotspot/core/sdk/api/http/response"
	"github.com/flarehotspot/core/sdk/api/http/router"
	"github.com/flarehotspot/core/sdk/api/http/views"
	"github.com/flarehotspot/core/sdk/utils/flash"
)

type HttpResponse struct {
	api *PluginApi
}

func NewHttpResponse(api *PluginApi) *HttpResponse {
	return &HttpResponse{api}
}

func (self *HttpResponse) SetFlashMsg(w http.ResponseWriter, t flash.FlashType, msg string) {
	flash.SetFlashMsg(w, t, msg)
}

func (self *HttpResponse) AdminView(w http.ResponseWriter, r *http.Request, view string, data any) {
	if data == nil {
		data = map[string]interface{}{}
	}
	helpers := NewViewHelpers(self.api, w, r)
	vdir := self.api.Resource("views/web-admin")
	viewfile := filepath.Join(vdir, view)
	layout := themes.WebAdminLayout()
	fmap := self.api.vfmap
	vdata := &views.ViewData{Helpers: helpers, Data: data}
	resp.ViewWithLayout(w, layout, viewfile, fmap, vdata)
}

func (self *HttpResponse) PortalView(w http.ResponseWriter, r *http.Request, view string, data any) {
	if data == nil {
		data = map[string]interface{}{}
	}
	helpers := NewViewHelpers(self.api, w, r)
	vdir := self.api.Resource("views/captive-portal")
	viewfile := filepath.Join(vdir, view)
	layout := themes.PortalLayout()

	fmap := self.api.vfmap
	vdata := &views.ViewData{Helpers: helpers, Data: data}
	resp.ViewWithLayout(w, layout, viewfile, fmap, vdata)
}

func (self *HttpResponse) View(w http.ResponseWriter, r *http.Request, view string, data any) {
	if data == nil {
		data = map[string]interface{}{}
	}
	helpers := NewViewHelpers(self.api, w, r)
	vdir := self.api.Resource("views")
	viewfile := filepath.Join(vdir, view)

	fmap := self.api.vfmap
	vdata := &views.ViewData{Helpers: helpers, Data: data}
	v := &v.ViewInput{File: viewfile}
	resp.View(w, v, fmap, vdata)
}

func (res *HttpResponse) Json(w http.ResponseWriter, data any, status int) {
	resp.Json(w, data, status)
}

func (res *HttpResponse) NewErrRoute(route router.PluginRouteName, pairs ...string) response.IErrorRedirect {
	muxroute := res.api.HttpApi().Router().MuxRouteName(route)
	return resp.NewErrRoute(muxroute, pairs...)
}

func (res *HttpResponse) NewErrUrl(url string) response.IErrorRedirect {
	return resp.NewErrUrl(url)
}
