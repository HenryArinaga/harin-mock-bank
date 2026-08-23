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
