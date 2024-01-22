package sdkhttp

type AdminNavCategory struct {
	Category string         `json:"menu_head"`
	Items    []AdminNavItem `json:"items"`
}

type AdminNavItem struct {
	Label     string `json:"label"`
	RouteName string `json:"route_name"`
}
