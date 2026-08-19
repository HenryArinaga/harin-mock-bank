CREATE VIEW customer_summary AS
SELECT
  c.id AS customer_id,
  c.first_name,
  c.last_name,
  u.email,
  c.phone,
  c.date_of_birth,
  COUNT(a.id) AS account_count,
  COALESCE(SUM(ab.balance), 0) AS total_balance
FROM customers c
JOIN users u ON u.id = c.user_id
LEFT JOIN accounts a ON a.customer_id = c.id
LEFT JOIN account_balances ab ON ab.account_id = a.id
GROUP BY
  c.id,
  c.first_name,
  c.last_name,
  u.email,
  c.phone,
  c.date_of_birth;