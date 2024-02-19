package apictrl

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/flarehotspot/core/internal/plugins"
)

type AdminNavJson struct {
	Text string `json:"text"`
	Href string `json:"href"`
}

type NavsCtrl struct {
	pmgr *plugins.PluginsMgr
}

func NewNavsCtrl(pmgr *plugins.PluginsMgr) *NavsCtrl {
	return &NavsCtrl{pmgr}
}

func (self *NavsCtrl) NavSearchJson(w http.ResponseWriter, r *http.Request) {
	// api := self.pmgr.AdminPluginApi()
	// helpers := plugins.NewViewHelpers(api, w, r)
	// navList := helpers.GetAdminNavs()

	navsJson := []AdminNavJson{}
	// for _, list := range navList {
	// 	for _, item := range list.Navs() {
	// 		nav := AdminNavJson{
	// 			Text: item.Text(),
	// 			Href: item.Href(),
	// 		}
	// 		navsJson = append(navsJson, nav)
	// 	}
	// }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(navsJson)
	if err != nil {
		log.Println(err)
	}
}
