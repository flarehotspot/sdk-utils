package routes

// import (
// 	"github.com/flarehotspot/core/globals"
// 	"github.com/flarehotspot/core/web/controllers/adminctrl"
// 	"github.com/flarehotspot/core/web/middlewares"
// 	"github.com/flarehotspot/core/web/router"
// )

// func AdminRoutes(g *globals.CoreGlobals) {
// 	adminAssetsCtrl := adminctrl.NewAdminAssetsCtrl(g)
// 	apiR := router.ApiRouter().PathPrefix("/v1/admin").Subrouter()
// 	apiR.Use(middlewares.AdminAuth)

// 	// portal assets subpath
// 	corePkg := g.CoreApi.Pkg()
// 	assetsR := router.AssetsRouter()
// 	assetsR.HandleFunc("/"+corePkg+"/web-admin/main.js", adminAssetsCtrl.MainJs).Methods("GET")
// 	assetsR.HandleFunc("/{pkg}/web-admin/helpers.js", adminAssetsCtrl.HelpersJs(g)).Methods("GET")
// }
