package navigation

import "net/http"

// INavApi is used to create navigation items in the application.
type INavApi interface {
	// Used to register a function that returns a slice of IAdminNavItem.
	// Items returned from this function is added to the admin navigation.
	AdminNavsFn(func(r *http.Request) []IAdminNavItem)

	// Used to register a function that returns a slice of IPortalItem.
	// Items returned from this function is added to the captive portal navigation.
	PortalNavsFn(func(r *http.Request) []IPortalItem)

	// Returns the admin navigation items returned by the function passed to AdminNavsFn().
	GetAdminNavs(r *http.Request) []IAdminNavItem

	// Returns the captive portal navigation items returned by the function passed to PortalNavsFn().
	GetPortalNavs(r *http.Request) []IPortalItem
}
