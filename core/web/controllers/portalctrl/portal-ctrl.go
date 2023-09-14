package portalctrl

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/web/helpers"
	"github.com/flarehotspot/core/web/response"
	"github.com/flarehotspot/core/web/routes/names"
)

type PortalCtrl struct {
	g *globals.CoreGlobals
}

func NewPortalCtrl(g *globals.CoreGlobals) PortalCtrl {
	return PortalCtrl{g}
}

func (self *PortalCtrl) IndexPage(w http.ResponseWriter, r *http.Request) {
	api := self.g.PluginMgr.PortalPluginApi()
	handler := api.ThemesAPI.GetPortalHandler()
	if handler == nil {
		self.Error(w, r, errors.New("captive Portal theme must implement handler function"))
		return
	}

	handler(w, r)
}

// TODO: Remove this test method
func (self *PortalCtrl) Test(w http.ResponseWriter, r *http.Request) {
	clnt, err := helpers.CurrentClient(r)
	if err != nil {
		self.Error(w, r, err)
		return
	}

	self.g.Models.Session().Create(r.Context(), clnt.Id(), 0, 30, 0, nil, 1, 1, false)

	w.WriteHeader(200)
}

type testpage struct {
	Title string
	Data  map[string]any
}

func (self *PortalCtrl) TestTemplate(w http.ResponseWriter, r *http.Request) {
	p := &testpage{
		Title: "Some page title",
		Data:  map[string]any{"data": "data value"},
	}
	t, err := template.New("page").Parse("Title is \"{{ .Title }}\" and data is \"{{ .Data.data }}\".")
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, p)
	if err != nil {
		panic(err)
	}
}

func (self *PortalCtrl) Error(w http.ResponseWriter, r *http.Request, err error) {
	e := response.NewErrRoute(names.RoutePortalIndex)
	e.Redirect(w, r, err)
}
