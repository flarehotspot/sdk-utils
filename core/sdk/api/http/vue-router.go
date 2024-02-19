package sdkhttp

import "net/http"

const (
	VueNotFoundPath = "/404"
)

type VueAdminNavsFunc func(r *http.Request) []VueAdminNav
type VuePortalItemsFunc func(r *http.Request) []VuePortalItem

// VueRouter is used to create navigation items in the application.
type VueRouter interface {
	// Set admin vue route
	RegisterAdminRoutes(routes ...VueAdminRoute)

	// Set portal vue routes
	RegisterPortalRoutes(routes ...VuePortalRoute)

	// Used to register a function that returns a slice of admin navs.
	// Items returned from this function is added to the admin navigation menu.
	AdminNavsFunc(VueAdminNavsFunc)

	// Used to register a function that returns a slice of *PortalNavItem.
	// Items returned from this function is added to the captive portal navigation items.
	PortalItemsFunc(VuePortalItemsFunc)
}
