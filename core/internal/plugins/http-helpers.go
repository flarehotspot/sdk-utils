package plugins

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"strings"
	texttemplate "text/template"

	"github.com/flarehotspot/core/internal/utils/flaretmpl"
	"github.com/flarehotspot/core/internal/web/helpers"
	"github.com/flarehotspot/core/internal/web/response"
	"github.com/flarehotspot/core/internal/web/router"
	rnames "github.com/flarehotspot/core/internal/web/routes/names"
	sdkconnmgr "github.com/flarehotspot/core/sdk/api/connmgr"
	sdkhttp "github.com/flarehotspot/core/sdk/api/http"
	plugin "github.com/flarehotspot/core/sdk/api/plugin"
)

type HttpHelpers struct {
	api *PluginApi
}

func NewViewHelpers(api *PluginApi) sdkhttp.HttpHelpers {
	return &HttpHelpers{api: api}
}

func (self *HttpHelpers) Translate(msgtype string, msgk string) string {
	return self.api.Utl.Translate(msgtype, msgk)
}

func (self *HttpHelpers) GetClientDevice(r *http.Request) (sdkconnmgr.ClientDevice, error) {
	return helpers.CurrentClient(self.api.ClntReg, r)
}

func (self *HttpHelpers) AssetPath(p string) string {
	return path.Join("/plugin", self.api.Pkg(), self.api.Version(), "assets", p)
}

func (self *HttpHelpers) AssetWithHelpersPath(path string) string {
	r := router.AssetsRouter.Get(rnames.AssetWithHelpers)
	pluginApi := self.api
	url, err := r.URL("pkg", pluginApi.Pkg(), "version", pluginApi.Version(), "path", path)
	if err != nil {
		log.Println("Error: ", err.Error())
		return ""
	}

	return url.String()
}

func (self *HttpHelpers) VueComponentPath(path string) string {
	r := router.AssetsRouter.Get(rnames.AssetVueComponent)
	pluginApi := self.api
	url, err := r.URL("pkg", pluginApi.Pkg(), "version", pluginApi.Version(), "path", path)
	if err != nil {
		log.Println("Error: ", err.Error())
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

	jstmpl.Execute(&output, vdata)

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

	csstmpl.Execute(&output, vdata)

	return template.CSS(output.String())
}

func (self *HttpHelpers) PluginMgr() plugin.PluginsMgrApi {
	return self.api.PluginsMgrApi
}

func (self *HttpHelpers) AdView() (html template.HTML) {
	return ""
}

func (self *HttpHelpers) MuxRouteName(name string) sdkhttp.MuxRouteName {
	return self.api.HttpAPI.HttpRouter().MuxRouteName(sdkhttp.PluginRouteName(name))
}

func (self *HttpHelpers) UrlForMuxRoute(name string, pairs ...string) string {
	return self.api.HttpAPI.HttpRouter().UrlForMuxRoute(sdkhttp.MuxRouteName(name), pairs...)
}

func (self *HttpHelpers) UrlForRoute(name string, pairs ...string) string {
	return self.api.HttpAPI.httpRouter.UrlForRoute(sdkhttp.PluginRouteName(name), pairs...)
}

func (self *HttpHelpers) VueRouteName(name string) string {
	return self.api.HttpAPI.vueRouter.VueRouteName(name)
}

func (self *HttpHelpers) VueRoutePath(name string, pairs ...string) string {
	var path VueRoutePath
	route, ok := self.api.HttpAPI.vueRouter.FindVueRoute(name)
	if !ok {
		path = sdkhttp.VueNotFoundPath
	}
	path = route.VueRoutePath
	return path.URL(pairs...)
}
