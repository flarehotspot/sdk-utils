package plugins

import (
	"github.com/flarehotspot/flarehotspot/core/config/plugincfg"
	"github.com/flarehotspot/flarehotspot/core/connmgr"
	"github.com/flarehotspot/flarehotspot/core/db"
	"github.com/flarehotspot/flarehotspot/core/db/models"
)

type PluginApi struct {
	info             *plugincfg.PluginInfo
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
