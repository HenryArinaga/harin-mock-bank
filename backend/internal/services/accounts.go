package services

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CustomerAccount struct {
	AccountID      int64
	CustomerID     int64
	FirstName      string
	LastName       string
	AccountStatus  string
	AccountType    string
	AccountNumber  string
	Currency       string
	AccountBalance string
}

type TransactionsByAccount struct {
	TransactionID          int64
	ToAccountNumber        pgtype.Text
	FromAccountNumber      pgtype.Text
	TransactionType        string
	TransactionDescription pgtype.Text
	TransactionStatus      string
	Currency               string
	Amount                 string
	CreatedAt              time.Time
}

type accountStatusUpdate struct {
	AccountID     int64
	CurrentStatus string
	NewStatus     string
	UserRole      string
}

var allowedAccountStatusTransitions = map[string][]string{
	"pending":   {"active", "frozen", "suspended", "closed"},
	"active":    {"frozen", "suspended", "closed"},
	"frozen":    {"active", "suspended", "closed"},
	"suspended": {"active", "frozen", "closed"},
	"closed":    {},
}

var allowedUserRole = map[string][]string{
	"admin":    {"active", "frozen", "suspended", "closed"},
	"support":  {"active", "frozen", "suspended"},
	"customer": {},
}

// helper functions for other functions in this file
func isAccountStatusTransitionAllowed(currentStatus string, newStatus string) bool {
	allowedNewStatus, ok := allowedAccountStatusTransitions[currentStatus]
	if !ok {
		return false
	}
	for _, allowedStatus := range allowedNewStatus {
		if allowedStatus == newStatus {
			return true
		}
	}
	return false
}

func isAllowedToChangeStatusOfAccount(userRole string, newStatus string) bool {
	userRoleAllowed, ok := allowedUserRole[userRole]
	if !ok {
		return false
	}
	for _, allowedStatus := range userRoleAllowed {
		if allowedStatus == newStatus {
			return true
		}
	}
	return false
}

func GenerateAccountNumber() (string, error) {
	limit := big.NewInt(9000000000)
	n, err := rand.Int(rand.Reader, limit)
	if err != nil {
		return "", err
	}

	n.Add(n, big.NewInt(1000000000))

	return n.String(), nil
}

