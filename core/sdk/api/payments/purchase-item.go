package sdkpayments

// PurchaseItem represents a purchase item to be included in the purchase.
type PurchaseItem struct {
	Sku         string  `query:"s"`
	Name        string  `query:"n"`
	Description string  `query:"d"`
	Price       float64 `query:"p"`
}
