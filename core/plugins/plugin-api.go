package plugins

import (
	"database/sql"
	"log"
	"path/filepath"

	"github.com/flarehotspot/flarehotspot/core/config/plugincfg"
	"github.com/flarehotspot/flarehotspot/core/network"
	acct "github.com/flarehotspot/flarehotspot/core/sdk/api/accounts"
	ads "github.com/flarehotspot/flarehotspot/core/sdk/api/ads"
	config "github.com/flarehotspot/flarehotspot/core/sdk/api/config"
	connmgr "github.com/flarehotspot/flarehotspot/core/sdk/api/connmgr"
	http "github.com/flarehotspot/flarehotspot/core/sdk/api/http"
	inappur "github.com/flarehotspot/flarehotspot/core/sdk/api/inappur"
	sdknet "github.com/flarehotspot/flarehotspot/core/sdk/api/network"
	paymentsApi "github.com/flarehotspot/flarehotspot/core/sdk/api/payments"
	plugin "github.com/flarehotspot/flarehotspot/core/sdk/api/plugin"
	themes "github.com/flarehotspot/flarehotspot/core/sdk/api/themes"
	uci "github.com/flarehotspot/flarehotspot/core/sdk/api/uci"
	"github.com/flarehotspot/flarehotspot/core/utils/migrate"
)

func (p *PluginApi) InitCoreApi(coreApi *PluginApi) {
	p.CoreAPI = coreApi
}

func (p *PluginApi) Migrate() error {
	migdir := filepath.Join(p.dir, "resources/migrations")
	err := migrate.MigrateUp(migdir, p.db.SqlDB())
	if err != nil {
		log.Println("Error in plugin migration "+p.Name(), ":", err.Error())
		return err
	}

	log.Println("Done migrating plugin:", p.Name())
	return nil
}

func (p *PluginApi) Name() string {
	return p.info.Name
}

func (p *PluginApi) Pkg() string {
	return p.info.Package
}

func (p *PluginApi) Version() string {
	return p.info.Version
}

func (p *PluginApi) Description() string {
	info, err := plugincfg.GetPluginInfo(p.dir)
	if err != nil {
		return ""
	}
	return info.Description
}

func (p *PluginApi) Dir() string {
	return p.dir
}

func (p *PluginApi) Translate(t string, msgk string, pairs ...any) string {
	return p.Utl.Translate(t, msgk, pairs...)
}

func (p *PluginApi) Resource(f string) (path string) {
	return p.Utl.Resource(f)
}

func (p *PluginApi) SqlDb() *sql.DB {
	return p.db.SqlDB()
}

func (p *PluginApi) Acct() acct.AccountsApi {
	return p.AcctAPI
}

func (p *PluginApi) Http() http.HttpApi {
	return p.HttpAPI
}

func (p *PluginApi) Config() config.ConfigApi {
	return p.ConfigAPI
}

func (p *PluginApi) Payments() paymentsApi.PaymentsApi {
	return p.PaymentsAPI
}

func (p *PluginApi) Ads() ads.AdsApi {
	return p.AdsAPI
}

func (p *PluginApi) InAppPurchases() inappur.InAppPurchasesApi {
	return p.InAppPurchaseAPI
}

func (p *PluginApi) PluginsMgr() plugin.PluginsMgrApi {
	return p.PluginsMgrApi
}

func (p *PluginApi) Network() sdknet.Network {
	return p.NetworkAPI
}

func (p *PluginApi) DeviceHooks() connmgr.DeviceHooksApi {
	return p.ClntReg
}

func (p *PluginApi) SessionsMgr() connmgr.SessionsMgr {
	return p.ClntMgr
}

func (p *PluginApi) Uci() uci.UciApi {
	return p.UciAPI
}

func (p *PluginApi) Themes() themes.ThemesApi {
	return p.ThemesAPI
}

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
