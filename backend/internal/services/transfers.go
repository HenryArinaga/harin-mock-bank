package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MakeTransferInput struct {
	FromAccountID          int64
	ToAccountID            int64
	Currency               string
	Amount                 string
	TransactionDescription string
}

type GetTransferStatus struct {
	TransactionID          int64
	FromAccountID          int64
	ToAccountID            int64
	Currency               string
	Amount                 string
	TransactionDescription string
	TransactionStatus      string
}

func InitiateTransfer(ctx context.Context, pool *pgxpool.Pool, input MakeTransferInput) (int64, error) {
	insertTransferSQL :=
		`INSERT INTO transactions (
		transaction_type, 
		from_account_id,
		to_account_id, 
		currency, 
		amount, 
		transaction_description
		) 
		VALUES (
		'transfer', 
		@fromAccount, 
		@toAccount, 
		@currency, 
		@amount, 
		@transactionDescription
		)
		RETURNING id`
	args := pgx.NamedArgs{
		"fromAccount":            input.FromAccountID,
		"toAccount":              input.ToAccountID,
		"currency":               input.Currency,
		"amount":                 input.Amount,
		"transactionDescription": input.TransactionDescription,
	}
	var transactionID int64
	row := pool.QueryRow(ctx, insertTransferSQL, args)
	err := row.Scan(&transactionID)
	if err != nil {
		return 0, fmt.Errorf("unable to insert row: %w", err)
	}

	return transactionID, nil
}

func GetPendingTransferByID(ctx context.Context, db DBRunner, transactionID int64) (GetTransferStatus, error) {
	getPendingTransferByID :=
		`SELECT 
		id, 
		from_account_id, 
		to_account_id, 
		currency, amount, 
		transaction_description, 
		transaction_status 
		FROM transactions
	WHERE id = @id
	AND transaction_type = 'transfer' 
	AND transaction_status = 'pending'`

	args := pgx.NamedArgs{
		"id": transactionID,
	}
	var transaction GetTransferStatus
	row := db.QueryRow(ctx, getPendingTransferByID, args)
	err := row.Scan(
		&transaction.TransactionID,
		&transaction.FromAccountID,
		&transaction.ToAccountID,
		&transaction.Currency,
		&transaction.Amount,
		&transaction.TransactionDescription,
		&transaction.TransactionStatus,
	)
	if err != nil {
		return transaction, fmt.Errorf("unable to get pending transfer by id: %w", err)
	}

	return transaction, nil
}

func GetCompletedTransferByID(ctx context.Context, pool *pgxpool.Pool, transactionID int64) (GetTransferStatus, error) {
	getCompletedTransferByID :=
		`SELECT 
		id, 
		from_account_id, 
		to_account_id, 
		currency, amount, 
		transaction_description, 
		transaction_status 
		FROM transactions
	WHERE id = @id
	AND transaction_type = 'transfer' 
	AND transaction_status = 'completed'`

	args := pgx.NamedArgs{
		"id": transactionID,
	}
	var transaction GetTransferStatus
	row := pool.QueryRow(ctx, getCompletedTransferByID, args)
	err := row.Scan(
		&transaction.TransactionID,
		&transaction.FromAccountID,
		&transaction.ToAccountID,
		&transaction.Currency,
		&transaction.Amount,
		&transaction.TransactionDescription,
		&transaction.TransactionStatus,
	)
	if err != nil {
		return transaction, fmt.Errorf("unable to get pending transfer by id: %w", err)
	}

	return transaction, nil
}

func UpdatePendingTransfer(ctx context.Context, db DBRunner, transactionID int64) error {
	updateTransferByID :=
		`UPDATE transactions
	SET transaction_status = 'completed'
	WHERE id = @id
	AND transaction_type = 'transfer'
	AND transaction_status = 'pending'`
	args := pgx.NamedArgs{
		"id": transactionID,
	}
	commandTag, err := db.Exec(ctx, updateTransferByID, args)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() < 1 {
		return errors.New("no pending transfer found to complete")
	}
	return nil
}

func CompleteTransfer(ctx context.Context, pool *pgxpool.Pool, transactionID int64) error {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	transaction, err := GetPendingTransferByID(ctx, tx, transactionID)
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
	err = InsertLedgerTransaction(ctx, tx, input)
	if err != nil {
		return err
	}

	input2 := LedgerInfo{
		AccountID:     transaction.ToAccountID,
		TransactionID: transaction.TransactionID,
		Direction:     "credit",
		Currency:      transaction.Currency,
		Amount:        transaction.Amount,
	}
	err = InsertLedgerTransaction(ctx, tx, input2)
	if err != nil {
		return err
	}

	err = UpdatePendingTransfer(ctx, tx, transactionID)
	if err != nil {
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}
	return nil

}
