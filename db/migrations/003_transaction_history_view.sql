CREATE VIEW transaction_history AS
SELECT
  t.id AS transaction_id,
  t.transaction_type,
  t.transaction_description,
  t.transaction_status,
  t.currency,
  t.amount,
  t.created_at,
  from_account.account_number AS from_account_number,
  to_account.account_number AS to_account_number
FROM transactions t
LEFT JOIN accounts from_account ON from_account.id = t.from_account_id
LEFT JOIN accounts to_account ON to_account.id = t.to_account_id;