package plugins

import (
	"fmt"
	"net/http"
	"strings"

	IRouter "github.com/flarehotspot/core/sdk/api/http/router"
	"github.com/gorilla/mux"
)

type HttpRouter struct {
	api *PluginApi
	mux *mux.Router
}

func (r *HttpRouter) Router() *mux.Router {
	return r.mux
}

func (r *HttpRouter) Get(path string, h http.HandlerFunc) IRouter.IHttpRoute {
	path = sanitizePath(path)
	route := r.mux.HandleFunc(path, h).Methods("GET")
	return &HttpRoute{r.api, route}
}
func (r *HttpRouter) Post(path string, h http.HandlerFunc) IRouter.IHttpRoute {
	path = sanitizePath(path)
	route := r.mux.HandleFunc(path, h).Methods("POST")
	return &HttpRoute{r.api, route}
}

func (r *HttpRouter) Put(path string, h http.HandlerFunc) IRouter.IHttpRoute {
	path = sanitizePath(path)
	route := r.mux.HandleFunc(path, h).Methods("PUT")
	return &HttpRoute{r.api, route}
}

func (r *HttpRouter) Delete(path string, h http.HandlerFunc) IRouter.IHttpRoute {
	path = sanitizePath(path)
	route := r.mux.HandleFunc(path, h).Methods("DELETE")
	return &HttpRoute{r.api, route}
}

func (r *HttpRouter) Options(path string, h http.HandlerFunc) IRouter.IHttpRoute {
	path = sanitizePath(path)
	route := r.mux.HandleFunc(path, h).Methods("OPTIONS")
	return &HttpRoute{r.api, route}
}

func (r *HttpRouter) Group(path string, fn func(IRouter.IHttpRouter)) {
	if fn == nil {
		panic(fmt.Sprintf("plugin-router: attempting to Route() a nil subrouter on '%s'", path))
	}
	path = sanitizePath(path)
	router := r.mux.PathPrefix(path).Subrouter()
	newrouter := HttpRouter{api: r.api, mux: router}
	fn(&newrouter)
}

func (r *HttpRouter) Use(middlewares ...func(http.Handler) http.Handler) {
	for _, mw := range middlewares {
		r.mux.Use(mw)
	}
}

func sanitizePath(path string) string {
	if !strings.HasPrefix(path, "/") {
		return "/" + path
	}
	return path
}
