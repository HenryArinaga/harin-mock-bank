package main

import (
	"context"
	"fmt"
	"harin-mock-bank/backend/internal/db"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	godotenv.Load("../.env")

	pool, err := db.NewPool(ctx)
	if err != nil {
		fmt.Printf("Could not connect to pool: %v\n", err)
		return
	}
	defer pool.Close()

	fmt.Printf("Success connecting to pool\n")

	/* if err := services.PrintCustomerBalanceSample(ctx, pool); err != nil {
		fmt.Printf("Load error: %v\n", err)
		return
	}

	if err := services.PrintCustomerAccountBalanceSample(ctx, pool); err != nil {
		fmt.Printf("Load error: %v\n", err)
		return
	}

	if err := services.PrintTransactionByAccountSample(ctx, pool); err != nil {
		fmt.Printf("Load error: %v\n", err)
		return
	} */

	/* err := services.PrintInitiateTransferSample(ctx, pool); err != nil {
		fmt.Printf("Transfer error: %v\n", err)
		return
	}*/

	/*
		err = services.GetPendingTransferByIDSample(ctx, pool)
		if err != nil {
			fmt.Printf("Error getting transfer status: %v\n", err)
			return
		} */

	/*err = services.UpdateTransferSample(ctx, pool)
	if err != nil {
		fmt.Printf("Error updating transfer status: %s\n", err)
		return
	}*/

	return
}
