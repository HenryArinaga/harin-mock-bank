package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type UserSignUp struct {
	Email    string
	Password string
	UserRole string
}

func SignUpUser(ctx context.Context, pool *pgxpool.Pool, input UserSignUp) error {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	signUpCustomerQuery := `
	INSERT INTO users (
	email,
	password_hash,
	user_role
	)
	VALUES (
	@email,
	@password,
	@userRole
	)
	RETURNING id`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	args := pgx.NamedArgs{
		"email":    input.Email,
		"password": string(hashedPassword),
		"userRole": input.UserRole,
	}

	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	var userID int64
	var pgErr *pgconn.PgError
	row := tx.QueryRow(ctx, signUpCustomerQuery, args)
	err = row.Scan(&userID)
	if err != nil {
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return fmt.Errorf("email already exists")
		}
		return fmt.Errorf("failed to insert user: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil

}
