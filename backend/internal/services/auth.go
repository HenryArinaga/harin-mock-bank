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

type SignUpUserInput struct {
	Email    string
	Password string
	UserRole string
}

type LoginUserInput struct {
	Email    string
	Password string
}

type AuthenticatedUser struct {
	ID       int64
	Email    string
	UserRole string
}

type LoginResult struct {
	User  AuthenticatedUser
	Token string
}

func SignUpUser(ctx context.Context, pool *pgxpool.Pool, input SignUpUserInput) (int64, error) {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("begin transaction: %w", err)
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
		return 0, fmt.Errorf("hash password: %w", err)
	}

	var userID int64
	var pgErr *pgconn.PgError
	row := tx.QueryRow(ctx, signUpCustomerQuery, args)
	err = row.Scan(&userID)
	if err != nil {
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return 0, fmt.Errorf("email already exists")
		}
		return 0, fmt.Errorf("failed to insert user: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return userID, nil

}

func LogInUser(ctx context.Context, pool *pgxpool.Pool, input LoginUserInput) (LoginResult, error) {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return LoginResult{}, fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	queryEmail := `
	SELECT id, email, password_hash, user_role
	FROM users
	WHERE email = @email`

	args := pgx.NamedArgs{
		"email": input.Email,
	}

	var userID int64
	var email string
	var storedPasswordHash string
	var userRole string

	row := tx.QueryRow(ctx, queryEmail, args)
	err = row.Scan(
		&userID,
		&email,
		&storedPasswordHash,
		&userRole,
	)
	if err != nil {
		return LoginResult{}, fmt.Errorf("unable to scan into row: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(input.Password))
	if err != nil {
		return LoginResult{}, fmt.Errorf("invalid email or password")
	}

	authenticatedUser := AuthenticatedUser{
		ID:       userID,
		Email:    email,
		UserRole: userRole,
	}
	token, err := GenerateAuthToken(authenticatedUser)
	if err != nil {
		return LoginResult{}, err
	}
	updateLastLogin := `
	UPDATE users
	SET last_login_at = CURRENT_TIMESTAMP
	WHERE id = @userID`

	args = pgx.NamedArgs{
		"userID": userID,
	}

	_, err = tx.Exec(ctx, updateLastLogin, args)
	if err != nil {
		return LoginResult{}, fmt.Errorf("update last login: %w", err)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return LoginResult{}, fmt.Errorf("commit login: %w", err)
	}

	return LoginResult{
		User:  authenticatedUser,
		Token: token,
	}, nil
}
