DROP VIEW IF EXISTS customer_account_balances;
CREATE VIEW customer_account_balances AS
SELECT 
accounts.id AS account_id,
accounts.customer_id,
customers.first_name,
customers.last_name,
accounts.account_number,
accounts.account_type,
accounts.account_status,
accounts.currency,
account_balances.balance
FROM accounts
LEFT JOIN account_balances ON account_balances.account_id = accounts.id
LEFT JOIN customers ON customers.id = accounts.customer_id;

