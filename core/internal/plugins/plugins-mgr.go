package plugins

import (
	"log"
	"path/filepath"

	"github.com/flarehotspot/core/internal/config"
	"github.com/flarehotspot/core/internal/connmgr"
	"github.com/flarehotspot/core/internal/db"
	"github.com/flarehotspot/core/internal/db/models"
	"github.com/flarehotspot/core/internal/network"
	"github.com/flarehotspot/core/internal/utils/migrate"
	plugin "github.com/flarehotspot/core/sdk/api/plugin"
)

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

func (self *PluginsMgr) InitCoreApi(coreApi *PluginApi) {
	self.CoreAPI = coreApi
	self.utils = NewPluginsMgrUtil(self, coreApi)
	self.RegisterPlugin(coreApi)
}

func (self *PluginsMgr) Plugins() []*PluginApi {
	return self.plugins
}

func (self *PluginsMgr) RegisterPlugin(p *PluginApi) {
	p.InitCoreApi(self.CoreAPI)
	self.plugins = append(self.plugins, p)

	if p.Pkg() != self.CoreAPI.Pkg() {
		err := p.Init()
		if err != nil {
			log.Println("Error initializing plugin: "+p.Dir(), err)
		}
	}
}

func (self *PluginsMgr) MigrateAll() {
	pluginDirs := config.PluginDirList()
	for _, pdir := range pluginDirs {
		migdir := filepath.Join(pdir, "resources/migrations")
		err := migrate.MigrateUp(migdir, self.db.SqlDB())
		if err != nil {
			log.Println("Error in plugin migration "+pdir, ":", err.Error())
		} else {
			log.Println("Done migrating plugin:", pdir)
		}
	}
}

func (self *PluginsMgr) FindByName(name string) (plugin.PluginApi, bool) {
	for _, p := range self.plugins {
		if p.Name() == name {
			return p, true
		}
	}
	return nil, false
}

func (self *PluginsMgr) FindByPkg(pkg string) (plugin.PluginApi, bool) {
	for _, p := range self.plugins {
		if p.Pkg() == pkg {
			return p, true
		}
	}
	return nil, false
}

func (self *PluginsMgr) All() []plugin.PluginApi {
	plugins := []plugin.PluginApi{}
	for _, p := range self.plugins {
		plugins = append(plugins, p)
	}
	return plugins
}

func (self *PluginsMgr) PaymentMethods() []plugin.PluginApi {
	methods := []plugin.PluginApi{}
	for _, p := range self.plugins {
		pmnt := p.Payments().(*PaymentsApi)
		if pmnt.paymentsMgr != nil {
			methods = append(methods, p)
		}
	}
	return methods
}

func (self *PluginsMgr) Utils() *PluginsMgrUtils {
	return self.utils
}
