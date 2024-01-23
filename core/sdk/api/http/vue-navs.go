package sdkhttp

type AdminNavCategory struct {
	Category string         `json:"category"`
	Items    []AdminNavItem `json:"items"`
}

type AdminNavItem struct {
	Category     INavCategory `json:"category"`
	Label        string       `json:"label"`
	VueRouteName string       `json:"route_name"`
	VueRoutePath string       `json:"route_path"`
}
