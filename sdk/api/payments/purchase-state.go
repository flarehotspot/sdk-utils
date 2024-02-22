package sdkpayments

type PurchaseState struct {
	TotalPayment       float64 `json:"total_payment"`
	WalletDebit        float64 `json:"wallet_debit"`
	WalletEndingBal    float64 `json:"wallet_ending_bal"`
	WalletRealBal      float64 `json:"wallet_real_bal"`
}
