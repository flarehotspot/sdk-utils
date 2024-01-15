package router

type VueAdminRoute struct {
	RouteName           string
	RoutePath           string
	ComponentPath       string
	PermissionsRequired []string
	PermissionsAnyOf    []string
}
