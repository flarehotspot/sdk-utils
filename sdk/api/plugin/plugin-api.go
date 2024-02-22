package sdkplugin

import (
	"database/sql"

	sdkacct "github.com/flarehotspot/sdk/api/accounts"
	sdkads "github.com/flarehotspot/sdk/api/ads"
	sdkcfg "github.com/flarehotspot/sdk/api/config"
	sdkconnmgr "github.com/flarehotspot/sdk/api/connmgr"
	sdkhttp "github.com/flarehotspot/sdk/api/http"
	sdkinappur "github.com/flarehotspot/sdk/api/inappur"
	sdknet "github.com/flarehotspot/sdk/api/network"
	sdkpayments "github.com/flarehotspot/sdk/api/payments"
	sdktheme "github.com/flarehotspot/sdk/api/themes"
	sdkuci "github.com/flarehotspot/sdk/api/uci"
)

// PluginApi is the root of all plugin APIs.
type PluginApi interface {

	// Returns the name of the plugin as defined in package.yml "name" field.
	Name() string

	// Returns the package name of the plugin as defined in package.yml "package" field.
	Pkg() string

	// Returns the version of the plugin as defined in package.yml "version" field.
	Version() string

	// Returns the description of plugin.
	Description() string

	// Returns the root directory of the plugin's installation path.
	Dir() string

	// Translate a message to the user's language.
	Translate(t string, msgk string, pairs ...any) string

	// Returns the absolute path to the given file in /resources folder of your plugin.
	// For example, if you have the following code:
	//  api.Utils().Resource("some-file.txt")
	// then it will return the absolute path to the file "[plugin_root_dir]/resources/some-file.txt" under the plugin's root directory.
	Resource(f string) (path string)

	// Returns an instance of database/sql package from go standard library.
	SqlDb() *sql.DB

	// Returns an instance of accounts api.
	Acct() sdkacct.AccountsApi

	// Returns an instance of http api.
	Http() sdkhttp.HttpApi

	// Returns an instance of config api.
	Config() sdkcfg.ConfigApi

	// Returns an instance of payments api.
	Payments() sdkpayments.PaymentsApi

	// Returns an instance of network api.
	Network() sdknet.Network

	// Returns an instance of ads api.
	Ads() sdkads.AdsApi

	// Returns an instance of in-app purchase api.
	InAppPurchases() sdkinappur.InAppPurchasesApi

	// Returns an instance of the plugin manager.
	PluginsMgr() PluginsMgrApi

	// Returns an instance of the client register.
	DeviceHooks() sdkconnmgr.DeviceHooksApi

	// Returns an instance of the client manager.
	SessionsMgr() sdkconnmgr.SessionsMgr

	// Returns an instance of the uci api.
	Uci() sdkuci.UciApi

	Themes() sdktheme.ThemesApi
}
