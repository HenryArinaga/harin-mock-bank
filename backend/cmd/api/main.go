package main

import (
	"context"
	"fmt"
	"harin-mock-bank/backend/internal/db"
	"harin-mock-bank/backend/internal/services"
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
	/*
		err = services.PrintInitiateTransferSample(ctx, pool)
		if err != nil {
			fmt.Printf("Transfer error: %v\n", err)
			return
		}
	*/
	/*
		err = services.GetPendingTransferByIDSample(ctx, pool)
		if err != nil {
			fmt.Printf("Error getting transfer status: %v\n", err)
			return
		} */

	/* err = services.GetCompletedTransferByIDSample(ctx, pool)
	if err != nil {
		fmt.Printf("Error getting transfer status: %v\n", err)
		return
	} */

	/*err = services.UpdateTransferSample(ctx, pool)
	if err != nil {
		fmt.Printf("Error updating transfer status: %s\n", err)
		return
	}*/

	/*
		err = services.InsertLedgerTransactionSample(ctx, pool)
		if err != nil {
			fmt.Printf("error inserting into ledger %v", err)
			return
		} */
	/*
		err = services.CompleteTransferSample(ctx, pool)
		if err != nil {
			fmt.Printf("error completing transfer: %v\n", err)
			return
		}
	*/
	/*
		err = services.CreateHashedAccountSample(ctx, pool)
		if err != nil {
			fmt.Printf("error creating account: %v\n", err)
			return
		}
	*/
	/*
		err = services.UserLogInSample(ctx, pool)
		if err != nil {
			fmt.Printf("error logging in: %v\n", err)
			return
		} */
	/*
		err = services.CreateCustomerProfileSample(ctx, pool)
		if err != nil {
			fmt.Printf("error creating customer profile: %v\n", err)
			return
		}
	*/
	err = services.GetCustomerProfileByUserIDSample(ctx, pool)
	if err != nil {
		fmt.Printf("error getting customer profile: %v\n", err)
		return
	}
}
