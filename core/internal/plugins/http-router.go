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

func (r *HttpRouter) Get(path string, h http.HandlerFunc, mw ...func(next http.Handler) http.Handler) sdkhttp.HttpRoute {
	finalHandler := http.Handler(h)
	for i := len(mw) - 1; i >= 0; i-- {
		finalHandler = mw[i](finalHandler)
	}
	route := r.mux.Handle(path, finalHandler).Methods("GET")
	return &HttpRoute{r.api, route}
}
func (r *HttpRouter) Post(path string, h http.HandlerFunc, mw ...func(next http.Handler) http.Handler) sdkhttp.HttpRoute {
	finalHandler := http.Handler(h)
	for i := len(mw) - 1; i >= 0; i-- {
		finalHandler = mw[i](finalHandler)
	}
	route := r.mux.Handle(path, finalHandler).Methods("POST")
	return &HttpRoute{r.api, route}
}

func (r *HttpRouter) Group(path string, fn func(sdkhttp.RouterInstance)) {
	router := r.mux.PathPrefix(path).Subrouter()
	newrouter := &HttpRouter{api: r.api, mux: router}
	fn(newrouter)
}

func (r *HttpRouter) Use(middlewares ...func(http.Handler) http.Handler) {
	for _, mw := range middlewares {
		r.mux.Use(mw)
	}
}
