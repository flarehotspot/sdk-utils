package router

type INavCategory uint8

// List of admin navigation categories.
const (
	CategorySystem INavCategory = iota
	CategoryPayments
	CategoryNetwork
	CategoryThemes
	CategoryTools
)

// VueAdminNav represents an admin navigation item.
// It implements IAdminNavItem and is used to create basic admin navigation items without translations.
// For advanced type of navigation, you have to implement IAdminNavItem yourself.
type VueAdminNav struct {
	Category       INavCategory
	TranslateLabel string
	RouteName      string
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
