package navigation

// IPortalItem represents a portal navigation item.
type IPortalItem interface {
	// Returns the path to the icon.
	IconPath() string

	// Returns the text to display.
	Text() string

	// Returns the URL to the item.
	Href() string
}
