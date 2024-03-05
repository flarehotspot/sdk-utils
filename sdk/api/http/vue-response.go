/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkhttp

import "net/http"

type VueResponse interface {

	// Sets a flash message. It does not send http response to the client.
	// It should be used along with methods that send actual http response like "Json" and "Redirect" methods.
	// Message types are "success", "error", "warning", "info".
	SetFlashMsg(msgType string, msg string)


    // Similar to SetFlashMsg, but it sends an HTTP response to the client.
	SendFlashMsg(w http.ResponseWriter, msgType string, msg string, status int)

	// Respond with json data.
    // It sends an HTTP response and must be put as last line in the handler function.
	Json(w http.ResponseWriter, data any, status int)

	// Respond with a vue component. Useful for rendering dynamic component templates.
    // It sends an HTTP response and must be put as last line in the handler function.
	Component(w http.ResponseWriter, vuefile string, data any)

	// Redirect to another vue route. "pairs" are param pairs, e.g. "id", "123", "name", "john".
    // It sends an HTTP response and must be put as last line in the handler function.
	Redirect(w http.ResponseWriter, routename string, pairs ...string)

	// Redirect portal index page
    // It sends an HTTP response and must be put as last line in the handler function.
	RedirectToPortal(w http.ResponseWriter)

    // Respond with an error flash message.
    // It sends an HTTP response and must be put as last line in the handler function.
	Error(w http.ResponseWriter, err string, status int)
}
