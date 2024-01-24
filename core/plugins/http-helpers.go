package plugins

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
	texttemplate "text/template"

	"github.com/flarehotspot/core/sdk/api/http"
	plugin "github.com/flarehotspot/core/sdk/api/plugin"
	"github.com/flarehotspot/core/web/response"
	"github.com/flarehotspot/core/web/router"
	rnames "github.com/flarehotspot/core/web/routes/names"
)

type HttpHelpers struct {
	api *PluginApi
}

func NewViewHelpers(api *PluginApi) sdkhttp.IHelpers {
	return &HttpHelpers{api: api}
}

func (h *HttpHelpers) Translate(msgtype string, msgk string) string {
	return h.api.Utl.Translate(msgtype, msgk)
}

func (self *HttpHelpers) AssetPath(path string) string {
	return filepath.Join("/plugin", self.api.Pkg(), self.api.Version(), "assets", path)
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

func (self *HttpHelpers) EmbedJs(path string, data any) template.HTML {
	jspath := self.api.Utl.Resource(filepath.Join("assets", path))
	tpljs, err := os.ReadFile(jspath)
	if err != nil {
		tpljs = []byte(fmt.Sprintf("console.error('%s: %s')", jspath, err.Error()))
	}

	var output strings.Builder

	jstmpl, err := texttemplate.New("").Parse(string(tpljs))
	if err != nil {
		jstmpl, _ = texttemplate.New("").Parse(fmt.Sprintf("console.log('%s: %s')", jspath, err.Error()))
	}

	vdata := &response.ViewData{
		ViewData:    data,
		ViewHelpers: self,
	}

	jstmpl.Execute(&output, vdata)

	scriptTag := fmt.Sprintf("<script>%s</script>", output.String())
	return template.HTML(scriptTag)
}

func (self *HttpHelpers) EmbedCss(path string, data any) template.HTML {
	csspath := self.api.Utl.Resource(filepath.Join("assets", path))
	tplcss, err := os.ReadFile(csspath)
	if err != nil {
		tplcss = []byte(fmt.Sprintf("/* %s: %s */", csspath, err.Error()))
	}

	var output strings.Builder

	csstmpl, err := texttemplate.New("").Parse(string(tplcss))
	if err != nil {
		csstmpl, _ = texttemplate.New("").Parse(fmt.Sprintf("/* %s: %s */", csspath, err.Error()))
	}

	vdata := &response.ViewData{
		ViewData:    data,
		ViewHelpers: self,
	}

	csstmpl.Execute(&output, vdata)

	styleTag := fmt.Sprintf("<style>%s</style>", output.String())
	return template.HTML(styleTag)
}

func (h *HttpHelpers) PluginMgr() plugin.IPluginMgr {
	return h.api.PluginsMgr
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

func (h *HttpHelpers) VueRoutePath(name string) string {
	var path string
	route, ok := h.api.HttpAPI.vueRouter.FindVueRoute(name)
	if !ok {
		path = sdkhttp.VueNotFoundPath
	}
	path = route.VueRoutePath
    return path
}
