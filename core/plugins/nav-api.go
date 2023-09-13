package plugins

import (
	"github.com/flarehotspot/core/sdk/api/http/navigation"
	"net/http"
	"sync"
)

type NavApi struct {
	mu           sync.RWMutex
	pmgr         *PluginsMgr
	plugin       *PluginApi
	adminNavsFn  func(r *http.Request) []navigation.IAdminNavItem
	portalNavsFn func(r *http.Request) []navigation.IPortalItem
}

func (self *NavApi) AdminNavsFn(fn func(r *http.Request) []navigation.IAdminNavItem) {
	self.mu.Lock()
	defer self.mu.Unlock()
	self.adminNavsFn = fn
}

func (self *NavApi) PortalNavsFn(fn func(r *http.Request) []navigation.IPortalItem) {
	self.mu.Lock()
	defer self.mu.Unlock()
	self.portalNavsFn = fn
}

func (self *NavApi) GetAdminNavs(r *http.Request) []navigation.IAdminNavItem {
	self.mu.RLock()
	defer self.mu.RUnlock()
	if self.adminNavsFn != nil {
		return self.adminNavsFn(r)
	} else {
		return []navigation.IAdminNavItem{}
	}
}

func (self *NavApi) GetPortalNavs(r *http.Request) []navigation.IPortalItem {
	self.mu.RLock()
	defer self.mu.RUnlock()
	if self.portalNavsFn != nil {
		return self.portalNavsFn(r)
	} else {
		return []navigation.IPortalItem{}
	}
}

func NewNavApi(pmgr *PluginsMgr, plugin *PluginApi) *NavApi {
	return &NavApi{
		pmgr:   pmgr,
		plugin: plugin,
	}
}
