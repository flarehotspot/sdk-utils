package plugins

import (
	"github.com/flarehotspot/core/db/models"
	"github.com/flarehotspot/core/sdk/api/models"
)

type PluginModels struct {
	models *models.Models
}

func (self *PluginModels) Device() sdkmdls.IDeviceModel {
	return self.models.Device()
}

func (self *PluginModels) Session() sdkmdls.ISessionModel {
	return self.models.Session()
}

func (self *PluginModels) Purchase() sdkmdls.IPurchaseModel {
	return self.models.Purchase()
}

func (self *PluginModels) Payment() sdkmdls.IPaymentModel {
	return self.models.Payment()
}

func (self *PluginModels) Wallet() sdkmdls.IWalletModel {
	return self.models.Wallet()
}

func (self *PluginModels) WalletTrns() sdkmdls.IWalletTrnsModel {
	return self.models.WalletTrns()
}

func NewPluginModels(m *models.Models) *PluginModels {
	return &PluginModels{m}
}
