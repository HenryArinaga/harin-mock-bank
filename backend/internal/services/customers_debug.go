package services

import (
	"context"
	"fmt"
	"time"

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

func CreateCustomerProfileSample(ctx context.Context, pool *pgxpool.Pool) error {
	email := fmt.Sprintf("test-%d@gmail.com", time.Now().UnixNano())

	signUp := SignUpUserInput{
		Email:    email,
		Password: "123456",
		UserRole: "customer",
	}

	userID, err := SignUpUser(ctx, pool, signUp)
	if err != nil {
		return err
	}

	input := CustomerProfileInformation{
		UserID:    userID,
		FirstName: "Henry",
		LastName:  "Arinaga",
		Phone:     "6614265690",
		DOB:       "1999-02-18",
	}
	_, err = CreateCustomerProfile(ctx, pool, input)
	if err != nil {
		return err
	}
	return nil
}

func GetCustomerProfileByUserIDSample(ctx context.Context, pool *pgxpool.Pool) error {

	profile, err := GetCustomerProfileByUserID(ctx, pool, 130)
	if err != nil {
		return err
	}
	fmt.Printf("Profile is %v\n", profile)
	return nil
}
