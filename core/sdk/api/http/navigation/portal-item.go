package navigation

// PortalItem is used to create basic navigation item in the captive portal.
// For advanced type of navigation item, you have to implement the IPortalItem interface.
type PortalItem struct {
	iconPath string
	text     string
	desc     string
	href     string
}

// Returns the icon path of the navigation item.
func (data *PortalItem) IconPath() string {
	return ""
}

// Returns the text of the navigation item.
func (data *PortalItem) Text() string {
	return data.text
}

// Returns the href of the navigation item.
func (data *PortalItem) Href() string {
	return data.href
}

// Creates a new instance of PortalItem.
func NewPortalItem(icon string, text string, desc string, href string) IPortalItem {
	return &PortalItem{
		iconPath: icon,
		text:     text,
		desc:     desc,
		href:     href,
	}
}
