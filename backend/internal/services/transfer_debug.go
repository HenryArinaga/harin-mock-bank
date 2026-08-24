package services

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func PrintInitiateTransferSample(ctx context.Context, pool *pgxpool.Pool) error {
	input := MakeTransferInput{
		FromAccountID:          28,
		ToAccountID:            29,
		Currency:               `USD`,
		Amount:                 `150`,
		TransactionDescription: `IPhone 18 Pro MAX`,
	}
	transactionID, err := InitiateTransfer(ctx, pool, input)
	if err != nil {
		return err
	}
	fmt.Printf("Success inserting into database\nCreated transfer transaction ID: %v\n", transactionID)
	return nil
}

func GetPendingTransferByIDSample(ctx context.Context, pool *pgxpool.Pool) error {
	transactionID := int64(302)
	transaction, err := GetPendingTransferByID(ctx, pool, transactionID)
	if err != nil {
		return err
	}
	fmt.Printf("%d is %s\n", transaction.TransactionID, transaction.TransactionStatus)

	return nil
}

func GetCompletedTransferByIDSample(ctx context.Context, pool *pgxpool.Pool) error {
	transactionID := int64(303)
	transaction, err := GetCompletedTransferByID(ctx, pool, transactionID)
	if err != nil {
		return err
	}
	fmt.Printf("%d is %s\n", transaction.TransactionID, transaction.TransactionStatus)

	return nil
}

func UpdatePendingTransferSample(ctx context.Context, pool *pgxpool.Pool) error {
	// transactionID values need to be updated depending
	// On if the current transaction is not pending
	transactionID := int64(303)
	err := UpdatePendingTransfer(ctx, pool, transactionID)
	if err != nil {
		return err
	}
	fmt.Printf("Transaction %d succesfully updated to completed\n", transactionID)

	return nil
}

func CompleteTransferSample(ctx context.Context, pool *pgxpool.Pool) error {
	transactionID := int64(306)
	err := CompleteTransfer(ctx, pool, transactionID)
	if err != nil {
		return err
	}
	fmt.Printf("Success completing transfer: %d", transactionID)
	return nil
}
