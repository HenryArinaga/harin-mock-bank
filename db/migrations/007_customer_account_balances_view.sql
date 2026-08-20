CREATE VIEW customer_account_balances AS
SELECT 
accounts.id AS account_id,
accounts.customer_id,
accounts.account_number,
accounts.account_type,
accounts.account_status,
accounts.currency,
account_balances.balance
FROM accounts
LEFT JOIN account_balances ON account_balances.account_id = accounts.id

