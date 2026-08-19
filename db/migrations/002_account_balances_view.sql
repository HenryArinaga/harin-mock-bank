CREATE VIEW account_balances AS
SELECT
  a.id AS account_id,
  a.account_number,
  a.account_type,
  a.account_status,
  a.currency,
  COALESCE(
    SUM(
      CASE
        WHEN le.direction = 'credit' THEN le.amount
        WHEN le.direction = 'debit' THEN -le.amount
        ELSE 0
      END
    ),
    0
  ) AS balance
FROM accounts a
LEFT JOIN ledger_entries le ON le.account_id = a.id
GROUP BY
  a.id,
  a.account_number,
  a.account_type,
  a.account_status,
  a.currency;