package router

import "net/http"

const (
	NotFoundVuePath = "/404"
)

// IVueRouterApi is used to create navigation items in the application.
type IVueRouterApi interface {
	AdminRoutes(func(r *http.Request) []*VueRoute)

	// Used to register a function that returns a slice of *AdminNav.
	// Items returned from this function is added to the admin navigation menu.
	AdminNavs(func(r *http.Request) []*VueAdminNav)

	// Used to register a function that returns a slice of *PortalRouteItem.
	// Items return from this function are added to the captive portal front-end (vuejs) routes.
	PortalRoutes(func(r *http.Request) []*VueRoute)

	// Used to register a function that returns a slice of *PortalNavItem.
	// Items returned from this function is added to the captive portal navigation items.
	PortalItems(func(r *http.Request) []*VuePortalItem)
}
