package plugins

import (
	"fmt"
	"net/http"
	sdkhttp "sdk/api/http"

	resp "core/internal/web/response"
	"core/resources/views"

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
		self.ErrorPage(w, r, err, http.StatusInternalServerError)
		return
	}

	assets := self.api.Utl.GetAdminAssetsForPage(v)
	data := sdkhttp.AdminLayoutData{
		Layout: sdkhttp.LayoutData{
			Assets:      assets,
			PageContent: v.PageContent,
		},
	}

	w.Header().Set("Content-Type", "text/html")
	page := themeApi.AdminTheme.LayoutFactory(w, r, data)
	page.Render(r.Context(), w)
}

func (self *HttpResponse) PortalView(w http.ResponseWriter, r *http.Request, v sdkhttp.ViewPage) {
	_, themeApi, err := self.api.PluginsMgrApi.GetPortalTheme()
	if err != nil {
		self.ErrorPage(w, r, err, http.StatusInternalServerError)
		return
	}

	assets := self.api.Utl.GetPortalAssetsForPage(v)
	data := sdkhttp.PortalLayoutData{
		Layout: sdkhttp.LayoutData{
			Assets:      assets,
			PageContent: v.PageContent,
		},
	}

	w.Header().Set("Content-Type", "text/html")
	page := themeApi.PortalTheme.LayoutFactory(w, r, data)
	page.Render(r.Context(), w)
}

func (self *HttpResponse) View(w http.ResponseWriter, r *http.Request, v sdkhttp.ViewPage) {
	w.Header().Set("Content-Type", "text/html")
	v.PageContent.Render(r.Context(), w)
}

func (self *HttpResponse) File(w http.ResponseWriter, r *http.Request, file string, data interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}

	file = self.api.Utl.Resource(file)

	fmt.Fprintf(w, "TODO: respond with file download")
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

func (self *HttpResponse) ErrorPage(w http.ResponseWriter, r *http.Request, err error, status int) {
	w.WriteHeader(status)
	page := views.ErrorPage(err)
	v := sdkhttp.ViewPage{PageContent: page}
	_, autherr := self.api.HttpAPI.auth.CurrentAcct(r)
	if autherr != nil {
		self.api.HttpAPI.HttpResponse().PortalView(w, r, v)
	} else {
		self.api.HttpAPI.HttpResponse().AdminView(w, r, v)
	}
}
