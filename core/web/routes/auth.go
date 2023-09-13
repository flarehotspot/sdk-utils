package routes

import (
	"github.com/flarehotspot/core/globals"
	"github.com/flarehotspot/core/web/controllers"
	"github.com/flarehotspot/core/web/middlewares"
	"github.com/flarehotspot/core/web/router"
	"github.com/flarehotspot/core/web/routes/names"
)

func AuthRoutes(g *globals.CoreGlobals) {
	ctrl := controllers.NewAuthCtrl(g)
	authRouter := router.RootRouter().PathPrefix("/auth").Subrouter()
	loginRouter := authRouter.PathPrefix("/login").Subrouter()
	loginRouter.Use(middlewares.MustBeLoggedOut)
	loginRouter.HandleFunc("", ctrl.GetLogin).Methods("GET").Name(names.RouteAuthLogin)
	loginRouter.HandleFunc("", ctrl.PostLogin).Methods("POST").Name(names.RouteAuthLogin)

	logoutRouter := authRouter.PathPrefix("/logout").Subrouter()
	logoutRouter.HandleFunc("", ctrl.GetLogout).Methods("GET").Name(names.RouteAuthLogout)
}
