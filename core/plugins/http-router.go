package plugins

import (
	"net/http"

	sdkhttp "github.com/flarehotspot/core/sdk/api/http"
	"github.com/gorilla/mux"
)

type HttpRouter struct {
	api *PluginApi
	mux *mux.Router
}

func (r *HttpRouter) Router() *mux.Router {
	return r.mux
}

func (r *HttpRouter) Get(path string, h http.HandlerFunc) sdkhttp.IHttpRoute {
	route := r.mux.HandleFunc(path, h).Methods("GET")
	return &HttpRoute{r.api, route}
}
func (r *HttpRouter) Post(path string, h http.HandlerFunc) sdkhttp.IHttpRoute {
	route := r.mux.HandleFunc(path, h).Methods("POST")
	return &HttpRoute{r.api, route}
}

func (r *HttpRouter) Group(path string, fn func(sdkhttp.IHttpRouter)) {
	router := r.mux.PathPrefix(path).Subrouter()
	newrouter := HttpRouter{api: r.api, mux: router}
	fn(&newrouter)
}

func (r *HttpRouter) Use(middlewares ...func(http.Handler) http.Handler) {
	for _, mw := range middlewares {
		r.mux.Use(mw)
	}
}
