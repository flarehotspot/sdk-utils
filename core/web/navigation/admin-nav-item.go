package navigation

import (
	coreR "github.com/flarehotspot/core/web/router"
	"github.com/flarehotspot/core/sdk/api/http/navigation"
	"github.com/flarehotspot/core/sdk/api/http/router"
)

type AdminNavItem struct {
	category  navigation.INavCategory
	text      string
	routeName router.MuxRouteName
	perms     []string
}

func (self *AdminNavItem) Category() navigation.INavCategory {
	return self.category
}
func (self *AdminNavItem) Text() string {
	return self.text
}

func (self *AdminNavItem) Href() string {
	url, err := coreR.UrlForRoute(self.routeName)
	if err != nil {
		return coreR.NotFoundRoute
	}
	return url
}

func NewAdminNavItem(cat navigation.INavCategory, text string, route router.MuxRouteName, perms []string) navigation.IAdminNavItem {
	return &AdminNavItem{cat, text, route, perms}
}
