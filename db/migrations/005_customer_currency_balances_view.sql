CREATE VIEW customer_currency_balances AS
SELECT
  c.id AS customer_id,
  c.first_name,
  c.last_name,
  u.email,
  ab.currency,
  COUNT(ab.account_id) AS account_count,
  COALESCE(SUM(ab.balance), 0) AS total_balance
FROM customers c
JOIN users u ON u.id = c.user_id
JOIN accounts a ON a.customer_id = c.id
JOIN account_balances ab ON ab.account_id = a.id
GROUP BY
  c.id,
  c.first_name,
  c.last_name,
  u.email,
  ab.currency;