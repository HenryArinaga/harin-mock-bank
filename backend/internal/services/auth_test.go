package services

import (
	"context"
	"fmt"
	"harin-mock-bank/backend/internal/db"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestSignUpPassLogInPass(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	godotenv.Load("../../../.env")
	pool, err := db.NewPool(ctx)
	if err != nil {
		t.Fatalf("could not connect to pool: %v", err)
	}
	defer pool.Close()

	t.Logf("Success connecting to pool\n")

	email := fmt.Sprintf("test-%d@gmail.com", time.Now().UnixNano())

	signUp := SignUpUserInput{
		Email:    email,
		Password: "123456",
		UserRole: "customer",
	}

	_, err = SignUpUser(ctx, pool, signUp)
	if err != nil {
		t.Fatalf("Error Signing Up: %v\n", err)
	}

	logIn := LoginUserInput{
		Email:    email,
		Password: "123456",
	}

	LoginResult, err := LogInUser(ctx, pool, logIn)
	if err != nil {
		t.Fatalf("error logging in: %v", err)
	}
	if LoginResult.User.Email != email {
		t.Fatalf("expected email %s, got %s", email, LoginResult.User.Email)
	}

	if LoginResult.User.UserRole != "customer" {
		t.Fatalf("expected user role customer, got %s", LoginResult.User.UserRole)
	}

	if LoginResult.User.ID == 0 {
		t.Fatalf("expected authenticated user ID to be set")
	}
}

func TestSignUpPassLogInFail(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	godotenv.Load("../../../.env")
	pool, err := db.NewPool(ctx)
	if err != nil {
		t.Fatalf("could not connect to pool: %v", err)
	}
	defer pool.Close()

	t.Logf("Success connecting to pool\n")

	email := fmt.Sprintf("test-%d@gmail.com", time.Now().UnixNano())

	signUp := SignUpUserInput{
		Email:    email,
		Password: "123456",
		UserRole: "customer",
	}

	_, err = SignUpUser(ctx, pool, signUp)
	if err != nil {
		t.Fatalf("Error Signing Up: %v\n", err)
	}

	logIn := LoginUserInput{
		Email:    email,
		Password: "12345",
	}

	_, err = LogInUser(ctx, pool, logIn)
	if err == nil {
		t.Fatalf("expected login to fail")
	}
}

func TestSignUpFail(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	godotenv.Load("../../../.env")
	pool, err := db.NewPool(ctx)
	if err != nil {
		t.Fatalf("could not connect to pool: %v", err)
	}
	defer pool.Close()

	t.Logf("Success connecting to pool\n")
	email := fmt.Sprintf("test-%d@gmail.com", time.Now().UnixNano())
	signUp := SignUpUserInput{
		Email:    email,
		Password: "123456",
		UserRole: "customer",
	}

	_, err = SignUpUser(ctx, pool, signUp)
	if err != nil {
		t.Fatalf("error signing up in: %v", err)
	}

	_, err = SignUpUser(ctx, pool, signUp)
	if err == nil {
		t.Fatalf("expected Sign up to fail")
	}
}

func TestLogInFail(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	godotenv.Load("../../../.env")
	pool, err := db.NewPool(ctx)
	if err != nil {
		t.Fatalf("could not connect to pool: %v", err)
	}
	defer pool.Close()

	t.Logf("Success connecting to pool\n")

	logIn := LoginUserInput{
		Email:    fmt.Sprintf("missing-%d@gmail.com", time.Now().UnixNano()),
		Password: "12345",
	}

	_, err = LogInUser(ctx, pool, logIn)
	if err != nil {
		fmt.Printf("Log in failed %v\n", err)
	}
	if err == nil {
		t.Fatalf("expected login to fail")
	}

}

func TestLogInPass(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	godotenv.Load("../../../.env")
	pool, err := db.NewPool(ctx)
	if err != nil {
		t.Fatalf("could not connect to pool: %v", err)
	}
	defer pool.Close()

	t.Logf("Success connecting to pool\n")

	logIn := LoginUserInput{
		Email:    "test1@gmail.com",
		Password: "123456",
	}

	LoginResult, err := LogInUser(ctx, pool, logIn)
	if err != nil {
		t.Fatalf("error logging in: %v", err)
	}
	if LoginResult.Token == "" {
		t.Fatalf("empty token: %s", err)
	}
	if LoginResult.User.Email != logIn.Email {
		t.Fatalf("expected email %s, got %s", logIn.Email, LoginResult.User.Email)
	}

	if LoginResult.User.UserRole != "customer" {
		t.Fatalf("expected user role customer, got %s", LoginResult.User.UserRole)
	}

	if LoginResult.User.ID == 0 {
		t.Fatalf("expected authenticated user ID to be set")
	}

}
