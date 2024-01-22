package sdkplugin

import (
	"database/sql"

	sdkacct "github.com/flarehotspot/core/sdk/api/accounts"
	sdkads "github.com/flarehotspot/core/sdk/api/ads"
	sdkcfg "github.com/flarehotspot/core/sdk/api/config"
	sdkconnmgr "github.com/flarehotspot/core/sdk/api/connmgr"
	sdkhttp "github.com/flarehotspot/core/sdk/api/http"
	sdkinappur "github.com/flarehotspot/core/sdk/api/inappur"
	sdkmodels "github.com/flarehotspot/core/sdk/api/models"
	sdknet "github.com/flarehotspot/core/sdk/api/network"
	sdkpayments "github.com/flarehotspot/core/sdk/api/payments"
	sdktheme "github.com/flarehotspot/core/sdk/api/themes"
	sdkuci "github.com/flarehotspot/core/sdk/api/uci"
)

// IPluginApi is the root of all plugin APIs.
type IPluginApi interface {

	// Returns the name of the plugin as defined in package.yml "name" field.
	Name() string

	// Returns the package name of the plugin as defined in package.yml "package" field.
	Pkg() string

	// Returns the version of the plugin as defined in package.yml "version" field.
	Version() string

	// Returns the slug name of the plugin
	Slug() string

	// Returns the description of plugin.
	Description() string

	// Returns the root directory of the plugin's installation path.
	Dir() string

	// Returns an instance of database/sql package from go standard library.
	DbApi() *sql.DB

	// Returns an instance of models api.
	ModelsApi() sdkmodels.IModelsApi

	// Returns an instance of accounts api.
	AcctApi() sdkacct.IAcctApi

	// Returns an instance of http api.
	HttpApi() sdkhttp.IHttpApi

	// Returns an instance of config api.
	ConfigApi() sdkcfg.IConfigApi

	// Returns an instance of payments api.
	PaymentsApi() sdkpayments.IPaymentsApi

	// Returns an instance of network api.
	NetworkApi() sdknet.INetworkApi

	// Returns an instance of ads api.
	AdsApi() sdkads.IAdsApi

	// Returns an instance of in-app purchase api.
	InAppPurchaseApi() sdkinappur.InAppPurchaseApi

	// Returns an instance of the plugin manager.
	PluginMgr() IPluginMgr

	// Returns an instance of the client register.
	ClientReg() sdkconnmgr.IClientRegister

	// Returns an instance of the client manager.
	ClientMgr() sdkconnmgr.IClientMgr

	// Returns an instance of the uci api.
	UciApi() sdkuci.IUciApi

	ThemesApi() sdktheme.IThemesApi

	Utils() IPluginUtils
}
