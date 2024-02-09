package plugins

import (
	"log"
	"path/filepath"

	"github.com/flarehotspot/core/config"
	"github.com/flarehotspot/core/connmgr"
	"github.com/flarehotspot/core/db"
	"github.com/flarehotspot/core/db/models"
	"github.com/flarehotspot/core/network"
	plugin "github.com/flarehotspot/core/sdk/api/plugin"
	"github.com/flarehotspot/core/utils/migrate"
)

type PluginsMgr struct {
	CoreAPI *PluginApi
	db      *db.Database
	models  *models.Models
	paymgr  *PaymentsMgr
	clntReg *connmgr.ClientRegister
	clntMgr *connmgr.SessionsMgr
	trfkMgr *network.TrafficMgr
	plugins []*PluginApi
	utils   *PluginsMgrUtils
}

func NewPluginMgr(d *db.Database, m *models.Models, paymgr *PaymentsMgr, clntReg *connmgr.ClientRegister, clntMgr *connmgr.SessionsMgr, trfkMgr *network.TrafficMgr) *PluginsMgr {
	pmgr := &PluginsMgr{
		db:      d,
		models:  m,
		paymgr:  paymgr,
		clntReg: clntReg,
		clntMgr: clntMgr,
		plugins: []*PluginApi{},
	}
	return pmgr
}

func (pmgr *PluginsMgr) InitCoreApi(coreApi *PluginApi) {
	pmgr.CoreAPI = coreApi
	pmgr.utils = NewPluginsMgrUtil(pmgr, coreApi)
	pmgr.RegisterPlugin(coreApi)
}

func (pmgr *PluginsMgr) Plugins() []*PluginApi {
	return pmgr.plugins
}

func (pmgr *PluginsMgr) RegisterPlugin(p *PluginApi) {
	p.InitCoreApi(pmgr.CoreAPI)
	pmgr.plugins = append(pmgr.plugins, p)

	if p.Pkg() != pmgr.CoreAPI.Pkg() {
		err := p.Init()
		if err != nil {
			log.Println("Error initializing plugin: "+p.Dir(), err)
		}
	}
}

func (pmgr *PluginsMgr) MigrateAll() {
	pluginDirs := config.PluginDirList()
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

func (pmgr *PluginsMgr) FindByName(name string) (plugin.PluginApi, bool) {
	for _, p := range pmgr.plugins {
		if p.Name() == name {
			return p, true
		}
	}
	return nil, false
}

func (pmgr *PluginsMgr) FindByPkg(pkg string) (plugin.PluginApi, bool) {
	for _, p := range pmgr.plugins {
		if p.Pkg() == pkg {
			return p, true
		}
	}
	return nil, false
}

func (pmgr *PluginsMgr) All() []plugin.PluginApi {
	plugins := []plugin.PluginApi{}
	for _, p := range pmgr.plugins {
		plugins = append(plugins, p)
	}
	return plugins
}

func (pmgr *PluginsMgr) PaymentMethods() []plugin.PluginApi {
	methods := []plugin.PluginApi{}
	for _, p := range pmgr.plugins {
		pmnt := p.Payments().(*PaymentsApi)
		if pmnt.paymentsMgr != nil {
			methods = append(methods, p)
		}
	}
	return methods
}

func (pmgr *PluginsMgr) Utils() *PluginsMgrUtils {
	return pmgr.utils
}
