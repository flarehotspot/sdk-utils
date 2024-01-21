package plugins

import (
	"log"
	"path/filepath"

	"github.com/flarehotspot/core/config/plugincfg"
	"github.com/flarehotspot/core/connmgr"
	"github.com/flarehotspot/core/db"
	"github.com/flarehotspot/core/db/models"
	"github.com/flarehotspot/core/network"
	"github.com/flarehotspot/core/payments"
	plugin "github.com/flarehotspot/core/sdk/api/plugin"
	"github.com/flarehotspot/core/utils/migrate"
)

type PluginsMgr struct {
	db      *db.Database
	models  *models.Models
	paymgr  *payments.PaymentsMgr
	clntReg *connmgr.ClientRegister
	clntMgr *connmgr.ClientMgr
	trfkMgr *network.TrafficMgr
	plugins []*PluginApi
}

func NewPluginMgr(d *db.Database, m *models.Models, paymgr *payments.PaymentsMgr, clntReg *connmgr.ClientRegister, clntMgr *connmgr.ClientMgr, trfkMgr *network.TrafficMgr) *PluginsMgr {
	return &PluginsMgr{
		db:      d,
		models:  m,
		paymgr:  paymgr,
		clntReg: clntReg,
		clntMgr: clntMgr,
		plugins: []*PluginApi{},
	}
}

func (pmgr *PluginsMgr) Plugins() []*PluginApi {
	return pmgr.plugins
}

func (pmgr *PluginsMgr) RegisterPlugin(p *PluginApi) {
	pmgr.plugins = append(pmgr.plugins, p)

	err := p.Init()
	if err != nil {
		log.Println("Error initializing plugin: "+p.Dir(), err)
	}
}

func (pmgr *PluginsMgr) MigrateAll() {
	pluginDirs := plugincfg.ListDirs()
	for _, pdir := range pluginDirs {
		migdir := filepath.Join(pdir, "resources/migrations")
		err := migrate.MigrateUp(migdir, pmgr.db.SqlDB())
		if err != nil {
			log.Println("Error in plugin migration "+pdir, ":", err.Error())
		} else {
			log.Println("Done migrating plugin:", pdir)
		}
	}
}

func (pmgr *PluginsMgr) FindByName(name string) (plugin.IPluginApi, bool) {
	for _, p := range pmgr.plugins {
		if p.Name() == name {
			return p, true
		}
	}
	return nil, false
}

func (pmgr *PluginsMgr) FindByPkg(pkg string) (plugin.IPluginApi, bool) {
	for _, p := range pmgr.plugins {
		if p.Pkg() == pkg {
			return p, true
		}
	}
	return nil, false
}

func (pmgr *PluginsMgr) All() []plugin.IPluginApi {
	plugins := []plugin.IPluginApi{}
	for _, p := range pmgr.plugins {
		plugins = append(plugins, p)
	}
	return plugins
}

// func (pmgr *PluginsMgr) PortalPluginApi() *PluginApi {
// 	themepkg := themecfg.Read().Portal
// 	api, ok := pmgr.FindByPkg(themepkg)
// 	return api.(*PluginApi)
// }

// func (pmgr *PluginsMgr) AdminPluginApi() *PluginApi {
// 	themepkg := themecfg.Read().Admin
// 	api, ok := pmgr.FindByPkg(themepkg)
// 	return api.(*PluginApi)
// }

// func (pmgr *PluginsMgr) AuthPluginApi() *PluginApi {
// 	themepkg := themecfg.Read().Auth
// 	api, ok := pmgr.FindByPkg(themepkg)
// 	return api.(*PluginApi)
// }

func (pmgr *PluginsMgr) PaymentMethods() []plugin.IPluginApi {
	methods := []plugin.IPluginApi{}
	for _, p := range pmgr.plugins {
		pmnt := p.PaymentsApi().(*PaymentsApi)
		if pmnt.paymentsMgr != nil {
			methods = append(methods, p)
		}
	}
	return methods
}
