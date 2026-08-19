package services

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CustomerBalance struct {
	CustomerID         int64
	FirstName          string
	LastName           string
	Email              string
	BalancesByCurrency []byte
}

func ListCustomerBalances(ctx context.Context, pool *pgxpool.Pool) ([]CustomerBalance, error) {
	rows, err := pool.Query(ctx, "select customer_id, first_name, last_name, email, balances_by_currency from customer_balances_json order by customer_id limit 20")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customers := make([]CustomerBalance, 0)
	for rows.Next() {
		var customer CustomerBalance
		err := rows.Scan(
			&customer.CustomerID,
			&customer.FirstName,
			&customer.LastName,
			&customer.Email,
			&customer.BalancesByCurrency,
		)
		if err != nil {
			return nil, err
		}

		customers = append(customers, customer)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return customers, nil
}
