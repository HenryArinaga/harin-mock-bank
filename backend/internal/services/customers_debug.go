package services

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func PrintCustomerBalanceSample(ctx context.Context, pool *pgxpool.Pool) error {
	customers, err := ListCustomerBalances(ctx, pool)
	if err != nil {
		return err
	}
	fmt.Printf("Loaded %d customers\n", len(customers))

	if len(customers) > 0 {
		firstCustomer := customers[0]
		fmt.Printf("%d, %s, %s, %s, %s\n",
			firstCustomer.CustomerID,
			firstCustomer.FirstName,
			firstCustomer.LastName,
			firstCustomer.Email,
			string(firstCustomer.BalancesByCurrency))
	}
	return nil
}

func PrintCustomerAccountBalanceSample(ctx context.Context, pool *pgxpool.Pool) error {
	customerID := int64(1)
	accounts, err := ListAccountsByCustomer(ctx, pool, customerID)
	if err != nil {
		return err
	}
	if len(accounts) > 0 {
		firstAccount := accounts[0]
		fmt.Printf("%d, %d, %s, %s, %s, %s, %s, %s, %s\n",
			firstAccount.AccountID,
			firstAccount.CustomerID,
			firstAccount.FirstName,
			firstAccount.LastName,
			firstAccount.AccountStatus,
			firstAccount.AccountType,
			firstAccount.AccountNumber,
			firstAccount.Currency,
			firstAccount.AccountBalance,
		)
	}
	return nil
}

func PrintTransactionByAccountSample(ctx context.Context, pool *pgxpool.Pool) error {
	accountNumber := "654218"
	transactions, err := ListTransactionsByAccount(ctx, pool, accountNumber)
	if err != nil {
		return err
	}
	fmt.Printf("Loaded %d transactions\n", len(transactions))
	for _, transaction := range transactions {
		fmt.Printf("%d, %s, %s, %s, %s, %s, %s, %s, %v\n",
			transaction.TransactionID,
			transaction.ToAccountNumber.String,
			transaction.FromAccountNumber.String,
			transaction.TransactionType,
			transaction.TransactionDescription.String,
			transaction.TransactionStatus,
			transaction.Currency,
			transaction.Amount,
			transaction.CreatedAt,
		)
	}
	return nil
}

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

func InsertLedgerTransactionSample(ctx context.Context, pool *pgxpool.Pool) error {
	transactionID := int64(303)
	transaction, err := GetCompletedTransferByID(ctx, pool, transactionID)
	if err != nil {
		return err
	}
	input := LedgerInfo{
		AccountID:     transaction.FromAccountID,
		TransactionID: transaction.TransactionID,
		Direction:     "debit",
		Currency:      transaction.Currency,
		Amount:        transaction.Amount,
	}

	err = InsertLedgerTransaction(ctx, pool, input)
	if err != nil {
		return err
	}
	fmt.Printf("Success inserting into ledger\n")
	return nil

}
