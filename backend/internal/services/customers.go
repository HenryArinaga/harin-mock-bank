package services

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CustomerBalance struct {
	CustomerID         int64
	FirstName          string
	LastName           string
	Email              string
	BalancesByCurrency []byte
}

type CustomerProfileInformation struct {
	UserID    int64
	FirstName string
	LastName  string
	Phone     string
	DOB       string
}

type CustomerProfile struct {
	ID        int64
	UserID    int64
	FirstName string
	LastName  string
	Phone     string
	DOB       time.Time
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

func CreateCustomerProfile(ctx context.Context, db DBRunner, input CustomerProfileInformation) (int64, error) {
	insertCustomerInformation := `
	INSERT INTO customers (
	user_id,
	first_name,
	last_name,
	phone,
	date_of_birth
	)
	VALUES (
	@userID,
	@firstName,
	@lastName,
	@phone,
	@DOB
	)
	RETURNING id`

	args := pgx.NamedArgs{
		"userID":    input.UserID,
		"firstName": input.FirstName,
		"lastName":  input.LastName,
		"phone":     input.Phone,
		"DOB":       input.DOB,
	}
	var customerID int64
	row := db.QueryRow(ctx, insertCustomerInformation, args)
	err := row.Scan(&customerID)
	if err != nil {
		return 0, fmt.Errorf("unable to create customer profile: %w", err)
	}

	return customerID, nil

}

func GetCustomerProfileByUserID(ctx context.Context, db DBRunner, userID int64) (CustomerProfile, error) {
	customerProfilesQuery := `
	SELECT 
	id, 
	user_id, 
	first_name, 
	last_name, 
	phone, 
	date_of_birth 
	FROM customers 
	WHERE user_id = @userID`
	args := pgx.NamedArgs{
		"userID": userID,
	}
	var customer CustomerProfile
	row := db.QueryRow(ctx, customerProfilesQuery, args)
	err := row.Scan(
		&customer.ID,
		&customer.UserID,
		&customer.FirstName,
		&customer.LastName,
		&customer.Phone,
		&customer.DOB,
	)
	if err != nil {
		return CustomerProfile{}, fmt.Errorf("get customer profile by user id: %w", err)
	}
	return customer, nil

}
