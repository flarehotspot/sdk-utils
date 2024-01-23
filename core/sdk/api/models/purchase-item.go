package sdkmodels

// IPurchaseItem represents a purchase item record in the purchase_items table in the database.
type IPurchaseItem interface {

	// Returns the purchase item ID.
	Id() int64

	// Returns the purchase ID the item belongs to.
	PurchaseId() int64

	// Returns the product SKU.
	Sku() string

	// Returns the product name.
	Name() string

	// Returns the product description.
	Description() string

	// Returns the product price.
	Price() float64
}
