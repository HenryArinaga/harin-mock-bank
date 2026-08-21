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

func InitiateTransfer(ctx context.Context, pool *pgxpool.Pool, input MakeTransferInput) error {
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
		)`
	args := pgx.NamedArgs{
		"fromAccount":            input.FromAccountID,
		"toAccount":              input.ToAccountID,
		"currency":               input.Currency,
		"amount":                 input.Amount,
		"transactionDescription": input.TransactionDescription,
	}
	_, err := pool.Exec(ctx, insertTransferSQL, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}
