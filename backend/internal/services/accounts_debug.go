package services

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

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

func UpdateAccountStatusSample(ctx context.Context, pool *pgxpool.Pool) error {
	input := accountStatusUpdate{
		NewStatus:     "active",
		AccountID:     166,
		CurrentStatus: "frozen",
	}
	err := updateAccountStatus(ctx, pool, input)
	if err != nil {
		return err
	}
	fmt.Printf("Account current status %s, account number %v: successfully updated to %s\n", input.CurrentStatus, input.AccountID, input.NewStatus)
	return nil
}

func ChangeAccountStatusSample(ctx context.Context, pool *pgxpool.Pool) error {

	newStatus := "closed"
	accountID := int64(166)
	userRole := "support"

	err := ChangeAccountStatus(ctx, pool, accountID, newStatus, userRole)
	if err != nil {
		return err
	}
	fmt.Printf("Account number %v: successfully updated to %s\n", accountID, newStatus)
	return nil
}
