package response

import (
	"net/http"
)

// IHttpResponse is used to respond to http requests.
type IHttpResponse interface {
	// Used to render views from /resources/views/web-admin directory in your plugin.
	// For example if you have a view in /resources/views/web-admin/dashboard/index.html,
	// then you can render it with AdminView(w, r, "dashboard/index.html", data).
	AdminView(w http.ResponseWriter, r *http.Request, view string, data any)

	// Used to render views from /resources/views/captive-portal directory in your plugin.
	// For example if you have a view in /resources/views/captive-portal/payment/index.html,
	// then you can render it with PortalView(w, r, "payment/index.html", data).
	PortalView(w http.ResponseWriter, r *http.Request, view string, data any)

	// Used to render views from /resources/views directory in your plugin.
	// For example if you have a view in /resources/views/index.html,
	// then you can render it with View(w, r, "index.html", data).
	View(w http.ResponseWriter, r *http.Request, view string, data any)

	Script(w http.ResponseWriter, r *http.Request, file string, data any)

	// Used to send json response.
	Json(w http.ResponseWriter, data any, status int)
}
