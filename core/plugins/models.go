package plugins

import (
	"github.com/flarehotspot/core/db/models"
	"github.com/flarehotspot/core/sdk/api/models"
)

type PluginModels struct {
	models *models.Models
}

func (self *PluginModels) Device() sdkmodels.IDeviceModel {
	return self.models.Device()
}

func (self *PluginModels) Session() sdkmodels.ISessionModel {
	return self.models.Session()
}

func (self *PluginModels) Purchase() sdkmodels.IPurchaseModel {
	return self.models.Purchase()
}

func (self *PluginModels) PurchaseItem() sdkmodels.IPurchaseItemModel {
	return self.models.PurchaseItem()
}

func (self *PluginModels) Payment() sdkmodels.IPaymentModel {
	return self.models.Payment()
}

func (self *PluginModels) Wallet() sdkmodels.IWalletModel {
	return self.models.Wallet()
}

func (self *PluginModels) WalletTrns() sdkmodels.IWalletTrnsModel {
	return self.models.WalletTrns()
}

func NewPluginModels(m *models.Models) *PluginModels {
	return &PluginModels{m}
}
