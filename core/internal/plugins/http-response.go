package plugins

import (
	"net/http"
	sdkhttp "sdk/api/http"

	"core/internal/utils/assets"
	"core/internal/web/response"
	resp "core/internal/web/response"

	sdkfs "github.com/flarehotspot/go-utils/fs"
	paths "github.com/flarehotspot/go-utils/paths"
)

type HttpResponse struct {
	api      *PluginApi
	viewroot string
}

func NewHttpResponse(api *PluginApi) *HttpResponse {
	viewroot := paths.StripRoot(api.Utl.Resource("views"))
	return &HttpResponse{api, viewroot}
}

func (self *HttpResponse) AdminView(w http.ResponseWriter, r *http.Request, v sdkhttp.ViewPage) {
	_, themeApi, err := self.api.PluginsMgrApi.GetAdminTheme()
	if err != nil {
		self.ErrorPage(w, err)
		return
	}

	data := sdkhttp.AdminLayoutData{
		Layout: sdkhttp.LayoutData{
			PageContent: v.PageContent,
		},
	}

	page := themeApi.AdminTheme.LayoutFactory(data)
	page.Render(r.Context(), w)
}

func (self *HttpResponse) PortalView(w http.ResponseWriter, r *http.Request, v sdkhttp.ViewPage) {
	themePlugin, themeApi, err := self.api.PluginsMgrApi.GetPortalTheme()
	if err != nil {
		self.ErrorPage(w, err)
		return
	}

	data := sdkhttp.PortalLayoutData{
		Layout: sdkhttp.LayoutData{
			PageContent: v.PageContent,
		},
	}

	var scripts []string
	for _, f := range themeApi.PortalTheme.GlobalScripts {
		jspath := themePlugin.Resource(f)
		if sdkfs.IsDir(jspath) {
			var files []string
			err := sdkfs.LsFiles(jspath, &files, true)
			if err != nil {
				self.ErrorPage(w, err)
				return
			}
			scripts = append(scripts, files...)
		} else {
			scripts = append(scripts, jspath)
		}
	}

	for _, f := range v.Assets.Scripts {
		jspath := self.api.Resource(f)
		if sdkfs.IsDir(jspath) {
			var files []string
			err := sdkfs.LsFiles(jspath, &files, true)
			if err != nil {
				self.ErrorPage(w, err)
				return
			}
			scripts = append(scripts, files...)
		} else {
			scripts = append(scripts, jspath)
		}
	}

	var styles []string
	for _, f := range themeApi.PortalTheme.GlobalStylesheets {
		csspath := themePlugin.Resource(f)
		if sdkfs.IsDir(csspath) {
			var files []string
			err := sdkfs.LsFiles(csspath, &files, true)
			if err != nil {
				self.ErrorPage(w, err)
				return
			}
			styles = append(styles, files...)
		} else {
			styles = append(styles, csspath)
		}
	}

	for _, f := range v.Assets.Stylesheets {
		csspath := self.api.Resource(f)
		if sdkfs.IsDir(csspath) {
			var files []string
			err := sdkfs.LsFiles(csspath, &files, true)
			if err != nil {
				self.ErrorPage(w, err)
				return
			}
			styles = append(styles, files...)
		} else {
			styles = append(styles, csspath)
		}
	}

	if len(scripts) > 0 {
		jsdata, err := assets.Bundle(scripts)
		if err != nil {
			self.ErrorPage(w, err)
			return
		}
		data.Layout.GlobalScriptSrc = jsdata.PublicPath
	}

	if len(styles) > 0 {
		cssdata, err := assets.Bundle(styles)
		if err != nil {
			self.ErrorPage(w, err)
			return
		}
		data.Layout.GlobalStylesheetSrc = cssdata.PublicPath
	}

	page := themeApi.PortalTheme.LayoutFactory(data)
	page.Render(r.Context(), w)
}

func (self *HttpResponse) View(w http.ResponseWriter, r *http.Request, v sdkhttp.ViewPage) {
	v.PageContent.Render(r.Context(), w)
}

func (self *HttpResponse) File(w http.ResponseWriter, r *http.Request, file string, data interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}

	helpers := NewHttpHelpers(self.api)
	file = self.api.Utl.Resource(file)

	response.File(w, file, helpers, data)
}

func (self *HttpResponse) Json(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	resp.Json(w, data, status)
}

func (self *HttpResponse) FlashMsg(w http.ResponseWriter, r *http.Request, msg string, t string) {

}

func (self *HttpResponse) Redirect(w http.ResponseWriter, r *http.Request, routeName string, pairs ...string) {
	url := self.api.HttpAPI.Helpers().UrlForRoute(routeName, pairs...)
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func (self *HttpResponse) ErrorPage(w http.ResponseWriter, err error) {
	// TODO: show error page
	w.WriteHeader(500)
	w.Write([]byte(err.Error()))
}
