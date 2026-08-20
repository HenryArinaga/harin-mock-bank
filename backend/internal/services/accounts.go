package services

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CustomerAccount struct {
	AccountID      int64
	CustomerID     int64
	AccountStatus  string
	AccountType    string
	AccountNumber  string
	Currency       string
	AccountBalance string
}

func ListAccountsByCustomer(ctx context.Context, pool *pgxpool.Pool, customerID int64) ([]CustomerAccount, error) {
	rows, err := pool.Query(ctx, "SELECT account_id,customer_id,account_status,account_type,account_number,currency,balance FROM customer_account_balances WHERE customer_id = $1", customerID)
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
