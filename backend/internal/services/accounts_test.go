package services

import (
	"context"
	"fmt"
	"harin-mock-bank/backend/internal/db"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
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

func TestIsAccountStatusTransitionAllowed(t *testing.T) {

	got := isAccountStatusTransitionAllowed("pending", "active")
	if !got {
		t.Fatalf("cannot update account status")
	}
}

func TestIsAccountStatusTransitionInvalid(t *testing.T) {

	got := isAccountStatusTransitionAllowed("active", "pending")
	if got == true {
		t.Fatalf("should not be able to update account")
	}
}

func TestWrongAccountStatusName(t *testing.T) {

	got := isAccountStatusTransitionAllowed("active", "banana")
	if got == true {
		t.Fatalf("should not be able to update account")
	}
}

func TestIsAllowedToChangeStatusOfAccount(t *testing.T) {
	result := isAllowedToChangeStatusOfAccount("admin", "closed")
	if !result {
		t.Fatalf("cannot update account status")
	}
}

func TestIsNotallowedToChangeStatusOfAccount(t *testing.T) {
	result := isAllowedToChangeStatusOfAccount("customer", "active")
	if result == true {
		t.Fatalf("should not be able to update account")
	}
}

func TestOpenAccountSuccess(t *testing.T) {
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
	t.Logf("Registration email: %s", email)

	signUp := RegistrationInput{
		Email:     email,
		Password:  "123456",
		FirstName: "Henry",
		LastName:  "Testing",
		Phone:     "6614265690",
		DOB:       "1999-02-18",
	}

	err = RegisterCustomer(ctx, pool, signUp)
	if err != nil {
		t.Fatalf("Error:: %v\n", err)
	}

	logIn := LoginUserInput{
		Email:    signUp.Email,
		Password: signUp.Password,
	}

	loginResult, err := LogInUser(ctx, pool, logIn)
	if err != nil {
		t.Fatalf("error logging in: %v", err)
	}

	customerAccount := CustomerAccount{
		AccountType: "checking",
		Currency:    "USD",
	}

	accountID, err := OpenAccount(ctx, pool, loginResult.Token, customerAccount)
	if err != nil {
		t.Fatalf("error in: %v", err)
	}

	query := `
	SELECT customer_id, account_type, currency, account_status
	FROM accounts
	WHERE id = @accountID`

	args := pgx.NamedArgs{
		"accountID": accountID,
	}
	var customerID int64
	var accountType string
	var currency string
	var accountStatus string
	row := pool.QueryRow(ctx, query, args)
	err = row.Scan(
		&customerID,
		&accountType,
		&currency,
		&accountStatus,
	)
	if err != nil {
		t.Fatalf("error in: %v", err)
	}

	if accountType != "checking" {
		t.Fatalf("expected checking got: %v", accountType)
	}

	if currency != "USD" {
		t.Fatalf("expected USD got: %v", currency)
	}

	if accountStatus != "pending" {
		t.Fatalf("expected Pending got: %v", accountStatus)
	}

	profile, err := GetCustomerProfileByUserID(ctx, pool, loginResult.User.ID)
	if err != nil {
		t.Fatalf("error in: %v", err)
	}
	if customerID != profile.ID {
		t.Fatalf("expected %v got: %v", profile.ID, customerID)
	}
}
func TestOpenAccountInvalidToken(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	godotenv.Load("../../../.env")
	pool, err := db.NewPool(ctx)
	if err != nil {
		t.Fatalf("could not connect to pool: %v", err)
	}
	defer pool.Close()

	t.Logf("Success connecting to pool\n")

	token := "1231"

	customerAccount := CustomerAccount{
		AccountType: "checking",
		Currency:    "USD",
	}

	accountID, err := OpenAccount(ctx, pool, token, customerAccount)
	if err == nil {
		t.Fatalf("expected an error for invalid token")
	}
	if accountID != 0 {
		t.Fatalf("expected account ID 0, got %d", accountID)
	}

}
