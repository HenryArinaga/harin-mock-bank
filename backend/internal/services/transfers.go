package services

import (
	"context"
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

func GetPendingTransferByID(ctx context.Context, pool *pgxpool.Pool, transactionID int64) (GetTransferStatus, error) {
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
	row := pool.QueryRow(ctx, getPendingTransferByID, args)
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

func UpdateTransfer(ctx context.Context, pool *pgxpool.Pool, transactionID int64) error {
	updateTransferByID :=
		`UPDATE transactions
	SET transaction_status = 'completed'
	WHERE id = @id
	AND transaction_type = 'transfer'
	AND transaction_status = 'pending'`
	args := pgx.NamedArgs{
		"id": transactionID,
	}
	_, err := pool.Exec(ctx, updateTransferByID, args)

	if err != nil {
		return err
	}

	return nil
}
