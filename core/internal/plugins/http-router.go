package plugins

import (
	"net/http"

	sdkhttp "github.com/flarehotspot/sdk/api/http"
	"github.com/gorilla/mux"
)

func NewHttpRouter(api *PluginApi, mux *mux.Router) *HttpRouter {
	return &HttpRouter{api, mux}
}

type HttpRouter struct {
	api *PluginApi
	mux *mux.Router
}

func (self *HttpRouter) Router() *mux.Router {
	return self.mux
}

func (self *HttpRouter) Get(path string, h http.HandlerFunc, mw ...func(next http.Handler) http.Handler) sdkhttp.HttpRoute {
	path = self.api.HttpAPI.vueRouter.VuePathToMuxPath(path)
	finalHandler := http.Handler(h)
	for i := len(mw) - 1; i >= 0; i-- {
		finalHandler = mw[i](finalHandler)
	}
	route := self.mux.Handle(path, finalHandler).Methods("GET")
	return NewHttpRoute(self.api, route)
}

func (self *HttpRouter) Post(path string, h http.HandlerFunc, mw ...func(next http.Handler) http.Handler) sdkhttp.HttpRoute {
	path = self.api.HttpAPI.vueRouter.VuePathToMuxPath(path)
	finalHandler := http.Handler(h)
	for i := len(mw) - 1; i >= 0; i-- {
		finalHandler = mw[i](finalHandler)
	}
	route := self.mux.Handle(path, finalHandler).Methods("POST")
	return NewHttpRoute(self.api, route)
}

func (self *HttpRouter) Group(path string, fn func(sdkhttp.RouterInstance)) {
	path = self.api.HttpAPI.vueRouter.VuePathToMuxPath(path)
	router := self.mux.PathPrefix(path).Subrouter()
	newrouter := NewHttpRouter(self.api, router)
	fn(newrouter)
}

func (self *HttpRouter) Use(middlewares ...func(http.Handler) http.Handler) {
	for _, mw := range middlewares {
		self.mux.Use(mw)
	}
}
