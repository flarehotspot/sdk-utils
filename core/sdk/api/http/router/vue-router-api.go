package router

const (
	VueNotFoundPath = "/404"
)

// IVueRouterApi is used to create navigation items in the application.
type IVueRouterApi interface {
	// Set admin vue routes
	AdminRoutes([]VueAdminRoute)

	// Used to register a function that returns a slice of *AdminNav.
	// Items returned from this function is added to the admin navigation menu.
	AdminNavs(VueAdminNavsHandler)

	// Set portal vue routes
	PortalRoutes([]VuePortalRoute)

	// Used to register a function that returns a slice of *PortalNavItem.
	// Items returned from this function is added to the captive portal navigation items.
	PortalItems(VuePortalItemsHandler)
}
