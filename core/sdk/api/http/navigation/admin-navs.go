package navigation

// List of admin navigation categories.
const (
	CategorySystem INavCategory = iota
	CategoryPayments
	CategoryNetwork
	CategoryThemes
	CategoryTools
)

type INavCategory uint8

// Represents an admin navigation item.
type IAdminNavItem interface {
	// Returns the category of the navigation item.
	Category() INavCategory

	// Returns the text of the navigation item.
	Text() string

	// Returns the href of the navigation item.
	Href() string
}

// IAdminNavList is used to group admin navigation items.
type IAdminNavList interface {
	// Returns the menu head of the navigation list.
	MenuHead() string

	// Returns the navigation items of the navigation list.
	Navs() []IAdminNavItem
}
