package views

import (
	"github.com/flarehotspot/core/sdk/api/accounts"
	"github.com/flarehotspot/core/sdk/api/connmgr"
	"github.com/flarehotspot/core/sdk/api/http/router"
)

// IViewHelpers are methods available in html templates as .Helpers.
// For example, to use the Translate() method in html templates, use {{ .Helpers.Translate "label" "network_settings" }}.
type IViewHelpers interface {

	// Translates a message into the current language settings from application config.
	// msgtype is the message type, e.g. "error", "success", "info", "warning".
	// For example, if the current language is "en", then the following code in your template:
	//  {{ .Helpers.Translate "error" "some-key" }}
	// will look for the file "/resources/translations/en/error/some-key.txt" under the plugin root directory
	// and displays the text inside that file.
	Translate(msgtype string, msgk string) string

	// Returns asset path prefixed with assets version/hash path
	AssetPath(path string) string

	// Returns the html for the flash message.
	// These are the messages set in flash.SetFlashMsg() inside your controllers.
	FlashMsgHtml() (html string)

	// Returns the html for the ads view.
	AdView() (html string)

	// Returns the muxnmame for the route name in your plugin.
	// "muxname" is a route name that can be used for the UrlForMuxRoute() method.
	MuxRouteName(name string) (muxname router.MuxRouteName)

	// Returns the url for the mux route.
	// The difference between UrlForMuxRoute() vs UrlForRoute() is that UrlForRoute() only accepts route names specific to your plugin.
	UrlForMuxRoute(muxname string, pairs ...string) (url string)

	// Returns the url for the route.
	// The difference between UrlForMuxRoute() vs UrlForRoute() is that UrlForMuxRoute() only accepts route names built-in to the core system.
	UrlForRoute(name string, pairs ...string) (url string)

	// Returns true if the link/url is active.
	IsLinkActive(href string) bool

	// Returns the current admin account user.
	CurrentUser() accounts.IAccount

	// Returns the current client device.
	CurrentClient() connmgr.IClientDevice

	// Returns true if the current admin user has any of the specified permissions.
	AdminHasAnyPerm(perms ...string) bool

	// Returns true if the current admin user has all of the specified permissions.
	AdminHasAllPerms(perms ...string) bool
}
