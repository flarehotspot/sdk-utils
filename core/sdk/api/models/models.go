package sdkmdls

// IModels is the database models API.
type IModels interface {

	// Returns the devices model.
	Device() IDeviceModel

	// Returns the sessions model.
	Session() ISessionModel

	// Returns the purchases model.
	Purchase() IPurchaseModel

	// Returns the payments model.
	Payment() IPaymentModel

	// Returns the wallets model.
	Wallet() IWalletModel

	// Returns the wallet transactions model.
	WalletTrns() IWalletTrnsModel
}
