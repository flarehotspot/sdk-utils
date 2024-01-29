package sdkhttp

type INavCategory string

// List of admin navigation menu categories.
const (
	NavCategorySystem   INavCategory = "system"
	NavCategoryPayments INavCategory = "payments"
	NavCategoryNetwork  INavCategory = "network"
	NavCategoryThemes   INavCategory = "themes"
	NavCategoryTools    INavCategory = "tools"
)

// VueAdminNav represents an admin navigation menu item.
type VueAdminNav struct {
	Category       INavCategory
	TranslateLabel string
	RouteName      string
	RouteParams    map[string]string
	PermitFn       func(perms []string) bool
}
