package services

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func PrintCustomerBalanceSample(ctx context.Context, pool *pgxpool.Pool) error {
	customers, err := ListCustomerBalances(ctx, pool)
	if err != nil {
		return err
	}
	fmt.Printf("Loaded %d customers\n", len(customers))

	if len(customers) > 0 {
		firstCustomer := customers[0]
		fmt.Printf("%d, %s, %s, %s, %s\n",
			firstCustomer.CustomerID,
			firstCustomer.FirstName,
			firstCustomer.LastName,
			firstCustomer.Email,
			string(firstCustomer.BalancesByCurrency))
	}
	return nil
}

func PrintCustomerAccountBalanceSample(ctx context.Context, pool *pgxpool.Pool) error {
	customerID := int64(1)
	accounts, err := ListAccountsByCustomer(ctx, pool, customerID)
	if err != nil {
		return err
	}
	if len(accounts) > 0 {
		firstAccount := accounts[0]
		fmt.Printf("%d, %d, %s, %s, %s, %s, %s\n",
			firstAccount.AccountID,
			firstAccount.CustomerID,
			firstAccount.AccountStatus,
			firstAccount.AccountType,
			firstAccount.AccountNumber,
			firstAccount.Currency,
			firstAccount.AccountBalance,
		)
	}
	return nil
}