func ListAccountsByCustomer(ctx context.Context, pool *pgxpool.Pool, customerID int64) ([]CustomerAccount, error) {
	rows, err := pool.Query(ctx,
		`SELECT account_id, customer_id, first_name,
		last_name, account_status, account_type,
		account_number, currency,balance 
		FROM customer_account_balances 
		WHERE customer_id = $1`, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accounts := make([]CustomerAccount, 0)
	for rows.Next() {
		var account CustomerAccount
		err := rows.Scan(
			&account.AccountID,
			&account.CustomerID,
			&account.FirstName,
			&account.LastName,
			&account.AccountStatus,
			&account.AccountType,
			&account.AccountNumber,
			&account.Currency,
			&account.AccountBalance,
		)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return accounts, nil
}

func ListTransactionsByAccount(ctx context.Context, pool *pgxpool.Pool, accountNumber string) ([]TransactionsByAccount, error) {
	rows, err := pool.Query(ctx,
		`SELECT transaction_id, to_account_number, from_account_number, 
		transaction_type, transaction_description, transaction_status, 
		currency, amount, created_at 
		FROM transaction_history 
		WHERE from_account_number = $1 OR to_account_number = $1`, accountNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := make([]TransactionsByAccount, 0)
	for rows.Next() {
		var transaction TransactionsByAccount
		err := rows.Scan(
			&transaction.TransactionID,
			&transaction.ToAccountNumber,
			&transaction.FromAccountNumber,
			&transaction.TransactionType,
			&transaction.TransactionDescription,
			&transaction.TransactionStatus,
			&transaction.Currency,
			&transaction.Amount,
			&transaction.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return transactions, nil
}

func EnsureAccountExists(ctx context.Context, db DBRunner, accountID int64) error {
	accountCheck := `
	SELECT EXISTS (
		SELECT 1 
		FROM accounts 
		WHERE id = @accountID
		)`
	args := pgx.NamedArgs{
		"accountID": accountID,
	}
	var accountExists bool
	row := db.QueryRow(ctx, accountCheck, args)
	err := row.Scan(
		&accountExists,
	)
	if err != nil {
		return err
	}
	if !accountExists {
		return fmt.Errorf("account does not exist")
	}
	return nil
}

func EnsureAccountActive(ctx context.Context, db DBRunner, accountID int64) error {
	accountActiveCheck := `
	SELECT account_status = 'active'
	FROM accounts
	WHERE id = @accountID`
	args := pgx.NamedArgs{
		"accountID": accountID,
	}
	var accountActive bool
	row := db.QueryRow(ctx, accountActiveCheck, args)
	err := row.Scan(
		&accountActive,
	)
	if err != nil {
		return err
	}
	if !accountActive {
		return fmt.Errorf("account is not active")
	}
	return nil
}

func EnsureAccountOwnerValid(ctx context.Context, db DBRunner, fromAccountID int64, customerID int64) error {
	accountOwner := `
	SELECT EXISTS (
	SELECT 1 
	FROM accounts
	WHERE id = @fromAccountId
	AND customer_id = @customerID
	)`
	args := pgx.NamedArgs{
		"fromAccountId": fromAccountID,
		"customerID":    customerID,
	}
	var accountOwnerCheck bool
	row := db.QueryRow(ctx, accountOwner, args)
	err := row.Scan(
		&accountOwnerCheck,
	)
	if err != nil {
		return err
	}
	if !accountOwnerCheck {
		return fmt.Errorf("customer is not account owner")
	}
	return nil
}

func ListMyAccounts(ctx context.Context, pool *pgxpool.Pool, tokenString string) ([]CustomerAccount, error) {

	validateAccount, err := GetCustomerProfileByJWT(ctx, pool, tokenString)
	if err != nil {
		return nil, err
	}

	accounts, err := ListAccountsByCustomer(ctx, pool, validateAccount.ID)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func CreateAccount(ctx context.Context, db DBRunner, input CustomerAccount) (int64, error) {
	var accountID int64
	var pgErr *pgconn.PgError
	for attempts := 0; attempts < 5; attempts++ {
		accountNumber, err := GenerateAccountNumber()
		if err != nil {
			return 0, err
		}

		insertAccountInformation := `
	INSERT into accounts (
	customer_id,
	currency,
	account_type,
	account_number
	)
	VALUES (
	@customerID,
	@currency,
	@accountType,
	@accountNumber
	)
	RETURNING id`

		args := pgx.NamedArgs{
			"customerID":    input.CustomerID,
			"currency":      input.Currency,
			"accountType":   input.AccountType,
			"accountNumber": accountNumber,
		}

		row := db.QueryRow(ctx, insertAccountInformation, args)
		err = row.Scan(&accountID)
		if err != nil {
			if errors.As(err, &pgErr) && pgErr.Code == "23505" {
				continue
			}
			return 0, fmt.Errorf("unable to generate unique account number")
		}
		return accountID, nil
	}
	return 0, fmt.Errorf("unable to generate unique account number")
}

func GetAccountStatusByID(ctx context.Context, db DBRunner, accountID int64) (string, error) {
	getAccountStatusQuery := `
	SELECT account_status 
	FROM accounts
	WHERE id = @accountID`

	args := pgx.NamedArgs{
		"accountID": accountID,
	}
	var status string
	row := db.QueryRow(ctx, getAccountStatusQuery, args)
	err := row.Scan(&status)
	if err != nil {
		return "", err
	}

	return status, nil
}

// lower case function to make it harder to misuse
func updateAccountStatus(ctx context.Context, db DBRunner, input accountStatusUpdate) error {

	allowed := isAccountStatusTransitionAllowed(input.CurrentStatus, input.NewStatus)

	if !allowed {
		return fmt.Errorf("invalid account status transition from %s to %s", input.CurrentStatus, input.NewStatus)
	}

	updateAccountbyID := `
	UPDATE accounts
	SET account_status = @newStatus
	WHERE id = @id
	AND account_status = @currentStatus`

	args := pgx.NamedArgs{
		"newStatus":     input.NewStatus,
		"id":            input.AccountID,
		"currentStatus": input.CurrentStatus,
	}
	commandTag, err := db.Exec(ctx, updateAccountbyID, args)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() < 1 {
		return errors.New("no accounts found to update status")
	}
	return nil
}

func ChangeAccountStatus(ctx context.Context, db DBRunner, accountID int64, newStatus string, userRole string) error {
	rolePermissionCheck := isAllowedToChangeStatusOfAccount(userRole, newStatus)
	if !rolePermissionCheck {
		return errors.New("user permissions not allowed to change account status")
	}

	accountStatus, err := GetAccountStatusByID(ctx, db, accountID)
	if err != nil {
		return err
	}

	accountUpdate := accountStatusUpdate{
		AccountID:     accountID,
		CurrentStatus: accountStatus,
		NewStatus:     newStatus,
	}

	err = updateAccountStatus(ctx, db, accountUpdate)
	if err != nil {
		return err
	}
	return nil

}
