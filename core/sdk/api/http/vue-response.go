package sdkhttp

type IVueResponse interface {

    // Send a flash message to the user along with json data or redirect response.
    // This does not send http response to the client. It should be used along with methods that send actual http response like "JsonData" and "Redirect".
   // Message types are "success", "error", "warning", "info".
	FlashMsg(msgType string, msg string)

    // Respond with json data.
	JsonData(data any)

    // Redirect to another vue route. "pairs" are param pairs, e.g. "id", "123", "name", "john".
	Redirect(routename string, pairs ...string)
}
