package services

import (
	"context"
	"harin-mock-bank/backend/internal/db"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func TestCompleteTransferWrongOwnerFails(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	godotenv.Load("../../../.env")
	pool, err := db.NewPool(ctx)
	if err != nil {
		t.Fatalf("could not connect to pool: %v", err)
	}
	defer pool.Close()

	t.Logf("Success connecting to pool\n")

	transfer := MakeTransferInput{
		FromAccountID:          2,
		ToAccountID:            9,
		Currency:               "USD",
		Amount:                 "1.00",
		TransactionDescription: "Test transfer",
	}

	transactionID, err := InitiateTransfer(ctx, pool, transfer)
	if err != nil {
		t.Fatalf("initiate transfer: %v", err)
	}
	customerID := int64(74)
	input := CompleteTransferInput{
		TransactionID: transactionID,
		CustomerID:    customerID,
	}
	err = CompleteTransfer(ctx, pool, input)
	if err == nil {
		t.Fatalf("expected wrong owner transfer to fail")
	}

	initiateTransferQuery := `
	SELECT transaction_status
	FROM transactions
	WHERE id = @transactionID`
	args := pgx.NamedArgs{
		"transactionID": transactionID,
	}

	var status string
	row := pool.QueryRow(ctx, initiateTransferQuery, args)
	err = row.Scan(&status)
	if err != nil {
		t.Fatalf("query transaction status: %v", err)
	}

	query2 := `
	SELECT COUNT(*)
	FROM ledger_entries
	WHERE transaction_id = @transactionID`

	args = pgx.NamedArgs{
		"transactionID": transactionID,
	}

	var ledgerCount int
	row = pool.QueryRow(ctx, query2, args)
	err = row.Scan(&ledgerCount)
	if err != nil {
		t.Fatalf("query ledger count: %v", err)
	}
	if ledgerCount != 0 {
		t.Fatalf("expected 0 ledger entries, got %d", ledgerCount)
	}

	t.Logf("Executing Query: %s\n", initiateTransferQuery)
	t.Logf("Executing Query: %s\n", query2)

}
