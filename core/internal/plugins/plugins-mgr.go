package plugins

import (
	"log"
	"path/filepath"

	"core/internal/config"
	"core/internal/connmgr"
	"core/internal/db"
	"core/internal/db/models"
	"core/internal/network"
	"core/internal/utils/migrate"
	"sdk/api/plugin"
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
		err := migrate.MigrateUp(self.db.SqlDB(), migdir)
		if err != nil {
			log.Println("Error in plugin migration "+pdir, ":", err.Error())
		} else {
			log.Println("Done migrating plugin:", pdir)
		}
	}
}

func (self *PluginsMgr) FindByName(name string) (sdkplugin.PluginApi, bool) {
	for _, p := range self.plugins {
		if p.Name() == name {
			return p, true
		}
	}
	return nil, false
}

func (self *PluginsMgr) FindByPkg(pkg string) (sdkplugin.PluginApi, bool) {
	for _, p := range self.plugins {
		if p.Pkg() == pkg {
			return p, true
		}
	}
	return nil, false
}

func (self *PluginsMgr) All() []sdkplugin.PluginApi {
	plugins := []sdkplugin.PluginApi{}
	for _, p := range self.plugins {
		plugins = append(plugins, p)
	}
	return plugins
}

func (self *PluginsMgr) PaymentMethods() []sdkplugin.PluginApi {
	methods := []sdkplugin.PluginApi{}
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
