package models

import (
	"context"
	"database/sql"
	"log"

	"github.com/flarehotspot/flarehotspot/core/db"
)

type PaymentModel struct {
	db     *db.Database
	models *Models
}

func NewPaymentModel(dtb *db.Database, mdls *Models) *PaymentModel {
	return &PaymentModel{dtb, mdls}
}

func (self *PaymentModel) CreateTx(tx *sql.Tx, ctx context.Context, purid int64, amt float64, mtd string) (*Payment, error) {
	query := "INSERT INTO payments (purchase_id, amount, optname) VALUES(?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, purid, amt, mtd)
	if err != nil {
		log.Println("SQL Exec Error: ", err)
		return nil, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		log.Println("SQL Exec Error: ", err)
		return nil, err
	}

	return self.FindTx(tx, ctx, lastId)
}

func (self *PaymentModel) FindTx(tx *sql.Tx, ctx context.Context, id int64) (*Payment, error) {
	payment := NewPayment(self.db, self.models)
	query := "SELECT id, purchase_id, amount, optname, created_at FROM payments WHERE id = ? LIMIT 1"
	err := tx.QueryRowContext(ctx, query, id).
		Scan(&payment.id, &payment.purchaseId, &payment.amount, &payment.optname, &payment.createdAt)

	return payment, err
}

func (self *PaymentModel) FindAllByPurchaseTx(tx *sql.Tx, ctx context.Context, purId int64) ([]*Payment, error) {
	payments := []*Payment{}
	query := "SELECT id, purchase_id, amount, optname, created_at FROM payments WHERE purchase_id = ?"
	rows, err := tx.QueryContext(ctx, query, purId)

	for rows.Next() {
		pmt := NewPayment(self.db, self.models)
		err = rows.Scan(&pmt.id, &pmt.purchaseId, &pmt.amount, &pmt.optname, &pmt.createdAt)
		if err != nil {
			return nil, err
		}
		payments = append(payments, pmt)
	}

	return payments, nil
}

func (self *PaymentModel) UpdateTx(tx *sql.Tx, ctx context.Context, id int64, amt float64) error {
	query := "UPDATE payments SET amount = ? WHERE id = ? LIMIT 1"
	_, err := tx.ExecContext(ctx, query, amt, id)
	return err
}

func (self *PaymentModel) Create(ctx context.Context, purid int64, amt float64, mtd string) (*Payment, error) {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	payment, err := self.CreateTx(tx, ctx, purid, amt, mtd)
	if err != nil {
		return nil, err
	}

	return payment, tx.Commit()
}

func (self *PaymentModel) Find(ctx context.Context, id int64) (*Payment, error) {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	pmnt, err := self.FindTx(tx, ctx, id)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	return pmnt, err
}

func (self *PaymentModel) FindAllByPurchase(ctx context.Context, purId int64) ([]*Payment, error) {
	sqlDB := self.db.SqlDB()
	tx, err := sqlDB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	payments, err := self.FindAllByPurchaseTx(tx, ctx, purId)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	return payments, err
}

func (self *PaymentModel) Update(ctx context.Context, id int64, amt float64, dbt *float64, txid *int64) error {
	tx, err := self.db.SqlDB().BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = self.UpdateTx(tx, ctx, id, amt)
	if err != nil {
		return err
	}

	return tx.Commit()
}
