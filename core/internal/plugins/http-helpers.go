package plugins

import (
	"fmt"
	"html/template"
	"log"
	"path"
	"path/filepath"
	"strings"
	texttemplate "text/template"

	"core/internal/utils/flaretmpl"
	"core/internal/web/response"
	"core/internal/web/router"
	rnames "core/internal/web/routes/names"
	sdkhttp "sdk/api/http"
	plugin "sdk/api/plugin"
)

func NewHttpHelpers(api *PluginApi) sdkhttp.HttpHelpers {
	return &HttpHelpers{api: api}
}

type HttpHelpers struct {
	api *PluginApi
}

func (self *HttpHelpers) Translate(msgtype string, msgk string, pairs ...interface{}) string {
	return self.api.Utl.Translate(msgtype, msgk, pairs...)
}

func (self *HttpHelpers) AssetPath(p string) string {
	return path.Join("/plugin", self.api.Pkg(), self.api.Version(), "assets", p)
}

func (self *HttpHelpers) AssetWithHelpersPath(path string) string {
	assetsR := router.AssetsRouter.Get(rnames.RouteAssetWithHelpers)
	pluginApi := self.api
	url, err := assetsR.URL("pkg", pluginApi.Pkg(), "version", pluginApi.Version(), "path", path)
	if err != nil {
		log.Println("Error generating URL: ", err.Error())
		return ""
	}

	return url.String()
}

func (self *HttpHelpers) VueComponentPath(path string) string {
	assetsR := router.AssetsRouter.Get(rnames.RouteAssetVueComponent)
	if assetsR == nil {
		log.Println("Route not found: ", rnames.RouteAssetVueComponent)
		return ""
	}

	pluginApi := self.api
	url, err := assetsR.URL("pkg", pluginApi.Pkg(), "version", pluginApi.Version(), "path", path)
	if err != nil {
		log.Println("Error generating URL: ", err.Error())
		return ""
	}

	return url.String()
}

func (self *HttpHelpers) EmbedJs(path string, data interface{}) template.JS {
	jspath := self.api.Utl.Resource(filepath.Join("assets", path))

	var output strings.Builder

	jstmpl, err := flaretmpl.GetTextTemplate(jspath)
	if err != nil {
		jstmpl, _ = texttemplate.New("").Parse(fmt.Sprintf("console.error('%s: %s')", jspath, err.Error()))
	}

	vdata := &response.ViewData{
		ViewData:    data,
		ViewHelpers: self,
	}

	err = jstmpl.Execute(&output, vdata)
	if err != nil {
		log.Println("Error executing template: ", err.Error())
		return template.JS(fmt.Sprintf("console.error('%s: %s')", jspath, err.Error()))
	}

	return template.JS(output.String())
}

func (self *HttpHelpers) EmbedCss(path string, data interface{}) template.CSS {
	csspath := self.api.Utl.Resource(filepath.Join("assets", path))

	var output strings.Builder

	csstmpl, err := flaretmpl.GetTextTemplate(csspath)
	if err != nil {
		csstmpl, _ = texttemplate.New("").Parse(fmt.Sprintf("/* %s: %s */", csspath, err.Error()))
	}

	vdata := &response.ViewData{
		ViewData:    data,
		ViewHelpers: self,
	}

	err = csstmpl.Execute(&output, vdata)
	if err != nil {
		log.Println("Error executing template: ", err.Error())
		return template.CSS(fmt.Sprintf("/* %s: %s */", csspath, err.Error()))
	}

	return template.CSS(output.String())
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
