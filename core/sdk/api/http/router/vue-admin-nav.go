package router

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
	PermitFn       *func(perms []string) bool
}

type AdminNavJson struct {
	Category INavCategory `json:"category"`
	Label    string       `json:"label"`
	Route    string       `json:"route"`
}

type AdminNavList struct {
	MenuHead string          `json:"menu_head"`
	Navs     []*AdminNavJson `json:"navs"`
	Perms    []string        `json:"perms"`
}

func (list *AdminNavList) AddNav(nav *AdminNavJson) {
	list.Navs = append(list.Navs, nav)
}

func NewAdminList(menuHead string, perms []string) *AdminNavList {
	return &AdminNavList{
		MenuHead: menuHead,
		Navs:     []*AdminNavJson{},
		Perms:    perms,
	}
}
