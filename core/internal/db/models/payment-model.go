package models

import (
	"context"
	"fmt"
	"log"

	"core/internal/db"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type PaymentModel struct {
	db     *db.Database
	models *Models
}

func NewPaymentModel(dtb *db.Database, mdls *Models) *PaymentModel {
	return &PaymentModel{dtb, mdls}
}

func (self *PaymentModel) CreateTx(tx pgx.Tx, ctx context.Context, purid uuid.UUID, amt float64, mtd string) (*Payment, error) {
	defer func() {
		if err := tx.Rollback(ctx); err != nil && err != pgx.ErrTxClosed {
			log.Printf("Rollback failed: %v", err)
		}
	}()

	query := "INSERT INTO payments (purchase_id, amount, optname) VALUES($1, $2, $3) RETURNING id"

	var lastInsertId int
	err := tx.QueryRow(ctx, query, purid, amt, mtd).Scan(&lastInsertId)
	if err != nil {
		log.Printf("SQL Execution failed: %v", err)
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		log.Printf("SQL transaction commit failed: %v", err)
		return nil, err
	}

	return self.FindTx(tx, ctx, int64(lastInsertId))
}

func (self *PaymentModel) FindTx(tx pgx.Tx, ctx context.Context, id int64) (*Payment, error) {
	payment := NewPayment(self.db, self.models)

	query := "SELECT id, purchase_id, amount, optname, created_at FROM payments WHERE id = $1 LIMIT 1"

	err := tx.QueryRow(ctx, query, id).
		Scan(&payment.id, &payment.purchaseId, &payment.amount, &payment.optname, &payment.createdAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			log.Printf("No payment found with id %d", id)
			return nil, nil
		}
		log.Printf("Error finding payment with id %d: %v", id, err)
		return nil, err
	}

	return payment, nil
}

func (self *PaymentModel) FindAllByPurchaseTx(tx pgx.Tx, ctx context.Context, purId uuid.UUID) ([]*Payment, error) {
	payments := []*Payment{}

	query := "SELECT id, purchase_id, amount, optname, created_at FROM payments WHERE purchase_id = $1"

	rows, err := tx.Query(ctx, query, purId)
	if err != nil {
		return nil, fmt.Errorf("query exection failed: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		pmt := NewPayment(self.db, self.models)
		err = rows.Scan(&pmt.id, &pmt.purchaseId, &pmt.amount, &pmt.optname, &pmt.createdAt)
		if err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}

		payments = append(payments, pmt)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("row iteration error: %w", rows.Err())
	}

	return payments, nil
}

func (self *PaymentModel) UpdateTx(tx pgx.Tx, ctx context.Context, id uuid.UUID, amt float64) error {
	query := "UPDATE payments SET amount = $1 WHERE id = $2 LIMIT 1"

	cmdTag, err := tx.Exec(ctx, query, amt, id)
	if err != nil {
		log.Printf("SQL Exec Error while updating payment ID %d: %v", id, err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		log.Printf("No payment found with id %d; update operation skipped", id)
		return fmt.Errorf("device with id %d not found", id)
	}

	log.Printf("Successfully updated device with id %d", id)
	return nil
}

func (self *PaymentModel) Create(ctx context.Context, purid uuid.UUID, amt float64, mtd string) (*Payment, error) {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	payment, err := self.CreateTx(tx, ctx, purid, amt, mtd)
	if err != nil {
		return nil, fmt.Errorf("failed to create device: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("could not commit transaction: %w", err)
	}

	return payment, nil
}

func (self *PaymentModel) Find(ctx context.Context, id int64) (*Payment, error) {
	tx, err := self.db.SqlDB().Begin(ctx)

	if err != nil {
		return nil, fmt.Errorf("could not begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	pmnt, err := self.FindTx(tx, ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find payment: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("could not commit transaction: %w", err)
	}

	return pmnt, nil
}

func (self *PaymentModel) FindAllByPurchase(ctx context.Context, purId uuid.UUID) ([]*Payment, error) {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	payments, err := self.FindAllByPurchaseTx(tx, ctx, purId)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve payments: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("could not commit transaction: %w", err)
	}

	return payments, nil
}

func (self *PaymentModel) Update(ctx context.Context, id uuid.UUID, amt float64, dbt *float64, txid *int64) error {
	tx, err := self.db.SqlDB().Begin(ctx)
	if err != nil {
		return fmt.Errorf("count not begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	err = self.UpdateTx(tx, ctx, id, amt)
	if err != nil {
		return fmt.Errorf("could not update payment: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	return nil
}
