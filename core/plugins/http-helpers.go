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

	sdkconnmgr "github.com/flarehotspot/flarehotspot/core/sdk/api/connmgr"
	sdkhttp "github.com/flarehotspot/flarehotspot/core/sdk/api/http"
	plugin "github.com/flarehotspot/flarehotspot/core/sdk/api/plugin"
	"github.com/flarehotspot/flarehotspot/core/utils/flaretmpl"
	"github.com/flarehotspot/flarehotspot/core/web/helpers"
	"github.com/flarehotspot/flarehotspot/core/web/response"
	"github.com/flarehotspot/flarehotspot/core/web/router"
	rnames "github.com/flarehotspot/flarehotspot/core/web/routes/names"
)

type HttpHelpers struct {
	api *PluginApi
}

func NewViewHelpers(api *PluginApi) sdkhttp.HttpHelpers {
	return &HttpHelpers{api: api}
}

func (h *HttpHelpers) Translate(msgtype string, msgk string) string {
	return h.api.Utl.Translate(msgtype, msgk)
}

func (h *HttpHelpers) GetClientDevice(r *http.Request) (sdkconnmgr.ClientDevice, error) {
	return helpers.CurrentClient(h.api.ClntReg, r)
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

func (self *HttpHelpers) EmbedJs(path string, data any) template.JS {
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

func (self *HttpHelpers) EmbedCss(path string, data any) template.CSS {
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

func (h *HttpHelpers) PluginMgr() plugin.PluginsMgrApi {
	return h.api.PluginsMgrApi
}

func (h *HttpHelpers) AdView() (html template.HTML) {
	return ""
}

func (h *HttpHelpers) MuxRouteName(name string) sdkhttp.MuxRouteName {
	return h.api.HttpAPI.HttpRouter().MuxRouteName(sdkhttp.PluginRouteName(name))
}

func (h *HttpHelpers) UrlForMuxRoute(name string, pairs ...string) string {
	return h.api.HttpAPI.HttpRouter().UrlForMuxRoute(sdkhttp.MuxRouteName(name), pairs...)
}

func (h *HttpHelpers) UrlForRoute(name string, pairs ...string) string {
	return h.api.HttpAPI.httpRouter.UrlForRoute(sdkhttp.PluginRouteName(name), pairs...)
}

func (h *HttpHelpers) VueRouteName(name string) string {
	return h.api.HttpAPI.vueRouter.VueRouteName(name)
}

func (h *HttpHelpers) VueRoutePath(name string, pairs ...string) string {
	var path VueRoutePath
	route, ok := h.api.HttpAPI.vueRouter.FindVueRoute(name)
	if !ok {
		path = sdkhttp.VueNotFoundPath
	}
	path = route.VueRoutePath
	return path.URL(pairs...)
}
