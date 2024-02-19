package plugins

import (
	"github.com/flarehotspot/core/internal/connmgr"
	"github.com/flarehotspot/core/internal/db"
	"github.com/flarehotspot/core/internal/db/models"
	sdkplugin "github.com/flarehotspot/core/sdk/api/plugin"
)

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
