package plugins

import (
	"fmt"
	"html/template"
	nethttp "net/http"
	"os"
	"path/filepath"
	"strings"
	texttemplate "text/template"

	"github.com/flarehotspot/core/connmgr"
	"github.com/flarehotspot/core/db/models"
	"github.com/flarehotspot/core/payments"
	"github.com/flarehotspot/core/sdk/api/http"
	"github.com/flarehotspot/core/sdk/api/http/middlewares"
	"github.com/flarehotspot/core/sdk/api/http/response"
	"github.com/flarehotspot/core/sdk/api/http/router"
	"github.com/gorilla/mux"
)

type HttpApi struct {
	api         *PluginApi
	httpRouter  *HttpRouterApi
	vueRouter   *VueRouterApi
	response    *HttpResponse
	middlewares *PluginMiddlewares
}

func NewHttpApi(api *PluginApi, mdls *models.Models, dmgr *connmgr.ClientRegister, pmgr *payments.PaymentsMgr) *HttpApi {
	httpRouter := NewRouterApi(api)
	vueRouter := NewVueRouterApi(api)
	response := NewHttpResponse(api)
	middlewares := NewPluginMiddlewares(api.db, mdls, dmgr, pmgr)
	return &HttpApi{
		api:         api,
		httpRouter:  httpRouter,
		vueRouter:   vueRouter,
		response:    response,
		middlewares: middlewares,
	}
}

func (self *HttpApi) HttpRouter() router.IHttpRouterApi {
	return self.httpRouter
}

func (self *HttpApi) VueRouter() router.IVueRouterApi {
	return self.vueRouter
}

func (self *HttpApi) Helpers(w nethttp.ResponseWriter, r *nethttp.Request) http.IHelpers {
	return NewViewHelpers(self.api, w, r)
}

func (self *HttpApi) AssetPath(path string) string {
	return filepath.Join("/plugin", self.api.Pkg(), self.api.Version(), "assets", path)
}

func (self *HttpApi) EmbedJs(path string, data any) template.HTML {
	jspath := self.api.Resource(filepath.Join("assets", path))
	tpljs, err := os.ReadFile(jspath)
	if err != nil {
		tpljs = []byte(fmt.Sprintf("console.log('%s: %s')", jspath, err.Error()))
	}

	jstmpl, err := texttemplate.New("").Parse(string(tpljs))
	if err != nil {
		jstmpl, _ = texttemplate.New("").Parse(fmt.Sprintf("console.log('%s: %s')", jspath, err.Error()))
	}

	var output strings.Builder
	jstmpl.Execute(&output, data)

	scriptTag := fmt.Sprintf("<script>%s</script>", output.String())
	return template.HTML(scriptTag)
}

func (self *HttpApi) EmbedCss(path string, data any) template.HTML {
	csspath := self.api.Resource(filepath.Join("assets", path))
	tplcss, err := os.ReadFile(csspath)
	if err != nil {
		tplcss = []byte(fmt.Sprintf("/* %s: %s */", csspath, err.Error()))
	}

	csstmpl, err := texttemplate.New("").Parse(string(tplcss))
	if err != nil {
		csstmpl, _ = texttemplate.New("").Parse(fmt.Sprintf("/* %s: %s */", csspath, err.Error()))
	}

	var output strings.Builder
	csstmpl.Execute(&output, data)

	styleTag := fmt.Sprintf("<style>%s</style>", output.String())
	return template.HTML(styleTag)
}

func (self *HttpApi) Middlewares() middlewares.Middlewares {
	return self.middlewares
}

func (self *HttpApi) Respond() response.IHttpResponse {
	return self.response
}

func (self *HttpApi) MuxVars(r *nethttp.Request) map[string]string {
	return mux.Vars(r)
}
