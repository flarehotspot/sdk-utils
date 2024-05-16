/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkhttp

import (
	"net/http"
)

// HttpResponse is used to respond to http requests.
type HttpResponse interface {

	// Used to render views from /resources/views/portal directory from your plugin.
	// For example if you have a view in /resources/views/portal/payment/index.html,
	// then you can render it with PortalView(w, r, "payment/index.html", data).
	// It uses the layout.html from your plugin directory /resources/views/portal/layout.html
	PortalView(w http.ResponseWriter, r *http.Request, view string, data any)

	// Used to render views from /resources/views/admin directory from your plugin.
	// For example if you have a view in /resources/views/admin/dashboard/index.html,
	// then you can render it with AdminView(w, r, "dashboard/index.html", data).
	// It uses the layout.html from your plugin directory /resources/views/admin/layout.html
	AdminView(w http.ResponseWriter, r *http.Request, view string, data any)

	// Used to render single file views (without layout) from /resources/views directory from your plugin.
	// For example if you have a view in /resources/views/index.html,
	// then you can render it with View(w, r, "index.html", data).
	View(w http.ResponseWriter, r *http.Request, view string, data any)

	// Used to render resource files  from the resources directory in your plugin.
	// For example if you have a view in /resources/views/js/index.tmpl.js,
	// then you can render it with File(w, r, "views/js/index.tmpl.js", data).
	File(w http.ResponseWriter, r *http.Request, jspath string, data any)

	// Used to send json response.
	Json(w http.ResponseWriter, data any, status int)
}
