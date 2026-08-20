CREATE VIEW account_balances AS
SELECT
  accounts.id AS account_id,
  accounts.account_number,
  accounts.account_type,
  accounts.account_status,
  accounts.currency,
  COALESCE(
    SUM(
      CASE
        WHEN ledger_entries.direction = 'credit' THEN ledger_entries.amount
        WHEN ledger_entries.direction = 'debit' THEN -ledger_entries.amount
        ELSE 0
      END
    ),
    0
  ) AS balance
FROM accounts
LEFT JOIN ledger_entries ON ledger_entries.account_id = accounts.id
GROUP BY
  accounts.id,
  accounts.account_number,
  accounts.account_type,
  accounts.account_status,
  accounts.currency;