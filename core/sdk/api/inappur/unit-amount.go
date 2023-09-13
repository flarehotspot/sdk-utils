package inappur

// UnitAmount represents the payment amount and currency of items.
type UnitAmount struct {
	// Currency of the amount. Examples: USD, PHP, CNY.
	CurrencyCode string `query:"code"`

	// Numerical value of the amount.
	Value float64 `query:"val"`
}
