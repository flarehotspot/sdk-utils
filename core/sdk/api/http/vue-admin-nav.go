package sdkhttp

type INavCategory string

// List of admin navigation menu categories.
const (
	CategorySystem   INavCategory = "system"
	CategoryPayments INavCategory = "payments"
	CategoryNetwork  INavCategory = "network"
	CategoryThemes   INavCategory = "themes"
	CategoryTools    INavCategory = "tools"
)

// VueAdminNav represents an admin navigation menu item.
type VueAdminNav struct {
	Category       INavCategory
	TranslateLabel string
	RouteName      string
	PermitFn       func(perms []string) bool
}
