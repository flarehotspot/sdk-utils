package sdkhttp

import (
	"net/http"
)

// IHttpResponse is used to respond to http requests.
type IHttpResponse interface {
	// Used to render views from /resources/views/admin directory from your plugin.
	// For example if you have a view in /resources/views/admin/dashboard/index.html,
	// then you can render it with AdminView(w, r, "dashboard/index.html", data).
    // It uses the layout.html from your plugin directory /resources/views/admin/http-layout.html
	AdminView(w http.ResponseWriter, r *http.Request, view string, data any)

	// Used to render views from /resources/views/portal directory from your plugin.
	// For example if you have a view in /resources/views/portal/payment/index.html,
	// then you can render it with PortalView(w, r, "payment/index.html", data).
    // It uses the layout.html from your plugin directory /resources/views/portal/http-layout.html
	PortalView(w http.ResponseWriter, r *http.Request, view string, data any)

	// Used to render views from /resources/views directory from your plugin.
	// For example if you have a view in /resources/views/index.html,
	// then you can render it with View(w, r, "index.html", data).
	View(w http.ResponseWriter, r *http.Request, view string, data any)

	// Used to javascript templates from /resources/views/js directory from your plugin.
	// For example if you have a view in /resources/views/js/index.tpl.js,
	// then you can render it with Script(w, r, "index.tpl.js", data).
	Script(w http.ResponseWriter, r *http.Request, jspath string, data any)

	// Used to send json response.
	Json(w http.ResponseWriter, data any, status int)
}
