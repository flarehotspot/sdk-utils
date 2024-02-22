package plugins

import (
	"database/sql"
	"log"
	"path/filepath"

	"github.com/flarehotspot/core/internal/config/plugincfg"
	"github.com/flarehotspot/core/internal/connmgr"
	"github.com/flarehotspot/core/internal/db"
	"github.com/flarehotspot/core/internal/db/models"
	"github.com/flarehotspot/core/internal/network"
	"github.com/flarehotspot/core/internal/utils/migrate"
	sdkacct "github.com/flarehotspot/core/sdk/api/accounts"
	sdkads "github.com/flarehotspot/core/sdk/api/ads"
	sdkcfg "github.com/flarehotspot/core/sdk/api/config"
	sdkconnmgr "github.com/flarehotspot/core/sdk/api/connmgr"
	sdkhttp "github.com/flarehotspot/core/sdk/api/http"
	sdkinappur "github.com/flarehotspot/core/sdk/api/inappur"
	sdknet "github.com/flarehotspot/core/sdk/api/network"
	sdkpayments "github.com/flarehotspot/core/sdk/api/payments"
	sdkplugin "github.com/flarehotspot/core/sdk/api/plugin"
	sdkthemes "github.com/flarehotspot/core/sdk/api/themes"
	sdkuci "github.com/flarehotspot/core/sdk/api/uci"
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

	info, err := plugincfg.GetPluginInfo(dir)
	if err != nil {
		log.Println("Error getting plugin info: ", err.Error())
		return nil
	}

	pluginApi.info = info
	pluginApi.models = pmgr.models
	pluginApi.AcctAPI = NewAcctApi(pluginApi)
	pluginApi.HttpAPI = NewHttpApi(pluginApi, pmgr.db, pmgr.clntReg, pmgr.models, pmgr.clntReg, pmgr.paymgr)
	pluginApi.ConfigAPI = NewConfigApi(pluginApi)
	pluginApi.PaymentsAPI = NewPaymentsApi(pluginApi, pmgr.paymgr)
	pluginApi.ThemesAPI = NewThemesApi(pluginApi)
	pluginApi.NetworkAPI = NewNetworkApi(trfkMgr)
	pluginApi.AdsAPI = NewAdsApi(pluginApi)
	pluginApi.InAppPurchaseAPI = NewInAppPurchaseApi(pluginApi)
	pluginApi.UciAPI = NewUciApi()

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
	ThemesAPI        *ThemesApi
	NetworkAPI       *NetworkApi
	AdsAPI           *AdsApi
	InAppPurchaseAPI *InAppPurchaseApi
	PluginsMgrApi    *PluginsMgr
	ClntReg          *connmgr.ClientRegister
	ClntMgr          *connmgr.SessionsMgr
	UciAPI           *UciApi
	Utl              *PluginUtils
}

func (self *PluginApi) InitCoreApi(coreApi *PluginApi) {
	self.CoreAPI = coreApi
}

func (self *PluginApi) Migrate() error {
	migdir := filepath.Join(self.dir, "resources/migrations")
	err := migrate.MigrateUp(migdir, self.db.SqlDB())
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
	info, err := plugincfg.GetPluginInfo(self.dir)
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

func (self *PluginApi) Network() sdknet.Network {
	return self.NetworkAPI
}

func (self *PluginApi) DeviceHooks() sdkconnmgr.DeviceHooksApi {
	return self.ClntReg
}

func (self *PluginApi) SessionsMgr() sdkconnmgr.SessionsMgr {
	return self.ClntMgr
}

func (self *PluginApi) Uci() sdkuci.UciApi {
	return self.UciAPI
}

func (self *PluginApi) Themes() sdkthemes.ThemesApi {
	return self.ThemesAPI
}
