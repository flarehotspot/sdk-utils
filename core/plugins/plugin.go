package plugins

import (
	"github.com/flarehotspot/core/config/plugincfg"
	"github.com/flarehotspot/core/connmgr"
	"github.com/flarehotspot/core/db"
)

type PluginApi struct {
	info             *plugincfg.PluginInfo
	dir              string
	db               *db.Database
	models           *PluginModels
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
	ClntMgr          *connmgr.ClientMgr
	UciAPI           *UciApi
	Utl              *PluginUtils
}
