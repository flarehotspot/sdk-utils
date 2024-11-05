package plugins

import (
	"html/template"
	"path"

	sdkhttp "sdk/api/http"
	plugin "sdk/api/plugin"

	"github.com/gorilla/csrf"
)

func NewHttpHelpers(api *PluginApi) sdkhttp.HttpHelpers {
	return &HttpHelpers{api: api}
}

type HttpHelpers struct {
	api *PluginApi
}

func (self *HttpHelpers) CsrfHtmlTag() string {
	t := csrf.TemplateTag
	return t
}

func (self *HttpHelpers) Translate(msgtype string, msgk string, pairs ...interface{}) string {
	return self.api.Utl.Translate(msgtype, msgk, pairs...)
}

func (self *HttpHelpers) AssetPath(p string) string {
	return path.Join("/plugin", self.api.Pkg(), self.api.Version(), "assets/dist", p)
}

func (self *HttpHelpers) PluginMgr() plugin.PluginsMgrApi {
	return self.api.PluginsMgrApi
}

func (self *HttpHelpers) AdsView() (html template.HTML) {
	return ""
}

func (self *HttpHelpers) UrlForRoute(name string, pairs ...string) string {
	return self.api.HttpAPI.httpRouter.UrlForRoute(sdkhttp.PluginRouteName(name), pairs...)
}

func (self *HttpHelpers) UrlForPkgRoute(pkg string, name string, pairs ...string) string {
	return self.api.HttpAPI.httpRouter.UrlForPkgRoute(pkg, name, pairs...)
}
