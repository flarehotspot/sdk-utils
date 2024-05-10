package controllers

import (
	"net/http"

	"github.com/flarehotspot/core/internal/plugins"
	"github.com/flarehotspot/core/internal/web/helpers"
)

func PaymentOptionsCtrl(g *plugins.CoreGlobals) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := g.CoreAPI.HttpAPI.VueResponse()
		clnt, err := helpers.CurrentClient(g.ClientRegister, r)
		if err != nil {
			res.SendFlashMsg(w, "error", err.Error(), http.StatusInternalServerError)
			return
		}

		methods := map[string]string{}
		for _, opt := range g.PaymentsMgr.Options(clnt) {
			methods[opt.Opt.OptName] = opt.VueRoutePath
		}

		res.Json(w, methods, http.StatusOK)

	}
}
