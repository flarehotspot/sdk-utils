package plugins

import (
	"fmt"
	"net/http"
	"strings"

	IRouter "github.com/flarehotspot/core/sdk/api/http/router"
	"github.com/gorilla/mux"
)

type PluginRouter struct {
	api *PluginApi
	mux *mux.Router
}

func (r *PluginRouter) Router() *mux.Router {
	return r.mux
}

func (r *PluginRouter) Get(path string, h http.HandlerFunc) IRouter.IRoute {
	path = sanitizePath(path)
	route := r.mux.HandleFunc(path, h).Methods("GET")
	return &PluginRoute{r.api, route}
}
func (r *PluginRouter) Post(path string, h http.HandlerFunc) IRouter.IRoute {
	path = sanitizePath(path)
	route := r.mux.HandleFunc(path, h).Methods("POST")
	return &PluginRoute{r.api, route}
}

func (r *PluginRouter) Put(path string, h http.HandlerFunc) IRouter.IRoute {
	path = sanitizePath(path)
	route := r.mux.HandleFunc(path, h).Methods("PUT")
	return &PluginRoute{r.api, route}
}

func (r *PluginRouter) Delete(path string, h http.HandlerFunc) IRouter.IRoute {
	path = sanitizePath(path)
	route := r.mux.HandleFunc(path, h).Methods("DELETE")
	return &PluginRoute{r.api, route}
}

func (r *PluginRouter) Options(path string, h http.HandlerFunc) IRouter.IRoute {
	path = sanitizePath(path)
	route := r.mux.HandleFunc(path, h).Methods("OPTIONS")
	return &PluginRoute{r.api, route}
}

func (r *PluginRouter) Group(path string, fn func(IRouter.IRouter)) {
	if fn == nil {
		panic(fmt.Sprintf("plugin-router: attempting to Route() a nil subrouter on '%s'", path))
	}
	path = sanitizePath(path)
	router := r.mux.PathPrefix(path).Subrouter()
	newrouter := PluginRouter{api: r.api, mux: router}
	fn(&newrouter)
}

func (r *PluginRouter) Use(middlewares ...func(http.Handler) http.Handler) {
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
