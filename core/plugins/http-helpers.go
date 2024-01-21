package plugins

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
	texttemplate "text/template"

	httpI "github.com/flarehotspot/core/sdk/api/http"
	routerI "github.com/flarehotspot/core/sdk/api/http/router"
	"github.com/flarehotspot/core/sdk/api/plugin"
	"github.com/flarehotspot/core/sdk/utils/translate"
	"github.com/flarehotspot/core/web/response"
	"github.com/flarehotspot/core/web/router"
	rnames "github.com/flarehotspot/core/web/routes/names"
)

type ViewHelpers struct {
	api *PluginApi
}

func NewViewHelpers(api *PluginApi) httpI.IHelpers {
	return &ViewHelpers{api: api}
}

func (h *ViewHelpers) Translate(msgtype string, msgk string) string {
	return h.api.Translate(translate.MsgType(msgtype), msgk)
}

func (self *ViewHelpers) AssetPath(path string) string {
	return self.api.HttpApi().AssetPath(path)
}

func (self *ViewHelpers) AssetWithHelpersPath(path string) string {
	r := router.AssetsRouter.Get(rnames.AssetWithHelpers)
	pluginApi := self.api
	url, err := r.URL("pkg", pluginApi.Pkg(), "version", pluginApi.Version(), "path", path)
	if err != nil {
		log.Println("Error: ", err.Error())
		return ""
	}

	return url.String()
}

func (self *ViewHelpers) EmbedJs(path string, data any) template.HTML {
	jspath := self.api.Resource(filepath.Join("assets", path))
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

func (self *ViewHelpers) EmbedCss(path string, data any) template.HTML {
	csspath := self.api.Resource(filepath.Join("assets", path))
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

func (h *ViewHelpers) PluginMgr() plugin.IPluginMgr {
	return h.api.PluginsMgr
}

func (h *ViewHelpers) AdView() (html string) {
	return ""
}

func (h *ViewHelpers) MuxRouteName(name string) routerI.MuxRouteName {
	return h.api.HttpAPI.HttpRouter().MuxRouteName(routerI.PluginRouteName(name))
}

func (h *ViewHelpers) UrlForMuxRoute(name string, params ...string) string {
	url, _ := router.UrlForRoute(routerI.MuxRouteName(name), params...)
	return url
}

func (h *ViewHelpers) UrlForRoute(name string, params ...string) string {
	return h.api.HttpApi().HttpRouter().UrlForRoute(routerI.PluginRouteName(name), params...)
}

func (h *ViewHelpers) VueRouteName(name string) string {
	return h.api.HttpAPI.vueRouter.VueRouteName(name)
}
