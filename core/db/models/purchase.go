package models

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/flarehotspot/core/db"
	models "github.com/flarehotspot/core/sdk/api/models"
)

func NewPurchase(dtb *db.Database, mdls *Models) *Purchase {
	return &Purchase{
		db:     dtb,
		models: mdls,
	}
}

type Purchase struct {
	db                   *db.Database
	models               *Models
	id                   int64
	deviceId             int64
	token                string
	sku                  string
	name                 string
	description          string
	price                float64
	anyPrice             bool
	callbackVueRouteName string
	walletDebit          float64
	walletTxId           *int64
	confirmedAt          *time.Time
	cancelledAt          *time.Time
	cancelledReason      *string
	createdAt            time.Time
}

func (self *Purchase) Id() int64 {
	return self.id
}

func (self *Purchase) DeviceId() int64 {
	return self.deviceId
}

func (self *Purchase) Token() string {
	return self.token
}

func (self *Purchase) Sku() string {
	return self.sku
}

func (self *Purchase) Name() string {
	return self.name
}

func (self *Purchase) Description() string {
	return self.description
}

func (self *Purchase) Price() float64 {
	return self.price
}

func (self *Purchase) AnyPrice() bool {
	return self.anyPrice
}

func (self *Purchase) WalletDebit() float64 {
	return self.walletDebit
}

func (self *Purchase) WalletTxId() *int64 {
	return self.walletTxId
}

func (self *Purchase) ConfirmedAt() *time.Time {
	return self.confirmedAt
}

func (self *Purchase) CancelledAt() *time.Time {
	return self.cancelledAt
}

func (self *Purchase) CancelledReason() *string {
	return self.cancelledReason
}

func (self *Purchase) CreatedAt() time.Time {
	return self.createdAt
}

func (self *Purchase) CallbackUrl() string {
	return self.callbackVueRouteName
}

func (self *Purchase) IsConfirmed() bool {
	return self.confirmedAt != nil
}

func (self *Purchase) IsCancelled() bool {
	return self.confirmedAt != nil
}

func (self *Purchase) IsProcessed() bool {
	return self.IsCancelled() || self.IsConfirmed()
}

func (self *Purchase) DeviceTx(tx *sql.Tx, ctx context.Context) (models.IDevice, error) {
	dev, err := self.models.deviceModel.FindTx(tx, ctx, self.deviceId)
	return dev, err
}

func (self *Purchase) ConfirmTx(tx *sql.Tx, ctx context.Context) error {
	dev, err := self.DeviceTx(tx, ctx)
	if err != nil {
		log.Println(err)
		return err
	}

	wallet, err := dev.WalletTx(tx, ctx)
	if err != nil {
		log.Println(err)
		return err
	}

	var txid *int64
	dbt := self.walletDebit
	if dbt > 0 {
		newBal := wallet.Balance() - dbt
		err = wallet.UpdateTx(tx, ctx, newBal)
		if err != nil {
			return nil
		}

		desc := "Partial payment for " + self.description
		trns, err := self.models.walletTrnsModel.CreateTx(tx, ctx, wallet.Id(), -dbt, newBal, desc)
		if err != nil {
			return err
		}

		id := trns.Id()
		txid = &id
	}

	now := time.Now()
	return self.UpdateTx(tx, ctx, dbt, txid, nil, &now, nil)
}

func (self *Purchase) CancelTx(tx *sql.Tx, ctx context.Context) error {
	dev, err := self.DeviceTx(tx, ctx)
	if err != nil {
		log.Println(err)
		return err
	}

	pmtTotal, err := self.PaymentsTotalTx(tx, ctx)
	if err != nil {
		return err
	}

	desc := "Cancelled purchase: " + self.description
	dbt := self.walletDebit
	cancelledAt := time.Now()

	if pmtTotal > 0 {
		wallet, err := dev.WalletTx(tx, ctx)
		if err != nil {
			log.Println(err)
			return err
		}

		err = wallet.IncBalanceTx(tx, ctx, pmtTotal)
		if err != nil {
			log.Println("Error updating wallet balance: ", err)
			return err
		}

		trns, err := self.models.WalletTrns().CreateTx(tx, ctx, wallet.Id(), pmtTotal, wallet.Balance(), "Refund for "+desc)
		if err != nil {
			log.Println(err)
			return err
		}

		trnsId := trns.Id()
		return self.UpdateTx(tx, ctx, dbt, &trnsId, &cancelledAt, nil, &desc)
	}

	return self.UpdateTx(tx, ctx, dbt, nil, &cancelledAt, nil, &desc)
}

