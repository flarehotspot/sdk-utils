package plugins

import (
	"net/http"
	sdkhttp "sdk/api/http"

	"core/internal/web/response"
	resp "core/internal/web/response"

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

	assets := self.api.Utl.GetAdminAssetsForPage(v)
	data := sdkhttp.AdminLayoutData{
		Layout: sdkhttp.LayoutData{
			Assets:      assets,
			PageContent: v.PageContent,
		},
	}

	page := themeApi.AdminTheme.LayoutFactory(data)
	page.Render(r.Context(), w)
}

func (self *HttpResponse) PortalView(w http.ResponseWriter, r *http.Request, v sdkhttp.ViewPage) {
	_, themeApi, err := self.api.PluginsMgrApi.GetPortalTheme()
	if err != nil {
		self.ErrorPage(w, err)
		return
	}

	assets := self.api.Utl.GetPortalAssetsForPage(v)
	data := sdkhttp.PortalLayoutData{
		Layout: sdkhttp.LayoutData{
			Assets:      assets,
			PageContent: v.PageContent,
		},
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
