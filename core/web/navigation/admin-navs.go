package navigation

import (
	"github.com/flarehotspot/core/sdk/api/http/navigation"
)

type AdminNavListItem struct {
	menuHead string
	navs     []navigation.IAdminNavItem
	perms    []string
}

func (self *AdminNavListItem) MenuHead() string {
	return self.menuHead
}

func (self *AdminNavListItem) Navs() []navigation.IAdminNavItem {
	return self.navs
}

func (self *AdminNavListItem) AddNav(nav navigation.IAdminNavItem) {
	self.navs = append(self.navs, nav)
}

func NewAdminListItem(menuHead string, perms []string) *AdminNavListItem {
	return &AdminNavListItem{
		menuHead: menuHead,
		navs:     []navigation.IAdminNavItem{},
		perms:    perms,
	}
}
