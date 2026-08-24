package services

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InsertLedgerTransactionSample(ctx context.Context, pool *pgxpool.Pool) error {
	transactionID := int64(305)
	transaction, err := GetCompletedTransferByID(ctx, pool, transactionID)
	if err != nil {
		return err
	}
	input := LedgerInfo{
		AccountID:     transaction.FromAccountID,
		TransactionID: transaction.TransactionID,
		Direction:     "debit",
		Currency:      transaction.Currency,
		Amount:        transaction.Amount,
	}

	input2 := LedgerInfo{
		AccountID:     transaction.ToAccountID,
		TransactionID: transaction.TransactionID,
		Direction:     "credit",
		Currency:      transaction.Currency,
		Amount:        transaction.Amount,
	}

	err = InsertLedgerTransaction(ctx, pool, input)
	if err != nil {
		return err
	}
	err = InsertLedgerTransaction(ctx, pool, input2)
	if err != nil {
		return err
	}
	fmt.Printf("Success inserting into ledger\n")
	return nil
}
