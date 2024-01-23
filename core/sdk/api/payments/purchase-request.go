package sdkpayments

import (
	"context"

	models "github.com/flarehotspot/core/sdk/api/models"
	qs "github.com/flarehotspot/core/sdk/libs/urlquery"
)

// PurchaseRequest represents a purchase to be made by the customer.
type PurchaseRequest struct {
	Items       []*PurchaseItem `query:"i"`
	VarPrice    bool            `query:"v"`
	CallbackUrl string          `query:"c"`
}

// Returns the total price of the purchase. If VarPrice is true,
// then the price depends on the amount paid by the customer.
func (self *PurchaseRequest) TotalPrice() float64 {
	if self.VarPrice {
		return 0
	}

	var total float64
	for _, item := range self.Items {
		total += item.Price
	}
	return total
}

// Returns the query parameters for the purchase request.
func (self *PurchaseRequest) ToQueryParams() (string, error) {
	b, err := qs.Marshal(self)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// Returns a PurchaseRequest from the given IPurchase.
func FromPurchase(ctx context.Context, p models.IPurchase) (*PurchaseRequest, error) {
	items, err := p.PurchaseItems(ctx)
	if err != nil {
		return nil, err
	}

	pItems := []*PurchaseItem{}
	for _, i := range items {
		pItems = append(pItems, &PurchaseItem{
			Name:        i.Name(),
			Description: i.Description(),
			Price:       i.Price(),
		})
	}

	return &PurchaseRequest{
		Items:       pItems,
		VarPrice:    p.VarPrice(),
		CallbackUrl: p.CallbackUrl(),
	}, nil
}
