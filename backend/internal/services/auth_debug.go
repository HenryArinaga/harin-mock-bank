package services

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateHashedAccountSample(ctx context.Context, pool *pgxpool.Pool) error {
	input := UserInformation{
		Email:    "test@gmail.com",
		Password: "123456",
		UserRole: "customer",
	}
	err := SignUpUser(ctx, pool, input)
	if err != nil {
		return err
	}
	return nil
}

func UserLogInSample(ctx context.Context, pool *pgxpool.Pool) error {
	input := UserInformation{
		Email:    "test@gmail.com",
		Password: "12345",
	}
	err := LogInUser(ctx, pool, input)
	if err != nil {
		return err
	}
	return nil
}
