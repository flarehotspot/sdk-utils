package plugin

import (
	"database/sql"

	"github.com/flarehotspot/core/sdk/api/accounts"
	"github.com/flarehotspot/core/sdk/api/ads"
	"github.com/flarehotspot/core/sdk/api/config"
	"github.com/flarehotspot/core/sdk/api/connmgr"
	"github.com/flarehotspot/core/sdk/api/http"
	"github.com/flarehotspot/core/sdk/api/http/navigation"
	"github.com/flarehotspot/core/sdk/api/inappur"
	"github.com/flarehotspot/core/sdk/api/models"
	"github.com/flarehotspot/core/sdk/api/network"
	"github.com/flarehotspot/core/sdk/api/payments"
	"github.com/flarehotspot/core/sdk/api/themes"
	"github.com/flarehotspot/core/sdk/api/uci"
	"github.com/flarehotspot/core/sdk/utils/translate"
)

// IPluginApi is the root of all plugin APIs.
type IPluginApi interface {

	// Returns the name of the plugin as defined in package.yml "name" field.
	Name() string

	// Returns the version of the plugin as defined in package.yml "version" field.
	Version() string

	// Returns the description of plugin.
	Description() string

	// Returns the root directory of the plugin's installation path.
	Dir() string

	// Translates the given message key to the current language.
	// This is the same and is identical to the view helper's "Translate()" method.
	// For example, if the current language is "en", then the following code:
	//  api.Translate(translate.Error, "some-key")
	// will look for the file "/resources/translations/en/error/some-key.txt" under the plugin's root directory
	// and displays the text inside that file.
	Translate(t translate.MsgType, msgk string) string

	// Returns the absolute path to the given file in /resources folder of your plugin.
	// For example, if you have the following code:
	//  api.Resource("some-file.txt")
	// then it will return the absolute path to the file "/resources/some-file.txt" under the plugin's root directory.
	Resource(f string) (path string)

	// Returns an instance of database/sql package from go standard library.
	DbApi() *sql.DB

	// Returns an instance of models api.
	ModelsApi() models.IModelsApi

	// Returns an instance of accounts api.
	AcctApi() accounts.IAcctApi

	// Returns an instance of http api.
	HttpApi() http.IHttpApi

	// Returns an instance of navigation api.
	NavApi() navigation.INavApi

	// Returns an instance of config api.
	ConfigApi() config.IConfigApi

	// Returns an instance of payments api.
	PaymentsApi() payments.IPaymentsApi

	// Returns an instance of themes api.
	ThemesApi() themes.IThemesApi

	// Returns an instance of network api.
	NetworkApi() network.INetworkApi

	// Returns an instance of ads api.
	AdsApi() ads.IAdsApi

	// Returns an instance of in-app purchase api.
	InAppPurchaseApi() inappur.InAppPurchaseApi

	// Returns an instance of the plugin manager.
	PluginMgr() IPluginMgr

	// Returns an instance of the client register.
	ClientReg() connmgr.IClientRegister

	// Returns an instance of the client manager.
	ClientMgr() connmgr.IClientMgr

	// Returns an instance of the uci api.
	UciApi() uci.IUciApi
}
