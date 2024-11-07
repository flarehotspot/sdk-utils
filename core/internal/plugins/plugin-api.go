package plugins

import (
	"database/sql"
	"log"
	"path/filepath"

	"core/internal/connmgr"
	"core/internal/db"
	"core/internal/db/models"
	"core/internal/network"
	"core/internal/utils/migrate"
	"core/internal/utils/pkg"
	sdkacct "sdk/api/accounts"
	sdkads "sdk/api/ads"
	sdkcfg "sdk/api/config"
	sdkconnmgr "sdk/api/connmgr"
	sdkhttp "sdk/api/http"
	sdkinappur "sdk/api/inappur"
	sdklogger "sdk/api/logger"
	sdknet "sdk/api/network"
	sdkpayments "sdk/api/payments"
	sdkplugin "sdk/api/plugin"
	sdkuci "sdk/api/uci"
)

func NewPluginApi(dir string, pmgr *PluginsMgr, trfkMgr *network.TrafficMgr) *PluginApi {
	pluginApi := &PluginApi{
		dir:           dir,
		db:            pmgr.db,
		PluginsMgrApi: pmgr,
		ClntReg:       pmgr.clntReg,
		ClntMgr:       pmgr.clntMgr,
	}

	pluginApi.Utl = NewPluginUtils(pluginApi)

	info, err := pkg.GetSrcInfo(dir)
	if err != nil {
		log.Println("Error getting plugin info: ", err.Error())
		return nil
	}

	pluginApi.info = &info
	pluginApi.models = pmgr.models

	NewAcctApi(pluginApi)
	NewHttpApi(pluginApi, pmgr.db, pmgr.clntReg, pmgr.models, pmgr.clntReg, pmgr.paymgr)
	NewConfigApi(pluginApi)
	NewPaymentsApi(pluginApi, pmgr.paymgr)
	NewThemesApi(pluginApi)
	NewNetworkApi(pluginApi, trfkMgr)
	NewAdsApi(pluginApi)
	NewInAppPurchaseApi(pluginApi)
	NewUciApi(pluginApi)
	NewLoggerApi(pluginApi)

	log.Println("NewPluginApi: ", dir, " - ", info.Package, " - ", info.Name, " - ", info.Version, " - ", info.Description)

	return pluginApi
}

type PluginApi struct {
	info             *sdkplugin.PluginInfo
	dir              string
	db               *db.Database
	models           *models.Models
	CoreAPI          *PluginApi
	AcctAPI          *AccountsApi
	HttpAPI          *HttpApi
	ConfigAPI        *ConfigApi
	PaymentsAPI      *PaymentsApi
	ThemesAPI        *HttpThemesApi
	NetworkAPI       *NetworkApi
	AdsAPI           *AdsApi
	InAppPurchaseAPI *InAppPurchaseApi
	PluginsMgrApi    *PluginsMgr
	ClntReg          *connmgr.ClientRegister
	ClntMgr          *connmgr.SessionsMgr
	UciAPI           *UciApi
	Utl              *PluginUtils
	LoggerAPI        *LoggerApi
	AssetsManifest   pkg.OutputManifest
}

func (self *PluginApi) InitCoreApi(coreApi *PluginApi) {
	self.CoreAPI = coreApi
}

func (self *PluginApi) Migrate() error {
	migdir := filepath.Join(self.dir, "resources/migrations")
	err := migrate.MigrateUp(self.db.SqlDB(), migdir)
	if err != nil {
		log.Println("Error in plugin migration "+self.Name(), ":", err.Error())
		return err
	}

	log.Println("Done migrating plugin:", self.Name())
	return nil
}

func (self *PluginApi) Name() string {
	return self.info.Name
}

func (self *PluginApi) Pkg() string {
	return self.info.Package
}

func (self *PluginApi) Version() string {
	return self.info.Version
}

func (self *PluginApi) Description() string {
	info, err := pkg.GetSrcInfo(self.dir)
	if err != nil {
		return ""
	}
	return info.Description
}

func (self *PluginApi) Dir() string {
	return self.dir
}

func (self *PluginApi) Translate(t string, msgk string, pairs ...interface{}) string {
	return self.Utl.Translate(t, msgk, pairs...)
}

func (self *PluginApi) Resource(f string) (path string) {
	return self.Utl.Resource(f)
}

func (self *PluginApi) SqlDb() *sql.DB {
	return self.db.SqlDB()
}

func (self *PluginApi) Acct() sdkacct.AccountsApi {
	return self.AcctAPI
}

func (self *PluginApi) Http() sdkhttp.HttpApi {
	return self.HttpAPI
}

func (self *PluginApi) Config() sdkcfg.ConfigApi {
	return self.ConfigAPI
}

func (self *PluginApi) Payments() sdkpayments.PaymentsApi {
	return self.PaymentsAPI
}

func (self *PluginApi) Ads() sdkads.AdsApi {
	return self.AdsAPI
}

func (self *PluginApi) InAppPurchases() sdkinappur.InAppPurchasesApi {
	return self.InAppPurchaseAPI
}

func (self *PluginApi) PluginsMgr() sdkplugin.PluginsMgrApi {
	return self.PluginsMgrApi
}

func (self *PluginApi) Network() sdknet.NetworkApi {
	return self.NetworkAPI
}

func (self *PluginApi) DeviceHooks() sdkconnmgr.DeviceHooksApi {
	return self.ClntReg
}

func (self *PluginApi) SessionsMgr() sdkconnmgr.SessionsMgrApi {
	return self.ClntMgr
}

func (self *PluginApi) Uci() sdkuci.UciApi {
	return self.UciAPI
}

func (self *PluginApi) Themes() sdkhttp.HttpThemesApi {
	return self.ThemesAPI
}

func (self *PluginApi) Features() []string {
	features := []string{}
	if self.ThemesAPI.AdminTheme != nil {
		features = append(features, "theme:admin")
	}
	if self.ThemesAPI.PortalTheme != nil {
		features = append(features, "theme:portal")
	}
	return features
}

func (self *PluginApi) Logger() sdklogger.LoggerApi {
	return self.LoggerAPI
}

func (self *PluginApi) LoadAssetsManifest() {
	self.AssetsManifest = pkg.GetAssetManifest(self.dir)
}
