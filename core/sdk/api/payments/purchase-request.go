package sdkpayments

// PurchaseRequest represents a purchase to be made by the customer.
type PurchaseRequest struct {
	Sku                  string
	Name                 string
	Description          string
	Price                float64
	AnyPrice             bool
	CallbackVueRouteName string
}
