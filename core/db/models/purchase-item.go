package models

import (
	"time"

	"github.com/flarehotspot/core/db"
)

type PurchaseItem struct {
	db          *db.Database
	id          int64
	purchaseId  int64
	sku         string
	name        string
	description string
	price       float64
	createdAt   time.Time
}

func NewPurchaseItem(dtb *db.Database) *PurchaseItem {
	return &PurchaseItem{
		db: dtb,
	}
}

func (self *PurchaseItem) Id() int64 {
	return self.id
}

func (self *PurchaseItem) PurchaseId() int64 {
	return self.purchaseId
}

func (self *PurchaseItem) Sku() string {
	return self.sku
}

func (self *PurchaseItem) Name() string {
	return self.name
}

func (self *PurchaseItem) Description() string {
	return self.description
}

func (self *PurchaseItem) Price() float64 {
	return self.price
}

func (self *PurchaseItem) CreatedAt() time.Time {
	return self.createdAt
}
