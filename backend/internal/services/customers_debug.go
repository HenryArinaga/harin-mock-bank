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
