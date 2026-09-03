package services

import (
	"context"
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

	t.Logf("Success connecting to pool\n")

	input := CustomerAccount{
		CustomerID:  76,
		Currency:    "USD",
		AccountType: "checking",
	}
	_, err = CreateAccount(ctx, pool, input)
	if err != nil {
		t.Fatalf("Create account err: %v", err)
	}
}
