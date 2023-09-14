package navigation

import (
	"github.com/flarehotspot/core/sdk/api/http/navigation"
	"github.com/flarehotspot/core/sdk/api/http/router"
)

type PortalItems []navigation.IPortalItem

type PortalItem struct {
	IconPath  string
	Label     string
	Translate bool
	RouteName router.PluginRouteName
}
