package navigation

// AdminNav represents an admin navigation item.
// It implements IAdminNavItem and is used to create basic admin navigation items without translations.
// For advanced type of navigation, you have to implement IAdminNavItem yourself.
type AdminNav struct {
	category INavCategory
	text     string
	href     string
}

// Returns the category of the navigation item.
func (self *AdminNav) Category() INavCategory {
	return self.category
}

// Returns the text of the navigation item.
func (self *AdminNav) Text() string {
	return self.text
}

// Returns the href of the navigation item.
func (self *AdminNav) Href() string {
	return self.href
}

// Returns a new admin navigation item.
func NewAdminNav(category INavCategory, text string, href string) IAdminNavItem {
	return &AdminNav{
		category: category,
		text:     text,
		href:     href,
	}
}
