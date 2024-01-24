package sdkhttp

import "net/http"

const (
	VueNotFoundPath = "/404"
)

type VueAdminNavsFunc func(r *http.Request) []VueAdminNav
type VuePortalItemsFunc func(r *http.Request) []VuePortalItem

// IVueRouterApi is used to create navigation items in the application.
type IVueRouterApi interface {
	// Set admin vue route
	RegisterAdminRoutes(routes ...VueAdminRoute)

	// Used to register a function that returns a slice of admin navs.
	// Items returned from this function is added to the admin navigation menu.
	AdminNavsFunc(VueAdminNavsFunc)

	// Set portal vue routes
	SetPortalRoutes([]VuePortalRoute)

	// Used to register a function that returns a slice of *PortalNavItem.
	// Items returned from this function is added to the captive portal navigation items.
	PortalItemsFunc(VuePortalItemsFunc)
}
