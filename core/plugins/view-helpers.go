package plugins

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/flarehotspot/core/accounts"
	"github.com/flarehotspot/core/config/appcfg"
	sdkacct "github.com/flarehotspot/core/sdk/api/accounts"
	"github.com/flarehotspot/core/sdk/api/connmgr"
	Irtr "github.com/flarehotspot/core/sdk/api/http/router"
	"github.com/flarehotspot/core/sdk/api/http/views"
	"github.com/flarehotspot/core/sdk/api/plugin"
	"github.com/flarehotspot/core/sdk/utils/flash"
	"github.com/flarehotspot/core/sdk/utils/translate"
	"github.com/flarehotspot/core/web/helpers"
	"github.com/flarehotspot/core/web/router"
)

var (
	flashTypes = []string{string(flash.Info), string(flash.Success), string(flash.Warning), string(flash.Error)}
)

type ViewHelpers struct {
	w        http.ResponseWriter
	r        *http.Request
	api      *PluginApi
	flashMsg map[string]string
}

func NewViewHelpers(api *PluginApi, w http.ResponseWriter, r *http.Request) views.IViewHelpers {
	h := &ViewHelpers{
		w:        w,
		r:        r,
		api:      api,
		flashMsg: make(map[string]string),
	}

	// Cache flash messages and write cookie to w first
	// because we'ere unable to write cookie once the view is executed.
	for _, t := range flashTypes {
		msg := flash.GetFlashMsg(w, r, t)
		h.flashMsg[t] = msg
	}

	return h
}

func (h *ViewHelpers) Translate(msgtype string, msgk string) string {
	return h.api.Translate(translate.MsgType(msgtype), msgk)
}

func (h *ViewHelpers) PluginMgr() plugin.IPluginMgr {
	return h.api.PluginsMgr
}

// func (h *ViewHelpers) GetAdminNavs() []*navigation.AdminNavList {
// 	return GetAdminNavs(h.api.PluginsMgr, h.r)
// }

func (h *ViewHelpers) AssetPath(path string) string {
	cfg, _ := appcfg.Read()
	return filepath.Join("/assets", cfg.AssetsVersion, h.api.Pkg(), path)
}

func (h *ViewHelpers) FlashMsgHtml() string {
	var s strings.Builder
	for _, t := range flashTypes {
		klass := t
		if t == "error" {
			klass = "danger"
		}

		if msg, ok := h.flashMsg[t]; ok && msg != "" {
			log.Println("GET FLASH MSG: ", msg)
			s.WriteString(fmt.Sprintf(`
      <div class="flash-msg alert alert-%s alert-dismissible fade show">
      %s
      <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
      </div>`, klass, msg))
		}
	}

	return s.String()
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
	acct, err := helpers.CurrentUser(h.r)
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
	acct, err := helpers.CurrentUser(h.r)
	if err != nil {
		return false
	}

	return accounts.HasAnyPerm(acct, perms...)
}

func (h *ViewHelpers) AdminHasAllPerms(perms ...string) bool {
	acct, err := helpers.CurrentUser(h.r)
	if err != nil {
		return false
	}
	return accounts.HasAllPerms(acct, perms...)
}
