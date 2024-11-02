/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkhttp

// ViewData represents the data passed to the views in your plugin.
// It is composed of two parts, the view helpers and view data you passed to the view.
// The helpers data can be accessed using ".Helpers()" while the data is ".Data()".
//
// For example, if you rendered a view using AdminView(w, r "dashboard/index.html", data),
// Then your custom data can be accessed in your view with: {{ .Data() }}.
// The helpers can be accessed with: {{ .Helpers() }}.

// type IViewData interface {
// 	Helpers() HttpHelpers
// 	Data() any
// }
