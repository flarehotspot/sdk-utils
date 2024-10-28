/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package sdkplugin

import (
	"database/sql"

	sdkacct "sdk/api/accounts"
	sdkads "sdk/api/ads"
	sdkcfg "sdk/api/config"
	sdkconnmgr "sdk/api/connmgr"
	sdkhttp "sdk/api/http"
	sdkinappur "sdk/api/inappur"
	sdklogger "sdk/api/logger"
	sdknet "sdk/api/network"
	sdkpayments "sdk/api/payments"
	sdktheme "sdk/api/themes"
	sdkuci "sdk/api/uci"
)

// PluginApi is the root of all plugin APIs.
type PluginApi interface {

	// Returns the package name of the plugin as defined in package.yml "package" field.
	Pkg() string

	// Returns the name of the plugin as defined in package.yml "name" field.
	Name() string

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
	//  api.Resource("some-file.txt")
	// then it will return the absolute path to the file "[plugin_root_dir]/resources/some-file.txt" under the plugin's root directory.
	Resource(f string) (path string)

	// Returns an instance of database/sql package from go standard library.
	SqlDb() *sql.DB

    // Run the plugin migration scripts in resources/migrations folder.
    Migrate() error

	// Returns an instance of accounts api.
	Acct() sdkacct.AccountsApi

	// Returns an instance of http api.
	Http() sdkhttp.HttpApi

	// Returns an instance of config api.
	Config() sdkcfg.ConfigApi

	// Returns an instance of payments api.
	Payments() sdkpayments.PaymentsApi

	// Returns an instance of network api.
	Network() sdknet.NetworkApi

	// Returns an instance of ads api.
	Ads() sdkads.AdsApi

	// Returns an instance of in-app purchase api.
	InAppPurchases() sdkinappur.InAppPurchasesApi

	// Returns an instance of the plugin manager.
	PluginsMgr() PluginsMgrApi

	// Returns an instance of the client register.
	DeviceHooks() sdkconnmgr.DeviceHooksApi

	// Returns an instance of the client manager.
	SessionsMgr() sdkconnmgr.SessionsMgrApi

	// Returns an instance of the uci api.
	Uci() sdkuci.UciApi

	// Returns an instance of the themes api.
	Themes() sdktheme.ThemesApi

	// Features returns a slice of strings representing the features supported by the plugin.
	Features() []string

	Logger() sdklogger.LoggerApi
}
