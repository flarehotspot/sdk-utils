package response

import "net/http"

type IErrorRedirect interface {

	// Redirects to url or named route with error message. The pairs params is used to override the default
	// route parameters or to add more route params. For example, if you want to override the default route parameter
	// "id", then you can do this:
  //  e := api.HttpApi().Respond().NewErrRoute("my-route", "id", 1)
  //  e.Redirect(w, r, err, "id", 2)
	// This will redirect to the route "my-route" with the url parameter "id" set to "2".
  // You can also add more pairs using this method:
  //  e.Redirect(w, r, err, "id", 2, "foo", "bar")
	Redirect(w http.ResponseWriter, r *http.Request, err error, pairs ...string)
}
