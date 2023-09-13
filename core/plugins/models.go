package plugins

import (
	coreM "github.com/flarehotspot/core/db/models"
	"github.com/flarehotspot/core/sdk/api/models"
)

type PluginModels struct {
	models *coreM.Models
}

func (self *PluginModels) Device() models.IDeviceModel {
	return self.models.Device()
}

func (self *PluginModels) Session() models.ISessionModel {
	return self.models.Session()
}

func (self *PluginModels) Purchase() models.IPurchaseModel {
	return self.models.Purchase()
}

func (self *PluginModels) PurchaseItem() models.IPurchaseItemModel {
	return self.models.PurchaseItem()
}

func (self *PluginModels) Payment() models.IPaymentModel {
	return self.models.Payment()
}

func (self *PluginModels) Wallet() models.IWalletModel {
	return self.models.Wallet()
}

func (self *PluginModels) WalletTrns() models.IWalletTrnsModel {
	return self.models.WalletTrns()
}

func NewPluginModels(m *coreM.Models) *PluginModels {
	return &PluginModels{m}
}
