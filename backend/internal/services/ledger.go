package services

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type LedgerInfo struct {
	AccountID     int64
	TransactionID int64
	Direction     string
	Currency      string
	Amount        string
}

func InsertLedgerTransaction(ctx context.Context, db DBRunner, input LedgerInfo) error {
	insertLedgerTransaction :=
		`INSERT INTO ledger_entries (
	account_id,
	transaction_id,
	direction,
	currency,
	amount
	)
	VALUES (
	@accountID,
	@transactionID,
	@direction,
	@currency,
	@amount
	) `
	args := pgx.NamedArgs{
		"accountID":     input.AccountID,
		"transactionID": input.TransactionID,
		"direction":     input.Direction,
		"currency":      input.Currency,
		"amount":        input.Amount,
	}
	_, err := db.Exec(ctx, insertLedgerTransaction, args)
	if err != nil {
		return fmt.Errorf("unable to insert ledger entry: %w", err)
	}

	return nil
}
