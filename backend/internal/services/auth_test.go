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

	loginResult, err := LogInUser(ctx, pool, logIn)
	if err != nil {
		t.Fatalf("error logging in: %v", err)
	}
	if loginResult.User.Email != email {
		t.Fatalf("expected email %s, got %s", email, loginResult.User.Email)
	}

	if loginResult.User.UserRole != "customer" {
		t.Fatalf("expected user role customer, got %s", loginResult.User.UserRole)
	}

	if loginResult.User.ID == 0 {
		t.Fatalf("expected authenticated user ID to be set")
	}
}

func TestSignUpPassLogInPassJWTPass(t *testing.T) {
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

	loginResult, err := LogInUser(ctx, pool, logIn)
	if err != nil {
		t.Fatalf("error logging in: %v", err)
	}
	if loginResult.User.Email != email {
		t.Fatalf("expected email %s, got %s", email, loginResult.User.Email)
	}

	if loginResult.User.UserRole != "customer" {
		t.Fatalf("expected user role customer, got %s", loginResult.User.UserRole)
	}

	if loginResult.User.ID == 0 {
		t.Fatalf("expected authenticated user ID to be set")
	}

	if loginResult.Token == "" {
		t.Fatalf("token empty")
	}

	validatedUser, err := ValidateAuthToken(loginResult.Token)
	if err != nil {
		t.Fatalf("Error validating token: %v\n", err)
	}

	if validatedUser.ID != loginResult.User.ID {
		t.Fatalf("User ID incorrect")
	}
	if validatedUser.Email != loginResult.User.Email {
		t.Fatalf("User Email incorrect")
	}
	if validatedUser.UserRole != loginResult.User.UserRole {
		t.Fatalf("User role incorrect")
	}

}

func TestValidateAuthTokenFail(t *testing.T) {
	godotenv.Load("../../../.env")
	_, err := ValidateAuthToken("FakeToken")
	if err == nil {
		t.Fatalf("expected invalid token to fail")
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

	loginResult, err := LogInUser(ctx, pool, logIn)
	if err != nil {
		t.Fatalf("error logging in: %v", err)
	}
	if loginResult.Token == "" {
		t.Fatalf("empty token: %s", err)
	}
	if loginResult.User.Email != logIn.Email {
		t.Fatalf("expected email %s, got %s", logIn.Email, loginResult.User.Email)
	}

	if loginResult.User.UserRole != "customer" {
		t.Fatalf("expected user role customer, got %s", loginResult.User.UserRole)
	}

	if loginResult.User.ID == 0 {
		t.Fatalf("expected authenticated user ID to be set")
	}

}
