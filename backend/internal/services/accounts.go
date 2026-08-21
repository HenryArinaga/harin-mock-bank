package services

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CustomerAccount struct {
	AccountID      int64
	CustomerID     int64
	FirstName      string
	LastName       string
	AccountStatus  string
	AccountType    string
	AccountNumber  string
	Currency       string
	AccountBalance string
}

type TransactionsByAccount struct {
	TransactionID          int64
	ToAccountID            pgtype.Text
	FromAccountID          pgtype.Text
	TransactionType        string
	TransactionDescription string
	TransactionStatus      string
	Currency               string
	Amount                 string
	created_at             time.Time
}

func ListAccountsByCustomer(ctx context.Context, pool *pgxpool.Pool, customerID int64) ([]CustomerAccount, error) {
	rows, err := pool.Query(ctx, "SELECT account_id,customer_id,first_name,last_name,account_status,account_type,account_number,currency,balance FROM customer_account_balances WHERE customer_id = $1", customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accounts := make([]CustomerAccount, 0)
	for rows.Next() {
		var account CustomerAccount
		err := rows.Scan(
			&account.AccountID,
			&account.CustomerID,
			&account.FirstName,
			&account.LastName,
			&account.AccountStatus,
			&account.AccountType,
			&account.AccountNumber,
			&account.Currency,
			&account.AccountBalance,
		)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return accounts, nil
}

func ListTransactionsByAccount(ctx context.Context, pool *pgxpool.Pool, transactionID int64) ([]TransactionsByAccount, error) {
	rows, err := pool.Query(ctx, "SELECT transaction_id, to_account_number, from_account_number, transaction_type, transaction_description,transaction_status, currency, amount,created_at FROM transaction_history WHERE transaction_id = $1", transactionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := make([]TransactionsByAccount, 0)
	for rows.Next() {
		var transaction TransactionsByAccount
		err := rows.Scan(
			&transaction.TransactionID,
			&transaction.ToAccountID,
			&transaction.FromAccountID,
			&transaction.TransactionType,
			&transaction.TransactionDescription,
			&transaction.TransactionStatus,
			&transaction.Currency,
			&transaction.Amount,
			&transaction.created_at,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return transactions, nil
}
