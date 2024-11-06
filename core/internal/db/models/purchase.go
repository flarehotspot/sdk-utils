package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"core/internal/db"

	"github.com/jackc/pgx/v5"
	// sdkpayments "sdk/api/payments"
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
	sku                  string
	name                 string
	description          string
	price                float64
	anyPrice             bool
	callbackPluginPkg    string
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

func (self *Purchase) CallbackPluginPkg() string {
	return self.callbackPluginPkg
}

func (self *Purchase) CallbackVueRouteName() string {
	return self.callbackVueRouteName
}

func (self *Purchase) IsConfirmed() bool {
	return self.confirmedAt != nil
}

func (self *Purchase) IsCancelled() bool {
	return self.confirmedAt != nil
}

func (self *Purchase) FixedPrice() (float64, bool) {
	return self.price, !self.anyPrice
}

func (self *Purchase) DeviceTx(tx pgx.Tx, ctx context.Context) (*Device, error) {
	dev, err := self.models.deviceModel.FindTx(tx, ctx, self.deviceId)
	return dev, err
}

func (self *Purchase) ConfirmTx(tx pgx.Tx, ctx context.Context) error {
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

func (self *Purchase) CancelTx(tx pgx.Tx, ctx context.Context) error {
	dev, err := self.DeviceTx(tx, ctx)
	if err != nil {
		log.Println(err)
		return err
	}

	pmtTotal, err := self.TotalPaymentsTx(tx, ctx)
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

		trns, err := self.models.walletTrnsModel.CreateTx(tx, ctx, wallet.Id(), pmtTotal, wallet.Balance(), "Refund for "+desc)
		if err != nil {
			log.Println(err)
			return err
		}

		trnsId := trns.Id()
		return self.UpdateTx(tx, ctx, dbt, &trnsId, &cancelledAt, nil, &desc)
	}

	return self.UpdateTx(tx, ctx, dbt, nil, &cancelledAt, nil, &desc)
}

func (self *Purchase) PaymentsTx(tx pgx.Tx, ctx context.Context) ([]*Payment, error) {
	return self.models.paymentModel.FindAllByPurchaseTx(tx, ctx, self.id)
}

func (self *Purchase) TotalPaymentsTx(tx pgx.Tx, ctx context.Context) (float64, error) {
	pmts, err := self.PaymentsTx(tx, ctx)
	if err != nil {
		return 0, err
	}

	var total float64

	for _, p := range pmts {
		total += p.Amount()
	}

	total += self.WalletDebit()

	return total, nil
}

func (self *Purchase) UpdateTx(tx pgx.Tx, ctx context.Context, dbt float64, txid *int64, cancelledAt *time.Time, confirmedAt *time.Time, reason *string) error {
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
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("could not begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	err = self.CancelTx(tx, ctx)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (self *Purchase) Confirm(ctx context.Context) error {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	err = self.ConfirmTx(tx, ctx)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	return nil
}

func (self *Purchase) TotalPayment(ctx context.Context) (float64, error) {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("could not begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	total, err := self.TotalPaymentsTx(tx, ctx)
	if err != nil {
		return 0, err
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, fmt.Errorf("could not commit transaction: %w", err)
	}

	return total, nil
}

func (self *Purchase) Update(ctx context.Context, dbt float64, txid *int64, cancelledAt *time.Time, confirmedAt *time.Time, reason *string) error {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	err = self.UpdateTx(tx, ctx, dbt, txid, cancelledAt, confirmedAt, reason)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	return nil
}
