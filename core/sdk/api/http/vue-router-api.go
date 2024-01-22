package sdkhttp

const (
	VueNotFoundPath = "/404"
)

// IVueRouterApi is used to create navigation items in the application.
type IVueRouterApi interface {
	// Set admin vue routes
	SetAdminRoutes([]VueAdminRoute)

	// Used to register a function that returns a slice of *AdminNav.
	// Items returned from this function is added to the admin navigation menu.
	AdminNavsFunc(VueAdminNavsFunc)

	// Set portal vue routes
	SetPortalRoutes([]VuePortalRoute)

	// Used to register a function that returns a slice of *PortalNavItem.
	// Items returned from this function is added to the captive portal navigation items.
	PortalItemsFunc(VuePortalItemsFunc)

	// Returns the vue route name for a named route which can be used in <flare-link>
	VueRouteName(name string) string
}
