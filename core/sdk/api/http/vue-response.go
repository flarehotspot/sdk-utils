package sdkhttp

import "net/http"

type VueResponse interface {

	// Sends a flash message to the user. This does not send http response to the client.
	// It should be used along with methods that send actual http response like "Data" and "Redirect" methods.
	// Message types are "success", "error", "warning", "info".
	FlashMsg(msgType string, msg string)

	// Respond with json data.
	Json(w http.ResponseWriter, data any, status int)

	// Respond with a vue component. Useful for rendering dynamic component templates.
	Component(w http.ResponseWriter, vuefile string, data any)

	// Redirect to another vue route. "pairs" are param pairs, e.g. "id", "123", "name", "john".
	Redirect(w http.ResponseWriter, routename string, pairs ...string)

	Error(w http.ResponseWriter, err string, status int)
}
