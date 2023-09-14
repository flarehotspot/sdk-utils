package themes

import (
	"net/http"
)

// IThemesApi is used to configure and process themes.
type IThemesApi interface {

	// Define the handler for the index view of the captive portal.
	PortalIndexHandler(func(w http.ResponseWriter, r *http.Request))
}
