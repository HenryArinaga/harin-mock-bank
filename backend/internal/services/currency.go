package services

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func EnsureCorrectCurrency(ctx context.Context, db DBRunner, accountID int64, currency string) error {
	currencyCheck := `
	SELECT currency = @currency
	FROM accounts
	WHERE id = @accountID`
	args := pgx.NamedArgs{
		"accountID": accountID,
		"currency":  currency,
	}
	var sameCurrency bool
	row := db.QueryRow(ctx, currencyCheck, args)
	err := row.Scan(
		&sameCurrency,
	)
	if err != nil {
		return err
	}
	if !sameCurrency {
		return fmt.Errorf("the currency types are not the same")
	}
	return nil
}
