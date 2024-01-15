package router

import "github.com/gorilla/mux"

const (
	NotFoundRoute string = "/404"
)

var (
	RootRouter    *mux.Router
	BootingRouter *mux.Router
	PluginRouter  *mux.Router
)
