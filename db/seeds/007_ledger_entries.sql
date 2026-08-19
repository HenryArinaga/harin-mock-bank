INSERT INTO ledger_entries (account_id, transaction_id, direction, amount, currency)
SELECT
  to_account_id,
  id,
  'credit',
  amount,
  currency
FROM transactions
WHERE transaction_type = 'deposit'
  AND transaction_status = 'completed';

INSERT INTO ledger_entries (account_id, transaction_id, direction, amount, currency)
SELECT
  from_account_id,
  id,
  'debit',
  amount,
  currency
FROM transactions
WHERE transaction_type = 'withdrawal'
  AND transaction_status = 'completed';

INSERT INTO ledger_entries (account_id, transaction_id, direction, amount, currency)
SELECT
  from_account_id,
  id,
  'debit',
  amount,
  currency
FROM transactions
WHERE transaction_type = 'transfer'
  AND transaction_status = 'completed';

INSERT INTO ledger_entries (account_id, transaction_id, direction, amount, currency)
SELECT
  to_account_id,
  id,
  'credit',
  amount,
  currency
FROM transactions
WHERE transaction_type = 'transfer'
  AND transaction_status = 'completed';