func (self *Purchase) AddPaymentTx(tx *sql.Tx, ctx context.Context, amount float64, mtd string) (models.IPayment, error) {
	pmnt, err := self.models.paymentModel.CreateTx(tx, ctx, self.id, amount, mtd)
	return pmnt, err
}

func (self *Purchase) PaymentsTx(tx *sql.Tx, ctx context.Context) ([]models.IPayment, error) {
	return self.models.paymentModel.FindAllByPurchaseTx(tx, ctx, self.id)
}

func (self *Purchase) PaymentsTotalTx(tx *sql.Tx, ctx context.Context) (float64, error) {
	pmts, err := self.PaymentsTx(tx, ctx)
	if err != nil {
		return 0, err
	}

	var total float64

	for _, p := range pmts {
		total += p.Amount()
	}

	return total, nil
}

func (self *Purchase) StatTx(tx *sql.Tx, ctx context.Context) (*models.PurchaseStat, error) {
	device, err := self.DeviceTx(tx, ctx)
	if err != nil {
		return nil, err
	}

	wallet, err := device.WalletTx(tx, ctx)
	if err != nil {
		return nil, err
	}

	availBal, err := wallet.AvailableBalTx(tx, ctx)
	if err != nil {
		return nil, err
	}

	total, err := self.PaymentsTotalTx(tx, ctx)
	if err != nil {
		return nil, err
	}

	return &models.PurchaseStat{
		PaymentTotal:   total + self.WalletDebit(),
		WalletDebit:    self.WalletDebit(),
		WalletBal:      wallet.Balance(),
		WalletAvailBal: availBal,
	}, nil
}

func (self *Purchase) UpdateTx(tx *sql.Tx, ctx context.Context, dbt float64, txid *int64, cancelledAt *time.Time, confirmedAt *time.Time, reason *string) error {
	err := self.models.purchaseModel.UpdateTx(tx, ctx, self.id, dbt, txid, cancelledAt, confirmedAt, reason)
	if err != nil {
		return err
	}

	self.walletDebit = dbt
	self.walletTxId = txid
	self.cancelledAt = cancelledAt
	self.confirmedAt = confirmedAt
	return nil
}

func (self *Purchase) Cancel(ctx context.Context) error {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer tx.Rollback()

	err = self.CancelTx(tx, ctx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (self *Purchase) Confirm(ctx context.Context) error {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = self.ConfirmTx(tx, ctx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (self *Purchase) AddPayment(ctx context.Context, amt float64, mtd string) (models.IPayment, error) {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	pmt, err := self.AddPaymentTx(tx, ctx, amt, mtd)
	if err != nil {
		return nil, err
	}

	return pmt, tx.Commit()
}

func (self *Purchase) Payments(ctx context.Context) ([]models.IPayment, error) {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	payments, err := self.PaymentsTx(tx, ctx)
	if err != nil {
		return nil, err
	}

	return payments, tx.Commit()
}

func (self *Purchase) PaymentsTotal(ctx context.Context) (float64, error) {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	total, err := self.PaymentsTotalTx(tx, ctx)
	if err != nil {
		return 0, err
	}

	return total, tx.Commit()
}

func (self *Purchase) Stat(ctx context.Context) (*models.PurchaseStat, error) {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	stat, err := self.StatTx(tx, ctx)
	if err != nil {
		return nil, err
	}

	return stat, tx.Commit()
}

func (self *Purchase) Update(ctx context.Context, dbt float64, txid *int64, cancelledAt *time.Time, confirmedAt *time.Time, reason *string) error {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = self.UpdateTx(tx, ctx, dbt, txid, cancelledAt, confirmedAt, reason)
	if err != nil {
		return err
	}

	return tx.Commit()
}
