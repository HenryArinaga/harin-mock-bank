package services

import (
	"context"
	"fmt"
	"harin-mock-bank/backend/internal/db"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestCreateAccount(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	godotenv.Load("../../../.env")
	pool, err := db.NewPool(ctx)
	if err != nil {
		t.Fatalf("could not connect to pool: %v", err)
	}
	defer pool.Close()
	email := fmt.Sprintf("test-%d@gmail.com", time.Now().UnixNano())

	signUp := SignUpUserInput{
		Email:    email,
		Password: "123456",
		UserRole: "customer",
	}

	userID, err := SignUpUser(ctx, pool, signUp)
	if err != nil {
		t.Fatalf("Error Signing Up: %v\n", err)
	}

	customerInput := CustomerProfileInformation{
		UserID:    userID,
		FirstName: "Henry",
		LastName:  "Arinaga",
		Phone:     "6614265690",
		DOB:       "1999-02-18",
	}
	_, err = CreateCustomerProfile(ctx, pool, customerInput)
	if err != nil {
		t.Fatalf("Error:: %v\n", err)
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

	customerProfile, err := GetCustomerProfileByJWT(ctx, pool, loginResult.Token)
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	if customerProfile.UserID != loginResult.User.ID {
		t.Fatalf("User ID mismatch")
	}

	t.Logf("Success connecting to pool\n")

	input := CustomerAccount{
		CustomerID:  customerProfile.ID,
		Currency:    "USD",
		AccountType: "checking",
	}
	accountID, err := CreateAccount(ctx, pool, input)
	if err != nil {
		t.Fatalf("Create account err: %v", err)
	}
	t.Logf("created account id: %d for customer id: %d", accountID, customerProfile.ID)
}
