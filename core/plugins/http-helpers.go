package plugins

import (
	nethttp "net/http"
	"strings"

	"github.com/flarehotspot/core/accounts"
	sdkacct "github.com/flarehotspot/core/sdk/api/accounts"
	"github.com/flarehotspot/core/sdk/api/connmgr"
	"github.com/flarehotspot/core/sdk/api/http"
	Irtr "github.com/flarehotspot/core/sdk/api/http/router"
	"github.com/flarehotspot/core/sdk/api/plugin"
	"github.com/flarehotspot/core/sdk/utils/translate"
	"github.com/flarehotspot/core/web/helpers"
	"github.com/flarehotspot/core/web/router"
)

type ViewHelpers struct {
	w   nethttp.ResponseWriter
	r   *nethttp.Request
	api *PluginApi
}

func NewViewHelpers(api *PluginApi, w nethttp.ResponseWriter, r *nethttp.Request) http.IHelpers {
	return &ViewHelpers{
		w:   w,
		r:   r,
		api: api,
	}
}

func (h *ViewHelpers) Translate(msgtype string, msgk string) string {
	return h.api.Translate(translate.MsgType(msgtype), msgk)
}

func (h *ViewHelpers) PluginMgr() plugin.IPluginMgr {
	return h.api.PluginsMgr
}

func (h *ViewHelpers) AssetPath(path string) string {
	return h.api.HttpAPI.AssetPath(path)
}

func (h *ViewHelpers) AdView() (html string) {
	return ""
}

func (h *ViewHelpers) MuxRouteName(name string) Irtr.MuxRouteName {
	return h.api.HttpAPI.HttpRouter().MuxRouteName(Irtr.PluginRouteName(name))
}

func (h *ViewHelpers) UrlForMuxRoute(name string, params ...string) string {
	url, _ := router.UrlForRoute(Irtr.MuxRouteName(name), params...)
	return url
}

func (h *ViewHelpers) UrlForRoute(name string, params ...string) string {
	return h.api.HttpApi().HttpRouter().UrlForRoute(Irtr.PluginRouteName(name), params...)
}

func (h *ViewHelpers) IsLinkActive(href string) bool {
	curr := h.r.URL.String()
	return strings.HasPrefix(curr, href)
}

func (h *ViewHelpers) CurrentUser() sdkacct.IAccount {
	acct, err := helpers.CurrentAdmin(h.r)
	if err != nil {
		return nil
	}
	return acct
}

func (h *ViewHelpers) CurrentClient() connmgr.IClientDevice {
	clnt, err := helpers.CurrentClient(h.r)
	if err != nil {
		return nil
	}
	return clnt
}

func (h *ViewHelpers) AdminHasAnyPerm(perms ...string) bool {
	acct, err := helpers.CurrentAdmin(h.r)
	if err != nil {
		return false
	}

	return accounts.HasAnyPerm(acct, perms...)
}

func (h *ViewHelpers) AdminHasAllPerms(perms ...string) bool {
	acct, err := helpers.CurrentAdmin(h.r)
	if err != nil {
		return false
	}
	return accounts.HasAllPerms(acct, perms...)
}
