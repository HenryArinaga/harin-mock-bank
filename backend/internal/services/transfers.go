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

type CompleteTransferInput struct {
	TransactionID int64
	CustomerID    int64
}

func InitiateTransfer(ctx context.Context, db DBRunner, input MakeTransferInput) (int64, error) {
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
	row := db.QueryRow(ctx, insertTransferSQL, args)
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

func GetCompletedTransferByID(ctx context.Context, db DBRunner, transactionID int64) (GetTransferStatus, error) {
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
	row := db.QueryRow(ctx, getCompletedTransferByID, args)
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

func UpdatePendingTransferFailed(ctx context.Context, db DBRunner, transactionID int64) error {
	updateTransferByID :=
		`UPDATE transactions
	SET transaction_status = 'failed'
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
		return errors.New("no pending transfer found to fail")
	}
	return nil
}

func EnsureSufficientBalance(ctx context.Context, db DBRunner, accountID int64, amount string) error {
	balanceQueryCheck := `
	SELECT balance >= @amount
	FROM account_balances
	WHERE account_id = @accountID`
	args := pgx.NamedArgs{
		"amount":    amount,
		"accountID": accountID,
	}

	var hasEnoughBalance bool
	row := db.QueryRow(ctx, balanceQueryCheck, args)
	err := row.Scan(
		&hasEnoughBalance,
	)
	if err != nil {
		return err
	}
	if !hasEnoughBalance {
		return fmt.Errorf("insufficient funds error\n")
	}
	return nil
}

func failTransfer(ctx context.Context, tx pgx.Tx, transactionID int64, originalErr error) error {
	failErr := UpdatePendingTransferFailed(ctx, tx, transactionID)
	if failErr != nil {
		return failErr
	}
	commitErr := tx.Commit(ctx)
	if commitErr != nil {
		return commitErr
	}
	return originalErr
}

func CompleteTransfer(ctx context.Context, pool *pgxpool.Pool, input CompleteTransferInput) error {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	transaction, err := GetPendingTransferByID(ctx, tx, input.TransactionID)
	if err != nil {
		return err
	}

	accountOwnerErr := EnsureAccountOwnerValid(ctx, tx, transaction.FromAccountID, input.CustomerID)
	if accountOwnerErr != nil {
		return failTransfer(ctx, tx, input.TransactionID, accountOwnerErr)
	}

	fromAccountCheckErr := EnsureAccountExists(ctx, tx, transaction.FromAccountID)
	if fromAccountCheckErr != nil {
		return failTransfer(ctx, tx, input.TransactionID, fromAccountCheckErr)
	}

	toAccountCheckErr := EnsureAccountExists(ctx, tx, transaction.ToAccountID)
	if toAccountCheckErr != nil {
		return failTransfer(ctx, tx, input.TransactionID, toAccountCheckErr)
	}

	fromActiveAccountErr := EnsureAccountActive(ctx, tx, transaction.FromAccountID)
	if fromActiveAccountErr != nil {
		return failTransfer(ctx, tx, input.TransactionID, fromActiveAccountErr)
	}

	toActiveAccountErr := EnsureAccountActive(ctx, tx, transaction.ToAccountID)
	if toActiveAccountErr != nil {
		return failTransfer(ctx, tx, input.TransactionID, toActiveAccountErr)
	}

	fromAccountCurrencyErr := EnsureCorrectCurrency(ctx, tx, transaction.FromAccountID, transaction.Currency)
	if fromAccountCurrencyErr != nil {
		return failTransfer(ctx, tx, input.TransactionID, fromAccountCurrencyErr)
	}

	toAccountCurrencyErr := EnsureCorrectCurrency(ctx, tx, transaction.ToAccountID, transaction.Currency)
	if toAccountCurrencyErr != nil {
		return failTransfer(ctx, tx, input.TransactionID, toAccountCurrencyErr)
	}

	validationErr := EnsureSufficientBalance(ctx, tx, transaction.FromAccountID, transaction.Amount)
	if validationErr != nil {
		return failTransfer(ctx, tx, input.TransactionID, validationErr)
	}

	debitEntry := LedgerInfo{
		AccountID:     transaction.FromAccountID,
		TransactionID: transaction.TransactionID,
		Direction:     "debit",
		Currency:      transaction.Currency,
		Amount:        transaction.Amount,
	}

	err = InsertLedgerTransaction(ctx, tx, debitEntry)
	if err != nil {
		return err
	}

	creditEntry := LedgerInfo{
		AccountID:     transaction.ToAccountID,
		TransactionID: transaction.TransactionID,
		Direction:     "credit",
		Currency:      transaction.Currency,
		Amount:        transaction.Amount,
	}
	err = InsertLedgerTransaction(ctx, tx, creditEntry)
	if err != nil {
		return err
	}

	err = UpdatePendingTransfer(ctx, tx, input.TransactionID)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}
	return nil

}
